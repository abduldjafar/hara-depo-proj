// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/campaign_extension_setting_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for
// [CampaignExtensionSettingService.GetCampaignExtensionSetting][google.ads.googleads.v2.services.CampaignExtensionSettingService.GetCampaignExtensionSetting].
type GetCampaignExtensionSettingRequest struct {
	// Required. The resource name of the campaign extension setting to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignExtensionSettingRequest) Reset()         { *m = GetCampaignExtensionSettingRequest{} }
func (m *GetCampaignExtensionSettingRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignExtensionSettingRequest) ProtoMessage()    {}
func (*GetCampaignExtensionSettingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49140070a305e12a, []int{0}
}

func (m *GetCampaignExtensionSettingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignExtensionSettingRequest.Unmarshal(m, b)
}
func (m *GetCampaignExtensionSettingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignExtensionSettingRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignExtensionSettingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignExtensionSettingRequest.Merge(m, src)
}
func (m *GetCampaignExtensionSettingRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignExtensionSettingRequest.Size(m)
}
func (m *GetCampaignExtensionSettingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignExtensionSettingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignExtensionSettingRequest proto.InternalMessageInfo

func (m *GetCampaignExtensionSettingRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CampaignExtensionSettingService.MutateCampaignExtensionSettings][google.ads.googleads.v2.services.CampaignExtensionSettingService.MutateCampaignExtensionSettings].
type MutateCampaignExtensionSettingsRequest struct {
	// Required. The ID of the customer whose campaign extension settings are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual campaign extension
	// settings.
	Operations []*CampaignExtensionSettingOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignExtensionSettingsRequest) Reset() {
	*m = MutateCampaignExtensionSettingsRequest{}
}
func (m *MutateCampaignExtensionSettingsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignExtensionSettingsRequest) ProtoMessage()    {}
func (*MutateCampaignExtensionSettingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49140070a305e12a, []int{1}
}

func (m *MutateCampaignExtensionSettingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignExtensionSettingsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignExtensionSettingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignExtensionSettingsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignExtensionSettingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignExtensionSettingsRequest.Merge(m, src)
}
func (m *MutateCampaignExtensionSettingsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignExtensionSettingsRequest.Size(m)
}
func (m *MutateCampaignExtensionSettingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignExtensionSettingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignExtensionSettingsRequest proto.InternalMessageInfo

func (m *MutateCampaignExtensionSettingsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignExtensionSettingsRequest) GetOperations() []*CampaignExtensionSettingOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignExtensionSettingsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignExtensionSettingsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a campaign extension setting.
type CampaignExtensionSettingOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignExtensionSettingOperation_Create
	//	*CampaignExtensionSettingOperation_Update
	//	*CampaignExtensionSettingOperation_Remove
	Operation            isCampaignExtensionSettingOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *CampaignExtensionSettingOperation) Reset()         { *m = CampaignExtensionSettingOperation{} }
func (m *CampaignExtensionSettingOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignExtensionSettingOperation) ProtoMessage()    {}
func (*CampaignExtensionSettingOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_49140070a305e12a, []int{2}
}

func (m *CampaignExtensionSettingOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignExtensionSettingOperation.Unmarshal(m, b)
}
func (m *CampaignExtensionSettingOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignExtensionSettingOperation.Marshal(b, m, deterministic)
}
func (m *CampaignExtensionSettingOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignExtensionSettingOperation.Merge(m, src)
}
func (m *CampaignExtensionSettingOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignExtensionSettingOperation.Size(m)
}
func (m *CampaignExtensionSettingOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignExtensionSettingOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignExtensionSettingOperation proto.InternalMessageInfo

func (m *CampaignExtensionSettingOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignExtensionSettingOperation_Operation interface {
	isCampaignExtensionSettingOperation_Operation()
}

type CampaignExtensionSettingOperation_Create struct {
	Create *resources.CampaignExtensionSetting `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignExtensionSettingOperation_Update struct {
	Update *resources.CampaignExtensionSetting `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignExtensionSettingOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignExtensionSettingOperation_Create) isCampaignExtensionSettingOperation_Operation() {}

func (*CampaignExtensionSettingOperation_Update) isCampaignExtensionSettingOperation_Operation() {}

func (*CampaignExtensionSettingOperation_Remove) isCampaignExtensionSettingOperation_Operation() {}

func (m *CampaignExtensionSettingOperation) GetOperation() isCampaignExtensionSettingOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignExtensionSettingOperation) GetCreate() *resources.CampaignExtensionSetting {
	if x, ok := m.GetOperation().(*CampaignExtensionSettingOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignExtensionSettingOperation) GetUpdate() *resources.CampaignExtensionSetting {
	if x, ok := m.GetOperation().(*CampaignExtensionSettingOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignExtensionSettingOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignExtensionSettingOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignExtensionSettingOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignExtensionSettingOperation_Create)(nil),
		(*CampaignExtensionSettingOperation_Update)(nil),
		(*CampaignExtensionSettingOperation_Remove)(nil),
	}
}

// Response message for a campaign extension setting mutate.
type MutateCampaignExtensionSettingsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignExtensionSettingResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *MutateCampaignExtensionSettingsResponse) Reset() {
	*m = MutateCampaignExtensionSettingsResponse{}
}
func (m *MutateCampaignExtensionSettingsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignExtensionSettingsResponse) ProtoMessage()    {}
func (*MutateCampaignExtensionSettingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49140070a305e12a, []int{3}
}

func (m *MutateCampaignExtensionSettingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignExtensionSettingsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignExtensionSettingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignExtensionSettingsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignExtensionSettingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignExtensionSettingsResponse.Merge(m, src)
}
func (m *MutateCampaignExtensionSettingsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignExtensionSettingsResponse.Size(m)
}
func (m *MutateCampaignExtensionSettingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignExtensionSettingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignExtensionSettingsResponse proto.InternalMessageInfo

func (m *MutateCampaignExtensionSettingsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignExtensionSettingsResponse) GetResults() []*MutateCampaignExtensionSettingResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the campaign extension setting mutate.
type MutateCampaignExtensionSettingResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignExtensionSettingResult) Reset()         { *m = MutateCampaignExtensionSettingResult{} }
func (m *MutateCampaignExtensionSettingResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignExtensionSettingResult) ProtoMessage()    {}
func (*MutateCampaignExtensionSettingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_49140070a305e12a, []int{4}
}

func (m *MutateCampaignExtensionSettingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignExtensionSettingResult.Unmarshal(m, b)
}
func (m *MutateCampaignExtensionSettingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignExtensionSettingResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignExtensionSettingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignExtensionSettingResult.Merge(m, src)
}
func (m *MutateCampaignExtensionSettingResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignExtensionSettingResult.Size(m)
}
func (m *MutateCampaignExtensionSettingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignExtensionSettingResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignExtensionSettingResult proto.InternalMessageInfo

func (m *MutateCampaignExtensionSettingResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignExtensionSettingRequest)(nil), "google.ads.googleads.v2.services.GetCampaignExtensionSettingRequest")
	proto.RegisterType((*MutateCampaignExtensionSettingsRequest)(nil), "google.ads.googleads.v2.services.MutateCampaignExtensionSettingsRequest")
	proto.RegisterType((*CampaignExtensionSettingOperation)(nil), "google.ads.googleads.v2.services.CampaignExtensionSettingOperation")
	proto.RegisterType((*MutateCampaignExtensionSettingsResponse)(nil), "google.ads.googleads.v2.services.MutateCampaignExtensionSettingsResponse")
	proto.RegisterType((*MutateCampaignExtensionSettingResult)(nil), "google.ads.googleads.v2.services.MutateCampaignExtensionSettingResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/campaign_extension_setting_service.proto", fileDescriptor_49140070a305e12a)
}

var fileDescriptor_49140070a305e12a = []byte{
	// 779 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xbf, 0x6f, 0xe3, 0x36,
	0x14, 0xae, 0xe5, 0x20, 0x6d, 0xe8, 0xa4, 0x05, 0x58, 0xb4, 0x35, 0x9c, 0x02, 0x71, 0x55, 0xa3,
	0x31, 0x8c, 0x42, 0x02, 0xd4, 0x4d, 0x46, 0x0a, 0xc8, 0x69, 0x9c, 0x04, 0x45, 0x7e, 0x40, 0x46,
	0x33, 0x14, 0x06, 0x54, 0x5a, 0x62, 0x14, 0x21, 0x92, 0xa8, 0x92, 0x94, 0xd1, 0x20, 0xc8, 0xd2,
	0xa1, 0x4b, 0xc7, 0x6e, 0x1d, 0x3b, 0x76, 0xbf, 0x7f, 0x22, 0x6b, 0xb6, 0x4c, 0x37, 0xdc, 0x74,
	0xfb, 0xed, 0x07, 0x89, 0xa2, 0x7f, 0xe4, 0x4e, 0xd1, 0x01, 0xc9, 0xf6, 0xcc, 0xf7, 0xe9, 0x7b,
	0x1f, 0xdf, 0xfb, 0xf8, 0x0c, 0x0e, 0x7d, 0x42, 0xfc, 0x10, 0xeb, 0xc8, 0x63, 0xba, 0x08, 0xb3,
	0x68, 0x6a, 0xe8, 0x0c, 0xd3, 0x69, 0xe0, 0x62, 0xa6, 0xbb, 0x28, 0x4a, 0x50, 0xe0, 0xc7, 0x0e,
	0xfe, 0x83, 0xe3, 0x98, 0x05, 0x24, 0x76, 0x18, 0xe6, 0x3c, 0x88, 0x7d, 0xa7, 0xc0, 0x68, 0x09,
	0x25, 0x9c, 0xc0, 0xb6, 0xf8, 0x5e, 0x43, 0x1e, 0xd3, 0x66, 0x54, 0xda, 0xd4, 0xd0, 0x24, 0x55,
	0x6b, 0x50, 0x56, 0x8c, 0x62, 0x46, 0x52, 0xfa, 0x78, 0x35, 0x51, 0xa5, 0xf5, 0xb5, 0xe4, 0x48,
	0x02, 0x1d, 0xc5, 0x31, 0xe1, 0x88, 0x07, 0x24, 0x66, 0x45, 0xf6, 0xab, 0x85, 0xac, 0x1b, 0x06,
	0x38, 0xe6, 0x45, 0x62, 0x6b, 0x21, 0x71, 0x1e, 0xe0, 0xd0, 0x73, 0x26, 0xf8, 0x02, 0x4d, 0x03,
	0x42, 0x0b, 0x40, 0xa1, 0x5e, 0xcf, 0x7f, 0x4d, 0xd2, 0xf3, 0x02, 0x15, 0x21, 0x76, 0xf9, 0x80,
	0x9b, 0x26, 0xae, 0xce, 0x38, 0xe2, 0x69, 0x51, 0x54, 0x3d, 0x06, 0xea, 0x3e, 0xe6, 0xbb, 0x85,
	0xf2, 0x3d, 0x29, 0x7c, 0x24, 0x74, 0xdb, 0xf8, 0xf7, 0x14, 0x33, 0x0e, 0xbb, 0x60, 0x43, 0x5e,
	0xd3, 0x89, 0x51, 0x84, 0x9b, 0xb5, 0x76, 0xad, 0xbb, 0x36, 0xa8, 0xbf, 0xb4, 0x14, 0x7b, 0x5d,
	0x66, 0x8e, 0x51, 0x84, 0xd5, 0xbf, 0x14, 0xf0, 0xdd, 0x51, 0xca, 0x11, 0xc7, 0x65, 0x9c, 0x4c,
	0x92, 0x76, 0x40, 0xc3, 0x4d, 0x19, 0x27, 0x11, 0xa6, 0x4e, 0xe0, 0x2d, 0x52, 0x02, 0x79, 0x7e,
	0xe8, 0xc1, 0x0b, 0x00, 0x48, 0x82, 0xa9, 0xe8, 0x54, 0x53, 0x69, 0xd7, 0xbb, 0x0d, 0x63, 0x57,
	0xab, 0x1a, 0x97, 0x56, 0x56, 0xfd, 0x44, 0x72, 0x15, 0x95, 0xe6, 0xdc, 0x70, 0x1b, 0x7c, 0x96,
	0x20, 0xca, 0x03, 0x14, 0x3a, 0xe7, 0x28, 0x08, 0x53, 0x8a, 0x9b, 0xf5, 0x76, 0xad, 0xfb, 0x89,
	0xfd, 0x69, 0x71, 0x3c, 0x14, 0xa7, 0xf0, 0x5b, 0xb0, 0x31, 0x45, 0x61, 0xe0, 0x21, 0x8e, 0x1d,
	0x12, 0x87, 0x57, 0xcd, 0x95, 0x1c, 0xb6, 0x2e, 0x0f, 0x4f, 0xe2, 0xf0, 0x4a, 0x7d, 0xa1, 0x80,
	0x6f, 0x2a, 0x45, 0xc0, 0x3e, 0x68, 0xa4, 0x49, 0x4e, 0x94, 0x0d, 0x2b, 0x27, 0x6a, 0x18, 0x2d,
	0x79, 0x3d, 0x39, 0x4f, 0x6d, 0x98, 0xcd, 0xf3, 0x08, 0xb1, 0x4b, 0x1b, 0x08, 0x78, 0x16, 0xc3,
	0x5f, 0xc0, 0xaa, 0x4b, 0x31, 0xe2, 0x62, 0x1c, 0x0d, 0xa3, 0x5f, 0xda, 0x96, 0x99, 0x47, 0x4b,
	0xfb, 0x72, 0xf0, 0x91, 0x5d, 0x90, 0x65, 0xb4, 0xa2, 0x48, 0x53, 0x79, 0x16, 0x5a, 0x41, 0x06,
	0x9b, 0x60, 0x95, 0xe2, 0x88, 0x4c, 0x45, 0x57, 0xd7, 0xb2, 0x8c, 0xf8, 0x3d, 0x68, 0x80, 0xb5,
	0xd9, 0x18, 0xd4, 0xbb, 0x1a, 0xd8, 0xae, 0x34, 0x10, 0x4b, 0x48, 0xcc, 0x30, 0x1c, 0x82, 0x2f,
	0x1e, 0x4c, 0xcc, 0xc1, 0x94, 0x12, 0x9a, 0x57, 0x68, 0x18, 0x50, 0x0a, 0xa7, 0x89, 0xab, 0x8d,
	0x72, 0xd7, 0xdb, 0x9f, 0x2f, 0xcf, 0x72, 0x2f, 0x83, 0xc3, 0xdf, 0xc0, 0xc7, 0x14, 0xb3, 0x34,
	0xe4, 0xd2, 0x60, 0xc3, 0x6a, 0x83, 0x3d, 0xae, 0xd1, 0xce, 0xe9, 0x6c, 0x49, 0xab, 0xfe, 0x0c,
	0x3a, 0x1f, 0xf2, 0x41, 0x66, 0xad, 0xf7, 0x3c, 0xb4, 0xe5, 0x37, 0x66, 0xdc, 0xad, 0x80, 0xad,
	0x32, 0x9e, 0x91, 0xd0, 0x07, 0xdf, 0xd4, 0xc0, 0xe6, 0x23, 0x0f, 0x1b, 0xfe, 0x54, 0x7d, 0xc3,
	0xea, 0xbd, 0xd0, 0x7a, 0x8a, 0x35, 0xd4, 0xd1, 0xbd, 0xb5, 0x7c, 0xd9, 0x3f, 0xef, 0x5e, 0xfd,
	0xa3, 0xec, 0xc0, 0x7e, 0xb6, 0x55, 0xaf, 0x97, 0x32, 0x3b, 0x72, 0x21, 0x30, 0xbd, 0x37, 0x5b,
	0xb3, 0xef, 0xf8, 0x42, 0xef, 0xdd, 0xc0, 0x7f, 0x15, 0xb0, 0x55, 0x61, 0x1f, 0x78, 0xf0, 0xd4,
	0xe9, 0xca, 0x15, 0xd6, 0x3a, 0x7c, 0x06, 0x26, 0xe1, 0x65, 0x75, 0x72, 0x6f, 0x7d, 0xb9, 0xb0,
	0x0e, 0xbf, 0x9f, 0x2f, 0xa6, 0xbc, 0x2d, 0xbb, 0xea, 0x8f, 0x59, 0x5b, 0xe6, 0x7d, 0xb8, 0x5e,
	0x00, 0xef, 0xf4, 0x6e, 0xca, 0xbb, 0x62, 0x46, 0xb9, 0x02, 0xb3, 0xd6, 0x6b, 0x6d, 0xde, 0x5a,
	0xcd, 0xb9, 0xca, 0x22, 0x4a, 0x02, 0xa6, 0xb9, 0x24, 0x1a, 0xfc, 0xad, 0x80, 0x8e, 0x4b, 0xa2,
	0xca, 0x1b, 0x0d, 0x3a, 0x15, 0xde, 0x3b, 0xcd, 0xb6, 0xd6, 0x69, 0xed, 0xd7, 0x83, 0x82, 0xc9,
	0x27, 0x21, 0x8a, 0x7d, 0x8d, 0x50, 0x5f, 0xf7, 0x71, 0x9c, 0xef, 0x34, 0x7d, 0x5e, 0xbb, 0xfc,
	0xdf, 0xbb, 0x2f, 0x83, 0xff, 0x94, 0xfa, 0xbe, 0x65, 0xfd, 0xaf, 0xb4, 0xf7, 0x05, 0xa1, 0xe5,
	0x31, 0x4d, 0x84, 0x59, 0x74, 0x66, 0x68, 0x45, 0x61, 0x76, 0x2b, 0x21, 0x63, 0xcb, 0x63, 0xe3,
	0x19, 0x64, 0x7c, 0x66, 0x8c, 0x25, 0xe4, 0xb5, 0xd2, 0x11, 0xe7, 0xa6, 0x69, 0x79, 0xcc, 0x34,
	0x67, 0x20, 0xd3, 0x3c, 0x33, 0x4c, 0x53, 0xc2, 0x26, 0xab, 0xb9, 0xce, 0x1f, 0xde, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xdb, 0xd5, 0xd3, 0x7f, 0x64, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignExtensionSettingServiceClient is the client API for CampaignExtensionSettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignExtensionSettingServiceClient interface {
	// Returns the requested campaign extension setting in full detail.
	GetCampaignExtensionSetting(ctx context.Context, in *GetCampaignExtensionSettingRequest, opts ...grpc.CallOption) (*resources.CampaignExtensionSetting, error)
	// Creates, updates, or removes campaign extension settings. Operation
	// statuses are returned.
	MutateCampaignExtensionSettings(ctx context.Context, in *MutateCampaignExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateCampaignExtensionSettingsResponse, error)
}

type campaignExtensionSettingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignExtensionSettingServiceClient(cc grpc.ClientConnInterface) CampaignExtensionSettingServiceClient {
	return &campaignExtensionSettingServiceClient{cc}
}

func (c *campaignExtensionSettingServiceClient) GetCampaignExtensionSetting(ctx context.Context, in *GetCampaignExtensionSettingRequest, opts ...grpc.CallOption) (*resources.CampaignExtensionSetting, error) {
	out := new(resources.CampaignExtensionSetting)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CampaignExtensionSettingService/GetCampaignExtensionSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignExtensionSettingServiceClient) MutateCampaignExtensionSettings(ctx context.Context, in *MutateCampaignExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateCampaignExtensionSettingsResponse, error) {
	out := new(MutateCampaignExtensionSettingsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CampaignExtensionSettingService/MutateCampaignExtensionSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignExtensionSettingServiceServer is the server API for CampaignExtensionSettingService service.
type CampaignExtensionSettingServiceServer interface {
	// Returns the requested campaign extension setting in full detail.
	GetCampaignExtensionSetting(context.Context, *GetCampaignExtensionSettingRequest) (*resources.CampaignExtensionSetting, error)
	// Creates, updates, or removes campaign extension settings. Operation
	// statuses are returned.
	MutateCampaignExtensionSettings(context.Context, *MutateCampaignExtensionSettingsRequest) (*MutateCampaignExtensionSettingsResponse, error)
}

// UnimplementedCampaignExtensionSettingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignExtensionSettingServiceServer struct {
}

func (*UnimplementedCampaignExtensionSettingServiceServer) GetCampaignExtensionSetting(ctx context.Context, req *GetCampaignExtensionSettingRequest) (*resources.CampaignExtensionSetting, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCampaignExtensionSetting not implemented")
}
func (*UnimplementedCampaignExtensionSettingServiceServer) MutateCampaignExtensionSettings(ctx context.Context, req *MutateCampaignExtensionSettingsRequest) (*MutateCampaignExtensionSettingsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaignExtensionSettings not implemented")
}

func RegisterCampaignExtensionSettingServiceServer(s *grpc.Server, srv CampaignExtensionSettingServiceServer) {
	s.RegisterService(&_CampaignExtensionSettingService_serviceDesc, srv)
}

func _CampaignExtensionSettingService_GetCampaignExtensionSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignExtensionSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignExtensionSettingServiceServer).GetCampaignExtensionSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CampaignExtensionSettingService/GetCampaignExtensionSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignExtensionSettingServiceServer).GetCampaignExtensionSetting(ctx, req.(*GetCampaignExtensionSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignExtensionSettingService_MutateCampaignExtensionSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignExtensionSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignExtensionSettingServiceServer).MutateCampaignExtensionSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CampaignExtensionSettingService/MutateCampaignExtensionSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignExtensionSettingServiceServer).MutateCampaignExtensionSettings(ctx, req.(*MutateCampaignExtensionSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignExtensionSettingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.CampaignExtensionSettingService",
	HandlerType: (*CampaignExtensionSettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignExtensionSetting",
			Handler:    _CampaignExtensionSettingService_GetCampaignExtensionSetting_Handler,
		},
		{
			MethodName: "MutateCampaignExtensionSettings",
			Handler:    _CampaignExtensionSettingService_MutateCampaignExtensionSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/campaign_extension_setting_service.proto",
}
