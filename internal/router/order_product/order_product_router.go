package order_product

import (
	present "app.eirc/internal/presenter/order_product"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("orders").Group("products")
	{
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.GET("", control.GetByList)
		v10.GET(":orderProductID", control.GetBySingle)
		v10.DELETE(":orderProductID", control.Delete)
		v10.PATCH(":orderProductID", control.Update)
	}

	return router
}
