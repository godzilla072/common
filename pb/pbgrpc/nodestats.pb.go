// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nodestats.proto

package pbgrpc

import (
	context "context"

	_ "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"

	. "storj.io/common/pb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeStatsClient is the client API for NodeStats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeStatsClient interface {
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
	DailyStorageUsage(ctx context.Context, in *DailyStorageUsageRequest, opts ...grpc.CallOption) (*DailyStorageUsageResponse, error)
	PricingModel(ctx context.Context, in *PricingModelRequest, opts ...grpc.CallOption) (*PricingModelResponse, error)
}

type nodeStatsClient struct {
	cc *grpc.ClientConn
}

func NewNodeStatsClient(cc *grpc.ClientConn) NodeStatsClient {
	return &nodeStatsClient{cc}
}

func (c *nodeStatsClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, "/nodestats.NodeStats/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeStatsClient) DailyStorageUsage(ctx context.Context, in *DailyStorageUsageRequest, opts ...grpc.CallOption) (*DailyStorageUsageResponse, error) {
	out := new(DailyStorageUsageResponse)
	err := c.cc.Invoke(ctx, "/nodestats.NodeStats/DailyStorageUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeStatsClient) PricingModel(ctx context.Context, in *PricingModelRequest, opts ...grpc.CallOption) (*PricingModelResponse, error) {
	out := new(PricingModelResponse)
	err := c.cc.Invoke(ctx, "/nodestats.NodeStats/PricingModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeStatsServer is the server API for NodeStats service.
type NodeStatsServer interface {
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	DailyStorageUsage(context.Context, *DailyStorageUsageRequest) (*DailyStorageUsageResponse, error)
	PricingModel(context.Context, *PricingModelRequest) (*PricingModelResponse, error)
}

func RegisterNodeStatsServer(s *grpc.Server, srv NodeStatsServer) {
	s.RegisterService(&_NodeStats_serviceDesc, srv)
}

func _NodeStats_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeStatsServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodestats.NodeStats/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeStatsServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeStats_DailyStorageUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DailyStorageUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeStatsServer).DailyStorageUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodestats.NodeStats/DailyStorageUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeStatsServer).DailyStorageUsage(ctx, req.(*DailyStorageUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeStats_PricingModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PricingModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeStatsServer).PricingModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodestats.NodeStats/PricingModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeStatsServer).PricingModel(ctx, req.(*PricingModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeStats_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodestats.NodeStats",
	HandlerType: (*NodeStatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStats",
			Handler:    _NodeStats_GetStats_Handler,
		},
		{
			MethodName: "DailyStorageUsage",
			Handler:    _NodeStats_DailyStorageUsage_Handler,
		},
		{
			MethodName: "PricingModel",
			Handler:    _NodeStats_PricingModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodestats.proto",
}
