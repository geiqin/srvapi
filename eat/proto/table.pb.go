// Code generated by protoc-gen-go. DO NOT EDIT.
// source: table.proto

package geiqin_srv_eat

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

type TableWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Top                  int32    `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	Id                   int64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Status               int32    `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Keywords             string   `protobuf:"bytes,7,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Names                string   `protobuf:"bytes,8,opt,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableWhere) Reset()         { *m = TableWhere{} }
func (m *TableWhere) String() string { return proto.CompactTextString(m) }
func (*TableWhere) ProtoMessage()    {}
func (*TableWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_448a2743262f7a00, []int{0}
}

func (m *TableWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableWhere.Unmarshal(m, b)
}
func (m *TableWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableWhere.Marshal(b, m, deterministic)
}
func (m *TableWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableWhere.Merge(m, src)
}
func (m *TableWhere) XXX_Size() int {
	return xxx_messageInfo_TableWhere.Size(m)
}
func (m *TableWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_TableWhere.DiscardUnknown(m)
}

var xxx_messageInfo_TableWhere proto.InternalMessageInfo

func (m *TableWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *TableWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *TableWhere) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *TableWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TableWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *TableWhere) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TableWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *TableWhere) GetNames() string {
	if m != nil {
		return m.Names
	}
	return ""
}

type Table struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Qrcode               string   `protobuf:"bytes,3,opt,name=qrcode,proto3" json:"qrcode,omitempty"`
	Status               int32    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_448a2743262f7a00, []int{1}
}

func (m *Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table.Unmarshal(m, b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table.Marshal(b, m, deterministic)
}
func (m *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(m, src)
}
func (m *Table) XXX_Size() int {
	return xxx_messageInfo_Table.Size(m)
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

func (m *Table) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Table) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Table) GetQrcode() string {
	if m != nil {
		return m.Qrcode
	}
	return ""
}

func (m *Table) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Table) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Table) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type TableResponse struct {
	Entity               *Table   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Table `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableResponse) Reset()         { *m = TableResponse{} }
func (m *TableResponse) String() string { return proto.CompactTextString(m) }
func (*TableResponse) ProtoMessage()    {}
func (*TableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_448a2743262f7a00, []int{2}
}

func (m *TableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableResponse.Unmarshal(m, b)
}
func (m *TableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableResponse.Marshal(b, m, deterministic)
}
func (m *TableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableResponse.Merge(m, src)
}
func (m *TableResponse) XXX_Size() int {
	return xxx_messageInfo_TableResponse.Size(m)
}
func (m *TableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TableResponse proto.InternalMessageInfo

func (m *TableResponse) GetEntity() *Table {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *TableResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *TableResponse) GetItems() []*Table {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *TableResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *TableResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*TableWhere)(nil), "geiqin.srv.eat.TableWhere")
	proto.RegisterType((*Table)(nil), "geiqin.srv.eat.Table")
	proto.RegisterType((*TableResponse)(nil), "geiqin.srv.eat.TableResponse")
}

func init() {
	proto.RegisterFile("table.proto", fileDescriptor_448a2743262f7a00)
}

var fileDescriptor_448a2743262f7a00 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0xcd, 0xe6, 0x8f, 0x9b, 0xb3, 0xb5, 0xc8, 0x50, 0x65, 0x58, 0x29, 0x84, 0xbd, 0x0a,
	0x88, 0xb9, 0x58, 0xaf, 0x05, 0xcb, 0x1a, 0x44, 0x11, 0x2a, 0xb3, 0x8a, 0x97, 0x25, 0x4d, 0x4e,
	0xdb, 0xc1, 0xdd, 0x4c, 0x3a, 0x33, 0xad, 0xb6, 0x6f, 0xe2, 0xbb, 0xe8, 0x33, 0xf9, 0x0a, 0x32,
	0x67, 0x62, 0x71, 0x25, 0xe2, 0x45, 0xee, 0xce, 0x39, 0xdf, 0x37, 0x73, 0x7e, 0x27, 0x67, 0x08,
	0xcc, 0x6c, 0x75, 0xba, 0xc1, 0xa2, 0xd3, 0xca, 0x2a, 0xb6, 0x7f, 0x8e, 0xf2, 0x52, 0xb6, 0x85,
	0xd1, 0xd7, 0x05, 0x56, 0x76, 0xbe, 0x57, 0xab, 0xed, 0x56, 0xb5, 0x5e, 0x5d, 0xfc, 0x08, 0x00,
	0x3e, 0x38, 0xf7, 0xa7, 0x0b, 0xd4, 0xc8, 0x0e, 0x20, 0xee, 0xaa, 0x73, 0x6c, 0x78, 0x90, 0x05,
	0x79, 0x2c, 0x7c, 0xc2, 0x9e, 0x40, 0xea, 0x82, 0x13, 0x23, 0x6f, 0x91, 0x4f, 0x48, 0x99, 0xba,
	0xc2, 0x5a, 0xde, 0x22, 0x7b, 0x08, 0xa1, 0x55, 0x1d, 0x0f, 0xa9, 0xec, 0x42, 0xb6, 0x0f, 0x13,
	0xd9, 0xf0, 0x28, 0x0b, 0xf2, 0x50, 0x4c, 0x64, 0xe3, 0x1c, 0xb2, 0x31, 0x3c, 0xce, 0xc2, 0x3c,
	0x14, 0x2e, 0x64, 0x8f, 0x21, 0x31, 0xb6, 0xb2, 0x57, 0x86, 0x27, 0x74, 0xac, 0xcf, 0xd8, 0x1c,
	0xa6, 0x9f, 0xf1, 0xe6, 0x8b, 0xd2, 0x8d, 0xe1, 0xf7, 0xb3, 0x20, 0x4f, 0xc5, 0x5d, 0xee, 0xd0,
	0xda, 0x6a, 0x8b, 0x86, 0x4f, 0x49, 0xf0, 0xc9, 0xe2, 0x5b, 0x00, 0x31, 0xf1, 0xf7, 0x5d, 0x83,
	0xbb, 0xae, 0x0c, 0x22, 0x67, 0x21, 0xde, 0x54, 0x50, 0xec, 0xfa, 0x5e, 0xea, 0x5a, 0x35, 0x48,
	0xb8, 0xa9, 0xe8, 0xb3, 0x3f, 0x78, 0xa2, 0x1d, 0x9e, 0x43, 0x80, 0x5a, 0x63, 0x65, 0xb1, 0x39,
	0xa9, 0x2c, 0x8f, 0xe9, 0x4c, 0xda, 0x57, 0x8e, 0xac, 0x93, 0xaf, 0xba, 0xe6, 0xb7, 0x9c, 0x78,
	0xb9, 0xaf, 0x1c, 0xd9, 0xc5, 0xcf, 0x00, 0x1e, 0x10, 0x9b, 0x40, 0xd3, 0xa9, 0xd6, 0x20, 0x7b,
	0x06, 0x09, 0xb6, 0x56, 0xda, 0x1b, 0xe2, 0x9c, 0x2d, 0x1f, 0x15, 0xbb, 0xcb, 0x29, 0xbc, 0xbd,
	0x37, 0xb1, 0xa7, 0x7e, 0x1b, 0x9a, 0x66, 0x18, 0x70, 0xbf, 0x77, 0xa2, 0x5f, 0x92, 0x76, 0x66,
	0x69, 0x71, 0x6b, 0x78, 0x98, 0x85, 0xff, 0xbe, 0xda, 0x7b, 0x9c, 0x19, 0xb5, 0x56, 0x9a, 0xe6,
	0x1d, 0x30, 0x97, 0x4e, 0x14, 0xde, 0xc3, 0x72, 0x88, 0x64, 0x7b, 0xa6, 0x68, 0xfe, 0xd9, 0xf2,
	0xe0, 0x6f, 0xef, 0x9b, 0xf6, 0x4c, 0x09, 0x72, 0x2c, 0xbf, 0x47, 0xb0, 0x47, 0x7d, 0xd6, 0xa8,
	0xaf, 0x65, 0x8d, 0xac, 0x84, 0x64, 0x8d, 0x95, 0xae, 0x2f, 0xd8, 0x7c, 0x90, 0x87, 0x5e, 0xdd,
	0xfc, 0x70, 0x98, 0xb5, 0xff, 0x6a, 0x8b, 0x7b, 0x6c, 0x05, 0xd1, 0x3b, 0x69, 0xec, 0xb8, 0x4b,
	0x5e, 0x40, 0xf8, 0x1a, 0x2d, 0x1b, 0xfe, 0x30, 0xff, 0x3f, 0x5e, 0x42, 0xb2, 0xa2, 0xcd, 0x8f,
	0xa3, 0x78, 0x09, 0xc9, 0x47, 0x7a, 0x21, 0x63, 0x40, 0x5e, 0xe1, 0x06, 0xc7, 0x82, 0xbc, 0x85,
	0xd9, 0x71, 0x87, 0xed, 0xb1, 0x5e, 0x6d, 0x94, 0x19, 0x79, 0x57, 0x09, 0x49, 0xf9, 0xb5, 0x53,
	0x7a, 0xdc, 0x86, 0x4e, 0x13, 0xfa, 0x27, 0x3d, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x03, 0xbc,
	0x2b, 0x89, 0xc0, 0x04, 0x00, 0x00,
}
