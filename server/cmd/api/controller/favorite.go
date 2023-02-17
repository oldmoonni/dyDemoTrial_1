package controller

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/dao"
	"log"
	"strconv"
	"time"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	video_id, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong user_id")
	}
	action_type, err := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err != nil {
		log.Fatal("wrong action_type")
	}

	duserlock, flag := dao.UserLockInfoByToken(token)
	if flag == false {
		c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	//action_type==1 点赞  action_type==2 取消点赞
	if action_type == 1 {
		if dao.VideoIsFavByToken(token, video_id) == true {
			c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: "you have done favoriteAction on this video"})
			return
		}
		//生成的id
		node, err := snowflake.NewNode(1)
		if err != nil {
			log.Fatal(err)
		}
		id := int64(node.Generate())
		user_id := duserlock.Id
		timeUnix := time.Now().Unix()
		dao.VideoAddFav(id, user_id, video_id, timeUnix)
		video := dao.GetVideosByVideoId(video_id)
		switch video.Title {
		case "dy1": dao.DrecomAdd(token, 1)
		case "dy2": dao.DrecomAdd(token, 2)
		case "dy3": dao.DrecomAdd(token, 3)
		}
		//有问题，重新写
		//dao.FavAddRedis(user_id, video.Title)
	} else {
		if dao.VideoIsFavByToken(token, video_id) == false {
			c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: "you have not done favoriteAction on this video"})
			return
		}
		user_id := duserlock.Id
		dao.VideoSubFav(user_id, video_id)
		video := dao.GetVideosByVideoId(video_id)
		switch video.Title {
		case "dy1": dao.DrecomSub(token, 1)
		case "dy2": dao.DrecomSub(token, 2)
		case "dy3": dao.DrecomSub(token, 3)
		}
	}
	c.JSON(consts.StatusOK, Response{StatusCode: 0})
}

// FavoriteList all users have same favorite video list
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	user_id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong user_id")
	}
	duserlock, _ := dao.UserLockInfoById(user_id)
	videos := feedv2v(dao.VideoGetFavList(user_id), duserlock.Token)
	c.JSON(consts.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
