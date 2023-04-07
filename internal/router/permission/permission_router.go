package permission

import (
	present "app.eirc/internal/presenter/permission"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	authority := router.Group("authority")
	v10 := authority.Group("v1.0")
	{
		v10.GET("test", control.Test)
	}

	return router
}
