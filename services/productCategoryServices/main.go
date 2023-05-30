package productCategoryServices

import (
	"errors"
	"simple-mall/global"
	"simple-mall/models/productCategory"
	"simple-mall/utils"
)

// CreateAndUpdate 创建或更新商品分类
func CreateAndUpdate(saveForm productCategory.SaveForm) (int32, error) {
	saveInfo := productCategory.ProductCategory{
		Code: saveForm.Code,
		Name: saveForm.Name,
		Sort: saveForm.Sort,
	}

	// id 不存在新增
	db := global.DB
	if saveForm.ID == 0 {
		db = db.Create(&saveInfo)
	} else {
		db = db.Model(&productCategory.ProductCategory{}).Where("id = ?", saveForm.ID).Updates(&saveInfo)
	}

	if db.Error != nil {
		return 0, db.Error
	}

	return saveInfo.ID, nil
}

// Delete 删除商品分类
func Delete(productCategoryId string) error {
	db := global.DB.Where("id = ?", productCategoryId).Delete(&productCategory.ProductCategory{})

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在")
	}

	return nil
}

// GetAllProductCategory 获取全部数据
func GetAllProductCategory() ([]productCategory.ProductCategory, error) {
	var list []productCategory.ProductCategory
	db := global.DB.Find(&list)
	return list, db.Error
}

// GetProductCategoryList 分页获取商品分类列表
func GetProductCategoryList(listForm productCategory.ListForm) ([]productCategory.ProductCategory, int32, error) {
	var list []productCategory.ProductCategory

	db := global.DB
	if listForm.Name != "" {
		db = db.Where("name LIKE", "%"+listForm.Name+"%")
	}
	db = db.Find(&list)

	if db.Error != nil {
		return list, 0, db.Error
	}

	// 获取总数
	total := int32(db.RowsAffected)

	// 分页
	db = db.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Find(&list)

	if db.Error != nil {
		return list, 0, db.Error
	}

	return list, total, nil
}
