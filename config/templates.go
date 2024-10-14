package config

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type RequestTemplate struct {
	Method  string
	URL     string
	Cookie  string
	Headers map[string]string
}

var Templates = make(map[string]RequestTemplate)

func LoadTemplates() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".tpl") {
			return nil
		}

		tempName := strings.Split(info.Name(), ".tpl")[0]
		bs, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		reqTemp := RequestTemplate{
			Method:  "GET",
			Headers: make(map[string]string),
		}

		lines := strings.Split(string(bs), "\n")
		for _, line := range lines {
			line := strings.TrimSpace(line)
			if len(line) < 1 {
				continue
			}

			fields := strings.SplitN(line, " ", 2)
			symbol, head := fields[0], fields[1]

			head = strings.TrimSpace(strings.TrimRight(head, "\\"))
			head = strings.TrimRight(strings.TrimLeft(head, "'"), "'")
			switch symbol {
			case "curl":
				reqTemp.URL = head
			case "-H":
				kv := strings.SplitN(head, ": ", 2)

				k, v := kv[0], kv[1]
				switch k {
				case "cookie":
					// Cookie由外部提供
					reqTemp.Cookie = v
				default:
					reqTemp.Headers[k] = v
				}
			case "--data-raw":
				reqTemp.Method = "POST"
			default:
				return fmt.Errorf("invalid Line %s", line)
			}
		}

		Templates[tempName] = reqTemp
		log.Printf("加载模版	%q	%#v", tempName, reqTemp)

		return nil
	}); err != nil {
		return err
	}

	return nil
}
