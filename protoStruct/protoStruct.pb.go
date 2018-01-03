// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoStruct.proto

/*
Package protoStruct is a generated protocol buffer package.

It is generated from these files:
	protoStruct.proto

It has these top-level messages:
	Agentstate
	Request
	Setup
*/
package protoStruct

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Regular message, that is sent on heartbeat
type Agentstate struct {
	Token        string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	AgentID      string `protobuf:"bytes,2,opt,name=agentID" json:"agentID,omitempty"`
	UserID       string `protobuf:"bytes,3,opt,name=userID" json:"userID,omitempty"`
	ProductID    string `protobuf:"bytes,4,opt,name=productID" json:"productID,omitempty"`
	Weight       int32  `protobuf:"varint,5,opt,name=weight" json:"weight,omitempty"`
	StateExpires string `protobuf:"bytes,6,opt,name=stateExpires" json:"stateExpires,omitempty"`
}

func (m *Agentstate) Reset()                    { *m = Agentstate{} }
func (m *Agentstate) String() string            { return proto.CompactTextString(m) }
func (*Agentstate) ProtoMessage()               {}
func (*Agentstate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Agentstate) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Agentstate) GetAgentID() string {
	if m != nil {
		return m.AgentID
	}
	return ""
}

func (m *Agentstate) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Agentstate) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

func (m *Agentstate) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Agentstate) GetStateExpires() string {
	if m != nil {
		return m.StateExpires
	}
	return ""
}

// Request
type Request struct {
	AgentID string `protobuf:"bytes,1,opt,name=agentID" json:"agentID,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetAgentID() string {
	if m != nil {
		return m.AgentID
	}
	return ""
}

// Setup message is sent by server to configure agent
type Setup struct {
	Token        string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	ProductID    string `protobuf:"bytes,2,opt,name=productID" json:"productID,omitempty"`
	UserID       string `protobuf:"bytes,3,opt,name=userID" json:"userID,omitempty"`
	Heartbeat    int32  `protobuf:"varint,4,opt,name=heartbeat" json:"heartbeat,omitempty"`
	StateExpires string `protobuf:"bytes,5,opt,name=stateExpires" json:"stateExpires,omitempty"`
}

func (m *Setup) Reset()                    { *m = Setup{} }
func (m *Setup) String() string            { return proto.CompactTextString(m) }
func (*Setup) ProtoMessage()               {}
func (*Setup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Setup) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Setup) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

func (m *Setup) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Setup) GetHeartbeat() int32 {
	if m != nil {
		return m.Heartbeat
	}
	return 0
}

func (m *Setup) GetStateExpires() string {
	if m != nil {
		return m.StateExpires
	}
	return ""
}

func init() {
	proto.RegisterType((*Agentstate)(nil), "protoStruct.agentstate")
	proto.RegisterType((*Request)(nil), "protoStruct.request")
	proto.RegisterType((*Setup)(nil), "protoStruct.setup")
}

func init() { proto.RegisterFile("protoStruct.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4e, 0x85, 0x30,
	0x10, 0x86, 0xd3, 0xa7, 0xe5, 0xe5, 0x8d, 0x6e, 0x6c, 0x8c, 0xe9, 0xc2, 0xc5, 0x4b, 0xdd, 0xb0,
	0x72, 0xe3, 0x15, 0x70, 0xc1, 0x16, 0x4f, 0x50, 0x70, 0x02, 0xc4, 0x84, 0xd6, 0x76, 0x1a, 0xbd,
	0x88, 0xd7, 0xf0, 0x8c, 0x86, 0x11, 0x03, 0x04, 0xd9, 0xf1, 0x7f, 0xc3, 0x9f, 0x7c, 0x7f, 0xe1,
	0xc6, 0x07, 0x47, 0xee, 0x85, 0x42, 0x6a, 0xe8, 0x91, 0xbf, 0xd5, 0xd5, 0x02, 0x99, 0x6f, 0x01,
	0x60, 0x5b, 0x1c, 0x28, 0x92, 0x25, 0x54, 0xb7, 0x20, 0xc9, 0xbd, 0xe1, 0xa0, 0xc5, 0x59, 0xe4,
	0xa7, 0xea, 0x37, 0x28, 0x0d, 0x47, 0xfe, 0xa7, 0x2c, 0xf4, 0x81, 0xf9, 0x5f, 0x54, 0x77, 0x90,
	0xa5, 0x88, 0xa1, 0x2c, 0xf4, 0x05, 0x1f, 0xa6, 0xa4, 0xee, 0xe1, 0xe4, 0x83, 0x7b, 0x4d, 0xcd,
	0xd8, 0xb9, 0xe4, 0xd3, 0x0c, 0xc6, 0xd6, 0x07, 0xf6, 0x6d, 0x47, 0x5a, 0x9e, 0x45, 0x2e, 0xab,
	0x29, 0x29, 0x03, 0xd7, 0xac, 0xf1, 0xfc, 0xe9, 0xfb, 0x80, 0x51, 0x67, 0x5c, 0x5c, 0x31, 0xf3,
	0x00, 0xc7, 0x80, 0xef, 0x09, 0x23, 0x2d, 0xb5, 0xc4, 0x4a, 0xcb, 0x7c, 0x09, 0x90, 0x11, 0x29,
	0xf9, 0x9d, 0x41, 0x2b, 0xbd, 0xc3, 0x3f, 0x7a, 0x7b, 0xa3, 0x3a, 0xb4, 0x81, 0x6a, 0xb4, 0xc4,
	0xa3, 0x64, 0x35, 0x83, 0x8d, 0xbc, 0xdc, 0xca, 0xd7, 0x19, 0x3f, 0xfd, 0xd3, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0xf8, 0xb3, 0x79, 0x96, 0x01, 0x00, 0x00,
}
