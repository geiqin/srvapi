// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eatOutOrder.proto

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

func init() {
	proto.RegisterFile("eatOutOrder.proto", fileDescriptor_264817d298c7294f)
}

var fileDescriptor_264817d298c7294f = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4d, 0x2c, 0xf1,
	0x2f, 0x2d, 0xf1, 0x2f, 0x4a, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4b,
	0x4f, 0xcd, 0x2c, 0xcc, 0xcc, 0xd3, 0x2b, 0x2e, 0x2a, 0xd3, 0xcb, 0x2f, 0x4a, 0xd1, 0x2b, 0x28,
	0xca, 0x2c, 0x4b, 0x2c, 0x49, 0x95, 0xe2, 0xce, 0x47, 0x28, 0x92, 0xe2, 0x49, 0x2a, 0xad, 0xcc,
	0xcc, 0x4b, 0x87, 0xf0, 0x8c, 0xce, 0x30, 0x72, 0x89, 0xf8, 0x56, 0xba, 0x22, 0x8c, 0x0a, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x8a, 0xe2, 0x62, 0x77, 0xce, 0xcf, 0x4b, 0xcb, 0x2c, 0xca,
	0x15, 0x52, 0xd5, 0xc3, 0x6e, 0xae, 0x9e, 0x13, 0xd8, 0xa4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x29, 0x35, 0x42, 0xca, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95, 0x18, 0x84, 0x22, 0xb8,
	0xd8, 0x82, 0x4b, 0x93, 0x72, 0x33, 0x4b, 0x88, 0x35, 0x1a, 0xa7, 0x32, 0xb0, 0x93, 0x11, 0x26,
	0x27, 0xb1, 0x81, 0x7d, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x70, 0x13, 0xce, 0x1d,
	0x01, 0x00, 0x00,
}
