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
	Solution             string         `protobuf:"bytes,1,opt,name=solution,proto3" json:"solution"`
	Changed              bool           `protobuf:"varint,2,opt,name=changed,proto3" json:"changed"`
	Count                int32          `protobuf:"varint,3,opt,name=count,proto3" json:"count"`
	Total                float32        `protobuf:"fixed32,4,opt,name=total,proto3" json:"total"`
	Discount             float32        `protobuf:"fixed32,5,opt,name=discount,proto3" json:"discount"`
	Freight              float32        `protobuf:"fixed32,6,opt,name=freight,proto3" json:"freight"`
	Amount               float32        `protobuf:"fixed32,7,opt,name=amount,proto3" json:"amount"`
	AddressId            int64          `protobuf:"varint,8,opt,name=address_id,json=addressId,proto3" json:"address_id"`
	CustomerId           int64          `protobuf:"varint,9,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	UseTicketId          int64          `protobuf:"varint,10,opt,name=use_ticket_id,json=useTicketId,proto3" json:"use_ticket_id"`
	Message              string         `protobuf:"bytes,11,opt,name=message,proto3" json:"message"`
	AvailableTicketIds   []int64        `protobuf:"varint,12,rep,packed,name=available_ticket_ids,json=availableTicketIds,proto3" json:"available_ticket_ids"`
	Items                []*BuyingItem  `protobuf:"bytes,13,rep,name=items,proto3" json:"items"`
	CartRowIds           []string       `protobuf:"bytes,14,rep,name=cart_row_ids,json=cartRowIds,proto3" json:"cart_row_ids"`
	Coupons              []*OrderCoupon `protobuf:"bytes,15,rep,name=coupons,proto3" json:"coupons"`
	Address              *OrderAddress  `protobuf:"bytes,16,opt,name=address,proto3" json:"address"`
	UseTicket            *CouponTicket  `protobuf:"bytes,17,opt,name=use_ticket,json=useTicket,proto3" json:"use_ticket"`
	PlatformSource       string         `protobuf:"bytes,18,opt,name=platform_source,json=platformSource,proto3" json:"platform_source"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
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

func (m *Buying) GetChanged() bool {
	if m != nil {
		return m.Changed
	}
	return false
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

func (m *Buying) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
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

func (m *Buying) GetCartRowIds() []string {
	if m != nil {
		return m.CartRowIds
	}
	return nil
}

func (m *Buying) GetCoupons() []*OrderCoupon {
	if m != nil {
		return m.Coupons
	}
	return nil
}

func (m *Buying) GetAddress() *OrderAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Buying) GetUseTicket() *CouponTicket {
	if m != nil {
		return m.UseTicket
	}
	return nil
}

func (m *Buying) GetPlatformSource() string {
	if m != nil {
		return m.PlatformSource
	}
	return ""
}

//购买清单明细
type BuyingItem struct {
	ItemId               int64      `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64      `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Num                  int32      `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
	Price                float32    `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	OriginPrice          float32    `protobuf:"fixed32,5,opt,name=origin_price,json=originPrice,proto3" json:"origin_price,omitempty"`
	SubTotal             float32    `protobuf:"fixed32,6,opt,name=sub_total,json=subTotal,proto3" json:"sub_total,omitempty"`
	IsGift               bool       `protobuf:"varint,7,opt,name=is_gift,json=isGift,proto3" json:"is_gift,omitempty"`
	PromotionId          int64      `protobuf:"varint,8,opt,name=promotion_id,json=promotionId,proto3" json:"promotion_id,omitempty"`
	Goods                *GoodsInfo `protobuf:"bytes,9,opt,name=goods,proto3" json:"goods,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
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

func (m *BuyingItem) GetIsGift() bool {
	if m != nil {
		return m.IsGift
	}
	return false
}

func (m *BuyingItem) GetPromotionId() int64 {
	if m != nil {
		return m.PromotionId
	}
	return 0
}

func (m *BuyingItem) GetGoods() *GoodsInfo {
	if m != nil {
		return m.Goods
	}
	return nil
}

//购买清单请求(订单下单使用)
type BuyingRequest struct {
	CustomerId           int64         `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Source               int32         `protobuf:"varint,2,opt,name=source,proto3" json:"source,omitempty"`
	AddressId            int64         `protobuf:"varint,3,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	Message              string        `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	VipcardId            int64         `protobuf:"varint,5,opt,name=vipcard_id,json=vipcardId,proto3" json:"vipcard_id,omitempty"`
	PayMethod            int32         `protobuf:"varint,6,opt,name=pay_method,json=payMethod,proto3" json:"pay_method,omitempty"`
	UseTicketId          int64         `protobuf:"varint,7,opt,name=use_ticket_id,json=useTicketId,proto3" json:"use_ticket_id,omitempty"`
	PlatformSource       string        `protobuf:"bytes,8,opt,name=platform_source,json=platformSource,proto3" json:"platform_source,omitempty"`
	Items                []*BuyingItem `protobuf:"bytes,9,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BuyingRequest) Reset()         { *m = BuyingRequest{} }
func (m *BuyingRequest) String() string { return proto.CompactTextString(m) }
func (*BuyingRequest) ProtoMessage()    {}
func (*BuyingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc0a9e5c6a9833d6, []int{2}
}

func (m *BuyingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyingRequest.Unmarshal(m, b)
}
func (m *BuyingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyingRequest.Marshal(b, m, deterministic)
}
func (m *BuyingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyingRequest.Merge(m, src)
}
func (m *BuyingRequest) XXX_Size() int {
	return xxx_messageInfo_BuyingRequest.Size(m)
}
func (m *BuyingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuyingRequest proto.InternalMessageInfo

func (m *BuyingRequest) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *BuyingRequest) GetSource() int32 {
	if m != nil {
		return m.Source
	}
	return 0
}

func (m *BuyingRequest) GetAddressId() int64 {
	if m != nil {
		return m.AddressId
	}
	return 0
}

func (m *BuyingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BuyingRequest) GetVipcardId() int64 {
	if m != nil {
		return m.VipcardId
	}
	return 0
}

func (m *BuyingRequest) GetPayMethod() int32 {
	if m != nil {
		return m.PayMethod
	}
	return 0
}

func (m *BuyingRequest) GetUseTicketId() int64 {
	if m != nil {
		return m.UseTicketId
	}
	return 0
}

func (m *BuyingRequest) GetPlatformSource() string {
	if m != nil {
		return m.PlatformSource
	}
	return ""
}

func (m *BuyingRequest) GetItems() []*BuyingItem {
	if m != nil {
		return m.Items
	}
	return nil
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
	return fileDescriptor_fc0a9e5c6a9833d6, []int{3}
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
	proto.RegisterType((*BuyingRequest)(nil), "geiqin.srv.ord.private.BuyingRequest")
	proto.RegisterType((*BuyingResponse)(nil), "geiqin.srv.ord.private.BuyingResponse")
}

func init() {
	proto.RegisterFile("buying.proto", fileDescriptor_fc0a9e5c6a9833d6)
}

var fileDescriptor_fc0a9e5c6a9833d6 = []byte{
	// 805 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x26, 0x71, 0xed, 0xc4, 0xc7, 0x69, 0xbb, 0x3b, 0x5a, 0x8a, 0x55, 0x28, 0x78, 0x03, 0x02,
	0x5f, 0x45, 0xab, 0x14, 0x01, 0x37, 0x20, 0x41, 0x85, 0x56, 0xb9, 0x40, 0xac, 0x66, 0x97, 0xeb,
	0x68, 0xe2, 0x99, 0xb8, 0xa3, 0xc6, 0x1e, 0xef, 0xcc, 0x38, 0xab, 0x3e, 0x0a, 0xbc, 0x05, 0x2f,
	0xc1, 0x73, 0xa1, 0x39, 0x63, 0x27, 0xa5, 0xbf, 0xe2, 0xae, 0xe7, 0x3b, 0x3f, 0x73, 0xf2, 0x9d,
	0xef, 0x73, 0x61, 0xb2, 0x6a, 0xaf, 0x65, 0x5d, 0xce, 0x1a, 0xad, 0xac, 0x22, 0x27, 0xa5, 0x90,
	0xef, 0x65, 0x3d, 0x33, 0x7a, 0x3b, 0x53, 0x9a, 0xcf, 0x1a, 0x2d, 0xb7, 0xcc, 0x8a, 0xd3, 0x49,
	0xa1, 0xaa, 0x4a, 0xd5, 0xbe, 0xea, 0xf4, 0xb8, 0x54, 0x8a, 0x9b, 0x45, 0xbd, 0x56, 0x1d, 0x40,
	0x94, 0xe6, 0x42, 0xff, 0xcc, 0xb9, 0x16, 0xc6, 0x74, 0xd8, 0x73, 0xc4, 0x2e, 0x54, 0xdb, 0xec,
	0xfa, 0x4e, 0x0a, 0x8c, 0xde, 0xc9, 0xe2, 0x4a, 0xd8, 0x7d, 0xfb, 0xf4, 0xef, 0x10, 0xa2, 0x5f,
	0x70, 0x0d, 0x72, 0x0a, 0x63, 0xa3, 0x36, 0xad, 0x95, 0xaa, 0x4e, 0x07, 0xd9, 0x20, 0x8f, 0xe9,
	0x2e, 0x26, 0x29, 0x8c, 0x8a, 0x4b, 0x56, 0x97, 0x82, 0xa7, 0xc3, 0x6c, 0x90, 0x8f, 0x69, 0x1f,
	0x92, 0x17, 0x10, 0x16, 0xaa, 0xad, 0x6d, 0x1a, 0x64, 0x83, 0x3c, 0xa4, 0x3e, 0x70, 0xa8, 0x55,
	0x96, 0x6d, 0xd2, 0x83, 0x6c, 0x90, 0x0f, 0xa9, 0x0f, 0xdc, 0x0b, 0x5c, 0x1a, 0x5f, 0x1e, 0x62,
	0x62, 0x17, 0xbb, 0x17, 0xd6, 0x5a, 0xc8, 0xf2, 0xd2, 0xa6, 0x11, 0xa6, 0xfa, 0x90, 0x9c, 0x40,
	0xc4, 0x2a, 0xec, 0x19, 0x61, 0xa2, 0x8b, 0xc8, 0x19, 0x00, 0xf3, 0x3f, 0x7b, 0x29, 0x79, 0x3a,
	0xce, 0x06, 0x79, 0x40, 0xe3, 0x0e, 0x59, 0x70, 0xf2, 0x05, 0x24, 0x45, 0x6b, 0xac, 0xaa, 0x84,
	0x76, 0xf9, 0x18, 0xf3, 0xd0, 0x43, 0x0b, 0x4e, 0xa6, 0x70, 0xd8, 0x1a, 0xb1, 0xb4, 0x48, 0x89,
	0x2b, 0x01, 0x2c, 0x49, 0x5a, 0x23, 0x3a, 0x9a, 0xb8, 0xdb, 0xaa, 0x12, 0xc6, 0xb0, 0x52, 0xa4,
	0x09, 0x52, 0xd2, 0x87, 0xe4, 0x15, 0xbc, 0x60, 0x5b, 0x26, 0x37, 0x6c, 0xb5, 0xb9, 0x31, 0xc3,
	0xa4, 0x93, 0x2c, 0xc8, 0x03, 0x4a, 0x76, 0xb9, 0x7e, 0x94, 0x21, 0x3f, 0x40, 0x28, 0xad, 0xa8,
	0x4c, 0x7a, 0x98, 0x05, 0x79, 0x32, 0x9f, 0xce, 0xee, 0x3f, 0xf8, 0xcc, 0x9f, 0x63, 0x61, 0x45,
	0x45, 0x7d, 0x03, 0xc9, 0x60, 0x52, 0x30, 0x6d, 0x97, 0x5a, 0x7d, 0xc0, 0x37, 0x8e, 0xb2, 0x20,
	0x8f, 0x29, 0x38, 0x8c, 0xaa, 0x0f, 0x6e, 0xf6, 0x8f, 0x30, 0xf2, 0x07, 0x36, 0xe9, 0x31, 0x4e,
	0xff, 0xf2, 0xa1, 0xe9, 0xbf, 0xef, 0xa5, 0x41, 0xfb, 0x1e, 0xf2, 0x13, 0x8c, 0x3a, 0xe2, 0xd2,
	0x67, 0xd9, 0x20, 0x4f, 0xe6, 0x5f, 0x3d, 0xda, 0xde, 0xa9, 0x8d, 0xf6, 0x4d, 0xe4, 0x02, 0x60,
	0x4f, 0x65, 0xfa, 0xfc, 0xf1, 0x11, 0x17, 0x37, 0x94, 0x48, 0xe3, 0x1d, 0xdb, 0xe4, 0x1b, 0x38,
	0x6e, 0x36, 0xcc, 0xae, 0x95, 0xae, 0x96, 0x46, 0xb5, 0xba, 0x10, 0x29, 0x41, 0xce, 0x8f, 0x7a,
	0xf8, 0x2d, 0xa2, 0xd3, 0x3f, 0x87, 0x00, 0x7b, 0x92, 0xc8, 0x27, 0x30, 0x72, 0x34, 0xb9, 0x0b,
	0x0e, 0xf0, 0x82, 0x91, 0x0b, 0x17, 0x9c, 0x7c, 0x0c, 0x91, 0xb9, 0x6a, 0x1d, 0x3e, 0x44, 0x3c,
	0x34, 0x57, 0xed, 0x82, 0x93, 0x67, 0x10, 0xd4, 0x6d, 0xd5, 0xe9, 0xd5, 0xfd, 0xe9, 0xd4, 0xda,
	0x68, 0x59, 0x88, 0x5e, 0xad, 0x18, 0x90, 0x97, 0x30, 0x51, 0x5a, 0x96, 0xb2, 0x5e, 0xfa, 0xa4,
	0x57, 0x6c, 0xe2, 0xb1, 0x37, 0x58, 0xf2, 0x29, 0xc4, 0xa6, 0x5d, 0x2d, 0xbd, 0xd4, 0xbd, 0x6c,
	0xc7, 0xa6, 0x5d, 0xbd, 0x43, 0xb5, 0xbb, 0xbd, 0xcc, 0xb2, 0x94, 0x6b, 0x2f, 0xdc, 0x31, 0x8d,
	0xa4, 0x79, 0x2d, 0xd7, 0xd6, 0x0d, 0x6e, 0xb4, 0xaa, 0x94, 0x73, 0xd6, 0x5e, 0xba, 0xc9, 0x0e,
	0x5b, 0x70, 0xf2, 0x3d, 0x84, 0x68, 0x74, 0x94, 0x6d, 0x32, 0x7f, 0xf9, 0x10, 0x97, 0xaf, 0xfb,
	0xaf, 0x01, 0xf5, 0xf5, 0xd3, 0x7f, 0x86, 0x70, 0xe8, 0xb9, 0xa1, 0xe2, 0x7d, 0x2b, 0x8c, 0xbd,
	0xed, 0x83, 0xc1, 0x1d, 0x1f, 0x9c, 0x40, 0xd4, 0xd1, 0x3d, 0x44, 0x4a, 0xba, 0xe8, 0x96, 0xbf,
	0x82, 0xdb, 0xfe, 0xba, 0x61, 0x8d, 0x83, 0xff, 0x5a, 0xe3, 0x0c, 0x60, 0x2b, 0x9b, 0x82, 0x69,
	0xee, 0x1a, 0x43, 0xdf, 0xd8, 0x21, 0x0b, 0xee, 0xd2, 0x0d, 0xbb, 0x5e, 0x56, 0xc2, 0x5e, 0x2a,
	0x8e, 0xac, 0x85, 0x34, 0x6e, 0xd8, 0xf5, 0x6f, 0x08, 0xdc, 0xb5, 0xe5, 0xe8, 0xae, 0x2d, 0xef,
	0x91, 0xca, 0xf8, 0x3e, 0xa9, 0xec, 0x3d, 0x17, 0xff, 0x4f, 0xcf, 0x4d, 0xff, 0x1a, 0xc2, 0x51,
	0x4f, 0xa4, 0x71, 0x26, 0x11, 0xe4, 0x3b, 0x88, 0x44, 0x6d, 0xa5, 0xbd, 0x46, 0x12, 0x93, 0xf9,
	0xe7, 0x8f, 0x4f, 0xa3, 0x5d, 0x35, 0x39, 0x87, 0xb0, 0x61, 0xa5, 0xd0, 0xc8, 0x6f, 0x32, 0x3f,
	0x7b, 0xa8, 0xed, 0x8d, 0x2b, 0xa2, 0xbe, 0x96, 0x7c, 0xdb, 0x6f, 0x1e, 0xe0, 0xe6, 0x4f, 0xbd,
	0xd5, 0x7d, 0x29, 0xce, 0x21, 0x14, 0x5a, 0x2b, 0x8d, 0x27, 0x79, 0xe4, 0xa9, 0x5f, 0x5d, 0x11,
	0xf5, 0xb5, 0xe4, 0x15, 0x1c, 0xc8, 0x7a, 0xad, 0xf0, 0x52, 0xc9, 0xfc, 0xb3, 0x87, 0x7a, 0x50,
	0x66, 0x58, 0x39, 0x5f, 0xf7, 0x22, 0x7b, 0x2b, 0xf4, 0xd6, 0x19, 0xe1, 0x0f, 0x88, 0x2f, 0xd8,
	0xa6, 0x68, 0x37, 0xcc, 0x0a, 0xf2, 0xc4, 0xae, 0xa7, 0x5f, 0x3f, 0xf1, 0x5b, 0x3a, 0xbe, 0xa7,
	0x1f, 0xad, 0x22, 0xfc, 0x27, 0x75, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x34, 0x76,
	0xaf, 0x2a, 0x07, 0x00, 0x00,
}
