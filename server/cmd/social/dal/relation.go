package dal

import "log"

type DFollow struct {
	Id      int64
	From_id int64
	To_id   int64
}

func (f DFollow) TableName() string {
	return "dfollow"
}

func RelationAdd(follow_id int64, from_id int64, to_id int64) {
	db := getDB()
	var from_duser DUser
	db.Where(&DUser{Id: from_id}).Find(&from_duser)
	follow_count := from_duser.FollowCount
	db.Model(&from_duser).Update("FollowCount", follow_count+1)
	var to_duser DUser
	db.Where(&DUser{Id: to_id}).Find(&to_duser)
	follower_count := to_duser.FollowerCount
	db.Model(&to_duser).Update("FollowerCount", follower_count+1)
	db.Create(&DFollow{Id: follow_id, From_id: from_id, To_id: to_id})
}

func RelationSub(from_id int64, to_id int64) {
	db := getDB()
	db.Where(map[string]interface{}{"From_id": from_id, "To_id": to_id}).Delete(DFollow{})
	var from_duser DUser
	db.Where(&DUser{Id: from_id}).Find(&from_duser)
	follow_count := from_duser.FollowCount
	if follow_count > 0 {
		db.Model(&from_duser).Update("FollowCount", follow_count-1)
	}
	var to_duser DUser
	db.Where(&DUser{Id: to_id}).Find(&to_duser)
	follower_count := to_duser.FollowerCount
	if follower_count > 0 {
		db.Model(&to_duser).Update("FollowerCount", follower_count-1)
	}
}

func GetFollowList(user_id int64) (dusers []DUser) {
	db := getDB()
	dfollows := make([]DFollow, 0)
	db.Where(&DFollow{From_id: user_id}).Find(&dfollows)
	dusers = make([]DUser, len(dfollows))
	for i := range dfollows {
		id := dfollows[i].To_id
		var flag bool
		dusers[i], flag = UserIsExistById(id)
		if flag == false {
			log.Fatal("user does not exist")
		}
	}
	return
}

func GetFollowerList(user_id int64) (dusers []DUser) {
	db := getDB()
	dfollows := make([]DFollow, 0)
	db.Where(&DFollow{To_id: user_id}).Find(&dfollows)
	dusers = make([]DUser, len(dfollows))
	for i := range dfollows {
		id := dfollows[i].From_id
		var flag bool
		dusers[i], flag = UserIsExistById(id)
		if flag == false {
			log.Fatal("user does not exist")
		}
	}
	return
}

func GetFriendList(user_id int64) (friends []DUser) {
	db := getDB()
	dfollows := make([]DFollow, 0)
	db.Where(&DFollow{To_id: user_id}).Find(&dfollows)
	friends = make([]DUser, len(dfollows))
	j := 0
	for i := range dfollows {
		id := dfollows[i].From_id
		flag := Isfollow(user_id, id)
		if flag == true {
			friends[j], _ = UserIsExistById(id)
			j++
		}
	}
	return
}

func Isfollow(from_id int64, to_id int64) (flag bool) {
	if from_id == to_id {
		flag = true
		return
	}
	var dfollows []DFollow
	db.Where(map[string]interface{}{"From_id": from_id, "To_id": to_id}).Find(&dfollows)
	if len(dfollows) == 0 {
		flag = false
	} else {
		flag = true
	}
	return
}
