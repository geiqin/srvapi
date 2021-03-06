// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shipper.proto

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

type Shipper struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	LogoId               string   `protobuf:"bytes,4,opt,name=logo_id,json=logoId,proto3" json:"logo_id,omitempty"`
	LogoUrl              string   `protobuf:"bytes,5,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	Tel                  string   `protobuf:"bytes,6,opt,name=tel,proto3" json:"tel,omitempty"`
	Memo                 string   `protobuf:"bytes,7,opt,name=memo,proto3" json:"memo,omitempty"`
	Foreign              bool     `protobuf:"varint,8,opt,name=foreign,proto3" json:"foreign,omitempty"`
	IsAccess             bool     `protobuf:"varint,9,opt,name=is_access,json=isAccess,proto3" json:"is_access,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Shipper) Reset()         { *m = Shipper{} }
func (m *Shipper) String() string { return proto.CompactTextString(m) }
func (*Shipper) ProtoMessage()    {}
func (*Shipper) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2a14d76411df457, []int{0}
}

func (m *Shipper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shipper.Unmarshal(m, b)
}
func (m *Shipper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shipper.Marshal(b, m, deterministic)
}
func (m *Shipper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shipper.Merge(m, src)
}
func (m *Shipper) XXX_Size() int {
	return xxx_messageInfo_Shipper.Size(m)
}
func (m *Shipper) XXX_DiscardUnknown() {
	xxx_messageInfo_Shipper.DiscardUnknown(m)
}

var xxx_messageInfo_Shipper proto.InternalMessageInfo

func (m *Shipper) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Shipper) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Shipper) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Shipper) GetLogoId() string {
	if m != nil {
		return m.LogoId
	}
	return ""
}

func (m *Shipper) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func (m *Shipper) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

func (m *Shipper) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Shipper) GetForeign() bool {
	if m != nil {
		return m.Foreign
	}
	return false
}

func (m *Shipper) GetIsAccess() bool {
	if m != nil {
		return m.IsAccess
	}
	return false
}

type ShipperResponse struct {
	Entity               *Shipper   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Shipper `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ShipperResponse) Reset()         { *m = ShipperResponse{} }
func (m *ShipperResponse) String() string { return proto.CompactTextString(m) }
func (*ShipperResponse) ProtoMessage()    {}
func (*ShipperResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2a14d76411df457, []int{1}
}

func (m *ShipperResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipperResponse.Unmarshal(m, b)
}
func (m *ShipperResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipperResponse.Marshal(b, m, deterministic)
}
func (m *ShipperResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipperResponse.Merge(m, src)
}
func (m *ShipperResponse) XXX_Size() int {
	return xxx_messageInfo_ShipperResponse.Size(m)
}
func (m *ShipperResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipperResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShipperResponse proto.InternalMessageInfo

func (m *ShipperResponse) GetEntity() *Shipper {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ShipperResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ShipperResponse) GetItems() []*Shipper {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ShipperResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ShipperResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Shipper)(nil), "geiqin.srv.tms.Shipper")
	proto.RegisterType((*ShipperResponse)(nil), "geiqin.srv.tms.ShipperResponse")
}

func init() {
	proto.RegisterFile("shipper.proto", fileDescriptor_b2a14d76411df457)
}

var fileDescriptor_b2a14d76411df457 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcd, 0x8a, 0xd4, 0x40,
	0x10, 0xc7, 0xcd, 0x77, 0xa6, 0x46, 0x47, 0x29, 0x94, 0xed, 0x5d, 0x0f, 0x86, 0x39, 0x05, 0xc4,
	0x08, 0xf1, 0x09, 0xd6, 0x4f, 0x16, 0x3c, 0x48, 0x86, 0xc5, 0xe3, 0x12, 0x93, 0x9a, 0x4c, 0x43,
	0xd2, 0x1d, 0xbb, 0xdb, 0x01, 0x5f, 0xd3, 0xf7, 0xf0, 0xe8, 0x5d, 0xba, 0x93, 0x39, 0x38, 0x38,
	0x22, 0xcc, 0xad, 0xaa, 0xfe, 0xbf, 0xf9, 0x4d, 0x57, 0x41, 0xe0, 0x81, 0xde, 0xf1, 0x71, 0x24,
	0x55, 0x8c, 0x4a, 0x1a, 0x89, 0xab, 0x8e, 0xf8, 0x57, 0x2e, 0x0a, 0xad, 0xf6, 0x85, 0x19, 0xf4,
	0xd5, 0xfd, 0x46, 0x0e, 0x83, 0x14, 0x53, 0xba, 0xfe, 0xe1, 0x41, 0xb2, 0x99, 0x78, 0x5c, 0x81,
	0xcf, 0x5b, 0xe6, 0x65, 0x5e, 0x1e, 0x55, 0x3e, 0x6f, 0x11, 0x21, 0x6c, 0x64, 0x4b, 0xcc, 0xcf,
	0xbc, 0x7c, 0x51, 0xb9, 0xda, 0xce, 0x44, 0x3d, 0x10, 0x0b, 0xa6, 0x99, 0xad, 0xf1, 0x02, 0x92,
	0x5e, 0x76, 0xf2, 0x8e, 0xb7, 0x2c, 0x74, 0xe3, 0xd8, 0xb6, 0x37, 0x2d, 0x5e, 0x42, 0xea, 0x82,
	0x6f, 0xaa, 0x67, 0x91, 0x4b, 0x1c, 0x78, 0xab, 0x7a, 0x7c, 0x04, 0x81, 0xa1, 0x9e, 0xc5, 0x6e,
	0x6a, 0x4b, 0x6b, 0x1e, 0x68, 0x90, 0x2c, 0x99, 0xcc, 0xb6, 0x46, 0x06, 0xc9, 0x56, 0x2a, 0xe2,
	0x9d, 0x60, 0x69, 0xe6, 0xe5, 0x69, 0x75, 0x68, 0xf1, 0x29, 0x2c, 0xb8, 0xbe, 0xab, 0x9b, 0x86,
	0xb4, 0x66, 0x0b, 0x97, 0xa5, 0x5c, 0x5f, 0xbb, 0x7e, 0xfd, 0xcb, 0x83, 0x87, 0xf3, 0x52, 0x15,
	0xe9, 0x51, 0x0a, 0x4d, 0xf8, 0x12, 0x62, 0x12, 0x86, 0x9b, 0xef, 0x6e, 0xc1, 0x65, 0x79, 0x51,
	0xfc, 0x79, 0x97, 0xe2, 0xf0, 0x83, 0x19, 0xc3, 0xe7, 0x10, 0x8d, 0x75, 0x47, 0xca, 0xad, 0xbf,
	0x2c, 0x9f, 0x1c, 0xf3, 0x9f, 0x6c, 0x58, 0x4d, 0x0c, 0xbe, 0x80, 0x88, 0x1b, 0x1a, 0x34, 0x0b,
	0xb2, 0xe0, 0x5f, 0xf2, 0x89, 0xb2, 0x6e, 0x52, 0x4a, 0x2a, 0x77, 0xaf, 0xbf, 0xb8, 0xdf, 0xd9,
	0xb0, 0x9a, 0x18, 0xcc, 0x21, 0xe4, 0x62, 0x2b, 0xdd, 0x05, 0x97, 0xe5, 0xe3, 0x63, 0xf6, 0x46,
	0x6c, 0x65, 0xe5, 0x88, 0xf2, 0xa7, 0x0f, 0xab, 0xf9, 0x9f, 0x36, 0xa4, 0xf6, 0xbc, 0x21, 0x7c,
	0x0b, 0xf1, 0x1b, 0x45, 0xb5, 0x21, 0x3c, 0xf5, 0xa6, 0xab, 0x67, 0xa7, 0x1e, 0x3b, 0x9f, 0x6e,
	0x7d, 0xcf, 0x5a, 0x6e, 0xc7, 0xf6, 0x5c, 0xcb, 0x35, 0x04, 0x1f, 0xc8, 0x9c, 0xa5, 0x78, 0x0f,
	0xf1, 0x86, 0x6a, 0xd5, 0xec, 0xf0, 0xf2, 0x18, 0x7e, 0x5d, 0x6b, 0xfa, 0xbc, 0x23, 0x45, 0xff,
	0xb7, 0x50, 0xf8, 0x91, 0x6b, 0x73, 0x9e, 0xe5, 0x4b, 0xec, 0xbe, 0xa1, 0x57, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x7b, 0xfb, 0xc3, 0x14, 0x72, 0x03, 0x00, 0x00,
}
