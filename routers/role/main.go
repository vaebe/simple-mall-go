package role

import (
	"github.com/gin-gonic/gin"
	"simple-mall/controllers/role"
)

// LoadRouter 加载角色路由
func LoadRouter(r *gin.RouterGroup) {
	routes := r.Group("role")
	{
		routes.GET("/getRoleList", role.GetRoleList)
	}
}
