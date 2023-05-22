// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: proto/movie.proto

package movie

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

// MovieApiClient is the client API for MovieApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieApiClient interface {
	GetMovies(ctx context.Context, in *RequestMovies, opts ...grpc.CallOption) (*ResponseMovies, error)
	AddMovieToFav(ctx context.Context, in *RequestAddMovieToFav, opts ...grpc.CallOption) (*ResponseAddMovieToFav, error)
	DeleteMovieFromFav(ctx context.Context, in *RequestDeleteMovieFromFav, opts ...grpc.CallOption) (*ResponseDeleteMovieFromFav, error)
	SearchForMovie(ctx context.Context, in *RequestSearchForMovie, opts ...grpc.CallOption) (*ResponseMovies, error)
	GetFavMovies(ctx context.Context, in *RequestMovies, opts ...grpc.CallOption) (*ResponseMovies, error)
}

type movieApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieApiClient(cc grpc.ClientConnInterface) MovieApiClient {
	return &movieApiClient{cc}
}

func (c *movieApiClient) GetMovies(ctx context.Context, in *RequestMovies, opts ...grpc.CallOption) (*ResponseMovies, error) {
	out := new(ResponseMovies)
	err := c.cc.Invoke(ctx, "/MovieApi/GetMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieApiClient) AddMovieToFav(ctx context.Context, in *RequestAddMovieToFav, opts ...grpc.CallOption) (*ResponseAddMovieToFav, error) {
	out := new(ResponseAddMovieToFav)
	err := c.cc.Invoke(ctx, "/MovieApi/AddMovieToFav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieApiClient) DeleteMovieFromFav(ctx context.Context, in *RequestDeleteMovieFromFav, opts ...grpc.CallOption) (*ResponseDeleteMovieFromFav, error) {
	out := new(ResponseDeleteMovieFromFav)
	err := c.cc.Invoke(ctx, "/MovieApi/DeleteMovieFromFav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieApiClient) SearchForMovie(ctx context.Context, in *RequestSearchForMovie, opts ...grpc.CallOption) (*ResponseMovies, error) {
	out := new(ResponseMovies)
	err := c.cc.Invoke(ctx, "/MovieApi/SearchForMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieApiClient) GetFavMovies(ctx context.Context, in *RequestMovies, opts ...grpc.CallOption) (*ResponseMovies, error) {
	out := new(ResponseMovies)
	err := c.cc.Invoke(ctx, "/MovieApi/GetFavMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieApiServer is the server API for MovieApi service.
// All implementations must embed UnimplementedMovieApiServer
// for forward compatibility
type MovieApiServer interface {
	GetMovies(context.Context, *RequestMovies) (*ResponseMovies, error)
	AddMovieToFav(context.Context, *RequestAddMovieToFav) (*ResponseAddMovieToFav, error)
	DeleteMovieFromFav(context.Context, *RequestDeleteMovieFromFav) (*ResponseDeleteMovieFromFav, error)
	SearchForMovie(context.Context, *RequestSearchForMovie) (*ResponseMovies, error)
	GetFavMovies(context.Context, *RequestMovies) (*ResponseMovies, error)
	mustEmbedUnimplementedMovieApiServer()
}

// UnimplementedMovieApiServer must be embedded to have forward compatible implementations.
type UnimplementedMovieApiServer struct {
}

func (UnimplementedMovieApiServer) GetMovies(context.Context, *RequestMovies) (*ResponseMovies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovies not implemented")
}
func (UnimplementedMovieApiServer) AddMovieToFav(context.Context, *RequestAddMovieToFav) (*ResponseAddMovieToFav, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMovieToFav not implemented")
}
func (UnimplementedMovieApiServer) DeleteMovieFromFav(context.Context, *RequestDeleteMovieFromFav) (*ResponseDeleteMovieFromFav, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovieFromFav not implemented")
}
func (UnimplementedMovieApiServer) SearchForMovie(context.Context, *RequestSearchForMovie) (*ResponseMovies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForMovie not implemented")
}
func (UnimplementedMovieApiServer) GetFavMovies(context.Context, *RequestMovies) (*ResponseMovies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavMovies not implemented")
}
func (UnimplementedMovieApiServer) mustEmbedUnimplementedMovieApiServer() {}

// UnsafeMovieApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieApiServer will
// result in compilation errors.
type UnsafeMovieApiServer interface {
	mustEmbedUnimplementedMovieApiServer()
}

func RegisterMovieApiServer(s grpc.ServiceRegistrar, srv MovieApiServer) {
	s.RegisterService(&MovieApi_ServiceDesc, srv)
}

func _MovieApi_GetMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMovies)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieApiServer).GetMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieApi/GetMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieApiServer).GetMovies(ctx, req.(*RequestMovies))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieApi_AddMovieToFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAddMovieToFav)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieApiServer).AddMovieToFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieApi/AddMovieToFav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieApiServer).AddMovieToFav(ctx, req.(*RequestAddMovieToFav))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieApi_DeleteMovieFromFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeleteMovieFromFav)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieApiServer).DeleteMovieFromFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieApi/DeleteMovieFromFav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieApiServer).DeleteMovieFromFav(ctx, req.(*RequestDeleteMovieFromFav))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieApi_SearchForMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSearchForMovie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieApiServer).SearchForMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieApi/SearchForMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieApiServer).SearchForMovie(ctx, req.(*RequestSearchForMovie))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieApi_GetFavMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMovies)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieApiServer).GetFavMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieApi/GetFavMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieApiServer).GetFavMovies(ctx, req.(*RequestMovies))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieApi_ServiceDesc is the grpc.ServiceDesc for MovieApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MovieApi",
	HandlerType: (*MovieApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovies",
			Handler:    _MovieApi_GetMovies_Handler,
		},
		{
			MethodName: "AddMovieToFav",
			Handler:    _MovieApi_AddMovieToFav_Handler,
		},
		{
			MethodName: "DeleteMovieFromFav",
			Handler:    _MovieApi_DeleteMovieFromFav_Handler,
		},
		{
			MethodName: "SearchForMovie",
			Handler:    _MovieApi_SearchForMovie_Handler,
		},
		{
			MethodName: "GetFavMovies",
			Handler:    _MovieApi_GetFavMovies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/movie.proto",
}
