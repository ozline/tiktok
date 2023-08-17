package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ozline/tiktok/cmd/interactive/dal/sensitive_words"
	"io"
	"net/http"
)

type ResponseBody struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	IsPass bool   `json:"isPass"`
}

func (s *CommentService) MatchSensitiveWords(text string) (bool, error) {

	ok := sensitive_words.St.Match(text)
	if ok {
		return ok, nil
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
	if err = json.Unmarshal([]byte(string(body)), &res); err != nil {
		return false, err
	}
	return res.IsPass, nil

}
