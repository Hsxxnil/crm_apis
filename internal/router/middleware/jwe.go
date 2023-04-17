package middleware

import (
	"app.eirc/config"
	"app.eirc/internal/interactor/pkg/jwx"
	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Verify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		j := &jwx.JWE{
			PrivateKey: config.AccessPrivateKey,
			Token:      ctx.GetHeader("Authorization"),
		}

		if len(j.Token) == 0 {
			ctx.AbortWithStatusJSON(http.StatusOK, code.GetCodeMessage(code.JWTRejected, "AccessToken is null"))
			return
		}

		j, err := j.Verify()
		if err != nil {
			log.Error(err)
			ctx.AbortWithStatusJSON(http.StatusOK, code.GetCodeMessage(code.JWTRejected, err.Error()))
			return
		}

		ctx.Set("user_id", j.Other["user_id"])
		ctx.Set("company_id", j.Other["company_id"])
		ctx.Next()
	}
}
