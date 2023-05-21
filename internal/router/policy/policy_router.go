package policy

import (
	present "app.eirc/internal/presenter/policy"
	"app.eirc/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	v10 := router.Group("crm").Group("v1.0").Group("policies")
	{
		v10.POST("", middleware.Verify(), present.AddPolicy)
		v10.GET("", middleware.Verify(), present.GetAllPolicies)
		v10.DELETE("", middleware.Verify(), present.DeletePolicy)
	}

	return router
}
