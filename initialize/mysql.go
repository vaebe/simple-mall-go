package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"simple-mall/global"
	"simple-mall/models/enum"
	"simple-mall/models/product"
	"simple-mall/models/productCategory"
	"simple-mall/models/role"
	"simple-mall/models/user"
	"time"
)

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.MysqlConfig.User,
		global.MysqlConfig.Password,
		global.MysqlConfig.Host,
		global.MysqlConfig.Port,
		global.MysqlConfig.Name)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	if global.MysqlConfig.AutoMigrate {
		err = global.DB.AutoMigrate(
			&user.User{},
			&enum.Enum{},
			&role.Role{},
			&product.Product{},
			&productCategory.ProductCategory{},
		)
		// 自动建表
		if err != nil {
			panic(err)
		}
	}
}
