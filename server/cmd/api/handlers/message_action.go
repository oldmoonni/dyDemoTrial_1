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
	"time"
)

func MessageAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	toUserId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong to_user_id")
	}
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err != nil {
		log.Fatal("wrong actionType")
	}
	content := c.Query("content")

	_, err = rpc.MessageAction(context.Background(), &social.MessageActionRequest{
		Token:      token,
		ToUserId:   toUserId,
		ActionType: int32(actionType),
		Content:    content,
		CreateTime: time.Now().Unix(),
	})
	if err != nil {
		c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: errno.ConvertErr(err).ErrMsg})
		return
	}

	c.JSON(consts.StatusOK, Response{StatusCode: 0})
}
