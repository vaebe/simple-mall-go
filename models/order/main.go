package order

import (
	"simple-mall/models"
	"simple-mall/models/address"
)

// OrderProducts 订单商品详情
type OrderProducts struct {
	models.BaseModel
	OrderId   int32  `gorm:"type:int; not null; comment '商品订单id'" json:"orderId"`
	Name      string `gorm:"type:varbinary(200); not null; comment '商品名称'" json:"name"`
	ProductId int32  `gorm:"type:int; not null; comment '商品id'" json:"productId"`
	Count     int32  `gorm:"type:int; not null; comment '商品数量'" json:"count"`
	Price     int32  `gorm:"type:int; not null; comment '商品价格'" json:"price"`
	Info      string `gorm:"type:varbinary(300); not null; comment '商品简介'" json:"info"`
	Picture   string `gorm:"type:varbinary(500); not null; comment '商品图片'" json:"picture"`
}

// Order 订单信息
// 订单状态 待支付 00 已支付 01 处理中 02 已发货 03 已完成 04 已取消 05 退款中 06 已退款 07 异常 08
type Order struct {
	models.BaseModel
	UserId         int32           `gorm:"type:int; not null; comment '用户id'" json:"userId"`
	State          string          `gorm:"type:varbinary(10); not null; comment '订单状态" json:"state"`
	TotalPrice     int64           `gorm:"type:int; not null; comment '订单总价'" json:"totalPrice"`
	Remark         string          `gorm:"type:varbinary(300); comment '订单留言'" json:"remark"`
	AddressId      int32           `gorm:"type:int; not null; comment '地址id'" json:"addressId"`
	PaymentMethod  string          `gorm:"type:varbinary(10); comment '支付方式 00 支付宝 01 微信 02 银行卡'" json:"paymentMethod"`
	ShippingMethod string          `gorm:"type:varbinary(10); default: '00'; comment '运送方式 00 普通配送" json:"shippingMethod"`
	Products       []OrderProducts `json:"products"`
}

// ListForm 订单列表
type ListForm struct {
	models.PaginationParameters
	State string `json:"state" form:"state"`
}

// SaveForm 订单保存
type SaveForm struct {
	ID            int32           `form:"id" json:"id"`
	UserId        int32           `json:"userId" form:"userId" binding:"required"`
	TotalPrice    int64           `json:"totalPrice" form:"totalPrice" binding:"required"`
	Remark        string          `json:"remark" form:"remark" binding:"required"`
	AddressId     int32           `json:"addressId" form:"addressId" binding:"required"`
	PaymentMethod string          `json:"paymentMethod" form:"paymentMethod"`
	Products      []OrderProducts `json:"products" form:"products" binding:"required"`
}

// UpdateOrderStatusForm 更新订单状态表单
type UpdateOrderStatusForm struct {
	ID    int32  `form:"id" json:"id" binding:"required"`
	State string `json:"state" form:"state"`
}

// DetailsInfo 详情信息
type DetailsInfo struct {
	Order
	AddressInfo address.Address `json:"addressInfo" form:"addressInfo"`
}
