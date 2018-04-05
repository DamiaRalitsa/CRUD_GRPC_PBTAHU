// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tunjangan.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	tunjangan.proto

It has these top-level messages:
	AddTunjanganReq
	ReadTunjanganByNamaReq
	ReadTunjanganByNamaResp
	ReadTunjanganResp
	UpdateTunjanganReq
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

type AddTunjanganReq struct {
	IDTunjangan        int64  `protobuf:"varint,1,opt,name=IDTunjangan" json:"IDTunjangan,omitempty"`
	DeskripsiTunjangan string `protobuf:"bytes,2,opt,name=DeskripsiTunjangan" json:"DeskripsiTunjangan,omitempty"`
	BesaranTunjangan   int64  `protobuf:"varint,3,opt,name=BesaranTunjangan" json:"BesaranTunjangan,omitempty"`
	Status             int32  `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	CreatedBy          string `protobuf:"bytes,5,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn          string `protobuf:"bytes,6,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdatedOn          string `protobuf:"bytes,7,opt,name=UpdatedOn" json:"UpdatedOn,omitempty"`
	UpdatedBy          string `protobuf:"bytes,8,opt,name=UpdatedBy" json:"UpdatedBy,omitempty"`
}

func (m *AddTunjanganReq) Reset()                    { *m = AddTunjanganReq{} }
func (m *AddTunjanganReq) String() string            { return proto.CompactTextString(m) }
func (*AddTunjanganReq) ProtoMessage()               {}
func (*AddTunjanganReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddTunjanganReq) GetIDTunjangan() int64 {
	if m != nil {
		return m.IDTunjangan
	}
	return 0
}

func (m *AddTunjanganReq) GetDeskripsiTunjangan() string {
	if m != nil {
		return m.DeskripsiTunjangan
	}
	return ""
}

func (m *AddTunjanganReq) GetBesaranTunjangan() int64 {
	if m != nil {
		return m.BesaranTunjangan
	}
	return 0
}

func (m *AddTunjanganReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddTunjanganReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *AddTunjanganReq) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *AddTunjanganReq) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *AddTunjanganReq) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

type ReadTunjanganByNamaReq struct {
	DeskripsiTunjangan string `protobuf:"bytes,1,opt,name=DeskripsiTunjangan" json:"DeskripsiTunjangan,omitempty"`
}

func (m *ReadTunjanganByNamaReq) Reset()                    { *m = ReadTunjanganByNamaReq{} }
func (m *ReadTunjanganByNamaReq) String() string            { return proto.CompactTextString(m) }
func (*ReadTunjanganByNamaReq) ProtoMessage()               {}
func (*ReadTunjanganByNamaReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadTunjanganByNamaReq) GetDeskripsiTunjangan() string {
	if m != nil {
		return m.DeskripsiTunjangan
	}
	return ""
}

type ReadTunjanganByNamaResp struct {
	IDTunjangan        int64  `protobuf:"varint,1,opt,name=IDTunjangan" json:"IDTunjangan,omitempty"`
	DeskripsiTunjangan string `protobuf:"bytes,2,opt,name=DeskripsiTunjangan" json:"DeskripsiTunjangan,omitempty"`
	BesaranTunjangan   int64  `protobuf:"varint,3,opt,name=BesaranTunjangan" json:"BesaranTunjangan,omitempty"`
	Status             int32  `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	CreatedBy          string `protobuf:"bytes,5,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn          string `protobuf:"bytes,6,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdatedOn          string `protobuf:"bytes,7,opt,name=UpdatedOn" json:"UpdatedOn,omitempty"`
	UpdatedBy          string `protobuf:"bytes,8,opt,name=UpdatedBy" json:"UpdatedBy,omitempty"`
}

func (m *ReadTunjanganByNamaResp) Reset()                    { *m = ReadTunjanganByNamaResp{} }
func (m *ReadTunjanganByNamaResp) String() string            { return proto.CompactTextString(m) }
func (*ReadTunjanganByNamaResp) ProtoMessage()               {}
func (*ReadTunjanganByNamaResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadTunjanganByNamaResp) GetIDTunjangan() int64 {
	if m != nil {
		return m.IDTunjangan
	}
	return 0
}

func (m *ReadTunjanganByNamaResp) GetDeskripsiTunjangan() string {
	if m != nil {
		return m.DeskripsiTunjangan
	}
	return ""
}

func (m *ReadTunjanganByNamaResp) GetBesaranTunjangan() int64 {
	if m != nil {
		return m.BesaranTunjangan
	}
	return 0
}

func (m *ReadTunjanganByNamaResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReadTunjanganByNamaResp) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *ReadTunjanganByNamaResp) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *ReadTunjanganByNamaResp) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *ReadTunjanganByNamaResp) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

type ReadTunjanganResp struct {
	AllTunjangan []*ReadTunjanganByNamaResp `protobuf:"bytes,1,rep,name=allTunjangan" json:"allTunjangan,omitempty"`
}

func (m *ReadTunjanganResp) Reset()                    { *m = ReadTunjanganResp{} }
func (m *ReadTunjanganResp) String() string            { return proto.CompactTextString(m) }
func (*ReadTunjanganResp) ProtoMessage()               {}
func (*ReadTunjanganResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadTunjanganResp) GetAllTunjangan() []*ReadTunjanganByNamaResp {
	if m != nil {
		return m.AllTunjangan
	}
	return nil
}

type UpdateTunjanganReq struct {
	IDTunjangan        int64  `protobuf:"varint,1,opt,name=IDTunjangan" json:"IDTunjangan,omitempty"`
	DeskripsiTunjangan string `protobuf:"bytes,2,opt,name=DeskripsiTunjangan" json:"DeskripsiTunjangan,omitempty"`
	BesaranTunjangan   int64  `protobuf:"varint,3,opt,name=BesaranTunjangan" json:"BesaranTunjangan,omitempty"`
	Status             int32  `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	CreatedBy          string `protobuf:"bytes,5,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn          string `protobuf:"bytes,6,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdatedOn          string `protobuf:"bytes,7,opt,name=UpdatedOn" json:"UpdatedOn,omitempty"`
	UpdatedBy          string `protobuf:"bytes,8,opt,name=UpdatedBy" json:"UpdatedBy,omitempty"`
}

func (m *UpdateTunjanganReq) Reset()                    { *m = UpdateTunjanganReq{} }
func (m *UpdateTunjanganReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateTunjanganReq) ProtoMessage()               {}
func (*UpdateTunjanganReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateTunjanganReq) GetIDTunjangan() int64 {
	if m != nil {
		return m.IDTunjangan
	}
	return 0
}

func (m *UpdateTunjanganReq) GetDeskripsiTunjangan() string {
	if m != nil {
		return m.DeskripsiTunjangan
	}
	return ""
}

func (m *UpdateTunjanganReq) GetBesaranTunjangan() int64 {
	if m != nil {
		return m.BesaranTunjangan
	}
	return 0
}

func (m *UpdateTunjanganReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateTunjanganReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *UpdateTunjanganReq) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *UpdateTunjanganReq) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *UpdateTunjanganReq) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func init() {
	proto.RegisterType((*AddTunjanganReq)(nil), "grpc.AddTunjanganReq")
	proto.RegisterType((*ReadTunjanganByNamaReq)(nil), "grpc.ReadTunjanganByNamaReq")
	proto.RegisterType((*ReadTunjanganByNamaResp)(nil), "grpc.ReadTunjanganByNamaResp")
	proto.RegisterType((*ReadTunjanganResp)(nil), "grpc.ReadTunjanganResp")
	proto.RegisterType((*UpdateTunjanganReq)(nil), "grpc.UpdateTunjanganReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for TunjanganService service

type TunjanganServiceClient interface {
	AddTunjangan(ctx context.Context, in *AddTunjanganReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadTunjanganByNama(ctx context.Context, in *ReadTunjanganByNamaReq, opts ...grpc1.CallOption) (*ReadTunjanganByNamaResp, error)
	ReadTunjangan(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadTunjanganResp, error)
	UpdateTunjangan(ctx context.Context, in *UpdateTunjanganReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
}

type tunjanganServiceClient struct {
	cc *grpc1.ClientConn
}

func NewTunjanganServiceClient(cc *grpc1.ClientConn) TunjanganServiceClient {
	return &tunjanganServiceClient{cc}
}

func (c *tunjanganServiceClient) AddTunjangan(ctx context.Context, in *AddTunjanganReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.TunjanganService/AddTunjangan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunjanganServiceClient) ReadTunjanganByNama(ctx context.Context, in *ReadTunjanganByNamaReq, opts ...grpc1.CallOption) (*ReadTunjanganByNamaResp, error) {
	out := new(ReadTunjanganByNamaResp)
	err := grpc1.Invoke(ctx, "/grpc.TunjanganService/ReadTunjanganByNama", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunjanganServiceClient) ReadTunjangan(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadTunjanganResp, error) {
	out := new(ReadTunjanganResp)
	err := grpc1.Invoke(ctx, "/grpc.TunjanganService/ReadTunjangan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunjanganServiceClient) UpdateTunjangan(ctx context.Context, in *UpdateTunjanganReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.TunjanganService/UpdateTunjangan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TunjanganService service

type TunjanganServiceServer interface {
	AddTunjangan(context.Context, *AddTunjanganReq) (*google_protobuf.Empty, error)
	ReadTunjanganByNama(context.Context, *ReadTunjanganByNamaReq) (*ReadTunjanganByNamaResp, error)
	ReadTunjangan(context.Context, *google_protobuf.Empty) (*ReadTunjanganResp, error)
	UpdateTunjangan(context.Context, *UpdateTunjanganReq) (*google_protobuf.Empty, error)
}

func RegisterTunjanganServiceServer(s *grpc1.Server, srv TunjanganServiceServer) {
	s.RegisterService(&_TunjanganService_serviceDesc, srv)
}

func _TunjanganService_AddTunjangan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTunjanganReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunjanganServiceServer).AddTunjangan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TunjanganService/AddTunjangan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunjanganServiceServer).AddTunjangan(ctx, req.(*AddTunjanganReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunjanganService_ReadTunjanganByNama_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadTunjanganByNamaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunjanganServiceServer).ReadTunjanganByNama(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TunjanganService/ReadTunjanganByNama",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunjanganServiceServer).ReadTunjanganByNama(ctx, req.(*ReadTunjanganByNamaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunjanganService_ReadTunjangan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunjanganServiceServer).ReadTunjangan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TunjanganService/ReadTunjangan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunjanganServiceServer).ReadTunjangan(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunjanganService_UpdateTunjangan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTunjanganReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunjanganServiceServer).UpdateTunjangan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TunjanganService/UpdateTunjangan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunjanganServiceServer).UpdateTunjangan(ctx, req.(*UpdateTunjanganReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TunjanganService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.TunjanganService",
	HandlerType: (*TunjanganServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddTunjangan",
			Handler:    _TunjanganService_AddTunjangan_Handler,
		},
		{
			MethodName: "ReadTunjanganByNama",
			Handler:    _TunjanganService_ReadTunjanganByNama_Handler,
		},
		{
			MethodName: "ReadTunjangan",
			Handler:    _TunjanganService_ReadTunjangan_Handler,
		},
		{
			MethodName: "UpdateTunjangan",
			Handler:    _TunjanganService_UpdateTunjangan_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "tunjangan.proto",
}

func init() { proto.RegisterFile("tunjangan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0xcf, 0x6e, 0x82, 0x40,
	0x10, 0xc6, 0x05, 0xff, 0xb4, 0x8e, 0x36, 0xda, 0x69, 0xaa, 0x84, 0xda, 0x84, 0x70, 0x22, 0x3d,
	0xac, 0x89, 0x7d, 0x80, 0x46, 0xaa, 0x49, 0x7b, 0xa9, 0x09, 0xda, 0xde, 0x57, 0xd9, 0x12, 0x5b,
	0x05, 0x04, 0x6c, 0xc2, 0xd3, 0x34, 0xe9, 0xa5, 0x2f, 0xd0, 0x07, 0x6c, 0x00, 0x15, 0x50, 0xf0,
	0x09, 0x3c, 0xf2, 0x7d, 0xb3, 0xb3, 0xcb, 0x6f, 0x26, 0x1f, 0x34, 0xbc, 0xb5, 0xf9, 0x41, 0x4d,
	0x83, 0x9a, 0xc4, 0x76, 0x2c, 0xcf, 0xc2, 0x92, 0xe1, 0xd8, 0x33, 0xf1, 0xc6, 0xb0, 0x2c, 0x63,
	0xc1, 0xba, 0xa1, 0x36, 0x5d, 0xbf, 0x77, 0xd9, 0xd2, 0xf6, 0xfc, 0xa8, 0x44, 0xfe, 0xe6, 0xa1,
	0xd1, 0xd7, 0xf5, 0xc9, 0xf6, 0xa4, 0xc6, 0x56, 0x28, 0x41, 0xed, 0x79, 0xb0, 0x53, 0x04, 0x4e,
	0xe2, 0x94, 0xa2, 0x96, 0x94, 0x90, 0x00, 0x0e, 0x98, 0xfb, 0xe9, 0xcc, 0x6d, 0x77, 0x1e, 0x17,
	0xf2, 0x12, 0xa7, 0x54, 0xb5, 0x0c, 0x07, 0xef, 0xa0, 0xa9, 0x32, 0x97, 0x3a, 0xd4, 0x8c, 0xab,
	0x8b, 0x61, 0xdb, 0x03, 0x1d, 0x5b, 0x50, 0x19, 0x7b, 0xd4, 0x5b, 0xbb, 0x42, 0x49, 0xe2, 0x94,
	0xb2, 0xb6, 0xf9, 0xc2, 0x0e, 0x54, 0x1f, 0x1d, 0x46, 0x3d, 0xa6, 0xab, 0xbe, 0x50, 0x0e, 0xaf,
	0x8a, 0x85, 0x84, 0x3b, 0x32, 0x85, 0x4a, 0xca, 0x1d, 0x99, 0x81, 0xfb, 0x6a, 0xeb, 0x1b, 0xf7,
	0x2c, 0x72, 0x77, 0x42, 0xc2, 0x55, 0x7d, 0xe1, 0x3c, 0xe5, 0xaa, 0xbe, 0xfc, 0x04, 0x2d, 0x8d,
	0xd1, 0x98, 0x90, 0xea, 0xbf, 0xd0, 0x25, 0x0d, 0x38, 0x65, 0x53, 0xe0, 0xf2, 0x28, 0xc8, 0xbf,
	0x3c, 0xb4, 0x33, 0x5b, 0xb9, 0xf6, 0x89, 0x79, 0x8a, 0xf9, 0x1b, 0x5c, 0xa6, 0x40, 0x85, 0x88,
	0xfa, 0x50, 0xa7, 0x8b, 0x45, 0x92, 0x51, 0x51, 0xa9, 0xf5, 0x6e, 0x49, 0xb0, 0xe4, 0x24, 0x87,
	0xab, 0x96, 0x3a, 0x22, 0xff, 0xf0, 0x80, 0xd1, 0x2d, 0xa7, 0x85, 0xcf, 0x83, 0xdf, 0xfb, 0xe3,
	0xa1, 0xb9, 0x7b, 0xdd, 0x98, 0x39, 0x5f, 0xf3, 0x19, 0xc3, 0x07, 0xa8, 0x27, 0x63, 0x02, 0xaf,
	0x23, 0xec, 0x7b, 0xd1, 0x21, 0xb6, 0x48, 0x14, 0x36, 0x64, 0x1b, 0x36, 0x64, 0x18, 0x84, 0x8d,
	0x5c, 0xc0, 0x09, 0x5c, 0x65, 0xcc, 0x08, 0x3b, 0x47, 0xc6, 0xb7, 0x12, 0x8f, 0x0f, 0x57, 0x2e,
	0xa0, 0x0a, 0x17, 0x29, 0x13, 0x73, 0x1e, 0x20, 0xb6, 0x33, 0x3a, 0x6d, 0x7a, 0x0c, 0xa1, 0xb1,
	0xb7, 0x13, 0x28, 0x44, 0xd5, 0x87, 0xab, 0x92, 0xff, 0x83, 0xd3, 0x4a, 0xa8, 0xdc, 0xff, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x9a, 0x4f, 0x59, 0x77, 0x86, 0x05, 0x00, 0x00,
}
