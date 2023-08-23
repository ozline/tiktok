package service

import (
	"strconv"
	"sync"
	"time"

	"github.com/ozline/tiktok/cmd/interaction/pack"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/kitex_gen/user"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"

	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) GetComments(req *interaction.CommentListRequest) ([]*interaction.Comment, error) {

	var comments []db.Comment
	key := strconv.FormatInt(req.VideoId, 10)
	exist, err := cache.IsExistComment(s.ctx, key)
	if err != nil {
		return nil, err
	}

	if exist == 1 {
		rComments, err := cache.GetComments(s.ctx, key)
		if err != nil {
			return nil, err
		}
		for _, rComment := range *rComments {
			var comment db.Comment
			_, err = comment.UnmarshalMsg([]byte(rComment.Member.(string)))
			if err != nil {
				return nil, err
			}
			comment.CreatedAt = time.Unix(int64(rComment.Score), 0)
			comments = append(comments, comment)
		}

	} else {
		comments, err = db.GetCommentsByVideoID(s.ctx, req.VideoId)
		if err != nil {
			return nil, err
		}

		if len(comments) != 0 {
			err = cache.AddComments(s.ctx, key, &comments)
			if err != nil {
				return nil, err
			}

		}
	}

	var wg sync.WaitGroup
	//users := make(map[int64]int) // 利用map避免重复查询
	commentList := make([]*interaction.Comment, len(comments))
	errs := make([]error, len(comments))
	for index, data := range comments {

		//index, ok := users[comment.UserId]
		//if !ok {
		//	userInfo, err := rpc.UserInfo(s.ctx, &user.InfoRequest{
		//		UserId: comment.UserId,
		//		Token:  req.Token,
		//	})
		//	if err != nil {
		//		return resp, nil
		//	}
		//	rComment.User = userInfo
		//	users[comment.UserId] = commentIndex
		//} else {
		//	rComment.User = commentList[index].User
		//}
		comment := data
		commentIndex := index
		wg.Add(1)
		go func() {
			defer wg.Done()
			userInfo, err := rpc.UserInfo(s.ctx, &user.InfoRequest{
				UserId: comment.UserId,
				Token:  req.Token,
			})
			rComment := pack.Comment(&comment, userInfo)
			commentList[commentIndex] = rComment
			errs[commentIndex] = err
		}()
	}
	wg.Wait()

	for _, err = range errs {
		if err != nil {
			return nil, err
		}
	}

	return commentList, nil
}
