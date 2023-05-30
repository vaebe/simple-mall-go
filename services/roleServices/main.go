package roleServices

import (
	"simple-mall/global"
	"simple-mall/models/role"
)

// GetRoleList 获取角色列表
func GetRoleList() ([]role.Role, error) {
	var roles []role.Role
	db := global.DB.Find(&roles)

	if db.Error != nil {
		return roles, db.Error
	}
	return roles, db.Error
}
