package modelbank

import (
	"encoding/json"
	"time"
)

type FinancialInstitution struct {
	Name                        string                `json:"name"`
	Country                     string                `json:"country"`
	AddressLine1                string                `json:"address1"`
	AddressLine2                string                `json:"address2"`
	Bic                         string                `json:"bic"`
	IdOther                     string                `json:"idother"`
	OpenTransactions            []*PaymentTransaction `json:"opentransactions"`
	ReceivedPaymentInstructions []*PaymentInstruction `json:"received"`
	PaymentInstructionsToBeSent []*PaymentInstruction `json:"tbsent"`
	Customers                   []*AccountHolder      `json:"customers"`
}

type financialInstitutionProcessing interface {
	onboardPerson(name string,
		address string,
		dateOfBirth time.Time,
		cityOfBirth string,
		countryOfBirth string) *AccountHolder
	onboardOrganisation(name string,
		address string,
		tin string) *AccountHolder

	// receivePaymentInstruction
	// receives incoming paymentinstructions from other banks and stores them
	// for further processing.
	receivePaymentInstruction(paymentInstruction PaymentInstruction)

	// sendPaymentInstruction
	// sends outgoing paymentinstructions to other banks.
	sendPaymentInstruction(paymentInstruction PaymentInstruction)

	// processIncomingPaymentInstructions
	// books all received incoming paymentinstructions to corresponding accounts' ledgers.
	processIncomingPaymentInstructions()

	// createOutgoingPaymentInstructions
	// takes all existing paymenttransactions created by accountholders and creates
	// outgoing paymentInstructions. Books all payment transactions on accounts's ledgers.
	createOutgoingPaymentInstructions()

	GetAccounts() []Account
}

func (fi FinancialInstitution) GetAccounts() []*Account {
	var accounts []*Account

	acholders := fi.Customers

	for i := 0; i < len(acholders); i++ {
		cust := acholders[i]
		custAccnts := cust.Accounts
		for j := 0; j < len(custAccnts); j++ {
			acct := custAccnts[j]

			accounts = append(accounts, acct)
		}
	}

	return accounts
}

func NewFinancialInstitution() FinancialInstitution {
	return FinancialInstitution{}
}

// Marshal interface implementation
func (m *FinancialInstitution) Marshal() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Unmarshal interface implementation
func (m *FinancialInstitution) Unmarshal(b []byte) error {
	var res FinancialInstitution
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
