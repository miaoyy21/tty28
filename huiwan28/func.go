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
func Do(tempName string, headers map[string]string, r io.Reader, w io.Writer) error {
	temp := config.Templates[tempName]

	// Request
	req, err := http.NewRequest(temp.Method, temp.URL, r)
	if err != nil {
		return err
	}

	// 添加基础的Header
	for k, v := range temp.Headers {
		req.Header.Set(k, v)
	}

	// 添加Cookie
	req.Header.Set("cookie", conf.Cookie)

	// 添加变更的Header
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	http.DefaultClient.Timeout = 15 * time.Second
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err := io.Copy(w, resp.Body); err != nil {
		return err
	}

	return nil
}
