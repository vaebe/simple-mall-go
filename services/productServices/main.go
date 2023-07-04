package productServices

import (
	"errors"
	"gorm.io/gorm"
	"simple-mall/global"
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
		DetailImages:      saveForm.DetailImages,
		ParameterImages:   saveForm.ParameterImages,
		ProductCategoryId: saveForm.ProductCategoryId,
	}

	// id 不存在新增
	if saveForm.ID == 0 {
		db := global.DB.Save(&saveInfo)
		return saveInfo.ID, db.Error
	} else {
		saveInfo.ID = saveForm.ID

		err := global.DB.Transaction(func(tx *gorm.DB) error {
			// 删除旧的关联数据，这里只是清除了关联关系
			// 如果旧数据不需要留存可以直接硬删除，下边添加新的关联数据的逻辑也不再需要，Updates 会自动增加关联数据
			if err := tx.Model(&saveInfo).Association("Pictures").Clear(); err != nil {
				return err
			}

			// 添加新的关联数据
			if len(saveForm.Pictures) > 0 {
				if err := tx.Model(&saveInfo).Association("Pictures").Append(saveForm.Pictures); err != nil {
					return err
				}
			}

			// 更新信息
			if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&saveInfo).Error; err != nil {
				return err
			}

			return nil
		})

		return saveInfo.ID, err
	}
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
	db := global.DB.Model(&product.Product{}).Where("id = ?", productId).Preload("Pictures").First(&info)
	return info, db.Error
}

// GetProductInfoInBulk 批量获取商品信息
func GetProductInfoInBulk(ids []int32) ([]product.Product, map[int32]product.Product, error) {
	var productList []product.Product
	db := global.DB.Model(&product.Product{}).Where("id IN ?", ids).Preload("Pictures").Find(&productList)

	if db.Error != nil {
		return nil, nil, db.Error
	}

	productInfoObj := make(map[int32]product.Product, len(productList))

	for _, v := range productList {
		v.DetailImages = ""
		v.ParameterImages = ""
		productInfoObj[v.ID] = v
	}

	return productList, productInfoObj, db.Error
}

// GetProductList 分页获取商品列表
func GetProductList(listForm product.ListForm) ([]product.Product, int32, error) {
	var list []product.Product

	db := global.DB
	if listForm.Name != "" {
		db = db.Where("name LIKE ?", "%"+listForm.Name+"%")
	}

	if listForm.ProductCategoryId != 0 {
		db = db.Where("product_category_id = ?", listForm.ProductCategoryId)
	}

	db = db.Find(&list)

	if db.Error != nil {
		return list, 0, db.Error
	}

	// 获取总数
	total := int32(db.RowsAffected)

	// 分页
	db = db.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Preload("Pictures").Find(&list)

	if db.Error != nil {
		return list, 0, db.Error
	}

	return list, total, nil
}

// GetRandomRecommendedProductList 获取随机推荐商品列表
func GetRandomRecommendedProductList(total int) ([]product.Product, error) {
	var list []product.Product
	db := global.DB.Order("RAND()").Limit(total).Preload("Pictures").Find(&list)
	return list, db.Error
}
