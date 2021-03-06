package CampaignExtensionSettingService

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

//
// The available application stores for app extensions.
//
type AppFeedItemAppStore string

const (
	AppFeedItemAppStoreAPPLE_ITUNES AppFeedItemAppStore = "APPLE_ITUNES"

	AppFeedItemAppStoreGOOGLE_PLAY AppFeedItemAppStore = "GOOGLE_PLAY"

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	AppFeedItemAppStoreUNKNOWN AppFeedItemAppStore = "UNKNOWN"
)

//
// The single reason for the authentication failure.
//
type AuthenticationErrorReason string

const (

	//
	// Authentication of the request failed.
	//
	AuthenticationErrorReasonAUTHENTICATION_FAILED AuthenticationErrorReason = "AUTHENTICATION_FAILED"

	//
	// Client Customer Id is required if CustomerIdMode is set to CLIENT_EXTERNAL_CUSTOMER_ID.
	// Starting version V201409 ClientCustomerId will be required for all requests except
	// for {@link CustomerService#get}
	//
	AuthenticationErrorReasonCLIENT_CUSTOMER_ID_IS_REQUIRED AuthenticationErrorReason = "CLIENT_CUSTOMER_ID_IS_REQUIRED"

	//
	// Client Email is required if CustomerIdMode is set to CLIENT_EXTERNAL_EMAIL_FIELD.
	//
	AuthenticationErrorReasonCLIENT_EMAIL_REQUIRED AuthenticationErrorReason = "CLIENT_EMAIL_REQUIRED"

	//
	// Client customer Id is not a number.
	//
	AuthenticationErrorReasonCLIENT_CUSTOMER_ID_INVALID AuthenticationErrorReason = "CLIENT_CUSTOMER_ID_INVALID"

	//
	// Client customer Id is not a number.
	//
	AuthenticationErrorReasonCLIENT_EMAIL_INVALID AuthenticationErrorReason = "CLIENT_EMAIL_INVALID"

	//
	// Client email is not a valid customer email.
	//
	AuthenticationErrorReasonCLIENT_EMAIL_FAILED_TO_AUTHENTICATE AuthenticationErrorReason = "CLIENT_EMAIL_FAILED_TO_AUTHENTICATE"

	//
	// No customer found for the customer id provided in the header.
	//
	AuthenticationErrorReasonCUSTOMER_NOT_FOUND AuthenticationErrorReason = "CUSTOMER_NOT_FOUND"

	//
	// Client's Google Account is deleted.
	//
	AuthenticationErrorReasonGOOGLE_ACCOUNT_DELETED AuthenticationErrorReason = "GOOGLE_ACCOUNT_DELETED"

	//
	// Google account login token in the cookie is invalid.
	//
	AuthenticationErrorReasonGOOGLE_ACCOUNT_COOKIE_INVALID AuthenticationErrorReason = "GOOGLE_ACCOUNT_COOKIE_INVALID"

	//
	// problem occurred during Google account authentication.
	//
	AuthenticationErrorReasonFAILED_TO_AUTHENTICATE_GOOGLE_ACCOUNT AuthenticationErrorReason = "FAILED_TO_AUTHENTICATE_GOOGLE_ACCOUNT"

	//
	// The user in the google account login token does not match the UserId in the cookie.
	//
	AuthenticationErrorReasonGOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH AuthenticationErrorReason = "GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH"

	//
	// Login cookie is required for authentication.
	//
	AuthenticationErrorReasonLOGIN_COOKIE_REQUIRED AuthenticationErrorReason = "LOGIN_COOKIE_REQUIRED"

	//
	// User in the cookie is not a valid Ads user.
	//
	AuthenticationErrorReasonNOT_ADS_USER AuthenticationErrorReason = "NOT_ADS_USER"

	//
	// Oauth token in the header is not valid.
	//
	AuthenticationErrorReasonOAUTH_TOKEN_INVALID AuthenticationErrorReason = "OAUTH_TOKEN_INVALID"

	//
	// Oauth token in the header has expired.
	//
	AuthenticationErrorReasonOAUTH_TOKEN_EXPIRED AuthenticationErrorReason = "OAUTH_TOKEN_EXPIRED"

	//
	// Oauth token in the header has been disabled.
	//
	AuthenticationErrorReasonOAUTH_TOKEN_DISABLED AuthenticationErrorReason = "OAUTH_TOKEN_DISABLED"

	//
	// Oauth token in the header has been revoked.
	//
	AuthenticationErrorReasonOAUTH_TOKEN_REVOKED AuthenticationErrorReason = "OAUTH_TOKEN_REVOKED"

	//
	// Oauth token HTTP header is malformed.
	//
	AuthenticationErrorReasonOAUTH_TOKEN_HEADER_INVALID AuthenticationErrorReason = "OAUTH_TOKEN_HEADER_INVALID"

	//
	// Login cookie is not valid.
	//
	AuthenticationErrorReasonLOGIN_COOKIE_INVALID AuthenticationErrorReason = "LOGIN_COOKIE_INVALID"

	//
	// Failed to decrypt the login cookie.
	//
	AuthenticationErrorReasonFAILED_TO_RETRIEVE_LOGIN_COOKIE AuthenticationErrorReason = "FAILED_TO_RETRIEVE_LOGIN_COOKIE"

	//
	// User Id in the header is not a valid id.
	//
	AuthenticationErrorReasonUSER_ID_INVALID AuthenticationErrorReason = "USER_ID_INVALID"
)

//
// The reasons for the authorization error.
//
type AuthorizationErrorReason string

const (

	//
	// Could not complete authorization due to an internal problem.
	//
	AuthorizationErrorReasonUNABLE_TO_AUTHORIZE AuthorizationErrorReason = "UNABLE_TO_AUTHORIZE"

	//
	// Customer has no AdWords account.
	//
	AuthorizationErrorReasonNO_ADWORDS_ACCOUNT_FOR_CUSTOMER AuthorizationErrorReason = "NO_ADWORDS_ACCOUNT_FOR_CUSTOMER"

	//
	// User doesn't have permission to access customer.
	//
	AuthorizationErrorReasonUSER_PERMISSION_DENIED AuthorizationErrorReason = "USER_PERMISSION_DENIED"

	//
	// Effective user doesn't have permission to access this customer.
	//
	AuthorizationErrorReasonEFFECTIVE_USER_PERMISSION_DENIED AuthorizationErrorReason = "EFFECTIVE_USER_PERMISSION_DENIED"

	//
	// Access denied since the customer is not active.
	//
	AuthorizationErrorReasonCUSTOMER_NOT_ACTIVE AuthorizationErrorReason = "CUSTOMER_NOT_ACTIVE"

	//
	// User has read-only permission cannot mutate.
	//
	AuthorizationErrorReasonUSER_HAS_READONLY_PERMISSION AuthorizationErrorReason = "USER_HAS_READONLY_PERMISSION"

	//
	// No customer found.
	//
	AuthorizationErrorReasonNO_CUSTOMER_FOUND AuthorizationErrorReason = "NO_CUSTOMER_FOUND"

	//
	// Developer doesn't have permission to access service.
	//
	AuthorizationErrorReasonSERVICE_ACCESS_DENIED AuthorizationErrorReason = "SERVICE_ACCESS_DENIED"
)

//
// Enums for the various reasons an error can be thrown as a result of
// ClientTerms violation.
//
type ClientTermsErrorReason string

const (

	//
	// Customer has not agreed to the latest AdWords Terms & Conditions
	//
	ClientTermsErrorReasonINCOMPLETE_SIGNUP_CURRENT_ADWORDS_TNC_NOT_AGREED ClientTermsErrorReason = "INCOMPLETE_SIGNUP_CURRENT_ADWORDS_TNC_NOT_AGREED"
)

//
// The reasons for the target error.
//
type CollectionSizeErrorReason string

const (
	CollectionSizeErrorReasonTOO_FEW CollectionSizeErrorReason = "TOO_FEW"

	CollectionSizeErrorReasonTOO_MANY CollectionSizeErrorReason = "TOO_MANY"
)

//
// The types of criteria.
//
type CriterionType string

const (

	//
	// Content label for exclusion.
	//
	CriterionTypeCONTENT_LABEL CriterionType = "CONTENT_LABEL"

	//
	// Keyword. e.g. 'mars cruise'
	//
	CriterionTypeKEYWORD CriterionType = "KEYWORD"

	//
	// Placement. aka Website. e.g. 'www.flowers4sale.com'
	//
	CriterionTypePLACEMENT CriterionType = "PLACEMENT"

	//
	// Vertical, e.g. 'category::Animals>Pets'  This is for vertical targeting on the content
	// network.
	//
	CriterionTypeVERTICAL CriterionType = "VERTICAL"

	//
	// User lists, are links to sets of users defined by the advertiser.
	//
	CriterionTypeUSER_LIST CriterionType = "USER_LIST"

	//
	// User interests, categories of interests the user is interested in.
	//
	CriterionTypeUSER_INTEREST CriterionType = "USER_INTEREST"

	//
	// Mobile applications to target.
	//
	CriterionTypeMOBILE_APPLICATION CriterionType = "MOBILE_APPLICATION"

	//
	// Mobile application categories to target.
	//
	CriterionTypeMOBILE_APP_CATEGORY CriterionType = "MOBILE_APP_CATEGORY"

	//
	// Product partition (product group) in a shopping campaign.
	//
	CriterionTypePRODUCT_PARTITION CriterionType = "PRODUCT_PARTITION"

	//
	// IP addresses to exclude.
	//
	CriterionTypeIP_BLOCK CriterionType = "IP_BLOCK"

	//
	// Webpages of an advertiser's website to target.
	//
	CriterionTypeWEBPAGE CriterionType = "WEBPAGE"

	//
	// Languages to target.
	//
	CriterionTypeLANGUAGE CriterionType = "LANGUAGE"

	//
	// Geographic regions to target.
	//
	CriterionTypeLOCATION CriterionType = "LOCATION"

	//
	// Age Range to exclude.
	//
	CriterionTypeAGE_RANGE CriterionType = "AGE_RANGE"

	//
	// Mobile carriers to target.
	//
	CriterionTypeCARRIER CriterionType = "CARRIER"

	//
	// Mobile operating system versions to target.
	//
	CriterionTypeOPERATING_SYSTEM_VERSION CriterionType = "OPERATING_SYSTEM_VERSION"

	//
	// Mobile devices to target.
	//
	CriterionTypeMOBILE_DEVICE CriterionType = "MOBILE_DEVICE"

	//
	// Gender to exclude.
	//
	CriterionTypeGENDER CriterionType = "GENDER"

	//
	// Parent to target and exclude.
	//
	CriterionTypePARENT CriterionType = "PARENT"

	//
	// Proximity (area within a radius) to target.
	//
	CriterionTypePROXIMITY CriterionType = "PROXIMITY"

	//
	// Platforms to target.
	//
	CriterionTypePLATFORM CriterionType = "PLATFORM"

	//
	// Representing preferred content bid modifier.
	//
	CriterionTypePREFERRED_CONTENT CriterionType = "PREFERRED_CONTENT"

	//
	// AdSchedule or specific days and time intervals to target.
	//
	CriterionTypeAD_SCHEDULE CriterionType = "AD_SCHEDULE"

	//
	// Targeting based on location groups.
	//
	CriterionTypeLOCATION_GROUPS CriterionType = "LOCATION_GROUPS"

	//
	// Scope of products. Contains a list of product dimensions, all of which a product has to match
	// to be included in the campaign.
	//
	CriterionTypePRODUCT_SCOPE CriterionType = "PRODUCT_SCOPE"

	//
	// YouTube video to target.
	//
	CriterionTypeYOUTUBE_VIDEO CriterionType = "YOUTUBE_VIDEO"

	//
	// YouTube channel to target.
	//
	CriterionTypeYOUTUBE_CHANNEL CriterionType = "YOUTUBE_CHANNEL"

	//
	// Enables advertisers to target paid apps.
	//
	CriterionTypeAPP_PAYMENT_MODEL CriterionType = "APP_PAYMENT_MODEL"

	//
	// Income range to target and exclude.
	//
	CriterionTypeINCOME_RANGE CriterionType = "INCOME_RANGE"

	//
	// Interaction type to bid modify.
	//
	CriterionTypeINTERACTION_TYPE CriterionType = "INTERACTION_TYPE"

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	CriterionTypeUNKNOWN CriterionType = "UNKNOWN"
)

type CriterionErrorReason string

const (

	//
	// Concrete type of criterion is required for ADD and SET operations.
	//
	CriterionErrorReasonCONCRETE_TYPE_REQUIRED CriterionErrorReason = "CONCRETE_TYPE_REQUIRED"

	//
	// The category requested for exclusion is invalid.
	//
	CriterionErrorReasonINVALID_EXCLUDED_CATEGORY CriterionErrorReason = "INVALID_EXCLUDED_CATEGORY"

	//
	// Invalid keyword criteria text.
	//
	CriterionErrorReasonINVALID_KEYWORD_TEXT CriterionErrorReason = "INVALID_KEYWORD_TEXT"

	//
	// Keyword text should be less than 80 chars.
	//
	CriterionErrorReasonKEYWORD_TEXT_TOO_LONG CriterionErrorReason = "KEYWORD_TEXT_TOO_LONG"

	//
	// Keyword text has too many words.
	//
	CriterionErrorReasonKEYWORD_HAS_TOO_MANY_WORDS CriterionErrorReason = "KEYWORD_HAS_TOO_MANY_WORDS"

	//
	// Keyword text has invalid characters or symbols.
	//
	CriterionErrorReasonKEYWORD_HAS_INVALID_CHARS CriterionErrorReason = "KEYWORD_HAS_INVALID_CHARS"

	//
	// Invalid placement URL.
	//
	CriterionErrorReasonINVALID_PLACEMENT_URL CriterionErrorReason = "INVALID_PLACEMENT_URL"

	//
	// Invalid user list criterion.
	//
	CriterionErrorReasonINVALID_USER_LIST CriterionErrorReason = "INVALID_USER_LIST"

	//
	// Invalid user interest criterion.
	//
	CriterionErrorReasonINVALID_USER_INTEREST CriterionErrorReason = "INVALID_USER_INTEREST"

	//
	// Placement URL has wrong format.
	//
	CriterionErrorReasonINVALID_FORMAT_FOR_PLACEMENT_URL CriterionErrorReason = "INVALID_FORMAT_FOR_PLACEMENT_URL"

	//
	// Placement URL is too long.
	//
	CriterionErrorReasonPLACEMENT_URL_IS_TOO_LONG CriterionErrorReason = "PLACEMENT_URL_IS_TOO_LONG"

	//
	// Indicates the URL contains an illegal character.
	//
	CriterionErrorReasonPLACEMENT_URL_HAS_ILLEGAL_CHAR CriterionErrorReason = "PLACEMENT_URL_HAS_ILLEGAL_CHAR"

	//
	// Indicates the URL contains multiple comma separated URLs.
	//
	CriterionErrorReasonPLACEMENT_URL_HAS_MULTIPLE_SITES_IN_LINE CriterionErrorReason = "PLACEMENT_URL_HAS_MULTIPLE_SITES_IN_LINE"

	//
	// Indicates the domain is blacklisted.
	//
	CriterionErrorReasonPLACEMENT_IS_NOT_AVAILABLE_FOR_TARGETING_OR_EXCLUSION CriterionErrorReason = "PLACEMENT_IS_NOT_AVAILABLE_FOR_TARGETING_OR_EXCLUSION"

	//
	// Invalid vertical path.
	//
	CriterionErrorReasonINVALID_VERTICAL_PATH CriterionErrorReason = "INVALID_VERTICAL_PATH"

	//
	// Indicates the placement is a YouTube vertical channel, which is no longer supported.
	//
	CriterionErrorReasonYOUTUBE_VERTICAL_CHANNEL_DEPRECATED CriterionErrorReason = "YOUTUBE_VERTICAL_CHANNEL_DEPRECATED"

	//
	// Indicates the placement is a YouTube demographic channel, which is no longer supported.
	//
	CriterionErrorReasonYOUTUBE_DEMOGRAPHIC_CHANNEL_DEPRECATED CriterionErrorReason = "YOUTUBE_DEMOGRAPHIC_CHANNEL_DEPRECATED"

	//
	// YouTube urls are not supported in Placement criterion. Use YouTubeChannel and
	// YouTubeVideo criterion instead.
	//
	CriterionErrorReasonYOUTUBE_URL_UNSUPPORTED CriterionErrorReason = "YOUTUBE_URL_UNSUPPORTED"

	//
	// Criteria type can not be excluded by the customer,
	// like AOL account type cannot target site type criteria.
	//
	CriterionErrorReasonCANNOT_EXCLUDE_CRITERIA_TYPE CriterionErrorReason = "CANNOT_EXCLUDE_CRITERIA_TYPE"

	//
	// Criteria type can not be targeted.
	//
	CriterionErrorReasonCANNOT_ADD_CRITERIA_TYPE CriterionErrorReason = "CANNOT_ADD_CRITERIA_TYPE"

	//
	// Product filter in the product criteria has invalid characters.
	// Operand and the argument in the filter can not have "==" or "&+".
	//
	CriterionErrorReasonINVALID_PRODUCT_FILTER CriterionErrorReason = "INVALID_PRODUCT_FILTER"

	//
	// Product filter in the product criteria is translated to a string as
	// operand1==argument1&+operand2==argument2, maximum allowed length for
	// the string is 255 chars.
	//
	CriterionErrorReasonPRODUCT_FILTER_TOO_LONG CriterionErrorReason = "PRODUCT_FILTER_TOO_LONG"

	//
	// Not allowed to exclude similar user list.
	//
	CriterionErrorReasonCANNOT_EXCLUDE_SIMILAR_USER_LIST CriterionErrorReason = "CANNOT_EXCLUDE_SIMILAR_USER_LIST"

	//
	// Not allowed to target a closed user list.
	//
	CriterionErrorReasonCANNOT_ADD_CLOSED_USER_LIST CriterionErrorReason = "CANNOT_ADD_CLOSED_USER_LIST"

	//
	// Not allowed to add display only UserLists to search only campaigns.
	//
	CriterionErrorReasonCANNOT_ADD_DISPLAY_ONLY_LISTS_TO_SEARCH_ONLY_CAMPAIGNS CriterionErrorReason = "CANNOT_ADD_DISPLAY_ONLY_LISTS_TO_SEARCH_ONLY_CAMPAIGNS"

	//
	// Not allowed to add display only UserLists to search plus campaigns.
	//
	CriterionErrorReasonCANNOT_ADD_DISPLAY_ONLY_LISTS_TO_SEARCH_CAMPAIGNS CriterionErrorReason = "CANNOT_ADD_DISPLAY_ONLY_LISTS_TO_SEARCH_CAMPAIGNS"

	//
	// Not allowed to add display only UserLists to shopping campaigns.
	//
	CriterionErrorReasonCANNOT_ADD_DISPLAY_ONLY_LISTS_TO_SHOPPING_CAMPAIGNS CriterionErrorReason = "CANNOT_ADD_DISPLAY_ONLY_LISTS_TO_SHOPPING_CAMPAIGNS"

	//
	// Not allowed to add User interests to search only campaigns.
	//
	CriterionErrorReasonCANNOT_ADD_USER_INTERESTS_TO_SEARCH_CAMPAIGNS CriterionErrorReason = "CANNOT_ADD_USER_INTERESTS_TO_SEARCH_CAMPAIGNS"

	//
	// Not allowed to set bids for this criterion type in search campaigns
	//
	CriterionErrorReasonCANNOT_SET_BIDS_ON_CRITERION_TYPE_IN_SEARCH_CAMPAIGNS CriterionErrorReason = "CANNOT_SET_BIDS_ON_CRITERION_TYPE_IN_SEARCH_CAMPAIGNS"

	//
	// Final URLs, URL Templates and CustomParameters cannot be set for the criterion
	// types of Gender, AgeRange, UserList, Placement, MobileApp, and MobileAppCategory
	// in search campaigns and shopping campaigns.
	//
	CriterionErrorReasonCANNOT_ADD_URLS_TO_CRITERION_TYPE_FOR_CAMPAIGN_TYPE CriterionErrorReason = "CANNOT_ADD_URLS_TO_CRITERION_TYPE_FOR_CAMPAIGN_TYPE"

	//
	// IP address is not valid.
	//
	CriterionErrorReasonINVALID_IP_ADDRESS CriterionErrorReason = "INVALID_IP_ADDRESS"

	//
	// IP format is not valid.
	//
	CriterionErrorReasonINVALID_IP_FORMAT CriterionErrorReason = "INVALID_IP_FORMAT"

	//
	// Mobile application is not valid.
	//
	CriterionErrorReasonINVALID_MOBILE_APP CriterionErrorReason = "INVALID_MOBILE_APP"

	//
	// Mobile application category is not valid.
	//
	CriterionErrorReasonINVALID_MOBILE_APP_CATEGORY CriterionErrorReason = "INVALID_MOBILE_APP_CATEGORY"

	//
	// The CriterionId does not exist or is of the incorrect type.
	//
	CriterionErrorReasonINVALID_CRITERION_ID CriterionErrorReason = "INVALID_CRITERION_ID"

	//
	// The Criterion is not allowed to be targeted.
	//
	CriterionErrorReasonCANNOT_TARGET_CRITERION CriterionErrorReason = "CANNOT_TARGET_CRITERION"

	//
	// The criterion is not allowed to be targeted as it is deprecated.
	//
	CriterionErrorReasonCANNOT_TARGET_OBSOLETE_CRITERION CriterionErrorReason = "CANNOT_TARGET_OBSOLETE_CRITERION"

	//
	// The CriterionId is not valid for the type.
	//
	CriterionErrorReasonCRITERION_ID_AND_TYPE_MISMATCH CriterionErrorReason = "CRITERION_ID_AND_TYPE_MISMATCH"

	//
	// Distance for the radius for the proximity criterion is invalid.
	//
	CriterionErrorReasonINVALID_PROXIMITY_RADIUS CriterionErrorReason = "INVALID_PROXIMITY_RADIUS"

	//
	// Units for the distance for the radius for the proximity criterion is invalid.
	//
	CriterionErrorReasonINVALID_PROXIMITY_RADIUS_UNITS CriterionErrorReason = "INVALID_PROXIMITY_RADIUS_UNITS"

	//
	// Street address is too short.
	//
	CriterionErrorReasonINVALID_STREETADDRESS_LENGTH CriterionErrorReason = "INVALID_STREETADDRESS_LENGTH"

	//
	// City name in the address is too short.
	//
	CriterionErrorReasonINVALID_CITYNAME_LENGTH CriterionErrorReason = "INVALID_CITYNAME_LENGTH"

	//
	// Region code in the address is too short.
	//
	CriterionErrorReasonINVALID_REGIONCODE_LENGTH CriterionErrorReason = "INVALID_REGIONCODE_LENGTH"

	//
	// Region name in the address is not valid.
	//
	CriterionErrorReasonINVALID_REGIONNAME_LENGTH CriterionErrorReason = "INVALID_REGIONNAME_LENGTH"

	//
	// Postal code in the address is not valid.
	//
	CriterionErrorReasonINVALID_POSTALCODE_LENGTH CriterionErrorReason = "INVALID_POSTALCODE_LENGTH"

	//
	// Country code in the address is not valid.
	//
	CriterionErrorReasonINVALID_COUNTRY_CODE CriterionErrorReason = "INVALID_COUNTRY_CODE"

	//
	// Latitude for the GeoPoint is not valid.
	//
	CriterionErrorReasonINVALID_LATITUDE CriterionErrorReason = "INVALID_LATITUDE"

	//
	// Longitude for the GeoPoint is not valid.
	//
	CriterionErrorReasonINVALID_LONGITUDE CriterionErrorReason = "INVALID_LONGITUDE"

	//
	// The Proximity input is not valid. Both address and geoPoint cannot be null.
	//
	CriterionErrorReasonPROXIMITY_GEOPOINT_AND_ADDRESS_BOTH_CANNOT_BE_NULL CriterionErrorReason = "PROXIMITY_GEOPOINT_AND_ADDRESS_BOTH_CANNOT_BE_NULL"

	//
	// The Proximity address cannot be geocoded to a valid lat/long.
	//
	CriterionErrorReasonINVALID_PROXIMITY_ADDRESS CriterionErrorReason = "INVALID_PROXIMITY_ADDRESS"

	//
	// User domain name is not valid.
	//
	CriterionErrorReasonINVALID_USER_DOMAIN_NAME CriterionErrorReason = "INVALID_USER_DOMAIN_NAME"

	//
	// Length of serialized criterion parameter exceeded size limit.
	//
	CriterionErrorReasonCRITERION_PARAMETER_TOO_LONG CriterionErrorReason = "CRITERION_PARAMETER_TOO_LONG"

	//
	// Time interval in the AdSchedule overlaps with another AdSchedule.
	//
	CriterionErrorReasonAD_SCHEDULE_TIME_INTERVALS_OVERLAP CriterionErrorReason = "AD_SCHEDULE_TIME_INTERVALS_OVERLAP"

	//
	// AdSchedule time interval cannot span multiple days.
	//
	CriterionErrorReasonAD_SCHEDULE_INTERVAL_CANNOT_SPAN_MULTIPLE_DAYS CriterionErrorReason = "AD_SCHEDULE_INTERVAL_CANNOT_SPAN_MULTIPLE_DAYS"

	//
	// AdSchedule time interval specified is invalid,
	// endTime cannot be earlier than startTime.
	//
	CriterionErrorReasonAD_SCHEDULE_INVALID_TIME_INTERVAL CriterionErrorReason = "AD_SCHEDULE_INVALID_TIME_INTERVAL"

	//
	// The number of AdSchedule entries in a day exceeds the limit.
	//
	CriterionErrorReasonAD_SCHEDULE_EXCEEDED_INTERVALS_PER_DAY_LIMIT CriterionErrorReason = "AD_SCHEDULE_EXCEEDED_INTERVALS_PER_DAY_LIMIT"

	//
	// CriteriaId does not match the interval of the AdSchedule specified.
	//
	CriterionErrorReasonAD_SCHEDULE_CRITERION_ID_MISMATCHING_FIELDS CriterionErrorReason = "AD_SCHEDULE_CRITERION_ID_MISMATCHING_FIELDS"

	//
	// Cannot set bid modifier for this criterion type.
	//
	CriterionErrorReasonCANNOT_BID_MODIFY_CRITERION_TYPE CriterionErrorReason = "CANNOT_BID_MODIFY_CRITERION_TYPE"

	//
	// Cannot bid modify criterion, since it is opted out of the campaign.
	//
	CriterionErrorReasonCANNOT_BID_MODIFY_CRITERION_CAMPAIGN_OPTED_OUT CriterionErrorReason = "CANNOT_BID_MODIFY_CRITERION_CAMPAIGN_OPTED_OUT"

	//
	// Cannot set bid modifier for a negative criterion.
	//
	CriterionErrorReasonCANNOT_BID_MODIFY_NEGATIVE_CRITERION CriterionErrorReason = "CANNOT_BID_MODIFY_NEGATIVE_CRITERION"

	//
	// Bid Modifier already exists. Use SET operation to update.
	//
	CriterionErrorReasonBID_MODIFIER_ALREADY_EXISTS CriterionErrorReason = "BID_MODIFIER_ALREADY_EXISTS"

	//
	// Feed Id is not allowed in these Location Groups.
	//
	CriterionErrorReasonFEED_ID_NOT_ALLOWED CriterionErrorReason = "FEED_ID_NOT_ALLOWED"

	//
	// The account may not use the requested criteria type. For example, some
	// accounts are restricted to keywords only.
	//
	CriterionErrorReasonACCOUNT_INELIGIBLE_FOR_CRITERIA_TYPE CriterionErrorReason = "ACCOUNT_INELIGIBLE_FOR_CRITERIA_TYPE"

	//
	// The requested criteria type cannot be used with campaign or ad group bidding strategy.
	//
	CriterionErrorReasonCRITERIA_TYPE_INVALID_FOR_BIDDING_STRATEGY CriterionErrorReason = "CRITERIA_TYPE_INVALID_FOR_BIDDING_STRATEGY"

	//
	// The Criterion is not allowed to be excluded.
	//
	CriterionErrorReasonCANNOT_EXCLUDE_CRITERION CriterionErrorReason = "CANNOT_EXCLUDE_CRITERION"

	//
	// The criterion is not allowed to be removed. For example, we cannot remove any
	// of the platform criterion.
	//
	CriterionErrorReasonCANNOT_REMOVE_CRITERION CriterionErrorReason = "CANNOT_REMOVE_CRITERION"

	//
	// The combined length of product dimension values of the product scope criterion is too long.
	//
	CriterionErrorReasonPRODUCT_SCOPE_TOO_LONG CriterionErrorReason = "PRODUCT_SCOPE_TOO_LONG"

	//
	// Product scope contains too many dimensions.
	//
	CriterionErrorReasonPRODUCT_SCOPE_TOO_MANY_DIMENSIONS CriterionErrorReason = "PRODUCT_SCOPE_TOO_MANY_DIMENSIONS"

	//
	// The combined length of product dimension values of the product partition criterion is too
	// long.
	//
	CriterionErrorReasonPRODUCT_PARTITION_TOO_LONG CriterionErrorReason = "PRODUCT_PARTITION_TOO_LONG"

	//
	// Product partition contains too many dimensions.
	//
	CriterionErrorReasonPRODUCT_PARTITION_TOO_MANY_DIMENSIONS CriterionErrorReason = "PRODUCT_PARTITION_TOO_MANY_DIMENSIONS"

	//
	// The product dimension is invalid (e.g. dimension contains illegal value, dimension type is
	// represented with wrong class, etc). Product dimension value can not contain "==" or "&+".
	//
	CriterionErrorReasonINVALID_PRODUCT_DIMENSION CriterionErrorReason = "INVALID_PRODUCT_DIMENSION"

	//
	// Product dimension type is either invalid for campaigns of this type or cannot be used in the
	// current context. BIDDING_CATEGORY_Lx and PRODUCT_TYPE_Lx product dimensions must be used in
	// ascending order of their levels: L1, L2, L3, L4, L5... The levels must be specified
	// sequentially and start from L1. Furthermore, an "others" product partition cannot be
	// subdivided with a dimension of the same type but of a higher level ("others"
	// BIDDING_CATEGORY_L3 can be subdivided with BRAND but not with BIDDING_CATEGORY_L4).
	//
	CriterionErrorReasonINVALID_PRODUCT_DIMENSION_TYPE CriterionErrorReason = "INVALID_PRODUCT_DIMENSION_TYPE"

	//
	// Bidding categories do not form a valid path in the Shopping bidding category taxonomy.
	//
	CriterionErrorReasonINVALID_PRODUCT_BIDDING_CATEGORY CriterionErrorReason = "INVALID_PRODUCT_BIDDING_CATEGORY"

	//
	// ShoppingSetting must be added to the campaign before ProductScope criteria can be added.
	//
	CriterionErrorReasonMISSING_SHOPPING_SETTING CriterionErrorReason = "MISSING_SHOPPING_SETTING"

	//
	// Matching function is invalid.
	//
	CriterionErrorReasonINVALID_MATCHING_FUNCTION CriterionErrorReason = "INVALID_MATCHING_FUNCTION"

	//
	// Filter parameters not allowed for location groups targeting.
	//
	CriterionErrorReasonLOCATION_FILTER_NOT_ALLOWED CriterionErrorReason = "LOCATION_FILTER_NOT_ALLOWED"

	//
	// Given location filter parameter is invalid for location groups targeting.
	//
	CriterionErrorReasonLOCATION_FILTER_INVALID CriterionErrorReason = "LOCATION_FILTER_INVALID"

	//
	// Criteria type cannot be associated with a campaign and its ad group(s) simultaneously.
	//
	CriterionErrorReasonCANNOT_ATTACH_CRITERIA_AT_CAMPAIGN_AND_ADGROUP CriterionErrorReason = "CANNOT_ATTACH_CRITERIA_AT_CAMPAIGN_AND_ADGROUP"

	CriterionErrorReasonUNKNOWN CriterionErrorReason = "UNKNOWN"
)

//
// The reasons for the database error.
//
type DatabaseErrorReason string

const (

	//
	// A concurrency problem occurred as two threads were attempting to modify same object.
	//
	DatabaseErrorReasonCONCURRENT_MODIFICATION DatabaseErrorReason = "CONCURRENT_MODIFICATION"

	//
	// The permission was denied to access an object.
	//
	DatabaseErrorReasonPERMISSION_DENIED DatabaseErrorReason = "PERMISSION_DENIED"

	//
	// The user's access to an object has been prohibited.
	//
	DatabaseErrorReasonACCESS_PROHIBITED DatabaseErrorReason = "ACCESS_PROHIBITED"

	//
	// Requested campaign belongs to a product that is not supported by the api.
	//
	DatabaseErrorReasonCAMPAIGN_PRODUCT_NOT_SUPPORTED DatabaseErrorReason = "CAMPAIGN_PRODUCT_NOT_SUPPORTED"

	//
	// a duplicate key was detected upon insertion
	//
	DatabaseErrorReasonDUPLICATE_KEY DatabaseErrorReason = "DUPLICATE_KEY"

	//
	// a database error has occurred
	//
	DatabaseErrorReasonDATABASE_ERROR DatabaseErrorReason = "DATABASE_ERROR"

	//
	// an unknown error has occurred
	//
	DatabaseErrorReasonUNKNOWN DatabaseErrorReason = "UNKNOWN"
)

//
// The reasons for the target error.
//
type DateErrorReason string

const (

	//
	// Given field values do not correspond to a valid date.
	//
	DateErrorReasonINVALID_FIELD_VALUES_IN_DATE DateErrorReason = "INVALID_FIELD_VALUES_IN_DATE"

	//
	// Given field values do not correspond to a valid date time.
	//
	DateErrorReasonINVALID_FIELD_VALUES_IN_DATE_TIME DateErrorReason = "INVALID_FIELD_VALUES_IN_DATE_TIME"

	//
	// The string date's format should be yyyymmdd.
	//
	DateErrorReasonINVALID_STRING_DATE DateErrorReason = "INVALID_STRING_DATE"

	//
	// The string date range's format should be yyyymmdd yyyymmdd.
	//
	DateErrorReasonINVALID_STRING_DATE_RANGE DateErrorReason = "INVALID_STRING_DATE_RANGE"

	//
	// The string date time's format should be yyyymmdd hhmmss [tz].
	//
	DateErrorReasonINVALID_STRING_DATE_TIME DateErrorReason = "INVALID_STRING_DATE_TIME"

	//
	// Date is before allowed minimum.
	//
	DateErrorReasonEARLIER_THAN_MINIMUM_DATE DateErrorReason = "EARLIER_THAN_MINIMUM_DATE"

	//
	// Date is after allowed maximum.
	//
	DateErrorReasonLATER_THAN_MAXIMUM_DATE DateErrorReason = "LATER_THAN_MAXIMUM_DATE"

	//
	// Date range bounds are not in order.
	//
	DateErrorReasonDATE_RANGE_MINIMUM_DATE_LATER_THAN_MAXIMUM_DATE DateErrorReason = "DATE_RANGE_MINIMUM_DATE_LATER_THAN_MAXIMUM_DATE"

	//
	// Both dates in range are null.
	//
	DateErrorReasonDATE_RANGE_MINIMUM_AND_MAXIMUM_DATES_BOTH_NULL DateErrorReason = "DATE_RANGE_MINIMUM_AND_MAXIMUM_DATES_BOTH_NULL"
)

//
// Days of the week.
//
type DayOfWeek string

const (

	//
	// The day of week named Monday.
	//
	DayOfWeekMONDAY DayOfWeek = "MONDAY"

	//
	// The day of week named Tuesday.
	//
	DayOfWeekTUESDAY DayOfWeek = "TUESDAY"

	//
	// The day of week named Wednesday.
	//
	DayOfWeekWEDNESDAY DayOfWeek = "WEDNESDAY"

	//
	// The day of week named Thursday.
	//
	DayOfWeekTHURSDAY DayOfWeek = "THURSDAY"

	//
	// The day of week named Friday.
	//
	DayOfWeekFRIDAY DayOfWeek = "FRIDAY"

	//
	// The day of week named Saturday.
	//
	DayOfWeekSATURDAY DayOfWeek = "SATURDAY"

	//
	// The day of week named Sunday.
	//
	DayOfWeekSUNDAY DayOfWeek = "SUNDAY"
)

//
// The reasons for the validation error.
//
type DistinctErrorReason string

const (
	DistinctErrorReasonDUPLICATE_ELEMENT DistinctErrorReason = "DUPLICATE_ELEMENT"

	DistinctErrorReasonDUPLICATE_TYPE DistinctErrorReason = "DUPLICATE_TYPE"
)

type EntityAccessDeniedReason string

const (

	//
	// User did not have read access.
	//
	EntityAccessDeniedReasonREAD_ACCESS_DENIED EntityAccessDeniedReason = "READ_ACCESS_DENIED"

	//
	// User did not have write access.
	//
	EntityAccessDeniedReasonWRITE_ACCESS_DENIED EntityAccessDeniedReason = "WRITE_ACCESS_DENIED"
)

//
// Limits at various levels of the account.
//
type EntityCountLimitExceededReason string

const (

	//
	// Indicates that this request would exceed the number of allowed entities for the AdWords
	// account. The exact entity type and limit being checked can be inferred from
	// {@link #accountLimitType}.
	//
	EntityCountLimitExceededReasonACCOUNT_LIMIT EntityCountLimitExceededReason = "ACCOUNT_LIMIT"

	//
	// Indicates that this request would exceed the number of allowed entities in a Campaign.
	// The exact entity type and limit being checked can be inferred from
	// {@link #accountLimitType}, and the numeric id of the Campaign involved is given by
	// {@link #enclosingId}.
	//
	EntityCountLimitExceededReasonCAMPAIGN_LIMIT EntityCountLimitExceededReason = "CAMPAIGN_LIMIT"

	//
	// Indicates that this request would exceed the number of allowed entities in
	// an ad group.  The exact entity type and limit being checked can be
	// inferred from {@link #accountLimitType}, and the numeric id of the ad group
	// involved is given by {@link #enclosingId}.
	//
	EntityCountLimitExceededReasonADGROUP_LIMIT EntityCountLimitExceededReason = "ADGROUP_LIMIT"

	//
	// Indicates that this request would exceed the number of allowed entities in an ad group ad.
	// The exact entity type and limit being checked can be inferred from {@link #accountLimitType},
	// and the {@link #enclosingId} contains the ad group id followed by the ad id, separated by a
	// single comma (,).
	//
	EntityCountLimitExceededReasonAD_GROUP_AD_LIMIT EntityCountLimitExceededReason = "AD_GROUP_AD_LIMIT"

	//
	// Indicates that this request would exceed the number of allowed entities in an ad group
	// criterion.  The exact entity type and limit being checked can be inferred from
	// {@link #accountLimitType}, and the {@link #enclosingId} contains the ad group id followed by
	// the criterion id, separated by a single comma (,).
	//
	EntityCountLimitExceededReasonAD_GROUP_CRITERION_LIMIT EntityCountLimitExceededReason = "AD_GROUP_CRITERION_LIMIT"

	//
	// Indicates that this request would exceed the number of allowed entities in
	// this shared set.  The exact entity type and limit being checked can be
	// inferred from {@link #accountLimitType}, and the numeric id of the shared
	// set involved is given by {@link #enclosingId}.
	//
	EntityCountLimitExceededReasonSHARED_SET_LIMIT EntityCountLimitExceededReason = "SHARED_SET_LIMIT"

	//
	// Exceeds a limit related to a matching function.
	//
	EntityCountLimitExceededReasonMATCHING_FUNCTION_LIMIT EntityCountLimitExceededReason = "MATCHING_FUNCTION_LIMIT"

	//
	// Specific limit that has been exceeded is unknown (the client may be of an
	// older version than the server).
	//
	EntityCountLimitExceededReasonUNKNOWN EntityCountLimitExceededReason = "UNKNOWN"
)

type EntityNotFoundReason string

const (

	//
	// The specified id refered to an entity which either doesn't exist or is not accessible to the
	// customer. e.g. campaign belongs to another customer.
	//
	EntityNotFoundReasonINVALID_ID EntityNotFoundReason = "INVALID_ID"
)

//
// Different levels of platform restrictions.
//
type ExtensionSettingPlatform string

const (

	//
	// Desktop and tablet devices only.
	//
	ExtensionSettingPlatformDESKTOP ExtensionSettingPlatform = "DESKTOP"

	//
	// Mobile only.
	//
	ExtensionSettingPlatformMOBILE ExtensionSettingPlatform = "MOBILE"

	//
	// No restriction.
	//
	ExtensionSettingPlatformNONE ExtensionSettingPlatform = "NONE"
)

//
// Error reasons.
//
type ExtensionSettingErrorReason string

const (

	//
	// A platform restriction was provided without input extensions or existing extensions.
	//
	ExtensionSettingErrorReasonEXTENSIONS_REQUIRED ExtensionSettingErrorReason = "EXTENSIONS_REQUIRED"

	//
	// The provided feed type does not correspond to the provided extensions.
	//
	ExtensionSettingErrorReasonFEED_TYPE_EXTENSION_TYPE_MISMATCH ExtensionSettingErrorReason = "FEED_TYPE_EXTENSION_TYPE_MISMATCH"

	//
	// The provided feed type cannot be used.
	//
	ExtensionSettingErrorReasonINVALID_FEED_TYPE ExtensionSettingErrorReason = "INVALID_FEED_TYPE"

	//
	// The provided feed type cannot be used at the customer level.
	//
	ExtensionSettingErrorReasonINVALID_FEED_TYPE_FOR_CUSTOMER_EXTENSION_SETTING ExtensionSettingErrorReason = "INVALID_FEED_TYPE_FOR_CUSTOMER_EXTENSION_SETTING"

	//
	// Can not change a feed item field on an ADD operation.
	//
	ExtensionSettingErrorReasonCANNOT_CHANGE_FEED_ITEM_ON_ADD ExtensionSettingErrorReason = "CANNOT_CHANGE_FEED_ITEM_ON_ADD"

	//
	// Can not update an extension that is not already in this setting.
	//
	ExtensionSettingErrorReasonCANNOT_UPDATE_NEWLY_ADDED_EXTENSION ExtensionSettingErrorReason = "CANNOT_UPDATE_NEWLY_ADDED_EXTENSION"

	//
	// There is no existing AdGroupExtensionSetting for this type.
	//
	ExtensionSettingErrorReasonNO_EXISTING_AD_GROUP_EXTENSION_SETTING_FOR_TYPE ExtensionSettingErrorReason = "NO_EXISTING_AD_GROUP_EXTENSION_SETTING_FOR_TYPE"

	//
	// There is no existing CampaignExtensionSetting for this type.
	//
	ExtensionSettingErrorReasonNO_EXISTING_CAMPAIGN_EXTENSION_SETTING_FOR_TYPE ExtensionSettingErrorReason = "NO_EXISTING_CAMPAIGN_EXTENSION_SETTING_FOR_TYPE"

	//
	// There is no existing CustomerExtensionSetting for this type.
	//
	ExtensionSettingErrorReasonNO_EXISTING_CUSTOMER_EXTENSION_SETTING_FOR_TYPE ExtensionSettingErrorReason = "NO_EXISTING_CUSTOMER_EXTENSION_SETTING_FOR_TYPE"

	//
	// The AdGroupExtensionSetting already exists. SET should be used to modify the existing
	// AdGroupExtensionSetting.
	//
	ExtensionSettingErrorReasonAD_GROUP_EXTENSION_SETTING_ALREADY_EXISTS ExtensionSettingErrorReason = "AD_GROUP_EXTENSION_SETTING_ALREADY_EXISTS"

	//
	// The CampaignExtensionSetting already exists. SET should be used to modify the existing
	// CampaignExtensionSetting.
	//
	ExtensionSettingErrorReasonCAMPAIGN_EXTENSION_SETTING_ALREADY_EXISTS ExtensionSettingErrorReason = "CAMPAIGN_EXTENSION_SETTING_ALREADY_EXISTS"

	//
	// The CustomerExtensionSetting already exists. SET should be used to modify the existing
	// CustomerExtensionSetting.
	//
	ExtensionSettingErrorReasonCUSTOMER_EXTENSION_SETTING_ALREADY_EXISTS ExtensionSettingErrorReason = "CUSTOMER_EXTENSION_SETTING_ALREADY_EXISTS"

	//
	// An active ad group feed already exists for this place holder type.
	//
	ExtensionSettingErrorReasonAD_GROUP_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE ExtensionSettingErrorReason = "AD_GROUP_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE"

	//
	// An active campaign feed already exists for this place holder type.
	//
	ExtensionSettingErrorReasonCAMPAIGN_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE ExtensionSettingErrorReason = "CAMPAIGN_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE"

	//
	// An active customer feed already exists for this place holder type.
	//
	ExtensionSettingErrorReasonCUSTOMER_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE ExtensionSettingErrorReason = "CUSTOMER_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE"

	//
	// Value is not within the accepted range.
	//
	ExtensionSettingErrorReasonVALUE_OUT_OF_RANGE ExtensionSettingErrorReason = "VALUE_OUT_OF_RANGE"

	//
	// Cannot simultaneously set sitelink field with final urls.
	//
	ExtensionSettingErrorReasonCANNOT_SET_WITH_FINAL_URLS ExtensionSettingErrorReason = "CANNOT_SET_WITH_FINAL_URLS"

	//
	// Must set field with final urls.
	//
	ExtensionSettingErrorReasonCANNOT_SET_WITHOUT_FINAL_URLS ExtensionSettingErrorReason = "CANNOT_SET_WITHOUT_FINAL_URLS"

	//
	// Cannot simultaneously set sitelink url field with tracking url template.
	//
	ExtensionSettingErrorReasonCANNOT_SET_BOTH_DESTINATION_URL_AND_TRACKING_URL_TEMPLATE ExtensionSettingErrorReason = "CANNOT_SET_BOTH_DESTINATION_URL_AND_TRACKING_URL_TEMPLATE"

	//
	// Phone number for a call extension is invalid.
	//
	ExtensionSettingErrorReasonINVALID_PHONE_NUMBER ExtensionSettingErrorReason = "INVALID_PHONE_NUMBER"

	//
	// Phone number for a call extension is not supported for the given country code.
	//
	ExtensionSettingErrorReasonPHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY ExtensionSettingErrorReason = "PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY"

	//
	// A carrier specific number in short format is not allowed for call extensions.
	//
	ExtensionSettingErrorReasonCARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED ExtensionSettingErrorReason = "CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED"

	//
	// Premium rate numbers are not allowed for call extensions.
	//
	ExtensionSettingErrorReasonPREMIUM_RATE_NUMBER_NOT_ALLOWED ExtensionSettingErrorReason = "PREMIUM_RATE_NUMBER_NOT_ALLOWED"

	//
	// Phone number type for a call extension is not allowed.
	//
	ExtensionSettingErrorReasonDISALLOWED_NUMBER_TYPE ExtensionSettingErrorReason = "DISALLOWED_NUMBER_TYPE"

	//
	// Phone number for a call extension does not meet domestic format requirements.
	//
	ExtensionSettingErrorReasonINVALID_DOMESTIC_PHONE_NUMBER_FORMAT ExtensionSettingErrorReason = "INVALID_DOMESTIC_PHONE_NUMBER_FORMAT"

	//
	// Vanity phone numbers (i.e. those including letters) are not allowed for call extensions.
	//
	ExtensionSettingErrorReasonVANITY_PHONE_NUMBER_NOT_ALLOWED ExtensionSettingErrorReason = "VANITY_PHONE_NUMBER_NOT_ALLOWED"

	//
	// Country code provided for a call extension is invalid.
	//
	ExtensionSettingErrorReasonINVALID_COUNTRY_CODE ExtensionSettingErrorReason = "INVALID_COUNTRY_CODE"

	//
	// Call conversion type id provided for a call extension is invalid.
	//
	ExtensionSettingErrorReasonINVALID_CALL_CONVERSION_TYPE_ID ExtensionSettingErrorReason = "INVALID_CALL_CONVERSION_TYPE_ID"

	//
	// For a call extension, the customer is not whitelisted for call tracking.
	//
	ExtensionSettingErrorReasonCUSTOMER_NOT_WHITELISTED_FOR_CALLTRACKING ExtensionSettingErrorReason = "CUSTOMER_NOT_WHITELISTED_FOR_CALLTRACKING"

	//
	// Call tracking is not supported for the given country for a call extension.
	//
	ExtensionSettingErrorReasonCALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY ExtensionSettingErrorReason = "CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY"

	//
	// App id provided for an app extension is invalid.
	//
	ExtensionSettingErrorReasonINVALID_APP_ID ExtensionSettingErrorReason = "INVALID_APP_ID"

	//
	// Quotation marks present in the review text for a review extension.
	//
	ExtensionSettingErrorReasonQUOTES_IN_REVIEW_EXTENSION_SNIPPET ExtensionSettingErrorReason = "QUOTES_IN_REVIEW_EXTENSION_SNIPPET"

	//
	// Hyphen character present in the review text for a review extension.
	//
	ExtensionSettingErrorReasonHYPHENS_IN_REVIEW_EXTENSION_SNIPPET ExtensionSettingErrorReason = "HYPHENS_IN_REVIEW_EXTENSION_SNIPPET"

	//
	// A blacklisted review source name or url was provided for a review extension.
	//
	ExtensionSettingErrorReasonREVIEW_EXTENSION_SOURCE_INELIGIBLE ExtensionSettingErrorReason = "REVIEW_EXTENSION_SOURCE_INELIGIBLE"

	//
	// Review source name should not be found in the review text.
	//
	ExtensionSettingErrorReasonSOURCE_NAME_IN_REVIEW_EXTENSION_TEXT ExtensionSettingErrorReason = "SOURCE_NAME_IN_REVIEW_EXTENSION_TEXT"

	//
	// Field must be set.
	//
	ExtensionSettingErrorReasonMISSING_FIELD ExtensionSettingErrorReason = "MISSING_FIELD"

	//
	// Inconsistent currency codes.
	//
	ExtensionSettingErrorReasonINCONSISTENT_CURRENCY_CODES ExtensionSettingErrorReason = "INCONSISTENT_CURRENCY_CODES"

	//
	// Price extension cannot have duplicated headers.
	//
	ExtensionSettingErrorReasonPRICE_EXTENSION_HAS_DUPLICATED_HEADERS ExtensionSettingErrorReason = "PRICE_EXTENSION_HAS_DUPLICATED_HEADERS"

	//
	// Price item cannot have duplicated header and description.
	//
	ExtensionSettingErrorReasonPRICE_ITEM_HAS_DUPLICATED_HEADER_AND_DESCRIPTION ExtensionSettingErrorReason = "PRICE_ITEM_HAS_DUPLICATED_HEADER_AND_DESCRIPTION"

	//
	// Price extension has too few items
	//
	ExtensionSettingErrorReasonPRICE_EXTENSION_HAS_TOO_FEW_ITEMS ExtensionSettingErrorReason = "PRICE_EXTENSION_HAS_TOO_FEW_ITEMS"

	//
	// Price extension has too many items
	//
	ExtensionSettingErrorReasonPRICE_EXTENSION_HAS_TOO_MANY_ITEMS ExtensionSettingErrorReason = "PRICE_EXTENSION_HAS_TOO_MANY_ITEMS"

	//
	// The input value is not currently supported.
	//
	ExtensionSettingErrorReasonUNSUPPORTED_VALUE ExtensionSettingErrorReason = "UNSUPPORTED_VALUE"

	//
	// The input value is not currently supported in the selected language of an extension.
	//
	ExtensionSettingErrorReasonUNSUPPORTED_VALUE_IN_SELECTED_LANGUAGE ExtensionSettingErrorReason = "UNSUPPORTED_VALUE_IN_SELECTED_LANGUAGE"

	//
	// Unknown or unsupported device preference.
	//
	ExtensionSettingErrorReasonINVALID_DEVICE_PREFERENCE ExtensionSettingErrorReason = "INVALID_DEVICE_PREFERENCE"

	//
	// Invalid feed item schedule end time (i.e., endHour = 24 and endMinute != 0).
	//
	ExtensionSettingErrorReasonINVALID_SCHEDULE_END ExtensionSettingErrorReason = "INVALID_SCHEDULE_END"

	//
	// Date time zone does not match the account's time zone.
	//
	ExtensionSettingErrorReasonDATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE ExtensionSettingErrorReason = "DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE"

	//
	// Overlapping feed item schedule times (e.g., 7-10AM and 8-11AM) are not allowed.
	//
	ExtensionSettingErrorReasonOVERLAPPING_SCHEDULES ExtensionSettingErrorReason = "OVERLAPPING_SCHEDULES"

	//
	// Feed item schedule end time must be after start time.
	//
	ExtensionSettingErrorReasonSCHEDULE_END_NOT_AFTER_START ExtensionSettingErrorReason = "SCHEDULE_END_NOT_AFTER_START"

	//
	// There are too many feed item schedules per day.
	//
	ExtensionSettingErrorReasonTOO_MANY_SCHEDULES_PER_DAY ExtensionSettingErrorReason = "TOO_MANY_SCHEDULES_PER_DAY"

	//
	// Cannot edit the same extension feed item id twice.
	//
	ExtensionSettingErrorReasonDUPLICATE_EXTENSION_FEED_ITEM_EDIT ExtensionSettingErrorReason = "DUPLICATE_EXTENSION_FEED_ITEM_EDIT"

	//
	// Invalid structured snippet header.
	//
	ExtensionSettingErrorReasonINVALID_SNIPPETS_HEADER ExtensionSettingErrorReason = "INVALID_SNIPPETS_HEADER"

	//
	// Phone number not supported with call tracking enabled for country.
	//
	ExtensionSettingErrorReasonPHONE_NUMBER_NOT_SUPPORTED_WITH_CALLTRACKING_FOR_COUNTRY ExtensionSettingErrorReason = "PHONE_NUMBER_NOT_SUPPORTED_WITH_CALLTRACKING_FOR_COUNTRY"

	//
	// Targeted adgroup's campaign does not match the targeted campaign.
	//
	ExtensionSettingErrorReasonCAMPAIGN_TARGETING_MISMATCH ExtensionSettingErrorReason = "CAMPAIGN_TARGETING_MISMATCH"

	//
	// The feed used by the ExtensionSetting is deleted and cannot be operated on. Remove the
	// ExtensionSetting to allow a new one to be created using an active feed.
	//
	ExtensionSettingErrorReasonCANNOT_OPERATE_ON_DELETED_FEED ExtensionSettingErrorReason = "CANNOT_OPERATE_ON_DELETED_FEED"

	//
	// Concrete sub type of ExtensionFeedItem is required for this operation.
	//
	ExtensionSettingErrorReasonCONCRETE_EXTENSION_TYPE_REQUIRED ExtensionSettingErrorReason = "CONCRETE_EXTENSION_TYPE_REQUIRED"

	//
	// The matching function that links the extension feed to the customer, campaign, or ad group
	// is not compatible with the ExtensionSetting services.
	//
	ExtensionSettingErrorReasonINCOMPATIBLE_UNDERLYING_MATCHING_FUNCTION ExtensionSettingErrorReason = "INCOMPATIBLE_UNDERLYING_MATCHING_FUNCTION"

	ExtensionSettingErrorReasonSTART_DATE_AFTER_END_DATE ExtensionSettingErrorReason = "START_DATE_AFTER_END_DATE"

	ExtensionSettingErrorReasonINVALID_PRICE_FORMAT ExtensionSettingErrorReason = "INVALID_PRICE_FORMAT"

	ExtensionSettingErrorReasonPROMOTION_INVALID_TIME ExtensionSettingErrorReason = "PROMOTION_INVALID_TIME"

	ExtensionSettingErrorReasonPROMOTION_CANNOT_SET_PERCENT_OFF_AND_MONEY_AMOUNT_OFF ExtensionSettingErrorReason = "PROMOTION_CANNOT_SET_PERCENT_OFF_AND_MONEY_AMOUNT_OFF"

	ExtensionSettingErrorReasonPROMOTION_CANNOT_SET_PROMOTION_CODE_AND_ORDERS_OVER_AMOUNT ExtensionSettingErrorReason = "PROMOTION_CANNOT_SET_PROMOTION_CODE_AND_ORDERS_OVER_AMOUNT"

	ExtensionSettingErrorReasonTOO_MANY_DECIMAL_PLACES_SPECIFIED ExtensionSettingErrorReason = "TOO_MANY_DECIMAL_PLACES_SPECIFIED"

	//
	// The language code is not valid.
	//
	ExtensionSettingErrorReasonINVALID_LANGUAGE_CODE ExtensionSettingErrorReason = "INVALID_LANGUAGE_CODE"

	//
	// The language is not supported.
	//
	ExtensionSettingErrorReasonUNSUPPORTED_LANGUAGE ExtensionSettingErrorReason = "UNSUPPORTED_LANGUAGE"

	ExtensionSettingErrorReasonUNKNOWN ExtensionSettingErrorReason = "UNKNOWN"
)

type FeedItemStatus string

const (

	//
	// Feed item is active
	//
	FeedItemStatusENABLED FeedItemStatus = "ENABLED"

	//
	// Feed item is removed
	//
	FeedItemStatusREMOVED FeedItemStatus = "REMOVED"

	//
	// Unknown status
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	FeedItemStatusUNKNOWN FeedItemStatus = "UNKNOWN"
)

//
// Feed item approval status.
//
type FeedItemApprovalStatus string

const (

	//
	// Pending review
	//
	FeedItemApprovalStatusUNCHECKED FeedItemApprovalStatus = "UNCHECKED"

	//
	// Approved
	//
	FeedItemApprovalStatusAPPROVED FeedItemApprovalStatus = "APPROVED"

	//
	// Disapproved
	//
	FeedItemApprovalStatusDISAPPROVED FeedItemApprovalStatus = "DISAPPROVED"
)

//
// Feed item quality evaluation approval status.
//
type FeedItemQualityApprovalStatus string

const (

	//
	// Quality evaluation pending
	//
	FeedItemQualityApprovalStatusUNKNOWN FeedItemQualityApprovalStatus = "UNKNOWN"

	//
	// Approved for quality
	//
	FeedItemQualityApprovalStatusAPPROVED FeedItemQualityApprovalStatus = "APPROVED"

	//
	// Disapproved for quality
	//
	FeedItemQualityApprovalStatusDISAPPROVED FeedItemQualityApprovalStatus = "DISAPPROVED"
)

//
// Feed item quality evaluation disapproval reasons.
//
type FeedItemQualityDisapprovalReasons string

const (
	FeedItemQualityDisapprovalReasonsUNKNOWN FeedItemQualityDisapprovalReasons = "UNKNOWN"

	//
	// Price contains repetitive headers
	//
	FeedItemQualityDisapprovalReasonsTABLE_REPETITIVE_HEADERS FeedItemQualityDisapprovalReasons = "TABLE_REPETITIVE_HEADERS"

	//
	// Price contains repetitive description
	//
	FeedItemQualityDisapprovalReasonsTABLE_REPETITIVE_DESCRIPTION FeedItemQualityDisapprovalReasons = "TABLE_REPETITIVE_DESCRIPTION"

	//
	// Price contains inconsistent items
	//
	FeedItemQualityDisapprovalReasonsTABLE_INCONSISTENT_ROWS FeedItemQualityDisapprovalReasons = "TABLE_INCONSISTENT_ROWS"

	//
	// Price qualifiers in description
	//
	FeedItemQualityDisapprovalReasonsDESCRIPTION_HAS_PRICE_QUALIFIERS FeedItemQualityDisapprovalReasons = "DESCRIPTION_HAS_PRICE_QUALIFIERS"

	//
	// Unsupported language
	//
	FeedItemQualityDisapprovalReasonsUNSUPPORTED_LANGUAGE FeedItemQualityDisapprovalReasons = "UNSUPPORTED_LANGUAGE"

	//
	// Price item header is not relevant to the price type
	//
	FeedItemQualityDisapprovalReasonsTABLE_ROW_HEADER_TABLE_TYPE_MISMATCH FeedItemQualityDisapprovalReasons = "TABLE_ROW_HEADER_TABLE_TYPE_MISMATCH"

	//
	// Price item header has promotional text
	//
	FeedItemQualityDisapprovalReasonsTABLE_ROW_HEADER_HAS_PROMOTIONAL_TEXT FeedItemQualityDisapprovalReasons = "TABLE_ROW_HEADER_HAS_PROMOTIONAL_TEXT"

	//
	// Price item description is not relevant to the item header
	//
	FeedItemQualityDisapprovalReasonsTABLE_ROW_DESCRIPTION_NOT_RELEVANT FeedItemQualityDisapprovalReasons = "TABLE_ROW_DESCRIPTION_NOT_RELEVANT"

	//
	// Price item description contains promotional text
	//
	FeedItemQualityDisapprovalReasonsTABLE_ROW_DESCRIPTION_HAS_PROMOTIONAL_TEXT FeedItemQualityDisapprovalReasons = "TABLE_ROW_DESCRIPTION_HAS_PROMOTIONAL_TEXT"

	//
	// Price item header and description are repetitive
	//
	FeedItemQualityDisapprovalReasonsTABLE_ROW_HEADER_DESCRIPTION_REPETITIVE FeedItemQualityDisapprovalReasons = "TABLE_ROW_HEADER_DESCRIPTION_REPETITIVE"

	//
	// Price item is in a foreign language, nonsense, or can't be rated
	//
	FeedItemQualityDisapprovalReasonsTABLE_ROW_UNRATEABLE FeedItemQualityDisapprovalReasons = "TABLE_ROW_UNRATEABLE"

	//
	// Price item price is invalid or inaccurate
	//
	FeedItemQualityDisapprovalReasonsTABLE_ROW_PRICE_INVALID FeedItemQualityDisapprovalReasons = "TABLE_ROW_PRICE_INVALID"

	//
	// Price item url is invalid or irrelevant
	//
	FeedItemQualityDisapprovalReasonsTABLE_ROW_URL_INVALID FeedItemQualityDisapprovalReasons = "TABLE_ROW_URL_INVALID"

	//
	// Header or description has price
	//
	FeedItemQualityDisapprovalReasonsHEADER_OR_DESCRIPTION_HAS_PRICE FeedItemQualityDisapprovalReasons = "HEADER_OR_DESCRIPTION_HAS_PRICE"

	//
	// Snippet values do not match the header
	//
	FeedItemQualityDisapprovalReasonsSTRUCTURED_SNIPPETS_HEADER_POLICY_VIOLATED FeedItemQualityDisapprovalReasons = "STRUCTURED_SNIPPETS_HEADER_POLICY_VIOLATED"

	//
	// Snippet values are repeated
	//
	FeedItemQualityDisapprovalReasonsSTRUCTURED_SNIPPETS_REPEATED_VALUES FeedItemQualityDisapprovalReasons = "STRUCTURED_SNIPPETS_REPEATED_VALUES"

	//
	// Snippet values violate editorial guidelines like punctuation
	//
	FeedItemQualityDisapprovalReasonsSTRUCTURED_SNIPPETS_EDITORIAL_GUIDELINES FeedItemQualityDisapprovalReasons = "STRUCTURED_SNIPPETS_EDITORIAL_GUIDELINES"

	//
	// Snippets contain promotional text
	//
	FeedItemQualityDisapprovalReasonsSTRUCTURED_SNIPPETS_HAS_PROMOTIONAL_TEXT FeedItemQualityDisapprovalReasons = "STRUCTURED_SNIPPETS_HAS_PROMOTIONAL_TEXT"
)

//
// Validation status of a FeedItem.
//
type FeedItemValidationStatus string

const (

	//
	// Validation pending.
	//
	FeedItemValidationStatusUNCHECKED FeedItemValidationStatus = "UNCHECKED"

	//
	// An error was found.
	//
	FeedItemValidationStatusERROR FeedItemValidationStatus = "ERROR"

	//
	// FeedItem is semantically well-formed.
	//
	FeedItemValidationStatusVALID FeedItemValidationStatus = "VALID"
)

//
// Feed hard type. Values coincide with placeholder type id.
//
type FeedType string

const (
	FeedTypeNONE FeedType = "NONE"

	//
	// Sitelink placeholder typed
	//
	FeedTypeSITELINK FeedType = "SITELINK"

	//
	// Call placeholder typed
	//
	FeedTypeCALL FeedType = "CALL"

	//
	// App placeholder typed
	//
	FeedTypeAPP FeedType = "APP"

	//
	// Review placeholder typed
	//
	FeedTypeREVIEW FeedType = "REVIEW"

	//
	// AdCustomizer placeholder typed
	//
	FeedTypeAD_CUSTOMIZER FeedType = "AD_CUSTOMIZER"

	//
	// Callout placeholder typed
	//
	FeedTypeCALLOUT FeedType = "CALLOUT"

	//
	// Structured snippets placeholder typed
	//
	FeedTypeSTRUCTURED_SNIPPET FeedType = "STRUCTURED_SNIPPET"

	//
	// Message placeholder typed
	//
	FeedTypeMESSAGE FeedType = "MESSAGE"

	//
	// Price placeholder typed
	//
	FeedTypePRICE FeedType = "PRICE"

	//
	// Promotion placeholder typed
	//
	FeedTypePROMOTION FeedType = "PROMOTION"
)

//
// A restriction used to determine if the request context's geo should be matched.
//
type GeoRestriction string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	GeoRestrictionUNKNOWN GeoRestriction = "UNKNOWN"

	//
	// Indicates that request context should match the physical location of the user.
	//
	GeoRestrictionLOCATION_OF_PRESENCE GeoRestriction = "LOCATION_OF_PRESENCE"
)

//
// The reasons for the target error.
//
type IdErrorReason string

const (

	//
	// Id not found
	//
	IdErrorReasonNOT_FOUND IdErrorReason = "NOT_FOUND"
)

//
// The single reason for the internal API error.
//
type InternalApiErrorReason string

const (

	//
	// API encountered an unexpected internal error.
	//
	InternalApiErrorReasonUNEXPECTED_INTERNAL_API_ERROR InternalApiErrorReason = "UNEXPECTED_INTERNAL_API_ERROR"

	//
	// A temporary error occurred during the request. Please retry.
	//
	InternalApiErrorReasonTRANSIENT_ERROR InternalApiErrorReason = "TRANSIENT_ERROR"

	//
	// The cause of the error is not known or only defined in newer versions.
	//
	InternalApiErrorReasonUNKNOWN InternalApiErrorReason = "UNKNOWN"

	//
	// The API is currently unavailable for a planned downtime.
	//
	InternalApiErrorReasonDOWNTIME InternalApiErrorReason = "DOWNTIME"

	//
	// Mutate succeeded but server was unable to build response. Client should not retry mutate.
	//
	InternalApiErrorReasonERROR_GENERATING_RESPONSE InternalApiErrorReason = "ERROR_GENERATING_RESPONSE"
)

//
// Match type of a keyword. i.e. the way we match a keyword string with
// search queries.
//
type KeywordMatchType string

const (

	//
	// Exact match
	//
	KeywordMatchTypeEXACT KeywordMatchType = "EXACT"

	//
	// Phrase match
	//
	KeywordMatchTypePHRASE KeywordMatchType = "PHRASE"

	//
	// Broad match
	//
	KeywordMatchTypeBROAD KeywordMatchType = "BROAD"
)

//
// Enum that represents the different Targeting Status values for a Location criterion.
//
type LocationTargetingStatus string

const (

	//
	// The location is active.
	//
	LocationTargetingStatusACTIVE LocationTargetingStatus = "ACTIVE"

	//
	// The location is not available for targeting.
	//
	LocationTargetingStatusOBSOLETE LocationTargetingStatus = "OBSOLETE"

	//
	// The location is phasing out, it will marked obsolete soon.
	//
	LocationTargetingStatusPHASING_OUT LocationTargetingStatus = "PHASING_OUT"
)

//
// Minutes in an hour.  Currently only 0, 15, 30, and 45 are supported
//
type MinuteOfHour string

const (

	//
	// Zero minutes past hour.
	//
	MinuteOfHourZERO MinuteOfHour = "ZERO"

	//
	// Fifteen minutes past hour.
	//
	MinuteOfHourFIFTEEN MinuteOfHour = "FIFTEEN"

	//
	// Thirty minutes past hour.
	//
	MinuteOfHourTHIRTY MinuteOfHour = "THIRTY"

	//
	// Forty-five minutes past hour.
	//
	MinuteOfHourFORTY_FIVE MinuteOfHour = "FORTY_FIVE"
)

type NewEntityCreationErrorReason string

const (

	//
	// Do not set the id field while creating new entities.
	//
	NewEntityCreationErrorReasonCANNOT_SET_ID_FOR_ADD NewEntityCreationErrorReason = "CANNOT_SET_ID_FOR_ADD"

	//
	// Creating more than one entity with the same temp ID is not allowed.
	//
	NewEntityCreationErrorReasonDUPLICATE_TEMP_IDS NewEntityCreationErrorReason = "DUPLICATE_TEMP_IDS"

	//
	// Parent object with specified temp id failed validation, so no deep
	// validation will be done for this child entity.
	//
	NewEntityCreationErrorReasonTEMP_ID_ENTITY_HAD_ERRORS NewEntityCreationErrorReason = "TEMP_ID_ENTITY_HAD_ERRORS"
)

//
// The reasons for the validation error.
//
type NotEmptyErrorReason string

const (
	NotEmptyErrorReasonEMPTY_LIST NotEmptyErrorReason = "EMPTY_LIST"
)

//
// The reasons for the validation error.
//
type NullErrorReason string

const (

	//
	// Specified list/container must not contain any null elements
	//
	NullErrorReasonNULL_CONTENT NullErrorReason = "NULL_CONTENT"
)

//
// The reasons for the operation access error.
//
type OperationAccessDeniedReason string

const (

	//
	// Unauthorized invocation of a service's method (get, mutate, etc.)
	//
	OperationAccessDeniedReasonACTION_NOT_PERMITTED OperationAccessDeniedReason = "ACTION_NOT_PERMITTED"

	//
	// Unauthorized ADD operation in invoking a service's mutate method.
	//
	OperationAccessDeniedReasonADD_OPERATION_NOT_PERMITTED OperationAccessDeniedReason = "ADD_OPERATION_NOT_PERMITTED"

	//
	// Unauthorized REMOVE operation in invoking a service's mutate method.
	//
	OperationAccessDeniedReasonREMOVE_OPERATION_NOT_PERMITTED OperationAccessDeniedReason = "REMOVE_OPERATION_NOT_PERMITTED"

	//
	// Unauthorized SET operation in invoking a service's mutate method.
	//
	OperationAccessDeniedReasonSET_OPERATION_NOT_PERMITTED OperationAccessDeniedReason = "SET_OPERATION_NOT_PERMITTED"

	//
	// A mutate action is not allowed on this campaign, from this client.
	//
	OperationAccessDeniedReasonMUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT OperationAccessDeniedReason = "MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT"

	//
	// This operation is not permitted on this campaign type
	//
	OperationAccessDeniedReasonOPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE OperationAccessDeniedReason = "OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE"

	//
	// An ADD operation may not set status to REMOVED.
	//
	OperationAccessDeniedReasonADD_AS_REMOVED_NOT_PERMITTED OperationAccessDeniedReason = "ADD_AS_REMOVED_NOT_PERMITTED"

	//
	// This operation is not allowed because the campaign or adgroup is removed.
	//
	OperationAccessDeniedReasonOPERATION_NOT_PERMITTED_FOR_REMOVED_ENTITY OperationAccessDeniedReason = "OPERATION_NOT_PERMITTED_FOR_REMOVED_ENTITY"

	//
	// This operation is not permitted on this ad group type.
	//
	OperationAccessDeniedReasonOPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE OperationAccessDeniedReason = "OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE"

	//
	// The reason the invoked method or operation is prohibited is not known
	// (the client may be of an older version than the server).
	//
	OperationAccessDeniedReasonUNKNOWN OperationAccessDeniedReason = "UNKNOWN"
)

//
// This represents an operator that may be presented to an adsapi service.
//
type Operator string

const (

	//
	// The ADD operator.
	//
	OperatorADD Operator = "ADD"

	//
	// The REMOVE operator.
	//
	OperatorREMOVE Operator = "REMOVE"

	//
	// The SET operator (used for updates).
	//
	OperatorSET Operator = "SET"
)

//
// The reasons for the validation error.
//
type OperatorErrorReason string

const (
	OperatorErrorReasonOPERATOR_NOT_SUPPORTED OperatorErrorReason = "OPERATOR_NOT_SUPPORTED"
)

//
// Defines the valid set of operators.
//
type PredicateOperator string

const (

	//
	// Checks if the field is equal to the given value.
	//
	// <p>This operator is used with integers, dates, booleans, strings,
	// enums, and sets.
	//
	PredicateOperatorEQUALS PredicateOperator = "EQUALS"

	//
	// Checks if the field does not equal the given value.
	//
	// <p>This operator is used with integers, booleans, strings, enums,
	// and sets.
	//
	PredicateOperatorNOT_EQUALS PredicateOperator = "NOT_EQUALS"

	//
	// Checks if the field is equal to one of the given values.
	//
	// <p>This operator accepts multiple operands and is used with
	// integers, booleans, strings, and enums.
	//
	PredicateOperatorIN PredicateOperator = "IN"

	//
	// Checks if the field does not equal any of the given values.
	//
	// <p>This operator accepts multiple operands and is used with
	// integers, booleans, strings, and enums.
	//
	PredicateOperatorNOT_IN PredicateOperator = "NOT_IN"

	//
	// Checks if the field is greater than the given value.
	//
	// <p>This operator is used with numbers and dates.
	//
	PredicateOperatorGREATER_THAN PredicateOperator = "GREATER_THAN"

	//
	// Checks if the field is greater or equal to the given value.
	//
	// <p>This operator is used with numbers and dates.
	//
	PredicateOperatorGREATER_THAN_EQUALS PredicateOperator = "GREATER_THAN_EQUALS"

	//
	// Checks if the field is less than the given value.
	//
	// <p>This operator is used with numbers and dates.
	//
	PredicateOperatorLESS_THAN PredicateOperator = "LESS_THAN"

	//
	// Checks if the field is less or equal to than the given value.
	//
	// <p>This operator is used with numbers and dates.
	//
	PredicateOperatorLESS_THAN_EQUALS PredicateOperator = "LESS_THAN_EQUALS"

	//
	// Checks if the field starts with the given value.
	//
	// <p>This operator is used with strings.
	//
	PredicateOperatorSTARTS_WITH PredicateOperator = "STARTS_WITH"

	//
	// Checks if the field starts with the given value, ignoring case.
	//
	// <p>This operator is used with strings.
	//
	PredicateOperatorSTARTS_WITH_IGNORE_CASE PredicateOperator = "STARTS_WITH_IGNORE_CASE"

	//
	// Checks if the field contains the given value as a substring.
	//
	// <p>This operator is used with strings.
	//
	PredicateOperatorCONTAINS PredicateOperator = "CONTAINS"

	//
	// Checks if the field contains the given value as a substring, ignoring
	// case.
	//
	// <p>This operator is used with strings.
	//
	PredicateOperatorCONTAINS_IGNORE_CASE PredicateOperator = "CONTAINS_IGNORE_CASE"

	//
	// Checks if the field does not contain the given value as a substring.
	//
	// <p>This operator is used with strings.
	//
	PredicateOperatorDOES_NOT_CONTAIN PredicateOperator = "DOES_NOT_CONTAIN"

	//
	// Checks if the field does not contain the given value as a substring,
	// ignoring case.
	//
	// <p>This operator is used with strings.
	//
	PredicateOperatorDOES_NOT_CONTAIN_IGNORE_CASE PredicateOperator = "DOES_NOT_CONTAIN_IGNORE_CASE"

	//
	// Checks if the field contains <em>any</em> of the given values.
	//
	// <p>This operator accepts multiple values and is used on sets of numbers
	// or strings.
	//
	PredicateOperatorCONTAINS_ANY PredicateOperator = "CONTAINS_ANY"

	//
	// Checks if the field contains <em>all</em> of the given values.
	//
	// <p>This operator accepts multiple values and is used on sets of numbers
	// or strings.
	//
	PredicateOperatorCONTAINS_ALL PredicateOperator = "CONTAINS_ALL"

	//
	// Checks if the field contains <em>none</em> of the given values.
	//
	// <p>This operator accepts multiple values and is used on sets of numbers
	// or strings.
	//
	PredicateOperatorCONTAINS_NONE PredicateOperator = "CONTAINS_NONE"

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	PredicateOperatorUNKNOWN PredicateOperator = "UNKNOWN"
)

//
// The qualifier on the price for all Price items.
//
type PriceExtensionPriceQualifier string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	PriceExtensionPriceQualifierUNKNOWN PriceExtensionPriceQualifier = "UNKNOWN"

	//
	// 'From' qualifier for the price.
	//
	PriceExtensionPriceQualifierFROM PriceExtensionPriceQualifier = "FROM"

	//
	// 'Up to' qualifier for the price.
	//
	PriceExtensionPriceQualifierUP_TO PriceExtensionPriceQualifier = "UP_TO"

	//
	// 'Average' qualifier for the price.
	//
	PriceExtensionPriceQualifierAVERAGE PriceExtensionPriceQualifier = "AVERAGE"

	//
	// None is used for clearing the qualifier.
	//
	PriceExtensionPriceQualifierNONE PriceExtensionPriceQualifier = "NONE"
)

//
// The price unit for a Price table item.
//
type PriceExtensionPriceUnit string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	PriceExtensionPriceUnitUNKNOWN PriceExtensionPriceUnit = "UNKNOWN"

	//
	// Per hour.
	//
	PriceExtensionPriceUnitPER_HOUR PriceExtensionPriceUnit = "PER_HOUR"

	//
	// Per day.
	//
	PriceExtensionPriceUnitPER_DAY PriceExtensionPriceUnit = "PER_DAY"

	//
	// Per week.
	//
	PriceExtensionPriceUnitPER_WEEK PriceExtensionPriceUnit = "PER_WEEK"

	//
	// Per month.
	//
	PriceExtensionPriceUnitPER_MONTH PriceExtensionPriceUnit = "PER_MONTH"

	//
	// Per year.
	//
	PriceExtensionPriceUnitPER_YEAR PriceExtensionPriceUnit = "PER_YEAR"

	//
	// Per night.
	//
	PriceExtensionPriceUnitPER_NIGHT PriceExtensionPriceUnit = "PER_NIGHT"

	//
	// None is used for clearing the unit.
	//
	PriceExtensionPriceUnitNONE PriceExtensionPriceUnit = "NONE"
)

//
// The type of a price extension represents.
//
type PriceExtensionType string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	PriceExtensionTypeUNKNOWN PriceExtensionType = "UNKNOWN"

	//
	// The type for showing a list of brands.
	//
	PriceExtensionTypeBRANDS PriceExtensionType = "BRANDS"

	//
	// The type for showing a list of events.
	//
	PriceExtensionTypeEVENTS PriceExtensionType = "EVENTS"

	//
	// The type for showing locations relevant to your business.
	//
	PriceExtensionTypeLOCATIONS PriceExtensionType = "LOCATIONS"

	//
	// The type for showing sub-regions or districts within a city or region.
	//
	PriceExtensionTypeNEIGHBORHOODS PriceExtensionType = "NEIGHBORHOODS"

	//
	// The type for showing a collection of product categories.
	//
	PriceExtensionTypePRODUCT_CATEGORIES PriceExtensionType = "PRODUCT_CATEGORIES"

	//
	// The type for showing a collection of related product tiers.
	//
	PriceExtensionTypePRODUCT_TIERS PriceExtensionType = "PRODUCT_TIERS"

	//
	// The type for showing a collection of services offered by your business.
	//
	PriceExtensionTypeSERVICES PriceExtensionType = "SERVICES"

	//
	// The type for showing a collection of service categories.
	//
	PriceExtensionTypeSERVICE_CATEGORIES PriceExtensionType = "SERVICE_CATEGORIES"

	//
	// The type for showing a collection of related service tiers.
	//
	PriceExtensionTypeSERVICE_TIERS PriceExtensionType = "SERVICE_TIERS"
)

//
// Qualification for a promotion extension discount.
//
type PromotionExtensionDiscountModifier string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	PromotionExtensionDiscountModifierUNKNOWN PromotionExtensionDiscountModifier = "UNKNOWN"

	//
	// /'Up to'/.
	//
	PromotionExtensionDiscountModifierUP_TO PromotionExtensionDiscountModifier = "UP_TO"

	//
	// None is used for clearing the discount modifier.
	//
	PromotionExtensionDiscountModifierNONE PromotionExtensionDiscountModifier = "NONE"
)

//
// The occasion of a promotion extension.
//
type PromotionExtensionOccasion string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	PromotionExtensionOccasionUNKNOWN PromotionExtensionOccasion = "UNKNOWN"

	//
	// New Year's.
	//
	PromotionExtensionOccasionNEW_YEARS PromotionExtensionOccasion = "NEW_YEARS"

	//
	// Valentine's Day.
	//
	PromotionExtensionOccasionVALENTINES_DAY PromotionExtensionOccasion = "VALENTINES_DAY"

	//
	// Easter.
	//
	PromotionExtensionOccasionEASTER PromotionExtensionOccasion = "EASTER"

	//
	// Mother's Day.
	//
	PromotionExtensionOccasionMOTHERS_DAY PromotionExtensionOccasion = "MOTHERS_DAY"

	//
	// Father's Day.
	//
	PromotionExtensionOccasionFATHERS_DAY PromotionExtensionOccasion = "FATHERS_DAY"

	//
	// Labor Day.
	//
	PromotionExtensionOccasionLABOR_DAY PromotionExtensionOccasion = "LABOR_DAY"

	//
	// Back To School.
	//
	PromotionExtensionOccasionBACK_TO_SCHOOL PromotionExtensionOccasion = "BACK_TO_SCHOOL"

	//
	// Halloween.
	//
	PromotionExtensionOccasionHALLOWEEN PromotionExtensionOccasion = "HALLOWEEN"

	//
	// Black Friday.
	//
	PromotionExtensionOccasionBLACK_FRIDAY PromotionExtensionOccasion = "BLACK_FRIDAY"

	//
	// Cyber Monday.
	//
	PromotionExtensionOccasionCYBER_MONDAY PromotionExtensionOccasion = "CYBER_MONDAY"

	//
	// Christmas.
	//
	PromotionExtensionOccasionCHRISTMAS PromotionExtensionOccasion = "CHRISTMAS"

	//
	// Boxing Day.
	//
	PromotionExtensionOccasionBOXING_DAY PromotionExtensionOccasion = "BOXING_DAY"

	//
	// None is used for clearing the occasion.
	//
	PromotionExtensionOccasionNONE PromotionExtensionOccasion = "NONE"
)

//
// The reason for the query error.
//
type QueryErrorReason string

const (

	//
	// Exception that happens when trying to parse a query that doesn't match the AWQL grammar.
	//
	QueryErrorReasonPARSING_FAILED QueryErrorReason = "PARSING_FAILED"

	//
	// The provided query is an empty string.
	//
	QueryErrorReasonMISSING_QUERY QueryErrorReason = "MISSING_QUERY"

	//
	// The query does not contain the required SELECT clause or it is not in the
	// correct location.
	//
	QueryErrorReasonMISSING_SELECT_CLAUSE QueryErrorReason = "MISSING_SELECT_CLAUSE"

	//
	// The query does not contain the required FROM clause or it is not in the correct location.
	//
	QueryErrorReasonMISSING_FROM_CLAUSE QueryErrorReason = "MISSING_FROM_CLAUSE"

	//
	// The SELECT clause could not be parsed.
	//
	QueryErrorReasonINVALID_SELECT_CLAUSE QueryErrorReason = "INVALID_SELECT_CLAUSE"

	//
	// The FROM clause could not be parsed.
	//
	QueryErrorReasonINVALID_FROM_CLAUSE QueryErrorReason = "INVALID_FROM_CLAUSE"

	//
	// The WHERE clause could not be parsed.
	//
	QueryErrorReasonINVALID_WHERE_CLAUSE QueryErrorReason = "INVALID_WHERE_CLAUSE"

	//
	// The ORDER BY clause could not be parsed.
	//
	QueryErrorReasonINVALID_ORDER_BY_CLAUSE QueryErrorReason = "INVALID_ORDER_BY_CLAUSE"

	//
	// The LIMIT clause could not be parsed.
	//
	QueryErrorReasonINVALID_LIMIT_CLAUSE QueryErrorReason = "INVALID_LIMIT_CLAUSE"

	//
	// The startIndex in the LIMIT clause does not contain a valid integer.
	//
	QueryErrorReasonINVALID_START_INDEX_IN_LIMIT_CLAUSE QueryErrorReason = "INVALID_START_INDEX_IN_LIMIT_CLAUSE"

	//
	// The pageSize in the LIMIT clause does not contain a valid integer.
	//
	QueryErrorReasonINVALID_PAGE_SIZE_IN_LIMIT_CLAUSE QueryErrorReason = "INVALID_PAGE_SIZE_IN_LIMIT_CLAUSE"

	//
	// The DURING clause could not be parsed.
	//
	QueryErrorReasonINVALID_DURING_CLAUSE QueryErrorReason = "INVALID_DURING_CLAUSE"

	//
	// The minimum date in the DURING clause is not a valid date in YYYYMMDD format.
	//
	QueryErrorReasonINVALID_MIN_DATE_IN_DURING_CLAUSE QueryErrorReason = "INVALID_MIN_DATE_IN_DURING_CLAUSE"

	//
	// The maximum date in the DURING clause is not a valid date in YYYYMMDD format.
	//
	QueryErrorReasonINVALID_MAX_DATE_IN_DURING_CLAUSE QueryErrorReason = "INVALID_MAX_DATE_IN_DURING_CLAUSE"

	//
	// The minimum date in the DURING is after the maximum date.
	//
	QueryErrorReasonMAX_LESS_THAN_MIN_IN_DURING_CLAUSE QueryErrorReason = "MAX_LESS_THAN_MIN_IN_DURING_CLAUSE"

	//
	// The query matched the grammar, but is invalid in some way such as using a service that
	// isn't supported.
	//
	QueryErrorReasonVALIDATION_FAILED QueryErrorReason = "VALIDATION_FAILED"
)

//
// Enums for all the reasons an error can be thrown to the user during
// billing quota checks.
//
type QuotaCheckErrorReason string

const (

	//
	// Customer passed in an invalid token in the header.
	//
	QuotaCheckErrorReasonINVALID_TOKEN_HEADER QuotaCheckErrorReason = "INVALID_TOKEN_HEADER"

	//
	// Customer is marked delinquent.
	//
	QuotaCheckErrorReasonACCOUNT_DELINQUENT QuotaCheckErrorReason = "ACCOUNT_DELINQUENT"

	//
	// Customer is a fraudulent.
	//
	QuotaCheckErrorReasonACCOUNT_INACCESSIBLE QuotaCheckErrorReason = "ACCOUNT_INACCESSIBLE"

	//
	// Inactive Account.
	//
	QuotaCheckErrorReasonACCOUNT_INACTIVE QuotaCheckErrorReason = "ACCOUNT_INACTIVE"

	//
	// Signup not complete
	//
	QuotaCheckErrorReasonINCOMPLETE_SIGNUP QuotaCheckErrorReason = "INCOMPLETE_SIGNUP"

	//
	// Developer token is not approved for production access, and the customer
	// is attempting to access a production account.
	//
	QuotaCheckErrorReasonDEVELOPER_TOKEN_NOT_APPROVED QuotaCheckErrorReason = "DEVELOPER_TOKEN_NOT_APPROVED"

	//
	// Terms and conditions are not signed.
	//
	QuotaCheckErrorReasonTERMS_AND_CONDITIONS_NOT_SIGNED QuotaCheckErrorReason = "TERMS_AND_CONDITIONS_NOT_SIGNED"

	//
	// Monthly budget quota reached.
	//
	QuotaCheckErrorReasonMONTHLY_BUDGET_REACHED QuotaCheckErrorReason = "MONTHLY_BUDGET_REACHED"

	//
	// Monthly budget quota exceeded.
	//
	QuotaCheckErrorReasonQUOTA_EXCEEDED QuotaCheckErrorReason = "QUOTA_EXCEEDED"
)

//
// The reasons for the target error.
//
type RangeErrorReason string

const (
	RangeErrorReasonTOO_LOW RangeErrorReason = "TOO_LOW"

	RangeErrorReasonTOO_HIGH RangeErrorReason = "TOO_HIGH"
)

//
// The reason for the rate exceeded error.
//
type RateExceededErrorReason string

const (

	//
	// Rate exceeded.
	//
	RateExceededErrorReasonRATE_EXCEEDED RateExceededErrorReason = "RATE_EXCEEDED"
)

//
// The reasons for the target error.
//
type ReadOnlyErrorReason string

const (
	ReadOnlyErrorReasonREAD_ONLY ReadOnlyErrorReason = "READ_ONLY"
)

//
// The reasons for the target error.
//
type RejectedErrorReason string

const (

	//
	// Unknown value encountered
	//
	RejectedErrorReasonUNKNOWN_VALUE RejectedErrorReason = "UNKNOWN_VALUE"
)

type RequestErrorReason string

const (

	//
	// Error reason is unknown.
	//
	RequestErrorReasonUNKNOWN RequestErrorReason = "UNKNOWN"

	//
	// Invalid input.
	//
	RequestErrorReasonINVALID_INPUT RequestErrorReason = "INVALID_INPUT"

	//
	// The api version in the request has been discontinued. Please update
	// to the new AdWords API version.
	//
	RequestErrorReasonUNSUPPORTED_VERSION RequestErrorReason = "UNSUPPORTED_VERSION"
)

//
// The reasons for the target error.
//
type RequiredErrorReason string

const (

	//
	// Missing required field.
	//
	RequiredErrorReasonREQUIRED RequiredErrorReason = "REQUIRED"
)

//
// The reasons for the target error.
//
type SelectorErrorReason string

const (

	//
	// The field name is not valid.
	//
	SelectorErrorReasonINVALID_FIELD_NAME SelectorErrorReason = "INVALID_FIELD_NAME"

	//
	// The list of fields is null or empty.
	//
	SelectorErrorReasonMISSING_FIELDS SelectorErrorReason = "MISSING_FIELDS"

	//
	// The list of predicates is null or empty.
	//
	SelectorErrorReasonMISSING_PREDICATES SelectorErrorReason = "MISSING_PREDICATES"

	//
	// Predicate operator does not support multiple values. Multiple values are
	// supported only for {@link Predicate.Operator#IN} and {@link Predicate.Operator#NOT_IN}.
	//
	SelectorErrorReasonOPERATOR_DOES_NOT_SUPPORT_MULTIPLE_VALUES SelectorErrorReason = "OPERATOR_DOES_NOT_SUPPORT_MULTIPLE_VALUES"

	//
	// The predicate enum value is not valid.
	//
	SelectorErrorReasonINVALID_PREDICATE_ENUM_VALUE SelectorErrorReason = "INVALID_PREDICATE_ENUM_VALUE"

	//
	// The predicate operator is empty.
	//
	SelectorErrorReasonMISSING_PREDICATE_OPERATOR SelectorErrorReason = "MISSING_PREDICATE_OPERATOR"

	//
	// The predicate values are empty.
	//
	SelectorErrorReasonMISSING_PREDICATE_VALUES SelectorErrorReason = "MISSING_PREDICATE_VALUES"

	//
	// The predicate field name is not valid.
	//
	SelectorErrorReasonINVALID_PREDICATE_FIELD_NAME SelectorErrorReason = "INVALID_PREDICATE_FIELD_NAME"

	//
	// The predicate operator is not valid.
	//
	SelectorErrorReasonINVALID_PREDICATE_OPERATOR SelectorErrorReason = "INVALID_PREDICATE_OPERATOR"

	//
	// Invalid selection of fields.
	//
	SelectorErrorReasonINVALID_FIELD_SELECTION SelectorErrorReason = "INVALID_FIELD_SELECTION"

	//
	// The predicate value is not valid.
	//
	SelectorErrorReasonINVALID_PREDICATE_VALUE SelectorErrorReason = "INVALID_PREDICATE_VALUE"

	//
	// The sort field name is not valid or the field is not sortable.
	//
	SelectorErrorReasonINVALID_SORT_FIELD_NAME SelectorErrorReason = "INVALID_SORT_FIELD_NAME"

	//
	// Standard error.
	//
	SelectorErrorReasonSELECTOR_ERROR SelectorErrorReason = "SELECTOR_ERROR"

	//
	// Filtering by date range is not supported.
	//
	SelectorErrorReasonFILTER_BY_DATE_RANGE_NOT_SUPPORTED SelectorErrorReason = "FILTER_BY_DATE_RANGE_NOT_SUPPORTED"

	//
	// Selector paging start index is too high.
	//
	SelectorErrorReasonSTART_INDEX_IS_TOO_HIGH SelectorErrorReason = "START_INDEX_IS_TOO_HIGH"

	//
	// The values list in a predicate was too long.
	//
	SelectorErrorReasonTOO_MANY_PREDICATE_VALUES SelectorErrorReason = "TOO_MANY_PREDICATE_VALUES"

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	SelectorErrorReasonUNKNOWN_ERROR SelectorErrorReason = "UNKNOWN_ERROR"
)

//
// The reasons for Ad Scheduling errors.
//
type SizeLimitErrorReason string

const (

	//
	// The number of entries in the request exceeds the system limit.
	//
	SizeLimitErrorReasonREQUEST_SIZE_LIMIT_EXCEEDED SizeLimitErrorReason = "REQUEST_SIZE_LIMIT_EXCEEDED"

	//
	// The number of entries in the response exceeds the system limit.
	//
	SizeLimitErrorReasonRESPONSE_SIZE_LIMIT_EXCEEDED SizeLimitErrorReason = "RESPONSE_SIZE_LIMIT_EXCEEDED"

	//
	// The account is too large to load.
	//
	SizeLimitErrorReasonINTERNAL_STORAGE_ERROR SizeLimitErrorReason = "INTERNAL_STORAGE_ERROR"

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	SizeLimitErrorReasonUNKNOWN SizeLimitErrorReason = "UNKNOWN"
)

//
// Possible orders of sorting.
//
type SortOrder string

const (
	SortOrderASCENDING SortOrder = "ASCENDING"

	SortOrderDESCENDING SortOrder = "DESCENDING"
)

//
// The reasons for the target error.
//
type StringFormatErrorReason string

const (
	StringFormatErrorReasonUNKNOWN StringFormatErrorReason = "UNKNOWN"

	//
	// The input string value contains disallowed characters.
	//
	StringFormatErrorReasonILLEGAL_CHARS StringFormatErrorReason = "ILLEGAL_CHARS"

	//
	// The input string value is invalid for the associated field.
	//
	StringFormatErrorReasonINVALID_FORMAT StringFormatErrorReason = "INVALID_FORMAT"
)

//
// The reasons for the target error.
//
type StringLengthErrorReason string

const (
	StringLengthErrorReasonTOO_SHORT StringLengthErrorReason = "TOO_SHORT"

	StringLengthErrorReasonTOO_LONG StringLengthErrorReason = "TOO_LONG"
)

//
// The reasons for the url error.
//
type UrlErrorReason string

const (

	//
	// The tracking url template is invalid.
	//
	UrlErrorReasonINVALID_TRACKING_URL_TEMPLATE UrlErrorReason = "INVALID_TRACKING_URL_TEMPLATE"

	//
	// The tracking url template contains invalid tag.
	//
	UrlErrorReasonINVALID_TAG_IN_TRACKING_URL_TEMPLATE UrlErrorReason = "INVALID_TAG_IN_TRACKING_URL_TEMPLATE"

	//
	// The tracking url template must contain at least one tag (e.g. {lpurl}),
	// This applies only to tracking url template associated with website ads or product ads.
	//
	UrlErrorReasonMISSING_TRACKING_URL_TEMPLATE_TAG UrlErrorReason = "MISSING_TRACKING_URL_TEMPLATE_TAG"

	//
	// The tracking url template must start with a valid protocol (or lpurl tag).
	//
	UrlErrorReasonMISSING_PROTOCOL_IN_TRACKING_URL_TEMPLATE UrlErrorReason = "MISSING_PROTOCOL_IN_TRACKING_URL_TEMPLATE"

	//
	// The tracking url template starts with an invalid protocol.
	//
	UrlErrorReasonINVALID_PROTOCOL_IN_TRACKING_URL_TEMPLATE UrlErrorReason = "INVALID_PROTOCOL_IN_TRACKING_URL_TEMPLATE"

	//
	// The tracking url template  contains illegal characters.
	//
	UrlErrorReasonMALFORMED_TRACKING_URL_TEMPLATE UrlErrorReason = "MALFORMED_TRACKING_URL_TEMPLATE"

	//
	// The tracking url template must contain a host name (or lpurl tag).
	//
	UrlErrorReasonMISSING_HOST_IN_TRACKING_URL_TEMPLATE UrlErrorReason = "MISSING_HOST_IN_TRACKING_URL_TEMPLATE"

	//
	// The tracking url template has an invalid or missing top level domain extension.
	//
	UrlErrorReasonINVALID_TLD_IN_TRACKING_URL_TEMPLATE UrlErrorReason = "INVALID_TLD_IN_TRACKING_URL_TEMPLATE"

	//
	// The tracking url template contains nested occurrences of the same conditional tag
	// (i.e. {ifmobile:{ifmobile:x}}).
	//
	UrlErrorReasonREDUNDANT_NESTED_TRACKING_URL_TEMPLATE_TAG UrlErrorReason = "REDUNDANT_NESTED_TRACKING_URL_TEMPLATE_TAG"

	//
	// The final url is invalid.
	//
	UrlErrorReasonINVALID_FINAL_URL UrlErrorReason = "INVALID_FINAL_URL"

	//
	// The final url contains invalid tag.
	//
	UrlErrorReasonINVALID_TAG_IN_FINAL_URL UrlErrorReason = "INVALID_TAG_IN_FINAL_URL"

	//
	// The final url contains nested occurrences of the same conditional tag
	// (i.e. {ifmobile:{ifmobile:x}}).
	//
	UrlErrorReasonREDUNDANT_NESTED_FINAL_URL_TAG UrlErrorReason = "REDUNDANT_NESTED_FINAL_URL_TAG"

	//
	// The final url must start with a valid protocol.
	//
	UrlErrorReasonMISSING_PROTOCOL_IN_FINAL_URL UrlErrorReason = "MISSING_PROTOCOL_IN_FINAL_URL"

	//
	// The final url starts with an invalid protocol.
	//
	UrlErrorReasonINVALID_PROTOCOL_IN_FINAL_URL UrlErrorReason = "INVALID_PROTOCOL_IN_FINAL_URL"

	//
	// The final url  contains illegal characters.
	//
	UrlErrorReasonMALFORMED_FINAL_URL UrlErrorReason = "MALFORMED_FINAL_URL"

	//
	// The final url must contain a host name.
	//
	UrlErrorReasonMISSING_HOST_IN_FINAL_URL UrlErrorReason = "MISSING_HOST_IN_FINAL_URL"

	//
	// The tracking url template has an invalid or missing top level domain extension.
	//
	UrlErrorReasonINVALID_TLD_IN_FINAL_URL UrlErrorReason = "INVALID_TLD_IN_FINAL_URL"

	//
	// The final mobile url is invalid.
	//
	UrlErrorReasonINVALID_FINAL_MOBILE_URL UrlErrorReason = "INVALID_FINAL_MOBILE_URL"

	//
	// The final mobile url contains invalid tag.
	//
	UrlErrorReasonINVALID_TAG_IN_FINAL_MOBILE_URL UrlErrorReason = "INVALID_TAG_IN_FINAL_MOBILE_URL"

	//
	// The final mobile url contains nested occurrences of the same conditional tag
	// (i.e. {ifmobile:{ifmobile:x}}).
	//
	UrlErrorReasonREDUNDANT_NESTED_FINAL_MOBILE_URL_TAG UrlErrorReason = "REDUNDANT_NESTED_FINAL_MOBILE_URL_TAG"

	//
	// The final mobile url must start with a valid protocol.
	//
	UrlErrorReasonMISSING_PROTOCOL_IN_FINAL_MOBILE_URL UrlErrorReason = "MISSING_PROTOCOL_IN_FINAL_MOBILE_URL"

	//
	// The final mobile url starts with an invalid protocol.
	//
	UrlErrorReasonINVALID_PROTOCOL_IN_FINAL_MOBILE_URL UrlErrorReason = "INVALID_PROTOCOL_IN_FINAL_MOBILE_URL"

	//
	// The final mobile url  contains illegal characters.
	//
	UrlErrorReasonMALFORMED_FINAL_MOBILE_URL UrlErrorReason = "MALFORMED_FINAL_MOBILE_URL"

	//
	// The final mobile url must contain a host name.
	//
	UrlErrorReasonMISSING_HOST_IN_FINAL_MOBILE_URL UrlErrorReason = "MISSING_HOST_IN_FINAL_MOBILE_URL"

	//
	// The tracking url template has an invalid or missing top level domain extension.
	//
	UrlErrorReasonINVALID_TLD_IN_FINAL_MOBILE_URL UrlErrorReason = "INVALID_TLD_IN_FINAL_MOBILE_URL"

	//
	// The final app url is invalid.
	//
	UrlErrorReasonINVALID_FINAL_APP_URL UrlErrorReason = "INVALID_FINAL_APP_URL"

	//
	// The final app url contains invalid tag.
	//
	UrlErrorReasonINVALID_TAG_IN_FINAL_APP_URL UrlErrorReason = "INVALID_TAG_IN_FINAL_APP_URL"

	//
	// The final app url contains nested occurrences of the same conditional tag
	// (i.e. {ifmobile:{ifmobile:x}}).
	//
	UrlErrorReasonREDUNDANT_NESTED_FINAL_APP_URL_TAG UrlErrorReason = "REDUNDANT_NESTED_FINAL_APP_URL_TAG"

	//
	// More than one app url found for the same OS type.
	//
	UrlErrorReasonMULTIPLE_APP_URLS_FOR_OSTYPE UrlErrorReason = "MULTIPLE_APP_URLS_FOR_OSTYPE"

	//
	// The OS type given for an app url is not valid.
	//
	UrlErrorReasonINVALID_OSTYPE UrlErrorReason = "INVALID_OSTYPE"

	//
	// The protocol given for an app url is not valid. (E.g. "android-app://")
	//
	UrlErrorReasonINVALID_PROTOCOL_FOR_APP_URL UrlErrorReason = "INVALID_PROTOCOL_FOR_APP_URL"

	//
	// The package id (app id) given for an app url is not valid.
	//
	UrlErrorReasonINVALID_PACKAGE_ID_FOR_APP_URL UrlErrorReason = "INVALID_PACKAGE_ID_FOR_APP_URL"

	//
	// The number of url custom parameters for an entity exceeds the maximum limit allowed.
	//
	UrlErrorReasonURL_CUSTOM_PARAMETERS_COUNT_EXCEEDS_LIMIT UrlErrorReason = "URL_CUSTOM_PARAMETERS_COUNT_EXCEEDS_LIMIT"

	//
	// The parameter has isRemove set to true but a value that is non-null.
	//
	UrlErrorReasonURL_CUSTOM_PARAMETER_REMOVAL_WITH_NON_NULL_VALUE UrlErrorReason = "URL_CUSTOM_PARAMETER_REMOVAL_WITH_NON_NULL_VALUE"

	//
	// For add operations, there will not be any existing parameters to delete.
	//
	UrlErrorReasonCANNOT_REMOVE_URL_CUSTOM_PARAMETER_IN_ADD_OPERATION UrlErrorReason = "CANNOT_REMOVE_URL_CUSTOM_PARAMETER_IN_ADD_OPERATION"

	//
	// When the doReplace flag is set to true, individual parameters cannot be deleted.
	//
	UrlErrorReasonCANNOT_REMOVE_URL_CUSTOM_PARAMETER_DURING_FULL_REPLACEMENT UrlErrorReason = "CANNOT_REMOVE_URL_CUSTOM_PARAMETER_DURING_FULL_REPLACEMENT"

	//
	// For ADD operations and when the doReplace flag is set to true, custom parameter values
	// cannot be null.
	//
	UrlErrorReasonNULL_CUSTOM_PARAMETER_VALUE_DURING_ADD_OR_FULL_REPLACEMENT UrlErrorReason = "NULL_CUSTOM_PARAMETER_VALUE_DURING_ADD_OR_FULL_REPLACEMENT"

	//
	// An invalid character appears in the parameter key.
	//
	UrlErrorReasonINVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_KEY UrlErrorReason = "INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_KEY"

	//
	// An invalid character appears in the parameter value.
	//
	UrlErrorReasonINVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_VALUE UrlErrorReason = "INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_VALUE"

	//
	// The url custom parameter value fails url tag validation.
	//
	UrlErrorReasonINVALID_TAG_IN_URL_CUSTOM_PARAMETER_VALUE UrlErrorReason = "INVALID_TAG_IN_URL_CUSTOM_PARAMETER_VALUE"

	//
	// The custom parameter contains nested occurrences of the same conditional tag
	// (i.e. {ifmobile:{ifmobile:x}}).
	//
	UrlErrorReasonREDUNDANT_NESTED_URL_CUSTOM_PARAMETER_TAG UrlErrorReason = "REDUNDANT_NESTED_URL_CUSTOM_PARAMETER_TAG"

	//
	// The protocol (http:// or https://) is missing.
	//
	UrlErrorReasonMISSING_PROTOCOL UrlErrorReason = "MISSING_PROTOCOL"

	//
	// The url is invalid.
	//
	UrlErrorReasonINVALID_URL UrlErrorReason = "INVALID_URL"

	//
	// Destination Url is deprecated.
	//
	UrlErrorReasonDESTINATION_URL_DEPRECATED UrlErrorReason = "DESTINATION_URL_DEPRECATED"

	//
	// The url contains invalid tag.
	//
	UrlErrorReasonINVALID_TAG_IN_URL UrlErrorReason = "INVALID_TAG_IN_URL"

	//
	// The url must contain at least one tag (e.g. {lpurl}),
	// This applies only to urls associated with website ads or product ads.
	//
	UrlErrorReasonMISSING_URL_TAG UrlErrorReason = "MISSING_URL_TAG"

	UrlErrorReasonDUPLICATE_URL_ID UrlErrorReason = "DUPLICATE_URL_ID"

	UrlErrorReasonINVALID_URL_ID UrlErrorReason = "INVALID_URL_ID"

	UrlErrorReasonURL_ERROR UrlErrorReason = "URL_ERROR"
)

//
// Membership status of the user list.
//
type CriterionUserListMembershipStatus string

const (

	//
	// Open status - list is accruing members and can be targeted to.
	//
	CriterionUserListMembershipStatusOPEN CriterionUserListMembershipStatus = "OPEN"

	//
	// Closed status - No new members being added. Can not be used for targeting a new campaign.
	// Existing campaigns can still work as long as the list is not removed as a targeting criteria.
	//
	CriterionUserListMembershipStatusCLOSED CriterionUserListMembershipStatus = "CLOSED"
)

type Get struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 get"`

	//
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Selector *Selector `xml:"selector,omitempty"`
}

type GetResponse struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 getResponse"`

	Rval *CampaignExtensionSettingPage `xml:"rval,omitempty"`
}

type Mutate struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 mutate"`

	//
	// <span class="constraint ContentsNotNull">This field must not contain {@code null} elements.</span>
	// <span class="constraint DistinctIds">Elements in this field must have distinct IDs for following {@link Operator}s : ADD, SET, REMOVE.</span>
	// <span class="constraint NotEmpty">This field must contain at least one element.</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Operations []*CampaignExtensionSettingOperation `xml:"operations,omitempty"`
}

type MutateResponse struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 mutateResponse"`

	Rval *CampaignExtensionSettingReturnValue `xml:"rval,omitempty"`
}

type Query struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 query"`

	//
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Query string `xml:"query,omitempty"`
}

type QueryResponse struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 queryResponse"`

	Rval *CampaignExtensionSettingPage `xml:"rval,omitempty"`
}

type ApiError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ApiError"`

	//
	// The OGNL field path to identify cause of error.
	//
	FieldPath string `xml:"fieldPath,omitempty"`

	//
	// A parsed copy of the field path. For example, the field path "operations[1].operand"
	// corresponds to this list: {FieldPathElement(field = "operations", index = 1),
	// FieldPathElement(field = "operand", index = null)}.
	//
	FieldPathElements []*FieldPathElement `xml:"fieldPathElements,omitempty"`

	//
	// The data that caused the error.
	//
	Trigger string `xml:"trigger,omitempty"`

	//
	// A simple string representation of the error and reason.
	//
	ErrorString string `xml:"errorString,omitempty"`

	//
	// Indicates that this instance is a subtype of ApiError.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	ApiErrorType string `xml:"ApiError.Type,omitempty"`
}

type ApiException struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ApiException"`

	*ApplicationException

	//
	// List of errors.
	//
	Errors []*ApiError `xml:"errors,omitempty"`
}

type AppFeedItem struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AppFeedItem"`

	*ExtensionFeedItem

	//
	// The application store that the target application belongs to.
	//
	AppStore *AppFeedItemAppStore `xml:"appStore,omitempty"`

	//
	// The store-specific ID for the target application.
	// <span class="constraint StringLength">This string must not be empty, (trimmed).</span>
	//
	AppId string `xml:"appId,omitempty"`

	//
	// The visible text displayed when the link is rendered in an ad.
	// <span class="constraint StringLength">The length of this string should be between 1 and 25, inclusive, (trimmed).</span>
	//
	AppLinkText string `xml:"appLinkText,omitempty"`

	//
	// The destination URL of the in-app link.
	// <span class="constraint StringLength">The length of this string should be between 0 and 2076, inclusive, (trimmed).</span>
	//
	AppUrl string `xml:"appUrl,omitempty"`

	//
	// A list of possible final URLs after all cross domain redirects.
	//
	AppFinalUrls *UrlList `xml:"appFinalUrls,omitempty"`

	//
	// A list of possible final mobile URLs after all cross domain redirects.
	//
	AppFinalMobileUrls *UrlList `xml:"appFinalMobileUrls,omitempty"`

	//
	// URL template for constructing a tracking URL. To clear this field, set its value to the empty
	// string.
	//
	AppTrackingUrlTemplate string `xml:"appTrackingUrlTemplate,omitempty"`

	//
	// A list of mappings to be used for substituting URL custom parameter tags in the
	// trackingUrlTemplate, finalUrls, and/or finalMobileUrls.
	//
	AppUrlCustomParameters *CustomParameters `xml:"appUrlCustomParameters,omitempty"`
}

type ApplicationException struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ApplicationException"`

	//
	// Error message.
	//
	Message string `xml:"message,omitempty"`

	//
	// Indicates that this instance is a subtype of ApplicationException.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	ApplicationExceptionType string `xml:"ApplicationException.Type,omitempty"`
}

type AuthenticationError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AuthenticationError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *AuthenticationErrorReason `xml:"reason,omitempty"`
}

type AuthorizationError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AuthorizationError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *AuthorizationErrorReason `xml:"reason,omitempty"`
}

type CallConversionType struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CallConversionType"`

	//
	// The ID of an AdCallMetricsConversion object. This object contains the phoneCallDuration field
	// which is the minimum duration (in seconds) of a call to be considered a conversion.
	//
	ConversionTypeId int64 `xml:"conversionTypeId,omitempty"`
}

type CallFeedItem struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CallFeedItem"`

	*ExtensionFeedItem

	//
	// The advertiser's phone number to append to the ad.
	// <span class="constraint StringLength">This string must not be empty, (trimmed).</span>
	//
	CallPhoneNumber string `xml:"callPhoneNumber,omitempty"`

	//
	// Uppercase two-letter country code of the advertiser's phone number.
	// <span class="constraint StringLength">This string must not be empty, (trimmed).</span>
	//
	CallCountryCode string `xml:"callCountryCode,omitempty"`

	//
	// Indicates whether call tracking is enabled. By default, call tracking is not enabled.
	//
	CallTracking bool `xml:"callTracking,omitempty"`

	//
	// Call conversion type. To clear this field, set a CallConversionType with a value of null in its
	// conversionTypeId field. This value should not be set if
	// {@linkPlain CallFeedItem#disableCallConversionTracking} is true.
	//
	CallConversionType *CallConversionType `xml:"callConversionType,omitempty"`

	//
	// If set, disable call conversion tracking. {@linkPlain CallFeedItem#callConversionType} should
	// not be set if this value is true.
	//
	DisableCallConversionTracking bool `xml:"disableCallConversionTracking,omitempty"`
}

type CalloutFeedItem struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CalloutFeedItem"`

	*ExtensionFeedItem

	//
	// The callout text.
	// <span class="constraint StringLength">The length of this string should be between 1 and 25, inclusive, (trimmed).</span>
	//
	CalloutText string `xml:"calloutText,omitempty"`
}

type CampaignExtensionSetting struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CampaignExtensionSetting"`

	//
	// The id of the campaign for the feed items being added or modified.
	// <span class="constraint Selectable">This field can be selected using the value "CampaignId".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	CampaignId int64 `xml:"campaignId,omitempty"`

	//
	// The extension type the extension setting applies to.
	// <span class="constraint Selectable">This field can be selected using the value "ExtensionType".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	ExtensionType *FeedType `xml:"extensionType,omitempty"`

	//
	// The extension setting specifying which extensions to serve for the specified campaign. If
	// extensionSetting is empty (i.e. has an empty list of feed items and null platformRestrictions),
	// extensions are disabled for the specified extensionType.
	//
	ExtensionSetting *ExtensionSetting `xml:"extensionSetting,omitempty"`
}

type CampaignExtensionSettingOperation struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CampaignExtensionSettingOperation"`

	*Operation

	//
	// CampaignExtensionSetting to operate on.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Operand *CampaignExtensionSetting `xml:"operand,omitempty"`
}

type CampaignExtensionSettingPage struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CampaignExtensionSettingPage"`

	*Page

	//
	// The result entries in this page.
	//
	Entries []*CampaignExtensionSetting `xml:"entries,omitempty"`
}

type CampaignExtensionSettingReturnValue struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CampaignExtensionSettingReturnValue"`

	*ListReturnValue

	//
	// The resulting CampaignExtensionSettings.
	//
	Value []*CampaignExtensionSetting `xml:"value,omitempty"`

	PartialFailureErrors []*ApiError `xml:"partialFailureErrors,omitempty"`
}

type ClientTermsError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ClientTermsError"`

	*ApiError

	Reason *ClientTermsErrorReason `xml:"reason,omitempty"`
}

type CollectionSizeError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CollectionSizeError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *CollectionSizeErrorReason `xml:"reason,omitempty"`
}

type ComparableValue struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ComparableValue"`

	//
	// Indicates that this instance is a subtype of ComparableValue.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	ComparableValueType string `xml:"ComparableValue.Type,omitempty"`
}

type Criterion struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Criterion"`

	//
	// ID of this criterion.
	//
	Id int64 `xml:"id,omitempty"`

	//
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	Type_ *CriterionType `xml:"type,omitempty"`

	//
	// Indicates that this instance is a subtype of Criterion.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	CriterionType string `xml:"Criterion.Type,omitempty"`
}

type CriterionError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CriterionError"`

	*ApiError

	Reason *CriterionErrorReason `xml:"reason,omitempty"`
}

type CustomParameter struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CustomParameter"`

	//
	// The parameter key to be mapped.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	// <span class="constraint StringLength">The length of this string should be between 1 and 16, inclusive, in UTF-8 bytes, (trimmed).</span>
	//
	Key string `xml:"key,omitempty"`

	//
	// The value this parameter should be mapped to. It should be null if isRemove is true.
	// <span class="constraint StringLength">The length of this string should be between 0 and 200, inclusive, in UTF-8 bytes, (trimmed).</span>
	//
	Value string `xml:"value,omitempty"`

	//
	// On SET operation, indicates that the parameter should be removed from the existing parameters.
	// If set to true, the value field must be null.
	//
	IsRemove bool `xml:"isRemove,omitempty"`
}

type CustomParameters struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CustomParameters"`

	//
	// The list of custom parameters.
	//
	// <p>On update, all parameters can be cleared by providing an empty or null list and setting
	// doReplace to true.
	//
	Parameters []*CustomParameter `xml:"parameters,omitempty"`

	//
	// On SET operation, indicates that the current parameters should be cleared and replaced
	// with these parameters.
	//
	DoReplace bool `xml:"doReplace,omitempty"`
}

type DatabaseError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 DatabaseError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *DatabaseErrorReason `xml:"reason,omitempty"`
}

type DateError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 DateError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *DateErrorReason `xml:"reason,omitempty"`
}

type DateRange struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 DateRange"`

	//
	// the lower bound of this date range, inclusive.
	//
	Min string `xml:"min,omitempty"`

	//
	// the upper bound of this date range, inclusive.
	//
	Max string `xml:"max,omitempty"`
}

type DisapprovalReason struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 DisapprovalReason"`

	//
	// Short description of the disapproval reason, localized for the specific advertiser.
	//
	ShortName string `xml:"shortName,omitempty"`
}

type DistinctError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 DistinctError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *DistinctErrorReason `xml:"reason,omitempty"`
}

type DoubleValue struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 DoubleValue"`

	*NumberValue

	//
	// the underlying double value.
	//
	Number float64 `xml:"number,omitempty"`
}

type EntityAccessDenied struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 EntityAccessDenied"`

	*ApiError

	//
	// Reason for this error.
	//
	Reason *EntityAccessDeniedReason `xml:"reason,omitempty"`
}

type EntityCountLimitExceeded struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 EntityCountLimitExceeded"`

	*ApiError

	//
	// Specifies which level's limit was exceeded.
	//
	Reason *EntityCountLimitExceededReason `xml:"reason,omitempty"`

	//
	// Id of the entity whose limit was exceeded.
	//
	EnclosingId string `xml:"enclosingId,omitempty"`

	//
	// The limit which was exceeded.
	//
	Limit int32 `xml:"limit,omitempty"`

	//
	// The account limit type which was exceeded.
	//
	AccountLimitType string `xml:"accountLimitType,omitempty"`

	//
	// The count of existing entities.
	//
	ExistingCount int32 `xml:"existingCount,omitempty"`
}

type EntityNotFound struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 EntityNotFound"`

	*ApiError

	//
	// Reason for this error.
	//
	Reason *EntityNotFoundReason `xml:"reason,omitempty"`
}

type ExtensionFeedItem struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ExtensionFeedItem"`

	//
	// Id of this feed item's feed.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	FeedId int64 `xml:"feedId,omitempty"`

	//
	// Id of the feed item.
	//
	FeedItemId int64 `xml:"feedItemId,omitempty"`

	//
	// Status of the feed item.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	Status *FeedItemStatus `xml:"status,omitempty"`

	//
	// The type of the feed containing this extension feed item data. The field will be set by a
	// subclass with a defined feed type.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	FeedType *FeedType `xml:"feedType,omitempty"`

	//
	// Start time in which this feed item is effective and can begin serving. The time zone
	// of startTime must either match the time zone of the account or be unspecified where
	// the time zone defaults to the account time zone.
	// This field may be null to indicate no start time restriction.
	// The special value "00000101 000000" may be used to clear an existing start time.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API for the following {@link Operator}s: REMOVE.</span>
	//
	StartTime string `xml:"startTime,omitempty"`

	//
	// End time in which this feed item is no longer effective and will stop serving. The time zone
	// of endTime must either match the time zone of the account or be unspecified where
	// the time zone defaults to the account time zone.
	// This field may be null to indicate no end time restriction.
	// The special value "00000101 000000" may be used to clear an existing end time.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API for the following {@link Operator}s: REMOVE.</span>
	//
	EndTime string `xml:"endTime,omitempty"`

	//
	// Device preference for the feed item.
	//
	DevicePreference *FeedItemDevicePreference `xml:"devicePreference,omitempty"`

	//
	// FeedItemScheduling specifying times for when the feed item may serve.
	// On retrieval or creation of the feed item, if the field is left null,
	// no feed item scheduling restrictions are placed on the feed item.
	// On update, if the field is left unspecified, the previous feedItemScheduling state will
	// not be changed.
	// On update, if the field is set with a FeedItemScheduling with an empty feedItemSchedules
	// list, the scheduling will be cleared of all FeedItemSchedules indicating the feed item
	// has no scheduling restrictions.
	//
	Scheduling *FeedItemScheduling `xml:"scheduling,omitempty"`

	//
	// Campaign targeting specifying what campaign this extension can serve with.
	// On update, if the field is left unspecified, the previous campaign targeting state
	// will not be changed.
	// On update, if the field is set with an empty FeedItemCampaignTargeting, the
	// campaign targeting will be cleared.
	// Note: If adGroupTargeting and campaignTargeting are set (either in the request or pre-existing
	// from a previous request), the targeted campaign must match the targeted adgroup's campaign.
	// If only adGroupTargeting is specified and there is no campaignTargeting, then a
	// campaignTargeting will be set to the targeted adgroup's campaign.
	//
	CampaignTargeting *FeedItemCampaignTargeting `xml:"campaignTargeting,omitempty"`

	//
	// Adgroup targeting specifying what adgroup this extension can serve with.
	// On update, if the field is left unspecified, the previous adgroup targeting state
	// will not be changed.
	// On update, if the field is set with an empty FeedItemAdGroupTargeting, the
	// adgroup targeting will be cleared.
	// Note: If adGroupTargeting and campaignTargeting are set (either in the request or pre-existing
	// from a previous request), the targeted campaign must match the targeted adgroup's campaign.
	// If only adGroupTargeting is specified and there is no campaignTargeting, then a
	// campaignTargeting will be set to the targeted adgroup's campaign.
	//
	AdGroupTargeting *FeedItemAdGroupTargeting `xml:"adGroupTargeting,omitempty"`

	//
	// Keyword targeting specifies what keyword this extension can serve with.
	// On update, if the field is left unspecified, the previous keyword targeting state
	// will not be changed.
	// On update, if the field is set with a Keyword and without Keyword.text set keyword targeting
	// will be cleared.
	//
	KeywordTargeting *Keyword `xml:"keywordTargeting,omitempty"`

	//
	// Geo targeting specifies what locations this extension can serve with.
	// On update, if the field is left unspecified, the previous geo targeting state will not
	// be changed.
	// On update, if the field is set with a null value for criterionId, the geo targeting will be
	// cleared.
	//
	GeoTargeting *Location `xml:"geoTargeting,omitempty"`

	//
	// Geo targeting restriction specifies the type of location that can be used for targeting.
	// On update, if the field is left unspecified, the previous geo targeting restriction state
	// will not be changed.
	// On update, if the field is set with a null GeoRestriction enum, the geo targeting restriction
	// will be cleared.
	//
	GeoTargetingRestriction *FeedItemGeoRestriction `xml:"geoTargetingRestriction,omitempty"`

	//
	// List of details about a feed item's validation and approval state.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	PolicyData []*FeedItemPolicyData `xml:"policyData,omitempty"`

	//
	// Indicates that this instance is a subtype of ExtensionFeedItem.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	ExtensionFeedItemType string `xml:"ExtensionFeedItem.Type,omitempty"`
}

type ExtensionSetting struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ExtensionSetting"`

	//
	// The list of feed items to add or modify.
	// <span class="constraint Selectable">This field can be selected using the value "Extensions".</span>
	//
	Extensions []*ExtensionFeedItem `xml:"extensions,omitempty"`

	//
	// Any platform (desktop, mobile) restrictions for feed items being served. If set to DESKTOP or
	// MOBILE, only those feed items with the appropriate device preference or no device preference
	// will serve.
	// <span class="constraint Selectable">This field can be selected using the value "PlatformRestrictions".</span>
	//
	PlatformRestrictions *ExtensionSettingPlatform `xml:"platformRestrictions,omitempty"`
}

type ExtensionSettingError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ExtensionSettingError"`

	*ApiError

	//
	// Error reason.
	//
	Reason *ExtensionSettingErrorReason `xml:"reason,omitempty"`
}

type FeedItemAdGroupTargeting struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 FeedItemAdGroupTargeting"`

	//
	// The ID of the adgroup to target.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API for the following {@link Operator}s: REMOVE.</span>
	//
	TargetingAdGroupId int64 `xml:"TargetingAdGroupId,omitempty"`
}

type FeedItemAttributeError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 FeedItemAttributeError"`

	//
	// Contains the set of feed attribute ids whose attributes together triggered the error.
	// Null or empty field means error code does not apply to a specific set of attributes.
	//
	FeedAttributeIds []int64 `xml:"feedAttributeIds,omitempty"`

	//
	// Validation error code. See the
	// <a href="/adwords/api/docs/appendix/feed-errors">list of error codes</a>.
	//
	ValidationErrorCode int32 `xml:"validationErrorCode,omitempty"`

	//
	// Extra information about the error, including related field IDs.
	//
	ErrorInformation string `xml:"errorInformation,omitempty"`
}

type FeedItemCampaignTargeting struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 FeedItemCampaignTargeting"`

	//
	// The ID of the campaign to target.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API for the following {@link Operator}s: REMOVE.</span>
	//
	TargetingCampaignId int64 `xml:"TargetingCampaignId,omitempty"`
}

type FeedItemDevicePreference struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 FeedItemDevicePreference"`

	//
	// CriterionId of the type of device the feed item is preferred to serve on.
	// Only CriterionId 30001 (mobile devices) is currently supported.
	// If unspecified, the device preference will be cleared indicating that the feed item
	// is not preferred for any device type.
	//
	DevicePreference int64 `xml:"devicePreference,omitempty"`
}

type FeedItemGeoRestriction struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 FeedItemGeoRestriction"`

	//
	// The geo targeting restriction of a feed item.  If null then the geo restriction is cleared.
	//
	GeoRestriction *GeoRestriction `xml:"geoRestriction,omitempty"`
}

type FeedItemPolicyData struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 FeedItemPolicyData"`

	*PolicyData

	//
	// Mapped placeholder type used in validation/approvals checks.
	//
	PlaceholderType int32 `xml:"placeholderType,omitempty"`

	//
	// Id of FeedMapping used in validation/approvals checks.
	//
	FeedMappingId int64 `xml:"feedMappingId,omitempty"`

	//
	// Validation status of feed item for a particular feed mapping.
	//
	ValidationStatus *FeedItemValidationStatus `xml:"validationStatus,omitempty"`

	//
	// Feed item approval status.
	//
	ApprovalStatus *FeedItemApprovalStatus `xml:"approvalStatus,omitempty"`

	//
	// List of error codes specifying what errors were found during validation.
	//
	ValidationErrors []*FeedItemAttributeError `xml:"validationErrors,omitempty"`

	//
	// Feed item quality evaluation approval status for a particular feed mapping.
	//
	QualityApprovalStatus *FeedItemQualityApprovalStatus `xml:"qualityApprovalStatus,omitempty"`

	//
	// Feed item quality evaluation disapproval reasons.
	//
	QualityDisapprovalReasons []*FeedItemQualityDisapprovalReasons `xml:"qualityDisapprovalReasons,omitempty"`
}

type FeedItemSchedule struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 FeedItemSchedule"`

	//
	// Day of the week the schedule applies to.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	DayOfWeek *DayOfWeek `xml:"dayOfWeek,omitempty"`

	//
	// Starting hour in 24 hour time.
	// <span class="constraint InRange">This field must be between 0 and 23, inclusive.</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	StartHour int32 `xml:"startHour,omitempty"`

	//
	// Interval starts these minutes after the starting hour.
	// The value can be 0, 15, 30, and 45.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	StartMinute *MinuteOfHour `xml:"startMinute,omitempty"`

	//
	// Ending hour in 24 hour time; <code>24</code> signifies
	// end of the day and subsequently endMinute must be 0.
	// <span class="constraint InRange">This field must be between 0 and 24, inclusive.</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	EndHour int32 `xml:"endHour,omitempty"`

	//
	// Interval ends these minutes after the ending hour.
	// The value can be 0, 15, 30, and 45.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	EndMinute *MinuteOfHour `xml:"endMinute,omitempty"`
}

type FeedItemScheduling struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 FeedItemScheduling"`

	//
	// List of non-overlapping feed item schedules indicating when the feed item may serve.
	// There can be a maximum of 6 FeedItemSchedules per day.
	// If empty, the scheduling will be cleared of all FeedItemSchedules indicating the feed item
	// has no scheduling restrictions.
	//
	FeedItemSchedules []*FeedItemSchedule `xml:"feedItemSchedules,omitempty"`
}

type FieldPathElement struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 FieldPathElement"`

	//
	// The name of a field in lower camelcase. (e.g. "biddingStrategy")
	//
	Field string `xml:"field,omitempty"`

	//
	// For list fields, this is a 0-indexed position in the list. Null for non-list fields.
	//
	Index int32 `xml:"index,omitempty"`
}

type IdError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 IdError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *IdErrorReason `xml:"reason,omitempty"`
}

type InternalApiError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 InternalApiError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *InternalApiErrorReason `xml:"reason,omitempty"`
}

type Keyword struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Keyword"`

	*Criterion

	//
	// Text of this keyword (at most 80 characters and ten words).
	// <span class="constraint MatchesRegex">Keyword text must not contain NUL (code point 0x0) characters. This is checked by the regular expression '[^\x00]*'.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	Text string `xml:"text,omitempty"`

	//
	// Match type of this keyword.
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	MatchType *KeywordMatchType `xml:"matchType,omitempty"`
}

type ListReturnValue struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ListReturnValue"`

	//
	// Indicates that this instance is a subtype of ListReturnValue.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	ListReturnValueType string `xml:"ListReturnValue.Type,omitempty"`
}

type Location struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Location"`

	*Criterion

	//
	// Name of the location criterion. <b> Note:</b> This field is filterable only in
	// LocationCriterionService. If used as a filter, a location name cannot be greater than 300
	// characters.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	LocationName string `xml:"locationName,omitempty"`

	//
	// Display type of the location criterion.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	DisplayType string `xml:"displayType,omitempty"`

	//
	// The targeting status of the location criterion.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	TargetingStatus *LocationTargetingStatus `xml:"targetingStatus,omitempty"`

	//
	// Ordered list of parents of the location criterion.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	ParentLocations []*Location `xml:"parentLocations,omitempty"`
}

type LongValue struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 LongValue"`

	*NumberValue

	//
	// the underlying long value.
	//
	Number int64 `xml:"number,omitempty"`
}

type MessageFeedItem struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 MessageFeedItem"`

	*ExtensionFeedItem

	//
	// The business name of the message. This will be prepended to the message text.
	// <span class="constraint StringLength">The length of this string should be between 1 and 25, inclusive, (trimmed).</span>
	//
	MessageBusinessName string `xml:"messageBusinessName,omitempty"`

	//
	// Uppercase two-letter country code of the advertiser's phone number to message.
	// <span class="constraint StringLength">This string must not be empty, (trimmed).</span>
	//
	MessageCountryCode string `xml:"messageCountryCode,omitempty"`

	//
	// The advertiser's phone number the message will be sent to.
	// <span class="constraint StringLength">This string must not be empty, (trimmed).</span>
	//
	MessagePhoneNumber string `xml:"messagePhoneNumber,omitempty"`

	//
	// The text to show in the ad.
	// <span class="constraint StringLength">The length of this string should be between 5 and 35, inclusive, (trimmed).</span>
	//
	MessageExtensionText string `xml:"messageExtensionText,omitempty"`

	//
	// The message text populated in the messaging app.
	// <span class="constraint StringLength">The length of this string should be between 10 and 100, inclusive, (trimmed).</span>
	//
	MessageText string `xml:"messageText,omitempty"`
}

type MobileAppCategory struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 MobileAppCategory"`

	*Criterion

	//
	// ID of this mobile app category. A complete list of the available mobile app categories is
	// available <a href="/adwords/api/docs/appendix/mobileappcategories">here</a>.
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	MobileAppCategoryId int32 `xml:"mobileAppCategoryId,omitempty"`

	//
	// Name of this mobile app category.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	DisplayName string `xml:"displayName,omitempty"`
}

type MobileApplication struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 MobileApplication"`

	*Criterion

	//
	// A string that uniquely identifies a mobile application to AdWords API. The format of this
	// string is "<code>{platform}-{platform_native_id}</code>", where <code>platform</code> is "1"
	// for iOS apps and "2" for Android apps, and where <code>platform_native_id</code> is the mobile
	// application identifier native to the corresponding platform.
	// For iOS, this native identifier is the 9 digit string that appears at the end of an App Store
	// URL (e.g., "476943146" for "Flood-It! 2" whose App Store link is
	// http://itunes.apple.com/us/app/flood-it!-2/id476943146).
	// For Android, this native identifier is the application's package name (e.g.,
	// "com.labpixies.colordrips" for "Color Drips" given Google Play link
	// https://play.google.com/store/apps/details?id=com.labpixies.colordrips).
	// A well formed app id for AdWords API would thus be "1-476943146" for iOS and
	// "2-com.labpixies.colordrips" for Android.
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	AppId string `xml:"appId,omitempty"`

	//
	// Title of this mobile application.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	DisplayName string `xml:"displayName,omitempty"`
}

type Money struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Money"`

	*ComparableValue

	//
	// Amount in micros. One million is equivalent to one unit.
	//
	MicroAmount int64 `xml:"microAmount,omitempty"`
}

type MoneyWithCurrency struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 MoneyWithCurrency"`

	*ComparableValue

	//
	// The amount of money.
	// <span class="constraint InRange">This field must be greater than or equal to 0.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : SET, ADD.</span>
	//
	Money *Money `xml:"money,omitempty"`

	//
	// Currency code.
	// <span class="constraint StringLength">The length of this string should be between 3 and 3, inclusive, (trimmed).</span>
	//
	CurrencyCode string `xml:"currencyCode,omitempty"`
}

type NewEntityCreationError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 NewEntityCreationError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *NewEntityCreationErrorReason `xml:"reason,omitempty"`
}

type NotEmptyError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 NotEmptyError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *NotEmptyErrorReason `xml:"reason,omitempty"`
}

type NullError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 NullError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *NullErrorReason `xml:"reason,omitempty"`
}

type NumberValue struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 NumberValue"`

	*ComparableValue
}

type Operation struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Operation"`

	//
	// Operator.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Operator *Operator `xml:"operator,omitempty"`

	//
	// Indicates that this instance is a subtype of Operation.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	OperationType string `xml:"Operation.Type,omitempty"`
}

type OperationAccessDenied struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 OperationAccessDenied"`

	*ApiError

	Reason *OperationAccessDeniedReason `xml:"reason,omitempty"`
}

type OperatorError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 OperatorError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *OperatorErrorReason `xml:"reason,omitempty"`
}

type OrderBy struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 OrderBy"`

	//
	// The field to sort the results on.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Field string `xml:"field,omitempty"`

	//
	// The order to sort the results on. The default sort order is {@link SortOrder#ASCENDING}.
	//
	SortOrder *SortOrder `xml:"sortOrder,omitempty"`
}

type Page struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Page"`

	//
	// Total number of entries in the result that this page is a part of.
	//
	TotalNumEntries int32 `xml:"totalNumEntries,omitempty"`

	//
	// Indicates that this instance is a subtype of Page.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	PageType string `xml:"Page.Type,omitempty"`
}

type Paging struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Paging"`

	//
	// Index of the first result to return in this page.
	// <span class="constraint InRange">This field must be greater than or equal to 0.</span>
	//
	StartIndex int32 `xml:"startIndex,omitempty"`

	//
	// Maximum number of results to return in this page. Set this to a reasonable value to limit
	// the number of results returned per page.
	// <span class="constraint InRange">This field must be greater than or equal to 0.</span>
	//
	NumberResults int32 `xml:"numberResults,omitempty"`
}

type Placement struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Placement"`

	*Criterion

	//
	// Url of the placement.
	//
	// <p>For example, "http://www.domain.com".
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	Url string `xml:"url,omitempty"`
}

type PolicyData struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 PolicyData"`

	//
	// List of disapproval reasons attached to the entity.
	//
	DisapprovalReasons []*DisapprovalReason `xml:"disapprovalReasons,omitempty"`

	//
	// Indicates that this instance is a subtype of PolicyData.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	PolicyDataType string `xml:"PolicyData.Type,omitempty"`
}

type Predicate struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Predicate"`

	//
	// The field by which to filter the returned data. Possible values are marked Filterable on
	// the entity's reference page. For example, for predicates for the
	// CampaignService {@link Selector selector}, refer to the filterable fields from the
	// {@link Campaign} reference page.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Field string `xml:"field,omitempty"`

	//
	// The operator to use for filtering the data returned.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Operator *PredicateOperator `xml:"operator,omitempty"`

	//
	// The values by which to filter the field. The {@link Operator#CONTAINS_ALL},
	// {@link Operator#CONTAINS_ANY}, {@link Operator#CONTAINS_NONE}, {@link Operator#IN}
	// and {@link Operator#NOT_IN} take multiple values. All others take a single value.
	// <span class="constraint ContentsNotNull">This field must not contain {@code null} elements.</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Values []string `xml:"values,omitempty"`
}

type PriceFeedItem struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 PriceFeedItem"`

	*ExtensionFeedItem

	//
	// Price extension type of this extension. Required.
	//
	PriceExtensionType *PriceExtensionType `xml:"priceExtensionType,omitempty"`

	//
	// Price qualifier for all rows of this price extension.
	//
	PriceQualifier *PriceExtensionPriceQualifier `xml:"priceQualifier,omitempty"`

	//
	// Tracking URL template for all rows of this price extension. To clear this field, set its value
	// to an empty string.
	//
	TrackingUrlTemplate string `xml:"trackingUrlTemplate,omitempty"`

	//
	// The language the advertiser is using for this price extension. Required.
	// Supported language codes:
	// <ul>
	// <li>de</li>
	// <li>en</li>
	// <li>es</li>
	// <li>es-419</li>
	// <li>fr</li>
	// <li>it</li>
	// <li>ja</li>
	// <li>nl</li>
	// <li>pl</li>
	// <li>pt-BR</li>
	// <li>pt-PT</li>
	// <li>ru</li>
	// <li>sv</li>
	// </ul>
	//
	Language string `xml:"language,omitempty"`

	//
	// The rows in this price extension. Minimum number of items allowed is 3 and the maximum number
	// is 8.
	//
	TableRows []*PriceTableRow `xml:"tableRows,omitempty"`
}

type PriceTableRow struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 PriceTableRow"`

	//
	// Header text of this row. Required.
	// <span class="constraint StringLength">The length of this string should be between 1 and 25, inclusive, (trimmed).</span>
	//
	Header string `xml:"header,omitempty"`

	//
	// Description text of this row. Required.
	// <span class="constraint StringLength">The length of this string should be between 1 and 25, inclusive, (trimmed).</span>
	//
	Description string `xml:"description,omitempty"`

	//
	// A list of possible final URLs after all cross domain redirects. Required.
	//
	FinalUrls *UrlList `xml:"finalUrls,omitempty"`

	//
	// A list of possible final mobile URLs after all cross domain redirects.
	//
	FinalMobileUrls *UrlList `xml:"finalMobileUrls,omitempty"`

	//
	// Price value of this row. Required.
	//
	Price *MoneyWithCurrency `xml:"price,omitempty"`

	//
	// Price unit for this row.
	//
	PriceUnit *PriceExtensionPriceUnit `xml:"priceUnit,omitempty"`
}

type PromotionFeedItem struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 PromotionFeedItem"`

	*ExtensionFeedItem

	//
	// Promotion target. Required.
	//
	PromotionTarget string `xml:"promotionTarget,omitempty"`

	//
	// Discount modifier. Optional.
	//
	DiscountModifier *PromotionExtensionDiscountModifier `xml:"discountModifier,omitempty"`

	//
	// Percent off in micros. One million is equivalent to one percent.
	// Either percentOff or moneyAmountOff is required.
	// Cannot set both percentOff and moneyAmountOff.
	//
	PercentOff int64 `xml:"percentOff,omitempty"`

	//
	// Money amount off. Either percentOff or moneyAmountOff is required.
	// Cannot set both moneyAmountOff and percentOff.
	//
	MoneyAmountOff *MoneyWithCurrency `xml:"moneyAmountOff,omitempty"`

	//
	// Promotion code. Optional.
	// Cannot set both promotionCode and ordersOverAmount.
	//
	PromotionCode string `xml:"promotionCode,omitempty"`

	//
	// Orders over amount. Optional.
	// Cannot set both ordersOverAmount and promotionCode.
	//
	OrdersOverAmount *MoneyWithCurrency `xml:"ordersOverAmount,omitempty"`

	//
	// Promotion start. Optional.
	// The time part must be set to midnight.
	// The special value "00000101 000000" may be used to clear an existing value.
	//
	PromotionStart string `xml:"promotionStart,omitempty"`

	//
	// Promotion end. Optional
	// The time part must be set to midnight.
	// The special value "00000101 000000" may be used to clear an existing value.
	//
	PromotionEnd string `xml:"promotionEnd,omitempty"`

	//
	// Occasion of the promotion. Optional.
	//
	Occasion *PromotionExtensionOccasion `xml:"occasion,omitempty"`

	//
	// Final URLs. Required.
	//
	FinalUrls *UrlList `xml:"finalUrls,omitempty"`

	//
	// Final mobile URLs. Optional.
	//
	FinalMobileUrls *UrlList `xml:"finalMobileUrls,omitempty"`

	//
	// Tracking URL template. Optional.
	//
	TrackingUrlTemplate string `xml:"trackingUrlTemplate,omitempty"`

	//
	// A list of mappings to be used for substituting URL custom parameter tags in the
	// trackingUrlTemplate, finalUrls, and/or finalMobileUrls.
	//
	PromotionUrlCustomParameters *CustomParameters `xml:"promotionUrlCustomParameters,omitempty"`

	//
	// The language of the promotion. Optional.
	// The default language is English.
	// Represented as a BCP 47 language tag.
	// Supported language codes:
	// <ul>
	// <li>ar</li>
	// <li>bg</li>
	// <li>ca</li>
	// <li>zh-HK</li>
	// <li>zh-CN</li>
	// <li>zh-TW</li>
	// <li>hr</li>
	// <li>cs</li>
	// <li>da</li>
	// <li>nl</li>
	// <li>en-AU</li>
	// <li>en</li>
	// <li>en-GB</li>
	// <li>en-US</li>
	// <li>et</li>
	// <li>fil</li>
	// <li>fi</li>
	// <li>fr</li>
	// <li>de</li>
	// <li>el</li>
	// <li>iw</li>
	// <li>hi</li>
	// <li>hu</li>
	// <li>id</li>
	// <li>it</li>
	// <li>ja</li>
	// <li>ko</li>
	// <li>lv</li>
	// <li>lt</li>
	// <li>ms</li>
	// <li>no</li>
	// <li>pl</li>
	// <li>pt-BR</li>
	// <li>pt-PT</li>
	// <li>ro</li>
	// <li>ru</li>
	// <li>sr</li>
	// <li>sk</li>
	// <li>sl</li>
	// <li>es</li>
	// <li>es-419</li>
	// <li>sv</li>
	// <li>th</li>
	// <li>tr</li>
	// <li>uk</li>
	// <li>vi</li>
	// </ul>
	//
	Language string `xml:"language,omitempty"`
}

type QueryError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 QueryError"`

	*ApiError

	Reason *QueryErrorReason `xml:"reason,omitempty"`

	Message string `xml:"message,omitempty"`
}

type QuotaCheckError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 QuotaCheckError"`

	*ApiError

	Reason *QuotaCheckErrorReason `xml:"reason,omitempty"`
}

type RangeError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 RangeError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *RangeErrorReason `xml:"reason,omitempty"`
}

type RateExceededError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 RateExceededError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *RateExceededErrorReason `xml:"reason,omitempty"`

	//
	// Cause of the rate exceeded error.
	//
	RateName string `xml:"rateName,omitempty"`

	//
	// The scope of the rate (ACCOUNT/DEVELOPER).
	//
	RateScope string `xml:"rateScope,omitempty"`

	//
	// The amount of time (in seconds) the client should wait before retrying the request.
	//
	RetryAfterSeconds int32 `xml:"retryAfterSeconds,omitempty"`
}

type ReadOnlyError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ReadOnlyError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *ReadOnlyErrorReason `xml:"reason,omitempty"`
}

type RejectedError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 RejectedError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *RejectedErrorReason `xml:"reason,omitempty"`
}

type RequestError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 RequestError"`

	*ApiError

	Reason *RequestErrorReason `xml:"reason,omitempty"`
}

type RequiredError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 RequiredError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *RequiredErrorReason `xml:"reason,omitempty"`
}

type ReviewFeedItem struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ReviewFeedItem"`

	*ExtensionFeedItem

	//
	// An exact quote or paraphrase from a third-party source.
	// <span class="constraint StringLength">This string must not be empty, (trimmed).</span>
	//
	ReviewText string `xml:"reviewText,omitempty"`

	//
	// Name of the third-party publisher of the review.
	// <span class="constraint StringLength">This string must not be empty, (trimmed).</span>
	//
	ReviewSourceName string `xml:"reviewSourceName,omitempty"`

	//
	// Landing page of the third-party website of the review.
	// <span class="constraint StringLength">This string must not be empty, (trimmed).</span>
	//
	ReviewSourceUrl string `xml:"reviewSourceUrl,omitempty"`

	//
	// Indicates if your review is formatted as an exact quote. Use a value of false to indicate that
	// the review is paraphrased. If not set, the value is treated as false.
	//
	ReviewTextExactlyQuoted bool `xml:"reviewTextExactlyQuoted,omitempty"`
}

type Selector struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Selector"`

	//
	// List of fields to select.
	// <a href="/adwords/api/docs/appendix/selectorfields">Possible values</a>
	// are marked {@code Selectable} on the entity's reference page.
	// For example, for the {@code CampaignService} selector, refer to the
	// selectable fields from the {@link Campaign} reference page.
	// <span class="constraint ContentsDistinct">This field must contain distinct elements.</span>
	// <span class="constraint ContentsNotNull">This field must not contain {@code null} elements.</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Fields []string `xml:"fields,omitempty"`

	//
	// Specifies how an entity (eg. adgroup, campaign, criterion, ad) should be filtered.
	// <span class="constraint ContentsNotNull">This field must not contain {@code null} elements.</span>
	//
	Predicates []*Predicate `xml:"predicates,omitempty"`

	//
	// Range of dates for which you want to include data. If this value is omitted,
	// results for all dates are returned.
	// <p class="note"><b>Note:</b> This field is only used by the report download
	// service. For all other services, it is ignored.</p>
	// <span class="constraint DateRangeWithinRange">This range must be contained within the range [19700101, 20380101].</span>
	//
	DateRange *DateRange `xml:"dateRange,omitempty"`

	//
	// The fields on which you want to sort, and the sort order. The order in the list is
	// significant: The first element in the list indicates the primary sort order, the next
	// specifies the secondary sort order and so on.
	//
	Ordering []*OrderBy `xml:"ordering,omitempty"`

	//
	// Pagination information.
	//
	Paging *Paging `xml:"paging,omitempty"`
}

type SelectorError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 SelectorError"`

	*ApiError

	//
	// The error reason represented by enum.
	//
	Reason *SelectorErrorReason `xml:"reason,omitempty"`
}

type SitelinkFeedItem struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 SitelinkFeedItem"`

	*ExtensionFeedItem

	//
	// URL display text for the sitelink.
	// <span class="constraint StringLength">The length of this string should be between 1 and 25, inclusive, (trimmed).</span>
	//
	SitelinkText string `xml:"sitelinkText,omitempty"`

	//
	// Destination URL for the sitelink.
	// <span class="constraint StringLength">The length of this string should be between 0 and 2076, inclusive, (trimmed).</span>
	//
	SitelinkUrl string `xml:"sitelinkUrl,omitempty"`

	//
	// First line of the description for the sitelink. To clear this field, set its value to the empty
	// string. If this value is set, sitelinkLine3 must also be set.
	// <span class="constraint StringLength">The length of this string should be between 0 and 35, inclusive, (trimmed).</span>
	//
	SitelinkLine2 string `xml:"sitelinkLine2,omitempty"`

	//
	// Second line of the description for the sitelink. To clear this field, set its value to the
	// empty string. If this value is set, sitelinkLine2 must also be set.
	// <span class="constraint StringLength">The length of this string should be between 0 and 35, inclusive, (trimmed).</span>
	//
	SitelinkLine3 string `xml:"sitelinkLine3,omitempty"`

	//
	// A list of possible final URLs after all cross domain redirects.
	//
	SitelinkFinalUrls *UrlList `xml:"sitelinkFinalUrls,omitempty"`

	//
	// A list of possible final mobile URLs after all cross domain redirects.
	//
	SitelinkFinalMobileUrls *UrlList `xml:"sitelinkFinalMobileUrls,omitempty"`

	//
	// URL template for constructing a tracking URL. To clear this field, set its value to the empty
	// string.
	//
	SitelinkTrackingUrlTemplate string `xml:"sitelinkTrackingUrlTemplate,omitempty"`

	//
	// A list of mappings to be used for substituting URL custom parameter tags in the
	// trackingUrlTemplate, finalUrls, and/or finalMobileUrls.
	//
	SitelinkUrlCustomParameters *CustomParameters `xml:"sitelinkUrlCustomParameters,omitempty"`
}

type SizeLimitError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 SizeLimitError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *SizeLimitErrorReason `xml:"reason,omitempty"`
}

type SoapHeader struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 SoapHeader"`

	//
	// The header identifies the customer id of the client of the AdWords manager, if an AdWords
	// manager is acting on behalf of their client or the customer id of the advertiser managing their
	// own account.
	//
	ClientCustomerId string `xml:"clientCustomerId,omitempty"`

	//
	// Developer token to identify that the person making the call has enough
	// quota.
	//
	DeveloperToken string `xml:"developerToken,omitempty"`

	//
	// UserAgent is used to track distribution of API client programs and
	// application usage. The client is responsible for putting in a meaningful
	// value for tracking purposes. To be clear this is not the same as an HTTP
	// user agent.
	//
	UserAgent string `xml:"userAgent,omitempty"`

	//
	// Used to validate the request without executing it.
	//
	ValidateOnly bool `xml:"validateOnly,omitempty"`

	//
	// If true, API will try to commit as many error free operations as possible and
	// report the other operations' errors.
	//
	// <p>Ignored for non-mutate calls.
	//
	PartialFailure bool `xml:"partialFailure,omitempty"`
}

type SoapResponseHeader struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 SoapResponseHeader"`

	//
	// Unique id that identifies this request. If developers have any support issues, sending us
	// this id will enable us to find their request more easily.
	//
	RequestId string `xml:"requestId,omitempty"`

	//
	// The name of the service being invoked.
	//
	ServiceName string `xml:"serviceName,omitempty"`

	//
	// The name of the method being invoked.
	//
	MethodName string `xml:"methodName,omitempty"`

	//
	// Number of operations performed for this SOAP request.
	//
	Operations int64 `xml:"operations,omitempty"`

	//
	// Elapsed time in milliseconds between the AdWords API receiving the request and sending the
	// response.
	//
	ResponseTime int64 `xml:"responseTime,omitempty"`
}

type StringFormatError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 StringFormatError"`

	*ApiError

	Reason *StringFormatErrorReason `xml:"reason,omitempty"`
}

type StringLengthError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 StringLengthError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *StringLengthErrorReason `xml:"reason,omitempty"`
}

type StructuredSnippetFeedItem struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 StructuredSnippetFeedItem"`

	*ExtensionFeedItem

	//
	// The header of the snippet. See the
	// <a href="https://developers.google.com/adwords/api/docs/appendix/structured-snippet-headers">
	// structured snippets header translations</a> page for supported localized headers.
	// <span class="constraint StringLength">This string must not be empty.</span>
	//
	Header string `xml:"header,omitempty"`

	//
	// The values in the snippet. A SET operation replaces the values in the list.
	// <span class="constraint CollectionSize">The maximum size of this collection is 10.</span>
	// <span class="constraint ContentsDistinct">This field must contain distinct elements.</span>
	// <span class="constraint ContentsNotNull">This field must not contain {@code null} elements.</span>
	//
	Values []string `xml:"values,omitempty"`
}

type UrlError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 UrlError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *UrlErrorReason `xml:"reason,omitempty"`
}

type UrlList struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 UrlList"`

	//
	// List of URLs.  On SET operation, empty list indicates to clear the list.
	// <span class="constraint CollectionSize">The maximum size of this collection is 10.</span>
	// <span class="constraint ContentsStringLength">Strings in this field must be non-empty (trimmed).</span>
	//
	Urls []string `xml:"urls,omitempty"`
}

type CriterionUserInterest struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CriterionUserInterest"`

	*Criterion

	//
	// Id of this user interest. This is a required field.
	//
	UserInterestId int64 `xml:"userInterestId,omitempty"`

	//
	// Parent Id of this user interest.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	UserInterestParentId int64 `xml:"userInterestParentId,omitempty"`

	//
	// Name of this user interest.
	//
	UserInterestName string `xml:"userInterestName,omitempty"`
}

type CriterionUserList struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CriterionUserList"`

	*Criterion

	//
	// Id of this user list. This is a required field.
	//
	UserListId int64 `xml:"userListId,omitempty"`

	UserListName string `xml:"userListName,omitempty"`

	UserListMembershipStatus *CriterionUserListMembershipStatus `xml:"userListMembershipStatus,omitempty"`

	//
	// Determines whether a user list is eligible for targeting in the google.com
	// (search) network.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	UserListEligibleForSearch bool `xml:"userListEligibleForSearch,omitempty"`

	//
	// Determines whether a user list is eligible for targeting in the display network.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	UserListEligibleForDisplay bool `xml:"userListEligibleForDisplay,omitempty"`
}

type Vertical struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Vertical"`

	*Criterion

	//
	// Id of this vertical.
	//
	VerticalId int64 `xml:"verticalId,omitempty"`

	//
	// Id of the parent of this vertical.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	VerticalParentId int64 `xml:"verticalParentId,omitempty"`

	//
	// The category to target or exclude. Each subsequent element in the array
	// describes a more specific sub-category. For example,
	// <code>{"Pets &amp; Animals", "Pets", "Dogs"}</code> represents the "Pets &amp;
	// Animals/Pets/Dogs" category. A complete list of available vertical categories
	// is available <a href="/adwords/api/docs/appendix/verticals">here</a>
	// This field is required and must not be empty.
	//
	Path []string `xml:"path,omitempty"`
}

type CampaignExtensionSettingServiceInterface struct {
	client *SOAPClient
}

func NewCampaignExtensionSettingServiceInterface(url string, tls bool, auth *BasicAuth) *CampaignExtensionSettingServiceInterface {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &CampaignExtensionSettingServiceInterface{
		client: client,
	}
}

func NewCampaignExtensionSettingServiceInterfaceWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *CampaignExtensionSettingServiceInterface {
	if url == "" {
		url = ""
	}
	client := NewSOAPClientWithTLSConfig(url, tlsCfg, auth)

	return &CampaignExtensionSettingServiceInterface{
		client: client,
	}
}

func (service *CampaignExtensionSettingServiceInterface) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *CampaignExtensionSettingServiceInterface) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Error can be either of the following types:
//
//   - ApiException
/*
   Returns a list of CampaignExtensionSettings that meet the selector criteria.

   @param selector Determines which CampaignExtensionSettings to return. If empty, all
   CampaignExtensionSettings are returned.
   @return The list of CampaignExtensionSettings specified by the selector.
   @throws ApiException Indicates a problem with the request.
*/
func (service *CampaignExtensionSettingServiceInterface) Get(request *Get) (*GetResponse, error) {
	response := new(GetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - ApiException
/*
   Applies the list of mutate operations (add, remove, and set).

   <p> Beginning in v201509, add and set operations are treated identically. Performing an add
   operation on a campaign with an existing ExtensionSetting will cause the operation to be
   treated like a set operation. Performing a set operation on a campaign with no
   ExtensionSetting will cause the operation to be treated like an add operation.

   @param operations The operations to apply. The same {@link CampaignExtensionSetting} cannot be
   specified in more than one operation.
   @return The changed {@link CampaignExtensionSetting}s.
   @throws ApiException Indicates a problem with the request.
*/
func (service *CampaignExtensionSettingServiceInterface) Mutate(request *Mutate) (*MutateResponse, error) {
	response := new(MutateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - ApiException
/*
   Returns a list of CampaignExtensionSettings that match the query.

   @param query The SQL-like AWQL query string.
   @return The list of CampaignExtensionSettings specified by the query.
   @throws ApiException Indicates a problem with the request.
*/
func (service *CampaignExtensionSettingServiceInterface) Query(request *Query) (*QueryResponse, error) {
	response := new(QueryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Items []interface{} `xml:",omitempty"`
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU  string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
)

type WSSSecurityHeader struct {
	XMLName   xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ wsse:Security"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	MustUnderstand string `xml:"mustUnderstand,attr,omitempty"`

	Token *WSSUsernameToken `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName   xml.Name `xml:"wsse:UsernameToken"`
	XmlNSWsu  string   `xml:"xmlns:wsu,attr"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Id string `xml:"wsu:Id,attr,omitempty"`

	Username *WSSUsername `xml:",omitempty"`
	Password *WSSPassword `xml:",omitempty"`
}

type WSSUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Data string `xml:",chardata"`
}

type WSSPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	XmlNSType string   `xml:"Type,attr"`

	Data string `xml:",chardata"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url     string
	tlsCfg  *tls.Config
	auth    *BasicAuth
	headers []interface{}
}

// **********
// Accepted solution from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
// Author: Icza - http://stackoverflow.com/users/1705598/icza

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrc(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// **********

func NewWSSSecurityHeader(user, pass, mustUnderstand string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{XmlNSWsse: WssNsWSSE, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{XmlNSWsu: WssNsWSU, XmlNSWsse: WssNsWSSE, Id: "UsernameToken-" + randStringBytesMaskImprSrc(9)}
	hdr.Token.Username = &WSSUsername{XmlNSWsse: WssNsWSSE, Data: user}
	hdr.Token.Password = &WSSPassword{XmlNSWsse: WssNsWSSE, XmlNSType: WssNsType, Data: pass}
	return hdr
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, insecureSkipVerify bool, auth *BasicAuth) *SOAPClient {
	tlsCfg := &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}
	return NewSOAPClientWithTLSConfig(url, tlsCfg, auth)
}

func NewSOAPClientWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:    url,
		tlsCfg: tlsCfg,
		auth:   auth,
	}
}

func (s *SOAPClient) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.headers != nil && len(s.headers) > 0 {
		soapHeader := &SOAPHeader{Items: make([]interface{}, len(s.headers))}
		copy(soapHeader.Items, s.headers)
		envelope.Header = soapHeader
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: s.tlsCfg,
		Dial:            dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
