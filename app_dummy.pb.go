// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app_dummy.proto

package lazyledger

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

type DummyAppTransaction struct {
	Puts                 map[string]string `protobuf:"bytes,1,rep,name=puts" json:"puts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DummyAppTransaction) Reset()         { *m = DummyAppTransaction{} }
func (m *DummyAppTransaction) String() string { return proto.CompactTextString(m) }
func (*DummyAppTransaction) ProtoMessage()    {}
func (*DummyAppTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fabccd89524743b, []int{0}
}

func (m *DummyAppTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DummyAppTransaction.Unmarshal(m, b)
}
func (m *DummyAppTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DummyAppTransaction.Marshal(b, m, deterministic)
}
func (m *DummyAppTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DummyAppTransaction.Merge(m, src)
}
func (m *DummyAppTransaction) XXX_Size() int {
	return xxx_messageInfo_DummyAppTransaction.Size(m)
}
func (m *DummyAppTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_DummyAppTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_DummyAppTransaction proto.InternalMessageInfo

func (m *DummyAppTransaction) GetPuts() map[string]string {
	if m != nil {
		return m.Puts
	}
	return nil
}

func init() {
	proto.RegisterType((*DummyAppTransaction)(nil), "lazyledger.DummyAppTransaction")
	proto.RegisterMapType((map[string]string)(nil), "lazyledger.DummyAppTransaction.PutsEntry")
}

func init() { proto.RegisterFile("app_dummy.proto", fileDescriptor_6fabccd89524743b) }

var fileDescriptor_6fabccd89524743b = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2c, 0x28, 0x88,
	0x4f, 0x29, 0xcd, 0xcd, 0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x49, 0xac,
	0xaa, 0xcc, 0x49, 0x4d, 0x49, 0x4f, 0x2d, 0x52, 0xea, 0x65, 0xe4, 0x12, 0x76, 0x01, 0xc9, 0x39,
	0x16, 0x14, 0x84, 0x14, 0x25, 0xe6, 0x15, 0x27, 0x26, 0x97, 0x64, 0xe6, 0xe7, 0x09, 0xd9, 0x72,
	0xb1, 0x14, 0x94, 0x96, 0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x69, 0xea, 0x21, 0xb4,
	0xe8, 0x61, 0x51, 0xae, 0x17, 0x50, 0x5a, 0x52, 0xec, 0x9a, 0x57, 0x52, 0x54, 0x19, 0x04, 0xd6,
	0x26, 0x65, 0xce, 0xc5, 0x09, 0x17, 0x12, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25,
	0x98, 0xc0, 0x62, 0x10, 0x8e, 0x15, 0x93, 0x05, 0x23, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xfd,
	0x97, 0xf6, 0xad, 0x00, 0x00, 0x00,
}
