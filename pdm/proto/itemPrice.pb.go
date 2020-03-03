// Code generated by protoc-gen-go. DO NOT EDIT.
// source: itemPrice.proto

package geiqin_srv_pdm

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

type ItemPrice struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	TypeValueId          int64    `protobuf:"varint,3,opt,name=type_value_id,json=typeValueId,proto3" json:"type_value_id,omitempty"`
	ItemId               int64    `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64    `protobuf:"varint,5,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Method               string   `protobuf:"bytes,6,opt,name=method,proto3" json:"method,omitempty"`
	Price                float32  `protobuf:"fixed32,7,opt,name=price,proto3" json:"price,omitempty"`
	Discount             float32  `protobuf:"fixed32,8,opt,name=discount,proto3" json:"discount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemPrice) Reset()         { *m = ItemPrice{} }
func (m *ItemPrice) String() string { return proto.CompactTextString(m) }
func (*ItemPrice) ProtoMessage()    {}
func (*ItemPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e4669783059ba7d, []int{0}
}

func (m *ItemPrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemPrice.Unmarshal(m, b)
}
func (m *ItemPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemPrice.Marshal(b, m, deterministic)
}
func (m *ItemPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemPrice.Merge(m, src)
}
func (m *ItemPrice) XXX_Size() int {
	return xxx_messageInfo_ItemPrice.Size(m)
}
func (m *ItemPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemPrice.DiscardUnknown(m)
}

var xxx_messageInfo_ItemPrice proto.InternalMessageInfo

func (m *ItemPrice) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemPrice) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ItemPrice) GetTypeValueId() int64 {
	if m != nil {
		return m.TypeValueId
	}
	return 0
}

func (m *ItemPrice) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ItemPrice) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *ItemPrice) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ItemPrice) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ItemPrice) GetDiscount() float32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

type ItemPriceResponse struct {
	Entity               *ItemPrice   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*ItemPrice `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error       `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info        `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ItemPriceResponse) Reset()         { *m = ItemPriceResponse{} }
func (m *ItemPriceResponse) String() string { return proto.CompactTextString(m) }
func (*ItemPriceResponse) ProtoMessage()    {}
func (*ItemPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e4669783059ba7d, []int{1}
}

func (m *ItemPriceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemPriceResponse.Unmarshal(m, b)
}
func (m *ItemPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemPriceResponse.Marshal(b, m, deterministic)
}
func (m *ItemPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemPriceResponse.Merge(m, src)
}
func (m *ItemPriceResponse) XXX_Size() int {
	return xxx_messageInfo_ItemPriceResponse.Size(m)
}
func (m *ItemPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ItemPriceResponse proto.InternalMessageInfo

func (m *ItemPriceResponse) GetEntity() *ItemPrice {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ItemPriceResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ItemPriceResponse) GetItems() []*ItemPrice {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ItemPriceResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ItemPriceResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemPrice)(nil), "geiqin.srv.pdm.ItemPrice")
	proto.RegisterType((*ItemPriceResponse)(nil), "geiqin.srv.pdm.ItemPriceResponse")
}

func init() { proto.RegisterFile("itemPrice.proto", fileDescriptor_3e4669783059ba7d) }

var fileDescriptor_3e4669783059ba7d = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xb5, 0x2d, 0x2d, 0x30, 0x55, 0xd4, 0x09, 0x68, 0xe5, 0x54, 0x7b, 0x6a, 0x62, 0x52, 0x63,
	0xfd, 0x0d, 0xc4, 0xf4, 0x46, 0x4a, 0xe2, 0x95, 0x20, 0xbb, 0xe0, 0x06, 0xdb, 0xad, 0xbb, 0x5b,
	0x12, 0xfe, 0x8d, 0x7f, 0xc8, 0x5f, 0xe4, 0xc5, 0xec, 0x16, 0x31, 0x7e, 0xe0, 0xc5, 0x53, 0x3b,
	0xf3, 0xde, 0xbc, 0x99, 0xf7, 0xb2, 0x70, 0xcc, 0x14, 0x2d, 0xc6, 0x82, 0xcd, 0x69, 0x52, 0x09,
	0xae, 0x38, 0xf6, 0x96, 0x94, 0x3d, 0xb3, 0x32, 0x91, 0x62, 0x9d, 0x54, 0xa4, 0x18, 0x1e, 0xce,
	0x79, 0x51, 0xf0, 0xb2, 0x41, 0xa3, 0x57, 0x0b, 0xba, 0xd9, 0xc7, 0x04, 0xf6, 0xc0, 0x66, 0x24,
	0xb0, 0x42, 0x2b, 0x76, 0x72, 0x9b, 0x11, 0x44, 0x68, 0xa9, 0x4d, 0x45, 0x03, 0x3b, 0xb4, 0xe2,
	0x6e, 0x6e, 0xfe, 0x31, 0x82, 0x23, 0xfd, 0x9d, 0xae, 0x67, 0x4f, 0x35, 0x9d, 0x32, 0x12, 0x38,
	0x86, 0xee, 0xeb, 0xe6, 0xbd, 0xee, 0x65, 0x04, 0xcf, 0xa1, 0xad, 0xcf, 0xd0, 0x68, 0xcb, 0xa0,
	0x9e, 0x2e, 0x33, 0x82, 0x03, 0xf0, 0xe4, 0xaa, 0xd6, 0x7d, 0xd7, 0xf4, 0x5d, 0xb9, 0xaa, 0x33,
	0x82, 0x67, 0xe0, 0x15, 0x54, 0x3d, 0x72, 0x12, 0x78, 0x66, 0xd3, 0xb6, 0xc2, 0x3e, 0xb8, 0x95,
	0x3e, 0x2c, 0x68, 0x87, 0x56, 0x6c, 0xe7, 0x4d, 0x81, 0x43, 0xe8, 0x10, 0x26, 0xe7, 0xbc, 0x2e,
	0x55, 0xd0, 0x31, 0xc0, 0xae, 0x8e, 0xde, 0x2c, 0x38, 0xdd, 0xf9, 0xc9, 0xa9, 0xac, 0x78, 0x29,
	0x29, 0xde, 0x80, 0x47, 0x4b, 0xc5, 0xd4, 0xc6, 0x78, 0xf3, 0xd3, 0x8b, 0xe4, 0x6b, 0x28, 0xc9,
	0xe7, 0xc8, 0x96, 0x88, 0x57, 0xe0, 0x56, 0xb3, 0x25, 0x15, 0xc6, 0xbb, 0x9f, 0x0e, 0xbe, 0x4f,
	0x8c, 0x35, 0x98, 0x37, 0x1c, 0xbc, 0x06, 0x57, 0x1b, 0x94, 0x81, 0x13, 0x3a, 0x7f, 0xcb, 0x37,
	0x3c, 0xad, 0x4e, 0x85, 0xe0, 0xc2, 0xc4, 0xf3, 0x8b, 0xfa, 0x48, 0x83, 0x79, 0xc3, 0xc1, 0x18,
	0x5a, 0xac, 0x5c, 0x70, 0x13, 0x99, 0x9f, 0xf6, 0x7f, 0x88, 0x97, 0x0b, 0x9e, 0x1b, 0x46, 0xfa,
	0x62, 0xc1, 0xc9, 0x6e, 0xd7, 0x84, 0x8a, 0xb5, 0x8e, 0x6b, 0x04, 0xce, 0x84, 0x2a, 0xdc, 0x7f,
	0xd4, 0xf0, 0x72, 0xff, 0xbd, 0xdb, 0x04, 0xa3, 0x03, 0x2d, 0x73, 0xf7, 0x7f, 0x99, 0x07, 0xcf,
	0xbc, 0xbb, 0xdb, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x98, 0xc3, 0x0f, 0xa8, 0x02, 0x00,
	0x00,
}
