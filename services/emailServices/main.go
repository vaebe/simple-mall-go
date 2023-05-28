package emailServices

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"simple-mall/global"
)

// SendTheVerificationCodeEmail 发送验证码邮件
func SendTheVerificationCodeEmail(VerificationCode string, emailAddress string) error {
	mailUserName := "209005801@qq.com"     // 邮箱账号
	mailPassword := global.EmailConfig.Key // 邮箱授权码
	addr := "smtp.qq.com:465"              // TLS地址
	host := "smtp.qq.com"                  // 邮件服务器地址
	Subject := "MK社区验证码"                   // 发送的主题

	e := email.NewEmail()
	e.From = "MK社区 <209005801@qq.com>"
	e.To = []string{emailAddress}
	e.Subject = Subject
	e.HTML = []byte("您的验证码为：<h1>" + VerificationCode + "</h1>")
	return e.SendWithTLS(addr, smtp.PlainAuth("", mailUserName, mailPassword, host),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
}
