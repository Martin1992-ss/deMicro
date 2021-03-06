// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vessel.proto

package vessel

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

// 每条货轮的熟悉
type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwerId               string   `protobuf:"bytes,6,opt,name=ower_id,json=owerId,proto3" json:"ower_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0c0e45ee319d513, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwerId() string {
	if m != nil {
		return m.OwerId
	}
	return ""
}

// 等待运送的货物
type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0c0e45ee319d513, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

// 货轮装得下的话
// 返回的多条货轮信息
type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	Created              bool      `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0c0e45ee319d513, []int{2}
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

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func init() {
	proto.RegisterType((*Vessel)(nil), "vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "vessel.Specification")
	proto.RegisterType((*Response)(nil), "vessel.Response")
}

func init() { proto.RegisterFile("vessel.proto", fileDescriptor_e0c0e45ee319d513) }

var fileDescriptor_e0c0e45ee319d513 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0xfd, 0x6d, 0xda, 0x6e, 0xd3, 0xf9, 0x99, 0x22, 0x03, 0xe2, 0x52, 0x14, 0x42, 0x0e, 0x92,
	0x83, 0xf4, 0x50, 0x6f, 0xde, 0x44, 0x10, 0xf4, 0xb8, 0x05, 0x3d, 0x96, 0xed, 0xee, 0xa8, 0x0b,
	0xf9, 0x47, 0x12, 0xd2, 0xf6, 0xcb, 0xf8, 0x59, 0xc5, 0x4d, 0x52, 0x69, 0x29, 0xde, 0xe6, 0xcd,
	0x7b, 0xfb, 0x78, 0xf3, 0x16, 0xce, 0x1a, 0xaa, 0x2a, 0x4a, 0xe6, 0x45, 0x99, 0xd7, 0x39, 0xf2,
	0x16, 0x45, 0x5f, 0x0c, 0xf8, 0xab, 0x1b, 0x71, 0x0a, 0x9e, 0x35, 0x82, 0x85, 0x2c, 0x9e, 0x48,
	0xcf, 0x1a, 0x9c, 0x81, 0xaf, 0x55, 0xa1, 0xb4, 0xad, 0x77, 0xc2, 0x0b, 0x59, 0x3c, 0x92, 0x7b,
	0x8c, 0xd7, 0x00, 0xa9, 0xda, 0xae, 0x36, 0x64, 0x3f, 0x3e, 0x6b, 0x31, 0x70, 0xec, 0x24, 0x55,
	0xdb, 0x37, 0xb7, 0x40, 0x84, 0x61, 0xa6, 0x52, 0x12, 0x43, 0x67, 0xe6, 0x66, 0xbc, 0x82, 0x89,
	0x6a, 0x94, 0x4d, 0xd4, 0x3a, 0x21, 0x31, 0x0a, 0x59, 0xec, 0xcb, 0xdf, 0x05, 0x5e, 0xc2, 0x38,
	0xdf, 0x50, 0xb9, 0xb2, 0x46, 0x70, 0xf7, 0x88, 0xff, 0xc0, 0x67, 0x13, 0xbd, 0x40, 0xb0, 0x2c,
	0x48, 0xdb, 0x77, 0xab, 0x55, 0x6d, 0xf3, 0xec, 0x20, 0x16, 0xfb, 0x33, 0x96, 0x77, 0x14, 0x2b,
	0x6a, 0xc0, 0x97, 0x54, 0x15, 0x79, 0x56, 0x11, 0xde, 0x40, 0x57, 0x81, 0x33, 0xf9, 0xbf, 0x98,
	0xce, 0xbb, 0x7e, 0xda, 0x36, 0x64, 0xc7, 0x62, 0x0c, 0xe3, 0x76, 0xaa, 0x84, 0x17, 0x0e, 0x4e,
	0x08, 0x7b, 0x1a, 0x05, 0x8c, 0x75, 0x49, 0xaa, 0x26, 0xe3, 0x0a, 0xf1, 0x65, 0x0f, 0x17, 0x3b,
	0x08, 0x5a, 0xf1, 0x92, 0xca, 0xc6, 0x6a, 0xc2, 0x7b, 0x08, 0x9e, 0x6c, 0x66, 0x1e, 0xf6, 0xe7,
	0x5f, 0xf4, 0xa6, 0x07, 0xb7, 0xce, 0xce, 0xfb, 0x75, 0x1f, 0x3b, 0xfa, 0x87, 0xb7, 0xc0, 0x1f,
	0x9d, 0x2f, 0x1e, 0x25, 0x39, 0xa5, 0x5e, 0x73, 0xf7, 0xdd, 0x77, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0x46, 0xce, 0x20, 0xfe, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	// 检查是否有能运送货物的轮船
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
	// 创建货轮
	Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselServiceClient) Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	// 检查是否有能运送货物的轮船
	FindAvailable(context.Context, *Specification, *Response) error
	// 创建货轮
	Create(context.Context, *Vessel, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func (h *VesselService) Create(ctx context.Context, in *Vessel, out *Response) error {
	return h.VesselServiceHandler.Create(ctx, in, out)
}
