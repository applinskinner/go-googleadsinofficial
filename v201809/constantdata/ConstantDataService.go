package constantdata

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "https://adwords.google.com/api/adwords/cm/v201809"

// NewConstantDataServiceInterface creates an initializes a ConstantDataServiceInterface.
func NewConstantDataServiceInterface(cli *soap.Client) ConstantDataServiceInterface {
	return &constantDataServiceInterface{cli}
}

// ConstantDataServiceInterface was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ConstantDataServiceInterface interface {
	// Returns a list of all age range criteria.                  @return
	// A list of age ranges.         @throws ApiException when there
	// is at least one error with the request.
	GetAgeRangeCriterion(GetAgeRangeCriterion *GetAgeRangeCriterion) (*GetAgeRangeCriterionResponse, error)

	// Returns a list of all carrier criteria.                  @return
	// A list of carriers.         @throws ApiException when there
	// is at least one error with the request.
	GetCarrierCriterion(GetCarrierCriterion *GetCarrierCriterion) (*GetCarrierCriterionResponse, error)

	// Returns a list of all gender criteria.                  @return
	// A list of genders.         @throws ApiException when there is
	// at least one error with the request.
	GetGenderCriterion(GetGenderCriterion *GetGenderCriterion) (*GetGenderCriterionResponse, error)

	// Returns a list of all language criteria.                  @return
	// A list of languages.         @throws ApiException when there
	// is at least one error with the request.
	GetLanguageCriterion(GetLanguageCriterion *GetLanguageCriterion) (*GetLanguageCriterionResponse, error)

	// Returns a list of all mobile app category criteria.
	//          @return A list of mobile app categories.         @throws
	// ApiException when there is at least one error with the request.
	GetMobileAppCategoryCriterion(GetMobileAppCategoryCriterion *GetMobileAppCategoryCriterion) (*GetMobileAppCategoryCriterionResponse, error)

	// Returns a list of all mobile devices.                  @return
	// A list of mobile devices.         @throws ApiException when
	// there is at least one error with the request.
	GetMobileDeviceCriterion(GetMobileDeviceCriterion *GetMobileDeviceCriterion) (*GetMobileDeviceCriterionResponse, error)

	// Returns a list of all operating system version criteria.
	//               @return A list of operating system versions.
	//        @throws ApiException when there is at least one error
	// with the request.
	GetOperatingSystemVersionCriterion(GetOperatingSystemVersionCriterion *GetOperatingSystemVersionCriterion) (*GetOperatingSystemVersionCriterionResponse, error)

	// Returns a list of shopping bidding categories.
	//     A country predicate must be included in the selector, only
	// {@link Predicate.Operator#EQUALS}         and {@link Predicate.Operator#IN}
	// with a single value are supported in the country predicate.
	//         An empty parentDimensionType predicate will filter for
	// root categories.                  @return A list of shopping
	// bidding categories.         @throws ApiException when there
	// is at least one error with the request.
	GetProductBiddingCategoryData(GetProductBiddingCategoryData *GetProductBiddingCategoryData) (*GetProductBiddingCategoryDataResponse, error)

	// Returns a list of user interests.                  @param userInterestTaxonomyType
	// The type of taxonomy to use when requesting user interests.
	//         @return A list of user interests.         @throws ApiException
	// when there is at least one error with the request.
	GetUserInterestCriterion(GetUserInterestCriterion *GetUserInterestCriterion) (*GetUserInterestCriterionResponse, error)

	// Returns a list of content verticals.                  @return
	// A list of verticals.         @throws ApiException when there
	// is at least one error with the request.
	GetVerticalCriterion(GetVerticalCriterion *GetVerticalCriterion) (*GetVerticalCriterionResponse, error)
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

// ConstantDataServiceUserInterestTaxonomyType was auto-generated
// from WSDL.
type ConstantDataServiceUserInterestTaxonomyType string

// Validate validates ConstantDataServiceUserInterestTaxonomyType.
func (v ConstantDataServiceUserInterestTaxonomyType) Validate() bool {
	for _, vv := range []string{
		"BRAND",
		"IN_MARKET",
		"MOBILE_APP_INSTALL_USER",
		"VERTICAL_GEO",
		"NEW_SMART_PHONE_USER",
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

// MobileDeviceDeviceType was auto-generated from WSDL.
type MobileDeviceDeviceType string

// Validate validates MobileDeviceDeviceType.
func (v MobileDeviceDeviceType) Validate() bool {
	for _, vv := range []string{
		"DEVICE_TYPE_MOBILE",
		"DEVICE_TYPE_TABLET",
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

// OperatingSystemVersionOperatorType was auto-generated from WSDL.
type OperatingSystemVersionOperatorType string

// Validate validates OperatingSystemVersionOperatorType.
func (v OperatingSystemVersionOperatorType) Validate() bool {
	for _, vv := range []string{
		"GREATER_THAN_EQUAL_TO",
		"EQUAL_TO",
		"UNKNOWN",
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

// ShoppingBiddingDimensionStatus was auto-generated from WSDL.
type ShoppingBiddingDimensionStatus string

// Validate validates ShoppingBiddingDimensionStatus.
func (v ShoppingBiddingDimensionStatus) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"ACTIVE",
		"OBSOLETE",
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
// of this type can only be created using an ID.             <span
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

// Represents a Carrier Criterion.             <p>A criterion of
// this type can only be created using an ID.             <span
// class="constraint AdxEnabled">This is enabled for AdX.</span>
type Carrier struct {
	Id            *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	Name          *string        `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	CountryCode   *string        `xml:"countryCode,omitempty" json:"countryCode,omitempty" yaml:"countryCode,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Carrier) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Carrier"
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

// Marker interface for ConstantDataService objects. This is primarily
// required for field             catalog generation.
type ConstantData struct {
	ConstantDataType *string `xml:"ConstantData.Type,omitempty" json:"ConstantData.Type,omitempty" yaml:"ConstantData.Type,omitempty"`
}

// Represents a criterion (such as a keyword, placement, or vertical).
//             <span class="constraint AdxEnabled">This is disabled
// for AdX when it is contained within Operators: ADD, SET.</span>
type Criterion struct {
	Id            *int64         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type          *CriterionType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType *string        `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
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

// Represents a Gender criterion.             <p>A criterion of
// this type can only be created using an ID.             <span
// class="constraint AdxEnabled">This is disabled for AdX when
// it is contained within Operators: ADD, SET.</span>
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

// Represents a Mobile Device Criterion.             <p>A criterion
// of this type can only be created using an ID.             <span
// class="constraint AdxEnabled">This is enabled for AdX.</span>
type MobileDevice struct {
	Id                  *int64                  `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type                *CriterionType          `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType       *string                 `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	DeviceName          *string                 `xml:"deviceName,omitempty" json:"deviceName,omitempty" yaml:"deviceName,omitempty"`
	ManufacturerName    *string                 `xml:"manufacturerName,omitempty" json:"manufacturerName,omitempty" yaml:"manufacturerName,omitempty"`
	DeviceType          *MobileDeviceDeviceType `xml:"deviceType,omitempty" json:"deviceType,omitempty" yaml:"deviceType,omitempty"`
	OperatingSystemName *string                 `xml:"operatingSystemName,omitempty" json:"operatingSystemName,omitempty" yaml:"operatingSystemName,omitempty"`
	TypeAttrXSI         string                  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string                  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MobileDevice) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MobileDevice"
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

// Represents an Operating System Version Criterion.
//   <a href="/adwords/api/docs/appendix/mobileplatforms">View
// the complete             list of available mobile platforms</a>.
// You can also get the list from             {@link ConstantDataService#getOperatingSystemVersionCriterion
// ConstantDataService}.             <p>A criterion of this type
// can only be created using an ID.             <span class="constraint
// AdxEnabled">This is enabled for AdX.</span>
type OperatingSystemVersion struct {
	Id             *int64                              `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Type           *CriterionType                      `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	CriterionType  *string                             `xml:"Criterion.Type,omitempty" json:"Criterion.Type,omitempty" yaml:"Criterion.Type,omitempty"`
	Name           *string                             `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	OsMajorVersion *int                                `xml:"osMajorVersion,omitempty" json:"osMajorVersion,omitempty" yaml:"osMajorVersion,omitempty"`
	OsMinorVersion *int                                `xml:"osMinorVersion,omitempty" json:"osMinorVersion,omitempty" yaml:"osMinorVersion,omitempty"`
	OperatorType   *OperatingSystemVersionOperatorType `xml:"operatorType,omitempty" json:"operatorType,omitempty" yaml:"operatorType,omitempty"`
	TypeAttrXSI    string                              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace  string                              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *OperatingSystemVersion) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:OperatingSystemVersion"
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

// Specifies how the resulting information should be sorted.
type OrderBy struct {
	Field     *string    `xml:"field,omitempty" json:"field,omitempty" yaml:"field,omitempty"`
	SortOrder *SortOrder `xml:"sortOrder,omitempty" json:"sortOrder,omitempty" yaml:"sortOrder,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
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

// The taxonomy of ProductBiddingCategory dimension values.
//                       <p>Clients use this to convert between
// human-readable category names / display strings and
//     ProductBiddingCategory instances.
type ProductBiddingCategoryData struct {
	ConstantDataType     *string                         `xml:"ConstantData.Type,omitempty" json:"ConstantData.Type,omitempty" yaml:"ConstantData.Type,omitempty"`
	DimensionValue       *ProductBiddingCategory         `xml:"dimensionValue,omitempty" json:"dimensionValue,omitempty" yaml:"dimensionValue,omitempty"`
	ParentDimensionValue *ProductBiddingCategory         `xml:"parentDimensionValue,omitempty" json:"parentDimensionValue,omitempty" yaml:"parentDimensionValue,omitempty"`
	Country              *string                         `xml:"country,omitempty" json:"country,omitempty" yaml:"country,omitempty"`
	Status               *ShoppingBiddingDimensionStatus `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	DisplayValue         []*String_StringMapEntry        `xml:"displayValue,omitempty" json:"displayValue,omitempty" yaml:"displayValue,omitempty"`
	TypeAttrXSI          string                          `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string                          `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ProductBiddingCategoryData) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ProductBiddingCategoryData"
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

// This represents an entry in a map with a key of type String
//             and value of type String.
type String_StringMapEntry struct {
	Key   *string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
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

// GetAgeRangeCriterion was auto-generated from WSDL.
type GetAgeRangeCriterion struct {
}

// GetAgeRangeCriterionResponse was auto-generated from WSDL.
type GetAgeRangeCriterionResponse struct {
	Rval []*AgeRange `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// GetCarrierCriterion was auto-generated from WSDL.
type GetCarrierCriterion struct {
}

// GetCarrierCriterionResponse was auto-generated from WSDL.
type GetCarrierCriterionResponse struct {
	Rval []*Carrier `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// GetGenderCriterion was auto-generated from WSDL.
type GetGenderCriterion struct {
}

// GetGenderCriterionResponse was auto-generated from WSDL.
type GetGenderCriterionResponse struct {
	Rval []*Gender `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// GetLanguageCriterion was auto-generated from WSDL.
type GetLanguageCriterion struct {
}

// GetLanguageCriterionResponse was auto-generated from WSDL.
type GetLanguageCriterionResponse struct {
	Rval []*string `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// GetMobileAppCategoryCriterion was auto-generated from WSDL.
type GetMobileAppCategoryCriterion struct {
}

// GetMobileAppCategoryCriterionResponse was auto-generated from
// WSDL.
type GetMobileAppCategoryCriterionResponse struct {
	Rval []*MobileAppCategory `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// GetMobileDeviceCriterion was auto-generated from WSDL.
type GetMobileDeviceCriterion struct {
}

// GetMobileDeviceCriterionResponse was auto-generated from WSDL.
type GetMobileDeviceCriterionResponse struct {
	Rval []*MobileDevice `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// GetOperatingSystemVersionCriterion was auto-generated from WSDL.
type GetOperatingSystemVersionCriterion struct {
}

// GetOperatingSystemVersionCriterionResponse was auto-generated
// from WSDL.
type GetOperatingSystemVersionCriterionResponse struct {
	Rval []*OperatingSystemVersion `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// GetProductBiddingCategoryData was auto-generated from WSDL.
type GetProductBiddingCategoryData struct {
	Selector *Selector `xml:"selector,omitempty" json:"selector,omitempty" yaml:"selector,omitempty"`
}

// GetProductBiddingCategoryDataResponse was auto-generated from
// WSDL.
type GetProductBiddingCategoryDataResponse struct {
	Rval []*ProductBiddingCategoryData `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// GetUserInterestCriterion was auto-generated from WSDL.
type GetUserInterestCriterion struct {
	UserInterestTaxonomyType *ConstantDataServiceUserInterestTaxonomyType `xml:"userInterestTaxonomyType,omitempty" json:"userInterestTaxonomyType,omitempty" yaml:"userInterestTaxonomyType,omitempty"`
}

// GetUserInterestCriterionResponse was auto-generated from WSDL.
type GetUserInterestCriterionResponse struct {
	Rval []*CriterionUserInterest `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// GetVerticalCriterion was auto-generated from WSDL.
type GetVerticalCriterion struct {
}

// GetVerticalCriterionResponse was auto-generated from WSDL.
type GetVerticalCriterionResponse struct {
	Rval []*Vertical `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Operation wrapper for GetAgeRangeCriterion.
// OperationGetAgeRangeCriterionRequest was auto-generated from
// WSDL.
type OperationGetAgeRangeCriterionRequest struct {
	GetAgeRangeCriterion *GetAgeRangeCriterion `xml:"getAgeRangeCriterion,omitempty" json:"getAgeRangeCriterion,omitempty" yaml:"getAgeRangeCriterion,omitempty"`
}

// Operation wrapper for GetAgeRangeCriterion.
// OperationGetAgeRangeCriterionResponse was auto-generated from
// WSDL.
type OperationGetAgeRangeCriterionResponse struct {
	GetAgeRangeCriterionResponse *GetAgeRangeCriterionResponse `xml:"getAgeRangeCriterionResponse,omitempty" json:"getAgeRangeCriterionResponse,omitempty" yaml:"getAgeRangeCriterionResponse,omitempty"`
}

// Operation wrapper for GetCarrierCriterion.
// OperationGetCarrierCriterionRequest was auto-generated from
// WSDL.
type OperationGetCarrierCriterionRequest struct {
	GetCarrierCriterion *GetCarrierCriterion `xml:"getCarrierCriterion,omitempty" json:"getCarrierCriterion,omitempty" yaml:"getCarrierCriterion,omitempty"`
}

// Operation wrapper for GetCarrierCriterion.
// OperationGetCarrierCriterionResponse was auto-generated from
// WSDL.
type OperationGetCarrierCriterionResponse struct {
	GetCarrierCriterionResponse *GetCarrierCriterionResponse `xml:"getCarrierCriterionResponse,omitempty" json:"getCarrierCriterionResponse,omitempty" yaml:"getCarrierCriterionResponse,omitempty"`
}

// Operation wrapper for GetGenderCriterion.
// OperationGetGenderCriterionRequest was auto-generated from WSDL.
type OperationGetGenderCriterionRequest struct {
	GetGenderCriterion *GetGenderCriterion `xml:"getGenderCriterion,omitempty" json:"getGenderCriterion,omitempty" yaml:"getGenderCriterion,omitempty"`
}

// Operation wrapper for GetGenderCriterion.
// OperationGetGenderCriterionResponse was auto-generated from
// WSDL.
type OperationGetGenderCriterionResponse struct {
	GetGenderCriterionResponse *GetGenderCriterionResponse `xml:"getGenderCriterionResponse,omitempty" json:"getGenderCriterionResponse,omitempty" yaml:"getGenderCriterionResponse,omitempty"`
}

// Operation wrapper for GetLanguageCriterion.
// OperationGetLanguageCriterionRequest was auto-generated from
// WSDL.
type OperationGetLanguageCriterionRequest struct {
	GetLanguageCriterion *GetLanguageCriterion `xml:"getLanguageCriterion,omitempty" json:"getLanguageCriterion,omitempty" yaml:"getLanguageCriterion,omitempty"`
}

// Operation wrapper for GetLanguageCriterion.
// OperationGetLanguageCriterionResponse was auto-generated from
// WSDL.
type OperationGetLanguageCriterionResponse struct {
	GetLanguageCriterionResponse *GetLanguageCriterionResponse `xml:"getLanguageCriterionResponse,omitempty" json:"getLanguageCriterionResponse,omitempty" yaml:"getLanguageCriterionResponse,omitempty"`
}

// Operation wrapper for GetMobileAppCategoryCriterion.
// OperationGetMobileAppCategoryCriterionRequest was auto-generated
// from WSDL.
type OperationGetMobileAppCategoryCriterionRequest struct {
	GetMobileAppCategoryCriterion *GetMobileAppCategoryCriterion `xml:"getMobileAppCategoryCriterion,omitempty" json:"getMobileAppCategoryCriterion,omitempty" yaml:"getMobileAppCategoryCriterion,omitempty"`
}

// Operation wrapper for GetMobileAppCategoryCriterion.
// OperationGetMobileAppCategoryCriterionResponse was auto-generated
// from WSDL.
type OperationGetMobileAppCategoryCriterionResponse struct {
	GetMobileAppCategoryCriterionResponse *GetMobileAppCategoryCriterionResponse `xml:"getMobileAppCategoryCriterionResponse,omitempty" json:"getMobileAppCategoryCriterionResponse,omitempty" yaml:"getMobileAppCategoryCriterionResponse,omitempty"`
}

// Operation wrapper for GetMobileDeviceCriterion.
// OperationGetMobileDeviceCriterionRequest was auto-generated
// from WSDL.
type OperationGetMobileDeviceCriterionRequest struct {
	GetMobileDeviceCriterion *GetMobileDeviceCriterion `xml:"getMobileDeviceCriterion,omitempty" json:"getMobileDeviceCriterion,omitempty" yaml:"getMobileDeviceCriterion,omitempty"`
}

// Operation wrapper for GetMobileDeviceCriterion.
// OperationGetMobileDeviceCriterionResponse was auto-generated
// from WSDL.
type OperationGetMobileDeviceCriterionResponse struct {
	GetMobileDeviceCriterionResponse *GetMobileDeviceCriterionResponse `xml:"getMobileDeviceCriterionResponse,omitempty" json:"getMobileDeviceCriterionResponse,omitempty" yaml:"getMobileDeviceCriterionResponse,omitempty"`
}

// Operation wrapper for GetOperatingSystemVersionCriterion.
// OperationGetOperatingSystemVersionCriterionRequest was auto-generated
// from WSDL.
type OperationGetOperatingSystemVersionCriterionRequest struct {
	GetOperatingSystemVersionCriterion *GetOperatingSystemVersionCriterion `xml:"getOperatingSystemVersionCriterion,omitempty" json:"getOperatingSystemVersionCriterion,omitempty" yaml:"getOperatingSystemVersionCriterion,omitempty"`
}

// Operation wrapper for GetOperatingSystemVersionCriterion.
// OperationGetOperatingSystemVersionCriterionResponse was auto-generated
// from WSDL.
type OperationGetOperatingSystemVersionCriterionResponse struct {
	GetOperatingSystemVersionCriterionResponse *GetOperatingSystemVersionCriterionResponse `xml:"getOperatingSystemVersionCriterionResponse,omitempty" json:"getOperatingSystemVersionCriterionResponse,omitempty" yaml:"getOperatingSystemVersionCriterionResponse,omitempty"`
}

// Operation wrapper for GetProductBiddingCategoryData.
// OperationGetProductBiddingCategoryDataRequest was auto-generated
// from WSDL.
type OperationGetProductBiddingCategoryDataRequest struct {
	GetProductBiddingCategoryData *GetProductBiddingCategoryData `xml:"getProductBiddingCategoryData,omitempty" json:"getProductBiddingCategoryData,omitempty" yaml:"getProductBiddingCategoryData,omitempty"`
}

// Operation wrapper for GetProductBiddingCategoryData.
// OperationGetProductBiddingCategoryDataResponse was auto-generated
// from WSDL.
type OperationGetProductBiddingCategoryDataResponse struct {
	GetProductBiddingCategoryDataResponse *GetProductBiddingCategoryDataResponse `xml:"getProductBiddingCategoryDataResponse,omitempty" json:"getProductBiddingCategoryDataResponse,omitempty" yaml:"getProductBiddingCategoryDataResponse,omitempty"`
}

// Operation wrapper for GetUserInterestCriterion.
// OperationGetUserInterestCriterionRequest was auto-generated
// from WSDL.
type OperationGetUserInterestCriterionRequest struct {
	GetUserInterestCriterion *GetUserInterestCriterion `xml:"getUserInterestCriterion,omitempty" json:"getUserInterestCriterion,omitempty" yaml:"getUserInterestCriterion,omitempty"`
}

// Operation wrapper for GetUserInterestCriterion.
// OperationGetUserInterestCriterionResponse was auto-generated
// from WSDL.
type OperationGetUserInterestCriterionResponse struct {
	GetUserInterestCriterionResponse *GetUserInterestCriterionResponse `xml:"getUserInterestCriterionResponse,omitempty" json:"getUserInterestCriterionResponse,omitempty" yaml:"getUserInterestCriterionResponse,omitempty"`
}

// Operation wrapper for GetVerticalCriterion.
// OperationGetVerticalCriterionRequest was auto-generated from
// WSDL.
type OperationGetVerticalCriterionRequest struct {
	GetVerticalCriterion *GetVerticalCriterion `xml:"getVerticalCriterion,omitempty" json:"getVerticalCriterion,omitempty" yaml:"getVerticalCriterion,omitempty"`
}

// Operation wrapper for GetVerticalCriterion.
// OperationGetVerticalCriterionResponse was auto-generated from
// WSDL.
type OperationGetVerticalCriterionResponse struct {
	GetVerticalCriterionResponse *GetVerticalCriterionResponse `xml:"getVerticalCriterionResponse,omitempty" json:"getVerticalCriterionResponse,omitempty" yaml:"getVerticalCriterionResponse,omitempty"`
}

// constantDataServiceInterface implements the ConstantDataServiceInterface interface.
type constantDataServiceInterface struct {
	cli *soap.Client
}

// Returns a list of all age range criteria.                  @return
// A list of age ranges.         @throws ApiException when there
// is at least one error with the request.
func (p *constantDataServiceInterface) GetAgeRangeCriterion(GetAgeRangeCriterion *GetAgeRangeCriterion) (*GetAgeRangeCriterionResponse, error) {
	α := struct {
		OperationGetAgeRangeCriterionRequest `xml:"tns:getAgeRangeCriterion"`
	}{
		OperationGetAgeRangeCriterionRequest{
			GetAgeRangeCriterion,
		},
	}

	γ := struct {
		OperationGetAgeRangeCriterionResponse `xml:"getAgeRangeCriterionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetAgeRangeCriterion", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetAgeRangeCriterionResponse, nil
}

// Returns a list of all carrier criteria.                  @return
// A list of carriers.         @throws ApiException when there
// is at least one error with the request.
func (p *constantDataServiceInterface) GetCarrierCriterion(GetCarrierCriterion *GetCarrierCriterion) (*GetCarrierCriterionResponse, error) {
	α := struct {
		OperationGetCarrierCriterionRequest `xml:"tns:getCarrierCriterion"`
	}{
		OperationGetCarrierCriterionRequest{
			GetCarrierCriterion,
		},
	}

	γ := struct {
		OperationGetCarrierCriterionResponse `xml:"getCarrierCriterionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetCarrierCriterion", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetCarrierCriterionResponse, nil
}

// Returns a list of all gender criteria.                  @return
// A list of genders.         @throws ApiException when there is
// at least one error with the request.
func (p *constantDataServiceInterface) GetGenderCriterion(GetGenderCriterion *GetGenderCriterion) (*GetGenderCriterionResponse, error) {
	α := struct {
		OperationGetGenderCriterionRequest `xml:"tns:getGenderCriterion"`
	}{
		OperationGetGenderCriterionRequest{
			GetGenderCriterion,
		},
	}

	γ := struct {
		OperationGetGenderCriterionResponse `xml:"getGenderCriterionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetGenderCriterion", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetGenderCriterionResponse, nil
}

// Returns a list of all language criteria.                  @return
// A list of languages.         @throws ApiException when there
// is at least one error with the request.
func (p *constantDataServiceInterface) GetLanguageCriterion(GetLanguageCriterion *GetLanguageCriterion) (*GetLanguageCriterionResponse, error) {
	α := struct {
		OperationGetLanguageCriterionRequest `xml:"tns:getLanguageCriterion"`
	}{
		OperationGetLanguageCriterionRequest{
			GetLanguageCriterion,
		},
	}

	γ := struct {
		OperationGetLanguageCriterionResponse `xml:"getLanguageCriterionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetLanguageCriterion", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetLanguageCriterionResponse, nil
}

// Returns a list of all mobile app category criteria.
//          @return A list of mobile app categories.         @throws
// ApiException when there is at least one error with the request.
func (p *constantDataServiceInterface) GetMobileAppCategoryCriterion(GetMobileAppCategoryCriterion *GetMobileAppCategoryCriterion) (*GetMobileAppCategoryCriterionResponse, error) {
	α := struct {
		OperationGetMobileAppCategoryCriterionRequest `xml:"tns:getMobileAppCategoryCriterion"`
	}{
		OperationGetMobileAppCategoryCriterionRequest{
			GetMobileAppCategoryCriterion,
		},
	}

	γ := struct {
		OperationGetMobileAppCategoryCriterionResponse `xml:"getMobileAppCategoryCriterionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetMobileAppCategoryCriterion", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetMobileAppCategoryCriterionResponse, nil
}

// Returns a list of all mobile devices.                  @return
// A list of mobile devices.         @throws ApiException when
// there is at least one error with the request.
func (p *constantDataServiceInterface) GetMobileDeviceCriterion(GetMobileDeviceCriterion *GetMobileDeviceCriterion) (*GetMobileDeviceCriterionResponse, error) {
	α := struct {
		OperationGetMobileDeviceCriterionRequest `xml:"tns:getMobileDeviceCriterion"`
	}{
		OperationGetMobileDeviceCriterionRequest{
			GetMobileDeviceCriterion,
		},
	}

	γ := struct {
		OperationGetMobileDeviceCriterionResponse `xml:"getMobileDeviceCriterionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetMobileDeviceCriterion", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetMobileDeviceCriterionResponse, nil
}

// Returns a list of all operating system version criteria.
//               @return A list of operating system versions.
//        @throws ApiException when there is at least one error
// with the request.
func (p *constantDataServiceInterface) GetOperatingSystemVersionCriterion(GetOperatingSystemVersionCriterion *GetOperatingSystemVersionCriterion) (*GetOperatingSystemVersionCriterionResponse, error) {
	α := struct {
		OperationGetOperatingSystemVersionCriterionRequest `xml:"tns:getOperatingSystemVersionCriterion"`
	}{
		OperationGetOperatingSystemVersionCriterionRequest{
			GetOperatingSystemVersionCriterion,
		},
	}

	γ := struct {
		OperationGetOperatingSystemVersionCriterionResponse `xml:"getOperatingSystemVersionCriterionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetOperatingSystemVersionCriterion", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetOperatingSystemVersionCriterionResponse, nil
}

// Returns a list of shopping bidding categories.
//     A country predicate must be included in the selector, only
// {@link Predicate.Operator#EQUALS}         and {@link Predicate.Operator#IN}
// with a single value are supported in the country predicate.
//         An empty parentDimensionType predicate will filter for
// root categories.                  @return A list of shopping
// bidding categories.         @throws ApiException when there
// is at least one error with the request.
func (p *constantDataServiceInterface) GetProductBiddingCategoryData(GetProductBiddingCategoryData *GetProductBiddingCategoryData) (*GetProductBiddingCategoryDataResponse, error) {
	α := struct {
		OperationGetProductBiddingCategoryDataRequest `xml:"tns:getProductBiddingCategoryData"`
	}{
		OperationGetProductBiddingCategoryDataRequest{
			GetProductBiddingCategoryData,
		},
	}

	γ := struct {
		OperationGetProductBiddingCategoryDataResponse `xml:"getProductBiddingCategoryDataResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetProductBiddingCategoryData", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetProductBiddingCategoryDataResponse, nil
}

// Returns a list of user interests.                  @param userInterestTaxonomyType
// The type of taxonomy to use when requesting user interests.
//         @return A list of user interests.         @throws ApiException
// when there is at least one error with the request.
func (p *constantDataServiceInterface) GetUserInterestCriterion(GetUserInterestCriterion *GetUserInterestCriterion) (*GetUserInterestCriterionResponse, error) {
	α := struct {
		OperationGetUserInterestCriterionRequest `xml:"tns:getUserInterestCriterion"`
	}{
		OperationGetUserInterestCriterionRequest{
			GetUserInterestCriterion,
		},
	}

	γ := struct {
		OperationGetUserInterestCriterionResponse `xml:"getUserInterestCriterionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetUserInterestCriterion", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetUserInterestCriterionResponse, nil
}

// Returns a list of content verticals.                  @return
// A list of verticals.         @throws ApiException when there
// is at least one error with the request.
func (p *constantDataServiceInterface) GetVerticalCriterion(GetVerticalCriterion *GetVerticalCriterion) (*GetVerticalCriterionResponse, error) {
	α := struct {
		OperationGetVerticalCriterionRequest `xml:"tns:getVerticalCriterion"`
	}{
		OperationGetVerticalCriterionRequest{
			GetVerticalCriterion,
		},
	}

	γ := struct {
		OperationGetVerticalCriterionResponse `xml:"getVerticalCriterionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetVerticalCriterion", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetVerticalCriterionResponse, nil
}
