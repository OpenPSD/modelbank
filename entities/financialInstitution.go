package entities

import (
	"encoding/json"
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

func (fi *FinancialInstitution) receivePaymentInstruction(paymentInstruction PaymentInstruction) {
	fi.ReceivedPaymentInstructions = append(fi.ReceivedPaymentInstructions, &paymentInstruction)
}
func (fi *FinancialInstitution) processIncomingPaymentInstructions() {
	openInstrustions := fi.ReceivedPaymentInstructions
	for i := 0; i < len(openInstrustions); i++ {
		pi := openInstrustions[i]
		txs := pi.Transactions
		for j := 0; j < len(txs); j++ {
			tx := txs[0]
			iban := tx.CreditorAccount.Iban
			ac := fi.GetAccount(iban)
			le := NewLedgerEntry(ac)
			le.EntryAmount = tx.Amount
			le.EntryDate = pi.RequestedExecutionDate
			le.EntryTitle = tx.Purpose
			ac.Balance = ac.Balance - le.EntryAmount
			ac.Ledger = append(ac.Ledger, &le)
			tx.Status = Processed

		}
	}
}

func (fi *FinancialInstitution) GetAccount(iban string) Account {
	var ac Account

	return ac

}

func (fi *FinancialInstitution) GetAccounts() []Account {
	var accounts []Account

	accntholders := fi.Customers

	for i := 0; i < len(accntholders); i++ {
		cust := accntholders[i]
		custAccnts := cust.Accounts
		for j := 0; j < len(custAccnts); j++ {
			acct := *custAccnts[j]
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
