package adgroupcriterion

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "https://adwords.google.com/api/adwords/cm/v201809"

// NewAdGroupCriterionServiceInterface creates an initializes a AdGroupCriterionServiceInterface.
func NewAdGroupCriterionServiceInterface(cli *soap.Client) AdGroupCriterionServiceInterface {
	return &adGroupCriterionServiceInterface{cli}
}

// AdGroupCriterionServiceInterface was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type AdGroupCriterionServiceInterface interface {
	// Gets adgroup criteria.                  @param serviceSelector
	// filters the adgroup criteria to be returned.         @return
	// a page (subset) view of the criteria selected         @throws
	// ApiException when there is at least one error with the request
	Get(Get *Get) (*GetResponse, error)

	// Adds, removes or updates adgroup criteria.
	// @param operations operations to do         during checks on
	// keywords to be added.         @return added and updated adgroup
	// criteria (without optional parts)         @throws ApiException
	// when there is at least one error with the request
	Mutate(Mutate *Mutate) (*MutateResponse, error)

	// Adds labels to the AdGroupCriterion or removes labels from the
	// AdGroupCriterion         <p>Add - Apply an existing label to
	// an existing         {@linkplain AdGroupCriterion ad group criterion}.
	// The {@code adGroupId} and         {@code criterionId}
	//   must reference an existing {@linkplain AdGroupCriterion ad
	// group criterion}. The         {@code labelId} must reference
	// an existing {@linkplain Label label}.         <p>Remove - Removes
	// the link between the specified         {@linkplain AdGroupCriterion
	// ad group criterion} and {@linkplain Label label}.</p>
	//   @param operations the operations to apply         @return
	// a list of AdGroupCriterionLabel where each entry in the list
	// is the result of         applying the operation in the input
	// list with the same index. For an         add operation, the
	// returned AdGroupCriterionLabel contains the AdGroupId, CriterionId
	// and the         LabelId. In the case of a remove operation,
	// the returned AdGroupCriterionLabel will only have         AdGroupId
	// and CriterionId.         @throws ApiException when there are
	// one or more errors with the request
	MutateLabel(MutateLabel *MutateLabel) (*MutateLabelResponse, error)

	// Returns the list of AdGroupCriterion that match the query.
	//                 @param query The SQL-like AWQL query string
	//         @returns A list of AdGroupCriterion         @throws
	// ApiException when the query is invalid or there are errors processing
	// the request.
	Query(Query string) (*QueryResponse, error)
}

// AdGroupCriterionErrorReason was auto-generated from WSDL.
type AdGroupCriterionErrorReason string

// Validate validates AdGroupCriterionErrorReason.
func (v AdGroupCriterionErrorReason) Validate() bool {
	for _, vv := range []string{
		"AD_GROUP_CRITERION_LABEL_DOES_NOT_EXIST",
		"AD_GROUP_CRITERION_LABEL_ALREADY_EXISTS",
		"CANNOT_ADD_LABEL_TO_NEGATIVE_CRITERION",
		"TOO_MANY_OPERATIONS",
		"CANT_UPDATE_NEGATIVE",
		"CONCRETE_TYPE_REQUIRED",
		"BID_INCOMPATIBLE_WITH_ADGROUP",
		"CANNOT_TARGET_AND_EXCLUDE",
		"ILLEGAL_URL",
		"INVALID_KEYWORD_TEXT",
		"INVALID_DESTINATION_URL",
		"MISSING_DESTINATION_URL_TAG",
		"KEYWORD_LEVEL_BID_NOT_SUPPORTED_FOR_MANUALCPM",
		"INVALID_USER_STATUS",
		"CANNOT_ADD_CRITERIA_TYPE",
		"CANNOT_EXCLUDE_CRITERIA_TYPE",
		"INVALID_PRODUCT_PARTITION_HIERARCHY",
		"PRODUCT_PARTITION_UNIT_CANNOT_HAVE_CHILDREN",
		"PRODUCT_PARTITION_SUBDIVISION_REQUIRES_OTHERS_CASE",
		"PRODUCT_PARTITION_REQUIRES_SAME_DIMENSION_TYPE_AS_SIBLINGS",
		"PRODUCT_PARTITION_ALREADY_EXISTS",
		"PRODUCT_PARTITION_DOES_NOT_EXIST",
		"PRODUCT_PARTITION_CANNOT_BE_REMOVED",
		"INVALID_PRODUCT_PARTITION_TYPE",
		"PRODUCT_PARTITION_ADD_MAY_ONLY_USE_TEMP_ID",
		"CAMPAIGN_TYPE_NOT_COMPATIBLE_WITH_PARTIAL_FAILURE",
		"OPERATIONS_FOR_TOO_MANY_SHOPPING_ADGROUPS",
		"CANNOT_MODIFY_URL_FIELDS_WITH_DUPLICATE_ELEMENTS",
		"CANNOT_SET_WITHOUT_FINAL_URLS",
		"CANNOT_CLEAR_FINAL_URLS_IF_FINAL_MOBILE_URLS_EXIST",
		"CANNOT_CLEAR_FINAL_URLS_IF_FINAL_APP_URLS_EXIST",
		"CANNOT_CLEAR_FINAL_URLS_IF_TRACKING_URL_TEMPLATE_EXISTS",
		"CANNOT_CLEAR_FINAL_URLS_IF_URL_CUSTOM_PARAMETERS_EXIST",
		"CANNOT_SET_BOTH_DESTINATION_URL_AND_FINAL_URLS",
		"CANNOT_SET_BOTH_DESTINATION_URL_AND_TRACKING_URL_TEMPLATE",
		"FINAL_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE",
		"FINAL_MOBILE_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AdGroupCriterionLimitExceededCriteriaLimitType was auto-generated
// from WSDL.
type AdGroupCriterionLimitExceededCriteriaLimitType string

// Validate validates AdGroupCriterionLimitExceededCriteriaLimitType.
func (v AdGroupCriterionLimitExceededCriteriaLimitType) Validate() bool {
	for _, vv := range []string{
		"ADGROUP_KEYWORD",
		"ADGROUP_WEBSITE",
		"ADGROUP_CRITERION",
		"UNKNOWN",
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

// AgeRangeAgeRangeType was auto-generated from WSDL.
type AgeRangeAgeRangeType string

// Validate validates AgeRangeAgeRangeType.
func (v AgeRangeAgeRangeType) Validate() bool {
	for _, vv := range []string{
		"AGE_RANGE_18_24",
		"AGE_RANGE_25_34",
		"AGE_RANGE_35_44",
		"AGE_RANGE_45_54",
		"AGE_RANGE_55_64",
		"AGE_RANGE_65_UP",
		"AGE_RANGE_UNDETERMINED",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AppPaymentModelAppPaymentModelType was auto-generated from WSDL.
type AppPaymentModelAppPaymentModelType string

// Validate validates AppPaymentModelAppPaymentModelType.
func (v AppPaymentModelAppPaymentModelType) Validate() bool {
	for _, vv := range []string{
		"APP_PAYMENT_MODEL_PAID",
		"UNKNOWN",
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

// ApprovalStatus was auto-generated from WSDL.
type ApprovalStatus string

// Validate validates ApprovalStatus.
func (v ApprovalStatus) Validate() bool {
	for _, vv := range []string{
		"APPROVED",
		"PENDING_REVIEW",
		"UNDER_REVIEW",
		"DISAPPROVED",
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

// BidSource was auto-generated from WSDL.
type BidSource string

// Validate validates BidSource.
func (v BidSource) Validate() bool {
	for _, vv := range []string{
		"ADGROUP",
		"CRITERION",
		"ADGROUP_BIDDING_STRATEGY",
		"CAMPAIGN_BIDDING_STRATEGY",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// BiddingErrorsReason was auto-generated from WSDL.
type BiddingErrorsReason string

// Validate validates BiddingErrorsReason.
func (v BiddingErrorsReason) Validate() bool {
	for _, vv := range []string{
		"BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED",
		"BIDDING_STRATEGY_NOT_COMPATIBLE_WITH_ADGROUP_OVERRIDES",
		"BIDDING_STRATEGY_NOT_COMPATIBLE_WITH_ADGROUP_CRITERIA_OVERRIDES",
		"CAMPAIGN_BIDDING_STRATEGY_CANNOT_BE_OVERRIDDEN",
		"ADGROUP_BIDDING_STRATEGY_CANNOT_BE_OVERRIDDEN",
		"CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN",
		"CANNOT_ATTACH_BIDDING_STRATEGY_TO_ADGROUP",
		"CANNOT_ATTACH_BIDDING_STRATEGY_TO_ADGROUP_CRITERIA",
		"INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE",
		"BIDS_NOT_ALLLOWED",
		"DUPLICATE_BIDS",
		"INVALID_BIDDING_SCHEME",
		"INVALID_BIDDING_STRATEGY_TYPE",
		"MISSING_BIDDING_STRATEGY_TYPE",
		"NULL_BID",
		"INVALID_BID",
		"BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE",
		"CONVERSION_TRACKING_NOT_ENABLED",
		"NOT_ENOUGH_CONVERSIONS",
		"CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY",
		"CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_AD_GROUP_LEVEL_POP_BIDDING_STRATEGY",
		"CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY",
		"BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE",
		"PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER",
		"PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA",
		"BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS",
		"BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS",
		"BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION",
		"BID_TOO_SMALL",
		"BID_TOO_BIG",
		"BID_TOO_MANY_FRACTIONAL_DIGITS",
		"ENHANCED_CPC_ENABLED_NOT_SUPPORTED_ON_PORTFOLIO_TARGET_SPEND_STRATEGY",
		"BIDDING_STRATEGY_TYPE_NOT_ALLOWED_FOR_UNIVERSAL_APP_BIDDING_STRATEGY_GOAL_TYPE",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// BiddingStrategySource was auto-generated from WSDL.
type BiddingStrategySource string

// Validate validates BiddingStrategySource.
func (v BiddingStrategySource) Validate() bool {
	for _, vv := range []string{
		"CAMPAIGN",
		"ADGROUP",
		"CRITERION",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// BiddingStrategyType was auto-generated from WSDL.
type BiddingStrategyType string

// Validate validates BiddingStrategyType.
func (v BiddingStrategyType) Validate() bool {
	for _, vv := range []string{
		"MANUAL_CPC",
		"MANUAL_CPM",
		"PAGE_ONE_PROMOTED",
		"TARGET_SPEND",
		"TARGET_CPA",
		"TARGET_ROAS",
		"MAXIMIZE_CONVERSIONS",
		"MAXIMIZE_CONVERSION_VALUE",
		"TARGET_OUTRANK_SHARE",
		"NONE",
		"UNKNOWN",
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

// CriterionType was auto-generated from WSDL.
type CriterionType string

// Validate validates CriterionType.
func (v CriterionType) Validate() bool {
	for _, vv := range []string{
		"CONTENT_LABEL",
		"KEYWORD",
		"PLACEMENT",
		"VERTICAL",
		"USER_LIST",
		"USER_INTEREST",
		"MOBILE_APPLICATION",
		"MOBILE_APP_CATEGORY",
		"PRODUCT_PARTITION",
		"IP_BLOCK",
		"WEBPAGE",
		"LANGUAGE",
		"LOCATION",
		"AGE_RANGE",
		"CARRIER",
		"OPERATING_SYSTEM_VERSION",
		"MOBILE_DEVICE",
		"GENDER",
		"PARENT",
		"PROXIMITY",
		"PLATFORM",
		"PREFERRED_CONTENT",
		"AD_SCHEDULE",
		"LOCATION_GROUPS",
		"PRODUCT_SCOPE",
		"CUSTOM_AFFINITY",
		"CUSTOM_INTENT",
		"YOUTUBE_VIDEO",
		"YOUTUBE_CHANNEL",
		"APP_PAYMENT_MODEL",
		"INCOME_RANGE",
		"INTERACTION_TYPE",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// CriterionErrorReason was auto-generated from WSDL.
type CriterionErrorReason string

// Validate validates CriterionErrorReason.
func (v CriterionErrorReason) Validate() bool {
	for _, vv := range []string{
		"CONCRETE_TYPE_REQUIRED",
		"INVALID_EXCLUDED_CATEGORY",
		"INVALID_KEYWORD_TEXT",
		"KEYWORD_TEXT_TOO_LONG",
		"KEYWORD_HAS_TOO_MANY_WORDS",
		"KEYWORD_HAS_INVALID_CHARS",
		"INVALID_PLACEMENT_URL",
		"INVALID_USER_LIST",
		"INVALID_USER_INTEREST",
		"INVALID_FORMAT_FOR_PLACEMENT_URL",
		"PLACEMENT_URL_IS_TOO_LONG",
		"PLACEMENT_URL_HAS_ILLEGAL_CHAR",
		"PLACEMENT_URL_HAS_MULTIPLE_SITES_IN_LINE",
		"PLACEMENT_IS_NOT_AVAILABLE_FOR_TARGETING_OR_EXCLUSION",
		"INVALID_VERTICAL_PATH",
		"INVALID_YOUTUBE_CHANNEL_ID",
		"INVALID_YOUTUBE_VIDEO_ID",
		"YOUTUBE_VERTICAL_CHANNEL_DEPRECATED",
		"YOUTUBE_DEMOGRAPHIC_CHANNEL_DEPRECATED",
		"YOUTUBE_URL_UNSUPPORTED",
		"CANNOT_EXCLUDE_CRITERIA_TYPE",
		"CANNOT_ADD_CRITERIA_TYPE",
		"INVALID_PRODUCT_FILTER",
		"PRODUCT_FILTER_TOO_LONG",
		"CANNOT_EXCLUDE_SIMILAR_USER_LIST",
		"CANNOT_ADD_CLOSED_USER_LIST",
		"CANNOT_ADD_DISPLAY_ONLY_LISTS_TO_SEARCH_ONLY_CAMPAIGNS",
		"CANNOT_ADD_DISPLAY_ONLY_LISTS_TO_SEARCH_CAMPAIGNS",
		"CANNOT_ADD_DISPLAY_ONLY_LISTS_TO_SHOPPING_CAMPAIGNS",
		"CANNOT_ADD_USER_INTERESTS_TO_SEARCH_CAMPAIGNS",
		"CANNOT_SET_BIDS_ON_CRITERION_TYPE_IN_SEARCH_CAMPAIGNS",
		"CANNOT_ADD_URLS_TO_CRITERION_TYPE_FOR_CAMPAIGN_TYPE",
		"CANNOT_ADD_POSITIVE_PLACEMENTS_TO_SEARCH_CAMPAIGNS",
		"INVALID_IP_ADDRESS",
		"INVALID_IP_FORMAT",
		"INVALID_MOBILE_APP",
		"INVALID_MOBILE_APP_CATEGORY",
		"INVALID_CRITERION_ID",
		"CANNOT_TARGET_CRITERION",
		"CANNOT_TARGET_OBSOLETE_CRITERION",
		"CRITERION_ID_AND_TYPE_MISMATCH",
		"INVALID_PROXIMITY_RADIUS",
		"INVALID_PROXIMITY_RADIUS_UNITS",
		"INVALID_STREETADDRESS_LENGTH",
		"INVALID_CITYNAME_LENGTH",
		"INVALID_REGIONCODE_LENGTH",
		"INVALID_REGIONNAME_LENGTH",
		"INVALID_POSTALCODE_LENGTH",
		"INVALID_COUNTRY_CODE",
		"INVALID_LATITUDE",
		"INVALID_LONGITUDE",
		"PROXIMITY_GEOPOINT_AND_ADDRESS_BOTH_CANNOT_BE_NULL",
		"INVALID_PROXIMITY_ADDRESS",
		"INVALID_USER_DOMAIN_NAME",
		"INVALID_WEBPAGE_CONDITION",
		"INVALID_WEBPAGE_CONDITION_URL",
		"WEBPAGE_CONDITION_URL_CANNOT_BE_EMPTY",
		"WEBPAGE_CONDITION_URL_UNSUPPORTED_PROTOCOL",
		"WEBPAGE_CONDITION_URL_CANNOT_BE_IP_ADDRESS",
		"WEBPAGE_CONDITION_URL_DOMAIN_NOT_CONSISTENT_WITH_CAMPAIGN_SETTING",
		"WEBPAGE_CONDITION_URL_CANNOT_BE_PUBLIC_SUFFIX",
		"WEBPAGE_CONDITION_URL_INVALID_PUBLIC_SUFFIX",
		"WEBPAGE_CONDITION_URL_VALUE_TRACK_VALUE_NOT_SUPPORTED",
		"WEBPAGE_CRITERION_URL_EQUALS_CAN_HAVE_ONLY_ONE_CONDITION",
		"WEBPAGE_CRITERION_CANNOT_ADD_NON_SMART_TARGETING_TO_NON_DSA_AD_GROUP",
		"CRITERION_PARAMETER_TOO_LONG",
		"AD_SCHEDULE_TIME_INTERVALS_OVERLAP",
		"AD_SCHEDULE_INTERVAL_CANNOT_SPAN_MULTIPLE_DAYS",
		"AD_SCHEDULE_INVALID_TIME_INTERVAL",
		"AD_SCHEDULE_EXCEEDED_INTERVALS_PER_DAY_LIMIT",
		"AD_SCHEDULE_CRITERION_ID_MISMATCHING_FIELDS",
		"CANNOT_BID_MODIFY_CRITERION_TYPE",
		"CANNOT_BID_MODIFY_CRITERION_CAMPAIGN_OPTED_OUT",
		"CANNOT_BID_MODIFY_NEGATIVE_CRITERION",
		"BID_MODIFIER_ALREADY_EXISTS",
		"FEED_ID_NOT_ALLOWED",
		"ACCOUNT_INELIGIBLE_FOR_CRITERIA_TYPE",
		"CRITERIA_TYPE_INVALID_FOR_BIDDING_STRATEGY",
		"CANNOT_EXCLUDE_CRITERION",
		"CANNOT_REMOVE_CRITERION",
		"PRODUCT_SCOPE_TOO_LONG",
		"PRODUCT_SCOPE_TOO_MANY_DIMENSIONS",
		"PRODUCT_PARTITION_TOO_LONG",
		"PRODUCT_PARTITION_TOO_MANY_DIMENSIONS",
		"INVALID_PRODUCT_DIMENSION",
		"INVALID_PRODUCT_DIMENSION_TYPE",
		"INVALID_PRODUCT_BIDDING_CATEGORY",
		"MISSING_SHOPPING_SETTING",
		"INVALID_MATCHING_FUNCTION",
		"LOCATION_FILTER_NOT_ALLOWED",
		"LOCATION_FILTER_INVALID",
		"CANNOT_ATTACH_CRITERIA_AT_CAMPAIGN_AND_ADGROUP",
		"ADSENSE_FOR_MOBILE_APPS_PLACEMENT_DEPRECATED",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// CriterionUse was auto-generated from WSDL.
type CriterionUse string

// Validate validates CriterionUse.
func (v CriterionUse) Validate() bool {
	for _, vv := range []string{
		"BIDDABLE",
		"NEGATIVE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// CriterionUserListMembershipStatus was auto-generated from WSDL.
type CriterionUserListMembershipStatus string

// Validate validates CriterionUserListMembershipStatus.
func (v CriterionUserListMembershipStatus) Validate() bool {
	for _, vv := range []string{
		"OPEN",
		"CLOSED",
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

// GenderGenderType was auto-generated from WSDL.
type GenderGenderType string

// Validate validates GenderGenderType.
func (v GenderGenderType) Validate() bool {
	for _, vv := range []string{
		"GENDER_MALE",
		"GENDER_FEMALE",
		"GENDER_UNDETERMINED",
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

// IncomeRangeIncomeRangeType was auto-generated from WSDL.
type IncomeRangeIncomeRangeType string

// Validate validates IncomeRangeIncomeRangeType.
func (v IncomeRangeIncomeRangeType) Validate() bool {
	for _, vv := range []string{
		"INCOME_RANGE_UNDETERMINED",
		"INCOME_RANGE_0_50",
		"INCOME_RANGE_50_60",
		"INCOME_RANGE_60_70",
		"INCOME_RANGE_70_80",
		"INCOME_RANGE_80_90",
		"INCOME_RANGE_90_UP",
		"UNKNOWN",
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

// KeywordMatchType was auto-generated from WSDL.
type KeywordMatchType string

// Validate validates KeywordMatchType.
func (v KeywordMatchType) Validate() bool {
	for _, vv := range []string{
		"EXACT",
		"PHRASE",
		"BROAD",
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

// MultiplierErrorReason was auto-generated from WSDL.
type MultiplierErrorReason string

// Validate validates MultiplierErrorReason.
func (v MultiplierErrorReason) Validate() bool {
	for _, vv := range []string{
		"MULTIPLIER_TOO_HIGH",
		"MULTIPLIER_TOO_LOW",
		"TOO_MANY_FRACTIONAL_DIGITS",
		"MULTIPLIER_NOT_ALLOWED_FOR_BIDDING_STRATEGY",
		"MULTIPLIER_NOT_ALLOWED_WHEN_BASE_BID_IS_MISSING",
		"NO_MULTIPLIER_SPECIFIED",
		"MULTIPLIER_CAUSES_BID_TO_EXCEED_DAILY_BUDGET",
		"MULTIPLIER_CAUSES_BID_TO_EXCEED_MONTHLY_BUDGET",
		"MULTIPLIER_CAUSES_BID_TO_EXCEED_CUSTOM_BUDGET",
		"MULTIPLIER_CAUSES_BID_TO_EXCEED_MAX_ALLOWED_BID",
		"BID_LESS_THAN_MIN_ALLOWED_BID_WITH_MULTIPLIER",
		"MULTIPLIER_AND_BIDDING_STRATEGY_TYPE_MISMATCH",
		"MULTIPLIER_ERROR",
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

// PageOnePromotedBiddingSchemeStrategyGoal was auto-generated
// from WSDL.
type PageOnePromotedBiddingSchemeStrategyGoal string

// Validate validates PageOnePromotedBiddingSchemeStrategyGoal.
func (v PageOnePromotedBiddingSchemeStrategyGoal) Validate() bool {
	for _, vv := range []string{
		"PAGE_ONE",
		"PAGE_ONE_PROMOTED",
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

// ParentParentType was auto-generated from WSDL.
type ParentParentType string

// Validate validates ParentParentType.
func (v ParentParentType) Validate() bool {
	for _, vv := range []string{
		"PARENT_PARENT",
		"PARENT_NOT_A_PARENT",
		"PARENT_UNDETERMINED",
		"UNKNOWN",
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

// ProductCanonicalConditionCondition was auto-generated from WSDL.
type ProductCanonicalConditionCondition string

// Validate validates ProductCanonicalConditionCondition.
func (v ProductCanonicalConditionCondition) Validate() bool {
	for _, vv := range []string{
		"NEW",
		"USED",
		"REFURBISHED",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ProductDimensionType was auto-generated from WSDL.
type ProductDimensionType string

// Validate validates ProductDimensionType.
func (v ProductDimensionType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"BIDDING_CATEGORY_L1",
		"BIDDING_CATEGORY_L2",
		"BIDDING_CATEGORY_L3",
		"BIDDING_CATEGORY_L4",
		"BIDDING_CATEGORY_L5",
		"BRAND",
		"CANONICAL_CONDITION",
		"CUSTOM_ATTRIBUTE_0",
		"CUSTOM_ATTRIBUTE_1",
		"CUSTOM_ATTRIBUTE_2",
		"CUSTOM_ATTRIBUTE_3",
		"CUSTOM_ATTRIBUTE_4",
		"OFFER_ID",
		"PRODUCT_TYPE_L1",
		"PRODUCT_TYPE_L2",
		"PRODUCT_TYPE_L3",
		"PRODUCT_TYPE_L4",
		"PRODUCT_TYPE_L5",
		"CHANNEL",
		"CHANNEL_EXCLUSIVITY",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ProductPartitionType was auto-generated from WSDL.
type ProductPartitionType string

// Validate validates ProductPartitionType.
func (v ProductPartitionType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"SUBDIVISION",
		"UNIT",
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

// ShoppingProductChannel was auto-generated from WSDL.
type ShoppingProductChannel string

// Validate validates ShoppingProductChannel.
func (v ShoppingProductChannel) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"ONLINE",
		"LOCAL",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ShoppingProductChannelExclusivity was auto-generated from WSDL.
type ShoppingProductChannelExclusivity string

// Validate validates ShoppingProductChannelExclusivity.
func (v ShoppingProductChannelExclusivity) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"SINGLE_CHANNEL",
		"MULTI_CHANNEL",
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

// SystemServingStatus was auto-generated from WSDL.
type SystemServingStatus string

// Validate validates SystemServingStatus.
func (v SystemServingStatus) Validate() bool {
	for _, vv := range []string{
		"ELIGIBLE",
		"RARELY_SERVED",
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

// UserStatus was auto-generated from WSDL.
type UserStatus string

// Validate validates UserStatus.
func (v UserStatus) Validate() bool {
	for _, vv := range []string{
		"ENABLED",
		"REMOVED",
		"PAUSED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// WebpageConditionOperand was auto-generated from WSDL.
type WebpageConditionOperand string

// Validate validates WebpageConditionOperand.
func (v WebpageConditionOperand) Validate() bool {
	for _, vv := range []string{
		"URL",
		"CATEGORY",
		"PAGE_TITLE",
		"PAGE_CONTENT",
		"CUSTOM_LABEL",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// WebpageConditionOperator was auto-generated from WSDL.
type WebpageConditionOperator string

// Validate validates WebpageConditionOperator.
func (v WebpageConditionOperator) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"EQUALS",
		"CONTAINS",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// Represents a criterion in an ad group, used with AdGroupCriterionService.
type AdGroupCriterion struct {
	AdGroupId               *int64                   `xml:"adGroupId,omitempty" json:"adGroupId,omitempty" yaml:"adGroupId,omitempty"`
	CriterionUse            *CriterionUse            `xml:"criterionUse,omitempty" json:"criterionUse,omitempty" yaml:"criterionUse,omitempty"`
	Criterion               *Criterion               `xml:"criterion,omitempty" json:"criterion,omitempty" yaml:"criterion,omitempty"`
	Labels                  []*Label                 `xml:"labels,omitempty" json:"labels,omitempty" yaml:"labels,omitempty"`
	ForwardCompatibilityMap []*String_StringMapEntry `xml:"forwardCompatibilityMap,omitempty" json:"forwardCompatibilityMap,omitempty" yaml:"forwardCompatibilityMap,omitempty"`
	BaseCampaignId          *int64                   `xml:"baseCampaignId,omitempty" json:"baseCampaignId,omitempty" yaml:"baseCampaignId,omitempty"`
	BaseAdGroupId           *int64                   `xml:"baseAdGroupId,omitempty" json:"baseAdGroupId,omitempty" yaml:"baseAdGroupId,omitempty"`
	AdGroupCriterionType    *string                  `xml:"AdGroupCriterion.Type,omitempty" json:"AdGroupCriterion.Type,omitempty" yaml:"AdGroupCriterion.Type,omitempty"`
}

// Base error class for Ad Group Criterion Service.
type AdGroupCriterionError struct {
	FieldPath         *string                      `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement          `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                      `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                      `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                      `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *AdGroupCriterionErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupCriterionError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupCriterionError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Manages the labels associated with an AdGroupCriterion.
type AdGroupCriterionLabel struct {
	AdGroupId   *int64 `xml:"adGroupId,omitempty" json:"adGroupId,omitempty" yaml:"adGroupId,omitempty"`
	CriterionId *int64 `xml:"criterionId,omitempty" json:"criterionId,omitempty" yaml:"criterionId,omitempty"`
	LabelId     *int64 `xml:"labelId,omitempty" json:"labelId,omitempty" yaml:"labelId,omitempty"`
}

// Operations for adding/removing labels from AdGroupCriterion.
type AdGroupCriterionLabelOperation struct {
	Operator      *Operator              `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	OperationType *string                `xml:"Operation.Type,omitempty" json:"Operation.Type,omitempty" yaml:"Operation.Type,omitempty"`
	Operand       *AdGroupCriterionLabel `xml:"operand,omitempty" json:"operand,omitempty" yaml:"operand,omitempty"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupCriterionLabelOperation) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupCriterionLabelOperation"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A container for return values from the {@link AdGroupCriterionService#mutateLabel}
// call.
type AdGroupCriterionLabelReturnValue struct {
	ListReturnValueType  *string                  `xml:"ListReturnValue.Type,omitempty" json:"ListReturnValue.Type,omitempty" yaml:"ListReturnValue.Type,omitempty"`
	Value                []*AdGroupCriterionLabel `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	PartialFailureErrors []*ApiError              `xml:"partialFailureErrors,omitempty" json:"partialFailureErrors,omitempty" yaml:"partialFailureErrors,omitempty"`
	TypeAttrXSI          string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupCriterionLabelReturnValue) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupCriterionLabelReturnValue"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Signals that too many criteria were added to some ad group.
type AdGroupCriterionLimitExceeded struct {
	FieldPath         *string                                         `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement                             `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                                         `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                                         `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                                         `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *EntityCountLimitExceededReason                 `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	EnclosingId       *string                                         `xml:"enclosingId,omitempty" json:"enclosingId,omitempty" yaml:"enclosingId,omitempty"`
	Limit             *int                                            `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
	AccountLimitType  *string                                         `xml:"accountLimitType,omitempty" json:"accountLimitType,omitempty" yaml:"accountLimitType,omitempty"`
	ExistingCount     *int                                            `xml:"existingCount,omitempty" json:"existingCount,omitempty" yaml:"existingCount,omitempty"`
	LimitType         *AdGroupCriterionLimitExceededCriteriaLimitType `xml:"limitType,omitempty" json:"limitType,omitempty" yaml:"limitType,omitempty"`
	TypeAttrXSI       string                                          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                                          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupCriterionLimitExceeded) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupCriterionLimitExceeded"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Operation (add, remove and set) on adgroup criteria.
//                   <p>If you try to ADD a criterion that already
// exists, it will be treated as a SET operation             on
// the existing criterion.
type AdGroupCriterionOperation struct {
	Operator          *Operator           `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	OperationType     *string             `xml:"Operation.Type,omitempty" json:"Operation.Type,omitempty" yaml:"Operation.Type,omitempty"`
	Operand           *AdGroupCriterion   `xml:"operand,omitempty" json:"operand,omitempty" yaml:"operand,omitempty"`
	ExemptionRequests []*ExemptionRequest `xml:"exemptionRequests,omitempty" json:"exemptionRequests,omitempty" yaml:"exemptionRequests,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupCriterionOperation) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupCriterionOperation"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Contains a subset of adgroup criteria resulting from a
//        {@link AdGroupCriterionService#get} call.
type AdGroupCriterionPage struct {
	TotalNumEntries *int                `xml:"totalNumEntries,omitempty" json:"totalNumEntries,omitempty" yaml:"totalNumEntries,omitempty"`
	PageType        *string             `xml:"Page.Type,omitempty" json:"Page.Type,omitempty" yaml:"Page.Type,omitempty"`
	Entries         []*AdGroupCriterion `xml:"entries,omitempty" json:"entries,omitempty" yaml:"entries,omitempty"`
	TypeAttrXSI     string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupCriterionPage) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupCriterionPage"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A container for return values from the AdGroupCriterionService.
type AdGroupCriterionReturnValue struct {
	ListReturnValueType  *string             `xml:"ListReturnValue.Type,omitempty" json:"ListReturnValue.Type,omitempty" yaml:"ListReturnValue.Type,omitempty"`
	Value                []*AdGroupCriterion `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	PartialFailureErrors []*ApiError         `xml:"partialFailureErrors,omitempty" json:"partialFailureErrors,omitempty" yaml:"partialFailureErrors,omitempty"`
	TypeAttrXSI          string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupCriterionReturnValue) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupCriterionReturnValue"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
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

// Represents an Age Range criterion.             <p>A criterion
// of this type can only be created using an ID. A criterion of
// this type can be either targeted or excluded.             <span
// class="constraint AdxEnabled">This is disabled for AdX when
// it is contained within Operators: ADD, SET.</span>
type AgeRange struct {
	Id            *int64                `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType        `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string               `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	AgeRangeType  *AgeRangeAgeRangeType `xml:"ageRangeType,omitempty" json:"ageRangeType,omitempty" yaml:"ageRangeType,omitempty"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AgeRange) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AgeRange"
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

// Represents a criterion for targeting paid apps.
//              <p>Possible IDs: {@code 30} ({@code APP_PAYMENT_MODEL_PAID}).</p>
//             <p>A criterion of this type can only be created
// using an ID. A criterion of this type is only targetable.
//           <span class="constraint AdxEnabled">This is disabled
// for AdX when it is contained within Operators: ADD, SET.</span>
type AppPaymentModel struct {
	Id                  *int64                              `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type                *CriterionType                      `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType       *string                             `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	AppPaymentModelType *AppPaymentModelAppPaymentModelType `xml:"appPaymentModelType,omitempty" json:"appPaymentModelType,omitempty" yaml:"appPaymentModelType,omitempty"`
	TypeAttrXSI         string                              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string                              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AppPaymentModel) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AppPaymentModel"
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

// Wrapper object for a list of AppUrls. The list can be cleared
// if a request contains             an AppUrlList with an empty
// urls list.
type AppUrlList struct {
	AppUrls []*AppUrl `xml:"appUrls,omitempty" json:"appUrls,omitempty" yaml:"appUrls,omitempty"`
}

// Base class for exceptions.
type ApplicationException struct {
	Message                  *string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
	ApplicationExceptionType *string `xml:"ApplicationException.Type,omitempty" json:"ApplicationException.Type,omitempty" yaml:"ApplicationException.Type,omitempty"`
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

// Represents a bid of a certain amount.
type Bid struct {
	Amount *Money `xml:"amount,omitempty" json:"amount,omitempty" yaml:"amount,omitempty"`
}

// A biddable (positive) criterion in an adgroup.
type BiddableAdGroupCriterion struct {
	AdGroupId                    *int64                        `xml:"adGroupId,omitempty" json:"adGroupId,omitempty" yaml:"adGroupId,omitempty"`
	CriterionUse                 *CriterionUse                 `xml:"criterionUse,omitempty" json:"criterionUse,omitempty" yaml:"criterionUse,omitempty"`
	Criterion                    *Criterion                    `xml:"criterion,omitempty" json:"criterion,omitempty" yaml:"criterion,omitempty"`
	Labels                       []*Label                      `xml:"labels,omitempty" json:"labels,omitempty" yaml:"labels,omitempty"`
	ForwardCompatibilityMap      []*String_StringMapEntry      `xml:"forwardCompatibilityMap,omitempty" json:"forwardCompatibilityMap,omitempty" yaml:"forwardCompatibilityMap,omitempty"`
	BaseCampaignId               *int64                        `xml:"baseCampaignId,omitempty" json:"baseCampaignId,omitempty" yaml:"baseCampaignId,omitempty"`
	BaseAdGroupId                *int64                        `xml:"baseAdGroupId,omitempty" json:"baseAdGroupId,omitempty" yaml:"baseAdGroupId,omitempty"`
	AdGroupCriterionType         *string                       `xml:"AdGroupCriterion.Type,omitempty" json:"AdGroupCriterion.Type,omitempty" yaml:"AdGroupCriterion.Type,omitempty"`
	UserStatus                   *UserStatus                   `xml:"userStatus,omitempty" json:"userStatus,omitempty" yaml:"userStatus,omitempty"`
	SystemServingStatus          *SystemServingStatus          `xml:"systemServingStatus,omitempty" json:"systemServingStatus,omitempty" yaml:"systemServingStatus,omitempty"`
	ApprovalStatus               *ApprovalStatus               `xml:"approvalStatus,omitempty" json:"approvalStatus,omitempty" yaml:"approvalStatus,omitempty"`
	DisapprovalReasons           []*string                     `xml:"disapprovalReasons,omitempty" json:"disapprovalReasons,omitempty" yaml:"disapprovalReasons,omitempty"`
	FirstPageCpc                 *Bid                          `xml:"firstPageCpc,omitempty" json:"firstPageCpc,omitempty" yaml:"firstPageCpc,omitempty"`
	TopOfPageCpc                 *Bid                          `xml:"topOfPageCpc,omitempty" json:"topOfPageCpc,omitempty" yaml:"topOfPageCpc,omitempty"`
	FirstPositionCpc             *Bid                          `xml:"firstPositionCpc,omitempty" json:"firstPositionCpc,omitempty" yaml:"firstPositionCpc,omitempty"`
	QualityInfo                  *QualityInfo                  `xml:"qualityInfo,omitempty" json:"qualityInfo,omitempty" yaml:"qualityInfo,omitempty"`
	BiddingStrategyConfiguration *BiddingStrategyConfiguration `xml:"biddingStrategyConfiguration,omitempty" json:"biddingStrategyConfiguration,omitempty" yaml:"biddingStrategyConfiguration,omitempty"`
	BidModifier                  *float64                      `xml:"bidModifier,omitempty" json:"bidModifier,omitempty" yaml:"bidModifier,omitempty"`
	FinalUrls                    *UrlList                      `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls              *UrlList                      `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	FinalAppUrls                 *AppUrlList                   `xml:"finalAppUrls,omitempty" json:"finalAppUrls,omitempty" yaml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate          *string                       `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix               *string                       `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters          *CustomParameters             `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	TypeAttrXSI                  string                        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                string                        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *BiddableAdGroupCriterion) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:BiddableAdGroupCriterion"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents error codes for bidding strategy entities.
type BiddingErrors struct {
	FieldPath         *string              `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement  `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string              `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string              `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string              `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *BiddingErrorsReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *BiddingErrors) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:BiddingErrors"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Base class for all bidding schemes.             <span class="constraint
// AdxEnabled">This is disabled for AdX.</span>
type BiddingScheme interface{}

// Encapsulates the information about bids and bidding strategies.
//                          <p class="note"><b>Note:</b> Starting
// with v201705, bidding strategies can only be set on
//     campaigns. In earlier versions, bidding strategies can be
// set on campaigns, ad groups and ad             group criteria.
//                          <p>A bidding strategy can be set using
// one of the following:             <ul>             <li>{@linkplain
// BiddingStrategyConfiguration#biddingScheme bidding scheme}</li>
//             <li>{@linkplain BiddingStrategyConfiguration#biddingStrategyType
// bidding strategy type}</li>             <li>{@linkplain BiddingStrategyConfiguration#biddingStrategyId
// bidding strategy ID} for             flexible bid strategies.</li>
//             </ul>             <p>If the bidding strategy type
// is used, then schemes are created using default values.
//                      <p>Bids can be set only on ad groups and
// ad group criteria. They cannot be set on campaigns.
//     Multiple bids can be set at the same time. Only the bids
// that apply to the effective             bidding strategy will
// be used. Effective bidding strategy is considered to be the
// directly             attached strategy or inherited strategy
// from above level(s) when there is no directly attached
//        strategy.                          <p>For more information
// on flexible bidding, visit the             <a href="https://support.google.com/adwords/answer/2979071">Help
// Center</a>.
type BiddingStrategyConfiguration struct {
	BiddingStrategyId     *int64                 `xml:"biddingStrategyId,omitempty" json:"biddingStrategyId,omitempty" yaml:"biddingStrategyId,omitempty"`
	BiddingStrategyName   *string                `xml:"biddingStrategyName,omitempty" json:"biddingStrategyName,omitempty" yaml:"biddingStrategyName,omitempty"`
	BiddingStrategyType   *BiddingStrategyType   `xml:"biddingStrategyType,omitempty" json:"biddingStrategyType,omitempty" yaml:"biddingStrategyType,omitempty"`
	BiddingStrategySource *BiddingStrategySource `xml:"biddingStrategySource,omitempty" json:"biddingStrategySource,omitempty" yaml:"biddingStrategySource,omitempty"`
	BiddingScheme         *BiddingScheme         `xml:"biddingScheme,omitempty" json:"biddingScheme,omitempty" yaml:"biddingScheme,omitempty"`
	Bids                  []*Bids                `xml:"bids,omitempty" json:"bids,omitempty" yaml:"bids,omitempty"`
	TargetRoasOverride    *float64               `xml:"targetRoasOverride,omitempty" json:"targetRoasOverride,omitempty" yaml:"targetRoasOverride,omitempty"`
}

// Base class for all bids.
type Bids interface{}

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

// Comparable types for constructing ranges with.
type ComparableValue interface{}

// CPA Bids.
type CpaBid struct {
	BidsType      *string    `xml:"Bids.Type,omitempty" json:"Bids.Type,omitempty" yaml:"Bids.Type,omitempty"`
	Bid           *Money     `xml:"bid,omitempty" json:"bid,omitempty" yaml:"bid,omitempty"`
	BidSource     *BidSource `xml:"bidSource,omitempty" json:"bidSource,omitempty" yaml:"bidSource,omitempty"`
	TypeAttrXSI   string     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CpaBid) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CpaBid"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Manual click based bids.
type CpcBid struct {
	BidsType      *string    `xml:"Bids.Type,omitempty" json:"Bids.Type,omitempty" yaml:"Bids.Type,omitempty"`
	Bid           *Money     `xml:"bid,omitempty" json:"bid,omitempty" yaml:"bid,omitempty"`
	CpcBidSource  *BidSource `xml:"cpcBidSource,omitempty" json:"cpcBidSource,omitempty" yaml:"cpcBidSource,omitempty"`
	TypeAttrXSI   string     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CpcBid) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CpcBid"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Manual impression based bids.
type CpmBid struct {
	BidsType      *string    `xml:"Bids.Type,omitempty" json:"Bids.Type,omitempty" yaml:"Bids.Type,omitempty"`
	Bid           *Money     `xml:"bid,omitempty" json:"bid,omitempty" yaml:"bid,omitempty"`
	CpmBidSource  *BidSource `xml:"cpmBidSource,omitempty" json:"cpmBidSource,omitempty" yaml:"cpmBidSource,omitempty"`
	TypeAttrXSI   string     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CpmBid) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CpmBid"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a criterion (such as a keyword, placement, or vertical).
//             <span class="constraint AdxEnabled">This is disabled
// for AdX when it is contained within Operators: ADD, SET.</span>
type Criterion struct {
	Id            *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
}

// A Custom Affinity criterion.             <p> A criterion of
// this type is only targetable.             <span class="constraint
// AdxEnabled">This is disabled for AdX when it is contained within
// Operators: ADD, SET.</span>
type CriterionCustomAffinity struct {
	Id               *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type             *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType    *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	CustomAffinityId *int64         `xml:"customAffinityId,omitempty" json:"customAffinityId,omitempty" yaml:"customAffinityId,omitempty"`
	TypeAttrXSI      string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace    string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CriterionCustomAffinity) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CriterionCustomAffinity"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A Custom Intent criterion.             <p> A criterion of this
// type is only targetable.             <span class="constraint
// AdxEnabled">This is disabled for AdX when it is contained within
// Operators: ADD, SET.</span>
type CriterionCustomIntent struct {
	Id             *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type           *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType  *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	CustomIntentId *int64         `xml:"customIntentId,omitempty" json:"customIntentId,omitempty" yaml:"customIntentId,omitempty"`
	TypeAttrXSI    string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace  string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CriterionCustomIntent) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CriterionCustomIntent"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Error class used for reporting criteria related errors.
type CriterionError struct {
	FieldPath         *string               `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement   `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string               `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string               `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string               `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *CriterionErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CriterionError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CriterionError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Base type of criterion parameters.
type CriterionParameter interface{}

// Contains the policy violations for a single BiddableAdGroupCriterion.
type CriterionPolicyError struct {
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
func (t *CriterionPolicyError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CriterionPolicyError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// User Interest represents a particular interest-based vertical
// to be targeted.             <span class="constraint AdxEnabled">This
// is enabled for AdX.</span>
type CriterionUserInterest struct {
	Id                   *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type                 *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType        *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	UserInterestId       *int64         `xml:"userInterestId,omitempty" json:"userInterestId,omitempty" yaml:"userInterestId,omitempty"`
	UserInterestParentId *int64         `xml:"userInterestParentId,omitempty" json:"userInterestParentId,omitempty" yaml:"userInterestParentId,omitempty"`
	UserInterestName     *string        `xml:"userInterestName,omitempty" json:"userInterestName,omitempty" yaml:"userInterestName,omitempty"`
	TypeAttrXSI          string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CriterionUserInterest) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CriterionUserInterest"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// UserList - represents a user list that is defined by the advertiser
// to be targeted.             <span class="constraint AdxEnabled">This
// is enabled for AdX.</span>
type CriterionUserList struct {
	Id                         *int64                             `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type                       *CriterionType                     `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType              *string                            `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	UserListId                 *int64                             `xml:"userListId,omitempty" json:"userListId,omitempty" yaml:"userListId,omitempty"`
	UserListName               *string                            `xml:"userListName,omitempty" json:"userListName,omitempty" yaml:"userListName,omitempty"`
	UserListMembershipStatus   *CriterionUserListMembershipStatus `xml:"userListMembershipStatus,omitempty" json:"userListMembershipStatus,omitempty" yaml:"userListMembershipStatus,omitempty"`
	UserListEligibleForSearch  *bool                              `xml:"userListEligibleForSearch,omitempty" json:"userListEligibleForSearch,omitempty" yaml:"userListEligibleForSearch,omitempty"`
	UserListEligibleForDisplay *bool                              `xml:"userListEligibleForDisplay,omitempty" json:"userListEligibleForDisplay,omitempty" yaml:"userListEligibleForDisplay,omitempty"`
	TypeAttrXSI                string                             `xml:"xsi:type,attr,omitempty"`
	TypeNamespace              string                             `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CriterionUserList) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CriterionUserList"
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

// Number value type for constructing double valued ranges.
type DoubleValue struct {
	ComparableValueType *string  `xml:"ComparableValue.Type,omitempty" json:"ComparableValue.Type,omitempty" yaml:"ComparableValue.Type,omitempty"`
	Number              *float64 `xml:"number,omitempty" json:"number,omitempty" yaml:"number,omitempty"`
	TypeAttrXSI         string   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DoubleValue) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DoubleValue"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
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

// Represents a Gender criterion.             <p>A criterion of
// this type can only be created using an ID. A criterion of this
// type can be either targeted or excluded.             <span class="constraint
// AdxEnabled">This is disabled for AdX when it is contained within
// Operators: ADD, SET.</span>
type Gender struct {
	Id            *int64            `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string           `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	GenderType    *GenderGenderType `xml:"genderType,omitempty" json:"genderType,omitempty" yaml:"genderType,omitempty"`
	TypeAttrXSI   string            `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string            `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Gender) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Gender"
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

// Income range criterion allows to target and exclude predefined
// income percentile ranges.             <p>A criterion of this
// type can only be created using an ID. A criterion of this type
// can be either targeted or excluded.             <span class="constraint
// AdxEnabled">This is disabled for AdX when it is contained within
// Operators: ADD, SET.</span>
type IncomeRange struct {
	Id              *int64                      `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type            *CriterionType              `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType   *string                     `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	IncomeRangeType *IncomeRangeIncomeRangeType `xml:"incomeRangeType,omitempty" json:"incomeRangeType,omitempty" yaml:"incomeRangeType,omitempty"`
	TypeAttrXSI     string                      `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string                      `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *IncomeRange) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:IncomeRange"
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

// Represents a keyword.             <span class="constraint AdxEnabled">This
// is disabled for AdX when it is contained within Operators: ADD,
// SET.</span>
type Keyword struct {
	Id            *int64            `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string           `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	Text          *string           `xml:"text,omitempty" json:"text,omitempty" yaml:"text,omitempty"`
	MatchType     *KeywordMatchType `xml:"matchType,omitempty" json:"matchType,omitempty" yaml:"matchType,omitempty"`
	TypeAttrXSI   string            `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string            `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Keyword) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Keyword"
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

// Number value type for constructing long valued ranges.
type LongValue struct {
	ComparableValueType *string `xml:"ComparableValue.Type,omitempty" json:"ComparableValue.Type,omitempty" yaml:"ComparableValue.Type,omitempty"`
	Number              *int64  `xml:"number,omitempty" json:"number,omitempty" yaml:"number,omitempty"`
	TypeAttrXSI         string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *LongValue) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:LongValue"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Manual click based bidding where user pays per click.
//       <span class="constraint AdxEnabled">This is disabled for
// AdX.</span>
type ManualCpcBiddingScheme struct {
	BiddingSchemeType  *string `xml:"BiddingScheme.Type,omitempty" json:"BiddingScheme.Type,omitempty" yaml:"BiddingScheme.Type,omitempty"`
	EnhancedCpcEnabled *bool   `xml:"enhancedCpcEnabled,omitempty" json:"enhancedCpcEnabled,omitempty" yaml:"enhancedCpcEnabled,omitempty"`
	TypeAttrXSI        string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace      string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ManualCpcBiddingScheme) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ManualCpcBiddingScheme"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Manual impression based bidding where user pays per thousand
// impressions.             <span class="constraint AdxEnabled">This
// is enabled for AdX.</span>
type ManualCpmBiddingScheme struct {
	BiddingSchemeType  *string `xml:"BiddingScheme.Type,omitempty" json:"BiddingScheme.Type,omitempty" yaml:"BiddingScheme.Type,omitempty"`
	ViewableCpmEnabled *bool   `xml:"viewableCpmEnabled,omitempty" json:"viewableCpmEnabled,omitempty" yaml:"viewableCpmEnabled,omitempty"`
	TypeAttrXSI        string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace      string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ManualCpmBiddingScheme) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ManualCpmBiddingScheme"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Maximize Conversion Value bidding strategy is an automated bidding
// strategy which tries to             maximize conversion value
// given a daily budget.             <span class="constraint AdxEnabled">This
// is disabled for AdX.</span>
type MaximizeConversionValueBiddingScheme struct {
	BiddingSchemeType *string  `xml:"BiddingScheme.Type,omitempty" json:"BiddingScheme.Type,omitempty" yaml:"BiddingScheme.Type,omitempty"`
	TargetRoas        *float64 `xml:"targetRoas,omitempty" json:"targetRoas,omitempty" yaml:"targetRoas,omitempty"`
	TypeAttrXSI       string   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MaximizeConversionValueBiddingScheme) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MaximizeConversionValueBiddingScheme"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Maximize conversions is an automated bidding strategy that automatically
// sets bids to help             get the most conversions for your
// campaign while spending your budget.             <span class="constraint
// AdxEnabled">This is disabled for AdX.</span>
type MaximizeConversionsBiddingScheme struct {
	BiddingSchemeType *string `xml:"BiddingScheme.Type,omitempty" json:"BiddingScheme.Type,omitempty" yaml:"BiddingScheme.Type,omitempty"`
	TypeAttrXSI       string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MaximizeConversionsBiddingScheme) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MaximizeConversionsBiddingScheme"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents the mobile app category to be targeted.
//    <a href="/adwords/api/docs/appendix/mobileappcategories">View
// the complete list of             available mobile app categories</a>.
//             <span class="constraint AdxEnabled">This is enabled
// for AdX.</span>
type MobileAppCategory struct {
	Id                  *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type                *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType       *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	MobileAppCategoryId *int           `xml:"mobileAppCategoryId,omitempty" json:"mobileAppCategoryId,omitempty" yaml:"mobileAppCategoryId,omitempty"`
	DisplayName         *string        `xml:"displayName,omitempty" json:"displayName,omitempty" yaml:"displayName,omitempty"`
	TypeAttrXSI         string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MobileAppCategory) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MobileAppCategory"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents the mobile application to be targeted.
//   <span class="constraint AdxEnabled">This is enabled for AdX.</span>
type MobileApplication struct {
	Id            *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	AppId         *string        `xml:"appId,omitempty" json:"appId,omitempty" yaml:"appId,omitempty"`
	DisplayName   *string        `xml:"displayName,omitempty" json:"displayName,omitempty" yaml:"displayName,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MobileApplication) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MobileApplication"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a money amount.
type Money struct {
	ComparableValueType *string `xml:"ComparableValue.Type,omitempty" json:"ComparableValue.Type,omitempty" yaml:"ComparableValue.Type,omitempty"`
	MicroAmount         *int64  `xml:"microAmount,omitempty" json:"microAmount,omitempty" yaml:"microAmount,omitempty"`
	TypeAttrXSI         string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Money) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Money"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents errors in bid multipliers.
type MultiplierError struct {
	FieldPath         *string                `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement    `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *MultiplierErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                 `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MultiplierError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MultiplierError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A negative criterion in an adgroup.
type NegativeAdGroupCriterion struct {
	AdGroupId               *int64                   `xml:"adGroupId,omitempty" json:"adGroupId,omitempty" yaml:"adGroupId,omitempty"`
	CriterionUse            *CriterionUse            `xml:"criterionUse,omitempty" json:"criterionUse,omitempty" yaml:"criterionUse,omitempty"`
	Criterion               *Criterion               `xml:"criterion,omitempty" json:"criterion,omitempty" yaml:"criterion,omitempty"`
	Labels                  []*Label                 `xml:"labels,omitempty" json:"labels,omitempty" yaml:"labels,omitempty"`
	ForwardCompatibilityMap []*String_StringMapEntry `xml:"forwardCompatibilityMap,omitempty" json:"forwardCompatibilityMap,omitempty" yaml:"forwardCompatibilityMap,omitempty"`
	BaseCampaignId          *int64                   `xml:"baseCampaignId,omitempty" json:"baseCampaignId,omitempty" yaml:"baseCampaignId,omitempty"`
	BaseAdGroupId           *int64                   `xml:"baseAdGroupId,omitempty" json:"baseAdGroupId,omitempty" yaml:"baseAdGroupId,omitempty"`
	AdGroupCriterionType    *string                  `xml:"AdGroupCriterion.Type,omitempty" json:"AdGroupCriterion.Type,omitempty" yaml:"AdGroupCriterion.Type,omitempty"`
	TypeAttrXSI             string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace           string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *NegativeAdGroupCriterion) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:NegativeAdGroupCriterion"
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

// Number value types for constructing number valued ranges.
type NumberValue interface{}

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

// Page-One Promoted bidding scheme, which sets max cpc bids to
//             target impressions on page one or page one promoted
// slots on google.com.             <span class="constraint AdxEnabled">This
// is disabled for AdX.</span>
type PageOnePromotedBiddingScheme struct {
	BiddingSchemeType             *string                                   `xml:"BiddingScheme.Type,omitempty" json:"BiddingScheme.Type,omitempty" yaml:"BiddingScheme.Type,omitempty"`
	StrategyGoal                  *PageOnePromotedBiddingSchemeStrategyGoal `xml:"strategyGoal,omitempty" json:"strategyGoal,omitempty" yaml:"strategyGoal,omitempty"`
	BidCeiling                    *Money                                    `xml:"bidCeiling,omitempty" json:"bidCeiling,omitempty" yaml:"bidCeiling,omitempty"`
	BidModifier                   *float64                                  `xml:"bidModifier,omitempty" json:"bidModifier,omitempty" yaml:"bidModifier,omitempty"`
	BidChangesForRaisesOnly       *bool                                     `xml:"bidChangesForRaisesOnly,omitempty" json:"bidChangesForRaisesOnly,omitempty" yaml:"bidChangesForRaisesOnly,omitempty"`
	RaiseBidWhenBudgetConstrained *bool                                     `xml:"raiseBidWhenBudgetConstrained,omitempty" json:"raiseBidWhenBudgetConstrained,omitempty" yaml:"raiseBidWhenBudgetConstrained,omitempty"`
	RaiseBidWhenLowQualityScore   *bool                                     `xml:"raiseBidWhenLowQualityScore,omitempty" json:"raiseBidWhenLowQualityScore,omitempty" yaml:"raiseBidWhenLowQualityScore,omitempty"`
	TypeAttrXSI                   string                                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                 string                                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *PageOnePromotedBiddingScheme) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:PageOnePromotedBiddingScheme"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

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

// Parent criterion.             <p>A criterion of this type can
// only be created using an ID. A criterion of this type can be
// either targeted or excluded.             <span class="constraint
// AdxEnabled">This is disabled for AdX when it is contained within
// Operators: ADD, SET.</span>
type Parent struct {
	Id            *int64            `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string           `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	ParentType    *ParentParentType `xml:"parentType,omitempty" json:"parentType,omitempty" yaml:"parentType,omitempty"`
	TypeAttrXSI   string            `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string            `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Parent) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Parent"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A placement used for modifying bids for sites when targeting
// the content             network.             <span class="constraint
// AdxEnabled">This is enabled for AdX.</span>
type Placement struct {
	Id            *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	Url           *string        `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Placement) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Placement"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
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

// An {@code adwords grouping} string. Not supported by campaigns
// of             {@link AdvertisingChannelType#SHOPPING} so is
// only used in {@link ProductScope}s.
type ProductAdwordsGrouping struct {
	ProductDimensionType *string `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	Value                *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI          string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductAdwordsGrouping) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductAdwordsGrouping"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// An {@code adwords labels} string. Not supported by campaigns
// of             {@link AdvertisingChannelType#SHOPPING} so is
// only used in {@link ProductScope}s.
type ProductAdwordsLabels struct {
	ProductDimensionType *string `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	Value                *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI          string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductAdwordsLabels) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductAdwordsLabels"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// One element of a bidding category at a certain level. Top-level
// categories are at level 1,             their children at level
// 2, and so on. We currently support up to 5 levels. The user
// must specify             a dimension type that indicates the
// level of the category. All cases of the same subdivision
//          must have the same dimension type (category level).
type ProductBiddingCategory struct {
	ProductDimensionType *string               `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	Type                 *ProductDimensionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Value                *int64                `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI          string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductBiddingCategory) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductBiddingCategory"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A brand string.
type ProductBrand struct {
	ProductDimensionType *string `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	Value                *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI          string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductBrand) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductBrand"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A canonical condition. Only supported by campaigns of
//       {@link AdvertisingChannelType#SHOPPING}.
type ProductCanonicalCondition struct {
	ProductDimensionType *string                             `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	Condition            *ProductCanonicalConditionCondition `xml:"condition,omitempty" json:"condition,omitempty" yaml:"condition,omitempty"`
	TypeAttrXSI          string                              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductCanonicalCondition) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductCanonicalCondition"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// The product channel dimension, which specifies the locality
// of an offer. Only supported by             campaigns of {@link
// AdvertisingChannelType#SHOPPING}.
type ProductChannel struct {
	ProductDimensionType *string                 `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	Channel              *ShoppingProductChannel `xml:"channel,omitempty" json:"channel,omitempty" yaml:"channel,omitempty"`
	TypeAttrXSI          string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductChannel) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductChannel"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// The product channel exclusivity dimension, which limits the
// availability of an offer between only             local, only
// online, or both. Only supported by campaigns of
// {@link AdvertisingChannelType#SHOPPING}.
type ProductChannelExclusivity struct {
	ProductDimensionType *string                            `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	ChannelExclusivity   *ShoppingProductChannelExclusivity `xml:"channelExclusivity,omitempty" json:"channelExclusivity,omitempty" yaml:"channelExclusivity,omitempty"`
	TypeAttrXSI          string                             `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                             `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductChannelExclusivity) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductChannelExclusivity"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A custom attribute value. As a product can have multiple custom
// attributes, the user must specify             a dimension type
// that indicates the index of the attribute by which to subdivide.
// All cases of             the same subdivision must have the
// same dimension type (attribute index).
type ProductCustomAttribute struct {
	ProductDimensionType *string               `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	Type                 *ProductDimensionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Value                *string               `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI          string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductCustomAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductCustomAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Dimension by which to subdivide or filter products.
type ProductDimension interface{}

// A plain condition string. Not supported by campaigns of
//         {@link AdvertisingChannelType#SHOPPING} so is only used
// in {@link ProductScope}s.
type ProductLegacyCondition struct {
	ProductDimensionType *string `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	Value                *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI          string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductLegacyCondition) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductLegacyCondition"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// An offer ID as specified by the merchant.
type ProductOfferId struct {
	ProductDimensionType *string `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	Value                *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI          string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductOfferId) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductOfferId"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Product partition (product group) in a shopping campaign. Depending
// on its type, a product             partition subdivides products
// along some product dimension, specifies a bid for products,
// or             excludes products from bidding.
//             <p>Inner nodes of a product partition hierarchy
// are always subdivisions. Each child is linked to
//  the subdivision via the {@code parentCriterionId} and defines
// a {@code caseValue}. For all             children of the same
// subdivision, the {@code caseValue}s must be mutually different
// but             instances of the same class.
//           To create a subdivision and child node in the same
// API request, they should refer to each other             using
// temporary criterion IDs in the {@code parentCriterionId} of
// the child, and ID field of the             subdivision. Temporary
// IDs are specified by using any negative integer. Temporary IDs
// only exist             within the scope of a single API request.
// The API will assign real criterion IDs, and replace
//     the temporary values, and the API response will reflect
// this.             <span class="constraint AdxEnabled">This is
// disabled for AdX when it is contained within Operators: ADD,
// SET.</span>
type ProductPartition struct {
	Id                *int64                `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type              *CriterionType        `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType     *string               `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	PartitionType     *ProductPartitionType `xml:"partitionType,omitempty" json:"partitionType,omitempty" yaml:"partitionType,omitempty"`
	ParentCriterionId *int64                `xml:"parentCriterionId,omitempty" json:"parentCriterionId,omitempty" yaml:"parentCriterionId,omitempty"`
	CaseValue         *ProductDimension     `xml:"caseValue,omitempty" json:"caseValue,omitempty" yaml:"caseValue,omitempty"`
	TypeAttrXSI       string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductPartition) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductPartition"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// One element of a product type string at a certain level. Top-level
// product types are at level 1,             their children at
// level 2, and so on. We currently support up to 5 levels. The
// user must specify             a dimension type that indicates
// the level of the product type. All cases of the same
//      subdivision must have the same dimension type (product
// type level).
type ProductType struct {
	ProductDimensionType *string               `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	Type                 *ProductDimensionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Value                *string               `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI          string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductType) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductType"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A full product type string. Category of the product according
// to the merchant's own             classification. Example:
//                         <pre>{@code "Home & Garden > Kitchen
// & Dining > Kitchen Appliances > Refrigerators"}</pre>
//                    <p>Not supported by campaigns of {@link AdvertisingChannelType#SHOPPING}
// so is only used in             {@link ProductScope}s.
type ProductTypeFull struct {
	ProductDimensionType *string `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	Value                *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI          string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductTypeFull) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductTypeFull"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Container for criterion quality information.
type QualityInfo struct {
	QualityScore *int `xml:"qualityScore,omitempty" json:"qualityScore,omitempty" yaml:"qualityScore,omitempty"`
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

// <a href="https://support.google.com/adwords/answer/6268632">Target
// CPA</a> is an automated bid             strategy that sets bids
// to help get as many conversions as possible at the target cost
// per             acquisition (CPA) you set.
//         <p>A {@linkplain #targetCpa target CPA} must be set
// for the strategy, but can also be optionally             set
// for individual ad groups in the strategy. Ad group targets,
// if set, will override strategy             targets.
//                  <p>Note that campaigns must meet
//   <a href="https://support.google.com/adwords/answer/2471188">specific
// eligibility requirements</a>             before they can use
// the Target CPA bid strategy.             <span class="constraint
// AdxEnabled">This is disabled for AdX.</span>
type TargetCpaBiddingScheme struct {
	BiddingSchemeType *string `xml:"BiddingScheme.Type,omitempty" json:"BiddingScheme.Type,omitempty" yaml:"BiddingScheme.Type,omitempty"`
	TargetCpa         *Money  `xml:"targetCpa,omitempty" json:"targetCpa,omitempty" yaml:"targetCpa,omitempty"`
	MaxCpcBidCeiling  *Money  `xml:"maxCpcBidCeiling,omitempty" json:"maxCpcBidCeiling,omitempty" yaml:"maxCpcBidCeiling,omitempty"`
	MaxCpcBidFloor    *Money  `xml:"maxCpcBidFloor,omitempty" json:"maxCpcBidFloor,omitempty" yaml:"maxCpcBidFloor,omitempty"`
	TypeAttrXSI       string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TargetCpaBiddingScheme) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TargetCpaBiddingScheme"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Target Outrank Share bidding strategy is an automated bidding
// strategy which automatically sets             bids so that the
// customer's ads appear above a specified competitors' ads for
// a specified target             fraction of the advertiser's
// eligible impressions on Google.com.             <span class="constraint
// AdxEnabled">This is disabled for AdX.</span>
type TargetOutrankShareBiddingScheme struct {
	BiddingSchemeType           *string `xml:"BiddingScheme.Type,omitempty" json:"BiddingScheme.Type,omitempty" yaml:"BiddingScheme.Type,omitempty"`
	TargetOutrankShare          *int    `xml:"targetOutrankShare,omitempty" json:"targetOutrankShare,omitempty" yaml:"targetOutrankShare,omitempty"`
	CompetitorDomain            *string `xml:"competitorDomain,omitempty" json:"competitorDomain,omitempty" yaml:"competitorDomain,omitempty"`
	MaxCpcBidCeiling            *Money  `xml:"maxCpcBidCeiling,omitempty" json:"maxCpcBidCeiling,omitempty" yaml:"maxCpcBidCeiling,omitempty"`
	BidChangesForRaisesOnly     *bool   `xml:"bidChangesForRaisesOnly,omitempty" json:"bidChangesForRaisesOnly,omitempty" yaml:"bidChangesForRaisesOnly,omitempty"`
	RaiseBidWhenLowQualityScore *bool   `xml:"raiseBidWhenLowQualityScore,omitempty" json:"raiseBidWhenLowQualityScore,omitempty" yaml:"raiseBidWhenLowQualityScore,omitempty"`
	TypeAttrXSI                 string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace               string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TargetOutrankShareBiddingScheme) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TargetOutrankShareBiddingScheme"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Target Roas bidding strategy helps you maximize revenue while
// averaging a specific target             return on average spend
// (ROAS).                          <p>For example: If TargetRoas
// is 1.5, the strategy will create as much revenue as possible
// while             ensuring that every $1.00 of clicks provides
// $1.50 in conversion value.                          <p>Note
// that campaigns must meet <a             href="//support.google.com/adwords/answer/6268637">specific
//             eligibility requirements</a> before they can use
// the <code>TargetRoasBiddingScheme</code>             bidding
// strategy.             <span class="constraint AdxEnabled">This
// is disabled for AdX.</span>
type TargetRoasBiddingScheme struct {
	BiddingSchemeType *string  `xml:"BiddingScheme.Type,omitempty" json:"BiddingScheme.Type,omitempty" yaml:"BiddingScheme.Type,omitempty"`
	TargetRoas        *float64 `xml:"targetRoas,omitempty" json:"targetRoas,omitempty" yaml:"targetRoas,omitempty"`
	BidCeiling        *Money   `xml:"bidCeiling,omitempty" json:"bidCeiling,omitempty" yaml:"bidCeiling,omitempty"`
	BidFloor          *Money   `xml:"bidFloor,omitempty" json:"bidFloor,omitempty" yaml:"bidFloor,omitempty"`
	TypeAttrXSI       string   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TargetRoasBiddingScheme) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TargetRoasBiddingScheme"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// <a href="https://support.google.com/adwords/answer/6268626">Target
// Spend</a> is an automated             bid strategy that sets
// your bids to help get as many clicks as possible within your
// budget.             <span class="constraint AdxEnabled">This
// is disabled for AdX.</span>
type TargetSpendBiddingScheme struct {
	BiddingSchemeType *string `xml:"BiddingScheme.Type,omitempty" json:"BiddingScheme.Type,omitempty" yaml:"BiddingScheme.Type,omitempty"`
	BidCeiling        *Money  `xml:"bidCeiling,omitempty" json:"bidCeiling,omitempty" yaml:"bidCeiling,omitempty"`
	SpendTarget       *Money  `xml:"spendTarget,omitempty" json:"spendTarget,omitempty" yaml:"spendTarget,omitempty"`
	TypeAttrXSI       string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TargetSpendBiddingScheme) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TargetSpendBiddingScheme"
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

// An unknown product dimension type which will be returned in
// place of any ProductDimension not             supported by the
// clients current API version.
type UnknownProductDimension struct {
	ProductDimensionType *string `xml:"ProductDimension.Type,omitempty" json:"ProductDimension.Type,omitempty" yaml:"ProductDimension.Type,omitempty"`
	TypeAttrXSI          string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *UnknownProductDimension) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UnknownProductDimension"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
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

// Use verticals to target or exclude placements in the Google
// Display Network             based on the category into which
// the placement falls (for example, "Pets &amp;             Animals/Pets/Dogs").
//             <a href="/adwords/api/docs/appendix/verticals">View
// the complete list             of available vertical categories.</a>
//             <span class="constraint AdxEnabled">This is enabled
// for AdX.</span>
type Vertical struct {
	Id               *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type             *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType    *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	VerticalId       *int64         `xml:"verticalId,omitempty" json:"verticalId,omitempty" yaml:"verticalId,omitempty"`
	VerticalParentId *int64         `xml:"verticalParentId,omitempty" json:"verticalParentId,omitempty" yaml:"verticalParentId,omitempty"`
	Path             []*string      `xml:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty"`
	TypeAttrXSI      string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace    string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Vertical) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Vertical"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Criterion for targeting webpages of an advertiser's website.
// The             website domain name is specified at the campaign
// level.             <span class="constraint AdxEnabled">This
// is disabled for AdX when it is contained within Operators: ADD,
// SET.</span>
type Webpage struct {
	Id               *int64            `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type             *CriterionType    `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType    *string           `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	Parameter        *WebpageParameter `xml:"parameter,omitempty" json:"parameter,omitempty" yaml:"parameter,omitempty"`
	CriteriaCoverage *float64          `xml:"criteriaCoverage,omitempty" json:"criteriaCoverage,omitempty" yaml:"criteriaCoverage,omitempty"`
	CriteriaSamples  []*string         `xml:"criteriaSamples,omitempty" json:"criteriaSamples,omitempty" yaml:"criteriaSamples,omitempty"`
	TypeAttrXSI      string            `xml:"xsi:type,attr,omitempty"`
	TypeNamespace    string            `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Webpage) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Webpage"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Logical expression for targeting webpages of an advertiser's
// website.                          <p>A condition is defined
// as {@code operand OP argument}             where {@code operand}
// is one of the values enumerated in             {@link WebpageConditionOperand},
// and, based on this value,             {@code OP} is either of
// {@code EQUALS} or {@code CONTAINS}.</p>
type WebpageCondition struct {
	Operand  *WebpageConditionOperand  `xml:"operand,omitempty" json:"operand,omitempty" yaml:"operand,omitempty"`
	Argument *string                   `xml:"argument,omitempty" json:"argument,omitempty" yaml:"argument,omitempty"`
	Operator *WebpageConditionOperator `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
}

// Parameter of Webpage criterion, expressed as a list of conditions,
// or             logical expressions, for targeting webpages of
// an advertiser's website.
type WebpageParameter struct {
	CriterionParameterType *string             `xml:"CriterionParameter.Type,omitempty" json:"CriterionParameter.Type,omitempty" yaml:"CriterionParameter.Type,omitempty"`
	CriterionName          *string             `xml:"criterionName,omitempty" json:"criterionName,omitempty" yaml:"criterionName,omitempty"`
	Conditions             []*WebpageCondition `xml:"conditions,omitempty" json:"conditions,omitempty" yaml:"conditions,omitempty"`
	TypeAttrXSI            string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace          string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *WebpageParameter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:WebpageParameter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// YouTube channel criterion.             <p> A criterion of this
// type can be either targeted or excluded.             <span class="constraint
// AdxEnabled">This is disabled for AdX when it is contained within
// Operators: ADD, SET.</span>
type YouTubeChannel struct {
	Id            *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	ChannelId     *string        `xml:"channelId,omitempty" json:"channelId,omitempty" yaml:"channelId,omitempty"`
	ChannelName   *string        `xml:"channelName,omitempty" json:"channelName,omitempty" yaml:"channelName,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *YouTubeChannel) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:YouTubeChannel"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// YouTube video criterion.             <p> A criterion of this
// type can be either targeted or excluded.             <span class="constraint
// AdxEnabled">This is disabled for AdX when it is contained within
// Operators: ADD, SET.</span>
type YouTubeVideo struct {
	Id            *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	VideoId       *string        `xml:"videoId,omitempty" json:"videoId,omitempty" yaml:"videoId,omitempty"`
	VideoName     *string        `xml:"videoName,omitempty" json:"videoName,omitempty" yaml:"videoName,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *YouTubeVideo) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:YouTubeVideo"
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
	Rval *AdGroupCriterionPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Mutate was auto-generated from WSDL.
type Mutate struct {
	Operations []*AdGroupCriterionOperation `xml:"operations,omitempty" json:"operations,omitempty" yaml:"operations,omitempty"`
}

// MutateLabel was auto-generated from WSDL.
type MutateLabel struct {
	Operations []*AdGroupCriterionLabelOperation `xml:"operations,omitempty" json:"operations,omitempty" yaml:"operations,omitempty"`
}

// MutateLabelResponse was auto-generated from WSDL.
type MutateLabelResponse struct {
	Rval *AdGroupCriterionLabelReturnValue `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// MutateResponse was auto-generated from WSDL.
type MutateResponse struct {
	Rval *AdGroupCriterionReturnValue `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Query was auto-generated from WSDL.
type Query struct {
	Query *string `xml:"query,omitempty" json:"query,omitempty" yaml:"query,omitempty"`
}

// QueryResponse was auto-generated from WSDL.
type QueryResponse struct {
	Rval *AdGroupCriterionPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
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

// adGroupCriterionServiceInterface implements the AdGroupCriterionServiceInterface interface.
type adGroupCriterionServiceInterface struct {
	cli *soap.Client
}

// Gets adgroup criteria.                  @param serviceSelector
// filters the adgroup criteria to be returned.         @return
// a page (subset) view of the criteria selected         @throws
// ApiException when there is at least one error with the request
func (p *adGroupCriterionServiceInterface) Get(Get *Get) (*GetResponse, error) {
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

// Adds, removes or updates adgroup criteria.
// @param operations operations to do         during checks on
// keywords to be added.         @return added and updated adgroup
// criteria (without optional parts)         @throws ApiException
// when there is at least one error with the request
func (p *adGroupCriterionServiceInterface) Mutate(Mutate *Mutate) (*MutateResponse, error) {
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

// Adds labels to the AdGroupCriterion or removes labels from the
// AdGroupCriterion         <p>Add - Apply an existing label to
// an existing         {@linkplain AdGroupCriterion ad group criterion}.
// The {@code adGroupId} and         {@code criterionId}
//   must reference an existing {@linkplain AdGroupCriterion ad
// group criterion}. The         {@code labelId} must reference
// an existing {@linkplain Label label}.         <p>Remove - Removes
// the link between the specified         {@linkplain AdGroupCriterion
// ad group criterion} and {@linkplain Label label}.</p>
//   @param operations the operations to apply         @return
// a list of AdGroupCriterionLabel where each entry in the list
// is the result of         applying the operation in the input
// list with the same index. For an         add operation, the
// returned AdGroupCriterionLabel contains the AdGroupId, CriterionId
// and the         LabelId. In the case of a remove operation,
// the returned AdGroupCriterionLabel will only have         AdGroupId
// and CriterionId.         @throws ApiException when there are
// one or more errors with the request
func (p *adGroupCriterionServiceInterface) MutateLabel(MutateLabel *MutateLabel) (*MutateLabelResponse, error) {
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

// Returns the list of AdGroupCriterion that match the query.
//                 @param query The SQL-like AWQL query string
//         @returns A list of AdGroupCriterion         @throws
// ApiException when the query is invalid or there are errors processing
// the request.
func (p *adGroupCriterionServiceInterface) Query(Query string) (*QueryResponse, error) {
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
