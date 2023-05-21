package opportunity

import (
	present "app.eirc/internal/presenter/opportunity"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("opportunities")
	{
		v10.POST("", middleware.Verify(), middleware.Transaction(db), control.Create)
		v10.POST("list", middleware.Verify(), control.GetByList)
		v10.GET(":opportunityID", middleware.Verify(), control.GetBySingle)
		v10.GET("campaigns/:opportunityID", middleware.Verify(), control.GetBySingleCampaigns)
		v10.DELETE(":opportunityID", middleware.Verify(), control.Delete)
		v10.PATCH(":opportunityID", middleware.Verify(), control.Update)
	}

	return router
}
