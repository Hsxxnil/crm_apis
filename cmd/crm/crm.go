package main

import (
	"app.eirc/internal/interactor/pkg/connect"
	"app.eirc/internal/interactor/pkg/util/log"
	"app.eirc/internal/router"
	"app.eirc/internal/router/account"
	"app.eirc/internal/router/account_type"
	"app.eirc/internal/router/contact"
	"app.eirc/internal/router/contract"
	"app.eirc/internal/router/industry"
	"app.eirc/internal/router/lead"
	"app.eirc/internal/router/login"
	"app.eirc/internal/router/order"
	"app.eirc/internal/router/order_product"
	"app.eirc/internal/router/product"
	"app.eirc/internal/router/user"
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
	engine = account_type.GetRouter(engine, db)
	engine = contract.GetRouter(engine, db)
	engine = order_product.GetRouter(engine, db)
	log.Fatal(gateway.ListenAndServe(":8080", engine))
}
