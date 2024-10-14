package main

import (
	"log"
	"tty28/config"
	"tty28/luck"
)

func main() {
	// 加载配置模版
	if err := config.LoadTemplates(); err != nil {
		log.Panicf("config.LoadTemplates() Failure :: %s", err.Error())
	}

	// 运行定时任务
	if err := luck.Run(); err != nil {
		log.Printf("luck.Run() Failure :: %s \n", err.Error())
	}
}
