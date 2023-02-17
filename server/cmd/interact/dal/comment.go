package dao

type DComment struct {
	Comment_id  int64
	Author      int64
	Video_id    int64
	Content     string
	Create_data string
}

func (c DComment) TableName() string {
	return "dcomment"
}

func VideoAddComment(comment_id int64, author int64, video_id int64, content string, create_data string) {
	db := getDB()
	var dvideo DVideo
	db.Where(&DVideo{Id: video_id}).Find(&dvideo)
	comment_count := dvideo.CommentCount
	db.Model(&dvideo).Update("CommentCount", comment_count+1)
	db.Create(&DComment{Comment_id: comment_id, Author: author, Video_id: video_id, Content: content, Create_data: create_data})
}

func VideoSubComment(comment_id int64, video_id int64) {
	db := getDB()
	db.Where(map[string]interface{}{"Comment_id": comment_id}).Delete(DComment{})
	var dvideo DVideo
	db.Where(&DVideo{Id: video_id}).Find(&dvideo)
	comment_count := dvideo.CommentCount
	if comment_count > 0 {
		db.Model(&dvideo).Update("CommentCount", comment_count-1)
	}
}

func VideoGetCommentList(video_id int64) (dcomments []DComment) {
	db := getDB()
	dcomments = make([]DComment, 0)
	db.Where(&DComment{Video_id: video_id}).Order("Create_data desc").Find(&dcomments)
	return
}
