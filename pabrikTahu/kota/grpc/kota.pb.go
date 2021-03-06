// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kota.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	kota.proto

It has these top-level messages:
	AddKotaReq
	ReadKotaByNamaReq
	ReadKotaByNamaResp
	ReadKotaResp
	UpdateKotaReq
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddKotaReq struct {
	IDKota    int64  `protobuf:"varint,1,opt,name=IDKota" json:"IDKota,omitempty"`
	NamaKota  string `protobuf:"bytes,2,opt,name=NamaKota" json:"NamaKota,omitempty"`
	Status    int32  `protobuf:"varint,3,opt,name=Status" json:"Status,omitempty"`
	CreatedBy string `protobuf:"bytes,4,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn string `protobuf:"bytes,5,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdatedOn string `protobuf:"bytes,6,opt,name=UpdatedOn" json:"UpdatedOn,omitempty"`
	UpdatedBy string `protobuf:"bytes,7,opt,name=UpdatedBy" json:"UpdatedBy,omitempty"`
}

func (m *AddKotaReq) Reset()                    { *m = AddKotaReq{} }
func (m *AddKotaReq) String() string            { return proto.CompactTextString(m) }
func (*AddKotaReq) ProtoMessage()               {}
func (*AddKotaReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddKotaReq) GetIDKota() int64 {
	if m != nil {
		return m.IDKota
	}
	return 0
}

func (m *AddKotaReq) GetNamaKota() string {
	if m != nil {
		return m.NamaKota
	}
	return ""
}

func (m *AddKotaReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddKotaReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *AddKotaReq) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *AddKotaReq) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *AddKotaReq) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

type ReadKotaByNamaReq struct {
	NamaKota string `protobuf:"bytes,1,opt,name=NamaKota" json:"NamaKota,omitempty"`
}

func (m *ReadKotaByNamaReq) Reset()                    { *m = ReadKotaByNamaReq{} }
func (m *ReadKotaByNamaReq) String() string            { return proto.CompactTextString(m) }
func (*ReadKotaByNamaReq) ProtoMessage()               {}
func (*ReadKotaByNamaReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadKotaByNamaReq) GetNamaKota() string {
	if m != nil {
		return m.NamaKota
	}
	return ""
}

type ReadKotaByNamaResp struct {
	IDKota    int64  `protobuf:"varint,1,opt,name=IDKota" json:"IDKota,omitempty"`
	NamaKota  string `protobuf:"bytes,2,opt,name=NamaKota" json:"NamaKota,omitempty"`
	Status    int32  `protobuf:"varint,3,opt,name=Status" json:"Status,omitempty"`
	CreatedBy string `protobuf:"bytes,4,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn string `protobuf:"bytes,5,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdatedOn string `protobuf:"bytes,6,opt,name=UpdatedOn" json:"UpdatedOn,omitempty"`
	UpdatedBy string `protobuf:"bytes,7,opt,name=UpdatedBy" json:"UpdatedBy,omitempty"`
}

func (m *ReadKotaByNamaResp) Reset()                    { *m = ReadKotaByNamaResp{} }
func (m *ReadKotaByNamaResp) String() string            { return proto.CompactTextString(m) }
func (*ReadKotaByNamaResp) ProtoMessage()               {}
func (*ReadKotaByNamaResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadKotaByNamaResp) GetIDKota() int64 {
	if m != nil {
		return m.IDKota
	}
	return 0
}

func (m *ReadKotaByNamaResp) GetNamaKota() string {
	if m != nil {
		return m.NamaKota
	}
	return ""
}

func (m *ReadKotaByNamaResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReadKotaByNamaResp) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *ReadKotaByNamaResp) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *ReadKotaByNamaResp) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *ReadKotaByNamaResp) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

type ReadKotaResp struct {
	AllKota []*ReadKotaByNamaResp `protobuf:"bytes,1,rep,name=allKota" json:"allKota,omitempty"`
}

func (m *ReadKotaResp) Reset()                    { *m = ReadKotaResp{} }
func (m *ReadKotaResp) String() string            { return proto.CompactTextString(m) }
func (*ReadKotaResp) ProtoMessage()               {}
func (*ReadKotaResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadKotaResp) GetAllKota() []*ReadKotaByNamaResp {
	if m != nil {
		return m.AllKota
	}
	return nil
}

type UpdateKotaReq struct {
	IDKota    int64  `protobuf:"varint,1,opt,name=IDKota" json:"IDKota,omitempty"`
	NamaKota  string `protobuf:"bytes,2,opt,name=NamaKota" json:"NamaKota,omitempty"`
	Status    int32  `protobuf:"varint,3,opt,name=Status" json:"Status,omitempty"`
	CreatedBy string `protobuf:"bytes,4,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn string `protobuf:"bytes,5,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdatedOn string `protobuf:"bytes,6,opt,name=UpdatedOn" json:"UpdatedOn,omitempty"`
	UpdatedBy string `protobuf:"bytes,7,opt,name=UpdatedBy" json:"UpdatedBy,omitempty"`
}

func (m *UpdateKotaReq) Reset()                    { *m = UpdateKotaReq{} }
func (m *UpdateKotaReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateKotaReq) ProtoMessage()               {}
func (*UpdateKotaReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateKotaReq) GetIDKota() int64 {
	if m != nil {
		return m.IDKota
	}
	return 0
}

func (m *UpdateKotaReq) GetNamaKota() string {
	if m != nil {
		return m.NamaKota
	}
	return ""
}

func (m *UpdateKotaReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateKotaReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *UpdateKotaReq) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *UpdateKotaReq) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *UpdateKotaReq) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func init() {
	proto.RegisterType((*AddKotaReq)(nil), "grpc.AddKotaReq")
	proto.RegisterType((*ReadKotaByNamaReq)(nil), "grpc.ReadKotaByNamaReq")
	proto.RegisterType((*ReadKotaByNamaResp)(nil), "grpc.ReadKotaByNamaResp")
	proto.RegisterType((*ReadKotaResp)(nil), "grpc.ReadKotaResp")
	proto.RegisterType((*UpdateKotaReq)(nil), "grpc.UpdateKotaReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for KotaService service

type KotaServiceClient interface {
	AddKota(ctx context.Context, in *AddKotaReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadKotaByNama(ctx context.Context, in *ReadKotaByNamaReq, opts ...grpc1.CallOption) (*ReadKotaByNamaResp, error)
	ReadKota(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadKotaResp, error)
	UpdateKota(ctx context.Context, in *UpdateKotaReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
}

type kotaServiceClient struct {
	cc *grpc1.ClientConn
}

func NewKotaServiceClient(cc *grpc1.ClientConn) KotaServiceClient {
	return &kotaServiceClient{cc}
}

func (c *kotaServiceClient) AddKota(ctx context.Context, in *AddKotaReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.KotaService/AddKota", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kotaServiceClient) ReadKotaByNama(ctx context.Context, in *ReadKotaByNamaReq, opts ...grpc1.CallOption) (*ReadKotaByNamaResp, error) {
	out := new(ReadKotaByNamaResp)
	err := grpc1.Invoke(ctx, "/grpc.KotaService/ReadKotaByNama", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kotaServiceClient) ReadKota(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadKotaResp, error) {
	out := new(ReadKotaResp)
	err := grpc1.Invoke(ctx, "/grpc.KotaService/ReadKota", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kotaServiceClient) UpdateKota(ctx context.Context, in *UpdateKotaReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.KotaService/UpdateKota", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KotaService service

type KotaServiceServer interface {
	AddKota(context.Context, *AddKotaReq) (*google_protobuf.Empty, error)
	ReadKotaByNama(context.Context, *ReadKotaByNamaReq) (*ReadKotaByNamaResp, error)
	ReadKota(context.Context, *google_protobuf.Empty) (*ReadKotaResp, error)
	UpdateKota(context.Context, *UpdateKotaReq) (*google_protobuf.Empty, error)
}

func RegisterKotaServiceServer(s *grpc1.Server, srv KotaServiceServer) {
	s.RegisterService(&_KotaService_serviceDesc, srv)
}

func _KotaService_AddKota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKotaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KotaServiceServer).AddKota(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KotaService/AddKota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KotaServiceServer).AddKota(ctx, req.(*AddKotaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KotaService_ReadKotaByNama_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadKotaByNamaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KotaServiceServer).ReadKotaByNama(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KotaService/ReadKotaByNama",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KotaServiceServer).ReadKotaByNama(ctx, req.(*ReadKotaByNamaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KotaService_ReadKota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KotaServiceServer).ReadKota(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KotaService/ReadKota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KotaServiceServer).ReadKota(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KotaService_UpdateKota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKotaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KotaServiceServer).UpdateKota(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KotaService/UpdateKota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KotaServiceServer).UpdateKota(ctx, req.(*UpdateKotaReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _KotaService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.KotaService",
	HandlerType: (*KotaServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddKota",
			Handler:    _KotaService_AddKota_Handler,
		},
		{
			MethodName: "ReadKotaByNama",
			Handler:    _KotaService_ReadKotaByNama_Handler,
		},
		{
			MethodName: "ReadKota",
			Handler:    _KotaService_ReadKota_Handler,
		},
		{
			MethodName: "UpdateKota",
			Handler:    _KotaService_UpdateKota_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "kota.proto",
}

func init() { proto.RegisterFile("kota.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x53, 0xcd, 0x4a, 0xc3, 0x40,
	0x18, 0xec, 0xf6, 0xbf, 0x5f, 0x55, 0xf4, 0x13, 0x6a, 0x88, 0x1e, 0x4a, 0x4e, 0x39, 0xa5, 0x50,
	0x11, 0x04, 0x4f, 0x46, 0x7b, 0x10, 0xc1, 0x42, 0x8a, 0x0f, 0xb0, 0x6d, 0xd6, 0x22, 0xb6, 0xdd,
	0x34, 0xdd, 0x0a, 0xb9, 0xfa, 0x84, 0x82, 0x17, 0x1f, 0x47, 0xf6, 0x27, 0x6d, 0x57, 0xcd, 0x0b,
	0xf4, 0xf8, 0xcd, 0xec, 0x24, 0xdf, 0xcc, 0xee, 0x00, 0xbc, 0x71, 0x41, 0x83, 0x24, 0xe5, 0x82,
	0x63, 0x75, 0x9a, 0x26, 0x13, 0xf7, 0x7c, 0xca, 0xf9, 0x74, 0xc6, 0x7a, 0x0a, 0x1b, 0xaf, 0x5f,
	0x7a, 0x6c, 0x9e, 0x88, 0x4c, 0x1f, 0xf1, 0x3e, 0x09, 0xc0, 0x6d, 0x1c, 0x3f, 0x72, 0x41, 0x23,
	0xb6, 0xc4, 0x0e, 0xd4, 0x1f, 0xee, 0xe5, 0xe0, 0x90, 0x2e, 0xf1, 0x2b, 0x91, 0x99, 0xd0, 0x85,
	0xe6, 0x13, 0x9d, 0x53, 0xc5, 0x94, 0xbb, 0xc4, 0x6f, 0x45, 0x9b, 0x59, 0x6a, 0x46, 0x82, 0x8a,
	0xf5, 0xca, 0xa9, 0x74, 0x89, 0x5f, 0x8b, 0xcc, 0x84, 0x17, 0xd0, 0xba, 0x4b, 0x19, 0x15, 0x2c,
	0x0e, 0x33, 0xa7, 0xaa, 0x44, 0x5b, 0x60, 0x87, 0x1d, 0x2e, 0x9c, 0x9a, 0xc5, 0x0e, 0x17, 0x92,
	0x7d, 0x4e, 0x62, 0xc3, 0xd6, 0x35, 0xbb, 0x01, 0x76, 0xd8, 0x30, 0x73, 0x1a, 0x16, 0x1b, 0x66,
	0x5e, 0x0f, 0x4e, 0x22, 0x46, 0x95, 0xa5, 0x30, 0x93, 0x5b, 0x4a, 0x63, 0xbb, 0x06, 0x88, 0x6d,
	0xc0, 0xfb, 0x26, 0x80, 0xbf, 0x15, 0xab, 0x64, 0x2f, 0xb2, 0x08, 0xe1, 0x20, 0x77, 0xa6, 0x3c,
	0xf5, 0xa1, 0x41, 0x67, 0x33, 0x63, 0xaa, 0xe2, 0xb7, 0xfb, 0x4e, 0x20, 0xdf, 0x48, 0xf0, 0xd7,
	0x7e, 0x94, 0x1f, 0xf4, 0xbe, 0x08, 0x1c, 0xea, 0x2f, 0xee, 0xd1, 0x2b, 0xe9, 0x7f, 0x94, 0xa1,
	0x2d, 0x17, 0x1b, 0xb1, 0xf4, 0xfd, 0x75, 0xc2, 0xf0, 0x0a, 0x1a, 0xa6, 0x07, 0x78, 0xac, 0x33,
	0xd9, 0xd6, 0xc2, 0xed, 0x04, 0xba, 0x43, 0x41, 0xde, 0xa1, 0x60, 0x20, 0x3b, 0xe4, 0x95, 0x70,
	0x00, 0x47, 0x76, 0x76, 0x78, 0xf6, 0x7f, 0xa2, 0x4b, 0xb7, 0x30, 0x6a, 0xaf, 0x84, 0xd7, 0xd0,
	0xcc, 0x71, 0x2c, 0xf8, 0x99, 0x8b, 0xb6, 0xde, 0x28, 0x6f, 0x00, 0xb6, 0x97, 0x83, 0xa7, 0xfa,
	0x8c, 0x75, 0x5d, 0xc5, 0xdb, 0x8f, 0xeb, 0x0a, 0xb9, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x70,
	0x10, 0x80, 0x83, 0x35, 0x04, 0x00, 0x00,
}
