package pay

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"simple-mall/global"
	"simple-mall/models/pay"
	"simple-mall/services/orderServices"
	"simple-mall/services/payServices"
	"simple-mall/utils"
	"strconv"
	"time"
)

// 获取退款签名对象
func getLTZFOrderRefundSinObj(params map[string]string) string {
	sinObj := map[string]string{
		"mch_id":        params["mch_id"],
		"out_trade_no":  params["out_trade_no"],
		"out_refund_no": params["out_refund_no"],
		"timestamp":     params["timestamp"],
		"refund_fee":    params["refund_fee"],
	}

	return LTZFGenerateSignature(sinObj)
}

// 获取退款参数
func getLTZFOrderRefundReq(orderRefundReq pay.OrderRefundReq) url.Values {
	opts := map[string]string{
		"mch_id":        global.LTZFConfig.MchId, // 商户号
		"out_trade_no":  orderRefundReq.OrderId,  // 商户订单号
		"out_refund_no": orderRefundReq.OrderId,  // 商户退款单号 todo 考虑是否生成退款订单信息，如果每次都全部退款则不需要
		"timestamp":     strconv.FormatInt(time.Now().Unix(), 10),
		//"refund_fee":    orderRefundReq.Price,
		"refund_fee":  "0.01",                                         // 支付设置为0.01 所以这里也是
		"refund_desc": orderRefundReq.Info,                            // 退款信息
		"notify_url":  "https://vaebe.top:53015/api/pay/refundNotify", // 退款通知地址
		"sign":        "",
	}

	// 设置接口签名
	opts["sign"] = getLTZFOrderRefundSinObj(opts)

	// 格式化参数
	req := url.Values{}
	for key, value := range opts {
		req.Add(key, value)
	}

	return req
}

// OrderRefund
//
//	@Summary	 订单退款
//	@Description	订单退款
//	@Tags			pay支付
//	@Accept			json
//	@Produce		json
//	@Param			param	body		pay.OrderRefundReq	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/pay/orderRefund [post]
func OrderRefund(ctx *gin.Context) {
	req := pay.OrderRefundReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	res, err := http.PostForm("https://api.ltzf.cn/api/wxpay/refund_order", getLTZFOrderRefundReq(req))
	if err != nil {
		// 处理请求错误
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		// 处理读取响应体错误
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	// 解析接口响应数据
	type Data struct {
		MchId       string `json:"mch_id"`        // 商户号
		OutTradeNo  string `json:"out_trade_no"`  // 商户订单号
		OutRefundNo string `json:"out_refund_no"` // 商户退款单号
		OrderNo     string `json:"order_no"`      // 系统退款单号
		PayRefundNo string `json:"pay_refund_no"` // 微信支付退款单号
	}

	type Response struct {
		Code      int    `json:"code"`
		Data      Data   `json:"data"`
		Msg       string `json:"msg"`
		RequestID string `json:"request_id"`
	}

	// 解析JSON数据
	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		// 处理解析JSON错误
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	if resp.Code != 0 {
		utils.ResponseResultsError(ctx, resp.Msg)
		return
	}

	// 更新订单状态为退款中
	orderId, err := strconv.Atoi(req.OrderId)
	if err != nil {
		zap.S().Debug("订单 id 转换错误：", err.Error())
		return
	}

	err = orderServices.UpdateOrderStatus(int32(orderId), "06")
	if err != nil {
		zap.S().Debug("更新退款订单为退款中错误：", err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, resp.Data)
}

// RefundNotify
//
//	@Summary	 退款通知
//	@Description	退款通知
//	@Tags			pay支付
//	@Accept			json
//	@Produce		json
//	@Param			param	body		pay.RefundNotifyReq	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Router			/pay/refundNotify [post]
func RefundNotify(ctx *gin.Context) {
	req := pay.RefundNotifyReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	zap.S().Debug("退款信息：", req)

	// 根据 id 获取订单状态
	orderInfo, _ := orderServices.Details(req.OutTradeNo)

	// todo 暂不考虑部分退款 退款时间 一般为一到三分钟 银行卡一到三天 不做及时推送，后续考虑增加用户消息推送
	// 状态等于退款中 修改状态为已退款并发送通知
	if orderInfo.State == "06" {
		orderId, err := strconv.Atoi(req.OutTradeNo)
		if err != nil {
			zap.S().Debug("订单 id 转换错误：", err.Error())
			return
		}

		err = orderServices.UpdateOrderStatus(int32(orderId), "07")
		if err != nil {
			zap.S().Debug("退款完成更新订单信息错误：", err.Error())
			return
		}

		// 创建退款信息
		err = payServices.CreateRefundInfo(req)
		if err != nil {
			zap.S().Debug("创建支付信息错误：", err.Error())
			return
		}
	}

	// 接收成功的处理逻辑
	ctx.String(http.StatusOK, "SUCCESS")
}

// QueryRefundResult
//
//	@Summary	 查询退款结果
//	@Description	查询退款结果
//	@Tags			pay支付
//	@Accept			json
//	@Produce		json
//	@Param			param	body		pay.RefundNotifyReq	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/pay/queryRefundResult [post]
func QueryRefundResult(ctx *gin.Context) {

}
