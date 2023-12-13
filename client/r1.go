package client

import (
	"log"
)

func r1Fn(sn int, ns int) (int, bool, error) {

	issue, total, stop, err := qIssueGold(sn, ns)
	if err != nil {
		return 0, false, err
	}

	log.Printf("(1) 最新开奖期数【%d】，资金池总量【%d】 ... \n", issue, total)

	return issue, stop, nil
}
