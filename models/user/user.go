package user

import (
	"simple-mall/models"
)

type User struct {
	models.BaseModel
	NickName    string `gorm:"type:varbinary(40);unique;not null;comment '昵称'" json:"nickName"`
	UserAccount string `gorm:"type:varbinary(50);unique;not null;comment '用户账号'" json:"userAccount"`
	Password    string `gorm:"type:varbinary(300);not null;comment '密码'" json:"password"`
	PhoneNumber string `gorm:"type:varbinary(15); comment '手机号'" json:"phoneNumber"`
	Gender      string `gorm:"type:varbinary(15);default:02;comment '性别 00女 01男 02未知'" json:"gender"`
	Avatar      string `gorm:"type:varbinary(300);not null;comment '用户头像'" json:"avatar"`
	Role        string `gorm:"type:varbinary(15);default:02;comment '角色表定义 00 管理员 02 普通用户'" json:"role"`
}

// VerificationCodeForm 发送验证码
type VerificationCodeForm struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

// SaveForm 用户信息保存表单
type SaveForm struct {
	ID          int32  `json:"id" form:"id"`
	NickName    string `json:"nickName" form:"nickName" binding:"required,min=4,max=40"`
	UserAccount string `json:"userAccount" form:"userAccount"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Gender      string `json:"gender" form:"gender"`
	Avatar      string `json:"avatar" form:"avatar"`
	Role        string `json:"role" form:"role"`
}

// RegisterForm 注册
type RegisterForm struct {
	UserAccount string `form:"userAccount" json:"userAccount" binding:"required,email"`
	PassWord    string `form:"password" json:"password" binding:"required,min=6,max=20"`
	Code        string `form:"code" json:"code" binding:"required,min=6,max=6"`
}

// LoginForm 登陆
type LoginForm struct {
	UserAccount string `form:"userAccount" json:"userAccount" binding:"required,email" example:"admin@163.com"`
	PassWord    string `form:"password" json:"password" binding:"required,min=6,max=300" example:"123456"`
}

// ListForm 获取用户列表查询参数
type ListForm struct {
	models.PaginationParameters
	UserAccount string `json:"userAccount" form:"userAccount"`
	NickName    string `json:"nickName" form:"nickName"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
}

// LoginResultsData 登录返回数据
type LoginResultsData struct {
	UserInfo  User   `json:"userInfo"`
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}
