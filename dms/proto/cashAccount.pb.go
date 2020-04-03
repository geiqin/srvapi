// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cashAccount.proto

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

type CashAccount struct {
	Id                   int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DistributorId        int64        `protobuf:"varint,2,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	Type                 string       `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	RealName             string       `protobuf:"bytes,4,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	PayeeName            string       `protobuf:"bytes,5,opt,name=payee_name,json=payeeName,proto3" json:"payee_name,omitempty"`
	PayeeAccount         string       `protobuf:"bytes,6,opt,name=payee_account,json=payeeAccount,proto3" json:"payee_account,omitempty"`
	PayeeBank            string       `protobuf:"bytes,7,opt,name=payee_bank,json=payeeBank,proto3" json:"payee_bank,omitempty"`
	PlatformAccount      string       `protobuf:"bytes,8,opt,name=platform_account,json=platformAccount,proto3" json:"platform_account,omitempty"`
	Status               string       `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string       `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string       `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Distributor          *Distributor `protobuf:"bytes,12,opt,name=distributor,proto3" json:"distributor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CashAccount) Reset()         { *m = CashAccount{} }
func (m *CashAccount) String() string { return proto.CompactTextString(m) }
func (*CashAccount) ProtoMessage()    {}
func (*CashAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b8a2214f459de5b, []int{0}
}

func (m *CashAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CashAccount.Unmarshal(m, b)
}
func (m *CashAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CashAccount.Marshal(b, m, deterministic)
}
func (m *CashAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CashAccount.Merge(m, src)
}
func (m *CashAccount) XXX_Size() int {
	return xxx_messageInfo_CashAccount.Size(m)
}
func (m *CashAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_CashAccount.DiscardUnknown(m)
}

var xxx_messageInfo_CashAccount proto.InternalMessageInfo

func (m *CashAccount) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CashAccount) GetDistributorId() int64 {
	if m != nil {
		return m.DistributorId
	}
	return 0
}

func (m *CashAccount) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CashAccount) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *CashAccount) GetPayeeName() string {
	if m != nil {
		return m.PayeeName
	}
	return ""
}

func (m *CashAccount) GetPayeeAccount() string {
	if m != nil {
		return m.PayeeAccount
	}
	return ""
}

func (m *CashAccount) GetPayeeBank() string {
	if m != nil {
		return m.PayeeBank
	}
	return ""
}

func (m *CashAccount) GetPlatformAccount() string {
	if m != nil {
		return m.PlatformAccount
	}
	return ""
}

func (m *CashAccount) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CashAccount) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *CashAccount) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *CashAccount) GetDistributor() *Distributor {
	if m != nil {
		return m.Distributor
	}
	return nil
}

type CashAccountResponse struct {
	Entity               *CashAccount   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager         `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*CashAccount `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error         `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info          `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CashAccountResponse) Reset()         { *m = CashAccountResponse{} }
func (m *CashAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CashAccountResponse) ProtoMessage()    {}
func (*CashAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b8a2214f459de5b, []int{1}
}

func (m *CashAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CashAccountResponse.Unmarshal(m, b)
}
func (m *CashAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CashAccountResponse.Marshal(b, m, deterministic)
}
func (m *CashAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CashAccountResponse.Merge(m, src)
}
func (m *CashAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CashAccountResponse.Size(m)
}
func (m *CashAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CashAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CashAccountResponse proto.InternalMessageInfo

func (m *CashAccountResponse) GetEntity() *CashAccount {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *CashAccountResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *CashAccountResponse) GetItems() []*CashAccount {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CashAccountResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CashAccountResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*CashAccount)(nil), "geiqin.srv.dms.CashAccount")
	proto.RegisterType((*CashAccountResponse)(nil), "geiqin.srv.dms.CashAccountResponse")
}

func init() { proto.RegisterFile("cashAccount.proto", fileDescriptor_4b8a2214f459de5b) }

var fileDescriptor_4b8a2214f459de5b = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x5f, 0x6b, 0x13, 0x41,
	0x14, 0xc5, 0xcd, 0x6e, 0xb3, 0x36, 0x77, 0xd3, 0x68, 0xc7, 0x3f, 0x8c, 0x29, 0x42, 0x48, 0x11,
	0x22, 0xc2, 0x82, 0xe9, 0xb3, 0x0f, 0x69, 0x2b, 0x12, 0x11, 0x91, 0xed, 0x83, 0x8f, 0x61, 0xb2,
	0x7b, 0xd3, 0x0c, 0xcd, 0xce, 0xac, 0x33, 0x93, 0x42, 0x1e, 0xfd, 0x16, 0x7e, 0x23, 0xbf, 0x96,
	0xec, 0x9d, 0x4d, 0xbb, 0x86, 0x52, 0x7d, 0xe8, 0x5b, 0x72, 0xce, 0xb9, 0xbf, 0x7b, 0x99, 0xc3,
	0xc2, 0x61, 0x26, 0xec, 0x72, 0x92, 0x65, 0x7a, 0xad, 0x5c, 0x52, 0x1a, 0xed, 0x34, 0xeb, 0x5d,
	0xa2, 0xfc, 0x21, 0x55, 0x62, 0xcd, 0x75, 0x92, 0x17, 0xb6, 0xdf, 0xcd, 0x74, 0x51, 0x68, 0xe5,
	0xdd, 0xfe, 0x61, 0x2e, 0xad, 0x33, 0x72, 0xbe, 0x76, 0xda, 0x78, 0x69, 0xf8, 0x2b, 0x84, 0xf8,
	0xec, 0x16, 0xc3, 0x7a, 0x10, 0xc8, 0x9c, 0xb7, 0x06, 0xad, 0x51, 0x98, 0x06, 0x32, 0x67, 0x6f,
	0xa0, 0xd7, 0x18, 0x9a, 0xc9, 0x9c, 0x07, 0xe4, 0x1d, 0x34, 0xd4, 0x69, 0xce, 0x18, 0xec, 0xb9,
	0x4d, 0x89, 0x3c, 0x1c, 0xb4, 0x46, 0x9d, 0x94, 0x7e, 0xb3, 0x23, 0xe8, 0x18, 0x14, 0xab, 0x99,
	0x12, 0x05, 0xf2, 0x3d, 0x32, 0xf6, 0x2b, 0xe1, 0xab, 0x28, 0x90, 0xbd, 0x06, 0x28, 0xc5, 0x06,
	0xd1, 0xbb, 0x6d, 0x72, 0x3b, 0xa4, 0x90, 0x7d, 0x0c, 0x07, 0xde, 0x16, 0xfe, 0x2e, 0x1e, 0x51,
	0xa2, 0x4b, 0xe2, 0xf6, 0xd6, 0x1b, 0xc6, 0x5c, 0xa8, 0x2b, 0xfe, 0xb8, 0xc1, 0x38, 0x15, 0xea,
	0x8a, 0xbd, 0x85, 0xa7, 0xe5, 0x4a, 0xb8, 0x85, 0x36, 0xc5, 0x0d, 0x66, 0x9f, 0x42, 0x4f, 0xb6,
	0xfa, 0x96, 0xf4, 0x12, 0x22, 0xeb, 0x84, 0x5b, 0x5b, 0xde, 0xa1, 0x40, 0xfd, 0xaf, 0xda, 0x90,
	0x19, 0x14, 0x0e, 0xf3, 0x99, 0x70, 0x1c, 0xfc, 0x86, 0x5a, 0x99, 0xd0, 0x01, 0xeb, 0x32, 0xdf,
	0xda, 0xb1, 0xb7, 0x6b, 0x65, 0xe2, 0xd8, 0x07, 0x88, 0x1b, 0xaf, 0xc4, 0xbb, 0x83, 0xd6, 0x28,
	0x1e, 0x1f, 0x25, 0x7f, 0x57, 0x94, 0x9c, 0xdf, 0x46, 0xd2, 0x66, 0x7e, 0xf8, 0x33, 0x80, 0x67,
	0x8d, 0x6a, 0x52, 0xb4, 0xa5, 0x56, 0x16, 0xd9, 0x09, 0x44, 0xa8, 0x9c, 0x74, 0x1b, 0xaa, 0xe9,
	0x0e, 0x62, 0x73, 0xa8, 0x8e, 0xb2, 0x77, 0xd0, 0x2e, 0xc5, 0x25, 0x1a, 0xaa, 0x2f, 0x1e, 0xbf,
	0xd8, 0x9d, 0xf9, 0x56, 0x99, 0xa9, 0xcf, 0xb0, 0xf7, 0xd0, 0x96, 0x0e, 0x0b, 0xcb, 0xc3, 0x41,
	0xf8, 0xaf, 0x05, 0x3e, 0x59, 0xf1, 0xd1, 0x18, 0x6d, 0xa8, 0xe8, 0x3b, 0xf8, 0x1f, 0x2b, 0x33,
	0xf5, 0x19, 0x36, 0x82, 0x3d, 0xa9, 0x16, 0x9a, 0x6a, 0x8f, 0xc7, 0xcf, 0x77, 0xb3, 0x53, 0xb5,
	0xd0, 0x29, 0x25, 0xc6, 0xbf, 0x03, 0x60, 0x8d, 0x6d, 0x17, 0x68, 0xae, 0x65, 0x86, 0xec, 0x0b,
	0x44, 0x67, 0xd4, 0x02, 0xbb, 0xef, 0xb6, 0xfe, 0xf1, 0x7d, 0x87, 0xd7, 0xcf, 0x39, 0x7c, 0x54,
	0xd1, 0xce, 0x71, 0x85, 0x0f, 0x44, 0x9b, 0x42, 0xf8, 0x09, 0xdd, 0x83, 0xa0, 0x3e, 0x43, 0x74,
	0x81, 0xc2, 0x64, 0x4b, 0xf6, 0x6a, 0x77, 0xe0, 0x54, 0x58, 0xfc, 0xbe, 0x44, 0x83, 0xff, 0xc9,
	0x9a, 0x47, 0xf4, 0xbd, 0x9f, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xad, 0x5a, 0x01, 0xd1, 0x35,
	0x04, 0x00, 0x00,
}
