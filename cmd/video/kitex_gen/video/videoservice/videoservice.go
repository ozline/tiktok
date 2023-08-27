// Code generated by Kitex v0.6.2. DO NOT EDIT.

package videoservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	video "github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed":                 kitex.NewMethodInfo(feedHandler, newFeedArgs, newFeedResult, false),
		"PutVideo":             kitex.NewMethodInfo(putVideoHandler, newPutVideoArgs, newPutVideoResult, false),
		"GetFavoriteVideoInfo": kitex.NewMethodInfo(getFavoriteVideoInfoHandler, newGetFavoriteVideoInfoArgs, newGetFavoriteVideoInfoResult, false),
		"GetPublishList":       kitex.NewMethodInfo(getPublishListHandler, newGetPublishListArgs, newGetPublishListResult, false),
		"GetWorkCount":         kitex.NewMethodInfo(getWorkCountHandler, newGetWorkCountArgs, newGetWorkCountResult, false),
		"GetVideoIDByUid":      kitex.NewMethodInfo(getVideoIDByUidHandler, newGetVideoIDByUidArgs, newGetVideoIDByUidResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.FeedRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).Feed(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FeedArgs:
		success, err := handler.(video.VideoService).Feed(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FeedResult)
		realResult.Success = success
	}
	return nil
}
func newFeedArgs() interface{} {
	return &FeedArgs{}
}

func newFeedResult() interface{} {
	return &FeedResult{}
}

type FeedArgs struct {
	Req *video.FeedRequest
}

func (p *FeedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.FeedRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FeedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FeedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FeedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FeedArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FeedArgs) Unmarshal(in []byte) error {
	msg := new(video.FeedRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FeedArgs_Req_DEFAULT *video.FeedRequest

func (p *FeedArgs) GetReq() *video.FeedRequest {
	if !p.IsSetReq() {
		return FeedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FeedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FeedArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FeedResult struct {
	Success *video.FeedResponse
}

var FeedResult_Success_DEFAULT *video.FeedResponse

func (p *FeedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.FeedResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FeedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FeedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FeedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FeedResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FeedResult) Unmarshal(in []byte) error {
	msg := new(video.FeedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FeedResult) GetSuccess() *video.FeedResponse {
	if !p.IsSetSuccess() {
		return FeedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.FeedResponse)
}

func (p *FeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FeedResult) GetResult() interface{} {
	return p.Success
}

func putVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.PutVideoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).PutVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PutVideoArgs:
		success, err := handler.(video.VideoService).PutVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PutVideoResult)
		realResult.Success = success
	}
	return nil
}
func newPutVideoArgs() interface{} {
	return &PutVideoArgs{}
}

func newPutVideoResult() interface{} {
	return &PutVideoResult{}
}

type PutVideoArgs struct {
	Req *video.PutVideoRequest
}

func (p *PutVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.PutVideoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PutVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PutVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PutVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PutVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PutVideoArgs) Unmarshal(in []byte) error {
	msg := new(video.PutVideoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PutVideoArgs_Req_DEFAULT *video.PutVideoRequest

func (p *PutVideoArgs) GetReq() *video.PutVideoRequest {
	if !p.IsSetReq() {
		return PutVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PutVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PutVideoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PutVideoResult struct {
	Success *video.PutVideoResponse
}

var PutVideoResult_Success_DEFAULT *video.PutVideoResponse

func (p *PutVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.PutVideoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PutVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PutVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PutVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PutVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PutVideoResult) Unmarshal(in []byte) error {
	msg := new(video.PutVideoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PutVideoResult) GetSuccess() *video.PutVideoResponse {
	if !p.IsSetSuccess() {
		return PutVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PutVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.PutVideoResponse)
}

func (p *PutVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PutVideoResult) GetResult() interface{} {
	return p.Success
}

func getFavoriteVideoInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.GetFavoriteVideoInfoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).GetFavoriteVideoInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFavoriteVideoInfoArgs:
		success, err := handler.(video.VideoService).GetFavoriteVideoInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFavoriteVideoInfoResult)
		realResult.Success = success
	}
	return nil
}
func newGetFavoriteVideoInfoArgs() interface{} {
	return &GetFavoriteVideoInfoArgs{}
}

func newGetFavoriteVideoInfoResult() interface{} {
	return &GetFavoriteVideoInfoResult{}
}

type GetFavoriteVideoInfoArgs struct {
	Req *video.GetFavoriteVideoInfoRequest
}

func (p *GetFavoriteVideoInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.GetFavoriteVideoInfoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetFavoriteVideoInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetFavoriteVideoInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetFavoriteVideoInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFavoriteVideoInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFavoriteVideoInfoArgs) Unmarshal(in []byte) error {
	msg := new(video.GetFavoriteVideoInfoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFavoriteVideoInfoArgs_Req_DEFAULT *video.GetFavoriteVideoInfoRequest

func (p *GetFavoriteVideoInfoArgs) GetReq() *video.GetFavoriteVideoInfoRequest {
	if !p.IsSetReq() {
		return GetFavoriteVideoInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFavoriteVideoInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetFavoriteVideoInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetFavoriteVideoInfoResult struct {
	Success *video.GetFavoriteVideoInfoResponse
}

var GetFavoriteVideoInfoResult_Success_DEFAULT *video.GetFavoriteVideoInfoResponse

func (p *GetFavoriteVideoInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.GetFavoriteVideoInfoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetFavoriteVideoInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetFavoriteVideoInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetFavoriteVideoInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFavoriteVideoInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFavoriteVideoInfoResult) Unmarshal(in []byte) error {
	msg := new(video.GetFavoriteVideoInfoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFavoriteVideoInfoResult) GetSuccess() *video.GetFavoriteVideoInfoResponse {
	if !p.IsSetSuccess() {
		return GetFavoriteVideoInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFavoriteVideoInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.GetFavoriteVideoInfoResponse)
}

func (p *GetFavoriteVideoInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetFavoriteVideoInfoResult) GetResult() interface{} {
	return p.Success
}

func getPublishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.GetPublishListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).GetPublishList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetPublishListArgs:
		success, err := handler.(video.VideoService).GetPublishList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPublishListResult)
		realResult.Success = success
	}
	return nil
}
func newGetPublishListArgs() interface{} {
	return &GetPublishListArgs{}
}

func newGetPublishListResult() interface{} {
	return &GetPublishListResult{}
}

type GetPublishListArgs struct {
	Req *video.GetPublishListRequest
}

func (p *GetPublishListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.GetPublishListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetPublishListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetPublishListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetPublishListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPublishListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetPublishListArgs) Unmarshal(in []byte) error {
	msg := new(video.GetPublishListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPublishListArgs_Req_DEFAULT *video.GetPublishListRequest

func (p *GetPublishListArgs) GetReq() *video.GetPublishListRequest {
	if !p.IsSetReq() {
		return GetPublishListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPublishListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetPublishListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetPublishListResult struct {
	Success *video.GetPublishListResponse
}

var GetPublishListResult_Success_DEFAULT *video.GetPublishListResponse

func (p *GetPublishListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.GetPublishListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetPublishListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetPublishListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetPublishListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPublishListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetPublishListResult) Unmarshal(in []byte) error {
	msg := new(video.GetPublishListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPublishListResult) GetSuccess() *video.GetPublishListResponse {
	if !p.IsSetSuccess() {
		return GetPublishListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPublishListResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.GetPublishListResponse)
}

func (p *GetPublishListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPublishListResult) GetResult() interface{} {
	return p.Success
}

func getWorkCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.GetWorkCountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).GetWorkCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetWorkCountArgs:
		success, err := handler.(video.VideoService).GetWorkCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetWorkCountResult)
		realResult.Success = success
	}
	return nil
}
func newGetWorkCountArgs() interface{} {
	return &GetWorkCountArgs{}
}

func newGetWorkCountResult() interface{} {
	return &GetWorkCountResult{}
}

type GetWorkCountArgs struct {
	Req *video.GetWorkCountRequest
}

func (p *GetWorkCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.GetWorkCountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetWorkCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetWorkCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetWorkCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetWorkCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetWorkCountArgs) Unmarshal(in []byte) error {
	msg := new(video.GetWorkCountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetWorkCountArgs_Req_DEFAULT *video.GetWorkCountRequest

func (p *GetWorkCountArgs) GetReq() *video.GetWorkCountRequest {
	if !p.IsSetReq() {
		return GetWorkCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetWorkCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetWorkCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetWorkCountResult struct {
	Success *video.GetWorkCountResponse
}

var GetWorkCountResult_Success_DEFAULT *video.GetWorkCountResponse

func (p *GetWorkCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.GetWorkCountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetWorkCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetWorkCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetWorkCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetWorkCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetWorkCountResult) Unmarshal(in []byte) error {
	msg := new(video.GetWorkCountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetWorkCountResult) GetSuccess() *video.GetWorkCountResponse {
	if !p.IsSetSuccess() {
		return GetWorkCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetWorkCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.GetWorkCountResponse)
}

func (p *GetWorkCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetWorkCountResult) GetResult() interface{} {
	return p.Success
}

func getVideoIDByUidHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.GetVideoIDByUidRequset)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).GetVideoIDByUid(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetVideoIDByUidArgs:
		success, err := handler.(video.VideoService).GetVideoIDByUid(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetVideoIDByUidResult)
		realResult.Success = success
	}
	return nil
}
func newGetVideoIDByUidArgs() interface{} {
	return &GetVideoIDByUidArgs{}
}

func newGetVideoIDByUidResult() interface{} {
	return &GetVideoIDByUidResult{}
}

type GetVideoIDByUidArgs struct {
	Req *video.GetVideoIDByUidRequset
}

func (p *GetVideoIDByUidArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.GetVideoIDByUidRequset)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetVideoIDByUidArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetVideoIDByUidArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetVideoIDByUidArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetVideoIDByUidArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetVideoIDByUidArgs) Unmarshal(in []byte) error {
	msg := new(video.GetVideoIDByUidRequset)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetVideoIDByUidArgs_Req_DEFAULT *video.GetVideoIDByUidRequset

func (p *GetVideoIDByUidArgs) GetReq() *video.GetVideoIDByUidRequset {
	if !p.IsSetReq() {
		return GetVideoIDByUidArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetVideoIDByUidArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetVideoIDByUidArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetVideoIDByUidResult struct {
	Success *video.GetVideoIDByUidResponse
}

var GetVideoIDByUidResult_Success_DEFAULT *video.GetVideoIDByUidResponse

func (p *GetVideoIDByUidResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.GetVideoIDByUidResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetVideoIDByUidResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetVideoIDByUidResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetVideoIDByUidResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetVideoIDByUidResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetVideoIDByUidResult) Unmarshal(in []byte) error {
	msg := new(video.GetVideoIDByUidResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetVideoIDByUidResult) GetSuccess() *video.GetVideoIDByUidResponse {
	if !p.IsSetSuccess() {
		return GetVideoIDByUidResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetVideoIDByUidResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.GetVideoIDByUidResponse)
}

func (p *GetVideoIDByUidResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetVideoIDByUidResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Feed(ctx context.Context, Req *video.FeedRequest) (r *video.FeedResponse, err error) {
	var _args FeedArgs
	_args.Req = Req
	var _result FeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PutVideo(ctx context.Context, Req *video.PutVideoRequest) (r *video.PutVideoResponse, err error) {
	var _args PutVideoArgs
	_args.Req = Req
	var _result PutVideoResult
	if err = p.c.Call(ctx, "PutVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFavoriteVideoInfo(ctx context.Context, Req *video.GetFavoriteVideoInfoRequest) (r *video.GetFavoriteVideoInfoResponse, err error) {
	var _args GetFavoriteVideoInfoArgs
	_args.Req = Req
	var _result GetFavoriteVideoInfoResult
	if err = p.c.Call(ctx, "GetFavoriteVideoInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPublishList(ctx context.Context, Req *video.GetPublishListRequest) (r *video.GetPublishListResponse, err error) {
	var _args GetPublishListArgs
	_args.Req = Req
	var _result GetPublishListResult
	if err = p.c.Call(ctx, "GetPublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetWorkCount(ctx context.Context, Req *video.GetWorkCountRequest) (r *video.GetWorkCountResponse, err error) {
	var _args GetWorkCountArgs
	_args.Req = Req
	var _result GetWorkCountResult
	if err = p.c.Call(ctx, "GetWorkCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideoIDByUid(ctx context.Context, Req *video.GetVideoIDByUidRequset) (r *video.GetVideoIDByUidResponse, err error) {
	var _args GetVideoIDByUidArgs
	_args.Req = Req
	var _result GetVideoIDByUidResult
	if err = p.c.Call(ctx, "GetVideoIDByUid", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
