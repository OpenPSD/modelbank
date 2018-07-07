package modelbank

import (
	"encoding/json"
	"time"
)

type FinancialInstitution struct {
	Name         string `json:"name"`
	Country      string `json:"country"`
	AddressLine1 string `json:"address1"`
	AddressLine2 string `json:"address2"`
	Bic          string `json:"bic"`
	IdOther      string `json:"idother"`
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
	receivePaymentInstruction(paymentInstruction PaymentInstruction)
	sendPaymentInstruction(paymentInstruction PaymentInstruction)
	processIncomingPaymentInstructions()
	createOutgoingPaymentInstructions()
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
