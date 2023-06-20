package slideshow

import (
	"github.com/gin-gonic/gin"
	"simple-mall/models"
	"simple-mall/models/slideshow"
	"simple-mall/services/slideshowServices"
	"simple-mall/utils"
)

// Save
//
//	@Summary		增加、编辑
//	@Description	增加、编辑
//	@Tags			slideshow轮播图
//	@Accept			json
//	@Produce		json
//	@Param			param	body		slideshow.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/slideshow/save [post]
func Save(ctx *gin.Context) {
	saveForm := slideshow.SaveForm{}
	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	id, err := slideshowServices.CreateAndUpdate(saveForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, map[string]any{"id": id})
}

// Delete
//
//	@Summary			根据id删除指定轮播图
//	@Description	根据id删除指定轮播图
//	@Tags			slideshow轮播图
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"轮播图id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/slideshow/delete [delete]
func Delete(ctx *gin.Context) {
	slideshowsId := ctx.Query("id")

	if slideshowsId == "" {
		utils.ResponseResultsError(ctx, "轮播图id不能为空！")
		return
	}

	err := slideshowServices.Delete(slideshowsId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}

// Details
//
//	@Summary			获取轮播图详情
//	@Description	获取轮播图详情
//	@Tags			slideshow轮播图
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"轮播图id"
//	@Success		200	{object}	utils.ResponseResultInfo{data=slideshow.SaveForm}
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/slideshow/details [get]
func Details(ctx *gin.Context) {
	slideshowsId := ctx.Query("id")

	if slideshowsId == "" {
		utils.ResponseResultsError(ctx, "轮播图id不能为空！")
		return
	}

	details, err := slideshowServices.Details(slideshowsId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, details)
}

// GetSlideshowsByType
//
//	@Summary			根据分类查询轮播图
//	@Description	根据分类查询轮播图
//	@Tags			slideshow轮播图
//	@Accept			json
//	@Produce		json
//	@Param			type	query		string	true	"轮播图类型code"
//	@Success		200		{object}	utils.ResponseResultInfo{data=[]slideshow.SaveForm}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Router			/slideshow/getSlideshowsByType [get]
func GetSlideshowsByType(ctx *gin.Context) {
	typeCode := ctx.Query("type")

	if typeCode == "" {
		utils.ResponseResultsError(ctx, "轮播图类型code不能为空！")
		return
	}

	slideshowList, err := slideshowServices.GetSlideshowsByType(typeCode)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, slideshowList)
}

// GetSlideshowsList
//
//	@Summary			分页获取轮播图列表
//	@Description	分页获取轮播图列表
//	@Tags			slideshow轮播图
//	@Accept			json
//	@Produce		json
//	@Param			param	body		slideshow.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo{data=models.PagingData{list=[]slideshow.Slideshow}}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/slideshow/getSlideshowsList [post]
func GetSlideshowsList(ctx *gin.Context) {
	listForm := slideshow.ListForm{}

	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	slideshowsList, total, err := slideshowServices.GetSlideshowsList(listForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: listForm.PageSize,
		PageNo:   listForm.PageNo,
		Total:    total,
		List:     slideshowsList,
	})
}
