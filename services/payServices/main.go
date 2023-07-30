package payServices

import (
	"simple-mall/global"
	"simple-mall/models/pay"
)

// CreatePayInfo 创建支付信息
func CreatePayInfo(info pay.WeChatPayNotifyReq) error {
	saveInfo := &pay.WeChatPayNotifyInfo{
		WeChatPayNotifyReq: pay.WeChatPayNotifyReq{
			Code:        info.Code,
			Timestamp:   info.Timestamp,
			MchId:       info.MchId,
			OrderNo:     info.OrderNo,
			OutTradeNo:  info.OutTradeNo,
			PayNo:       info.PayNo,
			TotalFee:    info.TotalFee,
			Sign:        info.Sign,
			PayChannel:  info.PayChannel,
			TradeType:   info.TradeType,
			SuccessTime: info.SuccessTime,
			Attach:      info.Attach,
			Openid:      info.Openid,
		},
	}
	db := global.DB.Create(&saveInfo)
	return db.Error
}

// CreateRefundInfo 创建退款信息
func CreateRefundInfo(info pay.RefundNotifyReq) error {
	saveInfo := &pay.RefundNotifyInfo{
		RefundNotifyReq: pay.RefundNotifyReq{
			Code:        info.Code,
			Timestamp:   info.Timestamp,
			MchId:       info.MchId,
			OrderNo:     info.OrderNo,
			OutTradeNo:  info.OutTradeNo,
			PayNo:       info.PayNo,
			RefundNo:    info.RefundNo,
			OutRefundNo: info.OutRefundNo,
			PayChannel:  info.PayChannel,
			RefundFee:   info.RefundFee,
			Sign:        info.Sign,
			SuccessTime: info.SuccessTime,
		},
	}

	db := global.DB.Create(&saveInfo)
	return db.Error
}
