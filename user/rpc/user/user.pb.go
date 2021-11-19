// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CaptchaRequest struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CaptchaRequest) Reset()         { *m = CaptchaRequest{} }
func (m *CaptchaRequest) String() string { return proto.CompactTextString(m) }
func (*CaptchaRequest) ProtoMessage()    {}
func (*CaptchaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *CaptchaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CaptchaRequest.Unmarshal(m, b)
}
func (m *CaptchaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CaptchaRequest.Marshal(b, m, deterministic)
}
func (m *CaptchaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CaptchaRequest.Merge(m, src)
}
func (m *CaptchaRequest) XXX_Size() int {
	return xxx_messageInfo_CaptchaRequest.Size(m)
}
func (m *CaptchaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CaptchaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CaptchaRequest proto.InternalMessageInfo

func (m *CaptchaRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type CaptchaResponse struct {
	CaptchaKey           string   `protobuf:"bytes,1,opt,name=captcha_key,json=captchaKey,proto3" json:"captcha_key,omitempty"`
	ExpiredAt            string   `protobuf:"bytes,2,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	CaptchaImageContent  string   `protobuf:"bytes,3,opt,name=captcha_image_content,json=captchaImageContent,proto3" json:"captcha_image_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CaptchaResponse) Reset()         { *m = CaptchaResponse{} }
func (m *CaptchaResponse) String() string { return proto.CompactTextString(m) }
func (*CaptchaResponse) ProtoMessage()    {}
func (*CaptchaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *CaptchaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CaptchaResponse.Unmarshal(m, b)
}
func (m *CaptchaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CaptchaResponse.Marshal(b, m, deterministic)
}
func (m *CaptchaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CaptchaResponse.Merge(m, src)
}
func (m *CaptchaResponse) XXX_Size() int {
	return xxx_messageInfo_CaptchaResponse.Size(m)
}
func (m *CaptchaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CaptchaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CaptchaResponse proto.InternalMessageInfo

func (m *CaptchaResponse) GetCaptchaKey() string {
	if m != nil {
		return m.CaptchaKey
	}
	return ""
}

func (m *CaptchaResponse) GetExpiredAt() string {
	if m != nil {
		return m.ExpiredAt
	}
	return ""
}

func (m *CaptchaResponse) GetCaptchaImageContent() string {
	if m != nil {
		return m.CaptchaImageContent
	}
	return ""
}

func init() {
	proto.RegisterType((*CaptchaRequest)(nil), "user.CaptchaRequest")
	proto.RegisterType((*CaptchaResponse)(nil), "user.CaptchaResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xd4, 0xb8, 0xf8, 0x9c, 0x13,
	0x0b, 0x4a, 0x92, 0x33, 0x12, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58,
	0x0b, 0x32, 0xf2, 0xf3, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0xa5, 0x56,
	0x46, 0x2e, 0x7e, 0xb8, 0xc2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x79, 0x2e, 0xee, 0x64,
	0x88, 0x50, 0x7c, 0x76, 0x6a, 0x25, 0x54, 0x3d, 0x17, 0x54, 0xc8, 0x3b, 0xb5, 0x52, 0x48, 0x96,
	0x8b, 0x2b, 0xb5, 0xa2, 0x20, 0xb3, 0x28, 0x35, 0x25, 0x3e, 0xb1, 0x44, 0x82, 0x09, 0x2c, 0xcf,
	0x09, 0x15, 0x71, 0x2c, 0x11, 0x32, 0xe2, 0x12, 0x85, 0xe9, 0xcf, 0xcc, 0x4d, 0x4c, 0x4f, 0x8d,
	0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x91, 0x60, 0x06, 0xab, 0x14, 0x86, 0x4a, 0x7a, 0x82,
	0xe4, 0x9c, 0x21, 0x52, 0x46, 0x8e, 0x5c, 0x2c, 0xa1, 0xc5, 0xa9, 0x45, 0x42, 0x96, 0x5c, 0x5c,
	0xee, 0xa9, 0x25, 0x50, 0x17, 0x09, 0x89, 0xe8, 0x81, 0x3d, 0x86, 0xea, 0x13, 0x29, 0x51, 0x34,
	0x51, 0x88, 0xb3, 0x93, 0xd8, 0xc0, 0xfe, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x04, 0x51,
	0x62, 0xde, 0x0d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetCaptcha(ctx context.Context, in *CaptchaRequest, opts ...grpc.CallOption) (*CaptchaResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetCaptcha(ctx context.Context, in *CaptchaRequest, opts ...grpc.CallOption) (*CaptchaResponse, error) {
	out := new(CaptchaResponse)
	err := c.cc.Invoke(ctx, "/user.User/GetCaptcha", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetCaptcha(context.Context, *CaptchaRequest) (*CaptchaResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetCaptcha(ctx context.Context, req *CaptchaRequest) (*CaptchaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaptcha not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetCaptcha",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetCaptcha(ctx, req.(*CaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCaptcha",
			Handler:    _User_GetCaptcha_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}