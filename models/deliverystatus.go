package models

// DeliveryStatus represents the delivery status of campaigns/ads/adgroups
type DeliveryStatus string

const (
	DeliveryStatusDelivering               DeliveryStatus = "DELIVERING"
	DeliveryStatusNotDelivering            DeliveryStatus = "NOT_DELIVERING"
	DeliveryStatusPendingStartDate         DeliveryStatus = "PENDING_START_DATE"
	DeliveryStatusPaused                   DeliveryStatus = "PAUSED"
	DeliveryStatusEnded                    DeliveryStatus = "ENDED"
	DeliveryStatusArchived                 DeliveryStatus = "ARCHIVED"
	DeliveryStatusCampaignOutOfBudget      DeliveryStatus = "CAMPAIGN_OUT_OF_BUDGET"
	DeliveryStatusAccountOutOfBudget       DeliveryStatus = "ACCOUNT_OUT_OF_BUDGET"
	DeliveryStatusAdGroupPaused            DeliveryStatus = "AD_GROUP_PAUSED"
	DeliveryStatusCampaignPaused           DeliveryStatus = "CAMPAIGN_PAUSED"
	DeliveryStatusAdGroupIncomplete        DeliveryStatus = "AD_GROUP_INCOMPLETE"
	DeliveryStatusCampaignIncomplete       DeliveryStatus = "CAMPAIGN_INCOMPLETE"
	DeliveryStatusAdvertiserPaymentFailure DeliveryStatus = "ADVERTISER_PAYMENT_FAILURE"
	DeliveryStatusAdvertiserOutOfBudget    DeliveryStatus = "ADVERTISER_OUT_OF_BUDGET"
)

var AllDeliveryStatuses = []DeliveryStatus{
	DeliveryStatusDelivering,
	DeliveryStatusNotDelivering,
	DeliveryStatusPendingStartDate,
	DeliveryStatusPaused,
	DeliveryStatusEnded,
	DeliveryStatusArchived,
	DeliveryStatusCampaignOutOfBudget,
	DeliveryStatusAccountOutOfBudget,
	DeliveryStatusAdGroupPaused,
	DeliveryStatusCampaignPaused,
	DeliveryStatusAdGroupIncomplete,
	DeliveryStatusCampaignIncomplete,
	DeliveryStatusAdvertiserPaymentFailure,
	DeliveryStatusAdvertiserOutOfBudget,
}

func (ds DeliveryStatus) String() string {
	return string(ds)
}

func (ds DeliveryStatus) IsValid() bool {
	for _, valid := range AllDeliveryStatuses {
		if ds == valid {
			return true
		}
	}
	return false
}
