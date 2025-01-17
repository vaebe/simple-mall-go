package product

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"simple-mall/models"
	"simple-mall/models/product"
	"simple-mall/services/productServices"
	"simple-mall/utils"
	"strconv"
)

// Save
//
//	@Summary		商品增加、编辑
//	@Description	商品增加、编辑
//	@Tags			product商品
//	@Accept			json
//	@Produce		json
//	@Param			param	body		product.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/product/save [post]
func Save(ctx *gin.Context) {
	saveForm := product.SaveForm{}
	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	id, err := productServices.CreateAndUpdate(saveForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, map[string]any{"id": id})
}

// Delete
//
//	@Summary		商品删除
//	@Description	商品删除
//	@Tags			product商品
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"商品id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/product/delete [delete]
func Delete(ctx *gin.Context) {
	productId := ctx.Query("id")

	if productId == "" {
		utils.ResponseResultsError(ctx, "商品id不能为空！")
		return
	}

	err := productServices.Delete(productId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}

// Details
//
//	@Summary		商品详情
//	@Description	商品详情
//	@Tags			product商品
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"商品id"
//	@Success		200	{object}	utils.ResponseResultInfo{data=product.Product}
//	@Failure		500	{object}	utils.EmptyInfo
//	@Router			/product/details [get]
func Details(ctx *gin.Context) {
	productId := ctx.Query("id")

	if productId == "" {
		utils.ResponseResultsError(ctx, "商品 id 不能为空！")
		return
	}

	details, err := productServices.Details(productId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, details)
}

// GetProductList
//
//	@Summary		获取商品分页列表
//	@Description	获取商品分页列表
//	@Tags			product商品
//	@Accept			json
//	@Produce		json
//	@Param			param	body		product.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo{data=models.PagingData{list=[]product.Product}}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/product/getProductList [post]
func GetProductList(ctx *gin.Context) {
	listForm := product.ListForm{}

	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	list, total, err := productServices.GetProductList(listForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: listForm.PageSize,
		PageNo:   listForm.PageNo,
		Total:    total,
		List:     list,
	})
}

// GetRandomRecommendedProductList
//
//	@Summary		获取随机推荐商品列表
//	@Description	获取随机推荐商品列表
//	@Tags			product商品
//	@Accept			json
//	@Produce		json
//	@Param			total	query		int	true	"数量" default(20) maximum(40)
//	@Success		200		{object}	utils.ResponseResultInfo{data=[]product.Product}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Router			/product/getRandomRecommendedProductList [get]
func GetRandomRecommendedProductList(ctx *gin.Context) {
	total, _ := strconv.Atoi(ctx.Query("total"))
	if total > 40 {
		utils.ResponseResultsError(ctx, "获取推荐商品总数不能大于 40")
		return
	}

	list, err := productServices.GetRandomRecommendedProductList(total)
	if err != nil {
		zap.S().Info("获取随机推荐商品列表:", err)
		utils.ResponseResultsError(ctx, "获取数据失败")
		return
	}

	utils.ResponseResultsSuccess(ctx, list)
}
