package product

// Product 商品表
type Product struct {
	Name              string `gorm:"type:varbinary(200); not null; comment '商品名称'" json:"name"`
	Price             int32  `gorm:"type:int; default:0; comment '商品价格'" json:"price"`
	Picture           string `gorm:"type:varbinary(300); not null; comment '商品图片'" json:"Picture"`
	Stock             int32  `gorm:"type:int; default:1; comment '商品库存'" json:"stock"`
	Info              string `gorm:"type:varbinary(300); not null; comment '商品简介'" json:"info"`
	ProductCategoryId int32  `gorm:"type:int; not null; comment '商品分类id'" json:"productCategoryId"`
}
