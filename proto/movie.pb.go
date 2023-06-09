// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.1
// source: proto/movie.proto

package movie

import (
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

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Poster      string `protobuf:"bytes,3,opt,name=poster,proto3" json:"poster,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Movie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Movie) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *Movie) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RequestMovies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestMovies) Reset() {
	*x = RequestMovies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMovies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMovies) ProtoMessage() {}

func (x *RequestMovies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMovies.ProtoReflect.Descriptor instead.
func (*RequestMovies) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{1}
}

type ResponseMovies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie []*Movie `protobuf:"bytes,1,rep,name=movie,proto3" json:"movie,omitempty"`
}

func (x *ResponseMovies) Reset() {
	*x = ResponseMovies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMovies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMovies) ProtoMessage() {}

func (x *ResponseMovies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMovies.ProtoReflect.Descriptor instead.
func (*ResponseMovies) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseMovies) GetMovie() []*Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

type RequestAddMovieToFav struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *RequestAddMovieToFav) Reset() {
	*x = RequestAddMovieToFav{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAddMovieToFav) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAddMovieToFav) ProtoMessage() {}

func (x *RequestAddMovieToFav) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAddMovieToFav.ProtoReflect.Descriptor instead.
func (*RequestAddMovieToFav) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{3}
}

func (x *RequestAddMovieToFav) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

type ResponseAddMovieToFav struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDone  bool   `protobuf:"varint,1,opt,name=isDone,proto3" json:"isDone,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseAddMovieToFav) Reset() {
	*x = ResponseAddMovieToFav{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseAddMovieToFav) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAddMovieToFav) ProtoMessage() {}

func (x *ResponseAddMovieToFav) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAddMovieToFav.ProtoReflect.Descriptor instead.
func (*ResponseAddMovieToFav) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseAddMovieToFav) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *ResponseAddMovieToFav) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RequestDeleteMovieFromFav struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *RequestDeleteMovieFromFav) Reset() {
	*x = RequestDeleteMovieFromFav{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDeleteMovieFromFav) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDeleteMovieFromFav) ProtoMessage() {}

func (x *RequestDeleteMovieFromFav) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDeleteMovieFromFav.ProtoReflect.Descriptor instead.
func (*RequestDeleteMovieFromFav) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{5}
}

func (x *RequestDeleteMovieFromFav) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ResponseDeleteMovieFromFav struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDone  bool   `protobuf:"varint,1,opt,name=isDone,proto3" json:"isDone,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseDeleteMovieFromFav) Reset() {
	*x = ResponseDeleteMovieFromFav{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDeleteMovieFromFav) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDeleteMovieFromFav) ProtoMessage() {}

func (x *ResponseDeleteMovieFromFav) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDeleteMovieFromFav.ProtoReflect.Descriptor instead.
func (*ResponseDeleteMovieFromFav) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseDeleteMovieFromFav) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *ResponseDeleteMovieFromFav) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RequestSearchForMovie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *RequestSearchForMovie) Reset() {
	*x = RequestSearchForMovie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_movie_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestSearchForMovie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestSearchForMovie) ProtoMessage() {}

func (x *RequestSearchForMovie) ProtoReflect() protoreflect.Message {
	mi := &file_proto_movie_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestSearchForMovie.ProtoReflect.Descriptor instead.
func (*RequestSearchForMovie) Descriptor() ([]byte, []int) {
	return file_proto_movie_proto_rawDescGZIP(), []int{7}
}

func (x *RequestSearchForMovie) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

var File_proto_movie_proto protoreflect.FileDescriptor

var file_proto_movie_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x0e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a,
	0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x34, 0x0a, 0x14, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x6f,
	0x46, 0x61, 0x76, 0x12, 0x1c, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x22, 0x49, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x64, 0x64,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x6f, 0x46, 0x61, 0x76, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73,
	0x44, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x44, 0x6f,
	0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a, 0x19,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x61, 0x76, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x1a, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x46, 0x61, 0x76, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x15, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6f, 0x72, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x32, 0xb3, 0x02, 0x0a,
	0x08, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x41, 0x70, 0x69, 0x12, 0x2c, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x0e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x54, 0x6f, 0x46, 0x61, 0x76, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x54, 0x6f, 0x46, 0x61, 0x76, 0x1a,
	0x16, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x54, 0x6f, 0x46, 0x61, 0x76, 0x12, 0x4d, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x61, 0x76, 0x12, 0x1a, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x61, 0x76, 0x1a, 0x1b, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x46, 0x61, 0x76, 0x12, 0x39, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x46, 0x6f, 0x72, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x16, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6f, 0x72, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x12, 0x2f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x12, 0x0e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x62, 0x64, 0x75, 0x6c, 0x72, 0x61, 0x68, 0x6d, 0x61, 0x6e, 0x2d, 0x41, 0x77, 0x61,
	0x64, 0x68, 0x2f, 0x4d, 0x6f, 0x76, 0x69, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_movie_proto_rawDescOnce sync.Once
	file_proto_movie_proto_rawDescData = file_proto_movie_proto_rawDesc
)

func file_proto_movie_proto_rawDescGZIP() []byte {
	file_proto_movie_proto_rawDescOnce.Do(func() {
		file_proto_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_movie_proto_rawDescData)
	})
	return file_proto_movie_proto_rawDescData
}

var file_proto_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_movie_proto_goTypes = []interface{}{
	(*Movie)(nil),                      // 0: Movie
	(*RequestMovies)(nil),              // 1: RequestMovies
	(*ResponseMovies)(nil),             // 2: ResponseMovies
	(*RequestAddMovieToFav)(nil),       // 3: RequestAddMovieToFav
	(*ResponseAddMovieToFav)(nil),      // 4: ResponseAddMovieToFav
	(*RequestDeleteMovieFromFav)(nil),  // 5: RequestDeleteMovieFromFav
	(*ResponseDeleteMovieFromFav)(nil), // 6: ResponseDeleteMovieFromFav
	(*RequestSearchForMovie)(nil),      // 7: RequestSearchForMovie
}
var file_proto_movie_proto_depIdxs = []int32{
	0, // 0: ResponseMovies.movie:type_name -> Movie
	0, // 1: RequestAddMovieToFav.movie:type_name -> Movie
	1, // 2: MovieApi.GetMovies:input_type -> RequestMovies
	3, // 3: MovieApi.AddMovieToFav:input_type -> RequestAddMovieToFav
	5, // 4: MovieApi.DeleteMovieFromFav:input_type -> RequestDeleteMovieFromFav
	7, // 5: MovieApi.SearchForMovie:input_type -> RequestSearchForMovie
	1, // 6: MovieApi.GetFavMovies:input_type -> RequestMovies
	2, // 7: MovieApi.GetMovies:output_type -> ResponseMovies
	4, // 8: MovieApi.AddMovieToFav:output_type -> ResponseAddMovieToFav
	6, // 9: MovieApi.DeleteMovieFromFav:output_type -> ResponseDeleteMovieFromFav
	2, // 10: MovieApi.SearchForMovie:output_type -> ResponseMovies
	2, // 11: MovieApi.GetFavMovies:output_type -> ResponseMovies
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_movie_proto_init() }
func file_proto_movie_proto_init() {
	if File_proto_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
		file_proto_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMovies); i {
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
		file_proto_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMovies); i {
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
		file_proto_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAddMovieToFav); i {
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
		file_proto_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseAddMovieToFav); i {
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
		file_proto_movie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDeleteMovieFromFav); i {
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
		file_proto_movie_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseDeleteMovieFromFav); i {
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
		file_proto_movie_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestSearchForMovie); i {
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
			RawDescriptor: file_proto_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_movie_proto_goTypes,
		DependencyIndexes: file_proto_movie_proto_depIdxs,
		MessageInfos:      file_proto_movie_proto_msgTypes,
	}.Build()
	File_proto_movie_proto = out.File
	file_proto_movie_proto_rawDesc = nil
	file_proto_movie_proto_goTypes = nil
	file_proto_movie_proto_depIdxs = nil
}
