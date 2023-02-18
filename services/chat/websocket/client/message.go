package main

type WebMessage struct {
	FromUserId int64
	ToUserId   int64
	Content    string
	CreateTime string
	SeqID      int64
}

type ServerAckResponse struct {
	Status bool
	AckID  int64
}

type ClientAckResponse struct {
	Status bool
	AckID  int64
	UserID int64
}
