package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tvb-sz/cloud-build-notify/app/console"
	"github.com/tvb-sz/cloud-build-notify/client/initializer"
	"os"
)

// RootCmd 基于cobra的命令行根节点定义
var (
	RootCmd = &cobra.Command{
		Use:   "notify",
		Short: "google cloud run ci notify",
		Long:  "google cloud run ci notify",
		Run: func(cmd *cobra.Command, args []string) {
			console.Help()
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// init global client handle
			initializer.Init()

			return nil
		},
	}
)

func init() {

}

// Start 启动应用
func Start() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err) // client.Logger可能尚未init，此处尚不能使用
		os.Exit(1)
	}
}
