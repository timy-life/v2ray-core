// Code generated by protoc-gen-go.
// source: v2ray.com/core/proxy/socks/config.proto
// DO NOT EDIT!

/*
Package socks is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/proxy/socks/config.proto

It has these top-level messages:
	Account
	ServerConfig
	ClientConfig
*/
package socks

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import com_v2ray_core_common_net "v2ray.com/core/common/net"
import com_v2ray_core_common_protocol1 "v2ray.com/core/common/protocol"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthType int32

const (
	AuthType_NO_AUTH  AuthType = 0
	AuthType_PASSWORD AuthType = 1
)

var AuthType_name = map[int32]string{
	0: "NO_AUTH",
	1: "PASSWORD",
}
var AuthType_value = map[string]int32{
	"NO_AUTH":  0,
	"PASSWORD": 1,
}

func (x AuthType) String() string {
	return proto.EnumName(AuthType_name, int32(x))
}
func (AuthType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Account struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServerConfig struct {
	AuthType   AuthType                             `protobuf:"varint,1,opt,name=auth_type,json=authType,enum=com.v2ray.core.proxy.socks.AuthType" json:"auth_type,omitempty"`
	Accounts   map[string]string                    `protobuf:"bytes,2,rep,name=accounts" json:"accounts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Address    *com_v2ray_core_common_net.AddressPB `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	UdpEnabled bool                                 `protobuf:"varint,4,opt,name=udp_enabled,json=udpEnabled" json:"udp_enabled,omitempty"`
	Timeout    uint32                               `protobuf:"varint,5,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServerConfig) GetAccounts() map[string]string {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *ServerConfig) GetAddress() *com_v2ray_core_common_net.AddressPB {
	if m != nil {
		return m.Address
	}
	return nil
}

type ClientConfig struct {
	Server []*com_v2ray_core_common_protocol1.ServerSpecPB `protobuf:"bytes,1,rep,name=server" json:"server,omitempty"`
}

func (m *ClientConfig) Reset()                    { *m = ClientConfig{} }
func (m *ClientConfig) String() string            { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()               {}
func (*ClientConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientConfig) GetServer() []*com_v2ray_core_common_protocol1.ServerSpecPB {
	if m != nil {
		return m.Server
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "com.v2ray.core.proxy.socks.Account")
	proto.RegisterType((*ServerConfig)(nil), "com.v2ray.core.proxy.socks.ServerConfig")
	proto.RegisterType((*ClientConfig)(nil), "com.v2ray.core.proxy.socks.ClientConfig")
	proto.RegisterEnum("com.v2ray.core.proxy.socks.AuthType", AuthType_name, AuthType_value)
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/socks/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0x5f, 0x8b, 0xd3, 0x40,
	0x14, 0xc5, 0x4d, 0x6b, 0x37, 0xd9, 0xdb, 0xae, 0x94, 0xc1, 0x87, 0x90, 0x17, 0x43, 0x51, 0x0c,
	0x22, 0x13, 0xa9, 0x20, 0xa2, 0x20, 0xa4, 0xbb, 0x0b, 0x3e, 0xb9, 0x65, 0x5a, 0x11, 0x7c, 0x29,
	0xb3, 0x93, 0xab, 0x5b, 0x36, 0x99, 0x19, 0x66, 0x26, 0xd5, 0x7c, 0x58, 0xbf, 0x8b, 0x6c, 0x26,
	0x59, 0xfc, 0xb7, 0xbe, 0xcd, 0xcd, 0x3d, 0xf7, 0xe4, 0x9e, 0xdf, 0x85, 0xa7, 0x87, 0xa5, 0xe1,
	0x2d, 0x15, 0xaa, 0xce, 0x85, 0x32, 0x98, 0x6b, 0xa3, 0xbe, 0xb7, 0xb9, 0x55, 0xe2, 0xda, 0xe6,
	0x42, 0xc9, 0x2f, 0xfb, 0xaf, 0x54, 0x1b, 0xe5, 0x14, 0x49, 0x84, 0xaa, 0xe9, 0x20, 0x36, 0x48,
	0x3b, 0x21, 0xed, 0x84, 0xc9, 0x9f, 0x26, 0x42, 0xd5, 0xb5, 0x92, 0xb9, 0x44, 0x97, 0xf3, 0xb2,
	0x34, 0x68, 0xad, 0x37, 0x49, 0x5e, 0xfc, 0x5b, 0xd8, 0x35, 0x85, 0xaa, 0x72, 0x8b, 0xe6, 0x80,
	0x66, 0x67, 0x35, 0x0a, 0x3f, 0xb1, 0x28, 0x20, 0x2c, 0x84, 0x50, 0x8d, 0x74, 0x24, 0x81, 0xa8,
	0xb1, 0x68, 0x24, 0xaf, 0x31, 0x0e, 0xd2, 0x20, 0x3b, 0x66, 0xb7, 0xf5, 0x4d, 0x4f, 0x73, 0x6b,
	0xbf, 0x29, 0x53, 0xc6, 0x23, 0xdf, 0x1b, 0xea, 0xc5, 0x8f, 0x11, 0xcc, 0x36, 0x9d, 0xf1, 0x69,
	0x17, 0x88, 0x14, 0x70, 0xcc, 0x1b, 0x77, 0xb5, 0x73, 0xad, 0xf6, 0x4e, 0x0f, 0x96, 0x8f, 0xe9,
	0xdd, 0xf1, 0x68, 0xd1, 0xb8, 0xab, 0x6d, 0xab, 0x91, 0x45, 0xbc, 0x7f, 0x11, 0x06, 0x11, 0xf7,
	0x6b, 0xd9, 0x78, 0x94, 0x8e, 0xb3, 0xe9, 0xf2, 0xd5, 0xff, 0x1c, 0x7e, 0xfd, 0x3d, 0xed, 0xf3,
	0xd8, 0x73, 0xe9, 0x4c, 0xcb, 0x6e, 0x7d, 0xc8, 0x3b, 0x08, 0x7b, 0x5a, 0xf1, 0x38, 0x0d, 0xb2,
	0xe9, 0xdf, 0x4b, 0x79, 0x5c, 0x54, 0xa2, 0xa3, 0x85, 0x57, 0xae, 0x57, 0x6c, 0x18, 0x22, 0x8f,
	0x60, 0xda, 0x94, 0x7a, 0x87, 0x92, 0x5f, 0x56, 0x58, 0xc6, 0xf7, 0xd3, 0x20, 0x8b, 0x18, 0x34,
	0xa5, 0x3e, 0xf7, 0x5f, 0x48, 0x0c, 0xa1, 0xdb, 0xd7, 0xa8, 0x1a, 0x17, 0x4f, 0xd2, 0x20, 0x3b,
	0x61, 0x43, 0x99, 0xbc, 0x85, 0x93, 0xdf, 0xb6, 0x22, 0x73, 0x18, 0x5f, 0x63, 0xdb, 0x63, 0xbe,
	0x79, 0x92, 0x87, 0x30, 0x39, 0xf0, 0xaa, 0xc1, 0x1e, 0xaf, 0x2f, 0xde, 0x8c, 0x5e, 0x07, 0x8b,
	0x2d, 0xcc, 0x4e, 0xab, 0x3d, 0x4a, 0xd7, 0xe3, 0x3d, 0x83, 0x23, 0x7f, 0xc7, 0x38, 0xe8, 0xc8,
	0x3c, 0xbf, 0x23, 0xc6, 0x70, 0xf5, 0x9e, 0xce, 0x46, 0xa3, 0x58, 0xaf, 0x58, 0x3f, 0xfb, 0xec,
	0x09, 0x44, 0x03, 0x77, 0x32, 0x85, 0xf0, 0xc3, 0xc5, 0xae, 0xf8, 0xb8, 0x7d, 0x3f, 0xbf, 0x47,
	0x66, 0x10, 0xad, 0x8b, 0xcd, 0xe6, 0xd3, 0x05, 0x3b, 0x9b, 0x07, 0xab, 0xf0, 0xf3, 0xa4, 0x43,
	0x7c, 0x79, 0xd4, 0xd9, 0xbd, 0xfc, 0x19, 0x00, 0x00, 0xff, 0xff, 0x13, 0x18, 0x73, 0xfe, 0xd1,
	0x02, 0x00, 0x00,
}
