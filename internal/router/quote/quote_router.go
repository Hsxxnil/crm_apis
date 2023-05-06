package quote

import (
	present "app.eirc/internal/presenter/quote"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("quotes")
	{
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.GET("", control.GetByList)
		v10.GET("products", control.GetByListProducts)
		v10.GET(":quoteID", control.GetBySingle)
		v10.GET("products/:quoteID", control.GetBySingleProducts)
		v10.DELETE(":quoteID", control.Delete)
		v10.PATCH(":quoteID", control.Update)
	}

	return router
}
