// Code generated by protoc-gen-go. DO NOT EDIT.
// source: present.proto

package geiqin_srv_ims_present

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

type Present struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ItemId               int64    `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64    `protobuf:"varint,4,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	FetchLimit           int32    `protobuf:"varint,5,opt,name=fetch_limit,json=fetchLimit,proto3" json:"fetch_limit,omitempty"`
	LongEffective        bool     `protobuf:"varint,6,opt,name=long_effective,json=longEffective,proto3" json:"long_effective,omitempty"`
	StartAt              string   `protobuf:"bytes,7,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt                string   `protobuf:"bytes,8,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	IssuedNum            int32    `protobuf:"varint,9,opt,name=issued_num,json=issuedNum,proto3" json:"issued_num,omitempty"`
	ExchangedNum         int32    `protobuf:"varint,10,opt,name=exchanged_num,json=exchangedNum,proto3" json:"exchanged_num,omitempty"`
	Status               int32    `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string   `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Ids                  []int64  `protobuf:"varint,14,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Present) Reset()         { *m = Present{} }
func (m *Present) String() string { return proto.CompactTextString(m) }
func (*Present) ProtoMessage()    {}
func (*Present) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cdd2afe20cfa69b, []int{0}
}

func (m *Present) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Present.Unmarshal(m, b)
}
func (m *Present) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Present.Marshal(b, m, deterministic)
}
func (m *Present) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Present.Merge(m, src)
}
func (m *Present) XXX_Size() int {
	return xxx_messageInfo_Present.Size(m)
}
func (m *Present) XXX_DiscardUnknown() {
	xxx_messageInfo_Present.DiscardUnknown(m)
}

var xxx_messageInfo_Present proto.InternalMessageInfo

func (m *Present) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Present) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Present) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *Present) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *Present) GetFetchLimit() int32 {
	if m != nil {
		return m.FetchLimit
	}
	return 0
}

func (m *Present) GetLongEffective() bool {
	if m != nil {
		return m.LongEffective
	}
	return false
}

func (m *Present) GetStartAt() string {
	if m != nil {
		return m.StartAt
	}
	return ""
}

func (m *Present) GetEndAt() string {
	if m != nil {
		return m.EndAt
	}
	return ""
}

func (m *Present) GetIssuedNum() int32 {
	if m != nil {
		return m.IssuedNum
	}
	return 0
}

func (m *Present) GetExchangedNum() int32 {
	if m != nil {
		return m.ExchangedNum
	}
	return 0
}

func (m *Present) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Present) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Present) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Present) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

//
type PresentResponse struct {
	Entity               *Present   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Present `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PresentResponse) Reset()         { *m = PresentResponse{} }
func (m *PresentResponse) String() string { return proto.CompactTextString(m) }
func (*PresentResponse) ProtoMessage()    {}
func (*PresentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cdd2afe20cfa69b, []int{1}
}

func (m *PresentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PresentResponse.Unmarshal(m, b)
}
func (m *PresentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PresentResponse.Marshal(b, m, deterministic)
}
func (m *PresentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PresentResponse.Merge(m, src)
}
func (m *PresentResponse) XXX_Size() int {
	return xxx_messageInfo_PresentResponse.Size(m)
}
func (m *PresentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PresentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PresentResponse proto.InternalMessageInfo

func (m *PresentResponse) GetEntity() *Present {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *PresentResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *PresentResponse) GetItems() []*Present {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *PresentResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *PresentResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Present)(nil), "geiqin.srv.ims.present.Present")
	proto.RegisterType((*PresentResponse)(nil), "geiqin.srv.ims.present.PresentResponse")
}

func init() {
	proto.RegisterFile("present.proto", fileDescriptor_7cdd2afe20cfa69b)
}

var fileDescriptor_7cdd2afe20cfa69b = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x8f, 0xd3, 0x30,
	0x14, 0xa4, 0x4d, 0x93, 0x6e, 0x5f, 0x3f, 0x40, 0x96, 0x58, 0xcc, 0x8a, 0xd5, 0x86, 0x22, 0x44,
	0x4f, 0x11, 0xea, 0x0a, 0x71, 0x2e, 0xb0, 0x42, 0x95, 0x56, 0x08, 0xbc, 0x7c, 0x1c, 0xab, 0x10,
	0xbf, 0xb6, 0xd6, 0x36, 0x4e, 0xb0, 0x9d, 0x0a, 0xce, 0xfc, 0x02, 0x7e, 0x00, 0xff, 0x15, 0xf9,
	0x25, 0xe5, 0x44, 0x59, 0x90, 0x7a, 0xb3, 0x67, 0xe6, 0xcd, 0xd8, 0x9e, 0x28, 0x30, 0x2c, 0x0d,
	0x5a, 0xd4, 0x2e, 0x29, 0x4d, 0xe1, 0x0a, 0x76, 0xbc, 0x42, 0xf5, 0x45, 0xe9, 0xc4, 0x9a, 0x6d,
	0xa2, 0x72, 0x9b, 0x34, 0xec, 0xc9, 0x20, 0x2b, 0xf2, 0xbc, 0xd0, 0xb5, 0x6a, 0xfc, 0x3d, 0x80,
	0xee, 0xdb, 0x9a, 0x61, 0x23, 0x68, 0x2b, 0xc9, 0x5b, 0x71, 0x6b, 0x12, 0x88, 0xb6, 0x92, 0x8c,
	0x41, 0x47, 0xa7, 0x39, 0xf2, 0x76, 0xdc, 0x9a, 0xf4, 0x04, 0xad, 0xd9, 0x3d, 0xe8, 0x2a, 0x87,
	0xf9, 0x42, 0x49, 0x1e, 0x90, 0x30, 0xf2, 0xdb, 0xb9, 0x64, 0x77, 0x21, 0xb2, 0xd7, 0x95, 0xc7,
	0x3b, 0x84, 0x87, 0xf6, 0xba, 0x9a, 0x4b, 0x76, 0x06, 0xfd, 0x25, 0xba, 0x6c, 0xbd, 0xd8, 0xa8,
	0x5c, 0x39, 0x1e, 0xc6, 0xad, 0x49, 0x28, 0x80, 0xa0, 0x4b, 0x8f, 0xb0, 0xc7, 0x30, 0xda, 0x14,
	0x7a, 0xb5, 0xc0, 0xe5, 0x12, 0x33, 0xa7, 0xb6, 0xc8, 0xa3, 0xb8, 0x35, 0x39, 0x12, 0x43, 0x8f,
	0x5e, 0xec, 0x40, 0x76, 0x1f, 0x8e, 0xac, 0x4b, 0x8d, 0x5b, 0xa4, 0x8e, 0x77, 0xe9, 0x3c, 0x5d,
	0xda, 0xcf, 0x9c, 0x4f, 0x46, 0x2d, 0x3d, 0x71, 0x44, 0x44, 0x88, 0x5a, 0xce, 0x1c, 0x3b, 0x05,
	0x50, 0xd6, 0x56, 0x28, 0x17, 0xba, 0xca, 0x79, 0x8f, 0x82, 0x7b, 0x35, 0xf2, 0xa6, 0xca, 0xd9,
	0x23, 0x18, 0xe2, 0xd7, 0x6c, 0x9d, 0xea, 0x55, 0xa3, 0x00, 0x52, 0x0c, 0x7e, 0x83, 0x5e, 0x74,
	0x0c, 0x91, 0x75, 0xa9, 0xab, 0x2c, 0xef, 0x13, 0xdb, 0xec, 0xbc, 0x77, 0x66, 0x30, 0x75, 0x48,
	0xb1, 0x03, 0x8a, 0xed, 0x35, 0x48, 0x1d, 0x5d, 0x95, 0x72, 0x47, 0x0f, 0x6b, 0xba, 0x41, 0x66,
	0x8e, 0xdd, 0x81, 0x40, 0x49, 0xcb, 0x47, 0x71, 0x30, 0x09, 0x84, 0x5f, 0x8e, 0x7f, 0xb6, 0xe1,
	0x76, 0xd3, 0x82, 0x40, 0x5b, 0x16, 0xda, 0x22, 0x7b, 0xee, 0xaf, 0xe5, 0x94, 0xfb, 0x46, 0x8d,
	0xf4, 0xa7, 0x67, 0xc9, 0x9f, 0x0b, 0x4d, 0x76, 0x83, 0x8d, 0x9c, 0x9d, 0x43, 0x58, 0xa6, 0x2b,
	0x34, 0xd4, 0x5b, 0x7f, 0x7a, 0xba, 0x77, 0xce, 0x8b, 0x44, 0xad, 0x65, 0xcf, 0x20, 0xf4, 0x45,
	0x5a, 0x1e, 0xc4, 0xc1, 0xbf, 0x84, 0xd5, 0x6a, 0x9f, 0x85, 0xc6, 0x14, 0x86, 0x4a, 0xff, 0x4b,
	0xd6, 0x85, 0x17, 0x89, 0x5a, 0xcb, 0x9e, 0x42, 0x47, 0xe9, 0x65, 0x41, 0x1f, 0x43, 0x7f, 0xfa,
	0x60, 0xdf, 0xcc, 0x5c, 0x2f, 0x0b, 0x41, 0xca, 0xe9, 0x8f, 0x0e, 0x8c, 0x9a, 0xe4, 0x2b, 0x34,
	0x5b, 0x95, 0x21, 0x7b, 0x0f, 0xd1, 0x4b, 0x7a, 0x70, 0x76, 0xd3, 0x59, 0x4f, 0x9e, 0xdc, 0x74,
	0x99, 0xe6, 0xc9, 0xc7, 0xb7, 0xbc, 0xeb, 0x07, 0xea, 0xe9, 0xd0, 0xae, 0xaf, 0x70, 0x83, 0x07,
	0x76, 0x7d, 0x07, 0xc1, 0x6b, 0x74, 0x07, 0xb5, 0x14, 0xd0, 0xb9, 0x54, 0xf6, 0xb0, 0x9e, 0x1f,
	0x21, 0xba, 0xc2, 0xd4, 0x64, 0x6b, 0xf6, 0x70, 0xdf, 0xd0, 0x8b, 0xd4, 0xe2, 0xa7, 0x35, 0x1a,
	0xfc, 0x0f, 0xdf, 0xcf, 0x11, 0xfd, 0xc0, 0xce, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x7c,
	0x79, 0xae, 0xf7, 0x04, 0x00, 0x00,
}
