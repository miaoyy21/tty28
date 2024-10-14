package huiwan28

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"tty28/config"
)

func sleepTo(s0 float64) {
	d0 := time.Now().Sub(time.Now().Truncate(time.Minute))
	if s0-d0.Seconds() < 0 {
		log.Printf(fmt.Sprintf("【网络延迟原因？？？？？？】目标第%.2f秒小于当前第%.2f秒\n", s0, d0.Seconds()))
	}

	log.Printf("等待%.2f秒后继续执行 ... \n", s0-d0.Seconds())
	time.Sleep(time.Second * time.Duration(s0-d0.Seconds()))
}

// Do 执行HTTP请求
func Do(tempName string, cookie string, adjust map[string]string, body io.Reader) {
	temp := config.Templates[tempName]

	http.NewRequest(temp.Method, temp.URL, body)
}
