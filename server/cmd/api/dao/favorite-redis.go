package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)

var (
	rdb   *redis.Client
	once1 = &sync.Once{}
)

func getRDB() *redis.Client {
	once1.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379", //连接地址
			Password: "",               //连接密码
			DB:       0,                //默认连接库
			PoolSize: 100,              //连接池大小
		})
		_, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal("redis连接失败: %w", err)
			return
		}
	})
	return rdb
}

func FavAddRedis(userId int64, title string) {
	rdb := getRDB()
	err := rdb.Set(context.Background(), string(userId), title, 100000000).Err()
	if err != nil {
		fmt.Println("redis点赞失败")
		return
	}
}

func FavSubRedis(userId int64) {
	rdb := getRDB()
	err1 := rdb.Del(context.Background(), string(userId)).Err()
	if err1 != nil {
		fmt.Println("redis取消点赞失败")
		return
	}
}