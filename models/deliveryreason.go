package models

// DeliveryReason represents reasons for delivery status
type DeliveryReason string

const (
	DeliveryReasonAdvertiserArchived         DeliveryReason = "ADVERTISER_ARCHIVED"
	DeliveryReasonAdvertiserOutOfBudget      DeliveryReason = "ADVERTISER_OUT_OF_BUDGET"
	DeliveryReasonAdvertiserPaymentFailure   DeliveryReason = "ADVERTISER_PAYMENT_FAILURE"
	DeliveryReasonCampaignArchived           DeliveryReason = "CAMPAIGN_ARCHIVED"
	DeliveryReasonCampaignDraftStatus        DeliveryReason = "CAMPAIGN_DRAFT_STATUS"
	DeliveryReasonCampaignEnded              DeliveryReason = "CAMPAIGN_ENDED"
	DeliveryReasonCampaignIncomplete         DeliveryReason = "CAMPAIGN_INCOMPLETE"
	DeliveryReasonCampaignOutOfBudget        DeliveryReason = "CAMPAIGN_OUT_OF_BUDGET"
	DeliveryReasonCampaignPaused             DeliveryReason = "CAMPAIGN_PAUSED"
	DeliveryReasonCampaignPendingStartDate   DeliveryReason = "CAMPAIGN_PENDING_START_DATE"
	DeliveryReasonAdGroupArchived            DeliveryReason = "AD_GROUP_ARCHIVED"
	DeliveryReasonAdGroupEnded               DeliveryReason = "AD_GROUP_ENDED"
	DeliveryReasonAdGroupIncomplete          DeliveryReason = "AD_GROUP_INCOMPLETE"
	DeliveryReasonAdGroupPaused              DeliveryReason = "AD_GROUP_PAUSED"
	DeliveryReasonAdGroupPendingStartDate    DeliveryReason = "AD_GROUP_PENDING_START_DATE"
	DeliveryReasonAdArchived                 DeliveryReason = "AD_ARCHIVED"
	DeliveryReasonAdIncomplete               DeliveryReason = "AD_INCOMPLETE"
	DeliveryReasonAdPaused                   DeliveryReason = "AD_PAUSED"
	DeliveryReasonAdPolicing                 DeliveryReason = "AD_POLICING"
	DeliveryReasonAdRejected                 DeliveryReason = "AD_REJECTED"
	DeliveryReasonAdUnderReview              DeliveryReason = "AD_UNDER_REVIEW"
	DeliveryReasonAccountOutOfBudget         DeliveryReason = "ACCOUNT_OUT_OF_BUDGET"
	DeliveryReasonBiddingControlsIneffective DeliveryReason = "BIDDING_CONTROLS_INEFFECTIVE"
	DeliveryReasonBudgetControlsIneffective  DeliveryReason = "BUDGET_CONTROLS_INEFFECTIVE"
	DeliveryReasonKeywordArchived            DeliveryReason = "KEYWORD_ARCHIVED"
	DeliveryReasonKeywordPaused              DeliveryReason = "KEYWORD_PAUSED"
	DeliveryReasonNoValidPaymentMethod       DeliveryReason = "NO_VALID_PAYMENT_METHOD"
	DeliveryReasonPortfolioEnded             DeliveryReason = "PORTFOLIO_ENDED"
	DeliveryReasonPortfolioPaused            DeliveryReason = "PORTFOLIO_PAUSED"
	DeliveryReasonTargetingClauseArchived    DeliveryReason = "TARGETING_CLAUSE_ARCHIVED"
	DeliveryReasonTargetingClausePaused      DeliveryReason = "TARGETING_CLAUSE_PAUSED"
	DeliveryReasonTargetingClausePolicing    DeliveryReason = "TARGETING_CLAUSE_POLICING"
	DeliveryReasonTargetingClauseRejected    DeliveryReason = "TARGETING_CLAUSE_REJECTED"
	DeliveryReasonTargetingClauseUnderReview DeliveryReason = "TARGETING_CLAUSE_UNDER_REVIEW"
)

var AllDeliveryReasons = []DeliveryReason{
	DeliveryReasonAdvertiserArchived,
	DeliveryReasonAdvertiserOutOfBudget,
	DeliveryReasonAdvertiserPaymentFailure,
	DeliveryReasonCampaignArchived,
	DeliveryReasonCampaignDraftStatus,
	DeliveryReasonCampaignEnded,
	DeliveryReasonCampaignIncomplete,
	DeliveryReasonCampaignOutOfBudget,
	DeliveryReasonCampaignPaused,
	DeliveryReasonCampaignPendingStartDate,
	DeliveryReasonAdGroupArchived,
	DeliveryReasonAdGroupEnded,
	DeliveryReasonAdGroupIncomplete,
	DeliveryReasonAdGroupPaused,
	DeliveryReasonAdGroupPendingStartDate,
	DeliveryReasonAdArchived,
	DeliveryReasonAdIncomplete,
	DeliveryReasonAdPaused,
	DeliveryReasonAdPolicing,
	DeliveryReasonAdRejected,
	DeliveryReasonAdUnderReview,
	DeliveryReasonAccountOutOfBudget,
	DeliveryReasonBiddingControlsIneffective,
	DeliveryReasonBudgetControlsIneffective,
	DeliveryReasonKeywordArchived,
	DeliveryReasonKeywordPaused,
	DeliveryReasonNoValidPaymentMethod,
	DeliveryReasonPortfolioEnded,
	DeliveryReasonPortfolioPaused,
	DeliveryReasonTargetingClauseArchived,
	DeliveryReasonTargetingClausePaused,
	DeliveryReasonTargetingClausePolicing,
	DeliveryReasonTargetingClauseRejected,
	DeliveryReasonTargetingClauseUnderReview,
}

func (dr DeliveryReason) String() string {
	return string(dr)
}

func (dr DeliveryReason) IsValid() bool {
	for _, valid := range AllDeliveryReasons {
		if dr == valid {
			return true
		}
	}
	return false
}
