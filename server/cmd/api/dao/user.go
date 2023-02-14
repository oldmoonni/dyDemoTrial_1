package dao

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

func (v DUser) TableName() string {
	return "duser"
}

func (v DUserLock) TableName() string {
	return "duserlock"
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
