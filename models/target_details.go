package models

// TargetDetails represents the union of all possible targeting detail types.
// Only one of the fields will be non-nil depending on the target type.
type TargetDetails struct {
	AdInitiationTarget             *AdInitiationTarget             `json:"adInitiationTarget"`
	AdPlayerSizeTarget             *AdPlayerSizeTarget             `json:"adPlayerSizeTarget"`
	AppTarget                      *AppTarget                      `json:"appTarget"`
	AudienceTarget                 *AudienceTarget                 `json:"audienceTarget"`
	BrandSafetyCategoryTarget      *BrandSafetyCategoryTarget      `json:"brandSafetyCategoryTarget"`
	BrandSafetyTierTarget          *BrandSafetyTierTarget          `json:"brandSafetyTierTarget"`
	ContentCategoryTarget          *ContentCategoryTarget          `json:"contentCategoryTarget"`
	ContentGenreTarget             *ContentGenreTarget             `json:"contentGenreTarget"`
	ContentInstreamPositionTarget  *ContentInstreamPositionTarget  `json:"contentInstreamPositionTarget"`
	ContentOutstreamPositionTarget *ContentOutstreamPositionTarget `json:"contentOutstreamPositionTarget"`
	ContentRatingTarget            *ContentRatingTarget            `json:"contentRatingTarget"`
	DayPartTarget                  *DayPartTarget                  `json:"dayPartTarget"`
	DeviceTarget                   *DeviceTarget                   `json:"deviceTarget"`
	DomainTarget                   *DomainTarget                   `json:"domainTarget"`
	FoldPositionTarget             *FoldPositionTarget             `json:"foldPositionTarget"`
	InventorySourceTarget          *InventorySourceTarget          `json:"inventorySourceTarget"`
	KeywordTarget                  *KeywordTargetDetails           `json:"keywordTarget"`
	LocationTarget                 *LocationTarget                 `json:"locationTarget"`
	NativeContentPositionTarget    *NativeContentPositionTarget    `json:"nativeContentPositionTarget"`
	PlacementTypeTarget            *PlacementTypeTarget            `json:"placementTypeTarget"`
	ProductAudienceTarget          *ProductAudienceTarget          `json:"productAudienceTarget"`
	ProductCategoryTarget          *ProductCategoryTarget          `json:"productCategoryTarget"`
	ProductTarget                  *ProductTargetDetails           `json:"productTarget"`
	ThemeTarget                    *ThemeTargetDetails             `json:"themeTarget"`
	ThirdPartyTarget               *ThirdPartyTarget               `json:"thirdPartyTarget"`
	VideoAdFormatTarget            *VideoAdFormatTarget            `json:"videoAdFormatTarget"`
	VideoContentDurationTarget     *VideoContentDurationTarget     `json:"videoContentDurationTarget"`
}

// AdInitiationTarget represents targeting by how the ad is initiated.
type AdInitiationTarget struct {
	AdInitiations []AdInitiationType `json:"adInitiations"`
}

type AdInitiationType string

const (
	AdInitiationAutoPlay     AdInitiationType = "AUTO_PLAY"
	AdInitiationClickToPlay  AdInitiationType = "CLICK_TO_PLAY"
	AdInitiationAutoPlayMute AdInitiationType = "AUTO_PLAY_MUTE"
)

// AdPlayerSizeTarget represents targeting by ad player size.
type AdPlayerSizeTarget struct {
	PlayerSizes []AdPlayerSize `json:"playerSizes"`
}

type AdPlayerSize string

const (
	AdPlayerSizeSmall      AdPlayerSize = "SMALL"
	AdPlayerSizeMedium     AdPlayerSize = "MEDIUM"
	AdPlayerSizeLarge      AdPlayerSize = "LARGE"
	AdPlayerSizeExtraLarge AdPlayerSize = "EXTRA_LARGE"
	AdPlayerSizeUnknown    AdPlayerSize = "UNKNOWN"
)

// AppTarget represents targeting by application.
type AppTarget struct {
	Apps []AppTargetValue `json:"apps"`
}

type AppTargetValue struct {
	AppID   string `json:"appId"`
	AppName string `json:"appName"`
}

// AudienceTarget represents targeting based on audience segments.
type AudienceTarget struct {
	Audiences     []AudienceSegment `json:"audiences"`
	LookbackDays  int               `json:"lookbackDays"`
	AudienceMatch AudienceMatchType `json:"audienceMatch"`
	Recency       *AudienceRecency  `json:"recency"`
}

type AudienceMatchType string

const (
	AudienceMatchExact   AudienceMatchType = "EXACT"
	AudienceMatchSimilar AudienceMatchType = "SIMILAR"
)

type AudienceRecency struct {
	Days int `json:"days"`
}

type AudienceSegment struct {
	AudienceID string `json:"audienceId"`
}

// BrandSafetyCategoryTarget represents targeting by brand safety category.
type BrandSafetyCategoryTarget struct {
	Categories []BrandSafetyCategory `json:"categories"`
}

type BrandSafetyCategory struct {
	CategoryID   string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

// BrandSafetyTierTarget represents targeting by brand safety tier.
type BrandSafetyTierTarget struct {
	Tier BrandSafetyTier `json:"tier"`
}

type BrandSafetyTier string

const (
	BrandSafetyTierHigh     BrandSafetyTier = "HIGH"
	BrandSafetyTierMedium   BrandSafetyTier = "MEDIUM"
	BrandSafetyTierLow      BrandSafetyTier = "LOW"
	BrandSafetyTierStandard BrandSafetyTier = "STANDARD"
)

// ContentCategoryTarget represents targeting by content category.
type ContentCategoryTarget struct {
	Categories []ContentCategory `json:"categories"`
}

type ContentCategory struct {
	CategoryID   string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

// ContentGenreTarget represents targeting by content genre.
type ContentGenreTarget struct {
	Genres []ContentGenre `json:"genres"`
}

type ContentGenre struct {
	GenreID   string `json:"genreId"`
	GenreName string `json:"genreName"`
}

// ContentInstreamPositionTarget represents targeting by instream ad position.
type ContentInstreamPositionTarget struct {
	Positions []ContentInstreamPosition `json:"positions"`
}

type ContentInstreamPosition string

const (
	ContentInstreamPositionPreRoll  ContentInstreamPosition = "PRE_ROLL"
	ContentInstreamPositionMidRoll  ContentInstreamPosition = "MID_ROLL"
	ContentInstreamPositionPostRoll ContentInstreamPosition = "POST_ROLL"
)

// ContentOutstreamPositionTarget represents targeting by outstream ad position.
type ContentOutstreamPositionTarget struct {
	Positions []ContentOutstreamPosition `json:"positions"`
}

type ContentOutstreamPosition string

const (
	ContentOutstreamPositionInArticle    ContentOutstreamPosition = "IN_ARTICLE"
	ContentOutstreamPositionInFeed       ContentOutstreamPosition = "IN_FEED"
	ContentOutstreamPositionInBanner     ContentOutstreamPosition = "IN_BANNER"
	ContentOutstreamPositionInterstitial ContentOutstreamPosition = "INTERSTITIAL"
)

// ContentRatingTarget represents targeting by content rating.
type ContentRatingTarget struct {
	Ratings []ContentRating `json:"ratings"`
}

type ContentRating string

const (
	ContentRatingG       ContentRating = "G"
	ContentRatingPG      ContentRating = "PG"
	ContentRatingPG13    ContentRating = "PG13"
	ContentRatingT       ContentRating = "T"
	ContentRatingMA      ContentRating = "MA"
	ContentRatingR       ContentRating = "R"
	ContentRatingUnrated ContentRating = "UNRATED"
)

// DayPartTarget represents targeting by day and time (dayparting).
type DayPartTarget struct {
	DayParts []DayPart `json:"dayParts"`
}

type DayPart struct {
	Day       DayOfWeek `json:"day"`
	StartHour int       `json:"startHour"`
	EndHour   int       `json:"endHour"`
}

type DayOfWeek string

const (
	DayMonday    DayOfWeek = "MONDAY"
	DayTuesday   DayOfWeek = "TUESDAY"
	DayWednesday DayOfWeek = "WEDNESDAY"
	DayThursday  DayOfWeek = "THURSDAY"
	DayFriday    DayOfWeek = "FRIDAY"
	DaySaturday  DayOfWeek = "SATURDAY"
	DaySunday    DayOfWeek = "SUNDAY"
)

// DeviceTarget represents targeting by device type.
type DeviceTarget struct {
	Devices []DeviceType `json:"devices"`
}

type DeviceType string

const (
	DeviceDesktop      DeviceType = "DESKTOP"
	DeviceMobile       DeviceType = "MOBILE"
	DeviceTablet       DeviceType = "TABLET"
	DeviceConnectedTV  DeviceType = "CONNECTED_TV"
	DeviceGameConsole  DeviceType = "GAME_CONSOLE"
	DeviceSetTopBox    DeviceType = "SET_TOP_BOX"
	DeviceSmartSpeaker DeviceType = "SMART_SPEAKER"
)

// DomainTarget represents targeting by domain/website.
type DomainTarget struct {
	Domains []DomainValue `json:"domains"`
}

type DomainValue struct {
	Domain string `json:"domain"`
}

// FoldPositionTarget represents targeting by fold position.
type FoldPositionTarget struct {
	Positions []FoldPosition `json:"positions"`
}

type FoldPosition string

const (
	FoldPositionAbove   FoldPosition = "ABOVE_THE_FOLD"
	FoldPositionBelow   FoldPosition = "BELOW_THE_FOLD"
	FoldPositionUnknown FoldPosition = "UNKNOWN"
)

// InventorySourceTarget represents targeting by inventory source.
type InventorySourceTarget struct {
	Sources []InventorySource `json:"sources"`
}

type InventorySource struct {
	SourceID   string `json:"sourceId"`
	SourceName string `json:"sourceName"`
}

// KeywordTargetDetails represents keyword-based targeting.
type KeywordTargetDetails struct {
	Keyword               string           `json:"keyword"`
	MatchType             KeywordMatchType `json:"matchType"`
	NativeLanguageKeyword *string          `json:"nativeLanguageKeyword"`
	NativeLanguageLocale  *string          `json:"nativeLanguageLocale"`
}

type KeywordMatchType string

const (
	KeywordMatchBroad          KeywordMatchType = "BROAD"
	KeywordMatchExact          KeywordMatchType = "EXACT"
	KeywordMatchPhrase         KeywordMatchType = "PHRASE"
	KeywordMatchNegativeExact  KeywordMatchType = "NEGATIVE_EXACT"
	KeywordMatchNegativePhrase KeywordMatchType = "NEGATIVE_PHRASE"
	KeywordMatchNegativeBroad  KeywordMatchType = "NEGATIVE_BROAD"
)

// LocationTarget represents geographic location targeting.
type LocationTarget struct {
	Locations []LocationValue `json:"locations"`
}

type LocationValue struct {
	LocationID   string       `json:"locationId"`
	LocationName string       `json:"locationName"`
	LocationType LocationType `json:"locationType"`
}

type LocationType string

const (
	LocationTypeCountry    LocationType = "COUNTRY"
	LocationTypeState      LocationType = "STATE"
	LocationTypeCity       LocationType = "CITY"
	LocationTypeDMA        LocationType = "DMA"
	LocationTypePostalCode LocationType = "POSTAL_CODE"
)

// NativeContentPositionTarget represents targeting by native content position.
type NativeContentPositionTarget struct {
	Positions []NativeContentPosition `json:"positions"`
}

type NativeContentPosition string

const (
	NativeContentPositionInFeed         NativeContentPosition = "IN_FEED"
	NativeContentPositionRecommendation NativeContentPosition = "RECOMMENDATION"
	NativeContentPositionInArticle      NativeContentPosition = "IN_ARTICLE"
)

// PlacementTypeTarget represents targeting by placement type.
type PlacementTypeTarget struct {
	PlacementTypes []TargetPlacementType `json:"placementTypes"`
}

type TargetPlacementType string

const (
	TargetPlacementTypeWebDisplay    TargetPlacementType = "WEB_DISPLAY"
	TargetPlacementTypeWebVideo      TargetPlacementType = "WEB_VIDEO"
	TargetPlacementTypeAppDisplay    TargetPlacementType = "APP_DISPLAY"
	TargetPlacementTypeAppVideo      TargetPlacementType = "APP_VIDEO"
	TargetPlacementTypeInStreamVideo TargetPlacementType = "IN_STREAM_VIDEO"
)

// ProductAudienceTarget represents targeting by product audience.
type ProductAudienceTarget struct {
	Audiences     []ProductAudience `json:"audiences"`
	LookbackDays  int               `json:"lookbackDays"`
	AudienceMatch AudienceMatchType `json:"audienceMatch"`
}

type ProductAudience struct {
	AudienceID string `json:"audienceId"`
	ASIN       string `json:"asin"`
}

// ProductCategoryTarget represents targeting by product category.
type ProductCategoryTarget struct {
	Expression         []ProductCategoryExpression `json:"expression"`
	ResolvedExpression []ProductCategoryExpression `json:"resolvedExpression"`
}

type ProductCategoryExpression struct {
	Type  ProductCategoryExpressionType `json:"type"`
	Value string                        `json:"value"`
}

type ProductCategoryExpressionType string

const (
	ProductCategoryExpressionASINCategorySameAs      ProductCategoryExpressionType = "ASIN_CATEGORY_SAME_AS"
	ProductCategoryExpressionASINBrandSameAs         ProductCategoryExpressionType = "ASIN_BRAND_SAME_AS"
	ProductCategoryExpressionASINPriceBetween        ProductCategoryExpressionType = "ASIN_PRICE_BETWEEN"
	ProductCategoryExpressionASINPriceGreaterThan    ProductCategoryExpressionType = "ASIN_PRICE_GREATER_THAN"
	ProductCategoryExpressionASINPriceLessThan       ProductCategoryExpressionType = "ASIN_PRICE_LESS_THAN"
	ProductCategoryExpressionASINReviewRatingBetween ProductCategoryExpressionType = "ASIN_REVIEW_RATING_BETWEEN"
	ProductCategoryExpressionASINReviewRatingGreater ProductCategoryExpressionType = "ASIN_REVIEW_RATING_GREATER_THAN"
	ProductCategoryExpressionASINReviewRatingLess    ProductCategoryExpressionType = "ASIN_REVIEW_RATING_LESS_THAN"
	ProductCategoryExpressionASINIsPrimeShipping     ProductCategoryExpressionType = "ASIN_IS_PRIME_SHIPPING_ELIGIBLE"
)

// ProductTargetDetails represents product-based targeting (ASIN, category, brand, etc.).
type ProductTargetDetails struct {
	Expression         []ProductExpression `json:"expression"`
	ResolvedExpression []ProductExpression `json:"resolvedExpression"`
}

type ProductExpression struct {
	Type  ProductExpressionType `json:"type"`
	Value string                `json:"value"`
}

type ProductExpressionType string

const (
	ProductExpressionASINBrandSameAs         ProductExpressionType = "ASIN_BRAND_SAME_AS"
	ProductExpressionASINCategorySameAs      ProductExpressionType = "ASIN_CATEGORY_SAME_AS"
	ProductExpressionASINGenreSameAs         ProductExpressionType = "ASIN_GENRE_SAME_AS"
	ProductExpressionASINIsPrimeShipping     ProductExpressionType = "ASIN_IS_PRIME_SHIPPING_ELIGIBLE"
	ProductExpressionASINPriceBetween        ProductExpressionType = "ASIN_PRICE_BETWEEN"
	ProductExpressionASINPriceGreaterThan    ProductExpressionType = "ASIN_PRICE_GREATER_THAN"
	ProductExpressionASINPriceLessThan       ProductExpressionType = "ASIN_PRICE_LESS_THAN"
	ProductExpressionASINReviewRatingBetween ProductExpressionType = "ASIN_REVIEW_RATING_BETWEEN"
	ProductExpressionASINReviewRatingGreater ProductExpressionType = "ASIN_REVIEW_RATING_GREATER_THAN"
	ProductExpressionASINReviewRatingLess    ProductExpressionType = "ASIN_REVIEW_RATING_LESS_THAN"
	ProductExpressionASINSameAs              ProductExpressionType = "ASIN_SAME_AS"
	ProductExpressionASINSubstituteRelated   ProductExpressionType = "ASIN_SUBSTITUTE_RELATED"
	ProductExpressionASINAccessoryRelated    ProductExpressionType = "ASIN_ACCESSORY_RELATED"
)

// ThemeTargetDetails represents theme-based targeting.
type ThemeTargetDetails struct {
	Theme     ThemeTargetValue `json:"theme"`
	MatchType ThemeMatchType   `json:"matchType"`
}

type ThemeTargetValue struct {
	Value string `json:"value"`
}

type ThemeMatchType string

const (
	ThemeMatchBroad  ThemeMatchType = "BROAD"
	ThemeMatchExact  ThemeMatchType = "EXACT"
	ThemeMatchPhrase ThemeMatchType = "PHRASE"
)

// ThirdPartyTarget represents targeting by third-party data segments.
type ThirdPartyTarget struct {
	Segments []ThirdPartySegment `json:"segments"`
}

type ThirdPartySegment struct {
	SegmentID   string `json:"segmentId"`
	SegmentName string `json:"segmentName"`
	Provider    string `json:"provider"`
}

// VideoAdFormatTarget represents targeting by video ad format.
type VideoAdFormatTarget struct {
	Formats []VideoAdFormat `json:"formats"`
}

type VideoAdFormat string

const (
	VideoAdFormatInStream      VideoAdFormat = "IN_STREAM"
	VideoAdFormatOutStream     VideoAdFormat = "OUT_STREAM"
	VideoAdFormatInterstitial  VideoAdFormat = "INTERSTITIAL"
	VideoAdFormatRewardedVideo VideoAdFormat = "REWARDED_VIDEO"
)

// VideoContentDurationTarget represents targeting by video content duration.
type VideoContentDurationTarget struct {
	Durations []VideoContentDuration `json:"durations"`
}

type VideoContentDuration string

const (
	VideoContentDurationShort   VideoContentDuration = "SHORT_FORM"
	VideoContentDurationLong    VideoContentDuration = "LONG_FORM"
	VideoContentDurationUnknown VideoContentDuration = "UNKNOWN"
)
