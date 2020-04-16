// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/campaign_feed.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A campaign feed.
type CampaignFeed struct {
	// Immutable. The resource name of the campaign feed.
	// Campaign feed resource names have the form:
	//
	// `customers/{customer_id}/campaignFeeds/{campaign_id}~{feed_id}
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The feed to which the CampaignFeed belongs.
	Feed *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed,proto3" json:"feed,omitempty"`
	// Immutable. The campaign to which the CampaignFeed belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Indicates which placeholder types the feed may populate under the connected
	// campaign. Required.
	PlaceholderTypes []enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,4,rep,packed,name=placeholder_types,json=placeholderTypes,proto3,enum=google.ads.googleads.v2.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_types,omitempty"`
	// Matching function associated with the CampaignFeed.
	// The matching function is used to filter the set of feed items selected.
	// Required.
	MatchingFunction *common.MatchingFunction `protobuf:"bytes,5,opt,name=matching_function,json=matchingFunction,proto3" json:"matching_function,omitempty"`
	// Output only. Status of the campaign feed.
	// This field is read-only.
	Status               enums.FeedLinkStatusEnum_FeedLinkStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.FeedLinkStatusEnum_FeedLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CampaignFeed) Reset()         { *m = CampaignFeed{} }
func (m *CampaignFeed) String() string { return proto.CompactTextString(m) }
func (*CampaignFeed) ProtoMessage()    {}
func (*CampaignFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d62b83685c9f63e, []int{0}
}

func (m *CampaignFeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignFeed.Unmarshal(m, b)
}
func (m *CampaignFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignFeed.Marshal(b, m, deterministic)
}
func (m *CampaignFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignFeed.Merge(m, src)
}
func (m *CampaignFeed) XXX_Size() int {
	return xxx_messageInfo_CampaignFeed.Size(m)
}
func (m *CampaignFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignFeed.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignFeed proto.InternalMessageInfo

func (m *CampaignFeed) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignFeed) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *CampaignFeed) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignFeed) GetPlaceholderTypes() []enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderTypes
	}
	return nil
}

func (m *CampaignFeed) GetMatchingFunction() *common.MatchingFunction {
	if m != nil {
		return m.MatchingFunction
	}
	return nil
}

func (m *CampaignFeed) GetStatus() enums.FeedLinkStatusEnum_FeedLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CampaignFeed)(nil), "google.ads.googleads.v2.resources.CampaignFeed")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/campaign_feed.proto", fileDescriptor_0d62b83685c9f63e)
}

var fileDescriptor_0d62b83685c9f63e = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0x93, 0x36, 0xfa, 0xbe, 0xa1, 0x54, 0x8d, 0x57, 0xa6, 0x2a, 0x34, 0x01, 0x45, 0x84,
	0x05, 0x33, 0xc8, 0xfc, 0x2c, 0xcc, 0x06, 0x1b, 0xd1, 0x4a, 0x08, 0x50, 0xe4, 0xa2, 0x20, 0x41,
	0x8a, 0x35, 0xb1, 0x27, 0x8e, 0x55, 0xcf, 0x8c, 0xe5, 0xb1, 0x83, 0x2a, 0xd4, 0x97, 0x61, 0xc9,
	0x86, 0xf7, 0xe0, 0x29, 0xba, 0xee, 0x23, 0xb0, 0x40, 0xc8, 0xe3, 0xb1, 0x93, 0x06, 0x85, 0x64,
	0x77, 0x73, 0xef, 0x39, 0xe7, 0xfe, 0xf8, 0x4c, 0xc0, 0xd3, 0x90, 0xf3, 0x30, 0x26, 0x08, 0x07,
	0x02, 0x95, 0x61, 0x11, 0xcd, 0x4c, 0x94, 0x12, 0xc1, 0xf3, 0xd4, 0x27, 0x02, 0xf9, 0x98, 0x26,
	0x38, 0x0a, 0x99, 0x37, 0x21, 0x24, 0x80, 0x49, 0xca, 0x33, 0xae, 0x77, 0x4b, 0x2c, 0xc4, 0x81,
	0x80, 0x35, 0x0d, 0xce, 0x4c, 0x58, 0xd3, 0xf6, 0x9f, 0xad, 0x52, 0xf6, 0x39, 0xa5, 0x9c, 0x21,
	0x8a, 0x33, 0x7f, 0x1a, 0xb1, 0xd0, 0x9b, 0xe4, 0xcc, 0xcf, 0x22, 0xce, 0x4a, 0xe9, 0xfd, 0x27,
	0xab, 0x78, 0x84, 0xe5, 0x54, 0xa0, 0x62, 0x08, 0x2f, 0x8e, 0xd8, 0x99, 0x27, 0x32, 0x9c, 0xe5,
	0x62, 0x33, 0x56, 0x12, 0x63, 0x9f, 0x4c, 0x79, 0x1c, 0x90, 0xd4, 0xcb, 0xce, 0x13, 0xa2, 0x58,
	0x87, 0x15, 0x2b, 0x89, 0xd0, 0x24, 0x22, 0x71, 0xe0, 0x8d, 0xc9, 0x14, 0xcf, 0x22, 0x9e, 0x2a,
	0xc0, 0xad, 0x05, 0x40, 0xb5, 0x9a, 0x2a, 0xdd, 0x51, 0x25, 0xf9, 0x6b, 0x9c, 0x4f, 0xd0, 0x97,
	0x14, 0x27, 0x09, 0x49, 0xab, 0x89, 0x0e, 0x16, 0xa8, 0x98, 0x31, 0x9e, 0xe1, 0x62, 0x49, 0x55,
	0xbd, 0xfb, 0x63, 0x1b, 0xec, 0xbc, 0x54, 0x87, 0x3d, 0x22, 0x24, 0xd0, 0x5d, 0x70, 0xb3, 0x6a,
	0xe0, 0x31, 0x4c, 0x89, 0xa1, 0x75, 0xb4, 0xfe, 0xff, 0xce, 0xc3, 0x4b, 0x7b, 0xfb, 0x97, 0x7d,
	0x1f, 0xf4, 0xe6, 0x57, 0x56, 0x51, 0x12, 0x09, 0xe8, 0x73, 0x8a, 0x16, 0x55, 0xdc, 0x9d, 0x4a,
	0xe3, 0x1d, 0xa6, 0x44, 0xff, 0x00, 0xb6, 0x8a, 0x73, 0x19, 0x8d, 0x8e, 0xd6, 0xbf, 0x61, 0x1e,
	0x28, 0x26, 0xac, 0x26, 0x86, 0x27, 0x59, 0x1a, 0xb1, 0x70, 0x88, 0xe3, 0x9c, 0x38, 0x3d, 0xd9,
	0xe8, 0x10, 0xdc, 0x5e, 0xd9, 0x48, 0x36, 0x90, 0x82, 0xba, 0x0f, 0xfe, 0xab, 0x5c, 0x61, 0x34,
	0x37, 0x10, 0x7f, 0x20, 0xc5, 0xef, 0x81, 0xee, 0xda, 0x2d, 0xdc, 0x5a, 0x58, 0xe7, 0xa0, 0xbd,
	0xfc, 0xd9, 0x84, 0xb1, 0xd5, 0x69, 0xf6, 0x77, 0x4d, 0x07, 0xae, 0xf2, 0x9f, 0xfc, 0xdc, 0x70,
	0x30, 0xe7, 0xbd, 0x3f, 0x4f, 0xc8, 0x2b, 0x96, 0xd3, 0xe5, 0x9c, 0xbb, 0x97, 0x5c, 0x4f, 0x08,
	0xfd, 0x14, 0xb4, 0xff, 0x32, 0xa5, 0xb1, 0x2d, 0xd7, 0x7b, 0xb4, 0xb2, 0x61, 0xe9, 0x66, 0xf8,
	0x56, 0x11, 0x8f, 0x14, 0xcf, 0xdd, 0xa3, 0x4b, 0x19, 0xfd, 0x14, 0xb4, 0x4a, 0xcb, 0x1a, 0xad,
	0x8e, 0xd6, 0xdf, 0x35, 0x5f, 0xac, 0x59, 0xa2, 0xb8, 0xf7, 0x9b, 0x88, 0x9d, 0x9d, 0x48, 0x92,
	0xdc, 0xe1, 0x7a, 0xca, 0x69, 0x5e, 0xda, 0x4d, 0x57, 0x89, 0x5a, 0x9f, 0xaf, 0xec, 0x4f, 0x1b,
	0xda, 0x44, 0x37, 0xfd, 0x5c, 0x64, 0x9c, 0x92, 0x54, 0xa0, 0xaf, 0x55, 0x78, 0x51, 0x3f, 0xf4,
	0x02, 0x52, 0x14, 0x16, 0xdf, 0xfd, 0x85, 0xf3, 0x5b, 0x03, 0x3d, 0x9f, 0x53, 0xb8, 0xf6, 0xe5,
	0x3b, 0xed, 0xc5, 0x5e, 0x83, 0xc2, 0x0f, 0x03, 0xed, 0xe3, 0x6b, 0xc5, 0x0b, 0x79, 0x8c, 0x59,
	0x08, 0x79, 0x1a, 0xa2, 0x90, 0x30, 0xe9, 0x16, 0x34, 0x1f, 0xf5, 0x1f, 0xff, 0x43, 0xcf, 0xeb,
	0xe8, 0x5b, 0xa3, 0x79, 0x6c, 0xdb, 0xdf, 0x1b, 0xdd, 0xe3, 0x52, 0xd2, 0x0e, 0x04, 0x2c, 0xc3,
	0x22, 0x1a, 0x9a, 0xd0, 0xad, 0x90, 0x3f, 0x2b, 0xcc, 0xc8, 0x0e, 0xc4, 0xa8, 0xc6, 0x8c, 0x86,
	0xe6, 0xa8, 0xc6, 0x5c, 0x35, 0x7a, 0x65, 0xc1, 0xb2, 0xec, 0x40, 0x58, 0x56, 0x8d, 0xb2, 0xac,
	0xa1, 0x69, 0x59, 0x35, 0x6e, 0xdc, 0x92, 0xc3, 0x3e, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x90,
	0x4c, 0xa3, 0xa9, 0x33, 0x05, 0x00, 0x00,
}