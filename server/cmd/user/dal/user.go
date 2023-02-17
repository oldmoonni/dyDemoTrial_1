package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

type DUser struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
}

type DUserLock struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string
}

type DRecommend struct {
	Token   	string `gorm:"primaryKey"`
	Type1 		int64
	Type2		int64
	Type3		int64
}

type DFollow struct {
	Id      int64
	From_id int64
	To_id   int64
}

func (v DUser) TableName() string {
	return "duser"
}

func (v DUserLock) TableName() string {
	return "duserlock"
}

func (m DRecommend) TableName() string {
	return "drecommend"
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

func UserIsExistByName(askname string) (user DUser, flag bool) {
	db := getDB()
	var dusers []DUser
	db.Where(map[string]interface{}{"Name": askname}).Find(&dusers)

	flag = false
	if len(dusers) != 0 {
		user = dusers[0]
		flag = true
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

func UserInsert(id int64, name string) {
	db := getDB()
	db.Create(&DUser{Id: id, Name: name, FollowCount: 0, FollowerCount: 0})
}

func UserLockInsert(id int64, name string, password string) {
	db := getDB()
	db.Create(&DUserLock{Id: id, Name: name, Password: password, Token: name + password})
}

func UserLockInfoById(id int64) (duserlock DUserLock, flag bool) {
	db := getDB()
	var duserlocks []DUserLock
	db.Where(map[string]interface{}{"Id": id}).Find(&duserlocks)

	if len(duserlocks) != 0 {
		duserlock = duserlocks[0]
		flag = true
	} else {
		flag = false
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

func DrecomInsert(token string)  {
	db := getDB()
	db.Create(&DRecommend{Token: token, Type1: 10, Type2: 10, Type3: 10})
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
