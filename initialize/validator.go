package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"

	"simple-mall/global"
	myValidator "simple-mall/validator"
)

func InitTrans(locale string) (err error) {
	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			return en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			return zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			return en_translations.RegisterDefaultTranslations(v, global.Trans)
		}
	}
	return
}

// CustomValidators 自定义表单验证器规则
func CustomValidators() {
	// 设置自定义表单验证
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 手机号自定义验证器
		err := v.RegisterValidation("mobile", myValidator.ValidatorMobile)
		if err != nil {
			zap.S().Debug("ValidatorMobile 验证错误", err.Error())
			return
		}

		// 自定义验证错误规则信息
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
}
