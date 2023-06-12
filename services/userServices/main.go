package userServices

import (
	"errors"
	"go.uber.org/zap"
	"simple-mall/global"
	"simple-mall/models/user"
	"simple-mall/utils"
)

// GetUserList 获取用户列表
func GetUserList(listForm user.ListForm) ([]user.User, int32, error) {
	var users []user.User
	db := global.DB
	if listForm.UserAccount != "" {
		db = db.Where("user_account LIKE ?", "%"+listForm.UserAccount+"%")
	}

	if listForm.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+listForm.NickName+"%")
	}

	if listForm.PhoneNumber != "" {
		db = db.Where("nick_name LIKE ?", "%"+listForm.PhoneNumber+"%")
	}

	res := db.Find(&users)

	// 存在错误
	if res.Error != nil {
		zap.S().Info(res.Error)
		return nil, 0, res.Error
	}

	// 获取总数
	total := int32(res.RowsAffected)

	// 分页
	res.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Find(&users)

	for i := range users {
		users[i].Password = ""
	}

	return users, total, nil
}

// GetUserDetails 获取用户详情
func GetUserDetails(userId string) (user.User, error) {
	userInfo := user.User{}
	res := global.DB.Where("id", userId).First(&userInfo)

	userInfo.Password = ""

	if res.Error != nil {
		return user.User{}, res.Error
	}

	return userInfo, nil
}

// CreateAndUpdate 创建或更新用户信息
func CreateAndUpdate(saveForm user.SaveForm) (int32, error) {
	saveInfo := user.User{
		NickName:    saveForm.NickName,
		UserAccount: saveForm.UserAccount,
		PhoneNumber: saveForm.PhoneNumber,
		Avatar:      saveForm.Avatar,
		Gender:      saveForm.Gender,
		Role:        saveForm.Role,
	}

	// id 不存在新增
	db := global.DB
	if saveForm.ID == 0 {
		saveInfo.Password = saveForm.Password
		db = db.Create(&saveInfo)
	} else {
		db = db.Model(&user.User{}).Where("id = ?", saveForm.ID).Updates(&saveInfo)
	}

	if db.Error != nil {
		return 0, db.Error
	}

	return saveInfo.ID, nil
}

// Delete 删除用户
func Delete(userId string) error {
	db := global.DB.Where("id = ?", userId).Delete(&user.User{})

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在")
	}

	return nil
}
