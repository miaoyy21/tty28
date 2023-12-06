package hdo

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Do(authority, origin, referer, secChUa, secChUaPlatform, userAgent string, url string, t interface{}) error {

	// Sync
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Sync Header
	req.Header.Set("authority", authority)
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("origin", origin)
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", referer)

	req.Header.Set("sec-ch-ua", secChUa)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", secChUaPlatform)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")

	req.Header.Set("user-agent", userAgent)

	// Response
	http.DefaultClient.Timeout = 5 * time.Second
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// JSON Decode
	if err := json.NewDecoder(resp.Body).Decode(t); err != nil {
		return err
	}

	return nil
}

func buildCookie(q string) string {
	qs := strings.Split(q, "; ")

	ns := make([]string, 0)
	for _, s := range qs {
		s0 := strings.Split(s, "=")
		k, v := s0[0], strings.Join(s0[1:], "=")

		nv := v
		if strings.Contains(k, "lpvt") {
			nv = strconv.Itoa(int(time.Now().Unix() - 1))
		}

		ns = append(ns, strings.Join([]string{k, nv}, "="))
	}

	return strings.Join(ns, "; ")
}
