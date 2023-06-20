package routers

import (
	"github.com/gin-gonic/gin"
	"simple-mall/routers/enum"
	"simple-mall/routers/file"
	"simple-mall/routers/product"
	"simple-mall/routers/productCategory"
	"simple-mall/routers/role"
	"simple-mall/routers/shoppingCart"
	"simple-mall/routers/user"
)

// GetRouterWhiteList 获取路由白名单
func GetRouterWhiteList() []string {
	return []string{
		"/api/user/login",
		"/api/user/register",
		"/api/user/details",
		"/api/user/getVerificationCode",
		"/api/enum/getAllEnums",
		"/swagger/index.html",
		"/favicon.ico",

		"/api/productCategory/getAllProductCategory",
	}
}

// LoadAllRouter 加载全部路由
func LoadAllRouter(r *gin.Engine) {
	baseRouter := r.Group("/api")
	{
		user.LoadRouter(baseRouter)
		enum.LoadRouter(baseRouter)
		file.LoadRouter(baseRouter)
		role.LoadRouter(baseRouter)
		product.LoadRouter(baseRouter)
		productCategory.LoadRouter(baseRouter)
		shoppingCart.LoadRouter(baseRouter)
	}
}
