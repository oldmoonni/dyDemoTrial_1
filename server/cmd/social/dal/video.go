package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

type DVideo struct {
	Id            int64  `json:"id,omitempty"`
	Author        int64  `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	Title         string `json:"title,omitempty"`
	Time          int64  `json:"time,omitempty"`
}

type DFavorite struct {
	Id       int64
	User_id  int64
	Video_id int64
	Time     int64
}

func (v DVideo) TableName() string {
	return "dvideo"
}

func (f DFavorite) TableName() string {
	return "dfavrite"
}

var (
	db   *gorm.DB
	once = &sync.Once{}
)

func getDB() *gorm.DB {
	once.Do(func() {
		dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database: %w", err)
		}
	})
	return db
}
