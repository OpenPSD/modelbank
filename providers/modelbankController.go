package providers

import (
	"github.com/openpsd/modelbank/entities"
	"github.com/openpsd/modelbank/usecases"
)

type ModelBankControllerOutputPort interface {
}

type ModelBankControllerInputPort interface {
	InitializeBank()
	CreateConsent(Consent)
}

type ModelBankController struct {
	entityManager usecases.ModelBankEntityGateway
	usecase       usecases.ModelBankUsecaseInputPort
}

func NewModelBankontroller(mgr usecases.ModelBankEntityGateway) ModelBankController {
	cntrlr := ModelBankController{
		entityManager: mgr,
	}
	cntrlr.usecase = usecases.NewModelBankUseCase(cntrlr.entityManager, cntrlr)
	return cntrlr
}

func (cntrlr ModelBankController) CreateConsent(cnsnt *Consent) (*Consent, error) {
	var req usecases.Request
	var resp usecases.Response
	var err error
	req = mapToReq(cnsnt)
	resp, err = cntrlr.usecase.CreateConsent(req)
	cnsnt = mapToConsent(&resp)
	return cnsnt, err
}

func (cntrlr ModelBankController) InitializeBank(req usecases.Request) {
	cntrlr.usecase.InitializeBank(req)
}

func mapToReq(cnst *Consent) usecases.Request {
	consent := *cnst
	cnsnt := entities.Consent{
		ID: consent.Id,
	}
	req := usecases.NewRequest()
	req.Container["consent"], _ = cnsnt.Marshal()
	return req
}

func mapToConsent(resp *usecases.Response) *Consent {
	cnsnt := entities.NewConsent()
	cnsnt.Unmarshal(resp.Container["consent"])
	consent := Consent{
		Id: cnsnt.ID,
	}
	return &consent
}

func (cntrlr ModelBankController) ReceiveResponse(resp usecases.Response) {

}
