package service

import (
	"context"
	dal2 "github.com/trial_1/dyDemoTrial_1/server/cmd/interact/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/interact"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

func (s *CommentListService) CommentList(req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	token := req.Token
	videoId := req.VideoId
	dcomments := dal2.VideoGetCommentList(videoId)
	comments := c2c(dcomments, token)
	resp = &interact.CommentListResponse{
		CommentList: comments,
	}
	return
}

func c2c(dcomments []dal2.DComment, token string) (comments []*interact.Comment) {
	comments = make([]*interact.Comment, len(dcomments))
	for i := range dcomments {
		duser, _ := dal2.UserIsExistById(dcomments[i].Author)
		user := u2uPlusToken(duser, token)
		comments[i] = &interact.Comment{
			Id: dcomments[i].Comment_id,
			User: user,
			Content: dcomments[i].Content,
			CreateDate: dcomments[i].Create_data,
		}
	}
	return
}