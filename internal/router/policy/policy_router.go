package policy

import (
	present "crm/internal/presenter/policy"
	"crm/internal/router/middleware"
	"crm/internal/router/middleware/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	v10 := router.Group("crm").Group("v1.0").Group("policies")
	{
		v10.POST("", middleware.Verify(), auth.AuthCheckRole(db), present.AddPolicy)
		v10.GET("", middleware.Verify(), auth.AuthCheckRole(db), present.GetAllPolicies)
		v10.DELETE("", middleware.Verify(), auth.AuthCheckRole(db), present.DeletePolicy)
	}

	return router
}
