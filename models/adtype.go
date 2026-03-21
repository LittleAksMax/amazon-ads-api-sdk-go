package models

// AdType represents the type of ad
type AdType string

const (
	AdTypeAudio AdType = "AUDIO"
	AdTypeVideo AdType = "VIDEO"
	AdTypeImage AdType = "IMAGE"
	AdTypeText  AdType = "TEXT"
)

var AllAdTypes = []AdType{
	AdTypeAudio,
	AdTypeVideo,
	AdTypeImage,
	AdTypeText,
}

func (at AdType) String() string {
	return string(at)
}

func (at AdType) IsValid() bool {
	for _, valid := range AllAdTypes {
		if at == valid {
			return true
		}
	}
	return false
}
