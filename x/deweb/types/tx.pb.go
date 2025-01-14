// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deweb/tx.proto

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

type MsgSaveWallet struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Address      string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	EncryptedKey string `protobuf:"bytes,3,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encrypted_key,omitempty"`
	Chain        string `protobuf:"bytes,4,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (m *MsgSaveWallet) Reset()         { *m = MsgSaveWallet{} }
func (m *MsgSaveWallet) String() string { return proto.CompactTextString(m) }
func (*MsgSaveWallet) ProtoMessage()    {}
func (*MsgSaveWallet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0add4e41aef49c7d, []int{0}
}
func (m *MsgSaveWallet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSaveWallet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSaveWallet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSaveWallet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSaveWallet.Merge(m, src)
}
func (m *MsgSaveWallet) XXX_Size() int {
	return m.Size()
}
func (m *MsgSaveWallet) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSaveWallet.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSaveWallet proto.InternalMessageInfo

func (m *MsgSaveWallet) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSaveWallet) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgSaveWallet) GetEncryptedKey() string {
	if m != nil {
		return m.EncryptedKey
	}
	return ""
}

func (m *MsgSaveWallet) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

type MsgSaveWalletResponse struct {
}

func (m *MsgSaveWalletResponse) Reset()         { *m = MsgSaveWalletResponse{} }
func (m *MsgSaveWalletResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSaveWalletResponse) ProtoMessage()    {}
func (*MsgSaveWalletResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0add4e41aef49c7d, []int{1}
}
func (m *MsgSaveWalletResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSaveWalletResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSaveWalletResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSaveWalletResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSaveWalletResponse.Merge(m, src)
}
func (m *MsgSaveWalletResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSaveWalletResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSaveWalletResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSaveWalletResponse proto.InternalMessageInfo

type MsgDeleteWallet struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgDeleteWallet) Reset()         { *m = MsgDeleteWallet{} }
func (m *MsgDeleteWallet) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteWallet) ProtoMessage()    {}
func (*MsgDeleteWallet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0add4e41aef49c7d, []int{2}
}
func (m *MsgDeleteWallet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteWallet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteWallet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteWallet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteWallet.Merge(m, src)
}
func (m *MsgDeleteWallet) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteWallet) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteWallet.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteWallet proto.InternalMessageInfo

func (m *MsgDeleteWallet) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteWallet) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgDeleteWalletResponse struct {
}

func (m *MsgDeleteWalletResponse) Reset()         { *m = MsgDeleteWalletResponse{} }
func (m *MsgDeleteWalletResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteWalletResponse) ProtoMessage()    {}
func (*MsgDeleteWalletResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0add4e41aef49c7d, []int{3}
}
func (m *MsgDeleteWalletResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteWalletResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteWalletResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteWalletResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteWalletResponse.Merge(m, src)
}
func (m *MsgDeleteWalletResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteWalletResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteWalletResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteWalletResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSaveWallet)(nil), "dewebservices.deweb.deweb.MsgSaveWallet")
	proto.RegisterType((*MsgSaveWalletResponse)(nil), "dewebservices.deweb.deweb.MsgSaveWalletResponse")
	proto.RegisterType((*MsgDeleteWallet)(nil), "dewebservices.deweb.deweb.MsgDeleteWallet")
	proto.RegisterType((*MsgDeleteWalletResponse)(nil), "dewebservices.deweb.deweb.MsgDeleteWalletResponse")
}

func init() { proto.RegisterFile("deweb/tx.proto", fileDescriptor_0add4e41aef49c7d) }

var fileDescriptor_0add4e41aef49c7d = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x49, 0x2d, 0x4f,
	0x4d, 0xd2, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x04, 0xf3, 0x8b, 0x53,
	0x8b, 0xca, 0x32, 0x93, 0x53, 0x8b, 0xf5, 0xc0, 0x3c, 0x08, 0xa9, 0x54, 0xc7, 0xc5, 0xeb, 0x5b,
	0x9c, 0x1e, 0x9c, 0x58, 0x96, 0x1a, 0x9e, 0x98, 0x93, 0x93, 0x5a, 0x22, 0x24, 0xc1, 0xc5, 0x9e,
	0x5c, 0x94, 0x9a, 0x58, 0x92, 0x5f, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x82,
	0x64, 0x12, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x98, 0x20, 0x32, 0x50, 0xae, 0x90, 0x32,
	0x17, 0x6f, 0x6a, 0x5e, 0x72, 0x51, 0x65, 0x41, 0x49, 0x6a, 0x4a, 0x7c, 0x76, 0x6a, 0xa5, 0x04,
	0x33, 0x58, 0x9e, 0x07, 0x2e, 0xe8, 0x9d, 0x5a, 0x29, 0x24, 0xc2, 0xc5, 0x9a, 0x9c, 0x91, 0x98,
	0x99, 0x27, 0xc1, 0x02, 0x96, 0x84, 0x70, 0x94, 0xc4, 0xb9, 0x44, 0x51, 0xec, 0x0f, 0x4a, 0x2d,
	0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x55, 0x72, 0xe5, 0xe2, 0xf7, 0x2d, 0x4e, 0x77, 0x49, 0xcd, 0x49,
	0x2d, 0xa1, 0xc0, 0x69, 0x4a, 0x92, 0x5c, 0xe2, 0x68, 0xc6, 0xc0, 0x6c, 0x30, 0xba, 0xcf, 0xc8,
	0xc5, 0xec, 0x5b, 0x9c, 0x2e, 0x94, 0xc1, 0xc5, 0x85, 0xe4, 0x7f, 0x0d, 0x3d, 0x9c, 0x81, 0xa5,
	0x87, 0xe2, 0x52, 0x29, 0x03, 0x62, 0x55, 0xc2, 0x6c, 0x14, 0xca, 0xe3, 0xe2, 0x41, 0xf1, 0x90,
	0x16, 0x7e, 0x13, 0x90, 0xd5, 0x4a, 0x19, 0x11, 0xaf, 0x16, 0x66, 0x9f, 0x93, 0xfb, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85,
	0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0xea, 0x83, 0xcd, 0xd0, 0x85, 0x19, 0x0c, 0xe1, 0xea, 0x57, 0x40, 0xe9, 0x92,
	0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x3a, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x68,
	0x7f, 0x8f, 0x7e, 0x59, 0x02, 0x00, 0x00,
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
	SaveWallet(ctx context.Context, in *MsgSaveWallet, opts ...grpc.CallOption) (*MsgSaveWalletResponse, error)
	DeleteWallet(ctx context.Context, in *MsgDeleteWallet, opts ...grpc.CallOption) (*MsgDeleteWalletResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SaveWallet(ctx context.Context, in *MsgSaveWallet, opts ...grpc.CallOption) (*MsgSaveWalletResponse, error) {
	out := new(MsgSaveWalletResponse)
	err := c.cc.Invoke(ctx, "/dewebservices.deweb.deweb.Msg/SaveWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteWallet(ctx context.Context, in *MsgDeleteWallet, opts ...grpc.CallOption) (*MsgDeleteWalletResponse, error) {
	out := new(MsgDeleteWalletResponse)
	err := c.cc.Invoke(ctx, "/dewebservices.deweb.deweb.Msg/DeleteWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SaveWallet(context.Context, *MsgSaveWallet) (*MsgSaveWalletResponse, error)
	DeleteWallet(context.Context, *MsgDeleteWallet) (*MsgDeleteWalletResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SaveWallet(ctx context.Context, req *MsgSaveWallet) (*MsgSaveWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveWallet not implemented")
}
func (*UnimplementedMsgServer) DeleteWallet(ctx context.Context, req *MsgDeleteWallet) (*MsgDeleteWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWallet not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SaveWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSaveWallet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SaveWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dewebservices.deweb.deweb.Msg/SaveWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SaveWallet(ctx, req.(*MsgSaveWallet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteWallet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dewebservices.deweb.deweb.Msg/DeleteWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteWallet(ctx, req.(*MsgDeleteWallet))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dewebservices.deweb.deweb.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveWallet",
			Handler:    _Msg_SaveWallet_Handler,
		},
		{
			MethodName: "DeleteWallet",
			Handler:    _Msg_DeleteWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deweb/tx.proto",
}

func (m *MsgSaveWallet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSaveWallet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSaveWallet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.EncryptedKey) > 0 {
		i -= len(m.EncryptedKey)
		copy(dAtA[i:], m.EncryptedKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.EncryptedKey)))
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
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSaveWalletResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSaveWalletResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSaveWalletResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteWallet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteWallet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteWallet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
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

func (m *MsgDeleteWalletResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteWalletResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteWalletResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSaveWallet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.EncryptedKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSaveWalletResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteWallet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeleteWalletResponse) Size() (n int) {
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
func (m *MsgSaveWallet) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSaveWallet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSaveWallet: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedKey", wireType)
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
			m.EncryptedKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
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
			m.Chain = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSaveWalletResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSaveWalletResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSaveWalletResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgDeleteWallet) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteWallet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteWallet: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgDeleteWalletResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteWalletResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteWalletResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
