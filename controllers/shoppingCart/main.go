package shoppingCart

import (
	"github.com/gin-gonic/gin"
	"simple-mall/models/shoppingCart"
	"simple-mall/services/shoppingCartServices"
	"simple-mall/utils"
)

// AddProductToShoppingCart
//
//	@Summary		添加商品到购物车
//	@Description	添加商品到购物车
//	@Tags			shoppingCart购物车
//	@Accept			json
//	@Produce		json
//	@Param			param	body		shoppingCart.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/shoppingCart/addProductToShoppingCart [post]
func AddProductToShoppingCart(ctx *gin.Context) {
	saveForm := shoppingCart.SaveForm{}
	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	id, err := shoppingCartServices.AddProductToShoppingCart(saveForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, map[string]any{"id": id})
}

// BatchUpdateShoppingCartProductInfo
//
//	@Summary		批量更新购物车商品信息
//	@Description	批量更新购物车商品信息
//	@Tags			shoppingCart购物车
//	@Accept			json
//	@Produce		json
//	@Param			param	body		[]shoppingCart.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/shoppingCart/batchUpdateShoppingCartProductInfo [post]
func BatchUpdateShoppingCartProductInfo(ctx *gin.Context) {
	var editFormList []shoppingCart.SaveForm
	if err := ctx.ShouldBind(&editFormList); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	err := shoppingCartServices.BatchUpdateShoppingCartProductInfo(editFormList)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "更新成功")
}

// DeleteShoppingCartProduct
//
//	@Summary		删除购物车商品
//	@Description	删除购物车商品
//	@Tags			shoppingCart购物车
//	@Accept			json
//	@Produce		json
//	@Param			productId	query		int	true	"商品id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/shoppingCart/deleteShoppingCartProduct [delete]
func DeleteShoppingCartProduct(ctx *gin.Context) {
	productId := ctx.Query("productId")
	if productId == "" {
		utils.ResponseResultsError(ctx, "商品 id 不能为空！")
		return
	}

	userId, ok := ctx.Get("userId")

	if !ok || userId == "" {
		utils.ResponseResultsError(ctx, "未获取到用户信息！")
		return
	}

	err := shoppingCartServices.DeleteShoppingCartProduct(userId.(int32), productId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}

// GetShoppingCartInfo
//
//	@Summary		根据用户 id 获取购物车信息
//	@Description	根据用户 id 获取购物车信息
//	@Tags			shoppingCart购物车
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.ResponseResultInfo{data=[]shoppingCart.ShoppingCart}
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/shoppingCart/getShoppingCartInfo [get]
func GetShoppingCartInfo(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")

	if !ok || userId == "" {
		utils.ResponseResultsError(ctx, "未获取到用户信息！")
		return
	}

	list, err := shoppingCartServices.GetShoppingCartInfoByUserId(userId.(int32))
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, list)
}

// GetTheNumberOfItemsInTheShoppingCart
//
//	@Summary	 根据用户 id 获取购物车商品数量
//	@Description	根据用户 id 获取购物车商品数量
//	@Tags			shoppingCart购物车
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.ResponseResultInfo{data=int32}
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/shoppingCart/getTheNumberOfItemsInTheShoppingCart [get]
func GetTheNumberOfItemsInTheShoppingCart(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")

	if !ok || userId == "" {
		utils.ResponseResultsError(ctx, "未获取到用户信息！")
		return
	}

	list, err := shoppingCartServices.GetTheNumberOfItemsInTheShoppingCart(userId.(int32))
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, list)
}
