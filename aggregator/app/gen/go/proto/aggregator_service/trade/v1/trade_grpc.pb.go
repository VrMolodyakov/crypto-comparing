// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/aggregator_service/trade/v1/trade.proto

package pb_trade

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

// TradeServiceClient is the client API for TradeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradeServiceClient interface {
	Create(ctx context.Context, opts ...grpc.CallOption) (TradeService_CreateClient, error)
}

type tradeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTradeServiceClient(cc grpc.ClientConnInterface) TradeServiceClient {
	return &tradeServiceClient{cc}
}

func (c *tradeServiceClient) Create(ctx context.Context, opts ...grpc.CallOption) (TradeService_CreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &TradeService_ServiceDesc.Streams[0], "/proto.aggregator_service.trade.v1.TradeService/Create", opts...)
	if err != nil {
		return nil, err
	}
	x := &tradeServiceCreateClient{stream}
	return x, nil
}

type TradeService_CreateClient interface {
	Send(*CreateRequest) error
	CloseAndRecv() (*CreateResponse, error)
	grpc.ClientStream
}

type tradeServiceCreateClient struct {
	grpc.ClientStream
}

func (x *tradeServiceCreateClient) Send(m *CreateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tradeServiceCreateClient) CloseAndRecv() (*CreateResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TradeServiceServer is the server API for TradeService service.
// All implementations must embed UnimplementedTradeServiceServer
// for forward compatibility
type TradeServiceServer interface {
	Create(TradeService_CreateServer) error
	mustEmbedUnimplementedTradeServiceServer()
}

// UnimplementedTradeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTradeServiceServer struct {
}

func (UnimplementedTradeServiceServer) Create(TradeService_CreateServer) error {
	return status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTradeServiceServer) mustEmbedUnimplementedTradeServiceServer() {}

// UnsafeTradeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TradeServiceServer will
// result in compilation errors.
type UnsafeTradeServiceServer interface {
	mustEmbedUnimplementedTradeServiceServer()
}

func RegisterTradeServiceServer(s grpc.ServiceRegistrar, srv TradeServiceServer) {
	s.RegisterService(&TradeService_ServiceDesc, srv)
}

func _TradeService_Create_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TradeServiceServer).Create(&tradeServiceCreateServer{stream})
}

type TradeService_CreateServer interface {
	SendAndClose(*CreateResponse) error
	Recv() (*CreateRequest, error)
	grpc.ServerStream
}

type tradeServiceCreateServer struct {
	grpc.ServerStream
}

func (x *tradeServiceCreateServer) SendAndClose(m *CreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tradeServiceCreateServer) Recv() (*CreateRequest, error) {
	m := new(CreateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TradeService_ServiceDesc is the grpc.ServiceDesc for TradeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TradeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.aggregator_service.trade.v1.TradeService",
	HandlerType: (*TradeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Create",
			Handler:       _TradeService_Create_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/aggregator_service/trade/v1/trade.proto",
}
