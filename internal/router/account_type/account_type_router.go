package account_type

import (
	present "app.eirc/internal/presenter/account_type"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("accounts").Group("types")
	{
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.GET("", control.GetByList)
		v10.GET(":accountTypeID", control.GetBySingle)
		v10.DELETE(":accountTypeID", control.Delete)
		v10.PATCH(":accountTypeID", control.Update)
	}

	return router
}
