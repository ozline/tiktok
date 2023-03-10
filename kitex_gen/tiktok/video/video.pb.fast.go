// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package video

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *BaseResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseResp[number], err)
}

func (x *BaseResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Msg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishActionResquest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PublishActionResquest[number], err)
}

func (x *PublishActionResquest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Data, offset, err = fastpb.ReadBytes(buf, _type)
	return offset, err
}

func (x *PublishActionResquest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Userid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *PublishActionResquest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishActionResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PublishActionResponse[number], err)
}

func (x *PublishActionResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Base = &v
	return offset, nil
}

func (x *PublishListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PublishListRequest[number], err)
}

func (x *PublishListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Userid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *PublishListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.PageNum, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *PublishListRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PageSize, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *PublishListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PublishListResponse[number], err)
}

func (x *PublishListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Base = &v
	return offset, nil
}

func (x *PublishListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *FeedRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FeedRequest[number], err)
}

func (x *FeedRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.LatestTime, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FeedRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FeedRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PageNum, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FeedRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.PageSize, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FeedResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FeedResponse[number], err)
}

func (x *FeedResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Base = &v
	return offset, nil
}

func (x *FeedResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *FeedResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.NextTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetInfoRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetInfoRequest[number], err)
}

func (x *GetInfoRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetInfoResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetInfoResponse[number], err)
}

func (x *GetInfoResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Base = &v
	return offset, nil
}

func (x *GetInfoResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Video = &v
	return offset, nil
}

func (x *Video) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Video[number], err)
}

func (x *Video) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PlayUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CoverUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *BaseResp) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Code)
	return offset
}

func (x *BaseResp) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.Msg)
	return offset
}

func (x *PublishActionResquest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *PublishActionResquest) fastWriteField1(buf []byte) (offset int) {
	if len(x.Data) == 0 {
		return offset
	}
	offset += fastpb.WriteBytes(buf[offset:], 1, x.Data)
	return offset
}

func (x *PublishActionResquest) fastWriteField2(buf []byte) (offset int) {
	if x.Userid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.Userid)
	return offset
}

func (x *PublishActionResquest) fastWriteField3(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.Title)
	return offset
}

func (x *PublishActionResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *PublishActionResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Base == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.Base)
	return offset
}

func (x *PublishListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *PublishListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Userid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Userid)
	return offset
}

func (x *PublishListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.PageNum == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.PageNum)
	return offset
}

func (x *PublishListRequest) fastWriteField3(buf []byte) (offset int) {
	if x.PageSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.PageSize)
	return offset
}

func (x *PublishListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *PublishListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Base == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.Base)
	return offset
}

func (x *PublishListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.VideoList {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.VideoList[i])
	}
	return offset
}

func (x *FeedRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *FeedRequest) fastWriteField1(buf []byte) (offset int) {
	if x.LatestTime == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.LatestTime)
	return offset
}

func (x *FeedRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.UserId)
	return offset
}

func (x *FeedRequest) fastWriteField3(buf []byte) (offset int) {
	if x.PageNum == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.PageNum)
	return offset
}

func (x *FeedRequest) fastWriteField4(buf []byte) (offset int) {
	if x.PageSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.PageSize)
	return offset
}

func (x *FeedResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *FeedResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Base == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.Base)
	return offset
}

func (x *FeedResponse) fastWriteField3(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.VideoList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.VideoList[i])
	}
	return offset
}

func (x *FeedResponse) fastWriteField4(buf []byte) (offset int) {
	if x.NextTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.NextTime)
	return offset
}

func (x *GetInfoRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetInfoRequest) fastWriteField1(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.VideoId)
	return offset
}

func (x *GetInfoResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetInfoResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Base == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.Base)
	return offset
}

func (x *GetInfoResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Video == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.Video)
	return offset
}

func (x *Video) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	return offset
}

func (x *Video) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *Video) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.UserId)
	return offset
}

func (x *Video) fastWriteField3(buf []byte) (offset int) {
	if x.PlayUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.PlayUrl)
	return offset
}

func (x *Video) fastWriteField4(buf []byte) (offset int) {
	if x.CoverUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.CoverUrl)
	return offset
}

func (x *Video) fastWriteField8(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.Title)
	return offset
}

func (x *BaseResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *BaseResp) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Code)
	return n
}

func (x *BaseResp) sizeField2() (n int) {
	if x.Msg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.Msg)
	return n
}

func (x *PublishActionResquest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *PublishActionResquest) sizeField1() (n int) {
	if len(x.Data) == 0 {
		return n
	}
	n += fastpb.SizeBytes(1, x.Data)
	return n
}

func (x *PublishActionResquest) sizeField2() (n int) {
	if x.Userid == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.Userid)
	return n
}

func (x *PublishActionResquest) sizeField3() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(3, x.Title)
	return n
}

func (x *PublishActionResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *PublishActionResponse) sizeField1() (n int) {
	if x.Base == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.Base)
	return n
}

func (x *PublishListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *PublishListRequest) sizeField1() (n int) {
	if x.Userid == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Userid)
	return n
}

func (x *PublishListRequest) sizeField2() (n int) {
	if x.PageNum == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.PageNum)
	return n
}

func (x *PublishListRequest) sizeField3() (n int) {
	if x.PageSize == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.PageSize)
	return n
}

func (x *PublishListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *PublishListResponse) sizeField1() (n int) {
	if x.Base == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.Base)
	return n
}

func (x *PublishListResponse) sizeField2() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.VideoList {
		n += fastpb.SizeMessage(2, x.VideoList[i])
	}
	return n
}

func (x *FeedRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *FeedRequest) sizeField1() (n int) {
	if x.LatestTime == "" {
		return n
	}
	n += fastpb.SizeString(1, x.LatestTime)
	return n
}

func (x *FeedRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.UserId)
	return n
}

func (x *FeedRequest) sizeField3() (n int) {
	if x.PageNum == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.PageNum)
	return n
}

func (x *FeedRequest) sizeField4() (n int) {
	if x.PageSize == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.PageSize)
	return n
}

func (x *FeedResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *FeedResponse) sizeField1() (n int) {
	if x.Base == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.Base)
	return n
}

func (x *FeedResponse) sizeField3() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.VideoList {
		n += fastpb.SizeMessage(3, x.VideoList[i])
	}
	return n
}

func (x *FeedResponse) sizeField4() (n int) {
	if x.NextTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.NextTime)
	return n
}

func (x *GetInfoRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetInfoRequest) sizeField1() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.VideoId)
	return n
}

func (x *GetInfoResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetInfoResponse) sizeField1() (n int) {
	if x.Base == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.Base)
	return n
}

func (x *GetInfoResponse) sizeField2() (n int) {
	if x.Video == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.Video)
	return n
}

func (x *Video) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField8()
	return n
}

func (x *Video) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

func (x *Video) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.UserId)
	return n
}

func (x *Video) sizeField3() (n int) {
	if x.PlayUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.PlayUrl)
	return n
}

func (x *Video) sizeField4() (n int) {
	if x.CoverUrl == "" {
		return n
	}
	n += fastpb.SizeString(4, x.CoverUrl)
	return n
}

func (x *Video) sizeField8() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(8, x.Title)
	return n
}

var fieldIDToName_BaseResp = map[int32]string{
	1: "Code",
	2: "Msg",
}

var fieldIDToName_PublishActionResquest = map[int32]string{
	1: "Data",
	2: "Userid",
	3: "Title",
}

var fieldIDToName_PublishActionResponse = map[int32]string{
	1: "Base",
}

var fieldIDToName_PublishListRequest = map[int32]string{
	1: "Userid",
	2: "PageNum",
	3: "PageSize",
}

var fieldIDToName_PublishListResponse = map[int32]string{
	1: "Base",
	2: "VideoList",
}

var fieldIDToName_FeedRequest = map[int32]string{
	1: "LatestTime",
	2: "UserId",
	3: "PageNum",
	4: "PageSize",
}

var fieldIDToName_FeedResponse = map[int32]string{
	1: "Base",
	3: "VideoList",
	4: "NextTime",
}

var fieldIDToName_GetInfoRequest = map[int32]string{
	1: "VideoId",
}

var fieldIDToName_GetInfoResponse = map[int32]string{
	1: "Base",
	2: "Video",
}

var fieldIDToName_Video = map[int32]string{
	1: "Id",
	2: "UserId",
	3: "PlayUrl",
	4: "CoverUrl",
	8: "Title",
}
