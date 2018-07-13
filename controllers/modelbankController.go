package controllers

import (
	"time"

	"github.com/openpsd/modelbank/entities"
)

type ModelBankOutputPort interface {
	FindFiByBic(bic string) (entities.FinancialInstitution, error)
	FindAccntHldrByNameAndBirthDate(name string,
		birthDate time.Time,
		fi *entities.FinancialInstitution) ([]entities.AccountHolder, error)

	FindAccntByIban(iban string) (entities.Account, error)
	FindAccntByAccntHldr(accntHldr *entities.AccountHolder) ([]entities.Account, error)
	StoreFi(fi *entities.FinancialInstitution) error
	StoreAccntHldr(accntHldr *entities.AccountHolder) error
	StoreAccnt(accnt *entities.Account) error
}

type ModelBankInputPort interface {
	CreateFi(bic string) (entities.FinancialInstitution, error)
	CreateAccntHldr(name string,
		birthDate time.Time,
		birthCity string,
		birthCountry string,
		agent entities.FinancialInstitution) (entities.AccountHolder, error)
	CreateAccnt(iban string,
		currency string,
		accntHldr entities.AccountHolder) (entities.Account, error)
}

type ModelBankController struct {
	outputPort ModelBankOutputPort
}

func NewModelBankontroller(outPutPort ModelBankOutputPort) ModelBankController {
	return ModelBankController{
		outputPort: outPutPort,
	}
}

func (fiController ModelBankController) Create(bic string) (entities.FinancialInstitution, error) {
	fi, error := fiController.outputPort.FindFiByBic(bic)
	if error != nil {
		fi = entities.NewFinancialInstitution()
		fi.Bic = bic
		error = fiController.outputPort.StoreFi(&fi)
	}
	return fi, error
}
