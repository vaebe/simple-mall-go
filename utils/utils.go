package utils

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math/rand"
	"net"
	"strings"
	"time"
)

// GetFreePort 获取可用端口
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer func(l *net.TCPListener) {
		err := l.Close()
		if err != nil {
			zap.S().Error("获取可用端口信息失败！")
		}
	}(l)
	return l.Addr().(*net.TCPAddr).Port, nil
}

// GenerateSmsCode 生成width长度的短信验证码
func GenerateSmsCode(width int) string {

	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.NewSource(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			return ""
		}
	}
	return sb.String()
}

// Paginate 分页
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// GetFileSuffixName 获取文件后缀名称
func GetFileSuffixName(filename string) string {
	indexOfDot := strings.LastIndex(filename, ".") //获取文件后缀名前的.的位置
	if indexOfDot < 0 {
		return ""
	}
	suffix := filename[indexOfDot+1 : len(filename)] //获取后缀名
	return strings.ToLower(suffix)                   //后缀名统一小写处理
}
