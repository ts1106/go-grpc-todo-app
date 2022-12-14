// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: todo/todo.proto

package todo

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

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoClient interface {
	GetItem(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Item, error)
	ListItem(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Todo_ListItemClient, error)
	CreateItem(ctx context.Context, opts ...grpc.CallOption) (Todo_CreateItemClient, error)
	ProgressChat(ctx context.Context, opts ...grpc.CallOption) (Todo_ProgressChatClient, error)
}

type todoClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoClient(cc grpc.ClientConnInterface) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) GetItem(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/todo.Todo/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) ListItem(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Todo_ListItemClient, error) {
	stream, err := c.cc.NewStream(ctx, &Todo_ServiceDesc.Streams[0], "/todo.Todo/ListItem", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoListItemClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Todo_ListItemClient interface {
	Recv() (*Item, error)
	grpc.ClientStream
}

type todoListItemClient struct {
	grpc.ClientStream
}

func (x *todoListItemClient) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoClient) CreateItem(ctx context.Context, opts ...grpc.CallOption) (Todo_CreateItemClient, error) {
	stream, err := c.cc.NewStream(ctx, &Todo_ServiceDesc.Streams[1], "/todo.Todo/CreateItem", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoCreateItemClient{stream}
	return x, nil
}

type Todo_CreateItemClient interface {
	Send(*Item) error
	CloseAndRecv() (*Summary, error)
	grpc.ClientStream
}

type todoCreateItemClient struct {
	grpc.ClientStream
}

func (x *todoCreateItemClient) Send(m *Item) error {
	return x.ClientStream.SendMsg(m)
}

func (x *todoCreateItemClient) CloseAndRecv() (*Summary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Summary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoClient) ProgressChat(ctx context.Context, opts ...grpc.CallOption) (Todo_ProgressChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &Todo_ServiceDesc.Streams[2], "/todo.Todo/ProgressChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoProgressChatClient{stream}
	return x, nil
}

type Todo_ProgressChatClient interface {
	Send(*ProgressNote) error
	Recv() (*ProgressNote, error)
	grpc.ClientStream
}

type todoProgressChatClient struct {
	grpc.ClientStream
}

func (x *todoProgressChatClient) Send(m *ProgressNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *todoProgressChatClient) Recv() (*ProgressNote, error) {
	m := new(ProgressNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoServer is the server API for Todo service.
// All implementations must embed UnimplementedTodoServer
// for forward compatibility
type TodoServer interface {
	GetItem(context.Context, *GetRequest) (*Item, error)
	ListItem(*ListRequest, Todo_ListItemServer) error
	CreateItem(Todo_CreateItemServer) error
	ProgressChat(Todo_ProgressChatServer) error
	mustEmbedUnimplementedTodoServer()
}

// UnimplementedTodoServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServer struct {
}

func (UnimplementedTodoServer) GetItem(context.Context, *GetRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedTodoServer) ListItem(*ListRequest, Todo_ListItemServer) error {
	return status.Errorf(codes.Unimplemented, "method ListItem not implemented")
}
func (UnimplementedTodoServer) CreateItem(Todo_CreateItemServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedTodoServer) ProgressChat(Todo_ProgressChatServer) error {
	return status.Errorf(codes.Unimplemented, "method ProgressChat not implemented")
}
func (UnimplementedTodoServer) mustEmbedUnimplementedTodoServer() {}

// UnsafeTodoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServer will
// result in compilation errors.
type UnsafeTodoServer interface {
	mustEmbedUnimplementedTodoServer()
}

func RegisterTodoServer(s grpc.ServiceRegistrar, srv TodoServer) {
	s.RegisterService(&Todo_ServiceDesc, srv)
}

func _Todo_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetItem(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_ListItem_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServer).ListItem(m, &todoListItemServer{stream})
}

type Todo_ListItemServer interface {
	Send(*Item) error
	grpc.ServerStream
}

type todoListItemServer struct {
	grpc.ServerStream
}

func (x *todoListItemServer) Send(m *Item) error {
	return x.ServerStream.SendMsg(m)
}

func _Todo_CreateItem_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TodoServer).CreateItem(&todoCreateItemServer{stream})
}

type Todo_CreateItemServer interface {
	SendAndClose(*Summary) error
	Recv() (*Item, error)
	grpc.ServerStream
}

type todoCreateItemServer struct {
	grpc.ServerStream
}

func (x *todoCreateItemServer) SendAndClose(m *Summary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *todoCreateItemServer) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Todo_ProgressChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TodoServer).ProgressChat(&todoProgressChatServer{stream})
}

type Todo_ProgressChatServer interface {
	Send(*ProgressNote) error
	Recv() (*ProgressNote, error)
	grpc.ServerStream
}

type todoProgressChatServer struct {
	grpc.ServerStream
}

func (x *todoProgressChatServer) Send(m *ProgressNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *todoProgressChatServer) Recv() (*ProgressNote, error) {
	m := new(ProgressNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Todo_ServiceDesc is the grpc.ServiceDesc for Todo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItem",
			Handler:    _Todo_GetItem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListItem",
			Handler:       _Todo_ListItem_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateItem",
			Handler:       _Todo_CreateItem_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ProgressChat",
			Handler:       _Todo_ProgressChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "todo/todo.proto",
}
