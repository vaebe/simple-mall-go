package pay

// WeChatPayReq 微信支付请求接口参数
type WeChatPayReq struct {
	OrderId string  `json:"orderId" form:"orderId" binding:"required"`
	Price   float64 `json:"price" form:"price" binding:"required"`
	Info    string  `json:"info" form:"info" binding:"required"`
}

// WeChatPayResp 微信支付响应参数
type WeChatPayResp struct {
}
