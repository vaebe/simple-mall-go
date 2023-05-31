package shoppingCart

import (
	"github.com/gin-gonic/gin"
	"simple-mall/controllers/shoppingCart"
)

// LoadRouter 加载商品分类路由
func LoadRouter(r *gin.RouterGroup) {
	routes := r.Group("shoppingCart")
	{
		routes.POST("/save", shoppingCart.Save)
		routes.DELETE("/delete", shoppingCart.Delete)
		routes.GET("/getShoppingCartInfoByUserId", shoppingCart.GetShoppingCartInfoByUserId)
	}
}
