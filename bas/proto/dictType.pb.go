// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dictType.proto

package geiqin_srv_bas

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

//字段类型
type DictType struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Editable             bool     `protobuf:"varint,4,opt,name=editable,proto3" json:"editable,omitempty"`
	Dicts                []*Dict  `protobuf:"bytes,5,rep,name=dicts,proto3" json:"dicts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DictType) Reset()         { *m = DictType{} }
func (m *DictType) String() string { return proto.CompactTextString(m) }
func (*DictType) ProtoMessage()    {}
func (*DictType) Descriptor() ([]byte, []int) {
	return fileDescriptor_25937c3ebe87c716, []int{0}
}

func (m *DictType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DictType.Unmarshal(m, b)
}
func (m *DictType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DictType.Marshal(b, m, deterministic)
}
func (m *DictType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DictType.Merge(m, src)
}
func (m *DictType) XXX_Size() int {
	return xxx_messageInfo_DictType.Size(m)
}
func (m *DictType) XXX_DiscardUnknown() {
	xxx_messageInfo_DictType.DiscardUnknown(m)
}

var xxx_messageInfo_DictType proto.InternalMessageInfo

func (m *DictType) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DictType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DictType) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DictType) GetEditable() bool {
	if m != nil {
		return m.Editable
	}
	return false
}

func (m *DictType) GetDicts() []*Dict {
	if m != nil {
		return m.Dicts
	}
	return nil
}

type DictTypeRequest struct {
	Types                []string `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DictTypeRequest) Reset()         { *m = DictTypeRequest{} }
func (m *DictTypeRequest) String() string { return proto.CompactTextString(m) }
func (*DictTypeRequest) ProtoMessage()    {}
func (*DictTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25937c3ebe87c716, []int{1}
}

func (m *DictTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DictTypeRequest.Unmarshal(m, b)
}
func (m *DictTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DictTypeRequest.Marshal(b, m, deterministic)
}
func (m *DictTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DictTypeRequest.Merge(m, src)
}
func (m *DictTypeRequest) XXX_Size() int {
	return xxx_messageInfo_DictTypeRequest.Size(m)
}
func (m *DictTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DictTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DictTypeRequest proto.InternalMessageInfo

func (m *DictTypeRequest) GetTypes() []string {
	if m != nil {
		return m.Types
	}
	return nil
}

type DictTypeResponse struct {
	Entity               *DictType   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*DictType `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DictTypeResponse) Reset()         { *m = DictTypeResponse{} }
func (m *DictTypeResponse) String() string { return proto.CompactTextString(m) }
func (*DictTypeResponse) ProtoMessage()    {}
func (*DictTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_25937c3ebe87c716, []int{2}
}

func (m *DictTypeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DictTypeResponse.Unmarshal(m, b)
}
func (m *DictTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DictTypeResponse.Marshal(b, m, deterministic)
}
func (m *DictTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DictTypeResponse.Merge(m, src)
}
func (m *DictTypeResponse) XXX_Size() int {
	return xxx_messageInfo_DictTypeResponse.Size(m)
}
func (m *DictTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DictTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DictTypeResponse proto.InternalMessageInfo

func (m *DictTypeResponse) GetEntity() *DictType {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *DictTypeResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *DictTypeResponse) GetItems() []*DictType {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *DictTypeResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DictTypeResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*DictType)(nil), "geiqin.srv.bas.DictType")
	proto.RegisterType((*DictTypeRequest)(nil), "geiqin.srv.bas.DictTypeRequest")
	proto.RegisterType((*DictTypeResponse)(nil), "geiqin.srv.bas.DictTypeResponse")
}

func init() { proto.RegisterFile("dictType.proto", fileDescriptor_25937c3ebe87c716) }

var fileDescriptor_25937c3ebe87c716 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0xbb, 0x96, 0x64, 0xdc, 0x71, 0x71, 0xcb, 0xe0, 0xc2, 0x56, 0x97, 0x0a, 0x5d, 0x2a,
	0x5a, 0x10, 0x45, 0x7d, 0x83, 0xfe, 0xc1, 0x94, 0xf6, 0x10, 0xd6, 0x81, 0x9c, 0x65, 0x69, 0x6c,
	0x2f, 0x58, 0x5a, 0x79, 0x77, 0x63, 0xf0, 0x3b, 0xe4, 0x51, 0xf3, 0x00, 0x39, 0x06, 0xad, 0x9c,
	0x98, 0x38, 0x16, 0xf8, 0x36, 0xa3, 0xf9, 0xcd, 0xf7, 0xcd, 0x8c, 0x16, 0x26, 0xa5, 0x2c, 0xec,
	0xf5, 0xbe, 0xa1, 0xb4, 0xd1, 0xca, 0x2a, 0x9c, 0xac, 0x48, 0x6e, 0x65, 0x9d, 0x1a, 0xbd, 0x4b,
	0x17, 0xb9, 0x09, 0xdf, 0x15, 0xaa, 0xaa, 0x54, 0xdd, 0x55, 0x43, 0x68, 0xe9, 0x2e, 0x8e, 0xef,
	0x18, 0x8c, 0x7e, 0x1f, 0x9a, 0x71, 0x02, 0x03, 0x59, 0x72, 0x16, 0xb1, 0x24, 0x10, 0x03, 0x59,
	0x22, 0x82, 0x5f, 0xe7, 0x15, 0xf1, 0x41, 0xc4, 0x92, 0xb7, 0xc2, 0xc5, 0x38, 0x85, 0xc0, 0x4a,
	0xbb, 0x21, 0xee, 0xb9, 0x8f, 0x5d, 0x82, 0x21, 0x8c, 0xa8, 0x94, 0x36, 0x5f, 0x6c, 0x88, 0xfb,
	0x11, 0x4b, 0x46, 0xe2, 0x39, 0xc7, 0xaf, 0x10, 0xb4, 0x86, 0x86, 0x07, 0x91, 0x97, 0x8c, 0xb3,
	0x69, 0xfa, 0x72, 0xb8, 0xb4, 0xb5, 0x17, 0x1d, 0x12, 0x7f, 0x81, 0xf7, 0x4f, 0xd3, 0x08, 0xda,
	0xde, 0x92, 0xb1, 0xce, 0x70, 0xdf, 0x90, 0xe1, 0x2c, 0xf2, 0x9c, 0x61, 0x9b, 0xc4, 0x0f, 0x0c,
	0x3e, 0x1c, 0x49, 0xd3, 0xa8, 0xda, 0x10, 0x7e, 0x87, 0x21, 0xd5, 0x56, 0xda, 0xbd, 0xdb, 0x61,
	0x9c, 0xf1, 0x73, 0x56, 0xae, 0xe3, 0xc0, 0xe1, 0x37, 0x08, 0x9a, 0x7c, 0x45, 0xda, 0xad, 0x38,
	0xce, 0x3e, 0x9e, 0x36, 0x5c, 0xb5, 0x45, 0xd1, 0x31, 0x98, 0x42, 0x20, 0x2d, 0x55, 0x86, 0x7b,
	0x6e, 0x91, 0x7e, 0xf5, 0x0e, 0x6b, 0xc5, 0x49, 0x6b, 0xa5, 0xdd, 0x45, 0xce, 0x88, 0xff, 0x69,
	0x8b, 0xa2, 0x63, 0x30, 0x01, 0x5f, 0xd6, 0x4b, 0xc5, 0x03, 0xc7, 0xbe, 0x3a, 0xd2, 0xdf, 0x7a,
	0xa9, 0x84, 0x23, 0xb2, 0x7b, 0x76, 0x3c, 0xd2, 0x9c, 0xf4, 0x4e, 0x16, 0x84, 0xbf, 0xc0, 0x9b,
	0x91, 0xc5, 0xde, 0x91, 0xc2, 0xa8, 0x77, 0xd8, 0xc3, 0xf1, 0xe2, 0x37, 0x38, 0x83, 0xe1, 0x9c,
	0x72, 0x5d, 0xac, 0xf1, 0xd3, 0x29, 0xfd, 0x33, 0x37, 0x74, 0xb3, 0x26, 0x7d, 0x99, 0xd0, 0x3f,
	0xf0, 0xff, 0x4b, 0x63, 0xf1, 0x73, 0x3f, 0xeb, 0xfe, 0xed, 0x25, 0x62, 0x8b, 0xa1, 0x7b, 0xa8,
	0x3f, 0x1e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x26, 0x64, 0xf7, 0x4f, 0xe4, 0x02, 0x00, 0x00,
}