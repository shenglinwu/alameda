// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/ai_service/ai_service.proto

package ai_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type RecommendationPolicy int32

const (
	RecommendationPolicy_STABLE  RecommendationPolicy = 0
	RecommendationPolicy_COMPACT RecommendationPolicy = 1
)

var RecommendationPolicy_name = map[int32]string{
	0: "STABLE",
	1: "COMPACT",
}
var RecommendationPolicy_value = map[string]int32{
	"STABLE":  0,
	"COMPACT": 1,
}

func (x RecommendationPolicy) String() string {
	return proto.EnumName(RecommendationPolicy_name, int32(x))
}
func (RecommendationPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ai_service_72ef39760db107b3, []int{0}
}

type Object_Type int32

const (
	Object_POD  Object_Type = 0
	Object_NODE Object_Type = 1
)

var Object_Type_name = map[int32]string{
	0: "POD",
	1: "NODE",
}
var Object_Type_value = map[string]int32{
	"POD":  0,
	"NODE": 1,
}

func (x Object_Type) String() string {
	return proto.EnumName(Object_Type_name, int32(x))
}
func (Object_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ai_service_72ef39760db107b3, []int{1, 0}
}

type Pod struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ResourceLink         string   `protobuf:"bytes,2,opt,name=resourceLink,proto3" json:"resourceLink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_ai_service_72ef39760db107b3, []int{0}
}
func (m *Pod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pod.Unmarshal(m, b)
}
func (m *Pod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pod.Marshal(b, m, deterministic)
}
func (dst *Pod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pod.Merge(dst, src)
}
func (m *Pod) XXX_Size() int {
	return xxx_messageInfo_Pod.Size(m)
}
func (m *Pod) XXX_DiscardUnknown() {
	xxx_messageInfo_Pod.DiscardUnknown(m)
}

var xxx_messageInfo_Pod proto.InternalMessageInfo

func (m *Pod) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Pod) GetResourceLink() string {
	if m != nil {
		return m.ResourceLink
	}
	return ""
}

type Object struct {
	Type                 Object_Type          `protobuf:"varint,1,opt,name=type,proto3,enum=Object_Type" json:"type,omitempty"`
	Policy               RecommendationPolicy `protobuf:"varint,2,opt,name=policy,proto3,enum=RecommendationPolicy" json:"policy,omitempty"`
	Uid                  string               `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Namespace            string               `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Pod                  *Pod                 `protobuf:"bytes,6,opt,name=pod,proto3" json:"pod,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_ai_service_72ef39760db107b3, []int{1}
}
func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (dst *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(dst, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetType() Object_Type {
	if m != nil {
		return m.Type
	}
	return Object_POD
}

func (m *Object) GetPolicy() RecommendationPolicy {
	if m != nil {
		return m.Policy
	}
	return RecommendationPolicy_STABLE
}

func (m *Object) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Object) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Object) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Object) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

type PredictionObjectListCreationRequest struct {
	Objects              []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PredictionObjectListCreationRequest) Reset()         { *m = PredictionObjectListCreationRequest{} }
func (m *PredictionObjectListCreationRequest) String() string { return proto.CompactTextString(m) }
func (*PredictionObjectListCreationRequest) ProtoMessage()    {}
func (*PredictionObjectListCreationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ai_service_72ef39760db107b3, []int{2}
}
func (m *PredictionObjectListCreationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionObjectListCreationRequest.Unmarshal(m, b)
}
func (m *PredictionObjectListCreationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionObjectListCreationRequest.Marshal(b, m, deterministic)
}
func (dst *PredictionObjectListCreationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionObjectListCreationRequest.Merge(dst, src)
}
func (m *PredictionObjectListCreationRequest) XXX_Size() int {
	return xxx_messageInfo_PredictionObjectListCreationRequest.Size(m)
}
func (m *PredictionObjectListCreationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionObjectListCreationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionObjectListCreationRequest proto.InternalMessageInfo

func (m *PredictionObjectListCreationRequest) GetObjects() []*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

type PredictionObjectListDeletionRequest struct {
	Objects              []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PredictionObjectListDeletionRequest) Reset()         { *m = PredictionObjectListDeletionRequest{} }
func (m *PredictionObjectListDeletionRequest) String() string { return proto.CompactTextString(m) }
func (*PredictionObjectListDeletionRequest) ProtoMessage()    {}
func (*PredictionObjectListDeletionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ai_service_72ef39760db107b3, []int{3}
}
func (m *PredictionObjectListDeletionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionObjectListDeletionRequest.Unmarshal(m, b)
}
func (m *PredictionObjectListDeletionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionObjectListDeletionRequest.Marshal(b, m, deterministic)
}
func (dst *PredictionObjectListDeletionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionObjectListDeletionRequest.Merge(dst, src)
}
func (m *PredictionObjectListDeletionRequest) XXX_Size() int {
	return xxx_messageInfo_PredictionObjectListDeletionRequest.Size(m)
}
func (m *PredictionObjectListDeletionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionObjectListDeletionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionObjectListDeletionRequest proto.InternalMessageInfo

func (m *PredictionObjectListDeletionRequest) GetObjects() []*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

type RequestResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestResponse) Reset()         { *m = RequestResponse{} }
func (m *RequestResponse) String() string { return proto.CompactTextString(m) }
func (*RequestResponse) ProtoMessage()    {}
func (*RequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ai_service_72ef39760db107b3, []int{4}
}
func (m *RequestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestResponse.Unmarshal(m, b)
}
func (m *RequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestResponse.Marshal(b, m, deterministic)
}
func (dst *RequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestResponse.Merge(dst, src)
}
func (m *RequestResponse) XXX_Size() int {
	return xxx_messageInfo_RequestResponse.Size(m)
}
func (m *RequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestResponse proto.InternalMessageInfo

func (m *RequestResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Pod)(nil), "Pod")
	proto.RegisterType((*Object)(nil), "Object")
	proto.RegisterType((*PredictionObjectListCreationRequest)(nil), "PredictionObjectListCreationRequest")
	proto.RegisterType((*PredictionObjectListDeletionRequest)(nil), "PredictionObjectListDeletionRequest")
	proto.RegisterType((*RequestResponse)(nil), "RequestResponse")
	proto.RegisterEnum("RecommendationPolicy", RecommendationPolicy_name, RecommendationPolicy_value)
	proto.RegisterEnum("Object_Type", Object_Type_name, Object_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlamendaAIServiceClient is the client API for AlamendaAIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlamendaAIServiceClient interface {
	CreatePredictionObjects(ctx context.Context, in *PredictionObjectListCreationRequest, opts ...grpc.CallOption) (*RequestResponse, error)
	DeletePredictionObjects(ctx context.Context, in *PredictionObjectListDeletionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type alamendaAIServiceClient struct {
	cc *grpc.ClientConn
}

func NewAlamendaAIServiceClient(cc *grpc.ClientConn) AlamendaAIServiceClient {
	return &alamendaAIServiceClient{cc}
}

func (c *alamendaAIServiceClient) CreatePredictionObjects(ctx context.Context, in *PredictionObjectListCreationRequest, opts ...grpc.CallOption) (*RequestResponse, error) {
	out := new(RequestResponse)
	err := c.cc.Invoke(ctx, "/AlamendaAIService/CreatePredictionObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alamendaAIServiceClient) DeletePredictionObjects(ctx context.Context, in *PredictionObjectListDeletionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/AlamendaAIService/DeletePredictionObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlamendaAIServiceServer is the server API for AlamendaAIService service.
type AlamendaAIServiceServer interface {
	CreatePredictionObjects(context.Context, *PredictionObjectListCreationRequest) (*RequestResponse, error)
	DeletePredictionObjects(context.Context, *PredictionObjectListDeletionRequest) (*empty.Empty, error)
}

func RegisterAlamendaAIServiceServer(s *grpc.Server, srv AlamendaAIServiceServer) {
	s.RegisterService(&_AlamendaAIService_serviceDesc, srv)
}

func _AlamendaAIService_CreatePredictionObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionObjectListCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlamendaAIServiceServer).CreatePredictionObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AlamendaAIService/CreatePredictionObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlamendaAIServiceServer).CreatePredictionObjects(ctx, req.(*PredictionObjectListCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlamendaAIService_DeletePredictionObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionObjectListDeletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlamendaAIServiceServer).DeletePredictionObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AlamendaAIService/DeletePredictionObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlamendaAIServiceServer).DeletePredictionObjects(ctx, req.(*PredictionObjectListDeletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlamendaAIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AlamendaAIService",
	HandlerType: (*AlamendaAIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePredictionObjects",
			Handler:    _AlamendaAIService_CreatePredictionObjects_Handler,
		},
		{
			MethodName: "DeletePredictionObjects",
			Handler:    _AlamendaAIService_DeletePredictionObjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alameda_api/v1alpha1/ai_service/ai_service.proto",
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/ai_service/ai_service.proto", fileDescriptor_ai_service_72ef39760db107b3)
}

var fileDescriptor_ai_service_72ef39760db107b3 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xd1, 0x8e, 0x93, 0x40,
	0x14, 0x2d, 0x82, 0xb0, 0xbd, 0xdd, 0xac, 0x78, 0xa3, 0xbb, 0xd8, 0xf5, 0xa1, 0xa2, 0x0f, 0x8d,
	0xc6, 0xc1, 0xc5, 0x2f, 0xa8, 0x6d, 0x8d, 0x26, 0x75, 0x4b, 0x68, 0x5f, 0x7c, 0x6a, 0xa6, 0x70,
	0xad, 0x28, 0x30, 0x23, 0x43, 0x37, 0xe9, 0xf7, 0xf9, 0x03, 0x7e, 0x92, 0x61, 0x68, 0xa3, 0x36,
	0x9a, 0xac, 0x6f, 0x77, 0xce, 0xb9, 0x39, 0x9c, 0x73, 0x39, 0xf0, 0x8a, 0xe7, 0xbc, 0xa0, 0x94,
	0xaf, 0xb8, 0xcc, 0x82, 0x9b, 0x2b, 0x9e, 0xcb, 0xcf, 0xfc, 0x2a, 0xe0, 0xd9, 0x4a, 0x51, 0x75,
	0x93, 0x25, 0xf4, 0xdb, 0xc8, 0x64, 0x25, 0x6a, 0xd1, 0xbf, 0xdc, 0x08, 0xb1, 0xc9, 0x29, 0xd0,
	0xaf, 0xf5, 0xf6, 0x53, 0x40, 0x85, 0xac, 0x77, 0x2d, 0xe9, 0xbf, 0x05, 0x33, 0x12, 0x29, 0x5e,
	0x42, 0xb7, 0x14, 0x29, 0xad, 0x4a, 0x5e, 0x90, 0x67, 0x0c, 0x8c, 0x61, 0x37, 0x3e, 0x69, 0x80,
	0x6b, 0x5e, 0x10, 0xfa, 0x70, 0x5a, 0x91, 0x12, 0xdb, 0x2a, 0xa1, 0x59, 0x56, 0x7e, 0xf5, 0xee,
	0x68, 0xfe, 0x0f, 0xcc, 0xff, 0x61, 0x80, 0x3d, 0x5f, 0x7f, 0xa1, 0xa4, 0xc6, 0x01, 0x58, 0xf5,
	0x4e, 0xb6, 0x32, 0x67, 0xe1, 0x29, 0x6b, 0x61, 0xb6, 0xdc, 0x49, 0x8a, 0x35, 0x83, 0x2f, 0xc1,
	0x96, 0x22, 0xcf, 0x92, 0x9d, 0x96, 0x3a, 0x0b, 0x1f, 0xb2, 0x98, 0x12, 0x51, 0x14, 0x54, 0xa6,
	0xbc, 0xce, 0x44, 0x19, 0x69, 0x32, 0xde, 0x2f, 0xa1, 0x0b, 0xe6, 0x36, 0x4b, 0x3d, 0x53, 0x7f,
	0xb6, 0x19, 0xf1, 0x31, 0x74, 0x1b, 0xa7, 0x4a, 0xf2, 0x84, 0x3c, 0x4b, 0xe3, 0xbf, 0x00, 0x44,
	0xb0, 0x74, 0x8e, 0xbb, 0x9a, 0xd0, 0x33, 0x9e, 0x83, 0x29, 0x45, 0xea, 0xd9, 0x03, 0x63, 0xd8,
	0x0b, 0x2d, 0x16, 0x89, 0x34, 0x6e, 0x00, 0xff, 0x11, 0x58, 0x8d, 0x31, 0x74, 0xc0, 0x8c, 0xe6,
	0x13, 0xb7, 0x83, 0x27, 0x60, 0x5d, 0xcf, 0x27, 0x53, 0xd7, 0xf0, 0xdf, 0xc1, 0xd3, 0xa8, 0xa2,
	0x34, 0x4b, 0x1a, 0x4b, 0x6d, 0x88, 0x59, 0xa6, 0xea, 0x71, 0x45, 0xda, 0x64, 0x4c, 0xdf, 0xb6,
	0xa4, 0x6a, 0x7c, 0x02, 0x8e, 0xd0, 0xa4, 0xf2, 0x8c, 0x81, 0x39, 0xec, 0x85, 0xce, 0x3e, 0x71,
	0x7c, 0xc0, 0xff, 0xa5, 0x34, 0xa1, 0x9c, 0xfe, 0x53, 0xe9, 0x05, 0xdc, 0xdb, 0x6f, 0xc7, 0xa4,
	0xa4, 0x28, 0x15, 0xa1, 0x07, 0x4e, 0x41, 0x4a, 0xf1, 0xcd, 0xe1, 0xc7, 0x1d, 0x9e, 0xcf, 0x03,
	0x78, 0xf0, 0xb7, 0xbb, 0x22, 0x80, 0xbd, 0x58, 0x8e, 0xde, 0xcc, 0xa6, 0x6e, 0x07, 0x7b, 0xe0,
	0x8c, 0xe7, 0x1f, 0xa2, 0xd1, 0x78, 0xe9, 0x1a, 0xe1, 0x77, 0x03, 0xee, 0x8f, 0x9a, 0x7a, 0x95,
	0x29, 0x1f, 0xbd, 0x5f, 0xb4, 0x2d, 0xc2, 0x05, 0x5c, 0xe8, 0xcc, 0x74, 0x9c, 0x41, 0xe1, 0x33,
	0x76, 0x8b, 0x0b, 0xf5, 0x5d, 0x76, 0xe4, 0xd9, 0xef, 0xe0, 0x47, 0xb8, 0xd0, 0xf1, 0x6f, 0x2d,
	0x7a, 0x74, 0xac, 0xfe, 0x39, 0x6b, 0x6b, 0xcd, 0x0e, 0xb5, 0x66, 0xd3, 0xa6, 0xd6, 0x7e, 0x67,
	0x6d, 0x6b, 0xe4, 0xf5, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xf2, 0x02, 0x96, 0x2a, 0x03,
	0x00, 0x00,
}