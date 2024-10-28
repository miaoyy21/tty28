package canada

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"
	"tty28/base"
)

func run() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	// 1728925740 => 2024-10-15 01:09:00
	seconds := int(time.Now().Sub(time.Unix(1728925740, 0)).Seconds())

	// 是否在执行时间段
	steps := float64(seconds%210) / 60
	if steps < 1 || steps > 2 {
		return
	}

	log.Println("//*********************************** 定时任务开始执行 ***********************************//")

	// 重新加载配置
	if err := base.InitConfig(); err != nil {
		log.Printf("重新加载配置文件错误：%s \n", err.Error())
		return
	}
	log.Printf("重载配置文件成功 ...\n")

	// 第1步：查询开奖历史
	time.Sleep(time.Duration(15*rand.Float64()) * time.Second)
	user, err := qCharts()
	if err != nil {
		log.Printf("【ERR-0】: %s \n", err.Error())
		return
	}

	log.Printf("【1】用户ID: %s，用户名: %s，剩余金额: %d ...\n", user.UserId, user.UserName, user.Gold)
	log.Printf("【1】最新开奖期数【%d】，开奖结果【%02d】", user.Issue, user.Result)

	// 开奖频率排序
	sort.Slice(user.Spaces, func(i, j int) bool {
		return user.Spaces[i].Rate < user.Spaces[j].Rate
	})

	// 第2步：挑选热门数字
	summery, sNums := uint32(0), make(map[int32]struct{})
	for _, space := range user.Spaces {
		if summery > 800 {
			continue
		}

		sNums[space.Num] = struct{}{}
		summery = summery + base.STDS1000[space.Num]
	}

	sBets, total := make([]string, 0), uint32(0)
	for _, num := range base.SN28 {
		iGold := uint32(base.Config.Base * float64(base.STDS1000[num]) / 1000)
		if _, ok := sNums[num]; ok && iGold > 0 {
			total = total + iGold
			sBets = append(sBets, fmt.Sprintf("%d", iGold))
			log.Printf("\t竞猜数字【%02d】：【%6d】；\n", num, iGold)
		} else {
			sBets = append(sBets, "0")
		}
	}
	nextIssue := fmt.Sprintf("%d", user.Issue+1)
	log.Printf("【2】投注期数【%s】，投注基数【%d】，投注总额【%d】...\n", nextIssue, int(base.Config.Base), total)

	// 第3步：执行投注
	time.Sleep(time.Duration(15*rand.Float64()) * time.Second)
	log.Printf("【3】执行投注 ...")
	if err := qRecord(nextIssue, sBets); err != nil {
		log.Printf("【ERR-3】: %s \n", err.Error())
		return
	}

	log.Println("【9】投注成功，全部执行结束 ...")
}
