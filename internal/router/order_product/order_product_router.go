package order_product

import (
	present "app.eirc/internal/presenter/order_product"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("orders-products")
	{
		v10.POST("", middleware.Verify(), middleware.Transaction(db), control.Create)
		v10.GET("", middleware.Verify(), control.GetByList)
		v10.GET(":orderProductID", middleware.Verify(), control.GetBySingle)
		v10.DELETE(":orderProductID", middleware.Verify(), control.Delete)
		v10.PATCH(":orderProductID", middleware.Verify(), control.Update)
	}

	return router
}
