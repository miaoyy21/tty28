package client

import (
	"fmt"
	"log"
	"tty28/hdo"
)

type QRiddleResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []struct {
			No  int32   `json:"no"`
			Odd float64 `json:"odd"`
		} `json:"list"`
	} `json:"data"`
}

func qRiddle(issue string, ns int) (map[int32]float64, error) {
	var resp QRiddleResponse

	qUrl := fmt.Sprintf("%s?utoken=%s&cid=%s&stylePath=happy&t=%d", conf.RiddleURL, conf.UToken, issue, ns)
	err := hdo.Do(conf.Authority, conf.Origin, conf.Referer, conf.SecChUa, conf.SecChUaPlatform, conf.UserAgent, qUrl, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Code, resp.Msg)
	}

	sig := 0.99
	log.Printf("赔率折算系数为【%.2f】\n", sig)
	rts := make(map[int32]float64)
	for _, r := range resp.Data.List {
		rts[r.No] = r.Odd * sig / (1000.0 / float64(STDS1000[r.No]))
		log.Printf("竞猜数字【 %02d 】，实际赔率【%.3f】，赔率系数【%.3f】 \n", r.No, r.Odd, rts[r.No])
	}

	return rts, nil
}
