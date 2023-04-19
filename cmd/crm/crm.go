package main

import (
	"app.eirc/internal/interactor/pkg/connect"
	"app.eirc/internal/interactor/pkg/util/log"
	"app.eirc/internal/router"
	"app.eirc/internal/router/account"
	"app.eirc/internal/router/contact"
	"app.eirc/internal/router/industry"
	"app.eirc/internal/router/lead"
	"app.eirc/internal/router/lead_contact"
	"app.eirc/internal/router/login"
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
	engine = lead_contact.GetRouter(engine, db)
	engine = account.GetRouter(engine, db)
	engine = contact.GetRouter(engine, db)
	engine = industry.GetRouter(engine, db)
	log.Fatal(gateway.ListenAndServe(":8080", engine))
}
