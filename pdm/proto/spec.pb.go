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
	Id                   int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SpecType             string       `protobuf:"bytes,3,opt,name=spec_type,json=specType,proto3" json:"spec_type,omitempty"`
	ShowStyle            string       `protobuf:"bytes,4,opt,name=show_style,json=showStyle,proto3" json:"show_style,omitempty"`
	Memo                 string       `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
	Sorting              int32        `protobuf:"varint,6,opt,name=sorting,proto3" json:"sorting,omitempty"`
	CreatedAt            string       `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string       `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Values               []*SpecValue `protobuf:"bytes,9,rep,name=values,proto3" json:"values,omitempty"`
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
	Entity               *Spec    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Spec  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
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
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x6d, 0xda, 0x66, 0xb7, 0xa7, 0xcb, 0x0a, 0x07, 0x85, 0xb1, 0x2a, 0x94, 0x5e, 0x15,
	0x95, 0x80, 0xf5, 0x56, 0x94, 0x75, 0x15, 0xe9, 0x9d, 0x24, 0xfe, 0xb9, 0x2c, 0x31, 0x39, 0xdb,
	0x0e, 0x34, 0x33, 0xe3, 0xcc, 0x6c, 0x25, 0x0f, 0xe0, 0x03, 0xfa, 0x1e, 0x3e, 0x84, 0x9c, 0x99,
	0xf4, 0x42, 0x37, 0xb0, 0xb0, 0x77, 0xd3, 0xef, 0xfb, 0xce, 0x6f, 0xce, 0x7c, 0x34, 0x00, 0xce,
	0x50, 0x95, 0x19, 0xab, 0xbd, 0xc6, 0xf3, 0x2d, 0xc9, 0x1f, 0x52, 0x65, 0xce, 0x1e, 0x32, 0x53,
	0x37, 0xb3, 0xb3, 0x4a, 0x37, 0x8d, 0x56, 0xd1, 0x9d, 0xdd, 0xe7, 0xe4, 0xd7, 0x72, 0x7f, 0x4d,
	0x51, 0x58, 0xfc, 0x4a, 0x60, 0x54, 0x18, 0xaa, 0xf0, 0x1c, 0x12, 0x59, 0x8b, 0xc1, 0x7c, 0xb0,
	0x1c, 0xe7, 0x89, 0xac, 0x11, 0x61, 0xa4, 0xca, 0x86, 0x44, 0x32, 0x1f, 0x2c, 0x27, 0x79, 0x38,
	0xe3, 0x63, 0x98, 0xf0, 0xfc, 0xc6, 0xb7, 0x86, 0xc4, 0x30, 0x18, 0xa7, 0x2c, 0x7c, 0x6e, 0x0d,
	0xe1, 0x53, 0x00, 0xb7, 0xd3, 0x3f, 0x37, 0xce, 0xb7, 0x7b, 0x12, 0xa3, 0xe0, 0x4e, 0x58, 0x29,
	0x58, 0x60, 0x5e, 0x43, 0x8d, 0x16, 0xe3, 0xc8, 0xe3, 0x33, 0x0a, 0x38, 0x71, 0xda, 0x7a, 0xa9,
	0xb6, 0x22, 0x0d, 0x17, 0x1f, 0x7f, 0x32, 0xac, 0xb2, 0x54, 0x7a, 0xaa, 0x37, 0xa5, 0x17, 0x27,
	0x11, 0xd6, 0x29, 0x17, 0x9e, 0xed, 0x6b, 0x53, 0x1f, 0xed, 0xd3, 0x68, 0x77, 0xca, 0x85, 0xc7,
	0x97, 0x90, 0x1e, 0xf8, 0x8d, 0x4e, 0x4c, 0xe6, 0xc3, 0xe5, 0x74, 0xf5, 0x28, 0xfb, 0xb7, 0x94,
	0xac, 0x38, 0xb6, 0x90, 0x77, 0xc1, 0xc5, 0x9f, 0x01, 0x9c, 0xb1, 0x9a, 0x93, 0x33, 0x5a, 0x39,
	0xc2, 0x17, 0x90, 0x92, 0xf2, 0xd2, 0xb7, 0xa1, 0x93, 0xe9, 0xea, 0x41, 0x1f, 0x23, 0xef, 0x32,
	0xf8, 0x1c, 0xc6, 0xa6, 0xdc, 0x92, 0x0d, 0x75, 0x4d, 0x57, 0x0f, 0xff, 0x0f, 0x7f, 0x62, 0x33,
	0x8f, 0x19, 0x7c, 0x06, 0x63, 0xe9, 0xa9, 0x71, 0x62, 0x18, 0xb6, 0xeb, 0x27, 0xc7, 0x08, 0x83,
	0xc9, 0x5a, 0x6d, 0x43, 0xa1, 0x3d, 0xe0, 0x0f, 0x6c, 0xe6, 0x31, 0x83, 0x4b, 0x18, 0x49, 0x75,
	0x15, 0x3b, 0xee, 0xe1, 0xae, 0xd5, 0x95, 0xce, 0x43, 0x62, 0xf5, 0x3b, 0x81, 0x29, 0x5f, 0x53,
	0x90, 0x3d, 0xc8, 0x8a, 0xf0, 0x0d, 0xa4, 0x97, 0xa1, 0x5d, 0xec, 0xdd, 0x66, 0xf6, 0xa4, 0x77,
	0xc7, 0xae, 0xab, 0xc5, 0x3d, 0x9e, 0xff, 0x12, 0xea, 0xbf, 0xe3, 0xfc, 0x5b, 0x48, 0xdf, 0xd3,
	0x9e, 0x3c, 0xe1, 0x8d, 0x17, 0xae, 0xeb, 0xb5, 0xf2, 0xb7, 0x02, 0x5e, 0xc3, 0xf0, 0x23, 0xf9,
	0xbb, 0x4e, 0x5f, 0x42, 0x5a, 0x50, 0x69, 0xab, 0x1d, 0xde, 0xf8, 0xab, 0xbc, 0x2b, 0x1d, 0x7d,
	0xdb, 0x91, 0xa5, 0xdb, 0x20, 0xdf, 0xd3, 0xf0, 0x45, 0xbd, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x4c, 0x4a, 0x45, 0x9b, 0x8e, 0x03, 0x00, 0x00,
}