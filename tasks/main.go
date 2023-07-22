package tasks

import "simple-mall/tasks/orderCron"

// InitTasks  初始化定时任务
func InitTasks() {
	orderCron.UpdateTheOrderStatusToTimeoutNotPaid()
}
