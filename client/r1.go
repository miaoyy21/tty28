package client

import (
	"errors"
	"log"
)

func r1Fn(ns int) (int, error) {

	issue, total, err := qIssueGold(ns)
	if err != nil {
		return 0, err
	}

	log.Printf("(1) 最新开奖期数【%d】，资金池总量【%d】 ... \n", issue, total)
	if total < 1000000000 {
		return 0, errors.New("********************************")
	}

	return issue, nil
}
