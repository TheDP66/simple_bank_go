// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_simple_bank.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() { proto.RegisterFile("service_simple_bank.proto", fileDescriptor_4b64e86b71303064) }

var fileDescriptor_4b64e86b71303064 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0xe9, 0x2e, 0x08, 0x46, 0x71, 0x21, 0xce, 0x28, 0x56, 0x0f, 0x71, 0x10, 0x0f, 0x03,
	0xdb, 0xb0, 0x55, 0xe6, 0x30, 0x17, 0xed, 0xee, 0x82, 0x08, 0x0a, 0xe3, 0x9f, 0xb9, 0x28, 0x52,
	0xd2, 0xce, 0x6b, 0x1b, 0xa6, 0x93, 0xc4, 0x24, 0x1d, 0xef, 0x22, 0x1e, 0x3d, 0x8c, 0x1f, 0xc1,
	0x6f, 0xe3, 0xd5, 0xaf, 0xe0, 0x07, 0x91, 0x24, 0xeb, 0xb4, 0x56, 0x4f, 0xe5, 0x7d, 0x9e, 0xa7,
	0x79, 0x7f, 0xc9, 0xfb, 0xa2, 0x5b, 0x06, 0xf4, 0x96, 0x97, 0x90, 0x1b, 0xbe, 0x51, 0x0d, 0xe4,
	0x05, 0x13, 0xeb, 0x44, 0x69, 0x69, 0x25, 0x3e, 0x50, 0x45, 0x7c, 0xa7, 0x92, 0xb2, 0x6a, 0x80,
	0x32, 0xc5, 0x29, 0x13, 0x42, 0x5a, 0x66, 0xb9, 0x14, 0x26, 0x24, 0xe2, 0xf0, 0x29, 0x8f, 0x2b,
	0x10, 0xc7, 0x52, 0x81, 0x60, 0x8a, 0x6f, 0x53, 0x2a, 0x95, 0xcf, 0xfc, 0x27, 0x3f, 0xd6, 0xaa,
	0xcc, 0x4b, 0x0d, 0xcc, 0x42, 0xde, 0x1a, 0xd0, 0x7d, 0xb9, 0x55, 0xab, 0x81, 0x3c, 0x72, 0x72,
	0x23, 0x2b, 0x2e, 0x7a, 0x6a, 0xfa, 0xe3, 0x10, 0xa1, 0x57, 0x9e, 0xf5, 0x94, 0x89, 0x35, 0xfe,
	0x1a, 0x21, 0x74, 0xe6, 0x4f, 0x5c, 0x1a, 0xd0, 0x78, 0x9c, 0xa8, 0x22, 0xe9, 0xea, 0x97, 0xf0,
	0xa1, 0x05, 0x63, 0xe3, 0x1b, 0x43, 0xd9, 0x28, 0x29, 0x0c, 0x4c, 0x5e, 0xec, 0xb2, 0x87, 0xf8,
	0x28, 0x18, 0x44, 0xc0, 0x47, 0xe2, 0x1a, 0xc5, 0x77, 0x97, 0x06, 0x88, 0xad, 0xb9, 0x21, 0xd9,
	0xe2, 0x29, 0xb1, 0x92, 0x04, 0x64, 0xc2, 0xf6, 0x91, 0x4f, 0x3f, 0x7f, 0x7d, 0x3b, 0x18, 0x4d,
	0x8e, 0xe8, 0xf6, 0x84, 0xf6, 0xee, 0x33, 0x8f, 0xa6, 0xf8, 0x73, 0x84, 0xd0, 0xd2, 0xdf, 0xa5,
	0x03, 0xea, 0xea, 0xbf, 0x80, 0xfa, 0xf2, 0x05, 0xd0, 0x93, 0x5d, 0x36, 0xc5, 0x57, 0x82, 0x11,
	0x60, 0x6e, 0x0f, 0x61, 0xda, 0xce, 0x0c, 0x18, 0xa9, 0xc7, 0xe8, 0xbd, 0x9f, 0xc3, 0xf8, 0x1e,
	0xa1, 0xcb, 0xcf, 0xdc, 0xdb, 0x79, 0x8a, 0x91, 0x6b, 0xb7, 0x2f, 0xff, 0x40, 0x8c, 0x07, 0xea,
	0x05, 0x03, 0xdf, 0x65, 0xcf, 0x31, 0xf2, 0x7a, 0x40, 0x78, 0x34, 0x44, 0x68, 0xf6, 0x1e, 0x61,
	0x62, 0x45, 0x2a, 0xb0, 0x84, 0x95, 0x25, 0x18, 0x43, 0xac, 0x5c, 0x83, 0x20, 0xf7, 0x89, 0x86,
	0xf7, 0x1a, 0x4c, 0x1d, 0x6a, 0x8f, 0x79, 0x7d, 0x72, 0xcd, 0x61, 0x76, 0xf3, 0x9c, 0x47, 0xd3,
	0xd3, 0x2f, 0xd1, 0x2e, 0x7b, 0x87, 0xdf, 0xa2, 0xa3, 0x30, 0x51, 0xe2, 0x46, 0xea, 0x9a, 0x4c,
	0xce, 0xd0, 0xd5, 0xf3, 0x9a, 0xe9, 0x0d, 0x23, 0x8b, 0xd6, 0x6a, 0x86, 0xe3, 0xda, 0x5a, 0x65,
	0xe6, 0x94, 0x56, 0xdc, 0xd6, 0x6d, 0x91, 0x94, 0x72, 0x43, 0x5f, 0xd7, 0x70, 0xbe, 0x98, 0xcd,
	0xe2, 0x9b, 0x2b, 0x9f, 0x9c, 0xcd, 0x94, 0x8b, 0x3e, 0xae, 0x36, 0x8c, 0x37, 0x2e, 0x90, 0x1e,
	0x9e, 0x24, 0xe9, 0x9b, 0x7b, 0xff, 0xfe, 0x41, 0x7b, 0xbb, 0x9e, 0x57, 0x92, 0xaa, 0xa2, 0xb8,
	0xe4, 0x97, 0xeb, 0xc1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x90, 0x50, 0x02, 0x0f, 0x03,
	0x00, 0x00,
}
