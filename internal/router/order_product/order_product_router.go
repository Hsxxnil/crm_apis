package order_product

import (
	present "crm/internal/presenter/order_product"
	"crm/internal/router/middleware"
	"crm/internal/router/middleware/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("orders-products")
	{
		v10.POST("", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Create)
		v10.GET("", middleware.Verify(), auth.AuthCheckRole(db), control.GetByList)
		v10.GET(":orderProductID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingle)
		v10.DELETE("", middleware.Verify(), auth.AuthCheckRole(db), control.Delete)
		v10.PATCH("", middleware.Verify(), auth.AuthCheckRole(db), control.Update)
	}

	return router
}
