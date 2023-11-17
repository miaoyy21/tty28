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

	// 输出权重值
	for _, n := range SN28 {
		sigma := 0.985
		if rds[n] <= sigma {
			log.Printf("竞猜数字【 %02d - 】； \n", n)
			continue
		}

		var sig float64
		if rds[n] > 1.0 {
			sig = rds[n]
			log.Printf("竞猜数字【 %02d H 】，权重值【%.2f】； \n", n, sig)
		} else {
			sig = (rds[n] - sigma) / (1.0 - sigma)
			log.Printf("竞猜数字【 %02d L 】，权重值【%.2f】； \n", n, sig)
		}
	}

	return rds, nil
}
