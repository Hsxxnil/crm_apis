package event_user_attendee

import (
	present "app.eirc/internal/presenter/event_user_attendee"
	"app.eirc/internal/router/middleware"
	"app.eirc/internal/router/middleware/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("crm").Group("v1.0").Group("events-users-attendee")
	{
		v10.POST("", middleware.Verify(), auth.AuthCheckRole(db), middleware.Transaction(db), control.Create)
		v10.GET("", middleware.Verify(), auth.AuthCheckRole(db), control.GetByList)
		v10.GET(":eventUserAttendeeID", middleware.Verify(), auth.AuthCheckRole(db), control.GetBySingle)
		v10.DELETE(":eventUserAttendeeID", middleware.Verify(), auth.AuthCheckRole(db), control.Delete)
		v10.PATCH(":eventUserAttendeeID", middleware.Verify(), auth.AuthCheckRole(db), control.Update)
	}

	return router
}
