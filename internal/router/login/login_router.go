package login

import (
	present "app.eirc/internal/presenter/login"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0")
	{
		v10.POST("login", control.Login)
		v10.POST("refresh", control.Refresh)
	}

	return router
}
