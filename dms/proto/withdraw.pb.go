// Code generated by protoc-gen-go. DO NOT EDIT.
// source: withdraw.proto

package geiqin_srv_dms

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

type Withdraw struct {
	Id                   int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WithdrawSn           string       `protobuf:"bytes,2,opt,name=withdraw_sn,json=withdrawSn,proto3" json:"withdraw_sn,omitempty"`
	DistributorId        int64        `protobuf:"varint,3,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	Money                float32      `protobuf:"fixed32,4,opt,name=money,proto3" json:"money,omitempty"`
	TransferId           int64        `protobuf:"varint,5,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"`
	TransactionIdId      string       `protobuf:"bytes,6,opt,name=transactionId_id,json=transactionIdId,proto3" json:"transactionId_id,omitempty"`
	Method               string       `protobuf:"bytes,7,opt,name=method,proto3" json:"method,omitempty"`
	CashAccountId        int32        `protobuf:"varint,8,opt,name=cash_account_id,json=cashAccountId,proto3" json:"cash_account_id,omitempty"`
	PayeeName            string       `protobuf:"bytes,9,opt,name=payee_name,json=payeeName,proto3" json:"payee_name,omitempty"`
	PayeeAccount         string       `protobuf:"bytes,10,opt,name=payee_account,json=payeeAccount,proto3" json:"payee_account,omitempty"`
	PayeeBank            string       `protobuf:"bytes,11,opt,name=payee_bank,json=payeeBank,proto3" json:"payee_bank,omitempty"`
	PlatformAccount      string       `protobuf:"bytes,12,opt,name=platform_account,json=platformAccount,proto3" json:"platform_account,omitempty"`
	Memo                 string       `protobuf:"bytes,13,opt,name=memo,proto3" json:"memo,omitempty"`
	Status               string       `protobuf:"bytes,14,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string       `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string       `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Distributor          *Distributor `protobuf:"bytes,17,opt,name=distributor,proto3" json:"distributor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Withdraw) Reset()         { *m = Withdraw{} }
func (m *Withdraw) String() string { return proto.CompactTextString(m) }
func (*Withdraw) ProtoMessage()    {}
func (*Withdraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0dd7acb611886fa, []int{0}
}

func (m *Withdraw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Withdraw.Unmarshal(m, b)
}
func (m *Withdraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Withdraw.Marshal(b, m, deterministic)
}
func (m *Withdraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Withdraw.Merge(m, src)
}
func (m *Withdraw) XXX_Size() int {
	return xxx_messageInfo_Withdraw.Size(m)
}
func (m *Withdraw) XXX_DiscardUnknown() {
	xxx_messageInfo_Withdraw.DiscardUnknown(m)
}

var xxx_messageInfo_Withdraw proto.InternalMessageInfo

func (m *Withdraw) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Withdraw) GetWithdrawSn() string {
	if m != nil {
		return m.WithdrawSn
	}
	return ""
}

func (m *Withdraw) GetDistributorId() int64 {
	if m != nil {
		return m.DistributorId
	}
	return 0
}

func (m *Withdraw) GetMoney() float32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *Withdraw) GetTransferId() int64 {
	if m != nil {
		return m.TransferId
	}
	return 0
}

func (m *Withdraw) GetTransactionIdId() string {
	if m != nil {
		return m.TransactionIdId
	}
	return ""
}

func (m *Withdraw) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Withdraw) GetCashAccountId() int32 {
	if m != nil {
		return m.CashAccountId
	}
	return 0
}

func (m *Withdraw) GetPayeeName() string {
	if m != nil {
		return m.PayeeName
	}
	return ""
}

func (m *Withdraw) GetPayeeAccount() string {
	if m != nil {
		return m.PayeeAccount
	}
	return ""
}

func (m *Withdraw) GetPayeeBank() string {
	if m != nil {
		return m.PayeeBank
	}
	return ""
}

func (m *Withdraw) GetPlatformAccount() string {
	if m != nil {
		return m.PlatformAccount
	}
	return ""
}

func (m *Withdraw) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Withdraw) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Withdraw) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Withdraw) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Withdraw) GetDistributor() *Distributor {
	if m != nil {
		return m.Distributor
	}
	return nil
}

type WithdrawResponse struct {
	Entity               *Withdraw   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Withdraw `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WithdrawResponse) Reset()         { *m = WithdrawResponse{} }
func (m *WithdrawResponse) String() string { return proto.CompactTextString(m) }
func (*WithdrawResponse) ProtoMessage()    {}
func (*WithdrawResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0dd7acb611886fa, []int{1}
}

func (m *WithdrawResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawResponse.Unmarshal(m, b)
}
func (m *WithdrawResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawResponse.Marshal(b, m, deterministic)
}
func (m *WithdrawResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawResponse.Merge(m, src)
}
func (m *WithdrawResponse) XXX_Size() int {
	return xxx_messageInfo_WithdrawResponse.Size(m)
}
func (m *WithdrawResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawResponse proto.InternalMessageInfo

func (m *WithdrawResponse) GetEntity() *Withdraw {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *WithdrawResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *WithdrawResponse) GetItems() []*Withdraw {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *WithdrawResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *WithdrawResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Withdraw)(nil), "geiqin.srv.dms.Withdraw")
	proto.RegisterType((*WithdrawResponse)(nil), "geiqin.srv.dms.WithdrawResponse")
}

func init() {
	proto.RegisterFile("withdraw.proto", fileDescriptor_b0dd7acb611886fa)
}

var fileDescriptor_b0dd7acb611886fa = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0xd1, 0x6e, 0xd3, 0x3e,
	0x14, 0xc6, 0xff, 0x69, 0x9b, 0xfe, 0xb7, 0x93, 0xb5, 0xdd, 0xac, 0x81, 0xcc, 0x10, 0x22, 0x1a,
	0x02, 0x15, 0x21, 0x55, 0x28, 0x5c, 0x73, 0xd1, 0x8d, 0x31, 0xe5, 0x02, 0x84, 0xb2, 0x8b, 0x5d,
	0x56, 0x5e, 0x7c, 0xba, 0x58, 0x5b, 0xec, 0x60, 0xbb, 0x9b, 0xfa, 0x52, 0xbc, 0x00, 0x2f, 0xc4,
	0x23, 0x70, 0x89, 0x62, 0x27, 0x5b, 0x57, 0x31, 0x84, 0xb4, 0x3b, 0xee, 0xea, 0xef, 0x7c, 0xe7,
	0xe7, 0x73, 0xdc, 0x4f, 0x81, 0xe1, 0xb5, 0xb0, 0x05, 0xd7, 0xec, 0x7a, 0x52, 0x69, 0x65, 0x15,
	0x19, 0x9e, 0xa3, 0xf8, 0x2a, 0xe4, 0xc4, 0xe8, 0xab, 0x09, 0x2f, 0xcd, 0xde, 0x56, 0xae, 0xca,
	0x52, 0x49, 0x5f, 0xdd, 0xdb, 0xe1, 0xc2, 0x58, 0x2d, 0xce, 0x16, 0x56, 0x69, 0x2f, 0xed, 0x7f,
	0xef, 0xc1, 0xc6, 0x69, 0xc3, 0x20, 0x43, 0xe8, 0x08, 0x4e, 0x83, 0x38, 0x18, 0x77, 0xb3, 0x8e,
	0xe0, 0xe4, 0x39, 0x44, 0x2d, 0x7f, 0x66, 0x24, 0xed, 0xc4, 0xc1, 0x78, 0x33, 0x83, 0x56, 0x3a,
	0x91, 0xe4, 0x25, 0x0c, 0x57, 0x90, 0x33, 0xc1, 0x69, 0xd7, 0x35, 0x0f, 0x56, 0xd4, 0x94, 0x93,
	0x5d, 0x08, 0x4b, 0x25, 0x71, 0x49, 0x7b, 0x71, 0x30, 0xee, 0x64, 0xfe, 0x50, 0xd3, 0xad, 0x66,
	0xd2, 0xcc, 0xd1, 0x75, 0x86, 0xae, 0x13, 0x5a, 0x29, 0xe5, 0xe4, 0x35, 0x6c, 0xbb, 0x13, 0xcb,
	0xad, 0x50, 0x32, 0xe5, 0xb5, 0xab, 0xef, 0x66, 0x18, 0xdd, 0xd1, 0x53, 0x4e, 0x1e, 0x43, 0xbf,
	0x44, 0x5b, 0x28, 0x4e, 0xff, 0x77, 0x86, 0xe6, 0x44, 0x5e, 0xc1, 0x28, 0x67, 0xa6, 0x98, 0xb1,
	0x3c, 0x57, 0x0b, 0x69, 0x6b, 0xc2, 0x46, 0x1c, 0x8c, 0xc3, 0x6c, 0x50, 0xcb, 0x53, 0xaf, 0xa6,
	0x9c, 0x3c, 0x03, 0xa8, 0xd8, 0x12, 0x71, 0x26, 0x59, 0x89, 0x74, 0xd3, 0x31, 0x36, 0x9d, 0xf2,
	0x99, 0x95, 0x48, 0x5e, 0xc0, 0xc0, 0x97, 0x1b, 0x0e, 0x05, 0xe7, 0xd8, 0x72, 0x62, 0x43, 0xb9,
	0x65, 0x9c, 0x31, 0x79, 0x41, 0xa3, 0x15, 0xc6, 0x01, 0x93, 0x17, 0xf5, 0x36, 0xd5, 0x25, 0xb3,
	0x73, 0xa5, 0xcb, 0x1b, 0xcc, 0x96, 0xdf, 0xa6, 0xd5, 0x5b, 0x12, 0x81, 0x5e, 0x89, 0xa5, 0xa2,
	0x03, 0x57, 0x76, 0xbf, 0xeb, 0x0d, 0x8d, 0x65, 0x76, 0x61, 0xe8, 0xd0, 0x6f, 0xe8, 0x4f, 0xf5,
	0xad, 0xb9, 0x46, 0x66, 0x91, 0xcf, 0x98, 0xa5, 0x23, 0x7f, 0x6b, 0xa3, 0x4c, 0xdd, 0x50, 0x8b,
	0x8a, 0xb7, 0xe5, 0x6d, 0x5f, 0x6e, 0x94, 0xa9, 0x25, 0xef, 0x21, 0x5a, 0xf9, 0xab, 0xe8, 0x4e,
	0x1c, 0x8c, 0xa3, 0xe4, 0xe9, 0xe4, 0x6e, 0x8a, 0x26, 0x1f, 0x6e, 0x2d, 0xd9, 0xaa, 0x7f, 0xff,
	0x67, 0x00, 0xdb, 0x6d, 0x7a, 0x32, 0x34, 0x95, 0x92, 0x06, 0xc9, 0x5b, 0xe8, 0xa3, 0xb4, 0xc2,
	0x2e, 0x5d, 0x92, 0xa2, 0x84, 0xae, 0xe3, 0x6e, 0x3a, 0x1a, 0x1f, 0x79, 0x03, 0x61, 0xc5, 0xce,
	0x51, 0xbb, 0x84, 0x45, 0xc9, 0xa3, 0xf5, 0x86, 0x2f, 0x75, 0x31, 0xf3, 0x1e, 0x32, 0x81, 0x50,
	0x58, 0x2c, 0x0d, 0xed, 0xc6, 0xdd, 0x3f, 0xd2, 0xbd, 0xad, 0x86, 0xa3, 0xd6, 0x4a, 0xbb, 0xf0,
	0xfd, 0x06, 0x7e, 0x54, 0x17, 0x33, 0xef, 0x21, 0x63, 0xe8, 0x09, 0x39, 0x57, 0x2e, 0x8c, 0x51,
	0xb2, 0xbb, 0xee, 0x4d, 0xe5, 0x5c, 0x65, 0xce, 0x91, 0xfc, 0x08, 0x60, 0xe7, 0xd3, 0xb2, 0xbd,
	0xec, 0x04, 0xf5, 0x95, 0xc8, 0x91, 0x1c, 0x41, 0x38, 0xad, 0xaa, 0xcb, 0x25, 0xb9, 0x77, 0xac,
	0xbd, 0xf8, 0xde, 0x81, 0x9b, 0x07, 0xdc, 0xff, 0x8f, 0x1c, 0x42, 0xf7, 0x18, 0xed, 0x03, 0x21,
	0xc7, 0xd0, 0x3f, 0x41, 0xa6, 0xf3, 0x82, 0x3c, 0x59, 0x77, 0x1f, 0x30, 0x83, 0xa7, 0x05, 0x6a,
	0xfc, 0x1b, 0x50, 0xf2, 0xad, 0x03, 0xa3, 0xf5, 0x45, 0x3f, 0x42, 0xff, 0xd0, 0x85, 0xec, 0x81,
	0x43, 0x1e, 0x41, 0x78, 0x58, 0x60, 0x7e, 0xf1, 0x2f, 0x3d, 0xd8, 0x59, 0xdf, 0x7d, 0x5b, 0xdf,
	0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x31, 0x02, 0x9a, 0x9e, 0x05, 0x00, 0x00,
}
