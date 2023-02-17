package dao

import (
	"fmt"
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

type DRecommend struct {
	Token   	string `gorm:"primaryKey"`
	Type1 		int64
	Type2		int64
	Type3		int64
}

type DUserLock struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string
}

type DUser struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
}

type DFollow struct {
	Id      int64
	From_id int64
	To_id   int64
}

func (v DVideo) TableName() string {
	return "dvideo"
}

func (f DFavorite) TableName() string {
	return "dfavrite"
}

func (m DRecommend) TableName() string {
	return "drecommend"
}

func (v DUserLock) TableName() string {
	return "duserlock"
}

func (v DUser) TableName() string {
	return "duser"
}

func (f DFollow) TableName() string {
	return "dfollow"
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

func GetFeed(latest_time int64) (dvideos []DVideo) {
	db := getDB()
	db.Where("time > ?", latest_time).Order("favorite_count desc").Limit(30).Find(&dvideos)
	return
}

func GetFeedByToken(latest_time int64, token string) (dvideos []DVideo) {
	db := getDB()
	drecom, flag := DrecomFindByToken(token)
	if flag == false {
		log.Fatal("can not find user drecommend")
	}
	sum := drecom.Type1 + drecom.Type2 +drecom.Type3
	s1 := int(float32(drecom.Type1)/float32(sum)*30)
	s2 := int(float32(drecom.Type2)/float32(sum)*30)
	s3 := int(float32(drecom.Type3)/float32(sum)*30)
	fmt.Println("s1,s2,s3分别为", s1, s2, s3)
	dvideos = make([]DVideo,30)
	db.Where("time > ? AND title = ?", latest_time, "dy1").Order("favorite_count desc").Limit(s1).Find(&dvideos)
	println("type1的长度为: ", len(dvideos))
	zj := make([]DVideo, 30)
	db.Where("time > ? AND title = ?", latest_time, "dy2").Order("favorite_count desc").Limit(s2).Find(&zj)
	dvideos = append(dvideos, zj...)
	println("type1和2的长度为: ", len(dvideos))
	db.Where("time > ? AND title = ?", latest_time, "dy3").Order("favorite_count desc").Limit(s3).Find(&zj)
	dvideos = append(dvideos, zj...)
	println("type1和2和3的长度为: ", len(dvideos))
	db.Where("time > ? AND title = ?", latest_time, "dy4").Order("favorite_count desc").Limit(30-len(dvideos)).Find(&zj)
	dvideos = append(dvideos, zj...)
	println("type1和2和3和其他的长度为: ", len(dvideos))
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

func GetVideosByUserId(user_id int64) (dvideos []DVideo) {
	db := getDB()
	db.Where(map[string]interface{}{"Author": user_id}).Find(&dvideos)
	return
}

func DrecomFindByToken(token string) (drecommend DRecommend, flag bool) {
	db := getDB()
	var drecommends []DRecommend
	db.Where(map[string]interface{}{"Token": token}).Find(&drecommends)

	if len(drecommends) != 0 {
		drecommend = drecommends[0]
		flag = true
	} else {
		flag = false
	}
	return
}

func UserIsExistById(id int64) (user DUser, flag bool) {
	db := getDB()
	var dusers []DUser
	db.Where(map[string]interface{}{"Id": id}).Find(&dusers)

	flag = false
	if len(dusers) != 0 {
		user = dusers[0]
		flag = true
	}
	return
}

func UserLockInfoByToken(token string) (duserlock DUserLock, flag bool) {
	db := getDB()
	var duserlocks []DUserLock
	db.Where(map[string]interface{}{"Token": token}).Find(&duserlocks)

	if len(duserlocks) != 0 {
		duserlock = duserlocks[0]
		flag = true
	} else {
		flag = false
	}
	return
}

func Isfollow(from_id int64, to_id int64) (flag bool) {
	var dfollows []DFollow
	db.Where(map[string]interface{}{"From_id": from_id, "To_id": to_id}).Find(&dfollows)
	if len(dfollows) == 0 {
		flag = false
	} else {
		flag = true
	}
	return
}