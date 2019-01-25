package reportdefinition

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "https://adwords.google.com/api/adwords/cm/v201809"

// NewReportDefinitionServiceInterface creates an initializes a ReportDefinitionServiceInterface.
func NewReportDefinitionServiceInterface(cli *soap.Client) ReportDefinitionServiceInterface {
	return &reportDefinitionServiceInterface{cli}
}

// ReportDefinitionServiceInterface was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ReportDefinitionServiceInterface interface {
	// Returns the available report fields for a given report type.
	//         When using this method the {@code clientCustomerId}
	// header field is         optional. Callers are discouraged from
	// setting the clientCustomerId         header field in calls to
	// this method as its presence will trigger an         authorization
	// error if the caller does not have access to the customer
	//      with the included ID.                  @param reportType
	// The type of report.         @return The list of available report
	// fields. Each         {@link ReportDefinitionField} encapsulates
	// the field name, the         field data type, and the enum values
	// (if the field's type is         {@code enum}).         @throws
	// ApiException if a problem occurred while fetching the
	//   ReportDefinitionField information.
	GetReportFields(GetReportFields *GetReportFields) (*GetReportFieldsResponse, error)
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

// NotWhitelistedErrorReason was auto-generated from WSDL.
type NotWhitelistedErrorReason string

// Validate validates NotWhitelistedErrorReason.
func (v NotWhitelistedErrorReason) Validate() bool {
	for _, vv := range []string{
		"CUSTOMER_NOT_WHITELISTED_FOR_API",
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

// ReportDefinitionReportType was auto-generated from WSDL.
type ReportDefinitionReportType string

// Validate validates ReportDefinitionReportType.
func (v ReportDefinitionReportType) Validate() bool {
	for _, vv := range []string{
		"KEYWORDS_PERFORMANCE_REPORT",
		"AD_PERFORMANCE_REPORT",
		"URL_PERFORMANCE_REPORT",
		"ADGROUP_PERFORMANCE_REPORT",
		"CAMPAIGN_PERFORMANCE_REPORT",
		"ACCOUNT_PERFORMANCE_REPORT",
		"GEO_PERFORMANCE_REPORT",
		"SEARCH_QUERY_PERFORMANCE_REPORT",
		"AUTOMATIC_PLACEMENTS_PERFORMANCE_REPORT",
		"CAMPAIGN_NEGATIVE_KEYWORDS_PERFORMANCE_REPORT",
		"CAMPAIGN_NEGATIVE_PLACEMENTS_PERFORMANCE_REPORT",
		"SHARED_SET_REPORT",
		"CAMPAIGN_SHARED_SET_REPORT",
		"SHARED_SET_CRITERIA_REPORT",
		"CREATIVE_CONVERSION_REPORT",
		"CALL_METRICS_CALL_DETAILS_REPORT",
		"KEYWORDLESS_QUERY_REPORT",
		"KEYWORDLESS_CATEGORY_REPORT",
		"CRITERIA_PERFORMANCE_REPORT",
		"CLICK_PERFORMANCE_REPORT",
		"BUDGET_PERFORMANCE_REPORT",
		"BID_GOAL_PERFORMANCE_REPORT",
		"DISPLAY_KEYWORD_PERFORMANCE_REPORT",
		"PLACEHOLDER_FEED_ITEM_REPORT",
		"PLACEMENT_PERFORMANCE_REPORT",
		"CAMPAIGN_NEGATIVE_LOCATIONS_REPORT",
		"GENDER_PERFORMANCE_REPORT",
		"AGE_RANGE_PERFORMANCE_REPORT",
		"CAMPAIGN_LOCATION_TARGET_REPORT",
		"CAMPAIGN_AD_SCHEDULE_TARGET_REPORT",
		"PAID_ORGANIC_QUERY_REPORT",
		"AUDIENCE_PERFORMANCE_REPORT",
		"DISPLAY_TOPICS_PERFORMANCE_REPORT",
		"USER_AD_DISTANCE_REPORT",
		"SHOPPING_PERFORMANCE_REPORT",
		"PRODUCT_PARTITION_REPORT",
		"PARENTAL_STATUS_PERFORMANCE_REPORT",
		"PLACEHOLDER_REPORT",
		"AD_CUSTOMIZERS_FEED_ITEM_REPORT",
		"LABEL_REPORT",
		"FINAL_URL_REPORT",
		"VIDEO_PERFORMANCE_REPORT",
		"TOP_CONTENT_PERFORMANCE_REPORT",
		"CAMPAIGN_CRITERIA_REPORT",
		"CAMPAIGN_GROUP_PERFORMANCE_REPORT",
		"LANDING_PAGE_REPORT",
		"MARKETPLACE_PERFORMANCE_REPORT",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ReportDefinitionErrorReason was auto-generated from WSDL.
type ReportDefinitionErrorReason string

// Validate validates ReportDefinitionErrorReason.
func (v ReportDefinitionErrorReason) Validate() bool {
	for _, vv := range []string{
		"INVALID_DATE_RANGE_FOR_REPORT",
		"INVALID_FIELD_NAME_FOR_REPORT",
		"UNABLE_TO_FIND_MAPPING_FOR_THIS_REPORT",
		"INVALID_COLUMN_NAME_FOR_REPORT",
		"INVALID_REPORT_DEFINITION_ID",
		"REPORT_SELECTOR_CANNOT_BE_NULL",
		"NO_ENUMS_FOR_THIS_COLUMN_NAME",
		"INVALID_VIEW",
		"SORTING_NOT_SUPPORTED",
		"PAGING_NOT_SUPPORTED",
		"CUSTOMER_SERVING_TYPE_REPORT_MISMATCH",
		"CLIENT_SELECTOR_NO_CUSTOMER_IDENTIFIER",
		"CLIENT_SELECTOR_INVALID_CUSTOMER_ID",
		"REPORT_DEFINITION_ERROR",
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

// Represents information about enum values.
type EnumValuePair struct {
	EnumValue        *string `xml:"enumValue,omitempty" json:"enumValue,omitempty" yaml:"enumValue,omitempty"`
	EnumDisplayValue *string `xml:"enumDisplayValue,omitempty" json:"enumDisplayValue,omitempty" yaml:"enumDisplayValue,omitempty"`
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

// Indicates that the customer is not whitelisted for accessing
// the API.
type NotWhitelistedError struct {
	FieldPath         *string                    `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement        `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                    `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                    `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                    `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *NotWhitelistedErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *NotWhitelistedError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:NotWhitelistedError"
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

// Encapsulates the errors that can be thrown during {@link ReportDefinition}
//             mutate operation.
type ReportDefinitionError struct {
	FieldPath         *string                      `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement          `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                      `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                      `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                      `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *ReportDefinitionErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                       `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                       `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ReportDefinitionError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ReportDefinitionError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Represents the fields that can be used to create a ReportDefinition.
//             This class allows the user to query the list of
// fields applicable to a             given report type. Consumers
// of reports will be able use the retrieved             fields
// through the {@link ReportDefinitionService#getReportFields}
//             api and run the corresponding reports.
type ReportDefinitionField struct {
	FieldName           *string          `xml:"fieldName,omitempty" json:"fieldName,omitempty" yaml:"fieldName,omitempty"`
	DisplayFieldName    *string          `xml:"displayFieldName,omitempty" json:"displayFieldName,omitempty" yaml:"displayFieldName,omitempty"`
	XmlAttributeName    *string          `xml:"xmlAttributeName,omitempty" json:"xmlAttributeName,omitempty" yaml:"xmlAttributeName,omitempty"`
	FieldType           *string          `xml:"fieldType,omitempty" json:"fieldType,omitempty" yaml:"fieldType,omitempty"`
	FieldBehavior       *string          `xml:"fieldBehavior,omitempty" json:"fieldBehavior,omitempty" yaml:"fieldBehavior,omitempty"`
	EnumValues          []*string        `xml:"enumValues,omitempty" json:"enumValues,omitempty" yaml:"enumValues,omitempty"`
	CanSelect           *bool            `xml:"canSelect,omitempty" json:"canSelect,omitempty" yaml:"canSelect,omitempty"`
	CanFilter           *bool            `xml:"canFilter,omitempty" json:"canFilter,omitempty" yaml:"canFilter,omitempty"`
	IsEnumType          *bool            `xml:"isEnumType,omitempty" json:"isEnumType,omitempty" yaml:"isEnumType,omitempty"`
	IsBeta              *bool            `xml:"isBeta,omitempty" json:"isBeta,omitempty" yaml:"isBeta,omitempty"`
	IsZeroRowCompatible *bool            `xml:"isZeroRowCompatible,omitempty" json:"isZeroRowCompatible,omitempty" yaml:"isZeroRowCompatible,omitempty"`
	EnumValuePairs      []*EnumValuePair `xml:"enumValuePairs,omitempty" json:"enumValuePairs,omitempty" yaml:"enumValuePairs,omitempty"`
	ExclusiveFields     []*string        `xml:"exclusiveFields,omitempty" json:"exclusiveFields,omitempty" yaml:"exclusiveFields,omitempty"`
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

// GetReportFields was auto-generated from WSDL.
type GetReportFields struct {
	ReportType *ReportDefinitionReportType `xml:"reportType,omitempty" json:"reportType,omitempty" yaml:"reportType,omitempty"`
}

// GetReportFieldsResponse was auto-generated from WSDL.
type GetReportFieldsResponse struct {
	Rval []*ReportDefinitionField `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Operation wrapper for GetReportFields.
// OperationGetReportFieldsRequest was auto-generated from WSDL.
type OperationGetReportFieldsRequest struct {
	GetReportFields *GetReportFields `xml:"getReportFields,omitempty" json:"getReportFields,omitempty" yaml:"getReportFields,omitempty"`
}

// Operation wrapper for GetReportFields.
// OperationGetReportFieldsResponse was auto-generated from WSDL.
type OperationGetReportFieldsResponse struct {
	GetReportFieldsResponse *GetReportFieldsResponse `xml:"getReportFieldsResponse,omitempty" json:"getReportFieldsResponse,omitempty" yaml:"getReportFieldsResponse,omitempty"`
}

// reportDefinitionServiceInterface implements the ReportDefinitionServiceInterface interface.
type reportDefinitionServiceInterface struct {
	cli *soap.Client
}

// Returns the available report fields for a given report type.
//         When using this method the {@code clientCustomerId}
// header field is         optional. Callers are discouraged from
// setting the clientCustomerId         header field in calls to
// this method as its presence will trigger an         authorization
// error if the caller does not have access to the customer
//      with the included ID.                  @param reportType
// The type of report.         @return The list of available report
// fields. Each         {@link ReportDefinitionField} encapsulates
// the field name, the         field data type, and the enum values
// (if the field's type is         {@code enum}).         @throws
// ApiException if a problem occurred while fetching the
//   ReportDefinitionField information.
func (p *reportDefinitionServiceInterface) GetReportFields(GetReportFields *GetReportFields) (*GetReportFieldsResponse, error) {
	α := struct {
		OperationGetReportFieldsRequest `xml:"tns:getReportFields"`
	}{
		OperationGetReportFieldsRequest{
			GetReportFields,
		},
	}

	γ := struct {
		OperationGetReportFieldsResponse `xml:"getReportFieldsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetReportFields", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetReportFieldsResponse, nil
}
