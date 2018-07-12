package entities

type accountProcessor interface {
	setLimit(value float64)
	debitAccount(amount float64)
	creditAccount(amount float64)
	isAccountCoverageSufficient() bool
}
