package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/ozline/tiktok/cmd/interaction/dal/sensitive_words"
)

type ResponseBody struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	IsPass bool   `json:"isPass"`
}

func (s *InteractionService) MatchSensitiveWords(text string) (bool, error) {
	fail := sensitive_words.St.Match(text)
	if fail {
		return fail, nil
	}

	requestBody := fmt.Sprintf(`{"ColaKey":"enB54tE5dI55Qv169215005005221MuejeU6z","wordStr": "%s","isStrict":"1"}`, text)
	var jsonStr = []byte(requestBody)
	url := "https://luckycola.com.cn/tools/sensiWords"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var res ResponseBody
	if err = json.Unmarshal(body, &res); err != nil {
		return false, err
	}
	if res.Code != 0 {
		klog.Warnf("sensitive-words api error : %v", res.Msg)
		return false, nil
	}
	return !res.IsPass, nil
}
