package main

import (
	"fmt"
	"net/http"

	"app.eirc/internal/interactor/pkg/connect"

	"app.eirc/internal/router/customer"

	_ "app.eirc/api"
	"app.eirc/internal/interactor/pkg/util/log"
	"app.eirc/internal/router"
	"app.eirc/internal/router/permission"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// main is run all api form localhost port 8080

//	@title			Whisky SYSTEM API
//	@version		0.1
//	@description	JG Server API
//	@termsOfService	https://inherited.app/

//	@contact.name	API System Support
//	@contact.url	https://inherited.app/
//	@contact.email	mingzong.lyu@gmail.com

//	@license.name	AGPL 3.0
//	@license.url	https://www.gnu.org/licenses/agpl-3.0.en.html

// @host		api.testing.whisky.inherited.app
// @BasePath	/
// @schemes	https
func main() {
	db, err := connect.PostgresSQL()
	if err != nil {
		log.Error(err)
		return
	}

	engine := router.Default()
	permission.GetRouter(engine, db)
	customer.GetRouter(engine, db)

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:8080/swagger/doc.json"))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	log.Fatal(http.ListenAndServe(":8080", engine))
}
