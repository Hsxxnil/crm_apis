package user

import (
	present "app.eirc/internal/presenter/user"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("users")
	{
		v10.POST("", middleware.Verify(), middleware.Transaction(db), control.Create)
		v10.GET("", control.GetByList)
		v10.GET(":userID", control.GetBySingle)
		v10.DELETE(":userID", control.Delete)
		v10.PATCH(":userID", control.Update)
	}

	return router
}
