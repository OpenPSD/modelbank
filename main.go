package main

import (
	"log"

	"github.com/openpsd/modelbank/providers"
)

func main() {
	log.Println("Create modelbank's database")

	providers.ModelbankDB_Model()

	log.Println("modelbank's database created")

}
