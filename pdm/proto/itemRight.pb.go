// Code generated by protoc-gen-go. DO NOT EDIT.
// source: itemRight.proto

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

type ItemRight struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	TypeValueId          int64    `protobuf:"varint,3,opt,name=type_value_id,json=typeValueId,proto3" json:"type_value_id,omitempty"`
	ItemId               int64    `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemRight) Reset()         { *m = ItemRight{} }
func (m *ItemRight) String() string { return proto.CompactTextString(m) }
func (*ItemRight) ProtoMessage()    {}
func (*ItemRight) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7691599d9360138, []int{0}
}

func (m *ItemRight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemRight.Unmarshal(m, b)
}
func (m *ItemRight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemRight.Marshal(b, m, deterministic)
}
func (m *ItemRight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemRight.Merge(m, src)
}
func (m *ItemRight) XXX_Size() int {
	return xxx_messageInfo_ItemRight.Size(m)
}
func (m *ItemRight) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemRight.DiscardUnknown(m)
}

var xxx_messageInfo_ItemRight proto.InternalMessageInfo

func (m *ItemRight) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemRight) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ItemRight) GetTypeValueId() int64 {
	if m != nil {
		return m.TypeValueId
	}
	return 0
}

func (m *ItemRight) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

type ItemRightResponse struct {
	Entity               *ItemRight   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*ItemRight `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error       `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info        `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ItemRightResponse) Reset()         { *m = ItemRightResponse{} }
func (m *ItemRightResponse) String() string { return proto.CompactTextString(m) }
func (*ItemRightResponse) ProtoMessage()    {}
func (*ItemRightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7691599d9360138, []int{1}
}

func (m *ItemRightResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemRightResponse.Unmarshal(m, b)
}
func (m *ItemRightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemRightResponse.Marshal(b, m, deterministic)
}
func (m *ItemRightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemRightResponse.Merge(m, src)
}
func (m *ItemRightResponse) XXX_Size() int {
	return xxx_messageInfo_ItemRightResponse.Size(m)
}
func (m *ItemRightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemRightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ItemRightResponse proto.InternalMessageInfo

func (m *ItemRightResponse) GetEntity() *ItemRight {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ItemRightResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ItemRightResponse) GetItems() []*ItemRight {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ItemRightResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ItemRightResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemRight)(nil), "geiqin.srv.pdm.ItemRight")
	proto.RegisterType((*ItemRightResponse)(nil), "geiqin.srv.pdm.ItemRightResponse")
}

func init() { proto.RegisterFile("itemRight.proto", fileDescriptor_e7691599d9360138) }

var fileDescriptor_e7691599d9360138 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4b, 0x4b, 0x03, 0x31,
	0x14, 0x85, 0x99, 0xa7, 0xf4, 0x8e, 0x56, 0xbc, 0x28, 0x46, 0x57, 0xc3, 0xac, 0x06, 0x84, 0x88,
	0xf5, 0x37, 0xb8, 0x98, 0x9d, 0x64, 0xe1, 0xb6, 0x54, 0x93, 0x8e, 0x81, 0x4e, 0x12, 0x33, 0xb1,
	0xd0, 0xdf, 0xee, 0x46, 0x6e, 0xc6, 0x07, 0x3e, 0xe8, 0x2a, 0xc9, 0x3d, 0xdf, 0x3d, 0xe7, 0x10,
	0x38, 0xd6, 0x41, 0x0d, 0x42, 0xf7, 0xcf, 0x81, 0x3b, 0x6f, 0x83, 0xc5, 0x79, 0xaf, 0xf4, 0x8b,
	0x36, 0x7c, 0xf4, 0x5b, 0xee, 0xe4, 0x70, 0x79, 0xf8, 0x64, 0x87, 0xc1, 0x9a, 0x49, 0x6d, 0x36,
	0x30, 0xeb, 0x3e, 0x17, 0x70, 0x0e, 0xa9, 0x96, 0x2c, 0xa9, 0x93, 0x36, 0x13, 0xa9, 0x96, 0x88,
	0x90, 0x87, 0x9d, 0x53, 0x2c, 0xad, 0x93, 0x76, 0x26, 0xe2, 0x1d, 0x1b, 0x38, 0xa2, 0x73, 0xb9,
	0x5d, 0x6d, 0x5e, 0xd5, 0x52, 0x4b, 0x96, 0x45, 0xbc, 0xa2, 0xe1, 0x03, 0xcd, 0x3a, 0x89, 0xe7,
	0x70, 0x40, 0x2d, 0x48, 0xcd, 0xa3, 0x5a, 0xd2, 0xb3, 0x93, 0xcd, 0x5b, 0x02, 0x27, 0x5f, 0x71,
	0x42, 0x8d, 0xce, 0x9a, 0x51, 0xe1, 0x0d, 0x94, 0xca, 0x04, 0x1d, 0x76, 0x31, 0xba, 0x5a, 0x5c,
	0xf0, 0x9f, 0x95, 0xf9, 0xf7, 0xca, 0x07, 0x88, 0x57, 0x50, 0xb8, 0x55, 0xaf, 0x7c, 0xac, 0x56,
	0x2d, 0xce, 0x7e, 0x6f, 0xdc, 0x93, 0x28, 0x26, 0x06, 0xaf, 0xa1, 0xa0, 0xfc, 0x91, 0x65, 0x75,
	0xb6, 0xdf, 0x7e, 0xe2, 0xc8, 0x5d, 0x79, 0x6f, 0x7d, 0x6c, 0xff, 0x8f, 0xfb, 0x1d, 0x89, 0x62,
	0x62, 0xb0, 0x85, 0x5c, 0x9b, 0xb5, 0x65, 0x45, 0x64, 0x4f, 0xff, 0x98, 0x9b, 0xb5, 0x15, 0x91,
	0x78, 0x2c, 0xe3, 0x97, 0xdf, 0xbe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x4f, 0x65, 0x91, 0xa3,
	0x01, 0x00, 0x00,
}
