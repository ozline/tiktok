package rpc

import (
	"context"

	"github.com/ozline/tiktok/pkg/errno"

	"github.com/ozline/tiktok/kitex_gen/tiktok/chat"
)

func SendChatMessage(ctx context.Context, req *chat.SendMessageRequest) (bool, error) {
	resp, err := chatClient.SendChatMessage(ctx, req)

	if err != nil {
		return false, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return false, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return true, nil
}

func GetChatMessage(ctx context.Context, req *chat.ReceiveMessageRequest) (bool, error) {

	resp, err := chatClient.AcceptChatMessage(ctx, req)

	if err != nil {
		return false, err
	}

	if resp.StatusCode != errno.SuccessCode {
		return false, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}

	return true, nil
}
