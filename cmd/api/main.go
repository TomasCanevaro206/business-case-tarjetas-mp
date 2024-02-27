package main

import (
	"database/sql"

	"github.com/TomasCanevaro206/business-case-tarjetas-mp.git/cmd/api/routes"
	"github.com/TomasCanevaro206/business-case-tarjetas-mp.git/pkg/log"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			MP Card Emission API
//	@version		1.0
//	@description	Emision de tarjetas en MP
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	// NO MODIFICAR
	db, err := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/emision_tarjetas?parseTime=true")
	if err != nil {
		panic(err)
	}

	eng := gin.Default()
	log.Logger = log.Init(db)

	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := eng.Run(); err != nil {
		panic(err)
	}
}
