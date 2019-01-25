package adwordsuserlist

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "https://adwords.google.com/api/adwords/rm/v201809"

// NewAdwordsUserListServiceInterface creates an initializes a AdwordsUserListServiceInterface.
func NewAdwordsUserListServiceInterface(cli *soap.Client) AdwordsUserListServiceInterface {
	return &adwordsUserListServiceInterface{cli}
}

// AdwordsUserListServiceInterface was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type AdwordsUserListServiceInterface interface {
	// Returns the list of user lists that meet the selector criteria.
	//                  @param serviceSelector the selector specifying
	// the {@link UserList}s to return.         @return a list of UserList
	// entities which meet the selector criteria.         @throws ApiException
	// if problems occurred while fetching UserList information.
	Get(Get *Get) (*GetResponse, error)

	// Applies a list of mutate operations (i.e. add, set):
	//           Add - creates a set of user lists         Set - updates
	// a set of user lists         Remove - not supported
	//         @param operations the operations to apply         @return
	// a list of UserList objects
	Mutate(Mutate *Mutate) (*MutateResponse, error)

	// Mutate members of user lists by either adding or removing their
	// lists of members.         The following {@link Operator}s are
	// supported: ADD and REMOVE. The SET operator         is not supported.
	//                  <p>Note that operations cannot have same user
	// list id but different operators.                  @param operations
	// the mutate members operations to apply         @return a list
	// of UserList objects         @throws ApiException when there
	// are one or more errors with the request
	MutateMembers(MutateMembers *MutateMembers) (*MutateMembersResponse, error)

	// Returns the list of user lists that match the query.
	//           @param query The SQL-like AWQL query string
	//   @return A list of UserList         @throws ApiException when
	// the query is invalid or there are errors processing the request.
	Query(Query string) (*QueryResponse, error)
}

// AccessReason was auto-generated from WSDL.
type AccessReason string

// Validate validates AccessReason.
func (v AccessReason) Validate() bool {
	for _, vv := range []string{
		"OWNED",
		"SHARED",
		"LICENSED",
		"SUBSCRIBED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AccountUserListStatus was auto-generated from WSDL.
type AccountUserListStatus string

// Validate validates AccountUserListStatus.
func (v AccountUserListStatus) Validate() bool {
	for _, vv := range []string{
		"ACTIVE",
		"INACTIVE",
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

// CombinedRuleUserListRuleOperator was auto-generated from WSDL.
type CombinedRuleUserListRuleOperator string

// Validate validates CombinedRuleUserListRuleOperator.
func (v CombinedRuleUserListRuleOperator) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"AND",
		"AND_NOT",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// CrmDataSourceType was auto-generated from WSDL.
type CrmDataSourceType string

// Validate validates CrmDataSourceType.
func (v CrmDataSourceType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"FIRST_PARTY",
		"THIRD_PARTY_CREDIT_BUREAU",
		"THIRD_PARTY_VOTER_FILE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// CustomerMatchUploadKeyType was auto-generated from WSDL.
type CustomerMatchUploadKeyType string

// Validate validates CustomerMatchUploadKeyType.
func (v CustomerMatchUploadKeyType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"CONTACT_INFO",
		"CRM_ID",
		"MOBILE_ADVERTISING_ID",
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

// DateRuleItemDateOperator was auto-generated from WSDL.
type DateRuleItemDateOperator string

// Validate validates DateRuleItemDateOperator.
func (v DateRuleItemDateOperator) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"EQUALS",
		"NOT_EQUAL",
		"BEFORE",
		"AFTER",
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

// MutateMembersErrorReason was auto-generated from WSDL.
type MutateMembersErrorReason string

// Validate validates MutateMembersErrorReason.
func (v MutateMembersErrorReason) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"UNSUPPORTED_METHOD",
		"INVALID_USER_LIST_ID",
		"INVALID_USER_LIST_TYPE",
		"INVALID_DATA_TYPE",
		"INVALID_SHA256_FORMAT",
		"OPERATOR_CONFLICT_FOR_SAME_USER_LIST_ID",
		"INVALID_REMOVEALL_OPERATION",
		"INVALID_OPERATION_ORDER",
		"MISSING_MEMBER_IDENTIFIER",
		"INCOMPATIBLE_UPLOAD_KEY_TYPE",
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

// NumberRuleItemNumberOperator was auto-generated from WSDL.
type NumberRuleItemNumberOperator string

// Validate validates NumberRuleItemNumberOperator.
func (v NumberRuleItemNumberOperator) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"GREATER_THAN",
		"GREATER_THAN_OR_EQUAL",
		"EQUALS",
		"NOT_EQUAL",
		"LESS_THAN",
		"LESS_THAN_OR_EQUAL",
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

// RuleBasedUserListPrepopulationStatus was auto-generated from
// WSDL.
type RuleBasedUserListPrepopulationStatus string

// Validate validates RuleBasedUserListPrepopulationStatus.
func (v RuleBasedUserListPrepopulationStatus) Validate() bool {
	for _, vv := range []string{
		"NONE",
		"REQUESTED",
		"FINISHED",
		"FAILED",
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

// SizeRange was auto-generated from WSDL.
type SizeRange string

// Validate validates SizeRange.
func (v SizeRange) Validate() bool {
	for _, vv := range []string{
		"LESS_THAN_FIVE_HUNDRED",
		"LESS_THAN_ONE_THOUSAND",
		"ONE_THOUSAND_TO_TEN_THOUSAND",
		"TEN_THOUSAND_TO_FIFTY_THOUSAND",
		"FIFTY_THOUSAND_TO_ONE_HUNDRED_THOUSAND",
		"ONE_HUNDRED_THOUSAND_TO_THREE_HUNDRED_THOUSAND",
		"THREE_HUNDRED_THOUSAND_TO_FIVE_HUNDRED_THOUSAND",
		"FIVE_HUNDRED_THOUSAND_TO_ONE_MILLION",
		"ONE_MILLION_TO_TWO_MILLION",
		"TWO_MILLION_TO_THREE_MILLION",
		"THREE_MILLION_TO_FIVE_MILLION",
		"FIVE_MILLION_TO_TEN_MILLION",
		"TEN_MILLION_TO_TWENTY_MILLION",
		"TWENTY_MILLION_TO_THIRTY_MILLION",
		"THIRTY_MILLION_TO_FIFTY_MILLION",
		"OVER_FIFTY_MILLION",
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

// StringRuleItemStringOperator was auto-generated from WSDL.
type StringRuleItemStringOperator string

// Validate validates StringRuleItemStringOperator.
func (v StringRuleItemStringOperator) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"CONTAINS",
		"EQUALS",
		"STARTS_WITH",
		"ENDS_WITH",
		"NOT_EQUAL",
		"NOT_CONTAIN",
		"NOT_START_WITH",
		"NOT_END_WITH",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// UserListClosingReason was auto-generated from WSDL.
type UserListClosingReason string

// Validate validates UserListClosingReason.
func (v UserListClosingReason) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"UNUSED_LIST",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// UserListConversionTypeCategory was auto-generated from WSDL.
type UserListConversionTypeCategory string

// Validate validates UserListConversionTypeCategory.
func (v UserListConversionTypeCategory) Validate() bool {
	for _, vv := range []string{
		"BOOMERANG_EVENT",
		"OTHER",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// UserListErrorReason was auto-generated from WSDL.
type UserListErrorReason string

// Validate validates UserListErrorReason.
func (v UserListErrorReason) Validate() bool {
	for _, vv := range []string{
		"EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED",
		"CONCRETE_TYPE_REQUIRED",
		"CONVERSION_TYPE_ID_REQUIRED",
		"DUPLICATE_CONVERSION_TYPES",
		"INVALID_CONVERSION_TYPE",
		"INVALID_DESCRIPTION",
		"INVALID_NAME",
		"INVALID_TYPE",
		"CAN_NOT_ADD_SIMILAR_LIST_AS_LOGICAL_LIST_NONE_OPERAND",
		"CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND",
		"INVALID_USER_LIST_LOGICAL_RULE_OPERAND",
		"NAME_ALREADY_USED",
		"NEW_CONVERSION_TYPE_NAME_REQUIRED",
		"CONVERSION_TYPE_NAME_ALREADY_USED",
		"OWNERSHIP_REQUIRED_FOR_SET",
		"REMOVE_NOT_SUPPORTED",
		"USER_LIST_MUTATE_NOT_SUPPORTED",
		"INVALID_RULE",
		"INVALID_DATE_RANGE",
		"CAN_NOT_MUTATE_SENSITIVE_USERLIST",
		"MAX_NUM_RULEBASED_USERLISTS",
		"CANNOT_MODIFY_BILLABLE_RECORD_COUNT",
		"APP_ID_NOT_ALLOWED",
		"APP_ID_NOT_SET",
		"USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST",
		"USER_LIST_SERVICE_ERROR",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// UserListLogicalRuleOperator was auto-generated from WSDL.
type UserListLogicalRuleOperator string

// Validate validates UserListLogicalRuleOperator.
func (v UserListLogicalRuleOperator) Validate() bool {
	for _, vv := range []string{
		"ALL",
		"ANY",
		"NONE",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// UserListMembershipStatus was auto-generated from WSDL.
type UserListMembershipStatus string

// Validate validates UserListMembershipStatus.
func (v UserListMembershipStatus) Validate() bool {
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

// UserListRuleTypeEnumsEnum was auto-generated from WSDL.
type UserListRuleTypeEnumsEnum string

// Validate validates UserListRuleTypeEnumsEnum.
func (v UserListRuleTypeEnumsEnum) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"CNF",
		"DNF",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// UserListType was auto-generated from WSDL.
type UserListType string

// Validate validates UserListType.
func (v UserListType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"REMARKETING",
		"LOGICAL",
		"EXTERNAL_REMARKETING",
		"RULE_BASED",
		"SIMILAR",
		"CRM_BASED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// UserListUploadStatus was auto-generated from WSDL.
type UserListUploadStatus string

// Validate validates UserListUploadStatus.
func (v UserListUploadStatus) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"IN_PROCESS",
		"SUCCESS",
		"FAILURE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// Address identifier of a user list member. Accessible for whitelisted
// customers only.
type AddressInfo struct {
	HashedFirstName *string `xml:"hashedFirstName,omitempty" json:"hashedFirstName,omitempty" yaml:"hashedFirstName,omitempty"`
	HashedLastName  *string `xml:"hashedLastName,omitempty" json:"hashedLastName,omitempty" yaml:"hashedLastName,omitempty"`
	CountryCode     *string `xml:"countryCode,omitempty" json:"countryCode,omitempty" yaml:"countryCode,omitempty"`
	ZipCode         *string `xml:"zipCode,omitempty" json:"zipCode,omitempty" yaml:"zipCode,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// User list targeting as a collection of conversion types.
type BasicUserList struct {
	Id                    *int64                    `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	IsReadOnly            *bool                     `xml:"isReadOnly,omitempty" json:"isReadOnly,omitempty" yaml:"isReadOnly,omitempty"`
	Name                  *string                   `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description           *string                   `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Status                *UserListMembershipStatus `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	IntegrationCode       *string                   `xml:"integrationCode,omitempty" json:"integrationCode,omitempty" yaml:"integrationCode,omitempty"`
	AccessReason          *AccessReason             `xml:"accessReason,omitempty" json:"accessReason,omitempty" yaml:"accessReason,omitempty"`
	AccountUserListStatus *AccountUserListStatus    `xml:"accountUserListStatus,omitempty" json:"accountUserListStatus,omitempty" yaml:"accountUserListStatus,omitempty"`
	MembershipLifeSpan    *int64                    `xml:"membershipLifeSpan,omitempty" json:"membershipLifeSpan,omitempty" yaml:"membershipLifeSpan,omitempty"`
	Size                  *int64                    `xml:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	SizeRange             *SizeRange                `xml:"sizeRange,omitempty" json:"sizeRange,omitempty" yaml:"sizeRange,omitempty"`
	SizeForSearch         *int64                    `xml:"sizeForSearch,omitempty" json:"sizeForSearch,omitempty" yaml:"sizeForSearch,omitempty"`
	SizeRangeForSearch    *SizeRange                `xml:"sizeRangeForSearch,omitempty" json:"sizeRangeForSearch,omitempty" yaml:"sizeRangeForSearch,omitempty"`
	ListType              *UserListType             `xml:"listType,omitempty" json:"listType,omitempty" yaml:"listType,omitempty"`
	IsEligibleForSearch   *bool                     `xml:"isEligibleForSearch,omitempty" json:"isEligibleForSearch,omitempty" yaml:"isEligibleForSearch,omitempty"`
	IsEligibleForDisplay  *bool                     `xml:"isEligibleForDisplay,omitempty" json:"isEligibleForDisplay,omitempty" yaml:"isEligibleForDisplay,omitempty"`
	ClosingReason         *UserListClosingReason    `xml:"closingReason,omitempty" json:"closingReason,omitempty" yaml:"closingReason,omitempty"`
	UserListType          *string                   `xml:"UserList.Type,omitempty" json:"UserList.Type,omitempty" yaml:"UserList.Type,omitempty"`
	ConversionTypes       []*UserListConversionType `xml:"conversionTypes,omitempty" json:"conversionTypes,omitempty" yaml:"conversionTypes,omitempty"`
	TypeAttrXSI           string                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace         string                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *BasicUserList) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:BasicUserList"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// Users defined by combining two rules, left operand and right
// operand. There are two operators:             AND where left
// operand and right operand have to be true; AND_NOT where left
// operand is true but             right operand is false.
type CombinedRuleUserList struct {
	Id                    *int64                                `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	IsReadOnly            *bool                                 `xml:"isReadOnly,omitempty" json:"isReadOnly,omitempty" yaml:"isReadOnly,omitempty"`
	Name                  *string                               `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description           *string                               `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Status                *UserListMembershipStatus             `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	IntegrationCode       *string                               `xml:"integrationCode,omitempty" json:"integrationCode,omitempty" yaml:"integrationCode,omitempty"`
	AccessReason          *AccessReason                         `xml:"accessReason,omitempty" json:"accessReason,omitempty" yaml:"accessReason,omitempty"`
	AccountUserListStatus *AccountUserListStatus                `xml:"accountUserListStatus,omitempty" json:"accountUserListStatus,omitempty" yaml:"accountUserListStatus,omitempty"`
	MembershipLifeSpan    *int64                                `xml:"membershipLifeSpan,omitempty" json:"membershipLifeSpan,omitempty" yaml:"membershipLifeSpan,omitempty"`
	Size                  *int64                                `xml:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	SizeRange             *SizeRange                            `xml:"sizeRange,omitempty" json:"sizeRange,omitempty" yaml:"sizeRange,omitempty"`
	SizeForSearch         *int64                                `xml:"sizeForSearch,omitempty" json:"sizeForSearch,omitempty" yaml:"sizeForSearch,omitempty"`
	SizeRangeForSearch    *SizeRange                            `xml:"sizeRangeForSearch,omitempty" json:"sizeRangeForSearch,omitempty" yaml:"sizeRangeForSearch,omitempty"`
	ListType              *UserListType                         `xml:"listType,omitempty" json:"listType,omitempty" yaml:"listType,omitempty"`
	IsEligibleForSearch   *bool                                 `xml:"isEligibleForSearch,omitempty" json:"isEligibleForSearch,omitempty" yaml:"isEligibleForSearch,omitempty"`
	IsEligibleForDisplay  *bool                                 `xml:"isEligibleForDisplay,omitempty" json:"isEligibleForDisplay,omitempty" yaml:"isEligibleForDisplay,omitempty"`
	ClosingReason         *UserListClosingReason                `xml:"closingReason,omitempty" json:"closingReason,omitempty" yaml:"closingReason,omitempty"`
	UserListType          *string                               `xml:"UserList.Type,omitempty" json:"UserList.Type,omitempty" yaml:"UserList.Type,omitempty"`
	PrepopulationStatus   *RuleBasedUserListPrepopulationStatus `xml:"prepopulationStatus,omitempty" json:"prepopulationStatus,omitempty" yaml:"prepopulationStatus,omitempty"`
	LeftOperand           *Rule                                 `xml:"leftOperand,omitempty" json:"leftOperand,omitempty" yaml:"leftOperand,omitempty"`
	RightOperand          *Rule                                 `xml:"rightOperand,omitempty" json:"rightOperand,omitempty" yaml:"rightOperand,omitempty"`
	RuleOperator          *CombinedRuleUserListRuleOperator     `xml:"ruleOperator,omitempty" json:"ruleOperator,omitempty" yaml:"ruleOperator,omitempty"`
	TypeAttrXSI           string                                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace         string                                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CombinedRuleUserList) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CombinedRuleUserList"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// UserList of CRM users provided by the advertiser.
type CrmBasedUserList struct {
	Id                    *int64                      `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	IsReadOnly            *bool                       `xml:"isReadOnly,omitempty" json:"isReadOnly,omitempty" yaml:"isReadOnly,omitempty"`
	Name                  *string                     `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description           *string                     `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Status                *UserListMembershipStatus   `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	IntegrationCode       *string                     `xml:"integrationCode,omitempty" json:"integrationCode,omitempty" yaml:"integrationCode,omitempty"`
	AccessReason          *AccessReason               `xml:"accessReason,omitempty" json:"accessReason,omitempty" yaml:"accessReason,omitempty"`
	AccountUserListStatus *AccountUserListStatus      `xml:"accountUserListStatus,omitempty" json:"accountUserListStatus,omitempty" yaml:"accountUserListStatus,omitempty"`
	MembershipLifeSpan    *int64                      `xml:"membershipLifeSpan,omitempty" json:"membershipLifeSpan,omitempty" yaml:"membershipLifeSpan,omitempty"`
	Size                  *int64                      `xml:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	SizeRange             *SizeRange                  `xml:"sizeRange,omitempty" json:"sizeRange,omitempty" yaml:"sizeRange,omitempty"`
	SizeForSearch         *int64                      `xml:"sizeForSearch,omitempty" json:"sizeForSearch,omitempty" yaml:"sizeForSearch,omitempty"`
	SizeRangeForSearch    *SizeRange                  `xml:"sizeRangeForSearch,omitempty" json:"sizeRangeForSearch,omitempty" yaml:"sizeRangeForSearch,omitempty"`
	ListType              *UserListType               `xml:"listType,omitempty" json:"listType,omitempty" yaml:"listType,omitempty"`
	IsEligibleForSearch   *bool                       `xml:"isEligibleForSearch,omitempty" json:"isEligibleForSearch,omitempty" yaml:"isEligibleForSearch,omitempty"`
	IsEligibleForDisplay  *bool                       `xml:"isEligibleForDisplay,omitempty" json:"isEligibleForDisplay,omitempty" yaml:"isEligibleForDisplay,omitempty"`
	ClosingReason         *UserListClosingReason      `xml:"closingReason,omitempty" json:"closingReason,omitempty" yaml:"closingReason,omitempty"`
	UserListType          *string                     `xml:"UserList.Type,omitempty" json:"UserList.Type,omitempty" yaml:"UserList.Type,omitempty"`
	AppId                 *string                     `xml:"appId,omitempty" json:"appId,omitempty" yaml:"appId,omitempty"`
	UploadKeyType         *CustomerMatchUploadKeyType `xml:"uploadKeyType,omitempty" json:"uploadKeyType,omitempty" yaml:"uploadKeyType,omitempty"`
	DataSourceType        *CrmDataSourceType          `xml:"dataSourceType,omitempty" json:"dataSourceType,omitempty" yaml:"dataSourceType,omitempty"`
	DataUploadResult      *DataUploadResult           `xml:"dataUploadResult,omitempty" json:"dataUploadResult,omitempty" yaml:"dataUploadResult,omitempty"`
	TypeAttrXSI           string                      `xml:"xsi:type,attr,omitempty"`
	TypeNamespace         string                      `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CrmBasedUserList) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CrmBasedUserList"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// A class represents the data upload result for CRM based lists.
type DataUploadResult struct {
	UploadStatus    *UserListUploadStatus `xml:"uploadStatus,omitempty" json:"uploadStatus,omitempty" yaml:"uploadStatus,omitempty"`
	RemoveAllStatus *UserListUploadStatus `xml:"removeAllStatus,omitempty" json:"removeAllStatus,omitempty" yaml:"removeAllStatus,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// A custom parameter of date type. Supported date formats are
// listed as follows:             <ul>             <li> 2011-03-31T12:20:19-05:00
//             <li> 03/31/2011 12:20:19-05:00             <li>
// Fri, Mar 31 2011 12:20:19 EST             <li> Fri, Mar 31 12:20:19
// EST 2011             </ul>             <p>             If time
// zone information is not present in the value,             it
// is assumed to be PST. If time value is not specified,
//       it defaults to midnight of the day.
type DateKey struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// Represents a range of dates that has either an upper or a lower
// bound.             The format for the date is YYYYMMDD.
type DateRange struct {
	Min *string `xml:"min,omitempty" json:"min,omitempty" yaml:"min,omitempty"`
	Max *string `xml:"max,omitempty" json:"max,omitempty" yaml:"max,omitempty"`
}

// An atomic rule fragment composing of date operation.
type DateRuleItem struct {
	Key           *DateKey                  `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Op            *DateRuleItemDateOperator `xml:"op,omitempty" json:"op,omitempty" yaml:"op,omitempty"`
	Value         *string                   `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	RelativeValue *RelativeDate             `xml:"relativeValue,omitempty" json:"relativeValue,omitempty" yaml:"relativeValue,omitempty"`
}

// Visitors of a page during specific dates. The visiting periods
// are defined as follows:             <ul>             <li> between
// {@code startDate} (inclusive) and {@code endDate} (inclusive);
//             <li> before {@code endDate} (exclusive) with {@code
// startDate} = 2000-01-01;             <li> after {@code startDate}
// (exclusive) with {@code endDate} = 2037-12-30.             </ul>
type DateSpecificRuleUserList struct {
	Id                    *int64                                `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	IsReadOnly            *bool                                 `xml:"isReadOnly,omitempty" json:"isReadOnly,omitempty" yaml:"isReadOnly,omitempty"`
	Name                  *string                               `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description           *string                               `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Status                *UserListMembershipStatus             `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	IntegrationCode       *string                               `xml:"integrationCode,omitempty" json:"integrationCode,omitempty" yaml:"integrationCode,omitempty"`
	AccessReason          *AccessReason                         `xml:"accessReason,omitempty" json:"accessReason,omitempty" yaml:"accessReason,omitempty"`
	AccountUserListStatus *AccountUserListStatus                `xml:"accountUserListStatus,omitempty" json:"accountUserListStatus,omitempty" yaml:"accountUserListStatus,omitempty"`
	MembershipLifeSpan    *int64                                `xml:"membershipLifeSpan,omitempty" json:"membershipLifeSpan,omitempty" yaml:"membershipLifeSpan,omitempty"`
	Size                  *int64                                `xml:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	SizeRange             *SizeRange                            `xml:"sizeRange,omitempty" json:"sizeRange,omitempty" yaml:"sizeRange,omitempty"`
	SizeForSearch         *int64                                `xml:"sizeForSearch,omitempty" json:"sizeForSearch,omitempty" yaml:"sizeForSearch,omitempty"`
	SizeRangeForSearch    *SizeRange                            `xml:"sizeRangeForSearch,omitempty" json:"sizeRangeForSearch,omitempty" yaml:"sizeRangeForSearch,omitempty"`
	ListType              *UserListType                         `xml:"listType,omitempty" json:"listType,omitempty" yaml:"listType,omitempty"`
	IsEligibleForSearch   *bool                                 `xml:"isEligibleForSearch,omitempty" json:"isEligibleForSearch,omitempty" yaml:"isEligibleForSearch,omitempty"`
	IsEligibleForDisplay  *bool                                 `xml:"isEligibleForDisplay,omitempty" json:"isEligibleForDisplay,omitempty" yaml:"isEligibleForDisplay,omitempty"`
	ClosingReason         *UserListClosingReason                `xml:"closingReason,omitempty" json:"closingReason,omitempty" yaml:"closingReason,omitempty"`
	UserListType          *string                               `xml:"UserList.Type,omitempty" json:"UserList.Type,omitempty" yaml:"UserList.Type,omitempty"`
	PrepopulationStatus   *RuleBasedUserListPrepopulationStatus `xml:"prepopulationStatus,omitempty" json:"prepopulationStatus,omitempty" yaml:"prepopulationStatus,omitempty"`
	Rule                  *Rule                                 `xml:"rule,omitempty" json:"rule,omitempty" yaml:"rule,omitempty"`
	StartDate             *string                               `xml:"startDate,omitempty" json:"startDate,omitempty" yaml:"startDate,omitempty"`
	EndDate               *string                               `xml:"endDate,omitempty" json:"endDate,omitempty" yaml:"endDate,omitempty"`
	TypeAttrXSI           string                                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace         string                                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DateSpecificRuleUserList) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DateSpecificRuleUserList"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// Visitors of a page. The page visit is defined by one boolean
// rule expression.
type ExpressionRuleUserList struct {
	Id                    *int64                                `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	IsReadOnly            *bool                                 `xml:"isReadOnly,omitempty" json:"isReadOnly,omitempty" yaml:"isReadOnly,omitempty"`
	Name                  *string                               `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description           *string                               `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Status                *UserListMembershipStatus             `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	IntegrationCode       *string                               `xml:"integrationCode,omitempty" json:"integrationCode,omitempty" yaml:"integrationCode,omitempty"`
	AccessReason          *AccessReason                         `xml:"accessReason,omitempty" json:"accessReason,omitempty" yaml:"accessReason,omitempty"`
	AccountUserListStatus *AccountUserListStatus                `xml:"accountUserListStatus,omitempty" json:"accountUserListStatus,omitempty" yaml:"accountUserListStatus,omitempty"`
	MembershipLifeSpan    *int64                                `xml:"membershipLifeSpan,omitempty" json:"membershipLifeSpan,omitempty" yaml:"membershipLifeSpan,omitempty"`
	Size                  *int64                                `xml:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	SizeRange             *SizeRange                            `xml:"sizeRange,omitempty" json:"sizeRange,omitempty" yaml:"sizeRange,omitempty"`
	SizeForSearch         *int64                                `xml:"sizeForSearch,omitempty" json:"sizeForSearch,omitempty" yaml:"sizeForSearch,omitempty"`
	SizeRangeForSearch    *SizeRange                            `xml:"sizeRangeForSearch,omitempty" json:"sizeRangeForSearch,omitempty" yaml:"sizeRangeForSearch,omitempty"`
	ListType              *UserListType                         `xml:"listType,omitempty" json:"listType,omitempty" yaml:"listType,omitempty"`
	IsEligibleForSearch   *bool                                 `xml:"isEligibleForSearch,omitempty" json:"isEligibleForSearch,omitempty" yaml:"isEligibleForSearch,omitempty"`
	IsEligibleForDisplay  *bool                                 `xml:"isEligibleForDisplay,omitempty" json:"isEligibleForDisplay,omitempty" yaml:"isEligibleForDisplay,omitempty"`
	ClosingReason         *UserListClosingReason                `xml:"closingReason,omitempty" json:"closingReason,omitempty" yaml:"closingReason,omitempty"`
	UserListType          *string                               `xml:"UserList.Type,omitempty" json:"UserList.Type,omitempty" yaml:"UserList.Type,omitempty"`
	PrepopulationStatus   *RuleBasedUserListPrepopulationStatus `xml:"prepopulationStatus,omitempty" json:"prepopulationStatus,omitempty" yaml:"prepopulationStatus,omitempty"`
	Rule                  *Rule                                 `xml:"rule,omitempty" json:"rule,omitempty" yaml:"rule,omitempty"`
	TypeAttrXSI           string                                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace         string                                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ExpressionRuleUserList) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ExpressionRuleUserList"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// A segment of a field path. Each dot in a field path defines
// a new segment.
type FieldPathElement struct {
	Field *string `xml:"field,omitempty" json:"field,omitempty" yaml:"field,omitempty"`
	Index *int    `xml:"index,omitempty" json:"index,omitempty" yaml:"index,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// Base list return value type.
type ListReturnValue interface{}

// Represents a user list that is a custom combination of user
// lists and user             interests.
type LogicalUserList struct {
	Id                    *int64                    `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	IsReadOnly            *bool                     `xml:"isReadOnly,omitempty" json:"isReadOnly,omitempty" yaml:"isReadOnly,omitempty"`
	Name                  *string                   `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description           *string                   `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Status                *UserListMembershipStatus `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	IntegrationCode       *string                   `xml:"integrationCode,omitempty" json:"integrationCode,omitempty" yaml:"integrationCode,omitempty"`
	AccessReason          *AccessReason             `xml:"accessReason,omitempty" json:"accessReason,omitempty" yaml:"accessReason,omitempty"`
	AccountUserListStatus *AccountUserListStatus    `xml:"accountUserListStatus,omitempty" json:"accountUserListStatus,omitempty" yaml:"accountUserListStatus,omitempty"`
	MembershipLifeSpan    *int64                    `xml:"membershipLifeSpan,omitempty" json:"membershipLifeSpan,omitempty" yaml:"membershipLifeSpan,omitempty"`
	Size                  *int64                    `xml:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	SizeRange             *SizeRange                `xml:"sizeRange,omitempty" json:"sizeRange,omitempty" yaml:"sizeRange,omitempty"`
	SizeForSearch         *int64                    `xml:"sizeForSearch,omitempty" json:"sizeForSearch,omitempty" yaml:"sizeForSearch,omitempty"`
	SizeRangeForSearch    *SizeRange                `xml:"sizeRangeForSearch,omitempty" json:"sizeRangeForSearch,omitempty" yaml:"sizeRangeForSearch,omitempty"`
	ListType              *UserListType             `xml:"listType,omitempty" json:"listType,omitempty" yaml:"listType,omitempty"`
	IsEligibleForSearch   *bool                     `xml:"isEligibleForSearch,omitempty" json:"isEligibleForSearch,omitempty" yaml:"isEligibleForSearch,omitempty"`
	IsEligibleForDisplay  *bool                     `xml:"isEligibleForDisplay,omitempty" json:"isEligibleForDisplay,omitempty" yaml:"isEligibleForDisplay,omitempty"`
	ClosingReason         *UserListClosingReason    `xml:"closingReason,omitempty" json:"closingReason,omitempty" yaml:"closingReason,omitempty"`
	UserListType          *string                   `xml:"UserList.Type,omitempty" json:"UserList.Type,omitempty" yaml:"UserList.Type,omitempty"`
	Rules                 []*UserListLogicalRule    `xml:"rules,omitempty" json:"rules,omitempty" yaml:"rules,omitempty"`
	TypeAttrXSI           string                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace         string                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *LogicalUserList) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:LogicalUserList"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// An interface for a logical user list operand. A logical user
// list is a             combination of logical rules. Each rule
// is defined as a logical operator and             a list of operands.
// Those operands can be of type UserList.
type LogicalUserListOperand struct {
	UserList *UserList `xml:"UserList" json:"UserList" yaml:"UserList"`
}

// Class that holds user list member identifiers. The following
// types of member             identifier are supported:
//                    <ul>             <li>Contact info (email,
// phone number, address)             <li>Mobile advertising ID
//             <li>User IDs generated and assigned by advertiser
//             </ul>                          A list can be uploaded
// with only one type of data and once uploaded will not
//       accept any other ID types.
type Member struct {
	HashedEmail       *string      `xml:"hashedEmail,omitempty" json:"hashedEmail,omitempty" yaml:"hashedEmail,omitempty"`
	MobileId          *string      `xml:"mobileId,omitempty" json:"mobileId,omitempty" yaml:"mobileId,omitempty"`
	HashedPhoneNumber *string      `xml:"hashedPhoneNumber,omitempty" json:"hashedPhoneNumber,omitempty" yaml:"hashedPhoneNumber,omitempty"`
	AddressInfo       *AddressInfo `xml:"addressInfo,omitempty" json:"addressInfo,omitempty" yaml:"addressInfo,omitempty"`
	UserId            *string      `xml:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty"`
}

// Represents possible error codes from {@code UserListService#mutateMembers}.
type MutateMembersError struct {
	FieldPath         *string                   `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement       `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string                   `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string                   `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string                   `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *MutateMembersErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MutateMembersError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MutateMembersError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// Operand containing user list id and members to be added or removed.
type MutateMembersOperand struct {
	UserListId  *int64    `xml:"userListId,omitempty" json:"userListId,omitempty" yaml:"userListId,omitempty"`
	RemoveAll   *bool     `xml:"removeAll,omitempty" json:"removeAll,omitempty" yaml:"removeAll,omitempty"`
	MembersList []*Member `xml:"membersList,omitempty" json:"membersList,omitempty" yaml:"membersList,omitempty"`
}

// Operation representing a request to add or remove members from
// a user list.             The following {@link Operator}s are
// supported: ADD and REMOVE. The SET operator             is not
// supported.
type MutateMembersOperation struct {
	Operator      *Operator             `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	OperationType *string               `xml:"Operation.Type,omitempty" json:"Operation.Type,omitempty" yaml:"Operation.Type,omitempty"`
	Operand       *MutateMembersOperand `xml:"operand,omitempty" json:"operand,omitempty" yaml:"operand,omitempty"`
	TypeAttrXSI   string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *MutateMembersOperation) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:MutateMembersOperation"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// A container for return value from {@code UserListService#mutateMembers}
//             method.
type MutateMembersReturnValue struct {
	UserLists []*UserList `xml:"userLists,omitempty" json:"userLists,omitempty" yaml:"userLists,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// A custom parameter of type number.
type NumberKey struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// An atomic rule fragment composing of number operation.
type NumberRuleItem struct {
	Key   *NumberKey                    `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Op    *NumberRuleItemNumberOperator `xml:"op,omitempty" json:"op,omitempty" yaml:"op,omitempty"`
	Value *float64                      `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// Date relative to NOW (the current date).
type RelativeDate struct {
	OffsetInDays *int `xml:"offsetInDays,omitempty" json:"offsetInDays,omitempty" yaml:"offsetInDays,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// A client defined rule based on custom parameters sent by web
// sites.
type Rule struct {
	Groups   []*RuleItemGroup           `xml:"groups,omitempty" json:"groups,omitempty" yaml:"groups,omitempty"`
	RuleType *UserListRuleTypeEnumsEnum `xml:"ruleType,omitempty" json:"ruleType,omitempty" yaml:"ruleType,omitempty"`
}

// Representation of a userlist that is generated by a rule.
type RuleBasedUserList struct {
	Id                    *int64                                `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	IsReadOnly            *bool                                 `xml:"isReadOnly,omitempty" json:"isReadOnly,omitempty" yaml:"isReadOnly,omitempty"`
	Name                  *string                               `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description           *string                               `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Status                *UserListMembershipStatus             `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	IntegrationCode       *string                               `xml:"integrationCode,omitempty" json:"integrationCode,omitempty" yaml:"integrationCode,omitempty"`
	AccessReason          *AccessReason                         `xml:"accessReason,omitempty" json:"accessReason,omitempty" yaml:"accessReason,omitempty"`
	AccountUserListStatus *AccountUserListStatus                `xml:"accountUserListStatus,omitempty" json:"accountUserListStatus,omitempty" yaml:"accountUserListStatus,omitempty"`
	MembershipLifeSpan    *int64                                `xml:"membershipLifeSpan,omitempty" json:"membershipLifeSpan,omitempty" yaml:"membershipLifeSpan,omitempty"`
	Size                  *int64                                `xml:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	SizeRange             *SizeRange                            `xml:"sizeRange,omitempty" json:"sizeRange,omitempty" yaml:"sizeRange,omitempty"`
	SizeForSearch         *int64                                `xml:"sizeForSearch,omitempty" json:"sizeForSearch,omitempty" yaml:"sizeForSearch,omitempty"`
	SizeRangeForSearch    *SizeRange                            `xml:"sizeRangeForSearch,omitempty" json:"sizeRangeForSearch,omitempty" yaml:"sizeRangeForSearch,omitempty"`
	ListType              *UserListType                         `xml:"listType,omitempty" json:"listType,omitempty" yaml:"listType,omitempty"`
	IsEligibleForSearch   *bool                                 `xml:"isEligibleForSearch,omitempty" json:"isEligibleForSearch,omitempty" yaml:"isEligibleForSearch,omitempty"`
	IsEligibleForDisplay  *bool                                 `xml:"isEligibleForDisplay,omitempty" json:"isEligibleForDisplay,omitempty" yaml:"isEligibleForDisplay,omitempty"`
	ClosingReason         *UserListClosingReason                `xml:"closingReason,omitempty" json:"closingReason,omitempty" yaml:"closingReason,omitempty"`
	UserListType          *string                               `xml:"UserList.Type,omitempty" json:"UserList.Type,omitempty" yaml:"UserList.Type,omitempty"`
	PrepopulationStatus   *RuleBasedUserListPrepopulationStatus `xml:"prepopulationStatus,omitempty" json:"prepopulationStatus,omitempty" yaml:"prepopulationStatus,omitempty"`
	TypeAttrXSI           string                                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace         string                                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *RuleBasedUserList) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:RuleBasedUserList"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// An atomic rule fragment.
type RuleItem struct {
	DateRuleItem   *DateRuleItem   `xml:"DateRuleItem" json:"DateRuleItem" yaml:"DateRuleItem"`
	NumberRuleItem *NumberRuleItem `xml:"NumberRuleItem" json:"NumberRuleItem" yaml:"NumberRuleItem"`
	StringRuleItem *StringRuleItem `xml:"StringRuleItem" json:"StringRuleItem" yaml:"StringRuleItem"`
}

// A group of rule items that are ANDed together before version
// V201705.             Starting from version V201705, rule item
// groups will be grouped together based on             {@link
// Rule#getRuleType()}.
type RuleItemGroup struct {
	Items []*RuleItem `xml:"items,omitempty" json:"items,omitempty" yaml:"items,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// SimilarUserList is a list of users which are similar to users
// from another UserList.             These lists are readonly
// and automatically created by google.
type SimilarUserList struct {
	Id                      *int64                    `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	IsReadOnly              *bool                     `xml:"isReadOnly,omitempty" json:"isReadOnly,omitempty" yaml:"isReadOnly,omitempty"`
	Name                    *string                   `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description             *string                   `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Status                  *UserListMembershipStatus `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	IntegrationCode         *string                   `xml:"integrationCode,omitempty" json:"integrationCode,omitempty" yaml:"integrationCode,omitempty"`
	AccessReason            *AccessReason             `xml:"accessReason,omitempty" json:"accessReason,omitempty" yaml:"accessReason,omitempty"`
	AccountUserListStatus   *AccountUserListStatus    `xml:"accountUserListStatus,omitempty" json:"accountUserListStatus,omitempty" yaml:"accountUserListStatus,omitempty"`
	MembershipLifeSpan      *int64                    `xml:"membershipLifeSpan,omitempty" json:"membershipLifeSpan,omitempty" yaml:"membershipLifeSpan,omitempty"`
	Size                    *int64                    `xml:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	SizeRange               *SizeRange                `xml:"sizeRange,omitempty" json:"sizeRange,omitempty" yaml:"sizeRange,omitempty"`
	SizeForSearch           *int64                    `xml:"sizeForSearch,omitempty" json:"sizeForSearch,omitempty" yaml:"sizeForSearch,omitempty"`
	SizeRangeForSearch      *SizeRange                `xml:"sizeRangeForSearch,omitempty" json:"sizeRangeForSearch,omitempty" yaml:"sizeRangeForSearch,omitempty"`
	ListType                *UserListType             `xml:"listType,omitempty" json:"listType,omitempty" yaml:"listType,omitempty"`
	IsEligibleForSearch     *bool                     `xml:"isEligibleForSearch,omitempty" json:"isEligibleForSearch,omitempty" yaml:"isEligibleForSearch,omitempty"`
	IsEligibleForDisplay    *bool                     `xml:"isEligibleForDisplay,omitempty" json:"isEligibleForDisplay,omitempty" yaml:"isEligibleForDisplay,omitempty"`
	ClosingReason           *UserListClosingReason    `xml:"closingReason,omitempty" json:"closingReason,omitempty" yaml:"closingReason,omitempty"`
	UserListType            *string                   `xml:"UserList.Type,omitempty" json:"UserList.Type,omitempty" yaml:"UserList.Type,omitempty"`
	SeedUserListId          *int64                    `xml:"seedUserListId,omitempty" json:"seedUserListId,omitempty" yaml:"seedUserListId,omitempty"`
	SeedUserListName        *string                   `xml:"seedUserListName,omitempty" json:"seedUserListName,omitempty" yaml:"seedUserListName,omitempty"`
	SeedUserListDescription *string                   `xml:"seedUserListDescription,omitempty" json:"seedUserListDescription,omitempty" yaml:"seedUserListDescription,omitempty"`
	SeedUserListStatus      *UserListMembershipStatus `xml:"seedUserListStatus,omitempty" json:"seedUserListStatus,omitempty" yaml:"seedUserListStatus,omitempty"`
	SeedListSize            *int64                    `xml:"seedListSize,omitempty" json:"seedListSize,omitempty" yaml:"seedListSize,omitempty"`
	TypeAttrXSI             string                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace           string                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *SimilarUserList) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SimilarUserList"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// Custom parameter of type string. For websites, there are two
// built-in parameters             URL (name = 'url__') and referrer
// URL (name = 'ref_url__').
type StringKey struct {
	Name *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// An atomic rule fragment composing of string operation.
type StringRuleItem struct {
	Key   *StringKey                    `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Op    *StringRuleItemStringOperator `xml:"op,omitempty" json:"op,omitempty" yaml:"op,omitempty"`
	Value *string                       `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// Represents a UserList object that is sent over the wire.
//          This is a list of users an account may target.
type UserList struct {
	Id                    *int64                    `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	IsReadOnly            *bool                     `xml:"isReadOnly,omitempty" json:"isReadOnly,omitempty" yaml:"isReadOnly,omitempty"`
	Name                  *string                   `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description           *string                   `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Status                *UserListMembershipStatus `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	IntegrationCode       *string                   `xml:"integrationCode,omitempty" json:"integrationCode,omitempty" yaml:"integrationCode,omitempty"`
	AccessReason          *AccessReason             `xml:"accessReason,omitempty" json:"accessReason,omitempty" yaml:"accessReason,omitempty"`
	AccountUserListStatus *AccountUserListStatus    `xml:"accountUserListStatus,omitempty" json:"accountUserListStatus,omitempty" yaml:"accountUserListStatus,omitempty"`
	MembershipLifeSpan    *int64                    `xml:"membershipLifeSpan,omitempty" json:"membershipLifeSpan,omitempty" yaml:"membershipLifeSpan,omitempty"`
	Size                  *int64                    `xml:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	SizeRange             *SizeRange                `xml:"sizeRange,omitempty" json:"sizeRange,omitempty" yaml:"sizeRange,omitempty"`
	SizeForSearch         *int64                    `xml:"sizeForSearch,omitempty" json:"sizeForSearch,omitempty" yaml:"sizeForSearch,omitempty"`
	SizeRangeForSearch    *SizeRange                `xml:"sizeRangeForSearch,omitempty" json:"sizeRangeForSearch,omitempty" yaml:"sizeRangeForSearch,omitempty"`
	ListType              *UserListType             `xml:"listType,omitempty" json:"listType,omitempty" yaml:"listType,omitempty"`
	IsEligibleForSearch   *bool                     `xml:"isEligibleForSearch,omitempty" json:"isEligibleForSearch,omitempty" yaml:"isEligibleForSearch,omitempty"`
	IsEligibleForDisplay  *bool                     `xml:"isEligibleForDisplay,omitempty" json:"isEligibleForDisplay,omitempty" yaml:"isEligibleForDisplay,omitempty"`
	ClosingReason         *UserListClosingReason    `xml:"closingReason,omitempty" json:"closingReason,omitempty" yaml:"closingReason,omitempty"`
	UserListType          *string                   `xml:"UserList.Type,omitempty" json:"UserList.Type,omitempty" yaml:"UserList.Type,omitempty"`
}

// Represents a conversion type used for building remarketing user
// lists.
type UserListConversionType struct {
	Id       *int64                          `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Name     *string                         `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Category *UserListConversionTypeCategory `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
}

// Represents possible error codes in UserListService.
type UserListError struct {
	FieldPath         *string              `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement  `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string              `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string              `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string              `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *UserListErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *UserListError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UserListError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// A user list logical rule. A rule has a logical operator (and/or/not)
// and a             list of operands that can be user lists or
// user interests.
type UserListLogicalRule struct {
	Operator     *UserListLogicalRuleOperator `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	RuleOperands []*LogicalUserListOperand    `xml:"ruleOperands,omitempty" json:"ruleOperands,omitempty" yaml:"ruleOperands,omitempty"`
}

// UserList operations for adding/updating UserList entities.
//            The following {@link Operator}s are supported: ADD
// and SET.             The REMOVE operator is not supported.
type UserListOperation struct {
	Operator      *Operator `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	OperationType *string   `xml:"Operation.Type,omitempty" json:"Operation.Type,omitempty" yaml:"Operation.Type,omitempty"`
	Operand       *UserList `xml:"operand,omitempty" json:"operand,omitempty" yaml:"operand,omitempty"`
	TypeAttrXSI   string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *UserListOperation) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UserListOperation"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// Contains a list of user lists resulting from the filtering and
// paging of the             {@link UserListService#get} call.
type UserListPage struct {
	TotalNumEntries *int        `xml:"totalNumEntries,omitempty" json:"totalNumEntries,omitempty" yaml:"totalNumEntries,omitempty"`
	PageType        *string     `xml:"Page.Type,omitempty" json:"Page.Type,omitempty" yaml:"Page.Type,omitempty"`
	Entries         []*UserList `xml:"entries,omitempty" json:"entries,omitempty" yaml:"entries,omitempty"`
	TypeAttrXSI     string      `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string      `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *UserListPage) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UserListPage"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// A container for return values from the UserListService.
type UserListReturnValue struct {
	ListReturnValueType *string     `xml:"ListReturnValue.Type,omitempty" json:"ListReturnValue.Type,omitempty" yaml:"ListReturnValue.Type,omitempty"`
	Value               []*UserList `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	TypeAttrXSI         string      `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string      `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *UserListReturnValue) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UserListReturnValue"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/rm/v201809"
	}
}

// Get was auto-generated from WSDL.
type Get struct {
	ServiceSelector *Selector `xml:"serviceSelector,omitempty" json:"serviceSelector,omitempty" yaml:"serviceSelector,omitempty"`
}

// GetResponse was auto-generated from WSDL.
type GetResponse struct {
	Rval *UserListPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Mutate was auto-generated from WSDL.
type Mutate struct {
	Operations []*UserListOperation `xml:"operations,omitempty" json:"operations,omitempty" yaml:"operations,omitempty"`
}

// MutateMembers was auto-generated from WSDL.
type MutateMembers struct {
	Operations []*MutateMembersOperation `xml:"operations,omitempty" json:"operations,omitempty" yaml:"operations,omitempty"`
}

// MutateMembersResponse was auto-generated from WSDL.
type MutateMembersResponse struct {
	Rval *MutateMembersReturnValue `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// MutateResponse was auto-generated from WSDL.
type MutateResponse struct {
	Rval *UserListReturnValue `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Query was auto-generated from WSDL.
type Query struct {
	Query *string `xml:"query,omitempty" json:"query,omitempty" yaml:"query,omitempty"`
}

// QueryResponse was auto-generated from WSDL.
type QueryResponse struct {
	Rval *UserListPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
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

// Operation wrapper for MutateMembers.
// OperationMutateMembersRequest was auto-generated from WSDL.
type OperationMutateMembersRequest struct {
	MutateMembers *MutateMembers `xml:"mutateMembers,omitempty" json:"mutateMembers,omitempty" yaml:"mutateMembers,omitempty"`
}

// Operation wrapper for MutateMembers.
// OperationMutateMembersResponse was auto-generated from WSDL.
type OperationMutateMembersResponse struct {
	MutateMembersResponse *MutateMembersResponse `xml:"mutateMembersResponse,omitempty" json:"mutateMembersResponse,omitempty" yaml:"mutateMembersResponse,omitempty"`
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

// adwordsUserListServiceInterface implements the AdwordsUserListServiceInterface interface.
type adwordsUserListServiceInterface struct {
	cli *soap.Client
}

// Returns the list of user lists that meet the selector criteria.
//                  @param serviceSelector the selector specifying
// the {@link UserList}s to return.         @return a list of UserList
// entities which meet the selector criteria.         @throws ApiException
// if problems occurred while fetching UserList information.
func (p *adwordsUserListServiceInterface) Get(Get *Get) (*GetResponse, error) {
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

// Applies a list of mutate operations (i.e. add, set):
//           Add - creates a set of user lists         Set - updates
// a set of user lists         Remove - not supported
//         @param operations the operations to apply         @return
// a list of UserList objects
func (p *adwordsUserListServiceInterface) Mutate(Mutate *Mutate) (*MutateResponse, error) {
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

// Mutate members of user lists by either adding or removing their
// lists of members.         The following {@link Operator}s are
// supported: ADD and REMOVE. The SET operator         is not supported.
//                  <p>Note that operations cannot have same user
// list id but different operators.                  @param operations
// the mutate members operations to apply         @return a list
// of UserList objects         @throws ApiException when there
// are one or more errors with the request
func (p *adwordsUserListServiceInterface) MutateMembers(MutateMembers *MutateMembers) (*MutateMembersResponse, error) {
	α := struct {
		OperationMutateMembersRequest `xml:"tns:mutateMembers"`
	}{
		OperationMutateMembersRequest{
			MutateMembers,
		},
	}

	γ := struct {
		OperationMutateMembersResponse `xml:"mutateMembersResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("MutateMembers", α, &γ); err != nil {
		return nil, err
	}
	return γ.MutateMembersResponse, nil
}

// Returns the list of user lists that match the query.
//           @param query The SQL-like AWQL query string
//   @return A list of UserList         @throws ApiException when
// the query is invalid or there are errors processing the request.
func (p *adwordsUserListServiceInterface) Query(Query string) (*QueryResponse, error) {
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
