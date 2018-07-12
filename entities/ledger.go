package entities

import (
	"encoding/json"
	"time"
)

type LedgerEntry struct {
	RelatedAccount *Account  `json:"related"`
	EntryType      string    `json:"type"` // credit,debit
	EntryDate      time.Time `json:"date"`
	EntryAmount    float64   `json:"amount"`
	EntryTitle     string    `json:"title"`
	EntryStatus    string    `json:"status"`
}

func NewLedgerEntry(account Account) LedgerEntry {
	a := &account

	return LedgerEntry{
		RelatedAccount: a,
	}
}

// Marshal interface implementation
func (m *LedgerEntry) Marshal() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Unmarshal interface implementation
func (m *LedgerEntry) Unmarshal(b []byte) error {
	var res LedgerEntry
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
