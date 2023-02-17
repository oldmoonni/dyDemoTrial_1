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
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/social"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"log"
	"strconv"
)

func FollowerList(ctx context.Context, c *app.RequestContext) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong user_id")
	}
	token := c.Query("token")

	resp := new(social.FollowerListResponse)
	resp, err = rpc.FollowerList(context.Background(), &social.FollowerListRequest{
		UserId: userId,
		Token: token,
	})
	if err != nil {
		c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: errno.ConvertErr(err).ErrMsg})
		return
	}

	c.JSON(consts.StatusOK, UserListResponse{
		Response: 	Response{StatusCode: 0},
		UserList: 	socialup2uList(resp.UserList),
	})
}