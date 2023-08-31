package eslogrus

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	elastic "github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/sirupsen/logrus"
)

// ElasticHook is a logrus
// hook for ElasticSearch
type ElasticHook struct {
	client    *elastic.Client    // es的客户端
	host      string             // es 的 host
	index     IndexNameFunc      // 获取索引的名字
	levels    []logrus.Level     // 日志的级别 info，error
	ctx       context.Context    // 上下文
	ctxCancel context.CancelFunc // 上下文cancel的函数，
	fireFunc  fireFunc           // 需要实现这个
}

// 发送到es的信息结构
type message struct {
	Host      string
	Timestamp string `json:"@timestamp"`
	Message   string
	Data      logrus.Fields
	Level     string
}

// IndexNameFunc get index name
type IndexNameFunc func() string

type fireFunc func(entry *logrus.Entry, hook *ElasticHook) error

// NewElasticHook 新建一个es hook对象
func NewElasticHook(client *elastic.Client, host string, level logrus.Level, index string) (*ElasticHook, error) {
	return NewElasticHookWithFunc(client, host, level, func() string { return index })
}

func NewElasticHookWithFunc(client *elastic.Client, host string, level logrus.Level, indexFunc IndexNameFunc) (*ElasticHook, error) {
	return newHookFuncAndFireFunc(client, host, level, indexFunc, syncFireFunc)
}

// 新建一个hook
func newHookFuncAndFireFunc(client *elastic.Client, host string, level logrus.Level, indexFunc IndexNameFunc, fireFunc fireFunc) (*ElasticHook, error) {
	var levels []logrus.Level
	for _, l := range []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	} {
		if l <= level {
			levels = append(levels, l)
		}
	}

	ctx, cancel := context.WithCancel(context.TODO())

	return &ElasticHook{
		client:    client,
		host:      host,
		index:     indexFunc,
		levels:    levels,
		ctx:       ctx,
		ctxCancel: cancel,
		fireFunc:  fireFunc,
	}, nil
}

// createMessage 创建信息
func createMessage(entry *logrus.Entry, hook *ElasticHook) *message {
	level := entry.Level.String()

	if e, ok := entry.Data[logrus.ErrorKey]; ok && e != nil {
		if err, ok := e.(error); ok {
			entry.Data[logrus.ErrorKey] = err.Error()
		}
	}

	return &message{
		hook.host,
		entry.Time.UTC().Format(time.RFC3339Nano),
		entry.Message,
		entry.Data,
		strings.ToUpper(level),
	}
}

// syncFireFunc 异步发送
func syncFireFunc(entry *logrus.Entry, hook *ElasticHook) error {
	data, err := json.Marshal(createMessage(entry, hook))

	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:   hook.index(),
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}

	res, err := req.Do(hook.ctx, hook.client)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
	} else {
		// Print the response status and indexed document version.
		log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
	}
	return err
}

// Fire 实现 logrus hook 必须要的接口函数
func (hook *ElasticHook) Fire(entry *logrus.Entry) error {
	return hook.fireFunc(entry, hook)
}

// Levels 实现 logrus hook 必须要的接口函数
func (hook *ElasticHook) Levels() []logrus.Level {
	return hook.levels
}
