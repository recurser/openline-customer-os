// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: contact.proto

package contact_grpc_service

import (
	common "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpsertContactGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant    string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Name      string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Prefix    string `protobuf:"bytes,6,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Deprecated: Marked as deprecated in contact.proto.
	AppSource string `protobuf:"bytes,7,opt,name=appSource,proto3" json:"appSource,omitempty"`
	// Deprecated: Marked as deprecated in contact.proto.
	Source string `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
	// Deprecated: Marked as deprecated in contact.proto.
	SourceOfTruth        string                       `protobuf:"bytes,9,opt,name=sourceOfTruth,proto3" json:"sourceOfTruth,omitempty"`
	CreatedAt            *timestamppb.Timestamp       `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamppb.Timestamp       `protobuf:"bytes,11,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Description          string                       `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	Timezone             string                       `protobuf:"bytes,13,opt,name=timezone,proto3" json:"timezone,omitempty"`
	ProfilePhotoUrl      string                       `protobuf:"bytes,14,opt,name=profilePhotoUrl,proto3" json:"profilePhotoUrl,omitempty"`
	SourceFields         *common.SourceFields         `protobuf:"bytes,15,opt,name=sourceFields,proto3" json:"sourceFields,omitempty"`
	ExternalSystemFields *common.ExternalSystemFields `protobuf:"bytes,16,opt,name=externalSystemFields,proto3" json:"externalSystemFields,omitempty"`
	LoggedInUserId       string                       `protobuf:"bytes,17,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
}

func (x *UpsertContactGrpcRequest) Reset() {
	*x = UpsertContactGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertContactGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertContactGrpcRequest) ProtoMessage() {}

func (x *UpsertContactGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertContactGrpcRequest.ProtoReflect.Descriptor instead.
func (*UpsertContactGrpcRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertContactGrpcRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

// Deprecated: Marked as deprecated in contact.proto.
func (x *UpsertContactGrpcRequest) GetAppSource() string {
	if x != nil {
		return x.AppSource
	}
	return ""
}

// Deprecated: Marked as deprecated in contact.proto.
func (x *UpsertContactGrpcRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

// Deprecated: Marked as deprecated in contact.proto.
func (x *UpsertContactGrpcRequest) GetSourceOfTruth() string {
	if x != nil {
		return x.SourceOfTruth
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UpsertContactGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UpsertContactGrpcRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetProfilePhotoUrl() string {
	if x != nil {
		return x.ProfilePhotoUrl
	}
	return ""
}

func (x *UpsertContactGrpcRequest) GetSourceFields() *common.SourceFields {
	if x != nil {
		return x.SourceFields
	}
	return nil
}

func (x *UpsertContactGrpcRequest) GetExternalSystemFields() *common.ExternalSystemFields {
	if x != nil {
		return x.ExternalSystemFields
	}
	return nil
}

func (x *UpsertContactGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

type LinkPhoneNumberToContactGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant         string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	ContactId      string `protobuf:"bytes,2,opt,name=contactId,proto3" json:"contactId,omitempty"`
	PhoneNumberId  string `protobuf:"bytes,3,opt,name=phoneNumberId,proto3" json:"phoneNumberId,omitempty"`
	Primary        bool   `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`
	Label          string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	LoggedInUserId string `protobuf:"bytes,6,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
}

func (x *LinkPhoneNumberToContactGrpcRequest) Reset() {
	*x = LinkPhoneNumberToContactGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkPhoneNumberToContactGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkPhoneNumberToContactGrpcRequest) ProtoMessage() {}

func (x *LinkPhoneNumberToContactGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkPhoneNumberToContactGrpcRequest.ProtoReflect.Descriptor instead.
func (*LinkPhoneNumberToContactGrpcRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{1}
}

func (x *LinkPhoneNumberToContactGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *LinkPhoneNumberToContactGrpcRequest) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

func (x *LinkPhoneNumberToContactGrpcRequest) GetPhoneNumberId() string {
	if x != nil {
		return x.PhoneNumberId
	}
	return ""
}

func (x *LinkPhoneNumberToContactGrpcRequest) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *LinkPhoneNumberToContactGrpcRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *LinkPhoneNumberToContactGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

type LinkEmailToContactGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant         string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	ContactId      string `protobuf:"bytes,2,opt,name=contactId,proto3" json:"contactId,omitempty"`
	EmailId        string `protobuf:"bytes,3,opt,name=emailId,proto3" json:"emailId,omitempty"`
	Primary        bool   `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`
	Label          string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	LoggedInUserId string `protobuf:"bytes,6,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
}

func (x *LinkEmailToContactGrpcRequest) Reset() {
	*x = LinkEmailToContactGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkEmailToContactGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkEmailToContactGrpcRequest) ProtoMessage() {}

func (x *LinkEmailToContactGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkEmailToContactGrpcRequest.ProtoReflect.Descriptor instead.
func (*LinkEmailToContactGrpcRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{2}
}

func (x *LinkEmailToContactGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *LinkEmailToContactGrpcRequest) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

func (x *LinkEmailToContactGrpcRequest) GetEmailId() string {
	if x != nil {
		return x.EmailId
	}
	return ""
}

func (x *LinkEmailToContactGrpcRequest) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *LinkEmailToContactGrpcRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *LinkEmailToContactGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

type LinkLocationToContactGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant         string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	ContactId      string `protobuf:"bytes,2,opt,name=contactId,proto3" json:"contactId,omitempty"`
	LocationId     string `protobuf:"bytes,3,opt,name=locationId,proto3" json:"locationId,omitempty"`
	LoggedInUserId string `protobuf:"bytes,4,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
}

func (x *LinkLocationToContactGrpcRequest) Reset() {
	*x = LinkLocationToContactGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkLocationToContactGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkLocationToContactGrpcRequest) ProtoMessage() {}

func (x *LinkLocationToContactGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkLocationToContactGrpcRequest.ProtoReflect.Descriptor instead.
func (*LinkLocationToContactGrpcRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{3}
}

func (x *LinkLocationToContactGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *LinkLocationToContactGrpcRequest) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

func (x *LinkLocationToContactGrpcRequest) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *LinkLocationToContactGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

type LinkWithOrganizationGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant         string                 `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	ContactId      string                 `protobuf:"bytes,2,opt,name=contactId,proto3" json:"contactId,omitempty"`
	OrganizationId string                 `protobuf:"bytes,3,opt,name=organizationId,proto3" json:"organizationId,omitempty"`
	LoggedInUserId string                 `protobuf:"bytes,4,opt,name=loggedInUserId,proto3" json:"loggedInUserId,omitempty"`
	SourceFields   *common.SourceFields   `protobuf:"bytes,5,opt,name=sourceFields,proto3" json:"sourceFields,omitempty"`
	StartedAt      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	EndedAt        *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=endedAt,proto3" json:"endedAt,omitempty"`
	JobTitle       string                 `protobuf:"bytes,8,opt,name=jobTitle,proto3" json:"jobTitle,omitempty"`
	Primary        bool                   `protobuf:"varint,9,opt,name=primary,proto3" json:"primary,omitempty"`
	Description    string                 `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *LinkWithOrganizationGrpcRequest) Reset() {
	*x = LinkWithOrganizationGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkWithOrganizationGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkWithOrganizationGrpcRequest) ProtoMessage() {}

func (x *LinkWithOrganizationGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkWithOrganizationGrpcRequest.ProtoReflect.Descriptor instead.
func (*LinkWithOrganizationGrpcRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{4}
}

func (x *LinkWithOrganizationGrpcRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *LinkWithOrganizationGrpcRequest) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

func (x *LinkWithOrganizationGrpcRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *LinkWithOrganizationGrpcRequest) GetLoggedInUserId() string {
	if x != nil {
		return x.LoggedInUserId
	}
	return ""
}

func (x *LinkWithOrganizationGrpcRequest) GetSourceFields() *common.SourceFields {
	if x != nil {
		return x.SourceFields
	}
	return nil
}

func (x *LinkWithOrganizationGrpcRequest) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *LinkWithOrganizationGrpcRequest) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

func (x *LinkWithOrganizationGrpcRequest) GetJobTitle() string {
	if x != nil {
		return x.JobTitle
	}
	return ""
}

func (x *LinkWithOrganizationGrpcRequest) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *LinkWithOrganizationGrpcRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LinkWithOrganizationGrpcRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *LinkWithOrganizationGrpcRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ContactIdGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ContactIdGrpcResponse) Reset() {
	*x = ContactIdGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactIdGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactIdGrpcResponse) ProtoMessage() {}

func (x *ContactIdGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactIdGrpcResponse.ProtoReflect.Descriptor instead.
func (*ContactIdGrpcResponse) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{5}
}

func (x *ContactIdGrpcResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_contact_proto protoreflect.FileDescriptor

var file_contact_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x92, 0x05, 0x0a, 0x18, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x20,
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0d,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f,
	0x66, 0x54, 0x72, 0x75, 0x74, 0x68, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55,
	0x72, 0x6c, 0x12, 0x31, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x49, 0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x14, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64,
	0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xd9, 0x01, 0x0a, 0x23, 0x4c, 0x69, 0x6e,
	0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x0e,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0xc7, 0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64,
	0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa0,
	0x01, 0x0a, 0x20, 0x4c, 0x69, 0x6e, 0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x96, 0x04, 0x0a, 0x1f, 0x4c, 0x69, 0x6e, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x0c, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x38,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6a, 0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xa6, 0x03, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47,
	0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58,
	0x0a, 0x18, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x12, 0x4c, 0x69, 0x6e, 0x6b,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1e,
	0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x15, 0x4c, 0x69, 0x6e, 0x6b, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12,
	0x21, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x14, 0x4c, 0x69,
	0x6e, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64,
	0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x42, 0x0c,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contact_proto_rawDescOnce sync.Once
	file_contact_proto_rawDescData = file_contact_proto_rawDesc
)

func file_contact_proto_rawDescGZIP() []byte {
	file_contact_proto_rawDescOnce.Do(func() {
		file_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_contact_proto_rawDescData)
	})
	return file_contact_proto_rawDescData
}

var file_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_contact_proto_goTypes = []interface{}{
	(*UpsertContactGrpcRequest)(nil),            // 0: UpsertContactGrpcRequest
	(*LinkPhoneNumberToContactGrpcRequest)(nil), // 1: LinkPhoneNumberToContactGrpcRequest
	(*LinkEmailToContactGrpcRequest)(nil),       // 2: LinkEmailToContactGrpcRequest
	(*LinkLocationToContactGrpcRequest)(nil),    // 3: LinkLocationToContactGrpcRequest
	(*LinkWithOrganizationGrpcRequest)(nil),     // 4: LinkWithOrganizationGrpcRequest
	(*ContactIdGrpcResponse)(nil),               // 5: ContactIdGrpcResponse
	(*timestamppb.Timestamp)(nil),               // 6: google.protobuf.Timestamp
	(*common.SourceFields)(nil),                 // 7: SourceFields
	(*common.ExternalSystemFields)(nil),         // 8: ExternalSystemFields
}
var file_contact_proto_depIdxs = []int32{
	6,  // 0: UpsertContactGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	6,  // 1: UpsertContactGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: UpsertContactGrpcRequest.sourceFields:type_name -> SourceFields
	8,  // 3: UpsertContactGrpcRequest.externalSystemFields:type_name -> ExternalSystemFields
	7,  // 4: LinkWithOrganizationGrpcRequest.sourceFields:type_name -> SourceFields
	6,  // 5: LinkWithOrganizationGrpcRequest.startedAt:type_name -> google.protobuf.Timestamp
	6,  // 6: LinkWithOrganizationGrpcRequest.endedAt:type_name -> google.protobuf.Timestamp
	6,  // 7: LinkWithOrganizationGrpcRequest.createdAt:type_name -> google.protobuf.Timestamp
	6,  // 8: LinkWithOrganizationGrpcRequest.updatedAt:type_name -> google.protobuf.Timestamp
	0,  // 9: contactGrpcService.UpsertContact:input_type -> UpsertContactGrpcRequest
	1,  // 10: contactGrpcService.LinkPhoneNumberToContact:input_type -> LinkPhoneNumberToContactGrpcRequest
	2,  // 11: contactGrpcService.LinkEmailToContact:input_type -> LinkEmailToContactGrpcRequest
	3,  // 12: contactGrpcService.LinkLocationToContact:input_type -> LinkLocationToContactGrpcRequest
	4,  // 13: contactGrpcService.LinkWithOrganization:input_type -> LinkWithOrganizationGrpcRequest
	5,  // 14: contactGrpcService.UpsertContact:output_type -> ContactIdGrpcResponse
	5,  // 15: contactGrpcService.LinkPhoneNumberToContact:output_type -> ContactIdGrpcResponse
	5,  // 16: contactGrpcService.LinkEmailToContact:output_type -> ContactIdGrpcResponse
	5,  // 17: contactGrpcService.LinkLocationToContact:output_type -> ContactIdGrpcResponse
	5,  // 18: contactGrpcService.LinkWithOrganization:output_type -> ContactIdGrpcResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_contact_proto_init() }
func file_contact_proto_init() {
	if File_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertContactGrpcRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkPhoneNumberToContactGrpcRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkEmailToContactGrpcRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkLocationToContactGrpcRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contact_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkWithOrganizationGrpcRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contact_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactIdGrpcResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contact_proto_goTypes,
		DependencyIndexes: file_contact_proto_depIdxs,
		MessageInfos:      file_contact_proto_msgTypes,
	}.Build()
	File_contact_proto = out.File
	file_contact_proto_rawDesc = nil
	file_contact_proto_goTypes = nil
	file_contact_proto_depIdxs = nil
}
