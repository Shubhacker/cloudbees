// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: invoicer.proto

package cloudbees

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

// InvoicerClient is the client API for Invoicer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoicerClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	ReadPost(ctx context.Context, in *ReadPostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*PostResponse, error)
}

type invoicerClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoicerClient(cc grpc.ClientConnInterface) InvoicerClient {
	return &invoicerClient{cc}
}

func (c *invoicerClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/Invoicer/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoicerClient) ReadPost(ctx context.Context, in *ReadPostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/Invoicer/ReadPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoicerClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/Invoicer/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoicerServer is the server API for Invoicer service.
// All implementations must embed UnimplementedInvoicerServer
// for forward compatibility
type InvoicerServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*PostResponse, error)
	ReadPost(context.Context, *ReadPostRequest) (*PostResponse, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*PostResponse, error)
	mustEmbedUnimplementedInvoicerServer()
}

// UnimplementedInvoicerServer must be embedded to have forward compatible implementations.
type UnimplementedInvoicerServer struct {
}

func (UnimplementedInvoicerServer) CreatePost(context.Context, *CreatePostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedInvoicerServer) ReadPost(context.Context, *ReadPostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPost not implemented")
}
func (UnimplementedInvoicerServer) UpdatePost(context.Context, *UpdatePostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedInvoicerServer) mustEmbedUnimplementedInvoicerServer() {}

// UnsafeInvoicerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoicerServer will
// result in compilation errors.
type UnsafeInvoicerServer interface {
	mustEmbedUnimplementedInvoicerServer()
}

func RegisterInvoicerServer(s grpc.ServiceRegistrar, srv InvoicerServer) {
	s.RegisterService(&Invoicer_ServiceDesc, srv)
}

func _Invoicer_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicerServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Invoicer/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicerServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoicer_ReadPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicerServer).ReadPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Invoicer/ReadPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicerServer).ReadPost(ctx, req.(*ReadPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoicer_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicerServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Invoicer/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicerServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Invoicer_ServiceDesc is the grpc.ServiceDesc for Invoicer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Invoicer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Invoicer",
	HandlerType: (*InvoicerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _Invoicer_CreatePost_Handler,
		},
		{
			MethodName: "ReadPost",
			Handler:    _Invoicer_ReadPost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Invoicer_UpdatePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoicer.proto",
}