package industry

import (
	present "app.eirc/internal/presenter/industry"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("industries")
	{
		v10.POST("", middleware.Verify(), middleware.Transaction(db), control.Create)
		v10.GET("", middleware.Verify(), control.GetByList)
		v10.GET(":industryID", middleware.Verify(), control.GetBySingle)
		v10.DELETE(":industryID", middleware.Verify(), control.Delete)
		v10.PATCH(":industryID", middleware.Verify(), control.Update)
	}

	return router
}
