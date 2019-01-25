// Code generated by gowsdl DO NOT EDIT.

package customersync

import (
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AuthenticationErrorReason string

const (
	AuthenticationErrorReasonAUTHENTICATION_FAILED AuthenticationErrorReason = "AUTHENTICATION_FAILED"

	AuthenticationErrorReasonCLIENT_CUSTOMER_ID_IS_REQUIRED AuthenticationErrorReason = "CLIENT_CUSTOMER_ID_IS_REQUIRED"

	AuthenticationErrorReasonCLIENT_EMAIL_REQUIRED AuthenticationErrorReason = "CLIENT_EMAIL_REQUIRED"

	AuthenticationErrorReasonCLIENT_CUSTOMER_ID_INVALID AuthenticationErrorReason = "CLIENT_CUSTOMER_ID_INVALID"

	AuthenticationErrorReasonCLIENT_EMAIL_INVALID AuthenticationErrorReason = "CLIENT_EMAIL_INVALID"

	AuthenticationErrorReasonCLIENT_EMAIL_FAILED_TO_AUTHENTICATE AuthenticationErrorReason = "CLIENT_EMAIL_FAILED_TO_AUTHENTICATE"

	AuthenticationErrorReasonCUSTOMER_NOT_FOUND AuthenticationErrorReason = "CUSTOMER_NOT_FOUND"

	AuthenticationErrorReasonGOOGLE_ACCOUNT_DELETED AuthenticationErrorReason = "GOOGLE_ACCOUNT_DELETED"

	AuthenticationErrorReasonGOOGLE_ACCOUNT_COOKIE_INVALID AuthenticationErrorReason = "GOOGLE_ACCOUNT_COOKIE_INVALID"

	AuthenticationErrorReasonFAILED_TO_AUTHENTICATE_GOOGLE_ACCOUNT AuthenticationErrorReason = "FAILED_TO_AUTHENTICATE_GOOGLE_ACCOUNT"

	AuthenticationErrorReasonGOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH AuthenticationErrorReason = "GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH"

	AuthenticationErrorReasonLOGIN_COOKIE_REQUIRED AuthenticationErrorReason = "LOGIN_COOKIE_REQUIRED"

	AuthenticationErrorReasonNOT_ADS_USER AuthenticationErrorReason = "NOT_ADS_USER"

	AuthenticationErrorReasonOAUTH_TOKEN_INVALID AuthenticationErrorReason = "OAUTH_TOKEN_INVALID"

	AuthenticationErrorReasonOAUTH_TOKEN_EXPIRED AuthenticationErrorReason = "OAUTH_TOKEN_EXPIRED"

	AuthenticationErrorReasonOAUTH_TOKEN_DISABLED AuthenticationErrorReason = "OAUTH_TOKEN_DISABLED"

	AuthenticationErrorReasonOAUTH_TOKEN_REVOKED AuthenticationErrorReason = "OAUTH_TOKEN_REVOKED"

	AuthenticationErrorReasonOAUTH_TOKEN_HEADER_INVALID AuthenticationErrorReason = "OAUTH_TOKEN_HEADER_INVALID"

	AuthenticationErrorReasonLOGIN_COOKIE_INVALID AuthenticationErrorReason = "LOGIN_COOKIE_INVALID"

	AuthenticationErrorReasonFAILED_TO_RETRIEVE_LOGIN_COOKIE AuthenticationErrorReason = "FAILED_TO_RETRIEVE_LOGIN_COOKIE"

	AuthenticationErrorReasonUSER_ID_INVALID AuthenticationErrorReason = "USER_ID_INVALID"
)

type AuthorizationErrorReason string

const (
	AuthorizationErrorReasonUNABLE_TO_AUTHORIZE AuthorizationErrorReason = "UNABLE_TO_AUTHORIZE"

	AuthorizationErrorReasonNO_ADWORDS_ACCOUNT_FOR_CUSTOMER AuthorizationErrorReason = "NO_ADWORDS_ACCOUNT_FOR_CUSTOMER"

	AuthorizationErrorReasonUSER_PERMISSION_DENIED AuthorizationErrorReason = "USER_PERMISSION_DENIED"

	AuthorizationErrorReasonEFFECTIVE_USER_PERMISSION_DENIED AuthorizationErrorReason = "EFFECTIVE_USER_PERMISSION_DENIED"

	AuthorizationErrorReasonCUSTOMER_NOT_ACTIVE AuthorizationErrorReason = "CUSTOMER_NOT_ACTIVE"

	AuthorizationErrorReasonUSER_HAS_READONLY_PERMISSION AuthorizationErrorReason = "USER_HAS_READONLY_PERMISSION"

	AuthorizationErrorReasonNO_CUSTOMER_FOUND AuthorizationErrorReason = "NO_CUSTOMER_FOUND"

	AuthorizationErrorReasonSERVICE_ACCESS_DENIED AuthorizationErrorReason = "SERVICE_ACCESS_DENIED"

	AuthorizationErrorReasonTWO_STEP_VERIFICATION_NOT_ENROLLED AuthorizationErrorReason = "TWO_STEP_VERIFICATION_NOT_ENROLLED"

	AuthorizationErrorReasonADVANCED_PROTECTION_NOT_ENROLLED AuthorizationErrorReason = "ADVANCED_PROTECTION_NOT_ENROLLED"
)

type ClientTermsErrorReason string

const (
	ClientTermsErrorReasonINCOMPLETE_SIGNUP_CURRENT_ADWORDS_TNC_NOT_AGREED ClientTermsErrorReason = "INCOMPLETE_SIGNUP_CURRENT_ADWORDS_TNC_NOT_AGREED"
)

type DatabaseErrorReason string

const (
	DatabaseErrorReasonCONCURRENT_MODIFICATION DatabaseErrorReason = "CONCURRENT_MODIFICATION"

	DatabaseErrorReasonPERMISSION_DENIED DatabaseErrorReason = "PERMISSION_DENIED"

	DatabaseErrorReasonACCESS_PROHIBITED DatabaseErrorReason = "ACCESS_PROHIBITED"

	DatabaseErrorReasonCAMPAIGN_PRODUCT_NOT_SUPPORTED DatabaseErrorReason = "CAMPAIGN_PRODUCT_NOT_SUPPORTED"

	DatabaseErrorReasonDUPLICATE_KEY DatabaseErrorReason = "DUPLICATE_KEY"

	DatabaseErrorReasonDATABASE_ERROR DatabaseErrorReason = "DATABASE_ERROR"

	DatabaseErrorReasonUNKNOWN DatabaseErrorReason = "UNKNOWN"
)

type DateErrorReason string

const (
	DateErrorReasonINVALID_FIELD_VALUES_IN_DATE DateErrorReason = "INVALID_FIELD_VALUES_IN_DATE"

	DateErrorReasonINVALID_FIELD_VALUES_IN_DATE_TIME DateErrorReason = "INVALID_FIELD_VALUES_IN_DATE_TIME"

	DateErrorReasonINVALID_STRING_DATE DateErrorReason = "INVALID_STRING_DATE"

	DateErrorReasonINVALID_STRING_DATE_RANGE DateErrorReason = "INVALID_STRING_DATE_RANGE"

	DateErrorReasonINVALID_STRING_DATE_TIME DateErrorReason = "INVALID_STRING_DATE_TIME"

	DateErrorReasonEARLIER_THAN_MINIMUM_DATE DateErrorReason = "EARLIER_THAN_MINIMUM_DATE"

	DateErrorReasonLATER_THAN_MAXIMUM_DATE DateErrorReason = "LATER_THAN_MAXIMUM_DATE"

	DateErrorReasonDATE_RANGE_MINIMUM_DATE_LATER_THAN_MAXIMUM_DATE DateErrorReason = "DATE_RANGE_MINIMUM_DATE_LATER_THAN_MAXIMUM_DATE"

	DateErrorReasonDATE_RANGE_MINIMUM_AND_MAXIMUM_DATES_BOTH_NULL DateErrorReason = "DATE_RANGE_MINIMUM_AND_MAXIMUM_DATES_BOTH_NULL"
)

type DistinctErrorReason string

const (
	DistinctErrorReasonDUPLICATE_ELEMENT DistinctErrorReason = "DUPLICATE_ELEMENT"

	DistinctErrorReasonDUPLICATE_TYPE DistinctErrorReason = "DUPLICATE_TYPE"
)

type IdErrorReason string

const (
	IdErrorReasonNOT_FOUND IdErrorReason = "NOT_FOUND"
)

type InternalApiErrorReason string

const (
	InternalApiErrorReasonUNEXPECTED_INTERNAL_API_ERROR InternalApiErrorReason = "UNEXPECTED_INTERNAL_API_ERROR"

	InternalApiErrorReasonTRANSIENT_ERROR InternalApiErrorReason = "TRANSIENT_ERROR"

	InternalApiErrorReasonUNKNOWN InternalApiErrorReason = "UNKNOWN"

	InternalApiErrorReasonDOWNTIME InternalApiErrorReason = "DOWNTIME"

	InternalApiErrorReasonERROR_GENERATING_RESPONSE InternalApiErrorReason = "ERROR_GENERATING_RESPONSE"
)

type NotEmptyErrorReason string

const (
	NotEmptyErrorReasonEMPTY_LIST NotEmptyErrorReason = "EMPTY_LIST"
)

type OperationAccessDeniedReason string

const (
	OperationAccessDeniedReasonACTION_NOT_PERMITTED OperationAccessDeniedReason = "ACTION_NOT_PERMITTED"

	OperationAccessDeniedReasonADD_OPERATION_NOT_PERMITTED OperationAccessDeniedReason = "ADD_OPERATION_NOT_PERMITTED"

	OperationAccessDeniedReasonREMOVE_OPERATION_NOT_PERMITTED OperationAccessDeniedReason = "REMOVE_OPERATION_NOT_PERMITTED"

	OperationAccessDeniedReasonSET_OPERATION_NOT_PERMITTED OperationAccessDeniedReason = "SET_OPERATION_NOT_PERMITTED"

	OperationAccessDeniedReasonMUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT OperationAccessDeniedReason = "MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT"

	OperationAccessDeniedReasonOPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE OperationAccessDeniedReason = "OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE"

	OperationAccessDeniedReasonADD_AS_REMOVED_NOT_PERMITTED OperationAccessDeniedReason = "ADD_AS_REMOVED_NOT_PERMITTED"

	OperationAccessDeniedReasonOPERATION_NOT_PERMITTED_FOR_REMOVED_ENTITY OperationAccessDeniedReason = "OPERATION_NOT_PERMITTED_FOR_REMOVED_ENTITY"

	OperationAccessDeniedReasonOPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE OperationAccessDeniedReason = "OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE"

	OperationAccessDeniedReasonUNKNOWN OperationAccessDeniedReason = "UNKNOWN"
)

type OperatorErrorReason string

const (
	OperatorErrorReasonOPERATOR_NOT_SUPPORTED OperatorErrorReason = "OPERATOR_NOT_SUPPORTED"
)

type QuotaCheckErrorReason string

const (
	QuotaCheckErrorReasonINVALID_TOKEN_HEADER QuotaCheckErrorReason = "INVALID_TOKEN_HEADER"

	QuotaCheckErrorReasonACCOUNT_DELINQUENT QuotaCheckErrorReason = "ACCOUNT_DELINQUENT"

	QuotaCheckErrorReasonACCOUNT_INACCESSIBLE QuotaCheckErrorReason = "ACCOUNT_INACCESSIBLE"

	QuotaCheckErrorReasonACCOUNT_INACTIVE QuotaCheckErrorReason = "ACCOUNT_INACTIVE"

	QuotaCheckErrorReasonINCOMPLETE_SIGNUP QuotaCheckErrorReason = "INCOMPLETE_SIGNUP"

	QuotaCheckErrorReasonDEVELOPER_TOKEN_NOT_APPROVED QuotaCheckErrorReason = "DEVELOPER_TOKEN_NOT_APPROVED"

	QuotaCheckErrorReasonTERMS_AND_CONDITIONS_NOT_SIGNED QuotaCheckErrorReason = "TERMS_AND_CONDITIONS_NOT_SIGNED"

	QuotaCheckErrorReasonMONTHLY_BUDGET_REACHED QuotaCheckErrorReason = "MONTHLY_BUDGET_REACHED"

	QuotaCheckErrorReasonQUOTA_EXCEEDED QuotaCheckErrorReason = "QUOTA_EXCEEDED"
)

type RangeErrorReason string

const (
	RangeErrorReasonTOO_LOW RangeErrorReason = "TOO_LOW"

	RangeErrorReasonTOO_HIGH RangeErrorReason = "TOO_HIGH"
)

type RateExceededErrorReason string

const (
	RateExceededErrorReasonRATE_EXCEEDED RateExceededErrorReason = "RATE_EXCEEDED"
)

type ReadOnlyErrorReason string

const (
	ReadOnlyErrorReasonREAD_ONLY ReadOnlyErrorReason = "READ_ONLY"
)

type RejectedErrorReason string

const (
	RejectedErrorReasonUNKNOWN_VALUE RejectedErrorReason = "UNKNOWN_VALUE"
)

type RequestErrorReason string

const (
	RequestErrorReasonUNKNOWN RequestErrorReason = "UNKNOWN"

	RequestErrorReasonINVALID_INPUT RequestErrorReason = "INVALID_INPUT"

	RequestErrorReasonUNSUPPORTED_VERSION RequestErrorReason = "UNSUPPORTED_VERSION"
)

type RequiredErrorReason string

const (
	RequiredErrorReasonREQUIRED RequiredErrorReason = "REQUIRED"
)

type SizeLimitErrorReason string

const (
	SizeLimitErrorReasonREQUEST_SIZE_LIMIT_EXCEEDED SizeLimitErrorReason = "REQUEST_SIZE_LIMIT_EXCEEDED"

	SizeLimitErrorReasonRESPONSE_SIZE_LIMIT_EXCEEDED SizeLimitErrorReason = "RESPONSE_SIZE_LIMIT_EXCEEDED"

	SizeLimitErrorReasonINTERNAL_STORAGE_ERROR SizeLimitErrorReason = "INTERNAL_STORAGE_ERROR"

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	SizeLimitErrorReasonUNKNOWN SizeLimitErrorReason = "UNKNOWN"
)

type StringFormatErrorReason string

const (
	StringFormatErrorReasonUNKNOWN StringFormatErrorReason = "UNKNOWN"

	StringFormatErrorReasonILLEGAL_CHARS StringFormatErrorReason = "ILLEGAL_CHARS"

	StringFormatErrorReasonINVALID_FORMAT StringFormatErrorReason = "INVALID_FORMAT"
)

type StringLengthErrorReason string

const (
	StringLengthErrorReasonTOO_SHORT StringLengthErrorReason = "TOO_SHORT"

	StringLengthErrorReasonTOO_LONG StringLengthErrorReason = "TOO_LONG"
)

type ApiError struct {
	FieldPath string `xml:"fieldPath,omitempty"`

	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty"`

	Trigger string `xml:"trigger,omitempty"`

	ErrorString string `xml:"errorString,omitempty"`

	//
	// Indicates that this instance is a subtype of ApiError.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	ApiErrorType string `xml:"ApiError.Type,omitempty"`
}

type ApiException struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201809 ApiExceptionFault"`

	*ApplicationException

	Errors []*ApiError `xml:"errors,omitempty"`
}

type ApplicationException struct {
	Message string `xml:"message,omitempty"`

	//
	// Indicates that this instance is a subtype of ApplicationException.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	ApplicationExceptionType string `xml:"ApplicationException.Type,omitempty"`
}

type AuthenticationError struct {
	*ApiError

	Reason *AuthenticationErrorReason `xml:"reason,omitempty"`
}

type AuthorizationError struct {
	*ApiError

	Reason *AuthorizationErrorReason `xml:"reason,omitempty"`
}

type ClientTermsError struct {
	*ApiError

	Reason *ClientTermsErrorReason `xml:"reason,omitempty"`
}

type DatabaseError struct {
	*ApiError

	Reason *DatabaseErrorReason `xml:"reason,omitempty"`
}

type DateError struct {
	*ApiError

	Reason *DateErrorReason `xml:"reason,omitempty"`
}

type DateTimeRange struct {
	Min string `xml:"min,omitempty"`

	Max string `xml:"max,omitempty"`
}

type DistinctError struct {
	*ApiError

	Reason *DistinctErrorReason `xml:"reason,omitempty"`
}

type FieldPathElement struct {
	Field string `xml:"field,omitempty"`

	Index int32 `xml:"index,omitempty"`
}

type IdError struct {
	*ApiError

	Reason *IdErrorReason `xml:"reason,omitempty"`
}

type InternalApiError struct {
	*ApiError

	Reason *InternalApiErrorReason `xml:"reason,omitempty"`
}

type NotEmptyError struct {
	*ApiError

	Reason *NotEmptyErrorReason `xml:"reason,omitempty"`
}

type OperationAccessDenied struct {
	*ApiError

	Reason *OperationAccessDeniedReason `xml:"reason,omitempty"`
}

type OperatorError struct {
	*ApiError

	Reason *OperatorErrorReason `xml:"reason,omitempty"`
}

type QuotaCheckError struct {
	*ApiError

	Reason *QuotaCheckErrorReason `xml:"reason,omitempty"`
}

type RangeError struct {
	*ApiError

	Reason *RangeErrorReason `xml:"reason,omitempty"`
}

type RateExceededError struct {
	*ApiError

	Reason *RateExceededErrorReason `xml:"reason,omitempty"`

	RateName string `xml:"rateName,omitempty"`

	RateScope string `xml:"rateScope,omitempty"`

	RetryAfterSeconds int32 `xml:"retryAfterSeconds,omitempty"`
}

type ReadOnlyError struct {
	*ApiError

	Reason *ReadOnlyErrorReason `xml:"reason,omitempty"`
}

type RejectedError struct {
	*ApiError

	Reason *RejectedErrorReason `xml:"reason,omitempty"`
}

type RequestError struct {
	*ApiError

	Reason *RequestErrorReason `xml:"reason,omitempty"`
}

type RequiredError struct {
	*ApiError

	Reason *RequiredErrorReason `xml:"reason,omitempty"`
}

type SizeLimitError struct {
	*ApiError

	Reason *SizeLimitErrorReason `xml:"reason,omitempty"`
}

type SoapHeader struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201809 RequestHeader"`

	ClientCustomerId string `xml:"clientCustomerId,omitempty"`

	DeveloperToken string `xml:"developerToken,omitempty"`

	UserAgent string `xml:"userAgent,omitempty"`

	ValidateOnly bool `xml:"validateOnly,omitempty"`

	PartialFailure bool `xml:"partialFailure,omitempty"`
}

type SoapResponseHeader struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201809 ResponseHeader"`

	RequestId string `xml:"requestId,omitempty"`

	ServiceName string `xml:"serviceName,omitempty"`

	MethodName string `xml:"methodName,omitempty"`

	Operations int64 `xml:"operations,omitempty"`

	ResponseTime int64 `xml:"responseTime,omitempty"`
}

type StringFormatError struct {
	*ApiError

	Reason *StringFormatErrorReason `xml:"reason,omitempty"`
}

type StringLengthError struct {
	*ApiError

	Reason *StringLengthErrorReason `xml:"reason,omitempty"`
}

//
// An enum used to classify the types of changes that have been made to an adgroup/campaign during a
// specified date range. This only refers to the field of the entity itself, and not its children.
//
// <p>For example, if an AdGroup name changed, this status would be FIELDS_CHANGED, but if only bids
// on keywords belonging an AdGroup were changed this status would be FIELDS_UNCHANGED.
//
type ChangeStatus string

const (

	//
	// The fields of this entity have not changed, but there may still be changes to its children.
	//
	ChangeStatusFIELDS_UNCHANGED ChangeStatus = "FIELDS_UNCHANGED"

	//
	// The fields of this entity have changed, for example the name of an adgroup.
	//
	ChangeStatusFIELDS_CHANGED ChangeStatus = "FIELDS_CHANGED"

	//
	// This entity was created during the time frame we're looking at. We will not enumerate all of
	// the individual changes to this entity and its children. Instead it should be loaded from the
	// appropriate service.
	//
	ChangeStatusNEW ChangeStatus = "NEW"
)

type CustomerSyncErrorReason string

const (

	//
	// The request attempted to access a campaign that either does not exist or belongs to a
	// different account.
	//
	CustomerSyncErrorReasonINVALID_CAMPAIGN_ID CustomerSyncErrorReason = "INVALID_CAMPAIGN_ID"

	//
	// The request attempted to access a feed that either does not exist or belongs to a different
	// account.
	//
	CustomerSyncErrorReasonINVALID_FEED_ID CustomerSyncErrorReason = "INVALID_FEED_ID"

	//
	// The request asked for an invalid date range
	//
	CustomerSyncErrorReasonINVALID_DATE_RANGE CustomerSyncErrorReason = "INVALID_DATE_RANGE"

	//
	// There have been too many changes to sync the campaign over the requested date/time range. To
	// avoid this error, try specifying a smaller date/time range. If this does not work, you should
	// assume that everything has changed and retrieve the objects using their respective services.
	//
	CustomerSyncErrorReasonTOO_MANY_CHANGES CustomerSyncErrorReason = "TOO_MANY_CHANGES"
)

type Get struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/ch/v201809 get"`

	//
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Selector *CustomerSyncSelector `xml:"selector,omitempty"`
}

type GetResponse struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/ch/v201809 getResponse"`

	Rval *CustomerChangeData `xml:"rval,omitempty"`
}

type AdGroupChangeData struct {

	//
	// The ad group ID.
	//
	AdGroupId int64 `xml:"adGroupId,omitempty"`

	//
	// Whether or not the fields of this adgroup have changed, for example the AdGroup name. Changes
	// to the Ads and Criteria are enumerated in their respective lists and will not be reflected in
	// this status.
	//
	AdGroupChangeStatus *ChangeStatus `xml:"adGroupChangeStatus,omitempty"`

	//
	// The IDs of any changed ads of this ad group. This includes ads that have been deleted.
	//
	ChangedAds []int64 `xml:"changedAds,omitempty"`

	//
	// The IDs of any changed criterion of this ad group.
	//
	ChangedCriteria []int64 `xml:"changedCriteria,omitempty"`

	//
	// The IDs of any deleted criterion of this ad group.
	//
	RemovedCriteria []int64 `xml:"removedCriteria,omitempty"`

	//
	// A list of feed IDs for AdGroupFeeds that have been changed in this ad group. If an AdGroupFeed
	// is deleted after a modification, it will not be included in this list.
	//
	ChangedFeeds []int64 `xml:"changedFeeds,omitempty"`

	//
	// A list of feed IDs for AdGroupFeeds that have been removed from the ad group.
	//
	RemovedFeeds []int64 `xml:"removedFeeds,omitempty"`

	//
	// Set of campaign criterion that have a bid modifier override at ad group level. If the
	// associated bid modifier override is deleted after a modification, it will not be included in
	// this list.
	//
	ChangedAdGroupBidModifierCriteria []int64 `xml:"changedAdGroupBidModifierCriteria,omitempty"`

	//
	// Set of campaign criterion whose bid modifier override at ad group level has been removed.
	//
	RemovedAdGroupBidModifierCriteria []int64 `xml:"removedAdGroupBidModifierCriteria,omitempty"`
}

type CampaignChangeData struct {

	//
	// The campaign ID.
	//
	CampaignId int64 `xml:"campaignId,omitempty"`

	//
	// Whether or not the fields of this campaign have changed. Changes to campaign level criteria and
	// ad extensions are enumerated in their respective lists and will not be reflected in this
	// status.
	//
	CampaignChangeStatus *ChangeStatus `xml:"campaignChangeStatus,omitempty"`

	//
	// A list of change information for all changed adgroups belonging to the campaign.
	//
	ChangedAdGroups []*AdGroupChangeData `xml:"changedAdGroups,omitempty"`

	//
	// A list of criteria IDs that have been added as campaign criteria. This list includes any
	// criteria that can be downloaded using CampaignCriterionService.
	//
	AddedCampaignCriteria []int64 `xml:"addedCampaignCriteria,omitempty"`

	//
	// A list of criteria IDs that have been deleted as campaign criteria. This list includes any
	// criteria that can be downloaded using CampaignCriterionService.
	//
	RemovedCampaignCriteria []int64 `xml:"removedCampaignCriteria,omitempty"`

	//
	// A list of feed IDs for CampaignFeeds that have been changed in this campaign. If a CampaignFeed
	// is deleted after a modification, it will not be included in this list.
	//
	ChangedFeeds []int64 `xml:"changedFeeds,omitempty"`

	//
	// A list of feed IDs for CampaignFeeds that have been removed from the campaign.
	//
	RemovedFeeds []int64 `xml:"removedFeeds,omitempty"`
}

type CustomerChangeData struct {

	//
	// A list of campaign changes for the customer, as specified by the selector. If a campaign is
	// included in the selector it will be included in this list, even if the campaign did not change.
	//
	ChangedCampaigns []*CampaignChangeData `xml:"changedCampaigns,omitempty"`

	//
	// A list of feed changes for the customer as specified in the selector. If a feed is included in
	// the selector then it will be included in this list, even if the feed did not change.
	//
	ChangedFeeds []*FeedChangeData `xml:"changedFeeds,omitempty"`

	//
	// The timestamp for the last changed processed for this customer. It is important that this
	// timestamp be used for subsequent requests to avoid missing any changes.
	//
	LastChangeTimestamp string `xml:"lastChangeTimestamp,omitempty"`
}

type CustomerSyncError struct {
	*ApiError

	Reason *CustomerSyncErrorReason `xml:"reason,omitempty"`

	CampaignId int64 `xml:"campaignId,omitempty"`
}

type CustomerSyncSelector struct {

	//
	// Only return entities that have changed during the specified time range. String Format: yyyyMMdd
	// HHmmss <Timezone ID> (for example, 20100609 150223 America/New_York). See the <a
	// href="https://developers.google.com/adwords/api/docs/appendix/timezones">Timezones</a> page for
	// the complete list of Timezone IDs.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	DateTimeRange *DateTimeRange `xml:"dateTimeRange,omitempty"`

	//
	// Return entities belonging to these campaigns.
	// <span class="constraint ContentsDistinct">This field must contain distinct elements.</span>
	//
	CampaignIds []int64 `xml:"campaignIds,omitempty"`

	//
	// Return entities belonging to these feeds.
	// <span class="constraint ContentsDistinct">This field must contain distinct elements.</span>
	//
	FeedIds []int64 `xml:"feedIds,omitempty"`
}

type FeedChangeData struct {

	//
	// The feed ID.
	//
	FeedId int64 `xml:"feedId,omitempty"`

	//
	// Whether or not the fields of this feed have changed.
	//
	FeedChangeStatus *ChangeStatus `xml:"feedChangeStatus,omitempty"`

	//
	// A list of feed item IDs that have been added or modified within the the feed. If a feed item is
	// deleted after a modification, it will not be included in this list.
	//
	ChangedFeedItems []int64 `xml:"changedFeedItems,omitempty"`

	//
	// A list feed item IDs that have been removed from the feed.
	//
	RemovedFeedItems []int64 `xml:"removedFeedItems,omitempty"`
}

type CustomerSyncServiceInterface interface {

	// Error can be either of the following types:
	//
	//   - ApiException
	/*
	   Returns information about changed entities inside a customer's account.

	   @param selector Specifies the filter for selecting changehistory events for a customer.
	   @return A Customer->Campaign->AdGroup hierarchy containing information about the objects
	   changed at each level. All Campaigns that are requested in the selector will be returned,
	   regardless of whether or not they have changed, but unchanged AdGroups will be ignored.
	*/
	Get(request *Get) (*GetResponse, error)
}

type customerSyncServiceInterface struct {
	client *soap.Client
}

func NewCustomerSyncServiceInterface(client *soap.Client) CustomerSyncServiceInterface {
	return &customerSyncServiceInterface{
		client: client,
	}
}

func (service *customerSyncServiceInterface) Get(request *Get) (*GetResponse, error) {
	response := new(GetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
