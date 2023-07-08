package addressServices

import (
	"errors"
	"simple-mall/global"
	"simple-mall/models/address"
	"simple-mall/models/product"
	"simple-mall/utils"
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

// CreateAndUpdate 创建或更新地址
func CreateAndUpdate(formData address.SaveForm) (int32, error) {
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

	db := global.DB

	// 根据 ID 判断是创建还是更新
	if formData.ID == 0 {
		if err := db.Create(&saveInfo).Error; err != nil {
			return 0, err
		}
	} else {
		if err := db.Model(&address.Address{}).Where("id = ?", formData.ID).Updates(&saveInfo).Error; err != nil {
			return 0, err
		}
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

	db := global.DB.Model(&address.Address{}).Where("user_id = ?", userId).Find(&addressInfoList)
	return addressInfoList, db.Error
}
