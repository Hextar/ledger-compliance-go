// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package schema

import (
	context "context"
	schema "github.com/codenotary/immudb/pkg/api/schema"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LcServiceClient is the client API for LcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LcServiceClient interface {
	// immudb primitives
	// setters and getters
	Set(ctx context.Context, in *schema.KeyValue, opts ...grpc.CallOption) (*schema.Index, error)
	Get(ctx context.Context, in *schema.Key, opts ...grpc.CallOption) (*schema.Item, error)
	SafeSet(ctx context.Context, in *schema.SafeSetOptions, opts ...grpc.CallOption) (*schema.Proof, error)
	SafeGet(ctx context.Context, in *schema.SafeGetOptions, opts ...grpc.CallOption) (*schema.SafeItem, error)
	// batch
	SetBatch(ctx context.Context, in *schema.KVList, opts ...grpc.CallOption) (*schema.Index, error)
	GetBatch(ctx context.Context, in *schema.KeyList, opts ...grpc.CallOption) (*schema.ItemList, error)
	// scanners
	Scan(ctx context.Context, in *schema.ScanOptions, opts ...grpc.CallOption) (*schema.ItemList, error)
	History(ctx context.Context, in *schema.Key, opts ...grpc.CallOption) (*schema.ItemList, error)
	ZAdd(ctx context.Context, in *schema.ZAddOptions, opts ...grpc.CallOption) (*schema.Index, error)
	SafeZAdd(ctx context.Context, in *schema.SafeZAddOptions, opts ...grpc.CallOption) (*schema.Proof, error)
	ZScan(ctx context.Context, in *schema.ZScanOptions, opts ...grpc.CallOption) (*schema.ItemList, error)
	// mixed
	CurrentRoot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*schema.Root, error)
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*schema.HealthResponse, error)
	// ledger compliance extensions
	ReportTamper(ctx context.Context, in *ReportOptions, opts ...grpc.CallOption) (*empty.Empty, error)
	SendData(ctx context.Context, opts ...grpc.CallOption) (LcService_SendDataClient, error)
}

type lcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLcServiceClient(cc grpc.ClientConnInterface) LcServiceClient {
	return &lcServiceClient{cc}
}

func (c *lcServiceClient) Set(ctx context.Context, in *schema.KeyValue, opts ...grpc.CallOption) (*schema.Index, error) {
	out := new(schema.Index)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) Get(ctx context.Context, in *schema.Key, opts ...grpc.CallOption) (*schema.Item, error) {
	out := new(schema.Item)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) SafeSet(ctx context.Context, in *schema.SafeSetOptions, opts ...grpc.CallOption) (*schema.Proof, error) {
	out := new(schema.Proof)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/SafeSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) SafeGet(ctx context.Context, in *schema.SafeGetOptions, opts ...grpc.CallOption) (*schema.SafeItem, error) {
	out := new(schema.SafeItem)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/SafeGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) SetBatch(ctx context.Context, in *schema.KVList, opts ...grpc.CallOption) (*schema.Index, error) {
	out := new(schema.Index)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/SetBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) GetBatch(ctx context.Context, in *schema.KeyList, opts ...grpc.CallOption) (*schema.ItemList, error) {
	out := new(schema.ItemList)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/GetBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) Scan(ctx context.Context, in *schema.ScanOptions, opts ...grpc.CallOption) (*schema.ItemList, error) {
	out := new(schema.ItemList)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/Scan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) History(ctx context.Context, in *schema.Key, opts ...grpc.CallOption) (*schema.ItemList, error) {
	out := new(schema.ItemList)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/History", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) ZAdd(ctx context.Context, in *schema.ZAddOptions, opts ...grpc.CallOption) (*schema.Index, error) {
	out := new(schema.Index)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/ZAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) SafeZAdd(ctx context.Context, in *schema.SafeZAddOptions, opts ...grpc.CallOption) (*schema.Proof, error) {
	out := new(schema.Proof)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/SafeZAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) ZScan(ctx context.Context, in *schema.ZScanOptions, opts ...grpc.CallOption) (*schema.ItemList, error) {
	out := new(schema.ItemList)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/ZScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) CurrentRoot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*schema.Root, error) {
	out := new(schema.Root)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/CurrentRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*schema.HealthResponse, error) {
	out := new(schema.HealthResponse)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) ReportTamper(ctx context.Context, in *ReportOptions, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/lc.schema.LcService/ReportTamper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcServiceClient) SendData(ctx context.Context, opts ...grpc.CallOption) (LcService_SendDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LcService_serviceDesc.Streams[0], "/lc.schema.LcService/SendData", opts...)
	if err != nil {
		return nil, err
	}
	x := &lcServiceSendDataClient{stream}
	return x, nil
}

type LcService_SendDataClient interface {
	Send(*Data) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type lcServiceSendDataClient struct {
	grpc.ClientStream
}

func (x *lcServiceSendDataClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *lcServiceSendDataClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LcServiceServer is the server API for LcService service.
// All implementations must embed UnimplementedLcServiceServer
// for forward compatibility
type LcServiceServer interface {
	// immudb primitives
	// setters and getters
	Set(context.Context, *schema.KeyValue) (*schema.Index, error)
	Get(context.Context, *schema.Key) (*schema.Item, error)
	SafeSet(context.Context, *schema.SafeSetOptions) (*schema.Proof, error)
	SafeGet(context.Context, *schema.SafeGetOptions) (*schema.SafeItem, error)
	// batch
	SetBatch(context.Context, *schema.KVList) (*schema.Index, error)
	GetBatch(context.Context, *schema.KeyList) (*schema.ItemList, error)
	// scanners
	Scan(context.Context, *schema.ScanOptions) (*schema.ItemList, error)
	History(context.Context, *schema.Key) (*schema.ItemList, error)
	ZAdd(context.Context, *schema.ZAddOptions) (*schema.Index, error)
	SafeZAdd(context.Context, *schema.SafeZAddOptions) (*schema.Proof, error)
	ZScan(context.Context, *schema.ZScanOptions) (*schema.ItemList, error)
	// mixed
	CurrentRoot(context.Context, *empty.Empty) (*schema.Root, error)
	Health(context.Context, *empty.Empty) (*schema.HealthResponse, error)
	// ledger compliance extensions
	ReportTamper(context.Context, *ReportOptions) (*empty.Empty, error)
	SendData(LcService_SendDataServer) error
	mustEmbedUnimplementedLcServiceServer()
}

// UnimplementedLcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLcServiceServer struct {
}

func (*UnimplementedLcServiceServer) Set(context.Context, *schema.KeyValue) (*schema.Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedLcServiceServer) Get(context.Context, *schema.Key) (*schema.Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedLcServiceServer) SafeSet(context.Context, *schema.SafeSetOptions) (*schema.Proof, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SafeSet not implemented")
}
func (*UnimplementedLcServiceServer) SafeGet(context.Context, *schema.SafeGetOptions) (*schema.SafeItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SafeGet not implemented")
}
func (*UnimplementedLcServiceServer) SetBatch(context.Context, *schema.KVList) (*schema.Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBatch not implemented")
}
func (*UnimplementedLcServiceServer) GetBatch(context.Context, *schema.KeyList) (*schema.ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatch not implemented")
}
func (*UnimplementedLcServiceServer) Scan(context.Context, *schema.ScanOptions) (*schema.ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scan not implemented")
}
func (*UnimplementedLcServiceServer) History(context.Context, *schema.Key) (*schema.ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method History not implemented")
}
func (*UnimplementedLcServiceServer) ZAdd(context.Context, *schema.ZAddOptions) (*schema.Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ZAdd not implemented")
}
func (*UnimplementedLcServiceServer) SafeZAdd(context.Context, *schema.SafeZAddOptions) (*schema.Proof, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SafeZAdd not implemented")
}
func (*UnimplementedLcServiceServer) ZScan(context.Context, *schema.ZScanOptions) (*schema.ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ZScan not implemented")
}
func (*UnimplementedLcServiceServer) CurrentRoot(context.Context, *empty.Empty) (*schema.Root, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentRoot not implemented")
}
func (*UnimplementedLcServiceServer) Health(context.Context, *empty.Empty) (*schema.HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedLcServiceServer) ReportTamper(context.Context, *ReportOptions) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportTamper not implemented")
}
func (*UnimplementedLcServiceServer) SendData(LcService_SendDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SendData not implemented")
}
func (*UnimplementedLcServiceServer) mustEmbedUnimplementedLcServiceServer() {}

func RegisterLcServiceServer(s *grpc.Server, srv LcServiceServer) {
	s.RegisterService(&_LcService_serviceDesc, srv)
}

func _LcService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).Set(ctx, req.(*schema.KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).Get(ctx, req.(*schema.Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_SafeSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.SafeSetOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).SafeSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/SafeSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).SafeSet(ctx, req.(*schema.SafeSetOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_SafeGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.SafeGetOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).SafeGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/SafeGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).SafeGet(ctx, req.(*schema.SafeGetOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_SetBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.KVList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).SetBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/SetBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).SetBatch(ctx, req.(*schema.KVList))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_GetBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.KeyList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).GetBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/GetBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).GetBatch(ctx, req.(*schema.KeyList))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_Scan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.ScanOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).Scan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/Scan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).Scan(ctx, req.(*schema.ScanOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_History_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).History(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/History",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).History(ctx, req.(*schema.Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_ZAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.ZAddOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).ZAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/ZAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).ZAdd(ctx, req.(*schema.ZAddOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_SafeZAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.SafeZAddOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).SafeZAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/SafeZAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).SafeZAdd(ctx, req.(*schema.SafeZAddOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_ZScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.ZScanOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).ZScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/ZScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).ZScan(ctx, req.(*schema.ZScanOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_CurrentRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).CurrentRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/CurrentRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).CurrentRoot(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_ReportTamper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcServiceServer).ReportTamper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lc.schema.LcService/ReportTamper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcServiceServer).ReportTamper(ctx, req.(*ReportOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _LcService_SendData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LcServiceServer).SendData(&lcServiceSendDataServer{stream})
}

type LcService_SendDataServer interface {
	Send(*Response) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type lcServiceSendDataServer struct {
	grpc.ServerStream
}

func (x *lcServiceSendDataServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *lcServiceSendDataServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lc.schema.LcService",
	HandlerType: (*LcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _LcService_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LcService_Get_Handler,
		},
		{
			MethodName: "SafeSet",
			Handler:    _LcService_SafeSet_Handler,
		},
		{
			MethodName: "SafeGet",
			Handler:    _LcService_SafeGet_Handler,
		},
		{
			MethodName: "SetBatch",
			Handler:    _LcService_SetBatch_Handler,
		},
		{
			MethodName: "GetBatch",
			Handler:    _LcService_GetBatch_Handler,
		},
		{
			MethodName: "Scan",
			Handler:    _LcService_Scan_Handler,
		},
		{
			MethodName: "History",
			Handler:    _LcService_History_Handler,
		},
		{
			MethodName: "ZAdd",
			Handler:    _LcService_ZAdd_Handler,
		},
		{
			MethodName: "SafeZAdd",
			Handler:    _LcService_SafeZAdd_Handler,
		},
		{
			MethodName: "ZScan",
			Handler:    _LcService_ZScan_Handler,
		},
		{
			MethodName: "CurrentRoot",
			Handler:    _LcService_CurrentRoot_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _LcService_Health_Handler,
		},
		{
			MethodName: "ReportTamper",
			Handler:    _LcService_ReportTamper_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendData",
			Handler:       _LcService_SendData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "lc.proto",
}
