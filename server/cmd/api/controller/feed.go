package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/dao"
	"log"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	latest_time, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		log.Fatal("wrong latest_time")
	}
	println("latest_time = ", latest_time)
	println("当前时间戳: ", time.Now().Unix())

	var dd []dao.DVideo
	if token == "" {
		dd = dao.GetFeed(latest_time-500000)
		println("token 为空")
	} else {
		dd = dao.GetFeedByToken(latest_time-500000, token)
		println("token 存在")
	}
	DemoVideos2 := feedv2v(dd, token)

	c.JSON(consts.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos2,
		NextTime:  time.Now().Unix(),
	})
}

func feedv2v(dvideos []dao.DVideo, token string) (videos []Video) {
	videos = make([]Video, len(dvideos))
	for i := range dvideos {
		videos[i].Id = dvideos[i].Id
		duser, _ := dao.UserIsExistById(dvideos[i].Author)
		user := u2uplustoken(duser, token)
		videos[i].Author = user
		videos[i].PlayUrl = dvideos[i].PlayUrl
		videos[i].CoverUrl = dvideos[i].CoverUrl
		videos[i].FavoriteCount = dvideos[i].FavoriteCount
		videos[i].CommentCount = dvideos[i].CommentCount
		videos[i].IsFavorite = dao.VideoIsFavByToken(token, dvideos[i].Id)
	}
	return
}
