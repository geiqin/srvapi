// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deliveryRange.proto

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

type DeliveryRange struct {
	Id                   int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DeliveryId           int64         `protobuf:"varint,2,opt,name=delivery_id,json=deliveryId,proto3" json:"delivery_id,omitempty"`
	AreaName             string        `protobuf:"bytes,3,opt,name=area_name,json=areaName,proto3" json:"area_name,omitempty"`
	StartPrice           float32       `protobuf:"fixed32,4,opt,name=start_price,json=startPrice,proto3" json:"start_price,omitempty"`
	DeliveryPrice        float32       `protobuf:"fixed32,5,opt,name=delivery_price,json=deliveryPrice,proto3" json:"delivery_price,omitempty"`
	Method               int32         `protobuf:"varint,6,opt,name=method,proto3" json:"method,omitempty"`
	RegionRadius         float32       `protobuf:"fixed32,7,opt,name=region_radius,json=regionRadius,proto3" json:"region_radius,omitempty"`
	RegionData           string        `protobuf:"bytes,8,opt,name=region_data,json=regionData,proto3" json:"region_data,omitempty"`
	CreatedAt            string        `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string        `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Coordinate           []*Coordinate `protobuf:"bytes,11,rep,name=coordinate,proto3" json:"coordinate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DeliveryRange) Reset()         { *m = DeliveryRange{} }
func (m *DeliveryRange) String() string { return proto.CompactTextString(m) }
func (*DeliveryRange) ProtoMessage()    {}
func (*DeliveryRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6ee5a0fb8289955, []int{0}
}

func (m *DeliveryRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryRange.Unmarshal(m, b)
}
func (m *DeliveryRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryRange.Marshal(b, m, deterministic)
}
func (m *DeliveryRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryRange.Merge(m, src)
}
func (m *DeliveryRange) XXX_Size() int {
	return xxx_messageInfo_DeliveryRange.Size(m)
}
func (m *DeliveryRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryRange.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryRange proto.InternalMessageInfo

func (m *DeliveryRange) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeliveryRange) GetDeliveryId() int64 {
	if m != nil {
		return m.DeliveryId
	}
	return 0
}

func (m *DeliveryRange) GetAreaName() string {
	if m != nil {
		return m.AreaName
	}
	return ""
}

func (m *DeliveryRange) GetStartPrice() float32 {
	if m != nil {
		return m.StartPrice
	}
	return 0
}

func (m *DeliveryRange) GetDeliveryPrice() float32 {
	if m != nil {
		return m.DeliveryPrice
	}
	return 0
}

func (m *DeliveryRange) GetMethod() int32 {
	if m != nil {
		return m.Method
	}
	return 0
}

func (m *DeliveryRange) GetRegionRadius() float32 {
	if m != nil {
		return m.RegionRadius
	}
	return 0
}

func (m *DeliveryRange) GetRegionData() string {
	if m != nil {
		return m.RegionData
	}
	return ""
}

func (m *DeliveryRange) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *DeliveryRange) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *DeliveryRange) GetCoordinate() []*Coordinate {
	if m != nil {
		return m.Coordinate
	}
	return nil
}

type Coordinate struct {
	Lng                  string   `protobuf:"bytes,1,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat                  string   `protobuf:"bytes,2,opt,name=lat,proto3" json:"lat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinate) Reset()         { *m = Coordinate{} }
func (m *Coordinate) String() string { return proto.CompactTextString(m) }
func (*Coordinate) ProtoMessage()    {}
func (*Coordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6ee5a0fb8289955, []int{1}
}

func (m *Coordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinate.Unmarshal(m, b)
}
func (m *Coordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinate.Marshal(b, m, deterministic)
}
func (m *Coordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinate.Merge(m, src)
}
func (m *Coordinate) XXX_Size() int {
	return xxx_messageInfo_Coordinate.Size(m)
}
func (m *Coordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinate.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinate proto.InternalMessageInfo

func (m *Coordinate) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *Coordinate) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func init() {
	proto.RegisterType((*DeliveryRange)(nil), "geiqin.srv.tms.DeliveryRange")
	proto.RegisterType((*Coordinate)(nil), "geiqin.srv.tms.Coordinate")
}

func init() {
	proto.RegisterFile("deliveryRange.proto", fileDescriptor_c6ee5a0fb8289955)
}

var fileDescriptor_c6ee5a0fb8289955 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x62, 0x6b, 0x33, 0xb5, 0x45, 0x56, 0x90, 0x45, 0x11, 0x43, 0x45, 0xc8, 0x29,
	0x88, 0xde, 0xbc, 0x15, 0x7b, 0xf1, 0x22, 0xb2, 0x2f, 0x10, 0xc6, 0xee, 0x10, 0x17, 0x9a, 0x6c,
	0xdd, 0x4c, 0x0b, 0x3e, 0xbc, 0x20, 0xbb, 0x49, 0x6a, 0xbd, 0x25, 0xdf, 0xff, 0xcd, 0x30, 0xfc,
	0x0b, 0x17, 0x9a, 0x36, 0x66, 0x4f, 0xee, 0x5b, 0x61, 0x53, 0x51, 0xb1, 0x75, 0x96, 0xad, 0x98,
	0x57, 0x64, 0xbe, 0x4c, 0x53, 0xb4, 0x6e, 0x5f, 0x70, 0xdd, 0x2e, 0x7e, 0x62, 0x98, 0xad, 0x8e,
	0x3d, 0x31, 0x87, 0xd8, 0x68, 0x19, 0x65, 0x51, 0x3e, 0x52, 0xb1, 0xd1, 0xe2, 0x16, 0xa6, 0xc3,
	0xa2, 0xd2, 0x68, 0x19, 0x67, 0x51, 0x9e, 0x28, 0x18, 0xd0, 0xab, 0x16, 0xd7, 0x90, 0xa2, 0x23,
	0x2c, 0x1b, 0xac, 0x49, 0x26, 0x59, 0x94, 0xa7, 0x6a, 0xe2, 0xc1, 0x1b, 0xd6, 0xe4, 0xa7, 0x5b,
	0x46, 0xc7, 0xe5, 0xd6, 0x99, 0x35, 0xc9, 0x93, 0x2c, 0xca, 0x63, 0x05, 0x01, 0xbd, 0x7b, 0x22,
	0xee, 0x61, 0x7e, 0x58, 0xdf, 0x39, 0xa3, 0xe0, 0xcc, 0x06, 0xda, 0x69, 0x97, 0x30, 0xae, 0x89,
	0x3f, 0xad, 0x96, 0xe3, 0x70, 0x59, 0xff, 0x27, 0xee, 0x60, 0xe6, 0xa8, 0x32, 0xb6, 0x29, 0x1d,
	0x6a, 0xb3, 0x6b, 0xe5, 0x69, 0x98, 0x3e, 0xeb, 0xa0, 0x0a, 0xcc, 0x1f, 0xd1, 0x4b, 0x1a, 0x19,
	0xe5, 0x24, 0xdc, 0x08, 0x1d, 0x5a, 0x21, 0xa3, 0xb8, 0x01, 0x58, 0x3b, 0x42, 0x26, 0x5d, 0x22,
	0xcb, 0x34, 0xe4, 0x69, 0x4f, 0x96, 0xec, 0xe3, 0xdd, 0x56, 0x0f, 0x31, 0x74, 0x71, 0x4f, 0x96,
	0x2c, 0x9e, 0x01, 0xd6, 0xd6, 0x3a, 0x6d, 0x1a, 0x64, 0x92, 0xd3, 0x2c, 0xc9, 0xa7, 0x8f, 0x57,
	0xc5, 0xff, 0xa2, 0x8b, 0x97, 0x83, 0xa1, 0x8e, 0xec, 0xc5, 0x03, 0xc0, 0x5f, 0x22, 0xce, 0x21,
	0xd9, 0x34, 0x55, 0x28, 0x3f, 0x55, 0xfe, 0x33, 0x10, 0xe4, 0xd0, 0xba, 0x27, 0xc8, 0x1f, 0xe3,
	0xf0, 0x90, 0x4f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xf9, 0x51, 0xdc, 0xdf, 0x01, 0x00,
	0x00,
}