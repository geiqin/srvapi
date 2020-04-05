// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supplierContact.proto

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

type SupplierContact struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SupplierId           int64    `protobuf:"varint,2,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Mobile               string   `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Tel                  string   `protobuf:"bytes,5,opt,name=tel,proto3" json:"tel,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Weixin               string   `protobuf:"bytes,7,opt,name=weixin,proto3" json:"weixin,omitempty"`
	Qq                   string   `protobuf:"bytes,8,opt,name=qq,proto3" json:"qq,omitempty"`
	Memo                 string   `protobuf:"bytes,9,opt,name=memo,proto3" json:"memo,omitempty"`
	Disabled             bool     `protobuf:"varint,10,opt,name=disabled,proto3" json:"disabled,omitempty"`
	CreatedAt            string   `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupplierContact) Reset()         { *m = SupplierContact{} }
func (m *SupplierContact) String() string { return proto.CompactTextString(m) }
func (*SupplierContact) ProtoMessage()    {}
func (*SupplierContact) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc715dceb827fce9, []int{0}
}

func (m *SupplierContact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupplierContact.Unmarshal(m, b)
}
func (m *SupplierContact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupplierContact.Marshal(b, m, deterministic)
}
func (m *SupplierContact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplierContact.Merge(m, src)
}
func (m *SupplierContact) XXX_Size() int {
	return xxx_messageInfo_SupplierContact.Size(m)
}
func (m *SupplierContact) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplierContact.DiscardUnknown(m)
}

var xxx_messageInfo_SupplierContact proto.InternalMessageInfo

func (m *SupplierContact) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SupplierContact) GetSupplierId() int64 {
	if m != nil {
		return m.SupplierId
	}
	return 0
}

func (m *SupplierContact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SupplierContact) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *SupplierContact) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

func (m *SupplierContact) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SupplierContact) GetWeixin() string {
	if m != nil {
		return m.Weixin
	}
	return ""
}

func (m *SupplierContact) GetQq() string {
	if m != nil {
		return m.Qq
	}
	return ""
}

func (m *SupplierContact) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *SupplierContact) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *SupplierContact) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *SupplierContact) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type SupplierContactResponse struct {
	Entity               *SupplierContact   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager             `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*SupplierContact `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error             `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info              `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SupplierContactResponse) Reset()         { *m = SupplierContactResponse{} }
func (m *SupplierContactResponse) String() string { return proto.CompactTextString(m) }
func (*SupplierContactResponse) ProtoMessage()    {}
func (*SupplierContactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc715dceb827fce9, []int{1}
}

func (m *SupplierContactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupplierContactResponse.Unmarshal(m, b)
}
func (m *SupplierContactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupplierContactResponse.Marshal(b, m, deterministic)
}
func (m *SupplierContactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplierContactResponse.Merge(m, src)
}
func (m *SupplierContactResponse) XXX_Size() int {
	return xxx_messageInfo_SupplierContactResponse.Size(m)
}
func (m *SupplierContactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplierContactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SupplierContactResponse proto.InternalMessageInfo

func (m *SupplierContactResponse) GetEntity() *SupplierContact {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *SupplierContactResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *SupplierContactResponse) GetItems() []*SupplierContact {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SupplierContactResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *SupplierContactResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*SupplierContact)(nil), "geiqin.srv.wms.SupplierContact")
	proto.RegisterType((*SupplierContactResponse)(nil), "geiqin.srv.wms.SupplierContactResponse")
}

func init() {
	proto.RegisterFile("supplierContact.proto", fileDescriptor_dc715dceb827fce9)
}

var fileDescriptor_dc715dceb827fce9 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x75, 0x93, 0xdd, 0xb8, 0x7b, 0x53, 0xaa, 0x5c, 0xda, 0x3a, 0x2e, 0x48, 0x97, 0x7d, 0x71,
	0x41, 0xc8, 0x43, 0x44, 0x7c, 0xae, 0x55, 0xa4, 0x2f, 0xa2, 0x59, 0xc5, 0xc7, 0x32, 0x9b, 0xdc,
	0xb6, 0x03, 0x99, 0x99, 0x64, 0x32, 0x6d, 0xf5, 0x23, 0xfc, 0x13, 0x7f, 0xc1, 0x7f, 0x93, 0x99,
	0xc9, 0x0a, 0x0d, 0x42, 0x15, 0xf6, 0x6d, 0xee, 0x39, 0xf7, 0x9c, 0x33, 0x39, 0x03, 0x81, 0xc3,
	0xee, 0xba, 0x69, 0x6a, 0x41, 0xe6, 0x54, 0x2b, 0xcb, 0x4b, 0x9b, 0x35, 0x46, 0x5b, 0x8d, 0xfb,
	0x97, 0x24, 0x5a, 0xa1, 0xb2, 0xce, 0xdc, 0x64, 0xb7, 0xb2, 0x9b, 0xef, 0x95, 0x5a, 0x4a, 0xad,
	0x02, 0xbb, 0xfc, 0x19, 0xc1, 0xa3, 0xf5, 0x5d, 0x1d, 0xee, 0x43, 0x24, 0x2a, 0x36, 0x5a, 0x8c,
	0x56, 0x71, 0x11, 0x89, 0x0a, 0x8f, 0x21, 0xdd, 0x5a, 0x9f, 0x8b, 0x8a, 0x45, 0x9e, 0x80, 0x2d,
	0x74, 0x56, 0x21, 0xc2, 0x58, 0x71, 0x49, 0x2c, 0x5e, 0x8c, 0x56, 0xb3, 0xc2, 0x9f, 0xf1, 0x08,
	0x12, 0xa9, 0x37, 0xa2, 0x26, 0x36, 0xf6, 0x68, 0x3f, 0xe1, 0x63, 0x88, 0x2d, 0xd5, 0x6c, 0xe2,
	0x41, 0x77, 0xc4, 0x03, 0x98, 0x90, 0xe4, 0xa2, 0x66, 0x89, 0xc7, 0xc2, 0xe0, 0xf4, 0xb7, 0x24,
	0xbe, 0x09, 0xc5, 0x1e, 0x06, 0x7d, 0x98, 0xdc, 0xe5, 0xda, 0x96, 0x4d, 0x3d, 0x16, 0xb5, 0xad,
	0xcb, 0x96, 0x24, 0x35, 0x9b, 0x85, 0x6c, 0x77, 0xc6, 0x39, 0x4c, 0x2b, 0xd1, 0xf1, 0x4d, 0x4d,
	0x15, 0x83, 0xc5, 0x68, 0x35, 0x2d, 0xfe, 0xcc, 0xf8, 0x0c, 0xa0, 0x34, 0xc4, 0x2d, 0x55, 0xe7,
	0xdc, 0xb2, 0xd4, 0xab, 0x66, 0x3d, 0x72, 0x62, 0x1d, 0x7d, 0xdd, 0x54, 0x5b, 0x7a, 0x2f, 0xd0,
	0x3d, 0x72, 0x62, 0x97, 0x3f, 0x22, 0x78, 0x32, 0xa8, 0xab, 0xa0, 0xae, 0xd1, 0xaa, 0x23, 0x7c,
	0x0d, 0x09, 0x29, 0x2b, 0xec, 0x77, 0x5f, 0x5d, 0x9a, 0x1f, 0x67, 0x77, 0x9b, 0xcf, 0x86, 0xc2,
	0x7e, 0x1d, 0x5f, 0xc0, 0xa4, 0xe1, 0x97, 0x64, 0x7c, 0xb3, 0x69, 0x7e, 0x38, 0xd4, 0x7d, 0x74,
	0x64, 0x11, 0x76, 0xf0, 0x15, 0x4c, 0x84, 0x25, 0xd9, 0xb1, 0x78, 0x11, 0xff, 0x4b, 0x48, 0xd8,
	0x76, 0x19, 0x64, 0x8c, 0x36, 0xfe, 0x35, 0xfe, 0x92, 0xf1, 0xce, 0x91, 0x45, 0xd8, 0xc1, 0x15,
	0x8c, 0x85, 0xba, 0xd0, 0xfe, 0x91, 0xd2, 0xfc, 0x60, 0xb8, 0x7b, 0xa6, 0x2e, 0x74, 0xe1, 0x37,
	0xf2, 0x5f, 0x31, 0x1c, 0x0d, 0x12, 0xd7, 0x64, 0x6e, 0x44, 0x49, 0xf8, 0x19, 0x92, 0x53, 0x5f,
	0x2b, 0xde, 0x77, 0xc7, 0xf9, 0xf3, 0xfb, 0x3e, 0xa2, 0xaf, 0x78, 0xf9, 0xc0, 0xb9, 0x7e, 0xf1,
	0xaf, 0xb1, 0x6b, 0xd7, 0xb7, 0x54, 0xd3, 0x8e, 0x5d, 0x3f, 0x41, 0xfc, 0x9e, 0xec, 0x4e, 0x2d,
	0x3f, 0x40, 0xb2, 0x26, 0x6e, 0xca, 0x2b, 0x7c, 0x3a, 0x14, 0xbd, 0xe1, 0x1d, 0x7d, 0xbd, 0x22,
	0x43, 0xff, 0xe1, 0xb7, 0x49, 0xfc, 0x5f, 0xe0, 0xe5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f,
	0x43, 0xf8, 0xe0, 0x3c, 0x04, 0x00, 0x00,
}
