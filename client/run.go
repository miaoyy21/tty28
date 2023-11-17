package client

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"time"
)

func run60(db *sql.DB, portGold, portBetting string, delta float64) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	ns := time.Now().Nanosecond()
	log.Println("//*********************************** 定时任务开始执行 ***********************************//")

	// 第一步 查询本账号的最新期数
	sleepTo(delta + 5*rand.Float64())
	issue, mrx, err := r1Fn(ns)
	if err != nil {
		log.Printf("【ERR-1】: %s", err.Error())
		return
	}

	// 第二步 查询托管账户的金额
	sleepTo(delta + 10 + 5*rand.Float64())
	users, err := r2Fn(db, portGold, ns)
	if err != nil {
		log.Printf("【ERR-2】: %s", err.Error())
		return
	}

	// 第三步 查询本账户的权重值
	sleepTo(delta + 26.5)
	rds, err := r3Fn(issue, ns)
	if err != nil {
		log.Printf("【ERR-3】: %s", err.Error())
		return
	}

	// 第四步 委托账户投注
	r4Fn(db, portBetting, issue, users, mrx, rds, ns)

	log.Println("全部执行结束 ...")
}
