package dal

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

func DrecomAdd(token string, n int) {
	db := getDB()
	var drecommend DRecommend
	db.Where(&DRecommend{Token: token}).Find(&drecommend)
	switch n {
	case 1:
		s1 := drecommend.Type1
		db.Model(&drecommend).Update("type1", s1+1)
	case 2:
		s2 := drecommend.Type2
		db.Model(&drecommend).Update("type2", s2+1)
	case 3:
		s3 := drecommend.Type3
		db.Model(&drecommend).Update("type3", s3+1)
	}
}


func DrecomSub(token string, n int) {
	db := getDB()
	var drecommend DRecommend
	db.Where(&DRecommend{Token: token}).Find(&drecommend)
	switch n {
	case 1:
		s1 := drecommend.Type1
		if s1 == 0 {
			return
		}
		db.Model(&drecommend).Update("Type1", s1-1)
	case 2:
		s2 := drecommend.Type2
		if s2 == 0 {
			return
		}
		db.Model(&drecommend).Update("Type2", s2-1)
	case 3:
		s3 := drecommend.Type3
		if s3 == 0 {
			return
		}
		db.Model(&drecommend).Update("Type3", s3-1)
	}
}