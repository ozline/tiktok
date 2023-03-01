package rpc

import (
	"context"

	"github.com/ozline/tiktok/pkg/errno"

	"github.com/ozline/tiktok/kitex_gen/tiktok/chat"
)

func SendChatMessage(ctx context.Context, req *chat.SendMessageRequest) error {
	resp, err := chatClient.SendChatMessage(ctx, req)

	if err != nil {
		return err
	}

	if resp.Base.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func GetChatMessage(ctx context.Context, req *chat.ReceiveMessageRequest) ([]*chat.ChatMsg, error) {

	resp, err := chatClient.AcceptChatMessage(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(int64(resp.Base.Code), resp.Base.Msg)
	}

	return resp.Data, nil
}
