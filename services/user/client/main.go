package main

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/kitex_gen/tiktok/user"
	"github.com/ozline/tiktok/kitex_gen/tiktok/user/tiktokuserservice"
)

func main() {
	//测试
	client, err := tiktokuserservice.NewClient("user-test", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	//--------------------运行测试-------------------------
	//req1 := &user.Request1{
	//	Message: "testPingPong",
	//}
	//resp, err := client.PingPong(context.Background(), req1)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//log.Println("PingPong Func Response", resp)
	//--------------------注册测试-------------------------
	//registerRequest := &user.UserRegisterRequest{
	//	Username: "test2",
	//	Password: "123456",
	//}
	//registerResponse, err := client.Register(context.Background(), registerRequest)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//log.Println(registerResponse)
	//--------------------登录测试-------------------------
	loginRequest := &user.UserLoginRequest{
		Username: "test",
		Password: "123456",
	}
	loginResponse, err := client.Login(context.Background(), loginRequest)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(loginResponse)
	//--------------------获取用户信息测试-------------------------
	//userRequest := &user.UserRequest{
	//	UserId: 411828220120268800,
	//	Token:  "",
	//}
	//userResponse, err := client.Info(context.Background(), userRequest)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//log.Println(userResponse)
	//--------------------获取token测试-------------------------
	//getTokenRequest := &user.GetTokenRequest{
	//	UserId:   411828220120268800,
	//	Username: "test",
	//}
	//getTokenResponse, err := client.GetToken(context.Background(), getTokenRequest)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//log.Println(getTokenResponse)
	//--------------------检查token测试-------------------------
	checkTokenRequest := &user.CheckTokenRequest{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0MTE4MjgyMjAxMjAyNjg4MDAsInVzZXJuYW1lIjoidGVzdCIsImV4cCI6MTY3NjEzNTU0NCwibmJmIjoxNjc2MTMxOTQ0fQ.14uZs0YOOO5eLzXjF6-DzADiz_TY28QR5QLBoKsmQTM",
	}
	checkTokenResponse, err := client.CheckToken(context.Background(), checkTokenRequest)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(checkTokenResponse)
}
