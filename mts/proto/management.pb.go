// Code generated by protoc-gen-go. DO NOT EDIT.
// source: management.proto

package geiqin_srv_mts

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

type Management struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StoreId              int64    `protobuf:"varint,3,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	RoleId               int64    `protobuf:"varint,4,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Store                *Store   `protobuf:"bytes,5,opt,name=store,proto3" json:"store,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Management) Reset()         { *m = Management{} }
func (m *Management) String() string { return proto.CompactTextString(m) }
func (*Management) ProtoMessage()    {}
func (*Management) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc174f991dc0a25, []int{0}
}

func (m *Management) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Management.Unmarshal(m, b)
}
func (m *Management) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Management.Marshal(b, m, deterministic)
}
func (m *Management) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Management.Merge(m, src)
}
func (m *Management) XXX_Size() int {
	return xxx_messageInfo_Management.Size(m)
}
func (m *Management) XXX_DiscardUnknown() {
	xxx_messageInfo_Management.DiscardUnknown(m)
}

var xxx_messageInfo_Management proto.InternalMessageInfo

func (m *Management) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Management) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Management) GetStoreId() int64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *Management) GetRoleId() int64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *Management) GetStore() *Store {
	if m != nil {
		return m.Store
	}
	return nil
}

//
type ManagementResponse struct {
	Entity               *Management   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager        `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Management `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error        `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info         `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ManagementResponse) Reset()         { *m = ManagementResponse{} }
func (m *ManagementResponse) String() string { return proto.CompactTextString(m) }
func (*ManagementResponse) ProtoMessage()    {}
func (*ManagementResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc174f991dc0a25, []int{1}
}

func (m *ManagementResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagementResponse.Unmarshal(m, b)
}
func (m *ManagementResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagementResponse.Marshal(b, m, deterministic)
}
func (m *ManagementResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagementResponse.Merge(m, src)
}
func (m *ManagementResponse) XXX_Size() int {
	return xxx_messageInfo_ManagementResponse.Size(m)
}
func (m *ManagementResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagementResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ManagementResponse proto.InternalMessageInfo

func (m *ManagementResponse) GetEntity() *Management {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ManagementResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ManagementResponse) GetItems() []*Management {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ManagementResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ManagementResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Management)(nil), "geiqin.srv.mts.Management")
	proto.RegisterType((*ManagementResponse)(nil), "geiqin.srv.mts.ManagementResponse")
}

func init() {
	proto.RegisterFile("management.proto", fileDescriptor_edc174f991dc0a25)
}

var fileDescriptor_edc174f991dc0a25 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xd1, 0x4a, 0xeb, 0x30,
	0x18, 0xc7, 0x4f, 0xd7, 0xad, 0x3b, 0x7c, 0x3b, 0x8c, 0x63, 0x50, 0xec, 0x7a, 0x35, 0x7a, 0x35,
	0x10, 0x8a, 0xd4, 0x27, 0x50, 0x90, 0x51, 0x41, 0x90, 0x0c, 0xf1, 0x52, 0xea, 0xf2, 0x6d, 0x0b,
	0x98, 0xa4, 0x26, 0x71, 0xe0, 0x1b, 0xf8, 0x04, 0x3e, 0xae, 0x48, 0x92, 0xe1, 0x74, 0x8c, 0xe9,
	0xc5, 0x2e, 0xbf, 0xfe, 0x7f, 0xdf, 0x2f, 0xff, 0x04, 0x0a, 0xff, 0x45, 0x2d, 0xeb, 0x39, 0x0a,
	0x94, 0xb6, 0x68, 0xb4, 0xb2, 0x8a, 0xf4, 0xe7, 0xc8, 0x9f, 0xb8, 0x2c, 0x8c, 0x5e, 0x16, 0xc2,
	0x9a, 0xec, 0xdf, 0x54, 0x09, 0xa1, 0x64, 0x48, 0xb3, 0x9e, 0xb1, 0x4a, 0x63, 0x18, 0xf2, 0xb7,
	0x08, 0xe0, 0xfa, 0x73, 0x9f, 0xf4, 0xa1, 0xc5, 0x59, 0x1a, 0x0d, 0xa3, 0x51, 0x4c, 0x5b, 0x9c,
	0x91, 0x63, 0xe8, 0x3e, 0x1b, 0xd4, 0xf7, 0x9c, 0xa5, 0x2d, 0xff, 0x31, 0x71, 0x63, 0xc5, 0xc8,
	0x00, 0xfe, 0x7a, 0x8d, 0x4b, 0x62, 0x9f, 0x74, 0xfd, 0x5c, 0xf9, 0x1d, 0xad, 0x1e, 0x7d, 0xd2,
	0x0e, 0x3b, 0x6e, 0xac, 0x18, 0x39, 0x81, 0x8e, 0x67, 0xd2, 0xce, 0x30, 0x1a, 0xf5, 0xca, 0xa3,
	0xe2, 0x7b, 0xcd, 0x62, 0xe2, 0x42, 0x1a, 0x98, 0xfc, 0x3d, 0x02, 0xb2, 0x2e, 0x46, 0xd1, 0x34,
	0x4a, 0x1a, 0x24, 0x25, 0x24, 0x28, 0x2d, 0xb7, 0x2f, 0xbe, 0x64, 0xaf, 0xcc, 0x36, 0x25, 0x5f,
	0x76, 0x56, 0xa4, 0x3b, 0xb7, 0xa9, 0xe7, 0xa8, 0xfd, 0x15, 0xb6, 0x9c, 0x7b, 0xe3, 0x42, 0x1a,
	0x18, 0x72, 0x0a, 0x1d, 0x6e, 0x51, 0x98, 0x34, 0x1e, 0xc6, 0x3f, 0xf8, 0x03, 0xe8, 0xf4, 0xa8,
	0xb5, 0xd2, 0xfe, 0xb6, 0x5b, 0xf4, 0x97, 0x2e, 0xa4, 0x81, 0x21, 0x23, 0x68, 0x73, 0x39, 0x53,
	0xab, 0x27, 0x38, 0xdc, 0x64, 0x2b, 0x39, 0x53, 0xd4, 0x13, 0xe5, 0x6b, 0x0c, 0x07, 0xeb, 0xc3,
	0x26, 0xa8, 0x97, 0x7c, 0x8a, 0x64, 0x0c, 0xf1, 0x39, 0x63, 0x64, 0x47, 0xad, 0x2c, 0xdf, 0x51,
	0x79, 0xf5, 0x8c, 0xf9, 0x1f, 0x72, 0x05, 0xc9, 0x6d, 0xc3, 0x6a, 0x8b, 0xfb, 0x71, 0x51, 0x14,
	0x6a, 0xb9, 0x0f, 0xd7, 0x18, 0xe2, 0x31, 0xda, 0x3d, 0x88, 0x2a, 0x48, 0x26, 0x58, 0xeb, 0xe9,
	0x82, 0x0c, 0x36, 0xf9, 0x8b, 0xda, 0xe0, 0xdd, 0x02, 0x35, 0xfe, 0x4e, 0xf5, 0x90, 0xf8, 0x7f,
	0xe5, 0xec, 0x23, 0x00, 0x00, 0xff, 0xff, 0x22, 0xd3, 0xb4, 0xca, 0x6a, 0x03, 0x00, 0x00,
}
