package modelbank

import (
	"encoding/json"
	"time"
)

type AccountHolder struct {
	Agent                 *FinancialInstitution `json:"agent"`
	Name                  string                `json:"name"`
	Country               string                `json:"country"`
	AddressLine1          string                `json:"address1"`
	AddressLine2          string                `json:"address2"`
	AccountHolderType     string                `json:"type"`
	IdOrgBic              string                `json:"orgbic"`
	IdOther               string                `json:"idother"`
	IdOtherSchemeName     string                `json:"idschemename"`
	IdOtherIssuer         string                `json:"idissuer"`
	IdPersBirthDate       string                `json:"birthdate"`
	IdPersProvinceOfBirth string                `json:"birthprovince"`
	IdPersCityOfBirth     string                `json:"birthcity"`
	IdPersCountryOfBirth  string                `json:"birthcountry"`
}

type Account struct {
	AccountHolder *AccountHolder `json:"accountholder"`
	Iban          string         `json:"iban"`
	Bic           string         `json:"bic"`
	Currency      string         `json:"currency"`
	Balance       float64        `json:"balance"`
	OpeningDate   time.Time      `json:"opening"`
}

func NewAccount(accountHolder AccountHolder) Account {
	a := &accountHolder
	return Account{
		AccountHolder: a,
	}
}

func NewAccountHolder(financialInstitution FinancialInstitution) AccountHolder {
	f := &financialInstitution
	return AccountHolder{
		Agent: f,
	}
}

// Marshal interface implementation
func (m *AccountHolder) Marshal() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Unmarshal interface implementation
func (m *AccountHolder) Unmarshal(b []byte) error {
	var res AccountHolder
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Marshal interface implementation
func (m *Account) Marshal() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Unmarshal interface implementation
func (m *Account) Unmarshal(b []byte) error {
	var res Account
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
