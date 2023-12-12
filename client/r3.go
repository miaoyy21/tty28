package client

import (
	"fmt"
	"log"
)

func r3Fn(issue int, dz float64, ns int) (map[int32]float64, float64, error) {
	rds, coverage, dev, err := qRiddle(fmt.Sprintf("%d", issue+1), dz, ns)
	if err != nil {
		return nil, 0, err
	}

	log.Printf("(3) 获取赔率覆盖率【%.2f%%】，平均方差【%.4f】 ... \n", coverage/10, dev)
	return rds, dev, nil
}
