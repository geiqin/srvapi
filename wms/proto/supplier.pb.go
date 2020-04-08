// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supplier.proto

package geiqin_srv_wms

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

type Supplier struct {
	Id                   int64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 string             `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Shortname            string             `protobuf:"bytes,4,opt,name=shortname,proto3" json:"shortname,omitempty"`
	RegistedCode         string             `protobuf:"bytes,5,opt,name=registed_code,json=registedCode,proto3" json:"registed_code,omitempty"`
	SupplyType           string             `protobuf:"bytes,6,opt,name=supply_type,json=supplyType,proto3" json:"supply_type,omitempty"`
	LogisticsType        int32              `protobuf:"varint,7,opt,name=logistics_type,json=logisticsType,proto3" json:"logistics_type,omitempty"`
	Website              string             `protobuf:"bytes,8,opt,name=website,proto3" json:"website,omitempty"`
	AreaId               int64              `protobuf:"varint,9,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Address              string             `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty"`
	Memo                 string             `protobuf:"bytes,11,opt,name=memo,proto3" json:"memo,omitempty"`
	Verified             bool               `protobuf:"varint,12,opt,name=verified,proto3" json:"verified,omitempty"`
	Status               string             `protobuf:"bytes,13,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string             `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string             `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Contacts             []*SupplierContact `protobuf:"bytes,16,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Supplier) Reset()         { *m = Supplier{} }
func (m *Supplier) String() string { return proto.CompactTextString(m) }
func (*Supplier) ProtoMessage()    {}
func (*Supplier) Descriptor() ([]byte, []int) {
	return fileDescriptor_86038919225469b2, []int{0}
}

func (m *Supplier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Supplier.Unmarshal(m, b)
}
func (m *Supplier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Supplier.Marshal(b, m, deterministic)
}
func (m *Supplier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supplier.Merge(m, src)
}
func (m *Supplier) XXX_Size() int {
	return xxx_messageInfo_Supplier.Size(m)
}
func (m *Supplier) XXX_DiscardUnknown() {
	xxx_messageInfo_Supplier.DiscardUnknown(m)
}

var xxx_messageInfo_Supplier proto.InternalMessageInfo

func (m *Supplier) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Supplier) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Supplier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Supplier) GetShortname() string {
	if m != nil {
		return m.Shortname
	}
	return ""
}

func (m *Supplier) GetRegistedCode() string {
	if m != nil {
		return m.RegistedCode
	}
	return ""
}

func (m *Supplier) GetSupplyType() string {
	if m != nil {
		return m.SupplyType
	}
	return ""
}

func (m *Supplier) GetLogisticsType() int32 {
	if m != nil {
		return m.LogisticsType
	}
	return 0
}

func (m *Supplier) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Supplier) GetAreaId() int64 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *Supplier) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Supplier) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Supplier) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *Supplier) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Supplier) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Supplier) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Supplier) GetContacts() []*SupplierContact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type SupplierWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting              string   `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Shortname            string   `protobuf:"bytes,6,opt,name=shortname,proto3" json:"shortname,omitempty"`
	RegistedCode         string   `protobuf:"bytes,7,opt,name=registed_code,json=registedCode,proto3" json:"registed_code,omitempty"`
	SupplyType           string   `protobuf:"bytes,8,opt,name=supply_type,json=supplyType,proto3" json:"supply_type,omitempty"`
	AreaId               int64    `protobuf:"varint,9,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Status               string   `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupplierWhere) Reset()         { *m = SupplierWhere{} }
func (m *SupplierWhere) String() string { return proto.CompactTextString(m) }
func (*SupplierWhere) ProtoMessage()    {}
func (*SupplierWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_86038919225469b2, []int{1}
}

func (m *SupplierWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupplierWhere.Unmarshal(m, b)
}
func (m *SupplierWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupplierWhere.Marshal(b, m, deterministic)
}
func (m *SupplierWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplierWhere.Merge(m, src)
}
func (m *SupplierWhere) XXX_Size() int {
	return xxx_messageInfo_SupplierWhere.Size(m)
}
func (m *SupplierWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplierWhere.DiscardUnknown(m)
}

var xxx_messageInfo_SupplierWhere proto.InternalMessageInfo

func (m *SupplierWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *SupplierWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SupplierWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
}

func (m *SupplierWhere) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SupplierWhere) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SupplierWhere) GetShortname() string {
	if m != nil {
		return m.Shortname
	}
	return ""
}

func (m *SupplierWhere) GetRegistedCode() string {
	if m != nil {
		return m.RegistedCode
	}
	return ""
}

func (m *SupplierWhere) GetSupplyType() string {
	if m != nil {
		return m.SupplyType
	}
	return ""
}

func (m *SupplierWhere) GetAreaId() int64 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *SupplierWhere) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type SupplierResponse struct {
	Entity               *Supplier   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Supplier `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SupplierResponse) Reset()         { *m = SupplierResponse{} }
func (m *SupplierResponse) String() string { return proto.CompactTextString(m) }
func (*SupplierResponse) ProtoMessage()    {}
func (*SupplierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86038919225469b2, []int{2}
}

func (m *SupplierResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupplierResponse.Unmarshal(m, b)
}
func (m *SupplierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupplierResponse.Marshal(b, m, deterministic)
}
func (m *SupplierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplierResponse.Merge(m, src)
}
func (m *SupplierResponse) XXX_Size() int {
	return xxx_messageInfo_SupplierResponse.Size(m)
}
func (m *SupplierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SupplierResponse proto.InternalMessageInfo

func (m *SupplierResponse) GetEntity() *Supplier {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *SupplierResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *SupplierResponse) GetItems() []*Supplier {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SupplierResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *SupplierResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Supplier)(nil), "geiqin.srv.wms.Supplier")
	proto.RegisterType((*SupplierWhere)(nil), "geiqin.srv.wms.SupplierWhere")
	proto.RegisterType((*SupplierResponse)(nil), "geiqin.srv.wms.SupplierResponse")
}

func init() {
	proto.RegisterFile("supplier.proto", fileDescriptor_86038919225469b2)
}

var fileDescriptor_86038919225469b2 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5f, 0x6b, 0xdb, 0x3e,
	0x14, 0xfd, 0x39, 0x8e, 0x1d, 0xe7, 0xba, 0x49, 0x8b, 0x68, 0x7f, 0x13, 0xd9, 0x4a, 0x4d, 0xc6,
	0xc0, 0x30, 0x30, 0x23, 0x7b, 0xdc, 0x53, 0xc9, 0xfe, 0x50, 0xf6, 0x32, 0x9c, 0x8d, 0x3d, 0x06,
	0xd7, 0xbe, 0x4d, 0x05, 0xb5, 0xe5, 0x49, 0x6a, 0x4b, 0xfa, 0x4d, 0xf6, 0xe1, 0xf6, 0x3d, 0x06,
	0x7b, 0x19, 0x92, 0xec, 0x74, 0xed, 0x6a, 0x3a, 0xc8, 0x9b, 0xee, 0x3d, 0xe7, 0x9e, 0x48, 0xe7,
	0xdc, 0x18, 0xc6, 0xf2, 0xb2, 0xae, 0x2f, 0x18, 0x8a, 0xa4, 0x16, 0x5c, 0x71, 0x32, 0x5e, 0x21,
	0xfb, 0xc6, 0xaa, 0x44, 0x8a, 0xab, 0xe4, 0xba, 0x94, 0x93, 0x9d, 0x9c, 0x97, 0x25, 0xaf, 0x2c,
	0x3a, 0x39, 0x68, 0xd9, 0x73, 0x5e, 0xa9, 0x2c, 0x57, 0xb6, 0x3d, 0xfd, 0xe1, 0x42, 0xb0, 0x68,
	0x10, 0x32, 0x86, 0x1e, 0x2b, 0xa8, 0x13, 0x39, 0xb1, 0x9b, 0xf6, 0x58, 0x41, 0x08, 0xf4, 0x73,
	0x5e, 0x20, 0xed, 0x45, 0x4e, 0x3c, 0x4c, 0xcd, 0x59, 0xf7, 0xaa, 0xac, 0x44, 0xea, 0xda, 0x9e,
	0x3e, 0x93, 0x67, 0x30, 0x94, 0xe7, 0x5c, 0x28, 0x03, 0xf4, 0x0d, 0x70, 0xdb, 0x20, 0xcf, 0x61,
	0x24, 0x70, 0xc5, 0xa4, 0xc2, 0x62, 0x69, 0xe4, 0x3c, 0xc3, 0xd8, 0x69, 0x9b, 0x73, 0x2d, 0x7b,
	0x04, 0xa1, 0xb9, 0xe0, 0x7a, 0xa9, 0xd6, 0x35, 0x52, 0xdf, 0x50, 0xc0, 0xb6, 0x3e, 0xaf, 0x6b,
	0x24, 0x2f, 0x60, 0x7c, 0xc1, 0xf5, 0x00, 0xcb, 0xa5, 0xe5, 0x0c, 0x22, 0x27, 0xf6, 0xd2, 0xd1,
	0xa6, 0x6b, 0x68, 0x14, 0x06, 0xd7, 0x78, 0x2a, 0x99, 0x42, 0x1a, 0x18, 0x8d, 0xb6, 0x24, 0x4f,
	0x60, 0x90, 0x09, 0xcc, 0x96, 0xac, 0xa0, 0x43, 0xf3, 0x42, 0x5f, 0x97, 0x27, 0x85, 0x1e, 0xc9,
	0x8a, 0x42, 0xa0, 0x94, 0x14, 0xec, 0x48, 0x53, 0xea, 0xb7, 0x96, 0x58, 0x72, 0x1a, 0xda, 0xb7,
	0xea, 0x33, 0x99, 0x40, 0x70, 0x85, 0x82, 0x9d, 0x31, 0x2c, 0xe8, 0x4e, 0xe4, 0xc4, 0x41, 0xba,
	0xa9, 0xc9, 0xff, 0xe0, 0x4b, 0x95, 0xa9, 0x4b, 0x49, 0x47, 0x66, 0xa2, 0xa9, 0xc8, 0x21, 0x40,
	0x2e, 0x30, 0xd3, 0x06, 0x64, 0x8a, 0x8e, 0xad, 0x41, 0x4d, 0xe7, 0x58, 0x69, 0xf8, 0xb2, 0x2e,
	0x5a, 0x78, 0xd7, 0xc2, 0x4d, 0xe7, 0x58, 0x91, 0x37, 0x10, 0xe4, 0x36, 0x33, 0x49, 0xf7, 0x22,
	0x37, 0x0e, 0x67, 0x47, 0xc9, 0xdd, 0xa8, 0x93, 0xc5, 0xdd, 0x6c, 0xd3, 0xcd, 0xc0, 0xf4, 0x7b,
	0x0f, 0x46, 0x2d, 0xfa, 0xf5, 0x1c, 0x05, 0x92, 0x7d, 0xf0, 0xea, 0x6c, 0x85, 0x36, 0x67, 0x2f,
	0xb5, 0x05, 0x79, 0x0a, 0x43, 0x7d, 0x58, 0x4a, 0x76, 0x63, 0xf3, 0xf6, 0xd2, 0x40, 0x37, 0x16,
	0xec, 0xc6, 0x98, 0x2a, 0xb9, 0x50, 0xac, 0x5a, 0x35, 0xb1, 0xb7, 0xe5, 0x66, 0x43, 0xfa, 0x0f,
	0x6c, 0x88, 0xd7, 0xb5, 0x21, 0xfe, 0xa3, 0x1b, 0x32, 0x78, 0x7c, 0x43, 0x82, 0xbf, 0x36, 0xa4,
	0x33, 0xe0, 0xdb, 0x58, 0xe0, 0xcf, 0x58, 0xa6, 0x3f, 0x1d, 0xd8, 0x6b, 0xbd, 0x49, 0x51, 0xd6,
	0xbc, 0x92, 0x48, 0x5e, 0x81, 0x8f, 0x95, 0x62, 0x6a, 0x6d, 0xfc, 0x09, 0x67, 0xb4, 0xcb, 0xeb,
	0xb4, 0xe1, 0x91, 0x97, 0xd6, 0x50, 0x61, 0x6c, 0x0b, 0x67, 0x07, 0xf7, 0x07, 0x3e, 0x69, 0xd0,
	0xfa, 0x2c, 0x48, 0x02, 0x1e, 0x53, 0x58, 0x4a, 0xea, 0x9a, 0x24, 0xbb, 0xd5, 0x2d, 0x4d, 0x8b,
	0xa3, 0x10, 0x5c, 0x18, 0x87, 0x1f, 0x10, 0x7f, 0xa7, 0xc1, 0xd4, 0x72, 0x48, 0x0c, 0x7d, 0x56,
	0x9d, 0x71, 0xe3, 0x7c, 0x38, 0xdb, 0xbf, 0xcf, 0x3d, 0xa9, 0xce, 0x78, 0x6a, 0x18, 0xb3, 0x5f,
	0x3d, 0xd8, 0x6d, 0x7f, 0x6a, 0x81, 0xe2, 0x8a, 0xe5, 0x48, 0xde, 0x83, 0x3f, 0x37, 0x3b, 0x49,
	0x3a, 0x6f, 0x35, 0x89, 0x3a, 0xef, 0xdb, 0xf8, 0x37, 0xfd, 0x4f, 0xeb, 0x7c, 0x31, 0xcb, 0xbb,
	0xbd, 0xce, 0x5b, 0xbc, 0xc0, 0xad, 0x75, 0xe6, 0xe0, 0x7e, 0x40, 0xb5, 0xa5, 0xc8, 0x47, 0xf0,
	0x17, 0x98, 0x89, 0xfc, 0x9c, 0x1c, 0x76, 0xb1, 0xcd, 0xdf, 0xeb, 0x5f, 0xc4, 0x4e, 0x7d, 0xf3,
	0xed, 0x7d, 0xfd, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xda, 0xcc, 0xad, 0xc2, 0x05, 0x00, 0x00,
}
