package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth"
	"github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth/tiktokauthservice"

	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := tiktokauthservice.NewClient("kitex-test", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	// Get Token
	req1 := &auth.GetTokenRequest{
		Username: "ozline",
		UserId:   1015901853,
	}
	resp1, err := client.GetToken(context.Background(), req1)
	if err != nil {
		log.Fatal("err", err.Error())
	}

	token := resp1.Token
	printWithJSON(resp1)

	req2 := &auth.CheckTokenRequest{
		Token: token,
	}
	resp2, err := client.CheckToken(context.Background(), req2)
	if err != nil {
		log.Fatal("err", err.Error())
	}
	printWithJSON(resp2)
}

func printWithJSON(v interface{}) {
	bs, _ := json.Marshal(v)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Println("\n\n", out.String())
}
