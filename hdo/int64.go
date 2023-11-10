package hdo

import (
	"strconv"
	"strings"
)

func Int64(s string) (int64, error) {
	s0 := strings.ReplaceAll(s, ",", "")

	i64, err := strconv.ParseInt(s0, 10, 64)
	if err != nil {
		return 0, err
	}

	return i64, nil
}
