package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidatorMobile  自定义手机号验证
func ValidatorMobile(mobile validator.FieldLevel) bool {
	str := mobile.Field().String()
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, str)
	return ok
}
