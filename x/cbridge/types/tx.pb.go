// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sgn/cbridge/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// this line is used by starport scaffolding # proto/tx/message
type MsgUpdateAsset struct {
	Creator string                `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Assets  *MultiChainAssetParam `protobuf:"bytes,2,opt,name=assets,proto3" json:"assets,omitempty"`
}

func (m *MsgUpdateAsset) Reset()         { *m = MsgUpdateAsset{} }
func (m *MsgUpdateAsset) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateAsset) ProtoMessage()    {}
func (*MsgUpdateAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8bb6ecae18aabd7, []int{0}
}
func (m *MsgUpdateAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateAsset.Merge(m, src)
}
func (m *MsgUpdateAsset) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateAsset.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateAsset proto.InternalMessageInfo

func (m *MsgUpdateAsset) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateAsset) GetAssets() *MultiChainAssetParam {
	if m != nil {
		return m.Assets
	}
	return nil
}

type MsgUpdateAssetResponse struct {
}

func (m *MsgUpdateAssetResponse) Reset()         { *m = MsgUpdateAssetResponse{} }
func (m *MsgUpdateAssetResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateAssetResponse) ProtoMessage()    {}
func (*MsgUpdateAssetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8bb6ecae18aabd7, []int{1}
}
func (m *MsgUpdateAssetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateAssetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateAssetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateAssetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateAssetResponse.Merge(m, src)
}
func (m *MsgUpdateAssetResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateAssetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateAssetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateAssetResponse proto.InternalMessageInfo

type MsgOnchainEvent struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Chainid uint64 `protobuf:"varint,2,opt,name=chainid,proto3" json:"chainid,omitempty"`
	Evtype  string `protobuf:"bytes,3,opt,name=evtype,proto3" json:"evtype,omitempty"`
	Elog    []byte `protobuf:"bytes,4,opt,name=elog,proto3" json:"elog,omitempty"`
}

func (m *MsgOnchainEvent) Reset()         { *m = MsgOnchainEvent{} }
func (m *MsgOnchainEvent) String() string { return proto.CompactTextString(m) }
func (*MsgOnchainEvent) ProtoMessage()    {}
func (*MsgOnchainEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8bb6ecae18aabd7, []int{2}
}
func (m *MsgOnchainEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOnchainEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOnchainEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOnchainEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOnchainEvent.Merge(m, src)
}
func (m *MsgOnchainEvent) XXX_Size() int {
	return m.Size()
}
func (m *MsgOnchainEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOnchainEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOnchainEvent proto.InternalMessageInfo

func (m *MsgOnchainEvent) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgOnchainEvent) GetChainid() uint64 {
	if m != nil {
		return m.Chainid
	}
	return 0
}

func (m *MsgOnchainEvent) GetEvtype() string {
	if m != nil {
		return m.Evtype
	}
	return ""
}

func (m *MsgOnchainEvent) GetElog() []byte {
	if m != nil {
		return m.Elog
	}
	return nil
}

type MsgOnchainEventResponse struct {
}

func (m *MsgOnchainEventResponse) Reset()         { *m = MsgOnchainEventResponse{} }
func (m *MsgOnchainEventResponse) String() string { return proto.CompactTextString(m) }
func (*MsgOnchainEventResponse) ProtoMessage()    {}
func (*MsgOnchainEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8bb6ecae18aabd7, []int{3}
}
func (m *MsgOnchainEventResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOnchainEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOnchainEventResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOnchainEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOnchainEventResponse.Merge(m, src)
}
func (m *MsgOnchainEventResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgOnchainEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOnchainEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOnchainEventResponse proto.InternalMessageInfo

type MsgOnchainManyEvents struct {
	Creator string             `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Events  []*MsgOnchainEvent `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *MsgOnchainManyEvents) Reset()         { *m = MsgOnchainManyEvents{} }
func (m *MsgOnchainManyEvents) String() string { return proto.CompactTextString(m) }
func (*MsgOnchainManyEvents) ProtoMessage()    {}
func (*MsgOnchainManyEvents) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8bb6ecae18aabd7, []int{4}
}
func (m *MsgOnchainManyEvents) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOnchainManyEvents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOnchainManyEvents.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOnchainManyEvents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOnchainManyEvents.Merge(m, src)
}
func (m *MsgOnchainManyEvents) XXX_Size() int {
	return m.Size()
}
func (m *MsgOnchainManyEvents) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOnchainManyEvents.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOnchainManyEvents proto.InternalMessageInfo

func (m *MsgOnchainManyEvents) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgOnchainManyEvents) GetEvents() []*MsgOnchainEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type MsgOnchainManyEventsResponse struct {
}

func (m *MsgOnchainManyEventsResponse) Reset()         { *m = MsgOnchainManyEventsResponse{} }
func (m *MsgOnchainManyEventsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgOnchainManyEventsResponse) ProtoMessage()    {}
func (*MsgOnchainManyEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8bb6ecae18aabd7, []int{5}
}
func (m *MsgOnchainManyEventsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOnchainManyEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOnchainManyEventsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOnchainManyEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOnchainManyEventsResponse.Merge(m, src)
}
func (m *MsgOnchainManyEventsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgOnchainManyEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOnchainManyEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOnchainManyEventsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateAsset)(nil), "sgn.cbridge.v1.MsgUpdateAsset")
	proto.RegisterType((*MsgUpdateAssetResponse)(nil), "sgn.cbridge.v1.MsgUpdateAssetResponse")
	proto.RegisterType((*MsgOnchainEvent)(nil), "sgn.cbridge.v1.MsgOnchainEvent")
	proto.RegisterType((*MsgOnchainEventResponse)(nil), "sgn.cbridge.v1.MsgOnchainEventResponse")
	proto.RegisterType((*MsgOnchainManyEvents)(nil), "sgn.cbridge.v1.MsgOnchainManyEvents")
	proto.RegisterType((*MsgOnchainManyEventsResponse)(nil), "sgn.cbridge.v1.MsgOnchainManyEventsResponse")
}

func init() { proto.RegisterFile("sgn/cbridge/v1/tx.proto", fileDescriptor_f8bb6ecae18aabd7) }

var fileDescriptor_f8bb6ecae18aabd7 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0x31, 0x20, 0xaa, 0x0e, 0x88, 0xaa, 0xab, 0x0a, 0x5c, 0x0b, 0x6d, 0x2d, 0x0b, 0xb5,
	0x3e, 0x14, 0x5b, 0xb8, 0x87, 0x5e, 0x7a, 0x69, 0xa2, 0x1c, 0x72, 0xb0, 0x12, 0x59, 0x42, 0x8a,
	0x72, 0x33, 0x66, 0xb5, 0x58, 0x81, 0xb5, 0xe3, 0x5d, 0x1c, 0x78, 0x8b, 0xbc, 0x41, 0x5e, 0x27,
	0x47, 0x8e, 0x39, 0x46, 0xf0, 0x22, 0x91, 0x2d, 0x9b, 0xff, 0x81, 0xdc, 0x3c, 0xfa, 0x7e, 0x33,
	0xdf, 0xe7, 0xb1, 0x07, 0x9a, 0x9c, 0x32, 0xd3, 0xeb, 0x47, 0xfe, 0x80, 0x12, 0x33, 0xee, 0x9a,
	0x62, 0x6a, 0x84, 0x51, 0x20, 0x02, 0x54, 0xe7, 0x94, 0x19, 0x99, 0x60, 0xc4, 0x5d, 0xa5, 0xb5,
	0x03, 0xe6, 0x52, 0x4a, 0x6b, 0x43, 0xa8, 0xdb, 0x9c, 0xf6, 0xc2, 0x81, 0x2b, 0xc8, 0x7f, 0xce,
	0x89, 0x40, 0x32, 0x7c, 0xf2, 0x22, 0xe2, 0x8a, 0x20, 0x92, 0x25, 0x55, 0xd2, 0x3f, 0x3b, 0x79,
	0x89, 0xfe, 0x41, 0xc5, 0x4d, 0x10, 0x2e, 0x17, 0x55, 0x49, 0xaf, 0x5a, 0x6d, 0x63, 0xdb, 0xca,
	0xb0, 0x27, 0x23, 0xe1, 0x9f, 0x0f, 0x5d, 0x9f, 0xa5, 0xa3, 0xae, 0xdd, 0xc8, 0x1d, 0x3b, 0x59,
	0x8f, 0x26, 0x43, 0x63, 0xdb, 0xc9, 0x21, 0x3c, 0x0c, 0x18, 0x27, 0xda, 0x3d, 0x7c, 0xb1, 0x39,
	0xbd, 0x62, 0x5e, 0xd2, 0x79, 0x11, 0x13, 0x76, 0x2c, 0x44, 0xa2, 0x24, 0x9c, 0x3f, 0x48, 0x53,
	0x94, 0x9d, 0xbc, 0x44, 0x0d, 0xa8, 0x90, 0x58, 0xcc, 0x42, 0x22, 0x97, 0xd2, 0x96, 0xac, 0x42,
	0x08, 0xca, 0x64, 0x14, 0x50, 0xb9, 0xac, 0x4a, 0x7a, 0xcd, 0x49, 0x9f, 0xb5, 0xef, 0xd0, 0xdc,
	0xb1, 0x5c, 0xa5, 0xf1, 0xe1, 0xdb, 0x5a, 0xb2, 0x5d, 0x36, 0x4b, 0x65, 0x7e, 0x24, 0xd2, 0xdf,
	0xc4, 0x38, 0x61, 0xe4, 0xa2, 0x5a, 0xd2, 0xab, 0xd6, 0x8f, 0xbd, 0xbd, 0xec, 0x58, 0x65, 0xb8,
	0x86, 0xa1, 0x75, 0xc8, 0x2a, 0x8f, 0x62, 0x3d, 0x15, 0xa1, 0x64, 0x73, 0x8a, 0x7a, 0x50, 0xdd,
	0xfc, 0x42, 0xf8, 0xc0, 0xfc, 0x0d, 0x5d, 0xf9, 0x79, 0x5c, 0xcf, 0xc7, 0xa3, 0x1b, 0xa8, 0x6d,
	0x2d, 0xfd, 0x54, 0x6e, 0xe5, 0xd7, 0xa9, 0x17, 0xcb, 0x27, 0x53, 0xf8, 0xba, 0xbf, 0xc0, 0xf6,
	0xfb, 0xdd, 0x6b, 0x4a, 0xf9, 0xfd, 0x11, 0x2a, 0x37, 0x3a, 0xbb, 0x7c, 0x5e, 0x60, 0x69, 0xbe,
	0xc0, 0xd2, 0xeb, 0x02, 0x4b, 0x8f, 0x4b, 0x5c, 0x98, 0x2f, 0x71, 0xe1, 0x65, 0x89, 0x0b, 0xb7,
	0x26, 0xf5, 0xc5, 0x70, 0xd2, 0x37, 0xbc, 0x60, 0x6c, 0x7a, 0x64, 0x44, 0xa2, 0x0e, 0x23, 0xe2,
	0x21, 0x88, 0xee, 0x4c, 0x4e, 0x59, 0x27, 0xb6, 0xcc, 0xe9, 0xea, 0x28, 0x92, 0xbf, 0x84, 0xf7,
	0x2b, 0xe9, 0x41, 0xfc, 0x79, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xdc, 0xab, 0xe8, 0x59, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	UpdateAsset(ctx context.Context, in *MsgUpdateAsset, opts ...grpc.CallOption) (*MsgUpdateAssetResponse, error)
	OnchainEvent(ctx context.Context, in *MsgOnchainEvent, opts ...grpc.CallOption) (*MsgOnchainEventResponse, error)
	OnchainManyEvents(ctx context.Context, in *MsgOnchainManyEvents, opts ...grpc.CallOption) (*MsgOnchainManyEventsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateAsset(ctx context.Context, in *MsgUpdateAsset, opts ...grpc.CallOption) (*MsgUpdateAssetResponse, error) {
	out := new(MsgUpdateAssetResponse)
	err := c.cc.Invoke(ctx, "/sgn.cbridge.v1.Msg/UpdateAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) OnchainEvent(ctx context.Context, in *MsgOnchainEvent, opts ...grpc.CallOption) (*MsgOnchainEventResponse, error) {
	out := new(MsgOnchainEventResponse)
	err := c.cc.Invoke(ctx, "/sgn.cbridge.v1.Msg/OnchainEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) OnchainManyEvents(ctx context.Context, in *MsgOnchainManyEvents, opts ...grpc.CallOption) (*MsgOnchainManyEventsResponse, error) {
	out := new(MsgOnchainManyEventsResponse)
	err := c.cc.Invoke(ctx, "/sgn.cbridge.v1.Msg/OnchainManyEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	UpdateAsset(context.Context, *MsgUpdateAsset) (*MsgUpdateAssetResponse, error)
	OnchainEvent(context.Context, *MsgOnchainEvent) (*MsgOnchainEventResponse, error)
	OnchainManyEvents(context.Context, *MsgOnchainManyEvents) (*MsgOnchainManyEventsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateAsset(ctx context.Context, req *MsgUpdateAsset) (*MsgUpdateAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAsset not implemented")
}
func (*UnimplementedMsgServer) OnchainEvent(ctx context.Context, req *MsgOnchainEvent) (*MsgOnchainEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnchainEvent not implemented")
}
func (*UnimplementedMsgServer) OnchainManyEvents(ctx context.Context, req *MsgOnchainManyEvents) (*MsgOnchainManyEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnchainManyEvents not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateAsset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgn.cbridge.v1.Msg/UpdateAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateAsset(ctx, req.(*MsgUpdateAsset))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_OnchainEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgOnchainEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).OnchainEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgn.cbridge.v1.Msg/OnchainEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).OnchainEvent(ctx, req.(*MsgOnchainEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_OnchainManyEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgOnchainManyEvents)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).OnchainManyEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgn.cbridge.v1.Msg/OnchainManyEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).OnchainManyEvents(ctx, req.(*MsgOnchainManyEvents))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sgn.cbridge.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateAsset",
			Handler:    _Msg_UpdateAsset_Handler,
		},
		{
			MethodName: "OnchainEvent",
			Handler:    _Msg_OnchainEvent_Handler,
		},
		{
			MethodName: "OnchainManyEvents",
			Handler:    _Msg_OnchainManyEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sgn/cbridge/v1/tx.proto",
}

func (m *MsgUpdateAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Assets != nil {
		{
			size, err := m.Assets.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateAssetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateAssetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateAssetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgOnchainEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOnchainEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOnchainEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Elog) > 0 {
		i -= len(m.Elog)
		copy(dAtA[i:], m.Elog)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Elog)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Evtype) > 0 {
		i -= len(m.Evtype)
		copy(dAtA[i:], m.Evtype)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Evtype)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Chainid != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Chainid))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgOnchainEventResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOnchainEventResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOnchainEventResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgOnchainManyEvents) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOnchainManyEvents) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOnchainManyEvents) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgOnchainManyEventsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOnchainManyEventsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOnchainManyEventsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Assets != nil {
		l = m.Assets.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateAssetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgOnchainEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Chainid != 0 {
		n += 1 + sovTx(uint64(m.Chainid))
	}
	l = len(m.Evtype)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Elog)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgOnchainEventResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgOnchainManyEvents) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgOnchainManyEventsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Assets == nil {
				m.Assets = &MultiChainAssetParam{}
			}
			if err := m.Assets.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateAssetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateAssetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateAssetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgOnchainEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgOnchainEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOnchainEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chainid", wireType)
			}
			m.Chainid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chainid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evtype", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Evtype = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elog", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elog = append(m.Elog[:0], dAtA[iNdEx:postIndex]...)
			if m.Elog == nil {
				m.Elog = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgOnchainEventResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgOnchainEventResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOnchainEventResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgOnchainManyEvents) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgOnchainManyEvents: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOnchainManyEvents: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &MsgOnchainEvent{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgOnchainManyEventsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgOnchainManyEventsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOnchainManyEventsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
