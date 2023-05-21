package user

import (
	present "app.eirc/internal/presenter/user"
	"app.eirc/internal/router/middleware"
	"app.eirc/internal/router/middleware/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("users")
	{
		v10.POST("", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Create)
		v10.GET("", middleware.Verify(), auth.AuthCheckRole(db), control.GetByList)
		v10.GET(":userID", middleware.Verify(), control.GetBySingle)
		v10.DELETE(":userID", middleware.Verify(), control.Delete)
		v10.PATCH(":userID", middleware.Verify(), control.Update)
	}

	return router
}
