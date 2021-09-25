// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sgn/farming/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// MsgClaimRewards claims the farming rewards from a single pool for an account.
type MsgClaimRewards struct {
	// pool_name defines the pool name to claim from.
	PoolName string `protobuf:"bytes,1,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	// address defines the account address to claim for.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// sender defines the SGN account sending the claim Msg.
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgClaimRewards) Reset()         { *m = MsgClaimRewards{} }
func (m *MsgClaimRewards) String() string { return proto.CompactTextString(m) }
func (*MsgClaimRewards) ProtoMessage()    {}
func (*MsgClaimRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d6ed81011549f9, []int{0}
}
func (m *MsgClaimRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimRewards.Merge(m, src)
}
func (m *MsgClaimRewards) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimRewards.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimRewards proto.InternalMessageInfo

// MsgClaimRewardsResponse defines the Msg/ClaimRewards response type.
type MsgClaimRewardsResponse struct {
}

func (m *MsgClaimRewardsResponse) Reset()         { *m = MsgClaimRewardsResponse{} }
func (m *MsgClaimRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimRewardsResponse) ProtoMessage()    {}
func (*MsgClaimRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d6ed81011549f9, []int{1}
}
func (m *MsgClaimRewardsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimRewardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimRewardsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimRewardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimRewardsResponse.Merge(m, src)
}
func (m *MsgClaimRewardsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimRewardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimRewardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimRewardsResponse proto.InternalMessageInfo

// MsgClaimRewards claims all the farming rewards for an account.
type MsgClaimAllRewards struct {
	// address defines the account address to claim for.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// sender defines the SGN account sending the claim Msg.
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgClaimAllRewards) Reset()         { *m = MsgClaimAllRewards{} }
func (m *MsgClaimAllRewards) String() string { return proto.CompactTextString(m) }
func (*MsgClaimAllRewards) ProtoMessage()    {}
func (*MsgClaimAllRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d6ed81011549f9, []int{2}
}
func (m *MsgClaimAllRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimAllRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimAllRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimAllRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimAllRewards.Merge(m, src)
}
func (m *MsgClaimAllRewards) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimAllRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimAllRewards.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimAllRewards proto.InternalMessageInfo

// MsgClaimRewardsResponse defines the Msg/ClaimRewards response type.
type MsgClaimAllRewardsResponse struct {
}

func (m *MsgClaimAllRewardsResponse) Reset()         { *m = MsgClaimAllRewardsResponse{} }
func (m *MsgClaimAllRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimAllRewardsResponse) ProtoMessage()    {}
func (*MsgClaimAllRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d6ed81011549f9, []int{3}
}
func (m *MsgClaimAllRewardsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimAllRewardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimAllRewardsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimAllRewardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimAllRewardsResponse.Merge(m, src)
}
func (m *MsgClaimAllRewardsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimAllRewardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimAllRewardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimAllRewardsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgClaimRewards)(nil), "sgn.farming.v1.MsgClaimRewards")
	proto.RegisterType((*MsgClaimRewardsResponse)(nil), "sgn.farming.v1.MsgClaimRewardsResponse")
	proto.RegisterType((*MsgClaimAllRewards)(nil), "sgn.farming.v1.MsgClaimAllRewards")
	proto.RegisterType((*MsgClaimAllRewardsResponse)(nil), "sgn.farming.v1.MsgClaimAllRewardsResponse")
}

func init() { proto.RegisterFile("sgn/farming/v1/tx.proto", fileDescriptor_10d6ed81011549f9) }

var fileDescriptor_10d6ed81011549f9 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4e, 0xcf, 0xd3,
	0x4f, 0x4b, 0x2c, 0xca, 0xcd, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2b, 0x4e, 0xcf, 0xd3, 0x83, 0x4a, 0xe8, 0x95, 0x19, 0x4a, 0x89,
	0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0xa5, 0xf4, 0x41, 0x2c, 0x88, 0x2a, 0xa5, 0x2c, 0x2e, 0x7e, 0xdf,
	0xe2, 0x74, 0xe7, 0x9c, 0xc4, 0xcc, 0xdc, 0xa0, 0xd4, 0xf2, 0xc4, 0xa2, 0x94, 0x62, 0x21, 0x69,
	0x2e, 0xce, 0x82, 0xfc, 0xfc, 0x9c, 0xf8, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x0e, 0x90, 0x80, 0x5f, 0x62, 0x6e, 0xaa, 0x90, 0x04, 0x17, 0x7b, 0x62, 0x4a, 0x4a,
	0x51, 0x6a, 0x71, 0xb1, 0x04, 0x13, 0x58, 0x0a, 0xc6, 0x15, 0x12, 0xe3, 0x62, 0x2b, 0x4e, 0xcd,
	0x4b, 0x49, 0x2d, 0x92, 0x60, 0x06, 0x4b, 0x40, 0x79, 0x56, 0x1c, 0x1d, 0x0b, 0xe4, 0x19, 0x5e,
	0x2c, 0x90, 0x67, 0x50, 0x92, 0xe4, 0x12, 0x47, 0xb3, 0x2b, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf,
	0x38, 0x55, 0x29, 0x80, 0x4b, 0x08, 0x26, 0xe5, 0x98, 0x93, 0x03, 0x73, 0x09, 0x92, 0x65, 0x8c,
	0xb8, 0x2c, 0x63, 0xc2, 0x61, 0x99, 0x0c, 0x97, 0x14, 0xa6, 0x89, 0x30, 0xfb, 0x8c, 0x0e, 0x31,
	0x72, 0x31, 0xfb, 0x16, 0xa7, 0x0b, 0x45, 0x70, 0xf1, 0xa0, 0xf8, 0x5d, 0x5e, 0x0f, 0x35, 0xd4,
	0xf4, 0xd0, 0x1c, 0x2c, 0xa5, 0x4e, 0x40, 0x01, 0xcc, 0x06, 0xa1, 0x44, 0x2e, 0x7e, 0x74, 0xef,
	0x28, 0xe1, 0xd2, 0x8b, 0x50, 0x23, 0xa5, 0x45, 0x58, 0x0d, 0xcc, 0x0a, 0x27, 0xdf, 0x15, 0x8f,
	0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x3f, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x39, 0x35, 0x27, 0xb5, 0x48, 0x37, 0x2f,
	0xb5, 0xa4, 0x3c, 0xbf, 0x28, 0x5b, 0xbf, 0x38, 0x3d, 0x4f, 0xb7, 0xcc, 0x48, 0xbf, 0x02, 0x9e,
	0x6c, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x29, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x65, 0xcd, 0x60, 0x6a, 0x52, 0x02, 0x00, 0x00,
}

func (this *MsgClaimRewardsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgClaimRewardsResponse)
	if !ok {
		that2, ok := that.(MsgClaimRewardsResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *MsgClaimAllRewardsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgClaimAllRewardsResponse)
	if !ok {
		that2, ok := that.(MsgClaimAllRewardsResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
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
	// ClaimRewards defines a method to claim the farming rewards from a single pool.
	ClaimRewards(ctx context.Context, in *MsgClaimRewards, opts ...grpc.CallOption) (*MsgClaimRewardsResponse, error)
	// ClaimAllRewards defines a method to claim the farming rewards from all the staked pools.
	ClaimAllRewards(ctx context.Context, in *MsgClaimAllRewards, opts ...grpc.CallOption) (*MsgClaimAllRewardsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ClaimRewards(ctx context.Context, in *MsgClaimRewards, opts ...grpc.CallOption) (*MsgClaimRewardsResponse, error) {
	out := new(MsgClaimRewardsResponse)
	err := c.cc.Invoke(ctx, "/sgn.farming.v1.Msg/ClaimRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimAllRewards(ctx context.Context, in *MsgClaimAllRewards, opts ...grpc.CallOption) (*MsgClaimAllRewardsResponse, error) {
	out := new(MsgClaimAllRewardsResponse)
	err := c.cc.Invoke(ctx, "/sgn.farming.v1.Msg/ClaimAllRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ClaimRewards defines a method to claim the farming rewards from a single pool.
	ClaimRewards(context.Context, *MsgClaimRewards) (*MsgClaimRewardsResponse, error)
	// ClaimAllRewards defines a method to claim the farming rewards from all the staked pools.
	ClaimAllRewards(context.Context, *MsgClaimAllRewards) (*MsgClaimAllRewardsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ClaimRewards(ctx context.Context, req *MsgClaimRewards) (*MsgClaimRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimRewards not implemented")
}
func (*UnimplementedMsgServer) ClaimAllRewards(ctx context.Context, req *MsgClaimAllRewards) (*MsgClaimAllRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimAllRewards not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ClaimRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimRewards)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgn.farming.v1.Msg/ClaimRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimRewards(ctx, req.(*MsgClaimRewards))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimAllRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimAllRewards)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimAllRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgn.farming.v1.Msg/ClaimAllRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimAllRewards(ctx, req.(*MsgClaimAllRewards))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sgn.farming.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClaimRewards",
			Handler:    _Msg_ClaimRewards_Handler,
		},
		{
			MethodName: "ClaimAllRewards",
			Handler:    _Msg_ClaimAllRewards_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sgn/farming/v1/tx.proto",
}

func (m *MsgClaimRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PoolName) > 0 {
		i -= len(m.PoolName)
		copy(dAtA[i:], m.PoolName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PoolName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimRewardsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimRewardsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimRewardsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgClaimAllRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimAllRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimAllRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimAllRewardsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimAllRewardsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimAllRewardsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgClaimRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgClaimRewardsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgClaimAllRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgClaimAllRewardsResponse) Size() (n int) {
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
func (m *MsgClaimRewards) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgClaimRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolName", wireType)
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
			m.PoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
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
func (m *MsgClaimRewardsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgClaimRewardsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgClaimAllRewards) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgClaimAllRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimAllRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
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
func (m *MsgClaimAllRewardsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgClaimAllRewardsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimAllRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
