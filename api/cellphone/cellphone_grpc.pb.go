// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cellphone

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

// CellphoneServiceClient is the client API for CellphoneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CellphoneServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type cellphoneServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCellphoneServiceClient(cc grpc.ClientConnInterface) CellphoneServiceClient {
	return &cellphoneServiceClient{cc}
}

func (c *cellphoneServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/cellphone.CellphoneService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CellphoneServiceServer is the server API for CellphoneService service.
// All implementations must embed UnimplementedCellphoneServiceServer
// for forward compatibility
type CellphoneServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedCellphoneServiceServer()
}

// UnimplementedCellphoneServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCellphoneServiceServer struct {
}

func (UnimplementedCellphoneServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCellphoneServiceServer) mustEmbedUnimplementedCellphoneServiceServer() {}

// UnsafeCellphoneServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CellphoneServiceServer will
// result in compilation errors.
type UnsafeCellphoneServiceServer interface {
	mustEmbedUnimplementedCellphoneServiceServer()
}

func RegisterCellphoneServiceServer(s grpc.ServiceRegistrar, srv CellphoneServiceServer) {
	s.RegisterService(&CellphoneService_ServiceDesc, srv)
}

func _CellphoneService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CellphoneServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cellphone.CellphoneService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CellphoneServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CellphoneService_ServiceDesc is the grpc.ServiceDesc for CellphoneService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CellphoneService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cellphone.CellphoneService",
	HandlerType: (*CellphoneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CellphoneService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cellphone.proto",
}
