package main

import (
	"log"

	"github.com/openpsd/modelbank/providers"
)

func main() {
	log.Println("Create modelbank's database")

	providers.ModelbankDB_Model()

	log.Println("modelbank's database created")

	repository := providers.NewDbProvider()
	fi, err := repository.Controller.Create("PBNKDEFF")
	if err != nil {
		panic(err)
	}

	log.Println("fi added: " + fi.Bic)

}
