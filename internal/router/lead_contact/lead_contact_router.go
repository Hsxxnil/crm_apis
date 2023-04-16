package lead_contact

import (
	present "app.eirc/internal/presenter/lead_contact"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("authority").Group("v1.0").Group("leads").Group("contacts")
	{
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.GET("", control.GetByList)
		v10.GET(":leadContactID", control.GetBySingle)
		v10.DELETE(":leadContactID", control.Delete)
		v10.PATCH(":leadContactID", control.Update)
	}

	return router
}
