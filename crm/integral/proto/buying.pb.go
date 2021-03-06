// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buying.proto

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

//购买清单
type Buying struct {
	Total                float32       `protobuf:"fixed32,1,opt,name=total,proto3" json:"total,omitempty"`
	Count                int32         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Freight              float32       `protobuf:"fixed32,3,opt,name=freight,proto3" json:"freight,omitempty"`
	Amount               float32       `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Points               int32         `protobuf:"varint,5,opt,name=points,proto3" json:"points,omitempty"`
	CashExchangeMoney    float32       `protobuf:"fixed32,6,opt,name=cash_exchange_money,json=cashExchangeMoney,proto3" json:"cash_exchange_money,omitempty"`
	CashExchangePoints   int32         `protobuf:"varint,7,opt,name=cash_exchange_points,json=cashExchangePoints,proto3" json:"cash_exchange_points,omitempty"`
	IntegralName         string        `protobuf:"bytes,8,opt,name=integral_name,json=integralName,proto3" json:"integral_name,omitempty"`
	Items                []*BuyingItem `protobuf:"bytes,9,rep,name=items,proto3" json:"items,omitempty"`
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

func (m *Buying) GetTotal() float32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Buying) GetCount() int32 {
	if m != nil {
		return m.Count
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

func (m *Buying) GetPoints() int32 {
	if m != nil {
		return m.Points
	}
	return 0
}

func (m *Buying) GetCashExchangeMoney() float32 {
	if m != nil {
		return m.CashExchangeMoney
	}
	return 0
}

func (m *Buying) GetCashExchangePoints() int32 {
	if m != nil {
		return m.CashExchangePoints
	}
	return 0
}

func (m *Buying) GetIntegralName() string {
	if m != nil {
		return m.IntegralName
	}
	return ""
}

func (m *Buying) GetItems() []*BuyingItem {
	if m != nil {
		return m.Items
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
	proto.RegisterType((*Buying)(nil), "geiqin.srv.crm.integral.Buying")
	proto.RegisterType((*BuyingItem)(nil), "geiqin.srv.crm.integral.BuyingItem")
	proto.RegisterType((*BuyingResponse)(nil), "geiqin.srv.crm.integral.BuyingResponse")
}

func init() {
	proto.RegisterFile("buying.proto", fileDescriptor_fc0a9e5c6a9833d6)
}

var fileDescriptor_fc0a9e5c6a9833d6 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x8e, 0xd3, 0x3c,
	0x14, 0x85, 0xff, 0x26, 0x93, 0x74, 0x7a, 0xd3, 0xf9, 0x05, 0x66, 0x60, 0xa2, 0x41, 0x40, 0xe9,
	0x2c, 0xe8, 0x2a, 0x82, 0x00, 0x42, 0x6c, 0x41, 0xb3, 0xe8, 0x02, 0x54, 0x19, 0x24, 0x96, 0x91,
	0x9b, 0xba, 0xa9, 0x35, 0xb5, 0x1d, 0x6c, 0x67, 0x44, 0x9f, 0x86, 0x0d, 0x0f, 0xc2, 0xa3, 0x21,
	0x5f, 0xa7, 0x02, 0x16, 0xa5, 0xbb, 0x9e, 0x73, 0xcf, 0xf1, 0xb5, 0xbf, 0x2a, 0x30, 0x5e, 0x76,
	0x3b, 0xa1, 0x9a, 0xa2, 0x35, 0xda, 0x69, 0x72, 0xd1, 0x70, 0xf1, 0x55, 0xa8, 0xc2, 0x9a, 0xdb,
	0xa2, 0x36, 0xb2, 0x10, 0xca, 0xf1, 0xc6, 0xb0, 0xed, 0xe5, 0xb8, 0xd6, 0x52, 0x6a, 0x15, 0x62,
	0xd3, 0x9f, 0x11, 0xa4, 0xef, 0xb0, 0x47, 0xce, 0x21, 0x71, 0xda, 0xb1, 0x6d, 0x3e, 0x98, 0x0c,
	0x66, 0x11, 0x0d, 0xc2, 0xbb, 0xb5, 0xee, 0x94, 0xcb, 0xa3, 0xc9, 0x60, 0x96, 0xd0, 0x20, 0x48,
	0x0e, 0xc3, 0xb5, 0xe1, 0xa2, 0xd9, 0xb8, 0x3c, 0xc6, 0xf4, 0x5e, 0x92, 0x07, 0x90, 0x32, 0x89,
	0x85, 0x13, 0x1c, 0xf4, 0xca, 0xfb, 0xad, 0x16, 0xca, 0xd9, 0x3c, 0xc1, 0x83, 0x7a, 0x45, 0x0a,
	0xb8, 0x57, 0x33, 0xbb, 0xa9, 0xf8, 0xb7, 0x7a, 0xc3, 0x54, 0xc3, 0x2b, 0xa9, 0x15, 0xdf, 0xe5,
	0x29, 0x96, 0xef, 0xfa, 0xd1, 0x75, 0x3f, 0xf9, 0xe0, 0x07, 0xe4, 0x39, 0x9c, 0xff, 0x9d, 0xef,
	0x4f, 0x1d, 0xe2, 0xa9, 0xe4, 0xcf, 0xc2, 0x22, 0x6c, 0xb8, 0x82, 0xb3, 0xfd, 0xe3, 0x2b, 0xc5,
	0x24, 0xcf, 0x4f, 0x27, 0x83, 0xd9, 0x88, 0x8e, 0xf7, 0xe6, 0x47, 0x26, 0x39, 0x79, 0x0b, 0x89,
	0x70, 0x5c, 0xda, 0x7c, 0x34, 0x89, 0x67, 0x59, 0x79, 0x55, 0x1c, 0xc0, 0x57, 0x04, 0x58, 0x73,
	0xc7, 0x25, 0x0d, 0x8d, 0xe9, 0x8f, 0x01, 0xc0, 0x6f, 0x97, 0x5c, 0xc0, 0xd0, 0xfb, 0x95, 0x58,
	0x21, 0xc8, 0x98, 0xa6, 0x5e, 0xce, 0x57, 0xe4, 0x3e, 0xa4, 0xf6, 0xa6, 0xf3, 0x7e, 0x84, 0x7e,
	0x62, 0x6f, 0xba, 0xf9, 0x8a, 0xdc, 0x81, 0x58, 0x75, 0x12, 0x31, 0x26, 0xd4, 0xff, 0xf4, 0xc8,
	0x5b, 0x23, 0x6a, 0xde, 0x13, 0x0c, 0x82, 0x3c, 0x85, 0xb1, 0x36, 0xa2, 0x11, 0xaa, 0x0a, 0xc3,
	0x04, 0x87, 0x59, 0xf0, 0x16, 0x18, 0x79, 0x08, 0x23, 0xdb, 0x2d, 0xab, 0xf0, 0x2f, 0x06, 0x82,
	0xa7, 0xb6, 0x5b, 0x7e, 0xf6, 0x7a, 0xfa, 0x3d, 0x82, 0xff, 0xc3, 0x35, 0x29, 0xb7, 0xad, 0x56,
	0x96, 0x93, 0x37, 0x90, 0x72, 0xe5, 0x84, 0xdb, 0xe1, 0x4d, 0xb3, 0xf2, 0xc9, 0x91, 0x57, 0xd3,
	0x3e, 0x4e, 0x5e, 0x41, 0xd2, 0xb2, 0x86, 0x1b, 0x7c, 0x49, 0x56, 0x3e, 0x3e, 0xd8, 0x5b, 0xf8,
	0x14, 0x0d, 0x61, 0xf2, 0x7a, 0xcf, 0x38, 0x46, 0xc6, 0x47, 0xb7, 0x85, 0xb4, 0x5f, 0xc6, 0x8d,
	0xd1, 0x06, 0x71, 0xfc, 0x6b, 0xd9, 0xb5, 0x4f, 0xd1, 0x10, 0x26, 0x2f, 0xe0, 0x44, 0xa8, 0xb5,
	0x46, 0x4c, 0x59, 0xf9, 0xe8, 0x60, 0x69, 0xae, 0xd6, 0x9a, 0x62, 0xb4, 0xdc, 0xc0, 0x59, 0xd8,
	0xfc, 0x89, 0x9b, 0x5b, 0xcf, 0xf3, 0x0b, 0x8c, 0xde, 0xb3, 0x6d, 0xdd, 0x6d, 0x99, 0xe3, 0xe4,
	0xd8, 0x75, 0x2f, 0x9f, 0x1d, 0x7b, 0x4f, 0x8f, 0x7d, 0xfa, 0xdf, 0x32, 0xc5, 0x8f, 0xef, 0xe5,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0x8b, 0x6e, 0x5e, 0xb3, 0x03, 0x00, 0x00,
}
