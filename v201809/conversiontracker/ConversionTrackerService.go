package conversiontracker

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "https://adwords.google.com/api/adwords/cm/v201809"

// NewConversionTrackerServiceInterface creates an initializes a ConversionTrackerServiceInterface.
func NewConversionTrackerServiceInterface(cli *soap.Client) ConversionTrackerServiceInterface {
	return &conversionTrackerServiceInterface{cli}
}

// ConversionTrackerServiceInterface was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ConversionTrackerServiceInterface interface {
	// Returns a list of the conversion trackers that match the selector.
	// The         actual objects contained in the page's list of entries
	// will be specific         subclasses of the abstract {@link ConversionTracker}
	// class.                  @param serviceSelector The selector
	// specifying the         {@link ConversionTracker}s to return.
	//         @return List of conversion trackers specified by the
	// selector.         @throws com.google.ads.api.services.common.error.ApiException
	// if problems         occurred while retrieving results.
	Get(Get *Get) (*GetResponse, error)

	// Applies the list of mutate operations such as adding or updating
	// conversion trackers.         <p class="note"><b>Note:</b> {@link
	// ConversionTrackerOperation} does not support the         <code>REMOVE</code>
	// operator. In order to 'disable' a conversion type, send a
	//       <code>SET</code> operation for the conversion tracker
	// with the <code>status</code>         property set to <code>DISABLED</code></p>
	//                  <p>You can mutate any ConversionTracker that
	// belongs to your account. You may not         mutate a ConversionTracker
	// that belongs to some other account. You may not directly
	//      mutate a system-defined ConversionTracker, but you can
	// create a mutable copy of it         in your account by sending
	// a mutate request with an ADD operation specifying         an
	// originalConversionTypeId matching a system-defined conversion
	// tracker's ID. That new         ADDed ConversionTracker will
	// inherit the statistics and properties         of the system-defined
	// type, but will be editable as usual.</p>                  @param
	// operations A list of mutate operations to perform.         @return
	// The list of the conversion trackers as they appear after mutation,
	//         in the same order as they appeared in the list of operations.
	//         @throws com.google.ads.api.services.common.error.ApiException
	// if problems         occurred while updating the data.
	Mutate(Mutate *Mutate) (*MutateResponse, error)

	// Returns a list of conversion trackers that match the query.
	//                  @param query The SQL-like AWQL query string.
	//         @return A list of conversion trackers.         @throws
	// ApiException if problems occur while parsing the query or fetching
	// conversion trackers.
	Query(Query string) (*QueryResponse, error)
}

// AdWordsConversionTrackerTrackingCodeType was auto-generated
// from WSDL.
type AdWordsConversionTrackerTrackingCodeType string

// Validate validates AdWordsConversionTrackerTrackingCodeType.
func (v AdWordsConversionTrackerTrackingCodeType) Validate() bool {
	for _, vv := range []string{
		"WEBPAGE",
		"WEBPAGE_ONCLICK",
		"CLICK_TO_CALL",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AppConversionAppConversionType was auto-generated from WSDL.
type AppConversionAppConversionType string

// Validate validates AppConversionAppConversionType.
func (v AppConversionAppConversionType) Validate() bool {
	for _, vv := range []string{
		"NONE",
		"DOWNLOAD",
		"IN_APP_PURCHASE",
		"FIRST_OPEN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AppConversionAppPlatform was auto-generated from WSDL.
type AppConversionAppPlatform string

// Validate validates AppConversionAppPlatform.
func (v AppConversionAppPlatform) Validate() bool {
	for _, vv := range []string{
		"NONE",
		"ITUNES",
		"ANDROID_MARKET",
		"MOBILE_APP_CHANNEL",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AppPostbackUrlErrorReason was auto-generated from WSDL.
type AppPostbackUrlErrorReason string

// Validate validates AppPostbackUrlErrorReason.
func (v AppPostbackUrlErrorReason) Validate() bool {
	for _, vv := range []string{
		"INVALID_URL_FORMAT",
		"INVALID_DOMAIN",
		"REQUIRED_MACRO_NOT_FOUND",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AttributionModelType was auto-generated from WSDL.
type AttributionModelType string

// Validate validates AttributionModelType.
func (v AttributionModelType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"LAST_CLICK",
		"FIRST_CLICK",
		"LINEAR",
		"TIME_DECAY",
		"U_SHAPED",
		"DATA_DRIVEN",
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

// ConversionDeduplicationMode was auto-generated from WSDL.
type ConversionDeduplicationMode string

// Validate validates ConversionDeduplicationMode.
func (v ConversionDeduplicationMode) Validate() bool {
	for _, vv := range []string{
		"ONE_PER_CLICK",
		"MANY_PER_CLICK",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ConversionTrackerCategory was auto-generated from WSDL.
type ConversionTrackerCategory string

// Validate validates ConversionTrackerCategory.
func (v ConversionTrackerCategory) Validate() bool {
	for _, vv := range []string{
		"DEFAULT",
		"PAGE_VIEW",
		"PURCHASE",
		"SIGNUP",
		"LEAD",
		"REMARKETING",
		"DOWNLOAD",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ConversionTrackerStatus was auto-generated from WSDL.
type ConversionTrackerStatus string

// Validate validates ConversionTrackerStatus.
func (v ConversionTrackerStatus) Validate() bool {
	for _, vv := range []string{
		"ENABLED",
		"DISABLED",
		"HIDDEN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ConversionTrackingErrorReason was auto-generated from WSDL.
type ConversionTrackingErrorReason string

// Validate validates ConversionTrackingErrorReason.
func (v ConversionTrackingErrorReason) Validate() bool {
	for _, vv := range []string{
		"ALREADY_CREATED_CUSTOM_CONVERSION_TYPE",
		"ANALYTICS_NOT_ALLOWED",
		"CANNOT_ADD_CONVERSION_TYPE_SUBCLASS",
		"CANNOT_ADD_EXTERNALLY_ATTRIBUTED_SALESFORCE_CONVERSION",
		"CANNOT_CHANGE_APP_CONVERSION_TYPE",
		"CANNOT_CHANGE_APP_PLATFORM",
		"CANNNOT_CHANGE_CONVERSION_SUBCLASS",
		"CANNOT_SET_HIDDEN_STATUS",
		"CATEGORY_IS_UNEDITABLE",
		"ATTRIBUTION_MODEL_IS_UNEDITABLE",
		"ATTRIBUTION_MODEL_CANNOT_BE_UNKNOWN",
		"DATA_DRIVEN_MODEL_WAS_NEVER_GENERATED",
		"DATA_DRIVEN_MODEL_IS_EXPIRED",
		"DATA_DRIVEN_MODEL_IS_STALE",
		"DATA_DRIVEN_MODEL_IS_UNKNOWN",
		"CONVERSION_TYPE_NOT_FOUND",
		"CTC_LOOKBACK_WINDOW_IS_UNEDITABLE",
		"DOMAIN_EXCEPTION",
		"INCONSISTENT_COUNTING_TYPE",
		"DUPLICATE_APP_ID",
		"TWO_CONVERSION_TYPES_BIDDING_ON_SAME_APP_DOWNLOAD",
		"CONVERSION_TYPE_BIDDING_ON_SAME_APP_DOWNLOAD_AS_GLOBAL_TYPE",
		"DUPLICATE_NAME",
		"EMAIL_FAILED",
		"EXCEEDED_CONVERSION_TYPE_LIMIT",
		"ID_IS_NULL",
		"INVALID_APP_ID",
		"CANNOT_SET_APP_ID",
		"INVALID_CATEGORY",
		"INVALID_COLOR",
		"INVALID_DATE_RANGE",
		"INVALID_EMAIL_ADDRESS",
		"INVALID_ORIGINAL_CONVERSION_TYPE_ID",
		"MUST_SET_APP_PLATFORM_AND_APP_CONVERSION_TYPE_TOGETHER",
		"NAME_ALREADY_EXISTS",
		"NO_RECIPIENTS",
		"NO_SNIPPET",
		"TOO_MANY_WEBPAGES",
		"UNKNOWN_SORTING_TYPE",
		"UNSUPPORTED_APP_CONVERSION_TYPE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// DataDrivenModelStatus was auto-generated from WSDL.
type DataDrivenModelStatus string

// Validate validates DataDrivenModelStatus.
func (v DataDrivenModelStatus) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"AVAILABLE",
		"STALE",
		"EXPIRED",
		"NEVER_GENERATED",
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

// A ConversionTracker for phone calls from conversion-tracked
// call extensions and             call-only ads.
//             <p>A call made from the call extension is reported
// as a conversion if it lasts longer             than N seconds.
// This duration is 60 seconds by default. Each call extension
// can             specify the desired conversion configuration.</p>
type AdCallMetricsConversion struct {
	Id                            *int64                       `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	OriginalConversionTypeId      *int64                       `xml:"originalConversionTypeId,omitempty" json:"originalConversionTypeId,omitempty" yaml:"originalConversionTypeId,omitempty"`
	Name                          *string                      `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Status                        *ConversionTrackerStatus     `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Category                      *ConversionTrackerCategory   `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
	GoogleEventSnippet            *string                      `xml:"googleEventSnippet,omitempty" json:"googleEventSnippet,omitempty" yaml:"googleEventSnippet,omitempty"`
	GoogleGlobalSiteTag           *string                      `xml:"googleGlobalSiteTag,omitempty" json:"googleGlobalSiteTag,omitempty" yaml:"googleGlobalSiteTag,omitempty"`
	DataDrivenModelStatus         *DataDrivenModelStatus       `xml:"dataDrivenModelStatus,omitempty" json:"dataDrivenModelStatus,omitempty" yaml:"dataDrivenModelStatus,omitempty"`
	ConversionTypeOwnerCustomerId *int64                       `xml:"conversionTypeOwnerCustomerId,omitempty" json:"conversionTypeOwnerCustomerId,omitempty" yaml:"conversionTypeOwnerCustomerId,omitempty"`
	ViewthroughLookbackWindow     *int                         `xml:"viewthroughLookbackWindow,omitempty" json:"viewthroughLookbackWindow,omitempty" yaml:"viewthroughLookbackWindow,omitempty"`
	CtcLookbackWindow             *int                         `xml:"ctcLookbackWindow,omitempty" json:"ctcLookbackWindow,omitempty" yaml:"ctcLookbackWindow,omitempty"`
	CountingType                  *ConversionDeduplicationMode `xml:"countingType,omitempty" json:"countingType,omitempty" yaml:"countingType,omitempty"`
	DefaultRevenueValue           *float64                     `xml:"defaultRevenueValue,omitempty" json:"defaultRevenueValue,omitempty" yaml:"defaultRevenueValue,omitempty"`
	DefaultRevenueCurrencyCode    *string                      `xml:"defaultRevenueCurrencyCode,omitempty" json:"defaultRevenueCurrencyCode,omitempty" yaml:"defaultRevenueCurrencyCode,omitempty"`
	AlwaysUseDefaultRevenueValue  *bool                        `xml:"alwaysUseDefaultRevenueValue,omitempty" json:"alwaysUseDefaultRevenueValue,omitempty" yaml:"alwaysUseDefaultRevenueValue,omitempty"`
	ExcludeFromBidding            *bool                        `xml:"excludeFromBidding,omitempty" json:"excludeFromBidding,omitempty" yaml:"excludeFromBidding,omitempty"`
	AttributionModelType          *AttributionModelType        `xml:"attributionModelType,omitempty" json:"attributionModelType,omitempty" yaml:"attributionModelType,omitempty"`
	MostRecentConversionDate      *string                      `xml:"mostRecentConversionDate,omitempty" json:"mostRecentConversionDate,omitempty" yaml:"mostRecentConversionDate,omitempty"`
	LastReceivedRequestTime       *string                      `xml:"lastReceivedRequestTime,omitempty" json:"lastReceivedRequestTime,omitempty" yaml:"lastReceivedRequestTime,omitempty"`
	ConversionTrackerType         *string                      `xml:"ConversionTracker.Type,omitempty" json:"ConversionTracker.Type,omitempty" yaml:"ConversionTracker.Type,omitempty"`
	PhoneCallDuration             *int64                       `xml:"phoneCallDuration,omitempty" json:"phoneCallDuration,omitempty" yaml:"phoneCallDuration,omitempty"`
	TypeAttrXSI                   string                       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                 string                       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdCallMetricsConversion) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdCallMetricsConversion"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A conversion tracker created through AdWords Conversion Tracking.
type AdWordsConversionTracker struct {
	Id                            *int64                                    `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	OriginalConversionTypeId      *int64                                    `xml:"originalConversionTypeId,omitempty" json:"originalConversionTypeId,omitempty" yaml:"originalConversionTypeId,omitempty"`
	Name                          *string                                   `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Status                        *ConversionTrackerStatus                  `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Category                      *ConversionTrackerCategory                `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
	GoogleEventSnippet            *string                                   `xml:"googleEventSnippet,omitempty" json:"googleEventSnippet,omitempty" yaml:"googleEventSnippet,omitempty"`
	GoogleGlobalSiteTag           *string                                   `xml:"googleGlobalSiteTag,omitempty" json:"googleGlobalSiteTag,omitempty" yaml:"googleGlobalSiteTag,omitempty"`
	DataDrivenModelStatus         *DataDrivenModelStatus                    `xml:"dataDrivenModelStatus,omitempty" json:"dataDrivenModelStatus,omitempty" yaml:"dataDrivenModelStatus,omitempty"`
	ConversionTypeOwnerCustomerId *int64                                    `xml:"conversionTypeOwnerCustomerId,omitempty" json:"conversionTypeOwnerCustomerId,omitempty" yaml:"conversionTypeOwnerCustomerId,omitempty"`
	ViewthroughLookbackWindow     *int                                      `xml:"viewthroughLookbackWindow,omitempty" json:"viewthroughLookbackWindow,omitempty" yaml:"viewthroughLookbackWindow,omitempty"`
	CtcLookbackWindow             *int                                      `xml:"ctcLookbackWindow,omitempty" json:"ctcLookbackWindow,omitempty" yaml:"ctcLookbackWindow,omitempty"`
	CountingType                  *ConversionDeduplicationMode              `xml:"countingType,omitempty" json:"countingType,omitempty" yaml:"countingType,omitempty"`
	DefaultRevenueValue           *float64                                  `xml:"defaultRevenueValue,omitempty" json:"defaultRevenueValue,omitempty" yaml:"defaultRevenueValue,omitempty"`
	DefaultRevenueCurrencyCode    *string                                   `xml:"defaultRevenueCurrencyCode,omitempty" json:"defaultRevenueCurrencyCode,omitempty" yaml:"defaultRevenueCurrencyCode,omitempty"`
	AlwaysUseDefaultRevenueValue  *bool                                     `xml:"alwaysUseDefaultRevenueValue,omitempty" json:"alwaysUseDefaultRevenueValue,omitempty" yaml:"alwaysUseDefaultRevenueValue,omitempty"`
	ExcludeFromBidding            *bool                                     `xml:"excludeFromBidding,omitempty" json:"excludeFromBidding,omitempty" yaml:"excludeFromBidding,omitempty"`
	AttributionModelType          *AttributionModelType                     `xml:"attributionModelType,omitempty" json:"attributionModelType,omitempty" yaml:"attributionModelType,omitempty"`
	MostRecentConversionDate      *string                                   `xml:"mostRecentConversionDate,omitempty" json:"mostRecentConversionDate,omitempty" yaml:"mostRecentConversionDate,omitempty"`
	LastReceivedRequestTime       *string                                   `xml:"lastReceivedRequestTime,omitempty" json:"lastReceivedRequestTime,omitempty" yaml:"lastReceivedRequestTime,omitempty"`
	ConversionTrackerType         *string                                   `xml:"ConversionTracker.Type,omitempty" json:"ConversionTracker.Type,omitempty" yaml:"ConversionTracker.Type,omitempty"`
	TrackingCodeType              *AdWordsConversionTrackerTrackingCodeType `xml:"trackingCodeType,omitempty" json:"trackingCodeType,omitempty" yaml:"trackingCodeType,omitempty"`
	TypeAttrXSI                   string                                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                 string                                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdWordsConversionTracker) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdWordsConversionTracker"
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

// A ConversionTracker for mobile apps.
type AppConversion struct {
	Id                            *int64                          `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	OriginalConversionTypeId      *int64                          `xml:"originalConversionTypeId,omitempty" json:"originalConversionTypeId,omitempty" yaml:"originalConversionTypeId,omitempty"`
	Name                          *string                         `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Status                        *ConversionTrackerStatus        `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Category                      *ConversionTrackerCategory      `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
	GoogleEventSnippet            *string                         `xml:"googleEventSnippet,omitempty" json:"googleEventSnippet,omitempty" yaml:"googleEventSnippet,omitempty"`
	GoogleGlobalSiteTag           *string                         `xml:"googleGlobalSiteTag,omitempty" json:"googleGlobalSiteTag,omitempty" yaml:"googleGlobalSiteTag,omitempty"`
	DataDrivenModelStatus         *DataDrivenModelStatus          `xml:"dataDrivenModelStatus,omitempty" json:"dataDrivenModelStatus,omitempty" yaml:"dataDrivenModelStatus,omitempty"`
	ConversionTypeOwnerCustomerId *int64                          `xml:"conversionTypeOwnerCustomerId,omitempty" json:"conversionTypeOwnerCustomerId,omitempty" yaml:"conversionTypeOwnerCustomerId,omitempty"`
	ViewthroughLookbackWindow     *int                            `xml:"viewthroughLookbackWindow,omitempty" json:"viewthroughLookbackWindow,omitempty" yaml:"viewthroughLookbackWindow,omitempty"`
	CtcLookbackWindow             *int                            `xml:"ctcLookbackWindow,omitempty" json:"ctcLookbackWindow,omitempty" yaml:"ctcLookbackWindow,omitempty"`
	CountingType                  *ConversionDeduplicationMode    `xml:"countingType,omitempty" json:"countingType,omitempty" yaml:"countingType,omitempty"`
	DefaultRevenueValue           *float64                        `xml:"defaultRevenueValue,omitempty" json:"defaultRevenueValue,omitempty" yaml:"defaultRevenueValue,omitempty"`
	DefaultRevenueCurrencyCode    *string                         `xml:"defaultRevenueCurrencyCode,omitempty" json:"defaultRevenueCurrencyCode,omitempty" yaml:"defaultRevenueCurrencyCode,omitempty"`
	AlwaysUseDefaultRevenueValue  *bool                           `xml:"alwaysUseDefaultRevenueValue,omitempty" json:"alwaysUseDefaultRevenueValue,omitempty" yaml:"alwaysUseDefaultRevenueValue,omitempty"`
	ExcludeFromBidding            *bool                           `xml:"excludeFromBidding,omitempty" json:"excludeFromBidding,omitempty" yaml:"excludeFromBidding,omitempty"`
	AttributionModelType          *AttributionModelType           `xml:"attributionModelType,omitempty" json:"attributionModelType,omitempty" yaml:"attributionModelType,omitempty"`
	MostRecentConversionDate      *string                         `xml:"mostRecentConversionDate,omitempty" json:"mostRecentConversionDate,omitempty" yaml:"mostRecentConversionDate,omitempty"`
	LastReceivedRequestTime       *string                         `xml:"lastReceivedRequestTime,omitempty" json:"lastReceivedRequestTime,omitempty" yaml:"lastReceivedRequestTime,omitempty"`
	ConversionTrackerType         *string                         `xml:"ConversionTracker.Type,omitempty" json:"ConversionTracker.Type,omitempty" yaml:"ConversionTracker.Type,omitempty"`
	AppId                         *string                         `xml:"appId,omitempty" json:"appId,omitempty" yaml:"appId,omitempty"`
	AppPlatform                   *AppConversionAppPlatform       `xml:"appPlatform,omitempty" json:"appPlatform,omitempty" yaml:"appPlatform,omitempty"`
	Snippet                       *string                         `xml:"snippet,omitempty" json:"snippet,omitempty" yaml:"snippet,omitempty"`
	AppConversionType             *AppConversionAppConversionType `xml:"appConversionType,omitempty" json:"appConversionType,omitempty" yaml:"appConversionType,omitempty"`
	AppPostbackUrl                *string                         `xml:"appPostbackUrl,omitempty" json:"appPostbackUrl,omitempty" yaml:"appPostbackUrl,omitempty"`
	TypeAttrXSI                   string                          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                 string                          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AppConversion) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AppConversion"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Errors returned when App Postback Url is invalid.
type AppPostbackUrlError struct {
	FieldPath         *string                    `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement        `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                    `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                    `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                    `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *AppPostbackUrlErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AppPostbackUrlError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AppPostbackUrlError"
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

// An abstract Conversion base class.
type ConversionTracker interface{}

// Operations for conversion tracker.
type ConversionTrackerOperation struct {
	Operator      *Operator          `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	OperationType *string            `xml:"Operation.Type,omitempty" json:"Operation.Type,omitempty" yaml:"Operation.Type,omitempty"`
	Operand       *ConversionTracker `xml:"operand,omitempty" json:"operand,omitempty" yaml:"operand,omitempty"`
	TypeAttrXSI   string             `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string             `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ConversionTrackerOperation) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ConversionTrackerOperation"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Contains a subset of {@link ConversionTracker}s resulting from
// the filtering             and paging of the {@link ConversionTrackerService#get}
// call.
type ConversionTrackerPage struct {
	TotalNumEntries *int                 `xml:"totalNumEntries,omitempty" json:"totalNumEntries,omitempty" yaml:"totalNumEntries,omitempty"`
	PageType        *string              `xml:"Page.Type,omitempty" json:"Page.Type,omitempty" yaml:"Page.Type,omitempty"`
	Entries         []*ConversionTracker `xml:"entries,omitempty" json:"entries,omitempty" yaml:"entries,omitempty"`
	TypeAttrXSI     string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ConversionTrackerPage) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ConversionTrackerPage"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A container for return values from the ConversionTrackerService.
type ConversionTrackerReturnValue struct {
	ListReturnValueType *string              `xml:"ListReturnValue.Type,omitempty" json:"ListReturnValue.Type,omitempty" yaml:"ListReturnValue.Type,omitempty"`
	Value               []*ConversionTracker `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI         string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ConversionTrackerReturnValue) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ConversionTrackerReturnValue"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// An error that can occur during calls to the ConversionTypeService.
type ConversionTrackingError struct {
	FieldPath         *string                        `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement            `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                        `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                        `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                        `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *ConversionTrackingErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ConversionTrackingError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ConversionTrackingError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
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

// Base list return value type.
type ListReturnValue interface{}

// NoStatsPage was auto-generated from WSDL.
type NoStatsPage interface{}

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

// Specifies how an entity (eg. adgroup, campaign, criterion, ad)
// should be filtered.
type Predicate struct {
	Field    *string            `xml:"field,omitempty" json:"field,omitempty" yaml:"field,omitempty"`
	Operator *PredicateOperator `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	Values   []*string          `xml:"values,omitempty" json:"values,omitempty" yaml:"values,omitempty"`
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

// A conversion that happens when a user performs one the following
// sequences of actions:             <ul>             <li>User
// clicks on an advertiser's ad which takes the user to the advertiser's
// website, where             special javascript installed on the
// page produces a dynamically-generated phone number.
//     Then, user calls that number from their home (or other)
// phone</li>             </li>User makes a phone call from conversion-tracked
// call extensions </li>             </ul>
//      After successfully creating a new UploadCallConversion,
// send the name of this conversion type             along with
// your conversion details to the OfflineCallConversionFeedService
//             to attribute those conversions to this conversion
// type.
type UploadCallConversion struct {
	Id                            *int64                       `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	OriginalConversionTypeId      *int64                       `xml:"originalConversionTypeId,omitempty" json:"originalConversionTypeId,omitempty" yaml:"originalConversionTypeId,omitempty"`
	Name                          *string                      `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Status                        *ConversionTrackerStatus     `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Category                      *ConversionTrackerCategory   `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
	GoogleEventSnippet            *string                      `xml:"googleEventSnippet,omitempty" json:"googleEventSnippet,omitempty" yaml:"googleEventSnippet,omitempty"`
	GoogleGlobalSiteTag           *string                      `xml:"googleGlobalSiteTag,omitempty" json:"googleGlobalSiteTag,omitempty" yaml:"googleGlobalSiteTag,omitempty"`
	DataDrivenModelStatus         *DataDrivenModelStatus       `xml:"dataDrivenModelStatus,omitempty" json:"dataDrivenModelStatus,omitempty" yaml:"dataDrivenModelStatus,omitempty"`
	ConversionTypeOwnerCustomerId *int64                       `xml:"conversionTypeOwnerCustomerId,omitempty" json:"conversionTypeOwnerCustomerId,omitempty" yaml:"conversionTypeOwnerCustomerId,omitempty"`
	ViewthroughLookbackWindow     *int                         `xml:"viewthroughLookbackWindow,omitempty" json:"viewthroughLookbackWindow,omitempty" yaml:"viewthroughLookbackWindow,omitempty"`
	CtcLookbackWindow             *int                         `xml:"ctcLookbackWindow,omitempty" json:"ctcLookbackWindow,omitempty" yaml:"ctcLookbackWindow,omitempty"`
	CountingType                  *ConversionDeduplicationMode `xml:"countingType,omitempty" json:"countingType,omitempty" yaml:"countingType,omitempty"`
	DefaultRevenueValue           *float64                     `xml:"defaultRevenueValue,omitempty" json:"defaultRevenueValue,omitempty" yaml:"defaultRevenueValue,omitempty"`
	DefaultRevenueCurrencyCode    *string                      `xml:"defaultRevenueCurrencyCode,omitempty" json:"defaultRevenueCurrencyCode,omitempty" yaml:"defaultRevenueCurrencyCode,omitempty"`
	AlwaysUseDefaultRevenueValue  *bool                        `xml:"alwaysUseDefaultRevenueValue,omitempty" json:"alwaysUseDefaultRevenueValue,omitempty" yaml:"alwaysUseDefaultRevenueValue,omitempty"`
	ExcludeFromBidding            *bool                        `xml:"excludeFromBidding,omitempty" json:"excludeFromBidding,omitempty" yaml:"excludeFromBidding,omitempty"`
	AttributionModelType          *AttributionModelType        `xml:"attributionModelType,omitempty" json:"attributionModelType,omitempty" yaml:"attributionModelType,omitempty"`
	MostRecentConversionDate      *string                      `xml:"mostRecentConversionDate,omitempty" json:"mostRecentConversionDate,omitempty" yaml:"mostRecentConversionDate,omitempty"`
	LastReceivedRequestTime       *string                      `xml:"lastReceivedRequestTime,omitempty" json:"lastReceivedRequestTime,omitempty" yaml:"lastReceivedRequestTime,omitempty"`
	ConversionTrackerType         *string                      `xml:"ConversionTracker.Type,omitempty" json:"ConversionTracker.Type,omitempty" yaml:"ConversionTracker.Type,omitempty"`
	TypeAttrXSI                   string                       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                 string                       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *UploadCallConversion) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UploadCallConversion"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A conversion type that receives conversions by having them uploaded
//             through the OfflineConversionFeedService.
//                    After successfully creating a new UploadConversion,
// send the name of this conversion type             along with
// your conversion details to the OfflineConversionFeedService
//             to attribute those conversions to this conversion
// type.
type UploadConversion struct {
	Id                            *int64                       `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	OriginalConversionTypeId      *int64                       `xml:"originalConversionTypeId,omitempty" json:"originalConversionTypeId,omitempty" yaml:"originalConversionTypeId,omitempty"`
	Name                          *string                      `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Status                        *ConversionTrackerStatus     `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Category                      *ConversionTrackerCategory   `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
	GoogleEventSnippet            *string                      `xml:"googleEventSnippet,omitempty" json:"googleEventSnippet,omitempty" yaml:"googleEventSnippet,omitempty"`
	GoogleGlobalSiteTag           *string                      `xml:"googleGlobalSiteTag,omitempty" json:"googleGlobalSiteTag,omitempty" yaml:"googleGlobalSiteTag,omitempty"`
	DataDrivenModelStatus         *DataDrivenModelStatus       `xml:"dataDrivenModelStatus,omitempty" json:"dataDrivenModelStatus,omitempty" yaml:"dataDrivenModelStatus,omitempty"`
	ConversionTypeOwnerCustomerId *int64                       `xml:"conversionTypeOwnerCustomerId,omitempty" json:"conversionTypeOwnerCustomerId,omitempty" yaml:"conversionTypeOwnerCustomerId,omitempty"`
	ViewthroughLookbackWindow     *int                         `xml:"viewthroughLookbackWindow,omitempty" json:"viewthroughLookbackWindow,omitempty" yaml:"viewthroughLookbackWindow,omitempty"`
	CtcLookbackWindow             *int                         `xml:"ctcLookbackWindow,omitempty" json:"ctcLookbackWindow,omitempty" yaml:"ctcLookbackWindow,omitempty"`
	CountingType                  *ConversionDeduplicationMode `xml:"countingType,omitempty" json:"countingType,omitempty" yaml:"countingType,omitempty"`
	DefaultRevenueValue           *float64                     `xml:"defaultRevenueValue,omitempty" json:"defaultRevenueValue,omitempty" yaml:"defaultRevenueValue,omitempty"`
	DefaultRevenueCurrencyCode    *string                      `xml:"defaultRevenueCurrencyCode,omitempty" json:"defaultRevenueCurrencyCode,omitempty" yaml:"defaultRevenueCurrencyCode,omitempty"`
	AlwaysUseDefaultRevenueValue  *bool                        `xml:"alwaysUseDefaultRevenueValue,omitempty" json:"alwaysUseDefaultRevenueValue,omitempty" yaml:"alwaysUseDefaultRevenueValue,omitempty"`
	ExcludeFromBidding            *bool                        `xml:"excludeFromBidding,omitempty" json:"excludeFromBidding,omitempty" yaml:"excludeFromBidding,omitempty"`
	AttributionModelType          *AttributionModelType        `xml:"attributionModelType,omitempty" json:"attributionModelType,omitempty" yaml:"attributionModelType,omitempty"`
	MostRecentConversionDate      *string                      `xml:"mostRecentConversionDate,omitempty" json:"mostRecentConversionDate,omitempty" yaml:"mostRecentConversionDate,omitempty"`
	LastReceivedRequestTime       *string                      `xml:"lastReceivedRequestTime,omitempty" json:"lastReceivedRequestTime,omitempty" yaml:"lastReceivedRequestTime,omitempty"`
	ConversionTrackerType         *string                      `xml:"ConversionTracker.Type,omitempty" json:"ConversionTracker.Type,omitempty" yaml:"ConversionTracker.Type,omitempty"`
	IsExternallyAttributed        *bool                        `xml:"isExternallyAttributed,omitempty" json:"isExternallyAttributed,omitempty" yaml:"isExternallyAttributed,omitempty"`
	TypeAttrXSI                   string                       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                 string                       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *UploadConversion) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UploadConversion"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A conversion that happens when a user performs the following
// sequence of actions:             <ul>             <li>Clicks
// on an advertiser's ad</li>             <li>             Proceeds
// to the advertiser's website, where special javascript installed
// on the page             produces a dynamically-generated phone
// number,             </li>             <li>Calls that number
// from their home (or other) phone</li>             </ul>
type WebsiteCallMetricsConversion struct {
	Id                            *int64                       `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	OriginalConversionTypeId      *int64                       `xml:"originalConversionTypeId,omitempty" json:"originalConversionTypeId,omitempty" yaml:"originalConversionTypeId,omitempty"`
	Name                          *string                      `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Status                        *ConversionTrackerStatus     `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Category                      *ConversionTrackerCategory   `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
	GoogleEventSnippet            *string                      `xml:"googleEventSnippet,omitempty" json:"googleEventSnippet,omitempty" yaml:"googleEventSnippet,omitempty"`
	GoogleGlobalSiteTag           *string                      `xml:"googleGlobalSiteTag,omitempty" json:"googleGlobalSiteTag,omitempty" yaml:"googleGlobalSiteTag,omitempty"`
	DataDrivenModelStatus         *DataDrivenModelStatus       `xml:"dataDrivenModelStatus,omitempty" json:"dataDrivenModelStatus,omitempty" yaml:"dataDrivenModelStatus,omitempty"`
	ConversionTypeOwnerCustomerId *int64                       `xml:"conversionTypeOwnerCustomerId,omitempty" json:"conversionTypeOwnerCustomerId,omitempty" yaml:"conversionTypeOwnerCustomerId,omitempty"`
	ViewthroughLookbackWindow     *int                         `xml:"viewthroughLookbackWindow,omitempty" json:"viewthroughLookbackWindow,omitempty" yaml:"viewthroughLookbackWindow,omitempty"`
	CtcLookbackWindow             *int                         `xml:"ctcLookbackWindow,omitempty" json:"ctcLookbackWindow,omitempty" yaml:"ctcLookbackWindow,omitempty"`
	CountingType                  *ConversionDeduplicationMode `xml:"countingType,omitempty" json:"countingType,omitempty" yaml:"countingType,omitempty"`
	DefaultRevenueValue           *float64                     `xml:"defaultRevenueValue,omitempty" json:"defaultRevenueValue,omitempty" yaml:"defaultRevenueValue,omitempty"`
	DefaultRevenueCurrencyCode    *string                      `xml:"defaultRevenueCurrencyCode,omitempty" json:"defaultRevenueCurrencyCode,omitempty" yaml:"defaultRevenueCurrencyCode,omitempty"`
	AlwaysUseDefaultRevenueValue  *bool                        `xml:"alwaysUseDefaultRevenueValue,omitempty" json:"alwaysUseDefaultRevenueValue,omitempty" yaml:"alwaysUseDefaultRevenueValue,omitempty"`
	ExcludeFromBidding            *bool                        `xml:"excludeFromBidding,omitempty" json:"excludeFromBidding,omitempty" yaml:"excludeFromBidding,omitempty"`
	AttributionModelType          *AttributionModelType        `xml:"attributionModelType,omitempty" json:"attributionModelType,omitempty" yaml:"attributionModelType,omitempty"`
	MostRecentConversionDate      *string                      `xml:"mostRecentConversionDate,omitempty" json:"mostRecentConversionDate,omitempty" yaml:"mostRecentConversionDate,omitempty"`
	LastReceivedRequestTime       *string                      `xml:"lastReceivedRequestTime,omitempty" json:"lastReceivedRequestTime,omitempty" yaml:"lastReceivedRequestTime,omitempty"`
	ConversionTrackerType         *string                      `xml:"ConversionTracker.Type,omitempty" json:"ConversionTracker.Type,omitempty" yaml:"ConversionTracker.Type,omitempty"`
	PhoneCallDuration             *int64                       `xml:"phoneCallDuration,omitempty" json:"phoneCallDuration,omitempty" yaml:"phoneCallDuration,omitempty"`
	TypeAttrXSI                   string                       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                 string                       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *WebsiteCallMetricsConversion) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:WebsiteCallMetricsConversion"
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
	Rval *ConversionTrackerPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Mutate was auto-generated from WSDL.
type Mutate struct {
	Operations []*ConversionTrackerOperation `xml:"operations,omitempty" json:"operations,omitempty" yaml:"operations,omitempty"`
}

// MutateResponse was auto-generated from WSDL.
type MutateResponse struct {
	Rval *ConversionTrackerReturnValue `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Query was auto-generated from WSDL.
type Query struct {
	Query *string `xml:"query,omitempty" json:"query,omitempty" yaml:"query,omitempty"`
}

// QueryResponse was auto-generated from WSDL.
type QueryResponse struct {
	Rval *ConversionTrackerPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
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

// conversionTrackerServiceInterface implements the ConversionTrackerServiceInterface interface.
type conversionTrackerServiceInterface struct {
	cli *soap.Client
}

// Returns a list of the conversion trackers that match the selector.
// The         actual objects contained in the page's list of entries
// will be specific         subclasses of the abstract {@link ConversionTracker}
// class.                  @param serviceSelector The selector
// specifying the         {@link ConversionTracker}s to return.
//         @return List of conversion trackers specified by the
// selector.         @throws com.google.ads.api.services.common.error.ApiException
// if problems         occurred while retrieving results.
func (p *conversionTrackerServiceInterface) Get(Get *Get) (*GetResponse, error) {
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

// Applies the list of mutate operations such as adding or updating
// conversion trackers.         <p class="note"><b>Note:</b> {@link
// ConversionTrackerOperation} does not support the         <code>REMOVE</code>
// operator. In order to 'disable' a conversion type, send a
//       <code>SET</code> operation for the conversion tracker
// with the <code>status</code>         property set to <code>DISABLED</code></p>
//                  <p>You can mutate any ConversionTracker that
// belongs to your account. You may not         mutate a ConversionTracker
// that belongs to some other account. You may not directly
//      mutate a system-defined ConversionTracker, but you can
// create a mutable copy of it         in your account by sending
// a mutate request with an ADD operation specifying         an
// originalConversionTypeId matching a system-defined conversion
// tracker's ID. That new         ADDed ConversionTracker will
// inherit the statistics and properties         of the system-defined
// type, but will be editable as usual.</p>                  @param
// operations A list of mutate operations to perform.         @return
// The list of the conversion trackers as they appear after mutation,
//         in the same order as they appeared in the list of operations.
//         @throws com.google.ads.api.services.common.error.ApiException
// if problems         occurred while updating the data.
func (p *conversionTrackerServiceInterface) Mutate(Mutate *Mutate) (*MutateResponse, error) {
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

// Returns a list of conversion trackers that match the query.
//                  @param query The SQL-like AWQL query string.
//         @return A list of conversion trackers.         @throws
// ApiException if problems occur while parsing the query or fetching
// conversion trackers.
func (p *conversionTrackerServiceInterface) Query(Query string) (*QueryResponse, error) {
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
