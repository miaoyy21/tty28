package client

import (
	"fmt"
	"log"
)

func r3Fn(issue int, ns int) (map[int32]float64, error) {
	rds, err := qRiddle(fmt.Sprintf("%d", issue+1), ns)
	if err != nil {
		return nil, err
	}

	log.Printf("(3) 查询赔率完成 ... \n")
	return rds, nil
}
