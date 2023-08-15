package db

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Message struct {
    Id              int64
	ToUserId        int64
    FromUserId      int64
    Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func GetMessageList(ctx context.Context,to_user_id int64,from_user_id int64)([]*Message,error){
	messageList := make([]*Message,0)
	//redis  ZSET 
	//RedisDB.WithContext(ctx)
	key := strconv.Itoa(to_user_id)+"-"+strconv.Itoa(from_user_id)
	revkey := strconv.Itoa(from_user_id)+"-"+strconv.Itoa(to_user_id)
    ok1,_ := RedisDB.Exists(context.Background(),key ).Result()
    if ok1{
        //查询 a->b的消息
    mem,err := RedisDB.ZRevRangeByScore(ctx,key,&redis.ZRangeBy{
        Min:strconv.Itoa(0),
        Max:strconv.Itoa(int(time.Now().Unix())),
	}).Result()
		if err!=nil{
			return nil,err
		}
	err := json.Unmarshal([]byte(mem),messageList)
		if err!=nil{
			return nil,err
		}   
    }
	ok2,_ :=db.RedisDB.Exists(context.Background(), revkey).Result()
	if ok2{
		 mem,err := RedisDB.ZRevRangeByScore(ctx,key,&redis.ZRangeBy{
        Min:strconv.Itoa(0),
        Max:strconv.Itoa(int(time.Now().Unix())),
	}).Result()
		if err!=nil{
			return nil,err
		}
	temp_message :=  make([]*Message,0)
	err := json.Unmarshal([]byte(mem),temp_message)
		if err!=nil{
			return nil,err
		}  
	messageList = append(messageList,temp_message)
	}
	if len(messageList)>0 {
		//合并排序

	}
    //mysql
    err := DB.WithContext(ctx).Where("(to_user_id=? AND from_user_id =?) OR (to_user_id=? AND from_user_id =?) ",to_user_id,from_user_id,from_user_id,to_user_id).Order("created_at desc").Find(&MessageList).Error
    if err != nil {
		// add some logs

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
    return MessageList,nil
}