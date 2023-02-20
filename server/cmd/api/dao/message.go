package dao

type DMessage struct {
	Message_id   int64
	To_user_id   int64
	From_user_id int64
	Content      string
	Create_time  int64
}

func (m DMessage) TableName() string {
	return "dmessage"
}
