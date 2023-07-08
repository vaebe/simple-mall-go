package addressServices

import (
	"errors"
	"gorm.io/gorm"
	"simple-mall/global"
	"simple-mall/models/address"
	"simple-mall/models/product"
	"simple-mall/utils"
)

// updateNonDefaultAddresses 将用户的其他地址设为非默认地址
func updateNonDefaultAddresses(tx *gorm.DB, userId int32) error {
	return tx.Model(&address.Address{}).Where("user_id = ?", userId).Update("default_address", "00").Error
}

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

// CreateAndUpdate 创建或更新地址
func CreateAndUpdate(userId int32, formData address.SaveForm) (int32, error) {
	saveInfo := address.Address{
		UserId:          formData.UserId,
		Name:            formData.Name,
		Phone:           formData.Phone,
		Province:        formData.Province,
		ProvinceName:    formData.ProvinceName,
		City:            formData.City,
		CityName:        formData.CityName,
		District:        formData.District,
		DistrictName:    formData.DistrictName,
		Street:          formData.Street,
		StreetName:      formData.StreetName,
		DetailedAddress: formData.DetailedAddress,
		ZipCode:         formData.ZipCode,
		DefaultAddress:  formData.DefaultAddress,
	}

	var err error
	err = global.DB.Transaction(func(tx *gorm.DB) error {

		// 根据 id 判断是新增还是更新
		if formData.ID == 0 {

			// 创建数据是默认地址 将其他地址更改为非默认地址
			if saveInfo.DefaultAddress == "01" {
				if err := updateNonDefaultAddresses(tx, userId); err != nil {
					return err
				}
			}

			if err := tx.Create(&saveInfo).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Model(&address.Address{}).Where("id = ?", formData.ID).Updates(&saveInfo).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return saveInfo.ID, nil
}

// Delete 删除地址
func Delete(id string) error {
	db := global.DB.Where("id = ?", id).Delete(&address.Address{})

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在")
	}

	return nil
}

// Details 获取地址详情
func Details(id string) (address.Address, error) {
	info := address.Address{}
	db := global.DB.Model(&product.Product{}).Where("id = ?", id).First(&info)
	return info, db.Error
}

// GetAddressInfoList 获取地址信息列表
func GetAddressInfoList(listForm address.ListForm) ([]address.Address, int64, error) {
	var list []address.Address

	query := global.DB.Model(&address.Address{}).
		Where(utils.BuildMySQLLikeQueryCondition("province_name", listForm.ProvinceName)).
		Where(utils.BuildMySQLLikeQueryCondition("city_name", listForm.CityName)).
		Where(utils.BuildMySQLLikeQueryCondition("district_name", listForm.DistrictName)).
		Where(utils.BuildMySQLLikeQueryCondition("street_name", listForm.StreetName))

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	if err := query.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// GetUserAddressInfoList 获取用户的地址信息列表
func GetUserAddressInfoList(userId int32) ([]address.Address, error) {
	var addressInfoList []address.Address

	db := global.DB.Model(&address.Address{}).Where("user_id = ?", userId).Order("default_address desc").Find(&addressInfoList)
	return addressInfoList, db.Error
}

// SetDefaultAddress 设置默认地址
func SetDefaultAddress(userId int32, id string) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := updateNonDefaultAddresses(tx, userId); err != nil {
			return err
		}

		if err := tx.Model(&address.Address{}).Where("user_id = ? AND id = ?", userId, id).Update("default_address", "01").Error; err != nil {
			return err
		}

		return nil
	})
}
