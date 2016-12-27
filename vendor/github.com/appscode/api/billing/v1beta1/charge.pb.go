// Code generated by protoc-gen-go.
// source: charge.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	charge.proto
	paymentmethod.proto
	purchase.proto
	quota.proto
	subscription.proto

It has these top-level messages:
	ChargeEstimateResponse
	PurchaseBeginRequest
	PurchaseBeginResponse
	PurchaseCompleteRequest
	PurchaseCloseRequest
	QuotaGetResponse
	SubscriptionSubscribeRequest
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

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

type ChargeEstimateResponse struct {
	Status  *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	CostUsd string                  `protobuf:"bytes,2,opt,name=cost_usd,json=costUsd" json:"cost_usd,omitempty"`
}

func (m *ChargeEstimateResponse) Reset()                    { *m = ChargeEstimateResponse{} }
func (m *ChargeEstimateResponse) String() string            { return proto.CompactTextString(m) }
func (*ChargeEstimateResponse) ProtoMessage()               {}
func (*ChargeEstimateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ChargeEstimateResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ChargeEstimateResponse) GetCostUsd() string {
	if m != nil {
		return m.CostUsd
	}
	return ""
}

func init() {
	proto.RegisterType((*ChargeEstimateResponse)(nil), "appscode.billing.v1beta1.ChargeEstimateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Charges service

type ChargesClient interface {
	Estimate(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*ChargeEstimateResponse, error)
}

type chargesClient struct {
	cc *grpc.ClientConn
}

func NewChargesClient(cc *grpc.ClientConn) ChargesClient {
	return &chargesClient{cc}
}

func (c *chargesClient) Estimate(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*ChargeEstimateResponse, error) {
	out := new(ChargeEstimateResponse)
	err := grpc.Invoke(ctx, "/appscode.billing.v1beta1.Charges/Estimate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Charges service

type ChargesServer interface {
	Estimate(context.Context, *appscode_dtypes.VoidRequest) (*ChargeEstimateResponse, error)
}

func RegisterChargesServer(s *grpc.Server, srv ChargesServer) {
	s.RegisterService(&_Charges_serviceDesc, srv)
}

func _Charges_Estimate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChargesServer).Estimate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.billing.v1beta1.Charges/Estimate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChargesServer).Estimate(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Charges_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.billing.v1beta1.Charges",
	HandlerType: (*ChargesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Estimate",
			Handler:    _Charges_Estimate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "charge.proto",
}

func init() { proto.RegisterFile("charge.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0xe5, 0x0e, 0x4d, 0x71, 0x99, 0x3c, 0x40, 0x88, 0x2a, 0x11, 0x75, 0x40, 0x99, 0x6c,
	0x5a, 0x36, 0xc6, 0x22, 0xf6, 0x2a, 0x08, 0x06, 0x96, 0xca, 0x89, 0xad, 0x60, 0x29, 0xcd, 0x99,
	0xde, 0x05, 0x89, 0x95, 0x91, 0x15, 0x89, 0x17, 0xe3, 0x15, 0x78, 0x10, 0xd4, 0x38, 0x05, 0xf1,
	0x6f, 0xb1, 0xe4, 0xbb, 0xfb, 0xee, 0x77, 0xdf, 0xc7, 0xf7, 0xcb, 0x3b, 0xbd, 0xa9, 0xac, 0xf4,
	0x1b, 0x20, 0x10, 0xb1, 0xf6, 0x1e, 0x4b, 0x30, 0x56, 0x16, 0xae, 0xae, 0x5d, 0x53, 0xc9, 0x87,
	0x59, 0x61, 0x49, 0xcf, 0x92, 0x49, 0x05, 0x50, 0xd5, 0x56, 0x69, 0xef, 0x94, 0x6e, 0x1a, 0x20,
	0x4d, 0x0e, 0x1a, 0x0c, 0xba, 0xe4, 0x78, 0xa7, 0xeb, 0xfa, 0x86, 0x1e, 0xbd, 0x45, 0xd5, 0xbd,
	0x61, 0x60, 0x6a, 0xf8, 0xc1, 0x45, 0x07, 0xba, 0x44, 0x72, 0x6b, 0x4d, 0x36, 0xb7, 0xe8, 0xa1,
	0x41, 0x2b, 0x14, 0x1f, 0x22, 0x69, 0x6a, 0x31, 0x66, 0x29, 0xcb, 0xc6, 0xf3, 0x43, 0xf9, 0x79,
	0x43, 0xd8, 0x23, 0xaf, 0xba, 0x76, 0xde, 0x8f, 0x89, 0x23, 0x3e, 0x2a, 0x01, 0x69, 0xd5, 0xa2,
	0x89, 0x07, 0x29, 0xcb, 0xf6, 0xf2, 0x68, 0xfb, 0xbf, 0x46, 0x33, 0x7f, 0x65, 0x3c, 0x0a, 0x18,
	0x14, 0xcf, 0x8c, 0x8f, 0x76, 0x30, 0x31, 0xf9, 0xb5, 0xf4, 0x06, 0x9c, 0xc9, 0xed, 0x7d, 0x6b,
	0x91, 0x92, 0x53, 0xf9, 0x9f, 0x6d, 0xf9, 0xf7, 0xd1, 0x53, 0xf9, 0xf4, 0xf6, 0xfe, 0x32, 0xc8,
	0xc4, 0x89, 0x5a, 0x7d, 0x73, 0xde, 0xcb, 0x55, 0x2f, 0x57, 0x21, 0x5c, 0x5c, 0x9c, 0xf3, 0xb4,
	0x84, 0xf5, 0x17, 0x46, 0x7b, 0xf7, 0x13, 0xb5, 0x18, 0x07, 0xd6, 0x72, 0x9b, 0xd7, 0x92, 0xdd,
	0x46, 0x7d, 0xbd, 0x18, 0x76, 0x09, 0x9e, 0x7d, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xaa, 0xd5,
	0xbe, 0xaa, 0x01, 0x00, 0x00,
}