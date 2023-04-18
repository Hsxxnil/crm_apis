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
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.GET("", control.GetByList)
		v10.GET(":industryID", control.GetBySingle)
		v10.DELETE(":industryID", control.Delete)
		v10.PATCH(":industryID", control.Update)
	}

	return router
}
