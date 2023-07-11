package order

import (
	present "app.eirc/internal/presenter/order"
	"app.eirc/internal/router/middleware"
	"app.eirc/internal/router/middleware/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("orders")
	{
		v10.POST("", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Create)
		v10.POST("list", middleware.Verify(), auth.AuthCheckRole(db), control.GetByList)
		v10.GET(":orderID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingle)
		v10.GET("products/:orderID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingleProducts)
		v10.DELETE(":orderID", middleware.Verify(), auth.AuthCheckRole(db), control.Delete)
		v10.PATCH(":orderID", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Update)
	}

	return router
}
