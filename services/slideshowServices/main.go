package slideshowServices

import (
	"errors"
	"simple-mall/global"
	"simple-mall/models/slideshow"
	"simple-mall/utils"
)

// CreateAndUpdate 创建或更新轮播图
func CreateAndUpdate(saveForm slideshow.SaveForm) (int32, error) {
	saveInfo := slideshow.Slideshow{
		ImageURL:    saveForm.ImageURL,
		Description: saveForm.Description,
		JumpLink:    saveForm.JumpLink,
		Type:        saveForm.Type,
	}

	// id 不存在新增
	db := global.DB
	if saveForm.ID == 0 {
		db = db.Create(&saveInfo)
	} else {
		db = db.Model(&slideshow.Slideshow{}).Where("id = ?", saveForm.ID).Updates(&saveInfo)
	}

	if db.Error != nil {
		return 0, db.Error
	}

	return saveInfo.ID, nil
}

// Delete 删除轮播图
func Delete(id string) error {
	db := global.DB.Where("id = ?", id).Delete(&slideshow.Slideshow{})

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在")
	}

	return nil
}

// Details 获取轮播图详情
func Details(id string) (slideshow.SaveForm, error) {
	slideshowInfo := slideshow.SaveForm{}
	db := global.DB.Model(&slideshow.Slideshow{}).Where("id = ?", id).First(&slideshowInfo)

	if db.Error != nil {
		return slideshowInfo, db.Error
	}

	return slideshowInfo, nil
}

// GetSlideshowsByType 根据分类查询轮播图
func GetSlideshowsByType(code string) ([]slideshow.SaveForm, error) {
	var list []slideshow.SaveForm
	db := global.DB.Model(&slideshow.Slideshow{}).Where("type", code).Find(&list)
	return list, db.Error
}

// GetSlideshowsList 分页获取轮播图列表
func GetSlideshowsList(listForm slideshow.ListForm) ([]slideshow.Slideshow, int32, error) {
	var slideshows []slideshow.Slideshow

	db := global.DB
	if listForm.Description != "" {
		db = db.Where("description LIKE ?", "%"+listForm.Description+"%")
	}

	if listForm.Type != "" {
		db = db.Where("type = ?", listForm.Type)
	}

	db = db.Find(&slideshows)

	if db.Error != nil {
		return slideshows, 0, db.Error
	}

	// 获取总数
	total := int32(db.RowsAffected)

	// 分页
	db = db.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Find(&slideshows)

	if db.Error != nil {
		return slideshows, 0, db.Error
	}

	return slideshows, total, nil
}
