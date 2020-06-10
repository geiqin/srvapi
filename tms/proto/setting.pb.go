// Code generated by protoc-gen-go. DO NOT EDIT.
// source: setting.proto

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

type Setting struct {
	Id                    int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsDelivery            bool     `protobuf:"varint,2,opt,name=is_delivery,json=isDelivery,proto3" json:"is_delivery,omitempty"`
	IsFetch               bool     `protobuf:"varint,3,opt,name=is_fetch,json=isFetch,proto3" json:"is_fetch,omitempty"`
	IsExpress             bool     `protobuf:"varint,4,opt,name=is_express,json=isExpress,proto3" json:"is_express,omitempty"`
	ExpressChargingMethod int32    `protobuf:"varint,5,opt,name=express_charging_method,json=expressChargingMethod,proto3" json:"express_charging_method,omitempty"`
	CreatedAt             string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt             string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Setting) Reset()         { *m = Setting{} }
func (m *Setting) String() string { return proto.CompactTextString(m) }
func (*Setting) ProtoMessage()    {}
func (*Setting) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fa78c2ae253ed30, []int{0}
}

func (m *Setting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Setting.Unmarshal(m, b)
}
func (m *Setting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Setting.Marshal(b, m, deterministic)
}
func (m *Setting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Setting.Merge(m, src)
}
func (m *Setting) XXX_Size() int {
	return xxx_messageInfo_Setting.Size(m)
}
func (m *Setting) XXX_DiscardUnknown() {
	xxx_messageInfo_Setting.DiscardUnknown(m)
}

var xxx_messageInfo_Setting proto.InternalMessageInfo

func (m *Setting) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Setting) GetIsDelivery() bool {
	if m != nil {
		return m.IsDelivery
	}
	return false
}

func (m *Setting) GetIsFetch() bool {
	if m != nil {
		return m.IsFetch
	}
	return false
}

func (m *Setting) GetIsExpress() bool {
	if m != nil {
		return m.IsExpress
	}
	return false
}

func (m *Setting) GetExpressChargingMethod() int32 {
	if m != nil {
		return m.ExpressChargingMethod
	}
	return 0
}

func (m *Setting) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Setting) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type SettingResponse struct {
	Entity               *Setting   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Setting `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SettingResponse) Reset()         { *m = SettingResponse{} }
func (m *SettingResponse) String() string { return proto.CompactTextString(m) }
func (*SettingResponse) ProtoMessage()    {}
func (*SettingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fa78c2ae253ed30, []int{1}
}

func (m *SettingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingResponse.Unmarshal(m, b)
}
func (m *SettingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingResponse.Marshal(b, m, deterministic)
}
func (m *SettingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingResponse.Merge(m, src)
}
func (m *SettingResponse) XXX_Size() int {
	return xxx_messageInfo_SettingResponse.Size(m)
}
func (m *SettingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SettingResponse proto.InternalMessageInfo

func (m *SettingResponse) GetEntity() *Setting {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *SettingResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *SettingResponse) GetItems() []*Setting {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SettingResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *SettingResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Setting)(nil), "geiqin.srv.tms.Setting")
	proto.RegisterType((*SettingResponse)(nil), "geiqin.srv.tms.SettingResponse")
}

func init() {
	proto.RegisterFile("setting.proto", fileDescriptor_5fa78c2ae253ed30)
}

var fileDescriptor_5fa78c2ae253ed30 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x86, 0x2b, 0xcb, 0xf2, 0xc7, 0xa8, 0x75, 0x61, 0xa9, 0xb1, 0x6a, 0x28, 0x16, 0x3e, 0x09,
	0x4a, 0x55, 0x50, 0xa1, 0x77, 0x93, 0x38, 0x21, 0x87, 0x40, 0x90, 0x7f, 0x80, 0x50, 0xa4, 0xb1,
	0x3c, 0x10, 0xed, 0x2a, 0xbb, 0x1b, 0x13, 0xff, 0xe5, 0xdc, 0x73, 0x0f, 0x5a, 0xad, 0x0f, 0x31,
	0x71, 0x8e, 0x9a, 0xe7, 0xd1, 0xbb, 0x3b, 0x2f, 0x0b, 0xdf, 0x14, 0x6a, 0x4d, 0xbc, 0x8a, 0x1b,
	0x29, 0xb4, 0x60, 0x93, 0x0a, 0xe9, 0x91, 0x78, 0xac, 0xe4, 0x3e, 0xd6, 0xb5, 0x9a, 0x7f, 0x2d,
	0x44, 0x5d, 0x0b, 0xde, 0xd1, 0xe5, 0x8b, 0x03, 0xc3, 0x4d, 0xe7, 0xb3, 0x09, 0xf4, 0xa8, 0x0c,
	0x9c, 0xd0, 0x89, 0xbc, 0xb4, 0x47, 0x25, 0x5b, 0x80, 0x4f, 0x2a, 0x2b, 0xf1, 0x81, 0xf6, 0x28,
	0x0f, 0x41, 0x2f, 0x74, 0xa2, 0x51, 0x0a, 0xa4, 0x2e, 0xed, 0x84, 0xfd, 0x84, 0x11, 0xa9, 0x6c,
	0x8b, 0xba, 0xd8, 0x05, 0xae, 0xa1, 0x43, 0x52, 0x57, 0xed, 0x27, 0xfb, 0x05, 0x40, 0x2a, 0xc3,
	0xe7, 0x46, 0xa2, 0x52, 0x41, 0xdf, 0xc0, 0x31, 0xa9, 0x75, 0x37, 0x60, 0xff, 0x61, 0x66, 0x59,
	0x56, 0xec, 0x72, 0x59, 0x11, 0xaf, 0xb2, 0x1a, 0xf5, 0x4e, 0x94, 0x81, 0x67, 0xce, 0x9f, 0x5a,
	0x7c, 0x61, 0xe9, 0xad, 0x81, 0x6d, 0x6c, 0x21, 0x31, 0xd7, 0x58, 0x66, 0xb9, 0x0e, 0x06, 0xa1,
	0x13, 0x8d, 0xd3, 0xb1, 0x9d, 0xac, 0x74, 0x8b, 0x9f, 0x9a, 0xf2, 0x88, 0x87, 0x1d, 0xb6, 0x93,
	0x95, 0x5e, 0xbe, 0x3a, 0xf0, 0xdd, 0x2e, 0x9b, 0xa2, 0x6a, 0x04, 0x57, 0xc8, 0xfe, 0xc2, 0x00,
	0xb9, 0x26, 0x7d, 0x30, 0x8b, 0xfb, 0xc9, 0x2c, 0x7e, 0xdf, 0x57, 0x7c, 0xfc, 0xc1, 0x6a, 0xec,
	0x37, 0x78, 0x4d, 0x5e, 0xa1, 0x34, 0x7d, 0xf8, 0xc9, 0xf4, 0xd4, 0xbf, 0x6b, 0x61, 0xda, 0x39,
	0xec, 0x0f, 0x78, 0xa4, 0xb1, 0x56, 0x81, 0x1b, 0xba, 0x9f, 0x85, 0x77, 0x56, 0x9b, 0x8d, 0x52,
	0x0a, 0x69, 0x0a, 0xfb, 0x20, 0x7b, 0xdd, 0xc2, 0xb4, 0x73, 0x58, 0x04, 0x7d, 0xe2, 0x5b, 0x61,
	0x0a, 0xf3, 0x93, 0x1f, 0xa7, 0xee, 0x0d, 0xdf, 0x8a, 0xd4, 0x18, 0xc9, 0x06, 0x26, 0xf6, 0xa0,
	0x0d, 0xca, 0x3d, 0x15, 0xc8, 0x56, 0xe0, 0x5e, 0xa3, 0x66, 0xe7, 0xee, 0x33, 0x5f, 0x9c, 0xbb,
	0xa8, 0xad, 0x6d, 0xf9, 0xe5, 0x7e, 0x60, 0x1e, 0xd0, 0xbf, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb8, 0xa3, 0xa8, 0x87, 0x6f, 0x02, 0x00, 0x00,
}
