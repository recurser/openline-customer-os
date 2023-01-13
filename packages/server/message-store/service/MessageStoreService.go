package service

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/proto/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/repository/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type messageService struct {
	msProto.UnimplementedMessageStoreServiceServer
	driver               *neo4j.Driver
	postgresRepositories *repository.PostgresRepositories
	customerOSService    *customerOSService
	commonStoreService   *commonStoreService
}

func (s *messageService) GetMessage(ctx context.Context, msgId *msProto.MessageId) (*msProto.Message, error) {
	if msgId == nil || msgId.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Message ID must be specified")
	}

	//TODO check tenant or conversation
	//exists, err := s.customerOSService.ConversationByIdExists("openline", msgId.GetId())
	//if err != nil {
	//	return nil, err
	//}
	//if !exists {
	//	return nil, status.Errorf(codes.NotFound, "Conversation not found")
	//}

	queryResult := s.postgresRepositories.ConversationEventRepository.GetEventById(msgId.GetId())
	if queryResult.Error != nil {
		return nil, status.Errorf(codes.Internal, queryResult.Error.Error())
	}

	return s.commonStoreService.EncodeConversationEventToMS(*queryResult.Result.(*entity.ConversationEvent)), nil
}

func (s *messageService) GetMessagesForFeed(ctx context.Context, feedIdRequest *msProto.FeedId) (*msProto.MessageListResponse, error) {
	if feedIdRequest == nil || feedIdRequest.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Feed ID must be specified")
	}

	exists, err := s.customerOSService.ConversationByIdExists("openline", feedIdRequest.GetId())
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Conversation not found")
	}

	queryResult := s.postgresRepositories.ConversationEventRepository.GetEventsForConversation(feedIdRequest.GetId())
	if queryResult.Error != nil {
		return nil, status.Errorf(codes.Internal, queryResult.Error.Error())
	}

	var messages []*msProto.Message

	for _, event := range *queryResult.Result.(*[]entity.ConversationEvent) {
		messages = append(messages, s.commonStoreService.EncodeConversationEventToMS(event))
	}

	return &msProto.MessageListResponse{
		Messages: messages,
	}, nil
}

func (s *messageService) GetFeeds(ctx context.Context, request *msProto.GetFeedsPagedRequest) (*msProto.FeedItemPagedResponse, error) {
	conversations, err := s.customerOSService.GetConversations("openline")
	if err != nil {
		return nil, err
	}

	//if feedRequest.GetStateIn() != nil {
	//	stateIn := make([]genConversation.State, 0, len(feedRequest.GetStateIn()))
	//	for _, state := range feedRequest.GetStateIn() {
	//		stateIn = append(stateIn, encodeConversationState(state))
	//	}
	//	query.Where(genConversation.StateIn(stateIn...))
	//}
	//
	//limit := 100 // default to 100 if no pagination is specified
	//if feedRequest.GetPageSize() != 0 {
	//	limit = int(feedRequest.GetPageSize())
	//}
	//offset := limit * int(feedRequest.GetPage())

	fl := &msProto.FeedItemPagedResponse{FeedItems: make([]*msProto.FeedItem, len(conversations))}
	fl.TotalElements = int32(len(conversations))

	for i, conversation := range conversations {
		fl.FeedItems[i] = s.commonStoreService.EncodeConversationToMS(conversation)
	}

	return fl, nil
}

func (s *messageService) GetFeed(ctx context.Context, feedIdRequest *msProto.FeedId) (*msProto.FeedItem, error) {
	if feedIdRequest == nil || feedIdRequest.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Feed ID must be specified")
	}

	conversation, err := s.customerOSService.GetConversationById(feedIdRequest.GetId())
	if err != nil {
		return nil, err
	}

	return s.commonStoreService.EncodeConversationToMS(*conversation), nil
}

func NewMessageService(driver *neo4j.Driver, postgresRepositories *repository.PostgresRepositories, customerOSService *customerOSService, commonStoreService *commonStoreService) *messageService {
	ms := new(messageService)
	ms.driver = driver
	ms.postgresRepositories = postgresRepositories
	ms.customerOSService = customerOSService
	ms.commonStoreService = commonStoreService
	return ms
}
