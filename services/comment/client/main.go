package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	kitexclient "github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	"github.com/ozline/tiktok/services/comment/kitex_gen/tiktok/comment"
	"github.com/ozline/tiktok/services/comment/kitex_gen/tiktok/comment/tiktokcommentservice"
)

var client tiktokcommentservice.Client

func init() {
	snowflake.NewSnowflake(0, 0)

	var err error
	client, err = tiktokcommentservice.NewClient("kitex-test", kitexclient.WithHostPorts(constants.CommentServiceListenAddress))
	if err != nil {
		panic(err)
	}
}

func main() {
	// resp, _ := client.Post(context.Background(), &comment.PostReq{
	// 	Uid:     2,
	// 	Vid:     2,
	// 	Content: "hello",
	// })
	// printWithJSON(resp)

	// resp, _ := client.SetLike(context.Background(), &comment.LikeReq{
	// 	Uid:       2,
	// 	CommentId: 413656653972373504,
	// 	IsLike:    true,
	// })
	// printWithJSON(resp)

	// resp, _ := client.GetLike(context.Background(), &comment.LikeReq{
	// 	Uid:       3,
	// 	CommentId: 413625386241359872,
	// })
	// printWithJSON(resp)

	resp, _ := client.List(context.Background(), &comment.ListReq{
		Uid: 2,
		// Vid: 2,
		Type: comment.ListType_like,
		// Type: comment.ListType_video,
		// Type:       comment.ListType_comment,
		PageSize:   1,
		PageNumber: 2,
	})
	printWithJSON(resp)
}

func printWithJSON(v interface{}) {
	bs, _ := json.Marshal(v)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Println("\n\n", out.String())
}
