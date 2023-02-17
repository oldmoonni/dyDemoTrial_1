package service

import (
	"context"
	dal2 "github.com/trial_1/dyDemoTrial_1/server/cmd/interact/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/interact"
	"log"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

func (s *FavoriteListService) FavoriteList(req *interact.FavoriteListRequest) (resp *interact.FavoriteListResponse, err error) {
	userId := req.UserId
	duserlock, _ := dal2.UserLockInfoById(userId)
	videos := feedv2v(dal2.VideoGetFavList(userId), duserlock.Token)

	resp = &interact.FavoriteListResponse{
		VideoList: videos,
	}
	return
}

func feedv2v(dvideos []dal2.DVideo, token string) (videos []*interact.Video) {
	videos = make([]*interact.Video, len(dvideos))
	for i := range dvideos {
		duser, _ := dal2.UserIsExistById(dvideos[i].Author)
		user := u2uPlusToken(duser, token)
		videos[i] = &interact.Video{
			Id: dvideos[i].Id,
			Author: user,
			PlayUrl: dvideos[i].PlayUrl,
			CoverUrl: dvideos[i].CoverUrl,
			FavoriteCount: dvideos[i].FavoriteCount,
			CommentCount: dvideos[i].CommentCount,
			IsFavorite: dal2.VideoIsFavByToken(token, dvideos[i].Id),
		}
	}
	return
}

func u2uPlusToken(duser dal2.DUser, token string) (user *interact.User) {
	duserlock, flag := dal2.UserLockInfoByToken(token)
	if flag == false {
		log.Fatal("wrong user information")
	}
	user = &interact.User{
		Id: duser.Id,
		Name: duser.Name,
		FollowCount: duser.FollowCount,
		FollowerCount: duser.FollowerCount,
		IsFollow: dal2.Isfollow(duserlock.Id, duser.Id),
	}
	return
}
