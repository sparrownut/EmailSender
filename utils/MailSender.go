package utils

import (
	"SimpleDecrypt/Var"
	"SimpleDecrypt/datastruct"
	"SimpleDecrypt/varfiliter"
	"gopkg.in/gomail.v2"
)

type MailConfig struct {
	/*
		邮件信息配置器
	*/
	Host     string
	Port     int
	Username string
	Password string
}

func SendMail(sendPackage datastruct.EmailStructs, mailconfig MailConfig) error {
	/*
		发送邮件函数
	*/
	mail := gomail.NewMessage()
	if sendPackage.AttachFile != "" { // 添加附件
		mail.Attach(sendPackage.AttachFile)
	}
	mail.SetHeader("From", sendPackage.From)
	mail.SetHeader("To", sendPackage.To...)
	mail.SetHeader("Subject", sendPackage.Subject)
	mail.SetBody("text/html", string(sendPackage.Text))
	sender := gomail.NewDialer(mailconfig.Host, mailconfig.Port, mailconfig.Username, mailconfig.Password)
	err := sender.DialAndSend(mail)
	if err != nil {
		return err
	}
	Printsuc("%v -> %v发送成功", sendPackage.From, sendPackage.To)
	return nil
}

func SendSingle(recver string, mailConfig MailConfig, mailBody []byte) {
	varfiliter.VarFiliter(mailConfig.Username, recver, &mailBody)
	varfiliter.FormatFiliter(&recver)
	sendPackage := datastruct.EmailStructs{
		From:       mailConfig.Username,
		To:         datastruct.StringToStringArray(recver),
		Text:       mailBody,
		AttachFile: Var.AttachFile,
	}

mailsend:
	err := SendMail(sendPackage, mailConfig)
	if err != nil {
		Printerr("%v->%v发送失败 重试中...", mailConfig.Username, recver)
		if Var.Dbg { // 调试模式报错
			panic(err)
		}
		goto mailsend
	}
}

//func SendMulti(scanner *bufio.Scanner, mailconfig MailConfig, mailBody []byte) {
//	var sendPackageList []datastruct.EmailStructs
//	i := 0
//	for scanner.Scan() {
//		//打包成发送包结构
//		if i < Var.GroupSize {
//			sendPackage := datastruct.EmailStructs{
//				From: mailconfig.Username,
//				To:   datastruct.StringToStringArray(scanner.Text()),
//				Text: mailBody,
//				AttachFile: asdfasdfsadfasdf
//			}
//			sendPackageList = append(sendPackageList, sendPackage) // 添加列表
//		} else { // 群发数量到了设定值
//
//			recvers := sendPackageList
//			varfiliter.GroupFiliter(&sendPackageList)
//			//varfiliter.VarFiliter(mailConfig.Username, recvers, &mailBodyTMP)
//		mailGroupsend:
//			err := SendMail()
//			if err != nil {
//				utils.Printerr("%v->%v发送失败 重试中...", mailConfig.Username, recvers)
//				if Dbg { // 调试模式报错
//					panic(err)
//				}
//				goto mailGroupsend
//			}
//			i = 0
//			sendPackageList = nil // init
//		}
//		i++
//	}
//}
