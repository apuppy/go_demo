// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

package protobuf

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

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=protobuf.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("protobuf.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "protobuf.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "protobuf.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "protobuf.AddressBook")
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor_1eb1a68c9dd6d429) }

var fileDescriptor_1eb1a68c9dd6d429 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdf, 0xa4, 0x69, 0x68, 0x27, 0x2f, 0x25, 0x0e, 0x22, 0x21, 0x08, 0x86, 0x9e, 0x02,
	0x42, 0x0a, 0x55, 0xf0, 0xe4, 0xc1, 0x42, 0x41, 0xd1, 0xda, 0xb2, 0xb4, 0x78, 0x94, 0x84, 0x8c,
	0x35, 0x34, 0xc9, 0x2e, 0xd9, 0xed, 0xa1, 0x9f, 0xcd, 0x2f, 0x27, 0xd9, 0x4d, 0x55, 0xc4, 0xd3,
	0xce, 0x9f, 0x1f, 0xcf, 0x3c, 0xfb, 0xc0, 0x49, 0x9a, 0xe7, 0x0d, 0x49, 0x99, 0x71, 0xbe, 0x4b,
	0x44, 0xc3, 0x15, 0xc7, 0x81, 0x7e, 0xb2, 0xfd, 0x5b, 0x78, 0xb1, 0xe5, 0x7c, 0x5b, 0xd2, 0xe4,
	0x38, 0x98, 0xa8, 0xa2, 0x22, 0xa9, 0xd2, 0x4a, 0x18, 0x74, 0xfc, 0x61, 0x83, 0xbb, 0xa2, 0x46,
	0xf2, 0x1a, 0x11, 0x9c, 0x3a, 0xad, 0x28, 0xb0, 0x22, 0x2b, 0x1e, 0x32, 0x5d, 0xe3, 0x08, 0xec,
	0x22, 0x0f, 0xec, 0xc8, 0x8a, 0xfb, 0xcc, 0x2e, 0x72, 0x3c, 0x85, 0x3e, 0x55, 0x69, 0x51, 0x06,
	0x3d, 0x0d, 0x99, 0x06, 0xaf, 0xc1, 0x15, 0xef, 0xbc, 0x26, 0x19, 0x38, 0x51, 0x2f, 0xf6, 0xa6,
	0xe7, 0xc9, 0xf1, 0x5e, 0x62, 0xb4, 0x93, 0x55, 0xbb, 0x7e, 0xde, 0x57, 0x19, 0x35, 0xac, 0x63,
	0xf1, 0x16, 0xfe, 0x97, 0xa9, 0x54, 0xaf, 0x7b, 0x91, 0xa7, 0x8a, 0xf2, 0xa0, 0x1f, 0x59, 0xb1,
	0x37, 0x0d, 0x13, 0x63, 0xf9, 0x5b, 0x62, 0x7d, 0xb4, 0xcc, 0xbc, 0x96, 0xdf, 0x18, 0x3c, 0xdc,
	0x80, 0xf7, 0x43, 0x15, 0xcf, 0xc0, 0xad, 0x75, 0xd5, 0xf9, 0xef, 0x3a, 0x4c, 0xc0, 0x51, 0x07,
	0x41, 0xfa, 0x0f, 0xa3, 0x69, 0xf8, 0xb7, 0xb3, 0xf5, 0x41, 0x10, 0xd3, 0xdc, 0xf8, 0x12, 0x86,
	0x5f, 0x23, 0x04, 0x70, 0x17, 0xcb, 0xd9, 0xc3, 0xd3, 0xdc, 0xff, 0x87, 0x03, 0x70, 0xee, 0x97,
	0x8b, 0xb9, 0x6f, 0xb5, 0xd5, 0xcb, 0x92, 0x3d, 0xfa, 0xf6, 0xf8, 0x06, 0xbc, 0x3b, 0x93, 0xfe,
	0x8c, 0xf3, 0x1d, 0xc6, 0xe0, 0x0a, 0xe2, 0xa2, 0x6c, 0x33, 0x6c, 0x73, 0xf0, 0x7f, 0x5f, 0x63,
	0xdd, 0x3e, 0x73, 0xf5, 0xe2, 0xea, 0x33, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x67, 0x16, 0x6e, 0xbd,
	0x01, 0x00, 0x00,
}