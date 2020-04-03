// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderAddress.proto

package geiqin_srv_ord_private

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

type OrderAddress struct {
	Id                   int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                 string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	AreaId               int64     `protobuf:"varint,4,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	Addr                 string    `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr"`
	Zip                  string    `protobuf:"bytes,6,opt,name=zip,proto3" json:"zip"`
	Phone                string    `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone"`
	Mobile               string    `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile"`
	Email                string    `protobuf:"bytes,9,opt,name=email,proto3" json:"email"`
	DeliveryTime         string    `protobuf:"bytes,10,opt,name=delivery_time,json=deliveryTime,proto3" json:"delivery_time"`
	Confirmed            bool      `protobuf:"varint,11,opt,name=confirmed,proto3" json:"confirmed"`
	OrderId              int64     `protobuf:"varint,12,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	AddressId            int64     `protobuf:"varint,13,opt,name=address_id,json=addressId,proto3" json:"address_id"`
	CreatedAt            string    `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string    `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Area                 *AreaInfo `protobuf:"bytes,16,opt,name=area,proto3" json:"area,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OrderAddress) Reset()         { *m = OrderAddress{} }
func (m *OrderAddress) String() string { return proto.CompactTextString(m) }
func (*OrderAddress) ProtoMessage()    {}
func (*OrderAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd915234d95bb37, []int{0}
}

func (m *OrderAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderAddress.Unmarshal(m, b)
}
func (m *OrderAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderAddress.Marshal(b, m, deterministic)
}
func (m *OrderAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderAddress.Merge(m, src)
}
func (m *OrderAddress) XXX_Size() int {
	return xxx_messageInfo_OrderAddress.Size(m)
}
func (m *OrderAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderAddress.DiscardUnknown(m)
}

var xxx_messageInfo_OrderAddress proto.InternalMessageInfo

func (m *OrderAddress) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderAddress) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrderAddress) GetAreaId() int64 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OrderAddress) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *OrderAddress) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *OrderAddress) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *OrderAddress) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *OrderAddress) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *OrderAddress) GetDeliveryTime() string {
	if m != nil {
		return m.DeliveryTime
	}
	return ""
}

func (m *OrderAddress) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *OrderAddress) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderAddress) GetAddressId() int64 {
	if m != nil {
		return m.AddressId
	}
	return 0
}

func (m *OrderAddress) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *OrderAddress) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *OrderAddress) GetArea() *AreaInfo {
	if m != nil {
		return m.Area
	}
	return nil
}

type OrderAddressResponse struct {
	Entity               *OrderAddress   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager          `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*OrderAddress `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error          `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info           `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OrderAddressResponse) Reset()         { *m = OrderAddressResponse{} }
func (m *OrderAddressResponse) String() string { return proto.CompactTextString(m) }
func (*OrderAddressResponse) ProtoMessage()    {}
func (*OrderAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd915234d95bb37, []int{1}
}

func (m *OrderAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderAddressResponse.Unmarshal(m, b)
}
func (m *OrderAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderAddressResponse.Marshal(b, m, deterministic)
}
func (m *OrderAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderAddressResponse.Merge(m, src)
}
func (m *OrderAddressResponse) XXX_Size() int {
	return xxx_messageInfo_OrderAddressResponse.Size(m)
}
func (m *OrderAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderAddressResponse proto.InternalMessageInfo

func (m *OrderAddressResponse) GetEntity() *OrderAddress {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *OrderAddressResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *OrderAddressResponse) GetItems() []*OrderAddress {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *OrderAddressResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *OrderAddressResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderAddress)(nil), "geiqin.srv.ord.private.OrderAddress")
	proto.RegisterType((*OrderAddressResponse)(nil), "geiqin.srv.ord.private.OrderAddressResponse")
}

func init() { proto.RegisterFile("orderAddress.proto", fileDescriptor_1cd915234d95bb37) }

var fileDescriptor_1cd915234d95bb37 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4f, 0x8b, 0x13, 0x4f,
	0x10, 0xfd, 0x25, 0x93, 0xcc, 0x26, 0x95, 0x6c, 0x7e, 0x4b, 0xb9, 0xac, 0x6d, 0xd8, 0x85, 0x31,
	0x7a, 0xc8, 0x41, 0x06, 0xc9, 0x7a, 0x12, 0x2f, 0xf1, 0x0f, 0xb2, 0x27, 0x65, 0x56, 0xf1, 0x20,
	0x18, 0x7a, 0xd3, 0x95, 0x4d, 0x43, 0x66, 0x7a, 0xec, 0x69, 0x03, 0xeb, 0x07, 0xf2, 0xe4, 0x77,
	0xf3, 0x2b, 0x48, 0x57, 0xcf, 0x62, 0x04, 0x23, 0xf1, 0x90, 0x5b, 0xd7, 0xab, 0xf7, 0x5e, 0x4d,
	0xd7, 0x6b, 0x06, 0xd0, 0x58, 0x45, 0x76, 0xaa, 0x94, 0xa5, 0xaa, 0x4a, 0x4b, 0x6b, 0x9c, 0xc1,
	0x93, 0x6b, 0xd2, 0x9f, 0x75, 0x91, 0x56, 0x76, 0x9d, 0x1a, 0xab, 0xd2, 0xd2, 0xea, 0xb5, 0x74,
	0x34, 0xec, 0xcf, 0x4d, 0x9e, 0x9b, 0x22, 0xb0, 0x86, 0x03, 0x69, 0x49, 0x5e, 0x14, 0x0b, 0x13,
	0xea, 0xd1, 0xb7, 0x08, 0xfa, 0x6f, 0x36, 0xcc, 0x70, 0x00, 0x4d, 0xad, 0x44, 0x23, 0x69, 0x8c,
	0xa3, 0xac, 0xa9, 0x15, 0x22, 0xb4, 0x0a, 0x99, 0x93, 0x88, 0x92, 0xc6, 0xb8, 0x9b, 0xf1, 0x19,
	0xef, 0xc2, 0x81, 0xb7, 0x99, 0x69, 0x25, 0x5a, 0x4c, 0x8c, 0xd9, 0x95, 0xc9, 0x52, 0x29, 0x2b,
	0xda, 0x81, 0xec, 0xcf, 0x78, 0x04, 0xd1, 0x57, 0x5d, 0x8a, 0x98, 0x21, 0x7f, 0xc4, 0x63, 0x68,
	0x97, 0x4b, 0x53, 0x90, 0x38, 0x60, 0x2c, 0x14, 0x78, 0x02, 0x71, 0x6e, 0xae, 0xf4, 0x8a, 0x44,
	0x87, 0xe1, 0xba, 0xf2, 0x6c, 0xca, 0xa5, 0x5e, 0x89, 0x6e, 0x60, 0x73, 0x81, 0x0f, 0xe0, 0x50,
	0xd1, 0x4a, 0xaf, 0xc9, 0xde, 0xcc, 0x9c, 0xce, 0x49, 0x00, 0x77, 0xfb, 0xb7, 0xe0, 0x3b, 0x9d,
	0x13, 0x9e, 0x42, 0x77, 0x6e, 0x8a, 0x85, 0xb6, 0x39, 0x29, 0xd1, 0x4b, 0x1a, 0xe3, 0x4e, 0xf6,
	0x0b, 0xc0, 0x7b, 0xd0, 0xe1, 0x35, 0xfa, 0x6b, 0xf4, 0xf9, 0x1a, 0x07, 0x5c, 0x5f, 0x28, 0x3c,
	0x03, 0x90, 0x61, 0x1f, 0xbe, 0x79, 0xc8, 0xcd, 0x6e, 0x8d, 0x84, 0xf6, 0xdc, 0x92, 0x74, 0xa4,
	0x66, 0xd2, 0x89, 0x01, 0x4f, 0xee, 0xd6, 0xc8, 0xd4, 0xf9, 0xf6, 0x97, 0x52, 0xdd, 0xb6, 0xff,
	0x0f, 0xed, 0x1a, 0x99, 0x3a, 0x7c, 0x02, 0x2d, 0xbf, 0x2e, 0x71, 0x94, 0x34, 0xc6, 0xbd, 0x49,
	0x92, 0xfe, 0x39, 0xb7, 0x74, 0x5a, 0x07, 0x95, 0x31, 0x7b, 0xf4, 0xbd, 0x09, 0xc7, 0x9b, 0x41,
	0x65, 0x54, 0x95, 0xa6, 0xa8, 0x08, 0x9f, 0x41, 0x4c, 0x85, 0xd3, 0xee, 0x86, 0x43, 0xeb, 0x4d,
	0x1e, 0x6e, 0x33, 0xfc, 0x4d, 0x5d, 0x6b, 0xf0, 0x1c, 0xda, 0xa5, 0xbc, 0x26, 0x2b, 0x9a, 0x2c,
	0x3e, 0xdb, 0x26, 0x7e, 0xeb, 0x49, 0x59, 0xe0, 0xe2, 0x53, 0x68, 0x6b, 0x47, 0x79, 0x25, 0xa2,
	0x24, 0xda, 0x79, 0x62, 0x90, 0xf8, 0x81, 0x64, 0xad, 0xb1, 0xfc, 0x72, 0xfe, 0x32, 0xf0, 0x95,
	0x27, 0x65, 0x81, 0x8b, 0x8f, 0xa1, 0xa5, 0x8b, 0x85, 0xe1, 0x77, 0xd5, 0x9b, 0x9c, 0x6e, 0xd3,
	0x84, 0x75, 0x79, 0xe6, 0xe4, 0x47, 0x04, 0x77, 0x36, 0xc7, 0x5f, 0x92, 0x5d, 0xeb, 0x39, 0xe1,
	0x27, 0x88, 0x5f, 0x70, 0x50, 0xb8, 0xd3, 0x57, 0x0f, 0x1f, 0xed, 0x74, 0xb7, 0x3a, 0x8b, 0xd1,
	0x7f, 0xde, 0xff, 0x3d, 0x27, 0xbd, 0x3f, 0xff, 0x97, 0xb4, 0xa2, 0xbd, 0xf9, 0x7f, 0x84, 0xe8,
	0x35, 0xb9, 0xbd, 0x99, 0xc7, 0x97, 0x24, 0xed, 0x7c, 0x89, 0xf7, 0xb7, 0x29, 0x9f, 0xcb, 0x8a,
	0x3e, 0x2c, 0xc9, 0xd2, 0xbf, 0x9a, 0x5f, 0xc5, 0xfc, 0x43, 0x3b, 0xff, 0x19, 0x00, 0x00, 0xff,
	0xff, 0xa8, 0x2d, 0xd7, 0xee, 0x1c, 0x05, 0x00, 0x00,
}