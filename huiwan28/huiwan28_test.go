package huiwan28

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestRun(t *testing.T) {
	path := "/Users/miaojingyi/Documents/dev/go/src/tty28/index.html"

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("os.Open() Failure :: %s", err.Error())
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		t.Fatalf("io.Copy() Failure :: %s", err.Error())
	}

	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		t.Fatalf("goquery.NewDocumentFromReader() Failure :: %s", err.Error())
	}

	// 获取开奖间隔系数、最新开奖结果和最新开奖期数
	spaces, result, issue := make(map[int32]float64), int32(0), int64(0)
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
						t.Fatalf("strconv.ParseFloat() Failure :: %s", err.Error())
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
						t.Fatalf("strconv.ParseInt() Failure :: %s", err.Error())
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

	//re0 := regexp.MustCompile(`<table>(.*?)</table>`)
	//
	//ss := re0.FindStringSubmatch(buf.String())
	//t.Log(ss)
}
