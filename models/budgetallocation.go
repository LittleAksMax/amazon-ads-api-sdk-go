package models

// BudgetAllocation represents the budget allocation strategy
type BudgetAllocation string

const (
	BudgetAllocationAuto   BudgetAllocation = "AUTO"
	BudgetAllocationManual BudgetAllocation = "MANUAL"
)

var AllBudgetAllocations = []BudgetAllocation{
	BudgetAllocationAuto,
	BudgetAllocationManual,
}

func (ba BudgetAllocation) String() string {
	return string(ba)
}

func (ba BudgetAllocation) IsValid() bool {
	for _, valid := range AllBudgetAllocations {
		if ba == valid {
			return true
		}
	}
	return false
}
