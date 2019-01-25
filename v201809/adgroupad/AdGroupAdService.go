package adgroupad

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "https://adwords.google.com/api/adwords/cm/v201809"

// NewAdGroupAdServiceInterface creates an initializes a AdGroupAdServiceInterface.
func NewAdGroupAdServiceInterface(cli *soap.Client) AdGroupAdServiceInterface {
	return &adGroupAdServiceInterface{cli}
}

// AdGroupAdServiceInterface was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type AdGroupAdServiceInterface interface {
	// Returns a list of AdGroupAds.         AdGroupAds that had been
	// removed are not returned by default.                  @param
	// serviceSelector The selector specifying the {@link AdGroupAd}s
	// to return.         @return The page containing the AdGroupAds
	// that meet the criteria specified by the selector.         @throws
	// ApiException when there is at least one error with the request.
	Get(Get *Get) (*GetResponse, error)

	// Applies the list of mutate operations (ie. add, set, remove):
	//         <p>Add - Creates a new {@linkplain AdGroupAd ad group
	// ad}. The         {@code adGroupId} must         reference an
	// existing ad group. The child {@code Ad} must be sufficiently
	//         specified by constructing a concrete ad type (such as
	// {@code TextAd})         and setting its fields accordingly.</p>
	//         <p>Set - Updates an ad group ad. Except for {@code status},
	//         ad group ad fields are not mutable. Status updates are
	//         straightforward - the status of the ad group ad is updated
	// as         specified. If any other field has changed, it will
	// be ignored. If         you want to change any of the fields
	// other than status, you must         make a new ad and then remove
	// the old one.</p>         <p>Remove - Removes the link between
	// the specified AdGroup and         Ad.</p>         @param operations
	// The operations to apply.         @return A list of AdGroupAds
	// where each entry in the list is the result of         applying
	// the operation in the input list with the same index. For an
	//         add/set operation, the return AdGroupAd will be what
	// is saved to the db.         In the case of the remove operation,
	// the return AdGroupAd will simply be         an AdGroupAd containing
	// an Ad with the id set to the Ad being removed from         the
	// AdGroup.
	Mutate(Mutate *Mutate) (*MutateResponse, error)

	// Adds labels to the AdGroupAd or removes labels from the AdGroupAd.
	//         <p>Add - Apply an existing label to an existing {@linkplain
	// AdGroupAd ad group ad}. The         {@code adGroupId} and {@code
	// adId} must reference an existing         {@linkplain AdGroupAd
	// ad group ad}. The {@code labelId} must reference an existing
	//         {@linkplain Label label}.         <p>Remove - Removes
	// the link between the specified {@linkplain AdGroupAd ad group
	// ad} and         {@linkplain Label label}.         @param operations
	// The operations to apply.         @return A list of AdGroupAdLabel
	// where each entry in the list is the result of         applying
	// the operation in the input list with the same index. For an
	//         add operation, the returned AdGroupAdLabel contains
	// the AdGroupId, AdId and the LabelId.         In the case of
	// a remove operation, the returned AdGroupAdLabel will only have
	// AdGroupId and         AdId.         @throws ApiException when
	// there are one or more errors with the request.
	MutateLabel(MutateLabel *MutateLabel) (*MutateLabelResponse, error)

	// Returns a list of AdGroupAds based on the query.
	//       @param query The SQL-like AWQL query string.         @return
	// A list of AdGroupAds.         @throws ApiException if problems
	// occur while parsing the query or fetching AdGroupAds.
	Query(Query string) (*QueryResponse, error)
}

// AdType was auto-generated from WSDL.
type AdType string

// Validate validates AdType.
func (v AdType) Validate() bool {
	for _, vv := range []string{
		"DEPRECATED_AD",
		"IMAGE_AD",
		"PRODUCT_AD",
		"TEMPLATE_AD",
		"TEXT_AD",
		"THIRD_PARTY_REDIRECT_AD",
		"DYNAMIC_SEARCH_AD",
		"CALL_ONLY_AD",
		"EXPANDED_TEXT_AD",
		"RESPONSIVE_DISPLAY_AD",
		"SHOWCASE_AD",
		"GOAL_OPTIMIZED_SHOPPING_AD",
		"EXPANDED_DYNAMIC_SEARCH_AD",
		"GMAIL_AD",
		"RESPONSIVE_SEARCH_AD",
		"MULTI_ASSET_RESPONSIVE_DISPLAY_AD",
		"UNIVERSAL_APP_AD",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AdCustomizerErrorReason was auto-generated from WSDL.
type AdCustomizerErrorReason string

// Validate validates AdCustomizerErrorReason.
func (v AdCustomizerErrorReason) Validate() bool {
	for _, vv := range []string{
		"COUNTDOWN_INVALID_DATE_FORMAT",
		"COUNTDOWN_DATE_IN_PAST",
		"COUNTDOWN_INVALID_LOCALE",
		"COUNTDOWN_INVALID_START_DAYS_BEFORE",
		"UNKNOWN_USER_LIST",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AdErrorReason was auto-generated from WSDL.
type AdErrorReason string

// Validate validates AdErrorReason.
func (v AdErrorReason) Validate() bool {
	for _, vv := range []string{
		"AD_CUSTOMIZERS_NOT_SUPPORTED_FOR_AD_TYPE",
		"APPROXIMATELY_TOO_LONG",
		"APPROXIMATELY_TOO_SHORT",
		"BAD_SNIPPET",
		"CANNOT_MODIFY_AD",
		"CANNOT_SET_BUSINESS_NAME_IF_URL_SET",
		"CANNOT_SET_FIELD",
		"CANNOT_SET_FIELD_WITH_ORIGIN_AD_ID_SET",
		"CANNOT_SET_FIELD_WITH_AD_ID_SET_FOR_SHARING",
		"CANNOT_SET_ALLOW_FLEXIBLE_COLOR_FALSE",
		"CANNOT_SET_COLOR_CONTROL_WHEN_NATIVE_FORMAT_SETTING",
		"CANNOT_SET_URL",
		"CANNOT_SET_WITHOUT_FINAL_URLS",
		"CANNOT_SET_WITH_FINAL_URLS",
		"CANNOT_SET_WITH_TRACKING_URL_TEMPLATE",
		"CANNOT_SET_WITH_URL_DATA",
		"CANNOT_USE_AD_SUBCLASS_FOR_OPERATOR",
		"CUSTOMER_NOT_APPROVED_MOBILEADS",
		"CUSTOMER_NOT_APPROVED_THIRDPARTY_ADS",
		"CUSTOMER_NOT_APPROVED_THIRDPARTY_REDIRECT_ADS",
		"CUSTOMER_NOT_ELIGIBLE",
		"CUSTOMER_NOT_ELIGIBLE_FOR_UPDATING_BEACON_URL",
		"DIMENSION_ALREADY_IN_UNION",
		"DIMENSION_MUST_BE_SET",
		"DIMENSION_NOT_IN_UNION",
		"DISPLAY_URL_CANNOT_BE_SPECIFIED",
		"DOMESTIC_PHONE_NUMBER_FORMAT",
		"EMERGENCY_PHONE_NUMBER",
		"EMPTY_FIELD",
		"FEED_ATTRIBUTE_MUST_HAVE_MAPPING_FOR_TYPE_ID",
		"FEED_ATTRIBUTE_MAPPING_TYPE_MISMATCH",
		"ILLEGAL_AD_CUSTOMIZER_TAG_USE",
		"ILLEGAL_TAG_USE",
		"INCONSISTENT_DIMENSIONS",
		"INCONSISTENT_STATUS_IN_TEMPLATE_UNION",
		"INCORRECT_LENGTH",
		"INELIGIBLE_FOR_UPGRADE",
		"INVALID_AD_ADDRESS_CAMPAIGN_TARGET",
		"INVALID_AD_TYPE",
		"INVALID_ATTRIBUTES_FOR_MOBILE_IMAGE",
		"INVALID_ATTRIBUTES_FOR_MOBILE_TEXT",
		"INVALID_CALL_TO_ACTION_TEXT",
		"INVALID_CHARACTER_FOR_URL",
		"INVALID_COUNTRY_CODE",
		"INVALID_DSA_URL_TAG",
		"INVALID_EXPANDED_DYNAMIC_SEARCH_AD_TAG",
		"INVALID_INPUT",
		"INVALID_MARKUP_LANGUAGE",
		"INVALID_MOBILE_CARRIER",
		"INVALID_MOBILE_CARRIER_TARGET",
		"INVALID_NUMBER_OF_ELEMENTS",
		"INVALID_PHONE_NUMBER_FORMAT",
		"INVALID_RICH_MEDIA_CERTIFIED_VENDOR_FORMAT_ID",
		"INVALID_TEMPLATE_DATA",
		"INVALID_TEMPLATE_ELEMENT_FIELD_TYPE",
		"INVALID_TEMPLATE_ID",
		"LINE_TOO_WIDE",
		"MISSING_AD_CUSTOMIZER_MAPPING",
		"MISSING_ADDRESS_COMPONENT",
		"MISSING_ADVERTISEMENT_NAME",
		"MISSING_BUSINESS_NAME",
		"MISSING_DESCRIPTION1",
		"MISSING_DESCRIPTION2",
		"MISSING_DESTINATION_URL_TAG",
		"MISSING_LANDING_PAGE_URL_TAG",
		"MISSING_DIMENSION",
		"MISSING_DISPLAY_URL",
		"MISSING_HEADLINE",
		"MISSING_HEIGHT",
		"MISSING_IMAGE",
		"MISSING_MARKETING_IMAGE_OR_PRODUCT_VIDEOS",
		"MISSING_MARKUP_LANGUAGES",
		"MISSING_MOBILE_CARRIER",
		"MISSING_PHONE",
		"MISSING_REQUIRED_TEMPLATE_FIELDS",
		"MISSING_TEMPLATE_FIELD_VALUE",
		"MISSING_TEXT",
		"MISSING_VISIBLE_URL",
		"MISSING_WIDTH",
		"MULTIPLE_DISTINCT_FEEDS_UNSUPPORTED",
		"MUST_USE_TEMP_AD_UNION_ID_ON_ADD",
		"TOO_LONG",
		"TOO_SHORT",
		"UNION_DIMENSIONS_CANNOT_CHANGE",
		"UNKNOWN_ADDRESS_COMPONENT",
		"UNKNOWN_FIELD_NAME",
		"UNKNOWN_UNIQUE_NAME",
		"UNSUPPORTED_DIMENSIONS",
		"URL_INVALID_SCHEME",
		"URL_INVALID_TOP_LEVEL_DOMAIN",
		"URL_MALFORMED",
		"URL_NO_HOST",
		"URL_NOT_EQUIVALENT",
		"URL_HOST_NAME_TOO_LONG",
		"URL_NO_SCHEME",
		"URL_NO_TOP_LEVEL_DOMAIN",
		"URL_PATH_NOT_ALLOWED",
		"URL_PORT_NOT_ALLOWED",
		"URL_QUERY_NOT_ALLOWED",
		"URL_SCHEME_BEFORE_DSA_TAG",
		"URL_SCHEME_BEFORE_EXPANDED_DYNAMIC_SEARCH_AD_TAG",
		"USER_DOES_NOT_HAVE_ACCESS_TO_TEMPLATE",
		"INCONSISTENT_EXPANDABLE_SETTINGS",
		"INVALID_FORMAT",
		"INVALID_FIELD_TEXT",
		"ELEMENT_NOT_PRESENT",
		"IMAGE_ERROR",
		"VALUE_NOT_IN_RANGE",
		"FIELD_NOT_PRESENT",
		"ADDRESS_NOT_COMPLETE",
		"ADDRESS_INVALID",
		"VIDEO_RETRIEVAL_ERROR",
		"AUDIO_ERROR",
		"INVALID_YOUTUBE_DISPLAY_URL",
		"TOO_MANY_PRODUCT_IMAGES",
		"TOO_MANY_PRODUCT_VIDEOS",
		"INCOMPATIBLE_AD_TYPE_AND_DEVICE_PREFERENCE",
		"CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY",
		"CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED",
		"DISALLOWED_NUMBER_TYPE",
		"PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY",
		"PHONE_NUMBER_NOT_SUPPORTED_WITH_CALLTRACKING_FOR_COUNTRY",
		"PREMIUM_RATE_NUMBER_NOT_ALLOWED",
		"VANITY_PHONE_NUMBER_NOT_ALLOWED",
		"INVALID_CALL_CONVERSION_TYPE_ID",
		"CANNOT_DISABLE_CALL_CONVERSION_AND_SET_CONVERSION_TYPE_ID",
		"CANNOT_SET_PATH2_WITHOUT_PATH1",
		"MISSING_DYNAMIC_SEARCH_ADS_SETTING_DOMAIN_NAME",
		"CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AdGroupAdStatus was auto-generated from WSDL.
type AdGroupAdStatus string

// Validate validates AdGroupAdStatus.
func (v AdGroupAdStatus) Validate() bool {
	for _, vv := range []string{
		"ENABLED",
		"PAUSED",
		"DISABLED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AdGroupAdErrorReason was auto-generated from WSDL.
type AdGroupAdErrorReason string

// Validate validates AdGroupAdErrorReason.
func (v AdGroupAdErrorReason) Validate() bool {
	for _, vv := range []string{
		"AD_GROUP_AD_LABEL_DOES_NOT_EXIST",
		"AD_GROUP_AD_LABEL_ALREADY_EXISTS",
		"AD_NOT_UNDER_ADGROUP",
		"CANNOT_OPERATE_ON_REMOVED_ADGROUPAD",
		"CANNOT_CREATE_DEPRECATED_ADS",
		"CANNOT_CREATE_TEXT_ADS",
		"EMPTY_FIELD",
		"ENTITY_REFERENCED_IN_MULTIPLE_OPS",
		"UNSUPPORTED_OPERATION",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AdSharingErrorReason was auto-generated from WSDL.
type AdSharingErrorReason string

// Validate validates AdSharingErrorReason.
func (v AdSharingErrorReason) Validate() bool {
	for _, vv := range []string{
		"AD_GROUP_ALREADY_CONTAINS_AD",
		"INCOMPATIBLE_AD_UNDER_AD_GROUP",
		"CANNOT_SHARE_INACTIVE_AD",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AdStrength was auto-generated from WSDL.
type AdStrength string

// Validate validates AdStrength.
func (v AdStrength) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"PENDING",
		"NO_ADS",
		"POOR",
		"AVERAGE",
		"GOOD",
		"EXCELLENT",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AdxErrorReason was auto-generated from WSDL.
type AdxErrorReason string

// Validate validates AdxErrorReason.
func (v AdxErrorReason) Validate() bool {
	for _, vv := range []string{
		"UNSUPPORTED_FEATURE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AppUrlOsType was auto-generated from WSDL.
type AppUrlOsType string

// Validate validates AppUrlOsType.
func (v AppUrlOsType) Validate() bool {
	for _, vv := range []string{
		"OS_TYPE_IOS",
		"OS_TYPE_ANDROID",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AssetType was auto-generated from WSDL.
type AssetType string

// Validate validates AssetType.
func (v AssetType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"TEXT",
		"IMAGE",
		"YOUTUBE_VIDEO",
		"MEDIA_BUNDLE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AssetErrorReason was auto-generated from WSDL.
type AssetErrorReason string

// Validate validates AssetErrorReason.
func (v AssetErrorReason) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"ASSET_TYPE_NOT_SUPPORTED",
		"CANNOT_REMOVE_ASSET_WITH_REMOVED_STATUS",
		"CANNOT_MODIFY_ASSET_NAME",
		"DUPLICATE_ASSET",
		"DUPLICATE_ASSET_NAME",
		"ASSET_DATA_IS_MISSING",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AssetLinkErrorReason was auto-generated from WSDL.
type AssetLinkErrorReason string

// Validate validates AssetLinkErrorReason.
func (v AssetLinkErrorReason) Validate() bool {
	for _, vv := range []string{
		"INVALID_ASSET_TYPE_FOR_FIELD",
		"PINNING_UNSUPPORTED",
		"INVALID_PINNED_FIELD",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AssetPerformanceLabel was auto-generated from WSDL.
type AssetPerformanceLabel string

// Validate validates AssetPerformanceLabel.
func (v AssetPerformanceLabel) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"PENDING",
		"LEARNING",
		"LOW",
		"GOOD",
		"BEST",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AssetStatus was auto-generated from WSDL.
type AssetStatus string

// Validate validates AssetStatus.
func (v AssetStatus) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"ENABLED",
		"REMOVED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AuthenticationErrorReason was auto-generated from WSDL.
type AuthenticationErrorReason string

// Validate validates AuthenticationErrorReason.
func (v AuthenticationErrorReason) Validate() bool {
	for _, vv := range []string{
		"AUTHENTICATION_FAILED",
		"CLIENT_CUSTOMER_ID_IS_REQUIRED",
		"CLIENT_EMAIL_REQUIRED",
		"CLIENT_CUSTOMER_ID_INVALID",
		"CLIENT_EMAIL_INVALID",
		"CLIENT_EMAIL_FAILED_TO_AUTHENTICATE",
		"CUSTOMER_NOT_FOUND",
		"GOOGLE_ACCOUNT_DELETED",
		"GOOGLE_ACCOUNT_COOKIE_INVALID",
		"FAILED_TO_AUTHENTICATE_GOOGLE_ACCOUNT",
		"GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH",
		"LOGIN_COOKIE_REQUIRED",
		"NOT_ADS_USER",
		"OAUTH_TOKEN_INVALID",
		"OAUTH_TOKEN_EXPIRED",
		"OAUTH_TOKEN_DISABLED",
		"OAUTH_TOKEN_REVOKED",
		"OAUTH_TOKEN_HEADER_INVALID",
		"LOGIN_COOKIE_INVALID",
		"FAILED_TO_RETRIEVE_LOGIN_COOKIE",
		"USER_ID_INVALID",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AuthorizationErrorReason was auto-generated from WSDL.
type AuthorizationErrorReason string

// Validate validates AuthorizationErrorReason.
func (v AuthorizationErrorReason) Validate() bool {
	for _, vv := range []string{
		"UNABLE_TO_AUTHORIZE",
		"NO_ADWORDS_ACCOUNT_FOR_CUSTOMER",
		"USER_PERMISSION_DENIED",
		"EFFECTIVE_USER_PERMISSION_DENIED",
		"CUSTOMER_NOT_ACTIVE",
		"USER_HAS_READONLY_PERMISSION",
		"NO_CUSTOMER_FOUND",
		"SERVICE_ACCESS_DENIED",
		"TWO_STEP_VERIFICATION_NOT_ENROLLED",
		"ADVANCED_PROTECTION_NOT_ENROLLED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ClientTermsErrorReason was auto-generated from WSDL.
type ClientTermsErrorReason string

// Validate validates ClientTermsErrorReason.
func (v ClientTermsErrorReason) Validate() bool {
	for _, vv := range []string{
		"INCOMPLETE_SIGNUP_CURRENT_ADWORDS_TNC_NOT_AGREED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// CollectionSizeErrorReason was auto-generated from WSDL.
type CollectionSizeErrorReason string

// Validate validates CollectionSizeErrorReason.
func (v CollectionSizeErrorReason) Validate() bool {
	for _, vv := range []string{
		"TOO_FEW",
		"TOO_MANY",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// DatabaseErrorReason was auto-generated from WSDL.
type DatabaseErrorReason string

// Validate validates DatabaseErrorReason.
func (v DatabaseErrorReason) Validate() bool {
	for _, vv := range []string{
		"CONCURRENT_MODIFICATION",
		"PERMISSION_DENIED",
		"ACCESS_PROHIBITED",
		"CAMPAIGN_PRODUCT_NOT_SUPPORTED",
		"DUPLICATE_KEY",
		"DATABASE_ERROR",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// DateErrorReason was auto-generated from WSDL.
type DateErrorReason string

// Validate validates DateErrorReason.
func (v DateErrorReason) Validate() bool {
	for _, vv := range []string{
		"INVALID_FIELD_VALUES_IN_DATE",
		"INVALID_FIELD_VALUES_IN_DATE_TIME",
		"INVALID_STRING_DATE",
		"INVALID_STRING_DATE_RANGE",
		"INVALID_STRING_DATE_TIME",
		"EARLIER_THAN_MINIMUM_DATE",
		"LATER_THAN_MAXIMUM_DATE",
		"DATE_RANGE_MINIMUM_DATE_LATER_THAN_MAXIMUM_DATE",
		"DATE_RANGE_MINIMUM_AND_MAXIMUM_DATES_BOTH_NULL",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// DeprecatedAdType was auto-generated from WSDL.
type DeprecatedAdType string

// Validate validates DeprecatedAdType.
func (v DeprecatedAdType) Validate() bool {
	for _, vv := range []string{
		"VIDEO",
		"CLICK_TO_CALL",
		"IN_STREAM_VIDEO",
		"FROOGLE",
		"TEXT_LINK",
		"GADGET",
		"PRINT",
		"TEXT_WIDE",
		"GADGET_TEMPLATE",
		"TEXT_WITH_VIDEO",
		"AUDIO",
		"LOCAL_BUSINESS_AD",
		"AUDIO_TEMPLATE",
		"MOBILE_AD",
		"MOBILE_IMAGE_AD",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// DisplayAdFormatSetting was auto-generated from WSDL.
type DisplayAdFormatSetting string

// Validate validates DisplayAdFormatSetting.
func (v DisplayAdFormatSetting) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"ALL_FORMATS",
		"NON_NATIVE",
		"NATIVE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// DistinctErrorReason was auto-generated from WSDL.
type DistinctErrorReason string

// Validate validates DistinctErrorReason.
func (v DistinctErrorReason) Validate() bool {
	for _, vv := range []string{
		"DUPLICATE_ELEMENT",
		"DUPLICATE_TYPE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EntityAccessDeniedReason was auto-generated from WSDL.
type EntityAccessDeniedReason string

// Validate validates EntityAccessDeniedReason.
func (v EntityAccessDeniedReason) Validate() bool {
	for _, vv := range []string{
		"READ_ACCESS_DENIED",
		"WRITE_ACCESS_DENIED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EntityCountLimitExceededReason was auto-generated from WSDL.
type EntityCountLimitExceededReason string

// Validate validates EntityCountLimitExceededReason.
func (v EntityCountLimitExceededReason) Validate() bool {
	for _, vv := range []string{
		"ACCOUNT_LIMIT",
		"CAMPAIGN_LIMIT",
		"ADGROUP_LIMIT",
		"AD_GROUP_AD_LIMIT",
		"AD_GROUP_CRITERION_LIMIT",
		"SHARED_SET_LIMIT",
		"MATCHING_FUNCTION_LIMIT",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EntityNotFoundReason was auto-generated from WSDL.
type EntityNotFoundReason string

// Validate validates EntityNotFoundReason.
func (v EntityNotFoundReason) Validate() bool {
	for _, vv := range []string{
		"INVALID_ID",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// FeedAttributeReferenceErrorReason was auto-generated from WSDL.
type FeedAttributeReferenceErrorReason string

// Validate validates FeedAttributeReferenceErrorReason.
func (v FeedAttributeReferenceErrorReason) Validate() bool {
	for _, vv := range []string{
		"CANNOT_REFERENCE_DELETED_FEED",
		"INVALID_FEED_NAME",
		"INVALID_FEED_ATTRIBUTE_NAME",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ForwardCompatibilityErrorReason was auto-generated from WSDL.
type ForwardCompatibilityErrorReason string

// Validate validates ForwardCompatibilityErrorReason.
func (v ForwardCompatibilityErrorReason) Validate() bool {
	for _, vv := range []string{
		"INVALID_FORWARD_COMPATIBILITY_MAP_VALUE",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// FunctionErrorReason was auto-generated from WSDL.
type FunctionErrorReason string

// Validate validates FunctionErrorReason.
func (v FunctionErrorReason) Validate() bool {
	for _, vv := range []string{
		"INVALID_FUNCTION_FORMAT",
		"DATA_TYPE_MISMATCH",
		"INVALID_CONJUNCTION_OPERANDS",
		"INVALID_NUMBER_OF_OPERANDS",
		"INVALID_OPERAND_TYPE",
		"INVALID_OPERATOR",
		"INVALID_REQUEST_CONTEXT_TYPE",
		"INVALID_FUNCTION_FOR_CALL_PLACEHOLDER",
		"INVALID_FUNCTION_FOR_PLACEHOLDER",
		"INVALID_OPERAND",
		"MISSING_CONSTANT_OPERAND_VALUE",
		"INVALID_CONSTANT_OPERAND_VALUE",
		"INVALID_NESTING",
		"MULTIPLE_FEED_IDS_NOT_SUPPORTED",
		"INVALID_FUNCTION_FOR_FEED_WITH_FIXED_SCHEMA",
		"INVALID_ATTRIBUTE_NAME",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// FunctionParsingErrorReason was auto-generated from WSDL.
type FunctionParsingErrorReason string

// Validate validates FunctionParsingErrorReason.
func (v FunctionParsingErrorReason) Validate() bool {
	for _, vv := range []string{
		"NO_MORE_INPUT",
		"EXPECTED_CHARACTER",
		"UNEXPECTED_SEPARATOR",
		"UNMATCHED_LEFT_BRACKET",
		"UNMATCHED_RIGHT_BRACKET",
		"TOO_MANY_NESTED_FUNCTIONS",
		"MISSING_RIGHT_HAND_OPERAND",
		"INVALID_OPERATOR_NAME",
		"FEED_ATTRIBUTE_OPERAND_ARGUMENT_NOT_INTEGER",
		"NO_OPERANDS",
		"TOO_MANY_OPERANDS",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// IdErrorReason was auto-generated from WSDL.
type IdErrorReason string

// Validate validates IdErrorReason.
func (v IdErrorReason) Validate() bool {
	for _, vv := range []string{
		"NOT_FOUND",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ImageErrorReason was auto-generated from WSDL.
type ImageErrorReason string

// Validate validates ImageErrorReason.
func (v ImageErrorReason) Validate() bool {
	for _, vv := range []string{
		"INVALID_IMAGE",
		"STORAGE_ERROR",
		"BAD_REQUEST",
		"UNEXPECTED_SIZE",
		"ANIMATED_NOT_ALLOWED",
		"ANIMATION_TOO_LONG",
		"SERVER_ERROR",
		"CMYK_JPEG_NOT_ALLOWED",
		"FLASH_NOT_ALLOWED",
		"FLASH_WITHOUT_CLICKTAG",
		"FLASH_ERROR_AFTER_FIXING_CLICK_TAG",
		"ANIMATED_VISUAL_EFFECT",
		"FLASH_ERROR",
		"LAYOUT_PROBLEM",
		"PROBLEM_READING_IMAGE_FILE",
		"ERROR_STORING_IMAGE",
		"ASPECT_RATIO_NOT_ALLOWED",
		"FLASH_HAS_NETWORK_OBJECTS",
		"FLASH_HAS_NETWORK_METHODS",
		"FLASH_HAS_URL",
		"FLASH_HAS_MOUSE_TRACKING",
		"FLASH_HAS_RANDOM_NUM",
		"FLASH_SELF_TARGETS",
		"FLASH_BAD_GETURL_TARGET",
		"FLASH_VERSION_NOT_SUPPORTED",
		"FLASH_WITHOUT_HARD_CODED_CLICK_URL",
		"INVALID_FLASH_FILE",
		"FAILED_TO_FIX_CLICK_TAG_IN_FLASH",
		"FLASH_ACCESSES_NETWORK_RESOURCES",
		"FLASH_EXTERNAL_JS_CALL",
		"FLASH_EXTERNAL_FS_CALL",
		"FILE_TOO_LARGE",
		"IMAGE_DATA_TOO_LARGE",
		"IMAGE_PROCESSING_ERROR",
		"IMAGE_TOO_SMALL",
		"INVALID_INPUT",
		"PROBLEM_READING_FILE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// InternalApiErrorReason was auto-generated from WSDL.
type InternalApiErrorReason string

// Validate validates InternalApiErrorReason.
func (v InternalApiErrorReason) Validate() bool {
	for _, vv := range []string{
		"UNEXPECTED_INTERNAL_API_ERROR",
		"TRANSIENT_ERROR",
		"UNKNOWN",
		"DOWNTIME",
		"ERROR_GENERATING_RESPONSE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// LabelStatus was auto-generated from WSDL.
type LabelStatus string

// Validate validates LabelStatus.
func (v LabelStatus) Validate() bool {
	for _, vv := range []string{
		"ENABLED",
		"REMOVED",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// MediaLegacyMimeType was auto-generated from WSDL.
type MediaLegacyMimeType string

// Validate validates MediaLegacyMimeType.
func (v MediaLegacyMimeType) Validate() bool {
	for _, vv := range []string{
		"IMAGE_JPEG",
		"IMAGE_GIF",
		"IMAGE_PNG",
		"FLASH",
		"TEXT_HTML",
		"PDF",
		"MSWORD",
		"MSEXCEL",
		"RTF",
		"AUDIO_WAV",
		"AUDIO_MP3",
		"HTML5_AD_ZIP",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// MediaMediaType was auto-generated from WSDL.
type MediaMediaType string

// Validate validates MediaMediaType.
func (v MediaMediaType) Validate() bool {
	for _, vv := range []string{
		"AUDIO",
		"DYNAMIC_IMAGE",
		"ICON",
		"IMAGE",
		"STANDARD_ICON",
		"VIDEO",
		"MEDIA_BUNDLE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// MediaSize was auto-generated from WSDL.
type MediaSize string

// Validate validates MediaSize.
func (v MediaSize) Validate() bool {
	for _, vv := range []string{
		"FULL",
		"SHRUNKEN",
		"PREVIEW",
		"VIDEO_THUMBNAIL",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// MediaBundleErrorReason was auto-generated from WSDL.
type MediaBundleErrorReason string

// Validate validates MediaBundleErrorReason.
func (v MediaBundleErrorReason) Validate() bool {
	for _, vv := range []string{
		"ENTRY_POINT_CANNOT_BE_SET_USING_MEDIA_SERVICE",
		"BAD_REQUEST",
		"DOUBLECLICK_BUNDLE_NOT_ALLOWED",
		"EXTERNAL_URL_NOT_ALLOWED",
		"FILE_TOO_LARGE",
		"GOOGLE_WEB_DESIGNER_ZIP_FILE_NOT_PUBLISHED",
		"INVALID_INPUT",
		"INVALID_MEDIA_BUNDLE",
		"INVALID_MEDIA_BUNDLE_ENTRY",
		"INVALID_MIME_TYPE",
		"INVALID_PATH",
		"INVALID_URL_REFERENCE",
		"MEDIA_DATA_TOO_LARGE",
		"MISSING_PRIMARY_MEDIA_BUNDLE_ENTRY",
		"SERVER_ERROR",
		"STORAGE_ERROR",
		"SWIFFY_BUNDLE_NOT_ALLOWED",
		"TOO_MANY_FILES",
		"UNEXPECTED_SIZE",
		"UNSUPPORTED_GOOGLE_WEB_DESIGNER_ENVIRONMENT",
		"UNSUPPORTED_HTML5_FEATURE",
		"URL_IN_MEDIA_BUNDLE_NOT_SSL_COMPLIANT",
		"CUSTOM_EXIT_NOT_ALLOWED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// MediaErrorReason was auto-generated from WSDL.
type MediaErrorReason string

// Validate validates MediaErrorReason.
func (v MediaErrorReason) Validate() bool {
	for _, vv := range []string{
		"CANNOT_ADD_STANDARD_ICON",
		"CANNOT_SELECT_STANDARD_ICON_WITH_OTHER_TYPES",
		"CANNOT_SPECIFY_MEDIA_ID_AND_DATA",
		"DUPLICATE_MEDIA",
		"EMPTY_FIELD",
		"ENTITY_REFERENCED_IN_MULTIPLE_OPS",
		"FIELD_NOT_SUPPORTED_FOR_MEDIA_SUB_TYPE",
		"INVALID_MEDIA_ID",
		"INVALID_MEDIA_SUB_TYPE",
		"INVALID_MEDIA_TYPE",
		"INVALID_MIME_TYPE",
		"INVALID_REFERENCE_ID",
		"INVALID_YOU_TUBE_ID",
		"MEDIA_FAILED_TRANSCODING",
		"MEDIA_NOT_TRANSCODED",
		"MEDIA_TYPE_DOES_NOT_MATCH_OBJECT_TYPE",
		"NO_FIELDS_SPECIFIED",
		"NULL_REFERENCE_ID_AND_MEDIA_ID",
		"TOO_LONG",
		"UNSUPPORTED_OPERATION",
		"UNSUPPORTED_TYPE",
		"YOU_TUBE_SERVICE_UNAVAILABLE",
		"YOU_TUBE_VIDEO_HAS_NON_POSITIVE_DURATION",
		"YOU_TUBE_VIDEO_NOT_FOUND",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// MediaMimeType was auto-generated from WSDL.
type MediaMimeType string

// Validate validates MediaMimeType.
func (v MediaMimeType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"IMAGE_JPEG",
		"IMAGE_GIF",
		"IMAGE_PNG",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// NewEntityCreationErrorReason was auto-generated from WSDL.
type NewEntityCreationErrorReason string

// Validate validates NewEntityCreationErrorReason.
func (v NewEntityCreationErrorReason) Validate() bool {
	for _, vv := range []string{
		"CANNOT_SET_ID_FOR_ADD",
		"DUPLICATE_TEMP_IDS",
		"TEMP_ID_ENTITY_HAD_ERRORS",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// NotEmptyErrorReason was auto-generated from WSDL.
type NotEmptyErrorReason string

// Validate validates NotEmptyErrorReason.
func (v NotEmptyErrorReason) Validate() bool {
	for _, vv := range []string{
		"EMPTY_LIST",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// NullErrorReason was auto-generated from WSDL.
type NullErrorReason string

// Validate validates NullErrorReason.
func (v NullErrorReason) Validate() bool {
	for _, vv := range []string{
		"NULL_CONTENT",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// OperationAccessDeniedReason was auto-generated from WSDL.
type OperationAccessDeniedReason string

// Validate validates OperationAccessDeniedReason.
func (v OperationAccessDeniedReason) Validate() bool {
	for _, vv := range []string{
		"ACTION_NOT_PERMITTED",
		"ADD_OPERATION_NOT_PERMITTED",
		"REMOVE_OPERATION_NOT_PERMITTED",
		"SET_OPERATION_NOT_PERMITTED",
		"MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT",
		"OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE",
		"ADD_AS_REMOVED_NOT_PERMITTED",
		"OPERATION_NOT_PERMITTED_FOR_REMOVED_ENTITY",
		"OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// Operator was auto-generated from WSDL.
type Operator string

// Validate validates Operator.
func (v Operator) Validate() bool {
	for _, vv := range []string{
		"ADD",
		"REMOVE",
		"SET",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// OperatorErrorReason was auto-generated from WSDL.
type OperatorErrorReason string

// Validate validates OperatorErrorReason.
func (v OperatorErrorReason) Validate() bool {
	for _, vv := range []string{
		"OPERATOR_NOT_SUPPORTED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PagingErrorReason was auto-generated from WSDL.
type PagingErrorReason string

// Validate validates PagingErrorReason.
func (v PagingErrorReason) Validate() bool {
	for _, vv := range []string{
		"START_INDEX_CANNOT_BE_NEGATIVE",
		"NUMBER_OF_RESULTS_CANNOT_BE_NEGATIVE",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PolicyApprovalStatus was auto-generated from WSDL.
type PolicyApprovalStatus string

// Validate validates PolicyApprovalStatus.
func (v PolicyApprovalStatus) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"APPROVED",
		"APPROVED_LIMITED",
		"ELIGIBLE",
		"UNDER_REVIEW",
		"DISAPPROVED",
		"SITE_SUSPENDED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PolicyFindingErrorReason was auto-generated from WSDL.
type PolicyFindingErrorReason string

// Validate validates PolicyFindingErrorReason.
func (v PolicyFindingErrorReason) Validate() bool {
	for _, vv := range []string{
		"POLICY_FINDING",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PolicySummaryDenormalizedStatus was auto-generated from WSDL.
type PolicySummaryDenormalizedStatus string

// Validate validates PolicySummaryDenormalizedStatus.
func (v PolicySummaryDenormalizedStatus) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"DISAPPROVED",
		"APPROVED_LIMITED",
		"APPROVED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PolicySummaryReviewState was auto-generated from WSDL.
type PolicySummaryReviewState string

// Validate validates PolicySummaryReviewState.
func (v PolicySummaryReviewState) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"REVIEW_IN_PROGRESS",
		"REVIEWED",
		"UNDER_APPEAL",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PolicyTopicConstraintPolicyTopicConstraintType was auto-generated
// from WSDL.
type PolicyTopicConstraintPolicyTopicConstraintType string

// Validate validates PolicyTopicConstraintPolicyTopicConstraintType.
func (v PolicyTopicConstraintPolicyTopicConstraintType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"COUNTRY",
		"RESELLER",
		"CERTIFICATE_MISSING_IN_COUNTRY",
		"CERTIFICATE_DOMAIN_MISMATCH_IN_COUNTRY",
		"CERTIFICATE_MISSING",
		"CERTIFICATE_DOMAIN_MISMATCH",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PolicyTopicEntryType was auto-generated from WSDL.
type PolicyTopicEntryType string

// Validate validates PolicyTopicEntryType.
func (v PolicyTopicEntryType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"PROHIBITED",
		"LIMITED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PolicyTopicEvidenceDestinationMismatchUrlType was auto-generated
// from WSDL.
type PolicyTopicEvidenceDestinationMismatchUrlType string

// Validate validates PolicyTopicEvidenceDestinationMismatchUrlType.
func (v PolicyTopicEvidenceDestinationMismatchUrlType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"DISPLAY_URL",
		"FINAL_URL",
		"FINAL_MOBILE_URL",
		"TRACKING_URL",
		"MOBILE_TRACKING_URL",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PolicyTopicEvidenceType was auto-generated from WSDL.
type PolicyTopicEvidenceType string

// Validate validates PolicyTopicEvidenceType.
func (v PolicyTopicEvidenceType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"AD_TEXT",
		"HTTP_CODE",
		"WEBSITES",
		"LANGUAGE",
		"AD_TEXT_LIST",
		"DESTINATION_TEXT_LIST",
		"DESTINATION_MISMATCH",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PredicateOperator was auto-generated from WSDL.
type PredicateOperator string

// Validate validates PredicateOperator.
func (v PredicateOperator) Validate() bool {
	for _, vv := range []string{
		"EQUALS",
		"NOT_EQUALS",
		"IN",
		"NOT_IN",
		"GREATER_THAN",
		"GREATER_THAN_EQUALS",
		"LESS_THAN",
		"LESS_THAN_EQUALS",
		"STARTS_WITH",
		"STARTS_WITH_IGNORE_CASE",
		"CONTAINS",
		"CONTAINS_IGNORE_CASE",
		"DOES_NOT_CONTAIN",
		"DOES_NOT_CONTAIN_IGNORE_CASE",
		"CONTAINS_ANY",
		"CONTAINS_ALL",
		"CONTAINS_NONE",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// QueryErrorReason was auto-generated from WSDL.
type QueryErrorReason string

// Validate validates QueryErrorReason.
func (v QueryErrorReason) Validate() bool {
	for _, vv := range []string{
		"PARSING_FAILED",
		"MISSING_QUERY",
		"MISSING_SELECT_CLAUSE",
		"MISSING_FROM_CLAUSE",
		"INVALID_SELECT_CLAUSE",
		"INVALID_FROM_CLAUSE",
		"INVALID_WHERE_CLAUSE",
		"INVALID_ORDER_BY_CLAUSE",
		"INVALID_LIMIT_CLAUSE",
		"INVALID_START_INDEX_IN_LIMIT_CLAUSE",
		"INVALID_PAGE_SIZE_IN_LIMIT_CLAUSE",
		"INVALID_DURING_CLAUSE",
		"INVALID_MIN_DATE_IN_DURING_CLAUSE",
		"INVALID_MAX_DATE_IN_DURING_CLAUSE",
		"MAX_LESS_THAN_MIN_IN_DURING_CLAUSE",
		"VALIDATION_FAILED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// QuotaCheckErrorReason was auto-generated from WSDL.
type QuotaCheckErrorReason string

// Validate validates QuotaCheckErrorReason.
func (v QuotaCheckErrorReason) Validate() bool {
	for _, vv := range []string{
		"INVALID_TOKEN_HEADER",
		"ACCOUNT_DELINQUENT",
		"ACCOUNT_INACCESSIBLE",
		"ACCOUNT_INACTIVE",
		"INCOMPLETE_SIGNUP",
		"DEVELOPER_TOKEN_NOT_APPROVED",
		"TERMS_AND_CONDITIONS_NOT_SIGNED",
		"MONTHLY_BUDGET_REACHED",
		"QUOTA_EXCEEDED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// RangeErrorReason was auto-generated from WSDL.
type RangeErrorReason string

// Validate validates RangeErrorReason.
func (v RangeErrorReason) Validate() bool {
	for _, vv := range []string{
		"TOO_LOW",
		"TOO_HIGH",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// RateExceededErrorReason was auto-generated from WSDL.
type RateExceededErrorReason string

// Validate validates RateExceededErrorReason.
func (v RateExceededErrorReason) Validate() bool {
	for _, vv := range []string{
		"RATE_EXCEEDED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ReadOnlyErrorReason was auto-generated from WSDL.
type ReadOnlyErrorReason string

// Validate validates ReadOnlyErrorReason.
func (v ReadOnlyErrorReason) Validate() bool {
	for _, vv := range []string{
		"READ_ONLY",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// RejectedErrorReason was auto-generated from WSDL.
type RejectedErrorReason string

// Validate validates RejectedErrorReason.
func (v RejectedErrorReason) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN_VALUE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// RequestErrorReason was auto-generated from WSDL.
type RequestErrorReason string

// Validate validates RequestErrorReason.
func (v RequestErrorReason) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"INVALID_INPUT",
		"UNSUPPORTED_VERSION",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// RequiredErrorReason was auto-generated from WSDL.
type RequiredErrorReason string

// Validate validates RequiredErrorReason.
func (v RequiredErrorReason) Validate() bool {
	for _, vv := range []string{
		"REQUIRED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// RichMediaAdAdAttribute was auto-generated from WSDL.
type RichMediaAdAdAttribute string

// Validate validates RichMediaAdAdAttribute.
func (v RichMediaAdAdAttribute) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"ROLL_OVER_TO_EXPAND",
		"SSL",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// RichMediaAdRichMediaAdType was auto-generated from WSDL.
type RichMediaAdRichMediaAdType string

// Validate validates RichMediaAdRichMediaAdType.
func (v RichMediaAdRichMediaAdType) Validate() bool {
	for _, vv := range []string{
		"STANDARD",
		"IN_STREAM_VIDEO",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// SelectorErrorReason was auto-generated from WSDL.
type SelectorErrorReason string

// Validate validates SelectorErrorReason.
func (v SelectorErrorReason) Validate() bool {
	for _, vv := range []string{
		"INVALID_FIELD_NAME",
		"MISSING_FIELDS",
		"MISSING_PREDICATES",
		"OPERATOR_DOES_NOT_SUPPORT_MULTIPLE_VALUES",
		"INVALID_PREDICATE_ENUM_VALUE",
		"MISSING_PREDICATE_OPERATOR",
		"MISSING_PREDICATE_VALUES",
		"INVALID_PREDICATE_FIELD_NAME",
		"INVALID_PREDICATE_OPERATOR",
		"INVALID_FIELD_SELECTION",
		"INVALID_PREDICATE_VALUE",
		"INVALID_SORT_FIELD_NAME",
		"SELECTOR_ERROR",
		"FILTER_BY_DATE_RANGE_NOT_SUPPORTED",
		"START_INDEX_IS_TOO_HIGH",
		"TOO_MANY_PREDICATE_VALUES",
		"UNKNOWN_ERROR",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ServedAssetFieldType was auto-generated from WSDL.
type ServedAssetFieldType string

// Validate validates ServedAssetFieldType.
func (v ServedAssetFieldType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"NONE",
		"HEADLINE_1",
		"HEADLINE_2",
		"HEADLINE_3",
		"DESCRIPTION_1",
		"DESCRIPTION_2",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// SizeLimitErrorReason was auto-generated from WSDL.
type SizeLimitErrorReason string

// Validate validates SizeLimitErrorReason.
func (v SizeLimitErrorReason) Validate() bool {
	for _, vv := range []string{
		"REQUEST_SIZE_LIMIT_EXCEEDED",
		"RESPONSE_SIZE_LIMIT_EXCEEDED",
		"INTERNAL_STORAGE_ERROR",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// SortOrder was auto-generated from WSDL.
type SortOrder string

// Validate validates SortOrder.
func (v SortOrder) Validate() bool {
	for _, vv := range []string{
		"ASCENDING",
		"DESCENDING",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// StatsQueryErrorReason was auto-generated from WSDL.
type StatsQueryErrorReason string

// Validate validates StatsQueryErrorReason.
func (v StatsQueryErrorReason) Validate() bool {
	for _, vv := range []string{
		"DATE_NOT_IN_VALID_RANGE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// StringFormatErrorReason was auto-generated from WSDL.
type StringFormatErrorReason string

// Validate validates StringFormatErrorReason.
func (v StringFormatErrorReason) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"ILLEGAL_CHARS",
		"INVALID_FORMAT",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// StringLengthErrorReason was auto-generated from WSDL.
type StringLengthErrorReason string

// Validate validates StringLengthErrorReason.
func (v StringLengthErrorReason) Validate() bool {
	for _, vv := range []string{
		"TOO_SHORT",
		"TOO_LONG",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// SystemManagedEntitySource was auto-generated from WSDL.
type SystemManagedEntitySource string

// Validate validates SystemManagedEntitySource.
func (v SystemManagedEntitySource) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"AD_VARIATIONS",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// TemplateElementFieldType was auto-generated from WSDL.
type TemplateElementFieldType string

// Validate validates TemplateElementFieldType.
func (v TemplateElementFieldType) Validate() bool {
	for _, vv := range []string{
		"ADDRESS",
		"AUDIO",
		"ENUM",
		"IMAGE",
		"BACKGROUND_IMAGE",
		"NUMBER",
		"TEXT",
		"URL",
		"VIDEO",
		"VISIBLE_URL",
		"MEDIA_BUNDLE",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ThirdPartyRedirectAdExpandingDirection was auto-generated from
// WSDL.
type ThirdPartyRedirectAdExpandingDirection string

// Validate validates ThirdPartyRedirectAdExpandingDirection.
func (v ThirdPartyRedirectAdExpandingDirection) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"EXPANDING_UP",
		"EXPANDING_DOWN",
		"EXPANDING_LEFT",
		"EXPANDING_RIGHT",
		"EXPANDING_UPLEFT",
		"EXPANDING_UPRIGHT",
		"EXPANDING_DOWNLEFT",
		"EXPANDING_DOWNRIGHT",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// UrlErrorReason was auto-generated from WSDL.
type UrlErrorReason string

// Validate validates UrlErrorReason.
func (v UrlErrorReason) Validate() bool {
	for _, vv := range []string{
		"INVALID_TRACKING_URL_TEMPLATE",
		"INVALID_TAG_IN_TRACKING_URL_TEMPLATE",
		"MISSING_TRACKING_URL_TEMPLATE_TAG",
		"MISSING_PROTOCOL_IN_TRACKING_URL_TEMPLATE",
		"INVALID_PROTOCOL_IN_TRACKING_URL_TEMPLATE",
		"MALFORMED_TRACKING_URL_TEMPLATE",
		"MISSING_HOST_IN_TRACKING_URL_TEMPLATE",
		"INVALID_TLD_IN_TRACKING_URL_TEMPLATE",
		"REDUNDANT_NESTED_TRACKING_URL_TEMPLATE_TAG",
		"INVALID_FINAL_URL",
		"INVALID_TAG_IN_FINAL_URL",
		"REDUNDANT_NESTED_FINAL_URL_TAG",
		"MISSING_PROTOCOL_IN_FINAL_URL",
		"INVALID_PROTOCOL_IN_FINAL_URL",
		"MALFORMED_FINAL_URL",
		"MISSING_HOST_IN_FINAL_URL",
		"INVALID_TLD_IN_FINAL_URL",
		"INVALID_FINAL_MOBILE_URL",
		"INVALID_TAG_IN_FINAL_MOBILE_URL",
		"REDUNDANT_NESTED_FINAL_MOBILE_URL_TAG",
		"MISSING_PROTOCOL_IN_FINAL_MOBILE_URL",
		"INVALID_PROTOCOL_IN_FINAL_MOBILE_URL",
		"MALFORMED_FINAL_MOBILE_URL",
		"MISSING_HOST_IN_FINAL_MOBILE_URL",
		"INVALID_TLD_IN_FINAL_MOBILE_URL",
		"INVALID_FINAL_APP_URL",
		"INVALID_TAG_IN_FINAL_APP_URL",
		"REDUNDANT_NESTED_FINAL_APP_URL_TAG",
		"MULTIPLE_APP_URLS_FOR_OSTYPE",
		"INVALID_OSTYPE",
		"INVALID_PROTOCOL_FOR_APP_URL",
		"INVALID_PACKAGE_ID_FOR_APP_URL",
		"URL_CUSTOM_PARAMETERS_COUNT_EXCEEDS_LIMIT",
		"URL_CUSTOM_PARAMETER_REMOVAL_WITH_NON_NULL_VALUE",
		"CANNOT_REMOVE_URL_CUSTOM_PARAMETER_IN_ADD_OPERATION",
		"CANNOT_REMOVE_URL_CUSTOM_PARAMETER_DURING_FULL_REPLACEMENT",
		"FINAL_URL_SUFFIX_MALFORMED",
		"INVALID_TAG_IN_FINAL_URL_SUFFIX",
		"NULL_CUSTOM_PARAMETER_VALUE_DURING_ADD_OR_FULL_REPLACEMENT",
		"INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_KEY",
		"INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_VALUE",
		"INVALID_TAG_IN_URL_CUSTOM_PARAMETER_VALUE",
		"REDUNDANT_NESTED_URL_CUSTOM_PARAMETER_TAG",
		"MISSING_PROTOCOL",
		"INVALID_URL",
		"DESTINATION_URL_DEPRECATED",
		"INVALID_TAG_IN_URL",
		"MISSING_URL_TAG",
		"DUPLICATE_URL_ID",
		"INVALID_URL_ID",
		"URL_ERROR",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// VideoType was auto-generated from WSDL.
type VideoType string

// Validate validates VideoType.
func (v VideoType) Validate() bool {
	for _, vv := range []string{
		"ADOBE",
		"REALPLAYER",
		"QUICKTIME",
		"WINDOWSMEDIA",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// The base class of all ad types. {@code Ad}s are created using
// the {@code AdGroupAdService}.             Some ad types such
// as {@code ExpandedTextAd}s may be modified using the {@code
// AdService}.                          <p>When calling {@code
// AdGroupAdService} to update the {@code status} of an {@code
// AdGroupAd},             you can construct an {@code Ad} object
// (instead of the {@code Ad}'s concrete type) with the
//      {@link #id} field set.             <span class="constraint
// AdxEnabled">This is enabled for AdX.</span>
type Ad struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
}

// An error indicating a problem with an ad customizer tag.
type AdCustomizerError struct {
	FieldPath         *string                  `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement      `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                  `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                  `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                  `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *AdCustomizerErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	FunctionString    *string                  `xml:"functionString,omitempty" json:"functionString,omitempty" yaml:"functionString,omitempty"`
	OperatorName      *string                  `xml:"operatorName,omitempty" json:"operatorName,omitempty" yaml:"operatorName,omitempty"`
	OperandIndex      *int                     `xml:"operandIndex,omitempty" json:"operandIndex,omitempty" yaml:"operandIndex,omitempty"`
	OperandValue      *string                  `xml:"operandValue,omitempty" json:"operandValue,omitempty" yaml:"operandValue,omitempty"`
	TypeAttrXSI       string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdCustomizerError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdCustomizerError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Base error class for Ad Service.
type AdError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *AdErrorReason      `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents an ad in an ad group.
type AdGroupAd struct {
	AdGroupId               *int64                   `xml:"adGroupId,omitempty" json:"adGroupId,omitempty" yaml:"adGroupId,omitempty"`
	Ad                      *Ad                      `xml:"ad,omitempty" json:"ad,omitempty" yaml:"ad,omitempty"`
	Status                  *AdGroupAdStatus         `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	PolicySummary           *AdGroupAdPolicySummary  `xml:"policySummary,omitempty" json:"policySummary,omitempty" yaml:"policySummary,omitempty"`
	Labels                  []*Label                 `xml:"labels,omitempty" json:"labels,omitempty" yaml:"labels,omitempty"`
	BaseCampaignId          *int64                   `xml:"baseCampaignId,omitempty" json:"baseCampaignId,omitempty" yaml:"baseCampaignId,omitempty"`
	BaseAdGroupId           *int64                   `xml:"baseAdGroupId,omitempty" json:"baseAdGroupId,omitempty" yaml:"baseAdGroupId,omitempty"`
	ForwardCompatibilityMap []*String_StringMapEntry `xml:"forwardCompatibilityMap,omitempty" json:"forwardCompatibilityMap,omitempty" yaml:"forwardCompatibilityMap,omitempty"`
	AdStrengthInfo          *AdStrengthInfo          `xml:"adStrengthInfo,omitempty" json:"adStrengthInfo,omitempty" yaml:"adStrengthInfo,omitempty"`
}

// Indicates too many ads were added/enabled under the specified
// adgroup.
type AdGroupAdCountLimitExceeded struct {
	FieldPath         *string                         `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement             `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                         `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                         `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                         `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *EntityCountLimitExceededReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	EnclosingId       *string                         `xml:"enclosingId,omitempty" json:"enclosingId,omitempty" yaml:"enclosingId,omitempty"`
	Limit             *int                            `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
	AccountLimitType  *string                         `xml:"accountLimitType,omitempty" json:"accountLimitType,omitempty" yaml:"accountLimitType,omitempty"`
	ExistingCount     *int                            `xml:"existingCount,omitempty" json:"existingCount,omitempty" yaml:"existingCount,omitempty"`
	TypeAttrXSI       string                          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupAdCountLimitExceeded) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupAdCountLimitExceeded"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Error information for AdGroupAdService.
type AdGroupAdError struct {
	FieldPath         *string               `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement   `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string               `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string               `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string               `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *AdGroupAdErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupAdError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupAdError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Manages the labels associated with an AdGroupAd.
type AdGroupAdLabel struct {
	AdGroupId *int64 `xml:"adGroupId,omitempty" json:"adGroupId,omitempty" yaml:"adGroupId,omitempty"`
	AdId      *int64 `xml:"adId,omitempty" json:"adId,omitempty" yaml:"adId,omitempty"`
	LabelId   *int64 `xml:"labelId,omitempty" json:"labelId,omitempty" yaml:"labelId,omitempty"`
}

// Operations for adding/removing labels from AdGroupAds.
type AdGroupAdLabelOperation struct {
	Operator      *Operator       `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	OperationType *string         `xml:"Operation.Type,omitempty" json:"Operation.Type,omitempty" yaml:"Operation.Type,omitempty"`
	Operand       *AdGroupAdLabel `xml:"operand,omitempty" json:"operand,omitempty" yaml:"operand,omitempty"`
	TypeAttrXSI   string          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupAdLabelOperation) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupAdLabelOperation"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A container for return values from the {@link AdGroupAdService#mutateLabel}
// call.
type AdGroupAdLabelReturnValue struct {
	ListReturnValueType  *string           `xml:"ListReturnValue.Type,omitempty" json:"ListReturnValue.Type,omitempty" yaml:"ListReturnValue.Type,omitempty"`
	Value                []*AdGroupAdLabel `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	PartialFailureErrors []*ApiError       `xml:"partialFailureErrors,omitempty" json:"partialFailureErrors,omitempty" yaml:"partialFailureErrors,omitempty"`
	TypeAttrXSI          string            `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string            `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupAdLabelReturnValue) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupAdLabelReturnValue"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// AdGroupAd service operations.
type AdGroupAdOperation struct {
	Operator                *Operator           `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	OperationType           *string             `xml:"Operation.Type,omitempty" json:"Operation.Type,omitempty" yaml:"Operation.Type,omitempty"`
	Operand                 *AdGroupAd          `xml:"operand,omitempty" json:"operand,omitempty" yaml:"operand,omitempty"`
	ExemptionRequests       []*ExemptionRequest `xml:"exemptionRequests,omitempty" json:"exemptionRequests,omitempty" yaml:"exemptionRequests,omitempty"`
	IgnorablePolicyTopicIds []*string           `xml:"ignorablePolicyTopicIds,omitempty" json:"ignorablePolicyTopicIds,omitempty" yaml:"ignorablePolicyTopicIds,omitempty"`
	TypeAttrXSI             string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace           string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupAdOperation) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupAdOperation"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a page of {@link AdGroupAd}s resulting from the query
// done by             {@link AdGroupAdService}.
type AdGroupAdPage struct {
	TotalNumEntries *int         `xml:"totalNumEntries,omitempty" json:"totalNumEntries,omitempty" yaml:"totalNumEntries,omitempty"`
	PageType        *string      `xml:"Page.Type,omitempty" json:"Page.Type,omitempty" yaml:"Page.Type,omitempty"`
	Entries         []*AdGroupAd `xml:"entries,omitempty" json:"entries,omitempty" yaml:"entries,omitempty"`
	TypeAttrXSI     string       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupAdPage) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupAdPage"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Contains policy information for an ad.
type AdGroupAdPolicySummary struct {
	PolicyTopicEntries     []*PolicyTopicEntry              `xml:"policyTopicEntries,omitempty" json:"policyTopicEntries,omitempty" yaml:"policyTopicEntries,omitempty"`
	ReviewState            *PolicySummaryReviewState        `xml:"reviewState,omitempty" json:"reviewState,omitempty" yaml:"reviewState,omitempty"`
	DenormalizedStatus     *PolicySummaryDenormalizedStatus `xml:"denormalizedStatus,omitempty" json:"denormalizedStatus,omitempty" yaml:"denormalizedStatus,omitempty"`
	CombinedApprovalStatus *PolicyApprovalStatus            `xml:"combinedApprovalStatus,omitempty" json:"combinedApprovalStatus,omitempty" yaml:"combinedApprovalStatus,omitempty"`
}

// A container for return values from the AdGroupAdService.
type AdGroupAdReturnValue struct {
	ListReturnValueType  *string      `xml:"ListReturnValue.Type,omitempty" json:"ListReturnValue.Type,omitempty" yaml:"ListReturnValue.Type,omitempty"`
	Value                []*AdGroupAd `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	PartialFailureErrors []*ApiError  `xml:"partialFailureErrors,omitempty" json:"partialFailureErrors,omitempty" yaml:"partialFailureErrors,omitempty"`
	TypeAttrXSI          string       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupAdReturnValue) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupAdReturnValue"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors related using the AdGroupAdService to create an AdGroupAd
//             with a reference to an existing AdId.
type AdSharingError struct {
	FieldPath         *string               `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement   `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string               `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string               `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string               `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *AdSharingErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	SharedAdError     *ApiError             `xml:"sharedAdError,omitempty" json:"sharedAdError,omitempty" yaml:"sharedAdError,omitempty"`
	TypeAttrXSI       string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdSharingError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdSharingError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// The strength info for this ad. This contains the overall ad
// strength and recommendations to             improve the strength.
type AdStrengthInfo struct {
	AdStrength  *AdStrength `xml:"adStrength,omitempty" json:"adStrength,omitempty" yaml:"adStrength,omitempty"`
	ActionItems []*string   `xml:"actionItems,omitempty" json:"actionItems,omitempty" yaml:"actionItems,omitempty"`
}

// Represents an id indicating a grouping of Ads under some heuristic.
type AdUnionId struct {
	Id            *int64  `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	AdUnionIdType *string `xml:"AdUnionId.Type,omitempty" json:"AdUnionId.Type,omitempty" yaml:"AdUnionId.Type,omitempty"`
}

// Errors that are thrown when a non-AdX feature is accessed by
// an AdX customer.
type AdxError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *AdxErrorReason     `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdxError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdxError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// The API error base class that provides details about an error
// that occurred             while processing a service request.
//                          <p>The OGNL field path is provided
// for parsers to identify the request data             element
// that may have caused the error.</p>
type ApiError interface{}

// Exception class for holding a list of service errors.
type ApiException struct {
	Message                  *string     `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
	ApplicationExceptionType *string     `xml:"ApplicationException.Type,omitempty" json:"ApplicationException.Type,omitempty" yaml:"ApplicationException.Type,omitempty"`
	Errors                   []*ApiError `xml:"errors,omitempty" json:"errors,omitempty" yaml:"errors,omitempty"`
	TypeAttrXSI              string      `xml:"xsi:type,attr,omitempty"`
	TypeNamespace            string      `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ApiException) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ApiException"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A URL for deep linking into an app for the given operating system.
type AppUrl struct {
	Url    *string       `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	OsType *AppUrlOsType `xml:"osType,omitempty" json:"osType,omitempty" yaml:"osType,omitempty"`
}

// Base class for exceptions.
type ApplicationException struct {
	Message                  *string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
	ApplicationExceptionType *string `xml:"ApplicationException.Type,omitempty" json:"ApplicationException.Type,omitempty" yaml:"ApplicationException.Type,omitempty"`
}

// Represents an asset.
type Asset struct {
	AssetId      *int64       `xml:"assetId,omitempty" json:"assetId,omitempty" yaml:"assetId,omitempty"`
	AssetName    *string      `xml:"assetName,omitempty" json:"assetName,omitempty" yaml:"assetName,omitempty"`
	AssetSubtype *AssetType   `xml:"assetSubtype,omitempty" json:"assetSubtype,omitempty" yaml:"assetSubtype,omitempty"`
	AssetStatus  *AssetStatus `xml:"assetStatus,omitempty" json:"assetStatus,omitempty" yaml:"assetStatus,omitempty"`
	AssetType    *string      `xml:"Asset.Type,omitempty" json:"Asset.Type,omitempty" yaml:"Asset.Type,omitempty"`
}

// Represents an error in an Asset.
type AssetError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *AssetErrorReason   `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AssetError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AssetError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents an asset link. This class contains an asset and information
// that is specific to an             asset-entity link (e.g. policy
// information).
type AssetLink struct {
	Asset                  *Asset                  `xml:"asset,omitempty" json:"asset,omitempty" yaml:"asset,omitempty"`
	PinnedField            *ServedAssetFieldType   `xml:"pinnedField,omitempty" json:"pinnedField,omitempty" yaml:"pinnedField,omitempty"`
	AssetPolicySummaryInfo *AssetPolicySummaryInfo `xml:"assetPolicySummaryInfo,omitempty" json:"assetPolicySummaryInfo,omitempty" yaml:"assetPolicySummaryInfo,omitempty"`
	AssetPerformanceLabel  *AssetPerformanceLabel  `xml:"assetPerformanceLabel,omitempty" json:"assetPerformanceLabel,omitempty" yaml:"assetPerformanceLabel,omitempty"`
}

// Represents an error in an AssetLink
type AssetLinkError struct {
	FieldPath         *string               `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement   `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string               `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string               `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string               `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *AssetLinkErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AssetLinkError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AssetLinkError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Policy summary information attached to an asset-entity link.
type AssetPolicySummaryInfo struct {
	PolicyTopicEntries     []*PolicyTopicEntry              `xml:"policyTopicEntries,omitempty" json:"policyTopicEntries,omitempty" yaml:"policyTopicEntries,omitempty"`
	ReviewState            *PolicySummaryReviewState        `xml:"reviewState,omitempty" json:"reviewState,omitempty" yaml:"reviewState,omitempty"`
	DenormalizedStatus     *PolicySummaryDenormalizedStatus `xml:"denormalizedStatus,omitempty" json:"denormalizedStatus,omitempty" yaml:"denormalizedStatus,omitempty"`
	CombinedApprovalStatus *PolicyApprovalStatus            `xml:"combinedApprovalStatus,omitempty" json:"combinedApprovalStatus,omitempty" yaml:"combinedApprovalStatus,omitempty"`
	PolicySummaryInfoType  *string                          `xml:"PolicySummaryInfo.Type,omitempty" json:"PolicySummaryInfo.Type,omitempty" yaml:"PolicySummaryInfo.Type,omitempty"`
	TypeAttrXSI            string                           `xml:"xsi:type,attr,omitempty"`
	TypeNamespace          string                           `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AssetPolicySummaryInfo) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AssetPolicySummaryInfo"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Encapsulates an Audio media identified by a MediaId.
type Audio struct {
	MediaId             *int64                           `xml:"mediaId,omitempty" json:"mediaId,omitempty" yaml:"mediaId,omitempty"`
	Type                *MediaMediaType                  `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	ReferenceId         *int64                           `xml:"referenceId,omitempty" json:"referenceId,omitempty" yaml:"referenceId,omitempty"`
	Dimensions          []*Media_Size_DimensionsMapEntry `xml:"dimensions,omitempty" json:"dimensions,omitempty" yaml:"dimensions,omitempty"`
	Urls                []*Media_Size_StringMapEntry     `xml:"urls,omitempty" json:"urls,omitempty" yaml:"urls,omitempty"`
	MimeType            *MediaLegacyMimeType             `xml:"mimeType,omitempty" json:"mimeType,omitempty" yaml:"mimeType,omitempty"`
	SourceUrl           *string                          `xml:"sourceUrl,omitempty" json:"sourceUrl,omitempty" yaml:"sourceUrl,omitempty"`
	Name                *string                          `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	FileSize            *int64                           `xml:"fileSize,omitempty" json:"fileSize,omitempty" yaml:"fileSize,omitempty"`
	CreationTime        *string                          `xml:"creationTime,omitempty" json:"creationTime,omitempty" yaml:"creationTime,omitempty"`
	MediaType           *string                          `xml:"Media.Type,omitempty" json:"Media.Type,omitempty" yaml:"Media.Type,omitempty"`
	DurationMillis      *int64                           `xml:"durationMillis,omitempty" json:"durationMillis,omitempty" yaml:"durationMillis,omitempty"`
	StreamingUrl        *string                          `xml:"streamingUrl,omitempty" json:"streamingUrl,omitempty" yaml:"streamingUrl,omitempty"`
	ReadyToPlayOnTheWeb *bool                            `xml:"readyToPlayOnTheWeb,omitempty" json:"readyToPlayOnTheWeb,omitempty" yaml:"readyToPlayOnTheWeb,omitempty"`
	TypeAttrXSI         string                           `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string                           `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Audio) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Audio"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors returned when Authentication failed.
type AuthenticationError struct {
	FieldPath         *string                    `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement        `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                    `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                    `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                    `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *AuthenticationErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AuthenticationError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AuthenticationError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors encountered when trying to authorize a user.
type AuthorizationError struct {
	FieldPath         *string                   `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement       `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                   `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                   `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                   `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *AuthorizationErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AuthorizationError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AuthorizationError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a CallOnlyAd.                          <p class="caution"><b>Caution:</b>
// Call only ads do not use {@link #url url},             {@link
// #finalUrls finalUrls}, {@link #finalMobileUrls finalMobileUrls},
//             {@link #finalAppUrls finalAppUrls}, {@link #urlCustomParameters
// urlCustomParameters},             or {@link #trackingUrlTemplate
// trackingUrlTemplate};             setting these fields on a
// call only ad will cause an error.             <span class="constraint
// AdxEnabled">This is enabled for AdX.</span>
type CallOnlyAd struct {
	Id                         *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                        *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                 *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                  []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls            []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls               []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate        *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix             *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters        *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                    []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                  *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                       *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference           *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource  *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                     *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	CountryCode                *string                    `xml:"countryCode,omitempty" json:"countryCode,omitempty" yaml:"countryCode,omitempty"`
	PhoneNumber                *string                    `xml:"phoneNumber,omitempty" json:"phoneNumber,omitempty" yaml:"phoneNumber,omitempty"`
	BusinessName               *string                    `xml:"businessName,omitempty" json:"businessName,omitempty" yaml:"businessName,omitempty"`
	Description1               *string                    `xml:"description1,omitempty" json:"description1,omitempty" yaml:"description1,omitempty"`
	Description2               *string                    `xml:"description2,omitempty" json:"description2,omitempty" yaml:"description2,omitempty"`
	CallTracked                *bool                      `xml:"callTracked,omitempty" json:"callTracked,omitempty" yaml:"callTracked,omitempty"`
	DisableCallConversion      *bool                      `xml:"disableCallConversion,omitempty" json:"disableCallConversion,omitempty" yaml:"disableCallConversion,omitempty"`
	ConversionTypeId           *int64                     `xml:"conversionTypeId,omitempty" json:"conversionTypeId,omitempty" yaml:"conversionTypeId,omitempty"`
	PhoneNumberVerificationUrl *string                    `xml:"phoneNumberVerificationUrl,omitempty" json:"phoneNumberVerificationUrl,omitempty" yaml:"phoneNumberVerificationUrl,omitempty"`
	TypeAttrXSI                string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace              string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CallOnlyAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CallOnlyAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Indicates that a certificate does not include the correct domain.
type CertificateDomainMismatchConstraint struct {
	ConstraintType            *PolicyTopicConstraintPolicyTopicConstraintType `xml:"constraintType,omitempty" json:"constraintType,omitempty" yaml:"constraintType,omitempty"`
	PolicyTopicConstraintType *string                                         `xml:"PolicyTopicConstraint.Type,omitempty" json:"PolicyTopicConstraint.Type,omitempty" yaml:"PolicyTopicConstraint.Type,omitempty"`
	TypeAttrXSI               string                                          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                                          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CertificateDomainMismatchConstraint) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CertificateDomainMismatchConstraint"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Information about countries that were targeted, but the certificate
// for those countries does not             include the correct
// domain.
type CertificateDomainMismatchInCountryConstraint struct {
	ConstraintType            *PolicyTopicConstraintPolicyTopicConstraintType `xml:"constraintType,omitempty" json:"constraintType,omitempty" yaml:"constraintType,omitempty"`
	PolicyTopicConstraintType *string                                         `xml:"PolicyTopicConstraint.Type,omitempty" json:"PolicyTopicConstraint.Type,omitempty" yaml:"PolicyTopicConstraint.Type,omitempty"`
	ConstrainedCountries      []*int64                                        `xml:"constrainedCountries,omitempty" json:"constrainedCountries,omitempty" yaml:"constrainedCountries,omitempty"`
	TotalTargetedCountries    *int                                            `xml:"totalTargetedCountries,omitempty" json:"totalTargetedCountries,omitempty" yaml:"totalTargetedCountries,omitempty"`
	TypeAttrXSI               string                                          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                                          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CertificateDomainMismatchInCountryConstraint) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CertificateDomainMismatchInCountryConstraint"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A policy finding was caused by a missing certificate.
type CertificateMissingConstraint struct {
	ConstraintType            *PolicyTopicConstraintPolicyTopicConstraintType `xml:"constraintType,omitempty" json:"constraintType,omitempty" yaml:"constraintType,omitempty"`
	PolicyTopicConstraintType *string                                         `xml:"PolicyTopicConstraint.Type,omitempty" json:"PolicyTopicConstraint.Type,omitempty" yaml:"PolicyTopicConstraint.Type,omitempty"`
	TypeAttrXSI               string                                          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                                          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CertificateMissingConstraint) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CertificateMissingConstraint"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Information about countries that were targeted but that do not
// have a policy certificate.
type CertificateMissingInCountryConstraint struct {
	ConstraintType            *PolicyTopicConstraintPolicyTopicConstraintType `xml:"constraintType,omitempty" json:"constraintType,omitempty" yaml:"constraintType,omitempty"`
	PolicyTopicConstraintType *string                                         `xml:"PolicyTopicConstraint.Type,omitempty" json:"PolicyTopicConstraint.Type,omitempty" yaml:"PolicyTopicConstraint.Type,omitempty"`
	ConstrainedCountries      []*int64                                        `xml:"constrainedCountries,omitempty" json:"constrainedCountries,omitempty" yaml:"constrainedCountries,omitempty"`
	TotalTargetedCountries    *int                                            `xml:"totalTargetedCountries,omitempty" json:"totalTargetedCountries,omitempty" yaml:"totalTargetedCountries,omitempty"`
	TypeAttrXSI               string                                          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                                          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CertificateMissingInCountryConstraint) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CertificateMissingInCountryConstraint"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Error due to user not accepting the AdWords terms of service.
type ClientTermsError struct {
	FieldPath         *string                 `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement     `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                 `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                 `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                 `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *ClientTermsErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ClientTermsError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ClientTermsError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors associated with the size of the given collection being
//             out of bounds.
type CollectionSizeError struct {
	FieldPath         *string                    `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement        `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                    `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                    `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                    `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *CollectionSizeErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CollectionSizeError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CollectionSizeError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Information about countries that were targeted that caused a
// policy finding. If a             CountryConstraint has 0 targeted
// countries and an empty list of constrained countries, it means
//             that the constraint applies globally.
type CountryConstraint struct {
	ConstraintType            *PolicyTopicConstraintPolicyTopicConstraintType `xml:"constraintType,omitempty" json:"constraintType,omitempty" yaml:"constraintType,omitempty"`
	PolicyTopicConstraintType *string                                         `xml:"PolicyTopicConstraint.Type,omitempty" json:"PolicyTopicConstraint.Type,omitempty" yaml:"PolicyTopicConstraint.Type,omitempty"`
	ConstrainedCountries      []*int64                                        `xml:"constrainedCountries,omitempty" json:"constrainedCountries,omitempty" yaml:"constrainedCountries,omitempty"`
	TotalTargetedCountries    *int                                            `xml:"totalTargetedCountries,omitempty" json:"totalTargetedCountries,omitempty" yaml:"totalTargetedCountries,omitempty"`
	TypeAttrXSI               string                                          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                                          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CountryConstraint) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CountryConstraint"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// CustomParameter is used to map a custom parameter key to its
// value.
type CustomParameter struct {
	Key      *string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value    *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	IsRemove *bool   `xml:"isRemove,omitempty" json:"isRemove,omitempty" yaml:"isRemove,omitempty"`
}

// CustomParameters holds a list of CustomParameters to be treated
// as a map.             It has a special field used to indicate
// that the current map should be cleared and replaced
//     with this new map.
type CustomParameters struct {
	Parameters []*CustomParameter `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
	DoReplace  *bool              `xml:"doReplace,omitempty" json:"doReplace,omitempty" yaml:"doReplace,omitempty"`
}

// Errors that are thrown due to a database access problem.
type DatabaseError struct {
	FieldPath         *string              `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement  `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string              `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string              `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string              `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *DatabaseErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DatabaseError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DatabaseError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors associated with invalid dates and date ranges.
type DateError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *DateErrorReason    `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DateError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DateError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a range of dates that has either an upper or a lower
// bound.             The format for the date is YYYYMMDD.
type DateRange struct {
	Min *string `xml:"min,omitempty" json:"min,omitempty" yaml:"min,omitempty"`
	Max *string `xml:"max,omitempty" json:"max,omitempty" yaml:"max,omitempty"`
}

// Represents a deprecated ad.                          Deprecated
// ads can be deleted, but cannot be created.             <span
// class="constraint AdxEnabled">This is disabled for AdX when
// it is contained within Operators: ADD, SET.</span>
type DeprecatedAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	Name                      *string                    `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	DeprecatedAdType          *DeprecatedAdType          `xml:"deprecatedAdType,omitempty" json:"deprecatedAdType,omitempty" yaml:"deprecatedAdType,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DeprecatedAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DeprecatedAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a simple height-width dimension.
type Dimensions struct {
	Width  *int `xml:"width,omitempty" json:"width,omitempty" yaml:"width,omitempty"`
	Height *int `xml:"height,omitempty" json:"height,omitempty" yaml:"height,omitempty"`
}

// Attributes for Text Labels.
type DisplayAttribute struct {
	LabelAttributeType *string `xml:"LabelAttribute.Type,omitempty" json:"LabelAttribute.Type,omitempty" yaml:"LabelAttribute.Type,omitempty"`
	BackgroundColor    *string `xml:"backgroundColor,omitempty" json:"backgroundColor,omitempty" yaml:"backgroundColor,omitempty"`
	Description        *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	TypeAttrXSI        string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace      string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DisplayAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DisplayAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents display-call-to-action specific data.
type DisplayCallToAction struct {
	Text      *string `xml:"text,omitempty" json:"text,omitempty" yaml:"text,omitempty"`
	TextColor *string `xml:"textColor,omitempty" json:"textColor,omitempty" yaml:"textColor,omitempty"`
	UrlId     *string `xml:"urlId,omitempty" json:"urlId,omitempty" yaml:"urlId,omitempty"`
}

// Errors related to distinct ids or content.
type DistinctError struct {
	FieldPath         *string              `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement  `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string              `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string              `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string              `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *DistinctErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DistinctError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DistinctError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a dynamic search ad. This ad will have its headline
// and final URL auto-generated at             serving time according
// to domain name specific information provided by DynamicSearchAdsSetting
// at             the campaign level.
// <p>Auto-generated fields: headline and final URL.</p>
//                    <p>Note: we recommend using the ExpandedDynamicSearchAd
// type, introduced in v201705, rather than             the DynamicSearchAd
// type.</p>                          <p><b>Required fields:</b>
// {@code description1}, {@code description2}, {@code displayUrl}.</p>
//                          <p>The tracking URL field must contain
// at least one of the following placeholder tags             (URL
// parameters):</p>             <ul>             <li>{unescapedlpurl}</li>
//             <li>{escapedlpurl}</li>             <li>{lpurl}</li>
//             <li>{lpurl+2}</li>             <li>{lpurl+3}</li>
//             </ul>                          <ul>
// <li>{unescapedlpurl} will be replaced with the full landing
// page URL of the displayed ad.             Extra query parameters
// can be added to the end, e.g.: "{unescapedlpurl}?lang=en".</li>
//                          <li>{escapedlpurl} will be replaced
// with the URL-encoded version of the full             landing
// page URL. This makes it suitable for use as a query parameter
//             value (e.g.: "http://www.3rdpartytracker.com/?lp={escapedlpurl}")
// but             not at the beginning of the URL field.</li>
//                          <li>{lpurl} encodes the "?" and "="
// of the landing page URL making it suitable             for use
// as a query parameter. If found at the beginning of the URL field,
// it is             replaced by the {unescapedlpurl} value.
//           E.g.: "http://tracking.com/redir.php?tracking=xyz&url={lpurl}".</li>
//                          <li>{lpurl+2} and {lpurl+3}  will be
// replaced with the landing page URL escaped two or three
//         times, respectively.  This makes it suitable if there
// is a chain of redirects in the tracking             URL.</li>
//             </ul>                          <p class="note">Note
// that {@code finalUrls} and {@code finalMobileUrls}
//    cannot be set for dynamic search ads.</p>
//           <p>For more information, see the article
//    <a href="//support.google.com/adwords/answer/2549100">Using
// dynamic tracking URLs</a>.             </p>             <span
// class="constraint AdxEnabled">This is disabled for AdX when
// it is contained within Operators: ADD, SET.</span>
type DynamicSearchAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	Description1              *string                    `xml:"description1,omitempty" json:"description1,omitempty" yaml:"description1,omitempty"`
	Description2              *string                    `xml:"description2,omitempty" json:"description2,omitempty" yaml:"description2,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DynamicSearchAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DynamicSearchAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents Dynamic Settings.
type DynamicSettings struct {
	LandscapeLogoImage *Image  `xml:"landscapeLogoImage,omitempty" json:"landscapeLogoImage,omitempty" yaml:"landscapeLogoImage,omitempty"`
	PricePrefix        *string `xml:"pricePrefix,omitempty" json:"pricePrefix,omitempty" yaml:"pricePrefix,omitempty"`
	PromoText          *string `xml:"promoText,omitempty" json:"promoText,omitempty" yaml:"promoText,omitempty"`
}

// Reports permission problems trying to access an entity.
type EntityAccessDenied struct {
	FieldPath         *string                   `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement       `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                   `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                   `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                   `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *EntityAccessDeniedReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *EntityAccessDenied) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:EntityAccessDenied"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Signals that an entity count limit was exceeded for some level.
//             For example, too many criteria for a campaign.
type EntityCountLimitExceeded struct {
	FieldPath         *string                         `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement             `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                         `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                         `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                         `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *EntityCountLimitExceededReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	EnclosingId       *string                         `xml:"enclosingId,omitempty" json:"enclosingId,omitempty" yaml:"enclosingId,omitempty"`
	Limit             *int                            `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
	AccountLimitType  *string                         `xml:"accountLimitType,omitempty" json:"accountLimitType,omitempty" yaml:"accountLimitType,omitempty"`
	ExistingCount     *int                            `xml:"existingCount,omitempty" json:"existingCount,omitempty" yaml:"existingCount,omitempty"`
	TypeAttrXSI       string                          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *EntityCountLimitExceeded) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:EntityCountLimitExceeded"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// An id did not correspond to an entity, or it referred to an
// entity which does not belong to the             customer.
type EntityNotFound struct {
	FieldPath         *string               `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement   `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string               `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string               `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string               `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *EntityNotFoundReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *EntityNotFound) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:EntityNotFound"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A request to be exempted from a {@link PolicyViolationError}.
type ExemptionRequest struct {
	Key *PolicyViolationKey `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
}

// Represents an ExpandedDynamicSearchAd. This ad will have its
// headline, final URLs and display URL             auto-generated
// at serving time according to domain name specific information
// provided by             DynamicSearchAdsSetting linked at the
// campaign level.                          <p>Auto-generated fields:
// headline, final URLs and display URL.</p>
//        <p><b>Required fields:</b> {@code description}.</p>
//                         <p>The tracking URL field must contain
// at least one of the following placeholder tags             (URL
// parameters):</p>             <ul>             <li>{unescapedlpurl}</li>
//             <li>{escapedlpurl}</li>             <li>{lpurl}</li>
//             <li>{lpurl+2}</li>             <li>{lpurl+3}</li>
//             </ul>                          <ul>
// <li>{unescapedlpurl} will be replaced with the full landing
// page URL of the displayed ad.             Extra query parameters
// can be added to the end, e.g.: "{unescapedlpurl}?lang=en".</li>
//                          <li>{escapedlpurl} will be replaced
// with the URL-encoded version of the full             landing
// page URL. This makes it suitable for use as a query parameter
//             value (e.g.: "http://www.3rdpartytracker.com/?lp={escapedlpurl}")
// but             not at the beginning of the URL field.</li>
//                          <li>{lpurl} encodes the "?" and "="
// of the landing page URL making it suitable             for use
// as a query parameter. If found at the beginning of the URL field,
// it is             replaced by the {unescapedlpurl} value.
//           E.g.: "http://tracking.com/redir.php?tracking=xyz&url={lpurl}".</li>
//                          <li>{lpurl+2} and {lpurl+3}  will be
// replaced with the landing page URL escaped two or three
//         times, respectively.  This makes it suitable if there
// is a chain of redirects in the tracking             URL.</li>
//             </ul>             <span class="constraint AdxEnabled">This
// is enabled for AdX.</span>
type ExpandedDynamicSearchAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	Description               *string                    `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Description2              *string                    `xml:"description2,omitempty" json:"description2,omitempty" yaml:"description2,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ExpandedDynamicSearchAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ExpandedDynamicSearchAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Enhanced text ad format.                          <p class="caution"><b>Caution:</b>
// Expanded text ads do not use {@link #url url},             {@link
// #displayUrl displayUrl}, {@link #finalAppUrls finalAppUrls},
// or             {@link #devicePreference devicePreference};
//            setting these fields on an expanded text ad will
// cause an error.             <span class="constraint AdxEnabled">This
// is enabled for AdX.</span>
type ExpandedTextAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	HeadlinePart1             *string                    `xml:"headlinePart1,omitempty" json:"headlinePart1,omitempty" yaml:"headlinePart1,omitempty"`
	HeadlinePart2             *string                    `xml:"headlinePart2,omitempty" json:"headlinePart2,omitempty" yaml:"headlinePart2,omitempty"`
	HeadlinePart3             *string                    `xml:"headlinePart3,omitempty" json:"headlinePart3,omitempty" yaml:"headlinePart3,omitempty"`
	Description               *string                    `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Description2              *string                    `xml:"description2,omitempty" json:"description2,omitempty" yaml:"description2,omitempty"`
	Path1                     *string                    `xml:"path1,omitempty" json:"path1,omitempty" yaml:"path1,omitempty"`
	Path2                     *string                    `xml:"path2,omitempty" json:"path2,omitempty" yaml:"path2,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ExpandedTextAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ExpandedTextAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// An error indicating a problem with a reference to a feed attribute
// in an ad.
type FeedAttributeReferenceError struct {
	FieldPath         *string                            `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement                `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                            `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                            `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                            `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *FeedAttributeReferenceErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	FeedName          *string                            `xml:"feedName,omitempty" json:"feedName,omitempty" yaml:"feedName,omitempty"`
	FeedAttributeName *string                            `xml:"feedAttributeName,omitempty" json:"feedAttributeName,omitempty" yaml:"feedAttributeName,omitempty"`
	TypeAttrXSI       string                             `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                             `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *FeedAttributeReferenceError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:FeedAttributeReferenceError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A segment of a field path. Each dot in a field path defines
// a new segment.
type FieldPathElement struct {
	Field *string `xml:"field,omitempty" json:"field,omitempty" yaml:"field,omitempty"`
	Index *int    `xml:"index,omitempty" json:"index,omitempty" yaml:"index,omitempty"`
}

// A ForwardComptibilityError represents possible errors when using
// the forwardCompatibilityMap             in some of the common
// services.
type ForwardCompatibilityError struct {
	FieldPath         *string                          `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement              `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                          `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                          `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                          `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *ForwardCompatibilityErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                           `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                           `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ForwardCompatibilityError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ForwardCompatibilityError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors that indicate issues with the function.
type FunctionError struct {
	FieldPath         *string              `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement  `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string              `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string              `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string              `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *FunctionErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *FunctionError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:FunctionError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// An error resulting from a failure to parse the textual representation
// of a function.
type FunctionParsingError struct {
	FieldPath          *string                     `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements  []*FieldPathElement         `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger            *string                     `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString        *string                     `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType       *string                     `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason             *FunctionParsingErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	OffendingText      *string                     `xml:"offendingText,omitempty" json:"offendingText,omitempty" yaml:"offendingText,omitempty"`
	OffendingTextIndex *int                        `xml:"offendingTextIndex,omitempty" json:"offendingTextIndex,omitempty" yaml:"offendingTextIndex,omitempty"`
	TypeAttrXSI        string                      `xml:"xsi:type,attr,omitempty"`
	TypeNamespace      string                      `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *FunctionParsingError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:FunctionParsingError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents Gmail ad.                          <p class="caution"><b>Caution:</b>
// Gmail ads do not use {@link #url url}, {@link #displayUrl
//           displayUrl}, {@link #finalAppUrls finalAppUrls}, or
// {@link #devicePreference devicePreference};             Setting
// these fields on a Gmail ad will cause an error.
// <span class="constraint AdxEnabled">This is enabled for AdX.</span>
type GmailAd struct {
	Id                                *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                               *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                        *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                         []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls                   []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls                      []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate               *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix                    *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters               *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                           []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                         *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                              *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference                  *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource         *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                            *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	Teaser                            *GmailTeaser               `xml:"teaser,omitempty" json:"teaser,omitempty" yaml:"teaser,omitempty"`
	HeaderImage                       *Image                     `xml:"headerImage,omitempty" json:"headerImage,omitempty" yaml:"headerImage,omitempty"`
	MarketingImage                    *Image                     `xml:"marketingImage,omitempty" json:"marketingImage,omitempty" yaml:"marketingImage,omitempty"`
	MarketingImageHeadline            *string                    `xml:"marketingImageHeadline,omitempty" json:"marketingImageHeadline,omitempty" yaml:"marketingImageHeadline,omitempty"`
	MarketingImageDescription         *string                    `xml:"marketingImageDescription,omitempty" json:"marketingImageDescription,omitempty" yaml:"marketingImageDescription,omitempty"`
	MarketingImageDisplayCallToAction *DisplayCallToAction       `xml:"marketingImageDisplayCallToAction,omitempty" json:"marketingImageDisplayCallToAction,omitempty" yaml:"marketingImageDisplayCallToAction,omitempty"`
	ProductImages                     []*ProductImage            `xml:"productImages,omitempty" json:"productImages,omitempty" yaml:"productImages,omitempty"`
	ProductVideoList                  []*Video                   `xml:"productVideoList,omitempty" json:"productVideoList,omitempty" yaml:"productVideoList,omitempty"`
	TypeAttrXSI                       string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                     string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *GmailAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:GmailAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents Gmail teaser specific data.
type GmailTeaser struct {
	Headline     *string `xml:"headline,omitempty" json:"headline,omitempty" yaml:"headline,omitempty"`
	Description  *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	BusinessName *string `xml:"businessName,omitempty" json:"businessName,omitempty" yaml:"businessName,omitempty"`
	LogoImage    *Image  `xml:"logoImage,omitempty" json:"logoImage,omitempty" yaml:"logoImage,omitempty"`
}

// Represents a Smart Shopping ad that optimizes towards your goals.
// A Smart Shopping ad targets             multiple advertising
// channels across Search, Google Display Network, and YouTube
// with a focus on             retail. This supports ads that display
// product data (managed using the Google Merchant Center) as
//            specified in the parent campaign's {@linkplain ShoppingSetting
// Shopping setting} as well as ads             using advertiser
// provided asset data.                          <p class="caution"><b>Caution:</b>
// Smart Shopping ads do not use {@link #url url}, {@link
//        #finalUrls finalUrls}, {@link #finalMobileUrls finalMobileUrls},
// {@link #finalAppUrls             finalAppUrls}, or {@link #displayUrl
// displayUrl}; setting these fields on a Smart Shopping ad
//          will cause an error.             <span class="constraint
// AdxEnabled">This is enabled for AdX.</span>
type GoalOptimizedShoppingAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *GoalOptimizedShoppingAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:GoalOptimizedShoppingAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors associated with the ids.
type IdError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *IdErrorReason      `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *IdError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:IdError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Encapsulates an Image media. For {@code SET},{@code REMOVE}
// operations in             MediaService, use {@code mediaId}.
type Image struct {
	MediaId       *int64                           `xml:"mediaId,omitempty" json:"mediaId,omitempty" yaml:"mediaId,omitempty"`
	Type          *MediaMediaType                  `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	ReferenceId   *int64                           `xml:"referenceId,omitempty" json:"referenceId,omitempty" yaml:"referenceId,omitempty"`
	Dimensions    []*Media_Size_DimensionsMapEntry `xml:"dimensions,omitempty" json:"dimensions,omitempty" yaml:"dimensions,omitempty"`
	Urls          []*Media_Size_StringMapEntry     `xml:"urls,omitempty" json:"urls,omitempty" yaml:"urls,omitempty"`
	MimeType      *MediaLegacyMimeType             `xml:"mimeType,omitempty" json:"mimeType,omitempty" yaml:"mimeType,omitempty"`
	SourceUrl     *string                          `xml:"sourceUrl,omitempty" json:"sourceUrl,omitempty" yaml:"sourceUrl,omitempty"`
	Name          *string                          `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	FileSize      *int64                           `xml:"fileSize,omitempty" json:"fileSize,omitempty" yaml:"fileSize,omitempty"`
	CreationTime  *string                          `xml:"creationTime,omitempty" json:"creationTime,omitempty" yaml:"creationTime,omitempty"`
	MediaType     *string                          `xml:"Media.Type,omitempty" json:"Media.Type,omitempty" yaml:"Media.Type,omitempty"`
	Data          *[]byte                          `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	TypeAttrXSI   string                           `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                           `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Image) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Image"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents an ImageAd.             <span class="constraint AdxEnabled">This
// is enabled for AdX.</span>
type ImageAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	Image                     *Image                     `xml:"image,omitempty" json:"image,omitempty" yaml:"image,omitempty"`
	Name                      *string                    `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	AdToCopyImageFrom         *int64                     `xml:"adToCopyImageFrom,omitempty" json:"adToCopyImageFrom,omitempty" yaml:"adToCopyImageFrom,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ImageAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ImageAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents an image asset.
type ImageAsset struct {
	AssetId       *int64              `xml:"assetId,omitempty" json:"assetId,omitempty" yaml:"assetId,omitempty"`
	AssetName     *string             `xml:"assetName,omitempty" json:"assetName,omitempty" yaml:"assetName,omitempty"`
	AssetSubtype  *AssetType          `xml:"assetSubtype,omitempty" json:"assetSubtype,omitempty" yaml:"assetSubtype,omitempty"`
	AssetStatus   *AssetStatus        `xml:"assetStatus,omitempty" json:"assetStatus,omitempty" yaml:"assetStatus,omitempty"`
	AssetType     *string             `xml:"Asset.Type,omitempty" json:"Asset.Type,omitempty" yaml:"Asset.Type,omitempty"`
	ImageData     *[]byte             `xml:"imageData,omitempty" json:"imageData,omitempty" yaml:"imageData,omitempty"`
	ImageFileSize *int64              `xml:"imageFileSize,omitempty" json:"imageFileSize,omitempty" yaml:"imageFileSize,omitempty"`
	ImageMimeType *MediaMimeType      `xml:"imageMimeType,omitempty" json:"imageMimeType,omitempty" yaml:"imageMimeType,omitempty"`
	FullSizeInfo  *ImageDimensionInfo `xml:"fullSizeInfo,omitempty" json:"fullSizeInfo,omitempty" yaml:"fullSizeInfo,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ImageAsset) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ImageAsset"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Information about an image asset in specfic dimensions, either
// original or resized.
type ImageDimensionInfo struct {
	ImageHeight *int    `xml:"imageHeight,omitempty" json:"imageHeight,omitempty" yaml:"imageHeight,omitempty"`
	ImageWidth  *int    `xml:"imageWidth,omitempty" json:"imageWidth,omitempty" yaml:"imageWidth,omitempty"`
	ImageUrl    *string `xml:"imageUrl,omitempty" json:"imageUrl,omitempty" yaml:"imageUrl,omitempty"`
}

// Error class for errors associated with parsing image data.
type ImageError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *ImageErrorReason   `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ImageError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ImageError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Indicates that a server-side error has occured. {@code InternalApiError}s
//             are generally not the result of an invalid request
// or message sent by the             client.
type InternalApiError struct {
	FieldPath         *string                 `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement     `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                 `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                 `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                 `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *InternalApiErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *InternalApiError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:InternalApiError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a label that can be attached to entities such as
// campaign, ad group, ads,             criterion etc.
type Label struct {
	Id        *int64          `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Name      *string         `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Status    *LabelStatus    `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Attribute *LabelAttribute `xml:"attribute,omitempty" json:"attribute,omitempty" yaml:"attribute,omitempty"`
	LabelType *string         `xml:"Label.Type,omitempty" json:"Label.Type,omitempty" yaml:"Label.Type,omitempty"`
}

// Base type for AdWords label attributes.
type LabelAttribute struct {
	LabelAttributeType *string `xml:"LabelAttribute.Type,omitempty" json:"LabelAttribute.Type,omitempty" yaml:"LabelAttribute.Type,omitempty"`
}

// Base list return value type.
type ListReturnValue interface{}

// Represents some kind of media.
type Media struct {
	MediaId      *int64                           `xml:"mediaId,omitempty" json:"mediaId,omitempty" yaml:"mediaId,omitempty"`
	Type         *MediaMediaType                  `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	ReferenceId  *int64                           `xml:"referenceId,omitempty" json:"referenceId,omitempty" yaml:"referenceId,omitempty"`
	Dimensions   []*Media_Size_DimensionsMapEntry `xml:"dimensions,omitempty" json:"dimensions,omitempty" yaml:"dimensions,omitempty"`
	Urls         []*Media_Size_StringMapEntry     `xml:"urls,omitempty" json:"urls,omitempty" yaml:"urls,omitempty"`
	MimeType     *MediaLegacyMimeType             `xml:"mimeType,omitempty" json:"mimeType,omitempty" yaml:"mimeType,omitempty"`
	SourceUrl    *string                          `xml:"sourceUrl,omitempty" json:"sourceUrl,omitempty" yaml:"sourceUrl,omitempty"`
	Name         *string                          `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	FileSize     *int64                           `xml:"fileSize,omitempty" json:"fileSize,omitempty" yaml:"fileSize,omitempty"`
	CreationTime *string                          `xml:"creationTime,omitempty" json:"creationTime,omitempty" yaml:"creationTime,omitempty"`
	MediaType    *string                          `xml:"Media.Type,omitempty" json:"Media.Type,omitempty" yaml:"Media.Type,omitempty"`
}

// Represents a ZIP archive media the content of which contains
// HTML5 assets.
type MediaBundle struct {
	MediaId        *int64                           `xml:"mediaId,omitempty" json:"mediaId,omitempty" yaml:"mediaId,omitempty"`
	Type           *MediaMediaType                  `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	ReferenceId    *int64                           `xml:"referenceId,omitempty" json:"referenceId,omitempty" yaml:"referenceId,omitempty"`
	Dimensions     []*Media_Size_DimensionsMapEntry `xml:"dimensions,omitempty" json:"dimensions,omitempty" yaml:"dimensions,omitempty"`
	Urls           []*Media_Size_StringMapEntry     `xml:"urls,omitempty" json:"urls,omitempty" yaml:"urls,omitempty"`
	MimeType       *MediaLegacyMimeType             `xml:"mimeType,omitempty" json:"mimeType,omitempty" yaml:"mimeType,omitempty"`
	SourceUrl      *string                          `xml:"sourceUrl,omitempty" json:"sourceUrl,omitempty" yaml:"sourceUrl,omitempty"`
	Name           *string                          `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	FileSize       *int64                           `xml:"fileSize,omitempty" json:"fileSize,omitempty" yaml:"fileSize,omitempty"`
	CreationTime   *string                          `xml:"creationTime,omitempty" json:"creationTime,omitempty" yaml:"creationTime,omitempty"`
	MediaType      *string                          `xml:"Media.Type,omitempty" json:"Media.Type,omitempty" yaml:"Media.Type,omitempty"`
	Data           *[]byte                          `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	MediaBundleUrl *string                          `xml:"mediaBundleUrl,omitempty" json:"mediaBundleUrl,omitempty" yaml:"mediaBundleUrl,omitempty"`
	EntryPoint     *string                          `xml:"entryPoint,omitempty" json:"entryPoint,omitempty" yaml:"entryPoint,omitempty"`
	TypeAttrXSI    string                           `xml:"xsi:type,attr,omitempty"`
	TypeNamespace  string                           `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MediaBundle) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MediaBundle"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a media bundle asset.
type MediaBundleAsset struct {
	AssetId         *int64       `xml:"assetId,omitempty" json:"assetId,omitempty" yaml:"assetId,omitempty"`
	AssetName       *string      `xml:"assetName,omitempty" json:"assetName,omitempty" yaml:"assetName,omitempty"`
	AssetSubtype    *AssetType   `xml:"assetSubtype,omitempty" json:"assetSubtype,omitempty" yaml:"assetSubtype,omitempty"`
	AssetStatus     *AssetStatus `xml:"assetStatus,omitempty" json:"assetStatus,omitempty" yaml:"assetStatus,omitempty"`
	AssetType       *string      `xml:"Asset.Type,omitempty" json:"Asset.Type,omitempty" yaml:"Asset.Type,omitempty"`
	MediaBundleData *[]byte      `xml:"mediaBundleData,omitempty" json:"mediaBundleData,omitempty" yaml:"mediaBundleData,omitempty"`
	TypeAttrXSI     string       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MediaBundleAsset) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MediaBundleAsset"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Error class for errors associated with parsing media bundle
// data.
type MediaBundleError struct {
	FieldPath         *string                 `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement     `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                 `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                 `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                 `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *MediaBundleErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MediaBundleError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MediaBundleError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Error class for media related errors.
type MediaError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *MediaErrorReason   `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MediaError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MediaError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// This represents an entry in a map with a key of type Size
//           and value of type Dimensions.
type Media_Size_DimensionsMapEntry struct {
	Key   *MediaSize  `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value *Dimensions `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// This represents an entry in a map with a key of type Size
//           and value of type String.
type Media_Size_StringMapEntry struct {
	Key   *MediaSize `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value *string    `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// Representation of multi-asset responsive display ad format.
//                          <p class="caution"><b>Caution:</b>
// multi-asset responsive display ads do not use {@link #url
//           url}, {@link #displayUrl displayUrl}, {@link #finalAppUrls
// finalAppUrls}, or {@link             #devicePreference devicePreference};
// setting these fields on a multi-asset responsive display ad
//             will cause an error.             <span class="constraint
// AdxEnabled">This is enabled for AdX.</span>
type MultiAssetResponsiveDisplayAd struct {
	Id                         *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                        *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                 *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                  []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls            []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls               []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate        *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix             *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters        *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                    []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                  *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                       *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference           *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource  *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                     *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	MarketingImages            []*AssetLink               `xml:"marketingImages,omitempty" json:"marketingImages,omitempty" yaml:"marketingImages,omitempty"`
	SquareMarketingImages      []*AssetLink               `xml:"squareMarketingImages,omitempty" json:"squareMarketingImages,omitempty" yaml:"squareMarketingImages,omitempty"`
	LogoImages                 []*AssetLink               `xml:"logoImages,omitempty" json:"logoImages,omitempty" yaml:"logoImages,omitempty"`
	LandscapeLogoImages        []*AssetLink               `xml:"landscapeLogoImages,omitempty" json:"landscapeLogoImages,omitempty" yaml:"landscapeLogoImages,omitempty"`
	Headlines                  []*AssetLink               `xml:"headlines,omitempty" json:"headlines,omitempty" yaml:"headlines,omitempty"`
	LongHeadline               *AssetLink                 `xml:"longHeadline,omitempty" json:"longHeadline,omitempty" yaml:"longHeadline,omitempty"`
	Descriptions               []*AssetLink               `xml:"descriptions,omitempty" json:"descriptions,omitempty" yaml:"descriptions,omitempty"`
	YouTubeVideos              []*AssetLink               `xml:"youTubeVideos,omitempty" json:"youTubeVideos,omitempty" yaml:"youTubeVideos,omitempty"`
	BusinessName               *string                    `xml:"businessName,omitempty" json:"businessName,omitempty" yaml:"businessName,omitempty"`
	MainColor                  *string                    `xml:"mainColor,omitempty" json:"mainColor,omitempty" yaml:"mainColor,omitempty"`
	AccentColor                *string                    `xml:"accentColor,omitempty" json:"accentColor,omitempty" yaml:"accentColor,omitempty"`
	AllowFlexibleColor         *bool                      `xml:"allowFlexibleColor,omitempty" json:"allowFlexibleColor,omitempty" yaml:"allowFlexibleColor,omitempty"`
	CallToActionText           *string                    `xml:"callToActionText,omitempty" json:"callToActionText,omitempty" yaml:"callToActionText,omitempty"`
	DynamicSettingsPricePrefix *string                    `xml:"dynamicSettingsPricePrefix,omitempty" json:"dynamicSettingsPricePrefix,omitempty" yaml:"dynamicSettingsPricePrefix,omitempty"`
	DynamicSettingsPromoText   *string                    `xml:"dynamicSettingsPromoText,omitempty" json:"dynamicSettingsPromoText,omitempty" yaml:"dynamicSettingsPromoText,omitempty"`
	FormatSetting              *DisplayAdFormatSetting    `xml:"formatSetting,omitempty" json:"formatSetting,omitempty" yaml:"formatSetting,omitempty"`
	TypeAttrXSI                string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace              string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MultiAssetResponsiveDisplayAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MultiAssetResponsiveDisplayAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Error associated with creation of new entities.
type NewEntityCreationError struct {
	FieldPath         *string                       `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement           `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                       `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                       `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                       `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *NewEntityCreationErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *NewEntityCreationError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:NewEntityCreationError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors corresponding with violation of a NOT EMPTY check.
type NotEmptyError struct {
	FieldPath         *string              `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement  `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string              `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string              `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string              `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *NotEmptyErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *NotEmptyError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:NotEmptyError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors associated with violation of a NOT NULL check.
type NullError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *NullErrorReason    `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *NullError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:NullError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// This represents an operation that includes an operator and an
// operand             specified type.
type Operation interface{}

// Operation not permitted due to the invoked service's access
// policy.
type OperationAccessDenied struct {
	FieldPath         *string                      `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement          `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                      `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                      `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                      `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *OperationAccessDeniedReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *OperationAccessDenied) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:OperationAccessDenied"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors due to the use of unsupported operations.
type OperatorError struct {
	FieldPath         *string              `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement  `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string              `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string              `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string              `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *OperatorErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *OperatorError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:OperatorError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Specifies how the resulting information should be sorted.
type OrderBy struct {
	Field     *string    `xml:"field,omitempty" json:"field,omitempty" yaml:"field,omitempty"`
	SortOrder *SortOrder `xml:"sortOrder,omitempty" json:"sortOrder,omitempty" yaml:"sortOrder,omitempty"`
}

// Contains the results from a get call.
type Page interface{}

// Specifies the page of results to return in the response. A page
// is specified             by the result position to start at
// and the maximum number of results to             return.
type Paging struct {
	StartIndex    *int `xml:"startIndex,omitempty" json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
	NumberResults *int `xml:"numberResults,omitempty" json:"numberResults,omitempty" yaml:"numberResults,omitempty"`
}

// Error codes for pagination.
type PagingError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *PagingErrorReason  `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *PagingError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:PagingError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Error indicating that an entity will be disapproved unless changes
// are made to it before it is             saved. This error occurs
// when the entity will have a policy summary that includes a PROHIBITED
//             policy topic entry.
type PolicyFindingError struct {
	FieldPath         *string                   `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement       `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                   `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                   `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                   `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *PolicyFindingErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	PolicySummary     *PolicySummary            `xml:"policySummary,omitempty" json:"policySummary,omitempty" yaml:"policySummary,omitempty"`
	TypeAttrXSI       string                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *PolicyFindingError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:PolicyFindingError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Summary of policy information for a single entity.
type PolicySummary struct {
	PolicyTopicEntries []*PolicyTopicEntry `xml:"policyTopicEntries,omitempty" json:"policyTopicEntries,omitempty" yaml:"policyTopicEntries,omitempty"`
}

// Common policy summary information attached to a generic entity.
type PolicySummaryInfo interface{}

// A target which caused a policy finding.
type PolicyTopicConstraint interface{}

// Contains Ads Policy decisions.
type PolicyTopicEntry struct {
	PolicyTopicEntryType     *PolicyTopicEntryType    `xml:"policyTopicEntryType,omitempty" json:"policyTopicEntryType,omitempty" yaml:"policyTopicEntryType,omitempty"`
	PolicyTopicEvidences     []*PolicyTopicEvidence   `xml:"policyTopicEvidences,omitempty" json:"policyTopicEvidences,omitempty" yaml:"policyTopicEvidences,omitempty"`
	PolicyTopicConstraints   []*PolicyTopicConstraint `xml:"policyTopicConstraints,omitempty" json:"policyTopicConstraints,omitempty" yaml:"policyTopicConstraints,omitempty"`
	PolicyTopicId            *string                  `xml:"policyTopicId,omitempty" json:"policyTopicId,omitempty" yaml:"policyTopicId,omitempty"`
	PolicyTopicName          *string                  `xml:"policyTopicName,omitempty" json:"policyTopicName,omitempty" yaml:"policyTopicName,omitempty"`
	PolicyTopicHelpCenterUrl *string                  `xml:"policyTopicHelpCenterUrl,omitempty" json:"policyTopicHelpCenterUrl,omitempty" yaml:"policyTopicHelpCenterUrl,omitempty"`
}

// Evidence that caused this policy topic to be reported.
type PolicyTopicEvidence struct {
	PolicyTopicEvidenceType                        *PolicyTopicEvidenceType                         `xml:"policyTopicEvidenceType,omitempty" json:"policyTopicEvidenceType,omitempty" yaml:"policyTopicEvidenceType,omitempty"`
	EvidenceTextList                               []*string                                        `xml:"evidenceTextList,omitempty" json:"evidenceTextList,omitempty" yaml:"evidenceTextList,omitempty"`
	PolicyTopicEvidenceDestinationMismatchUrlTypes []*PolicyTopicEvidenceDestinationMismatchUrlType `xml:"policyTopicEvidenceDestinationMismatchUrlTypes,omitempty" json:"policyTopicEvidenceDestinationMismatchUrlTypes,omitempty" yaml:"policyTopicEvidenceDestinationMismatchUrlTypes,omitempty"`
}

// Represents violations of a single policy by some text in a field.
//                          Violations of a single policy by the
// same string in multiple places             within a field is
// reported in one instance of this class and only one
//     exemption needs to be filed.             Violations of a
// single policy by two different strings is reported
//    as two separate instances of this class.
//          e.g. If 'ACME' violates 'capitalization' and occurs
// twice in a text ad it             would be represented by one
// instance. If the ad also contains 'INC' which             also
// violates 'capitalization' it would be represented in a separate
//             instance.
type PolicyViolationError struct {
	FieldPath                 *string                     `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements         []*FieldPathElement         `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger                   *string                     `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString               *string                     `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType              *string                     `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Key                       *PolicyViolationKey         `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	ExternalPolicyName        *string                     `xml:"externalPolicyName,omitempty" json:"externalPolicyName,omitempty" yaml:"externalPolicyName,omitempty"`
	ExternalPolicyUrl         *string                     `xml:"externalPolicyUrl,omitempty" json:"externalPolicyUrl,omitempty" yaml:"externalPolicyUrl,omitempty"`
	ExternalPolicyDescription *string                     `xml:"externalPolicyDescription,omitempty" json:"externalPolicyDescription,omitempty" yaml:"externalPolicyDescription,omitempty"`
	IsExemptable              *bool                       `xml:"isExemptable,omitempty" json:"isExemptable,omitempty" yaml:"isExemptable,omitempty"`
	ViolatingParts            []*PolicyViolationErrorPart `xml:"violatingParts,omitempty" json:"violatingParts,omitempty" yaml:"violatingParts,omitempty"`
	TypeAttrXSI               string                      `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                      `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *PolicyViolationError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:PolicyViolationError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Points to a substring within an ad field or criterion.
type PolicyViolationErrorPart struct {
	Index  *int `xml:"index,omitempty" json:"index,omitempty" yaml:"index,omitempty"`
	Length *int `xml:"length,omitempty" json:"length,omitempty" yaml:"length,omitempty"`
}

// Key of the violation. The key is used for referring to a violation
// when             filing an exemption request.
type PolicyViolationKey struct {
	PolicyName    *string `xml:"policyName,omitempty" json:"policyName,omitempty" yaml:"policyName,omitempty"`
	ViolatingText *string `xml:"violatingText,omitempty" json:"violatingText,omitempty" yaml:"violatingText,omitempty"`
}

// Specifies how an entity (eg. adgroup, campaign, criterion, ad)
// should be filtered.
type Predicate struct {
	Field    *string            `xml:"field,omitempty" json:"field,omitempty" yaml:"field,omitempty"`
	Operator *PredicateOperator `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	Values   []*string          `xml:"values,omitempty" json:"values,omitempty" yaml:"values,omitempty"`
}

// Represents a product ad (known as a <a href=             "//support.google.com/adwords/answer/2456103">product
//             listing ad</a> in the AdWords user interface). A
// product ad displays             product data (managed using
// the Google Merchant Center) that is             pulled from
// the Google base product feed specified in the parent campaign's
//             {@linkplain ShoppingSetting shopping setting}.
//                         <p class="caution"><b>Caution:</b> Product
// ads do not use {@link #url url},             {@link #finalUrls
// finalUrls}, {@link #finalMobileUrls finalMobileUrls},
//       {@link #finalAppUrls finalAppUrls}, or {@link #displayUrl
// displayUrl};             setting these fields on a product ad
// will cause an error.             {@link #urlCustomParameters
// urlCustomParameters} and             {@link #trackingUrlTemplate
// trackingUrlTemplate} can be set, but it is not             recommended,
// as they will not be used; they should be set at the ad group
// or             campaign level instead.</p>             <span
// class="constraint AdxEnabled">This is disabled for AdX when
// it is contained within Operators: ADD, SET.</span>
type ProductAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents product image specific data.
type ProductImage struct {
	ProductImage        *Image               `xml:"productImage,omitempty" json:"productImage,omitempty" yaml:"productImage,omitempty"`
	Description         *string              `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	DisplayCallToAction *DisplayCallToAction `xml:"displayCallToAction,omitempty" json:"displayCallToAction,omitempty" yaml:"displayCallToAction,omitempty"`
}

// A QueryError represents possible errors for query parsing and
// execution.
type QueryError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *QueryErrorReason   `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	Message           *string             `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *QueryError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:QueryError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Encapsulates the errors thrown during developer quota checks.
type QuotaCheckError struct {
	FieldPath         *string                `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement    `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *QuotaCheckErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                 `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *QuotaCheckError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:QuotaCheckError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A list of all errors associated with the Range constraint.
type RangeError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *RangeErrorReason   `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *RangeError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:RangeError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Signals that a call failed because a measured rate exceeded.
type RateExceededError struct {
	FieldPath         *string                  `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement      `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                  `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                  `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                  `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *RateExceededErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	RateName          *string                  `xml:"rateName,omitempty" json:"rateName,omitempty" yaml:"rateName,omitempty"`
	RateScope         *string                  `xml:"rateScope,omitempty" json:"rateScope,omitempty" yaml:"rateScope,omitempty"`
	RetryAfterSeconds *int                     `xml:"retryAfterSeconds,omitempty" json:"retryAfterSeconds,omitempty" yaml:"retryAfterSeconds,omitempty"`
	TypeAttrXSI       string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *RateExceededError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:RateExceededError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors from attempting to write to read-only fields.
type ReadOnlyError struct {
	FieldPath         *string              `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement  `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string              `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string              `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string              `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *ReadOnlyErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ReadOnlyError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ReadOnlyError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Indicates that a field was rejected due to compatibility issues.
type RejectedError struct {
	FieldPath         *string              `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement  `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string              `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string              `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string              `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *RejectedErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *RejectedError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:RejectedError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Encapsulates the generic errors thrown when there's an error
// with user             request.
type RequestError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *RequestErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *RequestError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:RequestError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors due to missing required field.
type RequiredError struct {
	FieldPath         *string              `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement  `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string              `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string              `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string              `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *RequiredErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *RequiredError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:RequiredError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Constraint indicating that the policy topic was constrained
// by disapproval of the website for             reseller purposes.
type ResellerConstraint struct {
	ConstraintType            *PolicyTopicConstraintPolicyTopicConstraintType `xml:"constraintType,omitempty" json:"constraintType,omitempty" yaml:"constraintType,omitempty"`
	PolicyTopicConstraintType *string                                         `xml:"PolicyTopicConstraint.Type,omitempty" json:"PolicyTopicConstraint.Type,omitempty" yaml:"PolicyTopicConstraint.Type,omitempty"`
	TypeAttrXSI               string                                          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                                          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ResellerConstraint) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ResellerConstraint"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Enhanced display ad format.                          <p class="caution"><b>Caution:</b>
// Responsive display ads do not use {@link #url url},
//     {@link #displayUrl displayUrl}, {@link #finalAppUrls finalAppUrls},
// or             {@link #devicePreference devicePreference};
//            setting these fields on a responsive display ad will
// cause an error.             <span class="constraint AdxEnabled">This
// is enabled for AdX.</span>
type ResponsiveDisplayAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	MarketingImage            *Image                     `xml:"marketingImage,omitempty" json:"marketingImage,omitempty" yaml:"marketingImage,omitempty"`
	LogoImage                 *Image                     `xml:"logoImage,omitempty" json:"logoImage,omitempty" yaml:"logoImage,omitempty"`
	SquareMarketingImage      *Image                     `xml:"squareMarketingImage,omitempty" json:"squareMarketingImage,omitempty" yaml:"squareMarketingImage,omitempty"`
	ShortHeadline             *string                    `xml:"shortHeadline,omitempty" json:"shortHeadline,omitempty" yaml:"shortHeadline,omitempty"`
	LongHeadline              *string                    `xml:"longHeadline,omitempty" json:"longHeadline,omitempty" yaml:"longHeadline,omitempty"`
	Description               *string                    `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	BusinessName              *string                    `xml:"businessName,omitempty" json:"businessName,omitempty" yaml:"businessName,omitempty"`
	MainColor                 *string                    `xml:"mainColor,omitempty" json:"mainColor,omitempty" yaml:"mainColor,omitempty"`
	AccentColor               *string                    `xml:"accentColor,omitempty" json:"accentColor,omitempty" yaml:"accentColor,omitempty"`
	AllowFlexibleColor        *bool                      `xml:"allowFlexibleColor,omitempty" json:"allowFlexibleColor,omitempty" yaml:"allowFlexibleColor,omitempty"`
	CallToActionText          *string                    `xml:"callToActionText,omitempty" json:"callToActionText,omitempty" yaml:"callToActionText,omitempty"`
	DynamicDisplayAdSettings  *DynamicSettings           `xml:"dynamicDisplayAdSettings,omitempty" json:"dynamicDisplayAdSettings,omitempty" yaml:"dynamicDisplayAdSettings,omitempty"`
	FormatSetting             *DisplayAdFormatSetting    `xml:"formatSetting,omitempty" json:"formatSetting,omitempty" yaml:"formatSetting,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ResponsiveDisplayAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ResponsiveDisplayAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a responsive search ad.
// <p><strong>Important</strong></p>                          <p>
//             <ul>             <li>Responsive search ads are in
// beta and may not be available to all AdWords advertisers.</li>
//             <li>Per the <a href="https://support.google.com/adwordspolicy/answer/54818">AdWords
// Terms &amp;             Conditions</a> for features in beta,
// you may not disclose any non-public information.</li>
//       <li>Responsive search ads will learn and improve over
// time, so make sure to regularly monitor the             performance
// and status of your ads.</li>             <li>Responsive search
// ads are currently available in English, French, German, Spanish,
//             Portuguese, Italian, Swedish, Dutch, Russian, Japanese,
// Polish, Turkish, Danish, and             Norwegian, with other
// languages on the way.</li>             <li>Assets can be shown
// in any order, so make sure that they make sense individually
// or in             combination, and don?t violate our policies
// or local law.</li>             <li>Even after ads are assembled,
// they may not serve.</li>             <li>If you have text that
// should appear in every ad, then you must pin it to either Headline
//             position 1, Headline position 2, or Description
// position 1, and also make sure it is less             than 80
// characters long.</li>             </ul>             </p>
//                       <p>To increase the likelihood that your
// ad shows, be sure to provide at least 5 distinct
//  headlines and 2 distinct descriptions that don't repeat the
// same or similar phrases. Providing             redundant content
// will restrict the system's ability to assemble combinations.
//             <a href="https://support.google.com/adwords/answer/9023565">See
// example ads</a></p>                          <p>Help center
// documentation             <ul>             <li><a href="https://support.google.com/adwords/answer/7684791">About
// responsive search             ads</a></li>             <li><a
// href="https://support.google.com/adwords/answer/9023565">Create
// a responsive search ad             (see examples)</a></li>
//            </ul>             </p>                          <p
// class="caution"><b>Caution:</b> Responsive search ads do not
// use {@link #url url}, {@link             #displayUrl displayUrl},
// {@link #finalAppUrls finalAppUrls}, or {@link #devicePreference
//             devicePreference}; setting these fields on a responsive
// search ad will cause an error.</p>             <span class="constraint
// AdxEnabled">This is enabled for AdX.</span>
type ResponsiveSearchAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	Headlines                 []*AssetLink               `xml:"headlines,omitempty" json:"headlines,omitempty" yaml:"headlines,omitempty"`
	Descriptions              []*AssetLink               `xml:"descriptions,omitempty" json:"descriptions,omitempty" yaml:"descriptions,omitempty"`
	Path1                     *string                    `xml:"path1,omitempty" json:"path1,omitempty" yaml:"path1,omitempty"`
	Path2                     *string                    `xml:"path2,omitempty" json:"path2,omitempty" yaml:"path2,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ResponsiveSearchAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ResponsiveSearchAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Data associated with a rich media ad.             <span class="constraint
// AdxEnabled">This is disabled for AdX when it is contained within
// Operators: ADD, SET.</span>
type RichMediaAd interface{}

// A generic selector to specify the type of information to return.
type Selector struct {
	Fields     []*string    `xml:"fields,omitempty" json:"fields,omitempty" yaml:"fields,omitempty"`
	Predicates []*Predicate `xml:"predicates,omitempty" json:"predicates,omitempty" yaml:"predicates,omitempty"`
	DateRange  *DateRange   `xml:"dateRange,omitempty" json:"dateRange,omitempty" yaml:"dateRange,omitempty"`
	Ordering   []*OrderBy   `xml:"ordering,omitempty" json:"ordering,omitempty" yaml:"ordering,omitempty"`
	Paging     *Paging      `xml:"paging,omitempty" json:"paging,omitempty" yaml:"paging,omitempty"`
}

// Represents possible error codes for {@link Selector}.
type SelectorError struct {
	FieldPath         *string              `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement  `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string              `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string              `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string              `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *SelectorErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *SelectorError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SelectorError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a Showcase shopping ad.             <span class="constraint
// AdxEnabled">This is enabled for AdX.</span>
type ShowcaseAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	Name                      *string                    `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Headline                  *string                    `xml:"headline,omitempty" json:"headline,omitempty" yaml:"headline,omitempty"`
	Description               *string                    `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	CollapsedImage            *Image                     `xml:"collapsedImage,omitempty" json:"collapsedImage,omitempty" yaml:"collapsedImage,omitempty"`
	ExpandedImage             *Image                     `xml:"expandedImage,omitempty" json:"expandedImage,omitempty" yaml:"expandedImage,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ShowcaseAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ShowcaseAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Indicates that the number of entries in the request or response
// exceeds the system limit.
type SizeLimitError struct {
	FieldPath         *string               `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement   `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string               `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string               `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string               `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *SizeLimitErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *SizeLimitError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SizeLimitError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Defines the required and optional elements within the header
// of a SOAP request.
type SoapHeader struct {
	ClientCustomerId *string `xml:"clientCustomerId,omitempty" json:"clientCustomerId,omitempty" yaml:"clientCustomerId,omitempty"`
	DeveloperToken   *string `xml:"developerToken,omitempty" json:"developerToken,omitempty" yaml:"developerToken,omitempty"`
	UserAgent        *string `xml:"userAgent,omitempty" json:"userAgent,omitempty" yaml:"userAgent,omitempty"`
	ValidateOnly     *bool   `xml:"validateOnly,omitempty" json:"validateOnly,omitempty" yaml:"validateOnly,omitempty"`
	PartialFailure   *bool   `xml:"partialFailure,omitempty" json:"partialFailure,omitempty" yaml:"partialFailure,omitempty"`
}

// Defines the elements within the header of a SOAP response.
type SoapResponseHeader struct {
	RequestId    *string `xml:"requestId,omitempty" json:"requestId,omitempty" yaml:"requestId,omitempty"`
	ServiceName  *string `xml:"serviceName,omitempty" json:"serviceName,omitempty" yaml:"serviceName,omitempty"`
	MethodName   *string `xml:"methodName,omitempty" json:"methodName,omitempty" yaml:"methodName,omitempty"`
	Operations   *int64  `xml:"operations,omitempty" json:"operations,omitempty" yaml:"operations,omitempty"`
	ResponseTime *int64  `xml:"responseTime,omitempty" json:"responseTime,omitempty" yaml:"responseTime,omitempty"`
}

// Represents possible error codes when querying for stats.
type StatsQueryError struct {
	FieldPath         *string                `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement    `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *StatsQueryErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                 `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *StatsQueryError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:StatsQueryError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A list of error code for reporting invalid content of input
// strings.
type StringFormatError struct {
	FieldPath         *string                  `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement      `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                  `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                  `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                  `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *StringFormatErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *StringFormatError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:StringFormatError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors associated with the length of the given string being
//             out of bounds.
type StringLengthError struct {
	FieldPath         *string                  `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement      `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                  `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                  `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                  `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *StringLengthErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *StringLengthError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:StringLengthError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// This represents an entry in a map with a key of type String
//             and value of type String.
type String_StringMapEntry struct {
	Key   *string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// Represents the temporary id for an ad union id, which the user
// can specify.             The temporary id can be used to group
// ads together during ad creation.
type TempAdUnionId struct {
	Id            *int64  `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	AdUnionIdType *string `xml:"AdUnionId.Type,omitempty" json:"AdUnionId.Type,omitempty" yaml:"AdUnionId.Type,omitempty"`
	TypeAttrXSI   string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TempAdUnionId) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TempAdUnionId"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a <a href=             "//www.google.com/adwords/displaynetwork/plan-creative-campaigns/display-ad-builder.html"
//             >Display Ad Builder</a> template ad. A template
// ad is             composed of a template (specified by its ID)
// and the data that populates             the template's fields.
// For a list of available templates and their required
//      fields, see <a href="/adwords/api/docs/appendix/templateads">Template
// Ads</a>.             <span class="constraint AdxEnabled">This
// is disabled for AdX when it is contained within Operators: ADD,
// SET.</span>
type TemplateAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	TemplateId                *int64                     `xml:"templateId,omitempty" json:"templateId,omitempty" yaml:"templateId,omitempty"`
	AdUnionId                 *AdUnionId                 `xml:"adUnionId,omitempty" json:"adUnionId,omitempty" yaml:"adUnionId,omitempty"`
	TemplateElements          []*TemplateElement         `xml:"templateElements,omitempty" json:"templateElements,omitempty" yaml:"templateElements,omitempty"`
	AdAsImage                 *Image                     `xml:"adAsImage,omitempty" json:"adAsImage,omitempty" yaml:"adAsImage,omitempty"`
	Dimensions                *Dimensions                `xml:"dimensions,omitempty" json:"dimensions,omitempty" yaml:"dimensions,omitempty"`
	Name                      *string                    `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Duration                  *int                       `xml:"duration,omitempty" json:"duration,omitempty" yaml:"duration,omitempty"`
	OriginAdId                *int64                     `xml:"originAdId,omitempty" json:"originAdId,omitempty" yaml:"originAdId,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TemplateAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TemplateAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents an element in a template. Each template element is
// composed             of a list of fields (actual value data).
type TemplateElement struct {
	UniqueName *string                 `xml:"uniqueName,omitempty" json:"uniqueName,omitempty" yaml:"uniqueName,omitempty"`
	Fields     []*TemplateElementField `xml:"fields,omitempty" json:"fields,omitempty" yaml:"fields,omitempty"`
}

// Represents a field in a template element.
type TemplateElementField struct {
	Name       *string                   `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Type       *TemplateElementFieldType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	FieldText  *string                   `xml:"fieldText,omitempty" json:"fieldText,omitempty" yaml:"fieldText,omitempty"`
	FieldMedia *Media                    `xml:"fieldMedia,omitempty" json:"fieldMedia,omitempty" yaml:"fieldMedia,omitempty"`
}

// Represents a TextAd.             <span class="constraint AdxEnabled">This
// is disabled for AdX when it is contained within Operators: ADD,
// SET.</span>
type TextAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	Headline                  *string                    `xml:"headline,omitempty" json:"headline,omitempty" yaml:"headline,omitempty"`
	Description1              *string                    `xml:"description1,omitempty" json:"description1,omitempty" yaml:"description1,omitempty"`
	Description2              *string                    `xml:"description2,omitempty" json:"description2,omitempty" yaml:"description2,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TextAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TextAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a text asset.
type TextAsset struct {
	AssetId       *int64       `xml:"assetId,omitempty" json:"assetId,omitempty" yaml:"assetId,omitempty"`
	AssetName     *string      `xml:"assetName,omitempty" json:"assetName,omitempty" yaml:"assetName,omitempty"`
	AssetSubtype  *AssetType   `xml:"assetSubtype,omitempty" json:"assetSubtype,omitempty" yaml:"assetSubtype,omitempty"`
	AssetStatus   *AssetStatus `xml:"assetStatus,omitempty" json:"assetStatus,omitempty" yaml:"assetStatus,omitempty"`
	AssetType     *string      `xml:"Asset.Type,omitempty" json:"Asset.Type,omitempty" yaml:"Asset.Type,omitempty"`
	AssetText     *string      `xml:"assetText,omitempty" json:"assetText,omitempty" yaml:"assetText,omitempty"`
	TypeAttrXSI   string       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TextAsset) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TextAsset"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represent a display label entry.
type TextLabel struct {
	Id            *int64          `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Name          *string         `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Status        *LabelStatus    `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Attribute     *LabelAttribute `xml:"attribute,omitempty" json:"attribute,omitempty" yaml:"attribute,omitempty"`
	LabelType     *string         `xml:"Label.Type,omitempty" json:"Label.Type,omitempty" yaml:"Label.Type,omitempty"`
	TypeAttrXSI   string          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TextLabel) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TextLabel"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Data associated with rich media extension attributes.
//       <span class="constraint AdxEnabled">This is enabled for
// AdX.</span>
type ThirdPartyRedirectAd struct {
	Id                        *int64                                    `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                                   `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                                   `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                                 `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                                 `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                                 `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                                   `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                                   `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters                         `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                                `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                                     `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                                   `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                                    `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource                `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                                   `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	Name                      *string                                   `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Dimensions                *Dimensions                               `xml:"dimensions,omitempty" json:"dimensions,omitempty" yaml:"dimensions,omitempty"`
	Snippet                   *string                                   `xml:"snippet,omitempty" json:"snippet,omitempty" yaml:"snippet,omitempty"`
	ImpressionBeaconUrl       *string                                   `xml:"impressionBeaconUrl,omitempty" json:"impressionBeaconUrl,omitempty" yaml:"impressionBeaconUrl,omitempty"`
	AdDuration                *int                                      `xml:"adDuration,omitempty" json:"adDuration,omitempty" yaml:"adDuration,omitempty"`
	CertifiedVendorFormatId   *int64                                    `xml:"certifiedVendorFormatId,omitempty" json:"certifiedVendorFormatId,omitempty" yaml:"certifiedVendorFormatId,omitempty"`
	SourceUrl                 *string                                   `xml:"sourceUrl,omitempty" json:"sourceUrl,omitempty" yaml:"sourceUrl,omitempty"`
	RichMediaAdType           *RichMediaAdRichMediaAdType               `xml:"richMediaAdType,omitempty" json:"richMediaAdType,omitempty" yaml:"richMediaAdType,omitempty"`
	AdAttributes              []*RichMediaAdAdAttribute                 `xml:"adAttributes,omitempty" json:"adAttributes,omitempty" yaml:"adAttributes,omitempty"`
	IsCookieTargeted          *bool                                     `xml:"isCookieTargeted,omitempty" json:"isCookieTargeted,omitempty" yaml:"isCookieTargeted,omitempty"`
	IsUserInterestTargeted    *bool                                     `xml:"isUserInterestTargeted,omitempty" json:"isUserInterestTargeted,omitempty" yaml:"isUserInterestTargeted,omitempty"`
	IsTagged                  *bool                                     `xml:"isTagged,omitempty" json:"isTagged,omitempty" yaml:"isTagged,omitempty"`
	VideoTypes                []*VideoType                              `xml:"videoTypes,omitempty" json:"videoTypes,omitempty" yaml:"videoTypes,omitempty"`
	ExpandingDirections       []*ThirdPartyRedirectAdExpandingDirection `xml:"expandingDirections,omitempty" json:"expandingDirections,omitempty" yaml:"expandingDirections,omitempty"`
	TypeAttrXSI               string                                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ThirdPartyRedirectAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ThirdPartyRedirectAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a universal app ad                          <p class="caution"><b>Caution:</b>
// Universal app ads do not use #displayUrl displayUrl},
//       \{@link #finalAppUrls finalAppUrls}, or {@link #devicePreference
// devicePreference}; setting these             fields on a universal
// app ad will cause an error.             <span class="constraint
// AdxEnabled">This is enabled for AdX.</span>
type UniversalAppAd struct {
	Id                        *int64                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Url                       *string                    `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	DisplayUrl                *string                    `xml:"displayUrl,omitempty" json:"displayUrl,omitempty" yaml:"displayUrl,omitempty"`
	FinalUrls                 []*string                  `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls           []*string                  `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []*AppUrl                  `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       *string                    `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            *string                    `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters          `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData                 `xml:"urlData,omitempty" json:"urlData,omitempty" yaml:"urlData,omitempty"`
	Automated                 *bool                      `xml:"automated,omitempty" json:"automated,omitempty" yaml:"automated,omitempty"`
	Type                      *AdType                    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DevicePreference          *int64                     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty" json:"systemManagedEntitySource,omitempty" yaml:"systemManagedEntitySource,omitempty"`
	AdType                    *string                    `xml:"Ad.Type,omitempty" json:"Ad.Type,omitempty" yaml:"Ad.Type,omitempty"`
	Headlines                 []*AssetLink               `xml:"headlines,omitempty" json:"headlines,omitempty" yaml:"headlines,omitempty"`
	Descriptions              []*AssetLink               `xml:"descriptions,omitempty" json:"descriptions,omitempty" yaml:"descriptions,omitempty"`
	MandatoryAdText           *AssetLink                 `xml:"mandatoryAdText,omitempty" json:"mandatoryAdText,omitempty" yaml:"mandatoryAdText,omitempty"`
	Images                    []*AssetLink               `xml:"images,omitempty" json:"images,omitempty" yaml:"images,omitempty"`
	Videos                    []*AssetLink               `xml:"videos,omitempty" json:"videos,omitempty" yaml:"videos,omitempty"`
	Html5MediaBundles         []*AssetLink               `xml:"html5MediaBundles,omitempty" json:"html5MediaBundles,omitempty" yaml:"html5MediaBundles,omitempty"`
	TypeAttrXSI               string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *UniversalAppAd) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UniversalAppAd"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Holds a set of final urls that are scoped within a namespace.
type UrlData struct {
	UrlId               *string  `xml:"urlId,omitempty" json:"urlId,omitempty" yaml:"urlId,omitempty"`
	FinalUrls           *UrlList `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls     *UrlList `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	TrackingUrlTemplate *string  `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
}

// Url Validation errors.
type UrlError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *UrlErrorReason     `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *UrlError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UrlError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Wrapper POJO for a list of URLs.  The list can be cleared if
// a request contains             a UrlList with an empty urls
// list.
type UrlList struct {
	Urls []*string `xml:"urls,omitempty" json:"urls,omitempty" yaml:"urls,omitempty"`
}

// Encapsulates a Video media identified by a MediaId.
type Video struct {
	MediaId                              *int64                           `xml:"mediaId,omitempty" json:"mediaId,omitempty" yaml:"mediaId,omitempty"`
	Type                                 *MediaMediaType                  `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	ReferenceId                          *int64                           `xml:"referenceId,omitempty" json:"referenceId,omitempty" yaml:"referenceId,omitempty"`
	Dimensions                           []*Media_Size_DimensionsMapEntry `xml:"dimensions,omitempty" json:"dimensions,omitempty" yaml:"dimensions,omitempty"`
	Urls                                 []*Media_Size_StringMapEntry     `xml:"urls,omitempty" json:"urls,omitempty" yaml:"urls,omitempty"`
	MimeType                             *MediaLegacyMimeType             `xml:"mimeType,omitempty" json:"mimeType,omitempty" yaml:"mimeType,omitempty"`
	SourceUrl                            *string                          `xml:"sourceUrl,omitempty" json:"sourceUrl,omitempty" yaml:"sourceUrl,omitempty"`
	Name                                 *string                          `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	FileSize                             *int64                           `xml:"fileSize,omitempty" json:"fileSize,omitempty" yaml:"fileSize,omitempty"`
	CreationTime                         *string                          `xml:"creationTime,omitempty" json:"creationTime,omitempty" yaml:"creationTime,omitempty"`
	MediaType                            *string                          `xml:"Media.Type,omitempty" json:"Media.Type,omitempty" yaml:"Media.Type,omitempty"`
	DurationMillis                       *int64                           `xml:"durationMillis,omitempty" json:"durationMillis,omitempty" yaml:"durationMillis,omitempty"`
	StreamingUrl                         *string                          `xml:"streamingUrl,omitempty" json:"streamingUrl,omitempty" yaml:"streamingUrl,omitempty"`
	ReadyToPlayOnTheWeb                  *bool                            `xml:"readyToPlayOnTheWeb,omitempty" json:"readyToPlayOnTheWeb,omitempty" yaml:"readyToPlayOnTheWeb,omitempty"`
	IndustryStandardCommercialIdentifier *string                          `xml:"industryStandardCommercialIdentifier,omitempty" json:"industryStandardCommercialIdentifier,omitempty" yaml:"industryStandardCommercialIdentifier,omitempty"`
	AdvertisingId                        *string                          `xml:"advertisingId,omitempty" json:"advertisingId,omitempty" yaml:"advertisingId,omitempty"`
	YouTubeVideoIdString                 *string                          `xml:"youTubeVideoIdString,omitempty" json:"youTubeVideoIdString,omitempty" yaml:"youTubeVideoIdString,omitempty"`
	TypeAttrXSI                          string                           `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                        string                           `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Video) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Video"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a YouTube video asset.
type YouTubeVideoAsset struct {
	AssetId        *int64       `xml:"assetId,omitempty" json:"assetId,omitempty" yaml:"assetId,omitempty"`
	AssetName      *string      `xml:"assetName,omitempty" json:"assetName,omitempty" yaml:"assetName,omitempty"`
	AssetSubtype   *AssetType   `xml:"assetSubtype,omitempty" json:"assetSubtype,omitempty" yaml:"assetSubtype,omitempty"`
	AssetStatus    *AssetStatus `xml:"assetStatus,omitempty" json:"assetStatus,omitempty" yaml:"assetStatus,omitempty"`
	AssetType      *string      `xml:"Asset.Type,omitempty" json:"Asset.Type,omitempty" yaml:"Asset.Type,omitempty"`
	YouTubeVideoId *string      `xml:"youTubeVideoId,omitempty" json:"youTubeVideoId,omitempty" yaml:"youTubeVideoId,omitempty"`
	TypeAttrXSI    string       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace  string       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *YouTubeVideoAsset) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:YouTubeVideoAsset"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Get was auto-generated from WSDL.
type Get struct {
	ServiceSelector *Selector `xml:"serviceSelector,omitempty" json:"serviceSelector,omitempty" yaml:"serviceSelector,omitempty"`
}

// GetResponse was auto-generated from WSDL.
type GetResponse struct {
	Rval *AdGroupAdPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Mutate was auto-generated from WSDL.
type Mutate struct {
	Operations []*AdGroupAdOperation `xml:"operations,omitempty" json:"operations,omitempty" yaml:"operations,omitempty"`
}

// MutateLabel was auto-generated from WSDL.
type MutateLabel struct {
	Operations []*AdGroupAdLabelOperation `xml:"operations,omitempty" json:"operations,omitempty" yaml:"operations,omitempty"`
}

// MutateLabelResponse was auto-generated from WSDL.
type MutateLabelResponse struct {
	Rval *AdGroupAdLabelReturnValue `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// MutateResponse was auto-generated from WSDL.
type MutateResponse struct {
	Rval *AdGroupAdReturnValue `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Query was auto-generated from WSDL.
type Query struct {
	Query *string `xml:"query,omitempty" json:"query,omitempty" yaml:"query,omitempty"`
}

// QueryResponse was auto-generated from WSDL.
type QueryResponse struct {
	Rval *AdGroupAdPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Operation wrapper for Get.
// OperationGetRequest was auto-generated from WSDL.
type OperationGetRequest struct {
	Get *Get `xml:"get,omitempty" json:"get,omitempty" yaml:"get,omitempty"`
}

// Operation wrapper for Get.
// OperationGetResponse was auto-generated from WSDL.
type OperationGetResponse struct {
	GetResponse *GetResponse `xml:"getResponse,omitempty" json:"getResponse,omitempty" yaml:"getResponse,omitempty"`
}

// Operation wrapper for Mutate.
// OperationMutateRequest was auto-generated from WSDL.
type OperationMutateRequest struct {
	Mutate *Mutate `xml:"mutate,omitempty" json:"mutate,omitempty" yaml:"mutate,omitempty"`
}

// Operation wrapper for Mutate.
// OperationMutateResponse was auto-generated from WSDL.
type OperationMutateResponse struct {
	MutateResponse *MutateResponse `xml:"mutateResponse,omitempty" json:"mutateResponse,omitempty" yaml:"mutateResponse,omitempty"`
}

// Operation wrapper for MutateLabel.
// OperationMutateLabelRequest was auto-generated from WSDL.
type OperationMutateLabelRequest struct {
	MutateLabel *MutateLabel `xml:"mutateLabel,omitempty" json:"mutateLabel,omitempty" yaml:"mutateLabel,omitempty"`
}

// Operation wrapper for MutateLabel.
// OperationMutateLabelResponse was auto-generated from WSDL.
type OperationMutateLabelResponse struct {
	MutateLabelResponse *MutateLabelResponse `xml:"mutateLabelResponse,omitempty" json:"mutateLabelResponse,omitempty" yaml:"mutateLabelResponse,omitempty"`
}

// Operation wrapper for Query.
// OperationQueryRequest was auto-generated from WSDL.
type OperationQueryRequest struct {
	Query *Query `xml:"query,omitempty" json:"query,omitempty" yaml:"query,omitempty"`
}

// Operation wrapper for Query.
// OperationQueryResponse was auto-generated from WSDL.
type OperationQueryResponse struct {
	QueryResponse *QueryResponse `xml:"queryResponse,omitempty" json:"queryResponse,omitempty" yaml:"queryResponse,omitempty"`
}

// adGroupAdServiceInterface implements the AdGroupAdServiceInterface interface.
type adGroupAdServiceInterface struct {
	cli *soap.Client
}

// Returns a list of AdGroupAds.         AdGroupAds that had been
// removed are not returned by default.                  @param
// serviceSelector The selector specifying the {@link AdGroupAd}s
// to return.         @return The page containing the AdGroupAds
// that meet the criteria specified by the selector.         @throws
// ApiException when there is at least one error with the request.
func (p *adGroupAdServiceInterface) Get(Get *Get) (*GetResponse, error) {
	α := struct {
		OperationGetRequest `xml:"tns:get"`
	}{
		OperationGetRequest{
			Get,
		},
	}

	γ := struct {
		OperationGetResponse `xml:"getResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("Get", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetResponse, nil
}

// Applies the list of mutate operations (ie. add, set, remove):
//         <p>Add - Creates a new {@linkplain AdGroupAd ad group
// ad}. The         {@code adGroupId} must         reference an
// existing ad group. The child {@code Ad} must be sufficiently
//         specified by constructing a concrete ad type (such as
// {@code TextAd})         and setting its fields accordingly.</p>
//         <p>Set - Updates an ad group ad. Except for {@code status},
//         ad group ad fields are not mutable. Status updates are
//         straightforward - the status of the ad group ad is updated
// as         specified. If any other field has changed, it will
// be ignored. If         you want to change any of the fields
// other than status, you must         make a new ad and then remove
// the old one.</p>         <p>Remove - Removes the link between
// the specified AdGroup and         Ad.</p>         @param operations
// The operations to apply.         @return A list of AdGroupAds
// where each entry in the list is the result of         applying
// the operation in the input list with the same index. For an
//         add/set operation, the return AdGroupAd will be what
// is saved to the db.         In the case of the remove operation,
// the return AdGroupAd will simply be         an AdGroupAd containing
// an Ad with the id set to the Ad being removed from         the
// AdGroup.
func (p *adGroupAdServiceInterface) Mutate(Mutate *Mutate) (*MutateResponse, error) {
	α := struct {
		OperationMutateRequest `xml:"tns:mutate"`
	}{
		OperationMutateRequest{
			Mutate,
		},
	}

	γ := struct {
		OperationMutateResponse `xml:"mutateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("Mutate", α, &γ); err != nil {
		return nil, err
	}
	return γ.MutateResponse, nil
}

// Adds labels to the AdGroupAd or removes labels from the AdGroupAd.
//         <p>Add - Apply an existing label to an existing {@linkplain
// AdGroupAd ad group ad}. The         {@code adGroupId} and {@code
// adId} must reference an existing         {@linkplain AdGroupAd
// ad group ad}. The {@code labelId} must reference an existing
//         {@linkplain Label label}.         <p>Remove - Removes
// the link between the specified {@linkplain AdGroupAd ad group
// ad} and         {@linkplain Label label}.         @param operations
// The operations to apply.         @return A list of AdGroupAdLabel
// where each entry in the list is the result of         applying
// the operation in the input list with the same index. For an
//         add operation, the returned AdGroupAdLabel contains
// the AdGroupId, AdId and the LabelId.         In the case of
// a remove operation, the returned AdGroupAdLabel will only have
// AdGroupId and         AdId.         @throws ApiException when
// there are one or more errors with the request.
func (p *adGroupAdServiceInterface) MutateLabel(MutateLabel *MutateLabel) (*MutateLabelResponse, error) {
	α := struct {
		OperationMutateLabelRequest `xml:"tns:mutateLabel"`
	}{
		OperationMutateLabelRequest{
			MutateLabel,
		},
	}

	γ := struct {
		OperationMutateLabelResponse `xml:"mutateLabelResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("MutateLabel", α, &γ); err != nil {
		return nil, err
	}
	return γ.MutateLabelResponse, nil
}

// Returns a list of AdGroupAds based on the query.
//       @param query The SQL-like AWQL query string.         @return
// A list of AdGroupAds.         @throws ApiException if problems
// occur while parsing the query or fetching AdGroupAds.
func (p *adGroupAdServiceInterface) Query(Query string) (*QueryResponse, error) {
	α := struct {
		OperationQueryRequest `xml:"tns:query"`
	}{
		OperationQueryRequest{
			&Query,
		},
	}

	γ := struct {
		OperationQueryResponse `xml:"queryResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("Query", α, &γ); err != nil {
		return nil, err
	}
	return γ.QueryResponse, nil
}
