package product

import (
	present "app.eirc/internal/presenter/product"
	"app.eirc/internal/router/middleware"
	"app.eirc/internal/router/middleware/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("products")
	{
		v10.POST("", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Create)
		v10.POST("list", middleware.Verify(), auth.AuthCheckRole(db), control.GetByList)
		v10.POST("get-by-order/:orderID", middleware.Verify(), control.GetByOrderIDList)
		v10.GET(":productID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingle)
		v10.DELETE(":productID", middleware.Verify(), auth.AuthCheckRole(db), control.Delete)
		v10.PATCH(":productID", middleware.Verify(), auth.AuthCheckRole(db), control.Update)
	}

	return router
}
