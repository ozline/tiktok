package utils

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	//使用viper可以让我们专注于自己的项目代码，而不用自己写那些配置解析代码
	viper.SetConfigName("kitex") //配置文件名
	viper.AddConfigPath("chat")  //查找配置文件的路径 这里直接写目录名 不然会发生错误
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app inited...")
}

func InitMySQL() {

	//自定义日志模块 打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	fmt.Println("config mysql inited...")

}

func InitRedis() {

	//要加载到Red中
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("reds.minIdleConn"),
		//Addr:     "localhost:6379",
		//Password: "", // no password set
		//DB:       0,  // use default DB
	})

}
