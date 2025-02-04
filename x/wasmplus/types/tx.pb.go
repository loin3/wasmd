// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/wasm/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_line_lbm_sdk_types "github.com/line/lbm-sdk/types"
	types1 "github.com/line/lbm-sdk/types"
	github_com_line_wasmd_x_wasm_types "github.com/line/wasmd/x/wasm/types"
	types "github.com/line/wasmd/x/wasm/types"
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

// MsgStoreCodeAndInstantiateContract submit Wasm code to the system and instantiate a contract using it.
type MsgStoreCodeAndInstantiateContract struct {
	// Sender is the that actor that signed the messages
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// WASMByteCode can be raw or gzip compressed
	WASMByteCode          []byte              `protobuf:"bytes,2,opt,name=wasm_byte_code,json=wasmByteCode,proto3" json:"wasm_byte_code,omitempty"`
	InstantiatePermission *types.AccessConfig `protobuf:"bytes,5,opt,name=instantiate_permission,json=instantiatePermission,proto3" json:"instantiate_permission,omitempty"`
	// Admin is an optional address that can execute migrations
	Admin string `protobuf:"bytes,6,opt,name=admin,proto3" json:"admin,omitempty"`
	// Label is optional metadata to be stored with a contract instance.
	Label string `protobuf:"bytes,7,opt,name=label,proto3" json:"label,omitempty"`
	// Msg json encoded message to be passed to the contract on instantiation
	Msg github_com_line_wasmd_x_wasm_types.RawContractMessage `protobuf:"bytes,8,opt,name=msg,proto3,casttype=github.com/line/wasmd/x/wasm/types.RawContractMessage" json:"msg,omitempty"`
	// Funds coins that are transferred to the contract on instantiation
	Funds github_com_line_lbm_sdk_types.Coins `protobuf:"bytes,9,rep,name=funds,proto3,castrepeated=github.com/line/lbm-sdk/types.Coins" json:"funds"`
}

func (m *MsgStoreCodeAndInstantiateContract) Reset()         { *m = MsgStoreCodeAndInstantiateContract{} }
func (m *MsgStoreCodeAndInstantiateContract) String() string { return proto.CompactTextString(m) }
func (*MsgStoreCodeAndInstantiateContract) ProtoMessage()    {}
func (*MsgStoreCodeAndInstantiateContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_751e1d2b9f9bf9e8, []int{0}
}
func (m *MsgStoreCodeAndInstantiateContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStoreCodeAndInstantiateContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStoreCodeAndInstantiateContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStoreCodeAndInstantiateContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStoreCodeAndInstantiateContract.Merge(m, src)
}
func (m *MsgStoreCodeAndInstantiateContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgStoreCodeAndInstantiateContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStoreCodeAndInstantiateContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStoreCodeAndInstantiateContract proto.InternalMessageInfo

// MsgStoreCodeAndInstantiateContractResponse returns store and instantiate result data.
type MsgStoreCodeAndInstantiateContractResponse struct {
	// CodeID is the reference to the stored WASM code
	CodeID uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	// Address is the bech32 address of the new contract instance
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Data contains base64-encoded bytes to returned from the contract
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgStoreCodeAndInstantiateContractResponse) Reset() {
	*m = MsgStoreCodeAndInstantiateContractResponse{}
}
func (m *MsgStoreCodeAndInstantiateContractResponse) String() string {
	return proto.CompactTextString(m)
}
func (*MsgStoreCodeAndInstantiateContractResponse) ProtoMessage() {}
func (*MsgStoreCodeAndInstantiateContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_751e1d2b9f9bf9e8, []int{1}
}
func (m *MsgStoreCodeAndInstantiateContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStoreCodeAndInstantiateContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStoreCodeAndInstantiateContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStoreCodeAndInstantiateContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStoreCodeAndInstantiateContractResponse.Merge(m, src)
}
func (m *MsgStoreCodeAndInstantiateContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStoreCodeAndInstantiateContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStoreCodeAndInstantiateContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStoreCodeAndInstantiateContractResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgStoreCodeAndInstantiateContract)(nil), "lbm.wasm.v1.MsgStoreCodeAndInstantiateContract")
	proto.RegisterType((*MsgStoreCodeAndInstantiateContractResponse)(nil), "lbm.wasm.v1.MsgStoreCodeAndInstantiateContractResponse")
}

func init() { proto.RegisterFile("lbm/wasm/v1/tx.proto", fileDescriptor_751e1d2b9f9bf9e8) }

var fileDescriptor_751e1d2b9f9bf9e8 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0x6e, 0x4c, 0xda, 0xdd, 0x9d, 0x5d, 0x64, 0x19, 0xea, 0x12, 0x8b, 0x24, 0xa5, 0x0b, 0x52,
	0x14, 0x27, 0xb4, 0xa2, 0xe2, 0xb1, 0xed, 0x5e, 0xba, 0x52, 0x90, 0x2c, 0x22, 0x88, 0x50, 0x26,
	0x99, 0xd9, 0x18, 0x4c, 0x66, 0x4a, 0xde, 0x69, 0x77, 0x7b, 0xf2, 0x07, 0x78, 0x11, 0xfc, 0x17,
	0x1e, 0xfc, 0x1d, 0x3d, 0xee, 0xd1, 0x53, 0xd5, 0xf6, 0x5f, 0x78, 0x92, 0x49, 0x9a, 0x75, 0xf1,
	0x60, 0xf1, 0x36, 0xcf, 0xfb, 0xf1, 0xbc, 0xef, 0x3c, 0xcf, 0x0c, 0xaa, 0x27, 0x41, 0xea, 0x5d,
	0x50, 0x48, 0xbd, 0x59, 0xc7, 0x53, 0x97, 0x64, 0x92, 0x49, 0x25, 0xf1, 0x7e, 0x12, 0xa4, 0x44,
	0x47, 0xc9, 0xac, 0xd3, 0xa8, 0x47, 0x32, 0x92, 0x79, 0xdc, 0xd3, 0xa7, 0xa2, 0xa4, 0xe1, 0x84,
	0x12, 0x52, 0x09, 0x5e, 0x40, 0x81, 0x7b, 0xb3, 0x4e, 0xc0, 0x15, 0xed, 0x78, 0xa1, 0x8c, 0xc5,
	0x26, 0x7f, 0x4f, 0xe7, 0x73, 0xe2, 0x6b, 0xf6, 0xf9, 0x84, 0x43, 0x91, 0x6d, 0x7d, 0x35, 0x51,
	0x6b, 0x04, 0xd1, 0x99, 0x92, 0x19, 0x1f, 0x48, 0xc6, 0x7b, 0x82, 0x0d, 0x05, 0x28, 0x2a, 0x54,
	0x4c, 0x15, 0x1f, 0x48, 0xa1, 0x32, 0x1a, 0x2a, 0x7c, 0x84, 0x6a, 0xc0, 0x05, 0xe3, 0x99, 0x6d,
	0x34, 0x8d, 0xf6, 0x9e, 0xbf, 0x41, 0xf8, 0x29, 0xba, 0xad, 0x59, 0xc7, 0xc1, 0x5c, 0xf1, 0x71,
	0x28, 0x19, 0xb7, 0x6f, 0x35, 0x8d, 0xf6, 0x41, 0xff, 0x70, 0xb5, 0x74, 0x0f, 0x5e, 0xf7, 0xce,
	0x46, 0xfd, 0xb9, 0xca, 0x79, 0xfd, 0x03, 0x5d, 0x57, 0x22, 0xfc, 0x0a, 0x1d, 0xc5, 0x7f, 0xc6,
	0x8c, 0x27, 0x3c, 0x4b, 0x63, 0x80, 0x58, 0x0a, 0xbb, 0xda, 0x34, 0xda, 0xfb, 0x5d, 0x87, 0x94,
	0x5b, 0x97, 0xb7, 0x27, 0xbd, 0x30, 0xe4, 0x00, 0x03, 0x29, 0xce, 0xe3, 0xc8, 0xbf, 0x73, 0xa3,
	0xfb, 0xe5, 0x75, 0x33, 0xae, 0xa3, 0x2a, 0x65, 0x69, 0x2c, 0xec, 0x5a, 0xbe, 0x65, 0x01, 0x74,
	0x34, 0xa1, 0x01, 0x4f, 0xec, 0x9d, 0x22, 0x9a, 0x03, 0xfc, 0x02, 0x99, 0x29, 0x44, 0xf6, 0x6e,
	0xbe, 0xef, 0xf3, 0x5f, 0x4b, 0xf7, 0x49, 0x14, 0xab, 0x77, 0xd3, 0x80, 0x84, 0x32, 0xf5, 0x92,
	0x58, 0xf0, 0x5c, 0x2f, 0xe6, 0x5d, 0x16, 0xba, 0x15, 0xa2, 0xf9, 0xf4, 0xa2, 0xd4, 0x64, 0xc4,
	0x01, 0x68, 0xc4, 0x7d, 0xcd, 0x82, 0xdf, 0xa2, 0xea, 0xf9, 0x54, 0x30, 0xb0, 0xf7, 0x9a, 0x66,
	0x7b, 0xbf, 0x7b, 0x97, 0x14, 0xa6, 0x10, 0x6d, 0x0a, 0xd9, 0x98, 0x42, 0x06, 0x32, 0x16, 0xfd,
	0x87, 0x8b, 0xa5, 0x5b, 0xf9, 0xf2, 0xdd, 0x3d, 0xfe, 0x7b, 0x5a, 0x12, 0xa4, 0x8f, 0x80, 0xbd,
	0xdf, 0x0c, 0xd2, 0xb5, 0xe0, 0x17, 0xa4, 0xa7, 0xd6, 0xae, 0x79, 0x68, 0x9d, 0x5a, 0xbb, 0xd6,
	0x61, 0xb5, 0xf5, 0x01, 0x3d, 0xd8, 0xee, 0x97, 0xcf, 0x61, 0x22, 0x05, 0x70, 0x7c, 0x8c, 0x76,
	0xb4, 0x2b, 0xe3, 0x98, 0xe5, 0xc6, 0x59, 0x7d, 0xb4, 0x5a, 0xba, 0x35, 0xdd, 0x38, 0x3c, 0xf1,
	0x6b, 0x3a, 0x35, 0x64, 0xd8, 0x46, 0x3b, 0x94, 0xb1, 0x8c, 0x03, 0xe4, 0xee, 0xed, 0xf9, 0x25,
	0xc4, 0x18, 0x59, 0x8c, 0x2a, 0x6a, 0x9b, 0x5a, 0x24, 0x3f, 0x3f, 0x77, 0x3f, 0x1b, 0xc8, 0x1c,
	0x41, 0x84, 0x3f, 0x1a, 0xc8, 0xdd, 0xf6, 0x6c, 0x3c, 0x72, 0xe3, 0xfd, 0x92, 0xed, 0x7b, 0x37,
	0x9e, 0xfd, 0x67, 0x43, 0x79, 0xd1, 0xfe, 0xc9, 0xe2, 0xa7, 0x53, 0x59, 0xac, 0x1c, 0xe3, 0x6a,
	0xe5, 0x18, 0x3f, 0x56, 0x8e, 0xf1, 0x69, 0xed, 0x54, 0xae, 0xd6, 0x4e, 0xe5, 0xdb, 0xda, 0xa9,
	0xbc, 0xb9, 0xff, 0x2f, 0x6b, 0x27, 0xc9, 0x14, 0x0a, 0xd5, 0x83, 0x5a, 0xfe, 0x29, 0x1e, 0xff,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x92, 0x19, 0xa9, 0x26, 0x8d, 0x03, 0x00, 0x00,
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
	// StoreCodeAndInstantiateContract upload code and instantiate a contract using it
	StoreCodeAndInstantiateContract(ctx context.Context, in *MsgStoreCodeAndInstantiateContract, opts ...grpc.CallOption) (*MsgStoreCodeAndInstantiateContractResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) StoreCodeAndInstantiateContract(ctx context.Context, in *MsgStoreCodeAndInstantiateContract, opts ...grpc.CallOption) (*MsgStoreCodeAndInstantiateContractResponse, error) {
	out := new(MsgStoreCodeAndInstantiateContractResponse)
	err := c.cc.Invoke(ctx, "/lbm.wasm.v1.Msg/StoreCodeAndInstantiateContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// StoreCodeAndInstantiateContract upload code and instantiate a contract using it
	StoreCodeAndInstantiateContract(context.Context, *MsgStoreCodeAndInstantiateContract) (*MsgStoreCodeAndInstantiateContractResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) StoreCodeAndInstantiateContract(ctx context.Context, req *MsgStoreCodeAndInstantiateContract) (*MsgStoreCodeAndInstantiateContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreCodeAndInstantiateContract not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_StoreCodeAndInstantiateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStoreCodeAndInstantiateContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StoreCodeAndInstantiateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lbm.wasm.v1.Msg/StoreCodeAndInstantiateContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StoreCodeAndInstantiateContract(ctx, req.(*MsgStoreCodeAndInstantiateContract))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lbm.wasm.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreCodeAndInstantiateContract",
			Handler:    _Msg_StoreCodeAndInstantiateContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lbm/wasm/v1/tx.proto",
}

func (m *MsgStoreCodeAndInstantiateContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStoreCodeAndInstantiateContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStoreCodeAndInstantiateContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Funds) > 0 {
		for iNdEx := len(m.Funds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Funds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0x32
	}
	if m.InstantiatePermission != nil {
		{
			size, err := m.InstantiatePermission.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.WASMByteCode) > 0 {
		i -= len(m.WASMByteCode)
		copy(dAtA[i:], m.WASMByteCode)
		i = encodeVarintTx(dAtA, i, uint64(len(m.WASMByteCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStoreCodeAndInstantiateContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStoreCodeAndInstantiateContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStoreCodeAndInstantiateContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
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
	if m.CodeID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x8
	}
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
func (m *MsgStoreCodeAndInstantiateContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.WASMByteCode)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.InstantiatePermission != nil {
		l = m.InstantiatePermission.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Funds) > 0 {
		for _, e := range m.Funds {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgStoreCodeAndInstantiateContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeID != 0 {
		n += 1 + sovTx(uint64(m.CodeID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgStoreCodeAndInstantiateContract) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgStoreCodeAndInstantiateContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStoreCodeAndInstantiateContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WASMByteCode", wireType)
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
			m.WASMByteCode = append(m.WASMByteCode[:0], dAtA[iNdEx:postIndex]...)
			if m.WASMByteCode == nil {
				m.WASMByteCode = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstantiatePermission", wireType)
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
			if m.InstantiatePermission == nil {
				m.InstantiatePermission = &types.AccessConfig{}
			}
			if err := m.InstantiatePermission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
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
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
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
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
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
			m.Msg = append(m.Msg[:0], dAtA[iNdEx:postIndex]...)
			if m.Msg == nil {
				m.Msg = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
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
			m.Funds = append(m.Funds, types1.Coin{})
			if err := m.Funds[len(m.Funds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgStoreCodeAndInstantiateContractResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgStoreCodeAndInstantiateContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStoreCodeAndInstantiateContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
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
