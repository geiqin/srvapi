// Code generated by protoc-gen-go. DO NOT EDIT.
// source: specValue.proto

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

type SpecValue struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SpecId               int32    `protobuf:"varint,2,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	Alias                string   `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	SpecValue            string   `protobuf:"bytes,4,opt,name=spec_value,json=specValue,proto3" json:"spec_value,omitempty"`
	SpecImgId            int64    `protobuf:"varint,5,opt,name=spec_img_id,json=specImgId,proto3" json:"spec_img_id,omitempty"`
	SpecImgUrl           string   `protobuf:"bytes,6,opt,name=spec_img_url,json=specImgUrl,proto3" json:"spec_img_url,omitempty"`
	Sorting              int32    `protobuf:"varint,7,opt,name=sorting,proto3" json:"sorting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecValue) Reset()         { *m = SpecValue{} }
func (m *SpecValue) String() string { return proto.CompactTextString(m) }
func (*SpecValue) ProtoMessage()    {}
func (*SpecValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3cfe65595187a36, []int{0}
}

func (m *SpecValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecValue.Unmarshal(m, b)
}
func (m *SpecValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecValue.Marshal(b, m, deterministic)
}
func (m *SpecValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecValue.Merge(m, src)
}
func (m *SpecValue) XXX_Size() int {
	return xxx_messageInfo_SpecValue.Size(m)
}
func (m *SpecValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecValue.DiscardUnknown(m)
}

var xxx_messageInfo_SpecValue proto.InternalMessageInfo

func (m *SpecValue) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SpecValue) GetSpecId() int32 {
	if m != nil {
		return m.SpecId
	}
	return 0
}

func (m *SpecValue) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SpecValue) GetSpecValue() string {
	if m != nil {
		return m.SpecValue
	}
	return ""
}

func (m *SpecValue) GetSpecImgId() int64 {
	if m != nil {
		return m.SpecImgId
	}
	return 0
}

func (m *SpecValue) GetSpecImgUrl() string {
	if m != nil {
		return m.SpecImgUrl
	}
	return ""
}

func (m *SpecValue) GetSorting() int32 {
	if m != nil {
		return m.Sorting
	}
	return 0
}

type SpecValueResponse struct {
	Entity               *SpecValue   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*SpecValue `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error       `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info        `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SpecValueResponse) Reset()         { *m = SpecValueResponse{} }
func (m *SpecValueResponse) String() string { return proto.CompactTextString(m) }
func (*SpecValueResponse) ProtoMessage()    {}
func (*SpecValueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3cfe65595187a36, []int{1}
}

func (m *SpecValueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecValueResponse.Unmarshal(m, b)
}
func (m *SpecValueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecValueResponse.Marshal(b, m, deterministic)
}
func (m *SpecValueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecValueResponse.Merge(m, src)
}
func (m *SpecValueResponse) XXX_Size() int {
	return xxx_messageInfo_SpecValueResponse.Size(m)
}
func (m *SpecValueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecValueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpecValueResponse proto.InternalMessageInfo

func (m *SpecValueResponse) GetEntity() *SpecValue {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *SpecValueResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *SpecValueResponse) GetItems() []*SpecValue {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SpecValueResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *SpecValueResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*SpecValue)(nil), "geiqin.srv.pdm.SpecValue")
	proto.RegisterType((*SpecValueResponse)(nil), "geiqin.srv.pdm.SpecValueResponse")
}

func init() { proto.RegisterFile("specValue.proto", fileDescriptor_d3cfe65595187a36) }

var fileDescriptor_d3cfe65595187a36 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x6f, 0x92, 0x26, 0xa5, 0x27, 0xa5, 0xf7, 0xde, 0x43, 0xc5, 0x58, 0x50, 0x62, 0x57,
	0x01, 0x21, 0x62, 0x7c, 0x04, 0xad, 0x1a, 0x70, 0x21, 0x29, 0xba, 0x95, 0xd8, 0x9c, 0x86, 0x81,
	0x24, 0x13, 0x27, 0x69, 0xc1, 0x07, 0xf3, 0x11, 0x7c, 0x2a, 0x37, 0x32, 0x93, 0x36, 0x62, 0xb1,
	0x6e, 0xea, 0xf2, 0xe4, 0xff, 0xcf, 0x97, 0xff, 0xfc, 0x30, 0xf0, 0xb7, 0x2a, 0x69, 0xf6, 0x10,
	0x67, 0x0b, 0xf2, 0x4b, 0xc1, 0x6b, 0x8e, 0x83, 0x94, 0xd8, 0x33, 0x2b, 0xfc, 0x4a, 0x2c, 0xfd,
	0x32, 0xc9, 0x47, 0xfd, 0x19, 0xcf, 0x73, 0x5e, 0x34, 0xea, 0xf8, 0x4d, 0x83, 0xde, 0x74, 0xbd,
	0x81, 0x03, 0xd0, 0x59, 0xe2, 0x68, 0xae, 0xe6, 0x99, 0x91, 0xce, 0x12, 0xdc, 0x87, 0xae, 0xc4,
	0x3d, 0xb2, 0xc4, 0xd1, 0xd5, 0x47, 0x4b, 0x8e, 0x61, 0x82, 0x43, 0x30, 0xe3, 0x8c, 0xc5, 0x95,
	0x63, 0xb8, 0x9a, 0xd7, 0x8b, 0x9a, 0x01, 0x0f, 0x01, 0x94, 0x7d, 0x29, 0x61, 0x4e, 0x47, 0x49,
	0xbd, 0x36, 0x0f, 0x1e, 0x81, 0xdd, 0xd0, 0xf2, 0x54, 0x12, 0x4d, 0x57, 0xf3, 0x8c, 0x46, 0x0f,
	0xf3, 0x34, 0x4c, 0xd0, 0x85, 0x7e, 0xab, 0x2f, 0x44, 0xe6, 0x58, 0x0a, 0x00, 0x2b, 0xc3, 0xbd,
	0xc8, 0xd0, 0x81, 0x6e, 0xc5, 0x45, 0xcd, 0x8a, 0xd4, 0xe9, 0xaa, 0x3c, 0xeb, 0x71, 0xfc, 0xae,
	0xc1, 0xff, 0xf6, 0x8e, 0x88, 0xaa, 0x92, 0x17, 0x15, 0xe1, 0x19, 0x58, 0x54, 0xd4, 0xac, 0x7e,
	0x51, 0x37, 0xd9, 0xc1, 0x81, 0xff, 0xb5, 0x0c, 0xff, 0x73, 0x65, 0x65, 0xc4, 0x13, 0x30, 0xcb,
	0x38, 0x25, 0xa1, 0x0e, 0xb6, 0x83, 0xbd, 0xcd, 0x8d, 0x3b, 0x29, 0x46, 0x8d, 0x07, 0x4f, 0xc1,
	0x64, 0x35, 0xe5, 0xb2, 0x06, 0xe3, 0x67, 0x7c, 0xe3, 0x93, 0x74, 0x12, 0x82, 0x0b, 0x55, 0xce,
	0x37, 0xf4, 0x89, 0x14, 0xa3, 0xc6, 0x83, 0x1e, 0x74, 0x58, 0x31, 0xe7, 0xaa, 0x28, 0x3b, 0x18,
	0x6e, 0x7a, 0xc3, 0x62, 0xce, 0x23, 0xe5, 0x08, 0x5e, 0x75, 0xf8, 0xd7, 0xfe, 0x6b, 0x4a, 0x62,
	0xc9, 0x66, 0x84, 0x37, 0x60, 0x5d, 0x08, 0x8a, 0x6b, 0xc2, 0xed, 0xb9, 0x46, 0xc7, 0xdb, 0x23,
	0xaf, 0x4a, 0x1c, 0xff, 0x91, 0xa4, 0x4b, 0xca, 0xe8, 0x17, 0x48, 0x13, 0x30, 0xae, 0xa9, 0xde,
	0x19, 0x73, 0x05, 0x9d, 0x5b, 0x56, 0xed, 0xcc, 0x79, 0xb2, 0xd4, 0x23, 0x38, 0xff, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xb9, 0x3a, 0x66, 0x77, 0x35, 0x03, 0x00, 0x00,
}
