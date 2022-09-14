package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tvb-sz/cloud-build-notify/app/service"
)

// init version子命令
func init() {
	var token, secret, title, content string

	feiShu := &cobra.Command{
		Use:   "feiShu",
		Short: "发送飞书机器人通知",
		Long:  "发送飞书机器人通知",
		Run: func(cmd *cobra.Command, args []string) {
			service.NotifyService.InitFeishu(token, secret)
			if err := service.NotifyService.FeishuNotifyOk(title, content); err != nil {
				fmt.Printf("send fail: %s", err)
			} else {
				fmt.Println("send ok")
			}
		},
	}

	feiShu.Flags().StringVar(&token, "token", "", "飞书令牌：即创建飞书机器人得到的URL里的内容")
	feiShu.Flags().StringVar(&secret, "secret", "", "飞书秘钥：签名校验用的秘钥")
	feiShu.Flags().StringVar(&title, "title", "", "消息体标题")
	feiShu.Flags().StringVar(&content, "content", "", "markdown格式的消息内容")

	RootCmd.AddCommand(feiShu)
}
