// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: esdt.proto

package systemSmartContracts

import (
	bytes "bytes"
	fmt "fmt"
	github_com_ElrondNetwork_elrond_go_data "github.com/ElrondNetwork/elrond-go/data"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type ESDTData struct {
	OwnerAddress   []byte        `protobuf:"bytes,1,opt,name=OwnerAddress,proto3" json:"OwnerAddress"`
	TokenName      []byte        `protobuf:"bytes,2,opt,name=TokenName,proto3" json:"TokenName"`
	TickerName     []byte        `protobuf:"bytes,3,opt,name=TickerName,proto3" json:"TickerName"`
	Mintable       bool          `protobuf:"varint,4,opt,name=Mintable,proto3" json:"Mintable"`
	Burnable       bool          `protobuf:"varint,5,opt,name=Burnable,proto3" json:"Burnable"`
	CanPause       bool          `protobuf:"varint,6,opt,name=CanPause,proto3" json:"CanPause"`
	CanFreeze      bool          `protobuf:"varint,7,opt,name=CanFreeze,proto3" json:"CanFreeze"`
	CanWipe        bool          `protobuf:"varint,8,opt,name=CanWipe,proto3" json:"CanWipe"`
	Upgradable     bool          `protobuf:"varint,9,opt,name=Upgradable,proto3" json:"CanUpgrade"`
	CanChangeOwner bool          `protobuf:"varint,10,opt,name=CanChangeOwner,proto3" json:"CanChangeOwner"`
	IsPaused       bool          `protobuf:"varint,11,opt,name=IsPaused,proto3" json:"IsPaused"`
	MintedValue    *math_big.Int `protobuf:"bytes,12,opt,name=MintedValue,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"MintedValue"`
	BurntValue     *math_big.Int `protobuf:"bytes,13,opt,name=BurntValue,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"BurntValue"`
	NumDecimals    uint32        `protobuf:"varint,14,opt,name=NumDecimals,proto3" json:"NumDecimals"`
}

func (m *ESDTData) Reset()      { *m = ESDTData{} }
func (*ESDTData) ProtoMessage() {}
func (*ESDTData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e413e402abc6a34c, []int{0}
}
func (m *ESDTData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ESDTData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ESDTData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ESDTData.Merge(m, src)
}
func (m *ESDTData) XXX_Size() int {
	return m.Size()
}
func (m *ESDTData) XXX_DiscardUnknown() {
	xxx_messageInfo_ESDTData.DiscardUnknown(m)
}

var xxx_messageInfo_ESDTData proto.InternalMessageInfo

func (m *ESDTData) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *ESDTData) GetTokenName() []byte {
	if m != nil {
		return m.TokenName
	}
	return nil
}

func (m *ESDTData) GetTickerName() []byte {
	if m != nil {
		return m.TickerName
	}
	return nil
}

func (m *ESDTData) GetMintable() bool {
	if m != nil {
		return m.Mintable
	}
	return false
}

func (m *ESDTData) GetBurnable() bool {
	if m != nil {
		return m.Burnable
	}
	return false
}

func (m *ESDTData) GetCanPause() bool {
	if m != nil {
		return m.CanPause
	}
	return false
}

func (m *ESDTData) GetCanFreeze() bool {
	if m != nil {
		return m.CanFreeze
	}
	return false
}

func (m *ESDTData) GetCanWipe() bool {
	if m != nil {
		return m.CanWipe
	}
	return false
}

func (m *ESDTData) GetUpgradable() bool {
	if m != nil {
		return m.Upgradable
	}
	return false
}

func (m *ESDTData) GetCanChangeOwner() bool {
	if m != nil {
		return m.CanChangeOwner
	}
	return false
}

func (m *ESDTData) GetIsPaused() bool {
	if m != nil {
		return m.IsPaused
	}
	return false
}

func (m *ESDTData) GetMintedValue() *math_big.Int {
	if m != nil {
		return m.MintedValue
	}
	return nil
}

func (m *ESDTData) GetBurntValue() *math_big.Int {
	if m != nil {
		return m.BurntValue
	}
	return nil
}

func (m *ESDTData) GetNumDecimals() uint32 {
	if m != nil {
		return m.NumDecimals
	}
	return 0
}

type ESDTConfig struct {
	OwnerAddress       []byte        `protobuf:"bytes,1,opt,name=OwnerAddress,proto3" json:"OwnerAddress"`
	BaseIssuingCost    *math_big.Int `protobuf:"bytes,2,opt,name=BaseIssuingCost,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"BaseIssuingCost"`
	MinTokenNameLength uint32        `protobuf:"varint,3,opt,name=MinTokenNameLength,proto3" json:"MinTokenNameLength"`
	MaxTokenNameLength uint32        `protobuf:"varint,4,opt,name=MaxTokenNameLength,proto3" json:"MaxTokenNameLength"`
}

func (m *ESDTConfig) Reset()      { *m = ESDTConfig{} }
func (*ESDTConfig) ProtoMessage() {}
func (*ESDTConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e413e402abc6a34c, []int{1}
}
func (m *ESDTConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ESDTConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ESDTConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ESDTConfig.Merge(m, src)
}
func (m *ESDTConfig) XXX_Size() int {
	return m.Size()
}
func (m *ESDTConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ESDTConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ESDTConfig proto.InternalMessageInfo

func (m *ESDTConfig) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *ESDTConfig) GetBaseIssuingCost() *math_big.Int {
	if m != nil {
		return m.BaseIssuingCost
	}
	return nil
}

func (m *ESDTConfig) GetMinTokenNameLength() uint32 {
	if m != nil {
		return m.MinTokenNameLength
	}
	return 0
}

func (m *ESDTConfig) GetMaxTokenNameLength() uint32 {
	if m != nil {
		return m.MaxTokenNameLength
	}
	return 0
}

func init() {
	proto.RegisterType((*ESDTData)(nil), "proto.ESDTData")
	proto.RegisterType((*ESDTConfig)(nil), "proto.ESDTConfig")
}

func init() { proto.RegisterFile("esdt.proto", fileDescriptor_e413e402abc6a34c) }

var fileDescriptor_e413e402abc6a34c = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xb3, 0xfd, 0xf5, 0x4f, 0xba, 0x49, 0xda, 0x9f, 0x56, 0x08, 0x59, 0x1c, 0xd6, 0x55,
	0x25, 0xa4, 0x48, 0xa8, 0x89, 0x10, 0x9c, 0xe0, 0x54, 0xbb, 0xad, 0x14, 0x89, 0x06, 0xb4, 0x0d,
	0x7f, 0xc4, 0x6d, 0x13, 0x6f, 0x1d, 0xab, 0xf1, 0xba, 0xda, 0x5d, 0x53, 0xe0, 0x84, 0x78, 0x02,
	0x5e, 0x02, 0x09, 0xf1, 0x24, 0x1c, 0x7b, 0xa3, 0x27, 0x43, 0xdd, 0x0b, 0xf2, 0xa9, 0x8f, 0x80,
	0x3c, 0x25, 0xb6, 0x1b, 0x72, 0x42, 0x3d, 0xe5, 0x3b, 0x9f, 0xf9, 0xee, 0x6c, 0x66, 0x36, 0x13,
	0x8c, 0x85, 0xf6, 0x4c, 0xe7, 0x58, 0x45, 0x26, 0x22, 0x4b, 0xf0, 0x71, 0x67, 0xcb, 0x0f, 0xcc,
	0x38, 0x1e, 0x76, 0x46, 0x51, 0xd8, 0xf5, 0x23, 0x3f, 0xea, 0x02, 0x1e, 0xc6, 0x87, 0x10, 0x41,
	0x00, 0xea, 0xea, 0xd4, 0xe6, 0xe7, 0x65, 0x5c, 0xdf, 0x3d, 0xd8, 0x19, 0xec, 0x70, 0xc3, 0xc9,
	0x43, 0xdc, 0x7c, 0x7a, 0x22, 0x85, 0xda, 0xf6, 0x3c, 0x25, 0xb4, 0xb6, 0xd0, 0x06, 0x6a, 0x37,
	0x9d, 0xff, 0xb3, 0xc4, 0xbe, 0xc6, 0xd9, 0xb5, 0x88, 0xdc, 0xc3, 0xab, 0x83, 0xe8, 0x48, 0xc8,
	0x3e, 0x0f, 0x85, 0xb5, 0x00, 0x47, 0x5a, 0x59, 0x62, 0x97, 0x90, 0x95, 0x92, 0x74, 0x30, 0x1e,
	0x04, 0xa3, 0x23, 0xa1, 0xc0, 0xfd, 0x1f, 0xb8, 0xd7, 0xb2, 0xc4, 0xae, 0x50, 0x56, 0xd1, 0xa4,
	0x8d, 0xeb, 0xfb, 0x81, 0x34, 0x7c, 0x38, 0x11, 0xd6, 0xe2, 0x06, 0x6a, 0xd7, 0x9d, 0x66, 0x96,
	0xd8, 0x05, 0x63, 0x85, 0xca, 0x9d, 0x4e, 0xac, 0x24, 0x38, 0x97, 0x4a, 0xe7, 0x94, 0xb1, 0x42,
	0xe5, 0x4e, 0x97, 0xcb, 0x67, 0x3c, 0xd6, 0xc2, 0x5a, 0x2e, 0x9d, 0x53, 0xc6, 0x0a, 0x95, 0xb7,
	0xe6, 0x72, 0xb9, 0xa7, 0x84, 0x78, 0x2f, 0xac, 0x15, 0xb0, 0x42, 0x6b, 0x05, 0x64, 0xa5, 0x24,
	0x77, 0xf1, 0x8a, 0xcb, 0xe5, 0xcb, 0xe0, 0x58, 0x58, 0x75, 0xb0, 0x36, 0xb2, 0xc4, 0x9e, 0x22,
	0x36, 0x15, 0xf9, 0x04, 0x9e, 0x1f, 0xfb, 0x8a, 0x7b, 0xf0, 0x4d, 0x57, 0xc1, 0x09, 0x13, 0x70,
	0xb9, 0xbc, 0x4a, 0x08, 0x56, 0x71, 0x90, 0x47, 0x78, 0xcd, 0xe5, 0xd2, 0x1d, 0x73, 0xe9, 0x0b,
	0x98, 0xbb, 0x85, 0xe1, 0x0c, 0xc9, 0x12, 0x7b, 0x26, 0xc3, 0x66, 0xe2, 0xbc, 0xd3, 0x9e, 0x86,
	0x56, 0x3c, 0xab, 0x51, 0x76, 0x3a, 0x65, 0xac, 0x50, 0xe4, 0x0d, 0x6e, 0xe4, 0x93, 0x14, 0xde,
	0x0b, 0x3e, 0x89, 0x85, 0xd5, 0x84, 0x87, 0x19, 0x64, 0x89, 0x5d, 0xc5, 0x5f, 0x7f, 0xd8, 0xdb,
	0x21, 0x37, 0xe3, 0xee, 0x30, 0xf0, 0x3b, 0x3d, 0x69, 0x1e, 0x57, 0x7e, 0x6b, 0xbb, 0x13, 0x15,
	0x49, 0xaf, 0x2f, 0xcc, 0x49, 0xa4, 0x8e, 0xba, 0x02, 0xa2, 0x2d, 0x3f, 0xea, 0x7a, 0xdc, 0xf0,
	0x8e, 0x13, 0xf8, 0x3d, 0x69, 0x5c, 0xae, 0x8d, 0x50, 0xac, 0x5a, 0x91, 0x68, 0x8c, 0xf3, 0x77,
	0x31, 0x57, 0xd7, 0xb6, 0xe0, 0xda, 0x83, 0x7c, 0x1a, 0x25, 0xbd, 0x99, 0x5b, 0x2b, 0x05, 0xc9,
	0x7d, 0xdc, 0xe8, 0xc7, 0xe1, 0x8e, 0x18, 0x05, 0x21, 0x9f, 0x68, 0x6b, 0x6d, 0x03, 0xb5, 0x5b,
	0xce, 0x7a, 0xde, 0x6c, 0x05, 0xb3, 0x6a, 0xb0, 0xf9, 0x7d, 0x01, 0xe3, 0x7c, 0x4f, 0xdc, 0x48,
	0x1e, 0x06, 0xfe, 0x3f, 0x6e, 0xca, 0x47, 0x84, 0xd7, 0x1d, 0xae, 0x45, 0x4f, 0xeb, 0x38, 0x90,
	0xbe, 0x1b, 0x69, 0xf3, 0x67, 0x61, 0x5e, 0x65, 0x89, 0x3d, 0x9b, 0xba, 0x99, 0xbe, 0x67, 0xab,
	0x92, 0x3d, 0x4c, 0xf6, 0x03, 0x59, 0x6c, 0xe4, 0x13, 0x21, 0x7d, 0x33, 0x86, 0x4d, 0x6c, 0x39,
	0xb7, 0xb3, 0xc4, 0x9e, 0x93, 0x65, 0x73, 0x18, 0xd4, 0xe1, 0x6f, 0x67, 0xeb, 0x2c, 0x56, 0xea,
	0xfc, 0x95, 0x65, 0x73, 0x98, 0xd3, 0x3f, 0x3d, 0xa7, 0xb5, 0xb3, 0x73, 0x5a, 0xbb, 0x3c, 0xa7,
	0xe8, 0x43, 0x4a, 0xd1, 0x97, 0x94, 0xa2, 0x6f, 0x29, 0x45, 0xa7, 0x29, 0x45, 0x67, 0x29, 0x45,
	0x3f, 0x53, 0x8a, 0x7e, 0xa5, 0xb4, 0x76, 0x99, 0x52, 0xf4, 0xe9, 0x82, 0xd6, 0x4e, 0x2f, 0x68,
	0xed, 0xec, 0x82, 0xd6, 0x5e, 0xdf, 0xd2, 0xef, 0xb4, 0x11, 0xe1, 0x41, 0xc8, 0x95, 0x71, 0x23,
	0x69, 0x14, 0x1f, 0x19, 0x3d, 0x5c, 0x86, 0x3f, 0xb6, 0x07, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x1d, 0x26, 0x99, 0x78, 0x1c, 0x05, 0x00, 0x00,
}

func (this *ESDTData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ESDTData)
	if !ok {
		that2, ok := that.(ESDTData)
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
	if !bytes.Equal(this.OwnerAddress, that1.OwnerAddress) {
		return false
	}
	if !bytes.Equal(this.TokenName, that1.TokenName) {
		return false
	}
	if !bytes.Equal(this.TickerName, that1.TickerName) {
		return false
	}
	if this.Mintable != that1.Mintable {
		return false
	}
	if this.Burnable != that1.Burnable {
		return false
	}
	if this.CanPause != that1.CanPause {
		return false
	}
	if this.CanFreeze != that1.CanFreeze {
		return false
	}
	if this.CanWipe != that1.CanWipe {
		return false
	}
	if this.Upgradable != that1.Upgradable {
		return false
	}
	if this.CanChangeOwner != that1.CanChangeOwner {
		return false
	}
	if this.IsPaused != that1.IsPaused {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.MintedValue, that1.MintedValue) {
			return false
		}
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.BurntValue, that1.BurntValue) {
			return false
		}
	}
	if this.NumDecimals != that1.NumDecimals {
		return false
	}
	return true
}
func (this *ESDTConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ESDTConfig)
	if !ok {
		that2, ok := that.(ESDTConfig)
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
	if !bytes.Equal(this.OwnerAddress, that1.OwnerAddress) {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.BaseIssuingCost, that1.BaseIssuingCost) {
			return false
		}
	}
	if this.MinTokenNameLength != that1.MinTokenNameLength {
		return false
	}
	if this.MaxTokenNameLength != that1.MaxTokenNameLength {
		return false
	}
	return true
}
func (this *ESDTData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 18)
	s = append(s, "&systemSmartContracts.ESDTData{")
	s = append(s, "OwnerAddress: "+fmt.Sprintf("%#v", this.OwnerAddress)+",\n")
	s = append(s, "TokenName: "+fmt.Sprintf("%#v", this.TokenName)+",\n")
	s = append(s, "TickerName: "+fmt.Sprintf("%#v", this.TickerName)+",\n")
	s = append(s, "Mintable: "+fmt.Sprintf("%#v", this.Mintable)+",\n")
	s = append(s, "Burnable: "+fmt.Sprintf("%#v", this.Burnable)+",\n")
	s = append(s, "CanPause: "+fmt.Sprintf("%#v", this.CanPause)+",\n")
	s = append(s, "CanFreeze: "+fmt.Sprintf("%#v", this.CanFreeze)+",\n")
	s = append(s, "CanWipe: "+fmt.Sprintf("%#v", this.CanWipe)+",\n")
	s = append(s, "Upgradable: "+fmt.Sprintf("%#v", this.Upgradable)+",\n")
	s = append(s, "CanChangeOwner: "+fmt.Sprintf("%#v", this.CanChangeOwner)+",\n")
	s = append(s, "IsPaused: "+fmt.Sprintf("%#v", this.IsPaused)+",\n")
	s = append(s, "MintedValue: "+fmt.Sprintf("%#v", this.MintedValue)+",\n")
	s = append(s, "BurntValue: "+fmt.Sprintf("%#v", this.BurntValue)+",\n")
	s = append(s, "NumDecimals: "+fmt.Sprintf("%#v", this.NumDecimals)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ESDTConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&systemSmartContracts.ESDTConfig{")
	s = append(s, "OwnerAddress: "+fmt.Sprintf("%#v", this.OwnerAddress)+",\n")
	s = append(s, "BaseIssuingCost: "+fmt.Sprintf("%#v", this.BaseIssuingCost)+",\n")
	s = append(s, "MinTokenNameLength: "+fmt.Sprintf("%#v", this.MinTokenNameLength)+",\n")
	s = append(s, "MaxTokenNameLength: "+fmt.Sprintf("%#v", this.MaxTokenNameLength)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringEsdt(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ESDTData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ESDTData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ESDTData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumDecimals != 0 {
		i = encodeVarintEsdt(dAtA, i, uint64(m.NumDecimals))
		i--
		dAtA[i] = 0x70
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.BurntValue)
		i -= size
		if _, err := __caster.MarshalTo(m.BurntValue, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEsdt(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.MintedValue)
		i -= size
		if _, err := __caster.MarshalTo(m.MintedValue, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEsdt(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if m.IsPaused {
		i--
		if m.IsPaused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if m.CanChangeOwner {
		i--
		if m.CanChangeOwner {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.Upgradable {
		i--
		if m.Upgradable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.CanWipe {
		i--
		if m.CanWipe {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.CanFreeze {
		i--
		if m.CanFreeze {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.CanPause {
		i--
		if m.CanPause {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Burnable {
		i--
		if m.Burnable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Mintable {
		i--
		if m.Mintable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.TickerName) > 0 {
		i -= len(m.TickerName)
		copy(dAtA[i:], m.TickerName)
		i = encodeVarintEsdt(dAtA, i, uint64(len(m.TickerName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenName) > 0 {
		i -= len(m.TokenName)
		copy(dAtA[i:], m.TokenName)
		i = encodeVarintEsdt(dAtA, i, uint64(len(m.TokenName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintEsdt(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ESDTConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ESDTConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ESDTConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxTokenNameLength != 0 {
		i = encodeVarintEsdt(dAtA, i, uint64(m.MaxTokenNameLength))
		i--
		dAtA[i] = 0x20
	}
	if m.MinTokenNameLength != 0 {
		i = encodeVarintEsdt(dAtA, i, uint64(m.MinTokenNameLength))
		i--
		dAtA[i] = 0x18
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.BaseIssuingCost)
		i -= size
		if _, err := __caster.MarshalTo(m.BaseIssuingCost, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEsdt(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintEsdt(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEsdt(dAtA []byte, offset int, v uint64) int {
	offset -= sovEsdt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ESDTData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovEsdt(uint64(l))
	}
	l = len(m.TokenName)
	if l > 0 {
		n += 1 + l + sovEsdt(uint64(l))
	}
	l = len(m.TickerName)
	if l > 0 {
		n += 1 + l + sovEsdt(uint64(l))
	}
	if m.Mintable {
		n += 2
	}
	if m.Burnable {
		n += 2
	}
	if m.CanPause {
		n += 2
	}
	if m.CanFreeze {
		n += 2
	}
	if m.CanWipe {
		n += 2
	}
	if m.Upgradable {
		n += 2
	}
	if m.CanChangeOwner {
		n += 2
	}
	if m.IsPaused {
		n += 2
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.MintedValue)
		n += 1 + l + sovEsdt(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.BurntValue)
		n += 1 + l + sovEsdt(uint64(l))
	}
	if m.NumDecimals != 0 {
		n += 1 + sovEsdt(uint64(m.NumDecimals))
	}
	return n
}

func (m *ESDTConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovEsdt(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.BaseIssuingCost)
		n += 1 + l + sovEsdt(uint64(l))
	}
	if m.MinTokenNameLength != 0 {
		n += 1 + sovEsdt(uint64(m.MinTokenNameLength))
	}
	if m.MaxTokenNameLength != 0 {
		n += 1 + sovEsdt(uint64(m.MaxTokenNameLength))
	}
	return n
}

func sovEsdt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEsdt(x uint64) (n int) {
	return sovEsdt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ESDTData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ESDTData{`,
		`OwnerAddress:` + fmt.Sprintf("%v", this.OwnerAddress) + `,`,
		`TokenName:` + fmt.Sprintf("%v", this.TokenName) + `,`,
		`TickerName:` + fmt.Sprintf("%v", this.TickerName) + `,`,
		`Mintable:` + fmt.Sprintf("%v", this.Mintable) + `,`,
		`Burnable:` + fmt.Sprintf("%v", this.Burnable) + `,`,
		`CanPause:` + fmt.Sprintf("%v", this.CanPause) + `,`,
		`CanFreeze:` + fmt.Sprintf("%v", this.CanFreeze) + `,`,
		`CanWipe:` + fmt.Sprintf("%v", this.CanWipe) + `,`,
		`Upgradable:` + fmt.Sprintf("%v", this.Upgradable) + `,`,
		`CanChangeOwner:` + fmt.Sprintf("%v", this.CanChangeOwner) + `,`,
		`IsPaused:` + fmt.Sprintf("%v", this.IsPaused) + `,`,
		`MintedValue:` + fmt.Sprintf("%v", this.MintedValue) + `,`,
		`BurntValue:` + fmt.Sprintf("%v", this.BurntValue) + `,`,
		`NumDecimals:` + fmt.Sprintf("%v", this.NumDecimals) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ESDTConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ESDTConfig{`,
		`OwnerAddress:` + fmt.Sprintf("%v", this.OwnerAddress) + `,`,
		`BaseIssuingCost:` + fmt.Sprintf("%v", this.BaseIssuingCost) + `,`,
		`MinTokenNameLength:` + fmt.Sprintf("%v", this.MinTokenNameLength) + `,`,
		`MaxTokenNameLength:` + fmt.Sprintf("%v", this.MaxTokenNameLength) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEsdt(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ESDTData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEsdt
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
			return fmt.Errorf("proto: ESDTData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ESDTData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
				return ErrInvalidLengthEsdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEsdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = append(m.OwnerAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.OwnerAddress == nil {
				m.OwnerAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenName", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
				return ErrInvalidLengthEsdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEsdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenName = append(m.TokenName[:0], dAtA[iNdEx:postIndex]...)
			if m.TokenName == nil {
				m.TokenName = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickerName", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
				return ErrInvalidLengthEsdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEsdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TickerName = append(m.TickerName[:0], dAtA[iNdEx:postIndex]...)
			if m.TickerName == nil {
				m.TickerName = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mintable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
			m.Mintable = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Burnable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
			m.Burnable = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanPause", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
			m.CanPause = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanFreeze", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
			m.CanFreeze = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanWipe", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
			m.CanWipe = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Upgradable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
			m.Upgradable = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanChangeOwner", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
			m.CanChangeOwner = bool(v != 0)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPaused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
			m.IsPaused = bool(v != 0)
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintedValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
				return ErrInvalidLengthEsdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEsdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.MintedValue = tmp
				}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurntValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
				return ErrInvalidLengthEsdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEsdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.BurntValue = tmp
				}
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumDecimals", wireType)
			}
			m.NumDecimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumDecimals |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEsdt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEsdt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEsdt
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
func (m *ESDTConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEsdt
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
			return fmt.Errorf("proto: ESDTConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ESDTConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
				return ErrInvalidLengthEsdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEsdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = append(m.OwnerAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.OwnerAddress == nil {
				m.OwnerAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseIssuingCost", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
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
				return ErrInvalidLengthEsdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEsdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.BaseIssuingCost = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTokenNameLength", wireType)
			}
			m.MinTokenNameLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinTokenNameLength |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTokenNameLength", wireType)
			}
			m.MaxTokenNameLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTokenNameLength |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEsdt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEsdt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEsdt
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
func skipEsdt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEsdt
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
					return 0, ErrIntOverflowEsdt
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
					return 0, ErrIntOverflowEsdt
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
				return 0, ErrInvalidLengthEsdt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEsdt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEsdt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEsdt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEsdt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEsdt = fmt.Errorf("proto: unexpected end of group")
)
