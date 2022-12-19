package main

import (
	"SimpleDecrypt/Var"
	"SimpleDecrypt/utils"
	"bufio"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	run()
}
func run() {

	mailConfig := utils.MailConfig{
		Host:     "",
		Port:     465,
		Username: "",
		Password: "",
	}
	app := &cli.App{
		Name:      "EmailSender",
		Usage:     "批量发送邮件",
		UsageText: "./EmailSender -h",
		Version:   "0.0.2",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "smtpServer", Aliases: []string{"s"}, Destination: &mailConfig.Host, Usage: "smtp服务器host", Required: true},
			&cli.IntFlag{Name: "smtpPort", Aliases: []string{"p"}, Destination: &mailConfig.Port, Value: 465, Usage: "smtp服务器端口", Required: false},
			&cli.StringFlag{Name: "username", Aliases: []string{"u"}, Destination: &mailConfig.Username, Usage: "发送者用户名", Required: true},
			&cli.StringFlag{Name: "password", Aliases: []string{"pa"}, Destination: &mailConfig.Password, Usage: "发送者密码", Required: true},
			&cli.StringFlag{Name: "recverList", Aliases: []string{"l"}, Destination: &Var.RecvListFile, Usage: "接收者列表", Required: true},
			&cli.StringFlag{Name: "mailbodyfile", Aliases: []string{"mbf"}, Destination: &Var.MailBodyFile, Usage: "邮件BODY 的文件", Required: true},
			//&cli.IntFlag{Name: "groupsize", Aliases: []string{"gs"}, Destination: &Var.GroupSize, Usage: "群发分组邮件人数(批量单发不用填写此选项)", Required: false},
			&cli.StringFlag{Name: "mailheader", Aliases: []string{"mh"}, Destination: &Var.MailTitle, Usage: "邮件标题", Required: true},
			&cli.StringFlag{Name: "attachfile", Aliases: []string{"af"}, Destination: &Var.AttachFile, Value: "", Usage: "附件文件(会大幅降低邮件发送速度)", Required: false},
			&cli.BoolFlag{Name: "DBGMOD", Aliases: []string{"DBG"}, Destination: &Var.Dbg, Value: false, Usage: "DBG MOD", Required: false},
		}, // 配置信息 fwcirjbvslnsdcgj
		HideHelpCommand: true,
		Action: func(c *cli.Context) error { // 发送
			utils.ShowLogo() // 主程序
			mailBody, filerederr := os.ReadFile(Var.MailBodyFile)
			if filerederr != nil {
				utils.Printerr("邮件正文文件读取失败")
				return filerederr
			}
			recvlist, recvListFileerr := os.Open(Var.RecvListFile)
			if recvListFileerr != nil {
				return recvListFileerr
			}
			scanner := bufio.NewScanner(recvlist)

			//if Var.GroupSize == 1 { // 如果是批量单发
			for scanner.Scan() { // 读取接收者列表
				recver := scanner.Text()
				utils.SendSingle(recver, mailConfig, mailBody) // 发送邮件
			}
			//} else if Var.GroupSize > 1 { //如果是群发
			//	utils.SendMulti(scanner, mailConfig, mailBody) // 群发
			//}

			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		utils.Printcritical("PROGRESS PANIC")
		os.Exit(0)
	}
}
