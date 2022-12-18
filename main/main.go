package main

import (
	"SimpleDecrypt/utils"
	"bufio"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func main() {
	run()
}
func run() {
	dbg := true
	recvListFile := "" // 接收者列表文件名
	mailBodyFile := ""
	mailTitle := ""
	mailConfig := utils.MailConfig{
		Host:     "",
		Port:     465,
		Username: "",
		Password: "",
	}
	app := &cli.App{
		Name:      "EmailSender",
		Usage:     "批量发送邮件",
		UsageText: "./EmailSender -s -p -u -pa -l -mbf -mh",
		Version:   "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "smtpServer", Aliases: []string{"s"}, Destination: &mailConfig.Host, Usage: "smtp服务器host", Required: true},
			&cli.IntFlag{Name: "smtpPort", Aliases: []string{"p"}, Destination: &mailConfig.Port, Value: 465, Usage: "smtp服务器端口", Required: false},
			&cli.StringFlag{Name: "username", Aliases: []string{"u"}, Destination: &mailConfig.Username, Usage: "发送者用户名", Required: true},
			&cli.StringFlag{Name: "password", Aliases: []string{"pa"}, Destination: &mailConfig.Password, Usage: "发送者密码", Required: true},
			&cli.StringFlag{Name: "recverList", Aliases: []string{"l"}, Destination: &recvListFile, Usage: "接收者列表", Required: true},
			&cli.StringFlag{Name: "mailbodyfile", Aliases: []string{"mbf"}, Destination: &mailBodyFile, Usage: "邮件BODY 的文件", Required: true},
			&cli.StringFlag{Name: "mailheader", Aliases: []string{"mh"}, Destination: &mailTitle, Usage: "邮件标题", Required: true},
			&cli.BoolFlag{Name: "DBGMOD", Aliases: []string{"DBG"}, Destination: &dbg, Value: false, Usage: "DBG MOD", Required: false},
		}, // 配置信息 fwcirjbvslnsdcgj
		HideHelpCommand: true,
		Action: func(c *cli.Context) error { // 发送
			utils.ShowLogo() // 主程序
			mailBody, filerederr := os.ReadFile(mailBodyFile)
			if filerederr != nil {
				utils.Printerr("邮件正文文件读取失败")
				return filerederr
			}
			recvlist, recvListFileerr := os.Open(recvListFile)
			if recvListFileerr != nil {
				return recvListFileerr
			}
			scanner := bufio.NewScanner(recvlist)
			for scanner.Scan() { // 读取接收者列表
				recver := scanner.Text()
				mailBodyTMP := strings.ReplaceAll(string(mailBody), "%RECV_EMAIL%", recver)
			mailsend:
				err := utils.SendMail(mailConfig.Username, mailTitle, mailBodyTMP, mailConfig, recver)
				if err != nil {
					utils.Printerr("%v->%v发送失败 重试中...", mailConfig.Username, recver)
					if dbg { // 调试模式报错
						panic(err)
					}
					goto mailsend
				}
			}

			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		utils.Printcritical("PROGRESS PANIC")
		os.Exit(0)
	}
}
