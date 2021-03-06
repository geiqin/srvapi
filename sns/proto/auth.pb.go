// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package geiqin_srv_sns

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

type WxMiniLogin struct {
	Signature            string   `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	RawData              string   `protobuf:"bytes,3,opt,name=rawData,proto3" json:"rawData,omitempty"`
	EncryptedData        string   `protobuf:"bytes,4,opt,name=encryptedData,proto3" json:"encryptedData,omitempty"`
	Iv                   string   `protobuf:"bytes,5,opt,name=iv,proto3" json:"iv,omitempty"`
	Scene                string   `protobuf:"bytes,6,opt,name=scene,proto3" json:"scene,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WxMiniLogin) Reset()         { *m = WxMiniLogin{} }
func (m *WxMiniLogin) String() string { return proto.CompactTextString(m) }
func (*WxMiniLogin) ProtoMessage()    {}
func (*WxMiniLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *WxMiniLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WxMiniLogin.Unmarshal(m, b)
}
func (m *WxMiniLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WxMiniLogin.Marshal(b, m, deterministic)
}
func (m *WxMiniLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WxMiniLogin.Merge(m, src)
}
func (m *WxMiniLogin) XXX_Size() int {
	return xxx_messageInfo_WxMiniLogin.Size(m)
}
func (m *WxMiniLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_WxMiniLogin.DiscardUnknown(m)
}

var xxx_messageInfo_WxMiniLogin proto.InternalMessageInfo

func (m *WxMiniLogin) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *WxMiniLogin) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *WxMiniLogin) GetRawData() string {
	if m != nil {
		return m.RawData
	}
	return ""
}

func (m *WxMiniLogin) GetEncryptedData() string {
	if m != nil {
		return m.EncryptedData
	}
	return ""
}

func (m *WxMiniLogin) GetIv() string {
	if m != nil {
		return m.Iv
	}
	return ""
}

func (m *WxMiniLogin) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

type AlipayLogin struct {
	Signature            string   `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlipayLogin) Reset()         { *m = AlipayLogin{} }
func (m *AlipayLogin) String() string { return proto.CompactTextString(m) }
func (*AlipayLogin) ProtoMessage()    {}
func (*AlipayLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AlipayLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlipayLogin.Unmarshal(m, b)
}
func (m *AlipayLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlipayLogin.Marshal(b, m, deterministic)
}
func (m *AlipayLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlipayLogin.Merge(m, src)
}
func (m *AlipayLogin) XXX_Size() int {
	return xxx_messageInfo_AlipayLogin.Size(m)
}
func (m *AlipayLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_AlipayLogin.DiscardUnknown(m)
}

var xxx_messageInfo_AlipayLogin proto.InternalMessageInfo

func (m *AlipayLogin) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *AlipayLogin) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto.RegisterType((*WxMiniLogin)(nil), "geiqin.srv.sns.WxMiniLogin")
	proto.RegisterType((*AlipayLogin)(nil), "geiqin.srv.sns.AlipayLogin")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0xbd, 0x33, 0x39, 0xc9, 0x9c, 0xa6, 0x18, 0x2c, 0x96, 0xa8, 0x20, 0x87, 0x85, 0xd5,
	0x16, 0xfa, 0x00, 0x12, 0x10, 0x2b, 0x6d, 0xa2, 0x60, 0xbd, 0x6e, 0x86, 0xcb, 0xa0, 0x99, 0x3d,
	0x77, 0xf7, 0x4e, 0xf3, 0x30, 0x96, 0xbe, 0xa7, 0xb8, 0x87, 0x98, 0x68, 0x21, 0xd8, 0xed, 0x7c,
	0x1f, 0xfc, 0xfc, 0x3b, 0x03, 0x60, 0xda, 0xb8, 0xd0, 0x8d, 0x77, 0xd1, 0xe1, 0xb8, 0x26, 0x7e,
	0x66, 0xd1, 0xc1, 0x77, 0x3a, 0x48, 0x98, 0xec, 0x5a, 0xb7, 0x5c, 0x3a, 0xe9, 0x6d, 0xf5, 0x9e,
	0x41, 0x79, 0xff, 0x7a, 0xc3, 0xc2, 0xd7, 0xae, 0x66, 0xc1, 0x43, 0x18, 0x05, 0xae, 0xc5, 0xc4,
	0xd6, 0x93, 0xca, 0x8e, 0xb3, 0xd3, 0xd1, 0xec, 0x1b, 0x20, 0xc2, 0xc0, 0xba, 0x39, 0xa9, 0x3c,
	0x89, 0xf4, 0x46, 0x05, 0x3b, 0xde, 0xbc, 0x5c, 0x9a, 0x68, 0xd4, 0x76, 0xc2, 0x5f, 0x23, 0x9e,
	0xc0, 0x1e, 0x89, 0xf5, 0xab, 0x26, 0xd2, 0x3c, 0xf9, 0x41, 0xf2, 0x9b, 0x10, 0xc7, 0x90, 0x73,
	0xa7, 0x86, 0x49, 0xe5, 0xdc, 0xe1, 0x3e, 0x0c, 0x83, 0x25, 0x21, 0x55, 0x24, 0xd4, 0x0f, 0xd5,
	0x05, 0x94, 0xd3, 0x27, 0x6e, 0xcc, 0xea, 0x9f, 0x35, 0xcf, 0xde, 0x32, 0x28, 0xa7, 0x6d, 0x5c,
	0xdc, 0x92, 0xef, 0xd8, 0x12, 0x5e, 0x41, 0xd1, 0xff, 0x1b, 0x0f, 0xf4, 0xe6, 0x86, 0xf4, 0xda,
	0x3e, 0x26, 0x47, 0x3f, 0xe5, 0x9d, 0x7b, 0x24, 0x99, 0x51, 0x68, 0x9c, 0x04, 0xaa, 0xb6, 0x3e,
	0x73, 0xfa, 0x62, 0xbf, 0x73, 0xd6, 0x0a, 0xff, 0x99, 0xf3, 0x50, 0xa4, 0x7b, 0x9c, 0x7f, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x14, 0xd4, 0x88, 0xd9, 0xbb, 0x01, 0x00, 0x00,
}
