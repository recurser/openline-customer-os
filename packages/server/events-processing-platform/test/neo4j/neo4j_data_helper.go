package neo4j

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/graph_db/entity"
	"github.com/stretchr/testify/require"
	"reflect"
	"sort"
	"testing"
)

func CleanupAllData(ctx context.Context, driver *neo4j.DriverWithContext) {
	ExecuteWriteQuery(ctx, driver, `MATCH (n) DETACH DELETE n`, map[string]any{})
}

func CreateTenant(ctx context.Context, driver *neo4j.DriverWithContext, tenant string) {
	query := `MERGE (t:Tenant {name:$tenant})`
	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant": tenant,
	})
}

func CreateOrganization(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, organization entity.OrganizationEntity) string {
	orgId := organization.ID
	if orgId == "" {
		orgId = uuid.New().String()
	}
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:ORGANIZATION_BELONGS_TO_TENANT]-(o:Organization {id:$id})
				SET o:Organization_%s,
					o.name=$name,
					o.hide=$hide,
					o.renewalLikelihood=$renewalLikelihood,
					o.renewalLikelihoodComment=$renewalLikelihoodComment,
					o.renewalLikelihoodUpdatedAt=$renewalLikelihoodUpdatedAt,
					o.renewalLikelihoodUpdatedBy=$renewalLikelihoodUpdatedBy,
					o.renewalForecastAmount=$renewalForecastAmount,
					o.renewalForecastPotentialAmount=$renewalForecastPotentialAmount,
					o.renewalForecastUpdatedAt=$renewalForecastUpdatedAt,
					o.renewalForecastUpdatedBy=$renewalForecastUpdatedBy,
					o.renewalForecastComment=$renewalForecastComment,
					o.billingDetailsAmount=$billingAmount, 
					o.billingDetailsFrequency=$billingFrequency, 
					o.billingDetailsRenewalCycle=$billingRenewalCycle, 
			 		o.billingDetailsRenewalCycleStart=$billingRenewalCycleStart,
			 		o.billingDetailsRenewalCycleNext=$billingRenewalCycleNext
				`, tenant)

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":                         tenant,
		"name":                           organization.Name,
		"hide":                           organization.Hide,
		"id":                             orgId,
		"renewalLikelihood":              organization.RenewalLikelihood.RenewalLikelihood,
		"renewalLikelihoodPrevious":      organization.RenewalLikelihood.PreviousRenewalLikelihood,
		"renewalLikelihoodComment":       organization.RenewalLikelihood.Comment,
		"renewalLikelihoodUpdatedAt":     organization.RenewalLikelihood.UpdatedAt,
		"renewalLikelihoodUpdatedBy":     organization.RenewalLikelihood.UpdatedBy,
		"renewalForecastAmount":          organization.RenewalForecast.Amount,
		"renewalForecastPotentialAmount": organization.RenewalForecast.PotentialAmount,
		"renewalForecastUpdatedBy":       organization.RenewalForecast.UpdatedBy,
		"renewalForecastUpdatedAt":       organization.RenewalForecast.UpdatedAt,
		"renewalForecastComment":         organization.RenewalForecast.Comment,
		"billingAmount":                  organization.BillingDetails.Amount,
		"billingFrequency":               organization.BillingDetails.Frequency,
		"billingRenewalCycle":            organization.BillingDetails.RenewalCycle,
		"billingRenewalCycleStart":       utils.TimePtrFirstNonNilNillableAsAny(organization.BillingDetails.RenewalCycleStart),
		"billingRenewalCycleNext":        utils.TimePtrFirstNonNilNillableAsAny(organization.BillingDetails.RenewalCycleNext),
	})
	return orgId
}

func CreateUser(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, user entity.UserEntity) string {
	userId := user.Id
	if userId == "" {
		userId = uuid.New().String()
	}
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:USER_BELONGS_TO_TENANT]-(u:User {id:$id})
				SET u:User_%s,
					u.firstName=$firstName,
					u.lastName=$lastName
				`, tenant)

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":    tenant,
		"id":        userId,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
	})
	return userId
}

func CreateSocial(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, social entity.SocialEntity) string {
	socialId := utils.NewUUIDIfEmpty(social.Id)
	query := fmt.Sprintf(`MERGE (s:Social:Social_%s {id: $id})
				SET s.url=$url,
					s.platformName=$platformName
				`, tenant)

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"id":           socialId,
		"url":          social.Url,
		"platformName": social.PlatformName,
	})
	return socialId
}

func CreateContact(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, contact entity.ContactEntity) string {
	contactId := contact.Id
	if contactId == "" {
		contactId = uuid.New().String()
	}
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:CONTACT_BELONGS_TO_TENANT]-(c:Contact {id:$id})
				SET c:Contact_%s,
					c.firstName=$firstName,
					c.lastName=$lastName
				`, tenant)

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":    tenant,
		"id":        contactId,
		"firstName": contact.FirstName,
		"lastName":  contact.LastName,
	})
	return contactId
}

func CreateLogEntryForOrg(ctx context.Context, driver *neo4j.DriverWithContext, tenant, orgId string, logEntry entity.LogEntryEntity) string {
	logEntryId := logEntry.Id
	if logEntryId == "" {
		logEntryId = uuid.New().String()
	}
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(o:Organization {id:$orgId})
			  MERGE (o)-[:LOGGED]->(l:LogEntry {id:$id})
				SET l:LogEntry_%s,
					l:TimelineEvent,
					l:TimelineEvent_%s,
					l.content=$content,
					l.contentType=$contentType,
					l.startedAt=$startedAt
				`, tenant, tenant)

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":      tenant,
		"orgId":       orgId,
		"id":          logEntryId,
		"content":     logEntry.Content,
		"contentType": logEntry.ContentType,
		"startedAt":   logEntry.StartedAt,
	})
	return logEntryId
}

func CreateIssue(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, issue entity.IssueEntity) string {
	issueId := utils.NewUUIDIfEmpty(issue.Id)
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:ISSUE_BELONGS_TO_TENANT]-(i:Issue {id:$id})
				SET i:Issue_%s,
					i:TimelineEvent,
					i:TimelineEvent_%s,
					i.subject=$subject,
					i.status=$status,
					i.priority=$priority,
					i.description=$description,
					i.source=$source,
					i.sourceOfTruth=$sourceOfTruth
				`, tenant, tenant)

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":        tenant,
		"id":            issueId,
		"subject":       issue.Subject,
		"status":        issue.Status,
		"priority":      issue.Priority,
		"description":   issue.Description,
		"source":        issue.Source,
		"sourceOfTruth": issue.SourceOfTruth,
	})
	return issueId
}

func CreateComment(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, comment entity.CommentEntity) string {
	commentId := utils.NewUUIDIfEmpty(comment.Id)
	query := fmt.Sprintf(`MERGE (c:Comment:Comment_%s {id:$id})
				SET c.content=$content,
					c.contentType=$contentType,
					c.source=$source,
					c.sourceOfTruth=$sourceOfTruth
				`, tenant)

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"id":            commentId,
		"content":       comment.Content,
		"contentType":   comment.ContentType,
		"source":        comment.Source,
		"sourceOfTruth": comment.SourceOfTruth,
	})
	return commentId
}

func CreatePhoneNumber(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, phoneNumber entity.PhoneNumberEntity) string {
	phoneNumberId := utils.NewUUIDIfEmpty(phoneNumber.Id)
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:PHONE_NUMBER_BELONGS_TO_TENANT]-(i:PhoneNumber {id:$id})
				SET i:PhoneNumber_%s,
					i.e164=$e164,
					i.validated=$validated,
					i.rawPhoneNumber=$rawPhoneNumber,
					i.source=$source,
					i.sourceOfTruth=$sourceOfTruth,
					i.appSource=$appSource,
					i.createdAt=$createdAt,
					i.updatedAt=$updatedAt`, tenant)

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":         tenant,
		"id":             phoneNumberId,
		"e164":           phoneNumber.E164,
		"validated":      phoneNumber.Validated,
		"rawPhoneNumber": phoneNumber.RawPhoneNumber,
		"source":         phoneNumber.Source,
		"sourceOfTruth":  phoneNumber.SourceOfTruth,
		"appSource":      phoneNumber.AppSource,
		"createdAt":      phoneNumber.CreatedAt,
		"updatedAt":      phoneNumber.UpdatedAt,
	})
	return phoneNumberId
}

func CreateLocation(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, location entity.LocationEntity) string {
	locationId := utils.NewUUIDIfEmpty(location.Id)
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:LOCATION_BELONGS_TO_TENANT]-(i:Location {id:$id})
				SET i:Location_%s,
					i.name=$name,
					i.createdAt=$createdAt,
					i.updatedAt=$updatedAt,
					i.country=$country,
					i.region=$region,    
					i.locality=$locality,    
					i.address=$address,    
					i.address2=$address2,    
					i.zip=$zip,    
					i.addressType=$addressType,    
					i.houseNumber=$houseNumber,    
					i.postalCode=$postalCode,    
					i.plusFour=$plusFour,    
					i.commercial=$commercial,    
					i.predirection=$predirection,    
					i.district=$district,    
					i.street=$street,    
					i.rawAddress=$rawAddress,    
					i.latitude=$latitude,    
					i.longitude=$longitude,    
					i.timeZone=$timeZone,    
					i.utcOffset=$utcOffset,    
					i.sourceOfTruth=$sourceOfTruth,
					i.source=$source,
					i.appSource=$appSource`, tenant)

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":        tenant,
		"id":            locationId,
		"name":          location.Name,
		"createdAt":     location.CreatedAt,
		"updatedAt":     location.UpdatedAt,
		"country":       location.Country,
		"region":        location.Region,
		"locality":      location.Locality,
		"address":       location.Address,
		"address2":      location.Address2,
		"zip":           location.Zip,
		"addressType":   location.AddressType,
		"houseNumber":   location.HouseNumber,
		"postalCode":    location.PostalCode,
		"plusFour":      location.PlusFour,
		"commercial":    location.Commercial,
		"predirection":  location.Predirection,
		"district":      location.District,
		"street":        location.Street,
		"rawAddress":    location.RawAddress,
		"latitude":      location.Latitude,
		"longitude":     location.Longitude,
		"timeZone":      location.TimeZone,
		"utcOffset":     location.UtcOffset,
		"sourceOfTruth": location.SourceOfTruth,
		"source":        location.Source,
		"appSource":     location.AppSource,
	})
	return locationId

}

func LinkIssueReportedBy(ctx context.Context, driver *neo4j.DriverWithContext, issueId, entityId string) {
	query := `MATCH (e {id:$entityId})
				MATCH (i:Issue {id:$issueId})
				MERGE (i)-[:REPORTED_BY]->(e)`

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"issueId":  issueId,
		"entityId": entityId,
	})
}

func LinkIssueAssignedTo(ctx context.Context, driver *neo4j.DriverWithContext, issueId, entityId string) {
	query := `MATCH (e {id:$entityId})
				MATCH (i:Issue {id:$issueId})
				MERGE (i)-[:ASSIGNED_TO]->(e)`

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"issueId":  issueId,
		"entityId": entityId,
	})
}

func LinkIssueFollowedBy(ctx context.Context, driver *neo4j.DriverWithContext, issueId, entityId string) {
	query := `MATCH (e {id:$entityId})
				MATCH (i:Issue {id:$issueId})
				MERGE (i)-[:FOLLOWED_BY]->(e)`

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"issueId":  issueId,
		"entityId": entityId,
	})
}

func CreateInteractionEvent(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, interactionEvent entity.InteractionEventEntity) string {
	interactionEventId := utils.NewUUIDIfEmpty(interactionEvent.Id)
	query := fmt.Sprintf(`MERGE (i:InteractionEvent {id:$id})
				SET i:InteractionEvent_%s,
					i:TimelineEvent,
					i:TimelineEvent_%s,
					i.content=$content,
					i.contentType=$contentType,
					i.channel=$channel,
					i.channelData=$channelData,
					i.identifier=$identifier,
					i.eventType=$eventType,
					i.source=$source,
					i.sourceOfTruth=$sourceOfTruth
				`, tenant, tenant)

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"id":            interactionEventId,
		"content":       interactionEvent.Content,
		"contentType":   interactionEvent.ContentType,
		"channel":       interactionEvent.Channel,
		"channelData":   interactionEvent.ChannelData,
		"identifier":    interactionEvent.Identifier,
		"eventType":     interactionEvent.EventType,
		"source":        interactionEvent.Source,
		"sourceOfTruth": interactionEvent.SourceOfTruth,
	})
	return interactionEventId
}

func CreateTag(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, tag entity.TagEntity) string {
	tagId := tag.Id
	if tagId == "" {
		tagId = uuid.New().String()
	}

	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:TAG_BELONGS_TO_TENANT]-(tag:Tag {id:$tagId})
				SET tag:Tag_%s,
					tag.name=$name`, tenant)

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant": tenant,
		"tagId":  tagId,
		"name":   tag.Name,
	})
	return tagId
}

func LinkTag(ctx context.Context, driver *neo4j.DriverWithContext, tagId, entityId string) {

	query := `MATCH (e {id:$entityId})
				MATCH (t:Tag {id:$tagId})
				MERGE (e)-[rel:TAGGED]->(t)
				SET rel.taggedAt=$now`

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tagId":    tagId,
		"entityId": entityId,
		"now":      utils.Now(),
	})
}

func LinkSocial(ctx context.Context, driver *neo4j.DriverWithContext, socialId, entityId string) {
	query := `MATCH (e {id:$entityId})
				MATCH (s:Social {id:$socialId})
				MERGE (e)-[:HAS]->(s)`

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"socialId": socialId,
		"entityId": entityId,
	})
}

func CreateExternalSystem(ctx context.Context, driver *neo4j.DriverWithContext, tenant, externalSystem string) {
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(ext:ExternalSystem {id:$externalSystemId})
				ON CREATE SET ext:ExternalSystem_%s`, tenant)

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":           tenant,
		"externalSystemId": externalSystem,
	})
}

func CreateWorkspace(ctx context.Context, driver *neo4j.DriverWithContext, workspace string, provider string, tenant string) {
	query := `MATCH (t:Tenant {name: $tenant})
			  MERGE (t)-[:HAS_WORKSPACE]->(w:Workspace {name:$workspace, provider:$provider})`

	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":    tenant,
		"provider":  provider,
		"workspace": workspace,
	})
}

func GetNodeById(ctx context.Context, driver *neo4j.DriverWithContext, label string, id string) (*dbtype.Node, error) {
	session := utils.NewNeo4jReadSession(ctx, *driver)
	defer session.Close(ctx)

	queryResult, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		result, err := tx.Run(ctx, fmt.Sprintf(`
			MATCH (n:%s {id:$id}) RETURN n`, label),
			map[string]interface{}{
				"id": id,
			})
		record, err := result.Single(ctx)
		if err != nil {
			return nil, err
		}
		return record.Values[0], nil
	})
	if err != nil {
		return nil, err
	}

	node := queryResult.(dbtype.Node)
	return &node, nil
}

func GetRelationship(ctx context.Context, driver *neo4j.DriverWithContext, fromNodeId, toNodeId string) (*dbtype.Relationship, error) {
	session := utils.NewNeo4jReadSession(ctx, *driver)
	defer session.Close(ctx)

	queryResult, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		result, err := tx.Run(ctx, `MATCH (n {id:$fromNodeId})-[rel]->(m {id:$toNodeId}) RETURN rel limit 1`,
			map[string]interface{}{
				"fromNodeId": fromNodeId,
				"toNodeId":   toNodeId,
			})
		record, err := result.Single(ctx)
		if err != nil {
			return nil, err
		}
		return record.Values[0], nil
	})
	if err != nil {
		return nil, err
	}

	node := queryResult.(dbtype.Relationship)
	return &node, nil
}

func GetFirstNodeByLabel(ctx context.Context, driver *neo4j.DriverWithContext, label string) (*dbtype.Node, error) {
	session := utils.NewNeo4jReadSession(ctx, *driver)
	defer session.Close(ctx)

	queryResult, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		result, err := tx.Run(ctx, fmt.Sprintf(`
			MATCH (n:%s) RETURN n limit 1`, label),
			map[string]interface{}{})
		record, err := result.Single(ctx)
		if err != nil {
			return nil, err
		}
		return record.Values[0], nil
	})
	if err != nil {
		return nil, err
	}

	node := queryResult.(dbtype.Node)
	return &node, nil
}

func GetCountOfNodes(ctx context.Context, driver *neo4j.DriverWithContext, nodeLabel string) int {
	query := fmt.Sprintf(`MATCH (n:%s) RETURN count(n)`, nodeLabel)
	result := ExecuteReadQueryWithSingleReturn(ctx, driver, query, map[string]any{})
	return int(result.(*db.Record).Values[0].(int64))
}

func GetCountOfRelationships(ctx context.Context, driver *neo4j.DriverWithContext, relationship string) int {
	query := fmt.Sprintf(`MATCH (a)-[r:%s]-(b) RETURN count(distinct r)`, relationship)
	result := ExecuteReadQueryWithSingleReturn(ctx, driver, query, map[string]any{})
	return int(result.(*db.Record).Values[0].(int64))
}

func GetTotalCountOfNodes(ctx context.Context, driver *neo4j.DriverWithContext) int {
	query := `MATCH (n) RETURN count(n)`
	result := ExecuteReadQueryWithSingleReturn(ctx, driver, query, map[string]any{})
	return int(result.(*db.Record).Values[0].(int64))
}

func GetAllLabels(ctx context.Context, driver *neo4j.DriverWithContext) []string {
	query := `MATCH (n) RETURN DISTINCT labels(n)`
	dbRecords := ExecuteReadQueryWithCollectionReturn(ctx, driver, query, map[string]any{})
	labels := []string{}
	for _, v := range dbRecords {
		for _, nodeLabels := range v.Values {
			for _, label := range nodeLabels.([]interface{}) {
				if !contains(labels, label.(string)) {
					labels = append(labels, label.(string))
				}
			}
		}
	}
	return labels
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func AssertNeo4jLabels(ctx context.Context, t *testing.T, driver *neo4j.DriverWithContext, expectedLabels []string) {
	actualLabels := GetAllLabels(ctx, driver)
	sort.Strings(expectedLabels)
	sort.Strings(actualLabels)
	if !reflect.DeepEqual(actualLabels, expectedLabels) {
		t.Errorf("Expected labels: %v, \nActual labels: %v", expectedLabels, actualLabels)
	}
}

func AssertNeo4jNodeCount(ctx context.Context, t *testing.T, driver *neo4j.DriverWithContext, nodes map[string]int) {
	for name, expectedCount := range nodes {
		actualCount := GetCountOfNodes(ctx, driver, name)
		require.Equal(t, expectedCount, actualCount, "Unexpected count for node: "+name)
	}
}

func AssertNeo4jRelationCount(ctx context.Context, t *testing.T, driver *neo4j.DriverWithContext, relations map[string]int) {
	for name, expectedCount := range relations {
		actualCount := GetCountOfRelationships(ctx, driver, name)
		require.Equal(t, expectedCount, actualCount, "Unexpected count for relationship: "+name)
	}
}

func AssertRelationship(ctx context.Context, t *testing.T, driver *neo4j.DriverWithContext, fromNodeId, relationshipType, toNodeId string) {
	rel, err := GetRelationship(ctx, driver, fromNodeId, toNodeId)
	require.Nil(t, err)
	require.NotNil(t, rel)
	require.Equal(t, relationshipType, rel.Type)
}

func AssertRelationshipWithProperties(ctx context.Context, t *testing.T, driver *neo4j.DriverWithContext, fromNodeId, relationshipType, toNodeId string, expectedProperties map[string]any) {
	rel, err := GetRelationship(ctx, driver, fromNodeId, toNodeId)
	require.Nil(t, err)
	require.NotNil(t, rel)
	require.Equal(t, relationshipType, rel.Type)
	for k, v := range expectedProperties {
		require.Equal(t, v, rel.Props[k])
	}
}

func CreateCountry(ctx context.Context, driver *neo4j.DriverWithContext, codeA2, codeA3, name, phoneCode string) {
	query := `MERGE (c:Country{codeA3: $codeA3}) 
				ON CREATE SET 
					c.phoneCode = $phoneCode,
					c.codeA2 = $codeA2,
					c.name = $name, 
					c.createdAt = $now, 
					c.updatedAt = $now`
	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"codeA2":    codeA2,
		"codeA3":    codeA3,
		"phoneCode": phoneCode,
		"name":      name,
		"now":       utils.Now(),
	})
}
