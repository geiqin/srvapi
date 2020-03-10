// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buying.proto

package geiqin_srv_ord_private

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

//购买清单
type Buying struct {
	Solution             string        `protobuf:"bytes,1,opt,name=solution,proto3" json:"solution,omitempty"`
	Count                int32         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Total                float32       `protobuf:"fixed32,3,opt,name=total,proto3" json:"total,omitempty"`
	Discount             float32       `protobuf:"fixed32,4,opt,name=discount,proto3" json:"discount,omitempty"`
	Freight              float32       `protobuf:"fixed32,5,opt,name=freight,proto3" json:"freight,omitempty"`
	Amount               float32       `protobuf:"fixed32,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Items                []*BuyingItem `protobuf:"bytes,7,rep,name=items,proto3" json:"items,omitempty"`
	Gifts                []*BuyingItem `protobuf:"bytes,9,rep,name=gifts,proto3" json:"gifts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Buying) Reset()         { *m = Buying{} }
func (m *Buying) String() string { return proto.CompactTextString(m) }
func (*Buying) ProtoMessage()    {}
func (*Buying) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc0a9e5c6a9833d6, []int{0}
}

func (m *Buying) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Buying.Unmarshal(m, b)
}
func (m *Buying) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Buying.Marshal(b, m, deterministic)
}
func (m *Buying) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Buying.Merge(m, src)
}
func (m *Buying) XXX_Size() int {
	return xxx_messageInfo_Buying.Size(m)
}
func (m *Buying) XXX_DiscardUnknown() {
	xxx_messageInfo_Buying.DiscardUnknown(m)
}

var xxx_messageInfo_Buying proto.InternalMessageInfo

func (m *Buying) GetSolution() string {
	if m != nil {
		return m.Solution
	}
	return ""
}

func (m *Buying) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Buying) GetTotal() float32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Buying) GetDiscount() float32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func (m *Buying) GetFreight() float32 {
	if m != nil {
		return m.Freight
	}
	return 0
}

func (m *Buying) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Buying) GetItems() []*BuyingItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Buying) GetGifts() []*BuyingItem {
	if m != nil {
		return m.Gifts
	}
	return nil
}

//购买清单明细
type BuyingItem struct {
	ItemId               int64    `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64    `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Num                  int32    `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
	Price                float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	OriginPrice          float32  `protobuf:"fixed32,5,opt,name=origin_price,json=originPrice,proto3" json:"origin_price,omitempty"`
	SubTotal             float32  `protobuf:"fixed32,6,opt,name=sub_total,json=subTotal,proto3" json:"sub_total,omitempty"`
	PromotionId          int64    `protobuf:"varint,7,opt,name=promotion_id,json=promotionId,proto3" json:"promotion_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyingItem) Reset()         { *m = BuyingItem{} }
func (m *BuyingItem) String() string { return proto.CompactTextString(m) }
func (*BuyingItem) ProtoMessage()    {}
func (*BuyingItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc0a9e5c6a9833d6, []int{1}
}

func (m *BuyingItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyingItem.Unmarshal(m, b)
}
func (m *BuyingItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyingItem.Marshal(b, m, deterministic)
}
func (m *BuyingItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyingItem.Merge(m, src)
}
func (m *BuyingItem) XXX_Size() int {
	return xxx_messageInfo_BuyingItem.Size(m)
}
func (m *BuyingItem) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyingItem.DiscardUnknown(m)
}

var xxx_messageInfo_BuyingItem proto.InternalMessageInfo

func (m *BuyingItem) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *BuyingItem) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *BuyingItem) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *BuyingItem) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *BuyingItem) GetOriginPrice() float32 {
	if m != nil {
		return m.OriginPrice
	}
	return 0
}

func (m *BuyingItem) GetSubTotal() float32 {
	if m != nil {
		return m.SubTotal
	}
	return 0
}

func (m *BuyingItem) GetPromotionId() int64 {
	if m != nil {
		return m.PromotionId
	}
	return 0
}

type BuyingResponse struct {
	Entity               *Buying   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Buying `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error    `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info     `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BuyingResponse) Reset()         { *m = BuyingResponse{} }
func (m *BuyingResponse) String() string { return proto.CompactTextString(m) }
func (*BuyingResponse) ProtoMessage()    {}
func (*BuyingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc0a9e5c6a9833d6, []int{2}
}

func (m *BuyingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyingResponse.Unmarshal(m, b)
}
func (m *BuyingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyingResponse.Marshal(b, m, deterministic)
}
func (m *BuyingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyingResponse.Merge(m, src)
}
func (m *BuyingResponse) XXX_Size() int {
	return xxx_messageInfo_BuyingResponse.Size(m)
}
func (m *BuyingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuyingResponse proto.InternalMessageInfo

func (m *BuyingResponse) GetEntity() *Buying {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *BuyingResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *BuyingResponse) GetItems() []*Buying {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *BuyingResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *BuyingResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Buying)(nil), "geiqin.srv.ord.private.Buying")
	proto.RegisterType((*BuyingItem)(nil), "geiqin.srv.ord.private.BuyingItem")
	proto.RegisterType((*BuyingResponse)(nil), "geiqin.srv.ord.private.BuyingResponse")
}

func init() { proto.RegisterFile("buying.proto", fileDescriptor_fc0a9e5c6a9833d6) }

var fileDescriptor_fc0a9e5c6a9833d6 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x49, 0xb2, 0x49, 0xb7, 0x93, 0x82, 0x90, 0x05, 0x4b, 0x54, 0xfe, 0xa8, 0x9b, 0x03,
	0xea, 0x29, 0x42, 0x29, 0x42, 0x9c, 0x41, 0x1c, 0x72, 0x5b, 0x19, 0x38, 0x57, 0x69, 0xe3, 0x04,
	0x6b, 0x13, 0x3b, 0xf8, 0x4f, 0xa5, 0x7d, 0x0d, 0x8e, 0xbc, 0x10, 0xaf, 0x85, 0x3c, 0xce, 0x96,
	0x0b, 0xbb, 0xdd, 0x5b, 0xbe, 0x99, 0xef, 0xb3, 0x3d, 0xbf, 0x51, 0x60, 0xb1, 0xb3, 0x37, 0x5c,
	0x74, 0xc5, 0xa8, 0xa4, 0x91, 0xe4, 0xa2, 0x63, 0xfc, 0x27, 0x17, 0x85, 0x56, 0x87, 0x42, 0xaa,
	0xa6, 0x18, 0x15, 0x3f, 0xd4, 0x86, 0x2d, 0x17, 0x7b, 0x39, 0x0c, 0x52, 0x78, 0x57, 0xfe, 0x2b,
	0x84, 0xe4, 0x13, 0xc6, 0xc8, 0x12, 0xce, 0xb5, 0xec, 0xad, 0xe1, 0x52, 0x64, 0xc1, 0x2a, 0x58,
	0xcf, 0xe9, 0x51, 0x93, 0x67, 0x10, 0xef, 0xa5, 0x15, 0x26, 0x0b, 0x57, 0xc1, 0x3a, 0xa6, 0x5e,
	0xb8, 0xaa, 0x91, 0xa6, 0xee, 0xb3, 0x68, 0x15, 0xac, 0x43, 0xea, 0x85, 0x3b, 0xa7, 0xe1, 0xda,
	0xdb, 0xcf, 0xb0, 0x71, 0xd4, 0x24, 0x83, 0x59, 0xab, 0x18, 0xef, 0x7e, 0x98, 0x2c, 0xc6, 0xd6,
	0xad, 0x24, 0x17, 0x90, 0xd4, 0x03, 0x66, 0x12, 0x6c, 0x4c, 0x8a, 0x7c, 0x84, 0x98, 0x1b, 0x36,
	0xe8, 0x6c, 0xb6, 0x8a, 0xd6, 0x69, 0x99, 0x17, 0xff, 0x1f, 0xab, 0xf0, 0x43, 0x54, 0x86, 0x0d,
	0xd4, 0x07, 0x5c, 0xb2, 0xe3, 0xad, 0xd1, 0xd9, 0xfc, 0xe1, 0x49, 0x0c, 0xe4, 0x7f, 0x02, 0x80,
	0x7f, 0x55, 0xf2, 0x02, 0x66, 0xee, 0xc4, 0x2d, 0x6f, 0x90, 0x4b, 0x44, 0x13, 0x27, 0xab, 0x86,
	0x3c, 0x87, 0x44, 0x5f, 0x5b, 0x57, 0x0f, 0xb1, 0x1e, 0xeb, 0x6b, 0x5b, 0x35, 0xe4, 0x29, 0x44,
	0xc2, 0x0e, 0x08, 0x25, 0xa6, 0xee, 0xd3, 0x81, 0x1a, 0x15, 0xdf, 0xb3, 0x89, 0x87, 0x17, 0xe4,
	0x12, 0x16, 0x52, 0xf1, 0x8e, 0x8b, 0xad, 0x6f, 0x7a, 0x22, 0xa9, 0xaf, 0x5d, 0xa1, 0xe5, 0x25,
	0xcc, 0xb5, 0xdd, 0x6d, 0x3d, 0x65, 0x0f, 0xe6, 0x5c, 0xdb, 0xdd, 0x37, 0x04, 0x7d, 0x09, 0x8b,
	0x51, 0xc9, 0x41, 0xba, 0x0d, 0xb9, 0x47, 0xcc, 0xf0, 0x11, 0xe9, 0xb1, 0x56, 0x35, 0xf9, 0xef,
	0x10, 0x9e, 0xf8, 0x49, 0x28, 0xd3, 0xa3, 0x14, 0x9a, 0x91, 0x0f, 0x90, 0x30, 0x61, 0xb8, 0xb9,
	0xc1, 0x61, 0xd2, 0xf2, 0xcd, 0xfd, 0x5c, 0xe8, 0xe4, 0x26, 0x1b, 0x88, 0xc7, 0xba, 0x63, 0x0a,
	0x67, 0x4d, 0xcb, 0xd7, 0x77, 0xc5, 0xae, 0x9c, 0x89, 0x7a, 0x2f, 0x79, 0x7f, 0xbb, 0xbd, 0x08,
	0x77, 0x70, 0xea, 0xae, 0x69, 0x73, 0x1b, 0x88, 0x99, 0x52, 0x52, 0x21, 0xae, 0x7b, 0xae, 0xfa,
	0xe2, 0x4c, 0xd4, 0x7b, 0xc9, 0x3b, 0x38, 0xe3, 0xa2, 0x95, 0x48, 0x31, 0x2d, 0x5f, 0xdd, 0x95,
	0xa9, 0x44, 0x2b, 0x29, 0x3a, 0xcb, 0x16, 0x1e, 0xfb, 0x7b, 0xbf, 0x32, 0x75, 0x70, 0xb4, 0xbf,
	0xc3, 0xfc, 0x73, 0xdd, 0xef, 0x6d, 0x5f, 0x1b, 0x46, 0x4e, 0xbc, 0x75, 0xf9, 0xf6, 0xc4, 0x2c,
	0x13, 0xef, 0xfc, 0xd1, 0x2e, 0xc1, 0x5f, 0x6d, 0xf3, 0x37, 0x00, 0x00, 0xff, 0xff, 0x75, 0x26,
	0x22, 0xcc, 0xa0, 0x03, 0x00, 0x00,
}
