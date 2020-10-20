// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: logistica.proto

// option go_package = "google.golang.org/grpc/examples/helloworld/helloworld";
// option java_multiple_files = true;
// option java_package = "io.grpc.examples.helloworld";
// option java_outer_classname = "HelloWorldProto";

package logistica

import (
	context "context"
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IdSeguimiento int32  `protobuf:"varint,2,opt,name=idSeguimiento,proto3" json:"idSeguimiento,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistica_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_logistica_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_logistica_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Response) GetIdSeguimiento() int32 {
	if x != nil {
		return x.IdSeguimiento
	}
	return 0
}

type Orden struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto string `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Precio   int32  `protobuf:"varint,3,opt,name=precio,proto3" json:"precio,omitempty"`
	Origen   string `protobuf:"bytes,4,opt,name=origen,proto3" json:"origen,omitempty"`
	Destino  string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
	Tipo     int32  `protobuf:"varint,6,opt,name=tipo,proto3" json:"tipo,omitempty"`
}

func (x *Orden) Reset() {
	*x = Orden{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistica_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orden) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orden) ProtoMessage() {}

func (x *Orden) ProtoReflect() protoreflect.Message {
	mi := &file_logistica_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orden.ProtoReflect.Descriptor instead.
func (*Orden) Descriptor() ([]byte, []int) {
	return file_logistica_proto_rawDescGZIP(), []int{1}
}

func (x *Orden) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Orden) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *Orden) GetPrecio() int32 {
	if x != nil {
		return x.Precio
	}
	return 0
}

func (x *Orden) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *Orden) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *Orden) GetTipo() int32 {
	if x != nil {
		return x.Tipo
	}
	return 0
}

var File_logistica_proto protoreflect.FileDescriptor

var file_logistica_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x22, 0x48, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x24, 0x0a, 0x0d, 0x69, 0x64, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x69, 0x64, 0x53, 0x65, 0x67, 0x75, 0x69,
	0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x65, 0x63, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x72,
	0x65, 0x63, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x32, 0x8f, 0x01, 0x0a, 0x10, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x65, 0x73, 0x12,
	0x38, 0x0a, 0x0d, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x61, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e,
	0x12, 0x10, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x6e, 0x1a, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x13, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x74, 0x61, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f,
	0x12, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x61, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logistica_proto_rawDescOnce sync.Once
	file_logistica_proto_rawDescData = file_logistica_proto_rawDesc
)

func file_logistica_proto_rawDescGZIP() []byte {
	file_logistica_proto_rawDescOnce.Do(func() {
		file_logistica_proto_rawDescData = protoimpl.X.CompressGZIP(file_logistica_proto_rawDescData)
	})
	return file_logistica_proto_rawDescData
}

var file_logistica_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_logistica_proto_goTypes = []interface{}{
	(*Response)(nil), // 0: logistica.response
	(*Orden)(nil),    // 1: logistica.orden
}
var file_logistica_proto_depIdxs = []int32{
	1, // 0: logistica.logisticaOrdenes.ingresarOrden:input_type -> logistica.orden
	0, // 1: logistica.logisticaOrdenes.consultaSeguimiento:input_type -> logistica.response
	0, // 2: logistica.logisticaOrdenes.ingresarOrden:output_type -> logistica.response
	0, // 3: logistica.logisticaOrdenes.consultaSeguimiento:output_type -> logistica.response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_logistica_proto_init() }
func file_logistica_proto_init() {
	if File_logistica_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logistica_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_logistica_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orden); i {
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
			RawDescriptor: file_logistica_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logistica_proto_goTypes,
		DependencyIndexes: file_logistica_proto_depIdxs,
		MessageInfos:      file_logistica_proto_msgTypes,
	}.Build()
	File_logistica_proto = out.File
	file_logistica_proto_rawDesc = nil
	file_logistica_proto_goTypes = nil
	file_logistica_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogisticaOrdenesClient is the client API for LogisticaOrdenes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogisticaOrdenesClient interface {
	IngresarOrden(ctx context.Context, in *Orden, opts ...grpc.CallOption) (*Response, error)
	ConsultaSeguimiento(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Response, error)
}

type logisticaOrdenesClient struct {
	cc grpc.ClientConnInterface
}

func NewLogisticaOrdenesClient(cc grpc.ClientConnInterface) LogisticaOrdenesClient {
	return &logisticaOrdenesClient{cc}
}

func (c *logisticaOrdenesClient) IngresarOrden(ctx context.Context, in *Orden, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/logistica.logisticaOrdenes/ingresarOrden", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticaOrdenesClient) ConsultaSeguimiento(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/logistica.logisticaOrdenes/consultaSeguimiento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogisticaOrdenesServer is the server API for LogisticaOrdenes service.
type LogisticaOrdenesServer interface {
	IngresarOrden(context.Context, *Orden) (*Response, error)
	ConsultaSeguimiento(context.Context, *Response) (*Response, error)
}

// UnimplementedLogisticaOrdenesServer can be embedded to have forward compatible implementations.
type UnimplementedLogisticaOrdenesServer struct {
}

func (*UnimplementedLogisticaOrdenesServer) IngresarOrden(context.Context, *Orden) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IngresarOrden not implemented")
}
func (*UnimplementedLogisticaOrdenesServer) ConsultaSeguimiento(context.Context, *Response) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsultaSeguimiento not implemented")
}

func RegisterLogisticaOrdenesServer(s *grpc.Server, srv LogisticaOrdenesServer) {
	s.RegisterService(&_LogisticaOrdenes_serviceDesc, srv)
}

func _LogisticaOrdenes_IngresarOrden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Orden)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticaOrdenesServer).IngresarOrden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistica.logisticaOrdenes/IngresarOrden",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticaOrdenesServer).IngresarOrden(ctx, req.(*Orden))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticaOrdenes_ConsultaSeguimiento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Response)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticaOrdenesServer).ConsultaSeguimiento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistica.logisticaOrdenes/ConsultaSeguimiento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticaOrdenesServer).ConsultaSeguimiento(ctx, req.(*Response))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogisticaOrdenes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logistica.logisticaOrdenes",
	HandlerType: (*LogisticaOrdenesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ingresarOrden",
			Handler:    _LogisticaOrdenes_IngresarOrden_Handler,
		},
		{
			MethodName: "consultaSeguimiento",
			Handler:    _LogisticaOrdenes_ConsultaSeguimiento_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logistica.proto",
}
