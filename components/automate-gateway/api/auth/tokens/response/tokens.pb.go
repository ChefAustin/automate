// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/auth/tokens/response/tokens.proto

package response

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Token struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Active               bool     `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	Created              string   `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated              string   `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_98ce77d66ffe9f1d, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Token) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Token) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Token) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Token) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

type Tokens struct {
	Tokens               []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tokens) Reset()         { *m = Tokens{} }
func (m *Tokens) String() string { return proto.CompactTextString(m) }
func (*Tokens) ProtoMessage()    {}
func (*Tokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_98ce77d66ffe9f1d, []int{1}
}

func (m *Tokens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tokens.Unmarshal(m, b)
}
func (m *Tokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tokens.Marshal(b, m, deterministic)
}
func (m *Tokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tokens.Merge(m, src)
}
func (m *Tokens) XXX_Size() int {
	return xxx_messageInfo_Tokens.Size(m)
}
func (m *Tokens) XXX_DiscardUnknown() {
	xxx_messageInfo_Tokens.DiscardUnknown(m)
}

var xxx_messageInfo_Tokens proto.InternalMessageInfo

func (m *Tokens) GetTokens() []*Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type DeleteTokenResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTokenResp) Reset()         { *m = DeleteTokenResp{} }
func (m *DeleteTokenResp) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenResp) ProtoMessage()    {}
func (*DeleteTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_98ce77d66ffe9f1d, []int{2}
}

func (m *DeleteTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenResp.Unmarshal(m, b)
}
func (m *DeleteTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenResp.Marshal(b, m, deterministic)
}
func (m *DeleteTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenResp.Merge(m, src)
}
func (m *DeleteTokenResp) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenResp.Size(m)
}
func (m *DeleteTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Token)(nil), "chef.automate.api.tokens.response.Token")
	proto.RegisterType((*Tokens)(nil), "chef.automate.api.tokens.response.Tokens")
	proto.RegisterType((*DeleteTokenResp)(nil), "chef.automate.api.tokens.response.DeleteTokenResp")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/auth/tokens/response/tokens.proto", fileDescriptor_98ce77d66ffe9f1d)
}

var fileDescriptor_98ce77d66ffe9f1d = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x69, 0xd7, 0x56, 0x9d, 0x05, 0xc5, 0x20, 0x92, 0x63, 0xed, 0xa9, 0x17, 0x13, 0xd0,
	0x17, 0x50, 0xf1, 0x24, 0x88, 0x50, 0x3c, 0x79, 0xcb, 0xa6, 0xe3, 0x36, 0xb8, 0x6d, 0x42, 0x33,
	0x5d, 0xf1, 0x59, 0x7c, 0x59, 0x69, 0x9a, 0x8a, 0x37, 0xc1, 0xe3, 0xf7, 0x7f, 0x33, 0x61, 0xf2,
	0xc3, 0x9d, 0xb6, 0x9d, 0xb3, 0x3d, 0xf6, 0xe4, 0xa5, 0x1a, 0xc9, 0x76, 0x8a, 0xf0, 0x6a, 0xab,
	0x08, 0x3f, 0xd4, 0xa7, 0x54, 0xce, 0x4c, 0x61, 0x2b, 0xc9, 0xbe, 0x63, 0xef, 0xe5, 0x80, 0xde,
	0xd9, 0xde, 0x63, 0x64, 0xe1, 0x06, 0x4b, 0x96, 0x5d, 0xea, 0x16, 0xdf, 0xc4, 0xb2, 0x2c, 0x94,
	0x33, 0x22, 0xfa, 0x65, 0xbe, 0xfc, 0x4a, 0x20, 0x7b, 0x99, 0x32, 0x76, 0x02, 0xa9, 0x69, 0x78,
	0x52, 0x24, 0xd5, 0x71, 0x9d, 0x9a, 0x86, 0x15, 0xb0, 0x6e, 0xd0, 0xeb, 0xc1, 0x38, 0x32, 0xb6,
	0xe7, 0x69, 0x10, 0xbf, 0x23, 0x76, 0x0e, 0xd9, 0x5e, 0xed, 0x46, 0xe4, 0xab, 0xe0, 0x66, 0x60,
	0x17, 0x90, 0x2b, 0x4d, 0x66, 0x8f, 0xfc, 0xa0, 0x48, 0xaa, 0xa3, 0x3a, 0x12, 0xe3, 0x70, 0xa8,
	0x07, 0x54, 0x84, 0x0d, 0xcf, 0xc2, 0xfc, 0x82, 0x93, 0x19, 0x5d, 0x13, 0x4c, 0x3e, 0x9b, 0x88,
	0xe5, 0x23, 0xe4, 0xe1, 0x38, 0xcf, 0x6e, 0x21, 0x9f, 0x4f, 0xe7, 0x49, 0xb1, 0xaa, 0xd6, 0xd7,
	0x95, 0xf8, 0xf3, 0x6f, 0x22, 0xac, 0xd6, 0x71, 0xaf, 0x3c, 0x83, 0xd3, 0x07, 0xdc, 0x21, 0xe1,
	0x1c, 0xa3, 0x77, 0xf7, 0xcf, 0xaf, 0x4f, 0x5b, 0x43, 0xed, 0xb8, 0x11, 0xda, 0x76, 0x72, 0x7a,
	0xf0, 0xa7, 0x69, 0xf9, 0x9f, 0xf6, 0x37, 0x79, 0xe8, 0xfd, 0xe6, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x26, 0xcd, 0xf1, 0xfe, 0xbc, 0x01, 0x00, 0x00,
}