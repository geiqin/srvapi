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
	ShipperId             int32    `protobuf:"varint,6,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	CreatedAt             string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt             string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
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

func (m *Setting) GetShipperId() int32 {
	if m != nil {
		return m.ShipperId
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
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x51, 0xab, 0xd3, 0x30,
	0x18, 0x86, 0x6d, 0x7b, 0xba, 0xad, 0x5f, 0x75, 0x42, 0xf0, 0x70, 0xe2, 0x40, 0x56, 0x76, 0x55,
	0x10, 0x2b, 0x54, 0xf0, 0x7e, 0xe8, 0x94, 0x5d, 0x08, 0xd2, 0xfd, 0x80, 0x52, 0xdb, 0x6f, 0xed,
	0x07, 0x36, 0xa9, 0x49, 0x1c, 0xee, 0x27, 0xf8, 0x83, 0xbd, 0x97, 0xa6, 0xd9, 0x85, 0xc3, 0x9d,
	0xcb, 0xbe, 0xcf, 0xd3, 0x37, 0xe1, 0x25, 0xf0, 0x4c, 0xa3, 0x31, 0x24, 0xda, 0x6c, 0x50, 0xd2,
	0x48, 0xb6, 0x6c, 0x91, 0x7e, 0x90, 0xc8, 0xb4, 0x3a, 0x65, 0xa6, 0xd7, 0xab, 0xa7, 0xb5, 0xec,
	0x7b, 0x29, 0x26, 0xba, 0xf9, 0xed, 0xc3, 0xfc, 0x30, 0xf9, 0x6c, 0x09, 0x3e, 0x35, 0xdc, 0x4b,
	0xbc, 0x34, 0x2c, 0x7c, 0x6a, 0xd8, 0x1a, 0x62, 0xd2, 0x65, 0x83, 0xdf, 0xe9, 0x84, 0xea, 0xcc,
	0xfd, 0xc4, 0x4b, 0x17, 0x05, 0x90, 0xfe, 0xe8, 0x12, 0xf6, 0x12, 0x16, 0xa4, 0xcb, 0x23, 0x9a,
	0xba, 0xe3, 0x81, 0xa5, 0x73, 0xd2, 0x9f, 0xc6, 0x4f, 0xf6, 0x0a, 0x80, 0x74, 0x89, 0xbf, 0x06,
	0x85, 0x5a, 0xf3, 0x3b, 0x0b, 0x23, 0xd2, 0xbb, 0x29, 0x60, 0xef, 0xe1, 0xc1, 0xb1, 0xb2, 0xee,
	0x2a, 0xd5, 0x92, 0x68, 0xcb, 0x1e, 0x4d, 0x27, 0x1b, 0x1e, 0xda, 0xf3, 0xef, 0x1d, 0xfe, 0xe0,
	0xe8, 0x17, 0x0b, 0xc7, 0x5a, 0xdd, 0xd1, 0x30, 0xa0, 0x2a, 0xa9, 0xe1, 0x33, 0xab, 0x46, 0x2e,
	0xd9, 0x5b, 0x5c, 0x2b, 0xac, 0x0c, 0x36, 0x65, 0x65, 0xf8, 0x3c, 0xf1, 0xd2, 0xa8, 0x88, 0x5c,
	0xb2, 0x35, 0x23, 0xfe, 0x39, 0x34, 0x17, 0xbc, 0x98, 0xb0, 0x4b, 0xb6, 0x66, 0xf3, 0xc7, 0x83,
	0xe7, 0x6e, 0x8b, 0x02, 0xf5, 0x20, 0x85, 0x46, 0xf6, 0x16, 0x66, 0x28, 0x0c, 0x99, 0xb3, 0xdd,
	0x25, 0xce, 0x1f, 0xb2, 0x7f, 0xe7, 0xcc, 0x2e, 0x3f, 0x38, 0x8d, 0xbd, 0x86, 0x70, 0xa8, 0x5a,
	0x54, 0x76, 0xae, 0x38, 0xbf, 0xbf, 0xf6, 0xbf, 0x8e, 0xb0, 0x98, 0x1c, 0xf6, 0x06, 0x42, 0x32,
	0xd8, 0x6b, 0x1e, 0x24, 0xc1, 0x63, 0xe5, 0x93, 0x35, 0x76, 0xa3, 0x52, 0x52, 0xd9, 0x3d, 0xff,
	0xd3, 0xbd, 0x1b, 0x61, 0x31, 0x39, 0x2c, 0x85, 0x3b, 0x12, 0x47, 0x69, 0xf7, 0x8c, 0xf3, 0x17,
	0xd7, 0xee, 0x5e, 0x1c, 0x65, 0x61, 0x8d, 0xfc, 0x00, 0x4b, 0x77, 0xd0, 0x01, 0xd5, 0x89, 0x6a,
	0x64, 0x5b, 0x08, 0x3e, 0xa3, 0x61, 0xb7, 0xee, 0xb3, 0x5a, 0xdf, 0xba, 0xa8, 0x9b, 0x6d, 0xf3,
	0xe4, 0xdb, 0xcc, 0xbe, 0xaf, 0x77, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x55, 0x99, 0x83,
	0x8e, 0x02, 0x00, 0x00,
}
