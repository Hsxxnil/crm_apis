package policy

import (
	present "app.eirc/internal/presenter/policy"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	v10 := router.Group("crm").Group("v1.0").Group("policies")
	{
		v10.POST("", present.AddPolicy)
		v10.GET("", present.GetAllPolicies)
		v10.DELETE("", present.DeletePolicy)
	}

	return router
}
