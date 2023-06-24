package slideshow

import "simple-mall/models"

// Slideshow 轮播图
type Slideshow struct {
	models.BaseModel
	ImageURL    string `gorm:"type:varbinary(300); not null; comment '图片地址'" json:"imageURL"`
	JumpLink    string `gorm:"type:varbinary(300); comment '跳转链接'" json:"jumpLink"`
	Description string `gorm:"type:varbinary(200); comment '图片描述'" json:"description"`
	Type        string `gorm:"type:varbinary(10); default:01; comment '图片类型 01 登录页 02 首页'" json:"type"`
}

// SaveForm 轮播图保存
type SaveForm struct {
	ID          int32  `form:"id" json:"id"`
	ImageURL    string `form:"imageURL" json:"imageURL" binding:"required"`
	JumpLink    string `form:"jumpLink" json:"jumpLink"`
	Type        string `form:"type" json:"type" binding:"required"`
	Description string `form:"description" json:"description"`
}

// ListForm 轮播图分页参数
type ListForm struct {
	models.PaginationParameters
	Description string `form:"description" json:"description"`
	Type        string `form:"type" json:"type"`
}
