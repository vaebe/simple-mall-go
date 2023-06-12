package productCategory

import (
	"github.com/gin-gonic/gin"
	"simple-mall/models"
	"simple-mall/models/productCategory"
	"simple-mall/services/productCategoryServices"
	"simple-mall/utils"
)

// Save
//
//	@Summary		商品分类增加、编辑
//	@Description	商品分类增加、编辑
//	@Tags			productCategory商品分类
//	@Accept			json
//	@Produce		json
//	@Param			param	body		productCategory.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/productCategory/save [post]
func Save(ctx *gin.Context) {
	saveForm := productCategory.SaveForm{}
	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	id, err := productCategoryServices.CreateAndUpdate(saveForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, map[string]any{"id": id})
}

// Delete
//
//	@Summary		商品分类删除
//	@Description	商品分类删除
//	@Tags			productCategory商品分类
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"商品分类id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/productCategory/delete [delete]
func Delete(ctx *gin.Context) {
	productCategoryId := ctx.Query("id")

	if productCategoryId == "" {
		utils.ResponseResultsError(ctx, "商品分类id不能为空！")
		return
	}

	err := productCategoryServices.Delete(productCategoryId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}

// GetAllProductCategory
//
//	@Summary		获取全部商品分类
//	@Description	获取全部商品分类
//	@Tags			productCategory商品分类
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.ResponseResultInfo{data=[]productCategory.ProductCategory}
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/productCategory/getAllProductCategory [get]
func GetAllProductCategory(ctx *gin.Context) {
	list, err := productCategoryServices.GetAllProductCategory()
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, list)
}

// GetProductCategoryList
//
//	@Summary		获取商品分类分页列表
//	@Description	获取商品分类分页列表
//	@Tags			productCategory商品分类
//	@Accept			json
//	@Produce		json
//	@Param			param	body		productCategory.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo{data=models.PagingData{list=[]productCategory.ProductCategory}}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/productCategory/getProductCategoryList [post]
func GetProductCategoryList(ctx *gin.Context) {
	listForm := productCategory.ListForm{}

	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	list, total, err := productCategoryServices.GetProductCategoryList(listForm)
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
