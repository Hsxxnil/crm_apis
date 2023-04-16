package lead

import (
	present "app.eirc/internal/presenter/lead"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("authority").Group("v1.0").Group("lead")
	{
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.GET("", control.GetByList)
		v10.GET(":leadID", control.GetBySingle)
		v10.DELETE(":leadID", control.Delete)
		v10.PATCH(":leadID", control.Update)
	}

	return router
}
