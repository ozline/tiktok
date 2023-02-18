package main

type WebMessage struct {
	FromUserId int64
	ToUserId   int64
	Content    string
	CreateTime string
	SeqID      int64
}

type WebServerResponse struct {
	Status bool
	AckID  int64
}
