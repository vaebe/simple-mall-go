package shoppingCartServices

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"simple-mall/global"
	"simple-mall/models/shoppingCart"
	"simple-mall/services/productServices"
	"strconv"
)

// CreateAndUpdate 创建或更新购物车
func CreateAndUpdate(saveForm shoppingCart.SaveForm) (int32, error) {
	saveInfo := shoppingCart.ShoppingCart{
		UserId:    saveForm.UserId,
		ProductId: saveForm.ProductId,
		Count:     saveForm.Count,
	}

	var oldInfo shoppingCart.ShoppingCart
	db := global.DB.Model(&shoppingCart.ShoppingCart{}).Where("user_id = ? AND product_id = ?", saveForm.UserId, saveInfo.ProductId).First(&oldInfo)

	if db.Error != gorm.ErrRecordNotFound && db.Error != nil {
		return 0, db.Error
	}

	if oldInfo.ID == 0 {
		db = global.DB.Create(&saveInfo)
	} else {
		// 如果数据存在则更新数量
		saveInfo.Count += oldInfo.Count
		db = db.Model(&shoppingCart.ShoppingCart{}).Where("id = ?", oldInfo.ID).Updates(&saveInfo)
	}

	if db.Error != nil {
		return 0, db.Error
	}

	return saveInfo.ID, nil
}

// DeleteShoppingCartProduct 删除购物车商品
func DeleteShoppingCartProduct(userId string, productId string) error {
	db := global.DB.Where("user_id = ? AND product_id = ?", userId, productId).Delete(&shoppingCart.ShoppingCart{})

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在")
	}

	return nil
}

// GetShoppingCartInfoByUserId 根据用户 id 获取购物车信息
func GetShoppingCartInfoByUserId(userId int32) ([]shoppingCart.Details, error) {

	var list []shoppingCart.Details
	rows, err := global.DB.Model(&shoppingCart.ShoppingCart{}).Where("user_id = ?", userId).Rows()

	if err != nil {
		return list, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var shoppingCartInfo shoppingCart.ShoppingCart
		err := global.DB.ScanRows(rows, &shoppingCartInfo)

		if err != nil {
			return nil, err
		}

		zap.S().Info(string(shoppingCartInfo.ProductId))

		productInfo, err := productServices.Details(strconv.Itoa(int(shoppingCartInfo.ProductId)))
		if err != nil {
			return nil, err
		}

		details := shoppingCart.Details{
			UserId:      shoppingCartInfo.UserId,
			ProductInfo: productInfo,
			ProductId:   shoppingCartInfo.ProductId,
			Count:       shoppingCartInfo.Count,
		}

		list = append(list, details)
	}

	return list, nil
}

// GetTheNumberOfItemsInTheShoppingCart 获取购物车商品数量
func GetTheNumberOfItemsInTheShoppingCart(userId int32) (int32, error) {
	var sum int32

	db := global.DB.Model(&shoppingCart.ShoppingCart{}).Where("user_id = ?", userId).Select("SUM(count)").Scan(&sum)

	return sum, db.Error
}
