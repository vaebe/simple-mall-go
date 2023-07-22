package pay

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"simple-mall/global"
	"simple-mall/models/pay"
	"simple-mall/utils"
	"sort"
	"strconv"
	"strings"
	"time"
)

func calculateMD5(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// LTZFGenerateSignature 生成蓝兔支付签名
func LTZFGenerateSignature(params map[string]string) string {
	sinObj := map[string]string{
		"mch_id":       params["mch_id"],
		"out_trade_no": params["out_trade_no"],
		"total_fee":    params["total_fee"],
		"body":         params["body"],
		"timestamp":    params["timestamp"],
		"notify_url":   params["notify_url"],
	}

	// 按照参数名进行字典序排序
	keys := make([]string, 0, len(sinObj))
	for k := range sinObj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接参数键值对
	var builder strings.Builder
	for _, key := range keys {
		value := sinObj[key]
		if value != "" {
			builder.WriteString(key + "=" + value + "&")
		}
	}

	// 拼接密钥
	builder.WriteString("key=" + global.LTZFConfig.SecretKey)

	// 计算MD5并转换为大写
	signature := calculateMD5(builder.String())

	return strings.ToUpper(signature)
}

// 获取蓝兔支付微信 api 参数
func getLTZFWeChatPayApiReq(payReq pay.WeChatPayReq) url.Values {
	// 请求支付接口参数
	opts := map[string]string{
		"mch_id":       global.LTZFConfig.MchId,
		"out_trade_no": payReq.OrderId,
		//"total_fee":    strconv.FormatFloat(payReq.Price, 'f', -1, 64),
		"total_fee":   "0.01", // 设置为 0.01 防止误支付
		"body":        payReq.Info,
		"timestamp":   strconv.FormatInt(time.Now().Unix(), 10),
		"notify_url":  "https://vaebe.top:53015/api/pay/weChatPayNotify",
		"attach":      payReq.OrderId,
		"time_expire": "15m",
		"sign":        "",
	}

	// 设置接口签名
	opts["sign"] = LTZFGenerateSignature(opts)

	// 格式化参数
	req := url.Values{}
	for key, value := range opts {
		req.Add(key, value)
	}

	return req
}

// WeChatPay
//
//	@Summary		微信支付
//	@Description	微信支付
//	@Tags			pay支付
//	@Accept			json
//	@Produce		json
//	@Param			param	body		pay.WeChatPayReq	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/pay/weChatPay [post]
func WeChatPay(ctx *gin.Context) {
	payReq := pay.WeChatPayReq{}
	if err := ctx.ShouldBind(&payReq); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	// todo 考虑此处是否有必要再次校验订单状态防止除未支付状态外订单调用接口

	res, err := http.PostForm("https://api.ltzf.cn/api/wxpay/native", getLTZFWeChatPayApiReq(payReq))
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
		CodeURL   string `json:"code_url"`
		QRCodeURL string `json:"QRcode_url"`
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

	utils.ResponseResultsSuccess(ctx, resp.Data)
}

// WeChatPayNotify
//
//	@Summary		微信支付通知
//	@Description	微信支付通知
//	@Tags			pay支付
//	@Accept			json
//	@Produce		json
//	@Param			param	body		pay.WeChatPayNotifyReq	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/pay/weChatPayNotify [post]
func WeChatPayNotify(ctx *gin.Context) {
	req := pay.WeChatPayNotifyReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	zap.S().Debug("支付信息", req)

	// 接收成功的处理逻辑
	ctx.String(http.StatusOK, "SUCCESS")
}
