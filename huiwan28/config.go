package huiwan28

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Authority       string `json:"authority"`
	Origin          string `json:"origin"`
	Referer         string `json:"referer"`
	SecChUa         string `json:"sec_ch_ua"`
	SecChUaPlatform string `json:"sec_ch_ua_platform"`
	UserAgent       string `json:"user_agent"`
	UToken          string `json:"u_token"`
	Style           string `json:"style"`

	IssueURL   string `json:"issue_url"`
	RiddleURL  string `json:"riddle_url"`
	GoldURL    string `json:"gold_url"`
	BettingURL string `json:"betting_url"`
}

var conf Config

func InitConfig() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	bs, err := os.ReadFile(filepath.Join(dir, "config/huiwan28.json"))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bs, &conf); err != nil {
		return err
	}

	return nil
}
