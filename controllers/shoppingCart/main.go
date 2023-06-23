package shoppingCart

import (
	"github.com/gin-gonic/gin"
	"simple-mall/models/shoppingCart"
	"simple-mall/services/shoppingCartServices"
	"simple-mall/utils"
)

// Save
//
//	@Summary		购物车增加、编辑
//	@Description	购物车增加、编辑
//	@Tags			shoppingCart购物车
//	@Accept			json
//	@Produce		json
//	@Param			param	body		shoppingCart.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/shoppingCart/save [post]
func Save(ctx *gin.Context) {
	saveForm := shoppingCart.SaveForm{}
	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	id, err := shoppingCartServices.CreateAndUpdate(saveForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, map[string]any{"id": id})
}

// DeleteShoppingCartProduct
//
//	@Summary		删除购物车商品
//	@Description	删除购物车商品
//	@Tags			shoppingCart购物车
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"购物车id"
//	@Param			productId	query		int	true	"商品id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/shoppingCart/deleteShoppingCartProduct [delete]
func DeleteShoppingCartProduct(ctx *gin.Context) {
	shoppingCartId := ctx.Query("id")
	if shoppingCartId == "" {
		utils.ResponseResultsError(ctx, "购物车 id 不能为空！")
		return
	}

	productId := ctx.Query("productId")
	if productId == "" {
		utils.ResponseResultsError(ctx, "商品 id 不能为空！")
		return
	}

	err := shoppingCartServices.DeleteShoppingCartProduct(shoppingCartId, productId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}

// GetShoppingCartInfoByUserId
//
//	@Summary		根据用户 id 获取购物车信息
//	@Description	根据用户 id 获取购物车信息
//	@Tags			shoppingCart购物车
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.ResponseResultInfo{data=[]shoppingCart.ShoppingCart}
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/shoppingCart/getShoppingCartInfoByUserId [get]
func GetShoppingCartInfoByUserId(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")

	if ok || userId == "" {
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
