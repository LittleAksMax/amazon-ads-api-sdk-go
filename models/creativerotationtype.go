package models

// CreativeRotationType represents how creatives rotate in ad groups
type CreativeRotationType string

const (
	CreativeRotationTypeRandom     CreativeRotationType = "RANDOM"
	CreativeRotationTypeOptimized  CreativeRotationType = "OPTIMIZED"
	CreativeRotationTypeSequential CreativeRotationType = "SEQUENTIAL"
)

var AllCreativeRotationTypes = []CreativeRotationType{
	CreativeRotationTypeRandom,
	CreativeRotationTypeOptimized,
	CreativeRotationTypeSequential,
}

func (crt CreativeRotationType) String() string {
	return string(crt)
}

func (crt CreativeRotationType) IsValid() bool {
	for _, valid := range AllCreativeRotationTypes {
		if crt == valid {
			return true
		}
	}
	return false
}
