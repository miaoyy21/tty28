package main

import (
	"log"
	"tty28/config"
	"tty28/huiwan28"
)

func main() {
	// 加载配置模版
	if err := config.LoadTemplates(); err != nil {
		log.Panicf("config.LoadTemplates() Failure :: %s", err.Error())
	}

	// 运行定时任务
	if err := huiwan28.Run(); err != nil {
		log.Printf("huiwan28.Run() Failure :: %s \n", err.Error())
	}
}
