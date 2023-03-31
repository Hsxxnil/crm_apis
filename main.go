package main

import (
	"fmt"
	"net/http"

	_ "app.inherited.caelus/api"
	"app.inherited.caelus/internal/interactor/pkg/util/log"
	"app.inherited.caelus/internal/router"
	"app.inherited.caelus/internal/router/permission"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
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
	db := &gorm.DB{}
	engine := router.Default()
	permission.GetRouter(engine, db)

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:8081/swagger/doc.json"))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	log.Fatal(http.ListenAndServe(":8081", engine))
}
