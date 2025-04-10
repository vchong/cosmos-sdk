// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cosmos/upgrade/v1beta1/tx.proto

package upgradev1beta1

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
	Msg_SoftwareUpgrade_FullMethodName = "/cosmos.upgrade.v1beta1.Msg/SoftwareUpgrade"
	Msg_CancelUpgrade_FullMethodName   = "/cosmos.upgrade.v1beta1.Msg/CancelUpgrade"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Msg defines the upgrade Msg service.
type MsgClient interface {
	// SoftwareUpgrade is a governance operation for initiating a software upgrade.
	SoftwareUpgrade(ctx context.Context, in *MsgSoftwareUpgrade, opts ...grpc.CallOption) (*MsgSoftwareUpgradeResponse, error)
	// CancelUpgrade is a governance operation for cancelling a previously
	// approved software upgrade.
	CancelUpgrade(ctx context.Context, in *MsgCancelUpgrade, opts ...grpc.CallOption) (*MsgCancelUpgradeResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SoftwareUpgrade(ctx context.Context, in *MsgSoftwareUpgrade, opts ...grpc.CallOption) (*MsgSoftwareUpgradeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgSoftwareUpgradeResponse)
	err := c.cc.Invoke(ctx, Msg_SoftwareUpgrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelUpgrade(ctx context.Context, in *MsgCancelUpgrade, opts ...grpc.CallOption) (*MsgCancelUpgradeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgCancelUpgradeResponse)
	err := c.cc.Invoke(ctx, Msg_CancelUpgrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility.
//
// Msg defines the upgrade Msg service.
type MsgServer interface {
	// SoftwareUpgrade is a governance operation for initiating a software upgrade.
	SoftwareUpgrade(context.Context, *MsgSoftwareUpgrade) (*MsgSoftwareUpgradeResponse, error)
	// CancelUpgrade is a governance operation for cancelling a previously
	// approved software upgrade.
	CancelUpgrade(context.Context, *MsgCancelUpgrade) (*MsgCancelUpgradeResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMsgServer struct{}

func (UnimplementedMsgServer) SoftwareUpgrade(context.Context, *MsgSoftwareUpgrade) (*MsgSoftwareUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoftwareUpgrade not implemented")
}
func (UnimplementedMsgServer) CancelUpgrade(context.Context, *MsgCancelUpgrade) (*MsgCancelUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelUpgrade not implemented")
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

func _Msg_SoftwareUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSoftwareUpgrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SoftwareUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SoftwareUpgrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SoftwareUpgrade(ctx, req.(*MsgSoftwareUpgrade))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelUpgrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelUpgrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelUpgrade(ctx, req.(*MsgCancelUpgrade))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.upgrade.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SoftwareUpgrade",
			Handler:    _Msg_SoftwareUpgrade_Handler,
		},
		{
			MethodName: "CancelUpgrade",
			Handler:    _Msg_CancelUpgrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/upgrade/v1beta1/tx.proto",
}
