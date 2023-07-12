package orderServices

import (
	"errors"
	"gorm.io/gorm"
	"simple-mall/global"
	"simple-mall/models/address"
	"simple-mall/models/order"
	"simple-mall/models/shoppingCart"
	"simple-mall/utils"
)

// Create 创建订单
func Create(info order.SaveForm) (int32, error) {

	saveInfo := order.Order{
		UserId:        info.UserId,
		TotalPrice:    info.TotalPrice,
		Remark:        info.Remark,
		AddressId:     info.AddressId,
		PaymentMethod: info.PaymentMethod,
		Products:      info.Products,
	}

	err := global.DB.Transaction(func(tx *gorm.DB) error {

		productIDs := make([]int32, len(info.Products))
		for i, v := range info.Products {
			productIDs[i] = v.ProductId
		}

		// 生成订单时移除用户购物车商品
		if err := tx.Where("user_id = ? AND product_id in ?", info.UserId, productIDs).Delete(&shoppingCart.ShoppingCart{}).Error; err != nil {
			return err
		}

		// 创建订单
		saveInfo.State = "00"
		if err := tx.Model(&order.Order{}).Create(&saveInfo).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

	if err != nil {
		return 0, err
	}

	return info.ID, nil
}

// UpdateOrderStatus 修改订单状态
func UpdateOrderStatus(userId int32, id int32, state string) error {
	db := global.DB.Model(&order.Order{}).Where("id = ? AND user_id = ?", id, userId).Update("state", state)

	if db.RowsAffected == 0 {
		return errors.New("需要更新的数据不存在")
	}

	return db.Error
}

// GetUserOrderList 获取用户订单列表
func GetUserOrderList(userId int32, listForm order.ListForm) ([]order.Order, int64, error) {
	var list []order.Order

	query := global.DB.Model(&order.Order{}).Where("user_id = ?", userId)

	if listForm.State != "" {
		query = query.Where("state = ?", listForm.State)
	}

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	if err := query.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Preload("Products").Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// GetOrderList 获取订单列表
func GetOrderList(listForm order.ListForm) ([]order.Order, int64, error) {
	var list []order.Order

	query := global.DB.Model(&order.Order{})
	if listForm.State != "" {
		query = query.Where("state = ?", listForm.State)
	}

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	if err := query.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Preload("Products").Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// Details 获取订单详情
func Details(id string) (order.DetailsInfo, error) {
	orderInfo := &order.Order{}
	if err := global.DB.Preload("Products").First(orderInfo, "id = ?", id).Error; err != nil {
		return order.DetailsInfo{}, err
	}

	addressInfo := &address.Address{}
	if err := global.DB.First(addressInfo, "id = ?", orderInfo.AddressId).Error; err != nil {
		return order.DetailsInfo{}, err
	}

	detailsInfo := order.DetailsInfo{
		Order:       *orderInfo,
		AddressInfo: *addressInfo,
	}

	return detailsInfo, nil
}
