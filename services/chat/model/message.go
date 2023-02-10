package model

type Message struct {
	ID           int64  // 消息id
	To_user_id   int64  // 该消息接收者的id
	From_user_id int64  // 该消息发送者的id
	Content      string // 消息内容
	Create_time  string // 消息创建时间
}
