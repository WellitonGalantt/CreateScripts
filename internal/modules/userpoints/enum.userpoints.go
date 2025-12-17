package userpoints

type TransactionType string

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)

func (t TransactionType) IsValid() bool {
	switch t {
	case Credit, Debit:
		return true
	}

	return false
}

type TransactionReason string

const (
	Generation   TransactionReason = "generation"
	Batch        TransactionReason = "batch"
	Manual       TransactionReason = "manual"
	MonthlyReset TransactionReason = "monthly_reset"
)

func (t TransactionReason) IsValid() bool {
	switch t {
	case Generation, Batch, Manual, MonthlyReset:
		return true
	}

	return false
}
