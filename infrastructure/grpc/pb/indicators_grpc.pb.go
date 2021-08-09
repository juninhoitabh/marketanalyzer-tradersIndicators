// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IndicatorsServiceClient is the client API for IndicatorsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndicatorsServiceClient interface {
	Ema(ctx context.Context, in *StandardRequest, opts ...grpc.CallOption) (*StandardResult, error)
	Sma(ctx context.Context, in *StandardRequest, opts ...grpc.CallOption) (*StandardResult, error)
	Macd(ctx context.Context, in *MacdRequest, opts ...grpc.CallOption) (*MacdResult, error)
	StochOscillator(ctx context.Context, in *StochRequest, opts ...grpc.CallOption) (*StochResult, error)
	Cci(ctx context.Context, in *StandardDatasRequest, opts ...grpc.CallOption) (*StandardResult, error)
	Atr(ctx context.Context, in *StandardDatasRequest, opts ...grpc.CallOption) (*StandardResult, error)
	Sar(ctx context.Context, in *StandardDatasOnlyRequest, opts ...grpc.CallOption) (*StandardResult, error)
	Dmi(ctx context.Context, in *StandardDatasRequest, opts ...grpc.CallOption) (*DmiResult, error)
	Adx(ctx context.Context, in *StandardDatasRequest, opts ...grpc.CallOption) (*StandardResult, error)
}

type indicatorsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIndicatorsServiceClient(cc grpc.ClientConnInterface) IndicatorsServiceClient {
	return &indicatorsServiceClient{cc}
}

func (c *indicatorsServiceClient) Ema(ctx context.Context, in *StandardRequest, opts ...grpc.CallOption) (*StandardResult, error) {
	out := new(StandardResult)
	err := c.cc.Invoke(ctx, "/indicators.IndicatorsService/Ema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indicatorsServiceClient) Sma(ctx context.Context, in *StandardRequest, opts ...grpc.CallOption) (*StandardResult, error) {
	out := new(StandardResult)
	err := c.cc.Invoke(ctx, "/indicators.IndicatorsService/Sma", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indicatorsServiceClient) Macd(ctx context.Context, in *MacdRequest, opts ...grpc.CallOption) (*MacdResult, error) {
	out := new(MacdResult)
	err := c.cc.Invoke(ctx, "/indicators.IndicatorsService/Macd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indicatorsServiceClient) StochOscillator(ctx context.Context, in *StochRequest, opts ...grpc.CallOption) (*StochResult, error) {
	out := new(StochResult)
	err := c.cc.Invoke(ctx, "/indicators.IndicatorsService/StochOscillator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indicatorsServiceClient) Cci(ctx context.Context, in *StandardDatasRequest, opts ...grpc.CallOption) (*StandardResult, error) {
	out := new(StandardResult)
	err := c.cc.Invoke(ctx, "/indicators.IndicatorsService/Cci", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indicatorsServiceClient) Atr(ctx context.Context, in *StandardDatasRequest, opts ...grpc.CallOption) (*StandardResult, error) {
	out := new(StandardResult)
	err := c.cc.Invoke(ctx, "/indicators.IndicatorsService/Atr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indicatorsServiceClient) Sar(ctx context.Context, in *StandardDatasOnlyRequest, opts ...grpc.CallOption) (*StandardResult, error) {
	out := new(StandardResult)
	err := c.cc.Invoke(ctx, "/indicators.IndicatorsService/Sar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indicatorsServiceClient) Dmi(ctx context.Context, in *StandardDatasRequest, opts ...grpc.CallOption) (*DmiResult, error) {
	out := new(DmiResult)
	err := c.cc.Invoke(ctx, "/indicators.IndicatorsService/Dmi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indicatorsServiceClient) Adx(ctx context.Context, in *StandardDatasRequest, opts ...grpc.CallOption) (*StandardResult, error) {
	out := new(StandardResult)
	err := c.cc.Invoke(ctx, "/indicators.IndicatorsService/Adx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndicatorsServiceServer is the server API for IndicatorsService service.
// All implementations must embed UnimplementedIndicatorsServiceServer
// for forward compatibility
type IndicatorsServiceServer interface {
	Ema(context.Context, *StandardRequest) (*StandardResult, error)
	Sma(context.Context, *StandardRequest) (*StandardResult, error)
	Macd(context.Context, *MacdRequest) (*MacdResult, error)
	StochOscillator(context.Context, *StochRequest) (*StochResult, error)
	Cci(context.Context, *StandardDatasRequest) (*StandardResult, error)
	Atr(context.Context, *StandardDatasRequest) (*StandardResult, error)
	Sar(context.Context, *StandardDatasOnlyRequest) (*StandardResult, error)
	Dmi(context.Context, *StandardDatasRequest) (*DmiResult, error)
	Adx(context.Context, *StandardDatasRequest) (*StandardResult, error)
	mustEmbedUnimplementedIndicatorsServiceServer()
}

// UnimplementedIndicatorsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIndicatorsServiceServer struct {
}

func (UnimplementedIndicatorsServiceServer) Ema(context.Context, *StandardRequest) (*StandardResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ema not implemented")
}
func (UnimplementedIndicatorsServiceServer) Sma(context.Context, *StandardRequest) (*StandardResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sma not implemented")
}
func (UnimplementedIndicatorsServiceServer) Macd(context.Context, *MacdRequest) (*MacdResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Macd not implemented")
}
func (UnimplementedIndicatorsServiceServer) StochOscillator(context.Context, *StochRequest) (*StochResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StochOscillator not implemented")
}
func (UnimplementedIndicatorsServiceServer) Cci(context.Context, *StandardDatasRequest) (*StandardResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cci not implemented")
}
func (UnimplementedIndicatorsServiceServer) Atr(context.Context, *StandardDatasRequest) (*StandardResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Atr not implemented")
}
func (UnimplementedIndicatorsServiceServer) Sar(context.Context, *StandardDatasOnlyRequest) (*StandardResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sar not implemented")
}
func (UnimplementedIndicatorsServiceServer) Dmi(context.Context, *StandardDatasRequest) (*DmiResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dmi not implemented")
}
func (UnimplementedIndicatorsServiceServer) Adx(context.Context, *StandardDatasRequest) (*StandardResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Adx not implemented")
}
func (UnimplementedIndicatorsServiceServer) mustEmbedUnimplementedIndicatorsServiceServer() {}

// UnsafeIndicatorsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndicatorsServiceServer will
// result in compilation errors.
type UnsafeIndicatorsServiceServer interface {
	mustEmbedUnimplementedIndicatorsServiceServer()
}

func RegisterIndicatorsServiceServer(s grpc.ServiceRegistrar, srv IndicatorsServiceServer) {
	s.RegisterService(&IndicatorsService_ServiceDesc, srv)
}

func _IndicatorsService_Ema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StandardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicatorsServiceServer).Ema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indicators.IndicatorsService/Ema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicatorsServiceServer).Ema(ctx, req.(*StandardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndicatorsService_Sma_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StandardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicatorsServiceServer).Sma(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indicators.IndicatorsService/Sma",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicatorsServiceServer).Sma(ctx, req.(*StandardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndicatorsService_Macd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MacdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicatorsServiceServer).Macd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indicators.IndicatorsService/Macd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicatorsServiceServer).Macd(ctx, req.(*MacdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndicatorsService_StochOscillator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StochRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicatorsServiceServer).StochOscillator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indicators.IndicatorsService/StochOscillator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicatorsServiceServer).StochOscillator(ctx, req.(*StochRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndicatorsService_Cci_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StandardDatasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicatorsServiceServer).Cci(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indicators.IndicatorsService/Cci",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicatorsServiceServer).Cci(ctx, req.(*StandardDatasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndicatorsService_Atr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StandardDatasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicatorsServiceServer).Atr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indicators.IndicatorsService/Atr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicatorsServiceServer).Atr(ctx, req.(*StandardDatasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndicatorsService_Sar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StandardDatasOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicatorsServiceServer).Sar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indicators.IndicatorsService/Sar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicatorsServiceServer).Sar(ctx, req.(*StandardDatasOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndicatorsService_Dmi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StandardDatasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicatorsServiceServer).Dmi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indicators.IndicatorsService/Dmi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicatorsServiceServer).Dmi(ctx, req.(*StandardDatasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndicatorsService_Adx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StandardDatasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicatorsServiceServer).Adx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indicators.IndicatorsService/Adx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicatorsServiceServer).Adx(ctx, req.(*StandardDatasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IndicatorsService_ServiceDesc is the grpc.ServiceDesc for IndicatorsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IndicatorsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "indicators.IndicatorsService",
	HandlerType: (*IndicatorsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ema",
			Handler:    _IndicatorsService_Ema_Handler,
		},
		{
			MethodName: "Sma",
			Handler:    _IndicatorsService_Sma_Handler,
		},
		{
			MethodName: "Macd",
			Handler:    _IndicatorsService_Macd_Handler,
		},
		{
			MethodName: "StochOscillator",
			Handler:    _IndicatorsService_StochOscillator_Handler,
		},
		{
			MethodName: "Cci",
			Handler:    _IndicatorsService_Cci_Handler,
		},
		{
			MethodName: "Atr",
			Handler:    _IndicatorsService_Atr_Handler,
		},
		{
			MethodName: "Sar",
			Handler:    _IndicatorsService_Sar_Handler,
		},
		{
			MethodName: "Dmi",
			Handler:    _IndicatorsService_Dmi_Handler,
		},
		{
			MethodName: "Adx",
			Handler:    _IndicatorsService_Adx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofile/indicators.proto",
}
