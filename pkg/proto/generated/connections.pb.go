// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: connections.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source
	SrcIP          string  `protobuf:"bytes,1,opt,name=srcIP,proto3" json:"srcIP,omitempty"`
	SrcPort        string  `protobuf:"bytes,2,opt,name=srcPort,proto3" json:"srcPort,omitempty"`
	SrcLatitude    float64 `protobuf:"fixed64,3,opt,name=srcLatitude,proto3" json:"srcLatitude,omitempty"`
	SrcLongitude   float64 `protobuf:"fixed64,4,opt,name=srcLongitude,proto3" json:"srcLongitude,omitempty"`
	SrcCityName    string  `protobuf:"bytes,5,opt,name=srcCityName,proto3" json:"srcCityName,omitempty"`
	SrcCountryCode string  `protobuf:"bytes,6,opt,name=srcCountryCode,proto3" json:"srcCountryCode,omitempty"`
	// Destination
	DstIP          string  `protobuf:"bytes,7,opt,name=dstIP,proto3" json:"dstIP,omitempty"`
	DstPort        string  `protobuf:"bytes,8,opt,name=dstPort,proto3" json:"dstPort,omitempty"`
	DstLatitude    float64 `protobuf:"fixed64,9,opt,name=dstLatitude,proto3" json:"dstLatitude,omitempty"`
	DstLongitude   float64 `protobuf:"fixed64,10,opt,name=dstLongitude,proto3" json:"dstLongitude,omitempty"`
	DstCityName    string  `protobuf:"bytes,11,opt,name=dstCityName,proto3" json:"dstCityName,omitempty"`
	DstCountryCode string  `protobuf:"bytes,12,opt,name=dstCountryCode,proto3" json:"dstCountryCode,omitempty"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connections_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_connections_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_connections_proto_rawDescGZIP(), []int{0}
}

func (x *Connection) GetSrcIP() string {
	if x != nil {
		return x.SrcIP
	}
	return ""
}

func (x *Connection) GetSrcPort() string {
	if x != nil {
		return x.SrcPort
	}
	return ""
}

func (x *Connection) GetSrcLatitude() float64 {
	if x != nil {
		return x.SrcLatitude
	}
	return 0
}

func (x *Connection) GetSrcLongitude() float64 {
	if x != nil {
		return x.SrcLongitude
	}
	return 0
}

func (x *Connection) GetSrcCityName() string {
	if x != nil {
		return x.SrcCityName
	}
	return ""
}

func (x *Connection) GetSrcCountryCode() string {
	if x != nil {
		return x.SrcCountryCode
	}
	return ""
}

func (x *Connection) GetDstIP() string {
	if x != nil {
		return x.DstIP
	}
	return ""
}

func (x *Connection) GetDstPort() string {
	if x != nil {
		return x.DstPort
	}
	return ""
}

func (x *Connection) GetDstLatitude() float64 {
	if x != nil {
		return x.DstLatitude
	}
	return 0
}

func (x *Connection) GetDstLongitude() float64 {
	if x != nil {
		return x.DstLongitude
	}
	return 0
}

func (x *Connection) GetDstCityName() string {
	if x != nil {
		return x.DstCityName
	}
	return ""
}

func (x *Connection) GetDstCountryCode() string {
	if x != nil {
		return x.DstCountryCode
	}
	return ""
}

var File_connections_proto protoreflect.FileDescriptor

var file_connections_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x61, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c,
	0x03, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x72, 0x63, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x72,
	0x63, 0x49, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x72, 0x63, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0b, 0x73, 0x72, 0x63, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x73, 0x72, 0x63, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x73, 0x72, 0x63, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x72, 0x63, 0x43, 0x69, 0x74, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x72, 0x63, 0x43, 0x69, 0x74,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x72, 0x63, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x72, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x64, 0x73, 0x74, 0x49, 0x50, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x73,
	0x74, 0x49, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x73, 0x74, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0b, 0x64, 0x73, 0x74, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x64, 0x73, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x64, 0x73, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x73, 0x74, 0x43, 0x69, 0x74, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x73, 0x74, 0x43, 0x69, 0x74,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64,
	0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x4e, 0x0a,
	0x0b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x61, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connections_proto_rawDescOnce sync.Once
	file_connections_proto_rawDescData = file_connections_proto_rawDesc
)

func file_connections_proto_rawDescGZIP() []byte {
	file_connections_proto_rawDescOnce.Do(func() {
		file_connections_proto_rawDescData = protoimpl.X.CompressGZIP(file_connections_proto_rawDescData)
	})
	return file_connections_proto_rawDescData
}

var file_connections_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_connections_proto_goTypes = []interface{}{
	(*Connection)(nil),  // 0: connaections.Connection
	(*empty.Empty)(nil), // 1: google.protobuf.Empty
}
var file_connections_proto_depIdxs = []int32{
	1, // 0: connaections.Connections.Subscribe:input_type -> google.protobuf.Empty
	0, // 1: connaections.Connections.Subscribe:output_type -> connaections.Connection
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_connections_proto_init() }
func file_connections_proto_init() {
	if File_connections_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connections_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_connections_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_connections_proto_goTypes,
		DependencyIndexes: file_connections_proto_depIdxs,
		MessageInfos:      file_connections_proto_msgTypes,
	}.Build()
	File_connections_proto = out.File
	file_connections_proto_rawDesc = nil
	file_connections_proto_goTypes = nil
	file_connections_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnectionsClient is the client API for Connections service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectionsClient interface {
	Subscribe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Connections_SubscribeClient, error)
}

type connectionsClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionsClient(cc grpc.ClientConnInterface) ConnectionsClient {
	return &connectionsClient{cc}
}

func (c *connectionsClient) Subscribe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Connections_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Connections_serviceDesc.Streams[0], "/connaections.Connections/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionsSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Connections_SubscribeClient interface {
	Recv() (*Connection, error)
	grpc.ClientStream
}

type connectionsSubscribeClient struct {
	grpc.ClientStream
}

func (x *connectionsSubscribeClient) Recv() (*Connection, error) {
	m := new(Connection)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnectionsServer is the server API for Connections service.
type ConnectionsServer interface {
	Subscribe(*empty.Empty, Connections_SubscribeServer) error
}

// UnimplementedConnectionsServer can be embedded to have forward compatible implementations.
type UnimplementedConnectionsServer struct {
}

func (*UnimplementedConnectionsServer) Subscribe(*empty.Empty, Connections_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterConnectionsServer(s *grpc.Server, srv ConnectionsServer) {
	s.RegisterService(&_Connections_serviceDesc, srv)
}

func _Connections_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConnectionsServer).Subscribe(m, &connectionsSubscribeServer{stream})
}

type Connections_SubscribeServer interface {
	Send(*Connection) error
	grpc.ServerStream
}

type connectionsSubscribeServer struct {
	grpc.ServerStream
}

func (x *connectionsSubscribeServer) Send(m *Connection) error {
	return x.ServerStream.SendMsg(m)
}

var _Connections_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connaections.Connections",
	HandlerType: (*ConnectionsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Connections_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "connections.proto",
}
