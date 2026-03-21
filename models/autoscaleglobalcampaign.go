package models

// AutoScaleGlobalCampaign represents auto scaling settings
type AutoScaleGlobalCampaign string

const (
	AutoScaleGlobalCampaignAuto   AutoScaleGlobalCampaign = "AUTO"
	AutoScaleGlobalCampaignManual AutoScaleGlobalCampaign = "MANUAL"
)

var AllAutoScaleGlobalCampaigns = []AutoScaleGlobalCampaign{
	AutoScaleGlobalCampaignAuto,
	AutoScaleGlobalCampaignManual,
}

func (asgc AutoScaleGlobalCampaign) String() string {
	return string(asgc)
}

func (asgc AutoScaleGlobalCampaign) IsValid() bool {
	for _, valid := range AllAutoScaleGlobalCampaigns {
		if asgc == valid {
			return true
		}
	}
	return false
}
