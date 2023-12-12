package client

import (
	"fmt"
	"log"
	"math"
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

func qRiddle(issue string, dz float64, ns int) (map[int32]float64, float64, float64, error) {
	var resp QRiddleResponse

	qUrl := fmt.Sprintf("%s?utoken=%s&cid=%s&stylePath=%s&t=%d", conf.RiddleURL, conf.UToken, issue, conf.Style, ns)
	err := hdo.Do(conf.Authority, conf.Origin, conf.Referer, conf.SecChUa, conf.SecChUaPlatform, conf.UserAgent, qUrl, &resp)
	if err != nil {
		return nil, 0, 0, err
	}

	if resp.Code != 0 {
		return nil, 0, 0, fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Code, resp.Msg)
	}

	var coverage float64
	var exp float64

	rts := make(map[int32]float64)
	for _, r := range resp.Data.List {
		rts[r.No] = r.Odd / (1000.0 / float64(STDS1000[r.No]))
		exp = exp + (float64(STDS1000[r.No])/1000)*rts[r.No]

		if rts[r.No] < dz {
			log.Printf("竞猜数字【   %02d】，实际赔率【%7.2f】，赔率系数【%.3f】 \n", r.No, r.Odd, rts[r.No])
			continue
		}

		coverage = coverage + float64(STDS1000[r.No])
		log.Printf("竞猜数字【 ✓ %02d】，实际赔率【%7.2f】，赔率系数【%.3f】 \n", r.No, r.Odd, rts[r.No])
	}

	var dev float64
	for n, r0 := range rts {
		r1 := (float64(STDS1000[n]) / 1000) * (r0 / (1000.0 / float64(STDS1000[n])))
		dev = dev + (r1-exp)*(r1-exp)*(float64(STDS1000[n])/1000)
	}

	return rts, coverage, math.Sqrt(dev), nil
}
