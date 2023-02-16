package main

import (
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: %w", err)
	}
	db.Migrator().CreateTable(&dao.DComment{})
	db.Migrator().CreateTable(&dao.DFollow{})
	db.Migrator().CreateTable(&dao.DUser{})
	db.Migrator().CreateTable(&dao.DFavorite{})
	db.Migrator().CreateTable(&dao.DUserLock{})
	db.Migrator().CreateTable(&dao.DVideo{})
	db.Migrator().CreateTable(&dao.DMessage{})
	db.Migrator().CreateTable(&dao.DRecommend{})

	//db.Create(&dao.DRecommend{Token: "zhangsan123456", Type1: 10, Type2: 10, Type3: 10})
	//db.Create(&dao.DRecommend{Token: "lisi123456", Type1: 10, Type2: 10, Type3: 10})

	//var videos []DVideo
	//db.Limit(30).Find(&videos)

	//user := User{Id: 4, Name: "zhangsan", FollowCount: 1, FollowerCount: 1, IsFollow: true}
	//db.Create(&DUser{5, "laowu", 6, 6, true})

	//userIsExist("laoliu")

	//ss := GetFeed()
	//println(ss[1].PlayUrl)

}
