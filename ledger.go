package modelbank

import (
	"encoding/json"
	"time"
)

type Ledger struct {
	RelatedAccount *Account  `json:"related"`
	EntryType      string    `json:"type"` // credit,debit
	EntryDate      time.Time `json:"date"`
	EntryAmount    float64   `json:"amount"`
	EntryTitle     string    `json:"title"`
}

func NewLedger(account Account) Ledger {
	a := &account

	return Ledger{
		RelatedAccount: a,
	}
}

// Marshal interface implementation
func (m *Ledger) Marshal() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Unmarshal interface implementation
func (m *Ledger) Unmarshal(b []byte) error {
	var res Ledger
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
