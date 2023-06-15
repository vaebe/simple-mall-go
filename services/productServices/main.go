package productServices

import (
	"errors"
	"simple-mall/global"
	"simple-mall/models/enum"
	"simple-mall/models/product"
	"simple-mall/utils"
)

// CreateAndUpdate 创建或更新商品
func CreateAndUpdate(saveForm product.SaveForm) (int32, error) {
	saveInfo := product.Product{
		Name:              saveForm.Name,
		Price:             saveForm.Price,
		Info:              saveForm.Info,
		Stock:             saveForm.Stock,
		Pictures:          saveForm.Pictures,
		ProductCategoryId: saveForm.ProductCategoryId,
	}

	// id 不存在新增
	db := global.DB
	if saveForm.ID == 0 {
		db = db.Create(&saveInfo)
	} else {
		db = db.Model(&product.Product{}).Where("id = ?", saveForm.ID).Updates(&saveInfo)
	}

	if db.Error != nil {
		return 0, db.Error
	}

	return saveInfo.ID, nil
}

// Delete 删除商品
func Delete(productId string) error {
	db := global.DB.Where("id = ?", productId).Delete(&product.Product{})

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在")
	}

	return nil
}

// Details 获取商品详情
func Details(productId string) (product.Product, error) {
	info := product.Product{}
	db := global.DB.Model(&enum.Enum{}).Preload("Pictures").Where("id = ?", productId).First(&info)
	return info, db.Error
}

// GetProductList 分页获取商品列表
func GetProductList(listForm product.ListForm) ([]product.Product, int32, error) {
	var list []product.Product

	db := global.DB.Preload("Pictures")
	if listForm.Name != "" {
		db = db.Where("name LIKE ?", "%"+listForm.Name+"%")
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
