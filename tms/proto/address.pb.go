// Code generated by protoc-gen-go. DO NOT EDIT.
// source: address.proto

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

type AddressWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting              string   `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Keywords             string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id                   int64    `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Type                 int32    `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
	IsDefault            bool     `protobuf:"varint,8,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressWhere) Reset()         { *m = AddressWhere{} }
func (m *AddressWhere) String() string { return proto.CompactTextString(m) }
func (*AddressWhere) ProtoMessage()    {}
func (*AddressWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{0}
}

func (m *AddressWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressWhere.Unmarshal(m, b)
}
func (m *AddressWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressWhere.Marshal(b, m, deterministic)
}
func (m *AddressWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressWhere.Merge(m, src)
}
func (m *AddressWhere) XXX_Size() int {
	return xxx_messageInfo_AddressWhere.Size(m)
}
func (m *AddressWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressWhere.DiscardUnknown(m)
}

var xxx_messageInfo_AddressWhere proto.InternalMessageInfo

func (m *AddressWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *AddressWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *AddressWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
}

func (m *AddressWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *AddressWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AddressWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *AddressWhere) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *AddressWhere) GetIsDefault() bool {
	if m != nil {
		return m.IsDefault
	}
	return false
}

type Address struct {
	Id                   int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AreaId               int64          `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Addr                 string         `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	Lng                  string         `protobuf:"bytes,4,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat                  string         `protobuf:"bytes,5,opt,name=lat,proto3" json:"lat,omitempty"`
	Contact              string         `protobuf:"bytes,6,opt,name=contact,proto3" json:"contact,omitempty"`
	Tel                  string         `protobuf:"bytes,7,opt,name=tel,proto3" json:"tel,omitempty"`
	Mobile               string         `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty"`
	CreatedAt            string         `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string         `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Area                 *AreaInfo      `protobuf:"bytes,11,opt,name=area,proto3" json:"area,omitempty"`
	AddressTypes         []*AddressType `protobuf:"bytes,12,rep,name=address_types,json=addressTypes,proto3" json:"address_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{1}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Address) GetAreaId() int64 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *Address) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Address) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *Address) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *Address) GetContact() string {
	if m != nil {
		return m.Contact
	}
	return ""
}

func (m *Address) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

func (m *Address) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Address) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Address) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Address) GetArea() *AreaInfo {
	if m != nil {
		return m.Area
	}
	return nil
}

func (m *Address) GetAddressTypes() []*AddressType {
	if m != nil {
		return m.AddressTypes
	}
	return nil
}

type AddressResponse struct {
	Entity               *Address   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Address `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AddressResponse) Reset()         { *m = AddressResponse{} }
func (m *AddressResponse) String() string { return proto.CompactTextString(m) }
func (*AddressResponse) ProtoMessage()    {}
func (*AddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{2}
}

func (m *AddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressResponse.Unmarshal(m, b)
}
func (m *AddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressResponse.Marshal(b, m, deterministic)
}
func (m *AddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressResponse.Merge(m, src)
}
func (m *AddressResponse) XXX_Size() int {
	return xxx_messageInfo_AddressResponse.Size(m)
}
func (m *AddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddressResponse proto.InternalMessageInfo

func (m *AddressResponse) GetEntity() *Address {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *AddressResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *AddressResponse) GetItems() []*Address {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *AddressResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *AddressResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*AddressWhere)(nil), "geiqin.srv.tms.AddressWhere")
	proto.RegisterType((*Address)(nil), "geiqin.srv.tms.Address")
	proto.RegisterType((*AddressResponse)(nil), "geiqin.srv.tms.AddressResponse")
}

func init() {
	proto.RegisterFile("address.proto", fileDescriptor_982c640dad8fe78e)
}

var fileDescriptor_982c640dad8fe78e = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0x1c, 0x27, 0x4e, 0x3c, 0xce, 0x17, 0x60, 0x55, 0xe8, 0x2a, 0x05, 0x61, 0xe5, 0xca,
	0x12, 0x10, 0x24, 0xf3, 0x02, 0x44, 0x14, 0xaa, 0x4a, 0x20, 0xa1, 0x2d, 0x88, 0xcb, 0xc8, 0xf5,
	0x4e, 0xd3, 0x15, 0x89, 0xd7, 0xec, 0x6e, 0x8b, 0xd2, 0xc7, 0xe0, 0xb9, 0x78, 0x04, 0x5e, 0x83,
	0x7b, 0xb4, 0x3f, 0x2d, 0x50, 0x88, 0x84, 0x94, 0xde, 0xcd, 0xcc, 0x39, 0x39, 0x7b, 0xe6, 0x27,
	0x86, 0xff, 0x2b, 0xce, 0x15, 0x6a, 0x3d, 0x6d, 0x95, 0x34, 0x92, 0x8c, 0x16, 0x28, 0x3e, 0x89,
	0x66, 0xaa, 0xd5, 0xf9, 0xd4, 0xac, 0xf4, 0x78, 0x58, 0xcb, 0xd5, 0x4a, 0x36, 0x1e, 0x1d, 0xdf,
	0x09, 0xe4, 0x77, 0xeb, 0x16, 0x43, 0x69, 0x54, 0x29, 0xac, 0x0e, 0x9b, 0x13, 0xe9, 0xf3, 0xc9,
	0xd7, 0x08, 0x86, 0x33, 0xcf, 0xfa, 0x70, 0x8a, 0x0a, 0xc9, 0x0e, 0xf4, 0xda, 0x6a, 0x81, 0x9c,
	0x46, 0x79, 0x54, 0xf4, 0x98, 0x4f, 0xc8, 0x1e, 0xa4, 0x36, 0x98, 0x6b, 0x71, 0x81, 0xb4, 0xe3,
	0x90, 0x81, 0x2d, 0x1c, 0x89, 0x0b, 0x24, 0x14, 0xfa, 0x5a, 0x2a, 0x23, 0x9a, 0x05, 0x8d, 0xf3,
	0xa8, 0x48, 0xd9, 0x65, 0x4a, 0xc6, 0x30, 0xf8, 0x88, 0xeb, 0xcf, 0x52, 0x71, 0x4d, 0xbb, 0x0e,
	0xba, 0xca, 0xc9, 0x08, 0x3a, 0x82, 0xd3, 0x5e, 0x1e, 0x15, 0x31, 0xeb, 0x08, 0x4e, 0x6e, 0x43,
	0x2c, 0xb8, 0xa6, 0x49, 0x1e, 0x17, 0x31, 0xb3, 0x21, 0x21, 0xd0, 0x35, 0xeb, 0x16, 0x69, 0xdf,
	0xbd, 0xe7, 0x62, 0xf2, 0x00, 0x40, 0xe8, 0x39, 0xc7, 0x93, 0xea, 0x6c, 0x69, 0xe8, 0x20, 0x8f,
	0x8a, 0x01, 0x4b, 0x85, 0xde, 0xf7, 0x85, 0xc9, 0xb7, 0x0e, 0xf4, 0x43, 0x3b, 0xe1, 0x81, 0xe8,
	0xea, 0x81, 0x5d, 0xe8, 0xdb, 0xe6, 0xe7, 0x82, 0xbb, 0x0e, 0x62, 0x96, 0xb8, 0x59, 0x70, 0xfb,
	0x8e, 0x1d, 0x54, 0x30, 0xef, 0x62, 0xeb, 0x66, 0xd9, 0x2c, 0x82, 0x69, 0x1b, 0xba, 0x4a, 0x65,
	0x9c, 0x61, 0x5b, 0xa9, 0x8c, 0xed, 0xbb, 0x96, 0x8d, 0xa9, 0x6a, 0x43, 0x13, 0xdf, 0x77, 0x48,
	0x2d, 0xd7, 0xe0, 0xd2, 0x19, 0x4f, 0x99, 0x0d, 0xc9, 0x3d, 0x48, 0x56, 0xf2, 0x58, 0x2c, 0xd1,
	0x79, 0x4e, 0x59, 0xc8, 0x6c, 0x3f, 0xb5, 0xc2, 0xca, 0x20, 0x9f, 0x57, 0x86, 0xa6, 0x0e, 0x4b,
	0x43, 0x65, 0x66, 0x2c, 0x7c, 0xd6, 0xf2, 0x4b, 0x18, 0x3c, 0x1c, 0x2a, 0x33, 0x43, 0x1e, 0x43,
	0xd7, 0xf6, 0x40, 0xb3, 0x3c, 0x2a, 0xb2, 0x92, 0x4e, 0x7f, 0xbf, 0x86, 0xe9, 0x2c, 0xec, 0x9a,
	0x39, 0x16, 0x79, 0x7e, 0x75, 0x3d, 0x73, 0x3b, 0x4b, 0x4d, 0x87, 0x79, 0x5c, 0x64, 0xe5, 0xde,
	0x1f, 0x3f, 0xfb, 0x79, 0x35, 0x6c, 0xf8, 0xcb, 0x09, 0xe9, 0xc9, 0xf7, 0x08, 0x6e, 0x05, 0x94,
	0xa1, 0x6e, 0x65, 0xa3, 0x91, 0x3c, 0x85, 0x04, 0x1b, 0x23, 0xcc, 0xda, 0x8d, 0x3a, 0x2b, 0x77,
	0x37, 0xc8, 0xb1, 0x40, 0x23, 0x8f, 0xfc, 0x85, 0x29, 0xb7, 0x85, 0xac, 0xbc, 0x7b, 0x9d, 0xff,
	0xd6, 0x82, 0xfe, 0xf0, 0x14, 0x79, 0x02, 0x3d, 0x61, 0x70, 0xa5, 0x69, 0xec, 0xbc, 0x6e, 0x14,
	0xf7, 0x2c, 0xab, 0x8d, 0x4a, 0x49, 0xe5, 0x16, 0xf7, 0x17, 0xed, 0x97, 0x16, 0x64, 0x9e, 0x43,
	0x0a, 0xe8, 0x8a, 0xe6, 0x44, 0xba, 0x95, 0x66, 0xe5, 0xce, 0x75, 0xae, 0x9f, 0x9c, 0x65, 0x94,
	0x5f, 0xba, 0x30, 0x0a, 0x2f, 0x1d, 0xa1, 0x3a, 0x17, 0x35, 0x92, 0x7d, 0x48, 0x5e, 0xb8, 0x35,
	0x91, 0x4d, 0x9e, 0xc6, 0x0f, 0x37, 0x99, 0x0d, 0xa3, 0x9b, 0xfc, 0x67, 0x55, 0xde, 0xbb, 0x6d,
	0x6e, 0xa5, 0xf2, 0x0a, 0xe2, 0x03, 0x34, 0xe4, 0xfe, 0x06, 0xa6, 0xfb, 0x63, 0xff, 0x8b, 0xce,
	0x21, 0x24, 0xfb, 0xb8, 0x44, 0x83, 0x37, 0x22, 0x75, 0x84, 0x95, 0xaa, 0x4f, 0xb7, 0x97, 0x3a,
	0x80, 0xee, 0x6b, 0xa1, 0x6f, 0xa0, 0xbd, 0x37, 0x00, 0x07, 0x68, 0xc2, 0xa7, 0x62, 0x6b, 0xb9,
	0xe3, 0xc4, 0x7d, 0x41, 0x9f, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xc9, 0x55, 0x26, 0x93,
	0x05, 0x00, 0x00,
}
