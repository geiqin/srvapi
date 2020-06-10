// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delivery.proto

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

type DeliveryWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting              string   `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Keywords             string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id                   int64    `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliveryWhere) Reset()         { *m = DeliveryWhere{} }
func (m *DeliveryWhere) String() string { return proto.CompactTextString(m) }
func (*DeliveryWhere) ProtoMessage()    {}
func (*DeliveryWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf387bcb4e23d880, []int{0}
}

func (m *DeliveryWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryWhere.Unmarshal(m, b)
}
func (m *DeliveryWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryWhere.Marshal(b, m, deterministic)
}
func (m *DeliveryWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryWhere.Merge(m, src)
}
func (m *DeliveryWhere) XXX_Size() int {
	return xxx_messageInfo_DeliveryWhere.Size(m)
}
func (m *DeliveryWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryWhere.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryWhere proto.InternalMessageInfo

func (m *DeliveryWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *DeliveryWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *DeliveryWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
}

func (m *DeliveryWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *DeliveryWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeliveryWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Delivery struct {
	Id                    int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LocationId            int64            `protobuf:"varint,3,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Method                int32            `protobuf:"varint,4,opt,name=method,proto3" json:"method,omitempty"`
	MethodArr             []int32          `protobuf:"varint,31,rep,packed,name=method_arr,json=methodArr,proto3" json:"method_arr,omitempty"`
	Template              int32            `protobuf:"varint,5,opt,name=template,proto3" json:"template,omitempty"`
	RangeName             string           `protobuf:"bytes,6,opt,name=range_name,json=rangeName,proto3" json:"range_name,omitempty"`
	RangeContent          string           `protobuf:"bytes,7,opt,name=range_content,json=rangeContent,proto3" json:"range_content,omitempty"`
	RangeImageUrl         string           `protobuf:"bytes,8,opt,name=range_image_url,json=rangeImageUrl,proto3" json:"range_image_url,omitempty"`
	RangeStartPrice       float32          `protobuf:"fixed32,9,opt,name=range_start_price,json=rangeStartPrice,proto3" json:"range_start_price,omitempty"`
	RangeFee              float32          `protobuf:"fixed32,10,opt,name=range_fee,json=rangeFee,proto3" json:"range_fee,omitempty"`
	Standard              int32            `protobuf:"varint,11,opt,name=standard,proto3" json:"standard,omitempty"`
	FirstKm               float32          `protobuf:"fixed32,12,opt,name=first_km,json=firstKm,proto3" json:"first_km,omitempty"`
	FirstKmFee            float32          `protobuf:"fixed32,13,opt,name=first_km_fee,json=firstKmFee,proto3" json:"first_km_fee,omitempty"`
	AdditionalKm          float32          `protobuf:"fixed32,14,opt,name=additional_km,json=additionalKm,proto3" json:"additional_km,omitempty"`
	AdditionalKmFee       float32          `protobuf:"fixed32,15,opt,name=additional_km_fee,json=additionalKmFee,proto3" json:"additional_km_fee,omitempty"`
	FirstWeight           float32          `protobuf:"fixed32,16,opt,name=first_weight,json=firstWeight,proto3" json:"first_weight,omitempty"`
	AdditionalWeight      float32          `protobuf:"fixed32,17,opt,name=additional_weight,json=additionalWeight,proto3" json:"additional_weight,omitempty"`
	AdditionalWeightFee   float32          `protobuf:"fixed32,18,opt,name=additional_weight_fee,json=additionalWeightFee,proto3" json:"additional_weight_fee,omitempty"`
	IsTimedArrival        bool             `protobuf:"varint,19,opt,name=is_timed_arrival,json=isTimedArrival,proto3" json:"is_timed_arrival,omitempty"`
	Delivery              string           `protobuf:"bytes,20,opt,name=delivery,proto3" json:"delivery,omitempty"`
	DeliveryRepeatWeeks   int32            `protobuf:"varint,21,opt,name=delivery_repeat_weeks,json=deliveryRepeatWeeks,proto3" json:"delivery_repeat_weeks,omitempty"`
	DeliveryRepeatWeekArr []int32          `protobuf:"varint,22,rep,packed,name=delivery_repeat_week_arr,json=deliveryRepeatWeekArr,proto3" json:"delivery_repeat_week_arr,omitempty"`
	SubDeliveryTime       string           `protobuf:"bytes,23,opt,name=sub_delivery_time,json=subDeliveryTime,proto3" json:"sub_delivery_time,omitempty"`
	Appointment           string           `protobuf:"bytes,24,opt,name=appointment,proto3" json:"appointment,omitempty"`
	AppointmentNum        int32            `protobuf:"varint,25,opt,name=appointment_num,json=appointmentNum,proto3" json:"appointment_num,omitempty"`
	MaxAppointmentNum     int32            `protobuf:"varint,26,opt,name=max_appointment_num,json=maxAppointmentNum,proto3" json:"max_appointment_num,omitempty"`
	CreatedAt             string           `protobuf:"bytes,27,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt             string           `protobuf:"bytes,28,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeliveryTimes         []*Times         `protobuf:"bytes,29,rep,name=delivery_times,json=deliveryTimes,proto3" json:"delivery_times,omitempty"`
	Ranges                []*DeliveryRange `protobuf:"bytes,30,rep,name=ranges,proto3" json:"ranges,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}         `json:"-"`
	XXX_unrecognized      []byte           `json:"-"`
	XXX_sizecache         int32            `json:"-"`
}

func (m *Delivery) Reset()         { *m = Delivery{} }
func (m *Delivery) String() string { return proto.CompactTextString(m) }
func (*Delivery) ProtoMessage()    {}
func (*Delivery) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf387bcb4e23d880, []int{1}
}

func (m *Delivery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Delivery.Unmarshal(m, b)
}
func (m *Delivery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Delivery.Marshal(b, m, deterministic)
}
func (m *Delivery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delivery.Merge(m, src)
}
func (m *Delivery) XXX_Size() int {
	return xxx_messageInfo_Delivery.Size(m)
}
func (m *Delivery) XXX_DiscardUnknown() {
	xxx_messageInfo_Delivery.DiscardUnknown(m)
}

var xxx_messageInfo_Delivery proto.InternalMessageInfo

func (m *Delivery) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Delivery) GetLocationId() int64 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

func (m *Delivery) GetMethod() int32 {
	if m != nil {
		return m.Method
	}
	return 0
}

func (m *Delivery) GetMethodArr() []int32 {
	if m != nil {
		return m.MethodArr
	}
	return nil
}

func (m *Delivery) GetTemplate() int32 {
	if m != nil {
		return m.Template
	}
	return 0
}

func (m *Delivery) GetRangeName() string {
	if m != nil {
		return m.RangeName
	}
	return ""
}

func (m *Delivery) GetRangeContent() string {
	if m != nil {
		return m.RangeContent
	}
	return ""
}

func (m *Delivery) GetRangeImageUrl() string {
	if m != nil {
		return m.RangeImageUrl
	}
	return ""
}

func (m *Delivery) GetRangeStartPrice() float32 {
	if m != nil {
		return m.RangeStartPrice
	}
	return 0
}

func (m *Delivery) GetRangeFee() float32 {
	if m != nil {
		return m.RangeFee
	}
	return 0
}

func (m *Delivery) GetStandard() int32 {
	if m != nil {
		return m.Standard
	}
	return 0
}

func (m *Delivery) GetFirstKm() float32 {
	if m != nil {
		return m.FirstKm
	}
	return 0
}

func (m *Delivery) GetFirstKmFee() float32 {
	if m != nil {
		return m.FirstKmFee
	}
	return 0
}

func (m *Delivery) GetAdditionalKm() float32 {
	if m != nil {
		return m.AdditionalKm
	}
	return 0
}

func (m *Delivery) GetAdditionalKmFee() float32 {
	if m != nil {
		return m.AdditionalKmFee
	}
	return 0
}

func (m *Delivery) GetFirstWeight() float32 {
	if m != nil {
		return m.FirstWeight
	}
	return 0
}

func (m *Delivery) GetAdditionalWeight() float32 {
	if m != nil {
		return m.AdditionalWeight
	}
	return 0
}

func (m *Delivery) GetAdditionalWeightFee() float32 {
	if m != nil {
		return m.AdditionalWeightFee
	}
	return 0
}

func (m *Delivery) GetIsTimedArrival() bool {
	if m != nil {
		return m.IsTimedArrival
	}
	return false
}

func (m *Delivery) GetDelivery() string {
	if m != nil {
		return m.Delivery
	}
	return ""
}

func (m *Delivery) GetDeliveryRepeatWeeks() int32 {
	if m != nil {
		return m.DeliveryRepeatWeeks
	}
	return 0
}

func (m *Delivery) GetDeliveryRepeatWeekArr() []int32 {
	if m != nil {
		return m.DeliveryRepeatWeekArr
	}
	return nil
}

func (m *Delivery) GetSubDeliveryTime() string {
	if m != nil {
		return m.SubDeliveryTime
	}
	return ""
}

func (m *Delivery) GetAppointment() string {
	if m != nil {
		return m.Appointment
	}
	return ""
}

func (m *Delivery) GetAppointmentNum() int32 {
	if m != nil {
		return m.AppointmentNum
	}
	return 0
}

func (m *Delivery) GetMaxAppointmentNum() int32 {
	if m != nil {
		return m.MaxAppointmentNum
	}
	return 0
}

func (m *Delivery) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Delivery) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Delivery) GetDeliveryTimes() []*Times {
	if m != nil {
		return m.DeliveryTimes
	}
	return nil
}

func (m *Delivery) GetRanges() []*DeliveryRange {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type DeliveryResponse struct {
	Entity               *Delivery   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Delivery `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeliveryResponse) Reset()         { *m = DeliveryResponse{} }
func (m *DeliveryResponse) String() string { return proto.CompactTextString(m) }
func (*DeliveryResponse) ProtoMessage()    {}
func (*DeliveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf387bcb4e23d880, []int{2}
}

func (m *DeliveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryResponse.Unmarshal(m, b)
}
func (m *DeliveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryResponse.Marshal(b, m, deterministic)
}
func (m *DeliveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryResponse.Merge(m, src)
}
func (m *DeliveryResponse) XXX_Size() int {
	return xxx_messageInfo_DeliveryResponse.Size(m)
}
func (m *DeliveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryResponse proto.InternalMessageInfo

func (m *DeliveryResponse) GetEntity() *Delivery {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *DeliveryResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *DeliveryResponse) GetItems() []*Delivery {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *DeliveryResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DeliveryResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*DeliveryWhere)(nil), "geiqin.srv.tms.DeliveryWhere")
	proto.RegisterType((*Delivery)(nil), "geiqin.srv.tms.Delivery")
	proto.RegisterType((*DeliveryResponse)(nil), "geiqin.srv.tms.DeliveryResponse")
}

func init() {
	proto.RegisterFile("delivery.proto", fileDescriptor_bf387bcb4e23d880)
}

var fileDescriptor_bf387bcb4e23d880 = []byte{
	// 919 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x2e, 0xcd, 0x48, 0x96, 0x47, 0xbf, 0x5e, 0xdb, 0xe9, 0x46, 0xa9, 0x6b, 0xd6, 0x01, 0x5a,
	0x22, 0x01, 0x84, 0x42, 0x45, 0xd1, 0x4b, 0x0f, 0x55, 0x9d, 0x26, 0x30, 0xd2, 0xa6, 0x01, 0xdd,
	0xc0, 0x47, 0x82, 0x16, 0x47, 0xf2, 0xc2, 0x5a, 0x92, 0xdd, 0x5d, 0x39, 0x71, 0x1e, 0xa5, 0xef,
	0xd6, 0x47, 0xe8, 0xbd, 0xa7, 0xa2, 0xd8, 0x59, 0x52, 0x92, 0x65, 0xeb, 0xe4, 0xdb, 0xce, 0x37,
	0xdf, 0x7c, 0x33, 0xb3, 0x9c, 0x1d, 0x42, 0x27, 0xc5, 0x99, 0xb8, 0x46, 0x75, 0x33, 0x28, 0x54,
	0x6e, 0x72, 0xd6, 0x99, 0xa2, 0xf8, 0x53, 0x64, 0x03, 0xad, 0xae, 0x07, 0x46, 0xea, 0x7e, 0x6b,
	0x9c, 0x4b, 0x99, 0x67, 0xce, 0xdb, 0x6f, 0x6b, 0x34, 0x46, 0x64, 0xd3, 0xd2, 0x6c, 0x1a, 0x21,
	0x51, 0x97, 0xc6, 0x5e, 0xa5, 0x14, 0x25, 0xd9, 0x14, 0x1d, 0x78, 0xfc, 0x97, 0x07, 0xed, 0x97,
	0x25, 0x7e, 0x7e, 0x89, 0x0a, 0xd9, 0x3e, 0xd4, 0x8a, 0x64, 0x8a, 0x29, 0xf7, 0x02, 0x2f, 0xac,
	0x45, 0xce, 0x60, 0x4f, 0x61, 0xc7, 0x1e, 0x62, 0x2d, 0x3e, 0x21, 0xdf, 0x22, 0x4f, 0xc3, 0x02,
	0x67, 0xe2, 0x13, 0x32, 0x0e, 0xdb, 0x3a, 0x57, 0x36, 0x2f, 0xf7, 0x03, 0x2f, 0xdc, 0x89, 0x2a,
	0x93, 0xf5, 0xa1, 0x71, 0x85, 0x37, 0x1f, 0x72, 0x95, 0x6a, 0xfe, 0x88, 0x5c, 0x0b, 0x9b, 0x75,
	0x60, 0x4b, 0xa4, 0xbc, 0x16, 0x78, 0xa1, 0x1f, 0x6d, 0x89, 0x94, 0xf5, 0xc0, 0x17, 0xa9, 0xe6,
	0xf5, 0xc0, 0x0f, 0xfd, 0xc8, 0x1e, 0x8f, 0xff, 0x6b, 0x40, 0xa3, 0x2a, 0xae, 0xa4, 0x7b, 0x0b,
	0xfa, 0x11, 0x34, 0x67, 0xf9, 0x38, 0x31, 0x22, 0xcf, 0x62, 0x91, 0x52, 0x62, 0x3f, 0x82, 0x0a,
	0x3a, 0x4d, 0xd9, 0x63, 0xa8, 0x4b, 0x34, 0x97, 0x79, 0x4a, 0x99, 0x6b, 0x51, 0x69, 0xb1, 0x43,
	0x00, 0x77, 0x8a, 0x13, 0xa5, 0xf8, 0x51, 0xe0, 0x87, 0xb5, 0x68, 0xc7, 0x21, 0x23, 0xa5, 0x6c,
	0xc9, 0x06, 0x65, 0x31, 0x4b, 0x0c, 0x52, 0x71, 0xb5, 0x68, 0x61, 0xdb, 0x50, 0x65, 0x2f, 0x2f,
	0xce, 0x12, 0x89, 0xbc, 0x4e, 0x0d, 0xed, 0x10, 0xf2, 0x36, 0x91, 0xc8, 0x9e, 0x41, 0xdb, 0xb9,
	0xc7, 0x79, 0x66, 0x30, 0x33, 0x7c, 0x9b, 0x18, 0x2d, 0x02, 0x4f, 0x1c, 0xc6, 0xbe, 0x86, 0xae,
	0x23, 0x09, 0x69, 0x2f, 0x74, 0xae, 0x66, 0xbc, 0x41, 0x34, 0x17, 0x7b, 0x6a, 0xd1, 0xf7, 0x6a,
	0xc6, 0x9e, 0xc3, 0xae, 0xe3, 0x69, 0x93, 0x28, 0x13, 0x17, 0x4a, 0x8c, 0x91, 0xef, 0x04, 0x5e,
	0xb8, 0x15, 0x39, 0x81, 0x33, 0x8b, 0xbf, 0xb3, 0xb0, 0xfd, 0x3a, 0x8e, 0x3b, 0x41, 0xe4, 0x40,
	0x9c, 0x06, 0x01, 0xaf, 0x10, 0x6d, 0x43, 0xda, 0x24, 0x59, 0x9a, 0xa8, 0x94, 0x37, 0x5d, 0x43,
	0x95, 0xcd, 0x9e, 0x40, 0x63, 0x22, 0x94, 0x36, 0xf1, 0x95, 0xe4, 0x2d, 0x8a, 0xdb, 0x26, 0xfb,
	0x8d, 0x64, 0x01, 0xb4, 0x2a, 0x17, 0xc9, 0xb6, 0xc9, 0x0d, 0xa5, 0xdb, 0x0a, 0x3f, 0x83, 0x76,
	0x92, 0xa6, 0xc2, 0x5e, 0x77, 0x32, 0xb3, 0x0a, 0x1d, 0xa2, 0xb4, 0x96, 0xe0, 0x1b, 0x69, 0xdb,
	0xb8, 0x45, 0x22, 0xad, 0xae, 0x6b, 0x63, 0x95, 0x68, 0x05, 0xbf, 0xaa, 0x52, 0x7e, 0x40, 0x31,
	0xbd, 0x34, 0xbc, 0x47, 0xb4, 0x26, 0x61, 0xe7, 0x04, 0xb1, 0x17, 0xb7, 0xe4, 0x4a, 0xde, 0x2e,
	0xf1, 0x7a, 0x4b, 0x47, 0x49, 0x1e, 0xc2, 0xc1, 0x1d, 0x32, 0xe5, 0x67, 0x14, 0xb0, 0xb7, 0x1e,
	0x60, 0x6b, 0x08, 0xa1, 0x27, 0x74, 0x6c, 0xdf, 0x0d, 0xcd, 0x87, 0xb8, 0x4e, 0x66, 0x7c, 0x2f,
	0xf0, 0xc2, 0x46, 0xd4, 0x11, 0xfa, 0x0f, 0x0b, 0x8f, 0x1c, 0x6a, 0xef, 0xb5, 0x7a, 0x51, 0x7c,
	0xdf, 0xcd, 0x76, 0x65, 0xdb, 0xcc, 0xd5, 0x39, 0x56, 0x58, 0x60, 0x62, 0x7b, 0xc2, 0x2b, 0xcd,
	0x0f, 0xe8, 0x03, 0x2c, 0x9f, 0x22, 0xf9, 0xce, 0xad, 0x8b, 0xfd, 0x00, 0xfc, 0xbe, 0x18, 0x9a,
	0xd2, 0xc7, 0x34, 0xa5, 0x07, 0x77, 0xc3, 0xec, 0xc4, 0x3e, 0x87, 0x5d, 0x3d, 0xbf, 0x88, 0x17,
	0xc1, 0xb6, 0x78, 0xfe, 0x39, 0x55, 0xd4, 0xd5, 0xf3, 0x8b, 0xea, 0x05, 0xd9, 0xe2, 0x59, 0x00,
	0xcd, 0xa4, 0x28, 0x72, 0x91, 0x19, 0x69, 0x07, 0x94, 0x13, 0x6b, 0x15, 0x62, 0xdf, 0x40, 0x77,
	0xc5, 0x8c, 0xb3, 0xb9, 0xe4, 0x4f, 0xa8, 0xe8, 0xce, 0x0a, 0xfc, 0x76, 0x2e, 0xd9, 0x00, 0xf6,
	0x64, 0xf2, 0x31, 0x5e, 0x27, 0xf7, 0x89, 0xbc, 0x2b, 0x93, 0x8f, 0xa3, 0xdb, 0xfc, 0x43, 0x80,
	0xb1, 0xc2, 0xc4, 0xd8, 0x8b, 0x35, 0xfc, 0xa9, 0x7b, 0x3c, 0x25, 0x32, 0x32, 0xd6, 0x3d, 0x2f,
	0xd2, 0xca, 0xfd, 0x85, 0x73, 0x97, 0xc8, 0xc8, 0xb0, 0x1f, 0x97, 0x9b, 0x90, 0x1a, 0xd4, 0xfc,
	0x30, 0xf0, 0xc3, 0xe6, 0xf0, 0x60, 0x70, 0x7b, 0x21, 0x0e, 0x6c, 0x9b, 0x3a, 0x6a, 0xa7, 0x2b,
	0x5d, 0x6b, 0xf6, 0x3d, 0xd4, 0xe9, 0x3d, 0x68, 0xfe, 0x25, 0x45, 0x1d, 0xae, 0x47, 0xbd, 0x5c,
	0xdd, 0x8d, 0x51, 0x49, 0x3e, 0xfe, 0xd7, 0x83, 0xde, 0xc2, 0x83, 0xba, 0xc8, 0x33, 0x8d, 0xec,
	0x5b, 0xa8, 0x63, 0x66, 0x84, 0xb9, 0xa1, 0x65, 0xd4, 0x1c, 0xf2, 0x8d, 0x5a, 0x25, 0x8f, 0xbd,
	0x70, 0x2b, 0x55, 0xd1, 0xe2, 0xbc, 0xa7, 0xe4, 0x77, 0xd6, 0xe9, 0x36, 0xad, 0x62, 0x03, 0xa8,
	0x09, 0x83, 0x52, 0x73, 0x9f, 0x2a, 0xdd, 0xac, 0xee, 0x68, 0x56, 0x1c, 0x95, 0xca, 0x15, 0x6d,
	0xb9, 0x7b, 0xc4, 0x7f, 0xb1, 0xce, 0xc8, 0x71, 0x58, 0x08, 0x8f, 0x44, 0x36, 0xc9, 0x69, 0xb1,
	0x35, 0x87, 0xfb, 0xeb, 0xdc, 0xd3, 0x6c, 0x92, 0x47, 0xc4, 0x18, 0xfe, 0xed, 0x01, 0xab, 0x52,
	0x9d, 0x4c, 0xa6, 0x67, 0xa8, 0xae, 0xed, 0xa6, 0xf9, 0x09, 0x1e, 0xfd, 0x5e, 0x60, 0xc6, 0xee,
	0xa6, 0x91, 0x85, 0xb9, 0xe9, 0x1f, 0xad, 0xc3, 0x67, 0xee, 0x7f, 0x54, 0x5d, 0xde, 0xf1, 0x67,
	0x6c, 0x04, 0xb5, 0x93, 0x59, 0xae, 0xf1, 0x01, 0x12, 0x3f, 0x43, 0xfd, 0x54, 0x3f, 0xac, 0x8c,
	0xe1, 0x3f, 0x1e, 0x74, 0xab, 0xfe, 0xaa, 0xe6, 0x4e, 0xc0, 0x7f, 0x8d, 0x86, 0x6d, 0xbc, 0xf2,
	0x7e, 0xb0, 0xf1, 0x63, 0x2c, 0x8b, 0x7b, 0x05, 0xf5, 0xf7, 0x34, 0xb5, 0x0f, 0xd4, 0xf9, 0x0d,
	0x1a, 0xaf, 0xd1, 0xb8, 0xf1, 0xdd, 0x38, 0xae, 0xf4, 0xcb, 0xbe, 0x47, 0x2e, 0x31, 0xf8, 0xab,
	0xd0, 0x66, 0x29, 0x77, 0x51, 0xa7, 0xff, 0xfd, 0x77, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x1c,
	0xac, 0x14, 0x79, 0x50, 0x08, 0x00, 0x00,
}
