package shoppingCartServices

import (
	"errors"
	"simple-mall/global"
	"simple-mall/models/shoppingCart"
)

// CreateAndUpdate 创建或更新购物车
func CreateAndUpdate(saveForm shoppingCart.SaveForm) (int32, error) {
	saveInfo := shoppingCart.SaveForm{
		UserId:    saveForm.UserId,
		ProductId: saveForm.ProductId,
		Count:     saveForm.Count,
	}

	// id 不存在新增
	db := global.DB
	if saveForm.ID == 0 {
		db = db.Create(&saveInfo)
	} else {
		db = db.Model(&shoppingCart.ShoppingCart{}).Where("id = ?", saveForm.ID).Updates(&saveInfo)
	}

	if db.Error != nil {
		return 0, db.Error
	}

	return saveInfo.ID, nil
}

// Delete 删除购物车
func Delete(shoppingCartId string) error {
	db := global.DB.Where("id = ?", shoppingCartId).Delete(&shoppingCart.ShoppingCart{})

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在")
	}

	return nil
}

// GetShoppingCartInfoByUserId 根据用户 id 获取购物车信息
func GetShoppingCartInfoByUserId(userId int32) ([]shoppingCart.ShoppingCart, error) {
	var list []shoppingCart.ShoppingCart
	db := global.DB.Model(&shoppingCart.ShoppingCart{}).Where("user_id = ?", userId).Find(&list)

	return list, db.Error
}
