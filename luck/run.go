package luck

import (
	"fmt"
	"log"
	"math/rand"
	"tty28/base"
)

var parsed = 3

func run() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	log.Println("//*********************************** 定时任务开始执行 ***********************************//")

	// 第1步：查询开奖历史
	base.SleepTo(30 + 5*rand.Float64())
	user, err := qCharts()
	if err != nil {
		log.Printf("【ERR-0】: %s \n", err.Error())
		return
	}

	log.Printf("【1】用户ID: %s，用户名: %s，剩余金额: %d ...\n", user.UserId, user.UserName, user.Gold)
	log.Printf("【1】最新开奖期数【%d】，开奖结果【%02d】", user.Issue, user.Result)

	// 第2步：打印各间隔频率
	sBets, total := make([]string, 0), uint32(0)
	log.Printf("【2】开奖历史的各数字间隔频率：\n")
	for _, num := range base.SN28 {
		space := user.Spaces[num]

		var rate float64
		if space <= 0.67 {
			rate = 1.00
		} else if space <= 1.00 {
			rate = 0.80
		} else if space <= 1.33 {
			rate = 0.40
		} else {
			rate = 0
		}

		iGold := uint32(base.Config.Base * rate * float64(base.STDS1000[num]) / 1000)
		if iGold > 0 {
			log.Printf("\t竞猜数字【%02d】：间隔频率【%5.3f】，投注系数【%4.2f】，投注金额【%6d】；\n", num, space, rate, iGold)
		} else {
			log.Printf("\t竞猜数字【%02d】：间隔频率【%5.3f】，投注系数【 - 】；\n", num, space)
		}

		total = total + iGold
		sBets = append(sBets, fmt.Sprintf("%d", iGold))
	}
	nextIssue := base.OfNextIssue(user.Issue)
	log.Printf("【2】投注期数【%s】，投注基数【%d】，投注总额【%d】...\n", nextIssue, int(base.Config.Base), total)

	// 第3步：执行投注
	base.SleepTo(40 + 5*rand.Float64())
	log.Printf("【3】执行投注 ...")
	if err := qRecord(nextIssue, sBets); err != nil {
		log.Printf("【ERR-3】: %s \n", err.Error())
		return
	}

	log.Println("【9】投注成功，全部执行结束 ...")
}
