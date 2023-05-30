package role

import "simple-mall/models"

// Role 角色表
type Role struct {
	models.BaseModel
	RoleName    string `gorm:"type:varbinary(50); not null; comment '角色名称'" json:"roleName"`
	Description string `gorm:"type:varbinary(300); not null; comment '角色描述'" json:"description"`
}
