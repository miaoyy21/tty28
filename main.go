package main

import (
	"log"
	"tty28/canada"
	"tty28/conf"
)

func main() {
	// 加载配置模版
	if err := conf.LoadTemplates(); err != nil {
		log.Panicf("config.LoadTemplates() Failure :: %s", err.Error())
	}

	// 幸运28
	//if err := luck.Run(); err != nil {
	//	log.Printf("luck.Run() Failure :: %s \n", err.Error())
	//}

	// 加拿大28
	if err := canada.Run(); err != nil {
		log.Printf("canada.Run() Failure :: %s \n", err.Error())
	}
}
