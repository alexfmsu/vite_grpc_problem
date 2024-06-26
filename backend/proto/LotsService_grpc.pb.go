// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: proto/LotsService.proto

package lotspb

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

const (
	LotsService_ActiveLots_FullMethodName = "/lots.LotsService/ActiveLots"
)

// LotsServiceClient is the client API for LotsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LotsServiceClient interface {
	ActiveLots(ctx context.Context, in *LotsRequest, opts ...grpc.CallOption) (LotsService_ActiveLotsClient, error)
}

type lotsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLotsServiceClient(cc grpc.ClientConnInterface) LotsServiceClient {
	return &lotsServiceClient{cc}
}

func (c *lotsServiceClient) ActiveLots(ctx context.Context, in *LotsRequest, opts ...grpc.CallOption) (LotsService_ActiveLotsClient, error) {
	stream, err := c.cc.NewStream(ctx, &LotsService_ServiceDesc.Streams[0], LotsService_ActiveLots_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &lotsServiceActiveLotsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LotsService_ActiveLotsClient interface {
	Recv() (*LotsResponse, error)
	grpc.ClientStream
}

type lotsServiceActiveLotsClient struct {
	grpc.ClientStream
}

func (x *lotsServiceActiveLotsClient) Recv() (*LotsResponse, error) {
	m := new(LotsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LotsServiceServer is the server API for LotsService service.
// All implementations should embed UnimplementedLotsServiceServer
// for forward compatibility
type LotsServiceServer interface {
	ActiveLots(*LotsRequest, LotsService_ActiveLotsServer) error
}

// UnimplementedLotsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLotsServiceServer struct {
}

func (UnimplementedLotsServiceServer) ActiveLots(*LotsRequest, LotsService_ActiveLotsServer) error {
	return status.Errorf(codes.Unimplemented, "method ActiveLots not implemented")
}

// UnsafeLotsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LotsServiceServer will
// result in compilation errors.
type UnsafeLotsServiceServer interface {
	mustEmbedUnimplementedLotsServiceServer()
}

func RegisterLotsServiceServer(s grpc.ServiceRegistrar, srv LotsServiceServer) {
	s.RegisterService(&LotsService_ServiceDesc, srv)
}

func _LotsService_ActiveLots_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LotsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LotsServiceServer).ActiveLots(m, &lotsServiceActiveLotsServer{stream})
}

type LotsService_ActiveLotsServer interface {
	Send(*LotsResponse) error
	grpc.ServerStream
}

type lotsServiceActiveLotsServer struct {
	grpc.ServerStream
}

func (x *lotsServiceActiveLotsServer) Send(m *LotsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// LotsService_ServiceDesc is the grpc.ServiceDesc for LotsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LotsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lots.LotsService",
	HandlerType: (*LotsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ActiveLots",
			Handler:       _LotsService_ActiveLots_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/LotsService.proto",
}
