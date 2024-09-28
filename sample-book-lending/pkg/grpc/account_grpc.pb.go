// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: account.proto

// packageの宣言

package grpc

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
	LendingBooksService_SendBorrow_FullMethodName      = "/myapp.LendingBooksService/SendBorrow"
	LendingBooksService_ShowAccountInfo_FullMethodName = "/myapp.LendingBooksService/ShowAccountInfo"
)

// LendingBooksServiceClient is the client API for LendingBooksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 本の貸し出しサービス
type LendingBooksServiceClient interface {
	// 本を借りるためのメソッド
	SendBorrow(ctx context.Context, in *BorrowRequest, opts ...grpc.CallOption) (*BorrrowResponse, error)
	// 名前を指定すると借りてる本一覧が表示される。
	ShowAccountInfo(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountInfo, error)
}

type lendingBooksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLendingBooksServiceClient(cc grpc.ClientConnInterface) LendingBooksServiceClient {
	return &lendingBooksServiceClient{cc}
}

func (c *lendingBooksServiceClient) SendBorrow(ctx context.Context, in *BorrowRequest, opts ...grpc.CallOption) (*BorrrowResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BorrrowResponse)
	err := c.cc.Invoke(ctx, LendingBooksService_SendBorrow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lendingBooksServiceClient) ShowAccountInfo(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountInfo)
	err := c.cc.Invoke(ctx, LendingBooksService_ShowAccountInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LendingBooksServiceServer is the server API for LendingBooksService service.
// All implementations must embed UnimplementedLendingBooksServiceServer
// for forward compatibility.
//
// 本の貸し出しサービス
type LendingBooksServiceServer interface {
	// 本を借りるためのメソッド
	SendBorrow(context.Context, *BorrowRequest) (*BorrrowResponse, error)
	// 名前を指定すると借りてる本一覧が表示される。
	ShowAccountInfo(context.Context, *Account) (*AccountInfo, error)
	mustEmbedUnimplementedLendingBooksServiceServer()
}

// UnimplementedLendingBooksServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLendingBooksServiceServer struct{}

func (UnimplementedLendingBooksServiceServer) SendBorrow(context.Context, *BorrowRequest) (*BorrrowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBorrow not implemented")
}
func (UnimplementedLendingBooksServiceServer) ShowAccountInfo(context.Context, *Account) (*AccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAccountInfo not implemented")
}
func (UnimplementedLendingBooksServiceServer) mustEmbedUnimplementedLendingBooksServiceServer() {}
func (UnimplementedLendingBooksServiceServer) testEmbeddedByValue()                             {}

// UnsafeLendingBooksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LendingBooksServiceServer will
// result in compilation errors.
type UnsafeLendingBooksServiceServer interface {
	mustEmbedUnimplementedLendingBooksServiceServer()
}

func RegisterLendingBooksServiceServer(s grpc.ServiceRegistrar, srv LendingBooksServiceServer) {
	// If the following call pancis, it indicates UnimplementedLendingBooksServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LendingBooksService_ServiceDesc, srv)
}

func _LendingBooksService_SendBorrow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BorrowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LendingBooksServiceServer).SendBorrow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LendingBooksService_SendBorrow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LendingBooksServiceServer).SendBorrow(ctx, req.(*BorrowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LendingBooksService_ShowAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LendingBooksServiceServer).ShowAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LendingBooksService_ShowAccountInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LendingBooksServiceServer).ShowAccountInfo(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

// LendingBooksService_ServiceDesc is the grpc.ServiceDesc for LendingBooksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LendingBooksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myapp.LendingBooksService",
	HandlerType: (*LendingBooksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendBorrow",
			Handler:    _LendingBooksService_SendBorrow_Handler,
		},
		{
			MethodName: "ShowAccountInfo",
			Handler:    _LendingBooksService_ShowAccountInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}
