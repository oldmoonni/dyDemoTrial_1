package controller

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/dao"
	"log"
	"strconv"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err != nil {
		log.Fatal("wrong actionType")
	}
	video_id, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong video_id")
	}
	comment_text := c.Query("comment_text")

	duserlock, flag := dao.UserLockInfoByToken(token)
	duser, flag2 := dao.UserIsExistById(duserlock.Id)

	if flag == true && flag2 == true {
		//actionType==1 发表评论 actionType==2删除评论
		if actionType == 1 {
			//生成的id
			node, err := snowflake.NewNode(1)
			if err != nil {
				log.Fatal(err)
			}
			comment_id := int64(node.Generate())
			author := duserlock.Id
			currentTime := time.Now()
			create_data := currentTime.Format("01-02")
			dao.VideoAddComment(comment_id, author, video_id, comment_text, create_data)
			c.JSON(consts.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         comment_id,
					User:       u2u(duser),
					Content:    comment_text,
					CreateDate: create_data,
				}})
		} else {
			comment_id, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
			if err != nil {
				log.Fatal("wrong comment_id")
			}
			dao.VideoSubComment(comment_id, video_id)
			c.JSON(consts.StatusOK, Response{StatusCode: 0})
		}
	} else {
		c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: "wrong User information"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	video_id, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong video_id")
	}
	dcomments := dao.VideoGetCommentList(video_id)
	comments := c2c(dcomments, token)
	c.JSON(consts.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: comments,
	})
}

//type Comment struct {
//	Id         int64  `json:"id,omitempty"`
//	User       User   `json:"user"`
//	Content    string `json:"content,omitempty"`
//	CreateDate string `json:"create_date,omitempty"`
//}

func c2c(dcomments []dao.DComment, token string) (comments []Comment) {
	comments = make([]Comment, len(dcomments))
	for i := range dcomments {
		comments[i].Id = dcomments[i].Comment_id
		duser, _ := dao.UserIsExistById(dcomments[i].Author)
		user := u2uplustoken(duser, token)
		comments[i].User = user
		comments[i].Content = dcomments[i].Content
		comments[i].CreateDate = dcomments[i].Create_data
	}
	return
}
