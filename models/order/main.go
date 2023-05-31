package order

import "simple-mall/models"

// Order 订单信息
type Order struct {
	models.BaseModel
	UserId          int32  `gorm:"type:int; not null; comment '用户id'" json:"userId"`
	State           string `gorm:"type:varbinary(50); not null; comment '订单状态'" json:"state"`
	TotalPrice      int64  `gorm:"type:int; not null; comment '订单总价'" json:"totalPrice"`
	Remark          string `gorm:"type:varbinary(300); comment '订单留言'" json:"remark"`
	ReceiverName    string `gorm:"type:varbinary(50); not null; comment '收货人姓名'" json:"receiverName"`
	ReceiverPhone   string `gorm:"type:varbinary(20); not null; comment '收货人手机号'" json:"receiverPhone"`
	ReceiverAddress string `gorm:"type:varbinary(500); not null; comment '收货人地址'" json:"receiverAddress"`
	PaymentMethod   string `gorm:"type:varbinary(50); not null; comment '支付方式'" json:"paymentMethod"`
}

// OrderDetails 订单详情
type OrderDetails struct {
	models.BaseModel
	OrderId   int32 `gorm:"type:int; not null; comment '商品订单id'" json:"orderId"`
	ProductId int32 `gorm:"type:int; not null; comment '商品id'" json:"productId"`
	Count     int32 `gorm:"type:int; not null; comment '商品数量'" json:"count"`
	Price     int32 `gorm:"type:int; not null; comment '商品价格'" json:"price"`
}
