package usecases

import (
	"time"

	"github.com/openpsd/modelbank/entities"
)

type ModelBankUsecaseOutputPort interface {
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

type ModelBankUsecaseInputPort interface {
	InitializeBank() error
	CreateConsent()
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
type ModelBankUseCase struct {
	outputPort ModelBankUsecaseOutputPort
}

func NewModelBankUseCase(ouput ModelBankUsecaseOutputPort) ModelBankUseCase {
	return ModelBankUseCase{
		outputPort: ouput,
	}
}

func (usecase ModelBankUseCase) InitializeBank() error {
	var err error
	usecase.outputPort.ReadTestdata()
	usecase.outputPort.CreateRepository()
	return err
}

func (usecase ModelBankUseCase) CreateFi(bic string) (entities.FinancialInstitution, error) {
	fi, error := usecase.outputPort.FindFiByBic(bic)
	if error != nil {
		fi = entities.NewFinancialInstitution()
		fi.Bic = bic
		error = usecase.outputPort.StoreFi(&fi)
	}
	return fi, error
}

func (usecase ModelBankUseCase) CreateAccntHldr(name string,
	birthDate time.Time,
	birthCity string,
	birthCountry string,
	agent entities.FinancialInstitution) (entities.AccountHolder, error) {
	var err error
	var accntHldr entities.AccountHolder
	//TODO CODE ME

	return accntHldr, err

}

func (usecase ModelBankUseCase) CreateAccnt(iban string,
	currency string,
	accntHldr entities.AccountHolder) (entities.Account, error) {
	var err error
	var accnt entities.Account
	//TODO CODE ME

	return accnt, err

}

func (usecase ModelBankUseCase) CreateConsent() {

}
