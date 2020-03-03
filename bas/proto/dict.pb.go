// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dict.proto

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

// 字典信息
type Dict struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DictTypeId           int32    `protobuf:"varint,2,opt,name=dict_type_id,json=dictTypeId,proto3" json:"dict_type_id,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Text                 string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Term                 string   `protobuf:"bytes,5,opt,name=term,proto3" json:"term,omitempty"`
	Memo                 string   `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
	Sorting              int32    `protobuf:"varint,7,opt,name=sorting,proto3" json:"sorting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dict) Reset()         { *m = Dict{} }
func (m *Dict) String() string { return proto.CompactTextString(m) }
func (*Dict) ProtoMessage()    {}
func (*Dict) Descriptor() ([]byte, []int) {
	return fileDescriptor_67812e90854f6714, []int{0}
}

func (m *Dict) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dict.Unmarshal(m, b)
}
func (m *Dict) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dict.Marshal(b, m, deterministic)
}
func (m *Dict) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dict.Merge(m, src)
}
func (m *Dict) XXX_Size() int {
	return xxx_messageInfo_Dict.Size(m)
}
func (m *Dict) XXX_DiscardUnknown() {
	xxx_messageInfo_Dict.DiscardUnknown(m)
}

var xxx_messageInfo_Dict proto.InternalMessageInfo

func (m *Dict) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Dict) GetDictTypeId() int32 {
	if m != nil {
		return m.DictTypeId
	}
	return 0
}

func (m *Dict) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Dict) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Dict) GetTerm() string {
	if m != nil {
		return m.Term
	}
	return ""
}

func (m *Dict) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Dict) GetSorting() int32 {
	if m != nil {
		return m.Sorting
	}
	return 0
}

//
type DictResponse struct {
	Entity               *Dict            `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager           `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Dict          `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error           `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info            `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	Maps                 map[string]*Dict `protobuf:"bytes,6,rep,name=maps,proto3" json:"maps,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DictResponse) Reset()         { *m = DictResponse{} }
func (m *DictResponse) String() string { return proto.CompactTextString(m) }
func (*DictResponse) ProtoMessage()    {}
func (*DictResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67812e90854f6714, []int{1}
}

func (m *DictResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DictResponse.Unmarshal(m, b)
}
func (m *DictResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DictResponse.Marshal(b, m, deterministic)
}
func (m *DictResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DictResponse.Merge(m, src)
}
func (m *DictResponse) XXX_Size() int {
	return xxx_messageInfo_DictResponse.Size(m)
}
func (m *DictResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DictResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DictResponse proto.InternalMessageInfo

func (m *DictResponse) GetEntity() *Dict {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *DictResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *DictResponse) GetItems() []*Dict {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *DictResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DictResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *DictResponse) GetMaps() map[string]*Dict {
	if m != nil {
		return m.Maps
	}
	return nil
}

func init() {
	proto.RegisterType((*Dict)(nil), "geiqin.srv.bas.Dict")
	proto.RegisterType((*DictResponse)(nil), "geiqin.srv.bas.DictResponse")
	proto.RegisterMapType((map[string]*Dict)(nil), "geiqin.srv.bas.DictResponse.MapsEntry")
}

func init() { proto.RegisterFile("dict.proto", fileDescriptor_67812e90854f6714) }

var fileDescriptor_67812e90854f6714 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdf, 0x6a, 0x14, 0x31,
	0x14, 0xc6, 0x9d, 0xbf, 0x65, 0xcf, 0x2e, 0x45, 0x0e, 0x15, 0xe2, 0xe2, 0xc5, 0xb0, 0x17, 0xb2,
	0xa8, 0xcc, 0xc5, 0x7a, 0x23, 0xb5, 0x57, 0x6a, 0x91, 0x82, 0x05, 0x49, 0x0b, 0x5e, 0x96, 0xe9,
	0xcc, 0xe9, 0x36, 0x68, 0x92, 0x31, 0x89, 0x8b, 0xf3, 0x0c, 0xbe, 0x86, 0xef, 0xe8, 0xad, 0x24,
	0x69, 0x05, 0x97, 0x71, 0x85, 0xbd, 0x3b, 0xf9, 0xce, 0x2f, 0xdf, 0x7c, 0xe7, 0x0c, 0x01, 0xe8,
	0x44, 0xeb, 0xea, 0xde, 0x68, 0xa7, 0xf1, 0x70, 0x4d, 0xe2, 0xab, 0x50, 0xb5, 0x35, 0x9b, 0xfa,
	0xba, 0xb1, 0xf3, 0x59, 0xab, 0xa5, 0xd4, 0x2a, 0x76, 0x17, 0x3f, 0x13, 0xc8, 0xdf, 0x89, 0xd6,
	0xe1, 0x21, 0xa4, 0xa2, 0x63, 0x49, 0x95, 0x2c, 0x0b, 0x9e, 0x8a, 0x0e, 0x2b, 0x98, 0x79, 0x93,
	0x2b, 0x37, 0xf4, 0x74, 0x25, 0x3a, 0x96, 0x86, 0x4e, 0x30, 0xbe, 0x1c, 0x7a, 0x3a, 0xeb, 0xf0,
	0x08, 0x8a, 0x4d, 0xf3, 0xe5, 0x1b, 0xb1, 0xac, 0x4a, 0x96, 0x13, 0x1e, 0x0f, 0x88, 0x90, 0x3b,
	0xfa, 0xee, 0x58, 0x1e, 0xc4, 0x50, 0x47, 0xcd, 0x48, 0x56, 0xdc, 0x6b, 0x46, 0x7a, 0x4d, 0x92,
	0xd4, 0xac, 0x8c, 0x9a, 0xaf, 0x91, 0xc1, 0x81, 0xd5, 0xc6, 0x09, 0xb5, 0x66, 0x07, 0xe1, 0x73,
	0xf7, 0xc7, 0xc5, 0xaf, 0x14, 0x66, 0x3e, 0x26, 0x27, 0xdb, 0x6b, 0x65, 0x09, 0x5f, 0x40, 0x49,
	0xca, 0x09, 0x37, 0x84, 0xc8, 0xd3, 0xd5, 0x51, 0xfd, 0xf7, 0x98, 0x75, 0xa0, 0xef, 0x18, 0x7c,
	0x0e, 0x45, 0xdf, 0xac, 0xc9, 0x84, 0x29, 0xa6, 0xab, 0x47, 0xdb, 0xf0, 0x47, 0xdf, 0xe4, 0x91,
	0xc1, 0x67, 0x50, 0x08, 0x47, 0xd2, 0xb2, 0xac, 0xca, 0xfe, 0xe9, 0x1c, 0x11, 0x6f, 0x4c, 0xc6,
	0x68, 0x13, 0xc6, 0x1d, 0x31, 0x3e, 0xf5, 0x4d, 0x1e, 0x19, 0x5c, 0x42, 0x2e, 0xd4, 0x8d, 0x0e,
	0x6b, 0x18, 0xf1, 0x3d, 0x53, 0x37, 0x9a, 0x07, 0x02, 0x8f, 0x21, 0x97, 0x4d, 0x6f, 0x59, 0x19,
	0x12, 0x3c, 0x1d, 0x4d, 0x70, 0xb7, 0x89, 0xfa, 0xbc, 0xe9, 0xed, 0xa9, 0x72, 0x66, 0xe0, 0xe1,
	0xce, 0xfc, 0x1c, 0x26, 0x7f, 0x24, 0x7c, 0x08, 0xd9, 0x67, 0x8a, 0x3b, 0x9a, 0x70, 0x5f, 0xfa,
	0xe9, 0xe2, 0x5f, 0x4b, 0x77, 0xec, 0x2d, 0x22, 0xc7, 0xe9, 0xab, 0x64, 0xf5, 0x23, 0x85, 0xa9,
	0xd7, 0x2e, 0xc8, 0x6c, 0x44, 0x4b, 0xf8, 0x1a, 0xb2, 0xf7, 0xe4, 0x70, 0xf4, 0xde, 0xfc, 0xc9,
	0xae, 0xa4, 0x8b, 0x07, 0xf8, 0x16, 0xca, 0x0b, 0x6a, 0x4c, 0x7b, 0x8b, 0x8f, 0xb7, 0xc9, 0x37,
	0x8d, 0xa5, 0x4f, 0xb7, 0x64, 0xe8, 0xbf, 0x26, 0x27, 0x90, 0x5f, 0x1a, 0xa2, 0x3d, 0x23, 0x9c,
	0x40, 0xfe, 0x41, 0xd8, 0x3d, 0x07, 0xb8, 0x2e, 0xc3, 0xab, 0x79, 0xf9, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xa5, 0x2c, 0xaa, 0x8d, 0x61, 0x03, 0x00, 0x00,
}
