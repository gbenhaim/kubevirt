// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package v1alpha2 is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	OnDefineDomainParams
	OnDefineDomainResult
	PreCloudInitIsoParams
	PreCloudInitIsoResult
*/
package v1alpha2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type OnDefineDomainParams struct {
	// domainXML is original libvirt domain specification
	DomainXML []byte `protobuf:"bytes,1,opt,name=domainXML,proto3" json:"domainXML,omitempty"`
	// vmi is VirtualMachineInstance is object of virtual machine currently processed by virt-launcher, it is encoded as JSON
	Vmi []byte `protobuf:"bytes,2,opt,name=vmi,proto3" json:"vmi,omitempty"`
}

func (m *OnDefineDomainParams) Reset()                    { *m = OnDefineDomainParams{} }
func (m *OnDefineDomainParams) String() string            { return proto.CompactTextString(m) }
func (*OnDefineDomainParams) ProtoMessage()               {}
func (*OnDefineDomainParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OnDefineDomainParams) GetDomainXML() []byte {
	if m != nil {
		return m.DomainXML
	}
	return nil
}

func (m *OnDefineDomainParams) GetVmi() []byte {
	if m != nil {
		return m.Vmi
	}
	return nil
}

type OnDefineDomainResult struct {
	// domainXML is processed libvirt domain specification
	DomainXML []byte `protobuf:"bytes,1,opt,name=domainXML,proto3" json:"domainXML,omitempty"`
}

func (m *OnDefineDomainResult) Reset()                    { *m = OnDefineDomainResult{} }
func (m *OnDefineDomainResult) String() string            { return proto.CompactTextString(m) }
func (*OnDefineDomainResult) ProtoMessage()               {}
func (*OnDefineDomainResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OnDefineDomainResult) GetDomainXML() []byte {
	if m != nil {
		return m.DomainXML
	}
	return nil
}

type PreCloudInitIsoParams struct {
	// cloudInitData is an object of CloudInitNoCloudSource encoded as JSON
	CloudInitData []byte `protobuf:"bytes,1,opt,name=cloudInitData,proto3" json:"cloudInitData,omitempty"`
	// vmi is VirtualMachineInstance is object of virtual machine currently processed by virt-launcher, it is encoded as JSON
	Vmi []byte `protobuf:"bytes,2,opt,name=vmi,proto3" json:"vmi,omitempty"`
}

func (m *PreCloudInitIsoParams) Reset()                    { *m = PreCloudInitIsoParams{} }
func (m *PreCloudInitIsoParams) String() string            { return proto.CompactTextString(m) }
func (*PreCloudInitIsoParams) ProtoMessage()               {}
func (*PreCloudInitIsoParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PreCloudInitIsoParams) GetCloudInitData() []byte {
	if m != nil {
		return m.CloudInitData
	}
	return nil
}

func (m *PreCloudInitIsoParams) GetVmi() []byte {
	if m != nil {
		return m.Vmi
	}
	return nil
}

type PreCloudInitIsoResult struct {
	// cloudInitData is an object of CloudInitNoCloudSource encoded as JSON
	CloudInitData []byte `protobuf:"bytes,1,opt,name=cloudInitData,proto3" json:"cloudInitData,omitempty"`
}

func (m *PreCloudInitIsoResult) Reset()                    { *m = PreCloudInitIsoResult{} }
func (m *PreCloudInitIsoResult) String() string            { return proto.CompactTextString(m) }
func (*PreCloudInitIsoResult) ProtoMessage()               {}
func (*PreCloudInitIsoResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PreCloudInitIsoResult) GetCloudInitData() []byte {
	if m != nil {
		return m.CloudInitData
	}
	return nil
}

func init() {
	proto.RegisterType((*OnDefineDomainParams)(nil), "kubevirt.hooks.v1alpha2.OnDefineDomainParams")
	proto.RegisterType((*OnDefineDomainResult)(nil), "kubevirt.hooks.v1alpha2.OnDefineDomainResult")
	proto.RegisterType((*PreCloudInitIsoParams)(nil), "kubevirt.hooks.v1alpha2.PreCloudInitIsoParams")
	proto.RegisterType((*PreCloudInitIsoResult)(nil), "kubevirt.hooks.v1alpha2.PreCloudInitIsoResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Callbacks service

type CallbacksClient interface {
	OnDefineDomain(ctx context.Context, in *OnDefineDomainParams, opts ...grpc.CallOption) (*OnDefineDomainResult, error)
	PreCloudInitIso(ctx context.Context, in *PreCloudInitIsoParams, opts ...grpc.CallOption) (*PreCloudInitIsoResult, error)
}

type callbacksClient struct {
	cc *grpc.ClientConn
}

func NewCallbacksClient(cc *grpc.ClientConn) CallbacksClient {
	return &callbacksClient{cc}
}

func (c *callbacksClient) OnDefineDomain(ctx context.Context, in *OnDefineDomainParams, opts ...grpc.CallOption) (*OnDefineDomainResult, error) {
	out := new(OnDefineDomainResult)
	err := grpc.Invoke(ctx, "/kubevirt.hooks.v1alpha2.Callbacks/OnDefineDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callbacksClient) PreCloudInitIso(ctx context.Context, in *PreCloudInitIsoParams, opts ...grpc.CallOption) (*PreCloudInitIsoResult, error) {
	out := new(PreCloudInitIsoResult)
	err := grpc.Invoke(ctx, "/kubevirt.hooks.v1alpha2.Callbacks/PreCloudInitIso", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Callbacks service

type CallbacksServer interface {
	OnDefineDomain(context.Context, *OnDefineDomainParams) (*OnDefineDomainResult, error)
	PreCloudInitIso(context.Context, *PreCloudInitIsoParams) (*PreCloudInitIsoResult, error)
}

func RegisterCallbacksServer(s *grpc.Server, srv CallbacksServer) {
	s.RegisterService(&_Callbacks_serviceDesc, srv)
}

func _Callbacks_OnDefineDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnDefineDomainParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbacksServer).OnDefineDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.hooks.v1alpha2.Callbacks/OnDefineDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbacksServer).OnDefineDomain(ctx, req.(*OnDefineDomainParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Callbacks_PreCloudInitIso_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreCloudInitIsoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbacksServer).PreCloudInitIso(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.hooks.v1alpha2.Callbacks/PreCloudInitIso",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbacksServer).PreCloudInitIso(ctx, req.(*PreCloudInitIsoParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Callbacks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kubevirt.hooks.v1alpha2.Callbacks",
	HandlerType: (*CallbacksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnDefineDomain",
			Handler:    _Callbacks_OnDefineDomain_Handler,
		},
		{
			MethodName: "PreCloudInitIso",
			Handler:    _Callbacks_PreCloudInitIso_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcf, 0x2e, 0x4d, 0x4a, 0x2d, 0xcb, 0x2c, 0x2a, 0xd1,
	0xcb, 0xc8, 0xcf, 0xcf, 0x2e, 0xd6, 0x2b, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0x52, 0x72,
	0xe3, 0x12, 0xf1, 0xcf, 0x73, 0x49, 0x4d, 0xcb, 0xcc, 0x4b, 0x75, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc,
	0x0b, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0x16, 0x92, 0xe1, 0xe2, 0x4c, 0x01, 0xf3, 0x23, 0x7c, 0x7d,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x10, 0x02, 0x42, 0x02, 0x5c, 0xcc, 0x65, 0xb9, 0x99,
	0x12, 0x4c, 0x60, 0x71, 0x10, 0x53, 0xc9, 0x04, 0xdd, 0x9c, 0xa0, 0xd4, 0xe2, 0xd2, 0x9c, 0x12,
	0xfc, 0xe6, 0x28, 0xf9, 0x73, 0x89, 0x06, 0x14, 0xa5, 0x3a, 0xe7, 0xe4, 0x97, 0xa6, 0x78, 0xe6,
	0x65, 0x96, 0x78, 0x16, 0xe7, 0x43, 0xad, 0x57, 0xe1, 0xe2, 0x4d, 0x86, 0x89, 0xba, 0x24, 0x96,
	0x24, 0x42, 0xb5, 0xa2, 0x0a, 0x62, 0x71, 0x86, 0x2d, 0x86, 0x81, 0x50, 0x77, 0x10, 0x65, 0xa0,
	0xd1, 0x3b, 0x46, 0x2e, 0x4e, 0xe7, 0xc4, 0x9c, 0x9c, 0xa4, 0xc4, 0xe4, 0xec, 0x62, 0xa1, 0x3c,
	0x2e, 0x3e, 0x54, 0x3f, 0x09, 0xe9, 0xea, 0xe1, 0x08, 0x47, 0x3d, 0x6c, 0x81, 0x28, 0x45, 0xac,
	0x72, 0xa8, 0x1b, 0x0b, 0xb9, 0xf8, 0xd1, 0x1c, 0x2f, 0xa4, 0x87, 0xd3, 0x04, 0xac, 0xe1, 0x26,
	0x45, 0xb4, 0x7a, 0x88, 0x95, 0x49, 0x6c, 0xe0, 0xe4, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x76, 0x80, 0xa3, 0xc6, 0x2b, 0x02, 0x00, 0x00,
}
