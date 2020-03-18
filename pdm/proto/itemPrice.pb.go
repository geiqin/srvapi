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

type ItemCustomPrice struct {
	ItemId               int64        `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Type                 string       `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Method               string       `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Details              []*ItemPrice `protobuf:"bytes,9,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ItemCustomPrice) Reset()         { *m = ItemCustomPrice{} }
func (m *ItemCustomPrice) String() string { return proto.CompactTextString(m) }
func (*ItemCustomPrice) ProtoMessage()    {}
func (*ItemCustomPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e4669783059ba7d, []int{0}
}

func (m *ItemCustomPrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemCustomPrice.Unmarshal(m, b)
}
func (m *ItemCustomPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemCustomPrice.Marshal(b, m, deterministic)
}
func (m *ItemCustomPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemCustomPrice.Merge(m, src)
}
func (m *ItemCustomPrice) XXX_Size() int {
	return xxx_messageInfo_ItemCustomPrice.Size(m)
}
func (m *ItemCustomPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemCustomPrice.DiscardUnknown(m)
}

var xxx_messageInfo_ItemCustomPrice proto.InternalMessageInfo

func (m *ItemCustomPrice) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ItemCustomPrice) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ItemCustomPrice) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ItemCustomPrice) GetDetails() []*ItemPrice {
	if m != nil {
		return m.Details
	}
	return nil
}

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
	return fileDescriptor_3e4669783059ba7d, []int{1}
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
	return fileDescriptor_3e4669783059ba7d, []int{2}
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

type ItemCustomPriceResponse struct {
	Entity               *ItemCustomPrice   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager             `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*ItemCustomPrice `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error             `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info              `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ItemCustomPriceResponse) Reset()         { *m = ItemCustomPriceResponse{} }
func (m *ItemCustomPriceResponse) String() string { return proto.CompactTextString(m) }
func (*ItemCustomPriceResponse) ProtoMessage()    {}
func (*ItemCustomPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e4669783059ba7d, []int{3}
}

func (m *ItemCustomPriceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemCustomPriceResponse.Unmarshal(m, b)
}
func (m *ItemCustomPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemCustomPriceResponse.Marshal(b, m, deterministic)
}
func (m *ItemCustomPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemCustomPriceResponse.Merge(m, src)
}
func (m *ItemCustomPriceResponse) XXX_Size() int {
	return xxx_messageInfo_ItemCustomPriceResponse.Size(m)
}
func (m *ItemCustomPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemCustomPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ItemCustomPriceResponse proto.InternalMessageInfo

func (m *ItemCustomPriceResponse) GetEntity() *ItemCustomPrice {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ItemCustomPriceResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ItemCustomPriceResponse) GetItems() []*ItemCustomPrice {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ItemCustomPriceResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ItemCustomPriceResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemCustomPrice)(nil), "geiqin.srv.pdm.ItemCustomPrice")
	proto.RegisterType((*ItemPrice)(nil), "geiqin.srv.pdm.ItemPrice")
	proto.RegisterType((*ItemPriceResponse)(nil), "geiqin.srv.pdm.ItemPriceResponse")
	proto.RegisterType((*ItemCustomPriceResponse)(nil), "geiqin.srv.pdm.ItemCustomPriceResponse")
}

func init() { proto.RegisterFile("itemPrice.proto", fileDescriptor_3e4669783059ba7d) }

var fileDescriptor_3e4669783059ba7d = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0xce, 0xd2, 0x30,
	0x18, 0xb5, 0x1b, 0x1b, 0x3f, 0xdf, 0xf4, 0xff, 0xb5, 0x01, 0x99, 0xdc, 0xb8, 0xec, 0xc6, 0x25,
	0x26, 0x33, 0x8e, 0x18, 0x1f, 0xc0, 0x18, 0xb3, 0x3b, 0x1c, 0x89, 0xb7, 0x04, 0x69, 0xc1, 0x06,
	0xb6, 0xce, 0xb6, 0x23, 0xe1, 0x19, 0x8c, 0xef, 0xe3, 0x4b, 0xf8, 0x44, 0xde, 0x98, 0x76, 0x30,
	0x19, 0x12, 0x97, 0x18, 0xbd, 0x82, 0xef, 0x7c, 0xa7, 0xa7, 0xe7, 0x9c, 0x34, 0x83, 0x3b, 0xa6,
	0x68, 0x3e, 0x13, 0x6c, 0x45, 0xe3, 0x52, 0x70, 0xc5, 0xf1, 0xed, 0x86, 0xb2, 0xcf, 0xac, 0x88,
	0xa5, 0xd8, 0xc7, 0x25, 0xc9, 0x27, 0xf7, 0x57, 0x3c, 0xcf, 0x79, 0x51, 0x6f, 0xc3, 0x2f, 0x08,
	0xee, 0x52, 0x45, 0xf3, 0x37, 0x95, 0x54, 0xbc, 0x3e, 0x87, 0xc7, 0xd0, 0xd7, 0x22, 0x0b, 0x46,
	0x7c, 0x14, 0xa0, 0xc8, 0xce, 0x5c, 0x3d, 0xa6, 0x04, 0x63, 0xe8, 0xa9, 0x43, 0x49, 0x7d, 0x2b,
	0x40, 0xd1, 0x20, 0x33, 0xff, 0xf1, 0x63, 0x70, 0x73, 0xaa, 0x3e, 0x71, 0xe2, 0xdb, 0x06, 0x3d,
	0x4e, 0x78, 0x0a, 0x7d, 0x42, 0xd5, 0x92, 0xed, 0xa4, 0x3f, 0x08, 0xec, 0xc8, 0x4b, 0x9e, 0xc4,
	0x6d, 0x23, 0x71, 0x7a, 0x32, 0x9a, 0x9d, 0x98, 0xe1, 0x77, 0x04, 0x83, 0x06, 0xc6, 0xb7, 0x60,
	0x35, 0x16, 0x2c, 0x76, 0xfd, 0xfa, 0x10, 0x1e, 0xe8, 0xdf, 0xc5, 0x7e, 0xb9, 0xab, 0xa8, 0x76,
	0x6c, 0x1b, 0xba, 0xa7, 0xc1, 0x0f, 0x1a, 0x4b, 0xc9, 0x79, 0x9e, 0x5e, 0x2b, 0xcf, 0x08, 0x5c,
	0xb9, 0xad, 0x34, 0xee, 0x18, 0xdc, 0x91, 0xdb, 0x2a, 0x25, 0x67, 0x91, 0xdc, 0x56, 0xa4, 0x21,
	0x38, 0xa5, 0x36, 0xe6, 0xf7, 0x03, 0x14, 0x59, 0x59, 0x3d, 0xe0, 0x09, 0xdc, 0x10, 0x26, 0x57,
	0xbc, 0x2a, 0x94, 0x7f, 0x63, 0x16, 0xcd, 0x1c, 0xfe, 0x40, 0xf0, 0xe8, 0x57, 0x4c, 0x2a, 0x4b,
	0x5e, 0x48, 0x8a, 0x5f, 0x82, 0x4b, 0x0b, 0xc5, 0xd4, 0xc1, 0x64, 0xfb, 0x63, 0x33, 0x47, 0x22,
	0x7e, 0x0e, 0x4e, 0xb9, 0xdc, 0x50, 0x61, 0xb2, 0x7b, 0xc9, 0xe8, 0xf2, 0xc4, 0x4c, 0x2f, 0xb3,
	0x9a, 0x83, 0x5f, 0x80, 0xa3, 0x03, 0x4a, 0xdf, 0xee, 0x2a, 0xbe, 0xe6, 0x69, 0x75, 0x2a, 0x04,
	0x17, 0xa6, 0x9e, 0x2b, 0xea, 0x6f, 0xf5, 0x32, 0xab, 0x39, 0x38, 0x82, 0x1e, 0x2b, 0xd6, 0xdc,
	0x54, 0xe6, 0x25, 0xc3, 0xdf, 0xc4, 0x8b, 0x35, 0xcf, 0x0c, 0x23, 0xfc, 0x6a, 0xc1, 0xf8, 0xe2,
	0x6d, 0x35, 0x1d, 0xbc, 0xbe, 0xe8, 0xe0, 0xe9, 0x35, 0x93, 0xe7, 0x07, 0xff, 0xaa, 0x89, 0x57,
	0xed, 0x26, 0x3a, 0x2f, 0xf9, 0xaf, 0x7d, 0x24, 0xdf, 0x10, 0x3c, 0x6c, 0xba, 0x9f, 0x53, 0xb1,
	0xd7, 0xcf, 0xe7, 0x3d, 0xd8, 0x73, 0xaa, 0x70, 0x97, 0xb5, 0xc9, 0xb3, 0x2e, 0xef, 0xc7, 0x66,
	0xc3, 0x7b, 0x5a, 0xf2, 0xdd, 0xbf, 0x95, 0xfc, 0xe8, 0x9a, 0xaf, 0xc5, 0xf4, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xca, 0x1e, 0x84, 0x7b, 0x5e, 0x04, 0x00, 0x00,
}
