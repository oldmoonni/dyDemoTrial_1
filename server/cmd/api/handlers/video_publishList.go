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

func VideoPublishList(ctx context.Context, c *app.RequestContext) {
	println("/publish/list:get")
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong user_id")
	}
	token := c.Query("token")

	if len(token) == 0 {
		c.JSON(consts.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "wrong Param"},
		})
		return
	}

	resp := new(video.PublishListResponse)
	resp, err = rpc.VideoPublishList(context.Background(), &video.PublishListRequest{
		UserId: userId,
		Token: token,
	})
	if err != nil {
		c.JSON(consts.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: errno.ConvertErr(err).ErrMsg},
		})
		return
	}

	c.JSON(consts.StatusOK, VideoListResponse{
		Response: Response{StatusCode: 0},
		VideoList: videovp2v(resp.VideoList),
	})
}