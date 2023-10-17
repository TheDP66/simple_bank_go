// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

// Group multiple proto files of the same topic together

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type User struct {
	Username             string               `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName             string               `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PasswordChangeAt     *timestamp.Timestamp `protobuf:"bytes,4,opt,name=password_change_at,json=passwordChangeAt,proto3" json:"password_change_at,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPasswordChangeAt() *timestamp.Timestamp {
	if m != nil {
		return m.PasswordChangeAt
	}
	return nil
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0x31, 0x4b, 0x03, 0x31,
	0x1c, 0xc5, 0xb9, 0xda, 0x8a, 0xf7, 0x77, 0x91, 0xe0, 0x70, 0x9c, 0x83, 0x45, 0x44, 0x3a, 0x25,
	0xa0, 0x50, 0x70, 0xac, 0x3a, 0x38, 0x89, 0x94, 0xba, 0xb8, 0x84, 0xe4, 0xfa, 0x6f, 0xee, 0x30,
	0xb9, 0x84, 0x24, 0x87, 0xdf, 0xd5, 0x4f, 0x23, 0x97, 0x78, 0xae, 0x1d, 0xdf, 0xfb, 0xbd, 0xf7,
	0x78, 0x00, 0x43, 0x40, 0x4f, 0x9d, 0xb7, 0xd1, 0x92, 0x99, 0x93, 0xf5, 0xb5, 0xb2, 0x56, 0x69,
	0x64, 0xc9, 0x91, 0xc3, 0x81, 0xc5, 0xce, 0x60, 0x88, 0xc2, 0xb8, 0x1c, 0xba, 0xf9, 0x29, 0x60,
	0xfe, 0x11, 0xd0, 0x93, 0x1a, 0xce, 0xc6, 0x6e, 0x2f, 0x0c, 0x56, 0xc5, 0xb2, 0x58, 0x95, 0xdb,
	0x7f, 0x4d, 0xae, 0xa0, 0x3c, 0x0c, 0x5a, 0xf3, 0x04, 0x67, 0x19, 0x8e, 0xc6, 0xdb, 0x08, 0x2f,
	0x61, 0x81, 0x46, 0x74, 0xba, 0x3a, 0x49, 0x20, 0x0b, 0xf2, 0x0a, 0xc4, 0x89, 0x10, 0xbe, 0xad,
	0xdf, 0xf3, 0xa6, 0x15, 0xbd, 0x42, 0x2e, 0x62, 0x35, 0x5f, 0x16, 0xab, 0xf3, 0xfb, 0x9a, 0xe6,
	0x57, 0x74, 0x7a, 0x45, 0x77, 0xd3, 0xab, 0xed, 0xc5, 0xd4, 0x7a, 0x4e, 0xa5, 0x4d, 0x24, 0x8f,
	0x00, 0x8d, 0x47, 0x11, 0x71, 0x3f, 0x2e, 0x2c, 0x8e, 0x2e, 0x94, 0x7f, 0xe9, 0x4d, 0x7c, 0xba,
	0xfb, 0xbc, 0x55, 0x5d, 0x6c, 0x07, 0x49, 0x1b, 0x6b, 0xd8, 0xae, 0xc5, 0x97, 0xf7, 0xf5, 0x9a,
	0x85, 0xce, 0x38, 0x8d, 0x5c, 0x8a, 0xfe, 0x8b, 0x2b, 0xcb, 0x9c, 0x94, 0xa7, 0x69, 0xe6, 0xe1,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x43, 0x77, 0xa2, 0x10, 0x3e, 0x01, 0x00, 0x00,
}
