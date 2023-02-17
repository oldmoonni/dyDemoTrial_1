package service

import (
	"context"
	dal2 "github.com/trial_1/dyDemoTrial_1/server/cmd/social/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/social"
	"log"
)

type FollowListService struct {
	ctx context.Context
}

func NewFollowListService(ctx context.Context) *FollowListService {
	return &FollowListService{ctx: ctx}
}

func (s *FollowListService) FollowList(req *social.FollowListRequest) (resp *social.FollowListResponse, err error) {
	userId := req.UserId
	token := req.Token

	dusers := dal2.GetFollowList(userId)
	users := u2uplustokenList(dusers, token)

	resp = &social.FollowListResponse{
		UserList: users,
	}
	return
}

func u2uplustokenList(dusers []dal2.DUser, token string) (users []*social.User) {
	users = make([]*social.User, len(dusers))
	for i := range dusers {
		duserlock, flag := dal2.UserLockInfoByToken(token)
		if flag == false {
			log.Fatal("wrong user information")
		}
		users[i] = &social.User{
			Id: dusers[i].Id,
			Name: dusers[i].Name,
			FollowCount: dusers[i].FollowCount,
			FollowerCount: dusers[i].FollowerCount,
			IsFollow: dal2.Isfollow(duserlock.Id, dusers[i].Id),
		}
	}
	return
}