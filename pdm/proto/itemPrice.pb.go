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
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xfd, 0xda, 0xd2, 0x02, 0xd3, 0xef, 0xe3, 0xd3, 0x09, 0x68, 0xe5, 0x54, 0x7b, 0x6a, 0x62,
	0x52, 0x63, 0xbd, 0x7b, 0x23, 0xa6, 0x37, 0x52, 0x12, 0xaf, 0x04, 0xd9, 0x05, 0x37, 0xd8, 0x6e,
	0xdd, 0xdd, 0x92, 0xf0, 0x53, 0xfc, 0x51, 0xfe, 0x22, 0x2f, 0x66, 0xb7, 0x88, 0x51, 0xc4, 0x78,
	0x6a, 0x67, 0xde, 0x9b, 0x99, 0xf7, 0x5e, 0x16, 0xfe, 0x33, 0x45, 0x8b, 0xb1, 0x60, 0x73, 0x9a,
	0x54, 0x82, 0x2b, 0x8e, 0xbd, 0x25, 0x65, 0x4f, 0xac, 0x4c, 0xa4, 0x58, 0x27, 0x15, 0x29, 0x86,
	0x7f, 0xe7, 0xbc, 0x28, 0x78, 0xd9, 0xa0, 0xd1, 0x8b, 0x05, 0xdd, 0xec, 0x7d, 0x02, 0x7b, 0x60,
	0x33, 0x12, 0x58, 0xa1, 0x15, 0x3b, 0xb9, 0xcd, 0x08, 0x22, 0xb4, 0xd4, 0xa6, 0xa2, 0x81, 0x1d,
	0x5a, 0x71, 0x37, 0x37, 0xff, 0x18, 0xc1, 0x3f, 0xfd, 0x9d, 0xae, 0x67, 0x8f, 0x35, 0x9d, 0x32,
	0x12, 0x38, 0x86, 0xee, 0xeb, 0xe6, 0x9d, 0xee, 0x65, 0x04, 0x4f, 0xa1, 0xad, 0x65, 0x68, 0xb4,
	0x65, 0x50, 0x4f, 0x97, 0x19, 0xc1, 0x01, 0x78, 0x72, 0x55, 0xeb, 0xbe, 0x6b, 0xfa, 0xae, 0x5c,
	0xd5, 0x19, 0xc1, 0x13, 0xf0, 0x0a, 0xaa, 0x1e, 0x38, 0x09, 0x3c, 0x73, 0x69, 0x5b, 0x61, 0x1f,
	0xdc, 0x4a, 0x0b, 0x0b, 0xda, 0xa1, 0x15, 0xdb, 0x79, 0x53, 0xe0, 0x10, 0x3a, 0x84, 0xc9, 0x39,
	0xaf, 0x4b, 0x15, 0x74, 0x0c, 0xb0, 0xab, 0xa3, 0x57, 0x0b, 0x8e, 0x77, 0x7e, 0x72, 0x2a, 0x2b,
	0x5e, 0x4a, 0x8a, 0x57, 0xe0, 0xd1, 0x52, 0x31, 0xb5, 0x31, 0xde, 0xfc, 0xf4, 0x2c, 0xf9, 0x1c,
	0x4a, 0xf2, 0x31, 0xb2, 0x25, 0xe2, 0x05, 0xb8, 0xd5, 0x6c, 0x49, 0x85, 0xf1, 0xee, 0xa7, 0x83,
	0xaf, 0x13, 0x63, 0x0d, 0xe6, 0x0d, 0x07, 0x2f, 0xc1, 0xd5, 0x06, 0x65, 0xe0, 0x84, 0xce, 0xcf,
	0xeb, 0x1b, 0x9e, 0xde, 0x4e, 0x85, 0xe0, 0xc2, 0xc4, 0xf3, 0xcd, 0xf6, 0x91, 0x06, 0xf3, 0x86,
	0x83, 0x31, 0xb4, 0x58, 0xb9, 0xe0, 0x26, 0x32, 0x3f, 0xed, 0xef, 0x2d, 0x2f, 0x17, 0x3c, 0x37,
	0x8c, 0xf4, 0xd9, 0x82, 0xa3, 0xdd, 0xad, 0x09, 0x15, 0x6b, 0x1d, 0xd7, 0x08, 0x9c, 0x09, 0x55,
	0x78, 0x58, 0xd4, 0xf0, 0xfc, 0xb0, 0xde, 0x6d, 0x82, 0xd1, 0x1f, 0xbc, 0x01, 0xe7, 0x96, 0x2a,
	0xc4, 0x3d, 0x2e, 0xf9, 0xd5, 0xfc, 0xbd, 0x67, 0x1e, 0xdc, 0xf5, 0x5b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xaa, 0x04, 0x37, 0x3a, 0xa1, 0x02, 0x00, 0x00,
}
