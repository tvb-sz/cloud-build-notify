package main

import (
	"github.com/tvb-sz/cloud-build-notify/app/console/command"
	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
	// 设置GMP模型中的P的最大值；对于实体机运行默认设置为实体机CPU核心数就够用
	// 对于docker运行因各种原因会导致这个P数量被错误的设置导致一些问题，此处引入这个库进行设置主要针对的就是docker环境
	undo, _ := maxprocs.Set()
	defer undo()

	command.Start()
}
