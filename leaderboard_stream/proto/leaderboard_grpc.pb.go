// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: leaderboard.proto

package proto

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

// LeaderboardServiceClient is the client API for LeaderboardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeaderboardServiceClient interface {
	AddNewPlayer(ctx context.Context, opts ...grpc.CallOption) (LeaderboardService_AddNewPlayerClient, error)
	GetLeaderboard(ctx context.Context, in *EmptyParam, opts ...grpc.CallOption) (*LeaderboardResponse, error)
}

type leaderboardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeaderboardServiceClient(cc grpc.ClientConnInterface) LeaderboardServiceClient {
	return &leaderboardServiceClient{cc}
}

func (c *leaderboardServiceClient) AddNewPlayer(ctx context.Context, opts ...grpc.CallOption) (LeaderboardService_AddNewPlayerClient, error) {
	stream, err := c.cc.NewStream(ctx, &LeaderboardService_ServiceDesc.Streams[0], "/leaderboard.LeaderboardService/AddNewPlayer", opts...)
	if err != nil {
		return nil, err
	}
	x := &leaderboardServiceAddNewPlayerClient{stream}
	return x, nil
}

type LeaderboardService_AddNewPlayerClient interface {
	Send(*StoreNewLeaderboardRequest) error
	Recv() (*LeaderboardResponse, error)
	grpc.ClientStream
}

type leaderboardServiceAddNewPlayerClient struct {
	grpc.ClientStream
}

func (x *leaderboardServiceAddNewPlayerClient) Send(m *StoreNewLeaderboardRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *leaderboardServiceAddNewPlayerClient) Recv() (*LeaderboardResponse, error) {
	m := new(LeaderboardResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *leaderboardServiceClient) GetLeaderboard(ctx context.Context, in *EmptyParam, opts ...grpc.CallOption) (*LeaderboardResponse, error) {
	out := new(LeaderboardResponse)
	err := c.cc.Invoke(ctx, "/leaderboard.LeaderboardService/GetLeaderboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeaderboardServiceServer is the server API for LeaderboardService service.
// All implementations must embed UnimplementedLeaderboardServiceServer
// for forward compatibility
type LeaderboardServiceServer interface {
	AddNewPlayer(LeaderboardService_AddNewPlayerServer) error
	GetLeaderboard(context.Context, *EmptyParam) (*LeaderboardResponse, error)
	mustEmbedUnimplementedLeaderboardServiceServer()
}

// UnimplementedLeaderboardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLeaderboardServiceServer struct {
}

func (UnimplementedLeaderboardServiceServer) AddNewPlayer(LeaderboardService_AddNewPlayerServer) error {
	return status.Errorf(codes.Unimplemented, "method AddNewPlayer not implemented")
}
func (UnimplementedLeaderboardServiceServer) GetLeaderboard(context.Context, *EmptyParam) (*LeaderboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderboard not implemented")
}
func (UnimplementedLeaderboardServiceServer) mustEmbedUnimplementedLeaderboardServiceServer() {}

// UnsafeLeaderboardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeaderboardServiceServer will
// result in compilation errors.
type UnsafeLeaderboardServiceServer interface {
	mustEmbedUnimplementedLeaderboardServiceServer()
}

func RegisterLeaderboardServiceServer(s grpc.ServiceRegistrar, srv LeaderboardServiceServer) {
	s.RegisterService(&LeaderboardService_ServiceDesc, srv)
}

func _LeaderboardService_AddNewPlayer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LeaderboardServiceServer).AddNewPlayer(&leaderboardServiceAddNewPlayerServer{stream})
}

type LeaderboardService_AddNewPlayerServer interface {
	Send(*LeaderboardResponse) error
	Recv() (*StoreNewLeaderboardRequest, error)
	grpc.ServerStream
}

type leaderboardServiceAddNewPlayerServer struct {
	grpc.ServerStream
}

func (x *leaderboardServiceAddNewPlayerServer) Send(m *LeaderboardResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *leaderboardServiceAddNewPlayerServer) Recv() (*StoreNewLeaderboardRequest, error) {
	m := new(StoreNewLeaderboardRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LeaderboardService_GetLeaderboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServiceServer).GetLeaderboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leaderboard.LeaderboardService/GetLeaderboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServiceServer).GetLeaderboard(ctx, req.(*EmptyParam))
	}
	return interceptor(ctx, in, info, handler)
}

// LeaderboardService_ServiceDesc is the grpc.ServiceDesc for LeaderboardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LeaderboardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "leaderboard.LeaderboardService",
	HandlerType: (*LeaderboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLeaderboard",
			Handler:    _LeaderboardService_GetLeaderboard_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddNewPlayer",
			Handler:       _LeaderboardService_AddNewPlayer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "leaderboard.proto",
}