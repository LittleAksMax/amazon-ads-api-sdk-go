package amazon_ads_api_go_sdk

//////////////////////////////////////////////////////////
//  THESE ARE THE URLS FOR THE AMAZON ADVERTISING API   //
//////////////////////////////////////////////////////////

var amazonAdsApiRegionURLs = struct {
	NorthAmerica string
	Europe       string
	FarEast      string
}{
	NorthAmerica: "advertising-api.amazon.com",
	Europe:       "advertising-api-eu.amazon.com",
	FarEast:      "advertising-api-fe.amazon.com",
}

var amazonAdsApiRegionToURLMap = map[string]string{
	AmazonRegions.Europe:       amazonAdsApiRegionURLs.Europe,
	AmazonRegions.NorthAmerica: amazonAdsApiRegionURLs.NorthAmerica,
	AmazonRegions.FarEast:      amazonAdsApiRegionURLs.FarEast,
}

//////////////////////////////////////////////////////////
// THESE ARE THE URLS FOR THE AMAZON AUTHENTICATION API //
//////////////////////////////////////////////////////////

var amazonAuthApiRegionURLs = struct {
	NorthAmerica string
	Europe       string
	FarEast      string
}{
	NorthAmerica: "api.amazon.com",
	Europe:       "api.amazon.co.uk",
	FarEast:      "api.amazon.co.jp",
}

var amazonAuthApiRegionToURLMap = map[string]string{
	AmazonRegions.NorthAmerica: amazonAuthApiRegionURLs.NorthAmerica,
	AmazonRegions.Europe:       amazonAuthApiRegionURLs.Europe,
	AmazonRegions.FarEast:      amazonAuthApiRegionURLs.FarEast,
}
