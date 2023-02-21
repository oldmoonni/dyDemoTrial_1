package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	dal2 "github.com/trial_1/dyDemoTrial_1/server/cmd/interact/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/interact"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"log"
	"sync"
	"time"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

var WaitGroup sync.WaitGroup
func (s *FavoriteActionService) FavoriteAction(req *interact.FavoriteActionRequest) (err error) {
	token := req.Token
	videoId := req.VideoId
	actionType := req.ActionType

	duserlock, flag := dal2.UserLockInfoByToken(token)
	if flag == false {
		err = errno.UserNotExistErr
		return
	}

	//action_type==1 点赞  action_type==2 取消点赞
	if actionType == 1 {
		if dal2.VideoIsFavByToken(token, videoId) == true {
			err = errno.HaveDoneErr
			return
		}
		//生成的id
		node, err := snowflake.NewNode(1)
		if err != nil {
			log.Fatal(err)
		}
		id := int64(node.Generate())
		userId := duserlock.Id
		timeUnix := time.Now().Unix()

		var video dal2.DVideo
		WaitGroup.Add(2)
		go func() {
			dal2.VideoAddFav(id, userId, videoId, timeUnix)
			println("第一个协程完成")
			WaitGroup.Done()
		}()
		go func() {
			video = dal2.GetVideosByVideoId(videoId)
			dal2.UserAddFav(video, userId)
			println("第二个协程完成")
			WaitGroup.Done()
		}()
		WaitGroup.Wait()

		switch video.Title {
		case "dy1": dal2.DrecomAdd(token, 1)
		case "dy2": dal2.DrecomAdd(token, 2)
		case "dy3": dal2.DrecomAdd(token, 3)
		}
		//有问题，重新写
		//dao.FavAddRedis(user_id, video.Title)
	} else {
		if dal2.VideoIsFavByToken(token, videoId) == false {
			err = errno.HaveDoneErr
			return
		}
		userId := duserlock.Id
		var video dal2.DVideo

		WaitGroup.Add(2)
		go func() {
			dal2.VideoSubFav(userId, videoId)
			println("第一个协程完成")
			WaitGroup.Done()
		}()
		go func() {
			video = dal2.GetVideosByVideoId(videoId)
			dal2.UserSubFav(video, userId)
			println("第二个协程完成")
			WaitGroup.Done()
		}()
		WaitGroup.Wait()

		switch video.Title {
		case "dy1": dal2.DrecomSub(token, 1)
		case "dy2": dal2.DrecomSub(token, 2)
		case "dy3": dal2.DrecomSub(token, 3)
		}
	}
	return
}