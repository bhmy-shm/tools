package middleware

import (
	"github.com/gin-gonic/gin"
)

func {{.name}}() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
