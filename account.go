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
	Accounts              []*Account            `json:"accounts"`
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

type accountProcessing interface {
	setLimit(value float64)
	debitAccount(amount float64)
	creditAccount(amount float64)
	isAccountCoverageSufficient() bool
}

type accountHolderProcessing interface {
	openAccount(accountHolder AccountHolder, currency string) *Account

	closeAccount()

	// intructPayment
	// This is a credit transfer initiated by accountholder. It
	// creates a paymenttransaction and stores it in financial institution's
	// context for further processing.
	intructPayment(crdtrName string,
		crdtrIban string,
		crdtBic string,
		crdtAmnt float64,
		crdtCrrncy string,
		purpose string,
		reqExecDate time.Time, dbtrAccnt *Account) *PaymentTransaction
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
