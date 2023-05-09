package contact

import (
	present "app.eirc/internal/presenter/contact"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("contacts")
	{
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.POST("list", control.GetByList)
		v10.GET("list", control.GetByList)
		v10.GET(":contactID", control.GetBySingle)
		v10.DELETE(":contactID", control.Delete)
		v10.PATCH(":contactID", control.Update)
	}

	return router
}
