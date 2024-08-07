// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: flagrant/flagrant.proto

package flagrant

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

// EmailReminderServiceClient is the client API for EmailReminderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailReminderServiceClient interface {
	SendNotificationsGroups(ctx context.Context, in *SendNotificationsGroupsRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	SendNotificationsGroup(ctx context.Context, in *SendNotificationsGroupRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	SendNotificationsCoaches(ctx context.Context, in *SendNotificationsCoachesRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	SendNotificationsCoach(ctx context.Context, in *SendNotificationsCoachRequest, opts ...grpc.CallOption) (*BaseResponse, error)
}

type emailReminderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailReminderServiceClient(cc grpc.ClientConnInterface) EmailReminderServiceClient {
	return &emailReminderServiceClient{cc}
}

func (c *emailReminderServiceClient) SendNotificationsGroups(ctx context.Context, in *SendNotificationsGroupsRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/flagrant.EmailReminderService/SendNotificationsGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailReminderServiceClient) SendNotificationsGroup(ctx context.Context, in *SendNotificationsGroupRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/flagrant.EmailReminderService/SendNotificationsGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailReminderServiceClient) SendNotificationsCoaches(ctx context.Context, in *SendNotificationsCoachesRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/flagrant.EmailReminderService/SendNotificationsCoaches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailReminderServiceClient) SendNotificationsCoach(ctx context.Context, in *SendNotificationsCoachRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/flagrant.EmailReminderService/SendNotificationsCoach", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailReminderServiceServer is the server API for EmailReminderService service.
// All implementations must embed UnimplementedEmailReminderServiceServer
// for forward compatibility
type EmailReminderServiceServer interface {
	SendNotificationsGroups(context.Context, *SendNotificationsGroupsRequest) (*BaseResponse, error)
	SendNotificationsGroup(context.Context, *SendNotificationsGroupRequest) (*BaseResponse, error)
	SendNotificationsCoaches(context.Context, *SendNotificationsCoachesRequest) (*BaseResponse, error)
	SendNotificationsCoach(context.Context, *SendNotificationsCoachRequest) (*BaseResponse, error)
	mustEmbedUnimplementedEmailReminderServiceServer()
}

// UnimplementedEmailReminderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmailReminderServiceServer struct {
}

func (UnimplementedEmailReminderServiceServer) SendNotificationsGroups(context.Context, *SendNotificationsGroupsRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotificationsGroups not implemented")
}
func (UnimplementedEmailReminderServiceServer) SendNotificationsGroup(context.Context, *SendNotificationsGroupRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotificationsGroup not implemented")
}
func (UnimplementedEmailReminderServiceServer) SendNotificationsCoaches(context.Context, *SendNotificationsCoachesRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotificationsCoaches not implemented")
}
func (UnimplementedEmailReminderServiceServer) SendNotificationsCoach(context.Context, *SendNotificationsCoachRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotificationsCoach not implemented")
}
func (UnimplementedEmailReminderServiceServer) mustEmbedUnimplementedEmailReminderServiceServer() {}

// UnsafeEmailReminderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailReminderServiceServer will
// result in compilation errors.
type UnsafeEmailReminderServiceServer interface {
	mustEmbedUnimplementedEmailReminderServiceServer()
}

func RegisterEmailReminderServiceServer(s grpc.ServiceRegistrar, srv EmailReminderServiceServer) {
	s.RegisterService(&EmailReminderService_ServiceDesc, srv)
}

func _EmailReminderService_SendNotificationsGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotificationsGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailReminderServiceServer).SendNotificationsGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flagrant.EmailReminderService/SendNotificationsGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailReminderServiceServer).SendNotificationsGroups(ctx, req.(*SendNotificationsGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailReminderService_SendNotificationsGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotificationsGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailReminderServiceServer).SendNotificationsGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flagrant.EmailReminderService/SendNotificationsGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailReminderServiceServer).SendNotificationsGroup(ctx, req.(*SendNotificationsGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailReminderService_SendNotificationsCoaches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotificationsCoachesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailReminderServiceServer).SendNotificationsCoaches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flagrant.EmailReminderService/SendNotificationsCoaches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailReminderServiceServer).SendNotificationsCoaches(ctx, req.(*SendNotificationsCoachesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailReminderService_SendNotificationsCoach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotificationsCoachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailReminderServiceServer).SendNotificationsCoach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flagrant.EmailReminderService/SendNotificationsCoach",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailReminderServiceServer).SendNotificationsCoach(ctx, req.(*SendNotificationsCoachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailReminderService_ServiceDesc is the grpc.ServiceDesc for EmailReminderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailReminderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flagrant.EmailReminderService",
	HandlerType: (*EmailReminderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendNotificationsGroups",
			Handler:    _EmailReminderService_SendNotificationsGroups_Handler,
		},
		{
			MethodName: "SendNotificationsGroup",
			Handler:    _EmailReminderService_SendNotificationsGroup_Handler,
		},
		{
			MethodName: "SendNotificationsCoaches",
			Handler:    _EmailReminderService_SendNotificationsCoaches_Handler,
		},
		{
			MethodName: "SendNotificationsCoach",
			Handler:    _EmailReminderService_SendNotificationsCoach_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flagrant/flagrant.proto",
}
