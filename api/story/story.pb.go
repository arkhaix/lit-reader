// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/story/story.proto

package story

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/arkhaix/lit-reader/api/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CreateStoryRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateStoryRequest) Reset()         { *m = CreateStoryRequest{} }
func (m *CreateStoryRequest) String() string { return proto.CompactTextString(m) }
func (*CreateStoryRequest) ProtoMessage()    {}
func (*CreateStoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_story_b693fe2c963d8e94, []int{0}
}
func (m *CreateStoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStoryRequest.Unmarshal(m, b)
}
func (m *CreateStoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStoryRequest.Marshal(b, m, deterministic)
}
func (dst *CreateStoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStoryRequest.Merge(dst, src)
}
func (m *CreateStoryRequest) XXX_Size() int {
	return xxx_messageInfo_CreateStoryRequest.Size(m)
}
func (m *CreateStoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStoryRequest proto.InternalMessageInfo

func (m *CreateStoryRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type CreateStoryResponse struct {
	Status               *common.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 *Story         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateStoryResponse) Reset()         { *m = CreateStoryResponse{} }
func (m *CreateStoryResponse) String() string { return proto.CompactTextString(m) }
func (*CreateStoryResponse) ProtoMessage()    {}
func (*CreateStoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_story_b693fe2c963d8e94, []int{1}
}
func (m *CreateStoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStoryResponse.Unmarshal(m, b)
}
func (m *CreateStoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStoryResponse.Marshal(b, m, deterministic)
}
func (dst *CreateStoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStoryResponse.Merge(dst, src)
}
func (m *CreateStoryResponse) XXX_Size() int {
	return xxx_messageInfo_CreateStoryResponse.Size(m)
}
func (m *CreateStoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStoryResponse proto.InternalMessageInfo

func (m *CreateStoryResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateStoryResponse) GetData() *Story {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetStoryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStoryRequest) Reset()         { *m = GetStoryRequest{} }
func (m *GetStoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetStoryRequest) ProtoMessage()    {}
func (*GetStoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_story_b693fe2c963d8e94, []int{2}
}
func (m *GetStoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoryRequest.Unmarshal(m, b)
}
func (m *GetStoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoryRequest.Marshal(b, m, deterministic)
}
func (dst *GetStoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoryRequest.Merge(dst, src)
}
func (m *GetStoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetStoryRequest.Size(m)
}
func (m *GetStoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoryRequest proto.InternalMessageInfo

func (m *GetStoryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetStoryResponse struct {
	Status               *common.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 *Story         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetStoryResponse) Reset()         { *m = GetStoryResponse{} }
func (m *GetStoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetStoryResponse) ProtoMessage()    {}
func (*GetStoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_story_b693fe2c963d8e94, []int{3}
}
func (m *GetStoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoryResponse.Unmarshal(m, b)
}
func (m *GetStoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoryResponse.Marshal(b, m, deterministic)
}
func (dst *GetStoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoryResponse.Merge(dst, src)
}
func (m *GetStoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetStoryResponse.Size(m)
}
func (m *GetStoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoryResponse proto.InternalMessageInfo

func (m *GetStoryResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetStoryResponse) GetData() *Story {
	if m != nil {
		return m.Data
	}
	return nil
}

type Story struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Author               string   `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	NumChapters          int32    `protobuf:"varint,5,opt,name=numChapters,proto3" json:"numChapters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Story) Reset()         { *m = Story{} }
func (m *Story) String() string { return proto.CompactTextString(m) }
func (*Story) ProtoMessage()    {}
func (*Story) Descriptor() ([]byte, []int) {
	return fileDescriptor_story_b693fe2c963d8e94, []int{4}
}
func (m *Story) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Story.Unmarshal(m, b)
}
func (m *Story) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Story.Marshal(b, m, deterministic)
}
func (dst *Story) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Story.Merge(dst, src)
}
func (m *Story) XXX_Size() int {
	return xxx_messageInfo_Story.Size(m)
}
func (m *Story) XXX_DiscardUnknown() {
	xxx_messageInfo_Story.DiscardUnknown(m)
}

var xxx_messageInfo_Story proto.InternalMessageInfo

func (m *Story) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Story) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Story) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Story) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Story) GetNumChapters() int32 {
	if m != nil {
		return m.NumChapters
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateStoryRequest)(nil), "story.CreateStoryRequest")
	proto.RegisterType((*CreateStoryResponse)(nil), "story.CreateStoryResponse")
	proto.RegisterType((*GetStoryRequest)(nil), "story.GetStoryRequest")
	proto.RegisterType((*GetStoryResponse)(nil), "story.GetStoryResponse")
	proto.RegisterType((*Story)(nil), "story.Story")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoryServiceClient is the client API for StoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoryServiceClient interface {
	// CreateStory returns the story id and metadata for the queried url, scraping it first if necessary.
	CreateStory(ctx context.Context, in *CreateStoryRequest, opts ...grpc.CallOption) (*CreateStoryResponse, error)
	// GetStory returns the metadata for a previously-scraped story.
	GetStory(ctx context.Context, in *GetStoryRequest, opts ...grpc.CallOption) (*GetStoryResponse, error)
}

type storyServiceClient struct {
	cc *grpc.ClientConn
}

func NewStoryServiceClient(cc *grpc.ClientConn) StoryServiceClient {
	return &storyServiceClient{cc}
}

func (c *storyServiceClient) CreateStory(ctx context.Context, in *CreateStoryRequest, opts ...grpc.CallOption) (*CreateStoryResponse, error) {
	out := new(CreateStoryResponse)
	err := c.cc.Invoke(ctx, "/story.StoryService/CreateStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyServiceClient) GetStory(ctx context.Context, in *GetStoryRequest, opts ...grpc.CallOption) (*GetStoryResponse, error) {
	out := new(GetStoryResponse)
	err := c.cc.Invoke(ctx, "/story.StoryService/GetStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoryServiceServer is the server API for StoryService service.
type StoryServiceServer interface {
	// CreateStory returns the story id and metadata for the queried url, scraping it first if necessary.
	CreateStory(context.Context, *CreateStoryRequest) (*CreateStoryResponse, error)
	// GetStory returns the metadata for a previously-scraped story.
	GetStory(context.Context, *GetStoryRequest) (*GetStoryResponse, error)
}

func RegisterStoryServiceServer(s *grpc.Server, srv StoryServiceServer) {
	s.RegisterService(&_StoryService_serviceDesc, srv)
}

func _StoryService_CreateStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).CreateStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/story.StoryService/CreateStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).CreateStory(ctx, req.(*CreateStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryService_GetStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).GetStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/story.StoryService/GetStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).GetStory(ctx, req.(*GetStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "story.StoryService",
	HandlerType: (*StoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStory",
			Handler:    _StoryService_CreateStory_Handler,
		},
		{
			MethodName: "GetStory",
			Handler:    _StoryService_GetStory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/story/story.proto",
}

func init() { proto.RegisterFile("api/story/story.proto", fileDescriptor_story_b693fe2c963d8e94) }

var fileDescriptor_story_b693fe2c963d8e94 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x4f, 0x83, 0x40,
	0x10, 0x15, 0x5a, 0x88, 0x0e, 0x4d, 0x6d, 0x46, 0x6d, 0x91, 0x13, 0x72, 0x68, 0x7a, 0xa2, 0x49,
	0x3d, 0x7b, 0x6a, 0xa2, 0x77, 0xb8, 0x9a, 0x98, 0xb5, 0x6c, 0x52, 0x92, 0xc2, 0xe2, 0xee, 0x60,
	0xd2, 0x1f, 0xe2, 0xff, 0x35, 0xec, 0x2e, 0xda, 0x0f, 0x8f, 0x5e, 0x08, 0xef, 0x63, 0x67, 0x5e,
	0xde, 0x2e, 0xdc, 0xb1, 0xa6, 0x5c, 0x2a, 0x12, 0x72, 0x6f, 0xbe, 0x69, 0x23, 0x05, 0x09, 0xf4,
	0x34, 0x88, 0x66, 0x9d, 0xba, 0x11, 0x55, 0x25, 0xea, 0xa5, 0x22, 0x46, 0xad, 0x32, 0x7a, 0x32,
	0x07, 0x5c, 0x4b, 0xce, 0x88, 0xe7, 0x9d, 0x2f, 0xe3, 0x1f, 0x2d, 0x57, 0x84, 0x13, 0x18, 0xb4,
	0x72, 0x17, 0x3a, 0xb1, 0xb3, 0xb8, 0xca, 0xba, 0xdf, 0xe4, 0x0d, 0x6e, 0x8e, 0x7c, 0xaa, 0x11,
	0xb5, 0xe2, 0x38, 0x07, 0xdf, 0x8c, 0xd3, 0xde, 0x60, 0x35, 0x4e, 0xcd, 0x92, 0x34, 0xd7, 0x6c,
	0x66, 0x55, 0x8c, 0x61, 0x58, 0x30, 0x62, 0xa1, 0xab, 0x5d, 0xa3, 0xd4, 0x44, 0x34, 0xb3, 0xb4,
	0x92, 0x3c, 0xc0, 0xf5, 0x0b, 0xa7, 0xa3, 0x14, 0x63, 0x70, 0xcb, 0xc2, 0x86, 0x70, 0xcb, 0x22,
	0x79, 0x85, 0xc9, 0xaf, 0xe5, 0xdf, 0x03, 0xec, 0xc1, 0xd3, 0xf0, 0x74, 0x6d, 0x5f, 0x86, 0xfb,
	0x53, 0x06, 0xde, 0x82, 0x47, 0x25, 0xed, 0x78, 0x38, 0xd0, 0x9c, 0x01, 0x38, 0x05, 0x9f, 0xb5,
	0xb4, 0x15, 0x32, 0x1c, 0x6a, 0xda, 0x22, 0x8c, 0x21, 0xa8, 0xdb, 0x6a, 0xbd, 0x65, 0x0d, 0x71,
	0xa9, 0x42, 0x2f, 0x76, 0x16, 0x5e, 0x76, 0x48, 0xad, 0xbe, 0x1c, 0x18, 0xe9, 0xdd, 0x39, 0x97,
	0x9f, 0xe5, 0x86, 0xe3, 0x33, 0x04, 0x07, 0x6d, 0xe3, 0xbd, 0x8d, 0x7b, 0x7e, 0x53, 0x51, 0xf4,
	0x97, 0x64, 0xba, 0x49, 0x2e, 0xf0, 0x09, 0x2e, 0xfb, 0xc6, 0x70, 0x6a, 0x9d, 0x27, 0x2d, 0x47,
	0xb3, 0x33, 0xbe, 0x3f, 0xfe, 0xee, 0xeb, 0x37, 0xf2, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf3,
	0xf8, 0xc6, 0x9f, 0x5c, 0x02, 0x00, 0x00,
}
