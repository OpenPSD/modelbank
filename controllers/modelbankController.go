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

func (cntrlr ModelBankController) ReceiveResponse(resp usecases.Response) {

}
