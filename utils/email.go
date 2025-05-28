package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
	"swiftDaily_myself/global"
)

func Email(to, subject, body string) error {
	To := strings.Split(to, ",") // 将收件人的地址安装逗号拆分成多个地址，因为支持多个收件人
	emailCfg := global.Config.Email
	from := emailCfg.From
	nickname := emailCfg.NickName
	secret := emailCfg.Secret
	host := emailCfg.Host
	port := emailCfg.Port
	isSSL := emailCfg.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = To
	e.Subject = subject
	e.HTML = []byte(body)
	hostAddr := fmt.Sprintf("%s:%s", host, port)
	var err error
	if isSSL {
		// 使用带Tls的邮箱进行发送
		// TLS 是一种安全传输协议，用于加密客户端与邮件服务器之间的通信；
		// 保证邮件发送过程中的数据安全，防止被窃听或篡改。
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{
			ServerName: host,
		})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
