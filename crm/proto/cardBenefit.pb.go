// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cardBenefit.proto

package geiqin_srv_crm

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

type CardBenefit struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CardId               int32    `protobuf:"varint,2,opt,name=card_id,json=cardId,proto3" json:"card_id"`
	BenefitId            int32    `protobuf:"varint,3,opt,name=benefit_id,json=benefitId,proto3" json:"benefit_id"`
	Params               string   `protobuf:"bytes,4,opt,name=params,proto3" json:"params"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Benefit              *Benefit `protobuf:"bytes,7,opt,name=benefit,proto3" json:"benefit"`
	Ids                  []int32  `protobuf:"varint,8,rep,packed,name=ids,proto3" json:"ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CardBenefit) Reset()         { *m = CardBenefit{} }
func (m *CardBenefit) String() string { return proto.CompactTextString(m) }
func (*CardBenefit) ProtoMessage()    {}
func (*CardBenefit) Descriptor() ([]byte, []int) {
	return fileDescriptor_f69b83a596fd8713, []int{0}
}

func (m *CardBenefit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CardBenefit.Unmarshal(m, b)
}
func (m *CardBenefit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CardBenefit.Marshal(b, m, deterministic)
}
func (m *CardBenefit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CardBenefit.Merge(m, src)
}
func (m *CardBenefit) XXX_Size() int {
	return xxx_messageInfo_CardBenefit.Size(m)
}
func (m *CardBenefit) XXX_DiscardUnknown() {
	xxx_messageInfo_CardBenefit.DiscardUnknown(m)
}

var xxx_messageInfo_CardBenefit proto.InternalMessageInfo

func (m *CardBenefit) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CardBenefit) GetCardId() int32 {
	if m != nil {
		return m.CardId
	}
	return 0
}

func (m *CardBenefit) GetBenefitId() int32 {
	if m != nil {
		return m.BenefitId
	}
	return 0
}

func (m *CardBenefit) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

func (m *CardBenefit) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *CardBenefit) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *CardBenefit) GetBenefit() *Benefit {
	if m != nil {
		return m.Benefit
	}
	return nil
}

func (m *CardBenefit) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

//
type CardBenefitResponse struct {
	Entity               *CardBenefit   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager                *Pager         `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items                []*CardBenefit `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error                *Error         `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info                 *Info          `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CardBenefitResponse) Reset()         { *m = CardBenefitResponse{} }
func (m *CardBenefitResponse) String() string { return proto.CompactTextString(m) }
func (*CardBenefitResponse) ProtoMessage()    {}
func (*CardBenefitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f69b83a596fd8713, []int{1}
}

func (m *CardBenefitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CardBenefitResponse.Unmarshal(m, b)
}
func (m *CardBenefitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CardBenefitResponse.Marshal(b, m, deterministic)
}
func (m *CardBenefitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CardBenefitResponse.Merge(m, src)
}
func (m *CardBenefitResponse) XXX_Size() int {
	return xxx_messageInfo_CardBenefitResponse.Size(m)
}
func (m *CardBenefitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CardBenefitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CardBenefitResponse proto.InternalMessageInfo

func (m *CardBenefitResponse) GetEntity() *CardBenefit {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *CardBenefitResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *CardBenefitResponse) GetItems() []*CardBenefit {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CardBenefitResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CardBenefitResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*CardBenefit)(nil), "geiqin.srv.crm.CardBenefit")
	proto.RegisterType((*CardBenefitResponse)(nil), "geiqin.srv.crm.CardBenefitResponse")
}

func init() { proto.RegisterFile("cardBenefit.proto", fileDescriptor_f69b83a596fd8713) }

var fileDescriptor_f69b83a596fd8713 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xdf, 0xaa, 0xd3, 0x40,
	0x10, 0xc6, 0x4d, 0x72, 0x92, 0xda, 0x89, 0x1e, 0x74, 0xfd, 0x73, 0x62, 0x45, 0x08, 0xf5, 0x26,
	0x20, 0x04, 0x9a, 0x3e, 0x41, 0xad, 0x22, 0x11, 0x2f, 0x64, 0x8b, 0x78, 0x59, 0xb6, 0xd9, 0x69,
	0xbb, 0x60, 0xb2, 0x71, 0xb3, 0x16, 0xbc, 0xf4, 0x21, 0x7c, 0x4f, 0x1f, 0x41, 0x76, 0xb3, 0xc5,
	0x58, 0xa5, 0x7a, 0xd1, 0xbb, 0xec, 0x7c, 0xdf, 0xfc, 0x66, 0x67, 0x26, 0x0b, 0xf7, 0x2b, 0xa6,
	0xf8, 0x4b, 0x6c, 0x70, 0x2b, 0x74, 0xde, 0x2a, 0xa9, 0x25, 0xb9, 0xde, 0xa1, 0xf8, 0x2c, 0x9a,
	0xbc, 0x53, 0x87, 0xbc, 0x52, 0xf5, 0xe4, 0x4e, 0x25, 0xeb, 0x5a, 0x36, 0xbd, 0x3a, 0xb9, 0xbb,
	0x19, 0x9a, 0xa7, 0x3f, 0x3c, 0x88, 0x97, 0xbf, 0x10, 0xe4, 0x1a, 0x7c, 0xc1, 0x13, 0x2f, 0xf5,
	0xb2, 0x90, 0xfa, 0x82, 0x93, 0x1b, 0x18, 0x99, 0x0a, 0x6b, 0xc1, 0x13, 0xdf, 0x06, 0x23, 0x73,
	0x2c, 0x39, 0x79, 0x06, 0xe0, 0x48, 0x46, 0x0b, 0xac, 0x36, 0x76, 0x91, 0x92, 0x93, 0xc7, 0x10,
	0xb5, 0x4c, 0xb1, 0xba, 0x4b, 0xae, 0x52, 0x2f, 0x1b, 0x53, 0x77, 0x32, 0x69, 0x95, 0x42, 0xa6,
	0x91, 0xaf, 0x99, 0x4e, 0x42, 0xab, 0x8d, 0x5d, 0x64, 0xa1, 0x8d, 0xfc, 0xa5, 0xe5, 0x47, 0x39,
	0xea, 0x65, 0x17, 0x59, 0x68, 0x32, 0x83, 0x91, 0x2b, 0x91, 0x8c, 0x52, 0x2f, 0x8b, 0x8b, 0x9b,
	0xfc, 0xf7, 0x66, 0x73, 0xd7, 0x07, 0x3d, 0xfa, 0xc8, 0x3d, 0x08, 0x04, 0xef, 0x92, 0xdb, 0x69,
	0x90, 0x85, 0xd4, 0x7c, 0x4e, 0xbf, 0xf9, 0xf0, 0x60, 0xd0, 0x32, 0xc5, 0xae, 0x95, 0x4d, 0x87,
	0x64, 0x0e, 0x11, 0x36, 0x5a, 0xe8, 0xaf, 0xb6, 0xfd, 0xb8, 0x78, 0x7a, 0xca, 0x1e, 0x26, 0x39,
	0x2b, 0x79, 0x01, 0x61, 0xcb, 0x76, 0xa8, 0xec, 0x74, 0xe2, 0xe2, 0xd1, 0x69, 0xce, 0x7b, 0x23,
	0xd2, 0xde, 0x43, 0x66, 0x10, 0x0a, 0x8d, 0x75, 0x97, 0x04, 0x69, 0xf0, 0xaf, 0x02, 0xbd, 0xd3,
	0xf0, 0x51, 0x29, 0xa9, 0xec, 0x18, 0xff, 0xc2, 0x7f, 0x6d, 0x44, 0xda, 0x7b, 0x48, 0x06, 0x57,
	0xa2, 0xd9, 0x4a, 0x3b, 0xd6, 0xb8, 0x78, 0x78, 0xea, 0x2d, 0x9b, 0xad, 0xa4, 0xd6, 0x51, 0x7c,
	0x0f, 0x80, 0x0c, 0xaa, 0xad, 0x50, 0x1d, 0x44, 0x85, 0xe4, 0x1d, 0x44, 0x4b, 0xbb, 0x0b, 0x72,
	0xee, 0x6e, 0x93, 0xe7, 0xe7, 0x2e, 0xee, 0xc6, 0x39, 0xbd, 0x65, 0x68, 0x1f, 0xec, 0xea, 0x2e,
	0x45, 0x7b, 0x85, 0x9f, 0xf0, 0x42, 0xb4, 0x12, 0x82, 0x37, 0xa8, 0x2f, 0x82, 0x7a, 0x0b, 0xd1,
	0x0a, 0x99, 0xaa, 0xf6, 0xe4, 0xc9, 0x1f, 0x7f, 0x23, 0xeb, 0xf0, 0xe3, 0x1e, 0x15, 0xfe, 0x27,
	0x6b, 0x13, 0xd9, 0x57, 0x39, 0xff, 0x19, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xe7, 0x90, 0xb6, 0xd7,
	0x03, 0x00, 0x00,
}
