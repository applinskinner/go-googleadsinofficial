package AdGroupAdService

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
// The types of ads.
//
type AdType string

const (
	AdTypeDEPRECATED_AD AdType = "DEPRECATED_AD"

	AdTypeIMAGE_AD AdType = "IMAGE_AD"

	AdTypePRODUCT_AD AdType = "PRODUCT_AD"

	AdTypeTEMPLATE_AD AdType = "TEMPLATE_AD"

	AdTypeTEXT_AD AdType = "TEXT_AD"

	AdTypeTHIRD_PARTY_REDIRECT_AD AdType = "THIRD_PARTY_REDIRECT_AD"

	AdTypeDYNAMIC_SEARCH_AD AdType = "DYNAMIC_SEARCH_AD"

	AdTypeCALL_ONLY_AD AdType = "CALL_ONLY_AD"

	AdTypeEXPANDED_TEXT_AD AdType = "EXPANDED_TEXT_AD"

	AdTypeRESPONSIVE_DISPLAY_AD AdType = "RESPONSIVE_DISPLAY_AD"

	AdTypeSHOWCASE_AD AdType = "SHOWCASE_AD"

	AdTypeUNIVERSAL_SHOPPING_AD AdType = "UNIVERSAL_SHOPPING_AD"

	AdTypeEXPANDED_DYNAMIC_SEARCH_AD AdType = "EXPANDED_DYNAMIC_SEARCH_AD"

	AdTypeGMAIL_AD AdType = "GMAIL_AD"

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	AdTypeUNKNOWN AdType = "UNKNOWN"
)

//
// Ad customizer error reasons.
//
type AdCustomizerErrorReason string

const (

	//
	// Invalid date argument in countdown function.
	//
	AdCustomizerErrorReasonCOUNTDOWN_INVALID_DATE_FORMAT AdCustomizerErrorReason = "COUNTDOWN_INVALID_DATE_FORMAT"

	//
	// Countdown end date is in the past.
	//
	AdCustomizerErrorReasonCOUNTDOWN_DATE_IN_PAST AdCustomizerErrorReason = "COUNTDOWN_DATE_IN_PAST"

	//
	// Invalid locale string in countdown function.
	//
	AdCustomizerErrorReasonCOUNTDOWN_INVALID_LOCALE AdCustomizerErrorReason = "COUNTDOWN_INVALID_LOCALE"

	//
	// Days-before argument to countdown function is not positive.
	//
	AdCustomizerErrorReasonCOUNTDOWN_INVALID_START_DAYS_BEFORE AdCustomizerErrorReason = "COUNTDOWN_INVALID_START_DAYS_BEFORE"

	//
	// A user list referenced in an IF function does not exist.
	//
	AdCustomizerErrorReasonUNKNOWN_USER_LIST AdCustomizerErrorReason = "UNKNOWN_USER_LIST"
)

//
// The reasons for the target error.
//
type AdErrorReason string

const (

	//
	// Ad customizers are not supported for ad type.
	//
	AdErrorReasonAD_CUSTOMIZERS_NOT_SUPPORTED_FOR_AD_TYPE AdErrorReason = "AD_CUSTOMIZERS_NOT_SUPPORTED_FOR_AD_TYPE"

	//
	// Estimating character sizes the string is too long.
	//
	AdErrorReasonAPPROXIMATELY_TOO_LONG AdErrorReason = "APPROXIMATELY_TOO_LONG"

	//
	// Estimating character sizes the string is too short.
	//
	AdErrorReasonAPPROXIMATELY_TOO_SHORT AdErrorReason = "APPROXIMATELY_TOO_SHORT"

	//
	// There is a problem with the snippet.
	//
	AdErrorReasonBAD_SNIPPET AdErrorReason = "BAD_SNIPPET"

	//
	// Cannot modify an ad.
	//
	AdErrorReasonCANNOT_MODIFY_AD AdErrorReason = "CANNOT_MODIFY_AD"

	//
	// business name and url cannot be set at the same time
	//
	AdErrorReasonCANNOT_SET_BUSINESS_NAME_IF_URL_SET AdErrorReason = "CANNOT_SET_BUSINESS_NAME_IF_URL_SET"

	//
	// The specified field is incompatible with this ad's type or settings.
	//
	AdErrorReasonCANNOT_SET_FIELD AdErrorReason = "CANNOT_SET_FIELD"

	//
	// Cannot set field when originAdId is set.
	//
	AdErrorReasonCANNOT_SET_FIELD_WITH_ORIGIN_AD_ID_SET AdErrorReason = "CANNOT_SET_FIELD_WITH_ORIGIN_AD_ID_SET"

	//
	// Cannot set field when an existing ad id is set for sharing.
	//
	AdErrorReasonCANNOT_SET_FIELD_WITH_AD_ID_SET_FOR_SHARING AdErrorReason = "CANNOT_SET_FIELD_WITH_AD_ID_SET_FOR_SHARING"

	//
	// Cannot set allowFlexibleColor false if no color is provided by user.
	//
	AdErrorReasonCANNOT_SET_ALLOW_FLEXIBLE_COLOR_FALSE AdErrorReason = "CANNOT_SET_ALLOW_FLEXIBLE_COLOR_FALSE"

	//
	// When user select native, no color control is allowed because we will always respect publisher
	// color for native format serving.
	//
	AdErrorReasonCANNOT_SET_COLOR_CONTROL_WHEN_NATIVE_FORMAT_SETTING AdErrorReason = "CANNOT_SET_COLOR_CONTROL_WHEN_NATIVE_FORMAT_SETTING"

	//
	// Cannot specify a url for the ad type
	//
	AdErrorReasonCANNOT_SET_URL AdErrorReason = "CANNOT_SET_URL"

	//
	// Cannot specify a tracking or mobile url without also setting final urls
	//
	AdErrorReasonCANNOT_SET_WITHOUT_FINAL_URLS AdErrorReason = "CANNOT_SET_WITHOUT_FINAL_URLS"

	//
	// Cannot specify a legacy url and a final url simultaneously
	//
	AdErrorReasonCANNOT_SET_WITH_FINAL_URLS AdErrorReason = "CANNOT_SET_WITH_FINAL_URLS"

	//
	// Cannot specify a legacy url and a tracking url template simultaneously in a DSA.
	//
	AdErrorReasonCANNOT_SET_WITH_TRACKING_URL_TEMPLATE AdErrorReason = "CANNOT_SET_WITH_TRACKING_URL_TEMPLATE"

	//
	// Cannot specify a urls in UrlData and in template fields simultaneously.
	//
	AdErrorReasonCANNOT_SET_WITH_URL_DATA AdErrorReason = "CANNOT_SET_WITH_URL_DATA"

	//
	// This operator cannot be used with a subclass of Ad.
	//
	AdErrorReasonCANNOT_USE_AD_SUBCLASS_FOR_OPERATOR AdErrorReason = "CANNOT_USE_AD_SUBCLASS_FOR_OPERATOR"

	//
	// Customer is not approved for mobile ads.
	//
	AdErrorReasonCUSTOMER_NOT_APPROVED_MOBILEADS AdErrorReason = "CUSTOMER_NOT_APPROVED_MOBILEADS"

	//
	// Customer is not approved for 3PAS richmedia ads.
	//
	AdErrorReasonCUSTOMER_NOT_APPROVED_THIRDPARTY_ADS AdErrorReason = "CUSTOMER_NOT_APPROVED_THIRDPARTY_ADS"

	//
	// Customer is not approved for 3PAS redirect richmedia (Ad Exchange) ads.
	//
	AdErrorReasonCUSTOMER_NOT_APPROVED_THIRDPARTY_REDIRECT_ADS AdErrorReason = "CUSTOMER_NOT_APPROVED_THIRDPARTY_REDIRECT_ADS"

	//
	// Not an eligible customer
	//
	AdErrorReasonCUSTOMER_NOT_ELIGIBLE AdErrorReason = "CUSTOMER_NOT_ELIGIBLE"

	//
	// Customer is not eligible for updating beacon url
	//
	AdErrorReasonCUSTOMER_NOT_ELIGIBLE_FOR_UPDATING_BEACON_URL AdErrorReason = "CUSTOMER_NOT_ELIGIBLE_FOR_UPDATING_BEACON_URL"

	//
	// There already exists an ad with the same dimensions in the union.
	//
	AdErrorReasonDIMENSION_ALREADY_IN_UNION AdErrorReason = "DIMENSION_ALREADY_IN_UNION"

	//
	// Ad's dimension must be set before setting union dimension.
	//
	AdErrorReasonDIMENSION_MUST_BE_SET AdErrorReason = "DIMENSION_MUST_BE_SET"

	//
	// Ad's dimension must be included in the union dimensions.
	//
	AdErrorReasonDIMENSION_NOT_IN_UNION AdErrorReason = "DIMENSION_NOT_IN_UNION"

	//
	// Display Url cannot be specified (applies to Ad Exchange Ads)
	//
	AdErrorReasonDISPLAY_URL_CANNOT_BE_SPECIFIED AdErrorReason = "DISPLAY_URL_CANNOT_BE_SPECIFIED"

	//
	// Telephone number contains invalid characters or invalid format.
	// Please re-enter your number using digits (0-9), dashes (-), and parentheses only.
	//
	AdErrorReasonDOMESTIC_PHONE_NUMBER_FORMAT AdErrorReason = "DOMESTIC_PHONE_NUMBER_FORMAT"

	//
	// Emergency telephone numbers are not allowed.
	// Please enter a valid domestic phone number to connect customers to your business.
	//
	AdErrorReasonEMERGENCY_PHONE_NUMBER AdErrorReason = "EMERGENCY_PHONE_NUMBER"

	//
	// A required field was not specified or is an empty string.
	//
	AdErrorReasonEMPTY_FIELD AdErrorReason = "EMPTY_FIELD"

	//
	// A feed attribute referenced in an ad customizer tag is not in the ad customizer mapping for
	// the feed.
	//
	AdErrorReasonFEED_ATTRIBUTE_MUST_HAVE_MAPPING_FOR_TYPE_ID AdErrorReason = "FEED_ATTRIBUTE_MUST_HAVE_MAPPING_FOR_TYPE_ID"

	//
	// The ad customizer field mapping for the feed attribute does not match the expected field
	// type.
	//
	AdErrorReasonFEED_ATTRIBUTE_MAPPING_TYPE_MISMATCH AdErrorReason = "FEED_ATTRIBUTE_MAPPING_TYPE_MISMATCH"

	//
	// The use of ad customizer tags in the ad text is disallowed. Details in trigger.
	//
	AdErrorReasonILLEGAL_AD_CUSTOMIZER_TAG_USE AdErrorReason = "ILLEGAL_AD_CUSTOMIZER_TAG_USE"

	//
	// Tags of the form {PH_x}, where x is a number, are disallowed in ad text.
	//
	AdErrorReasonILLEGAL_TAG_USE AdErrorReason = "ILLEGAL_TAG_USE"

	//
	// The dimensions of the ad are specified or derived in multiple ways and are not consistent.
	//
	AdErrorReasonINCONSISTENT_DIMENSIONS AdErrorReason = "INCONSISTENT_DIMENSIONS"

	//
	// The status cannot differ among template ads of the same union.
	//
	AdErrorReasonINCONSISTENT_STATUS_IN_TEMPLATE_UNION AdErrorReason = "INCONSISTENT_STATUS_IN_TEMPLATE_UNION"

	//
	// The length of the string is not valid.
	//
	AdErrorReasonINCORRECT_LENGTH AdErrorReason = "INCORRECT_LENGTH"

	//
	// The ad is ineligible for upgrade.
	//
	AdErrorReasonINELIGIBLE_FOR_UPGRADE AdErrorReason = "INELIGIBLE_FOR_UPGRADE"

	//
	// User cannot create mobile ad for countries targeted in specified campaign.
	//
	AdErrorReasonINVALID_AD_ADDRESS_CAMPAIGN_TARGET AdErrorReason = "INVALID_AD_ADDRESS_CAMPAIGN_TARGET"

	//
	// Invalid Ad type. A specific type of Ad is required.
	//
	AdErrorReasonINVALID_AD_TYPE AdErrorReason = "INVALID_AD_TYPE"

	//
	// Headline, description or phone cannot be present when creating mobile image ad.
	//
	AdErrorReasonINVALID_ATTRIBUTES_FOR_MOBILE_IMAGE AdErrorReason = "INVALID_ATTRIBUTES_FOR_MOBILE_IMAGE"

	//
	// Image cannot be present when creating mobile text ad.
	//
	AdErrorReasonINVALID_ATTRIBUTES_FOR_MOBILE_TEXT AdErrorReason = "INVALID_ATTRIBUTES_FOR_MOBILE_TEXT"

	//
	// Invalid call to action text.
	//
	AdErrorReasonINVALID_CALL_TO_ACTION_TEXT AdErrorReason = "INVALID_CALL_TO_ACTION_TEXT"

	//
	// Invalid character in URL.
	//
	AdErrorReasonINVALID_CHARACTER_FOR_URL AdErrorReason = "INVALID_CHARACTER_FOR_URL"

	//
	// Creative's country code is not valid.
	//
	AdErrorReasonINVALID_COUNTRY_CODE AdErrorReason = "INVALID_COUNTRY_CODE"

	//
	// Invalid use of Dynamic Search Ads tags ({lpurl} etc.)
	//
	AdErrorReasonINVALID_DSA_URL_TAG AdErrorReason = "INVALID_DSA_URL_TAG"

	//
	// Invalid use of Expanded Dynamic Search Ads tags ({lpurl} etc.)
	//
	AdErrorReasonINVALID_EXPANDED_DYNAMIC_SEARCH_AD_TAG AdErrorReason = "INVALID_EXPANDED_DYNAMIC_SEARCH_AD_TAG"

	//
	// An input error whose real reason was not properly mapped (should not happen).
	//
	AdErrorReasonINVALID_INPUT AdErrorReason = "INVALID_INPUT"

	//
	// An invalid markup language was entered.
	//
	AdErrorReasonINVALID_MARKUP_LANGUAGE AdErrorReason = "INVALID_MARKUP_LANGUAGE"

	//
	// An invalid mobile carrier was entered.
	//
	AdErrorReasonINVALID_MOBILE_CARRIER AdErrorReason = "INVALID_MOBILE_CARRIER"

	//
	// Specified mobile carriers target a country not targeted by the campaign.
	//
	AdErrorReasonINVALID_MOBILE_CARRIER_TARGET AdErrorReason = "INVALID_MOBILE_CARRIER_TARGET"

	//
	// Wrong number of elements for given element type
	//
	AdErrorReasonINVALID_NUMBER_OF_ELEMENTS AdErrorReason = "INVALID_NUMBER_OF_ELEMENTS"

	//
	// The format of the telephone number is incorrect.
	// Please re-enter the number using the correct format.
	//
	AdErrorReasonINVALID_PHONE_NUMBER_FORMAT AdErrorReason = "INVALID_PHONE_NUMBER_FORMAT"

	//
	// The certified vendor format id is incorrect.
	//
	AdErrorReasonINVALID_RICH_MEDIA_CERTIFIED_VENDOR_FORMAT_ID AdErrorReason = "INVALID_RICH_MEDIA_CERTIFIED_VENDOR_FORMAT_ID"

	//
	// The template ad data contains validation errors.
	//
	AdErrorReasonINVALID_TEMPLATE_DATA AdErrorReason = "INVALID_TEMPLATE_DATA"

	//
	// The template field doesn't have have the correct type.
	//
	AdErrorReasonINVALID_TEMPLATE_ELEMENT_FIELD_TYPE AdErrorReason = "INVALID_TEMPLATE_ELEMENT_FIELD_TYPE"

	//
	// Invalid template id.
	//
	AdErrorReasonINVALID_TEMPLATE_ID AdErrorReason = "INVALID_TEMPLATE_ID"

	//
	// After substituting replacement strings, the line is too wide.
	//
	AdErrorReasonLINE_TOO_WIDE AdErrorReason = "LINE_TOO_WIDE"

	//
	// The feed referenced must have ad customizer mapping to be used in a customizer tag.
	//
	AdErrorReasonMISSING_AD_CUSTOMIZER_MAPPING AdErrorReason = "MISSING_AD_CUSTOMIZER_MAPPING"

	//
	// Missing address component in template element address field.
	//
	AdErrorReasonMISSING_ADDRESS_COMPONENT AdErrorReason = "MISSING_ADDRESS_COMPONENT"

	//
	// An ad name must be entered.
	//
	AdErrorReasonMISSING_ADVERTISEMENT_NAME AdErrorReason = "MISSING_ADVERTISEMENT_NAME"

	//
	// Business name must be entered.
	//
	AdErrorReasonMISSING_BUSINESS_NAME AdErrorReason = "MISSING_BUSINESS_NAME"

	//
	// Description (line 2) must be entered.
	//
	AdErrorReasonMISSING_DESCRIPTION1 AdErrorReason = "MISSING_DESCRIPTION1"

	//
	// Description (line 3) must be entered.
	//
	AdErrorReasonMISSING_DESCRIPTION2 AdErrorReason = "MISSING_DESCRIPTION2"

	//
	// The destination url must contain at least one tag (e.g. {lpurl})
	//
	AdErrorReasonMISSING_DESTINATION_URL_TAG AdErrorReason = "MISSING_DESTINATION_URL_TAG"

	//
	// The tracking url template of ExpandedDynamicSearchAd must contain at least one tag.
	// (e.g. {lpurl})
	//
	AdErrorReasonMISSING_LANDING_PAGE_URL_TAG AdErrorReason = "MISSING_LANDING_PAGE_URL_TAG"

	//
	// A valid dimension must be specified for this ad.
	//
	AdErrorReasonMISSING_DIMENSION AdErrorReason = "MISSING_DIMENSION"

	//
	// A display URL must be entered.
	//
	AdErrorReasonMISSING_DISPLAY_URL AdErrorReason = "MISSING_DISPLAY_URL"

	//
	// Headline must be entered.
	//
	AdErrorReasonMISSING_HEADLINE AdErrorReason = "MISSING_HEADLINE"

	//
	// A height must be entered.
	//
	AdErrorReasonMISSING_HEIGHT AdErrorReason = "MISSING_HEIGHT"

	//
	// An image must be entered.
	//
	AdErrorReasonMISSING_IMAGE AdErrorReason = "MISSING_IMAGE"

	//
	// Marketing image or product videos are required.
	//
	AdErrorReasonMISSING_MARKETING_IMAGE_OR_PRODUCT_VIDEOS AdErrorReason = "MISSING_MARKETING_IMAGE_OR_PRODUCT_VIDEOS"

	//
	// The markup language in which your site is written must be entered.
	//
	AdErrorReasonMISSING_MARKUP_LANGUAGES AdErrorReason = "MISSING_MARKUP_LANGUAGES"

	//
	// A mobile carrier must be entered.
	//
	AdErrorReasonMISSING_MOBILE_CARRIER AdErrorReason = "MISSING_MOBILE_CARRIER"

	//
	// Phone number must be entered.
	//
	AdErrorReasonMISSING_PHONE AdErrorReason = "MISSING_PHONE"

	//
	// Missing required template fields
	//
	AdErrorReasonMISSING_REQUIRED_TEMPLATE_FIELDS AdErrorReason = "MISSING_REQUIRED_TEMPLATE_FIELDS"

	//
	// Missing a required field value
	//
	AdErrorReasonMISSING_TEMPLATE_FIELD_VALUE AdErrorReason = "MISSING_TEMPLATE_FIELD_VALUE"

	//
	// The ad must have text.
	//
	AdErrorReasonMISSING_TEXT AdErrorReason = "MISSING_TEXT"

	//
	// A visible URL must be entered.
	//
	AdErrorReasonMISSING_VISIBLE_URL AdErrorReason = "MISSING_VISIBLE_URL"

	//
	// A width must be entered.
	//
	AdErrorReasonMISSING_WIDTH AdErrorReason = "MISSING_WIDTH"

	//
	// Only 1 feed can be used as the source of ad customizer substitutions in a single ad.
	//
	AdErrorReasonMULTIPLE_DISTINCT_FEEDS_UNSUPPORTED AdErrorReason = "MULTIPLE_DISTINCT_FEEDS_UNSUPPORTED"

	//
	// TempAdUnionId must be use when adding template ads.
	//
	AdErrorReasonMUST_USE_TEMP_AD_UNION_ID_ON_ADD AdErrorReason = "MUST_USE_TEMP_AD_UNION_ID_ON_ADD"

	//
	// The string has too many characters.
	//
	AdErrorReasonTOO_LONG AdErrorReason = "TOO_LONG"

	//
	// The string has too few characters.
	//
	AdErrorReasonTOO_SHORT AdErrorReason = "TOO_SHORT"

	//
	// Ad union dimensions cannot change for saved ads.
	//
	AdErrorReasonUNION_DIMENSIONS_CANNOT_CHANGE AdErrorReason = "UNION_DIMENSIONS_CANNOT_CHANGE"

	//
	// Address component is not {country, lat, lng}.
	//
	AdErrorReasonUNKNOWN_ADDRESS_COMPONENT AdErrorReason = "UNKNOWN_ADDRESS_COMPONENT"

	//
	// Unknown unique field name
	//
	AdErrorReasonUNKNOWN_FIELD_NAME AdErrorReason = "UNKNOWN_FIELD_NAME"

	//
	// Unknown unique name (template element type specifier)
	//
	AdErrorReasonUNKNOWN_UNIQUE_NAME AdErrorReason = "UNKNOWN_UNIQUE_NAME"

	//
	// Unsupported ad dimension
	//
	AdErrorReasonUNSUPPORTED_DIMENSIONS AdErrorReason = "UNSUPPORTED_DIMENSIONS"

	//
	// URL starts with an invalid scheme.
	//
	AdErrorReasonURL_INVALID_SCHEME AdErrorReason = "URL_INVALID_SCHEME"

	//
	// URL ends with an invalid top-level domain name.
	//
	AdErrorReasonURL_INVALID_TOP_LEVEL_DOMAIN AdErrorReason = "URL_INVALID_TOP_LEVEL_DOMAIN"

	//
	// URL contains illegal characters.
	//
	AdErrorReasonURL_MALFORMED AdErrorReason = "URL_MALFORMED"

	//
	// URL must contain a host name.
	//
	AdErrorReasonURL_NO_HOST AdErrorReason = "URL_NO_HOST"

	//
	// URL not equivalent during upgrade.
	//
	AdErrorReasonURL_NOT_EQUIVALENT AdErrorReason = "URL_NOT_EQUIVALENT"

	//
	// URL host name too long to be stored as visible URL (applies to Ad Exchange ads)
	//
	AdErrorReasonURL_HOST_NAME_TOO_LONG AdErrorReason = "URL_HOST_NAME_TOO_LONG"

	//
	// URL must start with a scheme.
	//
	AdErrorReasonURL_NO_SCHEME AdErrorReason = "URL_NO_SCHEME"

	//
	// URL should end in a valid domain extension, such as .com or .net.
	//
	AdErrorReasonURL_NO_TOP_LEVEL_DOMAIN AdErrorReason = "URL_NO_TOP_LEVEL_DOMAIN"

	//
	// URL must not end with a path.
	//
	AdErrorReasonURL_PATH_NOT_ALLOWED AdErrorReason = "URL_PATH_NOT_ALLOWED"

	//
	// URL must not specify a port.
	//
	AdErrorReasonURL_PORT_NOT_ALLOWED AdErrorReason = "URL_PORT_NOT_ALLOWED"

	//
	// URL must not contain a query.
	//
	AdErrorReasonURL_QUERY_NOT_ALLOWED AdErrorReason = "URL_QUERY_NOT_ALLOWED"

	//
	// A url scheme is not allowed in front of tag in dest url (e.g. http://{lpurl})
	//
	AdErrorReasonURL_SCHEME_BEFORE_DSA_TAG AdErrorReason = "URL_SCHEME_BEFORE_DSA_TAG"

	//
	// A url scheme is not allowed in front of tag in tracking url template (e.g. http://{lpurl})
	//
	AdErrorReasonURL_SCHEME_BEFORE_EXPANDED_DYNAMIC_SEARCH_AD_TAG AdErrorReason = "URL_SCHEME_BEFORE_EXPANDED_DYNAMIC_SEARCH_AD_TAG"

	//
	// The user does not have permissions to create a template ad for the given
	// template.
	//
	AdErrorReasonUSER_DOES_NOT_HAVE_ACCESS_TO_TEMPLATE AdErrorReason = "USER_DOES_NOT_HAVE_ACCESS_TO_TEMPLATE"

	//
	// Expandable setting is inconsistent/wrong. For example, an AdX ad is
	// invalid if it has a expandable vendor format but no expanding directions
	// specified, or expanding directions is specified, but the vendor format
	// is not expandable.
	//
	AdErrorReasonINCONSISTENT_EXPANDABLE_SETTINGS AdErrorReason = "INCONSISTENT_EXPANDABLE_SETTINGS"

	//
	// Format is invalid
	//
	AdErrorReasonINVALID_FORMAT AdErrorReason = "INVALID_FORMAT"

	//
	// The text of this field did not match a pattern of allowed values.
	//
	AdErrorReasonINVALID_FIELD_TEXT AdErrorReason = "INVALID_FIELD_TEXT"

	//
	// Template element is mising
	//
	AdErrorReasonELEMENT_NOT_PRESENT AdErrorReason = "ELEMENT_NOT_PRESENT"

	//
	// Error occurred during image processing
	//
	AdErrorReasonIMAGE_ERROR AdErrorReason = "IMAGE_ERROR"

	//
	// The value is not within the valid range
	//
	AdErrorReasonVALUE_NOT_IN_RANGE AdErrorReason = "VALUE_NOT_IN_RANGE"

	//
	// Template element field is not present
	//
	AdErrorReasonFIELD_NOT_PRESENT AdErrorReason = "FIELD_NOT_PRESENT"

	//
	// Address is incomplete
	//
	AdErrorReasonADDRESS_NOT_COMPLETE AdErrorReason = "ADDRESS_NOT_COMPLETE"

	//
	// Invalid address
	//
	AdErrorReasonADDRESS_INVALID AdErrorReason = "ADDRESS_INVALID"

	//
	// Error retrieving specified video
	//
	AdErrorReasonVIDEO_RETRIEVAL_ERROR AdErrorReason = "VIDEO_RETRIEVAL_ERROR"

	//
	// Error processing audio
	//
	AdErrorReasonAUDIO_ERROR AdErrorReason = "AUDIO_ERROR"

	//
	// Display URL is incorrect for YouTube PYV ads
	//
	AdErrorReasonINVALID_YOUTUBE_DISPLAY_URL AdErrorReason = "INVALID_YOUTUBE_DISPLAY_URL"

	//
	// Too many product Images in GmailAd
	//
	AdErrorReasonTOO_MANY_PRODUCT_IMAGES AdErrorReason = "TOO_MANY_PRODUCT_IMAGES"

	//
	// Too many product Videos in GmailAd
	//
	AdErrorReasonTOO_MANY_PRODUCT_VIDEOS AdErrorReason = "TOO_MANY_PRODUCT_VIDEOS"

	//
	// The device preference is not compatible with the ad type
	//
	AdErrorReasonINCOMPATIBLE_AD_TYPE_AND_DEVICE_PREFERENCE AdErrorReason = "INCOMPATIBLE_AD_TYPE_AND_DEVICE_PREFERENCE"

	//
	// Call tracking is not supported for specified country.
	//
	AdErrorReasonCALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY AdErrorReason = "CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY"

	//
	// Carrier specific short number is not allowed.
	//
	AdErrorReasonCARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED AdErrorReason = "CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED"

	//
	// Specified phone number type is disallowed.
	//
	AdErrorReasonDISALLOWED_NUMBER_TYPE AdErrorReason = "DISALLOWED_NUMBER_TYPE"

	//
	// Phone number not supported for country.
	//
	AdErrorReasonPHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY AdErrorReason = "PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY"

	//
	// Phone number not supported with call tracking enabled for country.
	//
	AdErrorReasonPHONE_NUMBER_NOT_SUPPORTED_WITH_CALLTRACKING_FOR_COUNTRY AdErrorReason = "PHONE_NUMBER_NOT_SUPPORTED_WITH_CALLTRACKING_FOR_COUNTRY"

	//
	// Premium rate phone number is not allowed.
	//
	AdErrorReasonPREMIUM_RATE_NUMBER_NOT_ALLOWED AdErrorReason = "PREMIUM_RATE_NUMBER_NOT_ALLOWED"

	//
	// Vanity phone number is not allowed.
	//
	AdErrorReasonVANITY_PHONE_NUMBER_NOT_ALLOWED AdErrorReason = "VANITY_PHONE_NUMBER_NOT_ALLOWED"

	//
	// Invalid call conversion type id.
	//
	AdErrorReasonINVALID_CALL_CONVERSION_TYPE_ID AdErrorReason = "INVALID_CALL_CONVERSION_TYPE_ID"

	AdErrorReasonCANNOT_DISABLE_CALL_CONVERSION_AND_SET_CONVERSION_TYPE_ID AdErrorReason = "CANNOT_DISABLE_CALL_CONVERSION_AND_SET_CONVERSION_TYPE_ID"

	//
	// Cannot set path2 without path1.
	//
	AdErrorReasonCANNOT_SET_PATH2_WITHOUT_PATH1 AdErrorReason = "CANNOT_SET_PATH2_WITHOUT_PATH1"

	//
	// Missing domain name in campaign setting when adding expanded dynamic search ad.
	//
	AdErrorReasonMISSING_DYNAMIC_SEARCH_ADS_SETTING_DOMAIN_NAME AdErrorReason = "MISSING_DYNAMIC_SEARCH_ADS_SETTING_DOMAIN_NAME"

	//
	// An unexpected or unknown error occurred.
	//
	AdErrorReasonUNKNOWN AdErrorReason = "UNKNOWN"
)

//
// The current status of an Ad.
//
type AdGroupAdStatus string

const (

	//
	// Enabled.
	//
	AdGroupAdStatusENABLED AdGroupAdStatus = "ENABLED"

	//
	// Paused.
	//
	AdGroupAdStatusPAUSED AdGroupAdStatus = "PAUSED"

	//
	// Disabled.
	//
	AdGroupAdStatusDISABLED AdGroupAdStatus = "DISABLED"
)

type DeprecatedAdType string

const (

	//
	// Video ad.
	//
	DeprecatedAdTypeVIDEO DeprecatedAdType = "VIDEO"

	//
	// Click to call ad.
	//
	DeprecatedAdTypeCLICK_TO_CALL DeprecatedAdType = "CLICK_TO_CALL"

	//
	// Instream video ad.
	//
	DeprecatedAdTypeIN_STREAM_VIDEO DeprecatedAdType = "IN_STREAM_VIDEO"

	//
	// Froogle ad.
	//
	DeprecatedAdTypeFROOGLE DeprecatedAdType = "FROOGLE"

	//
	// Text link ad.
	//
	DeprecatedAdTypeTEXT_LINK DeprecatedAdType = "TEXT_LINK"

	//
	// Gadget ad.
	//
	DeprecatedAdTypeGADGET DeprecatedAdType = "GADGET"

	//
	// Print ad.
	//
	DeprecatedAdTypePRINT DeprecatedAdType = "PRINT"

	//
	// Wide text ad.
	//
	DeprecatedAdTypeTEXT_WIDE DeprecatedAdType = "TEXT_WIDE"

	//
	// Gadget template ad.
	//
	DeprecatedAdTypeGADGET_TEMPLATE DeprecatedAdType = "GADGET_TEMPLATE"

	//
	// Text ad with video.
	//
	DeprecatedAdTypeTEXT_WITH_VIDEO DeprecatedAdType = "TEXT_WITH_VIDEO"

	//
	// Audio ad.
	//
	DeprecatedAdTypeAUDIO DeprecatedAdType = "AUDIO"

	//
	// Local business ads.
	//
	DeprecatedAdTypeLOCAL_BUSINESS_AD DeprecatedAdType = "LOCAL_BUSINESS_AD"

	//
	// Audio based template ads.
	//
	DeprecatedAdTypeAUDIO_TEMPLATE DeprecatedAdType = "AUDIO_TEMPLATE"

	//
	// Mobile ads
	//
	DeprecatedAdTypeMOBILE_AD DeprecatedAdType = "MOBILE_AD"

	//
	// Mobile image ads
	//
	DeprecatedAdTypeMOBILE_IMAGE_AD DeprecatedAdType = "MOBILE_IMAGE_AD"

	DeprecatedAdTypeUNKNOWN DeprecatedAdType = "UNKNOWN"
)

//
// The reasons for the target error.
//
type AdGroupAdErrorReason string

const (

	//
	// No link found between the adgroup ad and the label.
	//
	AdGroupAdErrorReasonAD_GROUP_AD_LABEL_DOES_NOT_EXIST AdGroupAdErrorReason = "AD_GROUP_AD_LABEL_DOES_NOT_EXIST"

	//
	// The label has already been attached to the adgroup ad.
	//
	AdGroupAdErrorReasonAD_GROUP_AD_LABEL_ALREADY_EXISTS AdGroupAdErrorReason = "AD_GROUP_AD_LABEL_ALREADY_EXISTS"

	//
	// The specified ad was not found in the adgroup
	//
	AdGroupAdErrorReasonAD_NOT_UNDER_ADGROUP AdGroupAdErrorReason = "AD_NOT_UNDER_ADGROUP"

	//
	// Removed ads may not be modified
	//
	AdGroupAdErrorReasonCANNOT_OPERATE_ON_REMOVED_ADGROUPAD AdGroupAdErrorReason = "CANNOT_OPERATE_ON_REMOVED_ADGROUPAD"

	//
	// An ad of this type is deprecated and cannot be created. Only deletions
	// are permitted.
	//
	AdGroupAdErrorReasonCANNOT_CREATE_DEPRECATED_ADS AdGroupAdErrorReason = "CANNOT_CREATE_DEPRECATED_ADS"

	//
	// Text ads are deprecated and cannot be created. Use expanded text ads instead.
	//
	AdGroupAdErrorReasonCANNOT_CREATE_TEXT_ADS AdGroupAdErrorReason = "CANNOT_CREATE_TEXT_ADS"

	//
	// A required field was not specified or is an empty string.
	//
	AdGroupAdErrorReasonEMPTY_FIELD AdGroupAdErrorReason = "EMPTY_FIELD"

	//
	// An ad may only be modified once per call
	//
	AdGroupAdErrorReasonENTITY_REFERENCED_IN_MULTIPLE_OPS AdGroupAdErrorReason = "ENTITY_REFERENCED_IN_MULTIPLE_OPS"

	//
	// The specified operation is not supported.  Only ADD, SET, and REMOVE
	// are supported
	//
	AdGroupAdErrorReasonUNSUPPORTED_OPERATION AdGroupAdErrorReason = "UNSUPPORTED_OPERATION"
)

//
// Reasons for error.
//
type AdSharingErrorReason string

const (

	//
	// Error resulting in attempting to add an Ad to an AdGroup that already contains the Ad.
	//
	AdSharingErrorReasonAD_GROUP_ALREADY_CONTAINS_AD AdSharingErrorReason = "AD_GROUP_ALREADY_CONTAINS_AD"

	//
	// Ad is not compatible with the AdGroup it is being shared with. For more details, look
	// at {@link #sharedAdError}.
	//
	AdSharingErrorReasonINCOMPATIBLE_AD_UNDER_AD_GROUP AdSharingErrorReason = "INCOMPATIBLE_AD_UNDER_AD_GROUP"

	//
	// Cannot add AdGroupAd on inactive Ad.
	//
	AdSharingErrorReasonCANNOT_SHARE_INACTIVE_AD AdSharingErrorReason = "CANNOT_SHARE_INACTIVE_AD"
)

//
// The reasons for the AdX error.
//
type AdxErrorReason string

const (

	//
	// Attempt to use non-AdX feature by AdX customer.
	//
	AdxErrorReasonUNSUPPORTED_FEATURE AdxErrorReason = "UNSUPPORTED_FEATURE"
)

//
// The possible os types for an AppUrl
//
type AppUrlOsType string

const (

	//
	// The Apple IOS operating system,
	//
	AppUrlOsTypeOS_TYPE_IOS AppUrlOsType = "OS_TYPE_IOS"

	//
	// The Android operating system.
	//
	AppUrlOsTypeOS_TYPE_ANDROID AppUrlOsType = "OS_TYPE_ANDROID"

	AppUrlOsTypeUNKNOWN AppUrlOsType = "UNKNOWN"
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
// Serving format setting of this ad.
//
type DisplayAdFormatSetting string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	DisplayAdFormatSettingUNKNOWN DisplayAdFormatSetting = "UNKNOWN"

	//
	// Text, image and native formats
	//
	DisplayAdFormatSettingALL_FORMATS DisplayAdFormatSetting = "ALL_FORMATS"

	//
	// Text and image formats
	//
	DisplayAdFormatSettingNON_NATIVE DisplayAdFormatSetting = "NON_NATIVE"

	//
	// Native format
	//
	DisplayAdFormatSettingNATIVE DisplayAdFormatSetting = "NATIVE"
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
// Feed attribute reference error reasons.
//
type FeedAttributeReferenceErrorReason string

const (

	//
	// A feed referenced by ID has been deleted.
	//
	FeedAttributeReferenceErrorReasonCANNOT_REFERENCE_DELETED_FEED FeedAttributeReferenceErrorReason = "CANNOT_REFERENCE_DELETED_FEED"

	//
	// There is no active feed with the given name.
	//
	FeedAttributeReferenceErrorReasonINVALID_FEED_NAME FeedAttributeReferenceErrorReason = "INVALID_FEED_NAME"

	//
	// There is no feed attribute in an active feed with the given name.
	//
	FeedAttributeReferenceErrorReasonINVALID_FEED_ATTRIBUTE_NAME FeedAttributeReferenceErrorReason = "INVALID_FEED_ATTRIBUTE_NAME"
)

//
// The reason for the error.
//
type ForwardCompatibilityErrorReason string

const (

	//
	// Invalid value specified for a key in the forward compatibility map.
	//
	ForwardCompatibilityErrorReasonINVALID_FORWARD_COMPATIBILITY_MAP_VALUE ForwardCompatibilityErrorReason = "INVALID_FORWARD_COMPATIBILITY_MAP_VALUE"

	ForwardCompatibilityErrorReasonUNKNOWN ForwardCompatibilityErrorReason = "UNKNOWN"
)

//
// The reasons for the target error.
//
type FunctionErrorReason string

const (

	//
	// The format of the function is not recognized as a supported function format.
	//
	FunctionErrorReasonINVALID_FUNCTION_FORMAT FunctionErrorReason = "INVALID_FUNCTION_FORMAT"

	//
	// Operand data types do not match.
	//
	FunctionErrorReasonDATA_TYPE_MISMATCH FunctionErrorReason = "DATA_TYPE_MISMATCH"

	//
	// The operands cannot be used together in a conjunction.
	//
	FunctionErrorReasonINVALID_CONJUNCTION_OPERANDS FunctionErrorReason = "INVALID_CONJUNCTION_OPERANDS"

	//
	// Invalid numer of Operands.
	//
	FunctionErrorReasonINVALID_NUMBER_OF_OPERANDS FunctionErrorReason = "INVALID_NUMBER_OF_OPERANDS"

	//
	// Operand Type not supported.
	//
	FunctionErrorReasonINVALID_OPERAND_TYPE FunctionErrorReason = "INVALID_OPERAND_TYPE"

	//
	// Operator not supported.
	//
	FunctionErrorReasonINVALID_OPERATOR FunctionErrorReason = "INVALID_OPERATOR"

	//
	// Request context type not supported.
	//
	FunctionErrorReasonINVALID_REQUEST_CONTEXT_TYPE FunctionErrorReason = "INVALID_REQUEST_CONTEXT_TYPE"

	//
	// The matching function is not allowed for call placeholders
	//
	FunctionErrorReasonINVALID_FUNCTION_FOR_CALL_PLACEHOLDER FunctionErrorReason = "INVALID_FUNCTION_FOR_CALL_PLACEHOLDER"

	//
	// The matching function is not allowed for the specified placeholder
	//
	FunctionErrorReasonINVALID_FUNCTION_FOR_PLACEHOLDER FunctionErrorReason = "INVALID_FUNCTION_FOR_PLACEHOLDER"

	//
	// Invalid operand.
	//
	FunctionErrorReasonINVALID_OPERAND FunctionErrorReason = "INVALID_OPERAND"

	//
	// Missing value for the constant operand.
	//
	FunctionErrorReasonMISSING_CONSTANT_OPERAND_VALUE FunctionErrorReason = "MISSING_CONSTANT_OPERAND_VALUE"

	//
	// The value of the constant operand is invalid.
	//
	FunctionErrorReasonINVALID_CONSTANT_OPERAND_VALUE FunctionErrorReason = "INVALID_CONSTANT_OPERAND_VALUE"

	//
	// Invalid function nesting.
	//
	FunctionErrorReasonINVALID_NESTING FunctionErrorReason = "INVALID_NESTING"

	//
	// The Feed ID was different from another Feed ID in the same function.
	//
	FunctionErrorReasonMULTIPLE_FEED_IDS_NOT_SUPPORTED FunctionErrorReason = "MULTIPLE_FEED_IDS_NOT_SUPPORTED"

	//
	// The matching function is invalid for use with a feed with a fixed schema.
	//
	FunctionErrorReasonINVALID_FUNCTION_FOR_FEED_WITH_FIXED_SCHEMA FunctionErrorReason = "INVALID_FUNCTION_FOR_FEED_WITH_FIXED_SCHEMA"

	//
	// Invalid attribute name.
	//
	FunctionErrorReasonINVALID_ATTRIBUTE_NAME FunctionErrorReason = "INVALID_ATTRIBUTE_NAME"

	FunctionErrorReasonUNKNOWN FunctionErrorReason = "UNKNOWN"
)

//
// Function parsing error reason.
//
type FunctionParsingErrorReason string

const (

	//
	// Unexpected end of function string.
	//
	FunctionParsingErrorReasonNO_MORE_INPUT FunctionParsingErrorReason = "NO_MORE_INPUT"

	//
	// Could not find an expected character.
	//
	FunctionParsingErrorReasonEXPECTED_CHARACTER FunctionParsingErrorReason = "EXPECTED_CHARACTER"

	//
	// Unexpected separator character.
	//
	FunctionParsingErrorReasonUNEXPECTED_SEPARATOR FunctionParsingErrorReason = "UNEXPECTED_SEPARATOR"

	//
	// Unmatched left bracket or parenthesis.
	//
	FunctionParsingErrorReasonUNMATCHED_LEFT_BRACKET FunctionParsingErrorReason = "UNMATCHED_LEFT_BRACKET"

	//
	// Unmatched right bracket or parenthesis.
	//
	FunctionParsingErrorReasonUNMATCHED_RIGHT_BRACKET FunctionParsingErrorReason = "UNMATCHED_RIGHT_BRACKET"

	//
	// Functions are nested too deeply.
	//
	FunctionParsingErrorReasonTOO_MANY_NESTED_FUNCTIONS FunctionParsingErrorReason = "TOO_MANY_NESTED_FUNCTIONS"

	//
	// Missing right-hand-side operand.
	//
	FunctionParsingErrorReasonMISSING_RIGHT_HAND_OPERAND FunctionParsingErrorReason = "MISSING_RIGHT_HAND_OPERAND"

	//
	// Invalid operator/function name.
	//
	FunctionParsingErrorReasonINVALID_OPERATOR_NAME FunctionParsingErrorReason = "INVALID_OPERATOR_NAME"

	//
	// Feed attribute operand's argument is not an integer.
	//
	FunctionParsingErrorReasonFEED_ATTRIBUTE_OPERAND_ARGUMENT_NOT_INTEGER FunctionParsingErrorReason = "FEED_ATTRIBUTE_OPERAND_ARGUMENT_NOT_INTEGER"

	//
	// Missing function operands.
	//
	FunctionParsingErrorReasonNO_OPERANDS FunctionParsingErrorReason = "NO_OPERANDS"

	//
	// Function had too many operands.
	//
	FunctionParsingErrorReasonTOO_MANY_OPERANDS FunctionParsingErrorReason = "TOO_MANY_OPERANDS"

	FunctionParsingErrorReasonUNKNOWN FunctionParsingErrorReason = "UNKNOWN"
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

type ImageErrorReason string

const (

	//
	// The image is not valid.
	//
	ImageErrorReasonINVALID_IMAGE ImageErrorReason = "INVALID_IMAGE"

	//
	// The image could not be stored.
	//
	ImageErrorReasonSTORAGE_ERROR ImageErrorReason = "STORAGE_ERROR"

	//
	// There was a problem with the request.
	//
	ImageErrorReasonBAD_REQUEST ImageErrorReason = "BAD_REQUEST"

	//
	// The image is not of legal dimensions.
	//
	ImageErrorReasonUNEXPECTED_SIZE ImageErrorReason = "UNEXPECTED_SIZE"

	//
	// Animated image are not permitted.
	//
	ImageErrorReasonANIMATED_NOT_ALLOWED ImageErrorReason = "ANIMATED_NOT_ALLOWED"

	//
	// Animation is too long.
	//
	ImageErrorReasonANIMATION_TOO_LONG ImageErrorReason = "ANIMATION_TOO_LONG"

	//
	// There was an error on the server.
	//
	ImageErrorReasonSERVER_ERROR ImageErrorReason = "SERVER_ERROR"

	//
	// Image cannot be in CMYK color format.
	//
	ImageErrorReasonCMYK_JPEG_NOT_ALLOWED ImageErrorReason = "CMYK_JPEG_NOT_ALLOWED"

	//
	// Flash images are not permitted.
	//
	ImageErrorReasonFLASH_NOT_ALLOWED ImageErrorReason = "FLASH_NOT_ALLOWED"

	//
	// Flash images must support clickTag.
	//
	ImageErrorReasonFLASH_WITHOUT_CLICKTAG ImageErrorReason = "FLASH_WITHOUT_CLICKTAG"

	//
	// A flash error has occurred after fixing the click tag.
	//
	ImageErrorReasonFLASH_ERROR_AFTER_FIXING_CLICK_TAG ImageErrorReason = "FLASH_ERROR_AFTER_FIXING_CLICK_TAG"

	//
	// Unacceptable visual effects.
	//
	ImageErrorReasonANIMATED_VISUAL_EFFECT ImageErrorReason = "ANIMATED_VISUAL_EFFECT"

	//
	// There was a problem with the flash image.
	//
	ImageErrorReasonFLASH_ERROR ImageErrorReason = "FLASH_ERROR"

	//
	// Incorrect image layout.
	//
	ImageErrorReasonLAYOUT_PROBLEM ImageErrorReason = "LAYOUT_PROBLEM"

	//
	// There was a problem reading the image file.
	//
	ImageErrorReasonPROBLEM_READING_IMAGE_FILE ImageErrorReason = "PROBLEM_READING_IMAGE_FILE"

	//
	// There was an error storing the image.
	//
	ImageErrorReasonERROR_STORING_IMAGE ImageErrorReason = "ERROR_STORING_IMAGE"

	//
	// The aspect ratio of the image is not allowed.
	//
	ImageErrorReasonASPECT_RATIO_NOT_ALLOWED ImageErrorReason = "ASPECT_RATIO_NOT_ALLOWED"

	//
	// Flash cannot have network objects.
	//
	ImageErrorReasonFLASH_HAS_NETWORK_OBJECTS ImageErrorReason = "FLASH_HAS_NETWORK_OBJECTS"

	//
	// Flash cannot have network methods.
	//
	ImageErrorReasonFLASH_HAS_NETWORK_METHODS ImageErrorReason = "FLASH_HAS_NETWORK_METHODS"

	//
	// Flash cannot have a Url.
	//
	ImageErrorReasonFLASH_HAS_URL ImageErrorReason = "FLASH_HAS_URL"

	//
	// Flash cannot use mouse tracking.
	//
	ImageErrorReasonFLASH_HAS_MOUSE_TRACKING ImageErrorReason = "FLASH_HAS_MOUSE_TRACKING"

	//
	// Flash cannot have a random number.
	//
	ImageErrorReasonFLASH_HAS_RANDOM_NUM ImageErrorReason = "FLASH_HAS_RANDOM_NUM"

	//
	// Ad click target cannot be '_self'.
	//
	ImageErrorReasonFLASH_SELF_TARGETS ImageErrorReason = "FLASH_SELF_TARGETS"

	//
	// GetUrl method should only use '_blank'.
	//
	ImageErrorReasonFLASH_BAD_GETURL_TARGET ImageErrorReason = "FLASH_BAD_GETURL_TARGET"

	//
	// Flash version is not supported.
	//
	ImageErrorReasonFLASH_VERSION_NOT_SUPPORTED ImageErrorReason = "FLASH_VERSION_NOT_SUPPORTED"

	//
	// Flash movies need to have hard coded click URL or clickTAG
	//
	ImageErrorReasonFLASH_WITHOUT_HARD_CODED_CLICK_URL ImageErrorReason = "FLASH_WITHOUT_HARD_CODED_CLICK_URL"

	//
	// Uploaded flash file is corrupted.
	//
	ImageErrorReasonINVALID_FLASH_FILE ImageErrorReason = "INVALID_FLASH_FILE"

	//
	// Uploaded flash file can be parsed, but the click tag can not be fixed properly.
	//
	ImageErrorReasonFAILED_TO_FIX_CLICK_TAG_IN_FLASH ImageErrorReason = "FAILED_TO_FIX_CLICK_TAG_IN_FLASH"

	//
	// Flash movie accesses network resources
	//
	ImageErrorReasonFLASH_ACCESSES_NETWORK_RESOURCES ImageErrorReason = "FLASH_ACCESSES_NETWORK_RESOURCES"

	//
	// Flash movie attempts to call external javascript code
	//
	ImageErrorReasonFLASH_EXTERNAL_JS_CALL ImageErrorReason = "FLASH_EXTERNAL_JS_CALL"

	//
	// Flash movie attempts to call flash system commands
	//
	ImageErrorReasonFLASH_EXTERNAL_FS_CALL ImageErrorReason = "FLASH_EXTERNAL_FS_CALL"

	//
	// Image file is too large.
	//
	ImageErrorReasonFILE_TOO_LARGE ImageErrorReason = "FILE_TOO_LARGE"

	//
	// Image data is too large.
	//
	ImageErrorReasonIMAGE_DATA_TOO_LARGE ImageErrorReason = "IMAGE_DATA_TOO_LARGE"

	//
	// Error while processing the image.
	//
	ImageErrorReasonIMAGE_PROCESSING_ERROR ImageErrorReason = "IMAGE_PROCESSING_ERROR"

	//
	// Image is too small.
	//
	ImageErrorReasonIMAGE_TOO_SMALL ImageErrorReason = "IMAGE_TOO_SMALL"

	//
	// Input was invalid.
	//
	ImageErrorReasonINVALID_INPUT ImageErrorReason = "INVALID_INPUT"

	//
	// There was a problem reading the image file.
	//
	ImageErrorReasonPROBLEM_READING_FILE ImageErrorReason = "PROBLEM_READING_FILE"
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

type LabelStatus string

const (

	//
	// The label is enabled.
	//
	LabelStatusENABLED LabelStatus = "ENABLED"

	//
	// The label has been removed.
	//
	LabelStatusREMOVED LabelStatus = "REMOVED"

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	LabelStatusUNKNOWN LabelStatus = "UNKNOWN"
)

//
// Media types
//
type MediaMediaType string

const (

	//
	// Audio file.
	//
	MediaMediaTypeAUDIO MediaMediaType = "AUDIO"

	//
	// Animated image, such as animated GIF.
	//
	MediaMediaTypeDYNAMIC_IMAGE MediaMediaType = "DYNAMIC_IMAGE"

	//
	// Small image; used for map ad.
	//
	MediaMediaTypeICON MediaMediaType = "ICON"

	//
	// Static image; for image ad.
	//
	MediaMediaTypeIMAGE MediaMediaType = "IMAGE"

	//
	// Predefined standard icon; used for map ads.
	//
	MediaMediaTypeSTANDARD_ICON MediaMediaType = "STANDARD_ICON"

	//
	// Video file.
	//
	MediaMediaTypeVIDEO MediaMediaType = "VIDEO"

	//
	// ZIP file; used in fields of template ads.
	//
	MediaMediaTypeMEDIA_BUNDLE MediaMediaType = "MEDIA_BUNDLE"
)

//
// Mime types
//
type MediaMimeType string

const (

	//
	// MIME type of image/jpeg
	//
	MediaMimeTypeIMAGE_JPEG MediaMimeType = "IMAGE_JPEG"

	//
	// MIME type of image/gif
	//
	MediaMimeTypeIMAGE_GIF MediaMimeType = "IMAGE_GIF"

	//
	// MIME type of image/png
	//
	MediaMimeTypeIMAGE_PNG MediaMimeType = "IMAGE_PNG"

	//
	// MIME type of application/x-shockwave-flash
	//
	MediaMimeTypeFLASH MediaMimeType = "FLASH"

	//
	// MIME type of text/html
	//
	MediaMimeTypeTEXT_HTML MediaMimeType = "TEXT_HTML"

	//
	// MIME type of application/pdf
	//
	MediaMimeTypePDF MediaMimeType = "PDF"

	//
	// MIME type of application/msword
	//
	MediaMimeTypeMSWORD MediaMimeType = "MSWORD"

	//
	// MIME type of application/vnd.ms-excel
	//
	MediaMimeTypeMSEXCEL MediaMimeType = "MSEXCEL"

	//
	// MIME type of application/rtf
	//
	MediaMimeTypeRTF MediaMimeType = "RTF"

	//
	// MIME type of audio/wav
	//
	MediaMimeTypeAUDIO_WAV MediaMimeType = "AUDIO_WAV"

	//
	// MIME type of audio/mp3
	//
	MediaMimeTypeAUDIO_MP3 MediaMimeType = "AUDIO_MP3"

	//
	// MIME type of application/x-html5-ad-zip
	//
	MediaMimeTypeHTML5_AD_ZIP MediaMimeType = "HTML5_AD_ZIP"
)

//
// Sizes for retrieving the original media
//
type MediaSize string

const (

	//
	// Full size of Media.
	//
	MediaSizeFULL MediaSize = "FULL"

	//
	// Shunken size of media.
	//
	MediaSizeSHRUNKEN MediaSize = "SHRUNKEN"

	//
	// Preview size of media.
	//
	MediaSizePREVIEW MediaSize = "PREVIEW"

	//
	// Video thumbnail size of Media.
	//
	MediaSizeVIDEO_THUMBNAIL MediaSize = "VIDEO_THUMBNAIL"
)

//
// Enumeration of the reasons for the {@link MediaBundleError}
//
type MediaBundleErrorReason string

const (

	//
	// The entryPoint field cannot be set using the <code>MediaService</code>.
	//
	MediaBundleErrorReasonENTRY_POINT_CANNOT_BE_SET_USING_MEDIA_SERVICE MediaBundleErrorReason = "ENTRY_POINT_CANNOT_BE_SET_USING_MEDIA_SERVICE"

	//
	// There was a problem with the request.
	//
	MediaBundleErrorReasonBAD_REQUEST MediaBundleErrorReason = "BAD_REQUEST"

	//
	// HTML5 ads using DoubleClick Studio created ZIP files are not supported.
	//
	MediaBundleErrorReasonDOUBLECLICK_BUNDLE_NOT_ALLOWED MediaBundleErrorReason = "DOUBLECLICK_BUNDLE_NOT_ALLOWED"

	//
	// Cannot reference URL external to the media bundle.
	//
	MediaBundleErrorReasonEXTERNAL_URL_NOT_ALLOWED MediaBundleErrorReason = "EXTERNAL_URL_NOT_ALLOWED"

	//
	// Media bundle file is too large.
	//
	MediaBundleErrorReasonFILE_TOO_LARGE MediaBundleErrorReason = "FILE_TOO_LARGE"

	//
	// ZIP file from Google Web Designer is not published.
	//
	MediaBundleErrorReasonGOOGLE_WEB_DESIGNER_ZIP_FILE_NOT_PUBLISHED MediaBundleErrorReason = "GOOGLE_WEB_DESIGNER_ZIP_FILE_NOT_PUBLISHED"

	//
	// Input was invalid.
	//
	MediaBundleErrorReasonINVALID_INPUT MediaBundleErrorReason = "INVALID_INPUT"

	//
	// There was a problem with the media bundle.
	//
	MediaBundleErrorReasonINVALID_MEDIA_BUNDLE MediaBundleErrorReason = "INVALID_MEDIA_BUNDLE"

	//
	// There was a problem with one or more of the media bundle entries.
	//
	MediaBundleErrorReasonINVALID_MEDIA_BUNDLE_ENTRY MediaBundleErrorReason = "INVALID_MEDIA_BUNDLE_ENTRY"

	//
	// The media bundle contains a file with an unknown mime type
	//
	MediaBundleErrorReasonINVALID_MIME_TYPE MediaBundleErrorReason = "INVALID_MIME_TYPE"

	//
	// The media bundle contain an invalid asset path.
	//
	MediaBundleErrorReasonINVALID_PATH MediaBundleErrorReason = "INVALID_PATH"

	//
	// HTML5 ad is trying to reference an asset not in .ZIP file
	//
	MediaBundleErrorReasonINVALID_URL_REFERENCE MediaBundleErrorReason = "INVALID_URL_REFERENCE"

	//
	// Media data is too large.
	//
	MediaBundleErrorReasonMEDIA_DATA_TOO_LARGE MediaBundleErrorReason = "MEDIA_DATA_TOO_LARGE"

	//
	// The media bundle contains no primary entry.
	//
	MediaBundleErrorReasonMISSING_PRIMARY_MEDIA_BUNDLE_ENTRY MediaBundleErrorReason = "MISSING_PRIMARY_MEDIA_BUNDLE_ENTRY"

	//
	// There was an error on the server.
	//
	MediaBundleErrorReasonSERVER_ERROR MediaBundleErrorReason = "SERVER_ERROR"

	//
	// The image could not be stored.
	//
	MediaBundleErrorReasonSTORAGE_ERROR MediaBundleErrorReason = "STORAGE_ERROR"

	//
	// Media bundle created with the Swiffy tool is not allowed.
	//
	MediaBundleErrorReasonSWIFFY_BUNDLE_NOT_ALLOWED MediaBundleErrorReason = "SWIFFY_BUNDLE_NOT_ALLOWED"

	//
	// The media bundle contains too many files.
	//
	MediaBundleErrorReasonTOO_MANY_FILES MediaBundleErrorReason = "TOO_MANY_FILES"

	//
	// The media bundle is not of legal dimensions.
	//
	MediaBundleErrorReasonUNEXPECTED_SIZE MediaBundleErrorReason = "UNEXPECTED_SIZE"

	//
	// Google Web Designer not created for "AdWords" environment.
	//
	MediaBundleErrorReasonUNSUPPORTED_GOOGLE_WEB_DESIGNER_ENVIRONMENT MediaBundleErrorReason = "UNSUPPORTED_GOOGLE_WEB_DESIGNER_ENVIRONMENT"

	//
	// Unsupported HTML5 feature in HTML5 asset.
	//
	MediaBundleErrorReasonUNSUPPORTED_HTML5_FEATURE MediaBundleErrorReason = "UNSUPPORTED_HTML5_FEATURE"

	//
	// URL in HTML5 entry is not ssl compliant.
	//
	MediaBundleErrorReasonURL_IN_MEDIA_BUNDLE_NOT_SSL_COMPLIANT MediaBundleErrorReason = "URL_IN_MEDIA_BUNDLE_NOT_SSL_COMPLIANT"
)

//
// The reasons for the target error.
//
type MediaErrorReason string

const (

	//
	// Cannot add a standard icon type
	//
	MediaErrorReasonCANNOT_ADD_STANDARD_ICON MediaErrorReason = "CANNOT_ADD_STANDARD_ICON"

	//
	// May only select Standard Icons alone
	//
	MediaErrorReasonCANNOT_SELECT_STANDARD_ICON_WITH_OTHER_TYPES MediaErrorReason = "CANNOT_SELECT_STANDARD_ICON_WITH_OTHER_TYPES"

	//
	// Image contains both a media ID and media data.
	//
	MediaErrorReasonCANNOT_SPECIFY_MEDIA_ID_AND_DATA MediaErrorReason = "CANNOT_SPECIFY_MEDIA_ID_AND_DATA"

	//
	// A media with given type and reference id already exists
	//
	MediaErrorReasonDUPLICATE_MEDIA MediaErrorReason = "DUPLICATE_MEDIA"

	//
	// A required field was not specified or is an empty string.
	//
	MediaErrorReasonEMPTY_FIELD MediaErrorReason = "EMPTY_FIELD"

	//
	// A media may only be modified once per call
	//
	MediaErrorReasonENTITY_REFERENCED_IN_MULTIPLE_OPS MediaErrorReason = "ENTITY_REFERENCED_IN_MULTIPLE_OPS"

	//
	// Field is not supported for the media sub type.
	//
	MediaErrorReasonFIELD_NOT_SUPPORTED_FOR_MEDIA_SUB_TYPE MediaErrorReason = "FIELD_NOT_SUPPORTED_FOR_MEDIA_SUB_TYPE"

	//
	// The media id is invalid
	//
	MediaErrorReasonINVALID_MEDIA_ID MediaErrorReason = "INVALID_MEDIA_ID"

	//
	// The media subtype is invalid
	//
	MediaErrorReasonINVALID_MEDIA_SUB_TYPE MediaErrorReason = "INVALID_MEDIA_SUB_TYPE"

	//
	// The media type is invalid
	//
	MediaErrorReasonINVALID_MEDIA_TYPE MediaErrorReason = "INVALID_MEDIA_TYPE"

	//
	// The mimetype is invalid
	//
	MediaErrorReasonINVALID_MIME_TYPE MediaErrorReason = "INVALID_MIME_TYPE"

	//
	// The media reference id is invalid
	//
	MediaErrorReasonINVALID_REFERENCE_ID MediaErrorReason = "INVALID_REFERENCE_ID"

	//
	// The YouTube video id is invalid
	//
	MediaErrorReasonINVALID_YOU_TUBE_ID MediaErrorReason = "INVALID_YOU_TUBE_ID"

	//
	// Media has failed transcoding
	//
	MediaErrorReasonMEDIA_FAILED_TRANSCODING MediaErrorReason = "MEDIA_FAILED_TRANSCODING"

	//
	// Media has not been transcoded
	//
	MediaErrorReasonMEDIA_NOT_TRANSCODED MediaErrorReason = "MEDIA_NOT_TRANSCODED"

	//
	// The MediaType does not match the actual media object's type
	//
	MediaErrorReasonMEDIA_TYPE_DOES_NOT_MATCH_OBJECT_TYPE MediaErrorReason = "MEDIA_TYPE_DOES_NOT_MATCH_OBJECT_TYPE"

	//
	// None of the fields have been specified.
	//
	MediaErrorReasonNO_FIELDS_SPECIFIED MediaErrorReason = "NO_FIELDS_SPECIFIED"

	//
	// One of reference Id or media Id must be specified
	//
	MediaErrorReasonNULL_REFERENCE_ID_AND_MEDIA_ID MediaErrorReason = "NULL_REFERENCE_ID_AND_MEDIA_ID"

	//
	// The string has too many characters.
	//
	MediaErrorReasonTOO_LONG MediaErrorReason = "TOO_LONG"

	//
	// The specified operation is not supported.  Only ADD, SET, and REMOVE
	// are supported
	//
	MediaErrorReasonUNSUPPORTED_OPERATION MediaErrorReason = "UNSUPPORTED_OPERATION"

	//
	// The specified type is not supported.
	//
	MediaErrorReasonUNSUPPORTED_TYPE MediaErrorReason = "UNSUPPORTED_TYPE"

	//
	// YouTube is unavailable for requesting video data.
	//
	MediaErrorReasonYOU_TUBE_SERVICE_UNAVAILABLE MediaErrorReason = "YOU_TUBE_SERVICE_UNAVAILABLE"

	//
	// The YouTube video has a non positive duration.
	//
	MediaErrorReasonYOU_TUBE_VIDEO_HAS_NON_POSITIVE_DURATION MediaErrorReason = "YOU_TUBE_VIDEO_HAS_NON_POSITIVE_DURATION"

	//
	// The YouTube video id is syntactically valid but the video was not found.
	//
	MediaErrorReasonYOU_TUBE_VIDEO_NOT_FOUND MediaErrorReason = "YOU_TUBE_VIDEO_NOT_FOUND"
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
// The reasons for errors when using pagination.
//
type PagingErrorReason string

const (

	//
	// The start index value cannot be a negative number.
	//
	PagingErrorReasonSTART_INDEX_CANNOT_BE_NEGATIVE PagingErrorReason = "START_INDEX_CANNOT_BE_NEGATIVE"

	//
	// The number of results cannot be a negative number.
	//
	PagingErrorReasonNUMBER_OF_RESULTS_CANNOT_BE_NEGATIVE PagingErrorReason = "NUMBER_OF_RESULTS_CANNOT_BE_NEGATIVE"

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	PagingErrorReasonUNKNOWN PagingErrorReason = "UNKNOWN"
)

//
// Policy approval status.
//
type PolicyApprovalStatus string

const (
	PolicyApprovalStatusUNKNOWN PolicyApprovalStatus = "UNKNOWN"

	PolicyApprovalStatusAPPROVED PolicyApprovalStatus = "APPROVED"

	PolicyApprovalStatusAPPROVED_LIMITED PolicyApprovalStatus = "APPROVED_LIMITED"

	PolicyApprovalStatusELIGIBLE PolicyApprovalStatus = "ELIGIBLE"

	PolicyApprovalStatusUNDER_REVIEW PolicyApprovalStatus = "UNDER_REVIEW"

	PolicyApprovalStatusDISAPPROVED PolicyApprovalStatus = "DISAPPROVED"

	PolicyApprovalStatusSITE_SUSPENDED PolicyApprovalStatus = "SITE_SUSPENDED"
)

//
// The denormalized status of a reviewable, calculated based on the status of its
// individual policy entries.
//
type PolicySummaryDenormalizedStatus string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	PolicySummaryDenormalizedStatusUNKNOWN PolicySummaryDenormalizedStatus = "UNKNOWN"

	PolicySummaryDenormalizedStatusDISAPPROVED PolicySummaryDenormalizedStatus = "DISAPPROVED"

	PolicySummaryDenormalizedStatusAPPROVED_LIMITED PolicySummaryDenormalizedStatus = "APPROVED_LIMITED"

	PolicySummaryDenormalizedStatusAPPROVED PolicySummaryDenormalizedStatus = "APPROVED"
)

//
// The review state of a reviewable.
//
type PolicySummaryReviewState string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	PolicySummaryReviewStateUNKNOWN PolicySummaryReviewState = "UNKNOWN"

	PolicySummaryReviewStateREVIEW_IN_PROGRESS PolicySummaryReviewState = "REVIEW_IN_PROGRESS"

	PolicySummaryReviewStateREVIEWED PolicySummaryReviewState = "REVIEWED"

	//
	// Appeal requested for this reviewable.
	//
	PolicySummaryReviewStateUNDER_APPEAL PolicySummaryReviewState = "UNDER_APPEAL"
)

//
// Subtype of PolicyTopicConstraint.
//
type PolicyTopicConstraintPolicyTopicConstraintType string

const (
	PolicyTopicConstraintPolicyTopicConstraintTypeUNKNOWN PolicyTopicConstraintPolicyTopicConstraintType = "UNKNOWN"

	PolicyTopicConstraintPolicyTopicConstraintTypeCOUNTRY PolicyTopicConstraintPolicyTopicConstraintType = "COUNTRY"

	PolicyTopicConstraintPolicyTopicConstraintTypeRESELLER PolicyTopicConstraintPolicyTopicConstraintType = "RESELLER"

	PolicyTopicConstraintPolicyTopicConstraintTypeCERTIFICATE_MISSING_IN_COUNTRY PolicyTopicConstraintPolicyTopicConstraintType = "CERTIFICATE_MISSING_IN_COUNTRY"

	PolicyTopicConstraintPolicyTopicConstraintTypeCERTIFICATE_DOMAIN_MISMATCH_IN_COUNTRY PolicyTopicConstraintPolicyTopicConstraintType = "CERTIFICATE_DOMAIN_MISMATCH_IN_COUNTRY"

	PolicyTopicConstraintPolicyTopicConstraintTypeCERTIFICATE_MISSING PolicyTopicConstraintPolicyTopicConstraintType = "CERTIFICATE_MISSING"
)

//
// The summarized nature of a policy entry.
//
type PolicyTopicEntryType string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	PolicyTopicEntryTypeUNKNOWN PolicyTopicEntryType = "UNKNOWN"

	//
	// Will never serve.
	//
	PolicyTopicEntryTypePROHIBITED PolicyTopicEntryType = "PROHIBITED"

	//
	// Constrained for at least one value.
	//
	PolicyTopicEntryTypeLIMITED PolicyTopicEntryType = "LIMITED"
)

//
// Describes the type of evidence inside the policy topic evidence.
//
type PolicyTopicEvidenceType string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	PolicyTopicEvidenceTypeUNKNOWN PolicyTopicEvidenceType = "UNKNOWN"

	//
	// Evidence found in the text of the ad.
	//
	PolicyTopicEvidenceTypeAD_TEXT PolicyTopicEvidenceType = "AD_TEXT"

	//
	// HTTP code returned when the final URL was crawled.
	//
	PolicyTopicEvidenceTypeHTTP_CODE PolicyTopicEvidenceType = "HTTP_CODE"

	//
	// List of websites linked with this ad.
	//
	PolicyTopicEvidenceTypeWEBSITES PolicyTopicEvidenceType = "WEBSITES"

	//
	// The language the ad was detected to be written in.
	//
	PolicyTopicEvidenceTypeLANGUAGE PolicyTopicEvidenceType = "LANGUAGE"

	//
	// The text in the destination of the ad that is causing a policy violation.
	//
	PolicyTopicEvidenceTypeDESTINATION_TEXT_LIST PolicyTopicEvidenceType = "DESTINATION_TEXT_LIST"
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
// A set of attributes that describe the rich media ad capabilities.
//
type RichMediaAdAdAttribute string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	RichMediaAdAdAttributeUNKNOWN RichMediaAdAdAttribute = "UNKNOWN"

	//
	// Indicates that the ad supports mouse roll over to expand.
	//
	RichMediaAdAdAttributeROLL_OVER_TO_EXPAND RichMediaAdAdAttribute = "ROLL_OVER_TO_EXPAND"

	//
	// Indicates that the ad supports SSL.
	//
	RichMediaAdAdAttributeSSL RichMediaAdAdAttribute = "SSL"
)

//
// Different types of rich media ad that are available to customers.
//
type RichMediaAdRichMediaAdType string

const (

	//
	// Standard.
	//
	RichMediaAdRichMediaAdTypeSTANDARD RichMediaAdRichMediaAdType = "STANDARD"

	//
	// In stream video ad.
	//
	RichMediaAdRichMediaAdTypeIN_STREAM_VIDEO RichMediaAdRichMediaAdType = "IN_STREAM_VIDEO"
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
// The reasons for errors when querying for stats.
//
type StatsQueryErrorReason string

const (

	//
	// Date is outside of allowed range.
	//
	StatsQueryErrorReasonDATE_NOT_IN_VALID_RANGE StatsQueryErrorReason = "DATE_NOT_IN_VALID_RANGE"
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
// Indicates the source of a system-generated entity.
//
type SystemManagedEntitySource string

const (

	//
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	SystemManagedEntitySourceUNKNOWN SystemManagedEntitySource = "UNKNOWN"

	//
	// Generated ad variations experiment ad.
	//
	SystemManagedEntitySourceAD_VARIATIONS SystemManagedEntitySource = "AD_VARIATIONS"
)

//
// Possible field types of template element fields.
//
type TemplateElementFieldType string

const (

	//
	// Address field type (text).
	//
	TemplateElementFieldTypeADDRESS TemplateElementFieldType = "ADDRESS"

	//
	// Audio field type (Media).
	//
	TemplateElementFieldTypeAUDIO TemplateElementFieldType = "AUDIO"

	//
	// Enum field type (text).
	//
	TemplateElementFieldTypeENUM TemplateElementFieldType = "ENUM"

	//
	// Image field type (Media).
	//
	TemplateElementFieldTypeIMAGE TemplateElementFieldType = "IMAGE"

	//
	// Background Image field type (Media).
	//
	TemplateElementFieldTypeBACKGROUND_IMAGE TemplateElementFieldType = "BACKGROUND_IMAGE"

	//
	// Number field type (text).
	//
	TemplateElementFieldTypeNUMBER TemplateElementFieldType = "NUMBER"

	//
	// Text field type (text).
	//
	TemplateElementFieldTypeTEXT TemplateElementFieldType = "TEXT"

	//
	// URL field type (text).
	//
	TemplateElementFieldTypeURL TemplateElementFieldType = "URL"

	//
	// Video field type (Media).
	//
	TemplateElementFieldTypeVIDEO TemplateElementFieldType = "VIDEO"

	//
	// Visible URL field type (text).
	//
	TemplateElementFieldTypeVISIBLE_URL TemplateElementFieldType = "VISIBLE_URL"

	//
	// A ZIP file containing HTML5 assets.
	//
	TemplateElementFieldTypeMEDIA_BUNDLE TemplateElementFieldType = "MEDIA_BUNDLE"

	//
	// UNKNOWN type can not be passed as input.
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	TemplateElementFieldTypeUNKNOWN TemplateElementFieldType = "UNKNOWN"
)

//
// Allowed expanding directions for ads that are expandable.
//
type ThirdPartyRedirectAdExpandingDirection string

const (

	//
	// Whether the ad can be expanded is unknown.
	// <span class="constraint Rejected">Used for return value only. An enumeration could not be processed, typically due to incompatibility with your WSDL version.</span>
	//
	ThirdPartyRedirectAdExpandingDirectionUNKNOWN ThirdPartyRedirectAdExpandingDirection = "UNKNOWN"

	//
	// The ad is allowed to expand upward.
	//
	ThirdPartyRedirectAdExpandingDirectionEXPANDING_UP ThirdPartyRedirectAdExpandingDirection = "EXPANDING_UP"

	//
	// The ad is allowed to expand downward.
	//
	ThirdPartyRedirectAdExpandingDirectionEXPANDING_DOWN ThirdPartyRedirectAdExpandingDirection = "EXPANDING_DOWN"

	//
	// The ad is allowed to expand leftward.
	//
	ThirdPartyRedirectAdExpandingDirectionEXPANDING_LEFT ThirdPartyRedirectAdExpandingDirection = "EXPANDING_LEFT"

	//
	// The ad is allowed to expand rightward.
	//
	ThirdPartyRedirectAdExpandingDirectionEXPANDING_RIGHT ThirdPartyRedirectAdExpandingDirection = "EXPANDING_RIGHT"

	//
	// The ad is allowed to expand toward the up-left corner.
	//
	ThirdPartyRedirectAdExpandingDirectionEXPANDING_UPLEFT ThirdPartyRedirectAdExpandingDirection = "EXPANDING_UPLEFT"

	//
	// The ad is allowed to expand toward the up-right corner.
	//
	ThirdPartyRedirectAdExpandingDirectionEXPANDING_UPRIGHT ThirdPartyRedirectAdExpandingDirection = "EXPANDING_UPRIGHT"

	//
	// The ad is allowed to expand toward the down-left corner.
	//
	ThirdPartyRedirectAdExpandingDirectionEXPANDING_DOWNLEFT ThirdPartyRedirectAdExpandingDirection = "EXPANDING_DOWNLEFT"

	//
	// The ad is allowed to expand toward the down-right corner.
	//
	ThirdPartyRedirectAdExpandingDirectionEXPANDING_DOWNRIGHT ThirdPartyRedirectAdExpandingDirection = "EXPANDING_DOWNRIGHT"
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

type VideoType string

const (
	VideoTypeADOBE VideoType = "ADOBE"

	VideoTypeREALPLAYER VideoType = "REALPLAYER"

	VideoTypeQUICKTIME VideoType = "QUICKTIME"

	VideoTypeWINDOWSMEDIA VideoType = "WINDOWSMEDIA"
)

type Get struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 get"`

	//
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	ServiceSelector *Selector `xml:"serviceSelector,omitempty"`
}

type GetResponse struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 getResponse"`

	Rval *AdGroupAdPage `xml:"rval,omitempty"`
}

type Mutate struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 mutate"`

	//
	// <span class="constraint ContentsNotNull">This field must not contain {@code null} elements.</span>
	// <span class="constraint DistinctIds">Elements in this field must have distinct IDs for following {@link Operator}s : SET, REMOVE.</span>
	// <span class="constraint NotEmpty">This field must contain at least one element.</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Operations []*AdGroupAdOperation `xml:"operations,omitempty"`
}

type MutateResponse struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 mutateResponse"`

	Rval *AdGroupAdReturnValue `xml:"rval,omitempty"`
}

type MutateLabel struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 mutateLabel"`

	//
	// <span class="constraint ContentsNotNull">This field must not contain {@code null} elements.</span>
	// <span class="constraint DistinctIds">Elements in this field must have distinct IDs for following {@link Operator}s : ADD, REMOVE.</span>
	// <span class="constraint NotEmpty">This field must contain at least one element.</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	// <span class="constraint SupportedOperators">The following {@link Operator}s are supported: ADD, REMOVE.</span>
	//
	Operations []*AdGroupAdLabelOperation `xml:"operations,omitempty"`
}

type MutateLabelResponse struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 mutateLabelResponse"`

	Rval *AdGroupAdLabelReturnValue `xml:"rval,omitempty"`
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

	Rval *AdGroupAdPage `xml:"rval,omitempty"`
}

type Ad struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Ad"`

	//
	// ID of this ad. This field is ignored when creating
	// ads using {@code AdGroupAdService}.
	// <span class="constraint Selectable">This field can be selected using the value "Id".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Id int64 `xml:"id,omitempty"`

	//
	// Destination URL.
	// <p>Do not set this field if you are using upgraded URLs, as described at:
	// https://developers.google.com/adwords/api/docs/guides/upgraded-urls
	// <span class="constraint Selectable">This field can be selected using the value "Url".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Url string `xml:"url,omitempty"`

	//
	// Visible URL.
	// <span class="constraint Selectable">This field can be selected using the value "DisplayUrl".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	DisplayUrl string `xml:"displayUrl,omitempty"`

	//
	// A list of possible final URLs after all cross domain redirects.
	// <p>This field is used for upgraded urls only, as described at:
	// https://developers.google.com/adwords/api/docs/guides/upgraded-urls
	// <span class="constraint Selectable">This field can be selected using the value "CreativeFinalUrls".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint CollectionSize">The maximum size of this collection is 10.</span>
	//
	FinalUrls []string `xml:"finalUrls,omitempty"`

	//
	// A list of possible final mobile URLs after all cross domain redirects.
	// <p>This field is used for upgraded urls only, as described at:
	// https://developers.google.com/adwords/api/docs/guides/upgraded-urls
	// <span class="constraint Selectable">This field can be selected using the value "CreativeFinalMobileUrls".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint CollectionSize">The maximum size of this collection is 10.</span>
	//
	FinalMobileUrls []string `xml:"finalMobileUrls,omitempty"`

	//
	// A list of final app URLs that will be used on mobile if the user has the specific app
	// installed.
	// <p>This field is used for upgraded urls only, as described at:
	// https://developers.google.com/adwords/api/docs/guides/upgraded-urls
	// <span class="constraint Selectable">This field can be selected using the value "CreativeFinalAppUrls".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	FinalAppUrls []*AppUrl `xml:"finalAppUrls,omitempty"`

	//
	// URL template for constructing a tracking URL.
	// <p>This field is used for upgraded urls only, as described at:
	// https://developers.google.com/adwords/api/docs/guides/upgraded-urls
	// <span class="constraint Selectable">This field can be selected using the value "CreativeTrackingUrlTemplate".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	TrackingUrlTemplate string `xml:"trackingUrlTemplate,omitempty"`

	//
	// URL template for appending params to Final URL.
	//
	// <p>On update, empty string ("") indicates to clear the field.
	// <p>This field is supported only in test accounts.
	//
	FinalUrlSuffix string `xml:"finalUrlSuffix,omitempty"`

	//
	// A list of mappings to be used for substituting URL custom parameter tags in the
	// trackingUrlTemplate, finalUrls, and/or finalMobileUrls.
	// <p>This field is used for upgraded urls only, as described at:
	// https://developers.google.com/adwords/api/docs/guides/upgraded-urls
	// <span class="constraint Selectable">This field can be selected using the value "CreativeUrlCustomParameters".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	UrlCustomParameters *CustomParameters `xml:"urlCustomParameters,omitempty"`

	//
	// Additional urls for the ad that are tagged with a unique identifier. Currently only used for
	// TemplateAds for specific template IDs. For all other ad types, use finalUrls,
	// finalMobileUrls and finalAppUrls instead.
	// <span class="constraint Selectable">This field can be selected using the value "UrlData".</span>
	//
	UrlData []*UrlData `xml:"urlData,omitempty"`

	//
	// Indicates if this ad was added by AdWords.
	// <span class="constraint Selectable">This field can be selected using the value "Automated".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Automated bool `xml:"automated,omitempty"`

	//
	// Type of ad.
	// <span class="constraint Selectable">This field can be selected using the value "AdType".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Type_ *AdType `xml:"type,omitempty"`

	//
	// The device preference for the ad. You can only specify a preference for
	// mobile devices (CriterionId 30001). If unspecified (no device preference),
	// all devices are targeted.
	// <span class="constraint Selectable">This field can be selected using the value "DevicePreference".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	DevicePreference int64 `xml:"devicePreference,omitempty"`

	//
	// The source of this system-managed ad.
	// <span class="constraint Selectable">This field can be selected using the value "SystemManagedEntitySource".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	SystemManagedEntitySource *SystemManagedEntitySource `xml:"systemManagedEntitySource,omitempty"`

	//
	// Indicates that this instance is a subtype of Ad.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	AdType string `xml:"Ad.Type,omitempty"`
}

type AdCustomizerError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdCustomizerError"`

	*ApiError

	Reason *AdCustomizerErrorReason `xml:"reason,omitempty"`

	//
	// String form of the function that contained the error.
	//
	FunctionString string `xml:"functionString,omitempty"`

	//
	// Lowercased string representation of the ad customizer function's operator.
	//
	OperatorName string `xml:"operatorName,omitempty"`

	//
	// Index of the operand that caused the error.
	//
	OperandIndex int32 `xml:"operandIndex,omitempty"`

	//
	// Value of the operand that caused the error.
	//
	OperandValue string `xml:"operandValue,omitempty"`
}

type AdError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *AdErrorReason `xml:"reason,omitempty"`
}

type AdGroupAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdGroupAd"`

	//
	// The id of the adgroup containing this ad.
	// <span class="constraint Selectable">This field can be selected using the value "AdGroupId".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	AdGroupId int64 `xml:"adGroupId,omitempty"`

	//
	// The contents of the ad itself.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Ad *Ad `xml:"ad,omitempty"`

	//
	// The status of the ad.
	// This field is required and should not be {@code null} when it is contained within
	// {@link Operator}s : SET.
	// <span class="constraint Selectable">This field can be selected using the value "Status".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Status *AdGroupAdStatus `xml:"status,omitempty"`

	//
	// Summary of policy findings for this ad.
	// <span class="constraint Selectable">This field can be selected using the value "PolicySummary".</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	PolicySummary *AdGroupAdPolicySummary `xml:"policySummary,omitempty"`

	//
	// Labels that are attached to the AdGroupAd. To associate an existing {@link Label} to an
	// existing {@link AdGroupAd}, use {@link AdGroupAdService#mutateLabel} with ADD operator. To
	// remove an associated {@link Label} from the {@link AdGroupAd}, use
	// {@link AdGroupAdService#mutateLabel} with REMOVE operator. To filter on {@link Label}s,
	// use one of {@link Predicate.Operator#CONTAINS_ALL}, {@link Predicate.Operator#CONTAINS_ANY},
	// {@link Predicate.Operator#CONTAINS_NONE} operators with a list of {@link Label} ids.
	// <span class="constraint Selectable">This field can be selected using the value "Labels".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint CampaignType">This field may not be set for campaign channel subtype UNIVERSAL_APP_CAMPAIGN.</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API for the following {@link Operator}s: REMOVE and SET.</span>
	//
	Labels []*Label `xml:"labels,omitempty"`

	//
	// ID of the base campaign from which this draft/trial ad was created.
	// This field is only returned on get requests.
	// <span class="constraint Selectable">This field can be selected using the value "BaseCampaignId".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	BaseCampaignId int64 `xml:"baseCampaignId,omitempty"`

	//
	// ID of the base ad group from which this draft/trial ad was created. For
	// base ad groups this is equal to the ad group ID.  If the ad group was created
	// in the draft or trial and has no corresponding base ad group, this field is null.
	// This field is only returned on get requests.
	// <span class="constraint Selectable">This field can be selected using the value "BaseAdGroupId".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	BaseAdGroupId int64 `xml:"baseAdGroupId,omitempty"`

	//
	// This Map provides a place to put new features and settings in older versions
	// of the AdWords API in the rare instance we need to introduce a new feature in
	// an older version.
	//
	// It is presently unused.  Do not set a value.
	//
	ForwardCompatibilityMap []*String_StringMapEntry `xml:"forwardCompatibilityMap,omitempty"`
}

type AdGroupAdCountLimitExceeded struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdGroupAdCountLimitExceeded"`

	*EntityCountLimitExceeded
}

type AdGroupAdError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdGroupAdError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *AdGroupAdErrorReason `xml:"reason,omitempty"`
}

type AdGroupAdLabel struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdGroupAdLabel"`

	//
	// The id of the adgroup containing the ad that the label to be applied to.
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD, REMOVE.</span>
	//
	AdGroupId int64 `xml:"adGroupId,omitempty"`

	//
	// The id of the ad that the label to be applied to.
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD, REMOVE.</span>
	//
	AdId int64 `xml:"adId,omitempty"`

	//
	// The id of an existing label to be applied to the adgroup ad.
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD, REMOVE.</span>
	//
	LabelId int64 `xml:"labelId,omitempty"`
}

type AdGroupAdLabelOperation struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdGroupAdLabelOperation"`

	*Operation

	//
	// AdGroupAdLabel to operate on.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Operand *AdGroupAdLabel `xml:"operand,omitempty"`
}

type AdGroupAdLabelReturnValue struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdGroupAdLabelReturnValue"`

	*ListReturnValue

	Value []*AdGroupAdLabel `xml:"value,omitempty"`

	PartialFailureErrors []*ApiError `xml:"partialFailureErrors,omitempty"`
}

type AdGroupAdOperation struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdGroupAdOperation"`

	*Operation

	//
	// AdGroupAd to operate on.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Operand *AdGroupAd `xml:"operand,omitempty"`

	//
	// Exemption requests for any policy violations in this Ad.  This field is
	// only used for ADD operations
	//
	ExemptionRequests []*ExemptionRequest `xml:"exemptionRequests,omitempty"`
}

type AdGroupAdPage struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdGroupAdPage"`

	*Page

	//
	// The result entries in this page.
	//
	Entries []*AdGroupAd `xml:"entries,omitempty"`
}

type AdGroupAdPolicySummary struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdGroupAdPolicySummary"`

	//
	// List of policy findings.
	//
	PolicyTopicEntries []*PolicyTopicEntry `xml:"policyTopicEntries,omitempty"`

	//
	// Progress through the review process.
	//
	ReviewState *PolicySummaryReviewState `xml:"reviewState,omitempty"`

	//
	// Overall review status based on the policy topic entries.
	//
	DenormalizedStatus *PolicySummaryDenormalizedStatus `xml:"denormalizedStatus,omitempty"`

	//
	// Approval status that combines review state and status.
	// <span class="constraint Selectable">This field can be selected using the value "CombinedApprovalStatus".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	CombinedApprovalStatus *PolicyApprovalStatus `xml:"combinedApprovalStatus,omitempty"`
}

type AdGroupAdReturnValue struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdGroupAdReturnValue"`

	*ListReturnValue

	//
	// List of ads in an ad group.
	//
	Value []*AdGroupAd `xml:"value,omitempty"`

	PartialFailureErrors []*ApiError `xml:"partialFailureErrors,omitempty"`
}

type AdSharingError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdSharingError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *AdSharingErrorReason `xml:"reason,omitempty"`

	SharedAdError *ApiError `xml:"sharedAdError,omitempty"`
}

type AdUnionId struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdUnionId"`

	//
	// The ID of the ad union
	// <span class="constraint InRange">This field must be greater than or equal to 1.</span>
	//
	Id int64 `xml:"id,omitempty"`

	//
	// Indicates that this instance is a subtype of AdUnionId.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	AdUnionIdType string `xml:"AdUnionId.Type,omitempty"`
}

type AdxError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AdxError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *AdxErrorReason `xml:"reason,omitempty"`
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

type AppUrl struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 AppUrl"`

	//
	// The app deep link url. E.g. "android-app://com.my.App"
	//
	Url string `xml:"url,omitempty"`

	//
	// The operating system targeted by this url.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	OsType *AppUrlOsType `xml:"osType,omitempty"`
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

type LabelAttribute struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 LabelAttribute"`

	//
	// Indicates that this instance is a subtype of LabelAttribute.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	LabelAttributeType string `xml:"LabelAttribute.Type,omitempty"`
}

type Audio struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Audio"`

	*Media

	//
	// The duration of the associated audio, in milliseconds.
	// <span class="constraint Selectable">This field can be selected using the value "DurationMillis".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	DurationMillis int64 `xml:"durationMillis,omitempty"`

	//
	// The streaming URL of the audio.
	// <span class="constraint Selectable">This field can be selected using the value "StreamingUrl".</span>
	//
	StreamingUrl string `xml:"streamingUrl,omitempty"`

	//
	// Indicates whether the audio is ready to play on the web.
	// <span class="constraint Selectable">This field can be selected using the value "ReadyToPlayOnTheWeb".</span>
	//
	ReadyToPlayOnTheWeb bool `xml:"readyToPlayOnTheWeb,omitempty"`
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

type CallOnlyAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CallOnlyAd"`

	*Ad

	//
	// Two letter country code for the ad. Examples: 'US', 'GB'.
	// <span class="constraint Selectable">This field can be selected using the value "CallOnlyAdCountryCode".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	CountryCode string `xml:"countryCode,omitempty"`

	//
	// Phone number string for the ad.
	// Examples: '(800) 356-9377', "16502531234", "+442001234567"
	// <span class="constraint Selectable">This field can be selected using the value "CallOnlyAdPhoneNumber".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	PhoneNumber string `xml:"phoneNumber,omitempty"`

	//
	// Business name of the ad.
	// <span class="constraint Selectable">This field can be selected using the value "CallOnlyAdBusinessName".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	BusinessName string `xml:"businessName,omitempty"`

	//
	// First line of ad text.
	// <span class="constraint Selectable">This field can be selected using the value "CallOnlyAdDescription1".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Description1 string `xml:"description1,omitempty"`

	//
	// Second line of ad text.
	// <span class="constraint Selectable">This field can be selected using the value "CallOnlyAdDescription2".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Description2 string `xml:"description2,omitempty"`

	//
	// If set to true, enable call tracking for the creative. Enabling call
	// tracking also enables call conversions.
	// <span class="constraint Selectable">This field can be selected using the value "CallOnlyAdCallTracked".</span>
	//
	CallTracked bool `xml:"callTracked,omitempty"`

	//
	// By default, call conversions are enabled when callTracked is on.
	// To disable call conversions, set this field to true.
	// Only in effect if callTracked is also set to true. If callTracked is set
	// to false, this field is ignored.
	// <span class="constraint Selectable">This field can be selected using the value "CallOnlyAdDisableCallConversion".</span>
	//
	DisableCallConversion bool `xml:"disableCallConversion,omitempty"`

	//
	// Conversion type to attribute a call conversion to. If not set, then a
	// default conversion type id is used. Only in effect if callTracked is also
	// set to true otherwise this field is ignored.
	// <span class="constraint Selectable">This field can be selected using the value "CallOnlyAdConversionTypeId".</span>
	//
	ConversionTypeId int64 `xml:"conversionTypeId,omitempty"`

	//
	// Url to be used for phone number verification.
	// <span class="constraint Selectable">This field can be selected using the value "CallOnlyAdPhoneNumberVerificationUrl".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	PhoneNumberVerificationUrl string `xml:"phoneNumberVerificationUrl,omitempty"`
}

type TextLabel struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 TextLabel"`

	*Label
}

type DisplayAttribute struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 DisplayAttribute"`

	*LabelAttribute

	//
	// Background color of the label in RGB format.
	// <span class="constraint MatchesRegex">A background color string must begin with a '#' character followed by either 6 or 3 hexadecimal characters (24 vs. 12 bits). This is checked by the regular expression '^\#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$'.</span>
	//
	BackgroundColor string `xml:"backgroundColor,omitempty"`

	//
	// A short description of the label.
	// <span class="constraint StringLength">The length of this string should be between 0 and 200, inclusive.</span>
	//
	Description string `xml:"description,omitempty"`
}

type CertificateDomainMismatchInCountryConstraint struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CertificateDomainMismatchInCountryConstraint"`

	*CountryConstraint
}

type CertificateMissingConstraint struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CertificateMissingConstraint"`

	*PolicyTopicConstraint
}

type CertificateMissingInCountryConstraint struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CertificateMissingInCountryConstraint"`

	*CountryConstraint
}

type ClientTermsError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ClientTermsError"`

	*ApiError

	Reason *ClientTermsErrorReason `xml:"reason,omitempty"`
}

type CountryConstraint struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 CountryConstraint"`

	*PolicyTopicConstraint

	//
	// The set of targeted country criterion IDs to which a policy topic entry applies.
	//
	ConstrainedCountries []int64 `xml:"constrainedCountries,omitempty"`

	//
	// The total number of targeted countries.
	//
	TotalTargetedCountries int32 `xml:"totalTargetedCountries,omitempty"`
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

type DeprecatedAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 DeprecatedAd"`

	*Ad

	//
	// Name of the ad.
	// <span class="constraint Selectable">This field can be selected using the value "Name".</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	Name string `xml:"name,omitempty"`

	//
	// Type of the creative.
	// <span class="constraint Selectable">This field can be selected using the value "Type".</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	DeprecatedAdType *DeprecatedAdType `xml:"deprecatedAdType,omitempty"`
}

type Dimensions struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Dimensions"`

	//
	// Width of the dimension
	// <span class="constraint Selectable">This field can be selected using the value "Width".</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Width int32 `xml:"width,omitempty"`

	//
	// Height of the dimension
	// <span class="constraint Selectable">This field can be selected using the value "Height".</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Height int32 `xml:"height,omitempty"`
}

type DisplayCallToAction struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 DisplayCallToAction"`

	//
	// Text of the display-call-to-action. Maximum display width is 15 characters.
	// <span class="constraint Selectable">This field can be selected using the value "MarketingImageCallToActionText".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	Text string `xml:"text,omitempty"`

	//
	// Text color of the display-call-to-action. In hexadecimal, e.g. #ffffff for white.
	// <span class="constraint Selectable">This field can be selected using the value "MarketingImageCallToActionTextColor".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	TextColor string `xml:"textColor,omitempty"`

	//
	// Identifies the url data in Ad.urlData used for this DisplayCallToAction. If not set, the url
	// defaults to {@link Ad#finalUrls}.
	//
	UrlId string `xml:"urlId,omitempty"`
}

type DistinctError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 DistinctError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *DistinctErrorReason `xml:"reason,omitempty"`
}

type DynamicSettings struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 DynamicSettings"`

	//
	// Landscape logo image. This ad format does not allow the creation of an image using the
	// Image.data field. An image must first be created using the MediaService, and Image.mediaId must
	// be populated when creating a {@link "DynamicSettings"}. Valid image types are GIF, JPEG, and
	// PNG. The minimum size is 512x128 the aspect ratio must be 512:128 (+-1%).
	// <span class="constraint Selectable">This field can be selected using the value "LandscapeLogoImage".</span>
	//
	LandscapeLogoImage *Image `xml:"landscapeLogoImage,omitempty"`

	//
	// Prefix before price. Maximum display width is 10. example, "as low as".
	// <span class="constraint Selectable">This field can be selected using the value "PricePrefix".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	PricePrefix string `xml:"pricePrefix,omitempty"`

	//
	// Promotion text used for dynamic formats of responsive ads. Maximum display width is 25. For
	// example, "Free two-day shipping".
	// <span class="constraint Selectable">This field can be selected using the value "PromoText".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	PromoText string `xml:"promoText,omitempty"`
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

type ExemptionRequest struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ExemptionRequest"`

	//
	// Identifies the violation to request an exemption for.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Key *PolicyViolationKey `xml:"key,omitempty"`
}

type ExpandedDynamicSearchAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ExpandedDynamicSearchAd"`

	*Ad

	//
	// The descriptive text of the ad.
	// <span class="constraint Selectable">This field can be selected using the value "Description".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	Description string `xml:"description,omitempty"`
}

type ExpandedTextAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ExpandedTextAd"`

	*Ad

	//
	// First part of the headline.
	// <span class="constraint Selectable">This field can be selected using the value "HeadlinePart1".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	HeadlinePart1 string `xml:"headlinePart1,omitempty"`

	//
	// Second part of the headline.
	// <span class="constraint Selectable">This field can be selected using the value "HeadlinePart2".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	HeadlinePart2 string `xml:"headlinePart2,omitempty"`

	//
	// The descriptive text of the ad.
	// <span class="constraint Selectable">This field can be selected using the value "Description".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	Description string `xml:"description,omitempty"`

	//
	// Text that appears in the ad with the displayed URL.
	// <span class="constraint Selectable">This field can be selected using the value "Path1".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Path1 string `xml:"path1,omitempty"`

	//
	// In addition to {@link #path1}, more text that appears with the displayed URL.
	// <span class="constraint Selectable">This field can be selected using the value "Path2".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Path2 string `xml:"path2,omitempty"`
}

type FeedAttributeReferenceError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 FeedAttributeReferenceError"`

	*ApiError

	Reason *FeedAttributeReferenceErrorReason `xml:"reason,omitempty"`

	//
	// The referenced feed name.
	//
	FeedName string `xml:"feedName,omitempty"`

	//
	// The referenced feed attribute name.
	//
	FeedAttributeName string `xml:"feedAttributeName,omitempty"`
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

type ForwardCompatibilityError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ForwardCompatibilityError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *ForwardCompatibilityErrorReason `xml:"reason,omitempty"`
}

type FunctionError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 FunctionError"`

	*ApiError

	//
	// The error reason represented by an enum
	//
	Reason *FunctionErrorReason `xml:"reason,omitempty"`
}

type FunctionParsingError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 FunctionParsingError"`

	*ApiError

	Reason *FunctionParsingErrorReason `xml:"reason,omitempty"`

	OffendingText string `xml:"offendingText,omitempty"`

	OffendingTextIndex int32 `xml:"offendingTextIndex,omitempty"`
}

type GmailAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 GmailAd"`

	*Ad

	//
	// Gmail teaser info.
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	Teaser *GmailTeaser `xml:"teaser,omitempty"`

	//
	// Header image. An image must first be created using the MediaService, and Image.mediaId must be
	// populated when creating a {@link "GmailAd"}. Valid image types are GIF, JPEG, and PNG. The
	// minimum size is 300x100 and the aspect ratio must be in 3:1 to 5:1 (+-1%).
	// <span class="constraint Selectable">This field can be selected using the value "GmailHeaderImage".</span>
	//
	HeaderImage *Image `xml:"headerImage,omitempty"`

	//
	// Marketing image. An image must first be created using the MediaService, and Image.mediaId must
	// be populated when creating a {@link "GmailAd"}.Valid image types are GIF, JPEG, and PNG. The
	// minimum size is 600x314 and the aspect ratio must be 600:314 (+-1%). For square marketing
	// image, the minimum size is 300x300 and the aspect ratio must be 1:1 (+-1%). Either
	// productVideos or marketingImage must be specified.
	// <span class="constraint Selectable">This field can be selected using the value "GmailMarketingImage".</span>
	//
	MarketingImage *Image `xml:"marketingImage,omitempty"`

	//
	// Headline of the marketing image. Maximum display width is 25 characters.
	// <span class="constraint Selectable">This field can be selected using the value "MarketingImageHeadline".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	MarketingImageHeadline string `xml:"marketingImageHeadline,omitempty"`

	//
	// Description of the marketing image. Maximum display width is 90 characters.
	// <span class="constraint Selectable">This field can be selected using the value "MarketingImageDescription".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	MarketingImageDescription string `xml:"marketingImageDescription,omitempty"`

	//
	// Display-call-to-action of the marketing image. The DisplayCallToAction.urlId field cannot be
	// set when setting this field.
	//
	MarketingImageDisplayCallToAction *DisplayCallToAction `xml:"marketingImageDisplayCallToAction,omitempty"`

	//
	// Product images. Support up to 15 product images.
	// <span class="constraint Selectable">This field can be selected using the value "ProductImages".</span>
	//
	ProductImages []*ProductImage `xml:"productImages,omitempty"`

	//
	// Product Videos. Either productVideoList or marketingImage must be specified. Supports up to 7
	// product videos. It must be a YouTube hosted video and mediaId must be populated. The YouTube
	// hosted video can be added to AdWords through either the AdWords UI or through AdWords Scripts
	// (https://developers.google.com/adwords/scripts/docs/reference/adwordsapp/adwordsapp_videobuilder).
	// <span class="constraint Selectable">This field can be selected using the value "ProductVideoList".</span>
	//
	ProductVideoList []*Video `xml:"productVideoList,omitempty"`
}

type GmailTeaser struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 GmailTeaser"`

	//
	// Headline of the teaser. Maximum display width is 25 characters.
	// <span class="constraint Selectable">This field can be selected using the value "GmailTeaserHeadline". This field can be selected using the value "DisplayUploadAdGmailTeaserHeadline".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	Headline string `xml:"headline,omitempty"`

	//
	// Description of the teaser. Maximum display width is 90 characters.
	// <span class="constraint Selectable">This field can be selected using the value "GmailTeaserDescription". This field can be selected using the value "DisplayUploadAdGmailTeaserDescription".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	Description string `xml:"description,omitempty"`

	//
	// Business name of the advertiser. Maximum display width is 20 characters.
	// <span class="constraint Selectable">This field can be selected using the value "GmailTeaserBusinessName". This field can be selected using the value "DisplayUploadAdGmailTeaserBusinessName".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	BusinessName string `xml:"businessName,omitempty"`

	//
	// Required. Logo image. An image must first be created using the MediaService, and Image.mediaId
	// must be populated when creating a {@link "GmailTeaser"}. Valid image types are GIF, JPEG, and
	// PNG. The minimum size is 144x144 and the aspect ratio must be 1:1 (+-1%). Required.
	// <span class="constraint Selectable">This field can be selected using the value "GmailTeaserLogoImage". This field can be selected using the value "DisplayUploadAdGmailTeaserLogoImage".</span>
	//
	LogoImage *Image `xml:"logoImage,omitempty"`
}

type IdError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 IdError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *IdErrorReason `xml:"reason,omitempty"`
}

type Image struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Image"`

	*Media

	//
	// Raw image data.
	//
	Data []byte `xml:"data,omitempty"`
}

type ImageAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ImageAd"`

	*Ad

	//
	// The image data for the ad.
	//
	Image *Image `xml:"image,omitempty"`

	//
	// The name label for this ad.
	// <span class="constraint Required">
	// This field is required and should not be {@code null}.</span>
	// <span class="constraint Selectable">This field can be selected using the value "ImageCreativeName".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Name string `xml:"name,omitempty"`

	//
	// For ADD operations only: use this field to specify an existing
	// image ad to copy the image from, in which case the "image" field
	// can be left empty. This is the preferred way of copying images
	// over re-uploading the same image.
	// <span class="constraint ReadOnly">This field is read only and will be ignored
	// when sent to the API for the following {@link Operator}s: REMOVE and SET.</span>
	//
	AdToCopyImageFrom int64 `xml:"adToCopyImageFrom,omitempty"`
}

type ImageError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ImageError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *ImageErrorReason `xml:"reason,omitempty"`
}

type InternalApiError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 InternalApiError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *InternalApiErrorReason `xml:"reason,omitempty"`
}

type Label struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Label"`

	//
	// Id of label.
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : SET, REMOVE.</span>
	//
	Id int64 `xml:"id,omitempty"`

	//
	// Name of label.
	// <span class="constraint StringLength">The length of this string should be between 1 and 80, inclusive.</span>
	//
	Name string `xml:"name,omitempty"`

	//
	// Status of the label.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	Status *LabelStatus `xml:"status,omitempty"`

	//
	// Attributes of the label.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API for the following {@link Operator}s: REMOVE.</span>
	//
	Attribute *LabelAttribute `xml:"attribute,omitempty"`

	//
	// Indicates that this instance is a subtype of Label.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	LabelType string `xml:"Label.Type,omitempty"`
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

type Media struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Media"`

	//
	// ID of this media object.
	// <span class="constraint Selectable">This field can be selected using the value "MediaId".</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : SET, REMOVE.</span>
	//
	MediaId int64 `xml:"mediaId,omitempty"`

	//
	// Type of media object. Required when using {@link MediaService#upload} to upload a new media
	// file. MEDIA_BUNDLE, ICON, IMAGE, and DYNAMIC_IMAGE are the supported MediaTypes to upload.
	// <span class="constraint Selectable">This field can be selected using the value "Type".</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API for the following {@link Operator}s: REMOVE and SET.</span>
	//
	Type_ *MediaMediaType `xml:"type,omitempty"`

	//
	// Media reference ID key.
	// <span class="constraint Selectable">This field can be selected using the value "ReferenceId".</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API for the following {@link Operator}s: REMOVE and SET.</span>
	//
	ReferenceId int64 `xml:"referenceId,omitempty"`

	//
	// Various dimension sizes for the media. Only applies to image media (and video media for
	// video thumbnails).
	// <span class="constraint Selectable">This field can be selected using the value "Dimensions".</span>
	//
	Dimensions []*Media_Size_DimensionsMapEntry `xml:"dimensions,omitempty"`

	//
	// URLs pointing to the resized media for the given sizes. Only applies to image media.
	// <span class="constraint Selectable">This field can be selected using the value "Urls".</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	Urls []*Media_Size_StringMapEntry `xml:"urls,omitempty"`

	//
	// The mime type of the media.
	// <span class="constraint Selectable">This field can be selected using the value "MimeType".</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API for the following {@link Operator}s: REMOVE and SET.</span>
	//
	MimeType *MediaMimeType `xml:"mimeType,omitempty"`

	//
	// The URL of where the original media was downloaded from (or a file name).
	// <span class="constraint Selectable">This field can be selected using the value "SourceUrl".</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API for the following {@link Operator}s: REMOVE and SET.</span>
	//
	SourceUrl string `xml:"sourceUrl,omitempty"`

	//
	// The name of the media. The name can be used by clients to
	// help identify previously uploaded media.
	// <span class="constraint Selectable">This field can be selected using the value "Name".</span>
	//
	Name string `xml:"name,omitempty"`

	//
	// The size of the media file in bytes.
	// <span class="constraint Selectable">This field can be selected using the value "FileSize".</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API for the following {@link Operator}s: REMOVE and SET.</span>
	//
	FileSize int64 `xml:"fileSize,omitempty"`

	//
	// Media creation date in the format YYYY-MM-DD HH:MM:SS+TZ.
	// This is not updatable and not specifiable.
	// <span class="constraint Selectable">This field can be selected using the value "CreationTime".</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API for the following {@link Operator}s: REMOVE and SET.</span>
	//
	CreationTime string `xml:"creationTime,omitempty"`

	//
	// Indicates that this instance is a subtype of Media.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	MediaType string `xml:"Media.Type,omitempty"`
}

type MediaBundle struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 MediaBundle"`

	*Media

	//
	// Raw zipped data.
	//
	Data []byte `xml:"data,omitempty"`

	//
	// URL pointing to the data for the MediaBundle data.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	MediaBundleUrl string `xml:"mediaBundleUrl,omitempty"`

	//
	// Entry in the ZIP archive used to display the <code>MediaBundle</code> in an
	// <code>Ad</code>. This field can only be set and returned when the <code>MediaBundle</code> is
	// used with the <code>AdGroupAdService</code>. If this field is set when calling
	// <code>MediaService</code>, an error will be returned.
	//
	// <p>To use a <code>MediaBundle</code> that was created with the <code>MediaService</code> in
	// an <code>Ad</code>, create a bundle and set the <code>mediaId</code> and
	// <code>entryPoint</code> fields.
	//
	EntryPoint string `xml:"entryPoint,omitempty"`
}

type MediaBundleError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 MediaBundleError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *MediaBundleErrorReason `xml:"reason,omitempty"`
}

type MediaError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 MediaError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *MediaErrorReason `xml:"reason,omitempty"`
}

type Media_Size_DimensionsMapEntry struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Media_Size_DimensionsMapEntry"`

	Key *MediaSize `xml:"key,omitempty"`

	Value *Dimensions `xml:"value,omitempty"`
}

type Media_Size_StringMapEntry struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Media_Size_StringMapEntry"`

	Key *MediaSize `xml:"key,omitempty"`

	Value string `xml:"value,omitempty"`
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

type PagingError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 PagingError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *PagingErrorReason `xml:"reason,omitempty"`
}

type PolicyTopicConstraint struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 PolicyTopicConstraint"`

	ConstraintType *PolicyTopicConstraintPolicyTopicConstraintType `xml:"constraintType,omitempty"`

	//
	// Indicates that this instance is a subtype of PolicyTopicConstraint.
	// Although this field is returned in the response, it is ignored on input
	// and cannot be selected. Specify xsi:type instead.
	//
	PolicyTopicConstraintType string `xml:"PolicyTopicConstraint.Type,omitempty"`
}

type PolicyTopicEntry struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 PolicyTopicEntry"`

	//
	// The type of the policy topic entry.
	//
	PolicyTopicEntryType *PolicyTopicEntryType `xml:"policyTopicEntryType,omitempty"`

	//
	// The policy topic evidences associated with this policy topic entry.
	//
	PolicyTopicEvidences []*PolicyTopicEvidence `xml:"policyTopicEvidences,omitempty"`

	//
	// The targeting constraints to which this PolicyTopicEntry is related.
	//
	PolicyTopicConstraints []*PolicyTopicConstraint `xml:"policyTopicConstraints,omitempty"`

	//
	// The policy topic id.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	PolicyTopicId string `xml:"policyTopicId,omitempty"`

	//
	// The policy topic name (in English).
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	PolicyTopicName string `xml:"policyTopicName,omitempty"`

	//
	// URL of the help center article describing this policy topic entry.
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	PolicyTopicHelpCenterUrl string `xml:"policyTopicHelpCenterUrl,omitempty"`
}

type PolicyTopicEvidence struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 PolicyTopicEvidence"`

	//
	// The type of evidence for the policy topic.
	//
	PolicyTopicEvidenceType *PolicyTopicEvidenceType `xml:"policyTopicEvidenceType,omitempty"`

	//
	// The actual evidence that triggered this policy topic to be reported. This field is associated
	// with the policyTopicEvidenceType. So for example, when policyTopicEvidenceType is AD_TEXT,
	// the evidence is the texts associated with the Ad.
	//
	EvidenceTextList []string `xml:"evidenceTextList,omitempty"`
}

type PolicyViolationError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 PolicyViolationError"`

	*ApiError

	//
	// Unique identifier for the violation.
	//
	Key *PolicyViolationKey `xml:"key,omitempty"`

	//
	// Name of policy suitable for display to users. In the user's preferred
	// language.
	//
	ExternalPolicyName string `xml:"externalPolicyName,omitempty"`

	//
	// Url with writeup about the policy.
	//
	ExternalPolicyUrl string `xml:"externalPolicyUrl,omitempty"`

	//
	// Localized description of the violation.
	//
	ExternalPolicyDescription string `xml:"externalPolicyDescription,omitempty"`

	//
	// Whether user can file an exemption request for this violation.
	//
	IsExemptable bool `xml:"isExemptable,omitempty"`

	//
	// Lists the parts that violate the policy.
	//
	ViolatingParts []*PolicyViolationErrorPart `xml:"violatingParts,omitempty"`
}

type PolicyViolationErrorPart struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 PolicyViolationError.Part"`

	//
	// Index of the starting position of the violating text within the line.
	//
	Index int32 `xml:"index,omitempty"`

	//
	// The length of the violating text.
	//
	Length int32 `xml:"length,omitempty"`
}

type PolicyViolationKey struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 PolicyViolationKey"`

	//
	// Unique id of the violated policy.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	PolicyName string `xml:"policyName,omitempty"`

	//
	// The text that violates the policy if specified. Otherwise, refers to the
	// policy in general (e.g. when requesting to be exempt from the whole
	// policy).
	//
	// May be null for criterion exemptions, in which case this refers to the
	// whole policy. Must be specified for ad exemptions.
	//
	ViolatingText string `xml:"violatingText,omitempty"`
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

type ProductAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ProductAd"`

	*Ad
}

type ProductImage struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ProductImage"`

	//
	// Product image. An image must first be created using the MediaService, and Image.mediaId must be
	// populated when creating a {@link "ProductImage"}. Valid image types are GIF, JPEG, and PNG. The
	// minimum size is 300x300 and the aspect ratio must be 1:1 (+-1%).
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	ProductImage *Image `xml:"productImage,omitempty"`

	//
	// Description of the product. Maximum display width is 15 characters.
	//
	Description string `xml:"description,omitempty"`

	//
	// Display-call-to-action of the product image. The DisplayCallToAction.textColor field cannot be
	// set when setting this field.
	//
	DisplayCallToAction *DisplayCallToAction `xml:"displayCallToAction,omitempty"`
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

type ResellerConstraint struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ResellerConstraint"`

	*PolicyTopicConstraint
}

type ResponsiveDisplayAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ResponsiveDisplayAd"`

	*Ad

	//
	// Marketing image to be used in the ad. This ad format does not allow the creation of an image
	// using the Image.data field. An image must first be created using the MediaService, and
	// Image.mediaId must be populated when creating a {@link "ResponsiveDisplayAd"}. Valid image
	// types are GIF, JPEG, and PNG. The minimum size is 600x314 and the aspect ratio must be 600:314
	// (+-1%).
	// <span class="constraint Selectable">This field can be selected using the value "MarketingImage".</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	MarketingImage *Image `xml:"marketingImage,omitempty"`

	//
	// Logo image to be used in the ad. This ad format does not allow the creation of an image using
	// the Image.data field. An image must first be created using the MediaService, and Image.mediaId
	// must be populated when creating a {@link "ResponsiveDisplayAd"}. Valid image types are GIF,
	// JPEG, and PNG. The minimum size is 128x128 and the aspect ratio must be 1:1 (+-1%).
	// <span class="constraint Selectable">This field can be selected using the value "LogoImage".</span>
	//
	LogoImage *Image `xml:"logoImage,omitempty"`

	//
	// Square marketing image to be used in the ad. This image may be used when a square aspect ratio
	// is more appropriate than the aspect ratio of the {@link #marketingImage} image. This ad format
	// does not allow the creation of an image using the Image.data field. An image must first be
	// created using the MediaService, and Image.mediaId must be populated when creating a {@link
	// "ResponsiveDisplayAd"}. Valid image types are GIF, JPEG, and PNG. The minimum size is 300x300
	// the aspect ratio must be 1:1 (+-1%).
	// <span class="constraint Selectable">This field can be selected using the value "SquareMarketingImage".</span>
	//
	SquareMarketingImage *Image `xml:"squareMarketingImage,omitempty"`

	//
	// Short format of the headline of the ad. Maximum display width is 25.
	// <span class="constraint Selectable">This field can be selected using the value "ShortHeadline".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	ShortHeadline string `xml:"shortHeadline,omitempty"`

	//
	// Long format of the headline of the ad. Maximum display width is 90.
	// <span class="constraint Selectable">This field can be selected using the value "LongHeadline".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	LongHeadline string `xml:"longHeadline,omitempty"`

	//
	// The descriptive text of the ad. Maximum display width is 90.
	// <span class="constraint Selectable">This field can be selected using the value "Description".</span><span class="constraint Filterable">This field can be filtered on.</span>
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	Description string `xml:"description,omitempty"`

	//
	// The business name. Maximum display width is 25. <span class="constraint Required">This field is
	// required and should not be {@code null} when it is contained within {@link Operator}s :
	// ADD.</span>
	// <span class="constraint Selectable">This field can be selected using the value "BusinessName".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	BusinessName string `xml:"businessName,omitempty"`

	//
	// Main color. In hexadecimal, e.g. #ffffff for white. If one of mainColor and accentColor is set,
	// the other is required as well.
	// <span class="constraint Selectable">This field can be selected using the value "MainColor".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	MainColor string `xml:"mainColor,omitempty"`

	//
	// Accent color. In hexadecimal, e.g. #ffffff for white. If one of mainColor and accentColor is
	// set, the other is required as well.
	// <span class="constraint Selectable">This field can be selected using the value "AccentColor".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	AccentColor string `xml:"accentColor,omitempty"`

	//
	// Advertiser?s consent to allow flexible color. When true, we may serve the ad with different
	// color when necessary. When false, we will serve the ad with advertiser color or neutral color.
	// Must be true if mainColor and accentColor are not set. The default value is true.
	// <span class="constraint Selectable">This field can be selected using the value "AllowFlexibleColor".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	AllowFlexibleColor bool `xml:"allowFlexibleColor,omitempty"`

	//
	// Call to action text. Valid texts: https://support.google.com/adwords/answer/7005917
	// <span class="constraint Selectable">This field can be selected using the value "CallToActionText".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	CallToActionText string `xml:"callToActionText,omitempty"`

	//
	// Settings for serving dynamic ResponsiveDisplayAd.
	//
	DynamicDisplayAdSettings *DynamicSettings `xml:"dynamicDisplayAdSettings,omitempty"`

	//
	// Specifies which format the ad will be served in. The default value is ALL_FORMATS.
	// <span class="constraint Selectable">This field can be selected using the value "FormatSetting".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	FormatSetting *DisplayAdFormatSetting `xml:"formatSetting,omitempty"`
}

type RichMediaAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 RichMediaAd"`

	*Ad

	//
	// Name of the rich media ad.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	// <span class="constraint Selectable">This field can be selected using the value "RichMediaAdName".</span>
	//
	Name string `xml:"name,omitempty"`

	//
	// Dimensions (height and width) of the ad.
	//
	// This field is optional for ThirdPartyRedirectAd.  Ad Exchange traditional
	// yield management creatives do not specify the dimension on the
	// ThirdPartyRedirectAd; instead, the size is specified in the publisher
	// front end when creating a mediation chain.
	//
	Dimensions *Dimensions `xml:"dimensions,omitempty"`

	//
	// Snippet for this ad. Required for standard third-party ads.
	// <p>The length of the string should be between 1 and 3072, inclusive.
	// <span class="constraint Selectable">This field can be selected using the value "RichMediaAdSnippet".</span>
	//
	Snippet string `xml:"snippet,omitempty"`

	//
	// Impression beacon URL for the ad.
	// <span class="constraint Selectable">This field can be selected using the value "RichMediaAdImpressionBeaconUrl".</span>
	//
	ImpressionBeaconUrl string `xml:"impressionBeaconUrl,omitempty"`

	//
	// Duration for the ad (in milliseconds). Default is 0.
	// <span class="constraint Selectable">This field can be selected using the value "RichMediaAdDuration".</span>
	// <span class="constraint InRange">This field must be greater than or equal to 0.</span>
	//
	AdDuration int32 `xml:"adDuration,omitempty"`

	//
	// <a href="/adwords/api/docs/appendix/richmediacodes">
	// Certified Vendor Format ID</a>.
	// <span class="constraint Selectable">This field can be selected using the value "RichMediaAdCertifiedVendorFormatId".</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	CertifiedVendorFormatId int64 `xml:"certifiedVendorFormatId,omitempty"`

	//
	// SourceUrl pointing to the third party snippet.
	// For third party in-stream video ads, this stores the VAST URL. For DFA ads,
	// it stores the InRed URL.
	// <span class="constraint Selectable">This field can be selected using the value "RichMediaAdSourceUrl".</span>
	//
	SourceUrl string `xml:"sourceUrl,omitempty"`

	//
	// Type of this rich media ad, the default is Standard.
	// <span class="constraint Selectable">This field can be selected using the value "RichMediaAdType".</span>
	//
	RichMediaAdType *RichMediaAdRichMediaAdType `xml:"richMediaAdType,omitempty"`

	//
	// A list of attributes that describe the rich media ad capabilities.
	//
	AdAttributes []*RichMediaAdAdAttribute `xml:"adAttributes,omitempty"`
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

type ShowcaseAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ShowcaseAd"`

	*Ad

	//
	// The name label for this ad.
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	Name string `xml:"name,omitempty"`

	//
	// Headline displayed in the Showcase shopping ad.
	//
	Headline string `xml:"headline,omitempty"`

	//
	// Description displayed in the expanded view of the Showcase shopping ad.
	//
	Description string `xml:"description,omitempty"`

	//
	// Image displayed in the collapsed view of the Showcase shopping ad.
	// <p>The format of the image must be either JPEG or PNG and the size of the image must be
	// 270x270 px.
	//
	CollapsedImage *Image `xml:"collapsedImage,omitempty"`

	//
	// Image displayed in the expanded view of the Showcase shopping ad.
	// <p>The format of the image must be either JPEG or PNG and the size of the image must be
	// 1080x566 px.
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	//
	ExpandedImage *Image `xml:"expandedImage,omitempty"`
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

type StatsQueryError struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 StatsQueryError"`

	*ApiError

	//
	// The error reason represented by an enum.
	//
	Reason *StatsQueryErrorReason `xml:"reason,omitempty"`
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

type String_StringMapEntry struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 String_StringMapEntry"`

	Key string `xml:"key,omitempty"`

	Value string `xml:"value,omitempty"`
}

type TempAdUnionId struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 TempAdUnionId"`

	*AdUnionId
}

type TemplateAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 TemplateAd"`

	*Ad

	//
	// ID of the template to use.
	// <span class="constraint Selectable">This field can be selected using the value "TemplateId".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	TemplateId int64 `xml:"templateId,omitempty"`

	//
	// Group ID of all template ads, which should be created together.
	// Template ads in the same union reference the same data but have different
	// dimensions. Single ads do not have a union ID. If a template ad specifies
	// an ad union with only one ad, no union will be created.
	// <span class="constraint Selectable">This field can be selected using the value "TemplateAdUnionId".</span>
	//
	AdUnionId *AdUnionId `xml:"adUnionId,omitempty"`

	//
	// List of elements (each containing a set of fields) for the template
	// referenced by {@code templateId}. See
	// <a href="/adwords/api/docs/appendix/templateads">Template
	// Ads</a> for the elements and fields required for each template.
	//
	TemplateElements []*TemplateElement `xml:"templateElements,omitempty"`

	//
	// The template ad rendered as an image.
	//
	AdAsImage *Image `xml:"adAsImage,omitempty"`

	//
	// The dimensions for this ad.
	//
	Dimensions *Dimensions `xml:"dimensions,omitempty"`

	//
	// Name of this ad.
	// <span class="constraint Required">
	// This field is required and should not be {@code null}.</span>
	// <span class="constraint Selectable">This field can be selected using the value "TemplateAdName".</span>
	//
	Name string `xml:"name,omitempty"`

	//
	// Duration of this ad (if it contains playable media).
	// <span class="constraint Selectable">This field can be selected using the value "TemplateAdDuration".</span>
	// <span class="constraint ReadOnly">This field is read only and will be ignored when sent to the API.</span>
	//
	Duration int32 `xml:"duration,omitempty"`

	//
	// For copies, the ad id of the ad this was or should be copied from.
	// <span class="constraint Selectable">This field can be selected using the value "TemplateOriginAdId".</span>
	//
	OriginAdId int64 `xml:"originAdId,omitempty"`
}

type TemplateElement struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 TemplateElement"`

	//
	// Unique name for this element.
	// <span class="constraint Selectable">This field can be selected using the value "UniqueName".</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	UniqueName string `xml:"uniqueName,omitempty"`

	//
	// List of fields to use for this template element.
	// These must be the same for all template ads in the same template ad union.
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Fields []*TemplateElementField `xml:"fields,omitempty"`
}

type TemplateElementField struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 TemplateElementField"`

	//
	// The name of this field.
	// <span class="constraint Selectable">This field can be selected using the value "TemplateElementFieldName".</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Name string `xml:"name,omitempty"`

	//
	// The type of this field.
	// <span class="constraint Selectable">This field can be selected using the value "TemplateElementFieldType".</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	Type_ *TemplateElementFieldType `xml:"type,omitempty"`

	//
	// Text value for text field types. Null if not text field.
	// The field is a text field if type is ADDRESS, ENUM, TEXT, URL,
	// or VISIBLE_URL.
	// <span class="constraint Selectable">This field can be selected using the value "TemplateElementFieldText".</span>
	//
	FieldText string `xml:"fieldText,omitempty"`

	//
	// Media value for non-text field types. Null if a text field. This
	// fields must be specified if fieldText is null.
	//
	FieldMedia *Media `xml:"fieldMedia,omitempty"`
}

type TextAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 TextAd"`

	*Ad

	//
	// The headline of the ad.
	// <span class="constraint Selectable">This field can be selected using the value "Headline".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Headline string `xml:"headline,omitempty"`

	//
	// The first description line.
	// <span class="constraint Selectable">This field can be selected using the value "Description1".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Description1 string `xml:"description1,omitempty"`

	//
	// The second description line.
	// <span class="constraint Selectable">This field can be selected using the value "Description2".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Description2 string `xml:"description2,omitempty"`
}

type ThirdPartyRedirectAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 ThirdPartyRedirectAd"`

	*RichMediaAd

	//
	// Defines whether or not the ad is cookie targeted.
	// (i.e. user list targeting, or the network's equivalent).
	// <span class="constraint Selectable">This field can be selected using the value "IsCookieTargeted".</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	IsCookieTargeted bool `xml:"isCookieTargeted,omitempty"`

	//
	// Defines whether or not the ad is targeting user interest.
	// <span class="constraint Selectable">This field can be selected using the value "IsUserInterestTargeted".</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	IsUserInterestTargeted bool `xml:"isUserInterestTargeted,omitempty"`

	//
	// Defines whether or not the ad contains a tracking pixel of any kind.
	// <span class="constraint Selectable">This field can be selected using the value "IsTagged".</span>
	// <span class="constraint Required">This field is required and should not be {@code null}.</span>
	//
	IsTagged bool `xml:"isTagged,omitempty"`

	//
	// Video Types of the ad. (RealMedia, Quick Time etc.)
	// <span class="constraint Selectable">This field can be selected using the value "VideoTypes".</span>
	//
	VideoTypes []*VideoType `xml:"videoTypes,omitempty"`

	//
	// Allowed expanding directions. These directions are used to match
	// publishers' ad slots. For example, if a slot allows expansion toward the
	// right, only ads with EXPANDING_RIGHT specified will show up there.
	// <span class="constraint Selectable">This field can be selected using the value "ExpandingDirections".</span>
	//
	ExpandingDirections []*ThirdPartyRedirectAdExpandingDirection `xml:"expandingDirections,omitempty"`
}

type UniversalShoppingAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 UniversalShoppingAd"`

	*Ad
}

type UrlData struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 UrlData"`

	//
	// Unique identifier for this instance of UrlData. Refer to the
	// <a href="https://developers.google.com/adwords/api/docs/appendix/templateads">Template
	// Ads documentation</a> for the list of valid values.
	// <span class="constraint Required">This field is required and should not be {@code null} when it is contained within {@link Operator}s : ADD.</span>
	// <span class="constraint StringLength">This string must not be empty, (trimmed).</span>
	//
	UrlId string `xml:"urlId,omitempty"`

	//
	// A list of final landing page urls.
	//
	FinalUrls *UrlList `xml:"finalUrls,omitempty"`

	//
	// A list of final mobile landing page urls.
	//
	FinalMobileUrls *UrlList `xml:"finalMobileUrls,omitempty"`

	//
	// URL template for constructing a tracking URL.
	//
	TrackingUrlTemplate string `xml:"trackingUrlTemplate,omitempty"`
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

type Video struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 Video"`

	*Media

	//
	// The duration of the associated video, in milliseconds.
	// <span class="constraint Selectable">This field can be selected using the value "DurationMillis".</span>
	//
	DurationMillis int64 `xml:"durationMillis,omitempty"`

	//
	// Streaming URL for the video.
	// <span class="constraint Selectable">This field can be selected using the value "StreamingUrl".</span>
	//
	StreamingUrl string `xml:"streamingUrl,omitempty"`

	//
	// Indicates whether the video is ready to play on the web.
	// <span class="constraint Selectable">This field can be selected using the value "ReadyToPlayOnTheWeb".</span>
	//
	ReadyToPlayOnTheWeb bool `xml:"readyToPlayOnTheWeb,omitempty"`

	//
	// The Industry Standard Commercial Identifier code for this media, used
	// mainly for television commercials.
	// <span class="constraint Selectable">This field can be selected using the value "IndustryStandardCommercialIdentifier".</span>
	//
	IndustryStandardCommercialIdentifier string `xml:"industryStandardCommercialIdentifier,omitempty"`

	//
	// The Advertising Digital Identification code for this media, as defined by
	// the American Association of Advertising Agencies, used mainly for
	// television commercials.
	// <span class="constraint Selectable">This field can be selected using the value "AdvertisingId".</span>
	//
	AdvertisingId string `xml:"advertisingId,omitempty"`

	//
	// For YouTube-hosted videos, the YouTube video ID (as seen in YouTube URLs)
	// may also be filled in.
	// <span class="constraint Selectable">This field can be selected using the value "YouTubeVideoIdString".</span>
	//
	YouTubeVideoIdString string `xml:"youTubeVideoIdString,omitempty"`
}

type DynamicSearchAd struct {
	XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201802 DynamicSearchAd"`

	*Ad

	//
	// The first description line.
	// <span class="constraint Selectable">This field can be selected using the value "Description1".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Description1 string `xml:"description1,omitempty"`

	//
	// The second description line.
	// <span class="constraint Selectable">This field can be selected using the value "Description2".</span><span class="constraint Filterable">This field can be filtered on.</span>
	//
	Description2 string `xml:"description2,omitempty"`
}

type AdGroupAdServiceInterface struct {
	client *SOAPClient
}

func NewAdGroupAdServiceInterface(url string, tls bool, auth *BasicAuth) *AdGroupAdServiceInterface {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &AdGroupAdServiceInterface{
		client: client,
	}
}

func NewAdGroupAdServiceInterfaceWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *AdGroupAdServiceInterface {
	if url == "" {
		url = ""
	}
	client := NewSOAPClientWithTLSConfig(url, tlsCfg, auth)

	return &AdGroupAdServiceInterface{
		client: client,
	}
}

func (service *AdGroupAdServiceInterface) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *AdGroupAdServiceInterface) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Error can be either of the following types:
//
//   - ApiException
/*
   Returns a list of AdGroupAds.
   AdGroupAds that had been removed are not returned by default.

   @param serviceSelector The selector specifying the {@link AdGroupAd}s to return.
   @return The page containing the AdGroupAds that meet the criteria specified by the selector.
   @throws ApiException when there is at least one error with the request.
*/
func (service *AdGroupAdServiceInterface) Get(request *Get) (*GetResponse, error) {
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
   Applies the list of mutate operations (ie. add, set, remove):
   <p>Add - Creates a new {@linkplain AdGroupAd ad group ad}. The
   {@code adGroupId} must
   reference an existing ad group. The child {@code Ad} must be sufficiently
   specified by constructing a concrete ad type (such as {@code TextAd})
   and setting its fields accordingly.</p>
   <p>Set - Updates an ad group ad. Except for {@code status},
   ad group ad fields are not mutable. Status updates are
   straightforward - the status of the ad group ad is updated as
   specified. If any other field has changed, it will be ignored. If
   you want to change any of the fields other than status, you must
   make a new ad and then remove the old one.</p>
   <p>Remove - Removes the link between the specified AdGroup and
   Ad.</p>
   @param operations The operations to apply.
   @return A list of AdGroupAds where each entry in the list is the result of
   applying the operation in the input list with the same index. For an
   add/set operation, the return AdGroupAd will be what is saved to the db.
   In the case of the remove operation, the return AdGroupAd will simply be
   an AdGroupAd containing an Ad with the id set to the Ad being removed from
   the AdGroup.
*/
func (service *AdGroupAdServiceInterface) Mutate(request *Mutate) (*MutateResponse, error) {
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
   Adds labels to the AdGroupAd or removes labels from the AdGroupAd.
   <p>Add - Apply an existing label to an existing {@linkplain AdGroupAd ad group ad}. The
   {@code adGroupId} and {@code adId} must reference an existing
   {@linkplain AdGroupAd ad group ad}. The {@code labelId} must reference an existing
   {@linkplain Label label}.
   <p>Remove - Removes the link between the specified {@linkplain AdGroupAd ad group ad} and
   {@linkplain Label label}.
   @param operations The operations to apply.
   @return A list of AdGroupAdLabel where each entry in the list is the result of
   applying the operation in the input list with the same index. For an
   add operation, the returned AdGroupAdLabel contains the AdGroupId, AdId and the LabelId.
   In the case of a remove operation, the returned AdGroupAdLabel will only have AdGroupId and
   AdId.
   @throws ApiException when there are one or more errors with the request.
*/
func (service *AdGroupAdServiceInterface) MutateLabel(request *MutateLabel) (*MutateLabelResponse, error) {
	response := new(MutateLabelResponse)
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
   Returns a list of AdGroupAds based on the query.

   @param query The SQL-like AWQL query string.
   @return A list of AdGroupAds.
   @throws ApiException if problems occur while parsing the query or fetching AdGroupAds.
*/
func (service *AdGroupAdServiceInterface) Query(request *Query) (*QueryResponse, error) {
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