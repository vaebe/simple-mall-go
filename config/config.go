package config

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Expire   int    `mapstructure:"expire" json:"expire"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
}

type MysqlConfig struct {
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Name        string `mapstructure:"name" json:"name"`
	User        string `mapstructure:"user" json:"user"`
	Password    string `mapstructure:"password" json:"password"`
	AutoMigrate bool   `mapstructure:"autoMigrate" json:"autoMigrate"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"signingKey" json:"signingKey"`
}

type EmailConfig struct {
	Key string `mapstructure:"key" json:"key"`
}

type QiNiuConfig struct {
	Access  string `mapstructure:"access" json:"access"`
	Secret  string `mapstructure:"secret" json:"secret"`
	Bucket  string `mapstructure:"bucket" json:"bucket"`
	BaseUrl string `mapstructure:"baseUrl" json:"baseUrl"`
}
