package luck

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
)

type User struct {
	UserId   string
	UserName string
	Gold     int64

	Spaces map[int32]float64
	Result int32
	Issue  int64
}

func qCharts() (*User, error) {
	buf := new(bytes.Buffer)

	if err := Do("huiwan28_charts", map[string]string{}, nil, buf); err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return nil, err
	}

	user := &User{
		Spaces: make(map[int32]float64),
	}

	div := doc.Find("div").Find("div").Find("div").Find("div").Find("div")

	// 用户ID
	user.UserId = div.Eq(0).Find("a").Find("span").Text()

	// 用户名
	user.UserName = div.Eq(1).Find("a").Text()

	// 剩余金额
	sGold := div.Eq(2).Find("a").Find("span").Text()
	gold, err := strconv.ParseInt(strings.ReplaceAll(sGold, ",", ""), 10, 64)
	if err != nil {
		return nil, err
	}
	user.Gold = gold

	// 获取开奖间隔系数、最新开奖结果和最新开奖期数
	sTable := doc.Find("table").Eq(0)
	sTable.Find("tr").Each(func(itr int, str *goquery.Selection) {
		switch itr {
		case 1:
			str.Find("th").Each(func(ith int, sth *goquery.Selection) {
				if ith == 0 || ith >= 29 {
					return
				}

				f64, err := strconv.ParseFloat(sth.Text(), 32)
				if err != nil {
					log.Panicf("strconv.ParseFloat() Failure :: %s", err.Error())
				}

				num := int32(ith) - 1
				user.Spaces[num] = f64 / (1000 / float64(STDS1000[num]))

				if f64 == 0 {
					user.Result = num
				}
			})
		case 5:
			std := str.Find("td").Eq(0)

			i64, err := strconv.ParseInt(std.Text(), 10, 64)
			if err != nil {
				log.Panicf("strconv.ParseInt() Failure :: %s", err.Error())
			}

			user.Issue = i64
		}
	})

	return user, nil
}
