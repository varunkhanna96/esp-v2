// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/envoy/http/path_matcher/config.proto

package path_matcher

import (
	common "cloudesf.googlesource.com/gcpproxy/src/go/proto/api/envoy/http/common"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PathMatcherRule struct {
	Pattern              *common.Pattern `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Operation            string          `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PathMatcherRule) Reset()         { *m = PathMatcherRule{} }
func (m *PathMatcherRule) String() string { return proto.CompactTextString(m) }
func (*PathMatcherRule) ProtoMessage()    {}
func (*PathMatcherRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_bebe96b55b7e4dec, []int{0}
}

func (m *PathMatcherRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathMatcherRule.Unmarshal(m, b)
}
func (m *PathMatcherRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathMatcherRule.Marshal(b, m, deterministic)
}
func (m *PathMatcherRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathMatcherRule.Merge(m, src)
}
func (m *PathMatcherRule) XXX_Size() int {
	return xxx_messageInfo_PathMatcherRule.Size(m)
}
func (m *PathMatcherRule) XXX_DiscardUnknown() {
	xxx_messageInfo_PathMatcherRule.DiscardUnknown(m)
}

var xxx_messageInfo_PathMatcherRule proto.InternalMessageInfo

func (m *PathMatcherRule) GetPattern() *common.Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *PathMatcherRule) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

type FilterConfig struct {
	Rules                []*PathMatcherRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_bebe96b55b7e4dec, []int{1}
}

func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetRules() []*PathMatcherRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*PathMatcherRule)(nil), "google.api.envoy.http.path_matcher.PathMatcherRule")
	proto.RegisterType((*FilterConfig)(nil), "google.api.envoy.http.path_matcher.FilterConfig")
}

func init() {
	proto.RegisterFile("api/envoy/http/path_matcher/config.proto", fileDescriptor_bebe96b55b7e4dec)
}

var fileDescriptor_bebe96b55b7e4dec = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x2c, 0xc8, 0xd4,
	0x4f, 0xcd, 0x2b, 0xcb, 0xaf, 0xd4, 0xcf, 0x28, 0x29, 0x29, 0xd0, 0x2f, 0x48, 0x2c, 0xc9, 0x88,
	0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4a, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4b, 0x2c, 0xc8,
	0xd4, 0x03, 0x6b, 0xd0, 0x03, 0x69, 0xd0, 0x43, 0xd6, 0x20, 0xa5, 0x8c, 0x66, 0x5a, 0x72, 0x7e,
	0x6e, 0x6e, 0x7e, 0x1e, 0xc8, 0xd0, 0x92, 0xd4, 0xa2, 0x3c, 0x88, 0x41, 0x52, 0xe2, 0x65, 0x89,
	0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44, 0x42, 0xa9, 0x9a, 0x8b, 0x3f, 0x20,
	0xb1, 0x24, 0xc3, 0x17, 0x62, 0x58, 0x50, 0x69, 0x4e, 0xaa, 0x90, 0x3d, 0x17, 0x3b, 0x54, 0xb3,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xaa, 0x1e, 0x76, 0x67, 0x40, 0x6c, 0xd2, 0x0b, 0x80,
	0x28, 0x0e, 0x82, 0xe9, 0x12, 0x52, 0xe7, 0xe2, 0xcc, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc,
	0xcf, 0x93, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x74, 0xe2, 0xdc, 0xf5, 0xf2, 0x00, 0x33, 0x4b, 0x11,
	0x93, 0x02, 0x63, 0x10, 0x42, 0x4e, 0x29, 0x92, 0x8b, 0xc7, 0x2d, 0x33, 0xa7, 0x24, 0xb5, 0xc8,
	0x19, 0xec, 0x69, 0x21, 0x4f, 0x2e, 0xd6, 0xa2, 0xd2, 0x9c, 0xd4, 0x62, 0x09, 0x46, 0x05, 0x66,
	0x0d, 0x6e, 0x23, 0x63, 0x3d, 0xc2, 0xde, 0xd7, 0x43, 0x73, 0x7d, 0x10, 0xc4, 0x84, 0x24, 0x36,
	0xb0, 0xf7, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x8f, 0x2d, 0x76, 0x6c, 0x01, 0x00,
	0x00,
}