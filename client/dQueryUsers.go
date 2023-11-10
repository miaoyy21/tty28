package client

import (
	"database/sql"
)

type User struct {
	UserId   string
	UserName string
	Host     string
	Sigma    float64

	UToken          string
	SecChUa         string
	SecChUaPlatform string
	UserAgent       string

	Gold int64
}

func dQueryUsers(db *sql.DB) ([]*User, error) {
	query := `
		SELECT user_id,user_name, host, sigma, u_token, sec_ch_ua, sec_ch_ua_platform, user_agent,gold
		FROM user
		WHERE is_valid = 1
		ORDER BY gold DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		var userId, userName, host, uToken, secChUa, secChUaPlatform, userAgent string
		var sigma float64
		var gold int64
		if err := rows.Scan(&userId, &userName, &host, &sigma, &uToken, &secChUa, &secChUaPlatform, &userAgent, &gold); err != nil {
			return nil, err
		}

		user := &User{
			UserId:   userId,
			UserName: userName,
			Host:     host,
			Sigma:    sigma,

			UToken:          uToken,
			SecChUa:         secChUa,
			SecChUaPlatform: secChUaPlatform,
			UserAgent:       userAgent,

			Gold: gold,
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
