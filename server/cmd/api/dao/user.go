package dao

type DUser struct {
	Id            			int64
	Name          			string
	FollowCount   			int64
	FollowerCount 			int64
	Avatar					string
	BackgroundImage			string
	Signature				string
	TotalFavorited			int64
	WorkCount				int64
	FavoriteCount			int64
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

func (v DUser) TableName() string {
	return "duser"
}

func (v DUserLock) TableName() string {
	return "duserlock"
}

func (m DRecommend) TableName() string {
	return "drecommend"
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
