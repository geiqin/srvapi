// Code generated by protoc-gen-go. DO NOT EDIT.
// source: branch.proto

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

type Branch struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	StoreId              int64    `protobuf:"varint,4,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	AreaId               int64    `protobuf:"varint,5,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Address              string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Lng                  string   `protobuf:"bytes,7,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat                  string   `protobuf:"bytes,8,opt,name=lat,proto3" json:"lat,omitempty"`
	CreatorId            int64    `protobuf:"varint,9,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Locked               bool     `protobuf:"varint,10,opt,name=locked,proto3" json:"locked,omitempty"`
	CreatedAt            string   `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Ids                  []int64  `protobuf:"varint,13,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Branch) Reset()         { *m = Branch{} }
func (m *Branch) String() string { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()    {}
func (*Branch) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ce1f5884b7e047, []int{0}
}

func (m *Branch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Branch.Unmarshal(m, b)
}
func (m *Branch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Branch.Marshal(b, m, deterministic)
}
func (m *Branch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Branch.Merge(m, src)
}
func (m *Branch) XXX_Size() int {
	return xxx_messageInfo_Branch.Size(m)
}
func (m *Branch) XXX_DiscardUnknown() {
	xxx_messageInfo_Branch.DiscardUnknown(m)
}

var xxx_messageInfo_Branch proto.InternalMessageInfo

func (m *Branch) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Branch) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Branch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Branch) GetStoreId() int64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *Branch) GetAreaId() int64 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *Branch) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Branch) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *Branch) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *Branch) GetCreatorId() int64 {
	if m != nil {
		return m.CreatorId
	}
	return 0
}

func (m *Branch) GetLocked() bool {
	if m != nil {
		return m.Locked
	}
	return false
}

func (m *Branch) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Branch) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Branch) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

//
type BranchResponse struct {
	Entity               *Branch   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Branch `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error    `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info     `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BranchResponse) Reset()         { *m = BranchResponse{} }
func (m *BranchResponse) String() string { return proto.CompactTextString(m) }
func (*BranchResponse) ProtoMessage()    {}
func (*BranchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ce1f5884b7e047, []int{1}
}

func (m *BranchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BranchResponse.Unmarshal(m, b)
}
func (m *BranchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BranchResponse.Marshal(b, m, deterministic)
}
func (m *BranchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchResponse.Merge(m, src)
}
func (m *BranchResponse) XXX_Size() int {
	return xxx_messageInfo_BranchResponse.Size(m)
}
func (m *BranchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BranchResponse proto.InternalMessageInfo

func (m *BranchResponse) GetEntity() *Branch {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *BranchResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *BranchResponse) GetItems() []*Branch {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *BranchResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *BranchResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Branch)(nil), "geiqin.srv.mts.Branch")
	proto.RegisterType((*BranchResponse)(nil), "geiqin.srv.mts.BranchResponse")
}

func init() {
	proto.RegisterFile("branch.proto", fileDescriptor_20ce1f5884b7e047)
}

var fileDescriptor_20ce1f5884b7e047 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x49, 0xd2, 0xa6, 0xed, 0x69, 0x57, 0xa1, 0x23, 0x18, 0xde, 0x24, 0x50, 0xd4, 0xab,
	0x48, 0xa0, 0x5c, 0x94, 0x07, 0x40, 0x1b, 0x4c, 0xa8, 0x77, 0x28, 0x13, 0xe2, 0x72, 0xf2, 0xe2,
	0xb3, 0xd6, 0xa2, 0xb1, 0x8b, 0xed, 0x4d, 0xe2, 0xd1, 0x78, 0x1e, 0xde, 0x80, 0x27, 0x40, 0x3e,
	0xc9, 0x90, 0x18, 0x14, 0x2e, 0x7a, 0x77, 0xce, 0xff, 0xff, 0xfe, 0x62, 0xfd, 0x56, 0x60, 0x76,
	0xed, 0xa4, 0x69, 0x36, 0xd5, 0xce, 0xd9, 0x60, 0x71, 0xbe, 0x26, 0xfd, 0x45, 0x9b, 0xca, 0xbb,
	0xbb, 0xaa, 0x0d, 0xfe, 0x74, 0xd6, 0xd8, 0xb6, 0xb5, 0xa6, 0x73, 0x17, 0xdf, 0x52, 0xc8, 0xcf,
	0x39, 0x8e, 0x73, 0x48, 0xb5, 0x12, 0x49, 0x91, 0x94, 0x59, 0x9d, 0x6a, 0x85, 0x08, 0x03, 0xbf,
	0xbd, 0x5d, 0x8b, 0xb4, 0x48, 0xca, 0x49, 0xcd, 0x73, 0xd4, 0x8c, 0x6c, 0x49, 0x64, 0x9d, 0x16,
	0x67, 0x3c, 0x81, 0xb1, 0x0f, 0xd6, 0xd1, 0x95, 0x56, 0x62, 0xc0, 0xa7, 0x47, 0xbc, 0xaf, 0x14,
	0x3e, 0x83, 0x91, 0x74, 0x24, 0xa3, 0x33, 0x64, 0x27, 0x8f, 0xeb, 0x4a, 0xa1, 0x80, 0x91, 0x54,
	0xca, 0x91, 0xf7, 0x22, 0x67, 0xd4, 0xfd, 0x8a, 0x8f, 0x21, 0xdb, 0x9a, 0xb5, 0x18, 0xb1, 0x1a,
	0x47, 0x56, 0x64, 0x10, 0xe3, 0x5e, 0x91, 0x01, 0x9f, 0x03, 0x34, 0x8e, 0x64, 0xb0, 0x2e, 0x92,
	0x27, 0x4c, 0x9e, 0xf4, 0xca, 0x4a, 0xe1, 0x31, 0xe4, 0x5b, 0xdb, 0x7c, 0x26, 0x25, 0xa0, 0x48,
	0xca, 0x71, 0xdd, 0x6f, 0xbf, 0x8e, 0x91, 0xba, 0x92, 0x41, 0x4c, 0x99, 0x37, 0xe9, 0x95, 0x33,
	0xa6, 0xde, 0xee, 0xd4, 0xbd, 0x3d, 0xeb, 0xec, 0x5e, 0x39, 0x0b, 0xf1, 0x1a, 0x5a, 0x79, 0x71,
	0x54, 0x64, 0x65, 0x56, 0xc7, 0x71, 0xf1, 0x23, 0x81, 0x79, 0xd7, 0x5d, 0x4d, 0x7e, 0x67, 0x8d,
	0x27, 0xac, 0x20, 0x27, 0x13, 0x74, 0xf8, 0xca, 0x3d, 0x4e, 0x97, 0xc7, 0xd5, 0xef, 0xed, 0x57,
	0x7d, 0xbe, 0x4f, 0xe1, 0x4b, 0x18, 0xee, 0xe4, 0x9a, 0x1c, 0x97, 0x3c, 0x5d, 0x3e, 0x7d, 0x18,
	0xff, 0x10, 0xcd, 0xba, 0xcb, 0xe0, 0x2b, 0x18, 0xea, 0x40, 0xad, 0x17, 0x59, 0x91, 0xfd, 0x83,
	0xdd, 0x85, 0x22, 0x9a, 0x9c, 0xb3, 0x8e, 0xdf, 0xe4, 0x2f, 0xe8, 0x8b, 0x68, 0xd6, 0x5d, 0x06,
	0x4b, 0x18, 0x68, 0x73, 0x63, 0xf9, 0x95, 0xa6, 0xcb, 0x27, 0x0f, 0xb3, 0x2b, 0x73, 0x63, 0x6b,
	0x4e, 0x2c, 0xbf, 0xa7, 0x70, 0xd4, 0x7d, 0xe8, 0x92, 0xdc, 0x9d, 0x6e, 0x08, 0xcf, 0x21, 0x7f,
	0xcb, 0x25, 0xe2, 0x9e, 0x1b, 0x9d, 0xbe, 0xd8, 0x73, 0xd3, 0xbe, 0xb5, 0xc5, 0xa3, 0xc8, 0xf8,
	0xc8, 0x4d, 0x1f, 0xc6, 0x78, 0x47, 0x5b, 0x3a, 0x88, 0xf1, 0x06, 0xb2, 0xf7, 0x14, 0x0e, 0x00,
	0x5c, 0x40, 0x7e, 0x49, 0xd2, 0x35, 0x1b, 0x3c, 0xf9, 0x23, 0x2b, 0x3d, 0x7d, 0xda, 0x90, 0xa3,
	0xff, 0x63, 0xae, 0x73, 0xfe, 0x3b, 0x5f, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x48, 0x6e,
	0xd0, 0xcb, 0x03, 0x00, 0x00,
}
