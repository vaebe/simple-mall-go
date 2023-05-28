package enum

import (
	"github.com/gin-gonic/gin"
	"simple-mall/models"
	"simple-mall/models/enum"
	"simple-mall/services/enumServices"
	"simple-mall/utils"
)

// Save
//
//	@Summary		增加、编辑
//	@Description	增加、编辑
//	@Tags			enum枚举
//	@Accept			json
//	@Produce		json
//	@Param			param	body		enum.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/enum/save [post]
func Save(ctx *gin.Context) {
	saveForm := enum.SaveForm{}
	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	id, err := enumServices.CreateAndUpdate(saveForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, map[string]any{"id": id})
}

// Delete todo 考虑增加类型 如系统则不能被删除
//
//	@Summary		根据id删除指定枚举
//	@Description	根据id删除指定枚举
//	@Tags			enum枚举
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"枚举id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/enum/delete [delete]
func Delete(ctx *gin.Context) {
	enumsId := ctx.Query("id")

	if enumsId == "" {
		utils.ResponseResultsError(ctx, "枚举id不能为空！")
		return
	}

	err := enumServices.Delete(enumsId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}

// Details
//
//	@Summary		获取枚举详情
//	@Description	获取枚举详情
//	@Tags			enum枚举
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"枚举id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/enum/details [get]
func Details(ctx *gin.Context) {
	enumsId := ctx.Query("id")

	if enumsId == "" {
		utils.ResponseResultsError(ctx, "枚举id不能为空！")
		return
	}

	details, err := enumServices.Details(enumsId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, details)
}

// GetEnumsByType
//
//	@Summary		根据分类查询枚举
//	@Description	根据分类查询枚举
//	@Tags			enum枚举
//	@Accept			json
//	@Produce		json
//	@Param			type	query		string	true	"枚举类型code"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/enum/getEnumsByType [get]
func GetEnumsByType(ctx *gin.Context) {
	typeCode := ctx.Query("type")

	if typeCode == "" {
		utils.ResponseResultsError(ctx, "枚举类型code不能为空！")
		return
	}

	enumList, err := enumServices.GetEnumsByType(typeCode)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, enumList)
}

// GetAllEnums
//
//	@Summary		获取全部数据
//	@Description	获取全部数据
//	@Tags			enum枚举
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Router			/enum/getAllEnums [get]
func GetAllEnums(ctx *gin.Context) {
	enumsObj, err := enumServices.GetAllEnums()
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, enumsObj)
}

// GetEnumsList
//
//	@Summary		分页获取枚举列表
//	@Description	分页获取枚举列表
//	@Tags			enum枚举
//	@Accept			json
//	@Produce		json
//	@Param			param	body		enum.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Router			/enum/getEnumsList [post]
func GetEnumsList(ctx *gin.Context) {
	listForm := enum.ListForm{}

	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	enumsList, total, err := enumServices.GetEnumsList(listForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: listForm.PageSize,
		PageNo:   listForm.PageNo,
		Total:    total,
		List:     enumsList,
	})
}
