package huiwan28

import (
	"fmt"
	"net/url"
	"strings"
)

type RecordRequest struct {
	Number  string `json:"number"`
	Periods string `json:"periods"`
	Money   string `json:"money"`
}

func qRecord(issue string, sBets []string) error {
	sNums := make([]string, 0, 28)
	for _, num := range SN28 {
		sNums = append(sNums, fmt.Sprintf("%d", num))
	}

	// data-raw
	values := url.Values{}
	values.Add("number", strings.Join(sNums, ","))
	values.Add("periods", issue)
	values.Add("money", strings.Join(sBets, ","))

	// 更新Header
	headers := map[string]string{
		"referer": fmt.Sprintf("https://www.huiwan28.com/Culmination/record/periods/%s.html", issue),
	}

	if err := Do("huiwan28_record", headers, strings.NewReader(values.Encode()), nil); err != nil {
		return err
	}

	return nil
}
