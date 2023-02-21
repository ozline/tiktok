// Code generated by Kitex v0.4.4. DO NOT EDIT.

package tiktokchatservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	chat "github.com/ozline/tiktok/kitex_gen/tiktok/chat"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return tiktokChatServiceServiceInfo
}

var tiktokChatServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "tiktokChatService"
	handlerType := (*chat.TiktokChatService)(nil)
	methods := map[string]kitex.MethodInfo{
		"SendChatMessage":   kitex.NewMethodInfo(sendChatMessageHandler, newSendChatMessageArgs, newSendChatMessageResult, false),
		"AcceptChatMessage": kitex.NewMethodInfo(acceptChatMessageHandler, newAcceptChatMessageArgs, newAcceptChatMessageResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "chat",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func sendChatMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(chat.SendMessageRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(chat.TiktokChatService).SendChatMessage(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SendChatMessageArgs:
		success, err := handler.(chat.TiktokChatService).SendChatMessage(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SendChatMessageResult)
		realResult.Success = success
	}
	return nil
}
func newSendChatMessageArgs() interface{} {
	return &SendChatMessageArgs{}
}

func newSendChatMessageResult() interface{} {
	return &SendChatMessageResult{}
}

type SendChatMessageArgs struct {
	Req *chat.SendMessageRequest
}

func (p *SendChatMessageArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(chat.SendMessageRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SendChatMessageArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SendChatMessageArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SendChatMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SendChatMessageArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SendChatMessageArgs) Unmarshal(in []byte) error {
	msg := new(chat.SendMessageRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SendChatMessageArgs_Req_DEFAULT *chat.SendMessageRequest

func (p *SendChatMessageArgs) GetReq() *chat.SendMessageRequest {
	if !p.IsSetReq() {
		return SendChatMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendChatMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type SendChatMessageResult struct {
	Success *chat.SendMessageResponse
}

var SendChatMessageResult_Success_DEFAULT *chat.SendMessageResponse

func (p *SendChatMessageResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(chat.SendMessageResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SendChatMessageResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SendChatMessageResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SendChatMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SendChatMessageResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SendChatMessageResult) Unmarshal(in []byte) error {
	msg := new(chat.SendMessageResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendChatMessageResult) GetSuccess() *chat.SendMessageResponse {
	if !p.IsSetSuccess() {
		return SendChatMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendChatMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.SendMessageResponse)
}

func (p *SendChatMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func acceptChatMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(chat.ReceiveMessageRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(chat.TiktokChatService).AcceptChatMessage(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *AcceptChatMessageArgs:
		success, err := handler.(chat.TiktokChatService).AcceptChatMessage(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AcceptChatMessageResult)
		realResult.Success = success
	}
	return nil
}
func newAcceptChatMessageArgs() interface{} {
	return &AcceptChatMessageArgs{}
}

func newAcceptChatMessageResult() interface{} {
	return &AcceptChatMessageResult{}
}

type AcceptChatMessageArgs struct {
	Req *chat.ReceiveMessageRequest
}

func (p *AcceptChatMessageArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(chat.ReceiveMessageRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AcceptChatMessageArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AcceptChatMessageArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AcceptChatMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AcceptChatMessageArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *AcceptChatMessageArgs) Unmarshal(in []byte) error {
	msg := new(chat.ReceiveMessageRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AcceptChatMessageArgs_Req_DEFAULT *chat.ReceiveMessageRequest

func (p *AcceptChatMessageArgs) GetReq() *chat.ReceiveMessageRequest {
	if !p.IsSetReq() {
		return AcceptChatMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AcceptChatMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

type AcceptChatMessageResult struct {
	Success *chat.ReceiveMessageResponse
}

var AcceptChatMessageResult_Success_DEFAULT *chat.ReceiveMessageResponse

func (p *AcceptChatMessageResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(chat.ReceiveMessageResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AcceptChatMessageResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AcceptChatMessageResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AcceptChatMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AcceptChatMessageResult")
	}
	return proto.Marshal(p.Success)
}

func (p *AcceptChatMessageResult) Unmarshal(in []byte) error {
	msg := new(chat.ReceiveMessageResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AcceptChatMessageResult) GetSuccess() *chat.ReceiveMessageResponse {
	if !p.IsSetSuccess() {
		return AcceptChatMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AcceptChatMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*chat.ReceiveMessageResponse)
}

func (p *AcceptChatMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SendChatMessage(ctx context.Context, Req *chat.SendMessageRequest) (r *chat.SendMessageResponse, err error) {
	var _args SendChatMessageArgs
	_args.Req = Req
	var _result SendChatMessageResult
	if err = p.c.Call(ctx, "SendChatMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AcceptChatMessage(ctx context.Context, Req *chat.ReceiveMessageRequest) (r *chat.ReceiveMessageResponse, err error) {
	var _args AcceptChatMessageArgs
	_args.Req = Req
	var _result AcceptChatMessageResult
	if err = p.c.Call(ctx, "AcceptChatMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}