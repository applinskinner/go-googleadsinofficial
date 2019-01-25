package adgroupextensionsetting

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "https://adwords.google.com/api/adwords/cm/v201809"

// NewAdGroupExtensionSettingServiceInterface creates an initializes a AdGroupExtensionSettingServiceInterface.
func NewAdGroupExtensionSettingServiceInterface(cli *soap.Client) AdGroupExtensionSettingServiceInterface {
	return &adGroupExtensionSettingServiceInterface{cli}
}

// AdGroupExtensionSettingServiceInterface was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type AdGroupExtensionSettingServiceInterface interface {
	// Returns a list of AdGroupExtensionSettings that meet the selector
	// criteria.                  @param selector Determines which
	// AdGroupExtensionSettings to return. If empty, all         AdGroupExtensionSettings
	// are returned.         @return The list of AdGroupExtensionSettings
	// specified by the selector.         @throws ApiException Indicates
	// a problem with the request.
	Get(Get *Get) (*GetResponse, error)

	// Applies the list of mutate operations (add, remove, and set).
	//                  <p> Beginning in v201509, add and set operations
	// are treated identically. Performing an add         operation
	// on an ad group with an existing ExtensionSetting will cause
	// the operation to be         treated like a set operation. Performing
	// a set operation on an ad group with no         ExtensionSetting
	// will cause the operation to be treated like an add operation.
	//                  @param operations The operations to apply.
	// The same {@link AdGroupExtensionSetting} cannot be         specified
	// in more than one operation.         @return The changed {@link
	// AdGroupExtensionSetting}s.         @throws ApiException Indicates
	// a problem with the request.
	Mutate(Mutate *Mutate) (*MutateResponse, error)

	// Returns a list of AdGroupExtensionSettings that match the query.
	//                  @param query The SQL-like AWQL query string.
	//         @return The list of AdGroupExtensionSettings specified
	// by the query.         @throws ApiException Indicates a problem
	// with the request.
	Query(Query string) (*QueryResponse, error)
}

// AppFeedItemAppStore was auto-generated from WSDL.
type AppFeedItemAppStore string

// Validate validates AppFeedItemAppStore.
func (v AppFeedItemAppStore) Validate() bool {
	for _, vv := range []string{
		"APPLE_ITUNES",
		"GOOGLE_PLAY",
		"UNKNOWN",
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

// DayOfWeek was auto-generated from WSDL.
type DayOfWeek string

// Validate validates DayOfWeek.
func (v DayOfWeek) Validate() bool {
	for _, vv := range []string{
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
		"SUNDAY",
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

// ExtensionSettingPlatform was auto-generated from WSDL.
type ExtensionSettingPlatform string

// Validate validates ExtensionSettingPlatform.
func (v ExtensionSettingPlatform) Validate() bool {
	for _, vv := range []string{
		"DESKTOP",
		"MOBILE",
		"NONE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ExtensionSettingErrorReason was auto-generated from WSDL.
type ExtensionSettingErrorReason string

// Validate validates ExtensionSettingErrorReason.
func (v ExtensionSettingErrorReason) Validate() bool {
	for _, vv := range []string{
		"EXTENSIONS_REQUIRED",
		"FEED_TYPE_EXTENSION_TYPE_MISMATCH",
		"INVALID_FEED_TYPE",
		"INVALID_FEED_TYPE_FOR_CUSTOMER_EXTENSION_SETTING",
		"CANNOT_CHANGE_FEED_ITEM_ON_ADD",
		"CANNOT_UPDATE_NEWLY_ADDED_EXTENSION",
		"NO_EXISTING_AD_GROUP_EXTENSION_SETTING_FOR_TYPE",
		"NO_EXISTING_CAMPAIGN_EXTENSION_SETTING_FOR_TYPE",
		"NO_EXISTING_CUSTOMER_EXTENSION_SETTING_FOR_TYPE",
		"AD_GROUP_EXTENSION_SETTING_ALREADY_EXISTS",
		"CAMPAIGN_EXTENSION_SETTING_ALREADY_EXISTS",
		"CUSTOMER_EXTENSION_SETTING_ALREADY_EXISTS",
		"AD_GROUP_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
		"CAMPAIGN_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
		"CUSTOMER_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
		"VALUE_OUT_OF_RANGE",
		"CANNOT_SET_WITH_FINAL_URLS",
		"CANNOT_SET_WITHOUT_FINAL_URLS",
		"CANNOT_SET_BOTH_DESTINATION_URL_AND_TRACKING_URL_TEMPLATE",
		"INVALID_PHONE_NUMBER",
		"PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY",
		"CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED",
		"PREMIUM_RATE_NUMBER_NOT_ALLOWED",
		"DISALLOWED_NUMBER_TYPE",
		"INVALID_DOMESTIC_PHONE_NUMBER_FORMAT",
		"VANITY_PHONE_NUMBER_NOT_ALLOWED",
		"INVALID_COUNTRY_CODE",
		"INVALID_CALL_CONVERSION_TYPE_ID",
		"CUSTOMER_NOT_WHITELISTED_FOR_CALLTRACKING",
		"CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY",
		"INVALID_APP_ID",
		"QUOTES_IN_REVIEW_EXTENSION_SNIPPET",
		"HYPHENS_IN_REVIEW_EXTENSION_SNIPPET",
		"REVIEW_EXTENSION_SOURCE_INELIGIBLE",
		"SOURCE_NAME_IN_REVIEW_EXTENSION_TEXT",
		"MISSING_FIELD",
		"INCONSISTENT_CURRENCY_CODES",
		"PRICE_EXTENSION_HAS_DUPLICATED_HEADERS",
		"PRICE_ITEM_HAS_DUPLICATED_HEADER_AND_DESCRIPTION",
		"PRICE_EXTENSION_HAS_TOO_FEW_ITEMS",
		"PRICE_EXTENSION_HAS_TOO_MANY_ITEMS",
		"UNSUPPORTED_VALUE",
		"UNSUPPORTED_VALUE_IN_SELECTED_LANGUAGE",
		"INVALID_DEVICE_PREFERENCE",
		"INVALID_SCHEDULE_END",
		"DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE",
		"OVERLAPPING_SCHEDULES",
		"SCHEDULE_END_NOT_AFTER_START",
		"TOO_MANY_SCHEDULES_PER_DAY",
		"DUPLICATE_EXTENSION_FEED_ITEM_EDIT",
		"INVALID_SNIPPETS_HEADER",
		"PHONE_NUMBER_NOT_SUPPORTED_WITH_CALLTRACKING_FOR_COUNTRY",
		"CAMPAIGN_TARGETING_MISMATCH",
		"CANNOT_OPERATE_ON_DELETED_FEED",
		"CONCRETE_EXTENSION_TYPE_REQUIRED",
		"INCOMPATIBLE_UNDERLYING_MATCHING_FUNCTION",
		"START_DATE_AFTER_END_DATE",
		"INVALID_PRICE_FORMAT",
		"PROMOTION_INVALID_TIME",
		"PROMOTION_CANNOT_SET_PERCENT_OFF_AND_MONEY_AMOUNT_OFF",
		"PROMOTION_CANNOT_SET_PROMOTION_CODE_AND_ORDERS_OVER_AMOUNT",
		"TOO_MANY_DECIMAL_PLACES_SPECIFIED",
		"INVALID_LANGUAGE_CODE",
		"UNSUPPORTED_LANGUAGE",
		"CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// FeedType was auto-generated from WSDL.
type FeedType string

// Validate validates FeedType.
func (v FeedType) Validate() bool {
	for _, vv := range []string{
		"NONE",
		"SITELINK",
		"CALL",
		"APP",
		"REVIEW",
		"AD_CUSTOMIZER",
		"CALLOUT",
		"STRUCTURED_SNIPPET",
		"MESSAGE",
		"PRICE",
		"PROMOTION",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// FeedItemStatus was auto-generated from WSDL.
type FeedItemStatus string

// Validate validates FeedItemStatus.
func (v FeedItemStatus) Validate() bool {
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

// FeedItemQualityApprovalStatus was auto-generated from WSDL.
type FeedItemQualityApprovalStatus string

// Validate validates FeedItemQualityApprovalStatus.
func (v FeedItemQualityApprovalStatus) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"APPROVED",
		"DISAPPROVED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// FeedItemQualityDisapprovalReasons was auto-generated from WSDL.
type FeedItemQualityDisapprovalReasons string

// Validate validates FeedItemQualityDisapprovalReasons.
func (v FeedItemQualityDisapprovalReasons) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"TABLE_REPETITIVE_HEADERS",
		"TABLE_REPETITIVE_DESCRIPTION",
		"TABLE_INCONSISTENT_ROWS",
		"DESCRIPTION_HAS_PRICE_QUALIFIERS",
		"UNSUPPORTED_LANGUAGE",
		"TABLE_ROW_HEADER_TABLE_TYPE_MISMATCH",
		"TABLE_ROW_HEADER_HAS_PROMOTIONAL_TEXT",
		"TABLE_ROW_DESCRIPTION_NOT_RELEVANT",
		"TABLE_ROW_DESCRIPTION_HAS_PROMOTIONAL_TEXT",
		"TABLE_ROW_HEADER_DESCRIPTION_REPETITIVE",
		"TABLE_ROW_UNRATEABLE",
		"TABLE_ROW_PRICE_INVALID",
		"TABLE_ROW_URL_INVALID",
		"HEADER_OR_DESCRIPTION_HAS_PRICE",
		"STRUCTURED_SNIPPETS_HEADER_POLICY_VIOLATED",
		"STRUCTURED_SNIPPETS_REPEATED_VALUES",
		"STRUCTURED_SNIPPETS_EDITORIAL_GUIDELINES",
		"STRUCTURED_SNIPPETS_HAS_PROMOTIONAL_TEXT",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// FeedItemValidationStatus was auto-generated from WSDL.
type FeedItemValidationStatus string

// Validate validates FeedItemValidationStatus.
func (v FeedItemValidationStatus) Validate() bool {
	for _, vv := range []string{
		"UNCHECKED",
		"ERROR",
		"VALID",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// GeoRestriction was auto-generated from WSDL.
type GeoRestriction string

// Validate validates GeoRestriction.
func (v GeoRestriction) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"LOCATION_OF_PRESENCE",
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

// MinuteOfHour was auto-generated from WSDL.
type MinuteOfHour string

// Validate validates MinuteOfHour.
func (v MinuteOfHour) Validate() bool {
	for _, vv := range []string{
		"ZERO",
		"FIFTEEN",
		"THIRTY",
		"FORTY_FIVE",
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

// PriceExtensionPriceQualifier was auto-generated from WSDL.
type PriceExtensionPriceQualifier string

// Validate validates PriceExtensionPriceQualifier.
func (v PriceExtensionPriceQualifier) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"FROM",
		"UP_TO",
		"AVERAGE",
		"NONE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PriceExtensionPriceUnit was auto-generated from WSDL.
type PriceExtensionPriceUnit string

// Validate validates PriceExtensionPriceUnit.
func (v PriceExtensionPriceUnit) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"PER_HOUR",
		"PER_DAY",
		"PER_WEEK",
		"PER_MONTH",
		"PER_YEAR",
		"PER_NIGHT",
		"NONE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PriceExtensionType was auto-generated from WSDL.
type PriceExtensionType string

// Validate validates PriceExtensionType.
func (v PriceExtensionType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"BRANDS",
		"EVENTS",
		"LOCATIONS",
		"NEIGHBORHOODS",
		"PRODUCT_CATEGORIES",
		"PRODUCT_TIERS",
		"SERVICES",
		"SERVICE_CATEGORIES",
		"SERVICE_TIERS",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PromotionExtensionDiscountModifier was auto-generated from WSDL.
type PromotionExtensionDiscountModifier string

// Validate validates PromotionExtensionDiscountModifier.
func (v PromotionExtensionDiscountModifier) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"UP_TO",
		"NONE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PromotionExtensionOccasion was auto-generated from WSDL.
type PromotionExtensionOccasion string

// Validate validates PromotionExtensionOccasion.
func (v PromotionExtensionOccasion) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"NEW_YEARS",
		"VALENTINES_DAY",
		"EASTER",
		"MOTHERS_DAY",
		"FATHERS_DAY",
		"LABOR_DAY",
		"BACK_TO_SCHOOL",
		"HALLOWEEN",
		"BLACK_FRIDAY",
		"CYBER_MONDAY",
		"CHRISTMAS",
		"BOXING_DAY",
		"NONE",
		"INDEPENDENCE_DAY",
		"NATIONAL_DAY",
		"END_OF_SEASON",
		"WINTER_SALE",
		"SUMMER_SALE",
		"FALL_SALE",
		"SPRING_SALE",
		"RAMADAN",
		"EID_AL_FITR",
		"EID_AL_ADHA",
		"SINGLES_DAY",
		"WOMENS_DAY",
		"HOLI",
		"PARENTS_DAY",
		"ST_NICHOLAS_DAY",
		"CHINESE_NEW_YEAR",
		"CARNIVAL",
		"EPIPHANY",
		"ROSH_HASHANAH",
		"PASSOVER",
		"HANUKKAH",
		"DIWALI",
		"NAVRATRI",
		"SONGKRAN",
		"YEAR_END_GIFT",
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

// An AdGroupExtensionSetting is used to add or modify extensions
// being served for the specified             ad group.
type AdGroupExtensionSetting struct {
	AdGroupId        *int64            `xml:"adGroupId,omitempty" json:"adGroupId,omitempty" yaml:"adGroupId,omitempty"`
	ExtensionType    *FeedType         `xml:"extensionType,omitempty" json:"extensionType,omitempty" yaml:"extensionType,omitempty"`
	ExtensionSetting *ExtensionSetting `xml:"extensionSetting,omitempty" json:"extensionSetting,omitempty" yaml:"extensionSetting,omitempty"`
}

// Operation used to create or mutate an AdGroupExtensionSetting.
type AdGroupExtensionSettingOperation struct {
	Operator      *Operator                `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	OperationType *string                  `xml:"Operation.Type,omitempty" json:"Operation.Type,omitempty" yaml:"Operation.Type,omitempty"`
	Operand       *AdGroupExtensionSetting `xml:"operand,omitempty" json:"operand,omitempty" yaml:"operand,omitempty"`
	TypeAttrXSI   string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupExtensionSettingOperation) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupExtensionSettingOperation"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Contains a subset of AdGroupExtensionSetting objects resulting
// from a             {@code AdGroupExtensionSettingService#get}
// call.
type AdGroupExtensionSettingPage struct {
	TotalNumEntries *int                       `xml:"totalNumEntries,omitempty" json:"totalNumEntries,omitempty" yaml:"totalNumEntries,omitempty"`
	PageType        *string                    `xml:"Page.Type,omitempty" json:"Page.Type,omitempty" yaml:"Page.Type,omitempty"`
	Entries         []*AdGroupExtensionSetting `xml:"entries,omitempty" json:"entries,omitempty" yaml:"entries,omitempty"`
	TypeAttrXSI     string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupExtensionSettingPage) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupExtensionSettingPage"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A container for return values from a {@code AdGroupExtensionSettingService#mutate}
// call.
type AdGroupExtensionSettingReturnValue struct {
	ListReturnValueType  *string                    `xml:"ListReturnValue.Type,omitempty" json:"ListReturnValue.Type,omitempty" yaml:"ListReturnValue.Type,omitempty"`
	Value                []*AdGroupExtensionSetting `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	PartialFailureErrors []*ApiError                `xml:"partialFailureErrors,omitempty" json:"partialFailureErrors,omitempty" yaml:"partialFailureErrors,omitempty"`
	TypeAttrXSI          string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupExtensionSettingReturnValue) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupExtensionSettingReturnValue"
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

// Represents an App extension.
type AppFeedItem struct {
	FeedId                  *int64                     `xml:"feedId,omitempty" json:"feedId,omitempty" yaml:"feedId,omitempty"`
	FeedItemId              *int64                     `xml:"feedItemId,omitempty" json:"feedItemId,omitempty" yaml:"feedItemId,omitempty"`
	Status                  *FeedItemStatus            `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	FeedType                *FeedType                  `xml:"feedType,omitempty" json:"feedType,omitempty" yaml:"feedType,omitempty"`
	StartTime               *string                    `xml:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty"`
	EndTime                 *string                    `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	DevicePreference        *FeedItemDevicePreference  `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	Scheduling              *FeedItemScheduling        `xml:"scheduling,omitempty" json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	CampaignTargeting       *FeedItemCampaignTargeting `xml:"campaignTargeting,omitempty" json:"campaignTargeting,omitempty" yaml:"campaignTargeting,omitempty"`
	AdGroupTargeting        *FeedItemAdGroupTargeting  `xml:"adGroupTargeting,omitempty" json:"adGroupTargeting,omitempty" yaml:"adGroupTargeting,omitempty"`
	KeywordTargeting        *Keyword                   `xml:"keywordTargeting,omitempty" json:"keywordTargeting,omitempty" yaml:"keywordTargeting,omitempty"`
	GeoTargeting            *Location                  `xml:"geoTargeting,omitempty" json:"geoTargeting,omitempty" yaml:"geoTargeting,omitempty"`
	GeoTargetingRestriction *FeedItemGeoRestriction    `xml:"geoTargetingRestriction,omitempty" json:"geoTargetingRestriction,omitempty" yaml:"geoTargetingRestriction,omitempty"`
	PolicySummaries         []*FeedItemPolicySummary   `xml:"policySummaries,omitempty" json:"policySummaries,omitempty" yaml:"policySummaries,omitempty"`
	ExtensionFeedItemType   *string                    `xml:"ExtensionFeedItem.Type,omitempty" json:"ExtensionFeedItem.Type,omitempty" yaml:"ExtensionFeedItem.Type,omitempty"`
	AppStore                *AppFeedItemAppStore       `xml:"appStore,omitempty" json:"appStore,omitempty" yaml:"appStore,omitempty"`
	AppId                   *string                    `xml:"appId,omitempty" json:"appId,omitempty" yaml:"appId,omitempty"`
	AppLinkText             *string                    `xml:"appLinkText,omitempty" json:"appLinkText,omitempty" yaml:"appLinkText,omitempty"`
	AppUrl                  *string                    `xml:"appUrl,omitempty" json:"appUrl,omitempty" yaml:"appUrl,omitempty"`
	AppFinalUrls            *UrlList                   `xml:"appFinalUrls,omitempty" json:"appFinalUrls,omitempty" yaml:"appFinalUrls,omitempty"`
	AppFinalMobileUrls      *UrlList                   `xml:"appFinalMobileUrls,omitempty" json:"appFinalMobileUrls,omitempty" yaml:"appFinalMobileUrls,omitempty"`
	AppTrackingUrlTemplate  *string                    `xml:"appTrackingUrlTemplate,omitempty" json:"appTrackingUrlTemplate,omitempty" yaml:"appTrackingUrlTemplate,omitempty"`
	AppFinalUrlSuffix       *string                    `xml:"appFinalUrlSuffix,omitempty" json:"appFinalUrlSuffix,omitempty" yaml:"appFinalUrlSuffix,omitempty"`
	AppUrlCustomParameters  *CustomParameters          `xml:"appUrlCustomParameters,omitempty" json:"appUrlCustomParameters,omitempty" yaml:"appUrlCustomParameters,omitempty"`
	TypeAttrXSI             string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace           string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AppFeedItem) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AppFeedItem"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
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

// Conversion type for a call extension.
type CallConversionType struct {
	ConversionTypeId *int64 `xml:"conversionTypeId,omitempty" json:"conversionTypeId,omitempty" yaml:"conversionTypeId,omitempty"`
}

// Represents a Call extension.
type CallFeedItem struct {
	FeedId                        *int64                     `xml:"feedId,omitempty" json:"feedId,omitempty" yaml:"feedId,omitempty"`
	FeedItemId                    *int64                     `xml:"feedItemId,omitempty" json:"feedItemId,omitempty" yaml:"feedItemId,omitempty"`
	Status                        *FeedItemStatus            `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	FeedType                      *FeedType                  `xml:"feedType,omitempty" json:"feedType,omitempty" yaml:"feedType,omitempty"`
	StartTime                     *string                    `xml:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty"`
	EndTime                       *string                    `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	DevicePreference              *FeedItemDevicePreference  `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	Scheduling                    *FeedItemScheduling        `xml:"scheduling,omitempty" json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	CampaignTargeting             *FeedItemCampaignTargeting `xml:"campaignTargeting,omitempty" json:"campaignTargeting,omitempty" yaml:"campaignTargeting,omitempty"`
	AdGroupTargeting              *FeedItemAdGroupTargeting  `xml:"adGroupTargeting,omitempty" json:"adGroupTargeting,omitempty" yaml:"adGroupTargeting,omitempty"`
	KeywordTargeting              *Keyword                   `xml:"keywordTargeting,omitempty" json:"keywordTargeting,omitempty" yaml:"keywordTargeting,omitempty"`
	GeoTargeting                  *Location                  `xml:"geoTargeting,omitempty" json:"geoTargeting,omitempty" yaml:"geoTargeting,omitempty"`
	GeoTargetingRestriction       *FeedItemGeoRestriction    `xml:"geoTargetingRestriction,omitempty" json:"geoTargetingRestriction,omitempty" yaml:"geoTargetingRestriction,omitempty"`
	PolicySummaries               []*FeedItemPolicySummary   `xml:"policySummaries,omitempty" json:"policySummaries,omitempty" yaml:"policySummaries,omitempty"`
	ExtensionFeedItemType         *string                    `xml:"ExtensionFeedItem.Type,omitempty" json:"ExtensionFeedItem.Type,omitempty" yaml:"ExtensionFeedItem.Type,omitempty"`
	CallPhoneNumber               *string                    `xml:"callPhoneNumber,omitempty" json:"callPhoneNumber,omitempty" yaml:"callPhoneNumber,omitempty"`
	CallCountryCode               *string                    `xml:"callCountryCode,omitempty" json:"callCountryCode,omitempty" yaml:"callCountryCode,omitempty"`
	CallTracking                  *bool                      `xml:"callTracking,omitempty" json:"callTracking,omitempty" yaml:"callTracking,omitempty"`
	CallConversionType            *CallConversionType        `xml:"callConversionType,omitempty" json:"callConversionType,omitempty" yaml:"callConversionType,omitempty"`
	DisableCallConversionTracking *bool                      `xml:"disableCallConversionTracking,omitempty" json:"disableCallConversionTracking,omitempty" yaml:"disableCallConversionTracking,omitempty"`
	TypeAttrXSI                   string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                 string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CallFeedItem) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CallFeedItem"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a callout extension.
type CalloutFeedItem struct {
	FeedId                  *int64                     `xml:"feedId,omitempty" json:"feedId,omitempty" yaml:"feedId,omitempty"`
	FeedItemId              *int64                     `xml:"feedItemId,omitempty" json:"feedItemId,omitempty" yaml:"feedItemId,omitempty"`
	Status                  *FeedItemStatus            `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	FeedType                *FeedType                  `xml:"feedType,omitempty" json:"feedType,omitempty" yaml:"feedType,omitempty"`
	StartTime               *string                    `xml:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty"`
	EndTime                 *string                    `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	DevicePreference        *FeedItemDevicePreference  `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	Scheduling              *FeedItemScheduling        `xml:"scheduling,omitempty" json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	CampaignTargeting       *FeedItemCampaignTargeting `xml:"campaignTargeting,omitempty" json:"campaignTargeting,omitempty" yaml:"campaignTargeting,omitempty"`
	AdGroupTargeting        *FeedItemAdGroupTargeting  `xml:"adGroupTargeting,omitempty" json:"adGroupTargeting,omitempty" yaml:"adGroupTargeting,omitempty"`
	KeywordTargeting        *Keyword                   `xml:"keywordTargeting,omitempty" json:"keywordTargeting,omitempty" yaml:"keywordTargeting,omitempty"`
	GeoTargeting            *Location                  `xml:"geoTargeting,omitempty" json:"geoTargeting,omitempty" yaml:"geoTargeting,omitempty"`
	GeoTargetingRestriction *FeedItemGeoRestriction    `xml:"geoTargetingRestriction,omitempty" json:"geoTargetingRestriction,omitempty" yaml:"geoTargetingRestriction,omitempty"`
	PolicySummaries         []*FeedItemPolicySummary   `xml:"policySummaries,omitempty" json:"policySummaries,omitempty" yaml:"policySummaries,omitempty"`
	ExtensionFeedItemType   *string                    `xml:"ExtensionFeedItem.Type,omitempty" json:"ExtensionFeedItem.Type,omitempty" yaml:"ExtensionFeedItem.Type,omitempty"`
	CalloutText             *string                    `xml:"calloutText,omitempty" json:"calloutText,omitempty" yaml:"calloutText,omitempty"`
	TypeAttrXSI             string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace           string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CalloutFeedItem) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CalloutFeedItem"
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

// Comparable types for constructing ranges with.
type ComparableValue interface{}

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

// Represents a criterion (such as a keyword, placement, or vertical).
//             <span class="constraint AdxEnabled">This is disabled
// for AdX when it is contained within Operators: ADD, SET.</span>
type Criterion struct {
	Id            *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
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

// Contains base extension feed item data for an extension in an
// extension feed managed by AdWords.
type ExtensionFeedItem struct {
	FeedId                  *int64                     `xml:"feedId,omitempty" json:"feedId,omitempty" yaml:"feedId,omitempty"`
	FeedItemId              *int64                     `xml:"feedItemId,omitempty" json:"feedItemId,omitempty" yaml:"feedItemId,omitempty"`
	Status                  *FeedItemStatus            `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	FeedType                *FeedType                  `xml:"feedType,omitempty" json:"feedType,omitempty" yaml:"feedType,omitempty"`
	StartTime               *string                    `xml:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty"`
	EndTime                 *string                    `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	DevicePreference        *FeedItemDevicePreference  `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	Scheduling              *FeedItemScheduling        `xml:"scheduling,omitempty" json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	CampaignTargeting       *FeedItemCampaignTargeting `xml:"campaignTargeting,omitempty" json:"campaignTargeting,omitempty" yaml:"campaignTargeting,omitempty"`
	AdGroupTargeting        *FeedItemAdGroupTargeting  `xml:"adGroupTargeting,omitempty" json:"adGroupTargeting,omitempty" yaml:"adGroupTargeting,omitempty"`
	KeywordTargeting        *Keyword                   `xml:"keywordTargeting,omitempty" json:"keywordTargeting,omitempty" yaml:"keywordTargeting,omitempty"`
	GeoTargeting            *Location                  `xml:"geoTargeting,omitempty" json:"geoTargeting,omitempty" yaml:"geoTargeting,omitempty"`
	GeoTargetingRestriction *FeedItemGeoRestriction    `xml:"geoTargetingRestriction,omitempty" json:"geoTargetingRestriction,omitempty" yaml:"geoTargetingRestriction,omitempty"`
	PolicySummaries         []*FeedItemPolicySummary   `xml:"policySummaries,omitempty" json:"policySummaries,omitempty" yaml:"policySummaries,omitempty"`
	ExtensionFeedItemType   *string                    `xml:"ExtensionFeedItem.Type,omitempty" json:"ExtensionFeedItem.Type,omitempty" yaml:"ExtensionFeedItem.Type,omitempty"`
}

// A setting specifying when and which extensions should serve
// at a given level (customer, campaign,             or ad group).
type ExtensionSetting struct {
	Extensions           []*ExtensionFeedItem      `xml:"extensions,omitempty" json:"extensions,omitempty" yaml:"extensions,omitempty"`
	PlatformRestrictions *ExtensionSettingPlatform `xml:"platformRestrictions,omitempty" json:"platformRestrictions,omitempty" yaml:"platformRestrictions,omitempty"`
}

// Represents an error for various extension setting services.
type ExtensionSettingError struct {
	FieldPath         *string                      `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement          `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                      `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                      `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                      `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *ExtensionSettingErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ExtensionSettingError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ExtensionSettingError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Specifies the adgroup the request context must match in order
// for             the feed item to be considered eligible for
// serving (aka the targeted adgroup).             E.g., if the
// below adgroup targeting is set to adgroup = X, then the feed
//             item can only serve under adgroup X.
type FeedItemAdGroupTargeting struct {
	TargetingAdGroupId *int64 `xml:"TargetingAdGroupId,omitempty" json:"TargetingAdGroupId,omitempty" yaml:"TargetingAdGroupId,omitempty"`
}

// Contains validation error details for a set of feed attributes.
type FeedItemAttributeError struct {
	FeedAttributeIds    []*int64 `xml:"feedAttributeIds,omitempty" json:"feedAttributeIds,omitempty" yaml:"feedAttributeIds,omitempty"`
	ValidationErrorCode *int     `xml:"validationErrorCode,omitempty" json:"validationErrorCode,omitempty" yaml:"validationErrorCode,omitempty"`
	ErrorInformation    *string  `xml:"errorInformation,omitempty" json:"errorInformation,omitempty" yaml:"errorInformation,omitempty"`
}

// Specifies the campaign the request context must match in order
// for             the feed item to be considered eligible for
// serving (aka the targeted campaign).             E.g., if the
// below campaign targeting is set to campaignId = X, then the
// feed             item can only serve under campaign X.
type FeedItemCampaignTargeting struct {
	TargetingCampaignId *int64 `xml:"TargetingCampaignId,omitempty" json:"TargetingCampaignId,omitempty" yaml:"TargetingCampaignId,omitempty"`
}

// Represents a FeedItem device preference.
type FeedItemDevicePreference struct {
	DevicePreference *int64 `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
}

// Represents a FeedItem geo restriction.
type FeedItemGeoRestriction struct {
	GeoRestriction *GeoRestriction `xml:"geoRestriction,omitempty" json:"geoRestriction,omitempty" yaml:"geoRestriction,omitempty"`
}

// Contains offline validation, policy findings, and approval results
// for a FeedItem in the context             of a FeedMapping.
type FeedItemPolicySummary struct {
	PolicyTopicEntries        []*PolicyTopicEntry                  `xml:"policyTopicEntries,omitempty" json:"policyTopicEntries,omitempty" yaml:"policyTopicEntries,omitempty"`
	ReviewState               *PolicySummaryReviewState            `xml:"reviewState,omitempty" json:"reviewState,omitempty" yaml:"reviewState,omitempty"`
	DenormalizedStatus        *PolicySummaryDenormalizedStatus     `xml:"denormalizedStatus,omitempty" json:"denormalizedStatus,omitempty" yaml:"denormalizedStatus,omitempty"`
	CombinedApprovalStatus    *PolicyApprovalStatus                `xml:"combinedApprovalStatus,omitempty" json:"combinedApprovalStatus,omitempty" yaml:"combinedApprovalStatus,omitempty"`
	PolicySummaryInfoType     *string                              `xml:"PolicySummaryInfo.Type,omitempty" json:"PolicySummaryInfo.Type,omitempty" yaml:"PolicySummaryInfo.Type,omitempty"`
	FeedMappingId             *int64                               `xml:"feedMappingId,omitempty" json:"feedMappingId,omitempty" yaml:"feedMappingId,omitempty"`
	ValidationStatus          *FeedItemValidationStatus            `xml:"validationStatus,omitempty" json:"validationStatus,omitempty" yaml:"validationStatus,omitempty"`
	ValidationErrors          []*FeedItemAttributeError            `xml:"validationErrors,omitempty" json:"validationErrors,omitempty" yaml:"validationErrors,omitempty"`
	QualityApprovalStatus     *FeedItemQualityApprovalStatus       `xml:"qualityApprovalStatus,omitempty" json:"qualityApprovalStatus,omitempty" yaml:"qualityApprovalStatus,omitempty"`
	QualityDisapprovalReasons []*FeedItemQualityDisapprovalReasons `xml:"qualityDisapprovalReasons,omitempty" json:"qualityDisapprovalReasons,omitempty" yaml:"qualityDisapprovalReasons,omitempty"`
	TypeAttrXSI               string                               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string                               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *FeedItemPolicySummary) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:FeedItemPolicySummary"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a FeedItem schedule, which specifies a time interval
// on a given day             when the feed item may serve. The
// FeedItemSchedule times are in the account's time zone.
type FeedItemSchedule struct {
	DayOfWeek   *DayOfWeek    `xml:"dayOfWeek,omitempty" json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`
	StartHour   *int          `xml:"startHour,omitempty" json:"startHour,omitempty" yaml:"startHour,omitempty"`
	StartMinute *MinuteOfHour `xml:"startMinute,omitempty" json:"startMinute,omitempty" yaml:"startMinute,omitempty"`
	EndHour     *int          `xml:"endHour,omitempty" json:"endHour,omitempty" yaml:"endHour,omitempty"`
	EndMinute   *MinuteOfHour `xml:"endMinute,omitempty" json:"endMinute,omitempty" yaml:"endMinute,omitempty"`
}

// Represents a collection of FeedItem schedules specifying all
// time intervals for which             the feed item may serve.
// Any time range not covered by the specified FeedItemSchedules
// will             prevent the feed item from serving during those
// times.
type FeedItemScheduling struct {
	FeedItemSchedules []*FeedItemSchedule `xml:"feedItemSchedules,omitempty" json:"feedItemSchedules,omitempty" yaml:"feedItemSchedules,omitempty"`
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

// Base list return value type.
type ListReturnValue interface{}

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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents a Message extension.
type MessageFeedItem struct {
	FeedId                  *int64                     `xml:"feedId,omitempty" json:"feedId,omitempty" yaml:"feedId,omitempty"`
	FeedItemId              *int64                     `xml:"feedItemId,omitempty" json:"feedItemId,omitempty" yaml:"feedItemId,omitempty"`
	Status                  *FeedItemStatus            `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	FeedType                *FeedType                  `xml:"feedType,omitempty" json:"feedType,omitempty" yaml:"feedType,omitempty"`
	StartTime               *string                    `xml:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty"`
	EndTime                 *string                    `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	DevicePreference        *FeedItemDevicePreference  `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	Scheduling              *FeedItemScheduling        `xml:"scheduling,omitempty" json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	CampaignTargeting       *FeedItemCampaignTargeting `xml:"campaignTargeting,omitempty" json:"campaignTargeting,omitempty" yaml:"campaignTargeting,omitempty"`
	AdGroupTargeting        *FeedItemAdGroupTargeting  `xml:"adGroupTargeting,omitempty" json:"adGroupTargeting,omitempty" yaml:"adGroupTargeting,omitempty"`
	KeywordTargeting        *Keyword                   `xml:"keywordTargeting,omitempty" json:"keywordTargeting,omitempty" yaml:"keywordTargeting,omitempty"`
	GeoTargeting            *Location                  `xml:"geoTargeting,omitempty" json:"geoTargeting,omitempty" yaml:"geoTargeting,omitempty"`
	GeoTargetingRestriction *FeedItemGeoRestriction    `xml:"geoTargetingRestriction,omitempty" json:"geoTargetingRestriction,omitempty" yaml:"geoTargetingRestriction,omitempty"`
	PolicySummaries         []*FeedItemPolicySummary   `xml:"policySummaries,omitempty" json:"policySummaries,omitempty" yaml:"policySummaries,omitempty"`
	ExtensionFeedItemType   *string                    `xml:"ExtensionFeedItem.Type,omitempty" json:"ExtensionFeedItem.Type,omitempty" yaml:"ExtensionFeedItem.Type,omitempty"`
	MessageBusinessName     *string                    `xml:"messageBusinessName,omitempty" json:"messageBusinessName,omitempty" yaml:"messageBusinessName,omitempty"`
	MessageCountryCode      *string                    `xml:"messageCountryCode,omitempty" json:"messageCountryCode,omitempty" yaml:"messageCountryCode,omitempty"`
	MessagePhoneNumber      *string                    `xml:"messagePhoneNumber,omitempty" json:"messagePhoneNumber,omitempty" yaml:"messagePhoneNumber,omitempty"`
	MessageExtensionText    *string                    `xml:"messageExtensionText,omitempty" json:"messageExtensionText,omitempty" yaml:"messageExtensionText,omitempty"`
	MessageText             *string                    `xml:"messageText,omitempty" json:"messageText,omitempty" yaml:"messageText,omitempty"`
	TypeAttrXSI             string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace           string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MessageFeedItem) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MessageFeedItem"
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

// Represents a money amount with currency.
type MoneyWithCurrency struct {
	ComparableValueType *string `xml:"ComparableValue.Type,omitempty" json:"ComparableValue.Type,omitempty" yaml:"ComparableValue.Type,omitempty"`
	Money               *Money  `xml:"money,omitempty" json:"money,omitempty" yaml:"money,omitempty"`
	CurrencyCode        *string `xml:"currencyCode,omitempty" json:"currencyCode,omitempty" yaml:"currencyCode,omitempty"`
	TypeAttrXSI         string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MoneyWithCurrency) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MoneyWithCurrency"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
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

// Specifies how an entity (eg. adgroup, campaign, criterion, ad)
// should be filtered.
type Predicate struct {
	Field    *string            `xml:"field,omitempty" json:"field,omitempty" yaml:"field,omitempty"`
	Operator *PredicateOperator `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	Values   []*string          `xml:"values,omitempty" json:"values,omitempty" yaml:"values,omitempty"`
}

// Represents a price extension.
type PriceFeedItem struct {
	FeedId                  *int64                        `xml:"feedId,omitempty" json:"feedId,omitempty" yaml:"feedId,omitempty"`
	FeedItemId              *int64                        `xml:"feedItemId,omitempty" json:"feedItemId,omitempty" yaml:"feedItemId,omitempty"`
	Status                  *FeedItemStatus               `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	FeedType                *FeedType                     `xml:"feedType,omitempty" json:"feedType,omitempty" yaml:"feedType,omitempty"`
	StartTime               *string                       `xml:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty"`
	EndTime                 *string                       `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	DevicePreference        *FeedItemDevicePreference     `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	Scheduling              *FeedItemScheduling           `xml:"scheduling,omitempty" json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	CampaignTargeting       *FeedItemCampaignTargeting    `xml:"campaignTargeting,omitempty" json:"campaignTargeting,omitempty" yaml:"campaignTargeting,omitempty"`
	AdGroupTargeting        *FeedItemAdGroupTargeting     `xml:"adGroupTargeting,omitempty" json:"adGroupTargeting,omitempty" yaml:"adGroupTargeting,omitempty"`
	KeywordTargeting        *Keyword                      `xml:"keywordTargeting,omitempty" json:"keywordTargeting,omitempty" yaml:"keywordTargeting,omitempty"`
	GeoTargeting            *Location                     `xml:"geoTargeting,omitempty" json:"geoTargeting,omitempty" yaml:"geoTargeting,omitempty"`
	GeoTargetingRestriction *FeedItemGeoRestriction       `xml:"geoTargetingRestriction,omitempty" json:"geoTargetingRestriction,omitempty" yaml:"geoTargetingRestriction,omitempty"`
	PolicySummaries         []*FeedItemPolicySummary      `xml:"policySummaries,omitempty" json:"policySummaries,omitempty" yaml:"policySummaries,omitempty"`
	ExtensionFeedItemType   *string                       `xml:"ExtensionFeedItem.Type,omitempty" json:"ExtensionFeedItem.Type,omitempty" yaml:"ExtensionFeedItem.Type,omitempty"`
	PriceExtensionType      *PriceExtensionType           `xml:"priceExtensionType,omitempty" json:"priceExtensionType,omitempty" yaml:"priceExtensionType,omitempty"`
	PriceQualifier          *PriceExtensionPriceQualifier `xml:"priceQualifier,omitempty" json:"priceQualifier,omitempty" yaml:"priceQualifier,omitempty"`
	TrackingUrlTemplate     *string                       `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix          *string                       `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	Language                *string                       `xml:"language,omitempty" json:"language,omitempty" yaml:"language,omitempty"`
	TableRows               []*PriceTableRow              `xml:"tableRows,omitempty" json:"tableRows,omitempty" yaml:"tableRows,omitempty"`
	TypeAttrXSI             string                        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace           string                        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *PriceFeedItem) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:PriceFeedItem"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents one row in a price extension.
type PriceTableRow struct {
	Header          *string                  `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	Description     *string                  `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	FinalUrls       *UrlList                 `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls *UrlList                 `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	Price           *MoneyWithCurrency       `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
	PriceUnit       *PriceExtensionPriceUnit `xml:"priceUnit,omitempty" json:"priceUnit,omitempty" yaml:"priceUnit,omitempty"`
}

// Represents a promotion extension.
type PromotionFeedItem struct {
	FeedId                       *int64                              `xml:"feedId,omitempty" json:"feedId,omitempty" yaml:"feedId,omitempty"`
	FeedItemId                   *int64                              `xml:"feedItemId,omitempty" json:"feedItemId,omitempty" yaml:"feedItemId,omitempty"`
	Status                       *FeedItemStatus                     `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	FeedType                     *FeedType                           `xml:"feedType,omitempty" json:"feedType,omitempty" yaml:"feedType,omitempty"`
	StartTime                    *string                             `xml:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty"`
	EndTime                      *string                             `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	DevicePreference             *FeedItemDevicePreference           `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	Scheduling                   *FeedItemScheduling                 `xml:"scheduling,omitempty" json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	CampaignTargeting            *FeedItemCampaignTargeting          `xml:"campaignTargeting,omitempty" json:"campaignTargeting,omitempty" yaml:"campaignTargeting,omitempty"`
	AdGroupTargeting             *FeedItemAdGroupTargeting           `xml:"adGroupTargeting,omitempty" json:"adGroupTargeting,omitempty" yaml:"adGroupTargeting,omitempty"`
	KeywordTargeting             *Keyword                            `xml:"keywordTargeting,omitempty" json:"keywordTargeting,omitempty" yaml:"keywordTargeting,omitempty"`
	GeoTargeting                 *Location                           `xml:"geoTargeting,omitempty" json:"geoTargeting,omitempty" yaml:"geoTargeting,omitempty"`
	GeoTargetingRestriction      *FeedItemGeoRestriction             `xml:"geoTargetingRestriction,omitempty" json:"geoTargetingRestriction,omitempty" yaml:"geoTargetingRestriction,omitempty"`
	PolicySummaries              []*FeedItemPolicySummary            `xml:"policySummaries,omitempty" json:"policySummaries,omitempty" yaml:"policySummaries,omitempty"`
	ExtensionFeedItemType        *string                             `xml:"ExtensionFeedItem.Type,omitempty" json:"ExtensionFeedItem.Type,omitempty" yaml:"ExtensionFeedItem.Type,omitempty"`
	PromotionTarget              *string                             `xml:"promotionTarget,omitempty" json:"promotionTarget,omitempty" yaml:"promotionTarget,omitempty"`
	DiscountModifier             *PromotionExtensionDiscountModifier `xml:"discountModifier,omitempty" json:"discountModifier,omitempty" yaml:"discountModifier,omitempty"`
	PercentOff                   *int64                              `xml:"percentOff,omitempty" json:"percentOff,omitempty" yaml:"percentOff,omitempty"`
	MoneyAmountOff               *MoneyWithCurrency                  `xml:"moneyAmountOff,omitempty" json:"moneyAmountOff,omitempty" yaml:"moneyAmountOff,omitempty"`
	PromotionCode                *string                             `xml:"promotionCode,omitempty" json:"promotionCode,omitempty" yaml:"promotionCode,omitempty"`
	OrdersOverAmount             *MoneyWithCurrency                  `xml:"ordersOverAmount,omitempty" json:"ordersOverAmount,omitempty" yaml:"ordersOverAmount,omitempty"`
	PromotionStart               *string                             `xml:"promotionStart,omitempty" json:"promotionStart,omitempty" yaml:"promotionStart,omitempty"`
	PromotionEnd                 *string                             `xml:"promotionEnd,omitempty" json:"promotionEnd,omitempty" yaml:"promotionEnd,omitempty"`
	Occasion                     *PromotionExtensionOccasion         `xml:"occasion,omitempty" json:"occasion,omitempty" yaml:"occasion,omitempty"`
	FinalUrls                    *UrlList                            `xml:"finalUrls,omitempty" json:"finalUrls,omitempty" yaml:"finalUrls,omitempty"`
	FinalMobileUrls              *UrlList                            `xml:"finalMobileUrls,omitempty" json:"finalMobileUrls,omitempty" yaml:"finalMobileUrls,omitempty"`
	TrackingUrlTemplate          *string                             `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix               *string                             `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	PromotionUrlCustomParameters *CustomParameters                   `xml:"promotionUrlCustomParameters,omitempty" json:"promotionUrlCustomParameters,omitempty" yaml:"promotionUrlCustomParameters,omitempty"`
	Language                     *string                             `xml:"language,omitempty" json:"language,omitempty" yaml:"language,omitempty"`
	TypeAttrXSI                  string                              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                string                              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *PromotionFeedItem) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:PromotionFeedItem"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
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

// Represents a Review extension.
type ReviewFeedItem struct {
	FeedId                  *int64                     `xml:"feedId,omitempty" json:"feedId,omitempty" yaml:"feedId,omitempty"`
	FeedItemId              *int64                     `xml:"feedItemId,omitempty" json:"feedItemId,omitempty" yaml:"feedItemId,omitempty"`
	Status                  *FeedItemStatus            `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	FeedType                *FeedType                  `xml:"feedType,omitempty" json:"feedType,omitempty" yaml:"feedType,omitempty"`
	StartTime               *string                    `xml:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty"`
	EndTime                 *string                    `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	DevicePreference        *FeedItemDevicePreference  `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	Scheduling              *FeedItemScheduling        `xml:"scheduling,omitempty" json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	CampaignTargeting       *FeedItemCampaignTargeting `xml:"campaignTargeting,omitempty" json:"campaignTargeting,omitempty" yaml:"campaignTargeting,omitempty"`
	AdGroupTargeting        *FeedItemAdGroupTargeting  `xml:"adGroupTargeting,omitempty" json:"adGroupTargeting,omitempty" yaml:"adGroupTargeting,omitempty"`
	KeywordTargeting        *Keyword                   `xml:"keywordTargeting,omitempty" json:"keywordTargeting,omitempty" yaml:"keywordTargeting,omitempty"`
	GeoTargeting            *Location                  `xml:"geoTargeting,omitempty" json:"geoTargeting,omitempty" yaml:"geoTargeting,omitempty"`
	GeoTargetingRestriction *FeedItemGeoRestriction    `xml:"geoTargetingRestriction,omitempty" json:"geoTargetingRestriction,omitempty" yaml:"geoTargetingRestriction,omitempty"`
	PolicySummaries         []*FeedItemPolicySummary   `xml:"policySummaries,omitempty" json:"policySummaries,omitempty" yaml:"policySummaries,omitempty"`
	ExtensionFeedItemType   *string                    `xml:"ExtensionFeedItem.Type,omitempty" json:"ExtensionFeedItem.Type,omitempty" yaml:"ExtensionFeedItem.Type,omitempty"`
	ReviewText              *string                    `xml:"reviewText,omitempty" json:"reviewText,omitempty" yaml:"reviewText,omitempty"`
	ReviewSourceName        *string                    `xml:"reviewSourceName,omitempty" json:"reviewSourceName,omitempty" yaml:"reviewSourceName,omitempty"`
	ReviewSourceUrl         *string                    `xml:"reviewSourceUrl,omitempty" json:"reviewSourceUrl,omitempty" yaml:"reviewSourceUrl,omitempty"`
	ReviewTextExactlyQuoted *bool                      `xml:"reviewTextExactlyQuoted,omitempty" json:"reviewTextExactlyQuoted,omitempty" yaml:"reviewTextExactlyQuoted,omitempty"`
	TypeAttrXSI             string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace           string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ReviewFeedItem) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ReviewFeedItem"
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

// Represents a sitelink extension.
type SitelinkFeedItem struct {
	FeedId                      *int64                     `xml:"feedId,omitempty" json:"feedId,omitempty" yaml:"feedId,omitempty"`
	FeedItemId                  *int64                     `xml:"feedItemId,omitempty" json:"feedItemId,omitempty" yaml:"feedItemId,omitempty"`
	Status                      *FeedItemStatus            `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	FeedType                    *FeedType                  `xml:"feedType,omitempty" json:"feedType,omitempty" yaml:"feedType,omitempty"`
	StartTime                   *string                    `xml:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty"`
	EndTime                     *string                    `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	DevicePreference            *FeedItemDevicePreference  `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	Scheduling                  *FeedItemScheduling        `xml:"scheduling,omitempty" json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	CampaignTargeting           *FeedItemCampaignTargeting `xml:"campaignTargeting,omitempty" json:"campaignTargeting,omitempty" yaml:"campaignTargeting,omitempty"`
	AdGroupTargeting            *FeedItemAdGroupTargeting  `xml:"adGroupTargeting,omitempty" json:"adGroupTargeting,omitempty" yaml:"adGroupTargeting,omitempty"`
	KeywordTargeting            *Keyword                   `xml:"keywordTargeting,omitempty" json:"keywordTargeting,omitempty" yaml:"keywordTargeting,omitempty"`
	GeoTargeting                *Location                  `xml:"geoTargeting,omitempty" json:"geoTargeting,omitempty" yaml:"geoTargeting,omitempty"`
	GeoTargetingRestriction     *FeedItemGeoRestriction    `xml:"geoTargetingRestriction,omitempty" json:"geoTargetingRestriction,omitempty" yaml:"geoTargetingRestriction,omitempty"`
	PolicySummaries             []*FeedItemPolicySummary   `xml:"policySummaries,omitempty" json:"policySummaries,omitempty" yaml:"policySummaries,omitempty"`
	ExtensionFeedItemType       *string                    `xml:"ExtensionFeedItem.Type,omitempty" json:"ExtensionFeedItem.Type,omitempty" yaml:"ExtensionFeedItem.Type,omitempty"`
	SitelinkText                *string                    `xml:"sitelinkText,omitempty" json:"sitelinkText,omitempty" yaml:"sitelinkText,omitempty"`
	SitelinkUrl                 *string                    `xml:"sitelinkUrl,omitempty" json:"sitelinkUrl,omitempty" yaml:"sitelinkUrl,omitempty"`
	SitelinkLine2               *string                    `xml:"sitelinkLine2,omitempty" json:"sitelinkLine2,omitempty" yaml:"sitelinkLine2,omitempty"`
	SitelinkLine3               *string                    `xml:"sitelinkLine3,omitempty" json:"sitelinkLine3,omitempty" yaml:"sitelinkLine3,omitempty"`
	SitelinkFinalUrls           *UrlList                   `xml:"sitelinkFinalUrls,omitempty" json:"sitelinkFinalUrls,omitempty" yaml:"sitelinkFinalUrls,omitempty"`
	SitelinkFinalMobileUrls     *UrlList                   `xml:"sitelinkFinalMobileUrls,omitempty" json:"sitelinkFinalMobileUrls,omitempty" yaml:"sitelinkFinalMobileUrls,omitempty"`
	SitelinkTrackingUrlTemplate *string                    `xml:"sitelinkTrackingUrlTemplate,omitempty" json:"sitelinkTrackingUrlTemplate,omitempty" yaml:"sitelinkTrackingUrlTemplate,omitempty"`
	SitelinkFinalUrlSuffix      *string                    `xml:"sitelinkFinalUrlSuffix,omitempty" json:"sitelinkFinalUrlSuffix,omitempty" yaml:"sitelinkFinalUrlSuffix,omitempty"`
	SitelinkUrlCustomParameters *CustomParameters          `xml:"sitelinkUrlCustomParameters,omitempty" json:"sitelinkUrlCustomParameters,omitempty" yaml:"sitelinkUrlCustomParameters,omitempty"`
	TypeAttrXSI                 string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace               string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *SitelinkFeedItem) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SitelinkFeedItem"
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

// Represents a structured snippet extension.
type StructuredSnippetFeedItem struct {
	FeedId                  *int64                     `xml:"feedId,omitempty" json:"feedId,omitempty" yaml:"feedId,omitempty"`
	FeedItemId              *int64                     `xml:"feedItemId,omitempty" json:"feedItemId,omitempty" yaml:"feedItemId,omitempty"`
	Status                  *FeedItemStatus            `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	FeedType                *FeedType                  `xml:"feedType,omitempty" json:"feedType,omitempty" yaml:"feedType,omitempty"`
	StartTime               *string                    `xml:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty"`
	EndTime                 *string                    `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	DevicePreference        *FeedItemDevicePreference  `xml:"devicePreference,omitempty" json:"devicePreference,omitempty" yaml:"devicePreference,omitempty"`
	Scheduling              *FeedItemScheduling        `xml:"scheduling,omitempty" json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	CampaignTargeting       *FeedItemCampaignTargeting `xml:"campaignTargeting,omitempty" json:"campaignTargeting,omitempty" yaml:"campaignTargeting,omitempty"`
	AdGroupTargeting        *FeedItemAdGroupTargeting  `xml:"adGroupTargeting,omitempty" json:"adGroupTargeting,omitempty" yaml:"adGroupTargeting,omitempty"`
	KeywordTargeting        *Keyword                   `xml:"keywordTargeting,omitempty" json:"keywordTargeting,omitempty" yaml:"keywordTargeting,omitempty"`
	GeoTargeting            *Location                  `xml:"geoTargeting,omitempty" json:"geoTargeting,omitempty" yaml:"geoTargeting,omitempty"`
	GeoTargetingRestriction *FeedItemGeoRestriction    `xml:"geoTargetingRestriction,omitempty" json:"geoTargetingRestriction,omitempty" yaml:"geoTargetingRestriction,omitempty"`
	PolicySummaries         []*FeedItemPolicySummary   `xml:"policySummaries,omitempty" json:"policySummaries,omitempty" yaml:"policySummaries,omitempty"`
	ExtensionFeedItemType   *string                    `xml:"ExtensionFeedItem.Type,omitempty" json:"ExtensionFeedItem.Type,omitempty" yaml:"ExtensionFeedItem.Type,omitempty"`
	Header                  *string                    `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	Values                  []*string                  `xml:"values,omitempty" json:"values,omitempty" yaml:"values,omitempty"`
	TypeAttrXSI             string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace           string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *StructuredSnippetFeedItem) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:StructuredSnippetFeedItem"
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

// Get was auto-generated from WSDL.
type Get struct {
	Selector *Selector `xml:"selector,omitempty" json:"selector,omitempty" yaml:"selector,omitempty"`
}

// GetResponse was auto-generated from WSDL.
type GetResponse struct {
	Rval *AdGroupExtensionSettingPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Mutate was auto-generated from WSDL.
type Mutate struct {
	Operations []*AdGroupExtensionSettingOperation `xml:"operations,omitempty" json:"operations,omitempty" yaml:"operations,omitempty"`
}

// MutateResponse was auto-generated from WSDL.
type MutateResponse struct {
	Rval *AdGroupExtensionSettingReturnValue `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Query was auto-generated from WSDL.
type Query struct {
	Query *string `xml:"query,omitempty" json:"query,omitempty" yaml:"query,omitempty"`
}

// QueryResponse was auto-generated from WSDL.
type QueryResponse struct {
	Rval *AdGroupExtensionSettingPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
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

// adGroupExtensionSettingServiceInterface implements the AdGroupExtensionSettingServiceInterface interface.
type adGroupExtensionSettingServiceInterface struct {
	cli *soap.Client
}

// Returns a list of AdGroupExtensionSettings that meet the selector
// criteria.                  @param selector Determines which
// AdGroupExtensionSettings to return. If empty, all         AdGroupExtensionSettings
// are returned.         @return The list of AdGroupExtensionSettings
// specified by the selector.         @throws ApiException Indicates
// a problem with the request.
func (p *adGroupExtensionSettingServiceInterface) Get(Get *Get) (*GetResponse, error) {
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

// Applies the list of mutate operations (add, remove, and set).
//                  <p> Beginning in v201509, add and set operations
// are treated identically. Performing an add         operation
// on an ad group with an existing ExtensionSetting will cause
// the operation to be         treated like a set operation. Performing
// a set operation on an ad group with no         ExtensionSetting
// will cause the operation to be treated like an add operation.
//                  @param operations The operations to apply.
// The same {@link AdGroupExtensionSetting} cannot be         specified
// in more than one operation.         @return The changed {@link
// AdGroupExtensionSetting}s.         @throws ApiException Indicates
// a problem with the request.
func (p *adGroupExtensionSettingServiceInterface) Mutate(Mutate *Mutate) (*MutateResponse, error) {
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

// Returns a list of AdGroupExtensionSettings that match the query.
//                  @param query The SQL-like AWQL query string.
//         @return The list of AdGroupExtensionSettings specified
// by the query.         @throws ApiException Indicates a problem
// with the request.
func (p *adGroupExtensionSettingServiceInterface) Query(Query string) (*QueryResponse, error) {
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
