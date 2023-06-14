package productCategory

import "simple-mall/models"

// ProductCategory 商品分类表
type ProductCategory struct {
	models.BaseModel
	Name string `gorm:"type:varbinary(50); not null; comment '商品分类名称'" json:"name"`
	Code string `gorm:"type:varbinary(50); not null; comment '商品分类code'" json:"code"`
	Sort int32  `gorm:"type:int;default:1;unique;comment '排序'" json:"sort"`
	Icon string `gorm:"type:varbinary(300); comment '商品分类icon'" json:"icon"`
}

// SaveForm 商品分类表保存表单
type SaveForm struct {
	ID   int32  `form:"id" json:"id"`
	Name string `form:"name" json:"name" binding:"required"`
	Code string `form:"code" json:"code" binding:"required"`
	Sort int32  `form:"sort" json:"sort" binding:"required"`
}

// ListForm 商品分类表分页查询参数
type ListForm struct {
	models.PaginationParameters
	Name string `json:"name" form:"name"`
}
