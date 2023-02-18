package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
	//"golang.org/x/net/websocket"
	"log"
)

var seqid int64
var RDB *redis.Client

func main() {
	seqid = 1
	url := "ws://localhost:8888/ws" //服务器地址
	userInfo := make(map[string][]string)
	userInfo["UserId"] = append(userInfo["UserId"], "1")
	userInfo["SeqID"] = append(userInfo["SeqId"], string(seqid))
	ws, _, err := websocket.DefaultDialer.Dial(url, userInfo)
	if err != nil {
		log.Fatal(err)
	}

	message := WebMessage{
		FromUserId: 1,
		ToUserId:   2,
		Content:    "Hello Listen,This is Send",
		CreateTime: strconv.FormatInt(time.Now().UnixMicro(), 10), // 毫秒
		SeqID:      seqid,
	}
	seqid++
	str, _ := json.Marshal(message)
	RDB.Set(string(seqid), str, 0).Err()

	go sendMessage(string(str), ws)
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		var response ServerAckResponse
		json.Unmarshal(data, &response)

		if response.Status == true && response.AckID == seqid {
			RDB.FlushDB().Result()
			continue
		} else {
			go reSendMessage(response.AckID, seqid, ws)
		}
		time.Sleep(1 * time.Second)
	}
}

func sendMessage(message string, ws *websocket.Conn) {
	for {

		err := ws.WriteMessage(websocket.BinaryMessage, []byte(message))

		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 2)
	}
}

func reSendMessage(ackid int64, seqid int64, ws *websocket.Conn) {
	for id := ackid; id <= seqid; id++ {
		message, _ := RDB.Get(string(ackid)).Result()
		go sendMessage(message, ws)
	}
}
