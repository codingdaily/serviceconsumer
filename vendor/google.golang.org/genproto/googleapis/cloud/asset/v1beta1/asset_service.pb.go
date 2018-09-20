// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/asset/v1beta1/asset_service.proto

package asset

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	math "math"
)

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

// Asset content type.
type ContentType int32

const (
	// Unspecified content type.
	ContentType_CONTENT_TYPE_UNSPECIFIED ContentType = 0
	// Resource metadata.
	ContentType_RESOURCE ContentType = 1
	// The actual IAM policy set on a resource.
	ContentType_IAM_POLICY ContentType = 2
)

var ContentType_name = map[int32]string{
	0: "CONTENT_TYPE_UNSPECIFIED",
	1: "RESOURCE",
	2: "IAM_POLICY",
}

var ContentType_value = map[string]int32{
	"CONTENT_TYPE_UNSPECIFIED": 0,
	"RESOURCE":                 1,
	"IAM_POLICY":               2,
}

func (x ContentType) String() string {
	return proto.EnumName(ContentType_name, int32(x))
}

func (ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{0}
}

// Export asset request.
type ExportAssetsRequest struct {
	// Required. The relative name of the root asset. Can only be an organization
	// number (such as "organizations/123"), or a project id (such as
	// "projects/my-project-id") or a project number (such as "projects/12345").
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Timestamp to take an asset snapshot. This can only be set to a timestamp in
	// the past or of the current time. If not specified, the current time will be
	// used. Due to delays in resource data collection and indexing, there is a
	// volatile window during which running the same query may get different
	// results.
	ReadTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=read_time,json=readTime,proto3" json:"read_time,omitempty"`
	// A list of asset types of which to take a snapshot for. Example:
	// "google.compute.disk". If specified, only matching assets will be returned.
	AssetTypes []string `protobuf:"bytes,3,rep,name=asset_types,json=assetTypes,proto3" json:"asset_types,omitempty"`
	// Asset content type. If not specified, no content but the asset name will be
	// returned.
	ContentType ContentType `protobuf:"varint,4,opt,name=content_type,json=contentType,proto3,enum=google.cloud.asset.v1beta1.ContentType" json:"content_type,omitempty"`
	// Required. Output configuration indicating where the results will be output
	// to. All results will be in newline delimited JSON format.
	OutputConfig         *OutputConfig `protobuf:"bytes,5,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExportAssetsRequest) Reset()         { *m = ExportAssetsRequest{} }
func (m *ExportAssetsRequest) String() string { return proto.CompactTextString(m) }
func (*ExportAssetsRequest) ProtoMessage()    {}
func (*ExportAssetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{0}
}

func (m *ExportAssetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportAssetsRequest.Unmarshal(m, b)
}
func (m *ExportAssetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportAssetsRequest.Marshal(b, m, deterministic)
}
func (m *ExportAssetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportAssetsRequest.Merge(m, src)
}
func (m *ExportAssetsRequest) XXX_Size() int {
	return xxx_messageInfo_ExportAssetsRequest.Size(m)
}
func (m *ExportAssetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportAssetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportAssetsRequest proto.InternalMessageInfo

func (m *ExportAssetsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ExportAssetsRequest) GetReadTime() *timestamp.Timestamp {
	if m != nil {
		return m.ReadTime
	}
	return nil
}

func (m *ExportAssetsRequest) GetAssetTypes() []string {
	if m != nil {
		return m.AssetTypes
	}
	return nil
}

func (m *ExportAssetsRequest) GetContentType() ContentType {
	if m != nil {
		return m.ContentType
	}
	return ContentType_CONTENT_TYPE_UNSPECIFIED
}

func (m *ExportAssetsRequest) GetOutputConfig() *OutputConfig {
	if m != nil {
		return m.OutputConfig
	}
	return nil
}

// The export asset response. This message is returned by the
// [google.longrunning.Operations.GetOperation][google.longrunning.Operations.GetOperation] method in the returned
// [google.longrunning.Operation.response][google.longrunning.Operation.response] field.
type ExportAssetsResponse struct {
	// Time the snapshot was taken.
	ReadTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=read_time,json=readTime,proto3" json:"read_time,omitempty"`
	// Output configuration indicating where the results were output to.
	// All results are in JSON format.
	OutputConfig         *OutputConfig `protobuf:"bytes,2,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExportAssetsResponse) Reset()         { *m = ExportAssetsResponse{} }
func (m *ExportAssetsResponse) String() string { return proto.CompactTextString(m) }
func (*ExportAssetsResponse) ProtoMessage()    {}
func (*ExportAssetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{1}
}

func (m *ExportAssetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportAssetsResponse.Unmarshal(m, b)
}
func (m *ExportAssetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportAssetsResponse.Marshal(b, m, deterministic)
}
func (m *ExportAssetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportAssetsResponse.Merge(m, src)
}
func (m *ExportAssetsResponse) XXX_Size() int {
	return xxx_messageInfo_ExportAssetsResponse.Size(m)
}
func (m *ExportAssetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportAssetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportAssetsResponse proto.InternalMessageInfo

func (m *ExportAssetsResponse) GetReadTime() *timestamp.Timestamp {
	if m != nil {
		return m.ReadTime
	}
	return nil
}

func (m *ExportAssetsResponse) GetOutputConfig() *OutputConfig {
	if m != nil {
		return m.OutputConfig
	}
	return nil
}

// Batch get assets history request.
type BatchGetAssetsHistoryRequest struct {
	// Required. The relative name of the root asset. It can only be an
	// organization number (such as "organizations/123"), or a project id (such as
	// "projects/my-project-id")"or a project number (such as "projects/12345").
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// A list of the full names of the assets. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	// Example:
	// "//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1".
	//
	// The request becomes a no-op if the asset name list is empty, and the max
	// size of the asset name list is 100 in one request.
	AssetNames []string `protobuf:"bytes,2,rep,name=asset_names,json=assetNames,proto3" json:"asset_names,omitempty"`
	// Required. The content type.
	ContentType ContentType `protobuf:"varint,3,opt,name=content_type,json=contentType,proto3,enum=google.cloud.asset.v1beta1.ContentType" json:"content_type,omitempty"`
	// Required. The time window for the asset history. The start time is
	// required. The returned results contain all temporal assets whose time
	// window overlap with read_time_window.
	ReadTimeWindow       *TimeWindow `protobuf:"bytes,4,opt,name=read_time_window,json=readTimeWindow,proto3" json:"read_time_window,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BatchGetAssetsHistoryRequest) Reset()         { *m = BatchGetAssetsHistoryRequest{} }
func (m *BatchGetAssetsHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*BatchGetAssetsHistoryRequest) ProtoMessage()    {}
func (*BatchGetAssetsHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{2}
}

func (m *BatchGetAssetsHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchGetAssetsHistoryRequest.Unmarshal(m, b)
}
func (m *BatchGetAssetsHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchGetAssetsHistoryRequest.Marshal(b, m, deterministic)
}
func (m *BatchGetAssetsHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchGetAssetsHistoryRequest.Merge(m, src)
}
func (m *BatchGetAssetsHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_BatchGetAssetsHistoryRequest.Size(m)
}
func (m *BatchGetAssetsHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchGetAssetsHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchGetAssetsHistoryRequest proto.InternalMessageInfo

func (m *BatchGetAssetsHistoryRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *BatchGetAssetsHistoryRequest) GetAssetNames() []string {
	if m != nil {
		return m.AssetNames
	}
	return nil
}

func (m *BatchGetAssetsHistoryRequest) GetContentType() ContentType {
	if m != nil {
		return m.ContentType
	}
	return ContentType_CONTENT_TYPE_UNSPECIFIED
}

func (m *BatchGetAssetsHistoryRequest) GetReadTimeWindow() *TimeWindow {
	if m != nil {
		return m.ReadTimeWindow
	}
	return nil
}

// Batch get assets history response.
type BatchGetAssetsHistoryResponse struct {
	// A list of assets with valid time windows.
	Assets               []*TemporalAsset `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BatchGetAssetsHistoryResponse) Reset()         { *m = BatchGetAssetsHistoryResponse{} }
func (m *BatchGetAssetsHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*BatchGetAssetsHistoryResponse) ProtoMessage()    {}
func (*BatchGetAssetsHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{3}
}

func (m *BatchGetAssetsHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchGetAssetsHistoryResponse.Unmarshal(m, b)
}
func (m *BatchGetAssetsHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchGetAssetsHistoryResponse.Marshal(b, m, deterministic)
}
func (m *BatchGetAssetsHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchGetAssetsHistoryResponse.Merge(m, src)
}
func (m *BatchGetAssetsHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_BatchGetAssetsHistoryResponse.Size(m)
}
func (m *BatchGetAssetsHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchGetAssetsHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchGetAssetsHistoryResponse proto.InternalMessageInfo

func (m *BatchGetAssetsHistoryResponse) GetAssets() []*TemporalAsset {
	if m != nil {
		return m.Assets
	}
	return nil
}

// Output configuration for export assets destination.
type OutputConfig struct {
	// Asset export destination.
	//
	// Types that are valid to be assigned to Destination:
	//	*OutputConfig_GcsDestination
	Destination          isOutputConfig_Destination `protobuf_oneof:"destination"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *OutputConfig) Reset()         { *m = OutputConfig{} }
func (m *OutputConfig) String() string { return proto.CompactTextString(m) }
func (*OutputConfig) ProtoMessage()    {}
func (*OutputConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{4}
}

func (m *OutputConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputConfig.Unmarshal(m, b)
}
func (m *OutputConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputConfig.Marshal(b, m, deterministic)
}
func (m *OutputConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputConfig.Merge(m, src)
}
func (m *OutputConfig) XXX_Size() int {
	return xxx_messageInfo_OutputConfig.Size(m)
}
func (m *OutputConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OutputConfig proto.InternalMessageInfo

type isOutputConfig_Destination interface {
	isOutputConfig_Destination()
}

type OutputConfig_GcsDestination struct {
	GcsDestination *GcsDestination `protobuf:"bytes,1,opt,name=gcs_destination,json=gcsDestination,proto3,oneof"`
}

func (*OutputConfig_GcsDestination) isOutputConfig_Destination() {}

func (m *OutputConfig) GetDestination() isOutputConfig_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *OutputConfig) GetGcsDestination() *GcsDestination {
	if x, ok := m.GetDestination().(*OutputConfig_GcsDestination); ok {
		return x.GcsDestination
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*OutputConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _OutputConfig_OneofMarshaler, _OutputConfig_OneofUnmarshaler, _OutputConfig_OneofSizer, []interface{}{
		(*OutputConfig_GcsDestination)(nil),
	}
}

func _OutputConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*OutputConfig)
	// destination
	switch x := m.Destination.(type) {
	case *OutputConfig_GcsDestination:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GcsDestination); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("OutputConfig.Destination has unexpected type %T", x)
	}
	return nil
}

func _OutputConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*OutputConfig)
	switch tag {
	case 1: // destination.gcs_destination
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GcsDestination)
		err := b.DecodeMessage(msg)
		m.Destination = &OutputConfig_GcsDestination{msg}
		return true, err
	default:
		return false, nil
	}
}

func _OutputConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*OutputConfig)
	// destination
	switch x := m.Destination.(type) {
	case *OutputConfig_GcsDestination:
		s := proto.Size(x.GcsDestination)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A Google Cloud Storage (GCS) location.
type GcsDestination struct {
	// The path of the GCS objects. It's the same path that is used by gsutil, for
	// example: "gs://bucket_name/object_path". See:
	// https://cloud.google.com/storage/docs/viewing-editing-metadata for more
	// information.
	Uri                  string   `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GcsDestination) Reset()         { *m = GcsDestination{} }
func (m *GcsDestination) String() string { return proto.CompactTextString(m) }
func (*GcsDestination) ProtoMessage()    {}
func (*GcsDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{5}
}

func (m *GcsDestination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GcsDestination.Unmarshal(m, b)
}
func (m *GcsDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GcsDestination.Marshal(b, m, deterministic)
}
func (m *GcsDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GcsDestination.Merge(m, src)
}
func (m *GcsDestination) XXX_Size() int {
	return xxx_messageInfo_GcsDestination.Size(m)
}
func (m *GcsDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_GcsDestination.DiscardUnknown(m)
}

var xxx_messageInfo_GcsDestination proto.InternalMessageInfo

func (m *GcsDestination) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func init() {
	proto.RegisterType((*ExportAssetsRequest)(nil), "google.cloud.asset.v1beta1.ExportAssetsRequest")
	proto.RegisterType((*ExportAssetsResponse)(nil), "google.cloud.asset.v1beta1.ExportAssetsResponse")
	proto.RegisterType((*BatchGetAssetsHistoryRequest)(nil), "google.cloud.asset.v1beta1.BatchGetAssetsHistoryRequest")
	proto.RegisterType((*BatchGetAssetsHistoryResponse)(nil), "google.cloud.asset.v1beta1.BatchGetAssetsHistoryResponse")
	proto.RegisterType((*OutputConfig)(nil), "google.cloud.asset.v1beta1.OutputConfig")
	proto.RegisterType((*GcsDestination)(nil), "google.cloud.asset.v1beta1.GcsDestination")
	proto.RegisterEnum("google.cloud.asset.v1beta1.ContentType", ContentType_name, ContentType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssetServiceClient interface {
	// Exports assets with time and resource types to a given Google Cloud Storage
	// location. The output format is newline-delimited JSON.
	// This API implements the [google.longrunning.Operation][google.longrunning.Operation] API allowing users
	// to keep track of the export.
	ExportAssets(ctx context.Context, in *ExportAssetsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Batch gets assets update history that overlaps a time window.
	// For RESOURCE content, this API outputs history with asset in both
	// non-delete or deleted status.
	// For IAM_POLICY content, this API only outputs history when asset and its
	// attached IAM POLICY both exist. So there may be gaps in the output history.
	BatchGetAssetsHistory(ctx context.Context, in *BatchGetAssetsHistoryRequest, opts ...grpc.CallOption) (*BatchGetAssetsHistoryResponse, error)
}

type assetServiceClient struct {
	cc *grpc.ClientConn
}

func NewAssetServiceClient(cc *grpc.ClientConn) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) ExportAssets(ctx context.Context, in *ExportAssetsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.asset.v1beta1.AssetService/ExportAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) BatchGetAssetsHistory(ctx context.Context, in *BatchGetAssetsHistoryRequest, opts ...grpc.CallOption) (*BatchGetAssetsHistoryResponse, error) {
	out := new(BatchGetAssetsHistoryResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.asset.v1beta1.AssetService/BatchGetAssetsHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
type AssetServiceServer interface {
	// Exports assets with time and resource types to a given Google Cloud Storage
	// location. The output format is newline-delimited JSON.
	// This API implements the [google.longrunning.Operation][google.longrunning.Operation] API allowing users
	// to keep track of the export.
	ExportAssets(context.Context, *ExportAssetsRequest) (*longrunning.Operation, error)
	// Batch gets assets update history that overlaps a time window.
	// For RESOURCE content, this API outputs history with asset in both
	// non-delete or deleted status.
	// For IAM_POLICY content, this API only outputs history when asset and its
	// attached IAM POLICY both exist. So there may be gaps in the output history.
	BatchGetAssetsHistory(context.Context, *BatchGetAssetsHistoryRequest) (*BatchGetAssetsHistoryResponse, error)
}

func RegisterAssetServiceServer(s *grpc.Server, srv AssetServiceServer) {
	s.RegisterService(&_AssetService_serviceDesc, srv)
}

func _AssetService_ExportAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).ExportAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.asset.v1beta1.AssetService/ExportAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).ExportAssets(ctx, req.(*ExportAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_BatchGetAssetsHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetAssetsHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).BatchGetAssetsHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.asset.v1beta1.AssetService/BatchGetAssetsHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).BatchGetAssetsHistory(ctx, req.(*BatchGetAssetsHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AssetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.asset.v1beta1.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExportAssets",
			Handler:    _AssetService_ExportAssets_Handler,
		},
		{
			MethodName: "BatchGetAssetsHistory",
			Handler:    _AssetService_BatchGetAssetsHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/asset/v1beta1/asset_service.proto",
}

func init() {
	proto.RegisterFile("google/cloud/asset/v1beta1/asset_service.proto", fileDescriptor_565cfbddaa85b7d6)
}

var fileDescriptor_565cfbddaa85b7d6 = []byte{
	// 766 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0xce, 0xbd, 0x08, 0x26, 0x21, 0x37, 0x77, 0xfa, 0x23, 0x2b, 0x82, 0x12, 0xb9, 0x52,
	0x09, 0x59, 0xd8, 0x22, 0x48, 0x2d, 0xa5, 0xaa, 0x2a, 0x12, 0x52, 0x48, 0x55, 0x92, 0xc8, 0x04,
	0x2a, 0x10, 0x92, 0x35, 0x71, 0x06, 0xd7, 0x55, 0x32, 0xe3, 0x7a, 0x26, 0x50, 0x5a, 0x75, 0xc3,
	0x2b, 0x74, 0xdf, 0x45, 0x97, 0x5d, 0xf5, 0x35, 0xda, 0x6d, 0x1f, 0xa0, 0x9b, 0xbe, 0x42, 0xf7,
	0x95, 0x67, 0x1c, 0x70, 0x44, 0x30, 0x6a, 0xd9, 0xe5, 0xcc, 0x7c, 0xdf, 0x77, 0xe6, 0x7c, 0xe7,
	0xf8, 0x04, 0x18, 0x2e, 0xa5, 0x6e, 0x0f, 0x9b, 0x4e, 0x8f, 0x0e, 0xba, 0x26, 0x62, 0x0c, 0x73,
	0xf3, 0x68, 0xa9, 0x83, 0x39, 0x5a, 0x92, 0x91, 0xcd, 0x70, 0x70, 0xe4, 0x39, 0xd8, 0xf0, 0x03,
	0xca, 0x29, 0xcc, 0x4b, 0xbc, 0x21, 0xf0, 0x86, 0x40, 0x18, 0x11, 0x3e, 0x3f, 0x1b, 0x69, 0x21,
	0xdf, 0x33, 0x11, 0x21, 0x94, 0x23, 0xee, 0x51, 0xc2, 0x24, 0x33, 0xbf, 0x70, 0x55, 0xa6, 0x21,
	0xf0, 0x6e, 0x04, 0xec, 0x51, 0xe2, 0x06, 0x03, 0x42, 0x3c, 0xe2, 0x9a, 0xd4, 0xc7, 0xc1, 0x88,
	0xda, 0x7c, 0x04, 0x12, 0x51, 0x67, 0x70, 0x68, 0x72, 0xaf, 0x8f, 0x19, 0x47, 0x7d, 0x5f, 0x02,
	0xf4, 0x4f, 0x2a, 0xb8, 0x51, 0x7b, 0xe3, 0xd3, 0x80, 0xaf, 0x09, 0x71, 0x0b, 0xbf, 0x1e, 0x60,
	0xc6, 0xe1, 0x6d, 0x30, 0xe9, 0xa3, 0x00, 0x13, 0xae, 0x29, 0x05, 0xa5, 0x38, 0x6d, 0x45, 0x11,
	0x7c, 0x00, 0xa6, 0x03, 0x8c, 0xba, 0x76, 0xa8, 0xa3, 0xa9, 0x05, 0xa5, 0x98, 0x2e, 0xe7, 0x23,
	0x73, 0x8c, 0x61, 0x12, 0xa3, 0x3d, 0x4c, 0x62, 0x4d, 0x85, 0xe0, 0x30, 0x84, 0xf3, 0x20, 0x2d,
	0x8d, 0xe2, 0x27, 0x3e, 0x66, 0x5a, 0xaa, 0x90, 0x2a, 0x4e, 0x5b, 0x40, 0x1c, 0xb5, 0xc3, 0x13,
	0xf8, 0x0c, 0x64, 0x1c, 0x4a, 0x38, 0x26, 0x12, 0xa2, 0xfd, 0x53, 0x50, 0x8a, 0xd9, 0xf2, 0x82,
	0x71, 0xb9, 0x93, 0x46, 0x55, 0xe2, 0x43, 0xbe, 0x95, 0x76, 0xce, 0x03, 0xb8, 0x05, 0x66, 0xe8,
	0x80, 0xfb, 0x03, 0x6e, 0x3b, 0x94, 0x1c, 0x7a, 0xae, 0xf6, 0xaf, 0x78, 0x69, 0x31, 0x49, 0xac,
	0x29, 0x08, 0x55, 0x81, 0xb7, 0x32, 0x34, 0x16, 0xe9, 0x1f, 0x15, 0x70, 0x73, 0xd4, 0x24, 0xe6,
	0x53, 0xc2, 0xf0, 0xa8, 0x1b, 0xca, 0x1f, 0xb8, 0x71, 0xe1, 0x81, 0xea, 0xb5, 0x1e, 0xf8, 0x4b,
	0x01, 0xb3, 0x15, 0xc4, 0x9d, 0x97, 0x1b, 0x38, 0x7a, 0xe2, 0xa6, 0xc7, 0x38, 0x0d, 0x4e, 0xae,
	0x6a, 0xe7, 0x59, 0x57, 0x08, 0xea, 0x63, 0xa6, 0xa9, 0xb1, 0xae, 0x34, 0xc2, 0x93, 0x0b, 0x5d,
	0x49, 0x5d, 0xa3, 0x2b, 0x2d, 0x90, 0x3b, 0x73, 0xcb, 0x3e, 0xf6, 0x48, 0x97, 0x1e, 0x8b, 0x2e,
	0xa7, 0xcb, 0xf7, 0x92, 0xf4, 0x42, 0xc3, 0x5e, 0x08, 0xb4, 0x95, 0x1d, 0x1a, 0x28, 0x63, 0xbd,
	0x03, 0xe6, 0x2e, 0x29, 0x3b, 0x6a, 0xd0, 0x1a, 0x98, 0x94, 0x1f, 0x8d, 0xa6, 0x14, 0x52, 0xc5,
	0x74, 0x79, 0x31, 0x31, 0x11, 0xee, 0xfb, 0x34, 0x40, 0x3d, 0x21, 0x65, 0x45, 0x44, 0x9d, 0x83,
	0x4c, 0xdc, 0x79, 0xb8, 0x03, 0xfe, 0x73, 0x1d, 0x66, 0x77, 0x31, 0xe3, 0x1e, 0x11, 0x1f, 0x5b,
	0xd4, 0xf9, 0x52, 0x92, 0xf6, 0x86, 0xc3, 0xd6, 0xcf, 0x19, 0x9b, 0x13, 0x56, 0xd6, 0x1d, 0x39,
	0xa9, 0xcc, 0x80, 0x74, 0x4c, 0x52, 0xd7, 0x41, 0x76, 0x94, 0x02, 0x73, 0x20, 0x35, 0x08, 0xbc,
	0xa8, 0x7f, 0xe1, 0xcf, 0x52, 0x1d, 0xa4, 0x63, 0x5e, 0xc3, 0x59, 0xa0, 0x55, 0x9b, 0x8d, 0x76,
	0xad, 0xd1, 0xb6, 0xdb, 0x7b, 0xad, 0x9a, 0xbd, 0xd3, 0xd8, 0x6e, 0xd5, 0xaa, 0xf5, 0xa7, 0xf5,
	0xda, 0x7a, 0x6e, 0x02, 0x66, 0xc0, 0x94, 0x55, 0xdb, 0x6e, 0xee, 0x58, 0xd5, 0x5a, 0x4e, 0x81,
	0x59, 0x00, 0xea, 0x6b, 0x5b, 0x76, 0xab, 0xf9, 0xbc, 0x5e, 0xdd, 0xcb, 0xa9, 0xe5, 0x1f, 0x29,
	0x90, 0x11, 0x65, 0x6f, 0xcb, 0x35, 0x06, 0xbf, 0x2a, 0x20, 0x13, 0x1f, 0x79, 0x68, 0x26, 0x55,
	0x37, 0x66, 0x83, 0xe4, 0xe7, 0x86, 0x84, 0xd8, 0x82, 0x32, 0x9a, 0xc3, 0x05, 0xa5, 0x7b, 0xa7,
	0xdf, 0x7f, 0x7e, 0x50, 0x1d, 0x7d, 0xf1, 0x6c, 0xbb, 0xbd, 0x93, 0x33, 0xf9, 0xd8, 0x0f, 0xe8,
	0x2b, 0xec, 0x70, 0x66, 0x96, 0xde, 0xaf, 0xe2, 0x98, 0xf0, 0xaa, 0x52, 0xda, 0x5f, 0xd6, 0x8d,
	0x0b, 0x78, 0x1a, 0xb8, 0x88, 0x78, 0x6f, 0xe5, 0xde, 0x1b, 0x43, 0x82, 0xa7, 0x2a, 0xb8, 0x35,
	0x76, 0x4c, 0xe0, 0x4a, 0x52, 0x51, 0x49, 0x1f, 0x54, 0xfe, 0xe1, 0x5f, 0x30, 0xe5, 0x4c, 0xea,
	0x5c, 0x54, 0x4e, 0x60, 0x39, 0xb1, 0xf2, 0xce, 0x38, 0x8d, 0xfd, 0x15, 0x78, 0xff, 0xea, 0xfa,
	0x47, 0xa8, 0x11, 0xb3, 0xf2, 0x45, 0x01, 0x77, 0x1c, 0xda, 0x4f, 0x78, 0x76, 0xe5, 0xff, 0xf8,
	0x04, 0xb4, 0xc2, 0xf5, 0xd5, 0x52, 0xf6, 0x9f, 0x44, 0x04, 0x97, 0xf6, 0x10, 0x71, 0x0d, 0x1a,
	0xb8, 0xa6, 0x8b, 0x89, 0x58, 0x6e, 0xa6, 0xbc, 0x42, 0xbe, 0xc7, 0xc6, 0xfd, 0x5d, 0x3d, 0x12,
	0xd1, 0x67, 0x35, 0xbf, 0x21, 0x15, 0xaa, 0x22, 0xa5, 0xc8, 0x61, 0xec, 0x2e, 0x55, 0x42, 0xc8,
	0xb7, 0xe1, 0xe5, 0x81, 0xb8, 0x3c, 0x10, 0x97, 0x07, 0xbb, 0x92, 0xdf, 0x99, 0x14, 0x59, 0x96,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x92, 0x02, 0xa5, 0x7d, 0x07, 0x00, 0x00,
}
