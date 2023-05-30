package enum

import "simple-mall/models"

// Enum 枚举表
type Enum struct {
	models.BaseModel
	Name     string `gorm:"type:varbinary(100); not null; comment '枚举名称'" json:"name"`
	Code     string `gorm:"type:varbinary(100); not null; comment '枚举值'" json:"code"`
	TypeName string `gorm:"type:varbinary(100); not null; comment '枚举分类名称'" json:"typeName"`
	TypeCode string `gorm:"type:varbinary(100); not null; comment '枚举分类值'" json:"typeCode"`
	ParentId string `gorm:"type:varbinary(100); comment '枚举上级id'" json:"parentId"`
	IsDel    int32  `gorm:"type:int; comment '是否可删除 1可以 0 不可以'" json:"isDel"`
}

// SaveForm 枚举保存表单
type SaveForm struct {
	ID       int32  `form:"id" json:"id"`
	Name     string `form:"name" json:"name" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	TypeName string `form:"typeName" json:"typeName" binding:"required"`
	TypeCode string `form:"typeCode" json:"typeCode" binding:"required"`
	ParentId string `form:"parentId" json:"parentId"`
}

// ListForm 分页查询枚举参数
type ListForm struct {
	models.PaginationParameters
	Name     string `json:"name" form:"name"`
	TypeName string `json:"typeName" form:"typeName"`
}
