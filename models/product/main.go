package product

import "simple-mall/models"

// Product 商品表
type Product struct {
	models.BaseModel
	Name              string `gorm:"type:varbinary(200); not null; comment '商品名称'" json:"name"`
	Price             int32  `gorm:"type:int; default:0; comment '商品价格'" json:"price"`
	Picture           string `gorm:"type:varbinary(300); not null; comment '商品图片'" json:"Picture"`
	Stock             int32  `gorm:"type:int; default:1; comment '商品库存'" json:"stock"`
	Info              string `gorm:"type:varbinary(300); not null; comment '商品简介'" json:"info"`
	ProductCategoryId int32  `gorm:"type:int; not null; comment '商品分类id'" json:"productCategoryId"`
}

// SaveForm 商品分类表保存表单
type SaveForm struct {
	ID                int32  `form:"id" json:"id"`
	Name              string `form:"name" json:"name" binding:"required"`
	Price             int32  `form:"price" json:"price" binding:"required"`
	Picture           string `form:"picture" json:"picture" binding:"required"`
	Stock             int32  `form:"stock" json:"stock" binding:"required"`
	Info              string `form:"info" json:"info" binding:"required"`
	ProductCategoryId int32  `form:"productCategoryId" json:"productCategoryId" binding:"required"`
}

// ListForm 商品分类表分页查询参数
type ListForm struct {
	models.PaginationParameters
	Name string `json:"name" form:"name"`
}
