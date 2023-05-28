package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"simple-mall/config"
)

var (
	ENV         string
	Trans       ut.Translator
	DB          *gorm.DB
	RedisClient *redis.Client
	JWTConfig   *config.JWTConfig
	MysqlConfig *config.MysqlConfig
	RedisConfig *config.RedisConfig
	EmailConfig *config.EmailConfig
	QiNiuConfig *config.QiNiuConfig
	Logger      *zap.Logger
)
