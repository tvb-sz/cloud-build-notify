package initializer

import "github.com/tvb-sz/cloud-build-notify/client"

// region 全局句柄初始化相关

// Init 初始化
//go:noinline
func Init() {
	client.Logger = iniLogger() // 初始化logger，需要优先执行
}

// endregion
