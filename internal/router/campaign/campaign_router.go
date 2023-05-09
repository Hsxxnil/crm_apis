package campaign

import (
	present "app.eirc/internal/presenter/campaign"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("campaigns")
	{
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.POST("list", control.GetByList)
		v10.GET(":campaignID", control.GetBySingle)
		v10.GET("opportunities/:campaignID", control.GetBySingleOpportunities)
		v10.DELETE(":campaignID", control.Delete)
		v10.PATCH(":campaignID", control.Update)
	}

	return router
}
