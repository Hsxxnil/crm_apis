package user

import (
	present "crm/internal/presenter/user"
	"crm/internal/router/middleware"
	"crm/internal/router/middleware/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("users")
	{
		v10.POST("", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Create)
		v10.POST("list", middleware.Verify(), auth.AuthCheckRole(db), control.GetByList)
		v10.GET("", middleware.Verify(), auth.AuthCheckRole(db), control.GetByListNoPagination)
		v10.GET(":userID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingle)
		v10.DELETE(":userID", middleware.Verify(), auth.AuthCheckRole(db), control.Delete)
		v10.PATCH(":userID", middleware.Verify(), auth.AuthCheckRole(db), control.Update)
	}

	return router
}
