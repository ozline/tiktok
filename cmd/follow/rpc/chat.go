package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/kitex_gen/chat/messageservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/middleware"
)

func InitChatRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	c, err := messageservice.NewClient(
		constants.ChatServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithSuite(trace.NewDefaultClientSuite()), // 设置链路追踪
	)

	if err != nil {
		panic(err)
	}

	chatClient = c
}

func GetMessage(ctx context.Context, req *chat.MessageListRequest, uid, tid int64) (string, int64, error) {
	resp, err := chatClient.MessageList(ctx, req)

	if err != nil {
		return "", -1, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return "", -1, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	messageList := resp.MessageList
	len := len(messageList)

	for i := len - 1; i >= 0; i-- {
		switch {
		case messageList[i].FromUserId == uid && messageList[i].ToUserId == tid:
			return messageList[i].Content, 1, nil
		case messageList[i].FromUserId == tid && messageList[i].ToUserId == uid:
			return messageList[i].Content, 0, nil
		default:
			continue
		}
	}
	return "", -1, nil
}
