package role

import (
	present "app.eirc/internal/presenter/role"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("roles")
	{
		v10.POST("", middleware.Verify(), middleware.Transaction(db), control.Create)
		v10.GET("", middleware.Verify(), control.GetByList)
		v10.GET(":roleID", control.GetBySingle)
		v10.DELETE(":roleID", control.Delete)
		v10.PATCH(":roleID", control.Update)
	}

	return router
}
