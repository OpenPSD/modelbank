package entities

import (
	"encoding/json"
	"time"
)

const (
	New       = "new"
	Processed = "Processed"
)

type PaymentInstruction struct {
	PaymentInstructionId   string                `json:"id"`
	PaymentMethod          string                `json:"method"`
	BatchBooking           string                `json:"batchbooking"`
	NumberOfTansactions    int                   `json:"numberoftx"`
	ControlSum             float64               `json:"controlsum"`
	RequestedExecutionDate time.Time             `json:"reqexecdate"`
	Debtor                 *AccountHolder        `json:"debtor"`
	DebtorAccount          *Account              `json:"debtoraccount"`
	DebtorAgent            *FinancialInstitution `json:"debtoragent"`
	UltimateDebtor         *AccountHolder        `json:"ultimatedebtor"`
	ChargeBearer           string                `json:"chargebearer"`
	Transactions           []*PaymentTransaction `json:"transactions"`
	Status                 string                `json:"status"` //new, processed
}

type PaymentTransaction struct {
	PaymentInstruction    *PaymentInstruction   `json:"instruction"`
	PaymentId             string                `json:"id"`
	Amount                float64               `json:"amount"`
	ChargeBearer          string                `json:"chargebearer"`
	UltimateDebtor        *AccountHolder        `json:"ultimatedebtor"`
	Creditor              *AccountHolder        `json:"creditor"`
	CreditorAgent         *FinancialInstitution `json:"creditoragent"`
	CreditorAccount       *Account              `json:"creditoraccount"`
	UltimateCreditor      *AccountHolder        `json:"ultimatecreditor"`
	Purpose               string                `json:"purpose"`
	RemittanceInformation string                `json:"remittance"`
	Status                string                `json:"status"` // new, processed
}

func NewPaymentInstruction(debtor AccountHolder,
	debtorAccount Account,
	debtorAgent FinancialInstitution) PaymentInstruction {
	d := &debtor
	dac := &debtorAccount
	da := &debtorAgent

	return PaymentInstruction{
		Debtor:        d,
		DebtorAccount: dac,
		DebtorAgent:   da,
	}
}

func NewPaymentTransaction(creditor AccountHolder,
	creditorAccount Account,
	creditorAgent FinancialInstitution) PaymentTransaction {
	d := &creditor
	dac := &creditorAccount
	da := &creditorAgent

	return PaymentTransaction{
		Creditor:        d,
		CreditorAccount: dac,
		CreditorAgent:   da,
	}
}

// Marshal interface implementation
func (m *PaymentInstruction) Marshal() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Unmarshal interface implementation
func (m *PaymentInstruction) Unmarshal(b []byte) error {
	var res PaymentInstruction
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Marshal interface implementation
func (m *PaymentTransaction) Marshal() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Unmarshal interface implementation
func (m *PaymentTransaction) Unmarshal(b []byte) error {
	var res PaymentTransaction
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
