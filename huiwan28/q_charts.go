package huiwan28

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
)

func qCharts() (spaces map[int32]float64, result int32, issue int64, err error) {
	buf := new(bytes.Buffer)

	if err := Do("huiwan28_charts", map[string]string{}, nil, buf); err != nil {
		return spaces, result, issue, err
	}

	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return spaces, result, issue, err
	}

	// 获取开奖间隔系数、最新开奖结果和最新开奖期数
	doc.Find("table").Each(func(iTable int, sTable *goquery.Selection) {
		sTable.Find("tr").Each(func(itr int, str *goquery.Selection) {
			switch itr {
			case 1:
				str.Find("th").Each(func(ith int, sth *goquery.Selection) {
					if ith == 0 || ith >= 28 {
						return
					}

					f64, err := strconv.ParseFloat(sth.Text(), 32)
					if err != nil {
						log.Panicf("strconv.ParseFloat() Failure :: %s", err.Error())
						return
					}

					num := int32(ith) - 1
					spaces[num] = f64 / (1000 / float64(STDS1000[num]))

					if f64 == 0 {
						result = num
					}
				})
			case 5:
				str.Find("td").Each(func(itd int, std *goquery.Selection) {
					if itd != 0 {
						return
					}

					i64, err := strconv.ParseInt(std.Text(), 10, 64)
					if err != nil {
						log.Panicf("strconv.ParseInt() Failure :: %s", err.Error())
					}

					issue = i64
				})
			}
		})
	})

	// 获取最新开奖期数
	log.Printf("实际间隔：%v\n", spaces)
	log.Printf("最新开奖结果：%d\n", result)
	log.Printf("最新开奖期数：%d\n", issue)

	return spaces, result, issue, nil
}
