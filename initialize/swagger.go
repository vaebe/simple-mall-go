package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"os/exec"
	"simple-mall/docs"
	"simple-mall/global"
)

// InitSwagger 初始化swagger
func InitSwagger(r *gin.Engine, serviceAddress string) {
	// 非开发环境环境无需生成swagger
	if global.ENV == "dev" {
		// 执行命令生成 swagger
		cmd := exec.Command("swag", "init")
		if _, err := cmd.CombinedOutput(); err != nil {
			panic(err)
		}
	}

	docs.SwaggerInfo.Title = "simple-mall API"
	docs.SwaggerInfo.Description = "simple-mall API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = serviceAddress
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	zap.S().Infof("swagger访问地址:http://%s/swagger/index.html", serviceAddress)
}
