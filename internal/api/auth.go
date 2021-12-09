package api

import (
	"github.com/dollarkillerx/creeper/internal/conf"
	"github.com/gin-gonic/gin"
)

func authToken(ctx *gin.Context) {
	if conf.CONFIG.Token != "" {
		if ctx.GetHeader("token") != conf.CONFIG.Token {
			ctx.String(401, "401 token is null")
			ctx.Abort()
			return
		}
	}
}
