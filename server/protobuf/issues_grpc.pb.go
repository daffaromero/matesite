// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: issues.proto

package protobuf

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

// IssuesServiceClient is the client API for IssuesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IssuesServiceClient interface {
	CreateIssue(ctx context.Context, in *CreateIssueRequest, opts ...grpc.CallOption) (*CreateIssueResponse, error)
	GetIssue(ctx context.Context, in *GetIssueRequest, opts ...grpc.CallOption) (*GetIssueResponse, error)
	ListIssues(ctx context.Context, in *ListIssuesRequest, opts ...grpc.CallOption) (*ListIssuesResponse, error)
	UpdateIssue(ctx context.Context, in *UpdateIssueRequest, opts ...grpc.CallOption) (*UpdateIssueResponse, error)
	DeleteIssue(ctx context.Context, in *DeleteIssueRequest, opts ...grpc.CallOption) (*DeleteIssueResponse, error)
}

type issuesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIssuesServiceClient(cc grpc.ClientConnInterface) IssuesServiceClient {
	return &issuesServiceClient{cc}
}

func (c *issuesServiceClient) CreateIssue(ctx context.Context, in *CreateIssueRequest, opts ...grpc.CallOption) (*CreateIssueResponse, error) {
	out := new(CreateIssueResponse)
	err := c.cc.Invoke(ctx, "/IssuesService/CreateIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issuesServiceClient) GetIssue(ctx context.Context, in *GetIssueRequest, opts ...grpc.CallOption) (*GetIssueResponse, error) {
	out := new(GetIssueResponse)
	err := c.cc.Invoke(ctx, "/IssuesService/GetIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issuesServiceClient) ListIssues(ctx context.Context, in *ListIssuesRequest, opts ...grpc.CallOption) (*ListIssuesResponse, error) {
	out := new(ListIssuesResponse)
	err := c.cc.Invoke(ctx, "/IssuesService/ListIssues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issuesServiceClient) UpdateIssue(ctx context.Context, in *UpdateIssueRequest, opts ...grpc.CallOption) (*UpdateIssueResponse, error) {
	out := new(UpdateIssueResponse)
	err := c.cc.Invoke(ctx, "/IssuesService/UpdateIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issuesServiceClient) DeleteIssue(ctx context.Context, in *DeleteIssueRequest, opts ...grpc.CallOption) (*DeleteIssueResponse, error) {
	out := new(DeleteIssueResponse)
	err := c.cc.Invoke(ctx, "/IssuesService/DeleteIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IssuesServiceServer is the server API for IssuesService service.
// All implementations must embed UnimplementedIssuesServiceServer
// for forward compatibility
type IssuesServiceServer interface {
	CreateIssue(context.Context, *CreateIssueRequest) (*CreateIssueResponse, error)
	GetIssue(context.Context, *GetIssueRequest) (*GetIssueResponse, error)
	ListIssues(context.Context, *ListIssuesRequest) (*ListIssuesResponse, error)
	UpdateIssue(context.Context, *UpdateIssueRequest) (*UpdateIssueResponse, error)
	DeleteIssue(context.Context, *DeleteIssueRequest) (*DeleteIssueResponse, error)
	mustEmbedUnimplementedIssuesServiceServer()
}

// UnimplementedIssuesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIssuesServiceServer struct {
}

func (UnimplementedIssuesServiceServer) CreateIssue(context.Context, *CreateIssueRequest) (*CreateIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIssue not implemented")
}
func (UnimplementedIssuesServiceServer) GetIssue(context.Context, *GetIssueRequest) (*GetIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIssue not implemented")
}
func (UnimplementedIssuesServiceServer) ListIssues(context.Context, *ListIssuesRequest) (*ListIssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIssues not implemented")
}
func (UnimplementedIssuesServiceServer) UpdateIssue(context.Context, *UpdateIssueRequest) (*UpdateIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIssue not implemented")
}
func (UnimplementedIssuesServiceServer) DeleteIssue(context.Context, *DeleteIssueRequest) (*DeleteIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIssue not implemented")
}
func (UnimplementedIssuesServiceServer) mustEmbedUnimplementedIssuesServiceServer() {}

// UnsafeIssuesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IssuesServiceServer will
// result in compilation errors.
type UnsafeIssuesServiceServer interface {
	mustEmbedUnimplementedIssuesServiceServer()
}

func RegisterIssuesServiceServer(s grpc.ServiceRegistrar, srv IssuesServiceServer) {
	s.RegisterService(&IssuesService_ServiceDesc, srv)
}

func _IssuesService_CreateIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuesServiceServer).CreateIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IssuesService/CreateIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuesServiceServer).CreateIssue(ctx, req.(*CreateIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssuesService_GetIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuesServiceServer).GetIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IssuesService/GetIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuesServiceServer).GetIssue(ctx, req.(*GetIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssuesService_ListIssues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIssuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuesServiceServer).ListIssues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IssuesService/ListIssues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuesServiceServer).ListIssues(ctx, req.(*ListIssuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssuesService_UpdateIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuesServiceServer).UpdateIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IssuesService/UpdateIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuesServiceServer).UpdateIssue(ctx, req.(*UpdateIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssuesService_DeleteIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuesServiceServer).DeleteIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IssuesService/DeleteIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuesServiceServer).DeleteIssue(ctx, req.(*DeleteIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IssuesService_ServiceDesc is the grpc.ServiceDesc for IssuesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IssuesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "IssuesService",
	HandlerType: (*IssuesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIssue",
			Handler:    _IssuesService_CreateIssue_Handler,
		},
		{
			MethodName: "GetIssue",
			Handler:    _IssuesService_GetIssue_Handler,
		},
		{
			MethodName: "ListIssues",
			Handler:    _IssuesService_ListIssues_Handler,
		},
		{
			MethodName: "UpdateIssue",
			Handler:    _IssuesService_UpdateIssue_Handler,
		},
		{
			MethodName: "DeleteIssue",
			Handler:    _IssuesService_DeleteIssue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "issues.proto",
}
