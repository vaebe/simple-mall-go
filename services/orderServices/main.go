package orderServices

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"simple-mall/global"
	"simple-mall/models/address"
	"simple-mall/models/order"
	"simple-mall/models/shoppingCart"
	"simple-mall/utils"
	"time"
)

// Create 创建订单 todo 直接购买无需清空购物车商品
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

	return saveInfo.ID, nil
}

// UpdateOrderStatus 修改订单状态
func UpdateOrderStatus(id int32, state string) error {
	db := global.DB.Model(&order.Order{}).Where("id = ?", id).Update("state", state)

	if db.RowsAffected == 0 {
		return errors.New("需要更新的数据不存在")
	}

	return db.Error
}

// BulkUpdateOrderStatus 批量修改订单状态
func BulkUpdateOrderStatus(ids []int32, state string) error {
	db := global.DB.Model(&order.Order{}).Where("id IN ?", ids).Update("state", state)

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
	if err := query.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Preload("Products").Order("created_at desc").Find(&list).Error; err != nil {
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
	if err := query.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Preload("Products").Order("created_at desc").Find(&list).Error; err != nil {
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

// Delete 删除订单
func Delete(userid int32, orderId string, isAdmin bool) error {
	db := global.DB.Where("id = ?", orderId)

	// 非管理员用户删除订单需要验证是否是自己的订单
	if !isAdmin {
		db = db.Where(" user_id = ?", userid)
	}

	db.Delete(&order.Order{})

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在")
	}

	return nil
}

// UpdateTimedOutUnpaidOrderStatus 更新超时未支付订单状态
func UpdateTimedOutUnpaidOrderStatus() error {
	// 获取当前时间
	now := time.Now()

	// 计算半个小时前的时间
	halfHourAgo := now.Add(-30 * time.Minute)

	// 查询超时订单
	var timeoutOrders []order.Order
	if err := global.DB.Where("state = ? AND created_at <= ?", "00", halfHourAgo).Find(&timeoutOrders).Error; err != nil {
		zap.S().Error(err.Error())
		return err
	}

	orderIds := make([]int32, len(timeoutOrders))
	for i, v := range timeoutOrders {
		orderIds[i] = v.ID
	}

	zap.S().Debug("本次更新超时未支付订单ids：", orderIds)

	if len(orderIds) == 0 {
		return nil
	}

	err := BulkUpdateOrderStatus(orderIds, "09")
	if err != nil {
		zap.S().Error(err.Error())
		return err
	}
	return nil
}
