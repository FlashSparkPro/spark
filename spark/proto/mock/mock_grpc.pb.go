// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: mock.proto

package mock

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MockService_CleanUpPreimageShare_FullMethodName = "/mock.MockService/clean_up_preimage_share"
	MockService_InterruptTransfer_FullMethodName    = "/mock.MockService/interrupt_transfer"
	MockService_UpdateNodesStatus_FullMethodName    = "/mock.MockService/update_nodes_status"
)

// MockServiceClient is the client API for MockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MockServiceClient interface {
	CleanUpPreimageShare(ctx context.Context, in *CleanUpPreimageShareRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	InterruptTransfer(ctx context.Context, in *InterruptTransferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateNodesStatus(ctx context.Context, in *UpdateNodesStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMockServiceClient(cc grpc.ClientConnInterface) MockServiceClient {
	return &mockServiceClient{cc}
}

func (c *mockServiceClient) CleanUpPreimageShare(ctx context.Context, in *CleanUpPreimageShareRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MockService_CleanUpPreimageShare_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockServiceClient) InterruptTransfer(ctx context.Context, in *InterruptTransferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MockService_InterruptTransfer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockServiceClient) UpdateNodesStatus(ctx context.Context, in *UpdateNodesStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MockService_UpdateNodesStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MockServiceServer is the server API for MockService service.
// All implementations must embed UnimplementedMockServiceServer
// for forward compatibility.
type MockServiceServer interface {
	CleanUpPreimageShare(context.Context, *CleanUpPreimageShareRequest) (*emptypb.Empty, error)
	InterruptTransfer(context.Context, *InterruptTransferRequest) (*emptypb.Empty, error)
	UpdateNodesStatus(context.Context, *UpdateNodesStatusRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMockServiceServer()
}

// UnimplementedMockServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMockServiceServer struct{}

func (UnimplementedMockServiceServer) CleanUpPreimageShare(context.Context, *CleanUpPreimageShareRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanUpPreimageShare not implemented")
}
func (UnimplementedMockServiceServer) InterruptTransfer(context.Context, *InterruptTransferRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InterruptTransfer not implemented")
}
func (UnimplementedMockServiceServer) UpdateNodesStatus(context.Context, *UpdateNodesStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodesStatus not implemented")
}
func (UnimplementedMockServiceServer) mustEmbedUnimplementedMockServiceServer() {}
func (UnimplementedMockServiceServer) testEmbeddedByValue()                     {}

// UnsafeMockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MockServiceServer will
// result in compilation errors.
type UnsafeMockServiceServer interface {
	mustEmbedUnimplementedMockServiceServer()
}

func RegisterMockServiceServer(s grpc.ServiceRegistrar, srv MockServiceServer) {
	// If the following call pancis, it indicates UnimplementedMockServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MockService_ServiceDesc, srv)
}

func _MockService_CleanUpPreimageShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanUpPreimageShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockServiceServer).CleanUpPreimageShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MockService_CleanUpPreimageShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockServiceServer).CleanUpPreimageShare(ctx, req.(*CleanUpPreimageShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MockService_InterruptTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InterruptTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockServiceServer).InterruptTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MockService_InterruptTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockServiceServer).InterruptTransfer(ctx, req.(*InterruptTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MockService_UpdateNodesStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodesStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockServiceServer).UpdateNodesStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MockService_UpdateNodesStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockServiceServer).UpdateNodesStatus(ctx, req.(*UpdateNodesStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MockService_ServiceDesc is the grpc.ServiceDesc for MockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mock.MockService",
	HandlerType: (*MockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "clean_up_preimage_share",
			Handler:    _MockService_CleanUpPreimageShare_Handler,
		},
		{
			MethodName: "interrupt_transfer",
			Handler:    _MockService_InterruptTransfer_Handler,
		},
		{
			MethodName: "update_nodes_status",
			Handler:    _MockService_UpdateNodesStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mock.proto",
}
