package main

import (
	"log"

	"github.com/openpsd/modelbank/entities"
	"github.com/openpsd/modelbank/providers"
)

func main() {
	log.Println("Create modelbank's database")

	providers.ModelbankDB_Model()

	log.Println("modelbank's database created")

	repository := providers.NewDbProvider()
	response, err := repository.ModelbankController.Usecase.CreateFi("PBNKDEFF")
	if err != nil {
		panic(err)
	}

	var fi entities.FinancialInstitution
	fi.Unmarshal(response.Container["fi"])
	log.Println("fi added: " + fi.Bic)

}
