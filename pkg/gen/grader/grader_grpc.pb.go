// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api/proto/grader.proto

package grader

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

// GraderServiceClient is the client API for GraderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GraderServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (GraderService_UploadClient, error)
}

type graderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGraderServiceClient(cc grpc.ClientConnInterface) GraderServiceClient {
	return &graderServiceClient{cc}
}

func (c *graderServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (GraderService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &GraderService_ServiceDesc.Streams[0], "/GraderService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &graderServiceUploadClient{stream}
	return x, nil
}

type GraderService_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type graderServiceUploadClient struct {
	grpc.ClientStream
}

func (x *graderServiceUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *graderServiceUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GraderServiceServer is the server API for GraderService service.
// All implementations must embed UnimplementedGraderServiceServer
// for forward compatibility
type GraderServiceServer interface {
	Upload(GraderService_UploadServer) error
	mustEmbedUnimplementedGraderServiceServer()
}

// UnimplementedGraderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGraderServiceServer struct {
}

func (UnimplementedGraderServiceServer) Upload(GraderService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedGraderServiceServer) mustEmbedUnimplementedGraderServiceServer() {}

// UnsafeGraderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GraderServiceServer will
// result in compilation errors.
type UnsafeGraderServiceServer interface {
	mustEmbedUnimplementedGraderServiceServer()
}

func RegisterGraderServiceServer(s grpc.ServiceRegistrar, srv GraderServiceServer) {
	s.RegisterService(&GraderService_ServiceDesc, srv)
}

func _GraderService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GraderServiceServer).Upload(&graderServiceUploadServer{stream})
}

type GraderService_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type graderServiceUploadServer struct {
	grpc.ServerStream
}

func (x *graderServiceUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *graderServiceUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GraderService_ServiceDesc is the grpc.ServiceDesc for GraderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GraderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GraderService",
	HandlerType: (*GraderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _GraderService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/proto/grader.proto",
}
