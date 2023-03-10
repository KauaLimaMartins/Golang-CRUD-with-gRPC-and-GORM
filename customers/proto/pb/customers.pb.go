// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customers.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Customer struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt            string   `protobuf:"bytes,8,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_8869719da4d4f36a, []int{0}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Customer) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Customer) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Customer) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Customer) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Customer) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Customer) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Customer) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type GetAllCustomersReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllCustomersReq) Reset()         { *m = GetAllCustomersReq{} }
func (m *GetAllCustomersReq) String() string { return proto.CompactTextString(m) }
func (*GetAllCustomersReq) ProtoMessage()    {}
func (*GetAllCustomersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8869719da4d4f36a, []int{1}
}

func (m *GetAllCustomersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllCustomersReq.Unmarshal(m, b)
}
func (m *GetAllCustomersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllCustomersReq.Marshal(b, m, deterministic)
}
func (m *GetAllCustomersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllCustomersReq.Merge(m, src)
}
func (m *GetAllCustomersReq) XXX_Size() int {
	return xxx_messageInfo_GetAllCustomersReq.Size(m)
}
func (m *GetAllCustomersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllCustomersReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllCustomersReq proto.InternalMessageInfo

type GetAllCustomersRes struct {
	Customers            []*Customer `protobuf:"bytes,1,rep,name=Customers,proto3" json:"Customers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetAllCustomersRes) Reset()         { *m = GetAllCustomersRes{} }
func (m *GetAllCustomersRes) String() string { return proto.CompactTextString(m) }
func (*GetAllCustomersRes) ProtoMessage()    {}
func (*GetAllCustomersRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_8869719da4d4f36a, []int{2}
}

func (m *GetAllCustomersRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllCustomersRes.Unmarshal(m, b)
}
func (m *GetAllCustomersRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllCustomersRes.Marshal(b, m, deterministic)
}
func (m *GetAllCustomersRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllCustomersRes.Merge(m, src)
}
func (m *GetAllCustomersRes) XXX_Size() int {
	return xxx_messageInfo_GetAllCustomersRes.Size(m)
}
func (m *GetAllCustomersRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllCustomersRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllCustomersRes proto.InternalMessageInfo

func (m *GetAllCustomersRes) GetCustomers() []*Customer {
	if m != nil {
		return m.Customers
	}
	return nil
}

type GetCustomerByIdReq struct {
	CustomerId           string   `protobuf:"bytes,1,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerByIdReq) Reset()         { *m = GetCustomerByIdReq{} }
func (m *GetCustomerByIdReq) String() string { return proto.CompactTextString(m) }
func (*GetCustomerByIdReq) ProtoMessage()    {}
func (*GetCustomerByIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8869719da4d4f36a, []int{3}
}

func (m *GetCustomerByIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerByIdReq.Unmarshal(m, b)
}
func (m *GetCustomerByIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerByIdReq.Marshal(b, m, deterministic)
}
func (m *GetCustomerByIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerByIdReq.Merge(m, src)
}
func (m *GetCustomerByIdReq) XXX_Size() int {
	return xxx_messageInfo_GetCustomerByIdReq.Size(m)
}
func (m *GetCustomerByIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerByIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerByIdReq proto.InternalMessageInfo

func (m *GetCustomerByIdReq) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

type GetCustomerByIdRes struct {
	Customer             *Customer `protobuf:"bytes,1,opt,name=Customer,proto3" json:"Customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetCustomerByIdRes) Reset()         { *m = GetCustomerByIdRes{} }
func (m *GetCustomerByIdRes) String() string { return proto.CompactTextString(m) }
func (*GetCustomerByIdRes) ProtoMessage()    {}
func (*GetCustomerByIdRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_8869719da4d4f36a, []int{4}
}

func (m *GetCustomerByIdRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerByIdRes.Unmarshal(m, b)
}
func (m *GetCustomerByIdRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerByIdRes.Marshal(b, m, deterministic)
}
func (m *GetCustomerByIdRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerByIdRes.Merge(m, src)
}
func (m *GetCustomerByIdRes) XXX_Size() int {
	return xxx_messageInfo_GetCustomerByIdRes.Size(m)
}
func (m *GetCustomerByIdRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerByIdRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerByIdRes proto.InternalMessageInfo

func (m *GetCustomerByIdRes) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type CreateCustomerReq struct {
	Customer             *Customer `protobuf:"bytes,1,opt,name=Customer,proto3" json:"Customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateCustomerReq) Reset()         { *m = CreateCustomerReq{} }
func (m *CreateCustomerReq) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerReq) ProtoMessage()    {}
func (*CreateCustomerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8869719da4d4f36a, []int{5}
}

func (m *CreateCustomerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerReq.Unmarshal(m, b)
}
func (m *CreateCustomerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerReq.Marshal(b, m, deterministic)
}
func (m *CreateCustomerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerReq.Merge(m, src)
}
func (m *CreateCustomerReq) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerReq.Size(m)
}
func (m *CreateCustomerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerReq proto.InternalMessageInfo

func (m *CreateCustomerReq) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type CreateCustomerRes struct {
	Status               uint64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Customer             *Customer `protobuf:"bytes,2,opt,name=Customer,proto3" json:"Customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateCustomerRes) Reset()         { *m = CreateCustomerRes{} }
func (m *CreateCustomerRes) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerRes) ProtoMessage()    {}
func (*CreateCustomerRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_8869719da4d4f36a, []int{6}
}

func (m *CreateCustomerRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerRes.Unmarshal(m, b)
}
func (m *CreateCustomerRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerRes.Marshal(b, m, deterministic)
}
func (m *CreateCustomerRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerRes.Merge(m, src)
}
func (m *CreateCustomerRes) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerRes.Size(m)
}
func (m *CreateCustomerRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerRes proto.InternalMessageInfo

func (m *CreateCustomerRes) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CreateCustomerRes) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type UpdateCustomerReq struct {
	Customer             *Customer `protobuf:"bytes,1,opt,name=Customer,proto3" json:"Customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateCustomerReq) Reset()         { *m = UpdateCustomerReq{} }
func (m *UpdateCustomerReq) String() string { return proto.CompactTextString(m) }
func (*UpdateCustomerReq) ProtoMessage()    {}
func (*UpdateCustomerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8869719da4d4f36a, []int{7}
}

func (m *UpdateCustomerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCustomerReq.Unmarshal(m, b)
}
func (m *UpdateCustomerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCustomerReq.Marshal(b, m, deterministic)
}
func (m *UpdateCustomerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCustomerReq.Merge(m, src)
}
func (m *UpdateCustomerReq) XXX_Size() int {
	return xxx_messageInfo_UpdateCustomerReq.Size(m)
}
func (m *UpdateCustomerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCustomerReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCustomerReq proto.InternalMessageInfo

func (m *UpdateCustomerReq) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type UpdateCustomerRes struct {
	Status               uint64    `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Customer             *Customer `protobuf:"bytes,2,opt,name=Customer,proto3" json:"Customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateCustomerRes) Reset()         { *m = UpdateCustomerRes{} }
func (m *UpdateCustomerRes) String() string { return proto.CompactTextString(m) }
func (*UpdateCustomerRes) ProtoMessage()    {}
func (*UpdateCustomerRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_8869719da4d4f36a, []int{8}
}

func (m *UpdateCustomerRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCustomerRes.Unmarshal(m, b)
}
func (m *UpdateCustomerRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCustomerRes.Marshal(b, m, deterministic)
}
func (m *UpdateCustomerRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCustomerRes.Merge(m, src)
}
func (m *UpdateCustomerRes) XXX_Size() int {
	return xxx_messageInfo_UpdateCustomerRes.Size(m)
}
func (m *UpdateCustomerRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCustomerRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCustomerRes proto.InternalMessageInfo

func (m *UpdateCustomerRes) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateCustomerRes) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type DeleteCustomerReq struct {
	CustomersIds         []string `protobuf:"bytes,1,rep,name=CustomersIds,proto3" json:"CustomersIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCustomerReq) Reset()         { *m = DeleteCustomerReq{} }
func (m *DeleteCustomerReq) String() string { return proto.CompactTextString(m) }
func (*DeleteCustomerReq) ProtoMessage()    {}
func (*DeleteCustomerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8869719da4d4f36a, []int{9}
}

func (m *DeleteCustomerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCustomerReq.Unmarshal(m, b)
}
func (m *DeleteCustomerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCustomerReq.Marshal(b, m, deterministic)
}
func (m *DeleteCustomerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCustomerReq.Merge(m, src)
}
func (m *DeleteCustomerReq) XXX_Size() int {
	return xxx_messageInfo_DeleteCustomerReq.Size(m)
}
func (m *DeleteCustomerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCustomerReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCustomerReq proto.InternalMessageInfo

func (m *DeleteCustomerReq) GetCustomersIds() []string {
	if m != nil {
		return m.CustomersIds
	}
	return nil
}

type DeleteCustomerRes struct {
	Status               uint64   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCustomerRes) Reset()         { *m = DeleteCustomerRes{} }
func (m *DeleteCustomerRes) String() string { return proto.CompactTextString(m) }
func (*DeleteCustomerRes) ProtoMessage()    {}
func (*DeleteCustomerRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_8869719da4d4f36a, []int{10}
}

func (m *DeleteCustomerRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCustomerRes.Unmarshal(m, b)
}
func (m *DeleteCustomerRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCustomerRes.Marshal(b, m, deterministic)
}
func (m *DeleteCustomerRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCustomerRes.Merge(m, src)
}
func (m *DeleteCustomerRes) XXX_Size() int {
	return xxx_messageInfo_DeleteCustomerRes.Size(m)
}
func (m *DeleteCustomerRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCustomerRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCustomerRes proto.InternalMessageInfo

func (m *DeleteCustomerRes) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*Customer)(nil), "Customer")
	proto.RegisterType((*GetAllCustomersReq)(nil), "GetAllCustomersReq")
	proto.RegisterType((*GetAllCustomersRes)(nil), "GetAllCustomersRes")
	proto.RegisterType((*GetCustomerByIdReq)(nil), "GetCustomerByIdReq")
	proto.RegisterType((*GetCustomerByIdRes)(nil), "GetCustomerByIdRes")
	proto.RegisterType((*CreateCustomerReq)(nil), "CreateCustomerReq")
	proto.RegisterType((*CreateCustomerRes)(nil), "CreateCustomerRes")
	proto.RegisterType((*UpdateCustomerReq)(nil), "UpdateCustomerReq")
	proto.RegisterType((*UpdateCustomerRes)(nil), "UpdateCustomerRes")
	proto.RegisterType((*DeleteCustomerReq)(nil), "DeleteCustomerReq")
	proto.RegisterType((*DeleteCustomerRes)(nil), "DeleteCustomerRes")
}

func init() {
	proto.RegisterFile("customers.proto", fileDescriptor_8869719da4d4f36a)
}

var fileDescriptor_8869719da4d4f36a = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6a, 0xe2, 0x40,
	0x14, 0xc6, 0x4d, 0xfc, 0x9b, 0xb3, 0x8b, 0xe2, 0x59, 0x59, 0x42, 0x58, 0x16, 0x19, 0x58, 0x56,
	0x28, 0xe4, 0xc2, 0x16, 0x0a, 0x16, 0x2f, 0xd4, 0xfe, 0x21, 0x50, 0x4a, 0x89, 0xf4, 0xa6, 0x77,
	0xa9, 0x99, 0x0b, 0x21, 0x12, 0xcd, 0x8c, 0x2d, 0x7d, 0xbe, 0xbe, 0x48, 0x1f, 0xa5, 0x64, 0x26,
	0x19, 0x8d, 0x93, 0x96, 0xb6, 0x57, 0x72, 0x7e, 0x9f, 0x5f, 0xce, 0xc9, 0xf9, 0x0e, 0x81, 0xce,
	0x62, 0xcb, 0x78, 0xbc, 0xa2, 0x09, 0x73, 0xd7, 0x49, 0xcc, 0x63, 0xf2, 0x6a, 0x40, 0x6b, 0x96,
	0x31, 0x6c, 0x83, 0xe9, 0x85, 0xb6, 0xd1, 0x37, 0x06, 0x96, 0x6f, 0x7a, 0x21, 0xfe, 0x01, 0xeb,
	0x72, 0x99, 0x30, 0x7e, 0x13, 0xac, 0xa8, 0x6d, 0x0a, 0xbc, 0x03, 0xe8, 0x40, 0xeb, 0x3a, 0xc8,
	0xc4, 0xaa, 0x10, 0x55, 0x8d, 0x3d, 0xa8, 0x5f, 0xac, 0x82, 0x65, 0x64, 0xd7, 0x84, 0x20, 0x8b,
	0xd4, 0x71, 0x1b, 0x30, 0xf6, 0x14, 0x27, 0xa1, 0x5d, 0x97, 0x8e, 0xbc, 0x4e, 0x7b, 0xcd, 0x12,
	0x1a, 0x70, 0x1a, 0x4e, 0xb8, 0xdd, 0x90, 0xbd, 0x14, 0x48, 0xd5, 0xbb, 0x75, 0x98, 0xa9, 0x4d,
	0xa9, 0x2a, 0x90, 0xaa, 0xe7, 0x34, 0xa2, 0x52, 0x6d, 0x49, 0x55, 0x01, 0xd2, 0x03, 0xbc, 0xa2,
	0x7c, 0x12, 0x45, 0xf9, 0x7b, 0x32, 0x9f, 0x6e, 0xc8, 0xb8, 0x84, 0x32, 0xfc, 0x0f, 0x96, 0xaa,
	0x6d, 0xa3, 0x5f, 0x1d, 0xfc, 0x18, 0x5a, 0x6e, 0x4e, 0xfc, 0x9d, 0x46, 0x4e, 0x84, 0x3d, 0xaf,
	0xa7, 0xcf, 0x5e, 0xe8, 0xd3, 0x0d, 0xfe, 0x05, 0xc8, 0x91, 0x5a, 0xe4, 0x1e, 0x21, 0x67, 0x25,
	0x2e, 0x86, 0xff, 0x76, 0x11, 0x08, 0x4f, 0xa1, 0xa7, 0x92, 0xc8, 0x08, 0xba, 0x72, 0x21, 0x4a,
	0xa3, 0x9b, 0xcf, 0x7a, 0x7d, 0xdd, 0xcb, 0xf0, 0x37, 0x34, 0x18, 0x0f, 0xf8, 0x96, 0x09, 0x67,
	0xcd, 0xcf, 0xaa, 0xc2, 0x33, 0xcd, 0x0f, 0xe7, 0x91, 0x11, 0x7c, 0x6f, 0x9e, 0x43, 0xaf, 0x98,
	0x67, 0x5e, 0x98, 0x67, 0xfe, 0xa5, 0x79, 0x4e, 0xa1, 0x2b, 0x43, 0xdf, 0x9f, 0x87, 0xc0, 0x4f,
	0x15, 0x9a, 0x17, 0xca, 0x4c, 0x2d, 0xbf, 0xc0, 0xc8, 0x91, 0x6e, 0x7c, 0x77, 0x39, 0xc3, 0x17,
	0x13, 0x3a, 0xf9, 0xff, 0xe6, 0x34, 0x79, 0x5c, 0x2e, 0x28, 0x8e, 0xa1, 0x73, 0x70, 0x4b, 0xf8,
	0xcb, 0xd5, 0x6f, 0xce, 0x29, 0x81, 0x8c, 0x54, 0x32, 0xfb, 0xfe, 0x55, 0x48, 0xfb, 0xc1, 0x75,
	0x39, 0x25, 0x30, 0xb5, 0x8f, 0xa0, 0x5d, 0xcc, 0x16, 0xd1, 0xd5, 0x0e, 0xc5, 0xd1, 0x59, 0xe6,
	0x2d, 0xe6, 0x80, 0xe8, 0x6a, 0xa1, 0x3a, 0x3a, 0xcb, 0xbc, 0xc5, 0xb5, 0x21, 0xba, 0x5a, 0x00,
	0x8e, 0xce, 0x18, 0xa9, 0x4c, 0x9b, 0xf7, 0x75, 0xf1, 0xfd, 0x79, 0x68, 0x88, 0x9f, 0xe3, 0xb7,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0xfb, 0x4b, 0x4a, 0x99, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerServiceClient interface {
	GetAllCustomers(ctx context.Context, in *GetAllCustomersReq, opts ...grpc.CallOption) (*GetAllCustomersRes, error)
	GetCustomerById(ctx context.Context, in *GetCustomerByIdReq, opts ...grpc.CallOption) (*GetCustomerByIdRes, error)
	CreateCustomer(ctx context.Context, in *CreateCustomerReq, opts ...grpc.CallOption) (*CreateCustomerRes, error)
	UpdateCustomer(ctx context.Context, in *UpdateCustomerReq, opts ...grpc.CallOption) (*UpdateCustomerRes, error)
	DeleteCustomer(ctx context.Context, in *DeleteCustomerReq, opts ...grpc.CallOption) (*DeleteCustomerRes, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) GetAllCustomers(ctx context.Context, in *GetAllCustomersReq, opts ...grpc.CallOption) (*GetAllCustomersRes, error) {
	out := new(GetAllCustomersRes)
	err := c.cc.Invoke(ctx, "/CustomerService/GetAllCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomerById(ctx context.Context, in *GetCustomerByIdReq, opts ...grpc.CallOption) (*GetCustomerByIdRes, error) {
	out := new(GetCustomerByIdRes)
	err := c.cc.Invoke(ctx, "/CustomerService/GetCustomerById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) CreateCustomer(ctx context.Context, in *CreateCustomerReq, opts ...grpc.CallOption) (*CreateCustomerRes, error) {
	out := new(CreateCustomerRes)
	err := c.cc.Invoke(ctx, "/CustomerService/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) UpdateCustomer(ctx context.Context, in *UpdateCustomerReq, opts ...grpc.CallOption) (*UpdateCustomerRes, error) {
	out := new(UpdateCustomerRes)
	err := c.cc.Invoke(ctx, "/CustomerService/UpdateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) DeleteCustomer(ctx context.Context, in *DeleteCustomerReq, opts ...grpc.CallOption) (*DeleteCustomerRes, error) {
	out := new(DeleteCustomerRes)
	err := c.cc.Invoke(ctx, "/CustomerService/DeleteCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
type CustomerServiceServer interface {
	GetAllCustomers(context.Context, *GetAllCustomersReq) (*GetAllCustomersRes, error)
	GetCustomerById(context.Context, *GetCustomerByIdReq) (*GetCustomerByIdRes, error)
	CreateCustomer(context.Context, *CreateCustomerReq) (*CreateCustomerRes, error)
	UpdateCustomer(context.Context, *UpdateCustomerReq) (*UpdateCustomerRes, error)
	DeleteCustomer(context.Context, *DeleteCustomerReq) (*DeleteCustomerRes, error)
}

// UnimplementedCustomerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (*UnimplementedCustomerServiceServer) GetAllCustomers(ctx context.Context, req *GetAllCustomersReq) (*GetAllCustomersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCustomers not implemented")
}
func (*UnimplementedCustomerServiceServer) GetCustomerById(ctx context.Context, req *GetCustomerByIdReq) (*GetCustomerByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerById not implemented")
}
func (*UnimplementedCustomerServiceServer) CreateCustomer(ctx context.Context, req *CreateCustomerReq) (*CreateCustomerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (*UnimplementedCustomerServiceServer) UpdateCustomer(ctx context.Context, req *UpdateCustomerReq) (*UpdateCustomerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (*UnimplementedCustomerServiceServer) DeleteCustomer(ctx context.Context, req *DeleteCustomerReq) (*DeleteCustomerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}

func RegisterCustomerServiceServer(s *grpc.Server, srv CustomerServiceServer) {
	s.RegisterService(&_CustomerService_serviceDesc, srv)
}

func _CustomerService_GetAllCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCustomersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetAllCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/GetAllCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetAllCustomers(ctx, req.(*GetAllCustomersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetCustomerById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomerById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/GetCustomerById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomerById(ctx, req.(*GetCustomerByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, req.(*CreateCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/UpdateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).UpdateCustomer(ctx, req.(*UpdateCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_DeleteCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).DeleteCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/DeleteCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).DeleteCustomer(ctx, req.(*DeleteCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllCustomers",
			Handler:    _CustomerService_GetAllCustomers_Handler,
		},
		{
			MethodName: "GetCustomerById",
			Handler:    _CustomerService_GetCustomerById_Handler,
		},
		{
			MethodName: "CreateCustomer",
			Handler:    _CustomerService_CreateCustomer_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _CustomerService_UpdateCustomer_Handler,
		},
		{
			MethodName: "DeleteCustomer",
			Handler:    _CustomerService_DeleteCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customers.proto",
}
