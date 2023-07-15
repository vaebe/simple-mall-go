package order

import (
	"github.com/gin-gonic/gin"
	"simple-mall/models"
	"simple-mall/models/order"
	"simple-mall/services/orderServices"
	"simple-mall/utils"
)

// Create
//
//	@Summary		订单创建
//	@Description	订单创建
//	@Tags			order订单管理
//	@Accept			json
//	@Produce		json
//	@Param			param	body		order.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/order/create [post]
func Create(ctx *gin.Context) {
	saveForm := order.SaveForm{}
	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	id, err := orderServices.Create(saveForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, id)
}

// UpdateOrderStatus
//
//	@Summary		更新订单状态
//	@Description	更新订单状态
//	@Tags			order订单管理
//	@Accept			json
//	@Produce		json
//	@Param			param	body		order.UpdateOrderStatusForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/order/updateOrderStatus [post]
func UpdateOrderStatus(ctx *gin.Context) {
	form := order.UpdateOrderStatusForm{}
	if err := ctx.ShouldBind(&form); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, ok := ctx.Get("userId")
	if !ok || userId == "" {
		utils.ResponseResultsError(ctx, "未获取到用户信息！")
		return
	}

	err := orderServices.UpdateOrderStatus(userId.(int32), form.ID, form.State)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "更新状态成功！")
}

// Details
//
//	@Summary		订单详情
//	@Description	订单详情
//	@Tags			order订单管理
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"订单id"
//	@Success		200	{object}	utils.ResponseResultInfo{data=order.Order}
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/order/details [get]
func Details(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		utils.ResponseResultsError(ctx, "订单 id 不能为空！")
		return
	}

	details, err := orderServices.Details(id)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, details)
}

// GetUserOrderList
//
//	@Summary		获取用户订单列表
//	@Description	获取用户订单列表
//	@Tags			order订单管理
//	@Accept			json
//	@Produce		json
//	@Param			param	body		order.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo{data=models.PagingData{list=[]order.Order}}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/order/getUserOrderList [post]
func GetUserOrderList(ctx *gin.Context) {
	listForm := order.ListForm{}

	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, ok := ctx.Get("userId")

	if !ok || userId == "" {
		utils.ResponseResultsError(ctx, "未获取到用户信息！")
		return
	}

	list, total, err := orderServices.GetUserOrderList(userId.(int32), listForm)
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

// GetOrderList
//
//	@Summary		获取订单分页列表
//	@Description	获取订单分页列表
//	@Tags			order订单管理
//	@Accept			json
//	@Produce		json
//	@Param			param	body		order.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo{data=models.PagingData{list=[]order.Order}}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/order/getOrderList [post]
func GetOrderList(ctx *gin.Context) {
	listForm := order.ListForm{}

	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	list, total, err := orderServices.GetOrderList(listForm)
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
