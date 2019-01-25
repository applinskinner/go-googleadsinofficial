package data

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "https://adwords.google.com/api/adwords/cm/v201809"

// NewDataServiceInterface creates an initializes a DataServiceInterface.
func NewDataServiceInterface(cli *soap.Client) DataServiceInterface {
	return &dataServiceInterface{cli}
}

// DataServiceInterface was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type DataServiceInterface interface {
	// Returns a list of {@link AdGroupBidLandscape}s for the ad groups
	// specified in the selector.         In the result, the returned
	// {@link LandscapePoint}s are grouped into         {@link AdGroupBidLandscape}s
	// by their ad groups, and numberResults of paging limits the total
	//         number of {@link LandscapePoint}s instead of number
	// of {@link AdGroupBidLandscape}s.                  @param serviceSelector
	// Selects the entities to return bid landscapes for.         @return
	// A list of bid landscapes.         @throws ApiException when
	// there is at least one error with the request.
	GetAdGroupBidLandscape(GetAdGroupBidLandscape *GetAdGroupBidLandscape) (*GetAdGroupBidLandscapeResponse, error)

	// Returns a list of {@link CriterionBidLandscape}s for the campaign
	// criteria specified in the         selector. In the result, the
	// returned {@link LandscapePoint}s are grouped into         {@link
	// CriterionBidLandscape}s by their campaign id and criterion id,
	// and numberResults         of paging limits the total number
	// of {@link LandscapePoint}s instead of number of         {@link
	// CriterionBidLandscape}s.                  @param serviceSelector
	// Selects the entities to return bid landscapes for.         @return
	// A list of bid landscapes.         @throws ApiException when
	// there is at least one error with the request.
	GetCampaignCriterionBidLandscape(GetCampaignCriterionBidLandscape *GetCampaignCriterionBidLandscape) (*GetCampaignCriterionBidLandscapeResponse, error)

	// Returns a list of {@link CriterionBidLandscape}s for the criteria
	// specified in the selector.         In the result, the returned
	// {@link LandscapePoint}s are grouped into         {@link CriterionBidLandscape}s
	// by their criteria, and numberResults of paging limits the total
	//         number of {@link LandscapePoint}s instead of number
	// of {@link CriterionBidLandscape}s.                  @param serviceSelector
	// Selects the entities to return bid landscapes for.         @return
	// A list of bid landscapes.         @throws ApiException when
	// there is at least one error with the request.
	GetCriterionBidLandscape(GetCriterionBidLandscape *GetCriterionBidLandscape) (*GetCriterionBidLandscapeResponse, error)

	// Returns a list of domain categories that can be used to create
	// {@link WebPage} criterion.                  @param serviceSelector
	// Selects the entities to return domain categories for.
	//   @return A list of domain categories.         @throws ApiException
	// when there is at least one error with the request.
	GetDomainCategory(GetDomainCategory *GetDomainCategory) (*GetDomainCategoryResponse, error)

	// Returns a list of {@link AdGroupBidLandscape}s for the ad groups
	// that match the query. In the         result, the returned {@link
	// LandscapePoint}s are grouped into {@link AdGroupBidLandscape}s
	//         by their ad groups, and numberResults of paging limits
	// the total number of         {@link LandscapePoint}s instead
	// of number of {@link AdGroupBidLandscape}s.
	// @param query The SQL-like AWQL query string.         @return
	// A list of bid landscapes.         @throws ApiException if problems
	// occur while parsing the query or fetching bid landscapes.
	QueryAdGroupBidLandscape(QueryAdGroupBidLandscape *QueryAdGroupBidLandscape) (*QueryAdGroupBidLandscapeResponse, error)

	// Returns a list of {@link CriterionBidLandscape}s for the campaign
	// criteria that match the         query. In the result, the returned
	// {@link LandscapePoint}s are grouped into         {@link CriterionBidLandscape}s
	// by their campaign id and criterion id, and numberResults
	//      of paging limits the total number of {@link LandscapePoint}s
	// instead of number of         {@link CriterionBidLandscape}s.
	//                  @param query The SQL-like AWQL query string.
	//         @return A list of bid landscapes.         @throws ApiException
	// if problems occur while parsing the query or fetching bid landscapes.
	QueryCampaignCriterionBidLandscape(QueryCampaignCriterionBidLandscape *QueryCampaignCriterionBidLandscape) (*QueryCampaignCriterionBidLandscapeResponse, error)

	// Returns a list of {@link CriterionBidLandscape}s for the criteria
	// that match the query. In the         result, the returned {@link
	// LandscapePoint}s are grouped into {@link CriterionBidLandscape}s
	//         by their criteria, and numberResults of paging limits
	// the total number of         {@link LandscapePoint}s instead
	// of number of {@link CriterionBidLandscape}s.
	//   @param query The SQL-like AWQL query string.         @return
	// A list of bid landscapes.         @throws ApiException if problems
	// occur while parsing the query or fetching bid landscapes.
	QueryCriterionBidLandscape(QueryCriterionBidLandscape *QueryCriterionBidLandscape) (*QueryCriterionBidLandscapeResponse, error)

	// Returns a list of domain categories that can be used to create
	// {@link WebPage} criterion.                  @param query The
	// SQL-like AWQL query string.         @return A list of domain
	// categories.         @throws ApiException if problems occur while
	// parsing the query         or fetching domain categories.
	QueryDomainCategory(QueryDomainCategory *QueryDomainCategory) (*QueryDomainCategoryResponse, error)
}

// AdGroupBidLandscapeType was auto-generated from WSDL.
type AdGroupBidLandscapeType string

// Validate validates AdGroupBidLandscapeType.
func (v AdGroupBidLandscapeType) Validate() bool {
	for _, vv := range []string{
		"UNIFORM",
		"DEFAULT",
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

// DataErrorReason was auto-generated from WSDL.
type DataErrorReason string

// Validate validates DataErrorReason.
func (v DataErrorReason) Validate() bool {
	for _, vv := range []string{
		"CANNOT_CREATE_TABLE_ENTRY",
		"NO_TABLE_ENTRY_CLASS_FOR_VIEW_TYPE",
		"TABLE_SERVICE_ERROR",
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

// Represents data about a bidlandscape for an adgroup.
type AdGroupBidLandscape struct {
	DataEntryType    *string                       `xml:"DataEntry.Type,omitempty" json:"DataEntry.Type,omitempty" yaml:"DataEntry.Type,omitempty"`
	CampaignId       *int64                        `xml:"campaignId,omitempty" json:"campaignId,omitempty" yaml:"campaignId,omitempty"`
	AdGroupId        *int64                        `xml:"adGroupId,omitempty" json:"adGroupId,omitempty" yaml:"adGroupId,omitempty"`
	StartDate        *string                       `xml:"startDate,omitempty" json:"startDate,omitempty" yaml:"startDate,omitempty"`
	EndDate          *string                       `xml:"endDate,omitempty" json:"endDate,omitempty" yaml:"endDate,omitempty"`
	LandscapePoints  []*BidLandscapeLandscapePoint `xml:"landscapePoints,omitempty" json:"landscapePoints,omitempty" yaml:"landscapePoints,omitempty"`
	Type             *AdGroupBidLandscapeType      `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	LandscapeCurrent *bool                         `xml:"landscapeCurrent,omitempty" json:"landscapeCurrent,omitempty" yaml:"landscapeCurrent,omitempty"`
	TypeAttrXSI      string                        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace    string                        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupBidLandscape) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupBidLandscape"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Contains a subset of {@link AdGroupBidLandscape} objects resulting
// from the             filtering and paging of {@link DataService#getAdGroupBidLandscape}
// call.
type AdGroupBidLandscapePage struct {
	TotalNumEntries *int                   `xml:"totalNumEntries,omitempty" json:"totalNumEntries,omitempty" yaml:"totalNumEntries,omitempty"`
	PageType        *string                `xml:"Page.Type,omitempty" json:"Page.Type,omitempty" yaml:"Page.Type,omitempty"`
	Entries         []*AdGroupBidLandscape `xml:"entries,omitempty" json:"entries,omitempty" yaml:"entries,omitempty"`
	TypeAttrXSI     string                 `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string                 `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *AdGroupBidLandscapePage) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:AdGroupBidLandscapePage"
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

// Represents data about a bid landscape for an ad group or criterion.
type BidLandscape interface{}

// A set of estimates for a criterion's performance for a specific
// bid             amount.
type BidLandscapeLandscapePoint struct {
	Bid                           *Money   `xml:"bid,omitempty" json:"bid,omitempty" yaml:"bid,omitempty"`
	Clicks                        *int64   `xml:"clicks,omitempty" json:"clicks,omitempty" yaml:"clicks,omitempty"`
	Cost                          *Money   `xml:"cost,omitempty" json:"cost,omitempty" yaml:"cost,omitempty"`
	Impressions                   *int64   `xml:"impressions,omitempty" json:"impressions,omitempty" yaml:"impressions,omitempty"`
	PromotedImpressions           *int64   `xml:"promotedImpressions,omitempty" json:"promotedImpressions,omitempty" yaml:"promotedImpressions,omitempty"`
	RequiredBudget                *Money   `xml:"requiredBudget,omitempty" json:"requiredBudget,omitempty" yaml:"requiredBudget,omitempty"`
	BiddableConversions           *float64 `xml:"biddableConversions,omitempty" json:"biddableConversions,omitempty" yaml:"biddableConversions,omitempty"`
	BiddableConversionsValue      *float64 `xml:"biddableConversionsValue,omitempty" json:"biddableConversionsValue,omitempty" yaml:"biddableConversionsValue,omitempty"`
	BidModifier                   *float64 `xml:"bidModifier,omitempty" json:"bidModifier,omitempty" yaml:"bidModifier,omitempty"`
	TotalLocalImpressions         *int64   `xml:"totalLocalImpressions,omitempty" json:"totalLocalImpressions,omitempty" yaml:"totalLocalImpressions,omitempty"`
	TotalLocalClicks              *int64   `xml:"totalLocalClicks,omitempty" json:"totalLocalClicks,omitempty" yaml:"totalLocalClicks,omitempty"`
	TotalLocalCost                *Money   `xml:"totalLocalCost,omitempty" json:"totalLocalCost,omitempty" yaml:"totalLocalCost,omitempty"`
	TotalLocalPromotedImpressions *int64   `xml:"totalLocalPromotedImpressions,omitempty" json:"totalLocalPromotedImpressions,omitempty" yaml:"totalLocalPromotedImpressions,omitempty"`
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

// Comparable types for constructing ranges with.
type ComparableValue interface{}

// The bid landscape for a criterion.  A bid landscape estimates
// how a             a criterion will perform based on different
// bid amounts.
type CriterionBidLandscape struct {
	DataEntryType   *string                       `xml:"DataEntry.Type,omitempty" json:"DataEntry.Type,omitempty" yaml:"DataEntry.Type,omitempty"`
	CampaignId      *int64                        `xml:"campaignId,omitempty" json:"campaignId,omitempty" yaml:"campaignId,omitempty"`
	AdGroupId       *int64                        `xml:"adGroupId,omitempty" json:"adGroupId,omitempty" yaml:"adGroupId,omitempty"`
	StartDate       *string                       `xml:"startDate,omitempty" json:"startDate,omitempty" yaml:"startDate,omitempty"`
	EndDate         *string                       `xml:"endDate,omitempty" json:"endDate,omitempty" yaml:"endDate,omitempty"`
	LandscapePoints []*BidLandscapeLandscapePoint `xml:"landscapePoints,omitempty" json:"landscapePoints,omitempty" yaml:"landscapePoints,omitempty"`
	CriterionId     *int64                        `xml:"criterionId,omitempty" json:"criterionId,omitempty" yaml:"criterionId,omitempty"`
	TypeAttrXSI     string                        `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string                        `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CriterionBidLandscape) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CriterionBidLandscape"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Contains a subset of {@link CriterionBidLandscapePage} objects
// resulting from the             filtering and paging of {@link
// DataService#getCriterionBidLandscape} call.
type CriterionBidLandscapePage struct {
	TotalNumEntries *int                     `xml:"totalNumEntries,omitempty" json:"totalNumEntries,omitempty" yaml:"totalNumEntries,omitempty"`
	PageType        *string                  `xml:"Page.Type,omitempty" json:"Page.Type,omitempty" yaml:"Page.Type,omitempty"`
	Entries         []*CriterionBidLandscape `xml:"entries,omitempty" json:"entries,omitempty" yaml:"entries,omitempty"`
	TypeAttrXSI     string                   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string                   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CriterionBidLandscapePage) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CriterionBidLandscapePage"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// The base class of all return types of the table service.
type DataEntry interface{}

// Represents errors thrown by the get operation.
type DataError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *DataErrorReason    `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DataError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DataError"
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

// Top level class for Dimensions.
type DimensionProperties interface{}

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

// Represents categories that AdWords finds automatically for your
// website.                          <p>             No categories
// available means that AdWords couldn't automatically find categories
// for your             website. To control how categories are
// assigned, manually add breadcrumbs to your webpages.
//                   <p>             Categories can be filtered
// by domain name or by a set of campaign IDs.
type DomainCategory struct {
	DataEntryType  *string        `xml:"DataEntry.Type,omitempty" json:"DataEntry.Type,omitempty" yaml:"DataEntry.Type,omitempty"`
	LevelOfDetail  *LevelOfDetail `xml:"levelOfDetail,omitempty" json:"levelOfDetail,omitempty" yaml:"levelOfDetail,omitempty"`
	Category       *string        `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
	Coverage       *float64       `xml:"coverage,omitempty" json:"coverage,omitempty" yaml:"coverage,omitempty"`
	DomainName     *string        `xml:"domainName,omitempty" json:"domainName,omitempty" yaml:"domainName,omitempty"`
	IsoLanguage    *string        `xml:"isoLanguage,omitempty" json:"isoLanguage,omitempty" yaml:"isoLanguage,omitempty"`
	RecommendedCpc *Money         `xml:"recommendedCpc,omitempty" json:"recommendedCpc,omitempty" yaml:"recommendedCpc,omitempty"`
	HasChild       *bool          `xml:"hasChild,omitempty" json:"hasChild,omitempty" yaml:"hasChild,omitempty"`
	CategoryRank   *int           `xml:"categoryRank,omitempty" json:"categoryRank,omitempty" yaml:"categoryRank,omitempty"`
	TypeAttrXSI    string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace  string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DomainCategory) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DomainCategory"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Contains a subset of {@link DomainCategory} objects resulting
// from             the filtering and paging of {@link DataService#getDomainCategory}
// call.
type DomainCategoryPage struct {
	TotalNumEntries *int              `xml:"totalNumEntries,omitempty" json:"totalNumEntries,omitempty" yaml:"totalNumEntries,omitempty"`
	PageType        *string           `xml:"Page.Type,omitempty" json:"Page.Type,omitempty" yaml:"Page.Type,omitempty"`
	Entries         []*DomainCategory `xml:"entries,omitempty" json:"entries,omitempty" yaml:"entries,omitempty"`
	TypeAttrXSI     string            `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string            `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DomainCategoryPage) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DomainCategoryPage"
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

// Properties of the level of detail of the report being viewed.
type LevelOfDetail struct {
	CampaignId *int64 `xml:"campaignId,omitempty" json:"campaignId,omitempty" yaml:"campaignId,omitempty"`
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

// GetAdGroupBidLandscape was auto-generated from WSDL.
type GetAdGroupBidLandscape struct {
	ServiceSelector *Selector `xml:"serviceSelector,omitempty" json:"serviceSelector,omitempty" yaml:"serviceSelector,omitempty"`
}

// GetAdGroupBidLandscapeResponse was auto-generated from WSDL.
type GetAdGroupBidLandscapeResponse struct {
	Rval *AdGroupBidLandscapePage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// GetCampaignCriterionBidLandscape was auto-generated from WSDL.
type GetCampaignCriterionBidLandscape struct {
	ServiceSelector *Selector `xml:"serviceSelector,omitempty" json:"serviceSelector,omitempty" yaml:"serviceSelector,omitempty"`
}

// GetCampaignCriterionBidLandscapeResponse was auto-generated
// from WSDL.
type GetCampaignCriterionBidLandscapeResponse struct {
	Rval *CriterionBidLandscapePage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// GetCriterionBidLandscape was auto-generated from WSDL.
type GetCriterionBidLandscape struct {
	ServiceSelector *Selector `xml:"serviceSelector,omitempty" json:"serviceSelector,omitempty" yaml:"serviceSelector,omitempty"`
}

// GetCriterionBidLandscapeResponse was auto-generated from WSDL.
type GetCriterionBidLandscapeResponse struct {
	Rval *CriterionBidLandscapePage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// GetDomainCategory was auto-generated from WSDL.
type GetDomainCategory struct {
	ServiceSelector *Selector `xml:"serviceSelector,omitempty" json:"serviceSelector,omitempty" yaml:"serviceSelector,omitempty"`
}

// GetDomainCategoryResponse was auto-generated from WSDL.
type GetDomainCategoryResponse struct {
	Rval *DomainCategoryPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// QueryAdGroupBidLandscape was auto-generated from WSDL.
type QueryAdGroupBidLandscape struct {
	Query *string `xml:"query,omitempty" json:"query,omitempty" yaml:"query,omitempty"`
}

// QueryAdGroupBidLandscapeResponse was auto-generated from WSDL.
type QueryAdGroupBidLandscapeResponse struct {
	Rval *AdGroupBidLandscapePage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// QueryCampaignCriterionBidLandscape was auto-generated from WSDL.
type QueryCampaignCriterionBidLandscape struct {
	Query *string `xml:"query,omitempty" json:"query,omitempty" yaml:"query,omitempty"`
}

// QueryCampaignCriterionBidLandscapeResponse was auto-generated
// from WSDL.
type QueryCampaignCriterionBidLandscapeResponse struct {
	Rval *CriterionBidLandscapePage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// QueryCriterionBidLandscape was auto-generated from WSDL.
type QueryCriterionBidLandscape struct {
	Query *string `xml:"query,omitempty" json:"query,omitempty" yaml:"query,omitempty"`
}

// QueryCriterionBidLandscapeResponse was auto-generated from WSDL.
type QueryCriterionBidLandscapeResponse struct {
	Rval *CriterionBidLandscapePage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// QueryDomainCategory was auto-generated from WSDL.
type QueryDomainCategory struct {
	Query *string `xml:"query,omitempty" json:"query,omitempty" yaml:"query,omitempty"`
}

// QueryDomainCategoryResponse was auto-generated from WSDL.
type QueryDomainCategoryResponse struct {
	Rval *DomainCategoryPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Operation wrapper for GetAdGroupBidLandscape.
// OperationGetAdGroupBidLandscapeRequest was auto-generated from
// WSDL.
type OperationGetAdGroupBidLandscapeRequest struct {
	GetAdGroupBidLandscape *GetAdGroupBidLandscape `xml:"getAdGroupBidLandscape,omitempty" json:"getAdGroupBidLandscape,omitempty" yaml:"getAdGroupBidLandscape,omitempty"`
}

// Operation wrapper for GetAdGroupBidLandscape.
// OperationGetAdGroupBidLandscapeResponse was auto-generated from
// WSDL.
type OperationGetAdGroupBidLandscapeResponse struct {
	GetAdGroupBidLandscapeResponse *GetAdGroupBidLandscapeResponse `xml:"getAdGroupBidLandscapeResponse,omitempty" json:"getAdGroupBidLandscapeResponse,omitempty" yaml:"getAdGroupBidLandscapeResponse,omitempty"`
}

// Operation wrapper for GetCampaignCriterionBidLandscape.
// OperationGetCampaignCriterionBidLandscapeRequest was auto-generated
// from WSDL.
type OperationGetCampaignCriterionBidLandscapeRequest struct {
	GetCampaignCriterionBidLandscape *GetCampaignCriterionBidLandscape `xml:"getCampaignCriterionBidLandscape,omitempty" json:"getCampaignCriterionBidLandscape,omitempty" yaml:"getCampaignCriterionBidLandscape,omitempty"`
}

// Operation wrapper for GetCampaignCriterionBidLandscape.
// OperationGetCampaignCriterionBidLandscapeResponse was auto-generated
// from WSDL.
type OperationGetCampaignCriterionBidLandscapeResponse struct {
	GetCampaignCriterionBidLandscapeResponse *GetCampaignCriterionBidLandscapeResponse `xml:"getCampaignCriterionBidLandscapeResponse,omitempty" json:"getCampaignCriterionBidLandscapeResponse,omitempty" yaml:"getCampaignCriterionBidLandscapeResponse,omitempty"`
}

// Operation wrapper for GetCriterionBidLandscape.
// OperationGetCriterionBidLandscapeRequest was auto-generated
// from WSDL.
type OperationGetCriterionBidLandscapeRequest struct {
	GetCriterionBidLandscape *GetCriterionBidLandscape `xml:"getCriterionBidLandscape,omitempty" json:"getCriterionBidLandscape,omitempty" yaml:"getCriterionBidLandscape,omitempty"`
}

// Operation wrapper for GetCriterionBidLandscape.
// OperationGetCriterionBidLandscapeResponse was auto-generated
// from WSDL.
type OperationGetCriterionBidLandscapeResponse struct {
	GetCriterionBidLandscapeResponse *GetCriterionBidLandscapeResponse `xml:"getCriterionBidLandscapeResponse,omitempty" json:"getCriterionBidLandscapeResponse,omitempty" yaml:"getCriterionBidLandscapeResponse,omitempty"`
}

// Operation wrapper for GetDomainCategory.
// OperationGetDomainCategoryRequest was auto-generated from WSDL.
type OperationGetDomainCategoryRequest struct {
	GetDomainCategory *GetDomainCategory `xml:"getDomainCategory,omitempty" json:"getDomainCategory,omitempty" yaml:"getDomainCategory,omitempty"`
}

// Operation wrapper for GetDomainCategory.
// OperationGetDomainCategoryResponse was auto-generated from WSDL.
type OperationGetDomainCategoryResponse struct {
	GetDomainCategoryResponse *GetDomainCategoryResponse `xml:"getDomainCategoryResponse,omitempty" json:"getDomainCategoryResponse,omitempty" yaml:"getDomainCategoryResponse,omitempty"`
}

// Operation wrapper for QueryAdGroupBidLandscape.
// OperationQueryAdGroupBidLandscapeRequest was auto-generated
// from WSDL.
type OperationQueryAdGroupBidLandscapeRequest struct {
	QueryAdGroupBidLandscape *QueryAdGroupBidLandscape `xml:"queryAdGroupBidLandscape,omitempty" json:"queryAdGroupBidLandscape,omitempty" yaml:"queryAdGroupBidLandscape,omitempty"`
}

// Operation wrapper for QueryAdGroupBidLandscape.
// OperationQueryAdGroupBidLandscapeResponse was auto-generated
// from WSDL.
type OperationQueryAdGroupBidLandscapeResponse struct {
	QueryAdGroupBidLandscapeResponse *QueryAdGroupBidLandscapeResponse `xml:"queryAdGroupBidLandscapeResponse,omitempty" json:"queryAdGroupBidLandscapeResponse,omitempty" yaml:"queryAdGroupBidLandscapeResponse,omitempty"`
}

// Operation wrapper for QueryCampaignCriterionBidLandscape.
// OperationQueryCampaignCriterionBidLandscapeRequest was auto-generated
// from WSDL.
type OperationQueryCampaignCriterionBidLandscapeRequest struct {
	QueryCampaignCriterionBidLandscape *QueryCampaignCriterionBidLandscape `xml:"queryCampaignCriterionBidLandscape,omitempty" json:"queryCampaignCriterionBidLandscape,omitempty" yaml:"queryCampaignCriterionBidLandscape,omitempty"`
}

// Operation wrapper for QueryCampaignCriterionBidLandscape.
// OperationQueryCampaignCriterionBidLandscapeResponse was auto-generated
// from WSDL.
type OperationQueryCampaignCriterionBidLandscapeResponse struct {
	QueryCampaignCriterionBidLandscapeResponse *QueryCampaignCriterionBidLandscapeResponse `xml:"queryCampaignCriterionBidLandscapeResponse,omitempty" json:"queryCampaignCriterionBidLandscapeResponse,omitempty" yaml:"queryCampaignCriterionBidLandscapeResponse,omitempty"`
}

// Operation wrapper for QueryCriterionBidLandscape.
// OperationQueryCriterionBidLandscapeRequest was auto-generated
// from WSDL.
type OperationQueryCriterionBidLandscapeRequest struct {
	QueryCriterionBidLandscape *QueryCriterionBidLandscape `xml:"queryCriterionBidLandscape,omitempty" json:"queryCriterionBidLandscape,omitempty" yaml:"queryCriterionBidLandscape,omitempty"`
}

// Operation wrapper for QueryCriterionBidLandscape.
// OperationQueryCriterionBidLandscapeResponse was auto-generated
// from WSDL.
type OperationQueryCriterionBidLandscapeResponse struct {
	QueryCriterionBidLandscapeResponse *QueryCriterionBidLandscapeResponse `xml:"queryCriterionBidLandscapeResponse,omitempty" json:"queryCriterionBidLandscapeResponse,omitempty" yaml:"queryCriterionBidLandscapeResponse,omitempty"`
}

// Operation wrapper for QueryDomainCategory.
// OperationQueryDomainCategoryRequest was auto-generated from
// WSDL.
type OperationQueryDomainCategoryRequest struct {
	QueryDomainCategory *QueryDomainCategory `xml:"queryDomainCategory,omitempty" json:"queryDomainCategory,omitempty" yaml:"queryDomainCategory,omitempty"`
}

// Operation wrapper for QueryDomainCategory.
// OperationQueryDomainCategoryResponse was auto-generated from
// WSDL.
type OperationQueryDomainCategoryResponse struct {
	QueryDomainCategoryResponse *QueryDomainCategoryResponse `xml:"queryDomainCategoryResponse,omitempty" json:"queryDomainCategoryResponse,omitempty" yaml:"queryDomainCategoryResponse,omitempty"`
}

// dataServiceInterface implements the DataServiceInterface interface.
type dataServiceInterface struct {
	cli *soap.Client
}

// Returns a list of {@link AdGroupBidLandscape}s for the ad groups
// specified in the selector.         In the result, the returned
// {@link LandscapePoint}s are grouped into         {@link AdGroupBidLandscape}s
// by their ad groups, and numberResults of paging limits the total
//         number of {@link LandscapePoint}s instead of number
// of {@link AdGroupBidLandscape}s.                  @param serviceSelector
// Selects the entities to return bid landscapes for.         @return
// A list of bid landscapes.         @throws ApiException when
// there is at least one error with the request.
func (p *dataServiceInterface) GetAdGroupBidLandscape(GetAdGroupBidLandscape *GetAdGroupBidLandscape) (*GetAdGroupBidLandscapeResponse, error) {
	α := struct {
		OperationGetAdGroupBidLandscapeRequest `xml:"tns:getAdGroupBidLandscape"`
	}{
		OperationGetAdGroupBidLandscapeRequest{
			GetAdGroupBidLandscape,
		},
	}

	γ := struct {
		OperationGetAdGroupBidLandscapeResponse `xml:"getAdGroupBidLandscapeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetAdGroupBidLandscape", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetAdGroupBidLandscapeResponse, nil
}

// Returns a list of {@link CriterionBidLandscape}s for the campaign
// criteria specified in the         selector. In the result, the
// returned {@link LandscapePoint}s are grouped into         {@link
// CriterionBidLandscape}s by their campaign id and criterion id,
// and numberResults         of paging limits the total number
// of {@link LandscapePoint}s instead of number of         {@link
// CriterionBidLandscape}s.                  @param serviceSelector
// Selects the entities to return bid landscapes for.         @return
// A list of bid landscapes.         @throws ApiException when
// there is at least one error with the request.
func (p *dataServiceInterface) GetCampaignCriterionBidLandscape(GetCampaignCriterionBidLandscape *GetCampaignCriterionBidLandscape) (*GetCampaignCriterionBidLandscapeResponse, error) {
	α := struct {
		OperationGetCampaignCriterionBidLandscapeRequest `xml:"tns:getCampaignCriterionBidLandscape"`
	}{
		OperationGetCampaignCriterionBidLandscapeRequest{
			GetCampaignCriterionBidLandscape,
		},
	}

	γ := struct {
		OperationGetCampaignCriterionBidLandscapeResponse `xml:"getCampaignCriterionBidLandscapeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetCampaignCriterionBidLandscape", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetCampaignCriterionBidLandscapeResponse, nil
}

// Returns a list of {@link CriterionBidLandscape}s for the criteria
// specified in the selector.         In the result, the returned
// {@link LandscapePoint}s are grouped into         {@link CriterionBidLandscape}s
// by their criteria, and numberResults of paging limits the total
//         number of {@link LandscapePoint}s instead of number
// of {@link CriterionBidLandscape}s.                  @param serviceSelector
// Selects the entities to return bid landscapes for.         @return
// A list of bid landscapes.         @throws ApiException when
// there is at least one error with the request.
func (p *dataServiceInterface) GetCriterionBidLandscape(GetCriterionBidLandscape *GetCriterionBidLandscape) (*GetCriterionBidLandscapeResponse, error) {
	α := struct {
		OperationGetCriterionBidLandscapeRequest `xml:"tns:getCriterionBidLandscape"`
	}{
		OperationGetCriterionBidLandscapeRequest{
			GetCriterionBidLandscape,
		},
	}

	γ := struct {
		OperationGetCriterionBidLandscapeResponse `xml:"getCriterionBidLandscapeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetCriterionBidLandscape", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetCriterionBidLandscapeResponse, nil
}

// Returns a list of domain categories that can be used to create
// {@link WebPage} criterion.                  @param serviceSelector
// Selects the entities to return domain categories for.
//   @return A list of domain categories.         @throws ApiException
// when there is at least one error with the request.
func (p *dataServiceInterface) GetDomainCategory(GetDomainCategory *GetDomainCategory) (*GetDomainCategoryResponse, error) {
	α := struct {
		OperationGetDomainCategoryRequest `xml:"tns:getDomainCategory"`
	}{
		OperationGetDomainCategoryRequest{
			GetDomainCategory,
		},
	}

	γ := struct {
		OperationGetDomainCategoryResponse `xml:"getDomainCategoryResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetDomainCategory", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetDomainCategoryResponse, nil
}

// Returns a list of {@link AdGroupBidLandscape}s for the ad groups
// that match the query. In the         result, the returned {@link
// LandscapePoint}s are grouped into {@link AdGroupBidLandscape}s
//         by their ad groups, and numberResults of paging limits
// the total number of         {@link LandscapePoint}s instead
// of number of {@link AdGroupBidLandscape}s.
// @param query The SQL-like AWQL query string.         @return
// A list of bid landscapes.         @throws ApiException if problems
// occur while parsing the query or fetching bid landscapes.
func (p *dataServiceInterface) QueryAdGroupBidLandscape(QueryAdGroupBidLandscape *QueryAdGroupBidLandscape) (*QueryAdGroupBidLandscapeResponse, error) {
	α := struct {
		OperationQueryAdGroupBidLandscapeRequest `xml:"tns:queryAdGroupBidLandscape"`
	}{
		OperationQueryAdGroupBidLandscapeRequest{
			QueryAdGroupBidLandscape,
		},
	}

	γ := struct {
		OperationQueryAdGroupBidLandscapeResponse `xml:"queryAdGroupBidLandscapeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("QueryAdGroupBidLandscape", α, &γ); err != nil {
		return nil, err
	}
	return γ.QueryAdGroupBidLandscapeResponse, nil
}

// Returns a list of {@link CriterionBidLandscape}s for the campaign
// criteria that match the         query. In the result, the returned
// {@link LandscapePoint}s are grouped into         {@link CriterionBidLandscape}s
// by their campaign id and criterion id, and numberResults
//      of paging limits the total number of {@link LandscapePoint}s
// instead of number of         {@link CriterionBidLandscape}s.
//                  @param query The SQL-like AWQL query string.
//         @return A list of bid landscapes.         @throws ApiException
// if problems occur while parsing the query or fetching bid landscapes.
func (p *dataServiceInterface) QueryCampaignCriterionBidLandscape(QueryCampaignCriterionBidLandscape *QueryCampaignCriterionBidLandscape) (*QueryCampaignCriterionBidLandscapeResponse, error) {
	α := struct {
		OperationQueryCampaignCriterionBidLandscapeRequest `xml:"tns:queryCampaignCriterionBidLandscape"`
	}{
		OperationQueryCampaignCriterionBidLandscapeRequest{
			QueryCampaignCriterionBidLandscape,
		},
	}

	γ := struct {
		OperationQueryCampaignCriterionBidLandscapeResponse `xml:"queryCampaignCriterionBidLandscapeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("QueryCampaignCriterionBidLandscape", α, &γ); err != nil {
		return nil, err
	}
	return γ.QueryCampaignCriterionBidLandscapeResponse, nil
}

// Returns a list of {@link CriterionBidLandscape}s for the criteria
// that match the query. In the         result, the returned {@link
// LandscapePoint}s are grouped into {@link CriterionBidLandscape}s
//         by their criteria, and numberResults of paging limits
// the total number of         {@link LandscapePoint}s instead
// of number of {@link CriterionBidLandscape}s.
//   @param query The SQL-like AWQL query string.         @return
// A list of bid landscapes.         @throws ApiException if problems
// occur while parsing the query or fetching bid landscapes.
func (p *dataServiceInterface) QueryCriterionBidLandscape(QueryCriterionBidLandscape *QueryCriterionBidLandscape) (*QueryCriterionBidLandscapeResponse, error) {
	α := struct {
		OperationQueryCriterionBidLandscapeRequest `xml:"tns:queryCriterionBidLandscape"`
	}{
		OperationQueryCriterionBidLandscapeRequest{
			QueryCriterionBidLandscape,
		},
	}

	γ := struct {
		OperationQueryCriterionBidLandscapeResponse `xml:"queryCriterionBidLandscapeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("QueryCriterionBidLandscape", α, &γ); err != nil {
		return nil, err
	}
	return γ.QueryCriterionBidLandscapeResponse, nil
}

// Returns a list of domain categories that can be used to create
// {@link WebPage} criterion.                  @param query The
// SQL-like AWQL query string.         @return A list of domain
// categories.         @throws ApiException if problems occur while
// parsing the query         or fetching domain categories.
func (p *dataServiceInterface) QueryDomainCategory(QueryDomainCategory *QueryDomainCategory) (*QueryDomainCategoryResponse, error) {
	α := struct {
		OperationQueryDomainCategoryRequest `xml:"tns:queryDomainCategory"`
	}{
		OperationQueryDomainCategoryRequest{
			QueryDomainCategory,
		},
	}

	γ := struct {
		OperationQueryDomainCategoryResponse `xml:"queryDomainCategoryResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("QueryDomainCategory", α, &γ); err != nil {
		return nil, err
	}
	return γ.QueryDomainCategoryResponse, nil
}
