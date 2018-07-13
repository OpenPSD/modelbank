package controller

import "github.com/openpsd/modelbank/entities"

type FinancialInstitutionOutputPort interface {
	FindByBic(bic string) (entities.FinancialInstitution, error)
}

type FinancialInstitutionInputPort interface {
	Create() entities.FinancialInstitution
}

type FinancialInstituteController struct {
	outputPort FinancialInstitutionOutputPort
}

func NewFinancialInstituteController(outPutPort FinancialInstitutionOutputPort) FinancialInstituteController {
	return FinancialInstituteController{
		outputPort: outPutPort,
	}
}

func (fiController *FinancialInstituteController) Create(bic string) entities.FinancialInstitution {
	fi, error := fiController.outputPort.FindByBic(bic)
	if error != nil {
		fi = entities.NewFinancialInstitution()
		fi.Bic = bic
	}
	return fi
}
