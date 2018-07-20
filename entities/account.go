package entities

import (
	"encoding/json"
	"time"
)

type ConsentAccountEntry struct {
	Iban      string `json:"iban"`
	Bban      string `json:"bban"`
	Pan       string `json:"pan"`
	MaskedPan string `json:"maskedPan"`
	Msisdn    string `json:"msisdn"`
	Currency  string `json:"currency"`
}
type ConsentAccess struct {
	Accounts          []ConsentAccountEntry `json:"accounts"`
	Balances          []ConsentAccountEntry `json:"balances"`
	Transactions      []ConsentAccountEntry `json:"transactions"`
	AvailableAccounts string                `json:"availableAccounts"`
	AllPsd2           string                `json:"allPsd2"`
}
type Consent struct {
	ID                 string        `json:"id"`
	AccessGrantedToTpp string        `json:"tpp"`
	TppCertificate     string        `json:"tppCertificate"`
	Access             ConsentAccess `json:"access"`
	RecurringIndicator bool          `json:"recurringIndicator"`
	ValidUntil         time.Time     `json:"validUntil"`
	FrequencyPerDay    int           `json:"frequencyPerDay"`
	LastActionDate     time.Time     `json:"lastActionDate"`
	ConsentStatus      string        `json:"consentStatus"`
	ScaStatus          string        `json:"scaStatus"`
}

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
	Accounts              []*Account            `json:"accounts"`
	UserName              string                `json:"username"`
}

type Account struct {
	AccountHolder *AccountHolder `json:"accountholder"`
	Iban          string         `json:"iban"`
	Bic           string         `json:"bic"`
	Currency      string         `json:"currency"`
	Balance       float64        `json:"balance"`
	Limit         float64        `json:"limit"`
	OpeningDate   time.Time      `json:"opening"`
	ClosingDate   time.Time      `json:"closing"`
	Ledger        []*LedgerEntry `json:"Ledger"`
	Status        string         `json:"status"`
}

func (ah *AccountHolder) instructPayment(crdtrName string,
	crdtrIban string,
	crdtBic string,
	crdtAmnt float64,
	crdtCrrncy string,
	purpose string,
	reqExecDate time.Time, dbtrAccnt *Account) *PaymentTransaction {

	var pt PaymentTransaction

	crdtAgnt := NewFinancialInstitution()
	crdtAgnt.Bic = crdtBic
	crdtr := NewAccountHolder(crdtAgnt)
	crdtr.Name = crdtrName
	crdtrAccnt := NewAccount(crdtr)
	crdtrAccnt.Bic = crdtBic
	crdtrAccnt.Iban = crdtrIban
	crdtrAccnt.Currency = crdtCrrncy
	pt = NewPaymentTransaction(crdtr, crdtrAccnt, crdtAgnt)
	pt.Amount = crdtAmnt
	pt.Purpose = purpose

	pi := NewPaymentInstruction(*dbtrAccnt.AccountHolder, *dbtrAccnt, *dbtrAccnt.AccountHolder.Agent)
	pi.RequestedExecutionDate = reqExecDate
	pi.Status = New

	pi.Transactions = append(pi.Transactions, &pt)
	pi.Status = New
	pt.PaymentInstruction = &pi

	openTx := dbtrAccnt.AccountHolder.Agent.OpenTransactions
	openTx = append(openTx, &pt)

	return &pt
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

func NewConsent() Consent {
	return Consent{}
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

// Marshal interface implementation
func (m *Consent) Marshal() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Unmarshal interface implementation
func (m *Consent) Unmarshal(b []byte) error {
	var res Consent
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
