package historical_record

import (
	present "app.eirc/internal/presenter/historical_record"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("historical-records")
	{
		// Todo:加上auth.AuthCheckRole(db)
		v10.GET("/list/:sourceID", middleware.Verify(), control.GetByList)
		v10.GET(":historicalRecordID", middleware.Verify(), control.GetBySingle)
	}

	return router
}
