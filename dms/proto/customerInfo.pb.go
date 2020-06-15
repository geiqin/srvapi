// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customerInfo.proto

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

type CustomerInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerSn           string   `protobuf:"bytes,2,opt,name=customer_sn,json=customerSn,proto3" json:"customer_sn,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string   `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Realname             string   `protobuf:"bytes,5,opt,name=realname,proto3" json:"realname,omitempty"`
	Idcard               string   `protobuf:"bytes,6,opt,name=idcard,proto3" json:"idcard,omitempty"`
	Gender               int32    `protobuf:"varint,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Mobile               string   `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email                string   `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
	AvatarId             int64    `protobuf:"varint,10,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,11,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerInfo) Reset()         { *m = CustomerInfo{} }
func (m *CustomerInfo) String() string { return proto.CompactTextString(m) }
func (*CustomerInfo) ProtoMessage()    {}
func (*CustomerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_84596b75fdcaddfa, []int{0}
}

func (m *CustomerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerInfo.Unmarshal(m, b)
}
func (m *CustomerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerInfo.Marshal(b, m, deterministic)
}
func (m *CustomerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerInfo.Merge(m, src)
}
func (m *CustomerInfo) XXX_Size() int {
	return xxx_messageInfo_CustomerInfo.Size(m)
}
func (m *CustomerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerInfo proto.InternalMessageInfo

func (m *CustomerInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerInfo) GetCustomerSn() string {
	if m != nil {
		return m.CustomerSn
	}
	return ""
}

func (m *CustomerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomerInfo) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *CustomerInfo) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *CustomerInfo) GetIdcard() string {
	if m != nil {
		return m.Idcard
	}
	return ""
}

func (m *CustomerInfo) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *CustomerInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *CustomerInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CustomerInfo) GetAvatarId() int64 {
	if m != nil {
		return m.AvatarId
	}
	return 0
}

func (m *CustomerInfo) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*CustomerInfo)(nil), "geiqin.srv.dms.CustomerInfo")
}

func init() {
	proto.RegisterFile("customerInfo.proto", fileDescriptor_84596b75fdcaddfa)
}

var fileDescriptor_84596b75fdcaddfa = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0x47, 0x95, 0xb4, 0xc9, 0x97, 0xdc, 0x54, 0x1d, 0xae, 0x3e, 0x21, 0x0b, 0x84, 0x08, 0x4c,
	0x99, 0xb2, 0xf0, 0x08, 0x4c, 0x5d, 0x18, 0x82, 0x98, 0x23, 0xb7, 0x36, 0x95, 0x25, 0xff, 0x29,
	0xd7, 0x69, 0x25, 0xde, 0x89, 0x87, 0x44, 0xb9, 0x71, 0x11, 0x5b, 0xce, 0x39, 0xbf, 0x58, 0xb2,
	0x01, 0x0f, 0xe7, 0x38, 0x05, 0xa7, 0x69, 0xe7, 0x3f, 0x42, 0x7f, 0xa2, 0x30, 0x05, 0xdc, 0x1e,
	0xb5, 0xf9, 0x34, 0xbe, 0x8f, 0x74, 0xe9, 0x95, 0x8b, 0x4f, 0xdf, 0x39, 0x6c, 0x5e, 0xfe, 0xcc,
	0x70, 0x0b, 0xb9, 0x51, 0x22, 0x6b, 0xb3, 0x6e, 0x35, 0xe4, 0x46, 0xe1, 0x03, 0x34, 0xd7, 0x63,
	0xc6, 0xe8, 0x45, 0xde, 0x66, 0x5d, 0x3d, 0xc0, 0x55, 0xbd, 0x79, 0x44, 0x58, 0x7b, 0xe9, 0xb4,
	0x58, 0x71, 0xe1, 0x6f, 0x7c, 0x84, 0x8d, 0x32, 0xf1, 0x64, 0xe5, 0xd7, 0xc8, 0x6d, 0xcd, 0xad,
	0x49, 0xee, 0x75, 0x9e, 0xdc, 0x42, 0x45, 0x5a, 0x5a, 0xce, 0x05, 0xe7, 0x5f, 0xc6, 0x1b, 0x28,
	0x8d, 0x3a, 0x48, 0x52, 0xa2, 0xe4, 0x92, 0x68, 0xf6, 0x47, 0xed, 0x95, 0x26, 0xf1, 0xaf, 0xcd,
	0xba, 0x62, 0x48, 0x34, 0x7b, 0x17, 0xf6, 0xc6, 0x6a, 0x51, 0x2d, 0xfb, 0x85, 0xf0, 0x3f, 0x14,
	0xda, 0x49, 0x63, 0x45, 0xcd, 0x7a, 0x01, 0xbc, 0x83, 0x5a, 0x5e, 0xe4, 0x24, 0x69, 0x34, 0x4a,
	0x00, 0x5f, 0xb4, 0x5a, 0xc4, 0x4e, 0xe1, 0x3d, 0x40, 0x8a, 0x67, 0xb2, 0xa2, 0xe1, 0xff, 0xd2,
	0xfc, 0x9d, 0xec, 0xbe, 0xe4, 0x57, 0x7c, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x21, 0x71, 0x2e,
	0x8c, 0x5b, 0x01, 0x00, 0x00,
}
