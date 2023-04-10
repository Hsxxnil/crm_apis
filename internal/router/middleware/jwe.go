package middleware

import (
	"app.eirc/internal/interactor/pkg/jwx"
	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Verify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		privateKey := ""
		j := &jwx.JWE{
			PrivateKey: privateKey,
			Token:      ctx.GetHeader("Authorization"),
		}

		if len(j.Token) == 0 {
			ctx.AbortWithStatusJSON(http.StatusOK, code.GetCodeMessage(code.JWTRejected, "jwe is null"))
			return
		}

		j, err := j.Verify()
		if err != nil {
			log.Error(err)
			ctx.AbortWithStatusJSON(http.StatusOK, code.GetCodeMessage(code.JWTRejected, err.Error()))
			return
		}

		ctx.Set("account_id", j.Other["account_id"])
		ctx.Set("company_id", j.Other["company_id"])
		ctx.Next()
	}
}
