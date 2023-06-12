package user

import (
	"github.com/gin-gonic/gin"
	"simple-mall/controllers/user"
	middlewares "simple-mall/middleware"
)

// LoadRouter 加载用户信息路由
func LoadRouter(r *gin.RouterGroup) {
	routes := r.Group("user")
	{
		routes.POST("/login", user.Login)
		routes.POST("/register", user.Register)
		routes.POST("/getVerificationCode", user.GetVerificationCode)
		routes.GET("/details", user.Details)
		routes.POST("/save", user.Save)
		routes.DELETE("/delete", user.Delete)
		// 非管理员不能获取用户列表
		routes.POST("/getUserList", middlewares.IsAdmin(), user.GetUserList)
	}
}
