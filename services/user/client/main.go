package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user"
	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user/tiktokuserservice"
	"log"
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
	//registerRequest := &user.DouyinUserRegisterRequest{
	//	Username: "test",
	//	Password: "123456",
	//}
	//registerResponse, err := client.Register(context.Background(), registerRequest)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//log.Println(registerResponse)
	//--------------------登录测试-------------------------
	//loginRequest := &user.DouyinUserLoginRequest{
	//	Username: "test",
	//	Password: "123456",
	//}
	//loginResponse, err := client.Login(context.Background(), loginRequest)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//log.Println(loginResponse)
	//--------------------获取用户信息测试-------------------------
	userRequest := &user.DouyinUserRequest{
		UserId: 3974384737,
		Token:  "",
	}
	userResponse, err := client.Info(context.Background(), userRequest)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(userResponse)
}
