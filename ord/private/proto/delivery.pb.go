// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delivery.proto

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

type Delivery struct {
	Id                   int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	OrderId              int64             `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	DeliveryType         string            `protobuf:"bytes,4,opt,name=delivery_type,json=deliveryType,proto3" json:"delivery_type,omitempty"`
	DeliverySn           string            `protobuf:"bytes,5,opt,name=delivery_sn,json=deliverySn,proto3" json:"delivery_sn,omitempty"`
	Freight              float32           `protobuf:"fixed32,6,opt,name=freight,proto3" json:"freight,omitempty"`
	Protected            bool              `protobuf:"varint,7,opt,name=protected,proto3" json:"protected,omitempty"`
	ExpressId            int32             `protobuf:"varint,8,opt,name=express_id,json=expressId,proto3" json:"express_id,omitempty"`
	ExpressName          string            `protobuf:"bytes,9,opt,name=express_name,json=expressName,proto3" json:"express_name,omitempty"`
	ExpressCode          string            `protobuf:"bytes,10,opt,name=express_code,json=expressCode,proto3" json:"express_code,omitempty"`
	ExpressNo            string            `protobuf:"bytes,11,opt,name=express_no,json=expressNo,proto3" json:"express_no,omitempty"`
	CustomerId           int64             `protobuf:"varint,12,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ReceiverName         string            `protobuf:"bytes,13,opt,name=receiver_name,json=receiverName,proto3" json:"receiver_name,omitempty"`
	ReceiverAreaId       int64             `protobuf:"varint,14,opt,name=receiver_area_id,json=receiverAreaId,proto3" json:"receiver_area_id,omitempty"`
	ReceiverAddr         string            `protobuf:"bytes,15,opt,name=receiver_addr,json=receiverAddr,proto3" json:"receiver_addr,omitempty"`
	ReceiverZip          string            `protobuf:"bytes,16,opt,name=receiver_zip,json=receiverZip,proto3" json:"receiver_zip,omitempty"`
	ReceiverTel          string            `protobuf:"bytes,17,opt,name=receiver_tel,json=receiverTel,proto3" json:"receiver_tel,omitempty"`
	ReceiverMobile       string            `protobuf:"bytes,18,opt,name=receiver_mobile,json=receiverMobile,proto3" json:"receiver_mobile,omitempty"`
	ReceiverEmail        string            `protobuf:"bytes,19,opt,name=receiver_email,json=receiverEmail,proto3" json:"receiver_email,omitempty"`
	OpId                 int64             `protobuf:"varint,20,opt,name=op_id,json=opId,proto3" json:"op_id,omitempty"`
	Status               string            `protobuf:"bytes,21,opt,name=status,proto3" json:"status,omitempty"`
	Memo                 string            `protobuf:"bytes,22,opt,name=memo,proto3" json:"memo,omitempty"`
	ArrivedAt            string            `protobuf:"bytes,23,opt,name=arrived_at,json=arrivedAt,proto3" json:"arrived_at,omitempty"`
	CreatedAt            string            `protobuf:"bytes,24,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string            `protobuf:"bytes,25,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Details              []*DeliveryDetail `protobuf:"bytes,26,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Delivery) Reset()         { *m = Delivery{} }
func (m *Delivery) String() string { return proto.CompactTextString(m) }
func (*Delivery) ProtoMessage()    {}
func (*Delivery) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf387bcb4e23d880, []int{0}
}

func (m *Delivery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Delivery.Unmarshal(m, b)
}
func (m *Delivery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Delivery.Marshal(b, m, deterministic)
}
func (m *Delivery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delivery.Merge(m, src)
}
func (m *Delivery) XXX_Size() int {
	return xxx_messageInfo_Delivery.Size(m)
}
func (m *Delivery) XXX_DiscardUnknown() {
	xxx_messageInfo_Delivery.DiscardUnknown(m)
}

var xxx_messageInfo_Delivery proto.InternalMessageInfo

func (m *Delivery) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Delivery) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Delivery) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Delivery) GetDeliveryType() string {
	if m != nil {
		return m.DeliveryType
	}
	return ""
}

func (m *Delivery) GetDeliverySn() string {
	if m != nil {
		return m.DeliverySn
	}
	return ""
}

func (m *Delivery) GetFreight() float32 {
	if m != nil {
		return m.Freight
	}
	return 0
}

func (m *Delivery) GetProtected() bool {
	if m != nil {
		return m.Protected
	}
	return false
}

func (m *Delivery) GetExpressId() int32 {
	if m != nil {
		return m.ExpressId
	}
	return 0
}

func (m *Delivery) GetExpressName() string {
	if m != nil {
		return m.ExpressName
	}
	return ""
}

func (m *Delivery) GetExpressCode() string {
	if m != nil {
		return m.ExpressCode
	}
	return ""
}

func (m *Delivery) GetExpressNo() string {
	if m != nil {
		return m.ExpressNo
	}
	return ""
}

func (m *Delivery) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *Delivery) GetReceiverName() string {
	if m != nil {
		return m.ReceiverName
	}
	return ""
}

func (m *Delivery) GetReceiverAreaId() int64 {
	if m != nil {
		return m.ReceiverAreaId
	}
	return 0
}

func (m *Delivery) GetReceiverAddr() string {
	if m != nil {
		return m.ReceiverAddr
	}
	return ""
}

func (m *Delivery) GetReceiverZip() string {
	if m != nil {
		return m.ReceiverZip
	}
	return ""
}

func (m *Delivery) GetReceiverTel() string {
	if m != nil {
		return m.ReceiverTel
	}
	return ""
}

func (m *Delivery) GetReceiverMobile() string {
	if m != nil {
		return m.ReceiverMobile
	}
	return ""
}

func (m *Delivery) GetReceiverEmail() string {
	if m != nil {
		return m.ReceiverEmail
	}
	return ""
}

func (m *Delivery) GetOpId() int64 {
	if m != nil {
		return m.OpId
	}
	return 0
}

func (m *Delivery) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Delivery) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Delivery) GetArrivedAt() string {
	if m != nil {
		return m.ArrivedAt
	}
	return ""
}

func (m *Delivery) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Delivery) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Delivery) GetDetails() []*DeliveryDetail {
	if m != nil {
		return m.Details
	}
	return nil
}

type DeliveryDetail struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DeliveryId           int64    `protobuf:"varint,2,opt,name=delivery_id,json=deliveryId,proto3" json:"delivery_id,omitempty"`
	OrderDetailId        int64    `protobuf:"varint,3,opt,name=order_detail_id,json=orderDetailId,proto3" json:"order_detail_id,omitempty"`
	ItemId               int64    `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64    `protobuf:"varint,5,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Num                  int32    `protobuf:"varint,6,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliveryDetail) Reset()         { *m = DeliveryDetail{} }
func (m *DeliveryDetail) String() string { return proto.CompactTextString(m) }
func (*DeliveryDetail) ProtoMessage()    {}
func (*DeliveryDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf387bcb4e23d880, []int{1}
}

func (m *DeliveryDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryDetail.Unmarshal(m, b)
}
func (m *DeliveryDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryDetail.Marshal(b, m, deterministic)
}
func (m *DeliveryDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryDetail.Merge(m, src)
}
func (m *DeliveryDetail) XXX_Size() int {
	return xxx_messageInfo_DeliveryDetail.Size(m)
}
func (m *DeliveryDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryDetail.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryDetail proto.InternalMessageInfo

func (m *DeliveryDetail) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeliveryDetail) GetDeliveryId() int64 {
	if m != nil {
		return m.DeliveryId
	}
	return 0
}

func (m *DeliveryDetail) GetOrderDetailId() int64 {
	if m != nil {
		return m.OrderDetailId
	}
	return 0
}

func (m *DeliveryDetail) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *DeliveryDetail) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *DeliveryDetail) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type DeliveryResponse struct {
	Entity               *Delivery   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Delivery `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeliveryResponse) Reset()         { *m = DeliveryResponse{} }
func (m *DeliveryResponse) String() string { return proto.CompactTextString(m) }
func (*DeliveryResponse) ProtoMessage()    {}
func (*DeliveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf387bcb4e23d880, []int{2}
}

func (m *DeliveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryResponse.Unmarshal(m, b)
}
func (m *DeliveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryResponse.Marshal(b, m, deterministic)
}
func (m *DeliveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryResponse.Merge(m, src)
}
func (m *DeliveryResponse) XXX_Size() int {
	return xxx_messageInfo_DeliveryResponse.Size(m)
}
func (m *DeliveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryResponse proto.InternalMessageInfo

func (m *DeliveryResponse) GetEntity() *Delivery {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *DeliveryResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *DeliveryResponse) GetItems() []*Delivery {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *DeliveryResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DeliveryResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Delivery)(nil), "geiqin.srv.ord.private.Delivery")
	proto.RegisterType((*DeliveryDetail)(nil), "geiqin.srv.ord.private.DeliveryDetail")
	proto.RegisterType((*DeliveryResponse)(nil), "geiqin.srv.ord.private.DeliveryResponse")
}

func init() { proto.RegisterFile("delivery.proto", fileDescriptor_bf387bcb4e23d880) }

var fileDescriptor_bf387bcb4e23d880 = []byte{
	// 737 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdf, 0x6e, 0xfb, 0x34,
	0x14, 0xc7, 0x69, 0xd3, 0xa4, 0xed, 0x69, 0xd7, 0x16, 0x8f, 0x6d, 0xde, 0xb4, 0x89, 0xae, 0x88,
	0x91, 0xab, 0x0a, 0x75, 0x12, 0xe2, 0x92, 0xb2, 0x4d, 0x28, 0x17, 0x4c, 0x28, 0x1d, 0x4c, 0xe2,
	0xa6, 0xf2, 0xe2, 0xb3, 0xcd, 0x5a, 0x13, 0x07, 0xc7, 0xad, 0x18, 0x4f, 0xc3, 0x3d, 0x0f, 0xc2,
	0x1b, 0xf0, 0x3c, 0xc8, 0x76, 0xd2, 0xad, 0xc0, 0xa8, 0x7e, 0xd2, 0xee, 0xec, 0xef, 0xf9, 0x9c,
	0xbf, 0xf1, 0x51, 0xa0, 0xc7, 0x71, 0x21, 0x56, 0xa8, 0x9e, 0xc7, 0xb9, 0x92, 0x5a, 0x92, 0xfd,
	0x07, 0x14, 0xbf, 0x88, 0x6c, 0x5c, 0xa8, 0xd5, 0x58, 0x2a, 0x3e, 0xce, 0x95, 0x58, 0x31, 0x8d,
	0x47, 0xdd, 0x44, 0xa6, 0xa9, 0xcc, 0x1c, 0x35, 0xfa, 0x2b, 0x80, 0xd6, 0x65, 0xe9, 0x48, 0x7a,
	0x50, 0x17, 0x9c, 0xd6, 0x86, 0xb5, 0xd0, 0x8b, 0xeb, 0x82, 0x13, 0x02, 0x0d, 0xfd, 0x9c, 0x23,
	0xad, 0x0f, 0x6b, 0x61, 0x3b, 0xb6, 0x67, 0x72, 0x08, 0x2d, 0xa9, 0x38, 0xaa, 0xb9, 0xe0, 0xd4,
	0xb3, 0x64, 0xd3, 0xde, 0x23, 0x4e, 0x3e, 0x83, 0x9d, 0xaa, 0x86, 0xb9, 0xf5, 0x6b, 0x58, 0xbf,
	0x6e, 0x25, 0xde, 0x18, 0xff, 0x4f, 0xa1, 0xb3, 0x86, 0x8a, 0x8c, 0xfa, 0x16, 0x81, 0x4a, 0x9a,
	0x65, 0x84, 0x42, 0xf3, 0x5e, 0xa1, 0x78, 0x78, 0xd4, 0x34, 0x18, 0xd6, 0xc2, 0x7a, 0x5c, 0x5d,
	0xc9, 0x31, 0xb4, 0x4d, 0xd1, 0x98, 0x68, 0xe4, 0xb4, 0x39, 0xac, 0x85, 0xad, 0xf8, 0x45, 0x20,
	0x27, 0x00, 0xf8, 0x6b, 0xae, 0xb0, 0x28, 0x4c, 0x69, 0xad, 0x61, 0x2d, 0xf4, 0xe3, 0x76, 0xa9,
	0x44, 0x9c, 0x9c, 0x42, 0xb7, 0x32, 0x67, 0x2c, 0x45, 0xda, 0xb6, 0x89, 0x3b, 0xa5, 0x76, 0xcd,
	0x52, 0x7c, 0x8d, 0x24, 0x92, 0x23, 0x85, 0x0d, 0xe4, 0x42, 0x72, 0x7c, 0x9d, 0x24, 0x93, 0xb4,
	0x63, 0x81, 0x2a, 0xc9, 0xb5, 0x34, 0xcd, 0x25, 0xcb, 0x42, 0xcb, 0xd4, 0xcd, 0xa7, 0x6b, 0xe7,
	0x03, 0x95, 0xe4, 0x46, 0xa4, 0x30, 0x41, 0xd3, 0xab, 0x2b, 0x63, 0xc7, 0x8d, 0xa8, 0x12, 0x6d,
	0x1d, 0x21, 0x0c, 0xd6, 0x10, 0x53, 0xc8, 0x4c, 0xa8, 0x9e, 0x0d, 0xd5, 0xab, 0xf4, 0xa9, 0x42,
	0xf6, 0x8f, 0x70, 0x8c, 0x73, 0x45, 0xfb, 0x9b, 0xe1, 0xa6, 0x9c, 0x2b, 0xd3, 0xd6, 0x1a, 0xfa,
	0x4d, 0xe4, 0x74, 0xe0, 0xda, 0xaa, 0xb4, 0x9f, 0x45, 0xbe, 0x81, 0x68, 0x5c, 0xd0, 0x8f, 0x37,
	0x91, 0x1b, 0x5c, 0x90, 0x2f, 0xa0, 0xbf, 0x46, 0x52, 0x79, 0x27, 0x16, 0x48, 0x89, 0xa5, 0xd6,
	0x35, 0x7d, 0x6f, 0x55, 0xf2, 0x39, 0xac, 0x95, 0x39, 0xa6, 0x4c, 0x2c, 0xe8, 0xae, 0xe5, 0xd6,
	0x95, 0x5e, 0x19, 0x91, 0xec, 0x82, 0x2f, 0x73, 0xd3, 0xd9, 0x27, 0xb6, 0xb3, 0x86, 0xcc, 0x23,
	0x4e, 0xf6, 0x21, 0x28, 0x34, 0xd3, 0xcb, 0x82, 0xee, 0x59, 0x9f, 0xf2, 0x66, 0x1e, 0x62, 0x8a,
	0xa9, 0xa4, 0xfb, 0xee, 0x21, 0x9a, 0xb3, 0xf9, 0x14, 0x4c, 0x29, 0xb1, 0x42, 0x3e, 0x67, 0x9a,
	0x1e, 0xb8, 0x4f, 0x51, 0x2a, 0x53, 0x6d, 0xcc, 0x89, 0x42, 0xa6, 0x9d, 0x99, 0x3a, 0x73, 0xa9,
	0x38, 0xf3, 0x32, 0xe7, 0x95, 0xf9, 0xd0, 0x99, 0x4b, 0x65, 0xaa, 0xc9, 0x37, 0xd0, 0xe4, 0xa8,
	0x99, 0x58, 0x14, 0xf4, 0x68, 0xe8, 0x85, 0x9d, 0xc9, 0xd9, 0xf8, 0xbf, 0xd7, 0x69, 0x5c, 0x2d,
	0xcf, 0xa5, 0xc5, 0xe3, 0xca, 0x6d, 0xf4, 0x47, 0x0d, 0x7a, 0x9b, 0xb6, 0x7f, 0xad, 0xd7, 0xeb,
	0x55, 0x10, 0xdc, 0x6e, 0x99, 0xf7, 0xb2, 0x0a, 0x11, 0x27, 0x67, 0xd0, 0x77, 0xbb, 0xe6, 0x82,
	0xbe, 0xac, 0xdc, 0x8e, 0x95, 0x5d, 0xd8, 0x88, 0x93, 0x03, 0x68, 0x0a, 0x8d, 0xa9, 0xb1, 0x37,
	0xac, 0x3d, 0x30, 0xd7, 0x88, 0x93, 0x3d, 0x08, 0x8a, 0xa7, 0xa5, 0xd1, 0x7d, 0xab, 0xfb, 0xc5,
	0xd3, 0x32, 0xe2, 0x64, 0x00, 0x5e, 0xb6, 0x4c, 0xed, 0x7a, 0xf9, 0xb1, 0x39, 0x8e, 0x7e, 0xaf,
	0xc3, 0xa0, 0xaa, 0x36, 0xc6, 0x22, 0x97, 0x59, 0x81, 0xe4, 0x6b, 0x08, 0x30, 0xd3, 0x42, 0x3f,
	0xdb, 0x9a, 0x3b, 0x93, 0xe1, 0xb6, 0x19, 0xc4, 0x25, 0x4f, 0xce, 0xc1, 0xcf, 0xd9, 0x03, 0x2a,
	0xdb, 0x53, 0x67, 0x72, 0xf2, 0x96, 0xe3, 0x0f, 0x06, 0x8a, 0x1d, 0x4b, 0xbe, 0x02, 0xdf, 0x94,
	0x5d, 0x50, 0xcf, 0x4e, 0x7c, 0x7b, 0x36, 0x87, 0x9b, 0x64, 0xa8, 0x94, 0x54, 0xb6, 0xf7, 0xff,
	0x49, 0x76, 0x65, 0xa0, 0xd8, 0xb1, 0xe4, 0x4b, 0x68, 0x88, 0xec, 0x5e, 0xda, 0xb9, 0x74, 0x26,
	0xc7, 0x6f, 0xf9, 0x44, 0xd9, 0xbd, 0x8c, 0x2d, 0x39, 0xf9, 0xd3, 0x83, 0x7e, 0x95, 0x7a, 0x86,
	0x6a, 0x25, 0x12, 0x24, 0x3f, 0x41, 0x70, 0x61, 0x9f, 0x14, 0xd9, 0x5a, 0xed, 0x51, 0xb8, 0xb5,
	0x9f, 0x72, 0xee, 0xa3, 0x8f, 0x4c, 0xdc, 0x1f, 0xed, 0x5b, 0x7c, 0xff, 0xb8, 0x97, 0xb8, 0xc0,
	0x77, 0x8f, 0x3b, 0x03, 0xef, 0x3b, 0xd4, 0xef, 0x1c, 0xf4, 0x16, 0x82, 0x19, 0x32, 0x95, 0x3c,
	0x92, 0xd3, 0xb7, 0xbc, 0xbe, 0x65, 0x05, 0xde, 0x3e, 0xa2, 0xc2, 0x0f, 0x09, 0x7c, 0x17, 0xd8,
	0x5f, 0xdf, 0xf9, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xee, 0xfb, 0x09, 0xff, 0x32, 0x07, 0x00,
	0x00,
}