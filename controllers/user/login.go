package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"simple-mall/global"
	middlewares "simple-mall/middleware"
	"simple-mall/models/user"
	"simple-mall/services/emailServices"
	"simple-mall/services/userServices"
	"simple-mall/utils"
	"time"
)

// GetVerificationCode
//
//	@Summary		获取验证码
//	@Description	获取验证码
//	@Tags			user用户
//	@Accept			json
//	@Produce		json
//	@Param			param	body		user.VerificationCodeForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Router			/user/getVerificationCode [post]
func GetVerificationCode(ctx *gin.Context) {
	//表单验证
	verificationCodeForm := user.VerificationCodeForm{}

	if err := ctx.ShouldBind(&verificationCodeForm); err != nil {
		zap.S().Info(&verificationCodeForm)
		utils.HandleValidatorError(ctx, err)
		return
	}

	// 获取随机验证码
	verificationCode := utils.GenerateSmsCode(6)

	// 发送验证码邮件
	err := emailServices.SendTheVerificationCodeEmail(verificationCode, verificationCodeForm.Email)
	if err != nil {
		utils.ResponseResultsError(ctx, "发送邮件验证码失败")
	}

	// 将数据存储到redis
	global.RedisClient.Set(context.Background(), verificationCodeForm.Email, verificationCode, time.Duration(global.RedisConfig.Expire)*time.Second)
	utils.ResponseResultsSuccess(ctx, "发送验证码成功！")
}

// loginSuccess 登陆成功后的操作
func loginSuccess(ctx *gin.Context, userInfo user.User) {
	token, err := middlewares.GenerateLoginToken(userInfo)
	if err != nil {
		zap.S().Info("生成token错误", err.Error())
		utils.ResponseResultsError(ctx, "生成token错误!")
		return
	}

	resultsData := user.LoginResultsData{
		UserInfo:  userInfo,
		Token:     token,
		ExpiredAt: (time.Now().Unix() + 60*60*24*30) * 1000,
	}

	resultsData.UserInfo.Password = ""

	utils.ResponseResultsSuccess(ctx, resultsData)
}

// Register
//
//	@Summary		用户注册
//	@Description	用户注册
//	@Tags			user用户
//	@Accept			json
//	@Produce		json
//	@Param			param	body		user.RegisterForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo{data=user.LoginResultsData}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Router			/user/register [post]
func Register(ctx *gin.Context) {
	//表单验证
	registerForm := user.RegisterForm{}

	if err := ctx.ShouldBind(&registerForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userInfo, err := userServices.Register(registerForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	loginSuccess(ctx, userInfo)
}

// Login
//
//	@Summary		用户登陆
//	@Description	用户登陆
//	@Tags			user用户
//	@Accept			json
//	@Produce		json
//	@Param			param	body		user.LoginForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo{data=user.LoginResultsData}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Router			/user/login [post]
func Login(ctx *gin.Context) {
	//表单验证
	loginForm := user.LoginForm{}

	if err := ctx.ShouldBind(&loginForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userInfo, err := userServices.VerifyUserPassword(loginForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	loginSuccess(ctx, userInfo)
}
