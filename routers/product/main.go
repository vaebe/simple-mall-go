package product

import (
	"github.com/gin-gonic/gin"
	"simple-mall/controllers/product"
)

// LoadRouter 加载商品路由
func LoadRouter(r *gin.RouterGroup) {
	routes := r.Group("product")
	{
		routes.POST("/save", product.Save)
		routes.DELETE("/delete", product.Delete)
		routes.GET("/details", product.Details)
		routes.POST("/getProductList", product.GetProductList)
	}
}
