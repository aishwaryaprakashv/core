// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type TokenWrapper struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenWrapper) Reset()         { *m = TokenWrapper{} }
func (m *TokenWrapper) String() string { return proto.CompactTextString(m) }
func (*TokenWrapper) ProtoMessage()    {}
func (*TokenWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *TokenWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenWrapper.Unmarshal(m, b)
}
func (m *TokenWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenWrapper.Marshal(b, m, deterministic)
}
func (m *TokenWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenWrapper.Merge(m, src)
}
func (m *TokenWrapper) XXX_Size() int {
	return xxx_messageInfo_TokenWrapper.Size(m)
}
func (m *TokenWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_TokenWrapper proto.InternalMessageInfo

func (m *TokenWrapper) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type IsValidTokenRequest struct {
	Token                *TokenWrapper `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IsValidTokenRequest) Reset()         { *m = IsValidTokenRequest{} }
func (m *IsValidTokenRequest) String() string { return proto.CompactTextString(m) }
func (*IsValidTokenRequest) ProtoMessage()    {}
func (*IsValidTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *IsValidTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsValidTokenRequest.Unmarshal(m, b)
}
func (m *IsValidTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsValidTokenRequest.Marshal(b, m, deterministic)
}
func (m *IsValidTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsValidTokenRequest.Merge(m, src)
}
func (m *IsValidTokenRequest) XXX_Size() int {
	return xxx_messageInfo_IsValidTokenRequest.Size(m)
}
func (m *IsValidTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsValidTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsValidTokenRequest proto.InternalMessageInfo

func (m *IsValidTokenRequest) GetToken() *TokenWrapper {
	if m != nil {
		return m.Token
	}
	return nil
}

type IsWorkspaceAuthenticatedRequest struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Authority            string   `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	OriginalUri          string   `protobuf:"bytes,3,opt,name=originalUri,proto3" json:"originalUri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsWorkspaceAuthenticatedRequest) Reset()         { *m = IsWorkspaceAuthenticatedRequest{} }
func (m *IsWorkspaceAuthenticatedRequest) String() string { return proto.CompactTextString(m) }
func (*IsWorkspaceAuthenticatedRequest) ProtoMessage()    {}
func (*IsWorkspaceAuthenticatedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *IsWorkspaceAuthenticatedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsWorkspaceAuthenticatedRequest.Unmarshal(m, b)
}
func (m *IsWorkspaceAuthenticatedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsWorkspaceAuthenticatedRequest.Marshal(b, m, deterministic)
}
func (m *IsWorkspaceAuthenticatedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsWorkspaceAuthenticatedRequest.Merge(m, src)
}
func (m *IsWorkspaceAuthenticatedRequest) XXX_Size() int {
	return xxx_messageInfo_IsWorkspaceAuthenticatedRequest.Size(m)
}
func (m *IsWorkspaceAuthenticatedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsWorkspaceAuthenticatedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsWorkspaceAuthenticatedRequest proto.InternalMessageInfo

func (m *IsWorkspaceAuthenticatedRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *IsWorkspaceAuthenticatedRequest) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *IsWorkspaceAuthenticatedRequest) GetOriginalUri() string {
	if m != nil {
		return m.OriginalUri
	}
	return ""
}

func init() {
	proto.RegisterType((*TokenWrapper)(nil), "api.TokenWrapper")
	proto.RegisterType((*IsValidTokenRequest)(nil), "api.IsValidTokenRequest")
	proto.RegisterType((*IsWorkspaceAuthenticatedRequest)(nil), "api.IsWorkspaceAuthenticatedRequest")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0x4e, 0x41, 0x49, 0x18, 0xb8, 0xb8, 0x1a, 0xd2, 0x20, 0x89, 0xd8, 0x90, 0xc0, 0xa9, 0x0d,
	0x78, 0xf3, 0x60, 0x82, 0x91, 0x44, 0xae, 0xf8, 0xc3, 0x79, 0x81, 0x11, 0x26, 0x40, 0x77, 0xdd,
	0x4e, 0x49, 0xb8, 0xf2, 0x0a, 0x3e, 0x9a, 0xaf, 0xe0, 0x6b, 0x98, 0x98, 0x6e, 0x4b, 0xec, 0xc1,
	0x7a, 0xfc, 0xe6, 0x9b, 0x99, 0xef, 0x07, 0x40, 0xc6, 0xbc, 0xf2, 0xb5, 0x51, 0xac, 0x44, 0x59,
	0x6a, 0x6a, 0xb6, 0x96, 0x4a, 0x2d, 0x37, 0x18, 0x48, 0x4d, 0x81, 0x0c, 0x43, 0xc5, 0x92, 0x49,
	0x85, 0x51, 0xba, 0xd2, 0xbc, 0xcc, 0x58, 0x8b, 0x66, 0xf1, 0x5b, 0x80, 0x5b, 0xcd, 0xfb, 0x94,
	0xf4, 0x3a, 0x50, 0x7f, 0x56, 0x6b, 0x0c, 0xa7, 0x46, 0x6a, 0x8d, 0x46, 0x5c, 0xc0, 0x29, 0x27,
	0xd8, 0x75, 0xda, 0x4e, 0xaf, 0x3a, 0x49, 0x81, 0x77, 0x07, 0xe7, 0xe3, 0xe8, 0x55, 0x6e, 0x68,
	0x61, 0x97, 0x27, 0xf8, 0x1e, 0x63, 0xc4, 0xa2, 0x9b, 0x5f, 0xae, 0x0d, 0xce, 0x7c, 0xa9, 0xc9,
	0xcf, 0xbf, 0x3b, 0xde, 0xef, 0xe1, 0x6a, 0x1c, 0x4d, 0x95, 0x59, 0x47, 0x5a, 0xce, 0x71, 0x18,
	0xf3, 0x0a, 0x43, 0xa6, 0xb9, 0x64, 0x5c, 0x1c, 0x7f, 0x35, 0xa0, 0xb2, 0x45, 0x5e, 0xa9, 0x45,
	0xa6, 0x9c, 0x21, 0xd1, 0x82, 0x6a, 0x12, 0x57, 0x19, 0xe2, 0xbd, 0x5b, 0xb2, 0xd4, 0xef, 0x40,
	0xb4, 0xa1, 0xa6, 0x0c, 0x2d, 0x29, 0x94, 0x9b, 0x17, 0x43, 0x6e, 0xd9, 0xf2, 0xf9, 0xd1, 0xe0,
	0xdb, 0x81, 0x5a, 0x22, 0xf8, 0x84, 0x66, 0x47, 0x73, 0x14, 0x04, 0xf5, 0x7c, 0x14, 0xe1, 0x5a,
	0xd3, 0x7f, 0xa4, 0x6b, 0x36, 0xfc, 0xb4, 0x38, 0xff, 0x58, 0x9c, 0x3f, 0x4a, 0x8a, 0xf3, 0xba,
	0x87, 0xcf, 0xaf, 0x8f, 0xd2, 0xb5, 0xe7, 0x26, 0x7d, 0x47, 0xc1, 0xae, 0x3f, 0x43, 0x96, 0xfd,
	0x20, 0x31, 0x15, 0xd8, 0xb8, 0xb7, 0x69, 0x6a, 0x71, 0x70, 0xc0, 0x2d, 0x8a, 0x2d, 0x3a, 0x99,
	0xee, 0xbf, 0xad, 0x14, 0x7a, 0xe8, 0x59, 0x0f, 0xde, 0x7d, 0x1b, 0x4e, 0x1e, 0x47, 0xc3, 0x07,
	0x51, 0xe8, 0x65, 0x56, 0xb1, 0x97, 0x37, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x6b, 0x01,
	0x48, 0x35, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	IsValidToken(ctx context.Context, in *IsValidTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	IsWorkspaceAuthenticated(ctx context.Context, in *IsWorkspaceAuthenticatedRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) IsValidToken(ctx context.Context, in *IsValidTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.AuthService/IsValidToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) IsWorkspaceAuthenticated(ctx context.Context, in *IsWorkspaceAuthenticatedRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.AuthService/IsWorkspaceAuthenticated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	IsValidToken(context.Context, *IsValidTokenRequest) (*empty.Empty, error)
	IsWorkspaceAuthenticated(context.Context, *IsWorkspaceAuthenticatedRequest) (*empty.Empty, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) IsValidToken(ctx context.Context, req *IsValidTokenRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsValidToken not implemented")
}
func (*UnimplementedAuthServiceServer) IsWorkspaceAuthenticated(ctx context.Context, req *IsWorkspaceAuthenticatedRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsWorkspaceAuthenticated not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_IsValidToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsValidTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsValidToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/IsValidToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsValidToken(ctx, req.(*IsValidTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_IsWorkspaceAuthenticated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsWorkspaceAuthenticatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsWorkspaceAuthenticated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/IsWorkspaceAuthenticated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsWorkspaceAuthenticated(ctx, req.(*IsWorkspaceAuthenticatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsValidToken",
			Handler:    _AuthService_IsValidToken_Handler,
		},
		{
			MethodName: "IsWorkspaceAuthenticated",
			Handler:    _AuthService_IsWorkspaceAuthenticated_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
