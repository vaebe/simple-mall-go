package address

import "simple-mall/models"

// Address 地址管理表
type Address struct {
	models.BaseModel
	Name     string `gorm:"type:varbinary(100); not null; comment '枚举名称'" json:"name"`
	Code     string `gorm:"type:varbinary(100); not null; comment '枚举值'" json:"code"`
	TypeName string `gorm:"type:varbinary(100); not null; comment '枚举分类名称'" json:"typeName"`
	TypeCode string `gorm:"type:varbinary(100); not null; comment '枚举分类值'" json:"typeCode"`
	ParentId string `gorm:"type:varbinary(100); comment '枚举上级id'" json:"parentId"`
	IsDel    int32  `gorm:"type:int; comment '是否可删除 1可以 0 不可以'" json:"isDel"`
}

// AreaInfo 区域信息
type AreaInfo struct {
	ID           int     `json:"id"`
	PID          int     `json:"pid"`
	Deep         int     `json:"deep"`
	Name         string  `json:"name"`
	PinyinPrefix string  `json:"pinyinPrefix"`
	Pinyin       string  `json:"pinyin"`
	ExtID        float64 `json:"extID"`
	ExtName      string  `json:"ExtName"`
}
