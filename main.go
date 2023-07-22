package main

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"simple-mall/global"
	"simple-mall/initialize"
	middlewares "simple-mall/middleware"
	"simple-mall/routers"
	"simple-mall/tasks"
	"time"
)

// @contact.name				API Support
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	// 初始化配置
	initialize.InitConfig()

	r := gin.Default()
	// 替换 gin logger 为 zap
	r.Use(ginzap.Ginzap(global.Logger, time.RFC3339, true), ginzap.RecoveryWithZap(global.Logger, true))

	// 添加跨域中间件
	r.Use(middlewares.Cors())

	// 添加jwt中间件
	r.Use(middlewares.JWTAuth(routers.GetRouterWhiteList()))

	// 加载路由
	routers.LoadAllRouter(r)

	// 初始化定时任务
	tasks.InitTasks()

	port := 51015
	serviceAddress := fmt.Sprintf("%s:%d", "127.0.0.1", port)

	// 初始化swagger
	initialize.InitSwagger(r, serviceAddress)

	err := r.Run(serviceAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
}
