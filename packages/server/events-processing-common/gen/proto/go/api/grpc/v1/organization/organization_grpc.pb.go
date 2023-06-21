// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: organization.proto

package organization_grpc_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrganizationGrpcServiceClient is the client API for OrganizationGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationGrpcServiceClient interface {
	UpsertOrganization(ctx context.Context, in *UpsertOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	LinkPhoneNumberToOrganization(ctx context.Context, in *LinkPhoneNumberToOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
	LinkEmailToOrganization(ctx context.Context, in *LinkEmailToOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error)
}

type organizationGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationGrpcServiceClient(cc grpc.ClientConnInterface) OrganizationGrpcServiceClient {
	return &organizationGrpcServiceClient{cc}
}

func (c *organizationGrpcServiceClient) UpsertOrganization(ctx context.Context, in *UpsertOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/UpsertOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) LinkPhoneNumberToOrganization(ctx context.Context, in *LinkPhoneNumberToOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/LinkPhoneNumberToOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationGrpcServiceClient) LinkEmailToOrganization(ctx context.Context, in *LinkEmailToOrganizationGrpcRequest, opts ...grpc.CallOption) (*OrganizationIdGrpcResponse, error) {
	out := new(OrganizationIdGrpcResponse)
	err := c.cc.Invoke(ctx, "/organizationGrpcService/LinkEmailToOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationGrpcServiceServer is the server API for OrganizationGrpcService service.
// All implementations should embed UnimplementedOrganizationGrpcServiceServer
// for forward compatibility
type OrganizationGrpcServiceServer interface {
	UpsertOrganization(context.Context, *UpsertOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error)
	LinkPhoneNumberToOrganization(context.Context, *LinkPhoneNumberToOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error)
	LinkEmailToOrganization(context.Context, *LinkEmailToOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error)
}

// UnimplementedOrganizationGrpcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrganizationGrpcServiceServer struct {
}

func (UnimplementedOrganizationGrpcServiceServer) UpsertOrganization(context.Context, *UpsertOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) LinkPhoneNumberToOrganization(context.Context, *LinkPhoneNumberToOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkPhoneNumberToOrganization not implemented")
}
func (UnimplementedOrganizationGrpcServiceServer) LinkEmailToOrganization(context.Context, *LinkEmailToOrganizationGrpcRequest) (*OrganizationIdGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkEmailToOrganization not implemented")
}

// UnsafeOrganizationGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationGrpcServiceServer will
// result in compilation errors.
type UnsafeOrganizationGrpcServiceServer interface {
	mustEmbedUnimplementedOrganizationGrpcServiceServer()
}

func RegisterOrganizationGrpcServiceServer(s grpc.ServiceRegistrar, srv OrganizationGrpcServiceServer) {
	s.RegisterService(&OrganizationGrpcService_ServiceDesc, srv)
}

func _OrganizationGrpcService_UpsertOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).UpsertOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/UpsertOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).UpsertOrganization(ctx, req.(*UpsertOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_LinkPhoneNumberToOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkPhoneNumberToOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).LinkPhoneNumberToOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/LinkPhoneNumberToOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).LinkPhoneNumberToOrganization(ctx, req.(*LinkPhoneNumberToOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationGrpcService_LinkEmailToOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkEmailToOrganizationGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationGrpcServiceServer).LinkEmailToOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizationGrpcService/LinkEmailToOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationGrpcServiceServer).LinkEmailToOrganization(ctx, req.(*LinkEmailToOrganizationGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationGrpcService_ServiceDesc is the grpc.ServiceDesc for OrganizationGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "organizationGrpcService",
	HandlerType: (*OrganizationGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertOrganization",
			Handler:    _OrganizationGrpcService_UpsertOrganization_Handler,
		},
		{
			MethodName: "LinkPhoneNumberToOrganization",
			Handler:    _OrganizationGrpcService_LinkPhoneNumberToOrganization_Handler,
		},
		{
			MethodName: "LinkEmailToOrganization",
			Handler:    _OrganizationGrpcService_LinkEmailToOrganization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "organization.proto",
}
