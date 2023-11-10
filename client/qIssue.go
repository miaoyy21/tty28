package client

import (
	"errors"
	"fmt"
	"time"
	"tty28/hdo"
)

type QIssueResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []struct {
			Status   int    `json:"status"`
			Cid      int    `json:"cid"`
			TotalBet string `json:"total_bet"`
		} `json:"list"`
	} `json:"data"`
}

func qIssueGold() (int, int64, error) {
	var resp QIssueResponse

	qUrl := fmt.Sprintf("%s?utoken=%s&stylePath=happy&t=%d", conf.IssueURL, conf.UToken, time.Now().UnixNano())
	err := hdo.Do(conf.Origin, conf.Referer, conf.SecChUa, conf.SecChUaPlatform, conf.UserAgent, qUrl, &resp)
	if err != nil {
		return 0, 0, err
	}

	if resp.Code != 0 {
		return 0, 0, fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Code, resp.Msg)
	}

	if len(resp.Data.List) < 1 {
		return 0, 0, errors.New("没有收到返回结果")
	}

	var issue int
	var gold, min int64
	for _, l := range resp.Data.List {
		if l.Status == 0 {
			if min == 0 {
				min, err = toInt64(l.TotalBet)
				if err != nil {
					return 0, 0, err
				}
			}
		} else {
			issue = l.Cid
			gold, err = toInt64(l.TotalBet)
			if err != nil {
				return 0, 0, err
			}
			break
		}
	}

	return issue, gold - min, nil
}
