// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: spark_ssp_internal.proto

package spark_ssp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SparkSspInternalService_QueryLostNodes_FullMethodName = "/spark_ssp.SparkSspInternalService/query_lost_nodes"
)

// SparkSspInternalServiceClient is the client API for SparkSspInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SparkSspInternalServiceClient interface {
	QueryLostNodes(ctx context.Context, in *QueryLostNodesRequest, opts ...grpc.CallOption) (*QueryLostNodesResponse, error)
}

type sparkSspInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSparkSspInternalServiceClient(cc grpc.ClientConnInterface) SparkSspInternalServiceClient {
	return &sparkSspInternalServiceClient{cc}
}

func (c *sparkSspInternalServiceClient) QueryLostNodes(ctx context.Context, in *QueryLostNodesRequest, opts ...grpc.CallOption) (*QueryLostNodesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryLostNodesResponse)
	err := c.cc.Invoke(ctx, SparkSspInternalService_QueryLostNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SparkSspInternalServiceServer is the server API for SparkSspInternalService service.
// All implementations must embed UnimplementedSparkSspInternalServiceServer
// for forward compatibility.
type SparkSspInternalServiceServer interface {
	QueryLostNodes(context.Context, *QueryLostNodesRequest) (*QueryLostNodesResponse, error)
	mustEmbedUnimplementedSparkSspInternalServiceServer()
}

// UnimplementedSparkSspInternalServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSparkSspInternalServiceServer struct{}

func (UnimplementedSparkSspInternalServiceServer) QueryLostNodes(context.Context, *QueryLostNodesRequest) (*QueryLostNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryLostNodes not implemented")
}
func (UnimplementedSparkSspInternalServiceServer) mustEmbedUnimplementedSparkSspInternalServiceServer() {
}
func (UnimplementedSparkSspInternalServiceServer) testEmbeddedByValue() {}

// UnsafeSparkSspInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SparkSspInternalServiceServer will
// result in compilation errors.
type UnsafeSparkSspInternalServiceServer interface {
	mustEmbedUnimplementedSparkSspInternalServiceServer()
}

func RegisterSparkSspInternalServiceServer(s grpc.ServiceRegistrar, srv SparkSspInternalServiceServer) {
	// If the following call pancis, it indicates UnimplementedSparkSspInternalServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SparkSspInternalService_ServiceDesc, srv)
}

func _SparkSspInternalService_QueryLostNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLostNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparkSspInternalServiceServer).QueryLostNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SparkSspInternalService_QueryLostNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparkSspInternalServiceServer).QueryLostNodes(ctx, req.(*QueryLostNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SparkSspInternalService_ServiceDesc is the grpc.ServiceDesc for SparkSspInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SparkSspInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spark_ssp.SparkSspInternalService",
	HandlerType: (*SparkSspInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "query_lost_nodes",
			Handler:    _SparkSspInternalService_QueryLostNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spark_ssp_internal.proto",
}
