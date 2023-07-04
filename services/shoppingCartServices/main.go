package shoppingCartServices

import (
	"errors"
	"gorm.io/gorm"
	"simple-mall/global"
	"simple-mall/models/shoppingCart"
	"simple-mall/services/productServices"
)

// getProductInfo 获取商品信息
func getProductInfo(userId int32, productId int32) (shoppingCart.ShoppingCart, error) {
	var productInfo shoppingCart.ShoppingCart
	db := global.DB.Model(&shoppingCart.ShoppingCart{}).Where("user_id = ? AND product_id = ?", userId, productId).First(&productInfo)

	// 错误不等于空 且错误不是查不到数据
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		return productInfo, db.Error
	}

	return productInfo, nil
}

// AddProductToShoppingCart 添加商品到购物车
func AddProductToShoppingCart(info shoppingCart.SaveForm) (int32, error) {
	saveInfo := shoppingCart.ShoppingCart{
		UserId:    info.UserId,
		ProductId: info.ProductId,
		Count:     info.Count,
		Selected:  info.Selected,
	}

	oldInfo, err := getProductInfo(saveInfo.UserId, saveInfo.ProductId)
	if err != nil {
		return 0, err
	}

	db := global.DB
	if oldInfo.ID == 0 {
		db = db.Create(&saveInfo)
	} else {
		// 如果数据存在则更新数量
		saveInfo.Count += oldInfo.Count
		db = db.Model(&shoppingCart.ShoppingCart{}).Where("id = ?", oldInfo.ID).Updates(&saveInfo)
	}

	return saveInfo.ID, db.Error
}

// BatchUpdateShoppingCartProductInfo 批量更新购物车商品信息
func BatchUpdateShoppingCartProductInfo(infoList []shoppingCart.SaveForm) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		for _, info := range infoList {
			editInfo := shoppingCart.ShoppingCart{
				Count:    info.Count,
				Selected: info.Selected,
			}
			if res := tx.Model(&shoppingCart.ShoppingCart{}).Where("id = ?", info.ID).Select("count", "selected").Updates(&editInfo); res.Error != nil {
				return res.Error
			}
		}
		return nil
	})
}

// DeleteShoppingCartProduct 删除购物车商品
func DeleteShoppingCartProduct(userId int32, productId string) error {
	db := global.DB.Where("user_id = ? AND product_id = ?", userId, productId).Delete(&shoppingCart.ShoppingCart{})

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在")
	}

	return nil
}

// GetShoppingCartInfoByUserId 根据用户 id 获取购物车信息
func GetShoppingCartInfoByUserId(userId int32) ([]shoppingCart.Details, error) {
	var list []shoppingCart.Details

	// 查询购物车信息
	db := global.DB.Model(&shoppingCart.ShoppingCart{}).Where("user_id = ?", userId).Find(&list)
	if db.Error != nil {
		return nil, db.Error
	}

	// 获取所有商品ID
	productIDs := make([]int32, len(list))
	for i, v := range list {
		productIDs[i] = v.ProductId
	}

	// 批量获取商品信息
	_, productInfoObj, err := productServices.GetProductInfoInBulk(productIDs)
	if err != nil {
		return nil, err
	}

	// 更新购物车信息中的商品信息
	for i := range list {
		productInfo := productInfoObj[list[i].ProductId]
		list[i].ProductInfo = productInfo
	}

	return list, nil
}

// GetTheNumberOfItemsInTheShoppingCart 获取购物车商品数量
func GetTheNumberOfItemsInTheShoppingCart(userId int32) (int32, error) {
	var sum int32

	db := global.DB.Model(&shoppingCart.ShoppingCart{}).Where("user_id = ?", userId).Select("SUM(count)").Scan(&sum)

	return sum, db.Error
}
