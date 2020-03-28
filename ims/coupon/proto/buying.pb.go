// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buying.proto

package geiqin_srv_ims_coupon

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
	Solution             string        `protobuf:"bytes,1,opt,name=solution,proto3" json:"solution"`
	Count                int32         `protobuf:"varint,2,opt,name=count,proto3" json:"count"`
	Total                float32       `protobuf:"fixed32,3,opt,name=total,proto3" json:"total"`
	Discount             float32       `protobuf:"fixed32,4,opt,name=discount,proto3" json:"discount"`
	Freight              float32       `protobuf:"fixed32,5,opt,name=freight,proto3" json:"freight"`
	Amount               float32       `protobuf:"fixed32,6,opt,name=amount,proto3" json:"amount"`
	AddressId            int64         `protobuf:"varint,7,opt,name=address_id,json=addressId,proto3" json:"address_id"`
	CustomerId           int64         `protobuf:"varint,8,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	UseTicketId          int64         `protobuf:"varint,9,opt,name=use_ticket_id,json=useTicketId,proto3" json:"use_ticket_id"`
	AvailableTicketIds   []int64       `protobuf:"varint,10,rep,packed,name=available_ticket_ids,json=availableTicketIds,proto3" json:"available_ticket_ids"`
	Items                []*BuyingItem `protobuf:"bytes,11,rep,name=items,proto3" json:"items"`
	Changed              bool          `protobuf:"varint,12,opt,name=changed,proto3" json:"changed"`
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
	proto.RegisterType((*Buying)(nil), "geiqin.srv.ims.coupon.Buying")
	proto.RegisterType((*BuyingItem)(nil), "geiqin.srv.ims.coupon.BuyingItem")
	proto.RegisterType((*BuyingResponse)(nil), "geiqin.srv.ims.coupon.BuyingResponse")
}

func init() { proto.RegisterFile("buying.proto", fileDescriptor_fc0a9e5c6a9833d6) }

var fileDescriptor_fc0a9e5c6a9833d6 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x5c, 0xbb, 0xc9, 0x38, 0x45, 0x68, 0xd5, 0x82, 0x95, 0x52, 0xe1, 0x46, 0x42,
	0xf2, 0xc9, 0x20, 0x57, 0x88, 0x3b, 0x88, 0x83, 0x6f, 0xd5, 0xb6, 0xf7, 0xc8, 0xf1, 0x6e, 0xdc,
	0x55, 0xec, 0xdd, 0xb0, 0x7f, 0x22, 0xf5, 0x09, 0x78, 0x02, 0xde, 0x80, 0x07, 0x45, 0xbb, 0x6b,
	0x07, 0x0e, 0x24, 0xdc, 0xf2, 0xcd, 0xf7, 0x9b, 0xcc, 0xec, 0x37, 0x32, 0xcc, 0xd7, 0xe6, 0x99,
	0xf1, 0xb6, 0xd8, 0x49, 0xa1, 0x05, 0xba, 0x6a, 0x29, 0xfb, 0xce, 0x78, 0xa1, 0xe4, 0xbe, 0x60,
	0xbd, 0x2a, 0x1a, 0x61, 0x76, 0x82, 0x2f, 0xe6, 0x8d, 0xe8, 0x7b, 0xc1, 0x3d, 0xb4, 0xfc, 0x11,
	0x42, 0xfc, 0xc5, 0x75, 0xa1, 0x05, 0x4c, 0x95, 0xe8, 0x8c, 0x66, 0x82, 0xa7, 0x41, 0x16, 0xe4,
	0x33, 0x7c, 0xd0, 0xe8, 0x12, 0xa2, 0x46, 0x18, 0xae, 0xd3, 0x49, 0x16, 0xe4, 0x11, 0xf6, 0xc2,
	0x56, 0xb5, 0xd0, 0x75, 0x97, 0x86, 0x59, 0x90, 0x4f, 0xb0, 0x17, 0xf6, 0x7f, 0x08, 0x53, 0x1e,
	0x3f, 0x73, 0xc6, 0x41, 0xa3, 0x14, 0xce, 0x37, 0x92, 0xb2, 0xf6, 0x49, 0xa7, 0x91, 0xb3, 0x46,
	0x89, 0x5e, 0x43, 0x5c, 0xf7, 0xae, 0x27, 0x76, 0xc6, 0xa0, 0xd0, 0x0d, 0x40, 0x4d, 0x88, 0xa4,
	0x4a, 0xad, 0x18, 0x49, 0xcf, 0xb3, 0x20, 0x0f, 0xf1, 0x6c, 0xa8, 0x54, 0x04, 0xbd, 0x83, 0xa4,
	0x31, 0x4a, 0x8b, 0x9e, 0x4a, 0xeb, 0x4f, 0x9d, 0x0f, 0x63, 0xa9, 0x22, 0x68, 0x09, 0x17, 0x46,
	0xd1, 0x95, 0x66, 0xcd, 0x96, 0x6a, 0x8b, 0xcc, 0x1c, 0x92, 0x18, 0x45, 0x1f, 0x5d, 0xad, 0x22,
	0xe8, 0x23, 0x5c, 0xd6, 0xfb, 0x9a, 0x75, 0xf5, 0xba, 0xfb, 0x8b, 0x54, 0x29, 0x64, 0x61, 0x1e,
	0x62, 0x74, 0xf0, 0xc6, 0x06, 0x85, 0x3e, 0x43, 0xc4, 0x34, 0xed, 0x55, 0x9a, 0x64, 0x61, 0x9e,
	0x94, 0xb7, 0xc5, 0x3f, 0xb3, 0x2e, 0x7c, 0xb2, 0x95, 0xa6, 0x3d, 0xf6, 0xbc, 0x0d, 0xa0, 0x79,
	0xaa, 0x79, 0x4b, 0x49, 0x3a, 0xcf, 0x82, 0x7c, 0x8a, 0x47, 0xb9, 0xfc, 0x15, 0x00, 0xfc, 0xe1,
	0xd1, 0x1b, 0x38, 0xb7, 0x1d, 0x76, 0xe3, 0xc0, 0x6d, 0x1c, 0x5b, 0x59, 0x11, 0x74, 0x05, 0xb1,
	0xda, 0x1a, 0x5b, 0x9f, 0xb8, 0x7a, 0xa4, 0xb6, 0xa6, 0x22, 0xe8, 0x15, 0x84, 0xdc, 0xf4, 0xee,
	0x12, 0x11, 0xb6, 0x3f, 0xed, 0x75, 0x76, 0x92, 0x35, 0x74, 0x38, 0x82, 0x17, 0xe8, 0x16, 0xe6,
	0x42, 0xb2, 0x96, 0xf1, 0x95, 0x37, 0xfd, 0x19, 0x12, 0x5f, 0xbb, 0x77, 0xc8, 0x35, 0xcc, 0x94,
	0x59, 0xaf, 0xfc, 0x69, 0xfd, 0x35, 0xa6, 0xca, 0xac, 0x1f, 0xad, 0x5e, 0xfe, 0x9c, 0xc0, 0x4b,
	0xbf, 0x26, 0xa6, 0x6a, 0x27, 0xb8, 0xa2, 0xe8, 0x13, 0xc4, 0x94, 0x6b, 0xa6, 0x9f, 0xdd, 0xa6,
	0x49, 0x79, 0x73, 0x32, 0x0d, 0x3c, 0xc0, 0xa8, 0x84, 0x68, 0x57, 0xb7, 0x54, 0xba, 0x77, 0x24,
	0xe5, 0xdb, 0x23, 0x5d, 0xf7, 0x96, 0xc1, 0x1e, 0x45, 0x77, 0x63, 0xee, 0xa1, 0xcb, 0xfd, 0x3f,
	0x93, 0x86, 0xcc, 0x4b, 0x88, 0xa8, 0x94, 0x42, 0xba, 0x20, 0x8e, 0x0f, 0xfa, 0x66, 0x19, 0xec,
	0x51, 0xf4, 0x01, 0xce, 0x18, 0xdf, 0x08, 0x17, 0x4f, 0x52, 0x5e, 0x1f, 0x69, 0xa9, 0xf8, 0x46,
	0x60, 0x07, 0x96, 0x04, 0x2e, 0xfc, 0xd4, 0x07, 0x2a, 0xf7, 0x36, 0xc5, 0x07, 0x98, 0x7d, 0xad,
	0xbb, 0xc6, 0x74, 0xb5, 0xa6, 0xe8, 0xf4, 0xa2, 0x8b, 0xf7, 0xa7, 0xdf, 0x31, 0x04, 0xbd, 0x7c,
	0xb1, 0x8e, 0xdd, 0x57, 0x7b, 0xf7, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x27, 0x7f, 0x54, 0x2e, 0xea,
	0x03, 0x00, 0x00,
}
