package productCategory

import (
	"github.com/gin-gonic/gin"
	"simple-mall/controllers/productCategory"
)

// LoadRouter 加载商品分类路由
func LoadRouter(r *gin.RouterGroup) {
	routes := r.Group("productCategory")
	{
		routes.POST("/save", productCategory.Save)
		routes.DELETE("/delete", productCategory.Delete)
		routes.GET("/getAllProductCategory", productCategory.GetAllProductCategory)
		routes.POST("/getProductCategoryList", productCategory.GetProductCategoryList)
	}
}
