package client

import (
	"errors"
	"fmt"
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

			Bet   interface{} `json:"bet"`
			Prize interface{} `json:"prize"`
		} `json:"list"`
	} `json:"data"`
}

func qIssueGold(ns int) (int, int64, error) {
	var resp QIssueResponse

	qUrl := fmt.Sprintf("%s?utoken=%s&stylePath=%s&t=%d", conf.IssueURL, conf.UToken, conf.Style, ns)
	err := hdo.Do(conf.Authority, conf.Origin, conf.Referer, conf.SecChUa, conf.SecChUaPlatform, conf.UserAgent, qUrl, &resp)
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
	var wins, fails int
	var winGold int64
	for _, l := range resp.Data.List[:20] {
		if l.Status == 2 {
			if min == 0 {
				min, err = hdo.Int64(l.TotalBet)
				if err != nil {
					return 0, 0, err
				}
			}

			continue
		}

		if issue <= 0 {
			issue = l.Cid
			gold, err = hdo.Int64(l.TotalBet)
			if err != nil {
				return 0, 0, err
			}
		}

		var bet, prize int64

		bet, err = hdo.Int64(fmt.Sprint(l.Bet))
		if err != nil {
			return 0, 0, err
		}

		prize, err = hdo.Int64(fmt.Sprint(l.Prize))
		if err != nil {
			return 0, 0, err
		}

		if prize > bet {
			wins++
		} else if prize < bet {
			fails++
		}

		if wins+fails < 5 {
			winGold = winGold + prize - bet
		}
	}

	if winGold < 0 && wins+fails > 5 {
		return 0, 0, fmt.Errorf("一直在亏损，并且投注次数超过%d次，不进行投注 ... ", 5)
	}

	return issue, gold - min, nil
}
