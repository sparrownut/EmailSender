package utils

import "gopkg.in/gomail.v2"

type MailConfig struct {
	/*
		邮件信息配置器
	*/
	Host     string
	Port     int
	Username string
	Password string
}

func SendMail(from string, subject string, body string, mailconfig MailConfig, to ...string) error {
	/*
		发送邮件函数
	*/
	mail := gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", to...)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)
	sender := gomail.NewDialer(mailconfig.Host, mailconfig.Port, mailconfig.Username, mailconfig.Password)
	err := sender.DialAndSend(mail)
	if err != nil {
		return err
	}
	Printsuc("%v -> %v发送成功", from, to)
	return nil
}
