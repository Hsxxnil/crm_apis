package account

import (
	present "app.eirc/internal/presenter/account"
	"app.eirc/internal/router/middleware"
	"app.eirc/internal/router/middleware/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("accounts")
	{
		v10.POST("", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Create)
		v10.POST("list", middleware.Verify(), auth.AuthCheckRole(db), control.GetByList)
		v10.GET(":accountID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingle)
		v10.GET("contacts/:accountID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingleContacts)
		v10.DELETE(":accountID", middleware.Verify(), auth.AuthCheckRole(db), control.Delete)
		v10.PATCH(":accountID", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Update)
	}

	return router
}
