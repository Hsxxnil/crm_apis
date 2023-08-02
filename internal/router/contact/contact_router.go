package contact

import (
	present "app.eirc/internal/presenter/contact"
	"app.eirc/internal/router/middleware"
	"app.eirc/internal/router/middleware/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("contacts")
	{
		v10.POST("", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Create)
		v10.POST("list", middleware.Verify(), auth.AuthCheckRole(db), control.GetByList)
		v10.GET("get-by-account/:accountID", middleware.Verify(), auth.AuthCheckRole(db), control.GetByAccountIDListNoPagination)
		v10.GET(":contactID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingle)
		v10.DELETE(":contactID", middleware.Verify(), auth.AuthCheckRole(db), control.Delete)
		v10.PATCH(":contactID", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Update)
	}

	return router
}
