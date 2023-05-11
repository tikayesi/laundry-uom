// package delivery

// import (
// 	"fmt"

// 	"github.com/eulbyvan/go-enigma-laundry/config"
// 	"github.com/eulbyvan/go-enigma-laundry/repository"
// 	"github.com/eulbyvan/go-enigma-laundry/usecase"
// 	"github.com/jmoiron/sqlx"
// )

// func Run() {
// 	config := config.NewConfig()
// 	db := config.DbConnect()

// 	ProductCli(db)
// }

// func ProductCli(db *sqlx.DB) {
// 	uomRepo := repository.NewUomRepository(db)
// 	UomUc := usecase.NewUomUseCase(uomRepo)

// 	products, err := UomUc.GetAllUom()
// 	if err != nil {
// 		panic(err.Error())
// 	}

//		for _, product := range products {
//			fmt.Println(product)
//		}
//	}
package delivery
