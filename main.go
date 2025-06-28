package main

import (
	"fmt"
	"net/http"

	_ "crm/api"
	"crm/internal/interactor/pkg/connect"
	"crm/internal/interactor/pkg/util/log"
	"crm/internal/router"
	"crm/internal/router/account"
	"crm/internal/router/campaign"
	"crm/internal/router/contact"
	"crm/internal/router/contract"
	"crm/internal/router/event"
	"crm/internal/router/historical_record"
	"crm/internal/router/industry"
	"crm/internal/router/lead"
	"crm/internal/router/login"
	"crm/internal/router/opportunity"
	"crm/internal/router/opportunity_campaign"
	"crm/internal/router/order"
	"crm/internal/router/order_product"
	"crm/internal/router/policy"
	"crm/internal/router/product"
	"crm/internal/router/quote"
	"crm/internal/router/quote_product"
	"crm/internal/router/role"
	"crm/internal/router/user"

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

// @host		localhost:8080
// @BasePath	/crm/v1.0
// @schemes	http
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
	opportunity_campaign.GetRouter(engine, db)
	quote_product.GetRouter(engine, db)
	policy.GetRouter(engine, db)
	role.GetRouter(engine, db)
	historical_record.GetRouter(engine, db)
	event.GetRouter(engine, db)

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:8080/swagger/doc.json"))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	log.Fatal(http.ListenAndServe(":8080", engine))
}
