package main

import (
	"fmt"
	"net/http"

	"app.eirc/internal/interactor/pkg/connect"

	_ "app.eirc/api"
	"app.eirc/internal/interactor/pkg/util/log"
	"app.eirc/internal/router"
	"app.eirc/internal/router/account"
	"app.eirc/internal/router/campaign"
	"app.eirc/internal/router/contact"
	"app.eirc/internal/router/contract"
	"app.eirc/internal/router/industry"
	"app.eirc/internal/router/lead"
	"app.eirc/internal/router/login"
	"app.eirc/internal/router/opportunity"
	"app.eirc/internal/router/order"
	"app.eirc/internal/router/order_product"
	"app.eirc/internal/router/product"
	"app.eirc/internal/router/quote"
	"app.eirc/internal/router/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// main is run all api form localhost port 8080

//	@title			CRM APIs
//	@version		0.1
//	@description	CRM APIs
//	@termsOfService

//	@contact.name
//	@contact.url
//	@contact.email

//	@license.name	AGPL 3.0
//	@license.url	https://www.gnu.org/licenses/agpl-3.0.en.html

// @host		api.t.d2din.com
// @BasePath	/crm/v1.0
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
	account.GetRouter(engine, db)
	contact.GetRouter(engine, db)
	industry.GetRouter(engine, db)
	product.GetRouter(engine, db)
	order.GetRouter(engine, db)
	contract.GetRouter(engine, db)
	order_product.GetRouter(engine, db)
	campaign.GetRouter(engine, db)
	quote.GetRouter(engine, db)
	opportunity.GetRouter(engine, db)

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:8080/swagger/doc.json"))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	log.Fatal(http.ListenAndServe(":8080", engine))
}
