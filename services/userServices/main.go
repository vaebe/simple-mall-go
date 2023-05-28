package userServices

import (
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

// Edit 编辑用户信息
func Edit(editForm user.EditForm, userId int32) error {

	res := global.DB.Where("id = ?", userId).Updates(&user.User{
		NickName:    editForm.NickName,
		Avatar:      editForm.Avatar,
		PhoneNumber: editForm.PhoneNumber,
		Gender:      editForm.Gender,
	})

	if res.RowsAffected == 0 {
		return res.Error
	}

	return nil
}
