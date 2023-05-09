package product

import (
	present "app.eirc/internal/presenter/product"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("products")
	{
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.POST("list", control.GetByList)
		v10.GET("list", control.GetByList)
		v10.GET(":productID", control.GetBySingle)
		v10.DELETE(":productID", control.Delete)
		v10.PATCH(":productID", control.Update)
	}

	return router
}
