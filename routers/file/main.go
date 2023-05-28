package file

import (
	"github.com/gin-gonic/gin"
	"simple-mall/controllers/file"
)

// LoadRouter 加载文件操作路由
func LoadRouter(r *gin.RouterGroup) {
	routes := r.Group("file")
	{
		routes.POST("/upload", file.Upload)
	}
}
