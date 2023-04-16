package main

import (
	"fmt"
	"net/http"

	"app.eirc/internal/interactor/pkg/connect"

	_ "app.eirc/api"
	"app.eirc/internal/interactor/pkg/util/log"
	"app.eirc/internal/router"
	"app.eirc/internal/router/account"
	"app.eirc/internal/router/lead"
	"app.eirc/internal/router/lead_contact"
	"app.eirc/internal/router/login"
	"app.eirc/internal/router/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// main is run all api form localhost port 8080

//	@title			CRM API
//	@version		0.1
//	@description	CRM API
//	@termsOfService	https://eirc.app/

//	@contact.name	API System Support
//	@contact.url	https://eirc.app/
//	@contact.email	eirc8888@gmail.com

//	@license.name	AGPL 3.0
//	@license.url	https://www.gnu.org/licenses/agpl-3.0.en.html

// @host		api.testing.eirc
// @BasePath	/
// @schemes	https
func main() {
	db, err := connect.PostgresSQL()
	if err != nil {
		log.Error(err)
		return
	}

	engine := router.Default()
	user.GetRouter(engine, db)
	login.GetRouter(engine, db)
	lead.GetRouter(engine, db)
	lead_contact.GetRouter(engine, db)
	account.GetRouter(engine, db)

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:8080/swagger/doc.json"))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	log.Fatal(http.ListenAndServe(":8080", engine))
}
