package models

// BudgetType represents the type of budget
type BudgetType string

const (
	BudgetTypeMonetary BudgetType = "MONETARY"
)

var AllBudgetTypes = []BudgetType{
	BudgetTypeMonetary,
}

func (bt BudgetType) String() string {
	return string(bt)
}

func (bt BudgetType) IsValid() bool {
	for _, valid := range AllBudgetTypes {
		if bt == valid {
			return true
		}
	}
	return false
}
