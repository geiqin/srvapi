// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buying.proto

package geiqin_srv_tms

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
	Total                float32         `protobuf:"fixed32,1,opt,name=total,proto3" json:"total,omitempty"`
	Count                int32           `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	TotalWeight          float32         `protobuf:"fixed32,3,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight,omitempty"`
	Freight              float32         `protobuf:"fixed32,4,opt,name=freight,proto3" json:"freight,omitempty"`
	AddressId            int64           `protobuf:"varint,5,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	Items                []*BuyingItem   `protobuf:"bytes,6,rep,name=items,proto3" json:"items,omitempty"`
	DeliveryMethod       *DeliveryMethod `protobuf:"bytes,7,opt,name=delivery_method,json=deliveryMethod,proto3" json:"delivery_method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
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

func (m *Buying) GetTotalWeight() float32 {
	if m != nil {
		return m.TotalWeight
	}
	return 0
}

func (m *Buying) GetFreight() float32 {
	if m != nil {
		return m.Freight
	}
	return 0
}

func (m *Buying) GetAddressId() int64 {
	if m != nil {
		return m.AddressId
	}
	return 0
}

func (m *Buying) GetItems() []*BuyingItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Buying) GetDeliveryMethod() *DeliveryMethod {
	if m != nil {
		return m.DeliveryMethod
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

// 配送信息
type DeliveryMethod struct {
	Lng                  string    `protobuf:"bytes,1,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat                  string    `protobuf:"bytes,2,opt,name=lat,proto3" json:"lat,omitempty"`
	IsFetch              bool      `protobuf:"varint,3,opt,name=is_fetch,json=isFetch,proto3" json:"is_fetch,omitempty"`
	IsDelivery           bool      `protobuf:"varint,4,opt,name=is_delivery,json=isDelivery,proto3" json:"is_delivery,omitempty"`
	IsExpress            bool      `protobuf:"varint,5,opt,name=is_express,json=isExpress,proto3" json:"is_express,omitempty"`
	Method               string    `protobuf:"bytes,6,opt,name=method,proto3" json:"method,omitempty"`
	FetchCount           int32     `protobuf:"varint,7,opt,name=fetch_count,json=fetchCount,proto3" json:"fetch_count,omitempty"`
	FetchLocationId      int64     `protobuf:"varint,8,opt,name=fetch_location_id,json=fetchLocationId,proto3" json:"fetch_location_id,omitempty"`
	FetchAt              string    `protobuf:"bytes,9,opt,name=fetch_at,json=fetchAt,proto3" json:"fetch_at,omitempty"`
	Deliverable          bool      `protobuf:"varint,10,opt,name=deliverable,proto3" json:"deliverable,omitempty"`
	DeliveryAt           string    `protobuf:"bytes,11,opt,name=delivery_at,json=deliveryAt,proto3" json:"delivery_at,omitempty"`
	ExpressFee           float32   `protobuf:"fixed32,12,opt,name=express_fee,json=expressFee,proto3" json:"express_fee,omitempty"`
	DeliveryFee          float32   `protobuf:"fixed32,13,opt,name=delivery_fee,json=deliveryFee,proto3" json:"delivery_fee,omitempty"`
	DeliverInfo          *Delivery `protobuf:"bytes,14,opt,name=deliver_info,json=deliverInfo,proto3" json:"deliver_info,omitempty"`
	FetchInfo            *Fetch    `protobuf:"bytes,15,opt,name=fetch_info,json=fetchInfo,proto3" json:"fetch_info,omitempty"`
	ExpressInfo          *Express  `protobuf:"bytes,16,opt,name=express_info,json=expressInfo,proto3" json:"express_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeliveryMethod) Reset()         { *m = DeliveryMethod{} }
func (m *DeliveryMethod) String() string { return proto.CompactTextString(m) }
func (*DeliveryMethod) ProtoMessage()    {}
func (*DeliveryMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc0a9e5c6a9833d6, []int{2}
}

func (m *DeliveryMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryMethod.Unmarshal(m, b)
}
func (m *DeliveryMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryMethod.Marshal(b, m, deterministic)
}
func (m *DeliveryMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryMethod.Merge(m, src)
}
func (m *DeliveryMethod) XXX_Size() int {
	return xxx_messageInfo_DeliveryMethod.Size(m)
}
func (m *DeliveryMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryMethod.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryMethod proto.InternalMessageInfo

func (m *DeliveryMethod) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *DeliveryMethod) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *DeliveryMethod) GetIsFetch() bool {
	if m != nil {
		return m.IsFetch
	}
	return false
}

func (m *DeliveryMethod) GetIsDelivery() bool {
	if m != nil {
		return m.IsDelivery
	}
	return false
}

func (m *DeliveryMethod) GetIsExpress() bool {
	if m != nil {
		return m.IsExpress
	}
	return false
}

func (m *DeliveryMethod) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *DeliveryMethod) GetFetchCount() int32 {
	if m != nil {
		return m.FetchCount
	}
	return 0
}

func (m *DeliveryMethod) GetFetchLocationId() int64 {
	if m != nil {
		return m.FetchLocationId
	}
	return 0
}

func (m *DeliveryMethod) GetFetchAt() string {
	if m != nil {
		return m.FetchAt
	}
	return ""
}

func (m *DeliveryMethod) GetDeliverable() bool {
	if m != nil {
		return m.Deliverable
	}
	return false
}

func (m *DeliveryMethod) GetDeliveryAt() string {
	if m != nil {
		return m.DeliveryAt
	}
	return ""
}

func (m *DeliveryMethod) GetExpressFee() float32 {
	if m != nil {
		return m.ExpressFee
	}
	return 0
}

func (m *DeliveryMethod) GetDeliveryFee() float32 {
	if m != nil {
		return m.DeliveryFee
	}
	return 0
}

func (m *DeliveryMethod) GetDeliverInfo() *Delivery {
	if m != nil {
		return m.DeliverInfo
	}
	return nil
}

func (m *DeliveryMethod) GetFetchInfo() *Fetch {
	if m != nil {
		return m.FetchInfo
	}
	return nil
}

func (m *DeliveryMethod) GetExpressInfo() *Express {
	if m != nil {
		return m.ExpressInfo
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
	proto.RegisterType((*Buying)(nil), "geiqin.srv.tms.Buying")
	proto.RegisterType((*BuyingItem)(nil), "geiqin.srv.tms.BuyingItem")
	proto.RegisterType((*DeliveryMethod)(nil), "geiqin.srv.tms.DeliveryMethod")
	proto.RegisterType((*BuyingResponse)(nil), "geiqin.srv.tms.BuyingResponse")
}

func init() {
	proto.RegisterFile("buying.proto", fileDescriptor_fc0a9e5c6a9833d6)
}

var fileDescriptor_fc0a9e5c6a9833d6 = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0x9c, 0xd4, 0x4e, 0x32, 0x4e, 0xd3, 0x7e, 0xab, 0xfe, 0x98, 0x20, 0x4a, 0xc8, 0x55,
	0x44, 0x51, 0x84, 0x02, 0x57, 0x70, 0x55, 0xfa, 0xa7, 0x48, 0x20, 0x55, 0x4b, 0x25, 0x2e, 0x2d,
	0x27, 0xde, 0xa4, 0xab, 0xc6, 0xde, 0xe0, 0x5d, 0x17, 0xfa, 0x0c, 0xbc, 0x06, 0xaf, 0xc2, 0x0b,
	0xf1, 0x04, 0x68, 0x66, 0xd7, 0x81, 0x46, 0xcd, 0x9d, 0xe7, 0xcc, 0x99, 0xdd, 0x33, 0x67, 0x66,
	0x0d, 0xed, 0x49, 0x79, 0x2f, 0xf3, 0xf9, 0x70, 0x59, 0x28, 0xa3, 0x58, 0x67, 0x2e, 0xe4, 0x57,
	0x99, 0x0f, 0x75, 0x71, 0x37, 0x34, 0x99, 0xee, 0xb6, 0xa7, 0x2a, 0xcb, 0x54, 0x6e, 0xb3, 0xdd,
	0x70, 0x26, 0xcc, 0xf4, 0xc6, 0x05, 0x9d, 0x54, 0x2c, 0xe4, 0x9d, 0x28, 0xee, 0x5d, 0xbc, 0x2d,
	0xbe, 0x2f, 0x0b, 0xa1, 0xb5, 0x0d, 0xfb, 0x3f, 0x6a, 0x10, 0x7c, 0xa0, 0xa3, 0xd9, 0x1e, 0xf8,
	0x46, 0x99, 0x64, 0x11, 0x79, 0x3d, 0x6f, 0x50, 0xe3, 0x36, 0x40, 0x74, 0xaa, 0xca, 0xdc, 0x44,
	0xb5, 0x9e, 0x37, 0xf0, 0xb9, 0x0d, 0xd8, 0x0b, 0x68, 0x53, 0x3a, 0xfe, 0x26, 0xe4, 0xfc, 0xc6,
	0x44, 0x75, 0x2a, 0x09, 0x09, 0xfb, 0x42, 0x10, 0x8b, 0xa0, 0x31, 0x2b, 0x6c, 0x76, 0x8b, 0xb2,
	0x55, 0xc8, 0x9e, 0x01, 0x24, 0x69, 0x8a, 0x22, 0x62, 0x99, 0x46, 0x7e, 0xcf, 0x1b, 0xd4, 0x79,
	0xcb, 0x21, 0xe3, 0x94, 0xbd, 0x06, 0x5f, 0x1a, 0x91, 0xe9, 0x28, 0xe8, 0xd5, 0x07, 0xe1, 0xa8,
	0x3b, 0x7c, 0xd8, 0xec, 0xd0, 0xca, 0x1d, 0x1b, 0x91, 0x71, 0x4b, 0x64, 0x97, 0xb0, 0x53, 0x75,
	0x19, 0x67, 0xc2, 0xdc, 0xa8, 0x34, 0x6a, 0xf4, 0xbc, 0x41, 0x38, 0x3a, 0x5a, 0xaf, 0x3d, 0x73,
	0xb4, 0x4f, 0xc4, 0xe2, 0x2b, 0x73, 0x6c, 0xdc, 0xff, 0xe9, 0x01, 0xfc, 0x3d, 0x9e, 0x1d, 0x42,
	0x03, 0x2f, 0x40, 0x95, 0x1e, 0xa9, 0x0c, 0x30, 0x1c, 0xa7, 0x6c, 0x1f, 0x02, 0x7d, 0x5b, 0x22,
	0x5e, 0x23, 0xdc, 0xd7, 0xb7, 0xe5, 0x38, 0x65, 0xbb, 0x50, 0xcf, 0xcb, 0x8c, 0xcc, 0xf0, 0x39,
	0x7e, 0xa2, 0x7b, 0xcb, 0x42, 0x4e, 0x85, 0xb3, 0xc0, 0x06, 0xe8, 0x9e, 0x2a, 0xe4, 0x5c, 0xe6,
	0xb1, 0x4d, 0xfa, 0xd6, 0x3d, 0x8b, 0x5d, 0x11, 0xe5, 0x29, 0xb4, 0x74, 0x39, 0x89, 0xed, 0x40,
	0x02, 0xca, 0x37, 0x75, 0x39, 0xb9, 0xc6, 0xb8, 0xff, 0x6b, 0x0b, 0x3a, 0x0f, 0x3b, 0xc1, 0xab,
	0x17, 0xf9, 0x9c, 0x64, 0xb6, 0x38, 0x7e, 0x12, 0x92, 0xd8, 0xb1, 0x21, 0x92, 0x18, 0xf6, 0x04,
	0x9a, 0x52, 0xc7, 0xb4, 0x1c, 0xa4, 0xb1, 0xc9, 0x1b, 0x52, 0x5f, 0x60, 0xc8, 0x9e, 0x43, 0x28,
	0x75, 0x5c, 0xb9, 0x41, 0x6a, 0x9b, 0x1c, 0xa4, 0xae, 0x6e, 0xc1, 0x99, 0x49, 0x1d, 0xbb, 0xdd,
	0x21, 0xc1, 0x4d, 0xde, 0x92, 0xfa, 0xdc, 0x02, 0xec, 0x00, 0x02, 0x67, 0x7c, 0x40, 0xf7, 0xb9,
	0x08, 0xcf, 0xa5, 0xfb, 0x62, 0xbb, 0x43, 0x0d, 0x72, 0x06, 0x08, 0x3a, 0xa5, 0x45, 0x7a, 0x09,
	0xff, 0x5b, 0xc2, 0x42, 0x4d, 0x13, 0x23, 0x55, 0x8e, 0xa6, 0x36, 0xc9, 0xd4, 0x1d, 0x4a, 0x7c,
	0x74, 0xf8, 0x38, 0x45, 0xfd, 0x96, 0x9b, 0x98, 0xa8, 0x45, 0xd7, 0x34, 0x28, 0x3e, 0x31, 0xac,
	0x07, 0xa1, 0x13, 0x9f, 0x4c, 0x16, 0x22, 0x02, 0xd2, 0xf7, 0x2f, 0x84, 0x4a, 0x56, 0x3b, 0x92,
	0x98, 0x28, 0xa4, 0x7a, 0xa8, 0xa0, 0x13, 0x83, 0x04, 0xd7, 0x5e, 0x3c, 0x13, 0x22, 0x6a, 0x93,
	0xe7, 0xe0, 0xa0, 0x0b, 0x41, 0x53, 0x5b, 0x9d, 0x80, 0x8c, 0x6d, 0x3b, 0xb5, 0x0a, 0x43, 0xca,
	0xfb, 0x15, 0x25, 0x96, 0xf9, 0x4c, 0x45, 0x1d, 0xda, 0xc2, 0x68, 0xd3, 0x16, 0xae, 0x8a, 0xc7,
	0xf9, 0x4c, 0xb1, 0xb7, 0x60, 0x8d, 0xb1, 0xa5, 0x3b, 0x54, 0xba, 0xbf, 0x5e, 0x4a, 0xe3, 0xe2,
	0x2d, 0x22, 0x52, 0xd5, 0x3b, 0x68, 0x57, 0xb2, 0xa9, 0x6e, 0x97, 0xea, 0x0e, 0xd7, 0xeb, 0xdc,
	0xa0, 0x78, 0xd5, 0x23, 0xd6, 0xf6, 0x7f, 0x7b, 0xd0, 0xb1, 0xeb, 0xce, 0x85, 0x5e, 0xaa, 0x5c,
	0x0b, 0x36, 0x84, 0x40, 0xe4, 0x46, 0x9a, 0x7b, 0x5a, 0xa5, 0x70, 0x74, 0xf0, 0xf8, 0xeb, 0xe3,
	0x8e, 0xc5, 0x8e, 0xc1, 0x5f, 0x26, 0x73, 0x51, 0xd0, 0x9e, 0x3d, 0xa2, 0xf7, 0x0a, 0x93, 0xdc,
	0x72, 0xd8, 0xab, 0xea, 0x65, 0xd7, 0xe9, 0x65, 0x6f, 0x3a, 0xdb, 0xbd, 0xea, 0x63, 0xf0, 0x45,
	0x51, 0xa8, 0x82, 0xb6, 0xf1, 0x91, 0xa3, 0xcf, 0x31, 0xc9, 0x2d, 0x87, 0x0d, 0x60, 0x8b, 0xda,
	0xf7, 0x89, 0xbb, 0xb7, 0xce, 0xc5, 0x76, 0x39, 0x31, 0x46, 0xd7, 0xb0, 0x6d, 0xef, 0xf9, 0x2c,
	0x8a, 0x3b, 0x7c, 0x6a, 0xa7, 0xd0, 0xb8, 0x14, 0xe6, 0x2c, 0x31, 0x09, 0xdb, 0xa0, 0xa8, 0x7b,
	0xb4, 0x41, 0xa9, 0x73, 0xad, 0xff, 0xdf, 0x24, 0xa0, 0xdf, 0xe9, 0x9b, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x24, 0x5d, 0xd0, 0x31, 0xa8, 0x05, 0x00, 0x00,
}
