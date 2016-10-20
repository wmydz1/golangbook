// Code generated by protoc-gen-go.
// source: addressbook.proto
// DO NOT EDIT!

/*
Package tutorial is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Person
	AddressBook
*/
package tutorial

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// [START messages]
type Person struct {
	Name   string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id     int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email  string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Our address book file is just one of these.
type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdf, 0xa4, 0x69, 0x68, 0xa7, 0x50, 0xd2, 0xa5, 0xbc, 0x84, 0xe2, 0xa1, 0x04, 0x0f,
	0x01, 0x61, 0x0f, 0x55, 0xf0, 0x6c, 0x20, 0xa8, 0x68, 0x4d, 0x58, 0x14, 0xcf, 0xf9, 0xb3, 0xd6,
	0x60, 0x92, 0x09, 0xd9, 0x04, 0xf4, 0x2b, 0xf9, 0x19, 0x3d, 0xb8, 0xd9, 0xc4, 0x12, 0xc4, 0xdb,
	0x33, 0x33, 0x3f, 0x9e, 0x9d, 0x79, 0x16, 0x56, 0x51, 0x9a, 0xd6, 0x5c, 0x88, 0x18, 0xf1, 0x8d,
	0x56, 0x35, 0x36, 0x48, 0x66, 0x4d, 0xdb, 0x60, 0x9d, 0x45, 0xb9, 0xf3, 0xa5, 0x81, 0x19, 0xf2,
	0x5a, 0x60, 0x49, 0x08, 0x18, 0x65, 0x54, 0x70, 0x5b, 0xdb, 0x6a, 0xee, 0x9c, 0x29, 0x4d, 0x96,
	0xa0, 0x67, 0xa9, 0xad, 0xcb, 0xce, 0x94, 0x49, 0x45, 0xd6, 0x30, 0xe5, 0x45, 0x94, 0xe5, 0xf6,
	0x44, 0x41, 0x7d, 0x41, 0x2e, 0xc0, 0xac, 0x5e, 0xb1, 0xe4, 0xc2, 0x36, 0xb6, 0x13, 0x77, 0xb1,
	0x3b, 0xa1, 0x3f, 0xfe, 0xb4, 0xf7, 0xa6, 0x61, 0x37, 0x7e, 0x68, 0x8b, 0x98, 0xd7, 0x6c, 0x60,
	0x37, 0x4f, 0xb0, 0x18, 0xb5, 0xc9, 0x7f, 0x30, 0x4b, 0xa5, 0x86, 0x05, 0x86, 0x8a, 0x50, 0x30,
	0x9a, 0x8f, 0x8a, 0xab, 0x25, 0x96, 0xbb, 0xcd, 0xdf, 0xd6, 0x8f, 0x92, 0x60, 0x8a, 0x73, 0xce,
	0x60, 0x7e, 0x6c, 0x11, 0x00, 0x73, 0x1f, 0x78, 0xb7, 0xf7, 0xbe, 0xf5, 0x8f, 0xcc, 0xc0, 0xb8,
	0x09, 0xf6, 0xbe, 0xa5, 0x75, 0xea, 0x39, 0x60, 0x77, 0x96, 0xee, 0x5c, 0xc2, 0xe2, 0xaa, 0x4f,
	0xc7, 0x93, 0xe9, 0x10, 0x57, 0x1e, 0xc2, 0xb1, 0xca, 0xbb, 0x10, 0xba, 0x43, 0xac, 0xdf, 0xaf,
	0xb1, 0x61, 0xee, 0x85, 0xb0, 0x4e, 0xb0, 0xa0, 0x39, 0x1e, 0x10, 0x93, 0xe4, 0x88, 0x79, 0xab,
	0x91, 0x5d, 0xd8, 0x65, 0x2d, 0x3e, 0xf5, 0xd3, 0x6b, 0xc4, 0x43, 0xce, 0xa9, 0xaa, 0xe3, 0xf6,
	0x85, 0xfa, 0xef, 0x51, 0x21, 0x4d, 0x04, 0x1d, 0xc1, 0xb1, 0xa9, 0xbe, 0xe6, 0xfc, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0xb2, 0x9b, 0x09, 0x03, 0xaf, 0x01, 0x00, 0x00,
}
