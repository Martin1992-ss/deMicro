// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

// 用户信息
type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company              string   `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               *Error   `protobuf:"bytes,3,opt,name=errors,proto3" json:"errors,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() *Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Token) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*Token)(nil), "user.Token")
	proto.RegisterType((*Error)(nil), "user.Error")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4b, 0x6b, 0xdb, 0x40,
	0x10, 0xb6, 0x9e, 0xb6, 0x47, 0xd8, 0x87, 0xa1, 0xd0, 0xc5, 0x87, 0xa2, 0xca, 0x50, 0x5a, 0x0a,
	0x3e, 0xb8, 0xe7, 0x1e, 0x4c, 0x29, 0xa6, 0x57, 0xf5, 0x71, 0x2d, 0x8a, 0x76, 0x20, 0x4b, 0x64,
	0xad, 0xbc, 0xbb, 0x76, 0xc8, 0x0f, 0xcd, 0xff, 0x09, 0x3b, 0x92, 0x83, 0x6d, 0x12, 0x72, 0xd2,
	0x7c, 0x0f, 0x46, 0xdf, 0x7c, 0x12, 0xc0, 0xc1, 0x92, 0x59, 0x75, 0x46, 0x3b, 0x8d, 0xb1, 0x9f,
	0x8b, 0x23, 0xc4, 0x7f, 0x2d, 0x19, 0x9c, 0x43, 0xa8, 0xa4, 0x08, 0xf2, 0xe0, 0xf3, 0xb4, 0x0c,
	0x95, 0x44, 0x84, 0xb8, 0xad, 0x76, 0x24, 0x42, 0x66, 0x78, 0x46, 0x01, 0xe3, 0x5a, 0xef, 0xba,
	0xaa, 0x7d, 0x10, 0x11, 0xd3, 0x27, 0x88, 0xef, 0x20, 0xa1, 0x5d, 0xa5, 0x1a, 0x11, 0x33, 0xdf,
	0x03, 0x5c, 0xc0, 0xa4, 0xab, 0xac, 0xbd, 0xd7, 0x46, 0x8a, 0x84, 0x85, 0x67, 0x5c, 0x4c, 0x61,
	0x5c, 0xd2, 0xfe, 0x40, 0xd6, 0x15, 0x7b, 0x98, 0x94, 0x64, 0x3b, 0xdd, 0x5a, 0xc2, 0x0f, 0xc0,
	0xb1, 0x38, 0x48, 0xb6, 0x86, 0x15, 0xe7, 0xf5, 0x01, 0x4b, 0xe6, 0x31, 0x87, 0xc4, 0x3f, 0xad,
	0x08, 0xf3, 0xe8, 0xca, 0xd0, 0x0b, 0xb8, 0x84, 0x94, 0x8c, 0xd1, 0xc6, 0x8a, 0x88, 0x2d, 0x59,
	0x6f, 0xf9, 0xe9, 0xb9, 0x72, 0x90, 0x8a, 0x3d, 0x24, 0x7f, 0xf4, 0x1d, 0xb5, 0x3e, 0xb8, 0xf3,
	0xc3, 0x70, 0x79, 0x0f, 0x3c, 0x7b, 0xac, 0x1a, 0x25, 0xf9, 0xfa, 0x49, 0xd9, 0x83, 0x8b, 0xcd,
	0xc1, 0x2b, 0x9b, 0xf1, 0x3d, 0x8c, 0x3d, 0xfb, 0x5f, 0xc9, 0xa1, 0x8b, 0xd4, 0xc3, 0x5f, 0xb2,
	0xf8, 0x0e, 0x09, 0x3b, 0x7d, 0xb3, 0xb5, 0x96, 0xc4, 0x6f, 0x4c, 0x4a, 0x9e, 0x31, 0x87, 0x4c,
	0x92, 0xad, 0x8d, 0xea, 0x9c, 0xd2, 0xed, 0x50, 0xfa, 0x39, 0xb5, 0x7e, 0x0c, 0x20, 0xf3, 0x67,
	0xfe, 0x26, 0x73, 0x54, 0x35, 0xe1, 0x27, 0x48, 0x7f, 0x18, 0xaa, 0x1c, 0xe1, 0x59, 0x07, 0x8b,
	0x79, 0x3f, 0x9f, 0xea, 0x2c, 0x46, 0xb8, 0x84, 0x68, 0x4b, 0xee, 0x0d, 0xd3, 0x17, 0x48, 0xb7,
	0xe4, 0x36, 0x4d, 0x83, 0xb3, 0x93, 0xc6, 0x9f, 0xe6, 0x05, 0xeb, 0x47, 0x88, 0x37, 0x07, 0x77,
	0x7b, 0xb1, 0x70, 0x28, 0x82, 0x1b, 0x2d, 0x46, 0xf8, 0x15, 0x66, 0xff, 0x7c, 0x61, 0x95, 0xa3,
	0xbe, 0xe4, 0x73, 0xfd, 0xca, 0x7c, 0x93, 0xf2, 0xcf, 0xf8, 0xed, 0x29, 0x00, 0x00, 0xff, 0xff,
	0x8b, 0xa2, 0xd8, 0x9d, 0x9a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}
