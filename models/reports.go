package models

import (
	"encoding/json"
	"time"
)

type ReportTypeID string

const (
	ReportTypeSponsoredProductsCampaigns  ReportTypeID = "spCampaigns"
	ReportTypeSponsoredProductsTargeting  ReportTypeID = "spTargeting"
	ReportTypeSponsoredProductsSearchTerm ReportTypeID = "spSearchTerm"
)

var AllReportTypeIDs = []ReportTypeID{
	ReportTypeSponsoredProductsCampaigns,
	ReportTypeSponsoredProductsTargeting,
	ReportTypeSponsoredProductsSearchTerm,
}

func (rt ReportTypeID) String() string {
	return string(rt)
}

func (rt ReportTypeID) IsValid() bool {
	for _, valid := range AllReportTypeIDs {
		if rt == valid {
			return true
		}
	}
	return false
}

type ReportTimeUnit string

const (
	ReportTimeUnitSummary ReportTimeUnit = "SUMMARY"
	ReportTimeUnitDaily   ReportTimeUnit = "DAILY"
)

var AllReportTimeUnits = []ReportTimeUnit{
	ReportTimeUnitSummary,
	ReportTimeUnitDaily,
}

func (rtu ReportTimeUnit) String() string {
	return string(rtu)
}

func (rtu ReportTimeUnit) IsValid() bool {
	for _, valid := range AllReportTimeUnits {
		if rtu == valid {
			return true
		}
	}
	return false
}

type ReportFormat string

const (
	ReportFormatGZIPJSON ReportFormat = "GZIP_JSON"
)

var AllReportFormats = []ReportFormat{
	ReportFormatGZIPJSON,
}

func (rf ReportFormat) String() string {
	return string(rf)
}

func (rf ReportFormat) IsValid() bool {
	for _, valid := range AllReportFormats {
		if rf == valid {
			return true
		}
	}
	return false
}

type ReportGroupBy string

const (
	ReportGroupByCampaign          ReportGroupBy = "campaign"
	ReportGroupByAdGroup           ReportGroupBy = "adGroup"
	ReportGroupByCampaignPlacement ReportGroupBy = "campaignPlacement"
	ReportGroupByTargeting         ReportGroupBy = "targeting"
	ReportGroupBySearchTerm        ReportGroupBy = "searchTerm"
)

var AllReportGroupBy = []ReportGroupBy{
	ReportGroupByCampaign,
	ReportGroupByAdGroup,
	ReportGroupByCampaignPlacement,
	ReportGroupByTargeting,
	ReportGroupBySearchTerm,
}

func (rgb ReportGroupBy) String() string {
	return string(rgb)
}

func (rgb ReportGroupBy) IsValid() bool {
	for _, valid := range AllReportGroupBy {
		if rgb == valid {
			return true
		}
	}
	return false
}

type ReportFilterField string

const (
	ReportFilterFieldCampaignStatus  ReportFilterField = "campaignStatus"
	ReportFilterFieldAdStatus        ReportFilterField = "adStatus"
	ReportFilterFieldKeywordType     ReportFilterField = "keywordType"
	ReportFilterFieldAdKeywordStatus ReportFilterField = "adKeywordStatus"
)

func (rff ReportFilterField) String() string {
	return string(rff)
}

type ReportKeywordType string

const (
	ReportKeywordTypeBroad                         ReportKeywordType = "BROAD"
	ReportKeywordTypePhrase                        ReportKeywordType = "PHRASE"
	ReportKeywordTypeExact                         ReportKeywordType = "EXACT"
	ReportKeywordTypeTargetingExpression           ReportKeywordType = "TARGETING_EXPRESSION"
	ReportKeywordTypeTargetingExpressionPredefined ReportKeywordType = "TARGETING_EXPRESSION_PREDEFINED"
)

func (rkt ReportKeywordType) String() string {
	return string(rkt)
}

type RequestReportOptions struct {
	Name          string              `json:"name,omitempty"`
	StartDate     string              `json:"startDate"`
	EndDate       string              `json:"endDate"`
	Configuration ReportConfiguration `json:"configuration"`
}

type ReportConfiguration struct {
	AdProduct    AdProduct       `json:"adProduct"`
	GroupBy      []ReportGroupBy `json:"groupBy"`
	Columns      []string        `json:"columns"`
	Filters      []ReportFilter  `json:"filters,omitempty"`
	ReportTypeID ReportTypeID    `json:"reportTypeId"`
	TimeUnit     ReportTimeUnit  `json:"timeUnit"`
	Format       ReportFormat    `json:"format"`
}

type ReportFilter struct {
	Field  ReportFilterField `json:"field"`
	Values []string          `json:"values"`
}

type ReportStatus string

const (
	ReportStatusPending    ReportStatus = "PENDING"
	ReportStatusProcessing ReportStatus = "PROCESSING"
	ReportStatusCompleted  ReportStatus = "COMPLETED"
	ReportStatusFailed     ReportStatus = "FAILED"
)

var AllReportStatuses = []ReportStatus{
	ReportStatusPending,
	ReportStatusProcessing,
	ReportStatusCompleted,
	ReportStatusFailed,
}

func (rs ReportStatus) String() string {
	return string(rs)
}

func (rs ReportStatus) IsValid() bool {
	for _, valid := range AllReportStatuses {
		if rs == valid {
			return true
		}
	}
	return false
}

type ReportDetails struct {
	Configuration ReportConfiguration `json:"configuration"`
	CreatedAt     string              `json:"createdAt"`
	EndDate       string              `json:"endDate"`
	FailureReason *string             `json:"failureReason"`
	FileSize      *int64              `json:"fileSize"`
	GeneratedAt   *string             `json:"generatedAt"`
	Name          string              `json:"name"`
	ReportID      string              `json:"reportId"`
	StartDate     string              `json:"startDate"`
	Status        ReportStatus        `json:"status"`
	UpdatedAt     string              `json:"updatedAt"`
	URL           *string             `json:"url"`
	URLExpiresAt  *string             `json:"urlExpiresAt"`
}

func (rd ReportDetails) IsPending() bool {
	return rd.Status == ReportStatusPending
}

func (rd ReportDetails) IsProcessing() bool {
	return rd.Status == ReportStatusProcessing
}

func (rd ReportDetails) IsCompleted() bool {
	return rd.Status == ReportStatusCompleted
}

func (rd ReportDetails) IsFailed() bool {
	return rd.Status == ReportStatusFailed
}

func (rd ReportDetails) IsTerminal() bool {
	return rd.IsCompleted() || rd.IsFailed()
}

func (rd ReportDetails) HasDownloadURL() bool {
	return rd.URL != nil && *rd.URL != ""
}

type GeneratedReport struct {
	URL  string          `json:"-"`
	Body json.RawMessage `json:"-"`
}

func (gr *GeneratedReport) Decode(v interface{}) error {
	return json.Unmarshal(gr.Body, v)
}

func NewReportFilter(field ReportFilterField, values ...string) ReportFilter {
	return ReportFilter{
		Field:  field,
		Values: append([]string(nil), values...),
	}
}

const requiredDateFormat = "2006-01-02"

func FormatDate(date time.Time) string {
	return date.Format(requiredDateFormat)
}

/////////////////////////////////////
// THESE ARE THE AVAILABLE COLUMNS //
/////////////////////////////////////
// https://advertising.amazon.com/API/docs/en-us/offline-report-prod-3p#tag/Asynchronous-Reports/operation/createAsyncReport
// "adGroupId",
// "adGroupName",
// "advertisedAsin",
// "advertisedSku",
// "campaignBudgetCurrencyCode",
// "campaignId",
// "campaignName",
// "date",
// "keyword",
// "keywordId",
// "keywordType",
// "kindleEditionNormalizedPagesRead14d",
// "kindleEditionNormalizedPagesRoyalties14d",
// "matchType",
// "portfolioId",
// "purchasedAsin",
// "purchases14d",
// "purchases1d",
// "purchases30d",
// "purchases7d",
// "purchasesOtherSku14d",
// "purchasesOtherSku1d",
// "purchasesOtherSku30d",
// "purchasesOtherSku7d",
// "sales14d",
// "sales1d",
// "sales30d",
// "sales7d",
// "salesOtherSku14d",
// "salesOtherSku1d",
// "salesOtherSku30d",
// "salesOtherSku7d",
// "unitsSoldClicks14d",
// "unitsSoldClicks1d",
// "unitsSoldClicks30d",
// "unitsSoldClicks7d",
// "unitsSoldOtherSku14d",
// "unitsSoldOtherSku1d",
// "unitsSoldOtherSku30d",
// "unitsSoldOtherSku7d"
