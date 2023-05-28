package middlewares

import (
	"github.com/gin-gonic/gin"
	"simple-mall/utils"
)

// IsAdmin 验证登陆用户是否是admin
func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取jwt验证后设置的用户信息
		authorityId, _ := ctx.Get("authorityId")

		if authorityId.(int32) != 2 {
			utils.ResponseResultsError(ctx, "用户无权限！")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
