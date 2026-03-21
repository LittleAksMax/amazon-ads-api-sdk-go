package models

// CreativeType represents the creative type for ad groups
type CreativeType string

const (
	CreativeTypeImage CreativeType = "IMAGE"
	CreativeTypeVideo CreativeType = "VIDEO"
	CreativeTypeAudio CreativeType = "AUDIO"
	CreativeTypeText  CreativeType = "TEXT"
)

var AllCreativeTypes = []CreativeType{
	CreativeTypeImage,
	CreativeTypeVideo,
	CreativeTypeAudio,
	CreativeTypeText,
}

func (ct CreativeType) String() string {
	return string(ct)
}

func (ct CreativeType) IsValid() bool {
	for _, valid := range AllCreativeTypes {
		if ct == valid {
			return true
		}
	}
	return false
}
