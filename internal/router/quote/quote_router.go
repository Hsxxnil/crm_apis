package quote

import (
	present "app.eirc/internal/presenter/quote"
	"app.eirc/internal/router/middleware"
	"app.eirc/internal/router/middleware/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("quotes")
	{
		v10.POST("", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Create)
		v10.POST("list", middleware.Verify(), auth.AuthCheckRole(db), control.GetByList)
		v10.GET(":quoteID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingle)
		v10.GET("products/:quoteID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingleProducts)
		v10.DELETE(":quoteID", middleware.Verify(), auth.AuthCheckRole(db), control.Delete)
		v10.PATCH(":quoteID", middleware.Verify(), auth.AuthCheckRole(db), control.Update)
	}

	return router
}
