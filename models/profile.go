package models

import (
	url2 "net/url"
	"strings"
)

/*
  - https://advertising.amazon.com/API/docs/en-us/reference/2/profiles#tag/Profiles/operation/listProfiles
    {
    "profileId": 0,
    "countryCode": "BR",
    "currencyCode": "BRL",
    "dailyBudget": 0,
    "timezone": "Africa/Cairo",
    "accountInfo": {
    "marketplaceStringId": "string",
    "id": "string",
    "type": "vendor",
    "name": "string",
    "subType": "KDP_AUTHOR",
    "validPaymentMethod": true
    }
    }
*/

// Profile /
type Profile struct {
	ProfileID    int64              `json:"profileId"`
	CountryCode  string             `json:"countryCode"`
	CurrencyCode string             `json:"currencyCode"`
	DailyBudget  float64            `json:"dailyBudget"`
	TimeZone     string             `json:"timezone"`
	AccountInfo  ProfileAccountInfo `json:"accountInfo"`
}

type ProfileAccountInfo struct {
	MarketplaceStringID string `json:"marketplaceStringId"`
	ID                  string `json:"id"`
	Type                string `json:"type"`
	Name                string `json:"name"`
	SubType             string `json:"subType"`
	ValidPaymentMethod  bool   `json:"validPaymentMethod"`
}

type ListProfilesOptions struct {
	ApiProgram               string   // "billing", "campaign", "paymentMethod", "store", "report", "account", "posts" -- default is "campaign"
	AccessLevel              string   // "view", "edit" -- default is "edit"
	ProfileTypeFilter        []string // choice of "seller", "vendor", "agency" -- default is all
	ValidPaymentMethodFilter string   // "true", "false" -- default is all
}

func (lpo *ListProfilesOptions) ToQuery() url2.Values {
	queryParameters := url2.Values{}

	if lpo.ApiProgram != "" {
		queryParameters.Add("apiProgram", lpo.ApiProgram)
	}
	if lpo.AccessLevel != "" {
		queryParameters.Add("accessLevel", lpo.AccessLevel)
	}
	if lpo.ProfileTypeFilter != nil {
		// We have to make the comma separated list ourselves unfortunately
		queryParameters.Add("profileTypeFilter", strings.Join(lpo.ProfileTypeFilter, ","))
	}
	if lpo.ValidPaymentMethodFilter != "" {
		queryParameters.Add("validPaymentMethodFilter", lpo.ValidPaymentMethodFilter)
	}
	return queryParameters
}
