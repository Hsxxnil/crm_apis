package quote_product

import (
	present "app.eirc/internal/presenter/quote_product"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("quotes-products")
	{
		v10.POST("", middleware.Verify(), middleware.Transaction(db), control.Create)
		v10.GET("", middleware.Verify(), control.GetByList)
		v10.GET(":quoteProductID", middleware.Verify(), control.GetBySingle)
		v10.DELETE(":quoteProductID", middleware.Verify(), control.Delete)
		v10.PATCH("", middleware.Verify(), control.Update)
	}

	return router
}
