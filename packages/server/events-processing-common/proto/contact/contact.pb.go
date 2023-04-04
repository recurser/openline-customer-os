// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: proto/contact/contact.proto

package contactGrpcService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID  string `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	UUID      string `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contact_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contact_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_proto_contact_contact_proto_rawDescGZIP(), []int{0}
}

func (x *Contact) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *Contact) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *Contact) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Contact) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type CreateContactGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID      string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	TenantID  string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *CreateContactGrpcRequest) Reset() {
	*x = CreateContactGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contact_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContactGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContactGrpcRequest) ProtoMessage() {}

func (x *CreateContactGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contact_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContactGrpcRequest.ProtoReflect.Descriptor instead.
func (*CreateContactGrpcRequest) Descriptor() ([]byte, []int) {
	return file_proto_contact_contact_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContactGrpcRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *CreateContactGrpcRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *CreateContactGrpcRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateContactGrpcRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type CreateContactGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregateID string `protobuf:"bytes,1,opt,name=AggregateID,proto3" json:"AggregateID,omitempty"`
	UUID        string `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *CreateContactGrpcResponse) Reset() {
	*x = CreateContactGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contact_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContactGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContactGrpcResponse) ProtoMessage() {}

func (x *CreateContactGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contact_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContactGrpcResponse.ProtoReflect.Descriptor instead.
func (*CreateContactGrpcResponse) Descriptor() ([]byte, []int) {
	return file_proto_contact_contact_proto_rawDescGZIP(), []int{2}
}

func (x *CreateContactGrpcResponse) GetAggregateID() string {
	if x != nil {
		return x.AggregateID
	}
	return ""
}

func (x *CreateContactGrpcResponse) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

var File_proto_contact_contact_proto protoreflect.FileDescriptor

var file_proto_contact_contact_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x73, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x51, 0x0a,
	0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44,
	0x32, 0x82, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_contact_contact_proto_rawDescOnce sync.Once
	file_proto_contact_contact_proto_rawDescData = file_proto_contact_contact_proto_rawDesc
)

func file_proto_contact_contact_proto_rawDescGZIP() []byte {
	file_proto_contact_contact_proto_rawDescOnce.Do(func() {
		file_proto_contact_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_contact_contact_proto_rawDescData)
	})
	return file_proto_contact_contact_proto_rawDescData
}

var file_proto_contact_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_contact_contact_proto_goTypes = []interface{}{
	(*Contact)(nil),                   // 0: contactGrpcService.Contact
	(*CreateContactGrpcRequest)(nil),  // 1: contactGrpcService.CreateContactGrpcRequest
	(*CreateContactGrpcResponse)(nil), // 2: contactGrpcService.CreateContactGrpcResponse
}
var file_proto_contact_contact_proto_depIdxs = []int32{
	1, // 0: contactGrpcService.contactGrpcService.CreateContact:input_type -> contactGrpcService.CreateContactGrpcRequest
	2, // 1: contactGrpcService.contactGrpcService.CreateContact:output_type -> contactGrpcService.CreateContactGrpcResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_contact_contact_proto_init() }
func file_proto_contact_contact_proto_init() {
	if File_proto_contact_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_contact_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_proto_contact_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContactGrpcRequest); i {
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
		file_proto_contact_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContactGrpcResponse); i {
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
			RawDescriptor: file_proto_contact_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_contact_contact_proto_goTypes,
		DependencyIndexes: file_proto_contact_contact_proto_depIdxs,
		MessageInfos:      file_proto_contact_contact_proto_msgTypes,
	}.Build()
	File_proto_contact_contact_proto = out.File
	file_proto_contact_contact_proto_rawDesc = nil
	file_proto_contact_contact_proto_goTypes = nil
	file_proto_contact_contact_proto_depIdxs = nil
}