package customer

import (
	present "app.eirc/internal/presenter/customer"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("authority").Group("v1.0").Group("customer")
	{
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.GET("", control.GetByList)
		v10.GET(":customerID", control.GetBySingle)
		v10.DELETE(":customerID", control.Delete)
		v10.PATCH(":customerID", control.Update)
	}

	return router
}
