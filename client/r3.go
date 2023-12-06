package client

import (
	"fmt"
	"log"
)

func r3Fn(issue int, ns int) (map[int32]float64, error) {
	rds, coverage, err := qRiddle(fmt.Sprintf("%d", issue+1), ns)
	if err != nil {
		return nil, err
	}

	log.Printf("(3) 获取赔率覆盖率【%.2f%%】 ... \n", coverage/10)
	return rds, nil
}
