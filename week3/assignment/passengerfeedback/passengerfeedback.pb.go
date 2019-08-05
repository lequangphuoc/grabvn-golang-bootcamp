// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passengerfeedback.proto

package passengerfeedback

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

type PassengerFeedback struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	PassengerID          int32    `protobuf:"varint,2,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	Feedback             string   `protobuf:"bytes,3,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassengerFeedback) Reset()         { *m = PassengerFeedback{} }
func (m *PassengerFeedback) String() string { return proto.CompactTextString(m) }
func (*PassengerFeedback) ProtoMessage()    {}
func (*PassengerFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{0}
}

func (m *PassengerFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassengerFeedback.Unmarshal(m, b)
}
func (m *PassengerFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassengerFeedback.Marshal(b, m, deterministic)
}
func (m *PassengerFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassengerFeedback.Merge(m, src)
}
func (m *PassengerFeedback) XXX_Size() int {
	return xxx_messageInfo_PassengerFeedback.Size(m)
}
func (m *PassengerFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_PassengerFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_PassengerFeedback proto.InternalMessageInfo

func (m *PassengerFeedback) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

func (m *PassengerFeedback) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

func (m *PassengerFeedback) GetFeedback() string {
	if m != nil {
		return m.Feedback
	}
	return ""
}

type AddFeedbackRequest struct {
	Feedback             *PassengerFeedback `protobuf:"bytes,1,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AddFeedbackRequest) Reset()         { *m = AddFeedbackRequest{} }
func (m *AddFeedbackRequest) String() string { return proto.CompactTextString(m) }
func (*AddFeedbackRequest) ProtoMessage()    {}
func (*AddFeedbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{1}
}

func (m *AddFeedbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFeedbackRequest.Unmarshal(m, b)
}
func (m *AddFeedbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFeedbackRequest.Marshal(b, m, deterministic)
}
func (m *AddFeedbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFeedbackRequest.Merge(m, src)
}
func (m *AddFeedbackRequest) XXX_Size() int {
	return xxx_messageInfo_AddFeedbackRequest.Size(m)
}
func (m *AddFeedbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFeedbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddFeedbackRequest proto.InternalMessageInfo

func (m *AddFeedbackRequest) GetFeedback() *PassengerFeedback {
	if m != nil {
		return m.Feedback
	}
	return nil
}

type AddFeedbackResponse struct {
	ResponseMessage      string   `protobuf:"bytes,1,opt,name=responseMessage,proto3" json:"responseMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFeedbackResponse) Reset()         { *m = AddFeedbackResponse{} }
func (m *AddFeedbackResponse) String() string { return proto.CompactTextString(m) }
func (*AddFeedbackResponse) ProtoMessage()    {}
func (*AddFeedbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{2}
}

func (m *AddFeedbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFeedbackResponse.Unmarshal(m, b)
}
func (m *AddFeedbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFeedbackResponse.Marshal(b, m, deterministic)
}
func (m *AddFeedbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFeedbackResponse.Merge(m, src)
}
func (m *AddFeedbackResponse) XXX_Size() int {
	return xxx_messageInfo_AddFeedbackResponse.Size(m)
}
func (m *AddFeedbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFeedbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddFeedbackResponse proto.InternalMessageInfo

func (m *AddFeedbackResponse) GetResponseMessage() string {
	if m != nil {
		return m.ResponseMessage
	}
	return ""
}

type FeedbacksByPassengerIdRequest struct {
	PassengerID          int32    `protobuf:"varint,1,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedbacksByPassengerIdRequest) Reset()         { *m = FeedbacksByPassengerIdRequest{} }
func (m *FeedbacksByPassengerIdRequest) String() string { return proto.CompactTextString(m) }
func (*FeedbacksByPassengerIdRequest) ProtoMessage()    {}
func (*FeedbacksByPassengerIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{3}
}

func (m *FeedbacksByPassengerIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedbacksByPassengerIdRequest.Unmarshal(m, b)
}
func (m *FeedbacksByPassengerIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedbacksByPassengerIdRequest.Marshal(b, m, deterministic)
}
func (m *FeedbacksByPassengerIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedbacksByPassengerIdRequest.Merge(m, src)
}
func (m *FeedbacksByPassengerIdRequest) XXX_Size() int {
	return xxx_messageInfo_FeedbacksByPassengerIdRequest.Size(m)
}
func (m *FeedbacksByPassengerIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedbacksByPassengerIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FeedbacksByPassengerIdRequest proto.InternalMessageInfo

func (m *FeedbacksByPassengerIdRequest) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

type FeedbacksByPassengerIdResponse struct {
	PassengerFeedbacks   []*PassengerFeedback `protobuf:"bytes,1,rep,name=passengerFeedbacks,proto3" json:"passengerFeedbacks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FeedbacksByPassengerIdResponse) Reset()         { *m = FeedbacksByPassengerIdResponse{} }
func (m *FeedbacksByPassengerIdResponse) String() string { return proto.CompactTextString(m) }
func (*FeedbacksByPassengerIdResponse) ProtoMessage()    {}
func (*FeedbacksByPassengerIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{4}
}

func (m *FeedbacksByPassengerIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedbacksByPassengerIdResponse.Unmarshal(m, b)
}
func (m *FeedbacksByPassengerIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedbacksByPassengerIdResponse.Marshal(b, m, deterministic)
}
func (m *FeedbacksByPassengerIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedbacksByPassengerIdResponse.Merge(m, src)
}
func (m *FeedbacksByPassengerIdResponse) XXX_Size() int {
	return xxx_messageInfo_FeedbacksByPassengerIdResponse.Size(m)
}
func (m *FeedbacksByPassengerIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedbacksByPassengerIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FeedbacksByPassengerIdResponse proto.InternalMessageInfo

func (m *FeedbacksByPassengerIdResponse) GetPassengerFeedbacks() []*PassengerFeedback {
	if m != nil {
		return m.PassengerFeedbacks
	}
	return nil
}

type FeedbacksByBookingCodeRequest struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedbacksByBookingCodeRequest) Reset()         { *m = FeedbacksByBookingCodeRequest{} }
func (m *FeedbacksByBookingCodeRequest) String() string { return proto.CompactTextString(m) }
func (*FeedbacksByBookingCodeRequest) ProtoMessage()    {}
func (*FeedbacksByBookingCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{5}
}

func (m *FeedbacksByBookingCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedbacksByBookingCodeRequest.Unmarshal(m, b)
}
func (m *FeedbacksByBookingCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedbacksByBookingCodeRequest.Marshal(b, m, deterministic)
}
func (m *FeedbacksByBookingCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedbacksByBookingCodeRequest.Merge(m, src)
}
func (m *FeedbacksByBookingCodeRequest) XXX_Size() int {
	return xxx_messageInfo_FeedbacksByBookingCodeRequest.Size(m)
}
func (m *FeedbacksByBookingCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedbacksByBookingCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FeedbacksByBookingCodeRequest proto.InternalMessageInfo

func (m *FeedbacksByBookingCodeRequest) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

type FeedbacksByBookingCodeResponse struct {
	PassengerFeedbacks   []*PassengerFeedback `protobuf:"bytes,1,rep,name=passengerFeedbacks,proto3" json:"passengerFeedbacks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FeedbacksByBookingCodeResponse) Reset()         { *m = FeedbacksByBookingCodeResponse{} }
func (m *FeedbacksByBookingCodeResponse) String() string { return proto.CompactTextString(m) }
func (*FeedbacksByBookingCodeResponse) ProtoMessage()    {}
func (*FeedbacksByBookingCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{6}
}

func (m *FeedbacksByBookingCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedbacksByBookingCodeResponse.Unmarshal(m, b)
}
func (m *FeedbacksByBookingCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedbacksByBookingCodeResponse.Marshal(b, m, deterministic)
}
func (m *FeedbacksByBookingCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedbacksByBookingCodeResponse.Merge(m, src)
}
func (m *FeedbacksByBookingCodeResponse) XXX_Size() int {
	return xxx_messageInfo_FeedbacksByBookingCodeResponse.Size(m)
}
func (m *FeedbacksByBookingCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedbacksByBookingCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FeedbacksByBookingCodeResponse proto.InternalMessageInfo

func (m *FeedbacksByBookingCodeResponse) GetPassengerFeedbacks() []*PassengerFeedback {
	if m != nil {
		return m.PassengerFeedbacks
	}
	return nil
}

func init() {
	proto.RegisterType((*PassengerFeedback)(nil), "passengerfeedback.PassengerFeedback")
	proto.RegisterType((*AddFeedbackRequest)(nil), "passengerfeedback.AddFeedbackRequest")
	proto.RegisterType((*AddFeedbackResponse)(nil), "passengerfeedback.AddFeedbackResponse")
	proto.RegisterType((*FeedbacksByPassengerIdRequest)(nil), "passengerfeedback.FeedbacksByPassengerIdRequest")
	proto.RegisterType((*FeedbacksByPassengerIdResponse)(nil), "passengerfeedback.FeedbacksByPassengerIdResponse")
	proto.RegisterType((*FeedbacksByBookingCodeRequest)(nil), "passengerfeedback.FeedbacksByBookingCodeRequest")
	proto.RegisterType((*FeedbacksByBookingCodeResponse)(nil), "passengerfeedback.FeedbacksByBookingCodeResponse")
}

func init() { proto.RegisterFile("passengerfeedback.proto", fileDescriptor_fbe323b0704a1f6c) }

var fileDescriptor_fbe323b0704a1f6c = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x48, 0x2c, 0x2e,
	0x4e, 0xcd, 0x4b, 0x4f, 0x2d, 0x4a, 0x4b, 0x4d, 0x4d, 0x49, 0x4a, 0x4c, 0xce, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc4, 0x90, 0x50, 0x2a, 0xe6, 0x12, 0x0c, 0x80, 0x09, 0xba, 0x41,
	0x05, 0x85, 0x14, 0xb8, 0xb8, 0x93, 0xf2, 0xf3, 0xb3, 0x33, 0xf3, 0xd2, 0x9d, 0xf3, 0x53, 0x52,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x90, 0x85, 0x40, 0x2a, 0xe0, 0x66, 0x79, 0xba, 0x48,
	0x30, 0x29, 0x30, 0x6a, 0xb0, 0x06, 0x21, 0x0b, 0x09, 0x49, 0x71, 0x71, 0xc0, 0x2c, 0x91, 0x60,
	0x06, 0x1b, 0x00, 0xe7, 0x2b, 0x85, 0x71, 0x09, 0x39, 0xa6, 0xa4, 0xc0, 0xac, 0x0b, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x72, 0x40, 0xd2, 0x01, 0xb2, 0x92, 0xdb, 0x48, 0x45, 0x0f, 0xd3,
	0x27, 0x18, 0xae, 0x45, 0x32, 0xd7, 0x9e, 0x4b, 0x18, 0xc5, 0xdc, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0x21, 0x0d, 0x2e, 0xfe, 0x22, 0x28, 0xdb, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0xe6, 0x25,
	0x74, 0x61, 0x25, 0x47, 0x2e, 0x59, 0x98, 0xee, 0x62, 0xa7, 0x4a, 0xb8, 0x55, 0x9e, 0x29, 0x30,
	0x37, 0xa2, 0xf9, 0x9b, 0x11, 0xc3, 0xdf, 0x4a, 0x65, 0x5c, 0x72, 0xb8, 0x8c, 0x80, 0x3a, 0x27,
	0x84, 0x4b, 0xa8, 0x00, 0xdd, 0x13, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x44, 0xfb, 0x18, 0x8b, 0x7e,
	0x34, 0xa7, 0x3b, 0x21, 0xe2, 0x0a, 0xc9, 0xe9, 0xf8, 0x23, 0x15, 0xcd, 0xe9, 0x28, 0x46, 0xd0,
	0xd2, 0xe9, 0x46, 0xff, 0x98, 0xb8, 0x24, 0x30, 0x54, 0x06, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7,
	0x0a, 0xc5, 0x70, 0x71, 0x23, 0xc5, 0xa9, 0x90, 0x2a, 0x16, 0x5b, 0x30, 0xd3, 0x92, 0x94, 0x1a,
	0x21, 0x65, 0x50, 0x0f, 0x35, 0x30, 0x72, 0x49, 0xba, 0xa7, 0x96, 0x60, 0x8f, 0x31, 0x21, 0x03,
	0x2c, 0xa6, 0xe0, 0x4d, 0x1f, 0x52, 0x86, 0x24, 0xe8, 0xc0, 0xe9, 0x04, 0xa4, 0x90, 0x27, 0xe4,
	0x04, 0xcc, 0x78, 0x26, 0xe4, 0x04, 0x2c, 0xd1, 0x9a, 0xc4, 0x06, 0x2e, 0x1e, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x89, 0xff, 0xb3, 0x5e, 0x39, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PassengerFeedbackServiceClient is the client API for PassengerFeedbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PassengerFeedbackServiceClient interface {
	AddFeedback(ctx context.Context, in *AddFeedbackRequest, opts ...grpc.CallOption) (*AddFeedbackResponse, error)
	GetFeedbacksByPassengerId(ctx context.Context, in *FeedbacksByPassengerIdRequest, opts ...grpc.CallOption) (*FeedbacksByPassengerIdResponse, error)
	GetFeedbacksByBookingCode(ctx context.Context, in *FeedbacksByBookingCodeRequest, opts ...grpc.CallOption) (*FeedbacksByBookingCodeResponse, error)
}

type passengerFeedbackServiceClient struct {
	cc *grpc.ClientConn
}

func NewPassengerFeedbackServiceClient(cc *grpc.ClientConn) PassengerFeedbackServiceClient {
	return &passengerFeedbackServiceClient{cc}
}

func (c *passengerFeedbackServiceClient) AddFeedback(ctx context.Context, in *AddFeedbackRequest, opts ...grpc.CallOption) (*AddFeedbackResponse, error) {
	out := new(AddFeedbackResponse)
	err := c.cc.Invoke(ctx, "/passengerfeedback.PassengerFeedbackService/AddFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackServiceClient) GetFeedbacksByPassengerId(ctx context.Context, in *FeedbacksByPassengerIdRequest, opts ...grpc.CallOption) (*FeedbacksByPassengerIdResponse, error) {
	out := new(FeedbacksByPassengerIdResponse)
	err := c.cc.Invoke(ctx, "/passengerfeedback.PassengerFeedbackService/GetFeedbacksByPassengerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackServiceClient) GetFeedbacksByBookingCode(ctx context.Context, in *FeedbacksByBookingCodeRequest, opts ...grpc.CallOption) (*FeedbacksByBookingCodeResponse, error) {
	out := new(FeedbacksByBookingCodeResponse)
	err := c.cc.Invoke(ctx, "/passengerfeedback.PassengerFeedbackService/GetFeedbacksByBookingCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PassengerFeedbackServiceServer is the server API for PassengerFeedbackService service.
type PassengerFeedbackServiceServer interface {
	AddFeedback(context.Context, *AddFeedbackRequest) (*AddFeedbackResponse, error)
	GetFeedbacksByPassengerId(context.Context, *FeedbacksByPassengerIdRequest) (*FeedbacksByPassengerIdResponse, error)
	GetFeedbacksByBookingCode(context.Context, *FeedbacksByBookingCodeRequest) (*FeedbacksByBookingCodeResponse, error)
}

// UnimplementedPassengerFeedbackServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPassengerFeedbackServiceServer struct {
}

func (*UnimplementedPassengerFeedbackServiceServer) AddFeedback(ctx context.Context, req *AddFeedbackRequest) (*AddFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFeedback not implemented")
}
func (*UnimplementedPassengerFeedbackServiceServer) GetFeedbacksByPassengerId(ctx context.Context, req *FeedbacksByPassengerIdRequest) (*FeedbacksByPassengerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedbacksByPassengerId not implemented")
}
func (*UnimplementedPassengerFeedbackServiceServer) GetFeedbacksByBookingCode(ctx context.Context, req *FeedbacksByBookingCodeRequest) (*FeedbacksByBookingCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedbacksByBookingCode not implemented")
}

func RegisterPassengerFeedbackServiceServer(s *grpc.Server, srv PassengerFeedbackServiceServer) {
	s.RegisterService(&_PassengerFeedbackService_serviceDesc, srv)
}

func _PassengerFeedbackService_AddFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).AddFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerfeedback.PassengerFeedbackService/AddFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).AddFeedback(ctx, req.(*AddFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackService_GetFeedbacksByPassengerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedbacksByPassengerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).GetFeedbacksByPassengerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerfeedback.PassengerFeedbackService/GetFeedbacksByPassengerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).GetFeedbacksByPassengerId(ctx, req.(*FeedbacksByPassengerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackService_GetFeedbacksByBookingCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedbacksByBookingCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).GetFeedbacksByBookingCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerfeedback.PassengerFeedbackService/GetFeedbacksByBookingCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).GetFeedbacksByBookingCode(ctx, req.(*FeedbacksByBookingCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PassengerFeedbackService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "passengerfeedback.PassengerFeedbackService",
	HandlerType: (*PassengerFeedbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFeedback",
			Handler:    _PassengerFeedbackService_AddFeedback_Handler,
		},
		{
			MethodName: "GetFeedbacksByPassengerId",
			Handler:    _PassengerFeedbackService_GetFeedbacksByPassengerId_Handler,
		},
		{
			MethodName: "GetFeedbacksByBookingCode",
			Handler:    _PassengerFeedbackService_GetFeedbacksByBookingCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "passengerfeedback.proto",
}