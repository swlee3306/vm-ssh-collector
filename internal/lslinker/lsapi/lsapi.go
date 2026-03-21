package lsapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Logstash struct {
	env struct {
		urls []string
	}

	_url_idx int
}

func (r *Logstash) Init(urls string) {
	r.env.urls = strings.Split(strings.ReplaceAll(strings.Trim(urls, "{[]}"), " ", ""), ",")
	r._url_idx = 0
}

func (r *Logstash) Send(msg interface{}) (err error) {
	client := &http.Client{Timeout: 10 * time.Second}
	for i := 0; i < len(r.env.urls); i++ {
		sndIdx := (r._url_idx + i) % len(r.env.urls)
		payload, marshalErr := json.Marshal(msg)
		if marshalErr != nil {
			return fmt.Errorf("marshal logstash payload: %w", marshalErr)
		}

		req, reqErr := http.NewRequest(http.MethodPost, r.env.urls[sndIdx], bytes.NewReader(payload))
		if reqErr != nil {
			err = reqErr
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		resp, doErr := client.Do(req)
		if doErr != nil {
			err = doErr
			continue
		}
		resp.Body.Close()
		if resp.StatusCode >= http.StatusBadRequest {
			err = fmt.Errorf("logstash returned status %d", resp.StatusCode)
			continue
		}

		r._url_idx = sndIdx
		return nil
	}

	return err
}
