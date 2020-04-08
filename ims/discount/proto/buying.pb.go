// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buying.proto

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

//购买清单
type Buying struct {
	Solution             string        `protobuf:"bytes,1,opt,name=solution,proto3" json:"solution,omitempty"`
	Count                int32         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Total                float32       `protobuf:"fixed32,3,opt,name=total,proto3" json:"total,omitempty"`
	Discount             float32       `protobuf:"fixed32,4,opt,name=discount,proto3" json:"discount,omitempty"`
	Freight              float32       `protobuf:"fixed32,5,opt,name=freight,proto3" json:"freight,omitempty"`
	Amount               float32       `protobuf:"fixed32,6,opt,name=amount,proto3" json:"amount,omitempty"`
	AddressId            int64         `protobuf:"varint,7,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	CustomerId           int64         `protobuf:"varint,8,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	UseTicketId          int64         `protobuf:"varint,9,opt,name=use_ticket_id,json=useTicketId,proto3" json:"use_ticket_id,omitempty"`
	AvailableTicketIds   []int64       `protobuf:"varint,10,rep,packed,name=available_ticket_ids,json=availableTicketIds,proto3" json:"available_ticket_ids,omitempty"`
	Items                []*BuyingItem `protobuf:"bytes,11,rep,name=items,proto3" json:"items,omitempty"`
	Changed              bool          `protobuf:"varint,12,opt,name=changed,proto3" json:"changed,omitempty"`
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

func (m *Buying) GetAddressId() int64 {
	if m != nil {
		return m.AddressId
	}
	return 0
}

func (m *Buying) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *Buying) GetUseTicketId() int64 {
	if m != nil {
		return m.UseTicketId
	}
	return 0
}

func (m *Buying) GetAvailableTicketIds() []int64 {
	if m != nil {
		return m.AvailableTicketIds
	}
	return nil
}

func (m *Buying) GetItems() []*BuyingItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Buying) GetChanged() bool {
	if m != nil {
		return m.Changed
	}
	return false
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
	proto.RegisterType((*Buying)(nil), "geiqin.srv.ims.discount.Buying")
	proto.RegisterType((*BuyingItem)(nil), "geiqin.srv.ims.discount.BuyingItem")
	proto.RegisterType((*BuyingResponse)(nil), "geiqin.srv.ims.discount.BuyingResponse")
}

func init() {
	proto.RegisterFile("buying.proto", fileDescriptor_fc0a9e5c6a9833d6)
}

var fileDescriptor_fc0a9e5c6a9833d6 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x1c, 0xd7, 0x6e, 0x72, 0x9d, 0x7e, 0x42, 0xa3, 0x42, 0xad, 0xa0, 0x52, 0x37, 0x2c,
	0xf0, 0xca, 0x82, 0x50, 0x84, 0xd8, 0x82, 0x58, 0x78, 0x57, 0x0d, 0x95, 0x58, 0x46, 0x8e, 0x67,
	0xe2, 0x8c, 0x62, 0xcf, 0x84, 0xf9, 0x89, 0xd4, 0x47, 0xe0, 0x29, 0x78, 0x1c, 0x5e, 0x0b, 0xcd,
	0x8c, 0x9d, 0xb2, 0x09, 0xd9, 0xe5, 0x9c, 0x7b, 0x4e, 0xee, 0xdc, 0x73, 0x12, 0x98, 0xae, 0xcc,
	0x23, 0xe3, 0x4d, 0xb1, 0x93, 0x42, 0x0b, 0x74, 0xd5, 0x50, 0xf6, 0x83, 0xf1, 0x42, 0xc9, 0x7d,
	0xc1, 0x3a, 0x55, 0x10, 0xa6, 0x6a, 0x61, 0xb8, 0x9e, 0x4d, 0x6b, 0xd1, 0x75, 0x82, 0x7b, 0xd9,
	0xfc, 0x67, 0x08, 0xf1, 0x67, 0xe7, 0x43, 0x33, 0x18, 0x2b, 0xd1, 0x1a, 0xcd, 0x04, 0x4f, 0x83,
	0x2c, 0xc8, 0x27, 0xf8, 0x80, 0xd1, 0x25, 0x44, 0xce, 0x9d, 0x8e, 0xb2, 0x20, 0x8f, 0xb0, 0x07,
	0x96, 0xd5, 0x42, 0x57, 0x6d, 0x1a, 0x66, 0x41, 0x3e, 0xc2, 0x1e, 0xd8, 0xef, 0x19, 0x96, 0xa5,
	0x67, 0x6e, 0x70, 0xc0, 0x28, 0x85, 0xf3, 0xb5, 0xa4, 0xac, 0xd9, 0xe8, 0x34, 0x72, 0xa3, 0x01,
	0xa2, 0x17, 0x10, 0x57, 0x9d, 0xf3, 0xc4, 0x6e, 0xd0, 0x23, 0x74, 0x0d, 0x50, 0x11, 0x22, 0xa9,
	0x52, 0x4b, 0x46, 0xd2, 0xf3, 0x2c, 0xc8, 0x43, 0x3c, 0xe9, 0x99, 0x92, 0xa0, 0x1b, 0x48, 0x6a,
	0xa3, 0xb4, 0xe8, 0xa8, 0xb4, 0xf3, 0xb1, 0x9b, 0xc3, 0x40, 0x95, 0x04, 0xcd, 0xe1, 0xc2, 0x28,
	0xba, 0xd4, 0xac, 0xde, 0x52, 0x6d, 0x25, 0x13, 0x27, 0x49, 0x8c, 0xa2, 0x0f, 0x8e, 0x2b, 0x09,
	0x7a, 0x0b, 0x97, 0xd5, 0xbe, 0x62, 0x6d, 0xb5, 0x6a, 0xff, 0x52, 0xaa, 0x14, 0xb2, 0x30, 0x0f,
	0x31, 0x3a, 0xcc, 0x06, 0x83, 0x42, 0x9f, 0x20, 0x62, 0x9a, 0x76, 0x2a, 0x4d, 0xb2, 0x30, 0x4f,
	0x16, 0xaf, 0x8b, 0x23, 0x69, 0x17, 0x3e, 0xdb, 0x52, 0xd3, 0x0e, 0x7b, 0x87, 0x8d, 0xa0, 0xde,
	0x54, 0xbc, 0xa1, 0x24, 0x9d, 0x66, 0x41, 0x3e, 0xc6, 0x03, 0x9c, 0xff, 0x0e, 0x00, 0x9e, 0xf4,
	0xe8, 0x0a, 0xce, 0xad, 0xc3, 0xbe, 0x39, 0x70, 0x6f, 0x8e, 0x2d, 0x2c, 0x09, 0x7a, 0x0e, 0xb1,
	0xda, 0x1a, 0xcb, 0x8f, 0x1c, 0x1f, 0xa9, 0xad, 0x29, 0x09, 0x7a, 0x06, 0x21, 0x37, 0x9d, 0xeb,
	0x22, 0xc2, 0xf6, 0xa3, 0xed, 0x67, 0x27, 0x59, 0x4d, 0xfb, 0x1a, 0x3c, 0x40, 0xb7, 0x30, 0x15,
	0x92, 0x35, 0x8c, 0x2f, 0xfd, 0xd0, 0x17, 0x91, 0x78, 0xee, 0xde, 0x49, 0x5e, 0xc2, 0x44, 0x99,
	0xd5, 0xd2, 0x97, 0xeb, 0xfb, 0x18, 0x2b, 0xb3, 0x7a, 0x70, 0xfd, 0xde, 0xc2, 0x74, 0x27, 0x45,
	0x27, 0xec, 0x0f, 0xe3, 0xa9, 0x93, 0xe4, 0xc0, 0x95, 0x64, 0xfe, 0x6b, 0x04, 0xff, 0xfb, 0x4b,
	0x30, 0x55, 0x3b, 0xc1, 0x15, 0x45, 0x1f, 0x21, 0xa6, 0x5c, 0x33, 0xfd, 0xe8, 0x8e, 0x49, 0x16,
	0x37, 0x27, 0x22, 0xc3, 0xbd, 0x1c, 0xdd, 0x41, 0xb4, 0xab, 0x1a, 0x2a, 0xdd, 0xb1, 0xc9, 0xe2,
	0xd5, 0x51, 0xdf, 0xbd, 0x55, 0x61, 0x2f, 0x46, 0x1f, 0x86, 0x82, 0x42, 0x57, 0xd0, 0xc9, 0x6d,
	0x7d, 0x39, 0x77, 0x10, 0x51, 0x29, 0x85, 0x74, 0x89, 0xfd, 0x6b, 0xd9, 0x57, 0xab, 0xc2, 0x5e,
	0x8c, 0xde, 0xc1, 0x19, 0xe3, 0x6b, 0xe1, 0x92, 0x4c, 0x16, 0xd7, 0x47, 0x4d, 0x25, 0x5f, 0x0b,
	0xec, 0xa4, 0x8b, 0x0d, 0x5c, 0xf8, 0xcd, 0xdf, 0xa8, 0xdc, 0xdb, 0xc8, 0xbf, 0xc3, 0xe4, 0x4b,
	0xd5, 0xd6, 0xa6, 0xad, 0x34, 0x45, 0xa7, 0x9e, 0x3b, 0x7b, 0x73, 0xea, 0x9e, 0x3e, 0xf6, 0xf9,
	0x7f, 0xab, 0xd8, 0xfd, 0xd1, 0xdf, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x1e, 0xf3, 0xda,
	0x1f, 0x04, 0x00, 0x00,
}
