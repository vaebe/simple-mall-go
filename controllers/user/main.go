package user

import (
	"github.com/gin-gonic/gin"
	"simple-mall/models"
	"simple-mall/models/user"
	"simple-mall/services/userServices"
	"simple-mall/utils"
)

// GetUserList
//
//	@Summary		获取user用户列表
//	@Description	获取user用户列表
//	@Tags			user用户
//	@Accept			json
//	@Produce		json
//	@Param			param	body		user.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/user/getUserList [post]
func GetUserList(ctx *gin.Context) {
	listForm := user.ListForm{}
	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	users, total, err := userServices.GetUserList(listForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: listForm.PageSize,
		PageNo:   listForm.PageNo,
		Total:    total,
		List:     users,
	})
}

// Details
//
//	@Summary		获取用户详情
//	@Description	获取用户详情
//	@Tags			user用户
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"用户id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/user/details [get]
func Details(ctx *gin.Context) {
	userId := ctx.Query("id")

	if userId == "" {
		utils.ResponseResultsError(ctx, "用户id不能为空！")
		return
	}

	userInfo, err := userServices.GetUserDetails(userId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, userInfo)
}

// Edit
//
//	@Summary		编辑用户信息
//	@Description	编辑用户信息
//	@Tags			user用户
//	@Accept			json
//	@Produce		json
//	@Param			param	body		user.EditForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/user/edit [post]
func Edit(ctx *gin.Context) {
	editForm := user.EditForm{}
	if err := ctx.ShouldBind(&editForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")
	if editForm.ID != userId {
		utils.ResponseResultsError(ctx, "非本人不可修改用户信息！")
		return
	}

	err := userServices.Edit(editForm, userId.(int32))
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "更新用户信息成功！")
}
