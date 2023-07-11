package routers

import (
	"github.com/gin-gonic/gin"
	"simple-mall/controllers/address"
	"simple-mall/controllers/enum"
	"simple-mall/controllers/file"
	"simple-mall/controllers/order"
	"simple-mall/controllers/product"
	"simple-mall/controllers/productCategory"
	"simple-mall/controllers/role"
	"simple-mall/controllers/shoppingCart"
	"simple-mall/controllers/slideshow"
	"simple-mall/controllers/user"
	middlewares "simple-mall/middleware"
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

		"/api/product/getRandomRecommendedProductList",
		"/api/productCategory/getAllProductCategory",
		"/api/product/getProductList",
		"/api/product/details",
		"/api/enum/getEnumsByType",
		"/api/enum/getAllEnums",
		"/api/slideshow/getSlideshowsByType",
		"/api/slideshow/getSlideshowsByType",
	}
}

// enumLoadRouter 加载枚举路由
func enumLoadRouter(r *gin.RouterGroup) {
	routes := r.Group("enum")
	{
		routes.POST("/save", enum.Save)
		routes.DELETE("/delete", enum.Delete)
		routes.GET("/details", enum.Details)
		routes.GET("/getEnumsByType", enum.GetEnumsByType)
		routes.GET("/getAllEnums", enum.GetAllEnums)
		routes.POST("/getEnumsList", enum.GetEnumsList)
	}
}

// fileLoadRouter 加载文件操作路由
func fileLoadRouter(r *gin.RouterGroup) {
	routes := r.Group("file")
	{
		routes.POST("/upload", file.Upload)
	}
}

// productLoadRouter 加载商品路由
func productLoadRouter(r *gin.RouterGroup) {
	routes := r.Group("product")
	{
		routes.POST("/save", product.Save)
		routes.DELETE("/delete", product.Delete)
		routes.GET("/details", product.Details)
		routes.POST("/getProductList", product.GetProductList)
		routes.GET("/getRandomRecommendedProductList", product.GetRandomRecommendedProductList)
	}
}

// productCategoryLoadRouter 加载商品分类路由
func productCategoryLoadRouter(r *gin.RouterGroup) {
	routes := r.Group("productCategory")
	{
		routes.POST("/save", productCategory.Save)
		routes.DELETE("/delete", productCategory.Delete)
		routes.GET("/getAllProductCategory", productCategory.GetAllProductCategory)
		routes.POST("/getProductCategoryList", productCategory.GetProductCategoryList)
	}
}

// roleLoadRouter 加载角色路由
func roleLoadRouter(r *gin.RouterGroup) {
	routes := r.Group("role")
	{
		routes.GET("/getRoleList", role.GetRoleList)
	}
}

// shoppingCartLoadRouter 加载购物车
func shoppingCartLoadRouter(r *gin.RouterGroup) {
	routes := r.Group("shoppingCart")
	{
		routes.POST("/addProductToShoppingCart", shoppingCart.AddProductToShoppingCart)
		routes.POST("/batchUpdateShoppingCartProductInfo", shoppingCart.BatchUpdateShoppingCartProductInfo)
		routes.DELETE("/deleteShoppingCartProduct", shoppingCart.DeleteShoppingCartProduct)
		routes.GET("/getShoppingCartInfo", shoppingCart.GetShoppingCartInfo)
		routes.GET("/getTheNumberOfItemsInTheShoppingCart", shoppingCart.GetTheNumberOfItemsInTheShoppingCart)
	}
}

// userLoadRouter 加载用户信息路由
func userLoadRouter(r *gin.RouterGroup) {
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

// slideshowLoadRouter 加载轮播图路由
func slideshowLoadRouter(r *gin.RouterGroup) {
	routes := r.Group("slideshow")
	{
		routes.GET("/getSlideshowsByType", slideshow.GetSlideshowsByType)
		routes.POST("/getSlideshowsList", slideshow.GetSlideshowsList)
		routes.GET("/details", slideshow.Details)
		routes.POST("/save", slideshow.Save)
		routes.DELETE("/delete", slideshow.Delete)
	}
}

// addressLoadRouter 加载地址管理
func addressLoadRouter(r *gin.RouterGroup) {
	routes := r.Group("address")
	{
		routes.GET("/getAreasByParams", address.GetAreasByParams)
		routes.POST("/save", address.Save)
		routes.POST("/getAddressInfoList", address.GetAddressInfoList)
		routes.GET("/getUserAddressInfoList", address.GetUserAddressInfoList)
		routes.GET("/details", address.Details)
		routes.DELETE("/delete", address.Delete)
		routes.GET("/setDefaultAddress", address.SetDefaultAddress)
	}
}

// orderLoadRouter 加载订单管理
func orderLoadRouter(r *gin.RouterGroup) {
	routes := r.Group("order")
	{
		routes.POST("/getOrderList", order.GetOrderList)
		routes.POST("/create", order.Create)
		routes.POST("/updateOrderStatus", order.UpdateOrderStatus)
		routes.POST("/getUserOrderList", order.GetUserOrderList)
		routes.GET("/details", order.Details)
	}
}

// LoadAllRouter 加载全部路由
func LoadAllRouter(r *gin.Engine) {
	baseRouter := r.Group("/api")
	{
		enumLoadRouter(baseRouter)
		fileLoadRouter(baseRouter)
		productLoadRouter(baseRouter)
		productCategoryLoadRouter(baseRouter)
		roleLoadRouter(baseRouter)
		shoppingCartLoadRouter(baseRouter)
		userLoadRouter(baseRouter)
		slideshowLoadRouter(baseRouter)
		addressLoadRouter(baseRouter)
		orderLoadRouter(baseRouter)
	}
}
