// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spec.proto

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

type Spec struct {
	Id                   int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	SpecType             string       `protobuf:"bytes,3,opt,name=spec_type,json=specType,proto3" json:"spec_type"`
	ShowStyle            string       `protobuf:"bytes,4,opt,name=show_style,json=showStyle,proto3" json:"show_style"`
	Memo                 string       `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo"`
	Sorting              int32        `protobuf:"varint,6,opt,name=sorting,proto3" json:"sorting"`
	CreatedAt            string       `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt            string       `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Values               []*SpecValue `protobuf:"bytes,9,rep,name=values,proto3" json:"values"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
func (*Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{0}
}

func (m *Spec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spec.Unmarshal(m, b)
}
func (m *Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spec.Marshal(b, m, deterministic)
}
func (m *Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spec.Merge(m, src)
}
func (m *Spec) XXX_Size() int {
	return xxx_messageInfo_Spec.Size(m)
}
func (m *Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_Spec proto.InternalMessageInfo

func (m *Spec) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Spec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Spec) GetSpecType() string {
	if m != nil {
		return m.SpecType
	}
	return ""
}

func (m *Spec) GetShowStyle() string {
	if m != nil {
		return m.ShowStyle
	}
	return ""
}

func (m *Spec) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Spec) GetSorting() int32 {
	if m != nil {
		return m.Sorting
	}
	return 0
}

func (m *Spec) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Spec) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Spec) GetValues() []*SpecValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type SpecResponse struct {
	Entity               *Spec    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items                []*Spec  `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecResponse) Reset()         { *m = SpecResponse{} }
func (m *SpecResponse) String() string { return proto.CompactTextString(m) }
func (*SpecResponse) ProtoMessage()    {}
func (*SpecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{1}
}

func (m *SpecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecResponse.Unmarshal(m, b)
}
func (m *SpecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecResponse.Marshal(b, m, deterministic)
}
func (m *SpecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecResponse.Merge(m, src)
}
func (m *SpecResponse) XXX_Size() int {
	return xxx_messageInfo_SpecResponse.Size(m)
}
func (m *SpecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpecResponse proto.InternalMessageInfo

func (m *SpecResponse) GetEntity() *Spec {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *SpecResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *SpecResponse) GetItems() []*Spec {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SpecResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *SpecResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Spec)(nil), "geiqin.srv.pdm.Spec")
	proto.RegisterType((*SpecResponse)(nil), "geiqin.srv.pdm.SpecResponse")
}

func init() { proto.RegisterFile("spec.proto", fileDescriptor_423806180556987f) }

var fileDescriptor_423806180556987f = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x5f, 0x8b, 0xd3, 0x40,
	0x14, 0xc5, 0x6d, 0x93, 0x66, 0xb7, 0xb7, 0xcb, 0x0a, 0x17, 0x85, 0xb1, 0x2a, 0x94, 0x3e, 0x15,
	0x95, 0x80, 0xf1, 0x51, 0x14, 0xd6, 0x55, 0x44, 0xf0, 0x41, 0x12, 0xff, 0x3c, 0x96, 0x98, 0xdc,
	0x6d, 0x07, 0x9a, 0xcc, 0x38, 0x33, 0x5b, 0xc9, 0x07, 0xf0, 0xd5, 0x4f, 0xea, 0x87, 0x90, 0x3b,
	0x93, 0x3e, 0xe8, 0x06, 0x16, 0xfa, 0x36, 0x39, 0xe7, 0xdc, 0x33, 0x37, 0x3f, 0x12, 0x00, 0xab,
	0xa9, 0x4a, 0xb5, 0x51, 0x4e, 0xe1, 0xf9, 0x86, 0xe4, 0x0f, 0xd9, 0xa6, 0xd6, 0xec, 0x53, 0x5d,
	0x37, 0xf3, 0xb3, 0x4a, 0x35, 0x8d, 0x6a, 0x83, 0x3b, 0xbf, 0xcb, 0xc9, 0xaf, 0xe5, 0xee, 0x9a,
	0x82, 0xb0, 0xfc, 0x35, 0x86, 0xb8, 0xd0, 0x54, 0xe1, 0x39, 0x8c, 0x65, 0x2d, 0x46, 0x8b, 0xd1,
	0x6a, 0x92, 0x8f, 0x65, 0x8d, 0x08, 0x71, 0x5b, 0x36, 0x24, 0xc6, 0x8b, 0xd1, 0x6a, 0x9a, 0xfb,
	0x33, 0x3e, 0x84, 0x29, 0xcf, 0xaf, 0x5d, 0xa7, 0x49, 0x44, 0xde, 0x38, 0x65, 0xe1, 0x73, 0xa7,
	0x09, 0x1f, 0x03, 0xd8, 0xad, 0xfa, 0xb9, 0xb6, 0xae, 0xdb, 0x91, 0x88, 0xbd, 0x3b, 0x65, 0xa5,
	0x60, 0x81, 0xfb, 0x1a, 0x6a, 0x94, 0x98, 0x84, 0x3e, 0x3e, 0xa3, 0x80, 0x13, 0xab, 0x8c, 0x93,
	0xed, 0x46, 0x24, 0xfe, 0xe2, 0xc3, 0x23, 0x97, 0x55, 0x86, 0x4a, 0x47, 0xf5, 0xba, 0x74, 0xe2,
	0x24, 0x94, 0xf5, 0xca, 0x85, 0x63, 0xfb, 0x5a, 0xd7, 0x07, 0xfb, 0x34, 0xd8, 0xbd, 0x72, 0xe1,
	0xf0, 0x39, 0x24, 0x7b, 0x7e, 0x47, 0x2b, 0xa6, 0x8b, 0x68, 0x35, 0xcb, 0x1e, 0xa4, 0xff, 0x42,
	0x49, 0x8b, 0x03, 0x85, 0xbc, 0x0f, 0x2e, 0xff, 0x8c, 0xe0, 0x8c, 0xd5, 0x9c, 0xac, 0x56, 0xad,
	0x25, 0x7c, 0x06, 0x09, 0xb5, 0x4e, 0xba, 0xce, 0x33, 0x99, 0x65, 0xf7, 0x86, 0x3a, 0xf2, 0x3e,
	0x83, 0x4f, 0x61, 0xa2, 0xcb, 0x0d, 0x19, 0x8f, 0x6b, 0x96, 0xdd, 0xff, 0x3f, 0xfc, 0x89, 0xcd,
	0x3c, 0x64, 0xf0, 0x09, 0x4c, 0xa4, 0xa3, 0xc6, 0x8a, 0xc8, 0x6f, 0x37, 0xdc, 0x1c, 0x22, 0x5c,
	0x4c, 0xc6, 0x28, 0xe3, 0x81, 0x0e, 0x14, 0xbf, 0x63, 0x33, 0x0f, 0x19, 0x5c, 0x41, 0x2c, 0xdb,
	0xab, 0xc0, 0x78, 0xa0, 0xf7, 0x43, 0x7b, 0xa5, 0x72, 0x9f, 0xc8, 0x7e, 0x47, 0x30, 0xe3, 0x6b,
	0x0a, 0x32, 0x7b, 0x59, 0x11, 0xbe, 0x86, 0xe4, 0xd2, 0xd3, 0xc5, 0xc1, 0x6d, 0xe6, 0x8f, 0x06,
	0x77, 0xec, 0x59, 0x2d, 0xef, 0xf0, 0xfc, 0x17, 0x8f, 0xff, 0xf8, 0xf9, 0xb7, 0xb4, 0xa3, 0xa3,
	0xe7, 0x5f, 0x42, 0xf4, 0x9e, 0xdc, 0x91, 0xc3, 0xaf, 0x20, 0xfe, 0x28, 0xad, 0xc3, 0x9b, 0x70,
	0x1b, 0xed, 0xba, 0x5b, 0xc7, 0x2f, 0x21, 0x29, 0xa8, 0x34, 0xd5, 0x16, 0x6f, 0x7c, 0x67, 0x6f,
	0x4a, 0x4b, 0xdf, 0xb6, 0x64, 0xe8, 0xb6, 0x92, 0xef, 0x89, 0xff, 0x1d, 0x5f, 0xfc, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xcd, 0x6c, 0xb1, 0x74, 0xcb, 0x03, 0x00, 0x00,
}
