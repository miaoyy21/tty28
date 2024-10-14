package luck

import (
	"log"
	"time"
)

func Run() error {
	// 配置文件
	if err := InitConfig(); err != nil {
		return err
	}

	// 如果超过30秒，那么等待30秒后运行
	if time.Now().Second() >= 30 {
		log.Printf("当前时间为%q，等待30秒 ... \n", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(30 * time.Second)
	}

	sleepTo(30)
	go run()

	t := time.NewTicker(time.Minute)
	defer t.Stop()

	log.Println("启动定时器完成 ...")
	for {
		select {
		case <-t.C:
			// 长时间运行时，可能会产生时间偏移，自动调整
			d0 := time.Now().Sub(time.Now().Truncate(time.Minute))
			t.Reset(time.Duration(90-d0.Seconds()) * time.Second)
			log.Printf("【重置时钟】偏移量%.2f秒 ...\n", 30-d0.Seconds())

			// 重新加载配置
			if err := InitConfig(); err != nil {
				log.Printf("重新加载配置文件错误：%s \n", err.Error())
				continue
			}
			log.Printf("重载配置文件成功 ...\n")

			go run()
		}
	}
}
