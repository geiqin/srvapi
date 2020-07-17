// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shipment.proto

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

type ShipmentWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting              string   `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Keywords             string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id                   int64    `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Type                 string   `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Method               string   `protobuf:"bytes,8,opt,name=method,proto3" json:"method,omitempty"`
	OrderId              int64    `protobuf:"varint,9,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipmentWhere) Reset()         { *m = ShipmentWhere{} }
func (m *ShipmentWhere) String() string { return proto.CompactTextString(m) }
func (*ShipmentWhere) ProtoMessage()    {}
func (*ShipmentWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_832505ff9d8a41db, []int{0}
}

func (m *ShipmentWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipmentWhere.Unmarshal(m, b)
}
func (m *ShipmentWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipmentWhere.Marshal(b, m, deterministic)
}
func (m *ShipmentWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipmentWhere.Merge(m, src)
}
func (m *ShipmentWhere) XXX_Size() int {
	return xxx_messageInfo_ShipmentWhere.Size(m)
}
func (m *ShipmentWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipmentWhere.DiscardUnknown(m)
}

var xxx_messageInfo_ShipmentWhere proto.InternalMessageInfo

func (m *ShipmentWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *ShipmentWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ShipmentWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
}

func (m *ShipmentWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *ShipmentWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ShipmentWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ShipmentWhere) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ShipmentWhere) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ShipmentWhere) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type Shipment struct {
	Id                   int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DeliverySn           string            `protobuf:"bytes,2,opt,name=delivery_sn,json=deliverySn,proto3" json:"delivery_sn,omitempty"`
	OrderId              int64             `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Method               string            `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	Type                 string            `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Freight              float32           `protobuf:"fixed32,6,opt,name=freight,proto3" json:"freight,omitempty"`
	Protected            bool              `protobuf:"varint,7,opt,name=protected,proto3" json:"protected,omitempty"`
	IsDelivery           bool              `protobuf:"varint,8,opt,name=is_delivery,json=isDelivery,proto3" json:"is_delivery,omitempty"`
	ShipperId            int32             `protobuf:"varint,9,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	LogisticsNo          string            `protobuf:"bytes,10,opt,name=logistics_no,json=logisticsNo,proto3" json:"logistics_no,omitempty"`
	FetchCode            string            `protobuf:"bytes,11,opt,name=fetch_code,json=fetchCode,proto3" json:"fetch_code,omitempty"`
	LocationId           int64             `protobuf:"varint,12,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	FetchAt              string            `protobuf:"bytes,13,opt,name=fetch_at,json=fetchAt,proto3" json:"fetch_at,omitempty"`
	DeliveryAt           string            `protobuf:"bytes,14,opt,name=delivery_at,json=deliveryAt,proto3" json:"delivery_at,omitempty"`
	CustomerId           int64             `protobuf:"varint,15,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ReceiverName         string            `protobuf:"bytes,16,opt,name=receiver_name,json=receiverName,proto3" json:"receiver_name,omitempty"`
	ReceiverAreaId       int64             `protobuf:"varint,17,opt,name=receiver_area_id,json=receiverAreaId,proto3" json:"receiver_area_id,omitempty"`
	ReceiverAddr         string            `protobuf:"bytes,18,opt,name=receiver_addr,json=receiverAddr,proto3" json:"receiver_addr,omitempty"`
	ReceiverZip          string            `protobuf:"bytes,19,opt,name=receiver_zip,json=receiverZip,proto3" json:"receiver_zip,omitempty"`
	ReceiverTel          string            `protobuf:"bytes,20,opt,name=receiver_tel,json=receiverTel,proto3" json:"receiver_tel,omitempty"`
	ReceiverMobile       string            `protobuf:"bytes,21,opt,name=receiver_mobile,json=receiverMobile,proto3" json:"receiver_mobile,omitempty"`
	ReceiverEmail        string            `protobuf:"bytes,22,opt,name=receiver_email,json=receiverEmail,proto3" json:"receiver_email,omitempty"`
	OpId                 int64             `protobuf:"varint,23,opt,name=op_id,json=opId,proto3" json:"op_id,omitempty"`
	Status               string            `protobuf:"bytes,24,opt,name=status,proto3" json:"status,omitempty"`
	Memo                 string            `protobuf:"bytes,25,opt,name=memo,proto3" json:"memo,omitempty"`
	ArrivedAt            string            `protobuf:"bytes,26,opt,name=arrived_at,json=arrivedAt,proto3" json:"arrived_at,omitempty"`
	CreatedAt            string            `protobuf:"bytes,27,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string            `protobuf:"bytes,28,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Details              []*ShipmentDetail `protobuf:"bytes,29,rep,name=details,proto3" json:"details,omitempty"`
	Shipper              *Shipper          `protobuf:"bytes,30,opt,name=shipper,proto3" json:"shipper,omitempty"`
	StartTime            string            `protobuf:"bytes,31,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              string            `protobuf:"bytes,32,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Shipment) Reset()         { *m = Shipment{} }
func (m *Shipment) String() string { return proto.CompactTextString(m) }
func (*Shipment) ProtoMessage()    {}
func (*Shipment) Descriptor() ([]byte, []int) {
	return fileDescriptor_832505ff9d8a41db, []int{1}
}

func (m *Shipment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shipment.Unmarshal(m, b)
}
func (m *Shipment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shipment.Marshal(b, m, deterministic)
}
func (m *Shipment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shipment.Merge(m, src)
}
func (m *Shipment) XXX_Size() int {
	return xxx_messageInfo_Shipment.Size(m)
}
func (m *Shipment) XXX_DiscardUnknown() {
	xxx_messageInfo_Shipment.DiscardUnknown(m)
}

var xxx_messageInfo_Shipment proto.InternalMessageInfo

func (m *Shipment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Shipment) GetDeliverySn() string {
	if m != nil {
		return m.DeliverySn
	}
	return ""
}

func (m *Shipment) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Shipment) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Shipment) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Shipment) GetFreight() float32 {
	if m != nil {
		return m.Freight
	}
	return 0
}

func (m *Shipment) GetProtected() bool {
	if m != nil {
		return m.Protected
	}
	return false
}

func (m *Shipment) GetIsDelivery() bool {
	if m != nil {
		return m.IsDelivery
	}
	return false
}

func (m *Shipment) GetShipperId() int32 {
	if m != nil {
		return m.ShipperId
	}
	return 0
}

func (m *Shipment) GetLogisticsNo() string {
	if m != nil {
		return m.LogisticsNo
	}
	return ""
}

func (m *Shipment) GetFetchCode() string {
	if m != nil {
		return m.FetchCode
	}
	return ""
}

func (m *Shipment) GetLocationId() int64 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

func (m *Shipment) GetFetchAt() string {
	if m != nil {
		return m.FetchAt
	}
	return ""
}

func (m *Shipment) GetDeliveryAt() string {
	if m != nil {
		return m.DeliveryAt
	}
	return ""
}

func (m *Shipment) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *Shipment) GetReceiverName() string {
	if m != nil {
		return m.ReceiverName
	}
	return ""
}

func (m *Shipment) GetReceiverAreaId() int64 {
	if m != nil {
		return m.ReceiverAreaId
	}
	return 0
}

func (m *Shipment) GetReceiverAddr() string {
	if m != nil {
		return m.ReceiverAddr
	}
	return ""
}

func (m *Shipment) GetReceiverZip() string {
	if m != nil {
		return m.ReceiverZip
	}
	return ""
}

func (m *Shipment) GetReceiverTel() string {
	if m != nil {
		return m.ReceiverTel
	}
	return ""
}

func (m *Shipment) GetReceiverMobile() string {
	if m != nil {
		return m.ReceiverMobile
	}
	return ""
}

func (m *Shipment) GetReceiverEmail() string {
	if m != nil {
		return m.ReceiverEmail
	}
	return ""
}

func (m *Shipment) GetOpId() int64 {
	if m != nil {
		return m.OpId
	}
	return 0
}

func (m *Shipment) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Shipment) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Shipment) GetArrivedAt() string {
	if m != nil {
		return m.ArrivedAt
	}
	return ""
}

func (m *Shipment) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Shipment) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Shipment) GetDetails() []*ShipmentDetail {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Shipment) GetShipper() *Shipper {
	if m != nil {
		return m.Shipper
	}
	return nil
}

func (m *Shipment) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Shipment) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

type ShipmentDetail struct {
	Id                   int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ShipmentId           int64      `protobuf:"varint,2,opt,name=shipment_id,json=shipmentId,proto3" json:"shipment_id,omitempty"`
	OrderDetailId        int64      `protobuf:"varint,3,opt,name=order_detail_id,json=orderDetailId,proto3" json:"order_detail_id,omitempty"`
	ItemId               int64      `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64      `protobuf:"varint,5,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Num                  int32      `protobuf:"varint,6,opt,name=num,proto3" json:"num,omitempty"`
	CreatedAt            string     `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string     `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Goods                *GoodsInfo `protobuf:"bytes,9,opt,name=goods,proto3" json:"goods,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ShipmentDetail) Reset()         { *m = ShipmentDetail{} }
func (m *ShipmentDetail) String() string { return proto.CompactTextString(m) }
func (*ShipmentDetail) ProtoMessage()    {}
func (*ShipmentDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_832505ff9d8a41db, []int{2}
}

func (m *ShipmentDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipmentDetail.Unmarshal(m, b)
}
func (m *ShipmentDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipmentDetail.Marshal(b, m, deterministic)
}
func (m *ShipmentDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipmentDetail.Merge(m, src)
}
func (m *ShipmentDetail) XXX_Size() int {
	return xxx_messageInfo_ShipmentDetail.Size(m)
}
func (m *ShipmentDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipmentDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ShipmentDetail proto.InternalMessageInfo

func (m *ShipmentDetail) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ShipmentDetail) GetShipmentId() int64 {
	if m != nil {
		return m.ShipmentId
	}
	return 0
}

func (m *ShipmentDetail) GetOrderDetailId() int64 {
	if m != nil {
		return m.OrderDetailId
	}
	return 0
}

func (m *ShipmentDetail) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ShipmentDetail) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *ShipmentDetail) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *ShipmentDetail) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ShipmentDetail) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *ShipmentDetail) GetGoods() *GoodsInfo {
	if m != nil {
		return m.Goods
	}
	return nil
}

type ShipmentResponse struct {
	Entity               *Shipment   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Shipment `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ShipmentResponse) Reset()         { *m = ShipmentResponse{} }
func (m *ShipmentResponse) String() string { return proto.CompactTextString(m) }
func (*ShipmentResponse) ProtoMessage()    {}
func (*ShipmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_832505ff9d8a41db, []int{3}
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

func (m *ShipmentResponse) GetEntity() *Shipment {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ShipmentResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ShipmentResponse) GetItems() []*Shipment {
	if m != nil {
		return m.Items
	}
	return nil
}

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
	proto.RegisterType((*ShipmentWhere)(nil), "geiqin.srv.tms.ShipmentWhere")
	proto.RegisterType((*Shipment)(nil), "geiqin.srv.tms.Shipment")
	proto.RegisterType((*ShipmentDetail)(nil), "geiqin.srv.tms.ShipmentDetail")
	proto.RegisterType((*ShipmentResponse)(nil), "geiqin.srv.tms.ShipmentResponse")
}

func init() {
	proto.RegisterFile("shipment.proto", fileDescriptor_832505ff9d8a41db)
}

var fileDescriptor_832505ff9d8a41db = []byte{
	// 965 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xc6, 0xff, 0xf6, 0x71, 0xe2, 0x84, 0x69, 0xd2, 0x4c, 0xd2, 0xa6, 0x35, 0x41, 0x80, 0x25,
	0x24, 0x03, 0xe6, 0x86, 0xdb, 0xa8, 0x29, 0x65, 0x25, 0x5a, 0xa1, 0x4d, 0x11, 0x12, 0x37, 0xab,
	0xed, 0xce, 0x89, 0x33, 0x8a, 0x77, 0x67, 0x99, 0x19, 0x07, 0x25, 0x6f, 0xc0, 0x2b, 0xf0, 0x76,
	0xbc, 0x05, 0x17, 0x5c, 0xa0, 0x39, 0xb3, 0xb3, 0xed, 0x9a, 0x04, 0x55, 0x82, 0xbb, 0x39, 0xdf,
	0xf9, 0xe6, 0x9b, 0xf3, 0xe7, 0xb3, 0x86, 0x89, 0xb9, 0x94, 0x65, 0x8e, 0x85, 0x9d, 0x97, 0x5a,
	0x59, 0xc5, 0x26, 0x4b, 0x94, 0xbf, 0xc8, 0x62, 0x6e, 0xf4, 0xf5, 0xdc, 0xe6, 0xe6, 0x68, 0x2b,
	0x53, 0x79, 0xae, 0x0a, 0xef, 0x3d, 0xda, 0x76, 0xec, 0x12, 0x75, 0x65, 0xee, 0x2c, 0x95, 0x12,
	0x26, 0x2a, 0x2e, 0x94, 0x07, 0x4e, 0xfe, 0x68, 0xc1, 0xf6, 0x79, 0x25, 0xf8, 0xd3, 0x25, 0x6a,
	0x64, 0x7b, 0xd0, 0x2b, 0xd3, 0x25, 0x0a, 0xde, 0x9a, 0xb6, 0x66, 0xbd, 0xd8, 0x1b, 0xec, 0x11,
	0x8c, 0xdc, 0x21, 0x31, 0xf2, 0x16, 0x79, 0x9b, 0x3c, 0x43, 0x07, 0x9c, 0xcb, 0x5b, 0x64, 0x1c,
	0x06, 0x46, 0x69, 0x2b, 0x8b, 0x25, 0xef, 0x4c, 0x5b, 0xb3, 0x51, 0x1c, 0x4c, 0x76, 0x04, 0xc3,
	0x2b, 0xbc, 0xf9, 0x55, 0x69, 0x61, 0x78, 0x97, 0x5c, 0xb5, 0xcd, 0x26, 0xd0, 0x96, 0x82, 0xf7,
	0xa6, 0xad, 0x59, 0x27, 0x6e, 0x4b, 0xc1, 0x76, 0xa1, 0x23, 0x85, 0xe1, 0xfd, 0x69, 0x67, 0xd6,
	0x89, 0xdd, 0x91, 0x31, 0xe8, 0xda, 0x9b, 0x12, 0xf9, 0x80, 0x6e, 0xd2, 0x99, 0x3d, 0x84, 0x7e,
	0x8e, 0xf6, 0x52, 0x09, 0x3e, 0x24, 0xb4, 0xb2, 0xd8, 0x21, 0x0c, 0x95, 0x16, 0xa8, 0x13, 0x29,
	0xf8, 0x88, 0x34, 0x07, 0x64, 0x47, 0xe2, 0xe4, 0xb7, 0x21, 0x0c, 0x43, 0x8e, 0xd5, 0xab, 0xad,
	0xfa, 0xd5, 0xa7, 0x30, 0x16, 0xb8, 0x92, 0xd7, 0xa8, 0x6f, 0x12, 0x53, 0x50, 0x6a, 0xa3, 0x18,
	0x02, 0x74, 0x5e, 0x34, 0x84, 0x3b, 0x0d, 0xe1, 0x77, 0x62, 0xe9, 0x36, 0x62, 0x09, 0x71, 0xf7,
	0xde, 0x89, 0x9b, 0xc3, 0xe0, 0x42, 0xa3, 0x5c, 0x5e, 0x5a, 0xde, 0x9f, 0xb6, 0x66, 0xed, 0x38,
	0x98, 0xec, 0x31, 0x8c, 0x5c, 0x2f, 0x30, 0xb3, 0x28, 0x28, 0xd5, 0x61, 0xfc, 0x16, 0x70, 0xf1,
	0x49, 0x93, 0x84, 0x78, 0x28, 0xe9, 0x61, 0x0c, 0xd2, 0x9c, 0x55, 0x08, 0x3b, 0x06, 0xa8, 0x7a,
	0x1c, 0x52, 0xef, 0xc5, 0xa3, 0x0a, 0x89, 0x04, 0xfb, 0x08, 0xb6, 0x56, 0x6a, 0x29, 0x8d, 0x95,
	0x99, 0x49, 0x0a, 0xc5, 0x81, 0x62, 0x1a, 0xd7, 0xd8, 0x2b, 0xe5, 0x14, 0x2e, 0xd0, 0x66, 0x97,
	0x49, 0xa6, 0x04, 0xf2, 0x31, 0x11, 0x46, 0x84, 0x3c, 0x53, 0x02, 0x5d, 0x04, 0x2b, 0x95, 0xa5,
	0x56, 0xaa, 0xc2, 0xbd, 0xb0, 0x45, 0x35, 0x80, 0x00, 0x45, 0x54, 0x7a, 0x7f, 0x3f, 0xb5, 0x7c,
	0xdb, 0xf7, 0x9f, 0xec, 0x53, 0xdb, 0xa8, 0x6e, 0x6a, 0xf9, 0xa4, 0x59, 0x5d, 0x4f, 0xc8, 0xd6,
	0xc6, 0xaa, 0xdc, 0x87, 0xbf, 0xe3, 0xc5, 0x03, 0x14, 0x09, 0xf6, 0x31, 0x6c, 0x6b, 0xcc, 0xd0,
	0xf1, 0x93, 0x22, 0xcd, 0x91, 0xef, 0x92, 0xc6, 0x56, 0x00, 0x5f, 0xa5, 0x39, 0xb2, 0x19, 0xec,
	0xd6, 0xa4, 0x54, 0x63, 0xea, 0xa4, 0x3e, 0x24, 0xa9, 0x49, 0xc0, 0x4f, 0x35, 0xa6, 0x1b, 0x72,
	0xa9, 0x10, 0x9a, 0xb3, 0xa6, 0xdc, 0xa9, 0x10, 0xda, 0xd5, 0xac, 0x26, 0xdd, 0xca, 0x92, 0x3f,
	0xf0, 0x35, 0x0b, 0xd8, 0xcf, 0xb2, 0x6c, 0x50, 0x2c, 0xae, 0xf8, 0x5e, 0x93, 0xf2, 0x1a, 0x57,
	0xec, 0x33, 0xd8, 0xa9, 0x29, 0xb9, 0x7a, 0x23, 0x57, 0xc8, 0xf7, 0x89, 0x55, 0xc7, 0xf4, 0x92,
	0x50, 0xf6, 0x09, 0xd4, 0x48, 0x82, 0x79, 0x2a, 0x57, 0xfc, 0x21, 0xf1, 0xea, 0x48, 0x9f, 0x3b,
	0x90, 0x3d, 0x80, 0x9e, 0x2a, 0x5d, 0x66, 0x07, 0x94, 0x59, 0x57, 0x95, 0x7e, 0x04, 0x8d, 0x4d,
	0xed, 0xda, 0x70, 0xee, 0x47, 0xd0, 0x5b, 0x6e, 0x04, 0x73, 0xcc, 0x15, 0x3f, 0xf4, 0x23, 0xe8,
	0xce, 0xae, 0xcf, 0xa9, 0xd6, 0xf2, 0x1a, 0x85, 0xeb, 0xc5, 0x91, 0xef, 0x73, 0x85, 0x9c, 0x5a,
	0xe7, 0xce, 0x34, 0xa6, 0xd6, 0xbb, 0x1f, 0x79, 0x77, 0x85, 0x78, 0xf7, 0xba, 0x14, 0xc1, 0xfd,
	0xd8, 0xbb, 0x2b, 0xe4, 0xd4, 0xb2, 0x6f, 0x60, 0x20, 0xd0, 0xa6, 0x72, 0x65, 0xf8, 0xf1, 0xb4,
	0x33, 0x1b, 0x2f, 0x9e, 0xcc, 0x9b, 0x8b, 0x69, 0x1e, 0x7e, 0x82, 0x67, 0x44, 0x8b, 0x03, 0x9d,
	0x7d, 0x05, 0x83, 0x6a, 0x5c, 0xf9, 0x93, 0x69, 0x6b, 0x36, 0x5e, 0x1c, 0xdc, 0x75, 0xb3, 0x44,
	0x1d, 0x07, 0x1e, 0xcd, 0xbc, 0x4d, 0xb5, 0x4d, 0xac, 0xcc, 0x91, 0x3f, 0xf5, 0xb1, 0x10, 0xf2,
	0x5a, 0xe6, 0xe8, 0x06, 0x12, 0x0b, 0xe1, 0x9d, 0x53, 0x3f, 0x90, 0x58, 0x08, 0xe7, 0x3a, 0xf9,
	0xbd, 0x0d, 0x93, 0x66, 0x20, 0x77, 0x6d, 0x84, 0xb0, 0x62, 0x5d, 0xb5, 0xdb, 0x7e, 0x24, 0x03,
	0x14, 0x09, 0xf6, 0x29, 0xec, 0xf8, 0x8d, 0xe0, 0x33, 0x78, 0xbb, 0x18, 0xb6, 0x09, 0xf6, 0xb2,
	0x91, 0x60, 0x07, 0x30, 0x90, 0x16, 0x73, 0xe7, 0xef, 0x92, 0xbf, 0xef, 0xcc, 0x48, 0xb0, 0x7d,
	0xe8, 0x9b, 0xab, 0x75, 0x52, 0x6f, 0xbf, 0x9e, 0xb9, 0x5a, 0x47, 0xb4, 0x00, 0x8b, 0x75, 0x4e,
	0xeb, 0xa1, 0x17, 0xbb, 0xe3, 0x46, 0x4b, 0x06, 0xff, 0xde, 0x92, 0xe1, 0x66, 0x4b, 0xbe, 0x80,
	0x1e, 0xad, 0x7b, 0x5a, 0x0a, 0xe3, 0xc5, 0xe1, 0x66, 0x59, 0x5f, 0x84, 0x6f, 0x41, 0xec, 0x79,
	0x27, 0x7f, 0xb6, 0x60, 0x37, 0x14, 0x27, 0x46, 0x53, 0xaa, 0xc2, 0x20, 0xfb, 0x12, 0xfa, 0x58,
	0x58, 0x69, 0x6f, 0xa8, 0x44, 0xe3, 0x05, 0xbf, 0xaf, 0xaf, 0x71, 0xc5, 0x63, 0x9f, 0xfb, 0x2f,
	0x88, 0xa6, 0xd2, 0x8d, 0x17, 0xfb, 0x9b, 0x17, 0x7e, 0x70, 0x4e, 0xff, 0x61, 0xd1, 0x6c, 0x0e,
	0x3d, 0x57, 0x15, 0xc3, 0x3b, 0x34, 0x35, 0xf7, 0xab, 0x7b, 0x9a, 0x13, 0x47, 0xad, 0x95, 0xa6,
	0x92, 0xde, 0x21, 0xfe, 0xdc, 0x39, 0x63, 0xcf, 0x61, 0x33, 0xe8, 0xca, 0xe2, 0x42, 0x51, 0x99,
	0xc7, 0x8b, 0xbd, 0x4d, 0x2e, 0xe5, 0x4e, 0x8c, 0xc5, 0x5f, 0x6d, 0xd8, 0x09, 0x4f, 0x9d, 0xa3,
	0xbe, 0x96, 0x19, 0xb2, 0x33, 0x18, 0xbc, 0x94, 0x4b, 0x9d, 0x5a, 0x64, 0xff, 0x7c, 0x26, 0x2f,
	0xed, 0xcd, 0xd1, 0xf4, 0xde, 0x68, 0xab, 0xea, 0x9d, 0x7c, 0xc0, 0xbe, 0x85, 0xfe, 0x33, 0xea,
	0x18, 0xbb, 0x37, 0xb7, 0xf7, 0xd5, 0xf9, 0x91, 0x5a, 0xfb, 0x1f, 0x75, 0xbe, 0x83, 0xce, 0x0b,
	0xb4, 0xec, 0xf8, 0x3e, 0x2a, 0xfd, 0x0b, 0x78, 0x2f, 0xa5, 0x08, 0xba, 0xdf, 0x4b, 0xf3, 0x7f,
	0x48, 0xbd, 0xe9, 0xd3, 0xbf, 0x91, 0xaf, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xde, 0x70, 0x4e,
	0xf0, 0xdd, 0x08, 0x00, 0x00,
}
