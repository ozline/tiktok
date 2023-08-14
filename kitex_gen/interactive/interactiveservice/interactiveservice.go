// Code generated by Kitex v0.6.2. DO NOT EDIT.

package interactiveservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	interactive "github.com/ozline/tiktok/kitex_gen/interactive"
)

func serviceInfo() *kitex.ServiceInfo {
	return interactiveServiceServiceInfo
}

var interactiveServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "InteractiveService"
	handlerType := (*interactive.InteractiveService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction": kitex.NewMethodInfo(favoriteActionHandler, newInteractiveServiceFavoriteActionArgs, newInteractiveServiceFavoriteActionResult, false),
		"FavoriteList":   kitex.NewMethodInfo(favoriteListHandler, newInteractiveServiceFavoriteListArgs, newInteractiveServiceFavoriteListResult, false),
		"CommentAction":  kitex.NewMethodInfo(commentActionHandler, newInteractiveServiceCommentActionArgs, newInteractiveServiceCommentActionResult, false),
		"CommentList":    kitex.NewMethodInfo(commentListHandler, newInteractiveServiceCommentListArgs, newInteractiveServiceCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "interactive",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interactive.InteractiveServiceFavoriteActionArgs)
	realResult := result.(*interactive.InteractiveServiceFavoriteActionResult)
	success, err := handler.(interactive.InteractiveService).FavoriteAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractiveServiceFavoriteActionArgs() interface{} {
	return interactive.NewInteractiveServiceFavoriteActionArgs()
}

func newInteractiveServiceFavoriteActionResult() interface{} {
	return interactive.NewInteractiveServiceFavoriteActionResult()
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interactive.InteractiveServiceFavoriteListArgs)
	realResult := result.(*interactive.InteractiveServiceFavoriteListResult)
	success, err := handler.(interactive.InteractiveService).FavoriteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractiveServiceFavoriteListArgs() interface{} {
	return interactive.NewInteractiveServiceFavoriteListArgs()
}

func newInteractiveServiceFavoriteListResult() interface{} {
	return interactive.NewInteractiveServiceFavoriteListResult()
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interactive.InteractiveServiceCommentActionArgs)
	realResult := result.(*interactive.InteractiveServiceCommentActionResult)
	success, err := handler.(interactive.InteractiveService).CommentAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractiveServiceCommentActionArgs() interface{} {
	return interactive.NewInteractiveServiceCommentActionArgs()
}

func newInteractiveServiceCommentActionResult() interface{} {
	return interactive.NewInteractiveServiceCommentActionResult()
}

func commentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interactive.InteractiveServiceCommentListArgs)
	realResult := result.(*interactive.InteractiveServiceCommentListResult)
	success, err := handler.(interactive.InteractiveService).CommentList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractiveServiceCommentListArgs() interface{} {
	return interactive.NewInteractiveServiceCommentListArgs()
}

func newInteractiveServiceCommentListResult() interface{} {
	return interactive.NewInteractiveServiceCommentListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteAction(ctx context.Context, req *interactive.FavoriteActionRequest) (r *interactive.FavoriteActionResponse, err error) {
	var _args interactive.InteractiveServiceFavoriteActionArgs
	_args.Req = req
	var _result interactive.InteractiveServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, req *interactive.FavoriteListRequest) (r *interactive.FavoriteListResponse, err error) {
	var _args interactive.InteractiveServiceFavoriteListArgs
	_args.Req = req
	var _result interactive.InteractiveServiceFavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentAction(ctx context.Context, req *interactive.CommentActionRequest) (r *interactive.CommentActionResponse, err error) {
	var _args interactive.InteractiveServiceCommentActionArgs
	_args.Req = req
	var _result interactive.InteractiveServiceCommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentList(ctx context.Context, req *interactive.CommentListRequest) (r *interactive.CommentListResponse, err error) {
	var _args interactive.InteractiveServiceCommentListArgs
	_args.Req = req
	var _result interactive.InteractiveServiceCommentListResult
	if err = p.c.Call(ctx, "CommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
