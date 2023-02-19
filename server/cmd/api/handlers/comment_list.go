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
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/rpc"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/interact"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"log"
	"strconv"
)

func CommentList(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong video_id")
	}

	resp := new(interact.CommentListResponse)
	resp, err = rpc.CommentList(context.Background(), &interact.CommentListRequest{
		Token: token,
		VideoId: videoId,
	})
	if err != nil {
		c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: errno.ConvertErr(err).ErrMsg})
		return
	}

	c.JSON(consts.StatusOK, CommentListResponse{
		Response: 	Response{StatusCode: 0},
		CommentList: cp2cList(resp.CommentList),
	})
}

func cp2cList(pcomment []*interact.Comment) (comment []Comment) {
	comment = make([]Comment, len(pcomment))
	for i := range pcomment {
		comment[i] = Comment{
			Id: pcomment[i].Id,
			User: User{
				Id:           	 pcomment[i].User.Id,
				Name:        	 pcomment[i].User.Name,
				FollowCount:  	 pcomment[i].User.FollowCount,
				FollowerCount:   pcomment[i].User.FollowerCount,
				IsFollow:    	 pcomment[i].User.IsFollow,
				Avatar: 		 pcomment[i].User.Avatar,
				BackgroundImage: pcomment[i].User.BackgroundImage,
				Signature: 		 pcomment[i].User.Signature,
				TotalFavorited:  pcomment[i].User.TotalFavorited,
				WorkCount: 		 pcomment[i].User.WorkCount,
				FavoriteCount: 	 pcomment[i].User.FavoriteCount,
			},
			Content:    pcomment[i].Content,
			CreateDate: pcomment[i].CreateDate,
		}
	}
	return
}