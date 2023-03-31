package permission

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Control interface {
	Test(ctx *gin.Context)
}

type control struct {
}

func Init(db *gorm.DB) Control {
	return &control{}
}

func (c *control) Test(ctx *gin.Context) {

	ctx.JSON(200, "ok!")
}
