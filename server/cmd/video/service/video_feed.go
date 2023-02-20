package service

import (
	"context"
	dal2 "github.com/trial_1/dyDemoTrial_1/server/cmd/video/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/video"
	"time"
)

type VideoFeedService struct {
	ctx context.Context
}

func NewVideoFeedService(ctx context.Context) *VideoFeedService {
	return &VideoFeedService{ctx: ctx}
}

func (s *VideoFeedService) VideoFeed(req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	latestTime := req.LatestTime
	token := req.Token
	println("latest_time = ", latestTime)
	println("当前时间戳: ", time.Now().Unix())

	var dd []dal2.DVideo
	if token == "" {
		dd = dal2.GetFeed(latestTime-1000000)
		println("token 为空")
	} else {
		dd = dal2.GetFeedByToken(latestTime-1000000, token)
		println("token 存在")
	}
	videoList := feedv2v(dd, token)

	resp = &video.FeedResponse{
		NextTime: time.Now().Unix(),
		VideoList: videoList,
	}
	return
}

func feedv2v(dvideos []dal2.DVideo, token string) (videos []*video.Video) {
	videos = make([]*video.Video, len(dvideos))
	for i := range dvideos {
		duser, _ := dal2.UserIsExistById(dvideos[i].Author)
		user := u2uPlusToken(duser, token)
		videos[i] = &video.Video{
			Id: dvideos[i].Id,
			Author: user,
			PlayUrl: dvideos[i].PlayUrl,
			CoverUrl: dvideos[i].CoverUrl,
			FavoriteCount: dvideos[i].FavoriteCount,
			CommentCount: dvideos[i].CommentCount,
			IsFavorite: dal2.VideoIsFavByToken(token, dvideos[i].Id),
			Title: dvideos[i].Title,
		}
	}
	return
}

func u2uPlusToken(duser dal2.DUser, token string) (user *video.User) {
	duserlock, flag := dal2.UserLockInfoByToken(token)
	isFollow := dal2.Isfollow(duserlock.Id, duser.Id)
	if flag == false {
		isFollow = false
	}
	user = &video.User{
		Id: duser.Id,
		Name: duser.Name,
		FollowCount: duser.FollowCount,
		FollowerCount: duser.FollowerCount,
		IsFollow: isFollow,
		Avatar: duser.Avatar,
		BackgroundImage: duser.BackgroundImage,
		Signature: duser.Signature,
		TotalFavorited: duser.TotalFavorited,
		WorkCount: duser.WorkCount,
		FavoriteCount: duser.FavoriteCount,
	}
	return
}
