package dal

type DFollow struct {
	Id      int64
	From_id int64
	To_id   int64
}

func (f DFollow) TableName() string {
	return "dfollow"
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
