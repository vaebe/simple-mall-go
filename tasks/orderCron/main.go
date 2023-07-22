package orderCron

import (
	"go.uber.org/zap"
	"simple-mall/services/orderServices"
	"time"
)

// 定义执行的任务逻辑
func doTask() {
	err := orderServices.UpdateTimedOutUnpaidOrderStatus()
	if err != nil {
		zap.S().Error("更新超时订单错误：", err.Error())
	}
}

// UpdateTheOrderStatusToTimeoutNotPaid 更新订单状态为超时未支付
func UpdateTheOrderStatusToTimeoutNotPaid() {
	// 开始立即调用一次
	doTask()

	// 创建间隔为两分钟的定时器通道
	ticker := time.Tick(2 * time.Minute)

	go func() {
		for {
			select {
			case <-ticker:
				doTask()
			}
		}
	}()
}
