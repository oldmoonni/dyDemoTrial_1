package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	dal2 "github.com/trial_1/dyDemoTrial_1/server/cmd/interact/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/interact"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"log"
	"time"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

func (s *CommentActionService) CommentAction(req *interact.CommentActionRequest) (resp *interact.CommentActionResponse, err error) {
	token := req.Token
	actionType := req.ActionType
	videoId := req.VideoId
	commentText := req.CommentText

	duserlock, flag := dal2.UserLockInfoByToken(token)
	duser, flag2 := dal2.UserIsExistById(duserlock.Id)

	if flag == true && flag2 == true {
		//actionType==1 发表评论 actionType==2删除评论
		if actionType == 1 {
			//生成的id
			node, err := snowflake.NewNode(1)
			if err != nil {
				log.Fatal(err)
			}
			commentId := int64(node.Generate())
			author := duserlock.Id
			currentTime := time.Now()
			createData := currentTime.Format("01-02")
			dal2.VideoAddComment(commentId, author, videoId, commentText, createData)
			resp = &interact.CommentActionResponse{
				Comment: &interact.Comment{
					Id: commentId,
					User: u2uPlusToken(duser, token),
					Content: commentText,
					CreateDate: createData,
				},
			}
			return resp, nil
		} else {
			commentId := req.CommentId
			dal2.VideoSubComment(commentId, videoId)
			return
		}
	} else {
		err = errno.UserInfoWrongErr
	}
	return
}