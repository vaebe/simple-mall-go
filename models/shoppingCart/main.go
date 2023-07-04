package shoppingCart

import (
	"simple-mall/models"
	"simple-mall/models/product"
)

// ShoppingCart 购物车
type ShoppingCart struct {
	models.BaseModel
	UserId    int32 `gorm:"type:int; not null; comment '用户id'" json:"userId"`
	ProductId int32 `gorm:"type:int; not null; comment '商品id'" json:"productId"`
	Count     int32 `gorm:"type:int; not null; comment '商品数量'" json:"count"`
	Selected  bool  `gorm:"type:tinyint(1); not null; comment '选中状态'" json:"selected"`
}

// SaveForm 购物车保存表单
type SaveForm struct {
	ID        int32 `form:"id" json:"id"`
	UserId    int32 `form:"userId" json:"userId" binding:"required"`
	Selected  bool  `form:"selected" json:"selected"`
	ProductId int32 `form:"productId" json:"productId" binding:"required"`
	Count     int32 `form:"count" json:"count" binding:"required"`
}

// Details 购物车详情
type Details struct {
	ID          int32           `form:"id" json:"id"`
	UserId      int32           `form:"userId" json:"userId"`
	ProductId   int32           `form:"productId" json:"productId"`
	ProductInfo product.Product `form:"productInfo" json:"productInfo"`
	Count       int32           `form:"count" json:"count"`
	Selected    bool            `form:"selected" json:"selected"`
}
