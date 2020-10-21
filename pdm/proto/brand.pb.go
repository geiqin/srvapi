// Code generated by protoc-gen-go. DO NOT EDIT.
// source: brand.proto

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

type Brand struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Website              string   `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	LogoId               int64    `protobuf:"varint,4,opt,name=logo_id,json=logoId,proto3" json:"logo_id,omitempty"`
	LogoUrl              string   `protobuf:"bytes,5,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	BrandType            string   `protobuf:"bytes,6,opt,name=brand_type,json=brandType,proto3" json:"brand_type,omitempty"`
	ItemNum              int32    `protobuf:"varint,7,opt,name=item_num,json=itemNum,proto3" json:"item_num,omitempty"`
	Description          string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Brand) Reset()         { *m = Brand{} }
func (m *Brand) String() string { return proto.CompactTextString(m) }
func (*Brand) ProtoMessage()    {}
func (*Brand) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{0}
}

func (m *Brand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Brand.Unmarshal(m, b)
}
func (m *Brand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Brand.Marshal(b, m, deterministic)
}
func (m *Brand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Brand.Merge(m, src)
}
func (m *Brand) XXX_Size() int {
	return xxx_messageInfo_Brand.Size(m)
}
func (m *Brand) XXX_DiscardUnknown() {
	xxx_messageInfo_Brand.DiscardUnknown(m)
}

var xxx_messageInfo_Brand proto.InternalMessageInfo

func (m *Brand) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Brand) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Brand) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Brand) GetLogoId() int64 {
	if m != nil {
		return m.LogoId
	}
	return 0
}

func (m *Brand) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func (m *Brand) GetBrandType() string {
	if m != nil {
		return m.BrandType
	}
	return ""
}

func (m *Brand) GetItemNum() int32 {
	if m != nil {
		return m.ItemNum
	}
	return 0
}

func (m *Brand) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Brand) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Brand) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type BrandResponse struct {
	Entity               *Brand   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Brand `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BrandResponse) Reset()         { *m = BrandResponse{} }
func (m *BrandResponse) String() string { return proto.CompactTextString(m) }
func (*BrandResponse) ProtoMessage()    {}
func (*BrandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{1}
}

func (m *BrandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrandResponse.Unmarshal(m, b)
}
func (m *BrandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrandResponse.Marshal(b, m, deterministic)
}
func (m *BrandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrandResponse.Merge(m, src)
}
func (m *BrandResponse) XXX_Size() int {
	return xxx_messageInfo_BrandResponse.Size(m)
}
func (m *BrandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BrandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BrandResponse proto.InternalMessageInfo

func (m *BrandResponse) GetEntity() *Brand {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *BrandResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *BrandResponse) GetItems() []*Brand {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *BrandResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *BrandResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Brand)(nil), "geiqin.srv.pdm.Brand")
	proto.RegisterType((*BrandResponse)(nil), "geiqin.srv.pdm.BrandResponse")
}

func init() {
	proto.RegisterFile("brand.proto", fileDescriptor_0b213c818d79a472)
}

var fileDescriptor_0b213c818d79a472 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x5f, 0x6b, 0xd4, 0x40,
	0x14, 0xc5, 0xcd, 0x66, 0x37, 0xdb, 0xbd, 0x5b, 0xfb, 0x30, 0x28, 0x4e, 0x0b, 0x85, 0xb0, 0x4f,
	0x0b, 0x62, 0x1e, 0xd6, 0x67, 0xc5, 0x6a, 0x45, 0x0a, 0x22, 0x92, 0x5a, 0x7c, 0x5c, 0xb2, 0x99,
	0xdb, 0xed, 0x40, 0x32, 0x33, 0xce, 0x4c, 0x2a, 0xfb, 0x39, 0x7c, 0xf2, 0x53, 0xfa, 0x15, 0x64,
	0xee, 0xa4, 0xe0, 0x9f, 0x88, 0xd0, 0x7d, 0xcb, 0xbd, 0xe7, 0x77, 0x4f, 0x4e, 0x0e, 0x04, 0xe6,
	0x1b, 0x5b, 0x29, 0x51, 0x18, 0xab, 0xbd, 0x66, 0x47, 0x5b, 0x94, 0x5f, 0xa4, 0x2a, 0x9c, 0xbd,
	0x2d, 0x8c, 0x68, 0x4f, 0x0e, 0x6b, 0xdd, 0xb6, 0x5a, 0x45, 0x75, 0xf1, 0x6d, 0x04, 0x93, 0xd7,
	0x81, 0x66, 0x47, 0x30, 0x92, 0x82, 0x27, 0x79, 0xb2, 0x9c, 0x94, 0x23, 0x29, 0x18, 0x83, 0xb1,
	0xaa, 0x5a, 0xe4, 0xa3, 0x3c, 0x59, 0xce, 0x4a, 0x7a, 0x66, 0x1c, 0xa6, 0x5f, 0x71, 0xe3, 0xa4,
	0x47, 0x9e, 0xd2, 0xfa, 0x6e, 0x64, 0x4f, 0x60, 0xda, 0xe8, 0xad, 0x5e, 0x4b, 0xc1, 0xc7, 0x79,
	0xb2, 0x4c, 0xcb, 0x2c, 0x8c, 0x17, 0x82, 0x1d, 0xc3, 0x01, 0x09, 0x9d, 0x6d, 0xf8, 0x24, 0xde,
	0x84, 0xf9, 0xca, 0x36, 0xec, 0x14, 0x80, 0x82, 0xae, 0xfd, 0xce, 0x20, 0xcf, 0x48, 0x9c, 0xd1,
	0xe6, 0xd3, 0xce, 0x60, 0xb8, 0x94, 0x1e, 0xdb, 0xb5, 0xea, 0x5a, 0x3e, 0xa5, 0x58, 0xd3, 0x30,
	0x7f, 0xe8, 0x5a, 0x96, 0xc3, 0x5c, 0xa0, 0xab, 0xad, 0x34, 0x5e, 0x6a, 0xc5, 0x0f, 0xe8, 0xf4,
	0xd7, 0x55, 0xf0, 0xae, 0x2d, 0x56, 0x1e, 0xc5, 0xba, 0xf2, 0x7c, 0x16, 0xbd, 0xfb, 0xcd, 0x99,
	0x0f, 0x72, 0x67, 0xc4, 0x9d, 0x0c, 0x51, 0xee, 0x37, 0x67, 0x7e, 0xf1, 0x23, 0x81, 0x87, 0xd4,
	0x4a, 0x89, 0xce, 0x68, 0xe5, 0x90, 0x3d, 0x83, 0x0c, 0x95, 0x97, 0x7e, 0x47, 0x0d, 0xcd, 0x57,
	0x8f, 0x8b, 0xdf, 0x6b, 0x2d, 0x22, 0xde, 0x43, 0xec, 0x29, 0x4c, 0x4c, 0xb5, 0x45, 0x4b, 0xed,
	0x0d, 0xd0, 0x1f, 0x83, 0x58, 0x46, 0x26, 0xc0, 0xe1, 0xc3, 0x1c, 0x4f, 0xf3, 0xf4, 0xdf, 0xd6,
	0x91, 0x09, 0x30, 0x5a, 0xab, 0x2d, 0xd5, 0x3c, 0x00, 0xbf, 0x0d, 0x62, 0x19, 0x19, 0xb6, 0x84,
	0xb1, 0x54, 0xd7, 0x9a, 0x8a, 0x9f, 0xaf, 0x1e, 0xfd, 0xc9, 0x5e, 0xa8, 0x6b, 0x5d, 0x12, 0xb1,
	0xfa, 0x9e, 0xc2, 0x21, 0xbd, 0xe7, 0x12, 0xed, 0xad, 0xac, 0x91, 0xbd, 0x82, 0xec, 0x0d, 0xd5,
	0xc5, 0x86, 0xf3, 0x9c, 0x9c, 0x0e, 0xc7, 0xec, 0x0b, 0x5b, 0x3c, 0x08, 0x0e, 0x57, 0xd4, 0xe8,
	0x3e, 0x0e, 0xe7, 0xd8, 0xe0, 0x1e, 0x0e, 0x2f, 0x20, 0x7d, 0x87, 0xfe, 0xde, 0xe7, 0xe7, 0x90,
	0x5d, 0x62, 0x65, 0xeb, 0x1b, 0x76, 0xfc, 0x17, 0x5a, 0x39, 0xfc, 0x7c, 0x83, 0x16, 0xff, 0xef,
	0xf2, 0x12, 0xc6, 0xef, 0xa5, 0xbb, 0x77, 0x8a, 0x4d, 0x46, 0xbf, 0xea, 0xf3, 0x9f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xf2, 0xe4, 0xc7, 0xe2, 0xd7, 0x03, 0x00, 0x00,
}
