package enumServices

import (
	"errors"
	"simple-mall/global"
	"simple-mall/models/enum"
	"simple-mall/utils"
)

// CreateAndUpdate 创建或更新枚举
func CreateAndUpdate(saveForm enum.SaveForm) (int32, error) {
	saveInfo := enum.Enum{
		Code:     saveForm.Code,
		Name:     saveForm.Name,
		TypeCode: saveForm.TypeCode,
		TypeName: saveForm.TypeName,
	}

	// id 不存在新增
	db := global.DB
	if saveForm.ID == 0 {
		db = db.Create(&saveInfo)
	} else {
		db = db.Model(&enum.Enum{}).Where("id = ?", saveForm.ID).Updates(&saveInfo)
	}

	if db.Error != nil {
		return 0, db.Error
	}

	return saveInfo.ID, nil
}

// Delete 删除枚举
func Delete(enumsId string) error {
	db := global.DB.Where("id = ?", enumsId).Delete(&enum.Enum{})

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在")
	}

	return nil
}

// Details 获取枚举详情
func Details(enumsId string) (enum.SaveForm, error) {
	enumInfo := enum.SaveForm{}
	db := global.DB.Model(&enum.Enum{}).Where("id = ?", enumsId).First(&enumInfo)

	if db.Error != nil {
		return enumInfo, db.Error
	}

	return enumInfo, nil
}

// GetEnumsByType 根据分类查询枚举
func GetEnumsByType(typeCode string) ([]enum.SaveForm, error) {
	var enumsList []enum.SaveForm
	db := global.DB.Model(&enum.Enum{}).Where("type_code", typeCode).Find(&enumsList)

	if db.Error != nil {
		return enumsList, db.Error
	}
	return enumsList, nil
}

// GetAllEnums 获取全部数据
func GetAllEnums() (map[string][]enum.SaveForm, error) {
	var enumsList []enum.SaveForm
	db := global.DB.Model(&enum.Enum{}).Find(&enumsList)

	enums := make(map[string][]enum.SaveForm)

	if db.Error != nil {
		return enums, db.Error
	}

	for _, v := range enumsList {
		enums[v.TypeCode] = append(enums[v.TypeCode], v)
	}

	return enums, nil
}

// GetEnumsList 分页获取枚举列表
func GetEnumsList(listForm enum.ListForm) ([]enum.Enum, int32, error) {
	var enums []enum.Enum
	db := global.DB.Where("name LIKE ? AND type_name LIKE ?", "%"+listForm.Name+"%", "%"+listForm.TypeName+"%").Find(&enums)

	if db.Error != nil {
		return enums, 0, db.Error
	}

	// 获取总数
	total := int32(db.RowsAffected)

	// 分页
	db = db.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Find(&enums)

	if db.Error != nil {
		return enums, 0, db.Error
	}

	return enums, total, nil
}
