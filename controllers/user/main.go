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
//	@Success		200		{object}	utils.ResponseResultInfo{data=models.PagingData{list=[]user.User}}
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
//	@Success		200	{object}	utils.ResponseResultInfo{data=user.User}
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

// Save
//
//	@Summary		增加、编辑
//	@Description	增加、编辑
//	@Tags			user用户
//	@Accept			json
//	@Produce		json
//	@Param			param	body		user.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/user/save [post]
func Save(ctx *gin.Context) {
	saveForm := user.SaveForm{}
	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	id, err := userServices.CreateAndUpdate(saveForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, map[string]any{"id": id})
}

// Delete
//
//	@Summary		根据id删除用户
//	@Description	根据id删除用户
//	@Tags			user用户
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"用户id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/user/delete [delete]
func Delete(ctx *gin.Context) {
	userId := ctx.Query("id")

	if userId == "" {
		utils.ResponseResultsError(ctx, "用户id不能为空！")
		return
	}

	err := userServices.Delete(userId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}
