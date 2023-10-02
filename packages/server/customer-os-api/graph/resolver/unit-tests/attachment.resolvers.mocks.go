package unit_tests

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/stretchr/testify/mock"
)

type MockedAttachmentService struct {
	mock.Mock
}

func (s *MockedAttachmentService) GetAttachmentById(ctx context.Context, id string) (*entity.AttachmentEntity, error) {
	return nil, nil
}

func (s *MockedAttachmentService) Create(ctx context.Context, newEntity *entity.AttachmentEntity, source, sourceOfTruth entity.DataSource) (*entity.AttachmentEntity, error) {
	return &entity.AttachmentEntity{
		MimeType:  newEntity.MimeType,
		Name:      newEntity.Name,
		Size:      newEntity.Size,
		Extension: newEntity.Extension,
		AppSource: newEntity.AppSource,
	}, nil
}
func (s *MockedAttachmentService) GetAttachmentsForNode(ctx context.Context, linkedWith repository.LinkedWith, linkedNature *repository.LinkedNature, ids []string) (*entity.AttachmentEntities, error) {
	return nil, nil
}

func (s *MockedAttachmentService) LinkNodeWithAttachment(ctx context.Context, linkedWith repository.LinkedWith, linkedNature *repository.LinkedNature, attachmentId, includedById string) (*dbtype.Node, error) {
	return nil, nil
}
func (s *MockedAttachmentService) UnlinkNodeWithAttachment(ctx context.Context, linkedWith repository.LinkedWith, linkedNature *repository.LinkedNature, attachmentId, includedById string) (*dbtype.Node, error) {
	return nil, nil
}

func (s *MockedAttachmentService) MapDbNodeToAttachmentEntity(node dbtype.Node) *entity.AttachmentEntity {
	return nil
}
