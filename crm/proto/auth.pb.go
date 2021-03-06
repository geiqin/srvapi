// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package geiqin_srv_crm

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

type AuthRequest struct {
	CustomerId           int64    `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	OldPwd               string   `protobuf:"bytes,6,opt,name=old_pwd,json=oldPwd,proto3" json:"old_pwd,omitempty"`
	NewPwd               string   `protobuf:"bytes,7,opt,name=new_pwd,json=newPwd,proto3" json:"new_pwd,omitempty"`
	ConfirmPwd           string   `protobuf:"bytes,8,opt,name=confirm_pwd,json=confirmPwd,proto3" json:"confirm_pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *AuthRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *AuthRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthRequest) GetOldPwd() string {
	if m != nil {
		return m.OldPwd
	}
	return ""
}

func (m *AuthRequest) GetNewPwd() string {
	if m != nil {
		return m.NewPwd
	}
	return ""
}

func (m *AuthRequest) GetConfirmPwd() string {
	if m != nil {
		return m.ConfirmPwd
	}
	return ""
}

//
type AuthResponse struct {
	Entity               *Token   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetEntity() *Token {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *AuthResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *AuthResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *AuthResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "geiqin.srv.crm.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "geiqin.srv.crm.AuthResponse")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0xc9, 0x4c, 0x92, 0x99, 0x79, 0x19, 0xcd, 0xc2, 0x1a, 0x20, 0x0a, 0x20, 0xaa, 0xae,
	0x2a, 0x21, 0xb2, 0x28, 0x07, 0x40, 0x23, 0x34, 0x8b, 0x4a, 0xb3, 0xa8, 0x02, 0xac, 0xab, 0x34,
	0x7e, 0x4d, 0x2d, 0x62, 0xbf, 0xd4, 0x71, 0x1a, 0xf5, 0x3c, 0x9c, 0x84, 0x93, 0x70, 0x15, 0x14,
	0x3b, 0x45, 0xa5, 0x84, 0x15, 0xbb, 0xbc, 0xff, 0xfb, 0xf3, 0xcb, 0xff, 0xb3, 0x01, 0xf2, 0xd6,
	0x6c, 0xd3, 0x5a, 0x93, 0x21, 0x76, 0x57, 0xa2, 0xd8, 0x09, 0x95, 0x36, 0x7a, 0x9f, 0x16, 0x5a,
	0x26, 0xb7, 0x05, 0x49, 0x49, 0xca, 0xd1, 0xe4, 0xae, 0x68, 0x1b, 0x43, 0x12, 0xb5, 0x9b, 0xa7,
	0x3f, 0x3d, 0x88, 0x1e, 0x5a, 0xb3, 0xcd, 0x70, 0xd7, 0x62, 0x63, 0xd8, 0x5b, 0x88, 0x8e, 0x8e,
	0x95, 0xe0, 0xb1, 0x37, 0xf1, 0x66, 0x97, 0x19, 0x1c, 0xa5, 0x05, 0x67, 0x0c, 0x7c, 0x95, 0x4b,
	0x8c, 0x2f, 0x26, 0xde, 0xec, 0x26, 0xb3, 0xdf, 0xec, 0x05, 0x84, 0x92, 0xd6, 0xa2, 0xc2, 0xf8,
	0xd2, 0xaa, 0xc3, 0xc4, 0xee, 0x21, 0x40, 0x99, 0x8b, 0x2a, 0xf6, 0xad, 0xec, 0x06, 0x96, 0xc0,
	0x75, 0x9d, 0x37, 0x4d, 0x47, 0x9a, 0xc7, 0x81, 0x05, 0xbf, 0x67, 0xf6, 0x12, 0xae, 0xa8, 0xe2,
	0xab, 0xba, 0xe3, 0x71, 0xe8, 0xa2, 0xa8, 0xe2, 0xcb, 0xce, 0x02, 0x85, 0x9d, 0x05, 0x57, 0x0e,
	0x28, 0xec, 0x7a, 0xd0, 0x1f, 0x98, 0xd4, 0x46, 0x68, 0x69, 0xe1, 0xb5, 0x85, 0x30, 0x48, 0xcb,
	0x8e, 0x4f, 0x7f, 0x78, 0x70, 0xeb, 0x1a, 0x36, 0x35, 0xa9, 0x06, 0xd9, 0x7b, 0x08, 0x51, 0x19,
	0x61, 0x0e, 0xb6, 0x5d, 0x34, 0x7f, 0x9e, 0xfe, 0xb9, 0xb1, 0xf4, 0x0b, 0x7d, 0x43, 0x95, 0x0d,
	0x26, 0xf6, 0x0e, 0x82, 0x3a, 0x2f, 0x51, 0xdb, 0xc6, 0x23, 0xee, 0x65, 0x0f, 0x33, 0xe7, 0xe9,
	0xcd, 0xa8, 0x35, 0x69, 0xdb, 0x78, 0xc4, 0xfc, 0xd8, 0xc3, 0xcc, 0x79, 0xd8, 0x0c, 0x7c, 0xa1,
	0x36, 0x64, 0x97, 0x10, 0xcd, 0xef, 0xcf, 0xbd, 0x0b, 0xb5, 0xa1, 0xcc, 0x3a, 0xe6, 0xdf, 0x2f,
	0xdc, 0x2d, 0x7d, 0x46, 0xbd, 0x17, 0x05, 0xb2, 0x47, 0x08, 0x9e, 0xa8, 0x14, 0x8a, 0xbd, 0x3a,
	0xff, 0xe9, 0xe4, 0x2e, 0x93, 0x37, 0xe3, 0xc5, 0x86, 0x3d, 0x4c, 0x9f, 0xb1, 0x05, 0xdc, 0x7c,
	0xad, 0x79, 0x6e, 0xb0, 0x5f, 0xe4, 0xff, 0x45, 0x3d, 0x80, 0xdf, 0x9f, 0x97, 0xfd, 0xdd, 0x58,
	0xd6, 0xe6, 0x90, 0x4c, 0xce, 0xe5, 0x4f, 0xc3, 0x8b, 0x3a, 0x89, 0xf8, 0x08, 0xe1, 0x13, 0x95,
	0xd4, 0x9a, 0x7f, 0x85, 0xbc, 0x1e, 0x3f, 0xe1, 0x31, 0x60, 0x1d, 0xda, 0x27, 0xfd, 0xe1, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0xa2, 0xbd, 0xa6, 0x0e, 0x03, 0x00, 0x00,
}
