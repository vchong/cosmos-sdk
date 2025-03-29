// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cosmos/protocolpool/v1/tx.proto

package protocolpoolv1

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
	Msg_FundCommunityPool_FullMethodName    = "/cosmos.protocolpool.v1.Msg/FundCommunityPool"
	Msg_CommunityPoolSpend_FullMethodName   = "/cosmos.protocolpool.v1.Msg/CommunityPoolSpend"
	Msg_CreateContinuousFund_FullMethodName = "/cosmos.protocolpool.v1.Msg/CreateContinuousFund"
	Msg_CancelContinuousFund_FullMethodName = "/cosmos.protocolpool.v1.Msg/CancelContinuousFund"
	Msg_UpdateParams_FullMethodName         = "/cosmos.protocolpool.v1.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Msg defines the pool Msg service.
type MsgClient interface {
	// FundCommunityPool defines a method to allow an account to directly
	// fund the community pool.
	FundCommunityPool(ctx context.Context, in *MsgFundCommunityPool, opts ...grpc.CallOption) (*MsgFundCommunityPoolResponse, error)
	// CommunityPoolSpend defines a governance operation for sending tokens from
	// the community pool in the x/protocolpool module to another account, which
	// could be the governance module itself. The authority is defined in the
	// keeper.
	CommunityPoolSpend(ctx context.Context, in *MsgCommunityPoolSpend, opts ...grpc.CallOption) (*MsgCommunityPoolSpendResponse, error)
	// CreateContinuousFund defines a method to distribute a percentage of funds to an address continuously.
	// This ContinuousFund can be indefinite or run until a given expiry time.
	// Funds come from validator block rewards from x/distribution, but may also come from
	// any user who funds the ProtocolPoolEscrow module account directly through x/bank.
	CreateContinuousFund(ctx context.Context, in *MsgCreateContinuousFund, opts ...grpc.CallOption) (*MsgCreateContinuousFundResponse, error)
	// CancelContinuousFund defines a method for cancelling continuous fund.
	CancelContinuousFund(ctx context.Context, in *MsgCancelContinuousFund, opts ...grpc.CallOption) (*MsgCancelContinuousFundResponse, error)
	// UpdateParams defines a governance operation for updating the x/protocolpool module parameters.
	// The authority is defined in the keeper.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) FundCommunityPool(ctx context.Context, in *MsgFundCommunityPool, opts ...grpc.CallOption) (*MsgFundCommunityPoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgFundCommunityPoolResponse)
	err := c.cc.Invoke(ctx, Msg_FundCommunityPool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CommunityPoolSpend(ctx context.Context, in *MsgCommunityPoolSpend, opts ...grpc.CallOption) (*MsgCommunityPoolSpendResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgCommunityPoolSpendResponse)
	err := c.cc.Invoke(ctx, Msg_CommunityPoolSpend_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateContinuousFund(ctx context.Context, in *MsgCreateContinuousFund, opts ...grpc.CallOption) (*MsgCreateContinuousFundResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgCreateContinuousFundResponse)
	err := c.cc.Invoke(ctx, Msg_CreateContinuousFund_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelContinuousFund(ctx context.Context, in *MsgCancelContinuousFund, opts ...grpc.CallOption) (*MsgCancelContinuousFundResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgCancelContinuousFundResponse)
	err := c.cc.Invoke(ctx, Msg_CancelContinuousFund_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility.
//
// Msg defines the pool Msg service.
type MsgServer interface {
	// FundCommunityPool defines a method to allow an account to directly
	// fund the community pool.
	FundCommunityPool(context.Context, *MsgFundCommunityPool) (*MsgFundCommunityPoolResponse, error)
	// CommunityPoolSpend defines a governance operation for sending tokens from
	// the community pool in the x/protocolpool module to another account, which
	// could be the governance module itself. The authority is defined in the
	// keeper.
	CommunityPoolSpend(context.Context, *MsgCommunityPoolSpend) (*MsgCommunityPoolSpendResponse, error)
	// CreateContinuousFund defines a method to distribute a percentage of funds to an address continuously.
	// This ContinuousFund can be indefinite or run until a given expiry time.
	// Funds come from validator block rewards from x/distribution, but may also come from
	// any user who funds the ProtocolPoolEscrow module account directly through x/bank.
	CreateContinuousFund(context.Context, *MsgCreateContinuousFund) (*MsgCreateContinuousFundResponse, error)
	// CancelContinuousFund defines a method for cancelling continuous fund.
	CancelContinuousFund(context.Context, *MsgCancelContinuousFund) (*MsgCancelContinuousFundResponse, error)
	// UpdateParams defines a governance operation for updating the x/protocolpool module parameters.
	// The authority is defined in the keeper.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMsgServer struct{}

func (UnimplementedMsgServer) FundCommunityPool(context.Context, *MsgFundCommunityPool) (*MsgFundCommunityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundCommunityPool not implemented")
}
func (UnimplementedMsgServer) CommunityPoolSpend(context.Context, *MsgCommunityPoolSpend) (*MsgCommunityPoolSpendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommunityPoolSpend not implemented")
}
func (UnimplementedMsgServer) CreateContinuousFund(context.Context, *MsgCreateContinuousFund) (*MsgCreateContinuousFundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContinuousFund not implemented")
}
func (UnimplementedMsgServer) CancelContinuousFund(context.Context, *MsgCancelContinuousFund) (*MsgCancelContinuousFundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelContinuousFund not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}
func (UnimplementedMsgServer) testEmbeddedByValue()             {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	// If the following call pancis, it indicates UnimplementedMsgServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_FundCommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFundCommunityPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FundCommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_FundCommunityPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FundCommunityPool(ctx, req.(*MsgFundCommunityPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CommunityPoolSpend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCommunityPoolSpend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CommunityPoolSpend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CommunityPoolSpend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CommunityPoolSpend(ctx, req.(*MsgCommunityPoolSpend))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateContinuousFund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateContinuousFund)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateContinuousFund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateContinuousFund_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateContinuousFund(ctx, req.(*MsgCreateContinuousFund))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelContinuousFund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelContinuousFund)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelContinuousFund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelContinuousFund_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelContinuousFund(ctx, req.(*MsgCancelContinuousFund))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.protocolpool.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FundCommunityPool",
			Handler:    _Msg_FundCommunityPool_Handler,
		},
		{
			MethodName: "CommunityPoolSpend",
			Handler:    _Msg_CommunityPoolSpend_Handler,
		},
		{
			MethodName: "CreateContinuousFund",
			Handler:    _Msg_CreateContinuousFund_Handler,
		},
		{
			MethodName: "CancelContinuousFund",
			Handler:    _Msg_CancelContinuousFund_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/protocolpool/v1/tx.proto",
}
