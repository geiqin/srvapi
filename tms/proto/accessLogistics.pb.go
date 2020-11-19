// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accessLogistics.proto

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

type ExpressBill struct {
	LogisticsNo          string   `protobuf:"bytes,1,opt,name=logistics_no,json=logisticsNo,proto3" json:"logistics_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpressBill) Reset()         { *m = ExpressBill{} }
func (m *ExpressBill) String() string { return proto.CompactTextString(m) }
func (*ExpressBill) ProtoMessage()    {}
func (*ExpressBill) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b48d4c79702ca9c, []int{0}
}

func (m *ExpressBill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpressBill.Unmarshal(m, b)
}
func (m *ExpressBill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpressBill.Marshal(b, m, deterministic)
}
func (m *ExpressBill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpressBill.Merge(m, src)
}
func (m *ExpressBill) XXX_Size() int {
	return xxx_messageInfo_ExpressBill.Size(m)
}
func (m *ExpressBill) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpressBill.DiscardUnknown(m)
}

var xxx_messageInfo_ExpressBill proto.InternalMessageInfo

func (m *ExpressBill) GetLogisticsNo() string {
	if m != nil {
		return m.LogisticsNo
	}
	return ""
}

type ShipmentReq struct {
	ShipperId            int32                `protobuf:"varint,1,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	OrderId              int64                `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	DeliverySn           string               `protobuf:"bytes,3,opt,name=delivery_sn,json=deliverySn,proto3" json:"delivery_sn,omitempty"`
	Details              []*ShipmentDetailReq `protobuf:"bytes,4,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ShipmentReq) Reset()         { *m = ShipmentReq{} }
func (m *ShipmentReq) String() string { return proto.CompactTextString(m) }
func (*ShipmentReq) ProtoMessage()    {}
func (*ShipmentReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b48d4c79702ca9c, []int{1}
}

func (m *ShipmentReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipmentReq.Unmarshal(m, b)
}
func (m *ShipmentReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipmentReq.Marshal(b, m, deterministic)
}
func (m *ShipmentReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipmentReq.Merge(m, src)
}
func (m *ShipmentReq) XXX_Size() int {
	return xxx_messageInfo_ShipmentReq.Size(m)
}
func (m *ShipmentReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipmentReq.DiscardUnknown(m)
}

var xxx_messageInfo_ShipmentReq proto.InternalMessageInfo

func (m *ShipmentReq) GetShipperId() int32 {
	if m != nil {
		return m.ShipperId
	}
	return 0
}

func (m *ShipmentReq) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *ShipmentReq) GetDeliverySn() string {
	if m != nil {
		return m.DeliverySn
	}
	return ""
}

func (m *ShipmentReq) GetDetails() []*ShipmentDetailReq {
	if m != nil {
		return m.Details
	}
	return nil
}

type ShipmentDetailReq struct {
	OrderDetailId        int64    `protobuf:"varint,1,opt,name=order_detail_id,json=orderDetailId,proto3" json:"order_detail_id,omitempty"`
	ItemId               int64    `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64    `protobuf:"varint,3,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Num                  int32    `protobuf:"varint,4,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipmentDetailReq) Reset()         { *m = ShipmentDetailReq{} }
func (m *ShipmentDetailReq) String() string { return proto.CompactTextString(m) }
func (*ShipmentDetailReq) ProtoMessage()    {}
func (*ShipmentDetailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b48d4c79702ca9c, []int{2}
}

func (m *ShipmentDetailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipmentDetailReq.Unmarshal(m, b)
}
func (m *ShipmentDetailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipmentDetailReq.Marshal(b, m, deterministic)
}
func (m *ShipmentDetailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipmentDetailReq.Merge(m, src)
}
func (m *ShipmentDetailReq) XXX_Size() int {
	return xxx_messageInfo_ShipmentDetailReq.Size(m)
}
func (m *ShipmentDetailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipmentDetailReq.DiscardUnknown(m)
}

var xxx_messageInfo_ShipmentDetailReq proto.InternalMessageInfo

func (m *ShipmentDetailReq) GetOrderDetailId() int64 {
	if m != nil {
		return m.OrderDetailId
	}
	return 0
}

func (m *ShipmentDetailReq) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ShipmentDetailReq) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *ShipmentDetailReq) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type ExpressBillResponse struct {
	Entity               *ExpressBill   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager         `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*ExpressBill `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error         `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info          `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExpressBillResponse) Reset()         { *m = ExpressBillResponse{} }
func (m *ExpressBillResponse) String() string { return proto.CompactTextString(m) }
func (*ExpressBillResponse) ProtoMessage()    {}
func (*ExpressBillResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b48d4c79702ca9c, []int{3}
}

func (m *ExpressBillResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpressBillResponse.Unmarshal(m, b)
}
func (m *ExpressBillResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpressBillResponse.Marshal(b, m, deterministic)
}
func (m *ExpressBillResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpressBillResponse.Merge(m, src)
}
func (m *ExpressBillResponse) XXX_Size() int {
	return xxx_messageInfo_ExpressBillResponse.Size(m)
}
func (m *ExpressBillResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpressBillResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExpressBillResponse proto.InternalMessageInfo

func (m *ExpressBillResponse) GetEntity() *ExpressBill {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ExpressBillResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ExpressBillResponse) GetItems() []*ExpressBill {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ExpressBillResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ExpressBillResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*ExpressBill)(nil), "geiqin.srv.tms.ExpressBill")
	proto.RegisterType((*ShipmentReq)(nil), "geiqin.srv.tms.ShipmentReq")
	proto.RegisterType((*ShipmentDetailReq)(nil), "geiqin.srv.tms.ShipmentDetailReq")
	proto.RegisterType((*ExpressBillResponse)(nil), "geiqin.srv.tms.ExpressBillResponse")
}

func init() {
	proto.RegisterFile("accessLogistics.proto", fileDescriptor_3b48d4c79702ca9c)
}

var fileDescriptor_3b48d4c79702ca9c = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0x1d, 0x3b, 0x74, 0x5c, 0xfe, 0x0d, 0x04, 0x4c, 0x2b, 0x44, 0x6a, 0x24, 0x14,
	0x09, 0xc9, 0x82, 0xf4, 0xc8, 0x09, 0x44, 0x0f, 0x96, 0x2a, 0x84, 0x36, 0x0f, 0x10, 0xa5, 0xde,
	0x49, 0xba, 0xaa, 0xbd, 0xeb, 0xec, 0x6c, 0x22, 0x7a, 0xe0, 0xc0, 0xc3, 0xf0, 0x9e, 0xc8, 0xeb,
	0x18, 0xda, 0x40, 0x7b, 0xf3, 0x7e, 0xf3, 0xfb, 0x76, 0xbe, 0x19, 0x2f, 0x0c, 0xe7, 0x45, 0x41,
	0xcc, 0x67, 0x66, 0xa9, 0xd8, 0xa9, 0x82, 0xb3, 0xda, 0x1a, 0x67, 0xf0, 0xe1, 0x92, 0xd4, 0x4a,
	0xe9, 0x8c, 0xed, 0x26, 0x73, 0x15, 0x1f, 0x1e, 0x14, 0xa6, 0xaa, 0x8c, 0x6e, 0xab, 0xe9, 0x7b,
	0x88, 0x4f, 0xbf, 0xd7, 0x96, 0x98, 0x3f, 0xab, 0xb2, 0xc4, 0x63, 0x38, 0x28, 0x3b, 0xff, 0x4c,
	0x9b, 0xa4, 0x37, 0xea, 0x8d, 0xf7, 0x45, 0xfc, 0x47, 0xfb, 0x6a, 0xd2, 0x5f, 0x3d, 0x88, 0xa7,
	0x17, 0xaa, 0xae, 0x48, 0x3b, 0x41, 0x2b, 0x7c, 0x05, 0xc0, 0x17, 0xaa, 0xae, 0xc9, 0xce, 0x94,
	0xf4, 0x86, 0x50, 0xec, 0x6f, 0x95, 0x5c, 0xe2, 0x4b, 0xb8, 0x6f, 0xac, 0x6c, 0x8b, 0x7b, 0xa3,
	0xde, 0x38, 0x10, 0x03, 0x7f, 0xce, 0x25, 0xbe, 0x86, 0x58, 0x52, 0xa9, 0x36, 0x64, 0xaf, 0x66,
	0xac, 0x93, 0xc0, 0xf7, 0x82, 0x4e, 0x9a, 0x6a, 0xfc, 0x08, 0x03, 0x49, 0x6e, 0xae, 0x4a, 0x4e,
	0xfa, 0xa3, 0x60, 0x1c, 0x4f, 0x8e, 0xb3, 0x9b, 0xc3, 0x64, 0x5d, 0x90, 0x2f, 0x1e, 0x13, 0xb4,
	0x12, 0x9d, 0x23, 0xfd, 0x01, 0x4f, 0xfe, 0xa9, 0xe2, 0x5b, 0x78, 0xd4, 0xa6, 0x69, 0xa9, 0x2e,
	0x71, 0x20, 0x1e, 0x78, 0xb9, 0x05, 0x73, 0x89, 0x2f, 0x60, 0xa0, 0x1c, 0x55, 0x7f, 0x43, 0x47,
	0xcd, 0x31, 0x97, 0x38, 0x84, 0x88, 0x2f, 0xd7, 0x8d, 0x1e, 0x78, 0x3d, 0xe4, 0xcb, 0x75, 0x2e,
	0xf1, 0x31, 0x04, 0x7a, 0x5d, 0x25, 0x7d, 0x3f, 0x7d, 0xf3, 0x99, 0xfe, 0xdc, 0x83, 0xa7, 0xd7,
	0x36, 0x2b, 0x88, 0x6b, 0xa3, 0x99, 0xf0, 0x04, 0x22, 0xd2, 0x4e, 0xb9, 0x2b, 0xdf, 0x38, 0x9e,
	0x1c, 0xed, 0x8e, 0x74, 0xdd, 0xb4, 0x45, 0xf1, 0x1d, 0x84, 0xf5, 0x7c, 0x49, 0xd6, 0x87, 0x89,
	0x27, 0xc3, 0x5d, 0xcf, 0xb7, 0xa6, 0x28, 0x5a, 0x06, 0x3f, 0x40, 0xd8, 0x84, 0xe5, 0x24, 0xf0,
	0x3b, 0xbb, 0xb3, 0x41, 0x4b, 0x36, 0xf7, 0x93, 0xb5, 0xc6, 0xfa, 0x01, 0xfe, 0x73, 0xff, 0x69,
	0x53, 0x14, 0x2d, 0x83, 0x63, 0xe8, 0x2b, 0xbd, 0x30, 0x49, 0xe8, 0xd9, 0x67, 0xbb, 0x6c, 0xae,
	0x17, 0x46, 0x78, 0x62, 0xb2, 0x80, 0xe7, 0x9f, 0x6e, 0xbe, 0xc9, 0x29, 0xd9, 0x8d, 0x2a, 0x08,
	0xcf, 0x20, 0x9a, 0xae, 0xcf, 0x2b, 0xe5, 0xf0, 0xe8, 0xb6, 0x5f, 0x2a, 0x68, 0x75, 0xf8, 0xe6,
	0xae, 0xec, 0xdb, 0x8d, 0xa6, 0xf7, 0xce, 0x23, 0xff, 0x96, 0x4f, 0x7e, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x57, 0xa6, 0xb8, 0xeb, 0x02, 0x03, 0x00, 0x00,
}
