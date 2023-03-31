package permission

import (
	"app.inherited.caelus/internal/presenter/permission"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := permission.Init(db)
	authority := router.Group("authority")
	v10 := authority.Group("v1.0")
	{
		v10.GET("test", control.Test)
	}

	return router
}
