// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cat.proto

package geiqin_srv_cms_media

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
type Cat struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Taxonomy             string   `protobuf:"bytes,4,opt,name=taxonomy,proto3" json:"taxonomy,omitempty"`
	TotalNum             int32    `protobuf:"varint,5,opt,name=total_num,json=totalNum,proto3" json:"total_num,omitempty"`
	TotalSize            float32  `protobuf:"fixed32,6,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	CatId                int32    `protobuf:"varint,7,opt,name=cat_id,json=catId,proto3" json:"cat_id,omitempty"`
	UserId               int64    `protobuf:"varint,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Children             []*Cat   `protobuf:"bytes,9,rep,name=children,proto3" json:"children,omitempty"`
	Ids                  []int32  `protobuf:"varint,10,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cat) Reset()         { *m = Cat{} }
func (m *Cat) String() string { return proto.CompactTextString(m) }
func (*Cat) ProtoMessage()    {}
func (*Cat) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0a5ac8640cab35d, []int{0}
}

func (m *Cat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cat.Unmarshal(m, b)
}
func (m *Cat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cat.Marshal(b, m, deterministic)
}
func (m *Cat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cat.Merge(m, src)
}
func (m *Cat) XXX_Size() int {
	return xxx_messageInfo_Cat.Size(m)
}
func (m *Cat) XXX_DiscardUnknown() {
	xxx_messageInfo_Cat.DiscardUnknown(m)
}

var xxx_messageInfo_Cat proto.InternalMessageInfo

func (m *Cat) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Cat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cat) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Cat) GetTaxonomy() string {
	if m != nil {
		return m.Taxonomy
	}
	return ""
}

func (m *Cat) GetTotalNum() int32 {
	if m != nil {
		return m.TotalNum
	}
	return 0
}

func (m *Cat) GetTotalSize() float32 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *Cat) GetCatId() int32 {
	if m != nil {
		return m.CatId
	}
	return 0
}

func (m *Cat) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Cat) GetChildren() []*Cat {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Cat) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

//
type CatResponse struct {
	Entity               *Cat     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Cat   `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CatResponse) Reset()         { *m = CatResponse{} }
func (m *CatResponse) String() string { return proto.CompactTextString(m) }
func (*CatResponse) ProtoMessage()    {}
func (*CatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0a5ac8640cab35d, []int{1}
}

func (m *CatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatResponse.Unmarshal(m, b)
}
func (m *CatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatResponse.Marshal(b, m, deterministic)
}
func (m *CatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatResponse.Merge(m, src)
}
func (m *CatResponse) XXX_Size() int {
	return xxx_messageInfo_CatResponse.Size(m)
}
func (m *CatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CatResponse proto.InternalMessageInfo

func (m *CatResponse) GetEntity() *Cat {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *CatResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *CatResponse) GetItems() []*Cat {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CatResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CatResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Cat)(nil), "geiqin.srv.cms.media.Cat")
	proto.RegisterType((*CatResponse)(nil), "geiqin.srv.cms.media.CatResponse")
}

func init() {
	proto.RegisterFile("cat.proto", fileDescriptor_c0a5ac8640cab35d)
}

var fileDescriptor_c0a5ac8640cab35d = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdf, 0x8a, 0x13, 0x31,
	0x14, 0xc6, 0xed, 0xfc, 0xdb, 0xf6, 0x54, 0x44, 0x0e, 0x8a, 0xb1, 0x8b, 0x38, 0xf6, 0x6a, 0xae,
	0x46, 0xb6, 0xe2, 0x0b, 0x58, 0x57, 0x2d, 0xc8, 0x22, 0x53, 0xc5, 0xcb, 0x25, 0x4e, 0xce, 0x6e,
	0x03, 0x4d, 0x32, 0x26, 0xe9, 0x62, 0xf7, 0x05, 0xbc, 0x13, 0x7c, 0x63, 0x49, 0x66, 0xd9, 0xab,
	0x56, 0x85, 0xed, 0xdd, 0xc9, 0x37, 0xbf, 0xf3, 0xe5, 0x9c, 0x6f, 0x20, 0x30, 0x6a, 0xb9, 0xaf,
	0x3b, 0x6b, 0xbc, 0xc1, 0x47, 0x97, 0x24, 0xbf, 0x4b, 0x5d, 0x3b, 0x7b, 0x55, 0xb7, 0xca, 0xd5,
	0x8a, 0x84, 0xe4, 0x93, 0xfb, 0xad, 0x51, 0xca, 0xe8, 0x9e, 0x99, 0xfe, 0x4a, 0x20, 0x9d, 0x73,
	0x8f, 0x0f, 0x20, 0x91, 0x82, 0x0d, 0xca, 0x41, 0x95, 0x37, 0x89, 0x14, 0x88, 0x90, 0x69, 0xae,
	0x88, 0x25, 0xe5, 0xa0, 0x1a, 0x35, 0xb1, 0x0e, 0x9a, 0xdf, 0x76, 0xc4, 0xd2, 0x5e, 0x0b, 0x35,
	0x4e, 0x60, 0xe8, 0xf9, 0x0f, 0xa3, 0x8d, 0xda, 0xb2, 0x2c, 0xea, 0xb7, 0x67, 0x3c, 0x86, 0x91,
	0x37, 0x9e, 0xaf, 0xcf, 0xf5, 0x46, 0xb1, 0x3c, 0x5a, 0x0f, 0xa3, 0x70, 0xb6, 0x51, 0xf8, 0x0c,
	0xa0, 0xff, 0xe8, 0xe4, 0x35, 0xb1, 0xa2, 0x1c, 0x54, 0x49, 0xd3, 0xe3, 0x4b, 0x79, 0x4d, 0xf8,
	0x18, 0x8a, 0x96, 0xfb, 0x73, 0x29, 0xd8, 0x51, 0x6c, 0xcc, 0x5b, 0xee, 0x17, 0x02, 0x9f, 0xc0,
	0xd1, 0xc6, 0x91, 0x0d, 0xfa, 0xb0, 0x1c, 0x54, 0x69, 0x53, 0x84, 0xe3, 0x42, 0xe0, 0x6b, 0x18,
	0xb6, 0x2b, 0xb9, 0x16, 0x96, 0x34, 0x1b, 0x95, 0x69, 0x35, 0x9e, 0x3d, 0xad, 0x77, 0xad, 0x5f,
	0xcf, 0xb9, 0x6f, 0x6e, 0x51, 0x7c, 0x08, 0xa9, 0x14, 0x8e, 0x41, 0x99, 0x56, 0x79, 0x13, 0xca,
	0xe9, 0xcf, 0x04, 0xc6, 0x81, 0x21, 0xd7, 0x19, 0xed, 0x08, 0x4f, 0xa0, 0x20, 0xed, 0xa5, 0xdf,
	0xc6, 0x70, 0xfe, 0x6a, 0x7b, 0x03, 0xe2, 0x09, 0xe4, 0x1d, 0xbf, 0x24, 0x1b, 0xc3, 0x1b, 0xcf,
	0x8e, 0x77, 0x77, 0x7c, 0x0a, 0x48, 0xd3, 0x93, 0xf8, 0x12, 0x72, 0xe9, 0x49, 0x39, 0x96, 0xfe,
	0x6b, 0xf6, 0x9e, 0x0b, 0x77, 0x90, 0xb5, 0xc6, 0xc6, 0xd0, 0xf7, 0xde, 0x71, 0x1a, 0x90, 0xa6,
	0x27, 0xb1, 0x86, 0x4c, 0xea, 0x0b, 0x13, 0xff, 0xc4, 0x78, 0x36, 0xd9, 0xdd, 0xb1, 0xd0, 0x17,
	0xa6, 0x89, 0xdc, 0xec, 0x77, 0x06, 0x30, 0xe7, 0x7e, 0x49, 0xf6, 0x4a, 0xb6, 0x84, 0x1f, 0xa0,
	0x98, 0x5b, 0xe2, 0x9e, 0x70, 0xff, 0x74, 0x93, 0x17, 0xfb, 0x07, 0xbf, 0x09, 0x74, 0x7a, 0x2f,
	0x38, 0x7d, 0xe9, 0xc4, 0x81, 0x9c, 0xde, 0xd2, 0x9a, 0x0e, 0xe0, 0x74, 0x0a, 0xe9, 0x7b, 0xf2,
	0x77, 0xb6, 0x79, 0x07, 0xd9, 0x47, 0xe9, 0x0e, 0xe2, 0xf3, 0xd9, 0xd2, 0xdd, 0xd7, 0x3a, 0x83,
	0x62, 0x49, 0xdc, 0xb6, 0x2b, 0x7c, 0xbe, 0x1b, 0x7f, 0xc3, 0x1d, 0x7d, 0x5d, 0x91, 0xa5, 0xff,
	0xf2, 0xfb, 0x56, 0xc4, 0x57, 0xe3, 0xd5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xcc, 0x70,
	0xac, 0x66, 0x04, 0x00, 0x00,
}
