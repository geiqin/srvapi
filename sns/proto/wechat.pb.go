// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wechat.proto

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

type Wechat struct {
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

func (m *Wechat) Reset()         { *m = Wechat{} }
func (m *Wechat) String() string { return proto.CompactTextString(m) }
func (*Wechat) ProtoMessage()    {}
func (*Wechat) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74b134486e7456c, []int{0}
}

func (m *Wechat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wechat.Unmarshal(m, b)
}
func (m *Wechat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wechat.Marshal(b, m, deterministic)
}
func (m *Wechat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wechat.Merge(m, src)
}
func (m *Wechat) XXX_Size() int {
	return xxx_messageInfo_Wechat.Size(m)
}
func (m *Wechat) XXX_DiscardUnknown() {
	xxx_messageInfo_Wechat.DiscardUnknown(m)
}

var xxx_messageInfo_Wechat proto.InternalMessageInfo

func (m *Wechat) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *Wechat) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Wechat) GetRawData() string {
	if m != nil {
		return m.RawData
	}
	return ""
}

func (m *Wechat) GetEncryptedData() string {
	if m != nil {
		return m.EncryptedData
	}
	return ""
}

func (m *Wechat) GetIv() string {
	if m != nil {
		return m.Iv
	}
	return ""
}

func (m *Wechat) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

type WechatResponse struct {
	Entity               *Wechat   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Wechat `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error    `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info     `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *WechatResponse) Reset()         { *m = WechatResponse{} }
func (m *WechatResponse) String() string { return proto.CompactTextString(m) }
func (*WechatResponse) ProtoMessage()    {}
func (*WechatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74b134486e7456c, []int{1}
}

func (m *WechatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WechatResponse.Unmarshal(m, b)
}
func (m *WechatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WechatResponse.Marshal(b, m, deterministic)
}
func (m *WechatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WechatResponse.Merge(m, src)
}
func (m *WechatResponse) XXX_Size() int {
	return xxx_messageInfo_WechatResponse.Size(m)
}
func (m *WechatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WechatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WechatResponse proto.InternalMessageInfo

func (m *WechatResponse) GetEntity() *Wechat {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *WechatResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *WechatResponse) GetItems() []*Wechat {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *WechatResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *WechatResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Wechat)(nil), "geiqin.srv.sns.Wechat")
	proto.RegisterType((*WechatResponse)(nil), "geiqin.srv.sns.WechatResponse")
}

func init() {
	proto.RegisterFile("wechat.proto", fileDescriptor_e74b134486e7456c)
}

var fileDescriptor_e74b134486e7456c = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0x17, 0xe7, 0xcf, 0xd6, 0xe3, 0x36, 0x03, 0xd1, 0x0d, 0x11, 0xb6, 0x11, 0xc2, 0x2e,
	0x0a, 0x1b, 0xbe, 0xf0, 0x1e, 0x61, 0x6d, 0x60, 0x8c, 0x40, 0x71, 0x0a, 0xbd, 0x56, 0x95, 0x93,
	0x44, 0x24, 0x91, 0x3c, 0x49, 0x75, 0xc9, 0xdb, 0xec, 0x11, 0x76, 0xb7, 0xf7, 0xda, 0x13, 0x0c,
	0x1f, 0xb9, 0xb4, 0x36, 0x75, 0x2f, 0x9a, 0xbb, 0x9c, 0xf3, 0x7d, 0xf9, 0xf9, 0xe8, 0xd3, 0xb1,
	0xe1, 0xf8, 0x0e, 0xe5, 0x5a, 0xf8, 0x24, 0xb7, 0xc6, 0x1b, 0x36, 0x5c, 0xa1, 0xfa, 0xa5, 0x74,
	0xe2, 0x6c, 0x91, 0x38, 0xed, 0x46, 0xc7, 0xd2, 0xec, 0x76, 0x46, 0x07, 0x75, 0x34, 0x5c, 0x0a,
	0x3d, 0x5f, 0x0b, 0x8b, 0x8f, 0xea, 0x2b, 0x2b, 0xe4, 0x26, 0xd4, 0x93, 0xdf, 0x1d, 0x18, 0x5c,
	0x13, 0x8e, 0x7d, 0x80, 0x23, 0xa7, 0x56, 0x5a, 0xf8, 0x5b, 0x8b, 0xbc, 0x33, 0xee, 0x9c, 0x1d,
	0x65, 0x0f, 0x0d, 0xc6, 0xa0, 0x27, 0xcd, 0x02, 0x79, 0x44, 0x02, 0xfd, 0x66, 0x1c, 0x5e, 0x5b,
	0x71, 0x77, 0x2e, 0xbc, 0xe0, 0x5d, 0x6a, 0xdf, 0x97, 0xec, 0x33, 0x9c, 0xa0, 0x96, 0x76, 0x9f,
	0x7b, 0x5c, 0x90, 0xde, 0x23, 0xbd, 0xde, 0x64, 0x43, 0x88, 0x54, 0xc1, 0xfb, 0x24, 0x45, 0xaa,
	0x60, 0xa7, 0xd0, 0x77, 0x12, 0x35, 0xf2, 0x01, 0xb5, 0x42, 0x31, 0xf9, 0xd7, 0x81, 0x61, 0x18,
	0x31, 0x43, 0x97, 0x1b, 0xed, 0x90, 0x25, 0x30, 0x40, 0xed, 0x95, 0xdf, 0xd3, 0x9c, 0x71, 0xfa,
	0x3e, 0xa9, 0x87, 0x90, 0x54, 0xfe, 0xca, 0xc5, 0xbe, 0x40, 0x3f, 0x17, 0x2b, 0xb4, 0x34, 0x7d,
	0x9c, 0xbe, 0x6b, 0xda, 0x2f, 0x4b, 0x31, 0x0b, 0x1e, 0xf6, 0x15, 0xfa, 0xca, 0xe3, 0xce, 0xf1,
	0xee, 0xb8, 0xfb, 0x0c, 0x3b, 0x98, 0x4a, 0x34, 0x5a, 0x6b, 0x2c, 0x9d, 0xf0, 0x09, 0xf4, 0x45,
	0x29, 0x66, 0xc1, 0xc3, 0xce, 0xa0, 0xa7, 0xf4, 0xd2, 0xd0, 0x91, 0xe3, 0xf4, 0xb4, 0xe9, 0xfd,
	0xa1, 0x97, 0x26, 0x23, 0x47, 0xfa, 0x27, 0x82, 0xb7, 0xb3, 0x7d, 0x78, 0xd4, 0x1c, 0x6d, 0xa1,
	0x24, 0xb2, 0x29, 0xc4, 0x19, 0x96, 0xb7, 0x8b, 0x7a, 0x81, 0x96, 0xb5, 0x0c, 0x36, 0xfa, 0xd8,
	0xec, 0x5f, 0x99, 0x0d, 0xea, 0xfb, 0xec, 0x26, 0xaf, 0xd8, 0x05, 0xc0, 0xb5, 0x55, 0x1e, 0x69,
	0x2f, 0x5e, 0x8e, 0x99, 0x56, 0x18, 0x5a, 0xa7, 0x56, 0xcc, 0xa7, 0x96, 0xf8, 0x6a, 0x9c, 0xcb,
	0xdb, 0xed, 0x76, 0x66, 0x6e, 0xd4, 0x16, 0x5f, 0xce, 0x49, 0xff, 0x46, 0x70, 0x52, 0x0f, 0xec,
	0x1c, 0xde, 0xcc, 0xc4, 0x06, 0xbf, 0x97, 0xbb, 0x7a, 0xd0, 0x7c, 0x25, 0x61, 0x8e, 0xc2, 0xca,
	0xf5, 0x01, 0x9c, 0x9f, 0x10, 0x53, 0x54, 0x15, 0x88, 0x37, 0xff, 0x30, 0xad, 0xde, 0xcc, 0xd1,
	0xb8, 0x4d, 0xa9, 0xc3, 0xe8, 0xfa, 0x9e, 0x81, 0x91, 0xfe, 0x24, 0x8c, 0x94, 0x07, 0xd8, 0xcd,
	0x80, 0xbe, 0x05, 0xdf, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xee, 0xae, 0x0f, 0x59, 0x04,
	0x00, 0x00,
}
