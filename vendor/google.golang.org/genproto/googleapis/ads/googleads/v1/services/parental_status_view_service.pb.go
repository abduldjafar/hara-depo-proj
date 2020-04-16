// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/parental_status_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [ParentalStatusViewService.GetParentalStatusView][google.ads.googleads.v1.services.ParentalStatusViewService.GetParentalStatusView].
type GetParentalStatusViewRequest struct {
	// Required. The resource name of the parental status view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParentalStatusViewRequest) Reset()         { *m = GetParentalStatusViewRequest{} }
func (m *GetParentalStatusViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetParentalStatusViewRequest) ProtoMessage()    {}
func (*GetParentalStatusViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3875092f69544412, []int{0}
}

func (m *GetParentalStatusViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParentalStatusViewRequest.Unmarshal(m, b)
}
func (m *GetParentalStatusViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParentalStatusViewRequest.Marshal(b, m, deterministic)
}
func (m *GetParentalStatusViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParentalStatusViewRequest.Merge(m, src)
}
func (m *GetParentalStatusViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetParentalStatusViewRequest.Size(m)
}
func (m *GetParentalStatusViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParentalStatusViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetParentalStatusViewRequest proto.InternalMessageInfo

func (m *GetParentalStatusViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetParentalStatusViewRequest)(nil), "google.ads.googleads.v1.services.GetParentalStatusViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/parental_status_view_service.proto", fileDescriptor_3875092f69544412)
}

var fileDescriptor_3875092f69544412 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x6b, 0xd4, 0x40,
	0x18, 0x65, 0x53, 0x10, 0x0c, 0x7a, 0x09, 0x88, 0x75, 0x2d, 0xba, 0x94, 0x1e, 0x4a, 0x0f, 0x33,
	0x44, 0x29, 0xc2, 0xf8, 0x03, 0x66, 0x3d, 0x6c, 0x2f, 0xca, 0xd2, 0xc2, 0x1e, 0x64, 0x21, 0x4c,
	0x93, 0xcf, 0x38, 0x90, 0xcc, 0xc4, 0xf9, 0x26, 0xe9, 0x41, 0xbc, 0x08, 0xfe, 0x05, 0x5e, 0x3c,
	0x7b, 0xf4, 0x4f, 0xe9, 0xd5, 0x9b, 0x20, 0x78, 0xf0, 0xe4, 0x5f, 0x21, 0xd9, 0xc9, 0x64, 0x5b,
	0x6a, 0xba, 0xb7, 0xc7, 0xbc, 0xf7, 0xbd, 0xf7, 0xfd, 0x98, 0xf0, 0x65, 0xae, 0x75, 0x5e, 0x00,
	0x15, 0x19, 0x52, 0x07, 0x5b, 0xd4, 0xc4, 0x14, 0xc1, 0x34, 0x32, 0x05, 0xa4, 0x95, 0x30, 0xa0,
	0xac, 0x28, 0x12, 0xb4, 0xc2, 0xd6, 0x98, 0x34, 0x12, 0xce, 0x92, 0x8e, 0x25, 0x95, 0xd1, 0x56,
	0x47, 0x13, 0x57, 0x49, 0x44, 0x86, 0xa4, 0x37, 0x21, 0x4d, 0x4c, 0xbc, 0xc9, 0xf8, 0xd9, 0x50,
	0x8c, 0x01, 0xd4, 0xb5, 0x19, 0xca, 0x71, 0xfe, 0xe3, 0x1d, 0x5f, 0x5d, 0x49, 0x2a, 0x94, 0xd2,
	0x56, 0x58, 0xa9, 0x15, 0x76, 0xec, 0xdd, 0x0b, 0x6c, 0x5a, 0x48, 0x50, 0xb6, 0x23, 0x1e, 0x5e,
	0x20, 0xde, 0x4a, 0x28, 0xb2, 0xe4, 0x14, 0xde, 0x89, 0x46, 0x6a, 0xe3, 0x04, 0xbb, 0x47, 0xe1,
	0xce, 0x0c, 0xec, 0xbc, 0x0b, 0x3e, 0x59, 0xe5, 0x2e, 0x24, 0x9c, 0x1d, 0xc3, 0xfb, 0x1a, 0xd0,
	0x46, 0xfb, 0xe1, 0x6d, 0xdf, 0x5f, 0xa2, 0x44, 0x09, 0xdb, 0xa3, 0xc9, 0x68, 0xff, 0xe6, 0x74,
	0xeb, 0x37, 0x0f, 0x8e, 0x6f, 0x79, 0xe6, 0xb5, 0x28, 0xe1, 0xd1, 0xd7, 0x20, 0xbc, 0x77, 0xd5,
	0xe7, 0xc4, 0x8d, 0x1f, 0xfd, 0x1a, 0x85, 0x77, 0xfe, 0x1b, 0x14, 0xbd, 0x20, 0x9b, 0x56, 0x47,
	0xae, 0xeb, 0x70, 0x7c, 0x38, 0x58, 0xdf, 0x2f, 0x96, 0x5c, 0xad, 0xde, 0x7d, 0xf5, 0x93, 0x5f,
	0x9e, 0xec, 0xd3, 0x8f, 0x3f, 0x5f, 0x82, 0x27, 0xd1, 0x61, 0x7b, 0x92, 0x0f, 0x97, 0x98, 0xe7,
	0x69, 0x8d, 0x56, 0x97, 0x60, 0x90, 0x1e, 0xf4, 0x37, 0x5a, 0x5b, 0x21, 0x3d, 0xf8, 0x38, 0xbe,
	0x7f, 0xce, 0xb7, 0xd7, 0xe1, 0x1d, 0xaa, 0x24, 0x92, 0x54, 0x97, 0xd3, 0xcf, 0x41, 0xb8, 0x97,
	0xea, 0x72, 0xe3, 0xa0, 0xd3, 0x07, 0x83, 0x0b, 0x9c, 0xb7, 0xd7, 0x9a, 0x8f, 0xde, 0x1c, 0x75,
	0x1e, 0xb9, 0x2e, 0x84, 0xca, 0x89, 0x36, 0x39, 0xcd, 0x41, 0xad, 0x6e, 0x49, 0xd7, 0xa9, 0xc3,
	0x7f, 0xf9, 0xa9, 0x07, 0xdf, 0x82, 0xad, 0x19, 0xe7, 0xdf, 0x83, 0xc9, 0xcc, 0x19, 0xf2, 0x0c,
	0x89, 0x83, 0x2d, 0x5a, 0xc4, 0xa4, 0x0b, 0xc6, 0x73, 0x2f, 0x59, 0xf2, 0x0c, 0x97, 0xbd, 0x64,
	0xb9, 0x88, 0x97, 0x5e, 0xf2, 0x37, 0xd8, 0x73, 0xef, 0x8c, 0xf1, 0x0c, 0x19, 0xeb, 0x45, 0x8c,
	0x2d, 0x62, 0xc6, 0xbc, 0xec, 0xf4, 0xc6, 0xaa, 0xcf, 0xc7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xef, 0x82, 0x31, 0x38, 0x72, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ParentalStatusViewServiceClient is the client API for ParentalStatusViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ParentalStatusViewServiceClient interface {
	// Returns the requested parental status view in full detail.
	GetParentalStatusView(ctx context.Context, in *GetParentalStatusViewRequest, opts ...grpc.CallOption) (*resources.ParentalStatusView, error)
}

type parentalStatusViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParentalStatusViewServiceClient(cc grpc.ClientConnInterface) ParentalStatusViewServiceClient {
	return &parentalStatusViewServiceClient{cc}
}

func (c *parentalStatusViewServiceClient) GetParentalStatusView(ctx context.Context, in *GetParentalStatusViewRequest, opts ...grpc.CallOption) (*resources.ParentalStatusView, error) {
	out := new(resources.ParentalStatusView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ParentalStatusViewService/GetParentalStatusView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParentalStatusViewServiceServer is the server API for ParentalStatusViewService service.
type ParentalStatusViewServiceServer interface {
	// Returns the requested parental status view in full detail.
	GetParentalStatusView(context.Context, *GetParentalStatusViewRequest) (*resources.ParentalStatusView, error)
}

// UnimplementedParentalStatusViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedParentalStatusViewServiceServer struct {
}

func (*UnimplementedParentalStatusViewServiceServer) GetParentalStatusView(ctx context.Context, req *GetParentalStatusViewRequest) (*resources.ParentalStatusView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParentalStatusView not implemented")
}

func RegisterParentalStatusViewServiceServer(s *grpc.Server, srv ParentalStatusViewServiceServer) {
	s.RegisterService(&_ParentalStatusViewService_serviceDesc, srv)
}

func _ParentalStatusViewService_GetParentalStatusView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParentalStatusViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParentalStatusViewServiceServer).GetParentalStatusView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ParentalStatusViewService/GetParentalStatusView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParentalStatusViewServiceServer).GetParentalStatusView(ctx, req.(*GetParentalStatusViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ParentalStatusViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ParentalStatusViewService",
	HandlerType: (*ParentalStatusViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParentalStatusView",
			Handler:    _ParentalStatusViewService_GetParentalStatusView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/parental_status_view_service.proto",
}