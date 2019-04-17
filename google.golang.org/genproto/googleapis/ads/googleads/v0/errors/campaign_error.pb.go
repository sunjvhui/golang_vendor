// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/campaign_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible campaign errors.
type CampaignErrorEnum_CampaignError int32

const (
	// Enum unspecified.
	CampaignErrorEnum_UNSPECIFIED CampaignErrorEnum_CampaignError = 0
	// The received error code is not known in this version.
	CampaignErrorEnum_UNKNOWN CampaignErrorEnum_CampaignError = 1
	// Cannot target content network.
	CampaignErrorEnum_CANNOT_TARGET_CONTENT_NETWORK CampaignErrorEnum_CampaignError = 3
	// Cannot target search network.
	CampaignErrorEnum_CANNOT_TARGET_SEARCH_NETWORK CampaignErrorEnum_CampaignError = 4
	// Cannot cover search network without google search network.
	CampaignErrorEnum_CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH CampaignErrorEnum_CampaignError = 5
	// Cannot target Google Search network for a CPM campaign.
	CampaignErrorEnum_CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN CampaignErrorEnum_CampaignError = 6
	// Must target at least one network.
	CampaignErrorEnum_CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK CampaignErrorEnum_CampaignError = 7
	// Only some Google partners are allowed to target partner search network.
	CampaignErrorEnum_CANNOT_TARGET_PARTNER_SEARCH_NETWORK CampaignErrorEnum_CampaignError = 8
	// Cannot target content network only as campaign has criteria-level bidding
	// strategy.
	CampaignErrorEnum_CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY CampaignErrorEnum_CampaignError = 9
	// Cannot modify the start or end date such that the campaign duration would
	// not contain the durations of all runnable trials.
	CampaignErrorEnum_CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS CampaignErrorEnum_CampaignError = 10
	// Cannot modify dates, budget or campaign name of a trial campaign.
	CampaignErrorEnum_CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN CampaignErrorEnum_CampaignError = 11
	// Trying to modify the name of an active or paused campaign, where the name
	// is already assigned to another active or paused campaign.
	CampaignErrorEnum_DUPLICATE_CAMPAIGN_NAME CampaignErrorEnum_CampaignError = 12
	// Two fields are in conflicting modes.
	CampaignErrorEnum_INCOMPATIBLE_CAMPAIGN_FIELD CampaignErrorEnum_CampaignError = 13
	// Campaign name cannot be used.
	CampaignErrorEnum_INVALID_CAMPAIGN_NAME CampaignErrorEnum_CampaignError = 14
	// Given status is invalid.
	CampaignErrorEnum_INVALID_AD_SERVING_OPTIMIZATION_STATUS CampaignErrorEnum_CampaignError = 15
	// Error in the campaign level tracking url.
	CampaignErrorEnum_INVALID_TRACKING_URL CampaignErrorEnum_CampaignError = 16
	// Cannot set both tracking url template and tracking setting. An user has
	// to clear legacy tracking setting in order to add tracking url template.
	CampaignErrorEnum_CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING CampaignErrorEnum_CampaignError = 17
	// The maximum number of impressions for Frequency Cap should be an integer
	// greater than 0.
	CampaignErrorEnum_MAX_IMPRESSIONS_NOT_IN_RANGE CampaignErrorEnum_CampaignError = 18
	// Only the Day, Week and Month time units are supported.
	CampaignErrorEnum_TIME_UNIT_NOT_SUPPORTED CampaignErrorEnum_CampaignError = 19
	// Operation not allowed on a campaign whose serving status has ended
	CampaignErrorEnum_INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED CampaignErrorEnum_CampaignError = 20
	// This budget is exclusively linked to a Campaign that is using experiments
	// so it cannot be shared.
	CampaignErrorEnum_BUDGET_CANNOT_BE_SHARED CampaignErrorEnum_CampaignError = 21
	// Campaigns using experiments cannot use a shared budget.
	CampaignErrorEnum_CAMPAIGN_CANNOT_USE_SHARED_BUDGET CampaignErrorEnum_CampaignError = 22
	// A different budget cannot be assigned to a campaign when there are
	// running or scheduled trials.
	CampaignErrorEnum_CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS CampaignErrorEnum_CampaignError = 23
	// No link found between the campaign and the label.
	CampaignErrorEnum_CAMPAIGN_LABEL_DOES_NOT_EXIST CampaignErrorEnum_CampaignError = 24
	// The label has already been attached to the campaign.
	CampaignErrorEnum_CAMPAIGN_LABEL_ALREADY_EXISTS CampaignErrorEnum_CampaignError = 25
	// A ShoppingSetting was not found when creating a shopping campaign.
	CampaignErrorEnum_MISSING_SHOPPING_SETTING CampaignErrorEnum_CampaignError = 26
	// The country in shopping setting is not an allowed country.
	CampaignErrorEnum_INVALID_SHOPPING_SALES_COUNTRY CampaignErrorEnum_CampaignError = 27
	// A Campaign with channel sub type UNIVERSAL_APP_CAMPAIGN must have a
	// UniversalAppCampaignSetting specified.
	CampaignErrorEnum_MISSING_UNIVERSAL_APP_CAMPAIGN_SETTING CampaignErrorEnum_CampaignError = 30
	// The requested channel type is not available according to the customer's
	// account setting.
	CampaignErrorEnum_ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE CampaignErrorEnum_CampaignError = 31
	// The AdvertisingChannelSubType is not a valid subtype of the primary
	// channel type.
	CampaignErrorEnum_INVALID_ADVERTISING_CHANNEL_SUB_TYPE CampaignErrorEnum_CampaignError = 32
	// At least one conversion must be selected.
	CampaignErrorEnum_AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED CampaignErrorEnum_CampaignError = 33
	// Setting ad rotation mode for a campaign is not allowed. Ad rotation mode
	// at campaign is deprecated.
	CampaignErrorEnum_CANNOT_SET_AD_ROTATION_MODE CampaignErrorEnum_CampaignError = 34
	// Trying to change start date on a campaign that has started.
	CampaignErrorEnum_CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED CampaignErrorEnum_CampaignError = 35
	// Trying to modify a date into the past.
	CampaignErrorEnum_CANNOT_SET_DATE_TO_PAST CampaignErrorEnum_CampaignError = 36
	// Hotel center id in the hotel setting does not match any customer links.
	CampaignErrorEnum_MISSING_HOTEL_CUSTOMER_LINK CampaignErrorEnum_CampaignError = 37
	// Hotel center id in the hotel setting must match an active customer link.
	CampaignErrorEnum_INVALID_HOTEL_CUSTOMER_LINK CampaignErrorEnum_CampaignError = 38
	// Hotel setting was not found when creating a hotel ads campaign.
	CampaignErrorEnum_MISSING_HOTEL_SETTING CampaignErrorEnum_CampaignError = 39
	// A Campaign cannot use shared campaign budgets and be part of a campaign
	// group.
	CampaignErrorEnum_CANNOT_USE_SHARED_CAMPAIGN_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP CampaignErrorEnum_CampaignError = 40
)

var CampaignErrorEnum_CampaignError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "CANNOT_TARGET_CONTENT_NETWORK",
	4:  "CANNOT_TARGET_SEARCH_NETWORK",
	5:  "CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH",
	6:  "CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN",
	7:  "CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK",
	8:  "CANNOT_TARGET_PARTNER_SEARCH_NETWORK",
	9:  "CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY",
	10: "CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS",
	11: "CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN",
	12: "DUPLICATE_CAMPAIGN_NAME",
	13: "INCOMPATIBLE_CAMPAIGN_FIELD",
	14: "INVALID_CAMPAIGN_NAME",
	15: "INVALID_AD_SERVING_OPTIMIZATION_STATUS",
	16: "INVALID_TRACKING_URL",
	17: "CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING",
	18: "MAX_IMPRESSIONS_NOT_IN_RANGE",
	19: "TIME_UNIT_NOT_SUPPORTED",
	20: "INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED",
	21: "BUDGET_CANNOT_BE_SHARED",
	22: "CAMPAIGN_CANNOT_USE_SHARED_BUDGET",
	23: "CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS",
	24: "CAMPAIGN_LABEL_DOES_NOT_EXIST",
	25: "CAMPAIGN_LABEL_ALREADY_EXISTS",
	26: "MISSING_SHOPPING_SETTING",
	27: "INVALID_SHOPPING_SALES_COUNTRY",
	30: "MISSING_UNIVERSAL_APP_CAMPAIGN_SETTING",
	31: "ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE",
	32: "INVALID_ADVERTISING_CHANNEL_SUB_TYPE",
	33: "AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED",
	34: "CANNOT_SET_AD_ROTATION_MODE",
	35: "CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED",
	36: "CANNOT_SET_DATE_TO_PAST",
	37: "MISSING_HOTEL_CUSTOMER_LINK",
	38: "INVALID_HOTEL_CUSTOMER_LINK",
	39: "MISSING_HOTEL_SETTING",
	40: "CANNOT_USE_SHARED_CAMPAIGN_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP",
}
var CampaignErrorEnum_CampaignError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"CANNOT_TARGET_CONTENT_NETWORK": 3,
	"CANNOT_TARGET_SEARCH_NETWORK":  4,
	"CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH":                      5,
	"CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN":                            6,
	"CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK":                               7,
	"CANNOT_TARGET_PARTNER_SEARCH_NETWORK":                                    8,
	"CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY": 9,
	"CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS":                      10,
	"CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN":                                        11,
	"DUPLICATE_CAMPAIGN_NAME":                                                 12,
	"INCOMPATIBLE_CAMPAIGN_FIELD":                                             13,
	"INVALID_CAMPAIGN_NAME":                                                   14,
	"INVALID_AD_SERVING_OPTIMIZATION_STATUS":                                  15,
	"INVALID_TRACKING_URL":                                                    16,
	"CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING":              17,
	"MAX_IMPRESSIONS_NOT_IN_RANGE":                                            18,
	"TIME_UNIT_NOT_SUPPORTED":                                                 19,
	"INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED":                           20,
	"BUDGET_CANNOT_BE_SHARED":                                                 21,
	"CAMPAIGN_CANNOT_USE_SHARED_BUDGET":                                       22,
	"CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS":                            23,
	"CAMPAIGN_LABEL_DOES_NOT_EXIST":                                           24,
	"CAMPAIGN_LABEL_ALREADY_EXISTS":                                           25,
	"MISSING_SHOPPING_SETTING":                                                26,
	"INVALID_SHOPPING_SALES_COUNTRY":                                          27,
	"MISSING_UNIVERSAL_APP_CAMPAIGN_SETTING":                                  30,
	"ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE":                 31,
	"INVALID_ADVERTISING_CHANNEL_SUB_TYPE":                                    32,
	"AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED":                                33,
	"CANNOT_SET_AD_ROTATION_MODE":                                             34,
	"CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED":                             35,
	"CANNOT_SET_DATE_TO_PAST":                                                 36,
	"MISSING_HOTEL_CUSTOMER_LINK":                                             37,
	"INVALID_HOTEL_CUSTOMER_LINK":                                             38,
	"MISSING_HOTEL_SETTING":                                                   39,
	"CANNOT_USE_SHARED_CAMPAIGN_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP":          40,
}

func (x CampaignErrorEnum_CampaignError) String() string {
	return proto.EnumName(CampaignErrorEnum_CampaignError_name, int32(x))
}
func (CampaignErrorEnum_CampaignError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_campaign_error_67c91d16be8f048d, []int{0, 0}
}

// Container for enum describing possible campaign errors.
type CampaignErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignErrorEnum) Reset()         { *m = CampaignErrorEnum{} }
func (m *CampaignErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignErrorEnum) ProtoMessage()    {}
func (*CampaignErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_error_67c91d16be8f048d, []int{0}
}
func (m *CampaignErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignErrorEnum.Unmarshal(m, b)
}
func (m *CampaignErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignErrorEnum.Marshal(b, m, deterministic)
}
func (dst *CampaignErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignErrorEnum.Merge(dst, src)
}
func (m *CampaignErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignErrorEnum.Size(m)
}
func (m *CampaignErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CampaignErrorEnum)(nil), "google.ads.googleads.v0.errors.CampaignErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.CampaignErrorEnum_CampaignError", CampaignErrorEnum_CampaignError_name, CampaignErrorEnum_CampaignError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/campaign_error.proto", fileDescriptor_campaign_error_67c91d16be8f048d)
}

var fileDescriptor_campaign_error_67c91d16be8f048d = []byte{
	// 937 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0x6f, 0x6f, 0x1c, 0xb5,
	0x13, 0xfe, 0xa5, 0xfd, 0xd1, 0x82, 0x93, 0x52, 0xd7, 0xb4, 0x34, 0x25, 0x25, 0x4d, 0x42, 0x5a,
	0x02, 0xb4, 0x77, 0x81, 0x4a, 0x20, 0x51, 0xa9, 0xd2, 0xdc, 0x7a, 0x6e, 0xcf, 0x8a, 0xd7, 0x5e,
	0xd9, 0xde, 0x4b, 0xaf, 0x8a, 0x34, 0x0a, 0x4d, 0x74, 0xaa, 0xd4, 0xe4, 0xa2, 0x1c, 0xf4, 0x35,
	0x9f, 0x82, 0x0f, 0xc0, 0x4b, 0x3e, 0x09, 0xe2, 0x53, 0x21, 0xef, 0xbf, 0xcb, 0xb6, 0x25, 0xaf,
	0xce, 0xeb, 0xe7, 0x99, 0xf1, 0xcc, 0xe3, 0x67, 0xce, 0xec, 0xe9, 0x74, 0x36, 0x9b, 0xbe, 0x39,
	0xee, 0x1f, 0x1e, 0xcd, 0xfb, 0xd5, 0x32, 0xae, 0xde, 0xee, 0xf6, 0x8f, 0xcf, 0xcf, 0x67, 0xe7,
	0xf3, 0xfe, 0xab, 0xc3, 0x93, 0xb3, 0xc3, 0xd7, 0xd3, 0x53, 0x2a, 0xbf, 0x7b, 0x67, 0xe7, 0xb3,
	0x5f, 0x67, 0x62, 0xbd, 0x62, 0xf6, 0x0e, 0x8f, 0xe6, 0xbd, 0x36, 0xa8, 0xf7, 0x76, 0xb7, 0x57,
	0x05, 0x6d, 0xfd, 0xb1, 0xc2, 0x6e, 0x25, 0x75, 0x20, 0xc6, 0x2d, 0x3c, 0xfd, 0xed, 0x64, 0xeb,
	0xf7, 0x15, 0x76, 0xa3, 0xb3, 0x2b, 0x6e, 0xb2, 0xe5, 0xc2, 0xf8, 0x1c, 0x13, 0x35, 0x54, 0x28,
	0xf9, 0xff, 0xc4, 0x32, 0xbb, 0x5e, 0x98, 0x3d, 0x63, 0xf7, 0x0d, 0x5f, 0x12, 0x9b, 0xec, 0xcb,
	0x04, 0x8c, 0xb1, 0x81, 0x02, 0xb8, 0x14, 0x03, 0x25, 0xd6, 0x04, 0x34, 0x81, 0x0c, 0x86, 0x7d,
	0xeb, 0xf6, 0xf8, 0x55, 0xb1, 0xc1, 0xee, 0x77, 0x29, 0x1e, 0xc1, 0x25, 0xa3, 0x96, 0xf1, 0x7f,
	0xf1, 0x23, 0xfb, 0xe1, 0x32, 0x06, 0xed, 0xab, 0x30, 0xb2, 0x45, 0xa0, 0xd4, 0xda, 0x54, 0x63,
	0x8d, 0xf2, 0x8f, 0xc4, 0x2e, 0x7b, 0xdc, 0x8d, 0xeb, 0x10, 0x68, 0x68, 0x1d, 0x25, 0x79, 0x46,
	0x09, 0x64, 0x39, 0xa8, 0xd4, 0xf0, 0x6b, 0xe2, 0x09, 0xfb, 0xa6, 0xf9, 0xa2, 0xac, 0xf0, 0x6d,
	0x20, 0x04, 0xd2, 0x08, 0x3e, 0x90, 0x35, 0xd8, 0x16, 0x76, 0x5d, 0xec, 0xb0, 0xed, 0xee, 0x01,
	0x39, 0xb8, 0x60, 0xd0, 0xbd, 0xdb, 0xc2, 0xc7, 0x62, 0x8f, 0xa5, 0x97, 0xea, 0x40, 0xd6, 0xe8,
	0x49, 0xd9, 0x08, 0x25, 0x4e, 0x05, 0x74, 0x0a, 0x48, 0xe3, 0x18, 0x35, 0x0d, 0x94, 0x94, 0xca,
	0xa4, 0xe4, 0x83, 0x83, 0x80, 0xe9, 0x84, 0x7f, 0x52, 0xe9, 0x51, 0x57, 0x29, 0x0b, 0x07, 0x41,
	0xd9, 0xba, 0xdc, 0x98, 0x15, 0x94, 0x21, 0xd0, 0x9a, 0x5c, 0x61, 0x0c, 0x0c, 0x34, 0x52, 0x70,
	0x0a, 0xb4, 0xe7, 0x4c, 0x6c, 0xb3, 0x8d, 0xba, 0x88, 0xcc, 0x4a, 0x35, 0x9c, 0x94, 0x0a, 0x94,
	0xe8, 0x42, 0x83, 0x65, 0xb1, 0xc6, 0xee, 0xca, 0x22, 0xd7, 0x2a, 0x81, 0x80, 0xed, 0x3e, 0x19,
	0xc8, 0x90, 0xaf, 0x88, 0x07, 0x6c, 0x4d, 0x99, 0xc4, 0x66, 0x39, 0x04, 0x15, 0x73, 0xb7, 0xf8,
	0x50, 0xa1, 0x96, 0xfc, 0x86, 0xb8, 0xc7, 0xee, 0x28, 0x33, 0x06, 0xad, 0xe4, 0x3b, 0xb1, 0x9f,
	0x8a, 0x6f, 0xd9, 0xa3, 0x06, 0x02, 0x49, 0x1e, 0xdd, 0x38, 0xf6, 0x65, 0xf3, 0xa0, 0x32, 0xf5,
	0xb2, 0x6a, 0xc2, 0x07, 0x08, 0x85, 0xe7, 0x37, 0xc5, 0x2a, 0xbb, 0xdd, 0x70, 0x83, 0x83, 0x64,
	0x2f, 0x32, 0x0b, 0xa7, 0x39, 0x17, 0xcf, 0xd9, 0xcf, 0x75, 0x13, 0x1e, 0x03, 0x0d, 0x6c, 0x18,
	0x75, 0x18, 0x14, 0x30, 0xcb, 0x75, 0x2c, 0x1d, 0xcc, 0x85, 0x58, 0x8f, 0x21, 0x28, 0x93, 0xf2,
	0x5b, 0xd1, 0x6e, 0x19, 0xbc, 0x20, 0x95, 0xe5, 0x0e, 0xbd, 0x57, 0xd6, 0x78, 0x8a, 0xc9, 0x94,
	0x21, 0x07, 0x26, 0x45, 0x2e, 0xa2, 0x00, 0x41, 0x65, 0x48, 0x85, 0x51, 0xa1, 0xc4, 0x7c, 0x91,
	0xe7, 0xd6, 0x05, 0x94, 0xfc, 0x33, 0xf1, 0x3d, 0x7b, 0xd2, 0x14, 0x66, 0x73, 0xac, 0xb5, 0x57,
	0xc3, 0xb6, 0x9d, 0xaa, 0x03, 0x1a, 0x81, 0x27, 0x34, 0x12, 0x25, 0xbf, 0x1d, 0xf3, 0x0d, 0x0a,
	0x59, 0x5e, 0x7a, 0x55, 0xf8, 0x00, 0xc9, 0x8f, 0xc0, 0xa1, 0xe4, 0x77, 0xc4, 0x43, 0xb6, 0xd9,
	0xea, 0x54, 0xc3, 0x85, 0x6f, 0x70, 0xaa, 0xe2, 0xf8, 0xe7, 0x17, 0xac, 0x9c, 0x8c, 0x62, 0x99,
	0x35, 0x42, 0xd6, 0x2c, 0x64, 0x2e, 0xdd, 0x53, 0x5f, 0xf6, 0xdd, 0x6a, 0xf2, 0x6a, 0x44, 0xc3,
	0x00, 0x35, 0x49, 0x8b, 0x55, 0xaf, 0xf8, 0x42, 0xf9, 0xc0, 0x57, 0x3f, 0x40, 0x01, 0xed, 0x10,
	0xe4, 0xa4, 0x62, 0x78, 0x7e, 0x4f, 0xdc, 0x67, 0xab, 0x99, 0xf2, 0xbe, 0xec, 0x6c, 0x64, 0xf3,
	0xfc, 0xa2, 0x96, 0x5f, 0x88, 0x2d, 0xb6, 0xde, 0x88, 0xb1, 0x40, 0x41, 0xa3, 0xa7, 0xc4, 0x16,
	0x26, 0xb8, 0x09, 0x5f, 0x8b, 0xb7, 0xde, 0x64, 0x28, 0x8c, 0x1a, 0xa3, 0xf3, 0xa0, 0x09, 0xf2,
	0x7c, 0x51, 0x77, 0x93, 0x6f, 0x5d, 0x3c, 0x63, 0x3f, 0x81, 0x1c, 0xa3, 0x0b, 0xaa, 0xe4, 0xc7,
	0x56, 0x0d, 0x6a, 0x0a, 0x93, 0x1c, 0xcb, 0xc2, 0x61, 0x0c, 0x4a, 0x97, 0x9e, 0x8e, 0xce, 0x85,
	0xa4, 0x3c, 0xa5, 0x84, 0xf9, 0x83, 0x38, 0x8c, 0x0b, 0x7b, 0xbd, 0x9f, 0xc4, 0x17, 0x83, 0x8a,
	0xb9, 0x21, 0x1e, 0xb3, 0x9d, 0xce, 0x40, 0x27, 0xd6, 0xc4, 0xba, 0xda, 0x29, 0x8a, 0xf7, 0x83,
	0x1a, 0x93, 0x78, 0xe3, 0x9b, 0xd1, 0xf2, 0x17, 0x0c, 0x07, 0x92, 0x9c, 0x0d, 0xf5, 0xc8, 0x59,
	0x89, 0x7c, 0x4b, 0xf4, 0xd9, 0x77, 0xdd, 0xb1, 0xf2, 0x01, 0x5c, 0x20, 0x19, 0x4d, 0xa8, 0x86,
	0xad, 0xa6, 0xe5, 0x2e, 0x4a, 0xfe, 0x55, 0x34, 0xc4, 0x85, 0x8c, 0x25, 0x2f, 0x58, 0xca, 0xc1,
	0x07, 0xbe, 0x1d, 0x8f, 0x6b, 0xf4, 0x1a, 0xd9, 0x80, 0x9a, 0x92, 0xc2, 0x07, 0x9b, 0xa1, 0x23,
	0xad, 0xcc, 0x1e, 0x7f, 0x58, 0x8d, 0x60, 0xd5, 0xe7, 0x87, 0x08, 0x8f, 0xe2, 0x08, 0x76, 0x33,
	0x34, 0x02, 0x7f, 0x2d, 0x06, 0xec, 0xf9, 0xfb, 0x26, 0x6b, 0x2f, 0xa2, 0xf6, 0xd4, 0xfe, 0x48,
	0x69, 0x2c, 0xff, 0xca, 0xc8, 0x0e, 0x17, 0x68, 0xea, 0x6c, 0x91, 0xf3, 0x9d, 0xc1, 0xdf, 0x4b,
	0x6c, 0xeb, 0xd5, 0xec, 0xa4, 0x77, 0xf9, 0xfb, 0x31, 0x10, 0x9d, 0x67, 0x22, 0x8f, 0x6f, 0x4e,
	0xbe, 0xf4, 0x52, 0xd6, 0x51, 0xd3, 0xd9, 0x9b, 0xc3, 0xd3, 0x69, 0x6f, 0x76, 0x3e, 0xed, 0x4f,
	0x8f, 0x4f, 0xcb, 0x17, 0xa9, 0x79, 0xba, 0xce, 0x5e, 0xcf, 0xff, 0xeb, 0x25, 0x7b, 0x56, 0xfd,
	0xfc, 0x79, 0xe5, 0x6a, 0x0a, 0xf0, 0xd7, 0x95, 0xf5, 0xb4, 0x4a, 0x06, 0x47, 0xf3, 0x5e, 0xb5,
	0x8c, 0xab, 0xf1, 0x6e, 0xaf, 0x3c, 0x72, 0xfe, 0x4f, 0x43, 0x38, 0x80, 0xa3, 0xf9, 0x41, 0x4b,
	0x38, 0x18, 0xef, 0x1e, 0x54, 0x84, 0x5f, 0xae, 0x95, 0x07, 0x3f, 0xfd, 0x37, 0x00, 0x00, 0xff,
	0xff, 0xcf, 0xa0, 0xc1, 0x07, 0x41, 0x07, 0x00, 0x00,
}
