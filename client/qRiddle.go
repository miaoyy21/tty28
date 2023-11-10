package client

import (
	"fmt"
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

	rts := make(map[int32]float64)
	for _, r := range resp.Data.List {
		rts[r.No] = r.Odd / (1000.0 / float64(STDS1000[r.No]))
	}

	return rts, nil
}
