package contract

import (
	present "app.eirc/internal/presenter/contract"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("contracts")
	{
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.POST("list", control.GetByList)
		v10.GET(":contractID", control.GetBySingle)
		v10.DELETE(":contractID", control.Delete)
		v10.PATCH(":contractID", control.Update)
	}

	return router
}
