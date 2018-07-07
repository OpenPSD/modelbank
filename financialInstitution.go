package modelbank

import "encoding/json"

type FinancialInstitution struct {
	Name         string `json:"name"`
	Country      string `json:"country"`
	AddressLine1 string `json:"address1"`
	AddressLine2 string `json:"address2"`
	Bic          string `json:"bic"`
	IdOther      string `json:"idother"`
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
