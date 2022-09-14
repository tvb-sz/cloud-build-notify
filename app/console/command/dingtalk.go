package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tvb-sz/cloud-build-notify/app/service"
)

// init version子命令
func init() {
	var token, secret, title, content string

	dingTalk := &cobra.Command{
		Use:   "dingTalk",
		Short: "发送钉钉机器人通知",
		Long:  "发送钉钉机器人通知",
		Run: func(cmd *cobra.Command, args []string) {
			service.NotifyService.InitDing(token, secret)
			if err := service.NotifyService.DingNotifyOk(title, content); err != nil {
				fmt.Printf("send fail: %s", err)
			} else {
				fmt.Println("send ok")
			}
		},
	}

	dingTalk.Flags().StringVar(&token, "token", "", "钉钉令牌：即创建钉钉机器人得到的URL里的access_token")
	dingTalk.Flags().StringVar(&secret, "secret", "", "钉钉秘钥：签名校验用的秘钥")
	dingTalk.Flags().StringVar(&title, "title", "", "消息体标题")
	dingTalk.Flags().StringVar(&content, "content", "", "markdown格式的消息内容")

	RootCmd.AddCommand(dingTalk)
}
