// Code generated by protoc-gen-go. DO NOT EDIT.
// source: limitDiscountGoods.proto

package geiqin_srv_ims_discount

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

type LimitDiscountGoods struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LimitDiscountId      int64    `protobuf:"varint,2,opt,name=limit_discount_id,json=limitDiscountId,proto3" json:"limit_discount_id,omitempty"`
	ItemId               int64    `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64    `protobuf:"varint,4,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Discount             float32  `protobuf:"fixed32,6,opt,name=discount,proto3" json:"discount,omitempty"`
	Money                float32  `protobuf:"fixed32,7,opt,name=money,proto3" json:"money,omitempty"`
	CreatedAt            string   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LimitDiscountGoods) Reset()         { *m = LimitDiscountGoods{} }
func (m *LimitDiscountGoods) String() string { return proto.CompactTextString(m) }
func (*LimitDiscountGoods) ProtoMessage()    {}
func (*LimitDiscountGoods) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aed98e0b2304ae1, []int{0}
}

func (m *LimitDiscountGoods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LimitDiscountGoods.Unmarshal(m, b)
}
func (m *LimitDiscountGoods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LimitDiscountGoods.Marshal(b, m, deterministic)
}
func (m *LimitDiscountGoods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitDiscountGoods.Merge(m, src)
}
func (m *LimitDiscountGoods) XXX_Size() int {
	return xxx_messageInfo_LimitDiscountGoods.Size(m)
}
func (m *LimitDiscountGoods) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitDiscountGoods.DiscardUnknown(m)
}

var xxx_messageInfo_LimitDiscountGoods proto.InternalMessageInfo

func (m *LimitDiscountGoods) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LimitDiscountGoods) GetLimitDiscountId() int64 {
	if m != nil {
		return m.LimitDiscountId
	}
	return 0
}

func (m *LimitDiscountGoods) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *LimitDiscountGoods) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *LimitDiscountGoods) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *LimitDiscountGoods) GetDiscount() float32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func (m *LimitDiscountGoods) GetMoney() float32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *LimitDiscountGoods) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *LimitDiscountGoods) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

//
type LimitDiscountGoodsResponse struct {
	Entity               *LimitDiscountGoods   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager                `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*LimitDiscountGoods `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error                `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info                 `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LimitDiscountGoodsResponse) Reset()         { *m = LimitDiscountGoodsResponse{} }
func (m *LimitDiscountGoodsResponse) String() string { return proto.CompactTextString(m) }
func (*LimitDiscountGoodsResponse) ProtoMessage()    {}
func (*LimitDiscountGoodsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aed98e0b2304ae1, []int{1}
}

func (m *LimitDiscountGoodsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LimitDiscountGoodsResponse.Unmarshal(m, b)
}
func (m *LimitDiscountGoodsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LimitDiscountGoodsResponse.Marshal(b, m, deterministic)
}
func (m *LimitDiscountGoodsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitDiscountGoodsResponse.Merge(m, src)
}
func (m *LimitDiscountGoodsResponse) XXX_Size() int {
	return xxx_messageInfo_LimitDiscountGoodsResponse.Size(m)
}
func (m *LimitDiscountGoodsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitDiscountGoodsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LimitDiscountGoodsResponse proto.InternalMessageInfo

func (m *LimitDiscountGoodsResponse) GetEntity() *LimitDiscountGoods {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *LimitDiscountGoodsResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *LimitDiscountGoodsResponse) GetItems() []*LimitDiscountGoods {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *LimitDiscountGoodsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *LimitDiscountGoodsResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*LimitDiscountGoods)(nil), "geiqin.srv.ims.discount.LimitDiscountGoods")
	proto.RegisterType((*LimitDiscountGoodsResponse)(nil), "geiqin.srv.ims.discount.LimitDiscountGoodsResponse")
}

func init() {
	proto.RegisterFile("limitDiscountGoods.proto", fileDescriptor_5aed98e0b2304ae1)
}

var fileDescriptor_5aed98e0b2304ae1 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xad, 0x25, 0x4b, 0x89, 0xc7, 0xa5, 0xa5, 0x43, 0x4b, 0xb6, 0x86, 0x14, 0xa1, 0x93, 0x69,
	0x41, 0x50, 0xa5, 0x3f, 0xe0, 0x26, 0x25, 0x08, 0x7a, 0x28, 0xca, 0xa1, 0x47, 0xa3, 0x6a, 0x27,
	0xc9, 0x12, 0x6b, 0x57, 0xdd, 0x5d, 0x05, 0x7c, 0xef, 0x9f, 0xf5, 0xd0, 0xdf, 0x2a, 0xbb, 0x92,
	0x0b, 0xc5, 0x55, 0xc1, 0x39, 0xe4, 0xa6, 0x99, 0xf7, 0xe6, 0xbd, 0x9d, 0x37, 0x20, 0x60, 0x1b,
	0xd1, 0x08, 0x7b, 0x21, 0x4c, 0xad, 0x3a, 0x69, 0x2f, 0x95, 0xe2, 0x26, 0x6b, 0xb5, 0xb2, 0x0a,
	0x4f, 0x6e, 0x48, 0x7c, 0x17, 0x32, 0x33, 0xfa, 0x3e, 0x13, 0x8d, 0xc9, 0xf8, 0xc0, 0x59, 0x3c,
	0xad, 0x55, 0xd3, 0x28, 0xd9, 0xd3, 0xd2, 0x1f, 0x01, 0xe0, 0xe7, 0x3d, 0x0d, 0x7c, 0x06, 0x81,
	0xe0, 0x6c, 0x92, 0x4c, 0x96, 0x61, 0x19, 0x08, 0x8e, 0x6f, 0xe1, 0x85, 0x77, 0x5a, 0xef, 0x64,
	0xd6, 0x82, 0xb3, 0xc0, 0xc3, 0xcf, 0xff, 0x7a, 0x42, 0xc1, 0xf1, 0x04, 0x8e, 0x84, 0xa5, 0xc6,
	0x31, 0x42, 0xcf, 0x88, 0x5d, 0x59, 0x70, 0x7c, 0x05, 0xb1, 0xb9, 0xeb, 0x5c, 0x7f, 0xea, 0xfb,
	0x91, 0xb9, 0xeb, 0x0a, 0x8e, 0x08, 0x53, 0xbb, 0x6d, 0x89, 0x45, 0xc9, 0x64, 0x39, 0x2b, 0xfd,
	0x37, 0x2e, 0xe0, 0x78, 0xe7, 0xc4, 0xe2, 0x64, 0xb2, 0x0c, 0xca, 0x3f, 0x35, 0xbe, 0x84, 0xa8,
	0x51, 0x92, 0xb6, 0xec, 0xc8, 0x03, 0x7d, 0x81, 0xa7, 0x00, 0xb5, 0xa6, 0xca, 0x12, 0x5f, 0x57,
	0x96, 0x1d, 0x7b, 0xad, 0xd9, 0xd0, 0x59, 0x59, 0x07, 0x77, 0x2d, 0xdf, 0xc1, 0xb3, 0x1e, 0x1e,
	0x3a, 0x2b, 0x9b, 0xfe, 0x0c, 0x60, 0xb1, 0x1f, 0x43, 0x49, 0xa6, 0x55, 0xd2, 0x10, 0x9e, 0x43,
	0x4c, 0xd2, 0x0a, 0xbb, 0xf5, 0x91, 0xcc, 0xf3, 0x77, 0xd9, 0x48, 0xba, 0xd9, 0x3f, 0x44, 0x86,
	0x51, 0xfc, 0x00, 0x51, 0x5b, 0xdd, 0x90, 0xf6, 0xb9, 0xcd, 0xf3, 0x37, 0xa3, 0x1a, 0x5f, 0x1c,
	0xab, 0xec, 0xc9, 0xb8, 0x82, 0xc8, 0xc5, 0x67, 0x58, 0x98, 0x84, 0x87, 0x3a, 0xf7, 0x93, 0xce,
	0x98, 0xb4, 0x56, 0xda, 0xc7, 0xfe, 0x3f, 0xe3, 0x4f, 0x8e, 0x55, 0xf6, 0x64, 0x7c, 0x0f, 0x53,
	0x21, 0xaf, 0x95, 0x3f, 0xcb, 0x3c, 0x3f, 0x1d, 0x1d, 0x2a, 0xe4, 0xb5, 0x2a, 0x3d, 0x35, 0xff,
	0x15, 0xc2, 0xeb, 0xfd, 0x67, 0x5c, 0x91, 0xbe, 0x17, 0x35, 0xe1, 0x06, 0xe2, 0x73, 0x7f, 0x0f,
	0x3c, 0x64, 0x89, 0xc5, 0xd9, 0x21, 0x1b, 0x0f, 0x07, 0x4b, 0x9f, 0x38, 0xb7, 0x0b, 0xda, 0xd0,
	0x23, 0xb9, 0x09, 0x08, 0x2f, 0xc9, 0x3e, 0x8a, 0x55, 0x0d, 0xf1, 0x15, 0x55, 0xba, 0xbe, 0xc5,
	0x74, 0x54, 0xe0, 0x63, 0x65, 0xe8, 0xeb, 0x2d, 0x69, 0x7a, 0xa0, 0xc9, 0xb7, 0xd8, 0xff, 0x1d,
	0xce, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xb5, 0x26, 0x8d, 0x60, 0x04, 0x00, 0x00,
}
