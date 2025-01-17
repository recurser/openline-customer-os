package servicet

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	jobrolepb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/job_role"
	userpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/user"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/user/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/user/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/test/eventstore"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestUserService_UpsertUser(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()
	grpcConnection, err := dialFactory.GetEventsProcessingPlatformConn(testDatabase.Repositories, aggregateStore)
	if err != nil {
		t.Fatalf("Failed to connect to events processing platform: %v", err)
	}
	userClient := userpb.NewUserGrpcServiceClient(grpcConnection)
	timeNow := time.Now().UTC()
	userId, _ := uuid.NewUUID()

	response, err := userClient.UpsertUser(ctx, &userpb.UpsertUserGrpcRequest{
		Id:              userId.String(),
		Tenant:          "ziggy",
		FirstName:       "Bob",
		LastName:        "Dole",
		Name:            "Bob Dole",
		AppSource:       "unit-test",
		Source:          "N/A",
		SourceOfTruth:   "N/A",
		Internal:        true,
		ProfilePhotoUrl: "https://www.google.com",
		Timezone:        "America/Los_Angeles",
		CreatedAt:       timestamppb.New(timeNow),
		UpdatedAt:       timestamppb.New(timeNow),
	})

	require.Nil(t, err)
	require.NotNil(t, response)
	require.Equal(t, userId.String(), response.Id)
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 1, len(eventsMap))
	eventList := eventsMap[aggregate.NewUserAggregateWithTenantAndID("ziggy", userId.String()).ID]
	require.Equal(t, 1, len(eventList))
	require.Equal(t, events.UserCreateV1, eventList[0].EventType)
	var eventData events.UserCreateEvent
	err = eventList[0].GetJsonData(&eventData)
	fmt.Printf("Got an envent %s\n", string(eventList[0].GetData()))
	if err := eventList[0].GetJsonData(&eventData); err != nil {
		t.Errorf("Failed to unmarshal event data: %v", err)
	}
	require.Equal(t, "Bob", eventData.FirstName)
	require.Equal(t, "Dole", eventData.LastName)
	require.Equal(t, "Bob Dole", eventData.Name)
	require.Equal(t, "unit-test", eventData.SourceFields.AppSource)
	require.Equal(t, "N/A", eventData.SourceFields.Source)
	require.Equal(t, "N/A", eventData.SourceFields.SourceOfTruth)
	require.Equal(t, timeNow, eventData.CreatedAt)
	require.Equal(t, timeNow, eventData.UpdatedAt)
	require.Equal(t, "ziggy", eventData.Tenant)
	require.True(t, eventData.Internal)
	require.Equal(t, "https://www.google.com", eventData.ProfilePhotoUrl)
	require.Equal(t, "America/Los_Angeles", eventData.Timezone)
}

func TestUserService_UpsertUserAndLinkJobRole(t *testing.T) {
	ctx := context.TODO()
	defer tearDownTestCase(ctx, testDatabase)(t)

	aggregateStore := eventstore.NewTestAggregateStore()
	grpcConnection, err := dialFactory.GetEventsProcessingPlatformConn(testDatabase.Repositories, aggregateStore)
	if err != nil {
		t.Fatalf("Failed to connect to events processing platform: %v", err)
	}
	userClient := userpb.NewUserGrpcServiceClient(grpcConnection)
	jobRoleClient := jobrolepb.NewJobRoleGrpcServiceClient(grpcConnection)

	timeNow := time.Now().UTC()
	userId, _ := uuid.NewUUID()

	createUserResponse, err := userClient.UpsertUser(ctx, &userpb.UpsertUserGrpcRequest{
		Id:            userId.String(),
		Tenant:        "ziggy",
		FirstName:     "Bob",
		LastName:      "Dole",
		Name:          "Bob Dole",
		AppSource:     "unit-test",
		Source:        "N/A",
		SourceOfTruth: "N/A",
		CreatedAt:     timestamppb.New(timeNow),
		UpdatedAt:     timestamppb.New(timeNow),
	})

	require.Nil(t, err)
	require.NotNil(t, createUserResponse)
	require.Equal(t, userId.String(), createUserResponse.Id)

	timeStarted := time.Now().UTC().AddDate(0, -6, 0)
	timeEnded := time.Now().UTC().AddDate(0, 6, 0)
	description := "I clean things"
	createJobRoleResponse, err := jobRoleClient.CreateJobRole(ctx, &jobrolepb.CreateJobRoleGrpcRequest{
		Tenant:        "ziggy",
		JobTitle:      "Chief Janitor",
		Description:   &description,
		Source:        "N/A",
		SourceOfTruth: "N/A",
		AppSource:     "unit-test",
		CreatedAt:     timestamppb.New(timeNow),
		StartedAt:     timestamppb.New(timeStarted),
		EndedAt:       timestamppb.New(timeEnded),
	})
	if err != nil {
		t.Fatalf("Failed to create job role: %v", err)
	}
	require.Nil(t, err)
	require.NotNil(t, createJobRoleResponse)

	linkJobRoleResponse, err := userClient.LinkJobRoleToUser(ctx, &userpb.LinkJobRoleToUserGrpcRequest{
		UserId:    createUserResponse.Id,
		JobRoleId: createJobRoleResponse.Id,
		Tenant:    "ziggy",
	})

	if err != nil {
		t.Fatalf("Failed to link job role to user: %v", err)
	}
	require.Nil(t, err)
	require.NotNil(t, linkJobRoleResponse)
	eventsMap := aggregateStore.GetEventMap()
	require.Equal(t, 2, len(eventsMap))
	eventList := eventsMap[aggregate.NewUserAggregateWithTenantAndID("ziggy", userId.String()).ID]
	require.Equal(t, 2, len(eventList))
	require.Equal(t, events.UserCreateV1, eventList[0].EventType)
	require.Equal(t, events.UserJobRoleLinkV1, eventList[1].EventType)
}
