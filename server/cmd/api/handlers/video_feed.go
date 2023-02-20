// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/rpc"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/video"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"log"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

func VideoFeed(ctx context.Context, c *app.RequestContext) {
	println("/feed:get")
	token := c.Query("token")
	latestTime, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		log.Fatal("wrong latest_time")
	}

	resp := new(video.FeedResponse)
	resp, err = rpc.VideoFeed(context.Background(), &video.FeedRequest{
		LatestTime: latestTime,
		Token: token,
	})
	if err != nil {
		c.JSON(consts.StatusOK, FeedResponse{
			Response: Response{StatusCode: 1, StatusMsg: errno.ConvertErr(err).ErrMsg},
		})
		return
	}

	c.JSON(consts.StatusOK, FeedResponse{
		Response: Response{StatusCode: 0},
		VideoList:   videovp2v(resp.VideoList),
		NextTime:    resp.NextTime,
	})
}

func videovp2v(pvideos []*video.Video) (videos []Video) {
	videos = make([]Video, len(pvideos))
	for i := range pvideos{
		videos[i] = Video{
			Id: pvideos[i].Id,
			Author: User{
				Id: pvideos[i].Author.Id,
				Name: pvideos[i].Author.Name,
				FollowCount: pvideos[i].Author.FollowCount,
				FollowerCount: pvideos[i].Author.FollowerCount,
				IsFollow: pvideos[i].Author.IsFollow,
				Avatar: pvideos[i].Author.Avatar,
				BackgroundImage: pvideos[i].Author.BackgroundImage,
				Signature: pvideos[i].Author.Signature,
				TotalFavorited: pvideos[i].Author.TotalFavorited,
				WorkCount: pvideos[i].Author.WorkCount,
				FavoriteCount: pvideos[i].Author.FavoriteCount,
			},
			PlayUrl: pvideos[i].PlayUrl,
			CoverUrl: pvideos[i].CoverUrl,
			FavoriteCount: pvideos[i].FavoriteCount,
			CommentCount: pvideos[i].CommentCount,
			IsFavorite: pvideos[i].IsFavorite,
			Title: pvideos[i].Title,
		}
	}
	return
}
