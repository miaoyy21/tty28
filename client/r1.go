package client

import "log"

func r1Fn(ns int) (int, float64, error) {

	issue, total, err := qIssueGold(ns)
	if err != nil {
		return 0, 0, err
	}

	mrx := 1.0
	if total < 1<<27 {
		mrx = float64(total) / float64(1<<28) // 268,435,456
	}

	log.Printf("(1) 最新开奖期数【%d】，资金池总量【%d，系数%.2f】 ... \n", issue, total, mrx)
	return issue, mrx, nil
}
