package client

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

func run(db *sql.DB, portGold, portBetting string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	log.Println("//*********************************** 定时任务开始执行 ***********************************//")

	// 第一步 查询本账号的最新期数
	sleepTo(30.0 + 5*rand.Float64())
	log.Println("<一> 查询本账号的最新期数 >>> ")

	issue, total, err := qIssueGold()
	if err != nil {
		log.Printf("【ERR-11】: %s \n", err)
		return
	}

	mrx := 1.0
	if total < 1<<27 {
		mrx = float64(total) / float64(1<<27) // 134,217,728
	}
	log.Printf("最新开奖期数【%d】，总值系数【%.2f】 ... \n", issue, mrx)

	// 第二步 查询托管账户的金额
	sleepTo(40.0 + 5*rand.Float64())
	log.Println("<二> 查询托管账户的资金余额 >>> ")

	users, err := dQueryUsers(db)
	if err != nil {
		log.Printf("【ERR-21】: %s \n", err)
		return
	}

	for _, user := range users {
		gold, err := gGold(net.JoinHostPort(user.Host, portGold), user.Cookie, user.UserAgent, user.Unix, user.KeyCode, user.DeviceId, user.UserId, user.Token)
		if err != nil {
			log.Printf("【ERR-22】: [%s] %s \n", user.UserId, err)
			return
		}

		user.Gold = gold

		// Update User's Gold
		if _, err := db.Exec("UPDATE user SET gold = ?, update_at = ? WHERE user_id = ?", gold, time.Now().Format("2006-01-02 15:04"), user.UserId); err != nil {
			log.Printf("【ERR-23】: [%s] %s \n", user.UserId, err)
			return
		}

		// Insert User's Log
		if _, err := db.Exec("INSERT INTO user_log(user_id, time_at, gold) VALUES (?,?,?)", user.UserId, time.Now().Format("2006-01-02 15:04"), gold); err != nil {
			log.Printf("【ERR-24】: [%s] %s \n", user.UserId, err)
			return
		}

		log.Printf("托管账户%q的资金余额 %d ... \n", user.UserName, user.Gold)
	}

	// 第三步 查询本账户的权重值
	sleepTo(54.0)
	log.Println("<三> 查询本账户的权重值 >>> ")

	rds, err := qRiddle(fmt.Sprintf("%d", issue+1))
	if err != nil {
		log.Printf("【ERR-31】: %s \n", err)
		return
	}

	// 输出权重值
	for _, n := range SN28 {
		sigma := 0.985
		if rds[n] <= sigma {
			log.Printf("竞猜数字【 %02d - 】； \n", n)
			continue
		}

		var sig float64
		if rds[n] > 1.0 {
			sig = rds[n]
			log.Printf("竞猜数字【 %02d H 】，权重值【%.2f】； \n", n, sig)
		} else {
			sig = (rds[n] - sigma) / (1.0 - sigma)
			log.Printf("竞猜数字【 %02d L 】，权重值【%.2f】； \n", n, sig)
		}
	}

	// 第四步 委托账户投注
	var wg sync.WaitGroup

	wg.Add(len(users))
	log.Println("<四> 执行托管账户投注 >>> ")

	for _, user := range users {
		user := user
		m1Gold := ofM1Gold(user.Gold)
		go func() {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			bets := make(map[int32]int32)
			for _, n := range SN28 {
				if rds[n] <= user.Sigma {
					continue
				}

				var sig float64
				if rds[n] > 1.0 {
					sig = rds[n]
				} else {
					sig = (rds[n] - user.Sigma) / (1.0 - user.Sigma)
				}

				fGold := mrx * sig * float64(m1Gold) * float64(STDS1000[n]) / 1000
				iGold := ofGold(fGold)
				if iGold > 0 {
					bets[n] = iGold
				}
			}

			log.Printf("托管账户 %q 执行投注成功 ... \n", user.UserName)
			if err := gBetting(net.JoinHostPort(user.Host, portBetting), fmt.Sprintf("%d", issue+1), bets,
				user.Cookie, user.UserAgent, user.Unix, user.KeyCode, user.DeviceId, user.UserId, user.Token); err != nil {
				log.Printf("【ERR-41】:【%s】 %s \n", user.UserName, err)

				if _, err := db.Exec("UPDATE user SET msg = ? WHERE user_id = ?", err.Error(), user.UserId); err != nil {
					log.Printf("【ERR-42】: 【%s】%s \n", user.UserName, err)
					return
				}

				return
			}

			if _, err := db.Exec("UPDATE user SET msg = ? WHERE user_id = ?", "OK", user.UserId); err != nil {
				log.Printf("【ERR-43】: [%s] %s \n", user.UserName, err)
				return
			}
		}()
	}

	wg.Wait()
	log.Println("全部执行结束 ...")
}
