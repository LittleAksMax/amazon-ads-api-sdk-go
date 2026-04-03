# amazon-ads-api-sdk-go

A small Go SDK for a narrow slice of the Amazon Ads API.

This is not a full wrapper over the Amazon Ads surface described in the Amazon Ads API overview. Right now the implemented coverage is mostly:

- Profiles (`v2/profiles`)
- A few *Amazon Ads* resource families: campaigns, ad groups, ads, and targets
- Reporting API v3 async report workflows

The code tries to stay fairly general in its models, but the main implementation focus so far has been the Sponsored Products variants of these endpoints. That is also where the examples and current reporting models are most complete.

Our implementations are based on parts of:
- https://github.com/amazinsellers/amazon-ads-api-sdk-go
- https://github.com/vanling1111/amazon-ads-api-go-sdk

## Coverage at a glance

Relative to the broader Amazon Ads API, this repo currently covers a small and practical subset:

- `GET /v2/profiles`
- `POST /adsApi/v1/query/campaigns`
- `POST /adsApi/v1/query/adGroups`
- `POST /adsApi/v1/update/adGroups`
- `POST /adsApi/v1/query/ads`
- `POST /adsApi/v1/query/targets`
- `POST /adsApi/v1/update/targets`
- `POST /reporting/reports`
- `GET /reporting/reports/{reportId}`
- `DELETE /reporting/reports/{reportId}`
- Downloading and decoding the generated report file once Amazon marks the report complete

Everything else in the wider Amazon Ads API should be treated as unsupported unless it is explicitly wrapped in code here.

## Client

`AmazonAdsAPIClient` is the main Ads API client. It owns:

- Region selection for the Ads API host
- Shared request/header setup
- Automatic token refresh through the auth client
- Service access via `CampaignsService`, `AdGroupsService`, `AdsService`, `TargetsService`, and `ReportsService`
- Direct profile access through `GetProfiles(...)`

Typical setup looks like this:

```go
authCfg := sdk.NewAmazonAuthAPIConfig(clientID, clientSecret, redirectURI)
authClient, err := sdk.NewAmazonAuthClient(authCfg, sdk.AmazonRegions.Europe)
if err != nil {
	panic(err)
}

client, err := sdk.NewAmazonAdsAPIClient(&sdk.Configuration{
	AuthClient: authClient,
	Region:     sdk.AmazonRegions.Europe,
})
if err != nil {
	panic(err)
}

client.SetRefreshToken(refreshToken)
```

## Auth Client

`AmazonAPIAuthClient` is the Login with Amazon helper used by the main client.

It currently supports:

- Exchanging an authorisation code for access/refresh tokens
- Storing access credentials in memory
- Refreshing access tokens from a refresh token
- Basic concurrency protection around token refresh
- Separate auth-region host mapping from the Ads API host

In practice, most callers will either:

- call `ExchangeAuthorisationCode(...)` once and persist the returned refresh token, or
- restore a known refresh token with `SetRefreshToken(...)`

## CampaignsService

Current coverage is query-only:

- `GetCampaigns(profileID, options)` wraps `POST /adsApi/v1/query/campaigns`
- Results come back through the generic paginator

There is no campaign create/update/delete wrapper yet.

The models try to cover more than one ad product, but the examples and intended path here are mainly Sponsored Products campaign queries.

## AdGroupsService

Current coverage:

- `GetAdGroups(profileID, options)` wraps `POST /adsApi/v1/query/adGroups`
- `UpdateAdGroups(ctx, profileID, options)` wraps `POST /adsApi/v1/update/adGroups`

There is no ad group create/delete wrapper yet.

The models try to cover more than one ad product, but the examples and intended path here are mainly Sponsored Products campaign queries.

## AdsService

Current coverage is read-only:

- `GetAds(ctx, profileID, options)` wraps `POST /adsApi/v1/query/ads`

There is no ad create/update/delete wrapper yet.

The ad models include multiple creative shapes, so the types are broader than pure Sponsored Products, but this part of the SDK has not been fleshed out across the wider Ads API surface.

## TargetsService

Current coverage:

- `GetTargets(profileID, options)` wraps `POST /adsApi/v1/query/targets`
- `UpdateTargets(ctx, profileID, options)` wraps `POST /adsApi/v1/update/targets`

There is no target create/delete wrapper yet.

The target model is fairly broad and includes a large `TargetDetails` union, but the practical focus has still been Sponsored Products style targeting/update flows.

## ReportsService

`ReportsService` wraps the async Reporting API v3 flow:

- `RequestReport(...)` creates a report
- `GetReport(...)` fetches a single report by ID
- `CancelReport(...)` cancels a report
- `Report.Refresh(...)` polls the latest status
- `Report.GeneratedReport(...)` downloads the completed file

The helper also handles gzip decompression for `GZIP_JSON` report downloads and exposes a simple `Decode(...)` method on the downloaded report body.

## Profiles API coverage

Profiles are currently exposed directly on the main client via `GetProfiles(ctx, options)`.

We only cover `GET /v2/profiles`, and none of the other endpoints.

## Reporting API v3 coverage

Relative to the Reporting v3 docs, this repo covers the core async report lifecycle, but only a narrow slice of the model surface.

What is covered well:

- Create async report
- Poll/fetch a single report
- Cancel a report
- Download the completed report payload
- Decode the generated JSON payload

What is only partial right now:

- Report type enums and helper constants
- Group-by/filter convenience types
- Broader ad-product coverage

The current reporting models are especially Sponsored Products-oriented. For example, the baked-in report type IDs are currently:

- `spCampaigns`
- `spTargeting`
- `spSearchTerm`

## Pagination

The query endpoints for campaigns, ad groups, and targets use a small generic paginator built around Amazon's `nextToken` pattern.

Collect everything:

```go
campaigns, err := client.CampaignsService.
	GetCampaigns(profileID, opts).
	Collect(ctx)
```

Or iterate page by page:

```go
p := client.TargetsService.GetTargets(profileID, opts)
for p.HasNext() {
	page, err := p.Next(ctx)
	if err != nil {
		break
	}
	_ = page
}
```

It is intentionally simple: fetch one page, parse the top-level `nextToken`, and stop when Amazon stops returning one.

## References

- Amazon Ads API overview: https://advertising.amazon.com/API/docs/en-us/reference/amazon-ads/overview
- Profiles v2 docs: https://advertising.amazon.com/API/docs/en-us/reference/2/profiles
- Reporting API v3 overview: https://advertising.amazon.com/API/docs/en-us/guides/reporting/v3/overview
