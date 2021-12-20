// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deweb/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type WalletRecordResponse struct {
	Owner        string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Address      string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	EncryptedKey string `protobuf:"bytes,3,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encrypted_key,omitempty"`
	Chain        string `protobuf:"bytes,4,opt,name=chain,proto3" json:"chain,omitempty"`
	Deleted      bool   `protobuf:"varint,5,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (m *WalletRecordResponse) Reset()         { *m = WalletRecordResponse{} }
func (m *WalletRecordResponse) String() string { return proto.CompactTextString(m) }
func (*WalletRecordResponse) ProtoMessage()    {}
func (*WalletRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa3f738e95cca609, []int{0}
}
func (m *WalletRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WalletRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WalletRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WalletRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletRecordResponse.Merge(m, src)
}
func (m *WalletRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *WalletRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WalletRecordResponse proto.InternalMessageInfo

func (m *WalletRecordResponse) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *WalletRecordResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *WalletRecordResponse) GetEncryptedKey() string {
	if m != nil {
		return m.EncryptedKey
	}
	return ""
}

func (m *WalletRecordResponse) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *WalletRecordResponse) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

type QueryFilterUserWalletRecordsRequest struct {
	Owner   string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Chain   string `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	Deleted bool   `protobuf:"varint,4,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Limit   int32  `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  int32  `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (m *QueryFilterUserWalletRecordsRequest) Reset()         { *m = QueryFilterUserWalletRecordsRequest{} }
func (m *QueryFilterUserWalletRecordsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFilterUserWalletRecordsRequest) ProtoMessage()    {}
func (*QueryFilterUserWalletRecordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa3f738e95cca609, []int{1}
}
func (m *QueryFilterUserWalletRecordsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFilterUserWalletRecordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFilterUserWalletRecordsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFilterUserWalletRecordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFilterUserWalletRecordsRequest.Merge(m, src)
}
func (m *QueryFilterUserWalletRecordsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFilterUserWalletRecordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFilterUserWalletRecordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFilterUserWalletRecordsRequest proto.InternalMessageInfo

func (m *QueryFilterUserWalletRecordsRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *QueryFilterUserWalletRecordsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QueryFilterUserWalletRecordsRequest) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *QueryFilterUserWalletRecordsRequest) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *QueryFilterUserWalletRecordsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryFilterUserWalletRecordsRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type QueryFilterUserWalletRecordsResponse struct {
	Records []*WalletRecordResponse `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (m *QueryFilterUserWalletRecordsResponse) Reset()         { *m = QueryFilterUserWalletRecordsResponse{} }
func (m *QueryFilterUserWalletRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFilterUserWalletRecordsResponse) ProtoMessage()    {}
func (*QueryFilterUserWalletRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa3f738e95cca609, []int{2}
}
func (m *QueryFilterUserWalletRecordsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFilterUserWalletRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFilterUserWalletRecordsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFilterUserWalletRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFilterUserWalletRecordsResponse.Merge(m, src)
}
func (m *QueryFilterUserWalletRecordsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFilterUserWalletRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFilterUserWalletRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFilterUserWalletRecordsResponse proto.InternalMessageInfo

func (m *QueryFilterUserWalletRecordsResponse) GetRecords() []*WalletRecordResponse {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*WalletRecordResponse)(nil), "dewebservices.deweb.deweb.WalletRecordResponse")
	proto.RegisterType((*QueryFilterUserWalletRecordsRequest)(nil), "dewebservices.deweb.deweb.QueryFilterUserWalletRecordsRequest")
	proto.RegisterType((*QueryFilterUserWalletRecordsResponse)(nil), "dewebservices.deweb.deweb.QueryFilterUserWalletRecordsResponse")
}

func init() { proto.RegisterFile("deweb/query.proto", fileDescriptor_fa3f738e95cca609) }

var fileDescriptor_fa3f738e95cca609 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0x4d, 0x93, 0xea, 0xa8, 0x07, 0x87, 0xa2, 0x63, 0x90, 0x25, 0xa4, 0x45, 0x82,
	0xd8, 0x1d, 0x5a, 0xef, 0x0a, 0x1e, 0x14, 0xf1, 0xe4, 0x82, 0x08, 0x5e, 0xca, 0xec, 0xee, 0x6b,
	0x3a, 0x38, 0x99, 0xd9, 0xcc, 0x9b, 0x34, 0xdd, 0xab, 0x9f, 0x40, 0xf0, 0xe8, 0xc7, 0xf0, 0x4b,
	0xf4, 0x58, 0xf4, 0xe2, 0x51, 0x12, 0x3f, 0x88, 0x64, 0x26, 0x09, 0x0a, 0x26, 0x68, 0x2f, 0xbb,
	0xfc, 0xff, 0xfb, 0xde, 0x7f, 0x7e, 0x3b, 0xef, 0xd1, 0xdb, 0x25, 0x4c, 0x20, 0x17, 0xa3, 0x31,
	0xb8, 0x3a, 0xad, 0x9c, 0xf5, 0x96, 0xdd, 0x0b, 0x16, 0x82, 0x3b, 0x53, 0x05, 0x60, 0x1a, 0x54,
	0x7c, 0x76, 0xee, 0x0f, 0xac, 0x1d, 0x68, 0x10, 0xb2, 0x52, 0x42, 0x1a, 0x63, 0xbd, 0xf4, 0xca,
	0x1a, 0x8c, 0x8d, 0x9d, 0x87, 0x85, 0xc5, 0xa1, 0x45, 0x91, 0x4b, 0x84, 0x98, 0x28, 0xce, 0x0e,
	0x73, 0xf0, 0xf2, 0x50, 0x54, 0x72, 0xa0, 0x4c, 0x28, 0x8e, 0xb5, 0xbd, 0xcf, 0x84, 0xee, 0xbe,
	0x95, 0x5a, 0x83, 0xcf, 0xa0, 0xb0, 0xae, 0xcc, 0x00, 0x2b, 0x6b, 0x10, 0xd8, 0x2e, 0x6d, 0xd9,
	0x89, 0x01, 0xc7, 0x49, 0x97, 0xf4, 0xaf, 0x67, 0x51, 0x30, 0x4e, 0x77, 0x64, 0x59, 0x3a, 0x40,
	0xe4, 0x5b, 0xc1, 0x5f, 0x4a, 0xb6, 0x47, 0x6f, 0x81, 0x29, 0x5c, 0x5d, 0x79, 0x28, 0x8f, 0xdf,
	0x43, 0xcd, 0x9b, 0xe1, 0xfb, 0xcd, 0x95, 0xf9, 0x0a, 0xea, 0x79, 0x68, 0x71, 0x2a, 0x95, 0xe1,
	0xdb, 0x31, 0x34, 0x88, 0x79, 0x68, 0x09, 0x1a, 0x3c, 0x94, 0xbc, 0xd5, 0x25, 0xfd, 0x6b, 0xd9,
	0x52, 0xf6, 0xbe, 0x10, 0xba, 0xf7, 0x7a, 0xfe, 0x03, 0xcf, 0x95, 0xf6, 0xe0, 0xde, 0x20, 0xb8,
	0xdf, 0x61, 0x31, 0x83, 0xd1, 0x18, 0xd0, 0xff, 0x37, 0xec, 0x8a, 0xa3, 0xb9, 0x86, 0x63, 0xfb,
	0x0f, 0x8e, 0x79, 0xbd, 0x56, 0x43, 0xe5, 0x03, 0x5f, 0x2b, 0x8b, 0x82, 0xdd, 0xa1, 0x6d, 0x7b,
	0x72, 0x82, 0xe0, 0x79, 0x3b, 0xd8, 0x0b, 0xd5, 0x1b, 0xd1, 0xfd, 0xcd, 0xd0, 0x8b, 0x2b, 0x7e,
	0x49, 0x77, 0x5c, 0xb4, 0x38, 0xe9, 0x36, 0xfb, 0x37, 0x8e, 0x44, 0xba, 0x76, 0xe4, 0xe9, 0xdf,
	0x86, 0x94, 0x2d, 0xfb, 0x8f, 0xbe, 0x12, 0xda, 0x0a, 0x67, 0xb2, 0x0b, 0x42, 0xef, 0xae, 0x39,
	0x98, 0x3d, 0xd9, 0x90, 0xff, 0x0f, 0xd7, 0xdc, 0x79, 0x7a, 0xe5, 0xfe, 0xc8, 0xdb, 0x7b, 0xf4,
	0xe1, 0xdb, 0xcf, 0x4f, 0x5b, 0x0f, 0xd8, 0xbe, 0x88, 0xeb, 0x0e, 0xe7, 0x1e, 0x9c, 0x91, 0xfa,
	0x78, 0x12, 0xca, 0x71, 0xb5, 0xa7, 0x5a, 0xa1, 0x7f, 0xf6, 0xe2, 0x62, 0x9a, 0x90, 0xcb, 0x69,
	0x42, 0x7e, 0x4c, 0x13, 0xf2, 0x71, 0x96, 0x34, 0x2e, 0x67, 0x49, 0xe3, 0xfb, 0x2c, 0x69, 0xbc,
	0x3b, 0x18, 0x28, 0x7f, 0x3a, 0xce, 0xd3, 0xc2, 0x0e, 0x63, 0xd2, 0xc1, 0x92, 0x69, 0x11, 0x7c,
	0xbe, 0x78, 0xfb, 0xba, 0x02, 0xcc, 0xdb, 0x61, 0xd7, 0x1f, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x2e, 0x76, 0x0b, 0xc8, 0x65, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a list of filterUserKeyRecords items.
	FilterUserWalletRecords(ctx context.Context, in *QueryFilterUserWalletRecordsRequest, opts ...grpc.CallOption) (*QueryFilterUserWalletRecordsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) FilterUserWalletRecords(ctx context.Context, in *QueryFilterUserWalletRecordsRequest, opts ...grpc.CallOption) (*QueryFilterUserWalletRecordsResponse, error) {
	out := new(QueryFilterUserWalletRecordsResponse)
	err := c.cc.Invoke(ctx, "/dewebservices.deweb.deweb.Query/FilterUserWalletRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a list of filterUserKeyRecords items.
	FilterUserWalletRecords(context.Context, *QueryFilterUserWalletRecordsRequest) (*QueryFilterUserWalletRecordsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) FilterUserWalletRecords(ctx context.Context, req *QueryFilterUserWalletRecordsRequest) (*QueryFilterUserWalletRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterUserWalletRecords not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_FilterUserWalletRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFilterUserWalletRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FilterUserWalletRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dewebservices.deweb.deweb.Query/FilterUserWalletRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FilterUserWalletRecords(ctx, req.(*QueryFilterUserWalletRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dewebservices.deweb.deweb.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FilterUserWalletRecords",
			Handler:    _Query_FilterUserWalletRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deweb/query.proto",
}

func (m *WalletRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WalletRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WalletRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Deleted {
		i--
		if m.Deleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.EncryptedKey) > 0 {
		i -= len(m.EncryptedKey)
		copy(dAtA[i:], m.EncryptedKey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EncryptedKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFilterUserWalletRecordsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFilterUserWalletRecordsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFilterUserWalletRecordsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Offset != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x30
	}
	if m.Limit != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x28
	}
	if m.Deleted {
		i--
		if m.Deleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFilterUserWalletRecordsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFilterUserWalletRecordsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFilterUserWalletRecordsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WalletRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.EncryptedKey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Deleted {
		n += 2
	}
	return n
}

func (m *QueryFilterUserWalletRecordsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Deleted {
		n += 2
	}
	if m.Limit != 0 {
		n += 1 + sovQuery(uint64(m.Limit))
	}
	if m.Offset != 0 {
		n += 1 + sovQuery(uint64(m.Offset))
	}
	return n
}

func (m *QueryFilterUserWalletRecordsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WalletRecordResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: WalletRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WalletRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deleted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryFilterUserWalletRecordsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryFilterUserWalletRecordsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFilterUserWalletRecordsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deleted = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryFilterUserWalletRecordsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryFilterUserWalletRecordsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFilterUserWalletRecordsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, &WalletRecordResponse{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
