package pay

import "simple-mall/models"

// WeChatPayReq 微信支付请求接口参数
type WeChatPayReq struct {
	OrderId string `json:"orderId" form:"orderId" binding:"required"` // 订单id
	Price   string `json:"price" form:"price" binding:"required"`     // 商品金额
	Info    string `json:"info" form:"info" binding:"required"`       // 商品描述
}

// WeChatPayNotifyReq 微信支付通知信息参数
type WeChatPayNotifyReq struct {
	Code        string `gorm:"type:varbinary(4); not null; comment '支付结果'" json:"code" form:"code"` // 0：成功 1：失败
	Timestamp   string `gorm:"type:varbinary(40); not null; comment '时间戳'" json:"timestamp" form:"timestamp"`
	MchId       string `gorm:"type:varbinary(100); not null; comment '商户号'" json:"mch_id" form:"mch_id"`
	OrderNo     string `gorm:"type:varbinary(100); not null; comment '系统订单号'" json:"order_no" form:"order_no"`
	OutTradeNo  string `gorm:"type:varbinary(100); not null; comment '商户订单号'" json:"out_trade_no" form:"out_trade_no"`
	PayNo       string `gorm:"type:varbinary(100); not null; comment '支付宝或微信支付订单号'" json:"pay_no" form:"pay_no"`
	TotalFee    string `gorm:"type:varbinary(100); not null; comment '支付金额'" json:"total_fee" form:"total_fee"`
	Sign        string `gorm:"type:varbinary(300); not null; comment '签名'" json:"sign" form:"sign"`
	PayChannel  string `gorm:"type:varbinary(10); not null; comment '支付渠道''" json:"pay_channel" form:"pay_channel"` // 枚举值： alipay：支付宝 wxpay：微信支付'
	TradeType   string `gorm:"type:varbinary(10); not null; comment '支付类型''" json:"trade_type" form:"trade_type"`   // 枚举值： NATIVE：扫码支付 H5：H5支付 APP：APP支付 JSAPI：公众号支付 MINIPROGRAM：小程序支付'
	SuccessTime string `gorm:"type:varbinary(20); not null; comment '支付完成时间'" json:"success_time" form:"success_time"`
	Attach      string `gorm:"type:varbinary(300); not null; comment '附加数据'" json:"attach" form:"attach"`
	Openid      string `gorm:"type:varbinary(100); not null; comment '支付者信息'" json:"openid" form:"openid"`
}

// WeChatPayNotifyInfo 微信支付信息
type WeChatPayNotifyInfo struct {
	models.BaseModel
	WeChatPayNotifyReq
}

// OrderRefundReq 订单退款请求参数
type OrderRefundReq struct {
	OrderId string `json:"orderId" form:"orderId" binding:"required"` // 订单id
	Price   string `json:"price" form:"price" binding:"required"`     // 退款金额
	Info    string `json:"info" form:"info" binding:"required"`       // 退款描述
}

// RefundNotifyReq 退款通知信息参数
type RefundNotifyReq struct {
	Code        string `gorm:"type:varbinary(4); not null; comment '支付结果'" json:"code" form:"code"` // 0：成功 1：失败
	Timestamp   string `gorm:"type:varbinary(40); not null; comment '时间戳'" json:"timestamp" form:"timestamp"`
	MchId       string `gorm:"type:varbinary(100); not null; comment '商户号'" json:"mch_id" form:"mch_id"`
	OrderNo     string `gorm:"type:varbinary(100); not null; comment '系统订单号'" json:"order_no" form:"order_no"`
	OutTradeNo  string `gorm:"type:varbinary(100); not null; comment '商户订单号'" json:"out_trade_no" form:"out_trade_no"`
	PayNo       string `gorm:"type:varbinary(100); not null; comment '支付宝或微信支付订单号'" json:"pay_no" form:"pay_no"`
	RefundNo    string `gorm:"type:varbinary(100); not null; comment '系统退款单号'" json:"refund_no" form:"refund_no"`
	OutRefundNo string `gorm:"type:varbinary(100); not null; comment '商户退款单号'" json:"out_refund_no" form:"out_refund_no"`
	PayChannel  string `gorm:"type:varbinary(10); not null; comment '支付渠道'" json:"pay_channel" form:"pay_channel"` // 枚举值： alipay：支付宝 wxpay：微信支付
	RefundFee   string `gorm:"type:varbinary(100); not null; comment '退款金额'" json:"refund_fee" form:"refund_fee"`
	Sign        string `gorm:"type:varbinary(300); not null; comment '签名'" json:"sign" form:"sign"`
	SuccessTime string `gorm:"type:varbinary(100); not null; comment '支付完成时间'" json:"success_time" form:"success_time"`
}

// RefundNotifyInfo 退款信息
type RefundNotifyInfo struct {
	models.BaseModel
	RefundNotifyReq
}
