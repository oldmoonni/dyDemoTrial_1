package service

import (
	"context"
	dal2 "github.com/trial_1/dyDemoTrial_1/server/cmd/social/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/social"
)

type FollowerListService struct {
	ctx context.Context
}

func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{ctx: ctx}
}

func (s *FollowerListService) FollowerList(req *social.FollowerListRequest) (resp *social.FollowerListResponse, err error) {
	userId := req.UserId
	token := req.Token
	dusers := dal2.GetFollowerList(userId)
	users := u2uplustokenList(dusers, token)

	resp = &social.FollowerListResponse{
		UserList: users,
	}
	return
}