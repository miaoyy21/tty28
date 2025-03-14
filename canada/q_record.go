package canada

import (
	"fmt"
	"net/url"
	"strings"
	"tty28/base"
)

func qRecord(issue string, sBets []string) error {
	sNums := make([]string, 0, 28)
	for _, num := range base.SN28 {
		sNums = append(sNums, fmt.Sprintf("%d", num))
	}

	// data-raw
	values := url.Values{}
	values.Add("number", strings.Join(sNums, ","))
	values.Add("periods", issue)
	values.Add("money", strings.Join(sBets, ","))

	// 更新Header
	headers := map[string]string{
		"referer": fmt.Sprintf("https://www.huiwan28.com/Canada/record/periods/%s.html", issue),
	}

	if err := base.Do("canada_record", headers, strings.NewReader(values.Encode()), nil); err != nil {
		return err
	}

	return nil
}
