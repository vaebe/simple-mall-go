package shoppingCart

import "simple-mall/models"

// ShoppingCart 购物车
type ShoppingCart struct {
	models.BaseModel
	UserId    int32 `gorm:"type:int; not null; comment '用户id'" json:"userId"`
	ProductId int32 `gorm:"type:int; not null; comment '商品id'" json:"productId"`
	Count     int32 `gorm:"type:int; not null; comment '商品数量'" json:"count"`
}

// SaveForm 购物车保存表单
type SaveForm struct {
	ID        int32 `form:"id" json:"id"`
	UserId    int32 `form:"userId" json:"userId" binding:"required"`
	ProductId int32 `form:"productId" json:"productId" binding:"required"`
	Count     int32 `form:"count" json:"count" binding:"required"`
}
