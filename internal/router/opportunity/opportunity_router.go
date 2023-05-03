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
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.GET("", control.GetByList)
		v10.GET("campaigns", control.GetByListCampaigns)
		v10.GET(":opportunityID", control.GetBySingle)
		v10.GET("campaigns/:opportunityID", control.GetBySingleCampaigns)
		v10.DELETE(":opportunityID", control.Delete)
		v10.PATCH(":opportunityID", control.Update)
	}

	return router
}
