package huiwan28

import (
	"bytes"
	"log"
)

func qCharts() {
	buf := new(bytes.Buffer)

	if err := Do("huiwan28_charts", map[string]string{}, nil, buf); err != nil {
		return
	}

	log.Printf("Charts is %s", buf.String())
}
