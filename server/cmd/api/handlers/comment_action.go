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

func CommentAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err != nil {
		log.Fatal("wrong actionType")
	}
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong video_id")
	}
	commentText := c.Query("comment_text")
	commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)

	resp := new(interact.CommentActionResponse)
	resp, err = rpc.CommentAction(context.Background(), &interact.CommentActionRequest{
		Token: token,
		VideoId: videoId,
		ActionType: int32(actionType),
		CommentText: commentText,
		CommentId: commentId,
	})
	if err != nil {
		c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: errno.ConvertErr(err).ErrMsg})
		return
	}

	c.JSON(consts.StatusOK, CommentActionResponse{
		Response: Response{StatusCode: 0},
		Comment: cp2c(resp.Comment),
	})
}

func cp2c(pcomment *interact.Comment) (comment Comment) {
	comment = Comment{
		Id: pcomment.Id,
		User: User{
			Id: pcomment.User.Id,
			Name: pcomment.User.Name,
			FollowCount: pcomment.User.FollowCount,
			FollowerCount: pcomment.User.FollowerCount,
			IsFollow: pcomment.User.IsFollow,
		},
		Content: pcomment.Content,
		CreateDate: pcomment.CreateDate,
	}
	return
}