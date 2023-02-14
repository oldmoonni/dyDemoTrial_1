package dao

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

//func main()  {
//
//	dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal("failed to connect database: %w", err)
//	}
//	//db.Migrator().CreateTable(&DComment{})
//
//	//var videos []DVideo
//	//db.Limit(30).Find(&videos)
//
//	//user := User{Id: 4, Name: "zhangsan", FollowCount: 1, FollowerCount: 1, IsFollow: true}
//	//db.Create(&DUser{5, "laowu", 6, 6, true})
//
//	//userIsExist("laoliu")
//
//	//ss := GetFeed()
//	//println(ss[1].PlayUrl)
//
//}

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

func GetFeed(latest_time int64) (dvideos []DVideo) {
	db := getDB()
	db.Where("time > ?", latest_time).Order("favorite_count desc").Limit(30).Find(&dvideos)
	return
}

func VideoIsFavByToken(token string, video_id int64) (flag bool) {
	db := getDB()
	var duserlock DUserLock
	db.Where(map[string]interface{}{"Token": token}).Find(&duserlock)
	dfavs := make([]*DFavorite, 0)
	db.Where(map[string]interface{}{"User_id": duserlock.Id, "Video_id": video_id}).Find(&dfavs)
	if len(dfavs) != 0 {
		flag = true
	} else {
		flag = false
	}
	return
}

func GetVideosByVideoId(video_id int64) (dvideo DVideo) {
	db := getDB()
	db.Where(map[string]interface{}{"Id": video_id}).Find(&dvideo)
	return
}

func GetVideosByUserId(user_id int64) (dvideos []DVideo) {
	db := getDB()
	db.Where(map[string]interface{}{"Author": user_id}).Find(&dvideos)
	return
}

func VideoInsert(id int64, author int64, play_url string, cover_url string, title string, time int64) {
	db := getDB()
	db.Create(&DVideo{Id: id, Author: author, PlayUrl: play_url, CoverUrl: cover_url, Title: title, Time: time})
}

func VideoFindById(video_id int64) (dvideo DVideo) {
	db := getDB()
	db.Where(map[string]interface{}{"Id": video_id}).Find(&dvideo)
	return
}

func VideoAddFav(id int64, user_id int64, video_id int64, time int64) {
	db := getDB()
	var dvideo DVideo
	db.Where(&DVideo{Id: video_id}).Find(&dvideo)
	fav_count := dvideo.FavoriteCount
	db.Model(&dvideo).Update("FavoriteCount", fav_count+1)
	db.Create(&DFavorite{Id: id, User_id: user_id, Video_id: video_id, Time: time})
}

func VideoSubFav(user_id int64, video_id int64) {
	db := getDB()
	db.Where(map[string]interface{}{"User_id": user_id, "Video_id": video_id}).Delete(DFavorite{})
	var dvideo DVideo
	db.Where(&DVideo{Id: video_id}).Find(&dvideo)
	fav_count := dvideo.FavoriteCount
	if fav_count > 0 {
		db.Model(&dvideo).Update("FavoriteCount", fav_count-1)
	}
}

func VideoGetFavList(user_id int64) (dvideos []DVideo) {
	db := getDB()
	dfav := make([]*DFavorite, 0)
	db.Where(&DFavorite{User_id: user_id}).Order("Time desc").Find(&dfav)
	dvideos = make([]DVideo, len(dfav))
	for i := range dfav {
		dvideos[i] = GetVideosByVideoId(dfav[i].Video_id)
	}
	return
}
