package client

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"time"
)

func run(db *sql.DB, portGold, portBetting string) {
	log.Printf("//*********************************** 定时任务开始执行 %s ***********************************// \n", time.Now().Format("2006-01-02 15:04"))

	go run0(db, portGold, portBetting, 0)

	time.Sleep(30)
	run0(db, portGold, portBetting, 30)

	log.Println("<<<<<< 全部执行结束 >>>>>>")
}

func run0(db *sql.DB, portGold, portBetting string, delta float64) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	ns := time.Now().Nanosecond()

	// 第一步 查询本账号的最新期数
	sleepTo(delta + 5 + 5*rand.Float64())

	log.Println()
	issue, err := r1Fn(ns)
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
	sleepTo(delta + 26.25)

	dz := 1.0
	rds, dev, err := r3Fn(issue, dz, ns)
	if err != nil {
		log.Printf("【ERR-3】: %s", err.Error())
		return
	}

	if dev < 1.0 {
		log.Println("平均方差太低，不进行投注 ... ")
		return
	}

	// 第四步 委托账户投注
	r4Fn(db, portBetting, issue, users, rds, dz, ns)

	log.Println()
}
