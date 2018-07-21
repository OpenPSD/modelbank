package controllers

import (
	"github.com/openpsd/modelbank/usecases"
)

type ModelBankControllerOutputPort interface {
}

type ModelBankControllerInputPort interface {
	InitializeBank(request usecases.Request)
}

type ModelBankController struct {
	EntityManager usecases.ModelBankEntityGateway
	Usecase       usecases.ModelBankUsecaseInputPort
}

func NewModelBankontroller(mgr usecases.ModelBankEntityGateway) ModelBankController {
	cntrlr := ModelBankController{
		EntityManager: mgr,
	}
	cntrlr.Usecase = usecases.NewModelBankUseCase(cntrlr.EntityManager, cntrlr)
	return cntrlr
}

func (cntrlr ModelBankController) ReceiveResponse(resp usecases.Response) {

}
