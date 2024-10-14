package huiwan28

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Cookie string  `json:"cookie"`
	Base   float64 `json:"base"`
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
