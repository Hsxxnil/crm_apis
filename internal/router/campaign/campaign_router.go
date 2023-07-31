package campaign

import (
	present "app.eirc/internal/presenter/campaign"
	"app.eirc/internal/router/middleware"
	"app.eirc/internal/router/middleware/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("campaigns")
	{
		v10.POST("", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Create)
		v10.POST("list", middleware.Verify(), auth.AuthCheckRole(db), control.GetByList)
		v10.GET("", middleware.Verify(), auth.AuthCheckRole(db), control.GetByListNoPagination)
		v10.GET(":campaignID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingle)
		v10.GET("opportunities/:campaignID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingleOpportunities)
		v10.DELETE(":campaignID", middleware.Verify(), auth.AuthCheckRole(db), control.Delete)
		v10.PATCH(":campaignID", middleware.Verify(), auth.AuthCheckRole(db), control.Update)
	}

	return router
}
