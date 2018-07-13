package controllers

import "github.com/openpsd/modelbank/entities"

type FinancialInstitutionOutputPort interface {
	FindByBic(bic string) (entities.FinancialInstitution, error)
	Store(fi *entities.FinancialInstitution) error
}

type FinancialInstitutionInputPort interface {
	Create(bic string) (entities.FinancialInstitution, error)
}

type FinancialInstituteController struct {
	outputPort FinancialInstitutionOutputPort
}

func NewFinancialInstituteController(outPutPort FinancialInstitutionOutputPort) FinancialInstituteController {
	return FinancialInstituteController{
		outputPort: outPutPort,
	}
}

func (fiController FinancialInstituteController) Create(bic string) (entities.FinancialInstitution, error) {
	fi, error := fiController.outputPort.FindByBic(bic)
	if error != nil {
		fi = entities.NewFinancialInstitution()
		fi.Bic = bic
		error = fiController.outputPort.Store(&fi)
	}
	return fi, error
}
