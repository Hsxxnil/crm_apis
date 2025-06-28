package main

import (
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

	"github.com/apex/gateway"
)

func main() {
	db, err := connect.PostgresSQL()
	if err != nil {
		log.Error(err)
		return
	}
	engine := router.Default()
	engine = user.GetRouter(engine, db)
	engine = login.GetRouter(engine, db)
	engine = lead.GetRouter(engine, db)
	engine = account.GetRouter(engine, db)
	engine = contact.GetRouter(engine, db)
	engine = industry.GetRouter(engine, db)
	engine = product.GetRouter(engine, db)
	engine = order.GetRouter(engine, db)
	engine = contract.GetRouter(engine, db)
	engine = order_product.GetRouter(engine, db)
	engine = campaign.GetRouter(engine, db)
	engine = quote.GetRouter(engine, db)
	engine = opportunity.GetRouter(engine, db)
	engine = opportunity_campaign.GetRouter(engine, db)
	engine = quote_product.GetRouter(engine, db)
	engine = policy.GetRouter(engine, db)
	engine = role.GetRouter(engine, db)
	engine = historical_record.GetRouter(engine, db)
	engine = event.GetRouter(engine, db)
	log.Fatal(gateway.ListenAndServe(":8080", engine))
}
