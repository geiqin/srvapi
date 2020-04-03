// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statement.proto

package geiqin_srv_dms

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

type Statement struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SourceId             int64    `protobuf:"varint,2,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	DistributorId        int64    `protobuf:"varint,3,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	Money                float32  `protobuf:"fixed32,4,opt,name=money,proto3" json:"money,omitempty"`
	Source               string   `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Statement) Reset()         { *m = Statement{} }
func (m *Statement) String() string { return proto.CompactTextString(m) }
func (*Statement) ProtoMessage()    {}
func (*Statement) Descriptor() ([]byte, []int) {
	return fileDescriptor_4430318ac3171d1a, []int{0}
}

func (m *Statement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statement.Unmarshal(m, b)
}
func (m *Statement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statement.Marshal(b, m, deterministic)
}
func (m *Statement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statement.Merge(m, src)
}
func (m *Statement) XXX_Size() int {
	return xxx_messageInfo_Statement.Size(m)
}
func (m *Statement) XXX_DiscardUnknown() {
	xxx_messageInfo_Statement.DiscardUnknown(m)
}

var xxx_messageInfo_Statement proto.InternalMessageInfo

func (m *Statement) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Statement) GetSourceId() int64 {
	if m != nil {
		return m.SourceId
	}
	return 0
}

func (m *Statement) GetDistributorId() int64 {
	if m != nil {
		return m.DistributorId
	}
	return 0
}

func (m *Statement) GetMoney() float32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *Statement) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Statement) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Statement) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Statement) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type StatementResponse struct {
	Entity               *Statement   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Statement `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error       `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info        `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StatementResponse) Reset()         { *m = StatementResponse{} }
func (m *StatementResponse) String() string { return proto.CompactTextString(m) }
func (*StatementResponse) ProtoMessage()    {}
func (*StatementResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4430318ac3171d1a, []int{1}
}

func (m *StatementResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatementResponse.Unmarshal(m, b)
}
func (m *StatementResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatementResponse.Marshal(b, m, deterministic)
}
func (m *StatementResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatementResponse.Merge(m, src)
}
func (m *StatementResponse) XXX_Size() int {
	return xxx_messageInfo_StatementResponse.Size(m)
}
func (m *StatementResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatementResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatementResponse proto.InternalMessageInfo

func (m *StatementResponse) GetEntity() *Statement {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *StatementResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *StatementResponse) GetItems() []*Statement {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *StatementResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *StatementResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Statement)(nil), "geiqin.srv.dms.Statement")
	proto.RegisterType((*StatementResponse)(nil), "geiqin.srv.dms.StatementResponse")
}

func init() { proto.RegisterFile("statement.proto", fileDescriptor_4430318ac3171d1a) }

var fileDescriptor_4430318ac3171d1a = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xd1, 0xea, 0xd3, 0x30,
	0x14, 0xc6, 0x6d, 0xbb, 0xd6, 0x7f, 0xcf, 0x74, 0xea, 0xe1, 0xaf, 0xd4, 0x89, 0x50, 0x07, 0x42,
	0x41, 0xa8, 0x58, 0x9f, 0x60, 0xc2, 0x70, 0xbb, 0x93, 0xee, 0xc2, 0xcb, 0xd1, 0x35, 0x67, 0x5b,
	0xc0, 0x26, 0x35, 0xc9, 0x06, 0x7b, 0x5a, 0x9f, 0xc0, 0x27, 0xf0, 0x46, 0x9a, 0x74, 0x53, 0xa7,
	0x13, 0xc1, 0xbb, 0xe6, 0xfb, 0x7e, 0xe7, 0x6b, 0xce, 0x47, 0xe0, 0x81, 0x36, 0x95, 0xa1, 0x86,
	0x84, 0xc9, 0x5b, 0x25, 0x8d, 0xc4, 0xd1, 0x96, 0xf8, 0x67, 0x2e, 0x72, 0xad, 0x0e, 0x39, 0x6b,
	0xf4, 0xf8, 0x5e, 0x2d, 0x9b, 0x46, 0x0a, 0xe7, 0x4e, 0xbe, 0x78, 0x10, 0x2f, 0x4f, 0x13, 0x38,
	0x02, 0x9f, 0xb3, 0xc4, 0x4b, 0xbd, 0x2c, 0x28, 0x7d, 0xce, 0xf0, 0x19, 0xc4, 0x5a, 0xee, 0x55,
	0x4d, 0x2b, 0xce, 0x12, 0xdf, 0xca, 0x37, 0x4e, 0x58, 0x30, 0x7c, 0x09, 0x23, 0xc6, 0xb5, 0x51,
	0x7c, 0xbd, 0x37, 0x52, 0x75, 0x44, 0x60, 0x89, 0xfb, 0x3f, 0xa9, 0x0b, 0x86, 0xb7, 0x10, 0x36,
	0x52, 0xd0, 0x31, 0x19, 0xa4, 0x5e, 0xe6, 0x97, 0xee, 0x80, 0x4f, 0x20, 0x72, 0x41, 0x49, 0x98,
	0x7a, 0x59, 0x5c, 0xf6, 0x27, 0x44, 0x18, 0x98, 0x63, 0x4b, 0x49, 0x64, 0x55, 0xfb, 0x8d, 0xcf,
	0x01, 0x6a, 0x45, 0x95, 0x21, 0xb6, 0xaa, 0x4c, 0x72, 0xd7, 0x3a, 0x71, 0xaf, 0x4c, 0x4d, 0x67,
	0xef, 0x5b, 0x76, 0xb2, 0x6f, 0x9c, 0xdd, 0x2b, 0x53, 0x33, 0xf9, 0xe6, 0xc1, 0xa3, 0xf3, 0x86,
	0x25, 0xe9, 0x56, 0x0a, 0x4d, 0xf8, 0x06, 0x22, 0x12, 0x86, 0x9b, 0xa3, 0xdd, 0x76, 0x58, 0x3c,
	0xcd, 0x7f, 0xad, 0x29, 0xff, 0x31, 0xd2, 0x83, 0xf8, 0x0a, 0xc2, 0xb6, 0xda, 0x92, 0xb2, 0x45,
	0x0c, 0x8b, 0xc7, 0x97, 0x13, 0x1f, 0x3a, 0xb3, 0x74, 0x0c, 0xbe, 0x86, 0x90, 0x1b, 0x6a, 0x74,
	0x12, 0xa4, 0xc1, 0xdf, 0xe3, 0x1d, 0xd7, 0xa5, 0x93, 0x52, 0x52, 0xd9, 0x9a, 0xfe, 0x90, 0x3e,
	0xeb, 0xcc, 0xd2, 0x31, 0x98, 0xc1, 0x80, 0x8b, 0x8d, 0xb4, 0xdd, 0x0d, 0x8b, 0xdb, 0x4b, 0x76,
	0x21, 0x36, 0xb2, 0xb4, 0x44, 0xf1, 0xd5, 0x83, 0x87, 0xe7, 0x7f, 0x2d, 0x49, 0x1d, 0x78, 0x4d,
	0x38, 0x87, 0x68, 0x5e, 0x09, 0xf6, 0x89, 0xf0, 0xfa, 0xbd, 0xc6, 0x2f, 0xae, 0x5f, 0xb9, 0x2f,
	0x71, 0x72, 0x07, 0x67, 0x10, 0xbc, 0x27, 0xf3, 0xdf, 0x31, 0x73, 0x88, 0x96, 0x54, 0xa9, 0x7a,
	0xf7, 0x7b, 0xd2, 0xbb, 0x4a, 0xd3, 0xc7, 0x1d, 0x29, 0xfa, 0xa7, 0xa4, 0x75, 0x64, 0x9f, 0xf5,
	0xdb, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x79, 0x3d, 0xc0, 0x3a, 0x07, 0x03, 0x00, 0x00,
}