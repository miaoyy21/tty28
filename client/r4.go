package client

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

func r4Fn(db *sql.DB, portBetting string, issue int, users []*User, rds map[int32]float64, dz float64, ns int) {
	var wg sync.WaitGroup

	wg.Add(len(users))
	for _, user := range users {
		go func(user *User) {
			defer wg.Done()

			m1Gold := ofM1Gold(user.Gold)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			bets, nums := make([]string, 0, 27), 0
			for _, n := range SN28 {
				if rds[n] < dz {
					bets = append(bets, "0")
					continue
				}

				//var sig float64
				//if rds[n] > 1.0 {
				//	sig = rds[n]
				//} else {
				//	sig = (rds[n] - user.Sigma) / (1.0 - user.Sigma)
				//}

				sig := 1.0
				fGold := sig * float64(m1Gold) * float64(STDS1000[n]) / 1000

				iGold := ofGold(fGold)
				if iGold == 0 {
					bets = append(bets, "0")
					continue
				}

				nums++
				bets = append(bets, strconv.Itoa(iGold))
			}

			if nums == 0 || nums == 28 {
				log.Printf("(4) 托管账户 %q 没有符合条件的投注数字 ... \n", user.UserName)
				return
			}

			if err := gBetting(
				net.JoinHostPort(user.Host, portBetting), fmt.Sprintf("%d", issue+1), strings.Join(bets, ","),
				user.UToken, user.SecChUa, user.SecChUaPlatform, user.UserAgent, ns); err != nil {
				log.Printf("【ERR-41】:【%s】 %s \n", user.UserName, err)

				if _, err := db.Exec("UPDATE user SET msg = ? WHERE user_id = ?", err.Error(), user.UserId); err != nil {
					log.Printf("【ERR-42】: 【%s】%s \n", user.UserName, err)
					return
				}

				return
			}

			log.Printf("(4) 托管账户 %q 执行投注成功 ... \n", user.UserName)
			if _, err := db.Exec("UPDATE user SET msg = ? WHERE user_id = ?", "Success", user.UserId); err != nil {
				log.Printf("【ERR-43】: [%s] %s \n", user.UserName, err)
				return
			}
		}(user)
	}

	wg.Wait()
}
