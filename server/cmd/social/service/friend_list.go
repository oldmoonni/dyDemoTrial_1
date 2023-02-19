package service

import (
	"context"
	dal2 "github.com/trial_1/dyDemoTrial_1/server/cmd/social/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/social"
)

type FriendListService struct {
	ctx context.Context
}

func NewFriendListService(ctx context.Context) *FriendListService {
	return &FriendListService{ctx: ctx}
}

func (s *FriendListService) FriendList(req *social.FriendListRequest) (resp *social.FriendListResponse, err error) {
	userId := req.UserId
	token := req.Token

	dusers := dal2.GetFriendList(userId)
	friendUsers := u2fuplustokenList(dusers, token)

	resp = &social.FriendListResponse{
		UserList: friendUsers,
	}
	return
}

func u2fuplustokenList(dusers []dal2.DUser, token string) (users []*social.FriendUser) {
	users = make([]*social.FriendUser, len(dusers))
	for i := range dusers {
		duserlock, flag := dal2.UserLockInfoByToken(token)
		isFollow := dal2.Isfollow(duserlock.Id, dusers[i].Id)
		if flag == false {
			isFollow = false
		}
		users[i] = &social.FriendUser{
			Id: dusers[i].Id,
			Name: dusers[i].Name,
			FollowCount: dusers[i].FollowCount,
			FollowerCount: dusers[i].FollowerCount,
			IsFollow: isFollow,
			Avatar: dusers[i].Avatar,
			BackgroundImage: dusers[i].BackgroundImage,
			Signature: dusers[i].Signature,
			TotalFavorited: dusers[i].TotalFavorited,
			WorkCount: dusers[i].WorkCount,
			FavoriteCount: dusers[i].FavoriteCount,
			Message: "last message",
			MsgType: 1,
		}
	}
	return
}