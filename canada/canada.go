package canada

import (
	"log"
	"time"
	"tty28/base"
)

func Run() error {

	// 运行时间判定，超过 dt%210
	base.SleepTo(45)
	go run()

	t := time.NewTicker(time.Minute)
	defer t.Stop()

	log.Println("启动定时器完成 ...")
	for {
		select {
		case <-t.C:

			go run()
		}
	}
}
