package address

import (
	"github.com/gin-gonic/gin"
	"simple-mall/models"
	"simple-mall/models/address"
	"simple-mall/services/addressServices"
	"simple-mall/utils"
)

// GetAreasByParams
//
//	@Summary		根据参数获取区域数据
//	@Description	根据参数获取区域数据
//	@Tags				address地址管理
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	false	"区域id"
//	@Param			pid	query		int	false	"上级id"
//	@Param			deep	query		int	false	"层级"
//	@Success		200		{object}	utils.ResponseResultInfo{data=[]address.AreaInfo}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/address/getAreasByParams [get]
func GetAreasByParams(ctx *gin.Context) {
	id := ctx.Query("id")
	pid := ctx.Query("pid")
	deep := ctx.Query("deep")

	areas, err := addressServices.GetAreasByParams(id, pid, deep)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, areas)
}

// Save
//
//	@Summary		地址增加、编辑
//	@Description	地址增加、编辑
//	@Tags			address地址管理
//	@Accept			json
//	@Produce		json
//	@Param			param	body		address.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/address/save [post]
func Save(ctx *gin.Context) {
	saveForm := address.SaveForm{}
	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, ok := ctx.Get("userId")

	if !ok || userId == "" {
		utils.ResponseResultsError(ctx, "未获取到用户信息！")
		return
	}

	id, err := addressServices.CreateAndUpdate(userId.(int32), saveForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, map[string]any{"id": id})
}

// Delete
//
//	@Summary		地址删除
//	@Description	地址删除
//	@Tags			address地址管理
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"地址id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/address/delete [delete]
func Delete(ctx *gin.Context) {
	addressId := ctx.Query("id")

	if addressId == "" {
		utils.ResponseResultsError(ctx, "地址id不能为空！")
		return
	}

	err := addressServices.Delete(addressId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}

// Details
//
//	@Summary		地址详情
//	@Description	地址详情
//	@Tags			address地址管理
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"地址id"
//	@Success		200	{object}	utils.ResponseResultInfo{data=address.Address}
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/address/details [get]
func Details(ctx *gin.Context) {
	addressId := ctx.Query("id")

	if addressId == "" {
		utils.ResponseResultsError(ctx, "地址 id 不能为空！")
		return
	}

	details, err := addressServices.Details(addressId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, details)
}

// GetUserAddressInfoList
//
//	@Summary		获取用户地址信息列表
//	@Description	获取用户地址信息列表
//	@Tags			address地址管理
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.ResponseResultInfo{data=address.Address}
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/address/getUserAddressInfoList [get]
func GetUserAddressInfoList(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")

	if !ok || userId == "" {
		utils.ResponseResultsError(ctx, "未获取到用户信息！")
		return
	}

	details, err := addressServices.GetUserAddressInfoList(userId.(int32))
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, details)
}

// GetAddressInfoList
//
//	@Summary		获取地址分页列表
//	@Description	获取地址分页列表
//	@Tags			address地址管理
//	@Accept			json
//	@Produce		json
//	@Param			param	body		address.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo{data=models.PagingData{list=[]address.Address}}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/address/getAddressInfoList [post]
func GetAddressInfoList(ctx *gin.Context) {
	listForm := address.ListForm{}

	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	list, total, err := addressServices.GetAddressInfoList(listForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: listForm.PageSize,
		PageNo:   listForm.PageNo,
		Total:    int32(total),
		List:     list,
	})
}

// SetDefaultAddress
//
//	@Summary		设置默认地址
//	@Description	设置默认地址
//	@Tags			address地址管理
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"地址id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/address/setDefaultAddress [get]
func SetDefaultAddress(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		utils.ResponseResultsError(ctx, "地址 id 不能为空！")
		return
	}

	userId, ok := ctx.Get("userId")
	if !ok || userId == "" {
		utils.ResponseResultsError(ctx, "未获取到用户信息！")
		return
	}

	err := addressServices.SetDefaultAddress(userId.(int32), id)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "设置默认地址成功！")
}
