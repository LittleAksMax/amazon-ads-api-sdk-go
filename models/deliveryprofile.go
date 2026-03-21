package models

// DeliveryProfile represents the delivery pacing profile
type DeliveryProfile string

const (
	DeliveryProfileASAP DeliveryProfile = "ASAP"
	DeliveryProfileEven DeliveryProfile = "EVEN"
)

var AllDeliveryProfiles = []DeliveryProfile{
	DeliveryProfileASAP,
	DeliveryProfileEven,
}

func (dp DeliveryProfile) String() string {
	return string(dp)
}

func (dp DeliveryProfile) IsValid() bool {
	for _, valid := range AllDeliveryProfiles {
		if dp == valid {
			return true
		}
	}
	return false
}
