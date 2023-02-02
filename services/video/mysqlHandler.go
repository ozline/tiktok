package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Id            int64  // 用户id
	Name          string // 用户名称
	FollowCount   int64  // 关注总数
	FollowerCount int64  // 粉丝总数
	IsFollow      bool   // true-已关注，false-未关注
}

type Video struct {
	Id            int64  // 视频ID
	Author        *User  // 作者信息
	PlayUrl       string // 播放地址
	CoverUrl      string // 封面地址
	FavoriteCount int64  // 视频的点赞总数
	CommentCount  int64  // 视频的评论总数
	IsFavorite    bool   // 是否本人已点赞
	Title         string // 标题
}

type VideoInfo struct {
	gorm.Model
	Videoinfo *Video
}

func main() {
	db, err := gorm.Open(sqlite.Open("test1.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&VideoInfo{})
	// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//// Read
	//var product Product
	//db.First(&product) // 根据整形主键查找
	//fmt.Println("Code=", product.Code)
	//fmt.Println("prices=", product.Prce)
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//result := db.First(&product)
	//fmt.Println("Number =", result.RowsAffected) // 返回找到的记录数
	//fmt.Println("Price", product.Price)
	//fmt.Println("Code =", product.Code)
	//// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	////db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	////db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	//db.Delete(&product, 1)
}
