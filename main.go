package modelbank

import (
	"log"

	"github.com/openpsd/modelbank/providers"
)

func main() {
	log.Println("Create modelbank's database")

	providers.ModelbankDB_Model()

}
