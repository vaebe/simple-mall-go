package addressServices

import (
	"simple-mall/global"
	"simple-mall/models/address"
)

// GetAreasByParams 根据参数获取区域数据
func GetAreasByParams(id string, pid string, deep string) ([]address.AreaInfo, error) {

	var areas []address.AreaInfo

	db := global.DB.Table("areas")

	if id != "" {
		db = db.Where("id = ?", id)
	}

	if pid != "" {
		db = db.Where("pid = ?", pid)
	}

	if deep != "" {
		db = db.Where("deep = ?", deep)
	}

	// 参数都为空查询层级为 0 的数据
	if deep == "" && id == "" && pid == "" {
		db = db.Where("deep = ?", 0)
	}

	db = db.Find(&areas)

	return areas, db.Error
}
