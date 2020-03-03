// Code generated by protoc-gen-go. DO NOT EDIT.
// source: benefit.proto

package geiqin_srv_crm

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

type Benefit struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	IconId               int64    `protobuf:"varint,4,opt,name=icon_id,json=iconId,proto3" json:"icon_id,omitempty"`
	IconUrl              string   `protobuf:"bytes,5,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Selected             bool     `protobuf:"varint,7,opt,name=selected,proto3" json:"selected,omitempty"`
	Package              bool     `protobuf:"varint,8,opt,name=package,proto3" json:"package,omitempty"`
	Params               string   `protobuf:"bytes,9,opt,name=params,proto3" json:"params,omitempty"`
	Memo                 string   `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo,omitempty"`
	CreatedAt            string   `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Ids                  []int32  `protobuf:"varint,13,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Benefit) Reset()         { *m = Benefit{} }
func (m *Benefit) String() string { return proto.CompactTextString(m) }
func (*Benefit) ProtoMessage()    {}
func (*Benefit) Descriptor() ([]byte, []int) {
	return fileDescriptor_76aa0d60ada43779, []int{0}
}

func (m *Benefit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Benefit.Unmarshal(m, b)
}
func (m *Benefit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Benefit.Marshal(b, m, deterministic)
}
func (m *Benefit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Benefit.Merge(m, src)
}
func (m *Benefit) XXX_Size() int {
	return xxx_messageInfo_Benefit.Size(m)
}
func (m *Benefit) XXX_DiscardUnknown() {
	xxx_messageInfo_Benefit.DiscardUnknown(m)
}

var xxx_messageInfo_Benefit proto.InternalMessageInfo

func (m *Benefit) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Benefit) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Benefit) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Benefit) GetIconId() int64 {
	if m != nil {
		return m.IconId
	}
	return 0
}

func (m *Benefit) GetIconUrl() string {
	if m != nil {
		return m.IconUrl
	}
	return ""
}

func (m *Benefit) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Benefit) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

func (m *Benefit) GetPackage() bool {
	if m != nil {
		return m.Package
	}
	return false
}

func (m *Benefit) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

func (m *Benefit) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Benefit) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Benefit) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Benefit) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

//
type BenefitResponse struct {
	Entity               *Benefit   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Benefit `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BenefitResponse) Reset()         { *m = BenefitResponse{} }
func (m *BenefitResponse) String() string { return proto.CompactTextString(m) }
func (*BenefitResponse) ProtoMessage()    {}
func (*BenefitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76aa0d60ada43779, []int{1}
}

func (m *BenefitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BenefitResponse.Unmarshal(m, b)
}
func (m *BenefitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BenefitResponse.Marshal(b, m, deterministic)
}
func (m *BenefitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BenefitResponse.Merge(m, src)
}
func (m *BenefitResponse) XXX_Size() int {
	return xxx_messageInfo_BenefitResponse.Size(m)
}
func (m *BenefitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BenefitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BenefitResponse proto.InternalMessageInfo

func (m *BenefitResponse) GetEntity() *Benefit {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *BenefitResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *BenefitResponse) GetItems() []*Benefit {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *BenefitResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *BenefitResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Benefit)(nil), "geiqin.srv.crm.Benefit")
	proto.RegisterType((*BenefitResponse)(nil), "geiqin.srv.crm.BenefitResponse")
}

func init() { proto.RegisterFile("benefit.proto", fileDescriptor_76aa0d60ada43779) }

var fileDescriptor_76aa0d60ada43779 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xff, 0xb6, 0x63, 0x27, 0x99, 0xb4, 0xf9, 0xa3, 0x11, 0xd0, 0x6d, 0x24, 0x84, 0x95,
	0x93, 0x25, 0x84, 0x91, 0xc2, 0x13, 0x04, 0x0a, 0xa8, 0x37, 0xe4, 0xaa, 0xe2, 0x58, 0x6d, 0xed,
	0x49, 0xba, 0xc2, 0xde, 0x35, 0xbb, 0x9b, 0x4a, 0x7d, 0x36, 0x9e, 0x89, 0x23, 0x77, 0xb4, 0xbb,
	0x0e, 0x12, 0x85, 0x22, 0xa4, 0xdc, 0x66, 0xe6, 0xfb, 0xfc, 0x9b, 0xd5, 0x37, 0x32, 0x1c, 0x5f,
	0x93, 0xa4, 0x8d, 0xb0, 0x65, 0xaf, 0x95, 0x55, 0x38, 0xdf, 0x92, 0xf8, 0x22, 0x64, 0x69, 0xf4,
	0x6d, 0x59, 0xeb, 0x6e, 0x71, 0x54, 0xab, 0xae, 0x53, 0x32, 0xa8, 0xcb, 0xaf, 0x31, 0x8c, 0xdf,
	0x04, 0x3f, 0xce, 0x21, 0x16, 0x0d, 0x8b, 0xf2, 0xa8, 0x48, 0xab, 0x58, 0x34, 0x88, 0x30, 0x32,
	0xed, 0x6e, 0xcb, 0xe2, 0x3c, 0x2a, 0xa6, 0x95, 0xaf, 0xdd, 0x4c, 0xf2, 0x8e, 0x58, 0x12, 0x66,
	0xae, 0xc6, 0x13, 0x18, 0x8b, 0x5a, 0xc9, 0x2b, 0xd1, 0xb0, 0x51, 0x1e, 0x15, 0x49, 0x95, 0xb9,
	0xf6, 0xbc, 0xc1, 0x53, 0x98, 0x78, 0x61, 0xa7, 0x5b, 0x96, 0xfa, 0x0f, 0xbc, 0xf1, 0x52, 0xb7,
	0x8e, 0x63, 0xef, 0x7a, 0x62, 0x59, 0xe0, 0xb8, 0x1a, 0x17, 0x30, 0x31, 0xd4, 0x52, 0x6d, 0xa9,
	0x61, 0xe3, 0x3c, 0x2a, 0x26, 0xd5, 0xcf, 0x1e, 0x19, 0x8c, 0x7b, 0x5e, 0x7f, 0xe6, 0x5b, 0x62,
	0x13, 0x2f, 0xed, 0x5b, 0x7c, 0x0a, 0x59, 0xcf, 0x35, 0xef, 0x0c, 0x9b, 0x7a, 0xd6, 0xd0, 0xb9,
	0x0d, 0x1d, 0x75, 0x8a, 0x41, 0xd8, 0xe0, 0x6a, 0x7c, 0x06, 0x50, 0x6b, 0xe2, 0x96, 0x9a, 0x2b,
	0x6e, 0xd9, 0xcc, 0x2b, 0xd3, 0x61, 0xb2, 0xb6, 0x4e, 0xde, 0xf5, 0xcd, 0x5e, 0x3e, 0x0a, 0xf2,
	0x30, 0x59, 0x5b, 0x7c, 0x04, 0x89, 0x68, 0x0c, 0x3b, 0xce, 0x93, 0x22, 0xad, 0x5c, 0xb9, 0xfc,
	0x1e, 0xc1, 0xff, 0x43, 0x7a, 0x15, 0x99, 0x5e, 0x49, 0x43, 0xf8, 0x0a, 0x32, 0x92, 0x56, 0xd8,
	0x3b, 0x9f, 0xe4, 0x6c, 0x75, 0x52, 0xfe, 0x7a, 0x80, 0x72, 0xff, 0xc1, 0x60, 0xc3, 0x17, 0x90,
	0xf6, 0x7c, 0x4b, 0xda, 0xe7, 0x3c, 0x5b, 0x3d, 0xb9, 0xef, 0xff, 0xe8, 0xc4, 0x2a, 0x78, 0xf0,
	0x25, 0xa4, 0xc2, 0x52, 0x67, 0x58, 0x92, 0x27, 0x7f, 0x83, 0x07, 0x97, 0x63, 0x93, 0xd6, 0x4a,
	0xfb, 0xc3, 0xfc, 0x81, 0xfd, 0xce, 0x89, 0x55, 0xf0, 0x60, 0x01, 0x23, 0x21, 0x37, 0xca, 0x9f,
	0x6a, 0xb6, 0x7a, 0x7c, 0xdf, 0x7b, 0x2e, 0x37, 0xaa, 0xf2, 0x8e, 0xd5, 0xb7, 0x18, 0xe6, 0xc3,
	0xa6, 0x0b, 0xd2, 0xb7, 0xa2, 0x26, 0x3c, 0x83, 0xec, 0xad, 0x0f, 0x12, 0x1f, 0x7a, 0xd3, 0xe2,
	0xf9, 0x43, 0x8f, 0x1d, 0xa2, 0x5b, 0xfe, 0xe7, 0x28, 0x97, 0x3e, 0xef, 0x43, 0x29, 0x67, 0xd4,
	0xd2, 0x81, 0x94, 0x35, 0x24, 0x1f, 0xc8, 0x1e, 0x84, 0x78, 0x0f, 0xd9, 0x05, 0x71, 0x5d, 0xdf,
	0xe0, 0xe9, 0x6f, 0x66, 0x6e, 0xe8, 0xd3, 0x0d, 0x69, 0xfa, 0x07, 0xce, 0x75, 0xe6, 0x7f, 0xd6,
	0xd7, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x02, 0xcb, 0x4c, 0xdb, 0x03, 0x00, 0x00,
}
