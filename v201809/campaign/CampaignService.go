package campaign

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "https://adwords.google.com/api/adwords/cm/v201809"

// NewCampaignServiceInterface creates an initializes a CampaignServiceInterface.
func NewCampaignServiceInterface(cli *soap.Client) CampaignServiceInterface {
	return &campaignServiceInterface{cli}
}

// CampaignServiceInterface was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type CampaignServiceInterface interface {
	// Returns the list of campaigns that meet the selector criteria.
	//                  @param serviceSelector the selector specifying
	// the {@link Campaign}s to return.         @return A list of campaigns.
	//         @throws ApiException if problems occurred while fetching
	// campaign information.
	Get(Get *Get) (*GetResponse, error)

	// Adds, updates, or removes campaigns.         <p class="note"><b>Note:</b>
	// {@link CampaignOperation} does not support the         <code>REMOVE</code>
	// operator. To remove a campaign, set its         {@link Campaign#status
	// status} to {@code REMOVED}.</p>         @param operations A
	// list of unique operations.         The same campaign cannot
	// be specified in more than one operation.         @return The
	// list of updated campaigns, returned in the same order as the
	//         <code>operations</code> array.         @throws ApiException
	// if problems occurred while updating campaign information.
	Mutate(Mutate *Mutate) (*MutateResponse, error)

	// Adds labels to the {@linkplain Campaign campaign} or removes
	// {@linkplain Label label}s from the         {@linkplain Campaign
	// campaign}.         <p>Add - Apply an existing label to an existing
	// {@linkplain Campaign campaign}. The         {@code campaignId}
	// must reference an existing {@linkplain Campaign}. The {@code
	// labelId} must         reference an existing {@linkplain Label
	// label}.         <p>Remove - Removes the link between the specified
	// {@linkplain Campaign campaign} and         {@linkplain Label
	// label}.                  @param operations the operations to
	// apply.         @return a list of {@linkplain CampaignLabel}s
	// where each entry in the list is the result of         applying
	// the operation in the input list with the same index. For an
	//         add operation, the returned CampaignLabel contains the
	// CampaignId and the LabelId.         In the case of a remove
	// operation, the returned CampaignLabel will only have CampaignId.
	//         @throws ApiException when there are one or more errors
	// with the request.
	MutateLabel(MutateLabel *MutateLabel) (*MutateLabelResponse, error)

	// Returns the list of campaigns that match the query.
	//          @param query The SQL-like AWQL query string.
	//   @return A list of campaigns.         @throws ApiException
	// if problems occur while parsing the query or fetching campaign
	//         information.
	Query(Query string) (*QueryResponse, error)
}

// AdServingOptimizationStatus was auto-generated from WSDL.
type AdServingOptimizationStatus string

// Validate validates AdServingOptimizationStatus.
func (v AdServingOptimizationStatus) Validate() bool {
	for _, vv := range []string{
		"OPTIMIZE",
		"CONVERSION_OPTIMIZE",
		"ROTATE",
		"ROTATE_INDEFINITELY",
		"UNAVAILABLE",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AdvertisingChannelSubType was auto-generated from WSDL.
type AdvertisingChannelSubType string

// Validate validates AdvertisingChannelSubType.
func (v AdvertisingChannelSubType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"SEARCH_MOBILE_APP",
		"DISPLAY_MOBILE_APP",
		"SEARCH_EXPRESS",
		"DISPLAY_EXPRESS",
		"UNIVERSAL_APP_CAMPAIGN",
		"DISPLAY_SMART_CAMPAIGN",
		"SHOPPING_GOAL_OPTIMIZED_ADS",
		"DISPLAY_GMAIL_AD",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AdvertisingChannelType was auto-generated from WSDL.
type AdvertisingChannelType string

// Validate validates AdvertisingChannelType.
func (v AdvertisingChannelType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"SEARCH",
		"DISPLAY",
		"SHOPPING",
		"MULTI_CHANNEL",
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

// BudgetBudgetDeliveryMethod was auto-generated from WSDL.
type BudgetBudgetDeliveryMethod string

// Validate validates BudgetBudgetDeliveryMethod.
func (v BudgetBudgetDeliveryMethod) Validate() bool {
	for _, vv := range []string{
		"STANDARD",
		"ACCELERATED",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// BudgetBudgetStatus was auto-generated from WSDL.
type BudgetBudgetStatus string

// Validate validates BudgetBudgetStatus.
func (v BudgetBudgetStatus) Validate() bool {
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

// CampaignErrorReason was auto-generated from WSDL.
type CampaignErrorReason string

// Validate validates CampaignErrorReason.
func (v CampaignErrorReason) Validate() bool {
	for _, vv := range []string{
		"CANNOT_GO_BACK_TO_INCOMPLETE",
		"CANNOT_TARGET_CONTENT_NETWORK",
		"CANNOT_TARGET_SEARCH_NETWORK",
		"CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH",
		"CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN",
		"CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK",
		"CANNOT_TARGET_PARTNER_SEARCH_NETWORK",
		"CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY",
		"CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS",
		"CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN",
		"DUPLICATE_CAMPAIGN_NAME",
		"INCOMPATIBLE_CAMPAIGN_FIELD",
		"INVALID_CAMPAIGN_NAME",
		"INVALID_AD_SERVING_OPTIMIZATION_STATUS",
		"INVALID_TRACKING_URL",
		"CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING",
		"MAX_IMPRESSIONS_NOT_IN_RANGE",
		"TIME_UNIT_NOT_SUPPORTED",
		"INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED",
		"BUDGET_CANNOT_BE_SHARED",
		"CAMPAIGN_CANNOT_USE_SHARED_BUDGET",
		"CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS",
		"CAMPAIGN_LABEL_DOES_NOT_EXIST",
		"CAMPAIGN_LABEL_ALREADY_EXISTS",
		"MISSING_SHOPPING_SETTING",
		"INVALID_SHOPPING_SALES_COUNTRY",
		"MISSING_UNIVERSAL_APP_CAMPAIGN_SETTING",
		"ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE",
		"INVALID_ADVERTISING_CHANNEL_SUB_TYPE",
		"CAMPAIGN_CANNOT_USE_SHARED_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP",
		"AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED",
		"CANNOT_SET_AD_ROTATION_MODE",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// CampaignStatus was auto-generated from WSDL.
type CampaignStatus string

// Validate validates CampaignStatus.
func (v CampaignStatus) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"ENABLED",
		"PAUSED",
		"REMOVED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// CampaignTrialType was auto-generated from WSDL.
type CampaignTrialType string

// Validate validates CampaignTrialType.
func (v CampaignTrialType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"BASE",
		"DRAFT",
		"TRIAL",
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

// ConversionOptimizerEligibilityRejectionReason was auto-generated
// from WSDL.
type ConversionOptimizerEligibilityRejectionReason string

// Validate validates ConversionOptimizerEligibilityRejectionReason.
func (v ConversionOptimizerEligibilityRejectionReason) Validate() bool {
	for _, vv := range []string{
		"CAMPAIGN_IS_NOT_ACTIVE",
		"NOT_CPC_CAMPAIGN",
		"CONVERSION_TRACKING_NOT_ENABLED",
		"NOT_ENOUGH_CONVERSIONS",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// CriterionTypeGroup was auto-generated from WSDL.
type CriterionTypeGroup string

// Validate validates CriterionTypeGroup.
func (v CriterionTypeGroup) Validate() bool {
	for _, vv := range []string{
		"KEYWORD",
		"USER_INTEREST_AND_LIST",
		"VERTICAL",
		"GENDER",
		"AGE_RANGE",
		"PLACEMENT",
		"PARENT",
		"INCOME_RANGE",
		"NONE",
		"UNKNOWN",
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

// DateRangeErrorReason was auto-generated from WSDL.
type DateRangeErrorReason string

// Validate validates DateRangeErrorReason.
func (v DateRangeErrorReason) Validate() bool {
	for _, vv := range []string{
		"DATE_RANGE_ERROR",
		"INVALID_DATE",
		"START_DATE_AFTER_END_DATE",
		"CANNOT_SET_DATE_TO_PAST",
		"AFTER_MAXIMUM_ALLOWABLE_DATE",
		"CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED",
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

// GeoTargetTypeSettingNegativeGeoTargetType was auto-generated
// from WSDL.
type GeoTargetTypeSettingNegativeGeoTargetType string

// Validate validates GeoTargetTypeSettingNegativeGeoTargetType.
func (v GeoTargetTypeSettingNegativeGeoTargetType) Validate() bool {
	for _, vv := range []string{
		"DONT_CARE",
		"LOCATION_OF_PRESENCE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// GeoTargetTypeSettingPositiveGeoTargetType was auto-generated
// from WSDL.
type GeoTargetTypeSettingPositiveGeoTargetType string

// Validate validates GeoTargetTypeSettingPositiveGeoTargetType.
func (v GeoTargetTypeSettingPositiveGeoTargetType) Validate() bool {
	for _, vv := range []string{
		"DONT_CARE",
		"AREA_OF_INTEREST",
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

// Level was auto-generated from WSDL.
type Level string

// Validate validates Level.
func (v Level) Validate() bool {
	for _, vv := range []string{
		"CREATIVE",
		"ADGROUP",
		"CAMPAIGN",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ListErrorReason was auto-generated from WSDL.
type ListErrorReason string

// Validate validates ListErrorReason.
func (v ListErrorReason) Validate() bool {
	for _, vv := range []string{
		"CLEAR_UNSUPPORTED",
		"INVALID_OPERATOR",
		"INVALID_ELEMENT",
		"LIST_LENGTH_MISMATCH",
		"DUPLICATE_ELEMENT",
		"MUTATE_UNSUPPORTED",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ListOperationsListOperator was auto-generated from WSDL.
type ListOperationsListOperator string

// Validate validates ListOperationsListOperator.
func (v ListOperationsListOperator) Validate() bool {
	for _, vv := range []string{
		"PUT",
		"REMOVE",
		"UPDATE",
		"UNKNOWN",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// MobileApplicationVendor was auto-generated from WSDL.
type MobileApplicationVendor string

// Validate validates MobileApplicationVendor.
func (v MobileApplicationVendor) Validate() bool {
	for _, vv := range []string{
		"VENDOR_UNKNOWN",
		"VENDOR_APPLE_APP_STORE",
		"VENDOR_GOOGLE_MARKET",
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

// ServingStatus was auto-generated from WSDL.
type ServingStatus string

// Validate validates ServingStatus.
func (v ServingStatus) Validate() bool {
	for _, vv := range []string{
		"SERVING",
		"NONE",
		"ENDED",
		"PENDING",
		"SUSPENDED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// SettingErrorReason was auto-generated from WSDL.
type SettingErrorReason string

// Validate validates SettingErrorReason.
func (v SettingErrorReason) Validate() bool {
	for _, vv := range []string{
		"DUPLICATE_SETTING_TYPE",
		"SETTING_TYPE_IS_NOT_AVAILABLE",
		"SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN",
		"TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP",
		"TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL",
		"TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP",
		"DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT",
		"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME",
		"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME",
		"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE",
		"TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN",
		"UNIVERSAL_APP_CAMPAIGN_SETTING_DUPLICATE_DESCRIPTION",
		"UNIVERSAL_APP_CAMPAIGN_SETTING_DESCRIPTION_LINE_WIDTH_TOO_LONG",
		"UNIVERSAL_APP_CAMPAIGN_SETTING_APP_ID_CANNOT_BE_MODIFIED",
		"TOO_MANY_YOUTUBE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN",
		"TOO_MANY_IMAGE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN",
		"MEDIA_INCOMPATIBLE_FOR_UNIVERSAL_APP_CAMPAIGN",
		"TOO_MANY_EXCLAMATION_MARKS",
		"UNKNOWN",
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

// TimeUnit was auto-generated from WSDL.
type TimeUnit string

// Validate validates TimeUnit.
func (v TimeUnit) Validate() bool {
	for _, vv := range []string{
		"MINUTE",
		"HOUR",
		"DAY",
		"WEEK",
		"MONTH",
		"LIFETIME",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// UniversalAppBiddingStrategyGoalType was auto-generated from
// WSDL.
type UniversalAppBiddingStrategyGoalType string

// Validate validates UniversalAppBiddingStrategyGoalType.
func (v UniversalAppBiddingStrategyGoalType) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"OPTIMIZE_FOR_INSTALL_CONVERSION_VOLUME",
		"OPTIMIZE_FOR_IN_APP_CONVERSION_VOLUME",
		"OPTIMIZE_FOR_TOTAL_CONVERSION_VALUE",
		"OPTIMIZE_FOR_TARGET_IN_APP_CONVERSION",
		"OPTIMIZE_FOR_RETURN_ON_ADVERTISING_SPEND",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// UniversalAppCampaignAsset was auto-generated from WSDL.
type UniversalAppCampaignAsset string

// Validate validates UniversalAppCampaignAsset.
func (v UniversalAppCampaignAsset) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"COMBINATION",
		"APP_DESTINATION",
		"APP_ASSETS",
		"DESCRIPTION_1",
		"DESCRIPTION_2",
		"DESCRIPTION_3",
		"DESCRIPTION_4",
		"VIDEO",
		"IMAGE",
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

// VanityPharmaDisplayUrlMode was auto-generated from WSDL.
type VanityPharmaDisplayUrlMode string

// Validate validates VanityPharmaDisplayUrlMode.
func (v VanityPharmaDisplayUrlMode) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"MANUFACTURER_WEBSITE_URL",
		"WEBSITE_DESCRIPTION",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// VanityPharmaText was auto-generated from WSDL.
type VanityPharmaText string

// Validate validates VanityPharmaText.
func (v VanityPharmaText) Validate() bool {
	for _, vv := range []string{
		"UNKNOWN",
		"PRESCRIPTION_TREATMENT_WEBSITE_EN",
		"PRESCRIPTION_TREATMENT_WEBSITE_ES",
		"PRESCRIPTION_DEVICE_WEBSITE_EN",
		"PRESCRIPTION_DEVICE_WEBSITE_ES",
		"MEDICAL_DEVICE_WEBSITE_EN",
		"MEDICAL_DEVICE_WEBSITE_ES",
		"PREVENTATIVE_TREATMENT_WEBSITE_EN",
		"PREVENTATIVE_TREATMENT_WEBSITE_ES",
		"PRESCRIPTION_CONTRACEPTION_WEBSITE_EN",
		"PRESCRIPTION_CONTRACEPTION_WEBSITE_ES",
		"PRESCRIPTION_VACCINE_WEBSITE_EN",
		"PRESCRIPTION_VACCINE_WEBSITE_ES",
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

// Budgets are used for managing the amount of money spent on AdWords.
type Budget struct {
	BudgetId           *int64                      `xml:"budgetId,omitempty" json:"budgetId,omitempty" yaml:"budgetId,omitempty"`
	Name               *string                     `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Amount             *Money                      `xml:"amount,omitempty" json:"amount,omitempty" yaml:"amount,omitempty"`
	DeliveryMethod     *BudgetBudgetDeliveryMethod `xml:"deliveryMethod,omitempty" json:"deliveryMethod,omitempty" yaml:"deliveryMethod,omitempty"`
	ReferenceCount     *int                        `xml:"referenceCount,omitempty" json:"referenceCount,omitempty" yaml:"referenceCount,omitempty"`
	IsExplicitlyShared *bool                       `xml:"isExplicitlyShared,omitempty" json:"isExplicitlyShared,omitempty" yaml:"isExplicitlyShared,omitempty"`
	Status             *BudgetBudgetStatus         `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
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
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Data representing an AdWords campaign.
type Campaign struct {
	Id                             *int64                          `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	CampaignGroupId                *int64                          `xml:"campaignGroupId,omitempty" json:"campaignGroupId,omitempty" yaml:"campaignGroupId,omitempty"`
	Name                           *string                         `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Status                         *CampaignStatus                 `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	ServingStatus                  *ServingStatus                  `xml:"servingStatus,omitempty" json:"servingStatus,omitempty" yaml:"servingStatus,omitempty"`
	StartDate                      *string                         `xml:"startDate,omitempty" json:"startDate,omitempty" yaml:"startDate,omitempty"`
	EndDate                        *string                         `xml:"endDate,omitempty" json:"endDate,omitempty" yaml:"endDate,omitempty"`
	Budget                         *Budget                         `xml:"budget,omitempty" json:"budget,omitempty" yaml:"budget,omitempty"`
	ConversionOptimizerEligibility *ConversionOptimizerEligibility `xml:"conversionOptimizerEligibility,omitempty" json:"conversionOptimizerEligibility,omitempty" yaml:"conversionOptimizerEligibility,omitempty"`
	AdServingOptimizationStatus    *AdServingOptimizationStatus    `xml:"adServingOptimizationStatus,omitempty" json:"adServingOptimizationStatus,omitempty" yaml:"adServingOptimizationStatus,omitempty"`
	FrequencyCap                   *FrequencyCap                   `xml:"frequencyCap,omitempty" json:"frequencyCap,omitempty" yaml:"frequencyCap,omitempty"`
	Settings                       []*Setting                      `xml:"settings,omitempty" json:"settings,omitempty" yaml:"settings,omitempty"`
	AdvertisingChannelType         *AdvertisingChannelType         `xml:"advertisingChannelType,omitempty" json:"advertisingChannelType,omitempty" yaml:"advertisingChannelType,omitempty"`
	AdvertisingChannelSubType      *AdvertisingChannelSubType      `xml:"advertisingChannelSubType,omitempty" json:"advertisingChannelSubType,omitempty" yaml:"advertisingChannelSubType,omitempty"`
	NetworkSetting                 *NetworkSetting                 `xml:"networkSetting,omitempty" json:"networkSetting,omitempty" yaml:"networkSetting,omitempty"`
	Labels                         []*Label                        `xml:"labels,omitempty" json:"labels,omitempty" yaml:"labels,omitempty"`
	BiddingStrategyConfiguration   *BiddingStrategyConfiguration   `xml:"biddingStrategyConfiguration,omitempty" json:"biddingStrategyConfiguration,omitempty" yaml:"biddingStrategyConfiguration,omitempty"`
	CampaignTrialType              *CampaignTrialType              `xml:"campaignTrialType,omitempty" json:"campaignTrialType,omitempty" yaml:"campaignTrialType,omitempty"`
	BaseCampaignId                 *int64                          `xml:"baseCampaignId,omitempty" json:"baseCampaignId,omitempty" yaml:"baseCampaignId,omitempty"`
	ForwardCompatibilityMap        []*String_StringMapEntry        `xml:"forwardCompatibilityMap,omitempty" json:"forwardCompatibilityMap,omitempty" yaml:"forwardCompatibilityMap,omitempty"`
	TrackingUrlTemplate            *string                         `xml:"trackingUrlTemplate,omitempty" json:"trackingUrlTemplate,omitempty" yaml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix                 *string                         `xml:"finalUrlSuffix,omitempty" json:"finalUrlSuffix,omitempty" yaml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters            *CustomParameters               `xml:"urlCustomParameters,omitempty" json:"urlCustomParameters,omitempty" yaml:"urlCustomParameters,omitempty"`
	VanityPharma                   *VanityPharma                   `xml:"vanityPharma,omitempty" json:"vanityPharma,omitempty" yaml:"vanityPharma,omitempty"`
	UniversalAppCampaignInfo       *UniversalAppCampaignInfo       `xml:"universalAppCampaignInfo,omitempty" json:"universalAppCampaignInfo,omitempty" yaml:"universalAppCampaignInfo,omitempty"`
	SelectiveOptimization          *SelectiveOptimization          `xml:"selectiveOptimization,omitempty" json:"selectiveOptimization,omitempty" yaml:"selectiveOptimization,omitempty"`
}

// Base error class for Campaign Service.
type CampaignError struct {
	FieldPath         *string              `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement  `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string              `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string              `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string              `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *CampaignErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string               `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string               `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CampaignError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CampaignError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Manages the labels associated with a campaign.
type CampaignLabel struct {
	CampaignId *int64 `xml:"campaignId,omitempty" json:"campaignId,omitempty" yaml:"campaignId,omitempty"`
	LabelId    *int64 `xml:"labelId,omitempty" json:"labelId,omitempty" yaml:"labelId,omitempty"`
}

// Operations for adding/removing labels from Campaign.
type CampaignLabelOperation struct {
	Operator      *Operator      `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	OperationType *string        `xml:"Operation.Type,omitempty" json:"Operation.Type,omitempty" yaml:"Operation.Type,omitempty"`
	Operand       *CampaignLabel `xml:"operand,omitempty" json:"operand,omitempty" yaml:"operand,omitempty"`
	TypeAttrXSI   string         `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string         `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CampaignLabelOperation) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CampaignLabelOperation"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A container for return values from the {@link CampaignService#mutateLabel}
// call.
type CampaignLabelReturnValue struct {
	ListReturnValueType  *string          `xml:"ListReturnValue.Type,omitempty" json:"ListReturnValue.Type,omitempty" yaml:"ListReturnValue.Type,omitempty"`
	Value                []*CampaignLabel `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	PartialFailureErrors []*ApiError      `xml:"partialFailureErrors,omitempty" json:"partialFailureErrors,omitempty" yaml:"partialFailureErrors,omitempty"`
	TypeAttrXSI          string           `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string           `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CampaignLabelReturnValue) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CampaignLabelReturnValue"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// An operation on an AdWords campaign.             <p class="note"><b>Note:</b>
// The <code>REMOVE</code> operator is not             supported.
// To remove a campaign, set its {@link Campaign#status status}
//             to {@code REMOVED}.</p>
type CampaignOperation struct {
	Operator      *Operator `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	OperationType *string   `xml:"Operation.Type,omitempty" json:"Operation.Type,omitempty" yaml:"Operation.Type,omitempty"`
	Operand       *Campaign `xml:"operand,omitempty" json:"operand,omitempty" yaml:"operand,omitempty"`
	TypeAttrXSI   string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CampaignOperation) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CampaignOperation"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Contains a set of campaigns resulting from the filtering and
// paging of the             {@link CampaignService#get} call.
type CampaignPage struct {
	TotalNumEntries *int        `xml:"totalNumEntries,omitempty" json:"totalNumEntries,omitempty" yaml:"totalNumEntries,omitempty"`
	PageType        *string     `xml:"Page.Type,omitempty" json:"Page.Type,omitempty" yaml:"Page.Type,omitempty"`
	Entries         []*Campaign `xml:"entries,omitempty" json:"entries,omitempty" yaml:"entries,omitempty"`
	TypeAttrXSI     string      `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string      `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CampaignPage) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CampaignPage"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// A container for return values from the CampaignService.
type CampaignReturnValue struct {
	ListReturnValueType  *string     `xml:"ListReturnValue.Type,omitempty" json:"ListReturnValue.Type,omitempty" yaml:"ListReturnValue.Type,omitempty"`
	Value                []*Campaign `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	PartialFailureErrors []*ApiError `xml:"partialFailureErrors,omitempty" json:"partialFailureErrors,omitempty" yaml:"partialFailureErrors,omitempty"`
	TypeAttrXSI          string      `xml:"xsi:type,attr,omitempty"`
	TypeNamespace        string      `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *CampaignReturnValue) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:CampaignReturnValue"
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

// Comparable types for constructing ranges with.
type ComparableValue interface{}

// Eligibility data for Campaign to transition to Conversion Optimizer
type ConversionOptimizerEligibility struct {
	Eligible         *bool                                            `xml:"eligible,omitempty" json:"eligible,omitempty" yaml:"eligible,omitempty"`
	RejectionReasons []*ConversionOptimizerEligibilityRejectionReason `xml:"rejectionReasons,omitempty" json:"rejectionReasons,omitempty" yaml:"rejectionReasons,omitempty"`
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

// A list of all the error codes being used for date range error.
type DateRangeError struct {
	FieldPath         *string               `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement   `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string               `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string               `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string               `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *DateRangeErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string                `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string                `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DateRangeError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DateRangeError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
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

// Setting for controlling Dynamic Search Ads (DSA). Contains the
// domain name and the language             used by the DSA system
// to automatically generate landing pages and keywords for a campaign.
type DynamicSearchAdsSetting struct {
	SettingType         *string   `xml:"Setting.Type,omitempty" json:"Setting.Type,omitempty" yaml:"Setting.Type,omitempty"`
	DomainName          *string   `xml:"domainName,omitempty" json:"domainName,omitempty" yaml:"domainName,omitempty"`
	LanguageCode        *string   `xml:"languageCode,omitempty" json:"languageCode,omitempty" yaml:"languageCode,omitempty"`
	UseSuppliedUrlsOnly *bool     `xml:"useSuppliedUrlsOnly,omitempty" json:"useSuppliedUrlsOnly,omitempty" yaml:"useSuppliedUrlsOnly,omitempty"`
	PageFeed            *PageFeed `xml:"pageFeed,omitempty" json:"pageFeed,omitempty" yaml:"pageFeed,omitempty"`
	TypeAttrXSI         string    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace       string    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *DynamicSearchAdsSetting) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:DynamicSearchAdsSetting"
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

// A frequency cap is the maximum number of times an ad (or some
// set of ads) can             be shown to a user over a particular
// time period.
type FrequencyCap struct {
	Impressions *int64    `xml:"impressions,omitempty" json:"impressions,omitempty" yaml:"impressions,omitempty"`
	TimeUnit    *TimeUnit `xml:"timeUnit,omitempty" json:"timeUnit,omitempty" yaml:"timeUnit,omitempty"`
	Level       *Level    `xml:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
}

// Represents a collection of settings related to ads geotargeting.
//                          <p>AdWords ads can be geotargeted using
// <b>Location of Presence</b> (<b>LOP</b>),             <b>Area
// of Interest</b> (<b>AOI</b>), or both. LOP is the physical location
//             of the user performing the search; AOI is the geographical
// region             in which the searcher is interested. For
// example, if a user in             New York City performs a search
// "hotels california", their LOP             is New York City
// and their AOI is California.                          <p>Additionally,
// ads can be <b>positively</b> or <b>negatively</b> geotargeted.
//             An ad that is positively geotargeted to New York
// City only appears             to users whose location is related
// (via AOI or LOP) to New York City. An ad             that is
// negatively geotargeted to New York City appears for <i>all</i>
//             users <i>except</i> those whose location is related
// to New York City. Ads can             only be negatively geotargeted
// if a positive geotargeting is also supplied, and
//  the negatively geotargeted region must be contained within
// the positive             region.                          <p>Geotargeting
// settings allow ads to be targeted in the following way:
//         <ul>             <li> Positively geotargeted using solely
// AOI, solely LOP, or either.             <li> Negatively geotargeted
// using solely LOP, or both.             </ul>
type GeoTargetTypeSetting struct {
	SettingType           *string                                    `xml:"Setting.Type,omitempty" json:"Setting.Type,omitempty" yaml:"Setting.Type,omitempty"`
	PositiveGeoTargetType *GeoTargetTypeSettingPositiveGeoTargetType `xml:"positiveGeoTargetType,omitempty" json:"positiveGeoTargetType,omitempty" yaml:"positiveGeoTargetType,omitempty"`
	NegativeGeoTargetType *GeoTargetTypeSettingNegativeGeoTargetType `xml:"negativeGeoTargetType,omitempty" json:"negativeGeoTargetType,omitempty" yaml:"negativeGeoTargetType,omitempty"`
	TypeAttrXSI           string                                     `xml:"xsi:type,attr,omitempty"`
	TypeNamespace         string                                     `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *GeoTargetTypeSetting) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:GeoTargetTypeSetting"
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

// Errors that can occur when making changes to a list using {@link
// ListOperations}.
type ListError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *ListErrorReason    `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ListError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ListError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Describes the behavior of elements in a list. Instances of ListOperations
// will always be defined             alongside some list in an
// API POJO. The number of operators in the ListOperations must
// be             equal to the number of elements in the POJO list.
// Each operator, together with its corresponding             list
// element, describe an intended change.
//    <p>For example, if in a request Campaign.selectiveOptimization
// contains 2 conversionTypeIds,             and the conversionTypeIdsOps
// is non-null, it must contain 2 operators. If those operators
// are             {PUT, REMOVE} then the API will add the first
// conversionTypeId (if it doesn't             already exist) and
// remove the second conversionTypeId (if it exists).
type ListOperations struct {
	Clear     *bool                         `xml:"clear,omitempty" json:"clear,omitempty" yaml:"clear,omitempty"`
	Operators []*ListOperationsListOperator `xml:"operators,omitempty" json:"operators,omitempty" yaml:"operators,omitempty"`
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

// Network settings for a Campaign.
type NetworkSetting struct {
	TargetGoogleSearch         *bool `xml:"targetGoogleSearch,omitempty" json:"targetGoogleSearch,omitempty" yaml:"targetGoogleSearch,omitempty"`
	TargetSearchNetwork        *bool `xml:"targetSearchNetwork,omitempty" json:"targetSearchNetwork,omitempty" yaml:"targetSearchNetwork,omitempty"`
	TargetContentNetwork       *bool `xml:"targetContentNetwork,omitempty" json:"targetContentNetwork,omitempty" yaml:"targetContentNetwork,omitempty"`
	TargetPartnerSearchNetwork *bool `xml:"targetPartnerSearchNetwork,omitempty" json:"targetPartnerSearchNetwork,omitempty" yaml:"targetPartnerSearchNetwork,omitempty"`
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

// Page feeds associated with {@link DynamicSearchAdsSetting}.
type PageFeed struct {
	FeedIds []*int64 `xml:"feedIds,omitempty" json:"feedIds,omitempty" yaml:"feedIds,omitempty"`
}

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

// Settings for Real-Time Bidding, a feature only available for
// campaigns             targeting the Ad Exchange network.
type RealTimeBiddingSetting struct {
	SettingType   *string `xml:"Setting.Type,omitempty" json:"Setting.Type,omitempty" yaml:"Setting.Type,omitempty"`
	OptIn         *bool   `xml:"optIn,omitempty" json:"optIn,omitempty" yaml:"optIn,omitempty"`
	TypeAttrXSI   string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *RealTimeBiddingSetting) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:RealTimeBiddingSetting"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
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

// Selected set of conversion types or a selective optimization
// set id for optimizing campaigns. For             e.g. For universal
// app campaigns, these are the set of in-app actions to optimize
// the campaign             towards.
type SelectiveOptimization struct {
	ConversionTypeIds    []*int64        `xml:"conversionTypeIds,omitempty" json:"conversionTypeIds,omitempty" yaml:"conversionTypeIds,omitempty"`
	ConversionTypeIdsOps *ListOperations `xml:"conversionTypeIdsOps,omitempty" json:"conversionTypeIdsOps,omitempty" yaml:"conversionTypeIdsOps,omitempty"`
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

// Base type for AdWords campaign settings.
type Setting interface{}

// Indicates a problem with campaign settings.
type SettingError struct {
	FieldPath         *string             `xml:"fieldPath,omitempty" json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty" json:"fieldPathElements,omitempty" yaml:"fieldPathElements,omitempty"`
	Trigger           *string             `xml:"trigger,omitempty" json:"trigger,omitempty" yaml:"trigger,omitempty"`
	ErrorString       *string             `xml:"errorString,omitempty" json:"errorString,omitempty" yaml:"errorString,omitempty"`
	ApiErrorType      *string             `xml:"ApiError.Type,omitempty" json:"ApiError.Type,omitempty" yaml:"ApiError.Type,omitempty"`
	Reason            *SettingErrorReason `xml:"reason,omitempty" json:"reason,omitempty" yaml:"reason,omitempty"`
	TypeAttrXSI       string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *SettingError) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:SettingError"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Setting for shopping campaigns. Defines the universe of products
// covered by the campaign.             Encapsulates a merchant
// ID, sales country, and campaign priority.
type ShoppingSetting struct {
	SettingType      *string `xml:"Setting.Type,omitempty" json:"Setting.Type,omitempty" yaml:"Setting.Type,omitempty"`
	MerchantId       *int64  `xml:"merchantId,omitempty" json:"merchantId,omitempty" yaml:"merchantId,omitempty"`
	SalesCountry     *string `xml:"salesCountry,omitempty" json:"salesCountry,omitempty" yaml:"salesCountry,omitempty"`
	CampaignPriority *int    `xml:"campaignPriority,omitempty" json:"campaignPriority,omitempty" yaml:"campaignPriority,omitempty"`
	EnableLocal      *bool   `xml:"enableLocal,omitempty" json:"enableLocal,omitempty" yaml:"enableLocal,omitempty"`
	TypeAttrXSI      string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace    string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ShoppingSetting) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:ShoppingSetting"
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

// Setting for targeting related features.             This is
// applicable at Campaign and AdGroup level.
type TargetingSetting struct {
	SettingType   *string                   `xml:"Setting.Type,omitempty" json:"Setting.Type,omitempty" yaml:"Setting.Type,omitempty"`
	Details       []*TargetingSettingDetail `xml:"details,omitempty" json:"details,omitempty" yaml:"details,omitempty"`
	TypeAttrXSI   string                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TargetingSetting) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TargetingSetting"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Specifies if criteria of this type group should be used to restrict
//             targeting, or if ads can serve anywhere and criteria
// are only used in             determining the bid.
//   <p>For more information, see             <a href="https://support.google.com/adwords/answer/6056342">Targeting
// Settings</a>.</p>
type TargetingSettingDetail struct {
	CriterionTypeGroup *CriterionTypeGroup `xml:"criterionTypeGroup,omitempty" json:"criterionTypeGroup,omitempty" yaml:"criterionTypeGroup,omitempty"`
	TargetAll          *bool               `xml:"targetAll,omitempty" json:"targetAll,omitempty" yaml:"targetAll,omitempty"`
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

// Campaign level settings for tracking information.
type TrackingSetting struct {
	SettingType   *string `xml:"Setting.Type,omitempty" json:"Setting.Type,omitempty" yaml:"Setting.Type,omitempty"`
	TrackingUrl   *string `xml:"trackingUrl,omitempty" json:"trackingUrl,omitempty" yaml:"trackingUrl,omitempty"`
	TypeAttrXSI   string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *TrackingSetting) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:TrackingSetting"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://adwords.google.com/api/adwords/cm/v201809"
	}
}

// Contains Universal App Campaign Ads Policy decisions with asset
// identifier information, where             available.
type UniversalAppCampaignAdsPolicyDecisions struct {
	UniversalAppCampaignAsset *UniversalAppCampaignAsset `xml:"universalAppCampaignAsset,omitempty" json:"universalAppCampaignAsset,omitempty" yaml:"universalAppCampaignAsset,omitempty"`
	AssetId                   *string                    `xml:"assetId,omitempty" json:"assetId,omitempty" yaml:"assetId,omitempty"`
	PolicyTopicEntries        []*PolicyTopicEntry        `xml:"policyTopicEntries,omitempty" json:"policyTopicEntries,omitempty" yaml:"policyTopicEntries,omitempty"`
}

// Stores campaign-level info related to Universal App Campaigns.
type UniversalAppCampaignInfo struct {
	BiddingStrategyGoalType *UniversalAppBiddingStrategyGoalType `xml:"biddingStrategyGoalType,omitempty" json:"biddingStrategyGoalType,omitempty" yaml:"biddingStrategyGoalType,omitempty"`
	AppId                   *string                              `xml:"appId,omitempty" json:"appId,omitempty" yaml:"appId,omitempty"`
	AppVendor               *MobileApplicationVendor             `xml:"appVendor,omitempty" json:"appVendor,omitempty" yaml:"appVendor,omitempty"`
}

// Setting for storing the mobile app to advertise and creative
// assets for Universal app campaigns.             This setting
// is required for Campaigns with advertising channel subtype UNIVERSAL_APP_CAMPAIGN
//             and can only be attached to such Campaigns.
type UniversalAppCampaignSetting struct {
	SettingType                         *string                                   `xml:"Setting.Type,omitempty" json:"Setting.Type,omitempty" yaml:"Setting.Type,omitempty"`
	AppId                               *string                                   `xml:"appId,omitempty" json:"appId,omitempty" yaml:"appId,omitempty"`
	AppVendor                           *MobileApplicationVendor                  `xml:"appVendor,omitempty" json:"appVendor,omitempty" yaml:"appVendor,omitempty"`
	Description1                        *string                                   `xml:"description1,omitempty" json:"description1,omitempty" yaml:"description1,omitempty"`
	Description2                        *string                                   `xml:"description2,omitempty" json:"description2,omitempty" yaml:"description2,omitempty"`
	Description3                        *string                                   `xml:"description3,omitempty" json:"description3,omitempty" yaml:"description3,omitempty"`
	Description4                        *string                                   `xml:"description4,omitempty" json:"description4,omitempty" yaml:"description4,omitempty"`
	YoutubeVideoMediaIds                []*int64                                  `xml:"youtubeVideoMediaIds,omitempty" json:"youtubeVideoMediaIds,omitempty" yaml:"youtubeVideoMediaIds,omitempty"`
	ImageMediaIds                       []*int64                                  `xml:"imageMediaIds,omitempty" json:"imageMediaIds,omitempty" yaml:"imageMediaIds,omitempty"`
	UniversalAppBiddingStrategyGoalType *UniversalAppBiddingStrategyGoalType      `xml:"universalAppBiddingStrategyGoalType,omitempty" json:"universalAppBiddingStrategyGoalType,omitempty" yaml:"universalAppBiddingStrategyGoalType,omitempty"`
	YoutubeVideoMediaIdsOps             *ListOperations                           `xml:"youtubeVideoMediaIdsOps,omitempty" json:"youtubeVideoMediaIdsOps,omitempty" yaml:"youtubeVideoMediaIdsOps,omitempty"`
	ImageMediaIdsOps                    *ListOperations                           `xml:"imageMediaIdsOps,omitempty" json:"imageMediaIdsOps,omitempty" yaml:"imageMediaIdsOps,omitempty"`
	AdsPolicyDecisions                  []*UniversalAppCampaignAdsPolicyDecisions `xml:"adsPolicyDecisions,omitempty" json:"adsPolicyDecisions,omitempty" yaml:"adsPolicyDecisions,omitempty"`
	TypeAttrXSI                         string                                    `xml:"xsi:type,attr,omitempty"`
	TypeNamespace                       string                                    `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *UniversalAppCampaignSetting) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:UniversalAppCampaignSetting"
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

// Describes how unbranded pharma ads will be displayed.
type VanityPharma struct {
	VanityPharmaDisplayUrlMode *VanityPharmaDisplayUrlMode `xml:"vanityPharmaDisplayUrlMode,omitempty" json:"vanityPharmaDisplayUrlMode,omitempty" yaml:"vanityPharmaDisplayUrlMode,omitempty"`
	VanityPharmaText           *VanityPharmaText           `xml:"vanityPharmaText,omitempty" json:"vanityPharmaText,omitempty" yaml:"vanityPharmaText,omitempty"`
}

// Get was auto-generated from WSDL.
type Get struct {
	ServiceSelector *Selector `xml:"serviceSelector,omitempty" json:"serviceSelector,omitempty" yaml:"serviceSelector,omitempty"`
}

// GetResponse was auto-generated from WSDL.
type GetResponse struct {
	Rval *CampaignPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Mutate was auto-generated from WSDL.
type Mutate struct {
	Operations []*CampaignOperation `xml:"operations,omitempty" json:"operations,omitempty" yaml:"operations,omitempty"`
}

// MutateLabel was auto-generated from WSDL.
type MutateLabel struct {
	Operations []*CampaignLabelOperation `xml:"operations,omitempty" json:"operations,omitempty" yaml:"operations,omitempty"`
}

// MutateLabelResponse was auto-generated from WSDL.
type MutateLabelResponse struct {
	Rval *CampaignLabelReturnValue `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// MutateResponse was auto-generated from WSDL.
type MutateResponse struct {
	Rval *CampaignReturnValue `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
}

// Query was auto-generated from WSDL.
type Query struct {
	Query *string `xml:"query,omitempty" json:"query,omitempty" yaml:"query,omitempty"`
}

// QueryResponse was auto-generated from WSDL.
type QueryResponse struct {
	Rval *CampaignPage `xml:"rval,omitempty" json:"rval,omitempty" yaml:"rval,omitempty"`
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

// campaignServiceInterface implements the CampaignServiceInterface interface.
type campaignServiceInterface struct {
	cli *soap.Client
}

// Returns the list of campaigns that meet the selector criteria.
//                  @param serviceSelector the selector specifying
// the {@link Campaign}s to return.         @return A list of campaigns.
//         @throws ApiException if problems occurred while fetching
// campaign information.
func (p *campaignServiceInterface) Get(Get *Get) (*GetResponse, error) {
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

// Adds, updates, or removes campaigns.         <p class="note"><b>Note:</b>
// {@link CampaignOperation} does not support the         <code>REMOVE</code>
// operator. To remove a campaign, set its         {@link Campaign#status
// status} to {@code REMOVED}.</p>         @param operations A
// list of unique operations.         The same campaign cannot
// be specified in more than one operation.         @return The
// list of updated campaigns, returned in the same order as the
//         <code>operations</code> array.         @throws ApiException
// if problems occurred while updating campaign information.
func (p *campaignServiceInterface) Mutate(Mutate *Mutate) (*MutateResponse, error) {
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

// Adds labels to the {@linkplain Campaign campaign} or removes
// {@linkplain Label label}s from the         {@linkplain Campaign
// campaign}.         <p>Add - Apply an existing label to an existing
// {@linkplain Campaign campaign}. The         {@code campaignId}
// must reference an existing {@linkplain Campaign}. The {@code
// labelId} must         reference an existing {@linkplain Label
// label}.         <p>Remove - Removes the link between the specified
// {@linkplain Campaign campaign} and         {@linkplain Label
// label}.                  @param operations the operations to
// apply.         @return a list of {@linkplain CampaignLabel}s
// where each entry in the list is the result of         applying
// the operation in the input list with the same index. For an
//         add operation, the returned CampaignLabel contains the
// CampaignId and the LabelId.         In the case of a remove
// operation, the returned CampaignLabel will only have CampaignId.
//         @throws ApiException when there are one or more errors
// with the request.
func (p *campaignServiceInterface) MutateLabel(MutateLabel *MutateLabel) (*MutateLabelResponse, error) {
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

// Returns the list of campaigns that match the query.
//          @param query The SQL-like AWQL query string.
//   @return A list of campaigns.         @throws ApiException
// if problems occur while parsing the query or fetching campaign
//         information.
func (p *campaignServiceInterface) Query(Query string) (*QueryResponse, error) {
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
