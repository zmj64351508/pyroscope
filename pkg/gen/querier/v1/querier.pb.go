// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: querier/v1/querier.proto

package querierv1

import (
	v1 "github.com/grafana/phlare/pkg/gen/common/v1"
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

type ProfileTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProfileTypesRequest) Reset() {
	*x = ProfileTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileTypesRequest) ProtoMessage() {}

func (x *ProfileTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileTypesRequest.ProtoReflect.Descriptor instead.
func (*ProfileTypesRequest) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{0}
}

type ProfileTypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileTypes []*v1.ProfileType `protobuf:"bytes,1,rep,name=profile_types,json=profileTypes,proto3" json:"profile_types,omitempty"`
}

func (x *ProfileTypesResponse) Reset() {
	*x = ProfileTypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileTypesResponse) ProtoMessage() {}

func (x *ProfileTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileTypesResponse.ProtoReflect.Descriptor instead.
func (*ProfileTypesResponse) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileTypesResponse) GetProfileTypes() []*v1.ProfileType {
	if x != nil {
		return x.ProfileTypes
	}
	return nil
}

type LabelValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LabelValuesRequest) Reset() {
	*x = LabelValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelValuesRequest) ProtoMessage() {}

func (x *LabelValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelValuesRequest.ProtoReflect.Descriptor instead.
func (*LabelValuesRequest) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{2}
}

func (x *LabelValuesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LabelValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *LabelValuesResponse) Reset() {
	*x = LabelValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelValuesResponse) ProtoMessage() {}

func (x *LabelValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelValuesResponse.ProtoReflect.Descriptor instead.
func (*LabelValuesResponse) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{3}
}

func (x *LabelValuesResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type LabelNamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LabelNamesRequest) Reset() {
	*x = LabelNamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelNamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelNamesRequest) ProtoMessage() {}

func (x *LabelNamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelNamesRequest.ProtoReflect.Descriptor instead.
func (*LabelNamesRequest) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{4}
}

type LabelNamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *LabelNamesResponse) Reset() {
	*x = LabelNamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelNamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelNamesResponse) ProtoMessage() {}

func (x *LabelNamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelNamesResponse.ProtoReflect.Descriptor instead.
func (*LabelNamesResponse) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{5}
}

func (x *LabelNamesResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type SeriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matchers []string `protobuf:"bytes,1,rep,name=matchers,proto3" json:"matchers,omitempty"`
}

func (x *SeriesRequest) Reset() {
	*x = SeriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeriesRequest) ProtoMessage() {}

func (x *SeriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeriesRequest.ProtoReflect.Descriptor instead.
func (*SeriesRequest) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{6}
}

func (x *SeriesRequest) GetMatchers() []string {
	if x != nil {
		return x.Matchers
	}
	return nil
}

type SeriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LabelsSet []*v1.Labels `protobuf:"bytes,2,rep,name=labels_set,json=labelsSet,proto3" json:"labels_set,omitempty"`
}

func (x *SeriesResponse) Reset() {
	*x = SeriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeriesResponse) ProtoMessage() {}

func (x *SeriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeriesResponse.ProtoReflect.Descriptor instead.
func (*SeriesResponse) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{7}
}

func (x *SeriesResponse) GetLabelsSet() []*v1.Labels {
	if x != nil {
		return x.LabelsSet
	}
	return nil
}

type SelectMergeStacktracesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileTypeID string `protobuf:"bytes,1,opt,name=profile_typeID,json=profileTypeID,proto3" json:"profile_typeID,omitempty"`
	LabelSelector string `protobuf:"bytes,2,opt,name=label_selector,json=labelSelector,proto3" json:"label_selector,omitempty"`
	Start         int64  `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"` // milliseconds since epoch
	End           int64  `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`     // milliseconds since epoch
}

func (x *SelectMergeStacktracesRequest) Reset() {
	*x = SelectMergeStacktracesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectMergeStacktracesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectMergeStacktracesRequest) ProtoMessage() {}

func (x *SelectMergeStacktracesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectMergeStacktracesRequest.ProtoReflect.Descriptor instead.
func (*SelectMergeStacktracesRequest) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{8}
}

func (x *SelectMergeStacktracesRequest) GetProfileTypeID() string {
	if x != nil {
		return x.ProfileTypeID
	}
	return ""
}

func (x *SelectMergeStacktracesRequest) GetLabelSelector() string {
	if x != nil {
		return x.LabelSelector
	}
	return ""
}

func (x *SelectMergeStacktracesRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *SelectMergeStacktracesRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type SelectMergeStacktracesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flamegraph *FlameGraph `protobuf:"bytes,1,opt,name=flamegraph,proto3" json:"flamegraph,omitempty"`
}

func (x *SelectMergeStacktracesResponse) Reset() {
	*x = SelectMergeStacktracesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectMergeStacktracesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectMergeStacktracesResponse) ProtoMessage() {}

func (x *SelectMergeStacktracesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectMergeStacktracesResponse.ProtoReflect.Descriptor instead.
func (*SelectMergeStacktracesResponse) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{9}
}

func (x *SelectMergeStacktracesResponse) GetFlamegraph() *FlameGraph {
	if x != nil {
		return x.Flamegraph
	}
	return nil
}

type FlameGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names   []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	Levels  []*Level `protobuf:"bytes,2,rep,name=levels,proto3" json:"levels,omitempty"`
	Total   int64    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	MaxSelf int64    `protobuf:"varint,4,opt,name=max_self,json=maxSelf,proto3" json:"max_self,omitempty"`
}

func (x *FlameGraph) Reset() {
	*x = FlameGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlameGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlameGraph) ProtoMessage() {}

func (x *FlameGraph) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlameGraph.ProtoReflect.Descriptor instead.
func (*FlameGraph) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{10}
}

func (x *FlameGraph) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *FlameGraph) GetLevels() []*Level {
	if x != nil {
		return x.Levels
	}
	return nil
}

func (x *FlameGraph) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FlameGraph) GetMaxSelf() int64 {
	if x != nil {
		return x.MaxSelf
	}
	return 0
}

type Level struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []int64 `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *Level) Reset() {
	*x = Level{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Level) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Level) ProtoMessage() {}

func (x *Level) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Level.ProtoReflect.Descriptor instead.
func (*Level) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{11}
}

func (x *Level) GetValues() []int64 {
	if x != nil {
		return x.Values
	}
	return nil
}

type SelectSeriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileTypeID string   `protobuf:"bytes,1,opt,name=profile_typeID,json=profileTypeID,proto3" json:"profile_typeID,omitempty"`
	LabelSelector string   `protobuf:"bytes,2,opt,name=label_selector,json=labelSelector,proto3" json:"label_selector,omitempty"`
	Start         int64    `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"` // milliseconds since epoch
	End           int64    `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`     // milliseconds since epoch
	GroupBy       []string `protobuf:"bytes,5,rep,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
	Step          float64  `protobuf:"fixed64,6,opt,name=step,proto3" json:"step,omitempty"` // Query resolution step width in seconds
}

func (x *SelectSeriesRequest) Reset() {
	*x = SelectSeriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectSeriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectSeriesRequest) ProtoMessage() {}

func (x *SelectSeriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectSeriesRequest.ProtoReflect.Descriptor instead.
func (*SelectSeriesRequest) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{12}
}

func (x *SelectSeriesRequest) GetProfileTypeID() string {
	if x != nil {
		return x.ProfileTypeID
	}
	return ""
}

func (x *SelectSeriesRequest) GetLabelSelector() string {
	if x != nil {
		return x.LabelSelector
	}
	return ""
}

func (x *SelectSeriesRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *SelectSeriesRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *SelectSeriesRequest) GetGroupBy() []string {
	if x != nil {
		return x.GroupBy
	}
	return nil
}

func (x *SelectSeriesRequest) GetStep() float64 {
	if x != nil {
		return x.Step
	}
	return 0
}

type SelectSeriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Series []*v1.Series `protobuf:"bytes,1,rep,name=series,proto3" json:"series,omitempty"`
}

func (x *SelectSeriesResponse) Reset() {
	*x = SelectSeriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_querier_v1_querier_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectSeriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectSeriesResponse) ProtoMessage() {}

func (x *SelectSeriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_querier_v1_querier_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectSeriesResponse.ProtoReflect.Descriptor instead.
func (*SelectSeriesResponse) Descriptor() ([]byte, []int) {
	return file_querier_v1_querier_proto_rawDescGZIP(), []int{13}
}

func (x *SelectSeriesResponse) GetSeries() []*v1.Series {
	if x != nil {
		return x.Series
	}
	return nil
}

var File_querier_v1_querier_proto protoreflect.FileDescriptor

var file_querier_v1_querier_proto_rawDesc = []byte{
	0x0a, 0x18, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65,
	0x72, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15,
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x12, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x13, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x12, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x22, 0x2b, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22,
	0x42, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x0a, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x5f, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x53, 0x65, 0x74, 0x22, 0x95, 0x01, 0x0a, 0x1d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x72, 0x67, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x25, 0x0a, 0x0e,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x58, 0x0a, 0x1e, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x0a, 0x66, 0x6c, 0x61, 0x6d, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x6c, 0x61, 0x6d, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x0a, 0x66, 0x6c, 0x61, 0x6d, 0x65,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x22, 0x7e, 0x0a, 0x0a, 0x46, 0x6c, 0x61, 0x6d, 0x65, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61,
	0x78, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61,
	0x78, 0x53, 0x65, 0x6c, 0x66, 0x22, 0x1f, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xba, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x65, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x22, 0x41, 0x0a, 0x14, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x32, 0x91, 0x04, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x69,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1e, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4d, 0x0a, 0x0a, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1d,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x71, 0x0a, 0x16, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x72, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x12, 0x29, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x72, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x9f, 0x01, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x51,
	0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e,
	0x61, 0x2f, 0x70, 0x68, 0x6c, 0x61, 0x72, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x51, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x51, 0x75,
	0x65, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x69,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0b, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_querier_v1_querier_proto_rawDescOnce sync.Once
	file_querier_v1_querier_proto_rawDescData = file_querier_v1_querier_proto_rawDesc
)

func file_querier_v1_querier_proto_rawDescGZIP() []byte {
	file_querier_v1_querier_proto_rawDescOnce.Do(func() {
		file_querier_v1_querier_proto_rawDescData = protoimpl.X.CompressGZIP(file_querier_v1_querier_proto_rawDescData)
	})
	return file_querier_v1_querier_proto_rawDescData
}

var file_querier_v1_querier_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_querier_v1_querier_proto_goTypes = []interface{}{
	(*ProfileTypesRequest)(nil),            // 0: querier.v1.ProfileTypesRequest
	(*ProfileTypesResponse)(nil),           // 1: querier.v1.ProfileTypesResponse
	(*LabelValuesRequest)(nil),             // 2: querier.v1.LabelValuesRequest
	(*LabelValuesResponse)(nil),            // 3: querier.v1.LabelValuesResponse
	(*LabelNamesRequest)(nil),              // 4: querier.v1.LabelNamesRequest
	(*LabelNamesResponse)(nil),             // 5: querier.v1.LabelNamesResponse
	(*SeriesRequest)(nil),                  // 6: querier.v1.SeriesRequest
	(*SeriesResponse)(nil),                 // 7: querier.v1.SeriesResponse
	(*SelectMergeStacktracesRequest)(nil),  // 8: querier.v1.SelectMergeStacktracesRequest
	(*SelectMergeStacktracesResponse)(nil), // 9: querier.v1.SelectMergeStacktracesResponse
	(*FlameGraph)(nil),                     // 10: querier.v1.FlameGraph
	(*Level)(nil),                          // 11: querier.v1.Level
	(*SelectSeriesRequest)(nil),            // 12: querier.v1.SelectSeriesRequest
	(*SelectSeriesResponse)(nil),           // 13: querier.v1.SelectSeriesResponse
	(*v1.ProfileType)(nil),                 // 14: common.v1.ProfileType
	(*v1.Labels)(nil),                      // 15: common.v1.Labels
	(*v1.Series)(nil),                      // 16: common.v1.Series
}
var file_querier_v1_querier_proto_depIdxs = []int32{
	14, // 0: querier.v1.ProfileTypesResponse.profile_types:type_name -> common.v1.ProfileType
	15, // 1: querier.v1.SeriesResponse.labels_set:type_name -> common.v1.Labels
	10, // 2: querier.v1.SelectMergeStacktracesResponse.flamegraph:type_name -> querier.v1.FlameGraph
	11, // 3: querier.v1.FlameGraph.levels:type_name -> querier.v1.Level
	16, // 4: querier.v1.SelectSeriesResponse.series:type_name -> common.v1.Series
	0,  // 5: querier.v1.QuerierService.ProfileTypes:input_type -> querier.v1.ProfileTypesRequest
	2,  // 6: querier.v1.QuerierService.LabelValues:input_type -> querier.v1.LabelValuesRequest
	4,  // 7: querier.v1.QuerierService.LabelNames:input_type -> querier.v1.LabelNamesRequest
	6,  // 8: querier.v1.QuerierService.Series:input_type -> querier.v1.SeriesRequest
	8,  // 9: querier.v1.QuerierService.SelectMergeStacktraces:input_type -> querier.v1.SelectMergeStacktracesRequest
	12, // 10: querier.v1.QuerierService.SelectSeries:input_type -> querier.v1.SelectSeriesRequest
	1,  // 11: querier.v1.QuerierService.ProfileTypes:output_type -> querier.v1.ProfileTypesResponse
	3,  // 12: querier.v1.QuerierService.LabelValues:output_type -> querier.v1.LabelValuesResponse
	5,  // 13: querier.v1.QuerierService.LabelNames:output_type -> querier.v1.LabelNamesResponse
	7,  // 14: querier.v1.QuerierService.Series:output_type -> querier.v1.SeriesResponse
	9,  // 15: querier.v1.QuerierService.SelectMergeStacktraces:output_type -> querier.v1.SelectMergeStacktracesResponse
	13, // 16: querier.v1.QuerierService.SelectSeries:output_type -> querier.v1.SelectSeriesResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_querier_v1_querier_proto_init() }
func file_querier_v1_querier_proto_init() {
	if File_querier_v1_querier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_querier_v1_querier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileTypesRequest); i {
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
		file_querier_v1_querier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileTypesResponse); i {
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
		file_querier_v1_querier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelValuesRequest); i {
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
		file_querier_v1_querier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelValuesResponse); i {
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
		file_querier_v1_querier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelNamesRequest); i {
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
		file_querier_v1_querier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelNamesResponse); i {
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
		file_querier_v1_querier_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeriesRequest); i {
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
		file_querier_v1_querier_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeriesResponse); i {
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
		file_querier_v1_querier_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectMergeStacktracesRequest); i {
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
		file_querier_v1_querier_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectMergeStacktracesResponse); i {
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
		file_querier_v1_querier_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlameGraph); i {
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
		file_querier_v1_querier_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Level); i {
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
		file_querier_v1_querier_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectSeriesRequest); i {
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
		file_querier_v1_querier_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectSeriesResponse); i {
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
			RawDescriptor: file_querier_v1_querier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_querier_v1_querier_proto_goTypes,
		DependencyIndexes: file_querier_v1_querier_proto_depIdxs,
		MessageInfos:      file_querier_v1_querier_proto_msgTypes,
	}.Build()
	File_querier_v1_querier_proto = out.File
	file_querier_v1_querier_proto_rawDesc = nil
	file_querier_v1_querier_proto_goTypes = nil
	file_querier_v1_querier_proto_depIdxs = nil
}
