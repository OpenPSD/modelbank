package providers

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/openpsd/modelbank/controllers"
	"github.com/openpsd/modelbank/entities"
)

type DbProvider struct {
	Db         *pg.DB
	Controller controllers.FinancialInstitutionInputPort
}

func (dbp DbProvider) FindByBic(bic string) (entities.FinancialInstitution, error) {
	dbp.Db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "start123",
	})
	defer dbp.Db.Close()
	// Select fi by bic.
	fi := &entities.FinancialInstitution{}
	err := dbp.Db.Model(fi).Where("bic = ?", bic).Select()
	return *fi, err
}

func (dbp DbProvider) Store(fi *entities.FinancialInstitution) error {
	dbp.Db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "start123",
	})
	defer dbp.Db.Close()
	// Select fi by bic.
	err := dbp.Db.Insert(fi)
	return err
}

func NewDbProvider() DbProvider {
	dbp := DbProvider{}
	dbp.Controller = controllers.NewFinancialInstituteController(dbp)
	return dbp
}

func ModelbankDB_Model() {

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "start123",
	})
	defer db.Close()

	err := dropSchema(db)
	if err != nil {
		panic(err)
	}

	err = createSchema(db)
	if err != nil {
		panic(err)
	}

	fi := &entities.FinancialInstitution{
		Name:         "SampleBank",
		Country:      "DE",
		AddressLine1: "Street",
		AddressLine2: "Addresline 2",
	}
	err = db.Insert(fi)
	if err != nil {
		panic(err)
	}

	/*
		err = db.Insert(&User{
			Name:   "root",
			Emails: []string{"root1@root", "root2@root"},
		})
		if err != nil {
			panic(err)
		}

		story1 := &Story{
			Title:    "Cool story",
			AuthorId: user1.Id,
		}
		err = db.Insert(story1)
		if err != nil {
			panic(err)
		}

		// Select user by primary key.
		user := &User{Id: user1.Id}
		err = db.Select(user)
		if err != nil {
			panic(err)
		}

		// Select all users.
		var users []User
		err = db.Model(&users).Select()
		if err != nil {
			panic(err)
		}

		// Select story and associated author in one query.
		story := new(Story)
		err = db.Model(story).
			Relation("Author").
			Where("story.id = ?", story1.Id).
			Select()
		if err != nil {
			panic(err)
		}

		fmt.Println(user)
		fmt.Println(users)
		fmt.Println(story)
		// Output: User<1 admin [admin1@admin admin2@admin]>
		// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
		// Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
	*/
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*entities.FinancialInstitution)(nil), (*entities.AccountHolder)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func dropSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*entities.FinancialInstitution)(nil), (*entities.AccountHolder)(nil)} {
		err := db.DropTable(model, &orm.DropTableOptions{
			IfExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil

}
