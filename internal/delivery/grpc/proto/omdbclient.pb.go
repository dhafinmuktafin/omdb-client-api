// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: omdbclient.proto

package __

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

type GetMovieByKeywordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int64  `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Keyword string `protobuf:"bytes,2,opt,name=Keyword,proto3" json:"Keyword,omitempty"`
}

func (x *GetMovieByKeywordRequest) Reset() {
	*x = GetMovieByKeywordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdbclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieByKeywordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieByKeywordRequest) ProtoMessage() {}

func (x *GetMovieByKeywordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omdbclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieByKeywordRequest.ProtoReflect.Descriptor instead.
func (*GetMovieByKeywordRequest) Descriptor() ([]byte, []int) {
	return file_omdbclient_proto_rawDescGZIP(), []int{0}
}

func (x *GetMovieByKeywordRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetMovieByKeywordRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type GetMovieByKeywordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchList   []*GetMovieByKeywordResponse_Search `protobuf:"bytes,1,rep,name=SearchList,proto3" json:"SearchList,omitempty"`
	TotalResults string                              `protobuf:"bytes,2,opt,name=totalResults,proto3" json:"totalResults,omitempty"`
	Response     string                              `protobuf:"bytes,3,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *GetMovieByKeywordResponse) Reset() {
	*x = GetMovieByKeywordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdbclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieByKeywordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieByKeywordResponse) ProtoMessage() {}

func (x *GetMovieByKeywordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omdbclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieByKeywordResponse.ProtoReflect.Descriptor instead.
func (*GetMovieByKeywordResponse) Descriptor() ([]byte, []int) {
	return file_omdbclient_proto_rawDescGZIP(), []int{1}
}

func (x *GetMovieByKeywordResponse) GetSearchList() []*GetMovieByKeywordResponse_Search {
	if x != nil {
		return x.SearchList
	}
	return nil
}

func (x *GetMovieByKeywordResponse) GetTotalResults() string {
	if x != nil {
		return x.TotalResults
	}
	return ""
}

func (x *GetMovieByKeywordResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type GetMovieByMovieIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId string `protobuf:"bytes,1,opt,name=MovieId,proto3" json:"MovieId,omitempty"`
}

func (x *GetMovieByMovieIdRequest) Reset() {
	*x = GetMovieByMovieIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdbclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieByMovieIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieByMovieIdRequest) ProtoMessage() {}

func (x *GetMovieByMovieIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omdbclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieByMovieIdRequest.ProtoReflect.Descriptor instead.
func (*GetMovieByMovieIdRequest) Descriptor() ([]byte, []int) {
	return file_omdbclient_proto_rawDescGZIP(), []int{2}
}

func (x *GetMovieByMovieIdRequest) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

type GetMovieByMovieIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string                              `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year       string                              `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	Rated      string                              `protobuf:"bytes,3,opt,name=Rated,proto3" json:"Rated,omitempty"`
	Released   string                              `protobuf:"bytes,4,opt,name=Released,proto3" json:"Released,omitempty"`
	Runtime    string                              `protobuf:"bytes,5,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Genre      string                              `protobuf:"bytes,6,opt,name=Genre,proto3" json:"Genre,omitempty"`
	Director   string                              `protobuf:"bytes,7,opt,name=Director,proto3" json:"Director,omitempty"`
	Writer     string                              `protobuf:"bytes,8,opt,name=Writer,proto3" json:"Writer,omitempty"`
	Actors     string                              `protobuf:"bytes,9,opt,name=Actors,proto3" json:"Actors,omitempty"`
	Plot       string                              `protobuf:"bytes,10,opt,name=Plot,proto3" json:"Plot,omitempty"`
	Language   string                              `protobuf:"bytes,11,opt,name=Language,proto3" json:"Language,omitempty"`
	Country    string                              `protobuf:"bytes,12,opt,name=Country,proto3" json:"Country,omitempty"`
	Awards     string                              `protobuf:"bytes,13,opt,name=Awards,proto3" json:"Awards,omitempty"`
	Poster     string                              `protobuf:"bytes,14,opt,name=Poster,proto3" json:"Poster,omitempty"`
	Ratings    []*GetMovieByMovieIdResponse_Rating `protobuf:"bytes,15,rep,name=Ratings,proto3" json:"Ratings,omitempty"`
	Metascore  string                              `protobuf:"bytes,16,opt,name=Metascore,proto3" json:"Metascore,omitempty"`
	ImdbRating string                              `protobuf:"bytes,17,opt,name=imdbRating,proto3" json:"imdbRating,omitempty"`
	ImdbVotes  string                              `protobuf:"bytes,18,opt,name=imdbVotes,proto3" json:"imdbVotes,omitempty"`
	ImdbID     string                              `protobuf:"bytes,19,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type       string                              `protobuf:"bytes,20,opt,name=Type,proto3" json:"Type,omitempty"`
	DVD        string                              `protobuf:"bytes,21,opt,name=DVD,proto3" json:"DVD,omitempty"`
	BoxOffice  string                              `protobuf:"bytes,22,opt,name=BoxOffice,proto3" json:"BoxOffice,omitempty"`
	Production string                              `protobuf:"bytes,23,opt,name=Production,proto3" json:"Production,omitempty"`
	Website    string                              `protobuf:"bytes,24,opt,name=Website,proto3" json:"Website,omitempty"`
	Response   string                              `protobuf:"bytes,25,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *GetMovieByMovieIdResponse) Reset() {
	*x = GetMovieByMovieIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdbclient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieByMovieIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieByMovieIdResponse) ProtoMessage() {}

func (x *GetMovieByMovieIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omdbclient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieByMovieIdResponse.ProtoReflect.Descriptor instead.
func (*GetMovieByMovieIdResponse) Descriptor() ([]byte, []int) {
	return file_omdbclient_proto_rawDescGZIP(), []int{3}
}

func (x *GetMovieByMovieIdResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetRatings() []*GetMovieByMovieIdResponse_Rating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *GetMovieByMovieIdResponse) GetMetascore() string {
	if x != nil {
		return x.Metascore
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetImdbVotes() string {
	if x != nil {
		return x.ImdbVotes
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetDVD() string {
	if x != nil {
		return x.DVD
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetBoxOffice() string {
	if x != nil {
		return x.BoxOffice
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *GetMovieByMovieIdResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type GetMovieByKeywordResponse_Search struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=Poster,proto3" json:"Poster,omitempty"`
}

func (x *GetMovieByKeywordResponse_Search) Reset() {
	*x = GetMovieByKeywordResponse_Search{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdbclient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieByKeywordResponse_Search) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieByKeywordResponse_Search) ProtoMessage() {}

func (x *GetMovieByKeywordResponse_Search) ProtoReflect() protoreflect.Message {
	mi := &file_omdbclient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieByKeywordResponse_Search.ProtoReflect.Descriptor instead.
func (*GetMovieByKeywordResponse_Search) Descriptor() ([]byte, []int) {
	return file_omdbclient_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetMovieByKeywordResponse_Search) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetMovieByKeywordResponse_Search) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *GetMovieByKeywordResponse_Search) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *GetMovieByKeywordResponse_Search) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetMovieByKeywordResponse_Search) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type GetMovieByMovieIdResponse_Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *GetMovieByMovieIdResponse_Rating) Reset() {
	*x = GetMovieByMovieIdResponse_Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdbclient_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieByMovieIdResponse_Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieByMovieIdResponse_Rating) ProtoMessage() {}

func (x *GetMovieByMovieIdResponse_Rating) ProtoReflect() protoreflect.Message {
	mi := &file_omdbclient_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieByMovieIdResponse_Rating.ProtoReflect.Descriptor instead.
func (*GetMovieByMovieIdResponse_Rating) Descriptor() ([]byte, []int) {
	return file_omdbclient_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetMovieByMovieIdResponse_Rating) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *GetMovieByMovieIdResponse_Rating) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_omdbclient_proto protoreflect.FileDescriptor

var file_omdbclient_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x6d, 0x64, 0x62, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x6d, 0x64, 0x62, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x48,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xa1, 0x02, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6f, 0x6d, 0x64,
	0x62, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x42, 0x79, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x76, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x49, 0x64, 0x22, 0xfb, 0x05, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42,
	0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6c, 0x6f, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x41, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x46, 0x0a, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0f, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6f, 0x6d, 0x64, 0x62, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65,
	0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d,
	0x65, 0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x64, 0x62,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d,
	0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x64, 0x62,
	0x56, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x64,
	0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x56, 0x44, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x44, 0x56, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6f, 0x78, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x6f, 0x78, 0x4f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x36, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x32, 0xd4, 0x01, 0x0a, 0x0a, 0x6f, 0x6d, 0x64, 0x62, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x64, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x79, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x2e, 0x6f, 0x6d, 0x64, 0x62, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x6f, 0x6d, 0x64, 0x62, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x42, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x6f, 0x6d, 0x64,
	0x62, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x42, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x6f, 0x6d, 0x64, 0x62, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_omdbclient_proto_rawDescOnce sync.Once
	file_omdbclient_proto_rawDescData = file_omdbclient_proto_rawDesc
)

func file_omdbclient_proto_rawDescGZIP() []byte {
	file_omdbclient_proto_rawDescOnce.Do(func() {
		file_omdbclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_omdbclient_proto_rawDescData)
	})
	return file_omdbclient_proto_rawDescData
}

var file_omdbclient_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_omdbclient_proto_goTypes = []interface{}{
	(*GetMovieByKeywordRequest)(nil),         // 0: omdbClient.GetMovieByKeywordRequest
	(*GetMovieByKeywordResponse)(nil),        // 1: omdbClient.GetMovieByKeywordResponse
	(*GetMovieByMovieIdRequest)(nil),         // 2: omdbClient.GetMovieByMovieIdRequest
	(*GetMovieByMovieIdResponse)(nil),        // 3: omdbClient.GetMovieByMovieIdResponse
	(*GetMovieByKeywordResponse_Search)(nil), // 4: omdbClient.GetMovieByKeywordResponse.Search
	(*GetMovieByMovieIdResponse_Rating)(nil), // 5: omdbClient.GetMovieByMovieIdResponse.Rating
}
var file_omdbclient_proto_depIdxs = []int32{
	4, // 0: omdbClient.GetMovieByKeywordResponse.SearchList:type_name -> omdbClient.GetMovieByKeywordResponse.Search
	5, // 1: omdbClient.GetMovieByMovieIdResponse.Ratings:type_name -> omdbClient.GetMovieByMovieIdResponse.Rating
	0, // 2: omdbClient.omdbClient.GetMovieListByKeyword:input_type -> omdbClient.GetMovieByKeywordRequest
	2, // 3: omdbClient.omdbClient.GetMovieByMovieId:input_type -> omdbClient.GetMovieByMovieIdRequest
	1, // 4: omdbClient.omdbClient.GetMovieListByKeyword:output_type -> omdbClient.GetMovieByKeywordResponse
	3, // 5: omdbClient.omdbClient.GetMovieByMovieId:output_type -> omdbClient.GetMovieByMovieIdResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_omdbclient_proto_init() }
func file_omdbclient_proto_init() {
	if File_omdbclient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omdbclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieByKeywordRequest); i {
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
		file_omdbclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieByKeywordResponse); i {
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
		file_omdbclient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieByMovieIdRequest); i {
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
		file_omdbclient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieByMovieIdResponse); i {
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
		file_omdbclient_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieByKeywordResponse_Search); i {
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
		file_omdbclient_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieByMovieIdResponse_Rating); i {
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
			RawDescriptor: file_omdbclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_omdbclient_proto_goTypes,
		DependencyIndexes: file_omdbclient_proto_depIdxs,
		MessageInfos:      file_omdbclient_proto_msgTypes,
	}.Build()
	File_omdbclient_proto = out.File
	file_omdbclient_proto_rawDesc = nil
	file_omdbclient_proto_goTypes = nil
	file_omdbclient_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OmdbClientClient is the client API for OmdbClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OmdbClientClient interface {
	GetMovieListByKeyword(ctx context.Context, in *GetMovieByKeywordRequest, opts ...grpc.CallOption) (*GetMovieByKeywordResponse, error)
	GetMovieByMovieId(ctx context.Context, in *GetMovieByMovieIdRequest, opts ...grpc.CallOption) (*GetMovieByMovieIdResponse, error)
}

type omdbClientClient struct {
	cc grpc.ClientConnInterface
}

func NewOmdbClientClient(cc grpc.ClientConnInterface) OmdbClientClient {
	return &omdbClientClient{cc}
}

func (c *omdbClientClient) GetMovieListByKeyword(ctx context.Context, in *GetMovieByKeywordRequest, opts ...grpc.CallOption) (*GetMovieByKeywordResponse, error) {
	out := new(GetMovieByKeywordResponse)
	err := c.cc.Invoke(ctx, "/omdbClient.omdbClient/GetMovieListByKeyword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omdbClientClient) GetMovieByMovieId(ctx context.Context, in *GetMovieByMovieIdRequest, opts ...grpc.CallOption) (*GetMovieByMovieIdResponse, error) {
	out := new(GetMovieByMovieIdResponse)
	err := c.cc.Invoke(ctx, "/omdbClient.omdbClient/GetMovieByMovieId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OmdbClientServer is the server API for OmdbClient service.
type OmdbClientServer interface {
	GetMovieListByKeyword(context.Context, *GetMovieByKeywordRequest) (*GetMovieByKeywordResponse, error)
	GetMovieByMovieId(context.Context, *GetMovieByMovieIdRequest) (*GetMovieByMovieIdResponse, error)
}

// UnimplementedOmdbClientServer can be embedded to have forward compatible implementations.
type UnimplementedOmdbClientServer struct {
}

func (*UnimplementedOmdbClientServer) GetMovieListByKeyword(context.Context, *GetMovieByKeywordRequest) (*GetMovieByKeywordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieListByKeyword not implemented")
}
func (*UnimplementedOmdbClientServer) GetMovieByMovieId(context.Context, *GetMovieByMovieIdRequest) (*GetMovieByMovieIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieByMovieId not implemented")
}

func RegisterOmdbClientServer(s *grpc.Server, srv OmdbClientServer) {
	s.RegisterService(&_OmdbClient_serviceDesc, srv)
}

func _OmdbClient_GetMovieListByKeyword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieByKeywordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmdbClientServer).GetMovieListByKeyword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omdbClient.omdbClient/GetMovieListByKeyword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmdbClientServer).GetMovieListByKeyword(ctx, req.(*GetMovieByKeywordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmdbClient_GetMovieByMovieId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieByMovieIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmdbClientServer).GetMovieByMovieId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omdbClient.omdbClient/GetMovieByMovieId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmdbClientServer).GetMovieByMovieId(ctx, req.(*GetMovieByMovieIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OmdbClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "omdbClient.omdbClient",
	HandlerType: (*OmdbClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovieListByKeyword",
			Handler:    _OmdbClient_GetMovieListByKeyword_Handler,
		},
		{
			MethodName: "GetMovieByMovieId",
			Handler:    _OmdbClient_GetMovieByMovieId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "omdbclient.proto",
}