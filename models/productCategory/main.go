package productCategory

// ProductCategory 商品分类表
type ProductCategory struct {
	Code string `gorm:"type:varbinary(50); not null; comment '商品分类code'" json:"code"`
	Name string `gorm:"type:varbinary(50); not null; comment '商品分类名称'" json:"name"`
	Sort int32  `gorm:"type:int;default:0;unique;comment '排序'" json:"sort"`
}
