// Code generated by protoc-gen-go. DO NOT EDIT.
// source: refunding.proto

package geiqin_srv_crm_integral

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Refunding struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderDetailIds       []int64  `protobuf:"varint,2,rep,packed,name=order_detail_ids,json=orderDetailIds,proto3" json:"order_detail_ids,omitempty"`
	ReturnPoints         int32    `protobuf:"varint,3,opt,name=return_points,json=returnPoints,proto3" json:"return_points,omitempty"`
	DeductPoints         int32    `protobuf:"varint,4,opt,name=deduct_points,json=deductPoints,proto3" json:"deduct_points,omitempty"`
	DeductMoney          float32  `protobuf:"fixed32,5,opt,name=deduct_money,json=deductMoney,proto3" json:"deduct_money,omitempty"`
	PointsRate           int32    `protobuf:"varint,6,opt,name=points_rate,json=pointsRate,proto3" json:"points_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Refunding) Reset()         { *m = Refunding{} }
func (m *Refunding) String() string { return proto.CompactTextString(m) }
func (*Refunding) ProtoMessage()    {}
func (*Refunding) Descriptor() ([]byte, []int) {
	return fileDescriptor_128a081085c2ac4a, []int{0}
}

func (m *Refunding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Refunding.Unmarshal(m, b)
}
func (m *Refunding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Refunding.Marshal(b, m, deterministic)
}
func (m *Refunding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Refunding.Merge(m, src)
}
func (m *Refunding) XXX_Size() int {
	return xxx_messageInfo_Refunding.Size(m)
}
func (m *Refunding) XXX_DiscardUnknown() {
	xxx_messageInfo_Refunding.DiscardUnknown(m)
}

var xxx_messageInfo_Refunding proto.InternalMessageInfo

func (m *Refunding) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Refunding) GetOrderDetailIds() []int64 {
	if m != nil {
		return m.OrderDetailIds
	}
	return nil
}

func (m *Refunding) GetReturnPoints() int32 {
	if m != nil {
		return m.ReturnPoints
	}
	return 0
}

func (m *Refunding) GetDeductPoints() int32 {
	if m != nil {
		return m.DeductPoints
	}
	return 0
}

func (m *Refunding) GetDeductMoney() float32 {
	if m != nil {
		return m.DeductMoney
	}
	return 0
}

func (m *Refunding) GetPointsRate() int32 {
	if m != nil {
		return m.PointsRate
	}
	return 0
}

type RefundingResponse struct {
	Entity               *Refunding   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Refunding `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error       `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info        `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RefundingResponse) Reset()         { *m = RefundingResponse{} }
func (m *RefundingResponse) String() string { return proto.CompactTextString(m) }
func (*RefundingResponse) ProtoMessage()    {}
func (*RefundingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_128a081085c2ac4a, []int{1}
}

func (m *RefundingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefundingResponse.Unmarshal(m, b)
}
func (m *RefundingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefundingResponse.Marshal(b, m, deterministic)
}
func (m *RefundingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefundingResponse.Merge(m, src)
}
func (m *RefundingResponse) XXX_Size() int {
	return xxx_messageInfo_RefundingResponse.Size(m)
}
func (m *RefundingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefundingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefundingResponse proto.InternalMessageInfo

func (m *RefundingResponse) GetEntity() *Refunding {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *RefundingResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *RefundingResponse) GetItems() []*Refunding {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *RefundingResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *RefundingResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Refunding)(nil), "geiqin.srv.crm.integral.Refunding")
	proto.RegisterType((*RefundingResponse)(nil), "geiqin.srv.crm.integral.RefundingResponse")
}

func init() {
	proto.RegisterFile("refunding.proto", fileDescriptor_128a081085c2ac4a)
}

var fileDescriptor_128a081085c2ac4a = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0x9b, 0x30,
	0x14, 0xc7, 0x07, 0x84, 0x6c, 0x31, 0xd9, 0x96, 0xf9, 0x32, 0x16, 0x69, 0x1b, 0xa5, 0x17, 0xd4,
	0x03, 0x52, 0xd3, 0x1e, 0xaa, 0x5e, 0xdb, 0x1e, 0x72, 0xa8, 0x14, 0xb9, 0xe7, 0x0a, 0x51, 0xfc,
	0x82, 0x2c, 0x81, 0x4d, 0x1e, 0x26, 0x52, 0xbe, 0x51, 0x3f, 0x55, 0x3f, 0x4b, 0x85, 0x4d, 0xd2,
	0x53, 0xaa, 0x1c, 0xfd, 0xe3, 0xf7, 0xf7, 0xe3, 0xfd, 0x81, 0xfc, 0x44, 0x58, 0x77, 0x92, 0x0b,
	0x59, 0xa6, 0x0d, 0x2a, 0xad, 0xe8, 0xef, 0x12, 0xc4, 0x46, 0xc8, 0xb4, 0xc5, 0x6d, 0x5a, 0x60,
	0x9d, 0x0a, 0xa9, 0xa1, 0xc4, 0xbc, 0x9a, 0x4f, 0x0b, 0x55, 0xd7, 0x4a, 0x5a, 0x2d, 0x7e, 0x73,
	0xc8, 0x84, 0xed, 0xa3, 0xf4, 0x0f, 0xf9, 0xa6, 0x90, 0x03, 0x66, 0x82, 0x87, 0x4e, 0xe4, 0x24,
	0x1e, 0xfb, 0x6a, 0xce, 0x4b, 0x4e, 0x13, 0x32, 0xb3, 0x8f, 0x38, 0xe8, 0x5c, 0x54, 0x99, 0xe0,
	0x6d, 0xe8, 0x46, 0x5e, 0xe2, 0xb1, 0x1f, 0x86, 0xdf, 0x1b, 0xbc, 0xe4, 0x2d, 0x3d, 0x27, 0xdf,
	0x11, 0x74, 0x87, 0x32, 0x6b, 0x94, 0x90, 0xba, 0x0d, 0xbd, 0xc8, 0x49, 0x7c, 0x36, 0xb5, 0x70,
	0x65, 0x58, 0x2f, 0x71, 0xe0, 0x5d, 0xa1, 0xf7, 0xd2, 0xc8, 0x4a, 0x16, 0x0e, 0xd2, 0x19, 0x19,
	0xce, 0x59, 0xad, 0x24, 0xec, 0x42, 0x3f, 0x72, 0x12, 0x97, 0x05, 0x96, 0x3d, 0xf6, 0x88, 0xfe,
	0x27, 0x81, 0xbd, 0x20, 0xc3, 0x5c, 0x43, 0x38, 0x36, 0xb7, 0x10, 0x8b, 0x58, 0xae, 0x21, 0x7e,
	0x75, 0xc9, 0xaf, 0xc3, 0x82, 0x0c, 0xda, 0x46, 0xc9, 0x16, 0xe8, 0x2d, 0x19, 0x83, 0xd4, 0x42,
	0xef, 0xcc, 0x9a, 0xc1, 0x22, 0x4e, 0x8f, 0xd4, 0x95, 0x7e, 0x64, 0x87, 0x04, 0xbd, 0x26, 0x7e,
	0x93, 0x97, 0x80, 0xa1, 0x6b, 0xa2, 0xff, 0x8e, 0x46, 0x57, 0xbd, 0xc5, 0xac, 0x4c, 0x6f, 0x88,
	0x2f, 0x34, 0xd4, 0x7d, 0x1b, 0xde, 0x89, 0x03, 0x6d, 0xa0, 0x9f, 0x07, 0x88, 0x0a, 0x4d, 0x45,
	0x9f, 0xcd, 0x7b, 0xe8, 0x2d, 0x66, 0x65, 0x7a, 0x49, 0x46, 0x42, 0xae, 0x95, 0xe9, 0x2c, 0x58,
	0xfc, 0x3d, 0x1a, 0x5a, 0xca, 0xb5, 0x62, 0x46, 0x5d, 0x6c, 0xc8, 0xec, 0x30, 0xfc, 0x09, 0x70,
	0x2b, 0x0a, 0xa0, 0xcf, 0x64, 0x72, 0x97, 0x57, 0x45, 0x57, 0xe5, 0x1a, 0xe8, 0x09, 0x2f, 0x3d,
	0xbf, 0x38, 0x61, 0xb1, 0xe1, 0x2b, 0xc4, 0x5f, 0x5e, 0xc6, 0xe6, 0x2f, 0xbc, 0x7a, 0x0f, 0x00,
	0x00, 0xff, 0xff, 0x3c, 0x6e, 0x20, 0xc6, 0xbf, 0x02, 0x00, 0x00,
}
