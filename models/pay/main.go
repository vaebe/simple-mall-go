package pay

// WeChatPayReq 微信支付请求接口参数
type WeChatPayReq struct {
	OrderId string  `json:"orderId" form:"orderId" binding:"required"`
	Price   float64 `json:"price" form:"price" binding:"required"`
	Info    string  `json:"info" form:"info" binding:"required"`
}

// WeChatPayNotifyReq 微信支付通知接口
type WeChatPayNotifyReq struct {
	Code        string `json:"code" form:"code"`                 // 支付结果 0：成功 1：失败
	Timestamp   string `json:"timestamp" form:"timestamp"`       // 时间戳
	MchId       string `json:"mch_id" form:"mch_id"`             // 商户号
	OrderNo     string `json:"order_no" form:"order_no"`         // 系统订单号
	OutTradeNo  string `json:"out_trade_no" form:"out_trade_no"` // 商户订单号
	PayNo       string `json:"pay_no" form:"pay_no"`             // 支付宝或微信支付订单号
	TotalFee    string `json:"total_fee" form:"total_fee"`       // 支付金额
	Sign        string `json:"sign" form:"sign"`                 // 签名
	PayChannel  string `json:"pay_channel" form:"pay_channel"`   // 支付渠道，枚举值： alipay：支付宝 wxpay：微信支付
	TradeType   string `json:"trade_type" form:"trade_type"`     // 支付类型，枚举值： NATIVE：扫码支付 H5：H5支付 APP：APP支付 JSAPI：公众号支付 MINIPROGRAM：小程序支付
	SuccessTime string `json:"success_time" form:"success_time"` // 支付完成时间
	Attach      string `json:"attach" form:"attach"`             // 附加数据
	Openid      string `json:"openid" form:"openid"`             // 支付者信息
}
