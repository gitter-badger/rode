// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1alpha1/rode.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
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

type Foo struct {
	Bar                  string   `protobuf:"bytes,1,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce02dbf0dfcdb9d7, []int{0}
}

func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetBar() string {
	if m != nil {
		return m.Bar
	}
	return ""
}

func init() {
	proto.RegisterType((*Foo)(nil), "rode.v1alpha1.Foo")
}

func init() { proto.RegisterFile("proto/v1alpha1/rode.proto", fileDescriptor_ce02dbf0dfcdb9d7) }

var fileDescriptor_ce02dbf0dfcdb9d7 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x50, 0x3d, 0x4b, 0xc4, 0x40,
	0x10, 0x35, 0x1c, 0x08, 0x2e, 0x08, 0x12, 0xc4, 0x8f, 0xab, 0xe4, 0x2a, 0x9b, 0xec, 0x12, 0x6d,
	0xc4, 0xf2, 0x0e, 0x6c, 0x85, 0xd8, 0xd9, 0x4d, 0x26, 0xe3, 0x65, 0x61, 0xbd, 0x59, 0x67, 0x27,
	0xd7, 0xf8, 0x33, 0xfc, 0xc3, 0xe2, 0xc6, 0x2d, 0x04, 0x8b, 0xeb, 0xde, 0xbc, 0x79, 0x6f, 0xe6,
	0xf1, 0xcc, 0x75, 0x14, 0x56, 0x76, 0xfb, 0x16, 0x42, 0x1c, 0xa1, 0x75, 0xc2, 0x03, 0xd9, 0xcc,
	0xd5, 0xa7, 0x19, 0x97, 0xcd, 0xb2, 0xc9, 0xec, 0x40, 0x31, 0xb9, 0xad, 0xc0, 0x1b, 0x41, 0x72,
	0xc5, 0xdb, 0x93, 0x42, 0x5b, 0xd8, 0xd9, 0xbd, 0xba, 0x34, 0x8b, 0x27, 0xe6, 0xfa, 0xcc, 0x2c,
	0x7a, 0x90, 0xab, 0xea, 0xa6, 0xba, 0x3d, 0xe9, 0x7e, 0xe0, 0xdd, 0x57, 0x65, 0xce, 0x3b, 0x1e,
	0x68, 0xc3, 0x21, 0x10, 0x2a, 0xcb, 0x0b, 0xc9, 0xde, 0x23, 0xd5, 0x9f, 0xe6, 0x62, 0x0d, 0x8a,
	0xe3, 0x46, 0x08, 0x94, 0x9e, 0x11, 0x27, 0x11, 0xda, 0x21, 0xa5, 0xda, 0xda, 0x72, 0xfb, 0xf7,
	0x97, 0xfd, 0x5f, 0xd8, 0xd1, 0xc7, 0x44, 0x49, 0x97, 0xee, 0x60, 0x7d, 0x8a, 0xbc, 0x4b, 0xb4,
	0x3a, 0x5a, 0x3f, 0xbe, 0x3e, 0x6c, 0xbd, 0x8e, 0x53, 0x6f, 0x91, 0xdf, 0x5d, 0xf0, 0xa0, 0xe2,
	0x39, 0xb7, 0xd1, 0x60, 0x09, 0xda, 0xa4, 0x39, 0xa9, 0xfb, 0xdb, 0x59, 0x7f, 0x9c, 0xe7, 0xfb,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0x04, 0x6b, 0x7c, 0x4c, 0x01, 0x00, 0x00,
}
