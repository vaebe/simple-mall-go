# simple-mall-go

+ simple-mall 接口服务

## 主要技术栈

+ go 、 gin 、mysql 、redis

## redis

+ [文档](https://redis.uptrace.dev/zh/guide/go-redis.html)

## swagger

+ 生成 `swag init`
+ 注释格式化 `swag fmt`

## 打包
+ `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build`

## docker 打包
+ 打包 `docker build . -t mall`
+ 测试 `docker run -p 127.0.0.1:8080:8080/tcp mall`

# 订单状态

+ 待支付（Pending Payment）：顾客下单后，订单状态为待支付，等待顾客完成支付操作。
+ 已支付（Paid）：顾客完成支付后，订单状态更新为已支付，表示订单支付成功。
+ 处理中（Processing）：订单支付成功后，商家开始处理订单，执行备货、打包等操作，订单状态更新为处理中。
+ 已发货（Shipped）：商家将商品交付给物流公司并更新订单状态为已发货，同时提供物流追踪信息供顾客查看。
+ 已完成（Completed）：顾客收到商品后，确认无误并满意，将订单状态更新为已完成。
+ 已取消（Cancelled）：在任何阶段，顾客或系统都可以取消订单，将订单状态更新为已取消。
+ 退款中（Refunding）：顾客发起退款请求后，商家开始退款流程，将订单状态更新为退款中。
+ 已退款（Refunded）：退款流程完成后，商家将订单状态更新为已退款，并完成退款操作。
+ 异常（Exception）：在购物流程中，如果出现异常情况，例如库存不足、商品损坏等，订单状态可能被设置为异常状态，并进行相应的处理和调查。