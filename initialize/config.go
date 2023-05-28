package initialize

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"simple-mall/global"
)

// GetEnvInfo 获取设置的env变量, 变量设置完成需要重启ide
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

// 设置 config 数据
func setConfig() {
	//debug := GetEnvInfo("MK_DEBUG")

	// 配置文件路径
	configFileName := "./config.yaml"

	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// 设置环境

	if err := v.UnmarshalKey("env", &global.ENV); err != nil {
		panic(err)
	}
	zap.S().Infof("env配置信息: %v", global.ENV)

	// mysqlConfig - 全局变量
	if err := v.UnmarshalKey("mysqlConfig", &global.MysqlConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("MysqlConfig配置信息: %v", global.MysqlConfig)

	// redisConfig - 全局变量
	if err := v.UnmarshalKey("redisConfig", &global.RedisConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("redisConfig配置信息: %v", global.RedisConfig)

	// jwt 全局变量
	if err := v.UnmarshalKey("JWTConfig", &global.JWTConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("JWTConfig配置信息: %v", global.JWTConfig)

	// emailServices 邮箱
	if err := v.UnmarshalKey("emailConfig", &global.EmailConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("emailConfig配置信息: %v", global.EmailConfig)

	// 七牛云存储
	if err := v.UnmarshalKey("qiNiuConfig", &global.QiNiuConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("七牛云存储配置信息: %v", global.QiNiuConfig)
}

// InitConfig 初始化config配置
func InitConfig() {
	// 初始化日志
	InitLogger()

	// 设置配置信息
	setConfig()

	// 初始化mysql
	InitMysql()

	// 初始化redis
	InitRedis()

	// 初始化表单验证翻译
	err := InitTrans("zh")
	if err != nil {
		zap.S().Error("InitTrans:", err.Error())
	}

	// 初始化自定义表单验证规则
	CustomValidators()
}
