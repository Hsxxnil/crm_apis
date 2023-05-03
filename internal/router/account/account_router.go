package account

import (
	present "app.eirc/internal/presenter/account"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("accounts")
	{
		v10.POST("", middleware.Transaction(db), control.Create)
		v10.GET("", control.GetByList)
		v10.GET("contacts", control.GetByListContacts)
		v10.GET(":accountID", control.GetBySingle)
		v10.GET("contacts/:accountID", control.GetBySingleContacts)
		v10.DELETE(":accountID", control.Delete)
		v10.PATCH(":accountID", control.Update)
	}

	return router
}
