package models

import (
	"getcharzp.cn/define"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

var RDB = InitRedisDB()

// 这里并不是main函数那里的init函数会自动执行，这里只是把连接数据库的代码给抽象出来了而已
func Init() *gorm.DB {
	dsn := define.MysqlDNS + "/my_gin_project?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
	return db
}

// redis应该是主要来存储用户的token
// 在使用之前，要先开启电脑的redis服务端。
func InitRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
