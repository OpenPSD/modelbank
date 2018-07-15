package controllers

import (
	"time"

	"github.com/openpsd/modelbank/entities"
	"github.com/openpsd/modelbank/usecases"
)

type ModelBankControllerOutputPort interface {
	ReadTestdata()
	CreateRepository()
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

type ModelBankControllerInputPort interface {
	InitializeBank()
}

type ModelBankController struct {
	outputPort ModelBankControllerOutputPort
	usecase    usecases.ModelBankUsecaseInputPort
}

func NewModelBankontroller(outPutPort ModelBankControllerOutputPort) ModelBankController {
	cntrlr := ModelBankController{
		outputPort: outPutPort,
	}
	cntrlr.usecase = usecases.NewModelBankUseCase(cntrlr.outputPort)

	return cntrlr
}
