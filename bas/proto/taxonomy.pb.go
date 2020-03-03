// Code generated by protoc-gen-go. DO NOT EDIT.
// source: taxonomy.proto

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

//标准分类信息
type Taxonomy struct {
	Id                   int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string      `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	TaxonomyId           int64       `protobuf:"varint,4,opt,name=taxonomy_id,json=taxonomyId,proto3" json:"taxonomy_id,omitempty"`
	Depth                int32       `protobuf:"varint,5,opt,name=depth,proto3" json:"depth,omitempty"`
	Path                 string      `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Memo                 string      `protobuf:"bytes,7,opt,name=memo,proto3" json:"memo,omitempty"`
	Sorting              int32       `protobuf:"varint,8,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Children             []*Taxonomy `protobuf:"bytes,9,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Taxonomy) Reset()         { *m = Taxonomy{} }
func (m *Taxonomy) String() string { return proto.CompactTextString(m) }
func (*Taxonomy) ProtoMessage()    {}
func (*Taxonomy) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b3e90e04276df48, []int{0}
}

func (m *Taxonomy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Taxonomy.Unmarshal(m, b)
}
func (m *Taxonomy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Taxonomy.Marshal(b, m, deterministic)
}
func (m *Taxonomy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Taxonomy.Merge(m, src)
}
func (m *Taxonomy) XXX_Size() int {
	return xxx_messageInfo_Taxonomy.Size(m)
}
func (m *Taxonomy) XXX_DiscardUnknown() {
	xxx_messageInfo_Taxonomy.DiscardUnknown(m)
}

var xxx_messageInfo_Taxonomy proto.InternalMessageInfo

func (m *Taxonomy) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Taxonomy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Taxonomy) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Taxonomy) GetTaxonomyId() int64 {
	if m != nil {
		return m.TaxonomyId
	}
	return 0
}

func (m *Taxonomy) GetDepth() int32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *Taxonomy) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Taxonomy) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Taxonomy) GetSorting() int32 {
	if m != nil {
		return m.Sorting
	}
	return 0
}

func (m *Taxonomy) GetChildren() []*Taxonomy {
	if m != nil {
		return m.Children
	}
	return nil
}

type TaxonomyResponse struct {
	Entity               *Taxonomy   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Taxonomy `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TaxonomyResponse) Reset()         { *m = TaxonomyResponse{} }
func (m *TaxonomyResponse) String() string { return proto.CompactTextString(m) }
func (*TaxonomyResponse) ProtoMessage()    {}
func (*TaxonomyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b3e90e04276df48, []int{1}
}

func (m *TaxonomyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaxonomyResponse.Unmarshal(m, b)
}
func (m *TaxonomyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaxonomyResponse.Marshal(b, m, deterministic)
}
func (m *TaxonomyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaxonomyResponse.Merge(m, src)
}
func (m *TaxonomyResponse) XXX_Size() int {
	return xxx_messageInfo_TaxonomyResponse.Size(m)
}
func (m *TaxonomyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaxonomyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaxonomyResponse proto.InternalMessageInfo

func (m *TaxonomyResponse) GetEntity() *Taxonomy {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *TaxonomyResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *TaxonomyResponse) GetItems() []*Taxonomy {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *TaxonomyResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *TaxonomyResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Taxonomy)(nil), "geiqin.srv.bas.Taxonomy")
	proto.RegisterType((*TaxonomyResponse)(nil), "geiqin.srv.bas.TaxonomyResponse")
}

func init() { proto.RegisterFile("taxonomy.proto", fileDescriptor_1b3e90e04276df48) }

var fileDescriptor_1b3e90e04276df48 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xc7, 0xcd, 0xe7, 0xed, 0x3d, 0x91, 0x2a, 0x87, 0x2b, 0x8c, 0x77, 0x63, 0xe8, 0x2a, 0x20,
	0x04, 0x89, 0x3e, 0x81, 0x5e, 0xb9, 0x5c, 0x70, 0x21, 0x69, 0xc1, 0xa5, 0xa4, 0xc9, 0x69, 0x33,
	0x60, 0x66, 0xe2, 0xcc, 0x50, 0xec, 0x33, 0xf9, 0x74, 0xee, 0x5c, 0xca, 0xcc, 0x34, 0x05, 0x8b,
	0x55, 0xa1, 0xbb, 0x33, 0xff, 0xf3, 0x3b, 0xff, 0xf3, 0x01, 0x03, 0x73, 0xd3, 0x7c, 0x93, 0x42,
	0x0e, 0xfb, 0x72, 0x54, 0xd2, 0x48, 0x9c, 0x6f, 0x89, 0x7f, 0xe5, 0xa2, 0xd4, 0x6a, 0x57, 0xae,
	0x1b, 0x7d, 0xfb, 0xb8, 0x95, 0xc3, 0x20, 0x85, 0xcf, 0x2e, 0x7e, 0x04, 0x30, 0x5b, 0x1d, 0x0a,
	0x70, 0x0e, 0x21, 0xef, 0x58, 0x90, 0x07, 0x45, 0x54, 0x87, 0xbc, 0x43, 0x84, 0x58, 0x34, 0x03,
	0xb1, 0x30, 0x0f, 0x8a, 0xeb, 0xda, 0xc5, 0x56, 0x33, 0xfb, 0x91, 0x58, 0xe4, 0x35, 0x1b, 0xe3,
	0x0b, 0xc8, 0xa6, 0xa6, 0x9f, 0x79, 0xc7, 0x62, 0x67, 0x00, 0x93, 0xf4, 0xd0, 0xe1, 0x0d, 0x24,
	0x1d, 0x8d, 0xa6, 0x67, 0x49, 0x1e, 0x14, 0x49, 0xed, 0x1f, 0xd6, 0x6a, 0x6c, 0x4c, 0xcf, 0x52,
	0x6f, 0x65, 0x63, 0xab, 0x0d, 0x34, 0x48, 0x76, 0xe5, 0x35, 0x1b, 0x23, 0x83, 0x2b, 0x2d, 0x95,
	0xe1, 0x62, 0xcb, 0x66, 0xae, 0x7e, 0x7a, 0xe2, 0x1b, 0x98, 0xb5, 0x3d, 0xff, 0xd2, 0x29, 0x12,
	0xec, 0x3a, 0x8f, 0x8a, 0xac, 0x62, 0xe5, 0xef, 0xeb, 0x96, 0xd3, 0x72, 0xf5, 0x91, 0x5c, 0xfc,
	0x0c, 0xe0, 0xe9, 0x51, 0x26, 0x3d, 0x4a, 0xa1, 0x09, 0x5f, 0x41, 0x4a, 0xc2, 0x70, 0xb3, 0x77,
	0xfb, 0xff, 0xcd, 0xe8, 0xc0, 0xe1, 0x4b, 0x48, 0xc6, 0x66, 0x4b, 0xca, 0x9d, 0x27, 0xab, 0x9e,
	0x9d, 0x16, 0x7c, 0xb4, 0xc9, 0xda, 0x33, 0x58, 0x42, 0xc2, 0x0d, 0x0d, 0x9a, 0x45, 0xff, 0x18,
	0xd3, 0x63, 0xd6, 0x9c, 0x94, 0x92, 0xca, 0x1d, 0xf3, 0x0f, 0xe6, 0xef, 0x6d, 0xb2, 0xf6, 0x0c,
	0x16, 0x10, 0x73, 0xb1, 0x91, 0xee, 0xba, 0x59, 0x75, 0x73, 0xca, 0x3e, 0x88, 0x8d, 0xac, 0x1d,
	0x51, 0x7d, 0x0f, 0xe1, 0xc9, 0xd4, 0x6a, 0x49, 0x6a, 0xc7, 0x5b, 0xc2, 0x77, 0x10, 0xdd, 0x93,
	0xc1, 0xb3, 0x23, 0xdd, 0xe6, 0x67, 0x87, 0x3d, 0x1c, 0x6f, 0xf1, 0x08, 0xef, 0x21, 0x5d, 0x52,
	0xa3, 0xda, 0x1e, 0x9f, 0x9f, 0xd2, 0x6f, 0x1b, 0x4d, 0x9f, 0x7a, 0x52, 0xf4, 0x5f, 0x46, 0x77,
	0x10, 0xaf, 0x14, 0xd1, 0x85, 0xe3, 0xdc, 0x41, 0xfc, 0x81, 0xeb, 0x0b, 0x97, 0x5a, 0xa7, 0xee,
	0x8f, 0xbc, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x26, 0x1d, 0x4f, 0x00, 0x53, 0x03, 0x00, 0x00,
}
