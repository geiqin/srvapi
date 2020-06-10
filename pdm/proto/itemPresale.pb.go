// Code generated by protoc-gen-go. DO NOT EDIT.
// source: itemPresale.proto

package geiqin_srv_pdm

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

type ItemPresale struct {
	Id                      int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemId                  int64    `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	PresaleType             int32    `protobuf:"varint,3,opt,name=presale_type,json=presaleType,proto3" json:"presale_type,omitempty"`
	PayedDate               string   `protobuf:"bytes,4,opt,name=payed_date,json=payedDate,proto3" json:"payed_date,omitempty"`
	DepositPaymentStartDate string   `protobuf:"bytes,5,opt,name=deposit_payment_start_date,json=depositPaymentStartDate,proto3" json:"deposit_payment_start_date,omitempty"`
	DepositPaymentEndDate   string   `protobuf:"bytes,6,opt,name=deposit_payment_end_date,json=depositPaymentEndDate,proto3" json:"deposit_payment_end_date,omitempty"`
	PaymentStartDate        string   `protobuf:"bytes,7,opt,name=payment_start_date,json=paymentStartDate,proto3" json:"payment_start_date,omitempty"`
	PaymentEndDate          string   `protobuf:"bytes,8,opt,name=payment_end_date,json=paymentEndDate,proto3" json:"payment_end_date,omitempty"`
	DepositType             int32    `protobuf:"varint,9,opt,name=deposit_type,json=depositType,proto3" json:"deposit_type,omitempty"`
	DepositRatio            int32    `protobuf:"varint,10,opt,name=deposit_ratio,json=depositRatio,proto3" json:"deposit_ratio,omitempty"`
	DepositAmount           float32  `protobuf:"fixed32,11,opt,name=deposit_amount,json=depositAmount,proto3" json:"deposit_amount,omitempty"`
	RestPayment             float32  `protobuf:"fixed32,12,opt,name=rest_payment,json=restPayment,proto3" json:"rest_payment,omitempty"`
	DeliveryType            int32    `protobuf:"varint,13,opt,name=delivery_type,json=deliveryType,proto3" json:"delivery_type,omitempty"`
	DeliveryDate            string   `protobuf:"bytes,14,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	DeliveryDays            int32    `protobuf:"varint,15,opt,name=delivery_days,json=deliveryDays,proto3" json:"delivery_days,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ItemPresale) Reset()         { *m = ItemPresale{} }
func (m *ItemPresale) String() string { return proto.CompactTextString(m) }
func (*ItemPresale) ProtoMessage()    {}
func (*ItemPresale) Descriptor() ([]byte, []int) {
	return fileDescriptor_4582b57b3bd016d0, []int{0}
}

func (m *ItemPresale) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemPresale.Unmarshal(m, b)
}
func (m *ItemPresale) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemPresale.Marshal(b, m, deterministic)
}
func (m *ItemPresale) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemPresale.Merge(m, src)
}
func (m *ItemPresale) XXX_Size() int {
	return xxx_messageInfo_ItemPresale.Size(m)
}
func (m *ItemPresale) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemPresale.DiscardUnknown(m)
}

var xxx_messageInfo_ItemPresale proto.InternalMessageInfo

func (m *ItemPresale) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemPresale) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ItemPresale) GetPresaleType() int32 {
	if m != nil {
		return m.PresaleType
	}
	return 0
}

func (m *ItemPresale) GetPayedDate() string {
	if m != nil {
		return m.PayedDate
	}
	return ""
}

func (m *ItemPresale) GetDepositPaymentStartDate() string {
	if m != nil {
		return m.DepositPaymentStartDate
	}
	return ""
}

func (m *ItemPresale) GetDepositPaymentEndDate() string {
	if m != nil {
		return m.DepositPaymentEndDate
	}
	return ""
}

func (m *ItemPresale) GetPaymentStartDate() string {
	if m != nil {
		return m.PaymentStartDate
	}
	return ""
}

func (m *ItemPresale) GetPaymentEndDate() string {
	if m != nil {
		return m.PaymentEndDate
	}
	return ""
}

func (m *ItemPresale) GetDepositType() int32 {
	if m != nil {
		return m.DepositType
	}
	return 0
}

func (m *ItemPresale) GetDepositRatio() int32 {
	if m != nil {
		return m.DepositRatio
	}
	return 0
}

func (m *ItemPresale) GetDepositAmount() float32 {
	if m != nil {
		return m.DepositAmount
	}
	return 0
}

func (m *ItemPresale) GetRestPayment() float32 {
	if m != nil {
		return m.RestPayment
	}
	return 0
}

func (m *ItemPresale) GetDeliveryType() int32 {
	if m != nil {
		return m.DeliveryType
	}
	return 0
}

func (m *ItemPresale) GetDeliveryDate() string {
	if m != nil {
		return m.DeliveryDate
	}
	return ""
}

func (m *ItemPresale) GetDeliveryDays() int32 {
	if m != nil {
		return m.DeliveryDays
	}
	return 0
}

type ItemPresaleResponse struct {
	Entity               *ItemPresale `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Error                *Error       `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info        `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ItemPresaleResponse) Reset()         { *m = ItemPresaleResponse{} }
func (m *ItemPresaleResponse) String() string { return proto.CompactTextString(m) }
func (*ItemPresaleResponse) ProtoMessage()    {}
func (*ItemPresaleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4582b57b3bd016d0, []int{1}
}

func (m *ItemPresaleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemPresaleResponse.Unmarshal(m, b)
}
func (m *ItemPresaleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemPresaleResponse.Marshal(b, m, deterministic)
}
func (m *ItemPresaleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemPresaleResponse.Merge(m, src)
}
func (m *ItemPresaleResponse) XXX_Size() int {
	return xxx_messageInfo_ItemPresaleResponse.Size(m)
}
func (m *ItemPresaleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemPresaleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ItemPresaleResponse proto.InternalMessageInfo

func (m *ItemPresaleResponse) GetEntity() *ItemPresale {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ItemPresaleResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ItemPresaleResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ItemPresaleResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemPresale)(nil), "geiqin.srv.pdm.ItemPresale")
	proto.RegisterType((*ItemPresaleResponse)(nil), "geiqin.srv.pdm.ItemPresaleResponse")
}

func init() {
	proto.RegisterFile("itemPresale.proto", fileDescriptor_4582b57b3bd016d0)
}

var fileDescriptor_4582b57b3bd016d0 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x41, 0x6b, 0xdb, 0x40,
	0x10, 0x85, 0x2b, 0xc7, 0x56, 0xea, 0x91, 0xa3, 0xa6, 0xdb, 0x86, 0x08, 0x97, 0x82, 0xe3, 0x50,
	0x10, 0xb4, 0xe8, 0xe0, 0x1c, 0x7a, 0xe8, 0xa9, 0x34, 0x39, 0x04, 0x7a, 0x30, 0x4a, 0xef, 0x62,
	0x13, 0x4d, 0xc2, 0x42, 0xb4, 0xbb, 0xdd, 0xdd, 0x1a, 0xf4, 0x0b, 0x7b, 0xeb, 0x6f, 0x2a, 0x3b,
	0x2b, 0x25, 0x96, 0x13, 0x72, 0xd4, 0xdb, 0xef, 0xcd, 0xcc, 0x1b, 0x46, 0xf0, 0x56, 0x38, 0x6c,
	0xd6, 0x06, 0x2d, 0xbf, 0xc7, 0x42, 0x1b, 0xe5, 0x14, 0x4b, 0xef, 0x50, 0xfc, 0x16, 0xb2, 0xb0,
	0x66, 0x53, 0xe8, 0xba, 0x99, 0xcf, 0x6e, 0x54, 0xd3, 0x28, 0x19, 0x5e, 0x97, 0x7f, 0xc7, 0x90,
	0x5c, 0x3e, 0x7a, 0x58, 0x0a, 0x23, 0x51, 0x67, 0xd1, 0x22, 0xca, 0xf7, 0xca, 0x91, 0xa8, 0xd9,
	0x31, 0xec, 0xfb, 0x92, 0x95, 0xa8, 0xb3, 0x11, 0x89, 0xb1, 0xff, 0xbc, 0xac, 0xd9, 0x09, 0xcc,
	0x74, 0xf0, 0x54, 0xae, 0xd5, 0x98, 0xed, 0x2d, 0xa2, 0x7c, 0x52, 0x26, 0x9d, 0xf6, 0xab, 0xd5,
	0xc8, 0x3e, 0x02, 0x68, 0xde, 0x62, 0x5d, 0xd5, 0xdc, 0x61, 0x36, 0x5e, 0x44, 0xf9, 0xb4, 0x9c,
	0x92, 0x72, 0xce, 0x1d, 0xb2, 0x6f, 0x30, 0xaf, 0x51, 0x2b, 0x2b, 0x5c, 0xa5, 0x79, 0xdb, 0xa0,
	0x74, 0x95, 0x75, 0xdc, 0xb8, 0x80, 0x4f, 0x08, 0x3f, 0xee, 0x88, 0x75, 0x00, 0xae, 0xfc, 0x3b,
	0x99, 0xbf, 0x42, 0xb6, 0x6b, 0x46, 0xd9, 0x75, 0x8a, 0xc9, 0x7a, 0x34, 0xb4, 0x5e, 0xc8, 0xd0,
	0xf5, 0x0b, 0xb0, 0x67, 0xba, 0xed, 0x93, 0xe5, 0x50, 0xef, 0xb6, 0xc9, 0xe1, 0xf0, 0x49, 0xf9,
	0xd7, 0xc4, 0xa6, 0x7a, 0x58, 0xf7, 0x04, 0x66, 0xfd, 0x40, 0xb4, 0x8f, 0x69, 0xd8, 0x47, 0xa7,
	0xd1, 0x3e, 0x4e, 0xe1, 0xa0, 0x47, 0x0c, 0x77, 0x42, 0x65, 0x40, 0x4c, 0xef, 0x2b, 0xbd, 0xc6,
	0x3e, 0x41, 0xda, 0x43, 0xbc, 0x51, 0x7f, 0xa4, 0xcb, 0x92, 0x45, 0x94, 0x8f, 0xca, 0xde, 0xfa,
	0x9d, 0x44, 0xdf, 0xce, 0xa0, 0x7d, 0x08, 0x9f, 0xcd, 0x08, 0x4a, 0xbc, 0xd6, 0x05, 0x0e, 0xed,
	0xee, 0xc5, 0x06, 0x4d, 0x1b, 0x46, 0x3a, 0xe8, 0xdb, 0x05, 0xf1, 0x71, 0xa6, 0x0e, 0xa2, 0x74,
	0x29, 0xa5, 0x7b, 0x80, 0x28, 0xdb, 0x10, 0x6a, 0x6d, 0xf6, 0x66, 0x58, 0xe9, 0x9c, 0xb7, 0x76,
	0xf9, 0x2f, 0x82, 0x77, 0x5b, 0x97, 0x54, 0xa2, 0xd5, 0x4a, 0x5a, 0x64, 0x67, 0x10, 0xa3, 0x74,
	0xc2, 0xb5, 0x74, 0x55, 0xc9, 0xea, 0x43, 0x31, 0x3c, 0xc8, 0x62, 0xdb, 0xd4, 0xa1, 0xec, 0x33,
	0x4c, 0x34, 0xbf, 0x43, 0x43, 0x47, 0x97, 0xac, 0x8e, 0x76, 0x3d, 0x6b, 0xff, 0x58, 0x06, 0xc6,
	0xc3, 0x68, 0x8c, 0x32, 0x74, 0x83, 0xcf, 0xc0, 0x17, 0xfe, 0xb1, 0x0c, 0x0c, 0xcb, 0x61, 0x2c,
	0xe4, 0xad, 0xa2, 0x73, 0x4c, 0x56, 0xef, 0x9f, 0x0c, 0x23, 0x6f, 0x55, 0x49, 0xc4, 0xea, 0x1a,
	0xd8, 0xd6, 0x68, 0x57, 0x68, 0x36, 0xe2, 0x06, 0xd9, 0x4f, 0x88, 0x7f, 0x18, 0xf4, 0x5b, 0x79,
	0x29, 0xc8, 0xfc, 0xf4, 0xa5, 0x94, 0xdd, 0x6a, 0x96, 0xaf, 0xae, 0x63, 0xfa, 0x0b, 0xcf, 0xfe,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x72, 0x54, 0xd6, 0xd4, 0xb8, 0x03, 0x00, 0x00,
}
