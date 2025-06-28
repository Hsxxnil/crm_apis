package role

import (
	present "crm/internal/presenter/role"
	"crm/internal/router/middleware"
	"crm/internal/router/middleware/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("roles")
	{
		v10.POST("", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Create)
		v10.GET("", middleware.Verify(), auth.AuthCheckRole(db), control.GetByList)
		v10.GET(":roleID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingle)
		v10.DELETE(":roleID", middleware.Verify(), auth.AuthCheckRole(db), control.Delete)
		v10.PATCH(":roleID", middleware.Verify(), auth.AuthCheckRole(db), control.Update)
	}

	return router
}
