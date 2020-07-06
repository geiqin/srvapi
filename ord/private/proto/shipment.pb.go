// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shipment.proto

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

// 下单的配送信息
type OrderShipment struct {
	Lng                  string        `protobuf:"bytes,1,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat                  string        `protobuf:"bytes,2,opt,name=lat,proto3" json:"lat,omitempty"`
	IsFetch              bool          `protobuf:"varint,3,opt,name=is_fetch,json=isFetch,proto3" json:"is_fetch,omitempty"`
	IsDelivery           bool          `protobuf:"varint,4,opt,name=is_delivery,json=isDelivery,proto3" json:"is_delivery,omitempty"`
	IsExpress            bool          `protobuf:"varint,5,opt,name=is_express,json=isExpress,proto3" json:"is_express,omitempty"`
	Method               string        `protobuf:"bytes,6,opt,name=method,proto3" json:"method,omitempty"`
	FetchCount           int32         `protobuf:"varint,7,opt,name=fetch_count,json=fetchCount,proto3" json:"fetch_count,omitempty"`
	FetchLocationId      int64         `protobuf:"varint,8,opt,name=fetch_location_id,json=fetchLocationId,proto3" json:"fetch_location_id,omitempty"`
	FetchAt              string        `protobuf:"bytes,9,opt,name=fetch_at,json=fetchAt,proto3" json:"fetch_at,omitempty"`
	Deliverable          bool          `protobuf:"varint,10,opt,name=deliverable,proto3" json:"deliverable,omitempty"`
	DeliveryAt           string        `protobuf:"bytes,11,opt,name=delivery_at,json=deliveryAt,proto3" json:"delivery_at,omitempty"`
	ExpressFee           float32       `protobuf:"fixed32,12,opt,name=express_fee,json=expressFee,proto3" json:"express_fee,omitempty"`
	DeliveryFee          float32       `protobuf:"fixed32,13,opt,name=delivery_fee,json=deliveryFee,proto3" json:"delivery_fee,omitempty"`
	DeliveryInfo         *DeliveryInfo `protobuf:"bytes,14,opt,name=delivery_info,json=deliveryInfo,proto3" json:"delivery_info,omitempty"`
	FetchInfo            *FetchInfo    `protobuf:"bytes,15,opt,name=fetch_info,json=fetchInfo,proto3" json:"fetch_info,omitempty"`
	ExpressInfo          *ExpressInfo  `protobuf:"bytes,16,opt,name=express_info,json=expressInfo,proto3" json:"express_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *OrderShipment) Reset()         { *m = OrderShipment{} }
func (m *OrderShipment) String() string { return proto.CompactTextString(m) }
func (*OrderShipment) ProtoMessage()    {}
func (*OrderShipment) Descriptor() ([]byte, []int) {
	return fileDescriptor_832505ff9d8a41db, []int{0}
}

func (m *OrderShipment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderShipment.Unmarshal(m, b)
}
func (m *OrderShipment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderShipment.Marshal(b, m, deterministic)
}
func (m *OrderShipment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderShipment.Merge(m, src)
}
func (m *OrderShipment) XXX_Size() int {
	return xxx_messageInfo_OrderShipment.Size(m)
}
func (m *OrderShipment) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderShipment.DiscardUnknown(m)
}

var xxx_messageInfo_OrderShipment proto.InternalMessageInfo

func (m *OrderShipment) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *OrderShipment) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *OrderShipment) GetIsFetch() bool {
	if m != nil {
		return m.IsFetch
	}
	return false
}

func (m *OrderShipment) GetIsDelivery() bool {
	if m != nil {
		return m.IsDelivery
	}
	return false
}

func (m *OrderShipment) GetIsExpress() bool {
	if m != nil {
		return m.IsExpress
	}
	return false
}

func (m *OrderShipment) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *OrderShipment) GetFetchCount() int32 {
	if m != nil {
		return m.FetchCount
	}
	return 0
}

func (m *OrderShipment) GetFetchLocationId() int64 {
	if m != nil {
		return m.FetchLocationId
	}
	return 0
}

func (m *OrderShipment) GetFetchAt() string {
	if m != nil {
		return m.FetchAt
	}
	return ""
}

func (m *OrderShipment) GetDeliverable() bool {
	if m != nil {
		return m.Deliverable
	}
	return false
}

func (m *OrderShipment) GetDeliveryAt() string {
	if m != nil {
		return m.DeliveryAt
	}
	return ""
}

func (m *OrderShipment) GetExpressFee() float32 {
	if m != nil {
		return m.ExpressFee
	}
	return 0
}

func (m *OrderShipment) GetDeliveryFee() float32 {
	if m != nil {
		return m.DeliveryFee
	}
	return 0
}

func (m *OrderShipment) GetDeliveryInfo() *DeliveryInfo {
	if m != nil {
		return m.DeliveryInfo
	}
	return nil
}

func (m *OrderShipment) GetFetchInfo() *FetchInfo {
	if m != nil {
		return m.FetchInfo
	}
	return nil
}

func (m *OrderShipment) GetExpressInfo() *ExpressInfo {
	if m != nil {
		return m.ExpressInfo
	}
	return nil
}

// 上门自提信息
type FetchInfo struct {
	Id                   int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AreaId               int64     `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Addr                 string    `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	Lng                  string    `protobuf:"bytes,5,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat                  string    `protobuf:"bytes,6,opt,name=lat,proto3" json:"lat,omitempty"`
	Tel                  string    `protobuf:"bytes,7,opt,name=tel,proto3" json:"tel,omitempty"`
	Mobile               string    `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty"`
	IsFetchTime          bool      `protobuf:"varint,9,opt,name=is_fetch_time,json=isFetchTime,proto3" json:"is_fetch_time,omitempty"`
	Distance             float32   `protobuf:"fixed32,10,opt,name=distance,proto3" json:"distance,omitempty"`
	Galleries            []string  `protobuf:"bytes,11,rep,name=galleries,proto3" json:"galleries,omitempty"`
	Area                 *AreaInfo `protobuf:"bytes,12,opt,name=area,proto3" json:"area,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FetchInfo) Reset()         { *m = FetchInfo{} }
func (m *FetchInfo) String() string { return proto.CompactTextString(m) }
func (*FetchInfo) ProtoMessage()    {}
func (*FetchInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_832505ff9d8a41db, []int{1}
}

func (m *FetchInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchInfo.Unmarshal(m, b)
}
func (m *FetchInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchInfo.Marshal(b, m, deterministic)
}
func (m *FetchInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchInfo.Merge(m, src)
}
func (m *FetchInfo) XXX_Size() int {
	return xxx_messageInfo_FetchInfo.Size(m)
}
func (m *FetchInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FetchInfo proto.InternalMessageInfo

func (m *FetchInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FetchInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FetchInfo) GetAreaId() int64 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *FetchInfo) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *FetchInfo) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *FetchInfo) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *FetchInfo) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

func (m *FetchInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *FetchInfo) GetIsFetchTime() bool {
	if m != nil {
		return m.IsFetchTime
	}
	return false
}

func (m *FetchInfo) GetDistance() float32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *FetchInfo) GetGalleries() []string {
	if m != nil {
		return m.Galleries
	}
	return nil
}

func (m *FetchInfo) GetArea() *AreaInfo {
	if m != nil {
		return m.Area
	}
	return nil
}

// 同城配送信息
type DeliveryInfo struct {
	LocationId           int64    `protobuf:"varint,1,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Method               int32    `protobuf:"varint,2,opt,name=method,proto3" json:"method,omitempty"`
	Template             int32    `protobuf:"varint,3,opt,name=template,proto3" json:"template,omitempty"`
	RangeName            string   `protobuf:"bytes,4,opt,name=range_name,json=rangeName,proto3" json:"range_name,omitempty"`
	RangeContent         string   `protobuf:"bytes,5,opt,name=range_content,json=rangeContent,proto3" json:"range_content,omitempty"`
	RangeImageUrl        string   `protobuf:"bytes,6,opt,name=range_image_url,json=rangeImageUrl,proto3" json:"range_image_url,omitempty"`
	RangeStartPrice      float32  `protobuf:"fixed32,7,opt,name=range_start_price,json=rangeStartPrice,proto3" json:"range_start_price,omitempty"`
	RangeFee             float32  `protobuf:"fixed32,8,opt,name=range_fee,json=rangeFee,proto3" json:"range_fee,omitempty"`
	IsDeliveryTime       bool     `protobuf:"varint,9,opt,name=is_delivery_time,json=isDeliveryTime,proto3" json:"is_delivery_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliveryInfo) Reset()         { *m = DeliveryInfo{} }
func (m *DeliveryInfo) String() string { return proto.CompactTextString(m) }
func (*DeliveryInfo) ProtoMessage()    {}
func (*DeliveryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_832505ff9d8a41db, []int{2}
}

func (m *DeliveryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryInfo.Unmarshal(m, b)
}
func (m *DeliveryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryInfo.Marshal(b, m, deterministic)
}
func (m *DeliveryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryInfo.Merge(m, src)
}
func (m *DeliveryInfo) XXX_Size() int {
	return xxx_messageInfo_DeliveryInfo.Size(m)
}
func (m *DeliveryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryInfo proto.InternalMessageInfo

func (m *DeliveryInfo) GetLocationId() int64 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

func (m *DeliveryInfo) GetMethod() int32 {
	if m != nil {
		return m.Method
	}
	return 0
}

func (m *DeliveryInfo) GetTemplate() int32 {
	if m != nil {
		return m.Template
	}
	return 0
}

func (m *DeliveryInfo) GetRangeName() string {
	if m != nil {
		return m.RangeName
	}
	return ""
}

func (m *DeliveryInfo) GetRangeContent() string {
	if m != nil {
		return m.RangeContent
	}
	return ""
}

func (m *DeliveryInfo) GetRangeImageUrl() string {
	if m != nil {
		return m.RangeImageUrl
	}
	return ""
}

func (m *DeliveryInfo) GetRangeStartPrice() float32 {
	if m != nil {
		return m.RangeStartPrice
	}
	return 0
}

func (m *DeliveryInfo) GetRangeFee() float32 {
	if m != nil {
		return m.RangeFee
	}
	return 0
}

func (m *DeliveryInfo) GetIsDeliveryTime() bool {
	if m != nil {
		return m.IsDeliveryTime
	}
	return false
}

type ExpressInfo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpressInfo) Reset()         { *m = ExpressInfo{} }
func (m *ExpressInfo) String() string { return proto.CompactTextString(m) }
func (*ExpressInfo) ProtoMessage()    {}
func (*ExpressInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_832505ff9d8a41db, []int{3}
}

func (m *ExpressInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpressInfo.Unmarshal(m, b)
}
func (m *ExpressInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpressInfo.Marshal(b, m, deterministic)
}
func (m *ExpressInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpressInfo.Merge(m, src)
}
func (m *ExpressInfo) XXX_Size() int {
	return xxx_messageInfo_ExpressInfo.Size(m)
}
func (m *ExpressInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpressInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExpressInfo proto.InternalMessageInfo

// 订单发货的请求数据
type ShipmentRequest struct {
	Method               string                   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Type                 string                   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	OrderId              int64                    `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	IsDelivery           bool                     `protobuf:"varint,4,opt,name=is_delivery,json=isDelivery,proto3" json:"is_delivery,omitempty"`
	ShipperId            int32                    `protobuf:"varint,5,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	LogisticsNo          string                   `protobuf:"bytes,6,opt,name=logistics_no,json=logisticsNo,proto3" json:"logistics_no,omitempty"`
	LocationId           int64                    `protobuf:"varint,7,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	FetchAt              string                   `protobuf:"bytes,8,opt,name=fetch_at,json=fetchAt,proto3" json:"fetch_at,omitempty"`
	DeliveryAt           string                   `protobuf:"bytes,9,opt,name=delivery_at,json=deliveryAt,proto3" json:"delivery_at,omitempty"`
	Details              []*ShipmentRequestDetail `protobuf:"bytes,10,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ShipmentRequest) Reset()         { *m = ShipmentRequest{} }
func (m *ShipmentRequest) String() string { return proto.CompactTextString(m) }
func (*ShipmentRequest) ProtoMessage()    {}
func (*ShipmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_832505ff9d8a41db, []int{4}
}

func (m *ShipmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipmentRequest.Unmarshal(m, b)
}
func (m *ShipmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipmentRequest.Marshal(b, m, deterministic)
}
func (m *ShipmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipmentRequest.Merge(m, src)
}
func (m *ShipmentRequest) XXX_Size() int {
	return xxx_messageInfo_ShipmentRequest.Size(m)
}
func (m *ShipmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShipmentRequest proto.InternalMessageInfo

func (m *ShipmentRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ShipmentRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ShipmentRequest) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *ShipmentRequest) GetIsDelivery() bool {
	if m != nil {
		return m.IsDelivery
	}
	return false
}

func (m *ShipmentRequest) GetShipperId() int32 {
	if m != nil {
		return m.ShipperId
	}
	return 0
}

func (m *ShipmentRequest) GetLogisticsNo() string {
	if m != nil {
		return m.LogisticsNo
	}
	return ""
}

func (m *ShipmentRequest) GetLocationId() int64 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

func (m *ShipmentRequest) GetFetchAt() string {
	if m != nil {
		return m.FetchAt
	}
	return ""
}

func (m *ShipmentRequest) GetDeliveryAt() string {
	if m != nil {
		return m.DeliveryAt
	}
	return ""
}

func (m *ShipmentRequest) GetDetails() []*ShipmentRequestDetail {
	if m != nil {
		return m.Details
	}
	return nil
}

type ShipmentRequestDetail struct {
	OrderDetailId        int64    `protobuf:"varint,1,opt,name=order_detail_id,json=orderDetailId,proto3" json:"order_detail_id,omitempty"`
	ItemId               int64    `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64    `protobuf:"varint,3,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Num                  int32    `protobuf:"varint,4,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipmentRequestDetail) Reset()         { *m = ShipmentRequestDetail{} }
func (m *ShipmentRequestDetail) String() string { return proto.CompactTextString(m) }
func (*ShipmentRequestDetail) ProtoMessage()    {}
func (*ShipmentRequestDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_832505ff9d8a41db, []int{5}
}

func (m *ShipmentRequestDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipmentRequestDetail.Unmarshal(m, b)
}
func (m *ShipmentRequestDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipmentRequestDetail.Marshal(b, m, deterministic)
}
func (m *ShipmentRequestDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipmentRequestDetail.Merge(m, src)
}
func (m *ShipmentRequestDetail) XXX_Size() int {
	return xxx_messageInfo_ShipmentRequestDetail.Size(m)
}
func (m *ShipmentRequestDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipmentRequestDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ShipmentRequestDetail proto.InternalMessageInfo

func (m *ShipmentRequestDetail) GetOrderDetailId() int64 {
	if m != nil {
		return m.OrderDetailId
	}
	return 0
}

func (m *ShipmentRequestDetail) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ShipmentRequestDetail) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *ShipmentRequestDetail) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type ShipmentResponse struct {
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipmentResponse) Reset()         { *m = ShipmentResponse{} }
func (m *ShipmentResponse) String() string { return proto.CompactTextString(m) }
func (*ShipmentResponse) ProtoMessage()    {}
func (*ShipmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_832505ff9d8a41db, []int{6}
}

func (m *ShipmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipmentResponse.Unmarshal(m, b)
}
func (m *ShipmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipmentResponse.Marshal(b, m, deterministic)
}
func (m *ShipmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipmentResponse.Merge(m, src)
}
func (m *ShipmentResponse) XXX_Size() int {
	return xxx_messageInfo_ShipmentResponse.Size(m)
}
func (m *ShipmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShipmentResponse proto.InternalMessageInfo

func (m *ShipmentResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ShipmentResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderShipment)(nil), "geiqin.srv.ord.private.OrderShipment")
	proto.RegisterType((*FetchInfo)(nil), "geiqin.srv.ord.private.FetchInfo")
	proto.RegisterType((*DeliveryInfo)(nil), "geiqin.srv.ord.private.DeliveryInfo")
	proto.RegisterType((*ExpressInfo)(nil), "geiqin.srv.ord.private.ExpressInfo")
	proto.RegisterType((*ShipmentRequest)(nil), "geiqin.srv.ord.private.ShipmentRequest")
	proto.RegisterType((*ShipmentRequestDetail)(nil), "geiqin.srv.ord.private.ShipmentRequestDetail")
	proto.RegisterType((*ShipmentResponse)(nil), "geiqin.srv.ord.private.ShipmentResponse")
}

func init() {
	proto.RegisterFile("shipment.proto", fileDescriptor_832505ff9d8a41db)
}

var fileDescriptor_832505ff9d8a41db = []byte{
	// 897 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x8e, 0xdb, 0x36,
	0x10, 0x86, 0x65, 0xcb, 0xb6, 0x46, 0xf6, 0xae, 0x4b, 0x20, 0xa9, 0xba, 0x4d, 0x50, 0xc5, 0x29,
	0x0a, 0xa1, 0x40, 0x8d, 0x62, 0xd3, 0x07, 0xe8, 0x22, 0xe9, 0x16, 0x02, 0x8a, 0xb4, 0x60, 0xda,
	0xb3, 0xa0, 0xb5, 0x68, 0x2f, 0x11, 0x49, 0x54, 0x48, 0x7a, 0xd1, 0x3d, 0xf6, 0xc5, 0x7a, 0xe8,
	0x9b, 0xf4, 0xda, 0xa7, 0x28, 0x66, 0xa8, 0x1f, 0x27, 0x1b, 0x03, 0xb9, 0x71, 0x3e, 0x7e, 0x1c,
	0x91, 0x33, 0xdf, 0x7c, 0x82, 0x33, 0x73, 0x2b, 0x9b, 0x4a, 0xd4, 0x76, 0xd3, 0x68, 0x65, 0x15,
	0x7b, 0xbc, 0x17, 0xf2, 0x9d, 0xac, 0x37, 0x46, 0xdf, 0x6d, 0x94, 0x2e, 0x36, 0x8d, 0x96, 0x77,
	0xb9, 0x15, 0x17, 0x8b, 0xad, 0xaa, 0x2a, 0x55, 0x3b, 0xd6, 0xc5, 0x59, 0xae, 0x45, 0x9e, 0xd6,
	0x3b, 0xe5, 0xe2, 0xf5, 0x7f, 0x13, 0x58, 0xfe, 0xaa, 0x0b, 0xa1, 0xdf, 0xb4, 0xd9, 0xd8, 0x0a,
	0xc6, 0x65, 0xbd, 0x8f, 0x46, 0xf1, 0x28, 0x09, 0x38, 0x2e, 0x09, 0xc9, 0x6d, 0xe4, 0xb5, 0x48,
	0x6e, 0xd9, 0x17, 0x30, 0x97, 0x26, 0xdb, 0x09, 0xbb, 0xbd, 0x8d, 0xc6, 0xf1, 0x28, 0x99, 0xf3,
	0x99, 0x34, 0xd7, 0x18, 0xb2, 0xaf, 0x20, 0x94, 0x26, 0x2b, 0x44, 0x29, 0xef, 0x84, 0xbe, 0x8f,
	0x26, 0xb4, 0x0b, 0xd2, 0xbc, 0x6a, 0x11, 0xf6, 0x14, 0x40, 0x9a, 0x4c, 0xfc, 0xd9, 0x68, 0x61,
	0x4c, 0xe4, 0xd3, 0x7e, 0x20, 0xcd, 0x4f, 0x0e, 0x60, 0x8f, 0x61, 0x5a, 0x09, 0x7b, 0xab, 0x8a,
	0x68, 0x4a, 0xdf, 0x6b, 0x23, 0xcc, 0x4b, 0xdf, 0xcb, 0xb6, 0xea, 0x50, 0xdb, 0x68, 0x16, 0x8f,
	0x12, 0x9f, 0x03, 0x41, 0x2f, 0x11, 0x61, 0xdf, 0xc2, 0x67, 0x8e, 0x50, 0xaa, 0x6d, 0x6e, 0xa5,
	0xaa, 0x33, 0x59, 0x44, 0xf3, 0x78, 0x94, 0x8c, 0xf9, 0x39, 0x6d, 0xfc, 0xd2, 0xe2, 0x69, 0x81,
	0xf7, 0x77, 0xdc, 0xdc, 0x46, 0x01, 0x7d, 0x66, 0x46, 0xf1, 0x95, 0x65, 0x31, 0x84, 0xed, 0xe5,
	0xf3, 0x9b, 0x52, 0x44, 0x40, 0xf7, 0x3b, 0x86, 0xf0, 0x26, 0xdd, 0xf3, 0xf0, 0x7c, 0x48, 0xe7,
	0xa1, 0x83, 0xae, 0x2c, 0x12, 0xda, 0xe7, 0x65, 0x3b, 0x21, 0xa2, 0x45, 0x3c, 0x4a, 0x3c, 0x0e,
	0x2d, 0x74, 0x2d, 0x04, 0x7b, 0x06, 0x8b, 0x3e, 0x03, 0x32, 0x96, 0xc4, 0xe8, 0xb3, 0x22, 0x25,
	0x85, 0x65, 0x4f, 0x91, 0xf5, 0x4e, 0x45, 0x67, 0xf1, 0x28, 0x09, 0x2f, 0xbf, 0xde, 0x7c, 0xbc,
	0xcb, 0x9b, 0xae, 0xbc, 0xd8, 0x5a, 0xde, 0x67, 0xc7, 0x88, 0xfd, 0x08, 0xae, 0x4c, 0x2e, 0xcf,
	0x39, 0xe5, 0x79, 0x76, 0x2a, 0x0f, 0x35, 0x91, 0x92, 0x04, 0xbb, 0x6e, 0xc9, 0xae, 0x61, 0xd1,
	0x3d, 0x88, 0x72, 0xac, 0x28, 0xc7, 0xf3, 0x53, 0x39, 0xda, 0x56, 0x52, 0x96, 0xae, 0x12, 0x18,
	0xac, 0xff, 0xf6, 0x20, 0xe8, 0x3f, 0xc0, 0xce, 0xc0, 0x93, 0x05, 0xe9, 0x6c, 0xcc, 0x3d, 0x59,
	0x30, 0x06, 0x93, 0x3a, 0xaf, 0x44, 0xab, 0x33, 0x5a, 0xb3, 0xcf, 0x61, 0x86, 0x82, 0xc5, 0x56,
	0x8e, 0x89, 0x38, 0x25, 0xfd, 0x12, 0x39, 0x2f, 0x0a, 0x4d, 0xfa, 0x0a, 0x38, 0xad, 0x3b, 0xe5,
	0xfa, 0x0f, 0x94, 0x3b, 0x1d, 0x94, 0xbb, 0x82, 0xb1, 0x15, 0x25, 0xc9, 0x27, 0xe0, 0xb8, 0x24,
	0xc1, 0xa9, 0x1b, 0x59, 0x0a, 0x12, 0x0b, 0x0a, 0x8e, 0x22, 0xb6, 0x86, 0x65, 0xa7, 0xf1, 0xcc,
	0xca, 0x4a, 0x90, 0x50, 0xe6, 0x3c, 0x6c, 0x85, 0xfe, 0xbb, 0xac, 0x04, 0xbb, 0x80, 0x79, 0x21,
	0x8d, 0xcd, 0xeb, 0xad, 0x53, 0x8a, 0xc7, 0xfb, 0x98, 0x3d, 0x81, 0x60, 0x9f, 0x97, 0xa5, 0xd0,
	0x52, 0x98, 0x28, 0x8c, 0xc7, 0x49, 0xc0, 0x07, 0x80, 0xfd, 0x00, 0x13, 0x7c, 0x09, 0x89, 0x23,
	0xbc, 0x8c, 0x4f, 0x95, 0xf2, 0xaa, 0x9d, 0x56, 0x4e, 0xec, 0xf5, 0x3f, 0x1e, 0x2c, 0x8e, 0x3b,
	0x8d, 0x52, 0x3b, 0x96, 0xbb, 0x2b, 0x26, 0x94, 0x83, 0xd2, 0x87, 0x71, 0xf2, 0x68, 0x62, 0xba,
	0x71, 0xba, 0x80, 0xb9, 0x15, 0x55, 0x53, 0xe6, 0x56, 0x50, 0x65, 0x7d, 0xde, 0xc7, 0x38, 0xa1,
	0x3a, 0xaf, 0xf7, 0x22, 0xa3, 0x76, 0xb8, 0x0a, 0x07, 0x84, 0xbc, 0xc6, 0x9e, 0x3c, 0x87, 0xa5,
	0xdb, 0xde, 0xaa, 0xda, 0x8a, 0xda, 0xb6, 0x05, 0x5f, 0x10, 0xf8, 0xd2, 0x61, 0xec, 0x1b, 0x38,
	0x77, 0x24, 0x59, 0xe5, 0x7b, 0x91, 0x1d, 0x74, 0xd9, 0x76, 0xc1, 0x9d, 0x4d, 0x11, 0xfd, 0x43,
	0x97, 0x38, 0xb5, 0x8e, 0x67, 0x6c, 0xae, 0x6d, 0xd6, 0x68, 0xb9, 0x15, 0xd4, 0x1d, 0x8f, 0xbb,
	0x04, 0x6f, 0x10, 0xff, 0x0d, 0x61, 0xf6, 0x25, 0xb8, 0x5b, 0xd0, 0xcc, 0xcc, 0x5d, 0xb9, 0x09,
	0xc0, 0x81, 0x49, 0x60, 0x75, 0xe4, 0x3b, 0xc7, 0x1d, 0x3b, 0x1b, 0xcc, 0x07, 0x9b, 0xb6, 0x5e,
	0x42, 0x78, 0xa4, 0xd0, 0xf5, 0xbf, 0x1e, 0x9c, 0x77, 0xe6, 0xc7, 0xc5, 0xbb, 0x83, 0x30, 0xf6,
	0xa8, 0x6a, 0xa3, 0xf7, 0x4c, 0x88, 0xc1, 0xc4, 0xde, 0x37, 0xbd, 0x44, 0x71, 0x8d, 0x5e, 0xa2,
	0xd0, 0x40, 0x07, 0x8d, 0xce, 0x28, 0x4e, 0x8b, 0x4f, 0xf2, 0x42, 0x74, 0xf1, 0xc6, 0x9d, 0xf6,
	0xa9, 0x0f, 0x41, 0x8b, 0xa4, 0x05, 0xfa, 0x44, 0xa9, 0xf6, 0xd2, 0x58, 0xb9, 0x35, 0x59, 0xad,
	0xda, 0x0a, 0x86, 0x3d, 0xf6, 0xfa, 0x81, 0x00, 0x66, 0x0f, 0x04, 0x70, 0x6c, 0x75, 0xf3, 0xf7,
	0xad, 0xee, 0x03, 0x23, 0x0b, 0x1e, 0x18, 0xd9, 0xcf, 0x30, 0x2b, 0x84, 0xcd, 0x65, 0x69, 0x22,
	0x88, 0xc7, 0x49, 0x78, 0xf9, 0xdd, 0x29, 0x9d, 0x7e, 0x50, 0xc0, 0x57, 0x74, 0x8a, 0x77, 0xa7,
	0xd7, 0x7f, 0x8d, 0xe0, 0xd1, 0x47, 0x29, 0xa8, 0x13, 0x57, 0x3d, 0x47, 0x1d, 0x44, 0xbc, 0x24,
	0xd8, 0xb1, 0xd2, 0x02, 0x8d, 0x40, 0x5a, 0x51, 0xe1, 0xbe, 0xe7, 0x8c, 0x00, 0xc3, 0xb4, 0x60,
	0x8f, 0x60, 0x6a, 0xde, 0x1e, 0x86, 0xe2, 0xfb, 0xe6, 0xed, 0x21, 0x2d, 0x70, 0xce, 0xeb, 0x43,
	0x45, 0x25, 0xf7, 0x39, 0x2e, 0xd7, 0xf7, 0xb0, 0x1a, 0xae, 0x60, 0x1a, 0x55, 0x1b, 0xc1, 0x5e,
	0x80, 0x2f, 0xb4, 0x56, 0xce, 0x46, 0xc2, 0xcb, 0xa7, 0x27, 0x1d, 0x0d, 0x49, 0xdc, 0x71, 0xd9,
	0xf7, 0x30, 0x21, 0x17, 0xf4, 0xe9, 0xcc, 0x93, 0x53, 0x67, 0xdc, 0xd8, 0x22, 0xf3, 0x66, 0x4a,
	0xff, 0xda, 0x17, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x47, 0x28, 0x57, 0xb3, 0x07, 0x00,
	0x00,
}