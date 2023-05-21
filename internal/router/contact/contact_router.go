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
		v10.POST("", middleware.Verify(), middleware.Transaction(db), control.Create)
		v10.POST("list", middleware.Verify(), control.GetByList)
		v10.GET(":contactID", middleware.Verify(), control.GetBySingle)
		v10.DELETE(":contactID", middleware.Verify(), control.Delete)
		v10.PATCH(":contactID", middleware.Verify(), control.Update)
	}

	return router
}
