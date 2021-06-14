// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: odpf/metadata/optimus/Job.proto

package optimus

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type JobMetadataKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urn string `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
}

func (x *JobMetadataKey) Reset() {
	*x = JobMetadataKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobMetadataKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobMetadataKey) ProtoMessage() {}

func (x *JobMetadataKey) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobMetadataKey.ProtoReflect.Descriptor instead.
func (*JobMetadataKey) Descriptor() ([]byte, []int) {
	return file_odpf_metadata_optimus_Job_proto_rawDescGZIP(), []int{0}
}

func (x *JobMetadataKey) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

type JobMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urn            string               `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	Name           string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tenant         string               `protobuf:"bytes,3,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Version        int32                `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Description    string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Labels         []*JobLabel          `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
	Owner          string               `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	Task           *JobTask             `protobuf:"bytes,8,opt,name=task,proto3" json:"task,omitempty"`
	Schedule       *JobSchedule         `protobuf:"bytes,9,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Behaviour      *JobBehavior         `protobuf:"bytes,10,opt,name=behaviour,proto3" json:"behaviour,omitempty"`
	Hooks          []*JobHook           `protobuf:"bytes,11,rep,name=hooks,proto3" json:"hooks,omitempty"`
	Dependencies   []*JobDependency     `protobuf:"bytes,12,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Namespace      string               `protobuf:"bytes,13,opt,name=namespace,proto3" json:"namespace,omitempty"`
	EventTimestamp *timestamp.Timestamp `protobuf:"bytes,100,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
}

func (x *JobMetadata) Reset() {
	*x = JobMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobMetadata) ProtoMessage() {}

func (x *JobMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobMetadata.ProtoReflect.Descriptor instead.
func (*JobMetadata) Descriptor() ([]byte, []int) {
	return file_odpf_metadata_optimus_Job_proto_rawDescGZIP(), []int{1}
}

func (x *JobMetadata) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *JobMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobMetadata) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *JobMetadata) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *JobMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *JobMetadata) GetLabels() []*JobLabel {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *JobMetadata) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *JobMetadata) GetTask() *JobTask {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *JobMetadata) GetSchedule() *JobSchedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *JobMetadata) GetBehaviour() *JobBehavior {
	if x != nil {
		return x.Behaviour
	}
	return nil
}

func (x *JobMetadata) GetHooks() []*JobHook {
	if x != nil {
		return x.Hooks
	}
	return nil
}

func (x *JobMetadata) GetDependencies() []*JobDependency {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *JobMetadata) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *JobMetadata) GetEventTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

type JobTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image       string           `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Description string           `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Destination string           `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	Config      []*JobTaskConfig `protobuf:"bytes,5,rep,name=config,proto3" json:"config,omitempty"`
	Window      *JobTaskWindow   `protobuf:"bytes,6,opt,name=window,proto3" json:"window,omitempty"`
	Priority    int32            `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *JobTask) Reset() {
	*x = JobTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTask) ProtoMessage() {}

func (x *JobTask) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTask.ProtoReflect.Descriptor instead.
func (*JobTask) Descriptor() ([]byte, []int) {
	return file_odpf_metadata_optimus_Job_proto_rawDescGZIP(), []int{2}
}

func (x *JobTask) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobTask) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *JobTask) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *JobTask) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *JobTask) GetConfig() []*JobTaskConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *JobTask) GetWindow() *JobTaskWindow {
	if x != nil {
		return x.Window
	}
	return nil
}

func (x *JobTask) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type JobHook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image       string           `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Description string           `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Config      []*JobHookConfig `protobuf:"bytes,4,rep,name=config,proto3" json:"config,omitempty"`
	Type        string           `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	DependsOn   []string         `protobuf:"bytes,6,rep,name=depends_on,json=dependsOn,proto3" json:"depends_on,omitempty"`
}

func (x *JobHook) Reset() {
	*x = JobHook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobHook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobHook) ProtoMessage() {}

func (x *JobHook) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobHook.ProtoReflect.Descriptor instead.
func (*JobHook) Descriptor() ([]byte, []int) {
	return file_odpf_metadata_optimus_Job_proto_rawDescGZIP(), []int{3}
}

func (x *JobHook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobHook) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *JobHook) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *JobHook) GetConfig() []*JobHookConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *JobHook) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *JobHook) GetDependsOn() []string {
	if x != nil {
		return x.DependsOn
	}
	return nil
}

type JobDependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Job    string `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	Type   string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *JobDependency) Reset() {
	*x = JobDependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobDependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobDependency) ProtoMessage() {}

func (x *JobDependency) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobDependency.ProtoReflect.Descriptor instead.
func (*JobDependency) Descriptor() ([]byte, []int) {
	return file_odpf_metadata_optimus_Job_proto_rawDescGZIP(), []int{4}
}

func (x *JobDependency) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *JobDependency) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *JobDependency) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type JobTaskWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size       string `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	Offset     string `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	TruncateTo string `protobuf:"bytes,3,opt,name=truncate_to,json=truncateTo,proto3" json:"truncate_to,omitempty"`
}

func (x *JobTaskWindow) Reset() {
	*x = JobTaskWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobTaskWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTaskWindow) ProtoMessage() {}

func (x *JobTaskWindow) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTaskWindow.ProtoReflect.Descriptor instead.
func (*JobTaskWindow) Descriptor() ([]byte, []int) {
	return file_odpf_metadata_optimus_Job_proto_rawDescGZIP(), []int{5}
}

func (x *JobTaskWindow) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *JobTaskWindow) GetOffset() string {
	if x != nil {
		return x.Offset
	}
	return ""
}

func (x *JobTaskWindow) GetTruncateTo() string {
	if x != nil {
		return x.TruncateTo
	}
	return ""
}

type JobSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Interval  string               `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *JobSchedule) Reset() {
	*x = JobSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSchedule) ProtoMessage() {}

func (x *JobSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSchedule.ProtoReflect.Descriptor instead.
func (*JobSchedule) Descriptor() ([]byte, []int) {
	return file_odpf_metadata_optimus_Job_proto_rawDescGZIP(), []int{6}
}

func (x *JobSchedule) GetStartDate() *timestamp.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *JobSchedule) GetEndDate() *timestamp.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *JobSchedule) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

type JobBehavior struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DependsOnPast bool `protobuf:"varint,1,opt,name=depends_on_past,json=dependsOnPast,proto3" json:"depends_on_past,omitempty"`
	Catchup       bool `protobuf:"varint,2,opt,name=catchup,proto3" json:"catchup,omitempty"`
}

func (x *JobBehavior) Reset() {
	*x = JobBehavior{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobBehavior) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobBehavior) ProtoMessage() {}

func (x *JobBehavior) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobBehavior.ProtoReflect.Descriptor instead.
func (*JobBehavior) Descriptor() ([]byte, []int) {
	return file_odpf_metadata_optimus_Job_proto_rawDescGZIP(), []int{7}
}

func (x *JobBehavior) GetDependsOnPast() bool {
	if x != nil {
		return x.DependsOnPast
	}
	return false
}

func (x *JobBehavior) GetCatchup() bool {
	if x != nil {
		return x.Catchup
	}
	return false
}

type JobLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *JobLabel) Reset() {
	*x = JobLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobLabel) ProtoMessage() {}

func (x *JobLabel) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobLabel.ProtoReflect.Descriptor instead.
func (*JobLabel) Descriptor() ([]byte, []int) {
	return file_odpf_metadata_optimus_Job_proto_rawDescGZIP(), []int{8}
}

func (x *JobLabel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobLabel) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type JobTaskConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *JobTaskConfig) Reset() {
	*x = JobTaskConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobTaskConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTaskConfig) ProtoMessage() {}

func (x *JobTaskConfig) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTaskConfig.ProtoReflect.Descriptor instead.
func (*JobTaskConfig) Descriptor() ([]byte, []int) {
	return file_odpf_metadata_optimus_Job_proto_rawDescGZIP(), []int{9}
}

func (x *JobTaskConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobTaskConfig) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type JobHookConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *JobHookConfig) Reset() {
	*x = JobHookConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobHookConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobHookConfig) ProtoMessage() {}

func (x *JobHookConfig) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_metadata_optimus_Job_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobHookConfig.ProtoReflect.Descriptor instead.
func (*JobHookConfig) Descriptor() ([]byte, []int) {
	return file_odpf_metadata_optimus_Job_proto_rawDescGZIP(), []int{10}
}

func (x *JobHookConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobHookConfig) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_odpf_metadata_optimus_Job_proto protoreflect.FileDescriptor

var file_odpf_metadata_optimus_Job_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x2f, 0x4a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0e, 0x4a, 0x6f, 0x62,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x22, 0xef, 0x04,
	0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e,
	0x4a, 0x6f, 0x62, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e, 0x4a, 0x6f, 0x62,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f,
	0x64, 0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6d, 0x75, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x6f, 0x64, 0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x52, 0x09, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x12, 0x34, 0x0a, 0x05,
	0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x64,
	0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6d, 0x75, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x68, 0x6f, 0x6f,
	0x6b, 0x73, 0x12, 0x48, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73,
	0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0c,
	0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x8f, 0x02, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x64, 0x70, 0x66,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75,
	0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3c, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e,
	0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x52, 0x06, 0x77,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x22, 0xc6, 0x01, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x64, 0x70, 0x66,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75,
	0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x6e, 0x22, 0x4d, 0x0a, 0x0d, 0x4a, 0x6f,
	0x62, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5c, 0x0a, 0x0d, 0x4a, 0x6f, 0x62,
	0x54, 0x61, 0x73, 0x6b, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x75,
	0x6e, 0x63, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x4f, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x42, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x5f,
	0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x64,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x6e, 0x50, 0x61, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x61, 0x74, 0x63, 0x68, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63,
	0x61, 0x74, 0x63, 0x68, 0x75, 0x70, 0x22, 0x34, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x0d,
	0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x0d, 0x4a, 0x6f, 0x62, 0x48, 0x6f,
	0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x53, 0x0a, 0x1f, 0x69, 0x6f, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6d, 0x75, 0x73, 0x42, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x64, 0x70, 0x66, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_odpf_metadata_optimus_Job_proto_rawDescOnce sync.Once
	file_odpf_metadata_optimus_Job_proto_rawDescData = file_odpf_metadata_optimus_Job_proto_rawDesc
)

func file_odpf_metadata_optimus_Job_proto_rawDescGZIP() []byte {
	file_odpf_metadata_optimus_Job_proto_rawDescOnce.Do(func() {
		file_odpf_metadata_optimus_Job_proto_rawDescData = protoimpl.X.CompressGZIP(file_odpf_metadata_optimus_Job_proto_rawDescData)
	})
	return file_odpf_metadata_optimus_Job_proto_rawDescData
}

var file_odpf_metadata_optimus_Job_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_odpf_metadata_optimus_Job_proto_goTypes = []interface{}{
	(*JobMetadataKey)(nil),      // 0: odpf.metadata.optimus.JobMetadataKey
	(*JobMetadata)(nil),         // 1: odpf.metadata.optimus.JobMetadata
	(*JobTask)(nil),             // 2: odpf.metadata.optimus.JobTask
	(*JobHook)(nil),             // 3: odpf.metadata.optimus.JobHook
	(*JobDependency)(nil),       // 4: odpf.metadata.optimus.JobDependency
	(*JobTaskWindow)(nil),       // 5: odpf.metadata.optimus.JobTaskWindow
	(*JobSchedule)(nil),         // 6: odpf.metadata.optimus.JobSchedule
	(*JobBehavior)(nil),         // 7: odpf.metadata.optimus.JobBehavior
	(*JobLabel)(nil),            // 8: odpf.metadata.optimus.JobLabel
	(*JobTaskConfig)(nil),       // 9: odpf.metadata.optimus.JobTaskConfig
	(*JobHookConfig)(nil),       // 10: odpf.metadata.optimus.JobHookConfig
	(*timestamp.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_odpf_metadata_optimus_Job_proto_depIdxs = []int32{
	8,  // 0: odpf.metadata.optimus.JobMetadata.labels:type_name -> odpf.metadata.optimus.JobLabel
	2,  // 1: odpf.metadata.optimus.JobMetadata.task:type_name -> odpf.metadata.optimus.JobTask
	6,  // 2: odpf.metadata.optimus.JobMetadata.schedule:type_name -> odpf.metadata.optimus.JobSchedule
	7,  // 3: odpf.metadata.optimus.JobMetadata.behaviour:type_name -> odpf.metadata.optimus.JobBehavior
	3,  // 4: odpf.metadata.optimus.JobMetadata.hooks:type_name -> odpf.metadata.optimus.JobHook
	4,  // 5: odpf.metadata.optimus.JobMetadata.dependencies:type_name -> odpf.metadata.optimus.JobDependency
	11, // 6: odpf.metadata.optimus.JobMetadata.event_timestamp:type_name -> google.protobuf.Timestamp
	9,  // 7: odpf.metadata.optimus.JobTask.config:type_name -> odpf.metadata.optimus.JobTaskConfig
	5,  // 8: odpf.metadata.optimus.JobTask.window:type_name -> odpf.metadata.optimus.JobTaskWindow
	10, // 9: odpf.metadata.optimus.JobHook.config:type_name -> odpf.metadata.optimus.JobHookConfig
	11, // 10: odpf.metadata.optimus.JobSchedule.start_date:type_name -> google.protobuf.Timestamp
	11, // 11: odpf.metadata.optimus.JobSchedule.end_date:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_odpf_metadata_optimus_Job_proto_init() }
func file_odpf_metadata_optimus_Job_proto_init() {
	if File_odpf_metadata_optimus_Job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_odpf_metadata_optimus_Job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobMetadataKey); i {
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
		file_odpf_metadata_optimus_Job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobMetadata); i {
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
		file_odpf_metadata_optimus_Job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobTask); i {
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
		file_odpf_metadata_optimus_Job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobHook); i {
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
		file_odpf_metadata_optimus_Job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobDependency); i {
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
		file_odpf_metadata_optimus_Job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobTaskWindow); i {
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
		file_odpf_metadata_optimus_Job_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobSchedule); i {
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
		file_odpf_metadata_optimus_Job_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobBehavior); i {
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
		file_odpf_metadata_optimus_Job_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobLabel); i {
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
		file_odpf_metadata_optimus_Job_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobTaskConfig); i {
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
		file_odpf_metadata_optimus_Job_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobHookConfig); i {
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
			RawDescriptor: file_odpf_metadata_optimus_Job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_odpf_metadata_optimus_Job_proto_goTypes,
		DependencyIndexes: file_odpf_metadata_optimus_Job_proto_depIdxs,
		MessageInfos:      file_odpf_metadata_optimus_Job_proto_msgTypes,
	}.Build()
	File_odpf_metadata_optimus_Job_proto = out.File
	file_odpf_metadata_optimus_Job_proto_rawDesc = nil
	file_odpf_metadata_optimus_Job_proto_goTypes = nil
	file_odpf_metadata_optimus_Job_proto_depIdxs = nil
}