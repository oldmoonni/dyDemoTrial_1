package dal

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

func MessageInsert(messageId int64, fromUserId int64, toUserId int64, content string, createTime int64)  {
	db := getDB()
	db.Create(&DMessage{
		Message_id:   messageId,
		To_user_id:   toUserId,
		From_user_id: fromUserId,
		Content:      content,
		Create_time:  createTime,
	})
}

func MessageSearch(fromUserId int64, toUserId int64, preMsgTime int64) (messages []DMessage) {
	db := getDB()
	db.Where("create_time > ?", preMsgTime).Where("to_user_id = ? AND from_user_id = ? OR to_user_id = ? AND from_user_id = ?", toUserId, fromUserId, fromUserId, toUserId).Order("create_time").Find(&messages)
	//db.Where("to_user_id = ? AND from_user_id = ? OR to_user_id = ? AND from_user_id = ?", toUserId, fromUserId, fromUserId, toUserId).Order("create_time").Find(&messages)
	return
}

func MessageFindLatest(fromUserId int64, toUserId int64) (message DMessage) {
	db := getDB()
	var messages []DMessage
	db.Where("to_user_id = ? AND from_user_id = ? OR to_user_id = ? AND from_user_id = ?", toUserId, fromUserId, fromUserId, toUserId).Order("create_time desc").Find(&messages)
	if len(messages) > 0 {
		message = messages[0]
	}
	return
}
