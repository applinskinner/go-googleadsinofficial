package targetingidea

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "https://adwords.google.com/api/adwords/o/v201809"

// NewTargetingIdeaServiceInterface creates an initializes a TargetingIdeaServiceInterface.
func NewTargetingIdeaServiceInterface(cli *soap.Client) TargetingIdeaServiceInterface {
	return &targetingIdeaServiceInterface{cli}
}

// TargetingIdeaServiceInterface was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type TargetingIdeaServiceInterface interface {
	// Returns a page of ideas that match the query described by the
	// specified         {@link TargetingIdeaSelector}.
	//       <p>The selector must specify a {@code paging} value, with
	// {@code numberResults} set to 700 or         less.  Large result
	// sets must be composed through multiple calls to this method,
	// advancing the         paging {@code startIndex} value by {@code
	// numberResults} with each call.</p>                  @param selector
	// Query describing the types of results to return when
	//  finding matches (similar keyword ideas).         @return A
	// {@link TargetingIdeaPage} of results, that is a subset of the
	//         list of possible ideas meeting the criteria of the
	//        {@link TargetingIdeaSelector}.         @throws ApiException
	// If problems occurred while querying for ideas.
	Get(Get *Get) (*GetResponse, error)
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

// AttributeType was auto-generated from WSDL.
type AttributeType string

// Validate validates AttributeType.
func (v AttributeType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"CATEGORY_PRODUCTS_AND_SERVICES",
		"COMPETITION",
		"EXTRACTED_FROM_WEBPAGE",
		"IDEA_TYPE",
		"KEYWORD_TEXT",
		"SEARCH_VOLUME",
		"AVERAGE_CPC",
		"TARGETED_MONTHLY_SEARCHES",
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

// BudgetErrorReason was auto-generated from WSDL.
type BudgetErrorReason string

// Validate validates BudgetErrorReason.
func (v BudgetErrorReason) Validate() bool {
	for _, vv := range []string{
		"BUDGET_CANNOT_BE_EXPLICITLY_SHARED",
		"BUDGET_REMOVED",
		"BUDGET_ERROR",
		"BUDGET_IN_USE",
		"BUDGET_PERIOD_NOT_AVAILABLE",
		"CANNOT_EDIT_SHARED_BUDGET",
		"CANNOT_MODIFY_FIELD_OF_IMPLICITLY_SHARED_BUDGET",
		"CANNOT_UPDATE_BUDGET_TO_IMPLICITLY_SHARED",
		"CANNOT_UPDATE_BUDGET_TO_EXPLICITLY_SHARED_WITHOUT_NAME",
		"CANNOT_UPDATE_BUDGET_TO_EXPLICITLY_SHARED",
		"CANNOT_USE_IMPLICITLY_SHARED_BUDGET_WITH_MULTIPLE_CAMPAIGNS",
		"DUPLICATE_NAME",
		"MONEY_AMOUNT_IN_WRONG_CURRENCY",
		"MONEY_AMOUNT_LESS_THAN_CURRENCY_MINIMUM_CPC",
		"MONEY_AMOUNT_TOO_LARGE",
		"NEGATIVE_MONEY_AMOUNT",
		"NON_MULTIPLE_OF_MINIMUM_CURRENCY_UNIT",
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

// CompetitionSearchParameterLevel was auto-generated from WSDL.
type CompetitionSearchParameterLevel string

// Validate validates CompetitionSearchParameterLevel.
func (v CompetitionSearchParameterLevel) Validate() bool {
	for _, vv := range []string{
		"LOW",
		"MEDIUM",
		"HIGH",
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

// CurrencyCodeErrorReason was auto-generated from WSDL.
type CurrencyCodeErrorReason string

// Validate validates CurrencyCodeErrorReason.
func (v CurrencyCodeErrorReason) Validate() bool {
	for _, vv := range []string{
		"UNSUPPORTED_CURRENCY_CODE",
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

// IdeaType was auto-generated from WSDL.
type IdeaType string

// Validate validates IdeaType.
func (v IdeaType) Validate() bool {
	for _, vv := range []string{
		"KEYWORD",
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

// LocationTargetingStatus was auto-generated from WSDL.
type LocationTargetingStatus string

// Validate validates LocationTargetingStatus.
func (v LocationTargetingStatus) Validate() bool {
	for _, vv := range []string{
		"ACTIVE",
		"OBSOLETE",
		"PHASING_OUT",
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

// RegionCodeErrorReason was auto-generated from WSDL.
type RegionCodeErrorReason string

// Validate validates RegionCodeErrorReason.
func (v RegionCodeErrorReason) Validate() bool {
	for _, vv := range []string{
		"INVALID_REGION_CODE",
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

// RequestType was auto-generated from WSDL.
type RequestType string

// Validate validates RequestType.
func (v RequestType) Validate() bool {
	for _, vv := range []string{
		"IDEAS",
		"STATS",
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

// TargetingIdeaErrorReason was auto-generated from WSDL.
type TargetingIdeaErrorReason string

// Validate validates TargetingIdeaErrorReason.
func (v TargetingIdeaErrorReason) Validate() bool {
	for _, vv := range []string{
		"DUPLICATE_SEARCH_FILTER_TYPES_PRESENT",
		"INSUFFICIENT_SEARCH_PARAMETERS",
		"INVALID_ATTRIBUTE_TYPE",
		"INVALID_SEARCH_PARAMETERS",
		"INVALID_DOMAIN_SUFFIX",
		"MUTUALLY_EXCLUSIVE_SEARCH_PARAMETERS_IN_QUERY",
		"SERVICE_UNAVAILABLE",
		"INVALID_URL_IN_SEARCH_PARAMETER",
		"TOO_MANY_RESULTS_REQUESTED",
		"NO_PAGING_IN_SELECTOR",
		"INVALID_INCLUDED_EXCLUDED_KEYWORDS",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// TrafficEstimatorErrorReason was auto-generated from WSDL.
type TrafficEstimatorErrorReason string

// Validate validates TrafficEstimatorErrorReason.
func (v TrafficEstimatorErrorReason) Validate() bool {
	for _, vv := range []string{
		"NO_CAMPAIGN_FOR_AD_GROUP_ESTIMATE_REQUEST",
		"NO_AD_GROUP_FOR_KEYWORD_ESTIMATE_REQUEST",
		"NO_MAX_CPC_FOR_KEYWORD_ESTIMATE_REQUEST",
		"TOO_MANY_KEYWORD_ESTIMATE_REQUESTS",
		"TOO_MANY_CAMPAIGN_ESTIMATE_REQUESTS",
		"TOO_MANY_ADGROUP_ESTIMATE_REQUESTS",
		"TOO_MANY_TARGETS",
		"KEYWORD_TOO_LONG",
		"KEYWORD_CONTAINS_BROAD_MATCH_MODIFIERS",
		"INVALID_INPUT",
		"SERVICE_UNAVAILABLE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Base class for exceptions.
type ApplicationException struct {
	Message                  *string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
	ApplicationExceptionType *string `xml:"ApplicationException.Type,omitempty" json:"ApplicationException.Type,omitempty" yaml:"ApplicationException.Type,omitempty"`
}

// {@link Attribute}s encompass the core information about a particular
// {@link TargetingIdea}. Some             attributes are for {@code
// KEYWORD} {@link IdeaType}s, some are for {@code PLACEMENT}
//            {@link IdeaType}s, and some are for both. Ultimately,
// an {@link Attribute} instance simply wraps             an actual
// value of interest. For example, {@link KeywordAttribute} wraps
// the keyword itself,             while a {@link BooleanAttribute}
// simply wraps a boolean describing some information about the
//             keyword idea.
type Attribute interface{}

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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// {@link Attribute} type that contains a boolean value.
type BooleanAttribute struct {
	AttributeType *string `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         *bool   `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *BooleanAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:BooleanAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// A list of all the error codes being used by the common budget
// domain package.
type BudgetError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *BudgetErrorReason  `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *BudgetError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:BudgetError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// A {@link SearchParameter} for {@code KEYWORD} {@link IdeaType}s
// that             sets a keyword category that all search results
// should belong to.             Uses the newer "Products and Services"
// taxonomy.             <p>This element is supported by following
// {@link IdeaType}s: KEYWORD.             <p>This element is supported
// by following {@link RequestType}s: IDEAS.
type CategoryProductsAndServicesSearchParameter struct {
	SearchParameterType *string `xml:"SearchParameter.Type,omitempty" json:"SearchParameter.Type,omitempty" yaml:"SearchParameter.Type,omitempty"`
	CategoryId          *int    `xml:"categoryId,omitempty" json:"categoryId,omitempty" yaml:"categoryId,omitempty"`
	TypeAttrXSI         string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CategoryProductsAndServicesSearchParameter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CategoryProductsAndServicesSearchParameter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Comparable types for constructing ranges with.
type ComparableValue interface{}

// A {@link SearchParameter} for {@code KEYWORD} {@link IdeaType}s
// used to             filter the results by the amount of competition
// (eg: LOW, MEDIUM, HIGH).             <p>This element is supported
// by following {@link IdeaType}s: KEYWORD.             <p>This
// element is supported by following {@link RequestType}s: IDEAS,
// STATS.
type CompetitionSearchParameter struct {
	SearchParameterType *string                            `xml:"SearchParameter.Type,omitempty" json:"SearchParameter.Type,omitempty" yaml:"SearchParameter.Type,omitempty"`
	Levels              []*CompetitionSearchParameterLevel `xml:"levels,omitempty" json:"levels,omitempty" yaml:"levels,omitempty"`
	TypeAttrXSI         string                             `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string                             `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CompetitionSearchParameter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CompetitionSearchParameter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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

// {@link Attribute} type that contains a {@link Criterion} value.
type CriterionAttribute struct {
	AttributeType *string    `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         *Criterion `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CriterionAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CriterionAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Errors for currency codes.
type CurrencyCodeError struct {
	FieldPath         *string                  `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement      `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                  `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                  `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                  `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *CurrencyCodeErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CurrencyCodeError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CurrencyCodeError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// {@link Attribute} type that contains a double value.
type DoubleAttribute struct {
	AttributeType *string  `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         *float64 `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DoubleAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DoubleAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// A segment of a field path. Each dot in a field path defines
// a new segment.
type FieldPathElement struct {
	Field *string `xml:"field,omitempty" json:"field,omitempty" yaml:"field,omitempty"`
	Index *int    `xml:"index,omitempty" json:"index,omitempty" yaml:"index,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// A {@link SearchParameter} for {@code KEYWORD} {@link IdeaType}s
// that             specifies a collection of strings by which
// the results should be             constrained. This guarantees
// that each idea in the result will match             at least
// one of the {@code included} values.
//  For this {@link SearchParameter}, excluded items will always
// take             priority over included ones.
//            This can handle a maximum of 200 (included + excluded)
// elements.             <p>This element is supported by following
// {@link IdeaType}s: KEYWORD.             <p>This element is supported
// by following {@link RequestType}s: IDEAS, STATS.
type IdeaTextFilterSearchParameter struct {
	SearchParameterType *string   `xml:"SearchParameter.Type,omitempty" json:"SearchParameter.Type,omitempty" yaml:"SearchParameter.Type,omitempty"`
	Included            []*string `xml:"included,omitempty" json:"included,omitempty" yaml:"included,omitempty"`
	Excluded            []*string `xml:"excluded,omitempty" json:"excluded,omitempty" yaml:"excluded,omitempty"`
	TypeAttrXSI         string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *IdeaTextFilterSearchParameter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:IdeaTextFilterSearchParameter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// {@link Attribute} type that contains an {@link IdeaType} value.
// For example, if a             {@link TargetingIdea} represents
// a keyword idea, its {@link IdeaTypeAttribute} would contain
// a             {@code KEYWORD} {@link IdeaType}.
type IdeaTypeAttribute struct {
	AttributeType *string   `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         *IdeaType `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *IdeaTypeAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:IdeaTypeAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// {@link SearchParameter} that specifies whether adult content
// should be             returned.<p>
// Presence of this {@link SearchParameter} will allow adult keywords
//             to be included in the results.             <p>This
// element is supported by following {@link IdeaType}s: KEYWORD.
//             <p>This element is supported by following {@link
// RequestType}s: IDEAS, STATS.
type IncludeAdultContentSearchParameter struct {
	SearchParameterType *string `xml:"SearchParameter.Type,omitempty" json:"SearchParameter.Type,omitempty" yaml:"SearchParameter.Type,omitempty"`
	TypeAttrXSI         string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *IncludeAdultContentSearchParameter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:IncludeAdultContentSearchParameter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// {@link Attribute} type that contains an integer value.
type IntegerAttribute struct {
	AttributeType *string `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         *int    `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *IntegerAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:IntegerAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// {@link Attribute} type that contains a Set of integer values.
type IntegerSetAttribute struct {
	AttributeType *string `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         []*int  `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *IntegerSetAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:IntegerSetAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// {@link Attribute} type that contains a {@link Keyword} value.
type KeywordAttribute struct {
	AttributeType *string  `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         *Keyword `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *KeywordAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:KeywordAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Represents a Language criterion.             <p>A criterion
// of this type can only be created using an ID.             <span
// class="constraint AdxEnabled">This is enabled for AdX.</span>
type Language struct {
	Id            *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	Code          *string        `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	Name          *string        `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Language) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Language"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// A {@link SearchParameter} used to indicate multiple language
// being targeted.             This can be used, for example, to
// search for {@code KEYWORD}             {@link IdeaType}s that
// are best for Japanese language.                          <p>The
// service allows at most one language to be targeted for
//        {@code KEYWORD} requests.             <p>In the {@code
// KEYWORD} {@link IdeaType} {@code STATS} {@link RequestType}
//             requests, those keywords that are from different
// language than specified in             {@code LanguageSearchParameter}
// or have unknown language will be filtered             out in
// the response. To avoid filtering, do not include
//  {@code LanguageSearchParameter} in the request.
//  <p>This element is supported by following {@link IdeaType}s:
// KEYWORD.             <p>This element is supported by following
// {@link RequestType}s: IDEAS, STATS.
type LanguageSearchParameter struct {
	SearchParameterType *string   `xml:"SearchParameter.Type,omitempty" json:"SearchParameter.Type,omitempty" yaml:"SearchParameter.Type,omitempty"`
	Languages           []*string `xml:"languages,omitempty" json:"languages,omitempty" yaml:"languages,omitempty"`
	TypeAttrXSI         string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *LanguageSearchParameter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:LanguageSearchParameter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Represents Location criterion.             <p>A criterion of
// this type can only be created using an ID.             <span
// class="constraint AdxEnabled">This is enabled for AdX.</span>
type Location struct {
	Id              *int64                   `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type            *CriterionType           `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType   *string                  `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	LocationName    *string                  `xml:"locationName,omitempty" json:"locationName,omitempty" yaml:"locationName,omitempty"`
	DisplayType     *string                  `xml:"displayType,omitempty" json:"displayType,omitempty" yaml:"displayType,omitempty"`
	TargetingStatus *LocationTargetingStatus `xml:"targetingStatus,omitempty" json:"targetingStatus,omitempty" yaml:"targetingStatus,omitempty"`
	ParentLocations []*Location              `xml:"parentLocations,omitempty" json:"parentLocations,omitempty" yaml:"parentLocations,omitempty"`
	TypeAttrXSI     string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Location) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Location"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// A {@link SearchParameter} used to indicate the locations being
// targeted.             This can be used, for example, to search
// for {@code KEYWORD}             {@link IdeaType}s that are best
// for Japan and Los Angeles.                          <p>This
// parameter replaces the {@code CountryTargetSearchParameter}.</p>
//                          <p>See the             <a href="https://developers.google.com/adwords/api/docs/appendix/geotargeting">Geographical
//             Targeting</a> page for the complete list of supported
// geo target types for this service.</p>
//     <p>The service allows up to 10 locations to be targeted
// for KEYWORD requests and 50 locations             for PLACEMENT
// requests.</p>             <p>This element is supported by following
// {@link IdeaType}s: KEYWORD.             <p>This element is supported
// by following {@link RequestType}s: IDEAS, STATS.
type LocationSearchParameter struct {
	SearchParameterType *string     `xml:"SearchParameter.Type,omitempty" json:"SearchParameter.Type,omitempty" yaml:"SearchParameter.Type,omitempty"`
	Locations           []*Location `xml:"locations,omitempty" json:"locations,omitempty" yaml:"locations,omitempty"`
	TypeAttrXSI         string      `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string      `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *LocationSearchParameter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:LocationSearchParameter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// {@link Attribute} type that contains a long value.
type LongAttribute struct {
	AttributeType *string `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         *int64  `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *LongAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:LongAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Object representing integer comparison operations. This is usually
// used within             a particular {@link SearchParameter}
// to specify the valid values requested for the specific
//        {@link Attribute}.
type LongComparisonOperation struct {
	Minimum *int64 `xml:"minimum,omitempty" json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum *int64 `xml:"maximum,omitempty" json:"maximum,omitempty" yaml:"maximum,omitempty"`
}

// {@link Attribute} type that contains a {@link Range} of {@link
// LongValue}             values.
type LongRangeAttribute struct {
	AttributeType *string `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         *Range  `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *LongRangeAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:LongRangeAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// {@link Attribute} type that contains a {@link Money} value.
type MoneyAttribute struct {
	AttributeType *string `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         *Money  `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MoneyAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MoneyAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// A keyword response value representing search volume for a single
// month. An             instance with a {@code null} count is
// valid, indicating that the information             is unavailable.
type MonthlySearchVolume struct {
	Year  *int   `xml:"year,omitempty" json:"year,omitempty" yaml:"year,omitempty"`
	Month *int   `xml:"month,omitempty" json:"month,omitempty" yaml:"month,omitempty"`
	Count *int64 `xml:"count,omitempty" json:"count,omitempty" yaml:"count,omitempty"`
}

// {@link Attribute} type that contains a list of {@link MonthlySearchVolume}
//             values. The list contains the past 12 {@link MonthlySearchVolume}s
//             (excluding the current month). The first item is
// the data for the most             recent month and the last
// item is the data for the oldest month.
type MonthlySearchVolumeAttribute struct {
	AttributeType *string                `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         []*MonthlySearchVolume `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                 `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MonthlySearchVolumeAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MonthlySearchVolumeAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// <p>A {@link SearchParameter} for setting the search network.
// Currently we             support two network setting options:
//             <ul>             <li>Google search network</li>
//             <li>Google search network and AFS</li>
//    </ul>             <p>This element is supported by following
// {@link IdeaType}s: KEYWORD.             <p>This element is supported
// by following {@link RequestType}s: IDEAS, STATS.
type NetworkSearchParameter struct {
	SearchParameterType *string         `xml:"SearchParameter.Type,omitempty" json:"SearchParameter.Type,omitempty" yaml:"SearchParameter.Type,omitempty"`
	NetworkSetting      *NetworkSetting `xml:"networkSetting,omitempty" json:"networkSetting,omitempty" yaml:"networkSetting,omitempty"`
	TypeAttrXSI         string          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *NetworkSearchParameter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:NetworkSearchParameter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Network settings for a Campaign.
type NetworkSetting struct {
	TargetGoogleSearch         *bool `xml:"targetGoogleSearch,omitempty" json:"targetGoogleSearch,omitempty" yaml:"targetGoogleSearch,omitempty"`
	TargetSearchNetwork        *bool `xml:"targetSearchNetwork,omitempty" json:"targetSearchNetwork,omitempty" yaml:"targetSearchNetwork,omitempty"`
	TargetContentNetwork       *bool `xml:"targetContentNetwork,omitempty" json:"targetContentNetwork,omitempty" yaml:"targetContentNetwork,omitempty"`
	TargetPartnerSearchNetwork *bool `xml:"targetPartnerSearchNetwork,omitempty" json:"targetPartnerSearchNetwork,omitempty" yaml:"targetPartnerSearchNetwork,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Number value types for constructing number valued ranges.
type NumberValue interface{}

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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Specifies the page of results to return in the response. A page
// is specified             by the result position to start at
// and the maximum number of results to             return.
type Paging struct {
	StartIndex    *int `xml:"startIndex,omitempty" json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
	NumberResults *int `xml:"numberResults,omitempty" json:"numberResults,omitempty" yaml:"numberResults,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Represents Platform criterion.             <p>A criterion of
// this type can only be created using an ID.             <span
// class="constraint AdxEnabled">This is enabled for AdX.</span>
type Platform struct {
	Id            *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	PlatformName  *string        `xml:"platformName,omitempty" json:"platformName,omitempty" yaml:"platformName,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Platform) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Platform"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Represents a range of values that has either an upper or a lower
// bound.
type Range struct {
	Min *ComparableValue `xml:"min,omitempty" json:"min,omitempty" yaml:"min,omitempty"`
	Max *ComparableValue `xml:"max,omitempty" json:"max,omitempty" yaml:"max,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// A list of all errors associated with the @RegionCode constraints.
type RegionCodeError struct {
	FieldPath         *string                `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement    `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *RegionCodeErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                 `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *RegionCodeError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:RegionCodeError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// A {@link SearchParameter} for a query of {@code String}s.
//           <p>This element is supported by following {@link IdeaType}s:
// KEYWORD.             <p>This element is supported by following
// {@link RequestType}s: IDEAS, STATS.
type RelatedToQuerySearchParameter struct {
	SearchParameterType *string   `xml:"SearchParameter.Type,omitempty" json:"SearchParameter.Type,omitempty" yaml:"SearchParameter.Type,omitempty"`
	Queries             []*string `xml:"queries,omitempty" json:"queries,omitempty" yaml:"queries,omitempty"`
	TypeAttrXSI         string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *RelatedToQuerySearchParameter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:RelatedToQuerySearchParameter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// A {@link SearchParameter} that specifies a set of URLs that
// results should             in some way be related to. For example,
// keyword results would be             similar to content keywords
// found on the related URLs.             <p>This element is supported
// by following {@link IdeaType}s: KEYWORD.             <p>This
// element is supported by following {@link RequestType}s: IDEAS.
type RelatedToUrlSearchParameter struct {
	SearchParameterType *string   `xml:"SearchParameter.Type,omitempty" json:"SearchParameter.Type,omitempty" yaml:"SearchParameter.Type,omitempty"`
	Urls                []*string `xml:"urls,omitempty" json:"urls,omitempty" yaml:"urls,omitempty"`
	IncludeSubUrls      *bool     `xml:"includeSubUrls,omitempty" json:"includeSubUrls,omitempty" yaml:"includeSubUrls,omitempty"`
	TypeAttrXSI         string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *RelatedToUrlSearchParameter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:RelatedToUrlSearchParameter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// A set of {@link SearchParameter}s is supplied to the {@link
// TargetingIdeaSelector} to specify how             the user wants
// to filter the set of all possible {@link TargetingIdea}s.
//                        <p>There is a {@link SearchParameter}
// for each type of input.             {@link SearchParameter}s
// can conceptually be broken down into two types.</p>
//                  <ul>             <li>Input {@link SearchParameter}s
// provide the seed information from which             ideas or
// stats are to be generated (e.g., {@link RelatedToQuerySearchParameter},
//             {@link RelatedToUrlSearchParameter}, etc.). This
// type of {@link SearchParameters}             is required in
// requests.             <li>Filter {@link SearchParameter}s (e.g.,
// {@link CompetitionSearchParameter}, etc.)             are used
// to trim down the results based on {@link Attribute}-related
// information.             </ul>                          <p>A
// request should only contain one instance of each {@link SearchParameter}.
//             Requests containing multiple instances of the same
// search parameter will be             rejected.</p>
//    <p>One or more of the following {@link SearchParameter}s
// are required:<br/> <ul><li>{@link CategoryProductsAndServicesSearchParameter}</li>
// <li>{@link LocationSearchParameter}</li> <li>{@link RelatedToQuerySearchParameter}</li>
// <li>{@link RelatedToUrlSearchParameter}</li> <li>{@link SeedAdGroupIdSearchParameter}</li>
// </ul><p> <p><b>{@link IdeaType} KEYWORD supports following {@link
// SearchParameter}s:</b><br/> <ul> <li>{@link CategoryProductsAndServicesSearchParameter}</li>
// <li>{@link CompetitionSearchParameter}</li> <li>{@link IdeaTextFilterSearchParameter}</li>
// <li>{@link IncludeAdultContentSearchParameter}</li> <li>{@link
// LanguageSearchParameter}</li> <li>{@link LocationSearchParameter}</li>
// <li>{@link NetworkSearchParameter}</li> <li>{@link RelatedToQuerySearchParameter}</li>
// <li>{@link RelatedToUrlSearchParameter}</li> <li>{@link SearchVolumeSearchParameter}</li>
// <li>{@link SeedAdGroupIdSearchParameter}</li> </ul><p>
type SearchParameter interface{}

// A {@link SearchParameter} that specifies the level of search
// volume expected in results,             and it has a direct
// relationship to {@link AttributeType#SEARCH_VOLUME SEARCH_VOLUME}.
// Absence of             a {@link SearchVolumeSearchParameter}
// in a {@link TargetingIdeaSelector} is equivalent to having
//            no constraint on search volume specified.
//      <p>This element is supported by following {@link IdeaType}s:
// KEYWORD.             <p>This element is supported by following
// {@link RequestType}s: IDEAS, STATS.
type SearchVolumeSearchParameter struct {
	SearchParameterType *string                  `xml:"SearchParameter.Type,omitempty" json:"SearchParameter.Type,omitempty" yaml:"SearchParameter.Type,omitempty"`
	Operation           *LongComparisonOperation `xml:"operation,omitempty" json:"operation,omitempty" yaml:"operation,omitempty"`
	TypeAttrXSI         string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *SearchVolumeSearchParameter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SearchVolumeSearchParameter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// A {@link SearchParameter} for {@code KEYWORD} {@link IdeaType}s
//             that specifies that the supplied AdGroup should
// be used as a seed             for generating new ideas. For
// example, an AdGroup's keywords             would be used to
// generate new and related keywords.             <p>This element
// is supported by following {@link IdeaType}s: KEYWORD.
//       <p>This element is supported by following {@link RequestType}s:
// IDEAS.
type SeedAdGroupIdSearchParameter struct {
	SearchParameterType *string `xml:"SearchParameter.Type,omitempty" json:"SearchParameter.Type,omitempty" yaml:"SearchParameter.Type,omitempty"`
	AdGroupId           *int64  `xml:"adGroupId,omitempty" json:"adGroupId,omitempty" yaml:"adGroupId,omitempty"`
	TypeAttrXSI         string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *SeedAdGroupIdSearchParameter) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SeedAdGroupIdSearchParameter"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// {@link Attribute} type that contains a string value.
type StringAttribute struct {
	AttributeType *string `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *StringAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:StringAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Represents a {@link TargetingIdea} returned by search criteria
// specified in             the {@link TargetingIdeaSelector}.
// Targeting ideas are keywords or placements             that
// are similar to those the user inputs.
type TargetingIdea struct {
	Data []*Type_AttributeMapEntry `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
}

// Base error class for the {@link TargetingIdeaService}.
type TargetingIdeaError struct {
	FieldPath         *string                   `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement       `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                   `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                   `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                   `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *TargetingIdeaErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TargetingIdeaError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TargetingIdeaError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Contains a subset of {@link TargetingIdea}s from the search
// criteria             specified by a {@link TargetingIdeaSelector}.
type TargetingIdeaPage struct {
	TotalNumEntries *int             `xml:"totalNumEntries,omitempty" json:"totalNumEntries,omitempty" yaml:"totalNumEntries,omitempty"`
	Entries         []*TargetingIdea `xml:"entries,omitempty" json:"entries,omitempty" yaml:"entries,omitempty"`
}

// A descriptor for finding {@link TargetingIdea}s that match the
// specified criteria.
type TargetingIdeaSelector struct {
	SearchParameters        []*SearchParameter `xml:"searchParameters,omitempty" json:"searchParameters,omitempty" yaml:"searchParameters,omitempty"`
	IdeaType                *IdeaType          `xml:"ideaType,omitempty" json:"ideaType,omitempty" yaml:"ideaType,omitempty"`
	RequestType             *RequestType       `xml:"requestType,omitempty" json:"requestType,omitempty" yaml:"requestType,omitempty"`
	RequestedAttributeTypes []*AttributeType   `xml:"requestedAttributeTypes,omitempty" json:"requestedAttributeTypes,omitempty" yaml:"requestedAttributeTypes,omitempty"`
	Paging                  *Paging            `xml:"paging,omitempty" json:"paging,omitempty" yaml:"paging,omitempty"`
	LocaleCode              *string            `xml:"localeCode,omitempty" json:"localeCode,omitempty" yaml:"localeCode,omitempty"`
	CurrencyCode            *string            `xml:"currencyCode,omitempty" json:"currencyCode,omitempty" yaml:"currencyCode,omitempty"`
}

// Base error class for {@link TrafficEstimatorService}.
type TrafficEstimatorError struct {
	FieldPath         *string                      `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement          `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                      `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                      `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                      `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *TrafficEstimatorErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TrafficEstimatorError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TrafficEstimatorError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// This represents an entry in a map with a key of type Type
//           and value of type Attribute.
type Type_AttributeMapEntry struct {
	Key   *AttributeType `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value *Attribute     `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Basic information about a webpage.
type WebpageDescriptor struct {
	Url   *string `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	Title *string `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
}

// {@link Attribute} type that contains a {@link WebpageDescriptor}
// value.
type WebpageDescriptorAttribute struct {
	AttributeType *string            `xml:"Attribute.Type,omitempty" json:"Attribute.Type,omitempty" yaml:"Attribute.Type,omitempty"`
	Value         *WebpageDescriptor `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI   string             `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string             `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *WebpageDescriptorAttribute) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:WebpageDescriptorAttribute"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/o/v201809"
	}
}

// Get was auto-generated from WSDL.
type Get struct {
	Selector *TargetingIdeaSelector `xml:"selector,omitempty" json:"selector,omitempty" yaml:"selector,omitempty"`
}

// GetResponse was auto-generated from WSDL.
type GetResponse struct {
	Rval *TargetingIdeaPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
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

// targetingIdeaServiceInterface implements the TargetingIdeaServiceInterface interface.
type targetingIdeaServiceInterface struct {
	cli *soap.Client
}

// Returns a page of ideas that match the query described by the
// specified         {@link TargetingIdeaSelector}.
//       <p>The selector must specify a {@code paging} value, with
// {@code numberResults} set to 700 or         less.  Large result
// sets must be composed through multiple calls to this method,
// advancing the         paging {@code startIndex} value by {@code
// numberResults} with each call.</p>                  @param selector
// Query describing the types of results to return when
//  finding matches (similar keyword ideas).         @return A
// {@link TargetingIdeaPage} of results, that is a subset of the
//         list of possible ideas meeting the criteria of the
//        {@link TargetingIdeaSelector}.         @throws ApiException
// If problems occurred while querying for ideas.
func (p *targetingIdeaServiceInterface) Get(Get *Get) (*GetResponse, error) {
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
