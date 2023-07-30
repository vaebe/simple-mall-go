package pay

import (
	"crypto/md5"
	"encoding/hex"
	"simple-mall/global"
	"sort"
	"strings"
)

func calculateMD5(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// LTZFGenerateSignature 生成蓝兔支付签名
func LTZFGenerateSignature(sinObj map[string]string) string {
	// 按照参数名进行字典序排序
	keys := make([]string, 0, len(sinObj))
	for k := range sinObj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接参数键值对
	var builder strings.Builder
	for _, key := range keys {
		value := sinObj[key]
		if value != "" {
			builder.WriteString(key + "=" + value + "&")
		}
	}

	// 拼接密钥
	builder.WriteString("key=" + global.LTZFConfig.SecretKey)

	// 计算MD5并转换为大写
	signature := calculateMD5(builder.String())

	return strings.ToUpper(signature)
}
