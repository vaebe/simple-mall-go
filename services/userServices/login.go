package userServices

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"simple-mall/global"
	"simple-mall/models/user"
	"simple-mall/utils/password"
)

// Register 用户注册
func Register(registerForm user.RegisterForm) (user.User, error) {
	// 验证短信验证码
	redisCode := global.RedisClient.Get(context.Background(), registerForm.UserAccount)
	verificationCode, _ := redisCode.Result()

	if verificationCode == "redis" || verificationCode != registerForm.Code {
		zap.S().Infof("验证码不正确:应为%s实际为%s", verificationCode, registerForm.Code)

		return user.User{}, errors.New("验证码不正确")
	}

	// 生成不带 - 的uuid
	uuidObj := uuid.New()
	uuidStr := fmt.Sprintf("mall%x", uuidObj[:])

	pwdStr, err := password.EncryptByAes([]byte(registerForm.PassWord))

	if err != nil {
		return user.User{}, err
	}

	userInfo := user.User{
		NickName:    uuidStr,
		Avatar:      "https://cdn.qiniu.vaebe.top/simple-mall/avatar_default.png",
		UserAccount: registerForm.UserAccount,
		Password:    pwdStr,
		Gender:      "02",
	}
	res := global.DB.Create(&userInfo)

	if res.Error != nil {
		return user.User{}, res.Error
	}

	return userInfo, nil
}

// VerifyUserPassword 校验用户密码
func VerifyUserPassword(loginForm user.LoginForm) (user.User, error) {
	userInfo := user.User{}

	res := global.DB.Where("user_account = ?", loginForm.UserAccount).First(&userInfo)

	if res.Error != nil {
		return user.User{}, res.Error
	}

	if userInfo.Password != loginForm.PassWord {
		return user.User{}, errors.New("密码不正确")
	}

	userInfo.Password = ""

	return userInfo, nil
}
