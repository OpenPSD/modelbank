package usecases

import (
	"time"

	"github.com/openpsd/modelbank/entities"
)

type Response struct {
	container map[string][]byte
}

func NewResponse() Response {
	m := make(map[string][]byte)
	response := Response{
		container: m,
	}
	return response
}

type Request struct {
	container map[string][]byte
}

func NewRequest() Request {
	m := make(map[string][]byte)
	request := Request{
		container: m,
	}
	return request
}

type ModelBankEntityGateway interface {
	ReadTestdata()
	CreateRepository()
	FindFiByBic(bic string) (entities.FinancialInstitution, error)
	FindAccntHldrByNameAndBirthDate(name string,
		birthDate time.Time,
		fi *entities.FinancialInstitution) ([]entities.AccountHolder, error)

	FindAccntByIban(iban string) (entities.Account, error)
	FindAccntByAccntHldr(accntHldr *entities.AccountHolder) ([]entities.Account, error)
	FindCnsntById(id string) (entities.Consent, error)
	StoreFi(fi *entities.FinancialInstitution) error
	StoreAccntHldr(accntHldr *entities.AccountHolder) error
	StoreAccnt(accnt *entities.Account) error
	StoreCnsnt(cnsnt *entities.Consent) error
}

type ModelBankUsecaseOutputPort interface {
	ReceiveResponse(resp Response)
}

type ModelBankUsecaseInputPort interface {
	InitializeBank(request Request) error
	CreateConsent(request Request) (Response, error)
	GetConsent(id string) (Response, error)
	CreateFi(bic string) (Response, error)
	CreateAccntHldr(request Request) (Response, error)
	CreateAccnt(request Request) (Response, error)
}
type ModelBankUseCase struct {
	outputPort    ModelBankUsecaseOutputPort
	entityManager ModelBankEntityGateway
}

func NewModelBankUseCase(mngr ModelBankEntityGateway, output ModelBankUsecaseOutputPort) ModelBankUseCase {
	return ModelBankUseCase{
		outputPort:    output,
		entityManager: mngr,
	}
}

func (usecase *ModelBankUseCase) InitializeBank(request Request) error {
	var err error
	usecase.entityManager.ReadTestdata()
	usecase.entityManager.CreateRepository()
	return err
}

func (usecase *ModelBankUseCase) CreateFi(bic string) (Response, error) {
	fi, err := usecase.entityManager.FindFiByBic(bic)
	if err != nil {
		fi = entities.NewFinancialInstitution()
		fi.Bic = bic
		err = usecase.entityManager.StoreFi(&fi)
	}
	resp := NewResponse()
	resp.container["fi"], err = fi.Marshal()
	return resp, err
}

func (usecase *ModelBankUseCase) CreateAccntHldr(request Request) error {
	var err error

	/*
		name string,
		birthDate time.Time,
		birthCity string,
		birthCountry string,
		agent entities.FinancialInstitution) (entities.AccountHolder, error
	*/
	//TODO CODE ME

	return err

}

func (usecase *ModelBankUseCase) CreateAccnt(request Request) error {
	/*
		iban string,
		currency string) (entities.Account, error)
	*/
	var err error
	//TODO CODE ME

	return err

}

func (usecase *ModelBankUseCase) CreateConsent(request Request) (Response, error) {
	var err error
	var response Response
	consent := entities.NewConsent()
	err = consent.Unmarshal(request.container["consent"])
	if err != nil {
		return response, err
	}
	consent, err = usecase.entityManager.FindCnsntById(consent.ID)
	if err != nil {
		err = usecase.entityManager.StoreCnsnt(&consent)
	}
	resp := NewResponse()
	resp.container["consent"], err = consent.Marshal()
	return resp, err

}

/*
func mapConsent(request Request) (entities.Consent, error) {
	consent := entities.NewConsent()
	var err error
	layout := "2006-01-02T15:04:05.000Z"

	consent.ID = request.container["id"]
	consent.AccessGrantedToTpp = request.container["accessGrantedToTpp"]
	consent.ConsentStatus = request.container["consentStatus"]

	fpd := request.container["frequencyPerDay"]
	consent.FrequencyPerDay, err = strconv.Atoi(fpd)

	consent.LastActionDate, err = time.Parse(layout, request.container["lastActionDate"])

	consent.RecurringIndicator, err = strconv.ParseBool(request.container["recurringIndicator"])

	consent.ScaStatus = request.container["scaStatus"]
	consent.TppCertificate = request.container["tppCertificate"]

	consent.ValidUntil, err = time.Parse(layout, request.container["validUntil"])

	return consent, err

}
*/
