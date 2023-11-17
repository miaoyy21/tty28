package client

import (
	"fmt"
)

func r3Fn(issue int, ns int) (map[int32]float64, error) {
	rds, err := qRiddle(fmt.Sprintf("%d", issue+1), ns)
	if err != nil {
		return nil, err
	}

	return rds, nil
}
