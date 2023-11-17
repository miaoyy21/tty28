package client

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"
)

func r2Fn(db *sql.DB, portGold string, ns int) ([]*User, error) {

	users, err := dQueryUsers(db)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		gold, err := gGold(net.JoinHostPort(user.Host, portGold), user.UToken, user.SecChUa, user.SecChUaPlatform, user.UserAgent, ns)
		if err != nil {
			return nil, fmt.Errorf("[%s %s] %s", user.UserId, user.UserName, err.Error())
		}

		user.Gold = gold

		// Update User's Gold
		if _, err := db.Exec("UPDATE user SET gold = ?, update_at = ? WHERE user_id = ?", gold, time.Now().Format("2006-01-02 15:04"), user.UserId); err != nil {
			return nil, fmt.Errorf("[%s %s] %s \n", user.UserId, user.UserName, err.Error())
		}

		// Insert User's Log
		if _, err := db.Exec("INSERT INTO user_log(user_id, time_at, gold) VALUES (?,?,?)", user.UserId, time.Now().Format("2006-01-02 15:04"), gold); err != nil {
			return nil, fmt.Errorf("[%s %s] %s \n", user.UserId, user.UserName, err.Error())
		}

		log.Printf("(2) 托管账户%q的资金余额 %d ... \n", user.UserName, user.Gold)
	}

	return users, nil
}
