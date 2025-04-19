package middleware

import "github.com/gin-gonic/gin"

func TokenVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 校验token
	}
}
