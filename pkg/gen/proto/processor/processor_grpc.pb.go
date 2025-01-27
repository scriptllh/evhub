// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package processor

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EvhubProcessorClient is the client API for EvhubProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EvhubProcessorClient interface {
	Dispatch(ctx context.Context, in *DispatchReq, opts ...grpc.CallOption) (*DispatchRsp, error)
}

type evhubProcessorClient struct {
	cc grpc.ClientConnInterface
}

func NewEvhubProcessorClient(cc grpc.ClientConnInterface) EvhubProcessorClient {
	return &evhubProcessorClient{cc}
}

func (c *evhubProcessorClient) Dispatch(ctx context.Context, in *DispatchReq, opts ...grpc.CallOption) (*DispatchRsp, error) {
	out := new(DispatchRsp)
	err := c.cc.Invoke(ctx, "/evhub_processor.evhubProcessor/Dispatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvhubProcessorServer is the server API for EvhubProcessor service.
// All implementations should embed UnimplementedEvhubProcessorServer
// for forward compatibility
type EvhubProcessorServer interface {
	Dispatch(context.Context, *DispatchReq) (*DispatchRsp, error)
}

// UnimplementedEvhubProcessorServer should be embedded to have forward compatible implementations.
type UnimplementedEvhubProcessorServer struct {
}

func (UnimplementedEvhubProcessorServer) Dispatch(context.Context, *DispatchReq) (*DispatchRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dispatch not implemented")
}

// UnsafeEvhubProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EvhubProcessorServer will
// result in compilation errors.
type UnsafeEvhubProcessorServer interface {
	mustEmbedUnimplementedEvhubProcessorServer()
}

func RegisterEvhubProcessorServer(s *grpc.Server, srv EvhubProcessorServer) {
	s.RegisterService(&_EvhubProcessor_serviceDesc, srv)
}

func _EvhubProcessor_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvhubProcessorServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evhub_processor.evhubProcessor/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvhubProcessorServer).Dispatch(ctx, req.(*DispatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _EvhubProcessor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evhub_processor.evhubProcessor",
	HandlerType: (*EvhubProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dispatch",
			Handler:    _EvhubProcessor_Dispatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "processor/processor.proto",
}
