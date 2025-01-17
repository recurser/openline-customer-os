package graph

import (
	"context"
	"github.com/google/uuid"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/command_handler"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/models"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/graph_db"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/graph_db/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/test"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/test/eventstore"
	neo4jt "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/test/neo4j"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
	"time"
)

const customerOsIdPattern = `^C-[A-HJ-NP-Z2-9]{3}-[A-HJ-NP-Z2-9]{3}$`

func TestGraphOrganizationEventHandler_OnOrganizationCreate(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()

	// prepare neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	userId := neo4jt.CreateUser(ctx, testDatabase.Driver, tenantName, entity.UserEntity{
		FirstName: "logged-in",
		LastName:  "user",
	})
	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		"Organization": 0,
		"User":         1, "User_" + tenantName: 1,
		"Action": 0, "TimelineEvent": 0})

	orgId := uuid.New().String()

	// prepare event handler
	orgEventHandler := &OrganizationEventHandler{
		repositories:         testDatabase.Repositories,
		organizationCommands: command_handler.NewOrganizationCommands(testLogger, &config.Config{}, aggregateStore, testDatabase.Repositories),
	}
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	now := utils.Now()
	event, err := events.NewOrganizationCreateEvent(orgAggregate, &models.OrganizationFields{
		ID: orgId,
		OrganizationDataFields: models.OrganizationDataFields{
			Name: "test org",
		},
	}, now, now)
	require.Nil(t, err)
	metadata := make(map[string]string)
	metadata["user-id"] = userId
	err = event.SetMetadata(metadata)
	require.Nil(t, err)

	// EXECUTE
	err = orgEventHandler.OnOrganizationCreate(context.Background(), event)
	require.Nil(t, err)

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{
		"User": 1, "User_" + tenantName: 1,
		"Organization": 1, "Organization_" + tenantName: 1,
		"Action": 1, "Action_" + tenantName: 1,
		"TimelineEvent": 1, "TimelineEvent_" + tenantName: 1})
	neo4jt.AssertNeo4jRelationCount(ctx, t, testDatabase.Driver, map[string]int{
		"ACTION_ON": 1,
		"OWNS":      1,
	})

	orgDbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Organization_"+tenantName, orgId)
	require.Nil(t, err)
	require.NotNil(t, orgDbNode)

	// verify organization
	organization := graph_db.MapDbNodeToOrganizationEntity(*orgDbNode)
	require.Equal(t, orgId, organization.ID)
	require.Equal(t, "test org", organization.Name)
	require.Equal(t, now, organization.CreatedAt)
	require.NotNil(t, organization.UpdatedAt)

	// verify action
	actionDbNode, err := neo4jt.GetFirstNodeByLabel(ctx, testDatabase.Driver, "Action_"+tenantName)
	require.Nil(t, err)
	require.NotNil(t, actionDbNode)
	action := graph_db.MapDbNodeToActionEntity(*actionDbNode)
	require.NotNil(t, action.Id)
	require.Equal(t, entity.DataSource(constants.SourceOpenline), action.Source)
	require.Equal(t, entity.DataSource(constants.SourceOpenline), action.SourceOfTruth)
	require.Equal(t, constants.AppSourceEventProcessingPlatform, action.AppSource)
	require.Equal(t, now, action.CreatedAt)
	require.Equal(t, entity.ActionCreated, action.Type)
	require.Equal(t, "", action.Content)
	require.Equal(t, "", action.Metadata)

	// Check last touch point request was generated
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))
	eventList := eventsMap[orgAggregate.ID]
	require.Equal(t, 1, len(eventList))
	generatedEvent := eventList[0]
	require.Equal(t, events.OrganizationRefreshLastTouchpointV1, generatedEvent.EventType)
	var eventData events.OrganizationRefreshLastTouchpointEvent
	err = generatedEvent.GetJsonData(&eventData)
	require.Nil(t, err)
	require.Equal(t, tenantName, eventData.Tenant)
}

func TestGraphOrganizationEventHandler_OnRenewalLikelihoodUpdate(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()

	// prepare neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	userId := neo4jt.CreateUser(ctx, testDatabase.Driver, tenantName, entity.UserEntity{
		FirstName: "new",
		LastName:  "user",
	})
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{
		Name: "test org",
		RenewalLikelihood: entity.RenewalLikelihood{
			RenewalLikelihood:         string(entity.RenewalLikelihoodZero),
			PreviousRenewalLikelihood: string(entity.RenewalLikelihoodHigh),
			Comment:                   utils.StringPtr("old comment"),
			UpdatedBy:                 "old user",
		},
	})
	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1,
		"User": 1, "User_" + tenantName: 1, "Action": 0, "TimelineEvent": 0})

	// prepare event handler
	orgEventHandler := &OrganizationEventHandler{
		repositories:         testDatabase.Repositories,
		organizationCommands: command_handler.NewOrganizationCommands(testLogger, &config.Config{}, aggregateStore, testDatabase.Repositories),
	}
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	now := utils.Now()
	event, err := events.NewOrganizationUpdateRenewalLikelihoodEvent(orgAggregate, models.RenewalLikelihoodLOW, models.RenewalLikelihoodHIGH, userId, nil, now)
	require.Nil(t, err)

	// EXECUTE
	err = orgEventHandler.OnRenewalLikelihoodUpdate(context.Background(), event)
	require.Nil(t, err)

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1,
		"User": 1, "User_" + tenantName: 1,
		"Action": 1, "Action_" + tenantName: 1,
		"TimelineEvent": 1, "TimelineEvent_" + tenantName: 1})

	orgDbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Organization_"+tenantName, orgId)
	require.Nil(t, err)
	require.NotNil(t, orgDbNode)

	// verify organization
	organization := graph_db.MapDbNodeToOrganizationEntity(*orgDbNode)
	require.Equal(t, orgId, organization.ID)
	require.Equal(t, string(entity.RenewalLikelihoodZero), organization.RenewalLikelihood.PreviousRenewalLikelihood)
	require.Equal(t, string(entity.RenewalLikelihoodLow), organization.RenewalLikelihood.RenewalLikelihood)
	require.Equal(t, now, *organization.RenewalLikelihood.UpdatedAt)
	require.Nil(t, organization.RenewalLikelihood.Comment)
	require.Equal(t, userId, organization.RenewalLikelihood.UpdatedBy)
	require.Equal(t, entity.DataSourceOpenline, organization.SourceOfTruth)
	require.NotNil(t, organization.UpdatedAt)

	// verify action
	actionDbNode, err := neo4jt.GetFirstNodeByLabel(ctx, testDatabase.Driver, "Action_"+tenantName)
	require.Nil(t, err)
	require.NotNil(t, actionDbNode)
	action := graph_db.MapDbNodeToActionEntity(*actionDbNode)
	require.NotNil(t, action.Id)
	require.Equal(t, entity.DataSource(constants.SourceOpenline), action.Source)
	require.Equal(t, constants.AppSourceEventProcessingPlatform, action.AppSource)
	require.Equal(t, now, action.CreatedAt)
	require.Equal(t, entity.ActionRenewalLikelihoodUpdated, action.Type)
	require.Equal(t, "Renewal likelihood set to Low by new user", action.Content)
	require.Equal(t, `{"likelihood":"LOW","reason":""}`, action.Metadata)

	// Check request was generated
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))
	eventList := eventsMap[orgAggregate.ID]
	require.Equal(t, 1, len(eventList))
	generatedEvent := eventList[0]
	require.Equal(t, events.OrganizationRequestRenewalForecastV1, generatedEvent.EventType)
	var eventData events.OrganizationRequestRenewalForecastEvent
	err = generatedEvent.GetJsonData(&eventData)
	require.Nil(t, err)
	test.AssertRecentTime(t, eventData.RequestedAt)
	require.Equal(t, tenantName, eventData.Tenant)
}

func TestGraphOrganizationEventHandler_OnRenewalForecastUpdate_ByUser(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()

	// create neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	userId := neo4jt.CreateUser(ctx, testDatabase.Driver, tenantName, entity.UserEntity{
		FirstName: "new",
		LastName:  "user",
	})
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{
		Name: "test org",
		RenewalForecast: entity.RenewalForecast{
			Amount:          utils.Float64Ptr(100),
			PotentialAmount: utils.Float64Ptr(200),
			Comment:         utils.StringPtr("old comment"),
		},
	})

	// prepare event handler
	orgEventHandler := &OrganizationEventHandler{
		repositories:         testDatabase.Repositories,
		organizationCommands: command_handler.NewOrganizationCommands(testLogger, &config.Config{}, aggregateStore, testDatabase.Repositories),
	}
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	now := utils.Now()
	event, err := events.NewOrganizationUpdateRenewalForecastEvent(orgAggregate, utils.Float64Ptr(50), utils.Float64Ptr(60), utils.Float64Ptr(10), userId, utils.StringPtr("new comment"), now, "")
	require.Nil(t, err)

	// EXECUTE
	err = orgEventHandler.OnRenewalForecastUpdate(context.Background(), event)
	require.Nil(t, err)

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1,
		"User": 1, "User_" + tenantName: 1,
		"Action": 1, "Action_" + tenantName: 1,
		"TimelineEvent": 1, "TimelineEvent_" + tenantName: 1})

	dbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Organization_"+tenantName, orgId)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	// verify organization
	organization := graph_db.MapDbNodeToOrganizationEntity(*dbNode)
	require.Equal(t, orgId, organization.ID)
	require.Equal(t, float64(50), *organization.RenewalForecast.Amount)
	// potential should not be updated
	require.Equal(t, float64(200), *organization.RenewalForecast.PotentialAmount)
	require.Equal(t, now, *organization.RenewalForecast.UpdatedAt)
	require.Equal(t, "new comment", *organization.RenewalForecast.Comment)
	require.Equal(t, userId, organization.RenewalForecast.UpdatedBy)
	require.Equal(t, entity.DataSourceOpenline, organization.SourceOfTruth)
	require.NotNil(t, organization.UpdatedAt)

	// verify action
	actionDbNode, err := neo4jt.GetFirstNodeByLabel(ctx, testDatabase.Driver, "Action_"+tenantName)
	require.Nil(t, err)
	require.NotNil(t, actionDbNode)
	action := graph_db.MapDbNodeToActionEntity(*actionDbNode)
	require.NotNil(t, action.Id)
	require.Equal(t, entity.DataSource(constants.SourceOpenline), action.Source)
	require.Equal(t, constants.AppSourceEventProcessingPlatform, action.AppSource)
	require.Equal(t, now, action.CreatedAt)
	require.Equal(t, entity.ActionRenewalForecastUpdated, action.Type)
	require.Equal(t, "Renewal forecast set to $50 by new user", action.Content)
	require.Equal(t, `{"likelihood":"","reason":"new comment"}`, action.Metadata)

	// Check request was not generated
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 0, len(eventsMap))
}

func TestGraphOrganizationEventHandler_OnRenewalForecastUpdate_ByInternalProcess(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()

	// create neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{
		Name: "test org",
		RenewalForecast: entity.RenewalForecast{
			Amount:          utils.Float64Ptr(100),
			PotentialAmount: utils.Float64Ptr(200),
			Comment:         utils.StringPtr("old comment"),
			UpdatedBy:       "old-user",
		},
	})

	// prepare event handler
	orgEventHandler := &OrganizationEventHandler{
		repositories:         testDatabase.Repositories,
		organizationCommands: command_handler.NewOrganizationCommands(testLogger, &config.Config{}, aggregateStore, testDatabase.Repositories),
	}
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	now := utils.Now()
	event, err := events.NewOrganizationUpdateRenewalForecastEvent(orgAggregate, utils.Float64Ptr(5000), utils.Float64Ptr(10000), nil, "", utils.StringPtr("new comment"), now, models.RenewalLikelihoodMEDIUM)
	require.Nil(t, err)

	// EXECUTE
	err = orgEventHandler.OnRenewalForecastUpdate(context.Background(), event)
	require.Nil(t, err)

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1,
		"Action": 1, "Action_" + tenantName: 1,
		"TimelineEvent": 1, "TimelineEvent_" + tenantName: 1})

	dbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Organization_"+tenantName, orgId)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	// verify organization
	organization := graph_db.MapDbNodeToOrganizationEntity(*dbNode)
	require.Equal(t, orgId, organization.ID)
	require.Equal(t, float64(5000), *organization.RenewalForecast.Amount)
	require.Equal(t, float64(10000), *organization.RenewalForecast.PotentialAmount)
	require.Equal(t, now, *organization.RenewalForecast.UpdatedAt)
	require.Equal(t, "new comment", *organization.RenewalForecast.Comment)
	require.Equal(t, "", organization.RenewalForecast.UpdatedBy)
	require.Equal(t, entity.DataSourceOpenline, organization.SourceOfTruth)
	require.NotNil(t, organization.UpdatedAt)

	// verify action
	actionDbNode, err := neo4jt.GetFirstNodeByLabel(ctx, testDatabase.Driver, "Action_"+tenantName)
	require.Nil(t, err)
	require.NotNil(t, actionDbNode)
	action := graph_db.MapDbNodeToActionEntity(*actionDbNode)
	require.NotNil(t, action.Id)
	require.Equal(t, entity.DataSource(constants.SourceOpenline), action.Source)
	require.Equal(t, constants.AppSourceEventProcessingPlatform, action.AppSource)
	require.Equal(t, now, action.CreatedAt)
	require.Equal(t, entity.ActionRenewalForecastUpdated, action.Type)
	require.Equal(t, "Renewal forecast set by default to $5,000, by discounting the billing amount using the renewal likelihood", action.Content)
	require.Equal(t, `{"likelihood":"MEDIUM","reason":"new comment"}`, action.Metadata)

	// Check request was not generated
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 0, len(eventsMap))
}

func TestGraphOrganizationEventHandler_OnRenewalForecastUpdate_ResetAmount(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()

	// create neo4j data
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	userId := neo4jt.CreateUser(ctx, testDatabase.Driver, tenantName, entity.UserEntity{
		FirstName: "new",
		LastName:  "user",
	})
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{
		Name: "test org",
		RenewalForecast: entity.RenewalForecast{
			Amount:          utils.Float64Ptr(100),
			PotentialAmount: utils.Float64Ptr(200),
			Comment:         utils.StringPtr("old comment"),
			UpdatedBy:       "old-user",
		},
	})

	// prepare event handler
	orgEventHandler := &OrganizationEventHandler{
		repositories:         testDatabase.Repositories,
		organizationCommands: command_handler.NewOrganizationCommands(testLogger, &config.Config{}, aggregateStore, testDatabase.Repositories),
	}
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	now := utils.Now()
	event, err := events.NewOrganizationUpdateRenewalForecastEvent(orgAggregate, nil, nil, nil, userId, utils.StringPtr("new comment"), now, "")
	require.Nil(t, err)

	// EXECUTE
	err = orgEventHandler.OnRenewalForecastUpdate(context.Background(), event)
	require.Nil(t, err)

	// no actions created
	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1,
		"User": 1, "User_" + tenantName: 1,
		"Action": 0, "Action_" + tenantName: 0,
		"TimelineEvent": 0, "TimelineEvent_" + tenantName: 0})

	dbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Organization_"+tenantName, orgId)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	organization := graph_db.MapDbNodeToOrganizationEntity(*dbNode)
	require.Equal(t, orgId, organization.ID)
	require.Nil(t, organization.RenewalForecast.Amount)
	require.Equal(t, float64(200), *organization.RenewalForecast.PotentialAmount)
	require.Equal(t, now, *organization.RenewalForecast.UpdatedAt)
	require.Equal(t, "new comment", *organization.RenewalForecast.Comment)
	require.Equal(t, userId, organization.RenewalForecast.UpdatedBy)
	require.Equal(t, entity.DataSourceOpenline, organization.SourceOfTruth)
	require.NotNil(t, organization.UpdatedAt)

	// Check request was generated
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))
	eventList := eventsMap[orgAggregate.ID]
	require.Equal(t, 1, len(eventList))
	generatedEvent := eventList[0]
	require.Equal(t, events.OrganizationRequestRenewalForecastV1, generatedEvent.EventType)
	var eventData events.OrganizationRequestRenewalForecastEvent
	err = generatedEvent.GetJsonData(&eventData)
	require.Nil(t, err)
	test.AssertRecentTime(t, eventData.RequestedAt)
	require.Equal(t, tenantName, eventData.Tenant)
}

func TestGraphOrganizationEventHandler_OnBillingDetailsUpdate(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()

	hourAgo := utils.Now().Add(time.Duration(-1) * time.Hour)
	minAgo := utils.Now().Add(time.Duration(-1) * time.Minute)

	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{
		Name: "test org",
		BillingDetails: entity.BillingDetails{
			Amount:            utils.Float64Ptr(100),
			Frequency:         "WEEKLY",
			RenewalCycle:      "MONTHLY",
			RenewalCycleStart: utils.TimePtr(hourAgo),
			RenewalCycleNext:  utils.TimePtr(minAgo),
		},
	})
	orgEventHandler := &OrganizationEventHandler{
		repositories:         testDatabase.Repositories,
		organizationCommands: command_handler.NewOrganizationCommands(testLogger, &config.Config{}, aggregateStore, testDatabase.Repositories),
	}
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	now := utils.Now()

	event, err := events.NewOrganizationUpdateBillingDetailsEvent(orgAggregate, utils.Float64Ptr(50), "MONTHLY", "ANNUALLY", "new user", utils.TimePtr(now), utils.TimePtr(now))
	require.Nil(t, err)
	err = orgEventHandler.OnBillingDetailsUpdate(context.Background(), event)
	require.Nil(t, err)

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1})

	dbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Organization_"+tenantName, orgId)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	organization := graph_db.MapDbNodeToOrganizationEntity(*dbNode)
	require.Equal(t, orgId, organization.ID)
	require.Equal(t, float64(50), *organization.BillingDetails.Amount)
	require.Equal(t, "MONTHLY", organization.BillingDetails.Frequency)
	require.Equal(t, "ANNUALLY", organization.BillingDetails.RenewalCycle)
	require.Equal(t, now, *organization.BillingDetails.RenewalCycleStart)
	require.Equal(t, minAgo, *organization.BillingDetails.RenewalCycleNext)
	require.Equal(t, entity.DataSourceOpenline, organization.SourceOfTruth)
	require.NotNil(t, organization.UpdatedAt)

	// Check request was generated
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))
	eventList := eventsMap[orgAggregate.ID]
	require.Equal(t, 2, len(eventList))

	generatedEvent1 := eventList[0]
	require.Equal(t, events.OrganizationRequestRenewalForecastV1, generatedEvent1.EventType)
	var eventData1 events.OrganizationRequestRenewalForecastEvent
	err = generatedEvent1.GetJsonData(&eventData1)
	require.Nil(t, err)
	test.AssertRecentTime(t, eventData1.RequestedAt)
	require.Equal(t, tenantName, eventData1.Tenant)

	generatedEvent2 := eventList[1]
	require.Equal(t, events.OrganizationRequestNextCycleDateV1, generatedEvent2.EventType)
	var eventData2 events.OrganizationRequestNextCycleDateEvent
	err = generatedEvent2.GetJsonData(&eventData2)
	require.Nil(t, err)
	test.AssertRecentTime(t, eventData2.RequestedAt)
	require.Equal(t, tenantName, eventData2.Tenant)
}

func TestGraphOrganizationEventHandler_OnBillingDetailsUpdate_SetNotByUser(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()

	hourAgo := utils.Now().Add(time.Duration(-1) * time.Hour)
	minAgo := utils.Now().Add(time.Duration(-1) * time.Minute)

	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{
		Name: "test org",
		BillingDetails: entity.BillingDetails{
			Amount:            utils.Float64Ptr(100),
			Frequency:         "WEEKLY",
			RenewalCycle:      "MONTHLY",
			RenewalCycleStart: utils.TimePtr(hourAgo),
			RenewalCycleNext:  utils.TimePtr(minAgo),
		},
	})
	orgEventHandler := &OrganizationEventHandler{
		repositories:         testDatabase.Repositories,
		organizationCommands: command_handler.NewOrganizationCommands(testLogger, &config.Config{}, aggregateStore, testDatabase.Repositories),
	}
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	now := utils.Now()
	tomorrow := now.Add(time.Duration(24) * time.Hour)

	event, err := events.NewOrganizationUpdateBillingDetailsEvent(orgAggregate, utils.Float64Ptr(50), "MONTHLY", "ANNUALLY", "", utils.TimePtr(now), utils.TimePtr(tomorrow))
	require.Nil(t, err)
	err = orgEventHandler.OnBillingDetailsUpdate(context.Background(), event)
	require.Nil(t, err)

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1})

	dbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Organization_"+tenantName, orgId)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	organization := graph_db.MapDbNodeToOrganizationEntity(*dbNode)
	require.Equal(t, orgId, organization.ID)
	require.Equal(t, float64(50), *organization.BillingDetails.Amount)
	require.Equal(t, "MONTHLY", organization.BillingDetails.Frequency)
	require.Equal(t, "ANNUALLY", organization.BillingDetails.RenewalCycle)
	require.Equal(t, now, *organization.BillingDetails.RenewalCycleStart)
	require.Equal(t, tomorrow, *organization.BillingDetails.RenewalCycleNext)
	require.Equal(t, entity.DataSourceOpenline, organization.SourceOfTruth)
	require.NotNil(t, organization.UpdatedAt)

	// Check request was not generated
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 0, len(eventsMap))
}

func TestGraphOrganizationEventHandler_OnOrganizationHide(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()

	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{
		Name: "test org",
		Hide: false,
	})
	orgEventHandler := &OrganizationEventHandler{
		repositories:         testDatabase.Repositories,
		organizationCommands: command_handler.NewOrganizationCommands(testLogger, &config.Config{}, aggregateStore, testDatabase.Repositories),
	}
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)

	event, err := events.NewHideOrganizationEventEvent(orgAggregate)
	require.Nil(t, err)
	err = orgEventHandler.OnOrganizationHide(context.Background(), event)
	require.Nil(t, err)

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1})
	neo4jt.AssertNeo4jLabels(ctx, t, testDatabase.Driver, []string{"Organization", "Organization_" + tenantName, "Tenant"})

	dbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Organization_"+tenantName, orgId)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	organization := graph_db.MapDbNodeToOrganizationEntity(*dbNode)
	require.Equal(t, orgId, organization.ID)
	require.Equal(t, true, organization.Hide)
}

func TestGraphOrganizationEventHandler_OnOrganizationShow(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()

	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{
		Name: "test org",
		Hide: true,
	})
	orgEventHandler := &OrganizationEventHandler{
		repositories:         testDatabase.Repositories,
		organizationCommands: command_handler.NewOrganizationCommands(testLogger, &config.Config{}, aggregateStore, testDatabase.Repositories),
	}
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)

	event, err := events.NewShowOrganizationEventEvent(orgAggregate)
	require.Nil(t, err)
	err = orgEventHandler.OnOrganizationShow(context.Background(), event)
	require.Nil(t, err)

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1})
	neo4jt.AssertNeo4jLabels(ctx, t, testDatabase.Driver, []string{"Organization", "Organization_" + tenantName, "Tenant"})

	dbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Organization_"+tenantName, orgId)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	organization := graph_db.MapDbNodeToOrganizationEntity(*dbNode)
	require.Equal(t, orgId, organization.ID)
	require.Equal(t, false, organization.Hide)
	require.NotEqual(t, "", organization.CustomerOsId)
	require.True(t, regexp.MustCompile(customerOsIdPattern).MatchString(organization.CustomerOsId), "Valid CustomerOsId should match the format")
}

func TestGraphOrganizationEventHandler_OnSocialAddedToOrganization_New(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()

	socialId := uuid.New().String()
	socialUrl := "https://www.facebook.com/organization"
	platformName := "facebook"
	now := utils.Now()
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{
		Name: "test org",
	})
	neo4jt.CreateSocial(ctx, testDatabase.Driver, tenantName, entity.SocialEntity{
		Url: socialUrl,
	})
	orgEventHandler := &OrganizationEventHandler{
		repositories:         testDatabase.Repositories,
		organizationCommands: command_handler.NewOrganizationCommands(testLogger, &config.Config{}, aggregateStore, testDatabase.Repositories),
	}
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)

	event, err := events.NewOrganizationAddSocialEvent(orgAggregate, socialId, platformName, socialUrl, constants.SourceOpenline, constants.SourceOpenline, constants.AppSourceEventProcessingPlatform, now, now)
	require.Nil(t, err)
	err = orgEventHandler.OnSocialAddedToOrganization(context.Background(), event)
	require.Nil(t, err)

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1, "Social": 2, "Social_" + tenantName: 2})
	neo4jt.AssertNeo4jLabels(ctx, t, testDatabase.Driver, []string{"Organization", "Organization_" + tenantName, "Tenant", "Social", "Social_" + tenantName})

	dbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Social_"+tenantName, socialId)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	social := graph_db.MapDbNodeToSocialEntity(*dbNode)
	require.Equal(t, socialId, social.Id)
	require.Equal(t, socialUrl, social.Url)
	require.Equal(t, platformName, social.PlatformName)
	require.Equal(t, entity.DataSource(constants.SourceOpenline), social.SourceFields.Source)
	require.Equal(t, entity.DataSource(constants.SourceOpenline), social.SourceFields.SourceOfTruth)
	require.Equal(t, constants.AppSourceEventProcessingPlatform, social.SourceFields.AppSource)
	require.Equal(t, now, social.CreatedAt)
	require.Equal(t, now, social.UpdatedAt)
}

func TestGraphOrganizationEventHandler_OnSocialAddedToOrganization_SocialUrlAlreadyExistsForOrg_NoChanges(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()

	socialId := uuid.New().String()
	socialUrl := "https://www.facebook.com/organization"
	platformName := "facebook"
	now := utils.Now()
	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{
		Name: "test org",
	})
	existingSocialId := neo4jt.CreateSocial(ctx, testDatabase.Driver, tenantName, entity.SocialEntity{
		Url:          socialUrl,
		PlatformName: platformName,
	})
	neo4jt.LinkSocial(ctx, testDatabase.Driver, existingSocialId, orgId)

	orgEventHandler := &OrganizationEventHandler{
		repositories:         testDatabase.Repositories,
		organizationCommands: command_handler.NewOrganizationCommands(testLogger, &config.Config{}, aggregateStore, testDatabase.Repositories),
	}
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)

	event, err := events.NewOrganizationAddSocialEvent(orgAggregate, socialId, "other platform name", socialUrl, constants.SourceOpenline, constants.SourceOpenline, constants.AppSourceEventProcessingPlatform, now, now)
	require.Nil(t, err)
	err = orgEventHandler.OnSocialAddedToOrganization(context.Background(), event)
	require.Nil(t, err)

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1, "Social": 1, "Social_" + tenantName: 1})
	neo4jt.AssertNeo4jLabels(ctx, t, testDatabase.Driver, []string{"Organization", "Organization_" + tenantName, "Tenant", "Social", "Social_" + tenantName})

	dbNode, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Social_"+tenantName, existingSocialId)
	require.Nil(t, err)
	require.NotNil(t, dbNode)

	social := graph_db.MapDbNodeToSocialEntity(*dbNode)
	require.Equal(t, existingSocialId, social.Id)
	require.Equal(t, socialUrl, social.Url)
	require.Equal(t, platformName, social.PlatformName)
}

func TestGraphOrganizationEventHandler_OnLocationLinkedToOrganization(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	neo4jt.CreateTenant(ctx, testDatabase.Driver, tenantName)

	organizationName := "test_org_name"
	orgId := neo4jt.CreateOrganization(ctx, testDatabase.Driver, tenantName, entity.OrganizationEntity{
		Name: organizationName,
	})

	neo4jt.AssertNeo4jNodeCount(ctx, t, testDatabase.Driver, map[string]int{"Organization": 1, "Organization_" + tenantName: 1})
	dbNodeAfterOrganizationCreate, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Organization_"+tenantName, orgId)
	require.Nil(t, err)
	require.NotNil(t, dbNodeAfterOrganizationCreate)
	propsAfterOrganizationCreate := utils.GetPropsFromNode(*dbNodeAfterOrganizationCreate)
	require.Equal(t, organizationName, utils.GetStringPropOrEmpty(propsAfterOrganizationCreate, "name"))

	locationName := "test_location_name"
	locationId := neo4jt.CreateLocation(ctx, testDatabase.Driver, tenantName, entity.LocationEntity{
		Name: locationName,
	})

	dbNodeAfterLocationCreate, err := neo4jt.GetNodeById(ctx, testDatabase.Driver, "Location_"+tenantName, locationId)
	require.Nil(t, err)
	require.NotNil(t, dbNodeAfterLocationCreate)
	propsAfterLocationCreate := utils.GetPropsFromNode(*dbNodeAfterLocationCreate)
	require.Equal(t, locationName, utils.GetStringPropOrEmpty(propsAfterLocationCreate, "name"))

	orgEventHandler := &OrganizationEventHandler{
		repositories: testDatabase.Repositories,
	}
	orgAggregate := aggregate.NewOrganizationAggregateWithTenantAndID(tenantName, orgId)
	now := utils.Now()
	event, err := events.NewOrganizationLinkLocationEvent(orgAggregate, locationId, now)
	require.Nil(t, err)
	err = orgEventHandler.OnLocationLinkedToOrganization(context.Background(), event)
	require.Nil(t, err)

	require.Equal(t, 1, neo4jt.GetCountOfRelationships(ctx, testDatabase.Driver, "ASSOCIATED_WITH"), "Incorrect number of ASSOCIATED_WITH relationships in Neo4j")
	neo4jt.AssertRelationship(ctx, t, testDatabase.Driver, orgId, "ASSOCIATED_WITH", locationId)
}
