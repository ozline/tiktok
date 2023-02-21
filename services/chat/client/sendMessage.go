package main

import (
	"fmt"
)

func main() {
	fmt.Println("----- Chat Client -----")
	userId := 1
	// client, err := tiktokchatservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8891"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //go receive_message_by_id(int64(userId), seqId)
	// startTime := time.Now().UnixMilli()

	// request := &chat.SendMessageRequest{
	// 	FromUserId: int64(userId),
	// 	ToUserId:   2,
	// 	Content:    "Hello 2,This is 1",
	// }

	// response, err := client.SendChatMessage(context.Background(), request)
	// if err != nil {
	// 	log.Fatal("error", err.Error())
	// }
	// endTime := time.Now().UnixMilli()
	// if response.StatusMsg == "Success message" {
	// 	fmt.Println("----- Success To Send Message,Use ", endTime-startTime, " ms -----")
	// } else {
	// 	fmt.Println("----- Failure To Send Message -----")
	// }

	// userId = 1
	receive_message_by_id(int64(userId))

}

func receive_message_by_id(userId int64) {
	// ticker := time.NewTicker(5 * time.Second)
	// defer ticker.Stop()

	// for range ticker.C {
	// 	client, err := tiktokchatservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8891"))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	request := &chat.DouyinReceiveMessageRequest{
	// 		ToUserId: userId,
	// 	}
	// 	response, _ := client.AcceptChatMessage(context.Background(), request)
	// 	if len(response.ToUserIds) > 0 {
	// 		fmt.Println("----- User ", userId, "Receive ", len(response.ToUserIds), " Messages -----")
	// 		for index, _ := range response.ToUserIds {
	// 			fmt.Println("Message", index+1, ":", response.FromUserIds[index], "->", response.ToUserIds[index], "——", response.Contents[index])
	// 		}
	// 	}
	// }
}
