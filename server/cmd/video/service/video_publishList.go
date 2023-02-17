package service

import (
	"context"
	dal2 "github.com/trial_1/dyDemoTrial_1/server/cmd/video/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/video"
)

type VideoPublishListService struct {
	ctx context.Context
}

func NewVideoPublishListService(ctx context.Context) *VideoPublishListService {
	return &VideoPublishListService{ctx: ctx}
}

func (s *VideoPublishListService) VideoPublishList(req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	userId := req.UserId
	token := req.Token

	dvideos := dal2.GetVideosByUserId(userId)
	videoList := feedv2v(dvideos, token)

	resp = &video.PublishListResponse{
		VideoList: videoList,
	}
	return
}