// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: test.proto

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

const (
	Goods_CreateGoods_FullMethodName = "/stream.Goods/CreateGoods"
	Goods_UploadGoods_FullMethodName = "/stream.Goods/UploadGoods"
	Goods_DeleteGoods_FullMethodName = "/stream.Goods/DeleteGoods"
	Goods_WhereGoods_FullMethodName  = "/stream.Goods/WhereGoods"
	Goods_UploadFile_FullMethodName  = "/stream.Goods/UploadFile"
)

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsClient interface {
	CreateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...grpc.CallOption) (*CreateGoodsResponse, error)
	UploadGoods(ctx context.Context, in *UploadGoodsRequest, opts ...grpc.CallOption) (*UploadGoodsResponse, error)
	DeleteGoods(ctx context.Context, in *DeleteGoodsRequest, opts ...grpc.CallOption) (*DeleteGoodsResponse, error)
	WhereGoods(ctx context.Context, in *WhereGoodsRequest, opts ...grpc.CallOption) (*WhereGoodsResponse, error)
	UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error)
}

type goodsClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsClient(cc grpc.ClientConnInterface) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) CreateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...grpc.CallOption) (*CreateGoodsResponse, error) {
	out := new(CreateGoodsResponse)
	err := c.cc.Invoke(ctx, Goods_CreateGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UploadGoods(ctx context.Context, in *UploadGoodsRequest, opts ...grpc.CallOption) (*UploadGoodsResponse, error) {
	out := new(UploadGoodsResponse)
	err := c.cc.Invoke(ctx, Goods_UploadGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteGoods(ctx context.Context, in *DeleteGoodsRequest, opts ...grpc.CallOption) (*DeleteGoodsResponse, error) {
	out := new(DeleteGoodsResponse)
	err := c.cc.Invoke(ctx, Goods_DeleteGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) WhereGoods(ctx context.Context, in *WhereGoodsRequest, opts ...grpc.CallOption) (*WhereGoodsResponse, error) {
	out := new(WhereGoodsResponse)
	err := c.cc.Invoke(ctx, Goods_WhereGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error) {
	out := new(UploadFileResponse)
	err := c.cc.Invoke(ctx, Goods_UploadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
// All implementations must embed UnimplementedGoodsServer
// for forward compatibility
type GoodsServer interface {
	CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsResponse, error)
	UploadGoods(context.Context, *UploadGoodsRequest) (*UploadGoodsResponse, error)
	DeleteGoods(context.Context, *DeleteGoodsRequest) (*DeleteGoodsResponse, error)
	WhereGoods(context.Context, *WhereGoodsRequest) (*WhereGoodsResponse, error)
	UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error)
	mustEmbedUnimplementedGoodsServer()
}

// UnimplementedGoodsServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServer struct {
}

func (UnimplementedGoodsServer) CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoods not implemented")
}
func (UnimplementedGoodsServer) UploadGoods(context.Context, *UploadGoodsRequest) (*UploadGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadGoods not implemented")
}
func (UnimplementedGoodsServer) DeleteGoods(context.Context, *DeleteGoodsRequest) (*DeleteGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoods not implemented")
}
func (UnimplementedGoodsServer) WhereGoods(context.Context, *WhereGoodsRequest) (*WhereGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhereGoods not implemented")
}
func (UnimplementedGoodsServer) UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedGoodsServer) mustEmbedUnimplementedGoodsServer() {}

// UnsafeGoodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServer will
// result in compilation errors.
type UnsafeGoodsServer interface {
	mustEmbedUnimplementedGoodsServer()
}

func RegisterGoodsServer(s grpc.ServiceRegistrar, srv GoodsServer) {
	s.RegisterService(&Goods_ServiceDesc, srv)
}

func _Goods_CreateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_CreateGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateGoods(ctx, req.(*CreateGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UploadGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UploadGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_UploadGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UploadGoods(ctx, req.(*UploadGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_DeleteGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteGoods(ctx, req.(*DeleteGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_WhereGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhereGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).WhereGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_WhereGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).WhereGoods(ctx, req.(*WhereGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_UploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UploadFile(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Goods_ServiceDesc is the grpc.ServiceDesc for Goods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Goods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGoods",
			Handler:    _Goods_CreateGoods_Handler,
		},
		{
			MethodName: "UploadGoods",
			Handler:    _Goods_UploadGoods_Handler,
		},
		{
			MethodName: "DeleteGoods",
			Handler:    _Goods_DeleteGoods_Handler,
		},
		{
			MethodName: "WhereGoods",
			Handler:    _Goods_WhereGoods_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _Goods_UploadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
