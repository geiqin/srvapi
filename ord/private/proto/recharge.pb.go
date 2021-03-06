// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recharge.proto

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

type Recharge struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RechargeSn           string   `protobuf:"bytes,2,opt,name=recharge_sn,json=rechargeSn,proto3" json:"recharge_sn,omitempty"`
	BranchId             int64    `protobuf:"varint,3,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Currency             string   `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float32  `protobuf:"fixed32,7,opt,name=amount,proto3" json:"amount,omitempty"`
	BackType             string   `protobuf:"bytes,8,opt,name=back_type,json=backType,proto3" json:"back_type,omitempty"`
	BackValue            string   `protobuf:"bytes,9,opt,name=back_value,json=backValue,proto3" json:"back_value,omitempty"`
	CustomerId           int64    `protobuf:"varint,10,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Meno                 string   `protobuf:"bytes,11,opt,name=meno,proto3" json:"meno,omitempty"`
	Metadata             string   `protobuf:"bytes,12,opt,name=metadata,proto3" json:"metadata,omitempty"`
	PayStatus            string   `protobuf:"bytes,13,opt,name=pay_status,json=payStatus,proto3" json:"pay_status,omitempty"`
	Status               string   `protobuf:"bytes,14,opt,name=status,proto3" json:"status,omitempty"`
	PaidAt               string   `protobuf:"bytes,15,opt,name=paid_at,json=paidAt,proto3" json:"paid_at,omitempty"`
	CreatedAt            string   `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Recharge) Reset()         { *m = Recharge{} }
func (m *Recharge) String() string { return proto.CompactTextString(m) }
func (*Recharge) ProtoMessage()    {}
func (*Recharge) Descriptor() ([]byte, []int) {
	return fileDescriptor_97761fa03f464eb0, []int{0}
}

func (m *Recharge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recharge.Unmarshal(m, b)
}
func (m *Recharge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recharge.Marshal(b, m, deterministic)
}
func (m *Recharge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recharge.Merge(m, src)
}
func (m *Recharge) XXX_Size() int {
	return xxx_messageInfo_Recharge.Size(m)
}
func (m *Recharge) XXX_DiscardUnknown() {
	xxx_messageInfo_Recharge.DiscardUnknown(m)
}

var xxx_messageInfo_Recharge proto.InternalMessageInfo

func (m *Recharge) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Recharge) GetRechargeSn() string {
	if m != nil {
		return m.RechargeSn
	}
	return ""
}

func (m *Recharge) GetBranchId() int64 {
	if m != nil {
		return m.BranchId
	}
	return 0
}

func (m *Recharge) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Recharge) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Recharge) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Recharge) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Recharge) GetBackType() string {
	if m != nil {
		return m.BackType
	}
	return ""
}

func (m *Recharge) GetBackValue() string {
	if m != nil {
		return m.BackValue
	}
	return ""
}

func (m *Recharge) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *Recharge) GetMeno() string {
	if m != nil {
		return m.Meno
	}
	return ""
}

func (m *Recharge) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Recharge) GetPayStatus() string {
	if m != nil {
		return m.PayStatus
	}
	return ""
}

func (m *Recharge) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Recharge) GetPaidAt() string {
	if m != nil {
		return m.PaidAt
	}
	return ""
}

func (m *Recharge) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Recharge) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type RechargeResponse struct {
	Entity               *Recharge   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Recharge `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RechargeResponse) Reset()         { *m = RechargeResponse{} }
func (m *RechargeResponse) String() string { return proto.CompactTextString(m) }
func (*RechargeResponse) ProtoMessage()    {}
func (*RechargeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97761fa03f464eb0, []int{1}
}

func (m *RechargeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RechargeResponse.Unmarshal(m, b)
}
func (m *RechargeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RechargeResponse.Marshal(b, m, deterministic)
}
func (m *RechargeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RechargeResponse.Merge(m, src)
}
func (m *RechargeResponse) XXX_Size() int {
	return xxx_messageInfo_RechargeResponse.Size(m)
}
func (m *RechargeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RechargeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RechargeResponse proto.InternalMessageInfo

func (m *RechargeResponse) GetEntity() *Recharge {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *RechargeResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *RechargeResponse) GetItems() []*Recharge {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *RechargeResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *RechargeResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Recharge)(nil), "geiqin.srv.ord.private.Recharge")
	proto.RegisterType((*RechargeResponse)(nil), "geiqin.srv.ord.private.RechargeResponse")
}

func init() {
	proto.RegisterFile("recharge.proto", fileDescriptor_97761fa03f464eb0)
}

var fileDescriptor_97761fa03f464eb0 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x26, 0x76, 0xe2, 0x24, 0x93, 0x92, 0x96, 0x3d, 0x94, 0x55, 0xa0, 0x22, 0xe4, 0x94, 0x93,
	0x85, 0x52, 0x09, 0x71, 0x2d, 0x3f, 0x42, 0xbd, 0x21, 0x07, 0xb5, 0xc7, 0x68, 0xbb, 0x9e, 0x26,
	0x2b, 0xea, 0x5d, 0xb3, 0x5e, 0x47, 0xf2, 0xdb, 0xf0, 0x52, 0x3c, 0x07, 0xaf, 0x80, 0x76, 0xd6,
	0xee, 0x89, 0x50, 0x90, 0x7a, 0xf3, 0x7c, 0x3f, 0xf3, 0xcd, 0xee, 0x8e, 0x0c, 0x53, 0x8b, 0x72,
	0x27, 0xec, 0x16, 0xd3, 0xd2, 0x1a, 0x67, 0xd8, 0xe9, 0x16, 0xd5, 0x77, 0xa5, 0xd3, 0xca, 0xee,
	0x53, 0x63, 0xf3, 0xb4, 0xb4, 0x6a, 0x2f, 0x1c, 0xce, 0x8e, 0xa4, 0x29, 0x0a, 0xa3, 0x83, 0x6a,
	0xf1, 0x33, 0x86, 0x51, 0xd6, 0x1a, 0xd9, 0x14, 0x22, 0x95, 0xf3, 0xde, 0xbc, 0xb7, 0x8c, 0xb3,
	0x48, 0xe5, 0xec, 0x15, 0x4c, 0xba, 0xa6, 0x9b, 0x4a, 0xf3, 0x68, 0xde, 0x5b, 0x8e, 0x33, 0xe8,
	0xa0, 0xb5, 0x66, 0x2f, 0x60, 0x7c, 0x63, 0x85, 0x96, 0xbb, 0x8d, 0xca, 0x79, 0x4c, 0xbe, 0x51,
	0x00, 0x2e, 0x73, 0xc6, 0xa0, 0xef, 0x9a, 0x12, 0x79, 0x9f, 0x6c, 0xf4, 0xcd, 0x38, 0x0c, 0xa5,
	0xd1, 0x0e, 0xb5, 0xe3, 0x03, 0x82, 0xbb, 0x92, 0xcd, 0x60, 0x24, 0x6b, 0x6b, 0x51, 0xcb, 0x86,
	0x27, 0x44, 0xdd, 0xd7, 0xec, 0x14, 0x12, 0x51, 0x98, 0x5a, 0x3b, 0x3e, 0x9c, 0xf7, 0x96, 0x51,
	0xd6, 0x56, 0x14, 0x2f, 0xe4, 0xb7, 0x0d, 0xc5, 0x8c, 0x82, 0xc9, 0x03, 0x5f, 0x7d, 0xd4, 0x19,
	0x00, 0x91, 0x7b, 0x71, 0x57, 0x23, 0x1f, 0x13, 0x4b, 0xf2, 0x2b, 0x0f, 0xf8, 0xb3, 0xc9, 0xba,
	0x72, 0xa6, 0x40, 0xeb, 0x87, 0x07, 0x1a, 0x1e, 0x3a, 0x28, 0x8c, 0x5f, 0xa0, 0x36, 0x7c, 0x12,
	0xc6, 0xf7, 0xdf, 0x7e, 0xc8, 0x02, 0x9d, 0xc8, 0x85, 0x13, 0xfc, 0x28, 0xe4, 0x75, 0xb5, 0xcf,
	0x2b, 0x45, 0xb3, 0xa9, 0x9c, 0x70, 0x75, 0xc5, 0x9f, 0x86, 0xbc, 0x52, 0x34, 0x6b, 0x02, 0xfc,
	0x19, 0x5a, 0x6a, 0x4a, 0x54, 0x5b, 0xb1, 0xe7, 0x30, 0x2c, 0x85, 0xca, 0x37, 0xc2, 0xf1, 0xe3,
	0x40, 0xf8, 0xf2, 0xc2, 0xf9, 0x7e, 0xd2, 0xa2, 0x70, 0x48, 0xdc, 0x49, 0xe8, 0xd7, 0x22, 0x81,
	0xae, 0xcb, 0xbc, 0xa3, 0x9f, 0x05, 0xba, 0x45, 0x2e, 0xdc, 0xe2, 0x47, 0x04, 0x27, 0xdd, 0xbb,
	0x66, 0x58, 0x95, 0x46, 0x57, 0xc8, 0xde, 0x41, 0x82, 0xda, 0x29, 0xd7, 0xd0, 0x1b, 0x4f, 0x56,
	0xf3, 0xf4, 0xcf, 0x3b, 0x92, 0xde, 0x3b, 0x5b, 0x3d, 0x3b, 0x87, 0x41, 0x29, 0xb6, 0x68, 0x69,
	0x07, 0x26, 0xab, 0xb3, 0x43, 0xc6, 0x2f, 0x5e, 0x94, 0x05, 0x2d, 0x7b, 0x0b, 0x03, 0xe5, 0xb0,
	0xa8, 0x78, 0x3c, 0x8f, 0xff, 0x29, 0x2d, 0xc8, 0x7d, 0x18, 0x5a, 0x6b, 0x2c, 0x6d, 0xce, 0x5f,
	0xc2, 0x3e, 0x79, 0x51, 0x16, 0xb4, 0xec, 0x0d, 0xf4, 0x95, 0xbe, 0x35, 0xb4, 0x56, 0x93, 0xd5,
	0xcb, 0x43, 0x9e, 0x4b, 0x7d, 0x6b, 0x32, 0x52, 0xae, 0x7e, 0x45, 0x70, 0xdc, 0x45, 0xaf, 0xd1,
	0xee, 0x95, 0x44, 0x76, 0x05, 0xc9, 0x07, 0xba, 0x62, 0xf6, 0xe0, 0xb4, 0xb3, 0xe5, 0x83, 0xe7,
	0x69, 0xef, 0x7d, 0xf1, 0xc4, 0xf7, 0xfd, 0x88, 0x77, 0xf8, 0xe8, 0x7d, 0xd7, 0x10, 0x7f, 0x46,
	0xf7, 0xc8, 0x4d, 0xaf, 0x21, 0x59, 0xa3, 0xb0, 0x72, 0xc7, 0x5e, 0x1f, 0x72, 0xbd, 0x17, 0x15,
	0x5e, 0xef, 0xd0, 0xfe, 0x57, 0xe3, 0x9b, 0x84, 0xfe, 0x39, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x20, 0xc9, 0x45, 0xba, 0xab, 0x04, 0x00, 0x00,
}
