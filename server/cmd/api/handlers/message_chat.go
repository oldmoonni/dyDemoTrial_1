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

func MessageChat(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	toUserId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong to_user_id")
	}
	preMsgTime, err := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)
	if err != nil {
		log.Fatal("wrong pre_msg_time")
	}

	resp := new(social.MessageChatResponse)
	resp, err = rpc.MessageChat(context.Background(), &social.MessageChatRequest{
		Token:      token,
		ToUserId:   toUserId,
		PreMsgTime: preMsgTime,
	})
	if err != nil {
		c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: errno.ConvertErr(err).ErrMsg})
		return
	}

	messages := mp2mList(resp.MessageList)
	c.JSON(consts.StatusOK, MessageChatResponse{
		MessageList: messages,
		Response:    Response{
			StatusCode: 0,
		},
	})
}

func mp2mList(smessages []*social.Message) (messages []Message) {
	messages = make([]Message, len(smessages))
	for i := range smessages{
		messages[i] = Message{
			Id:         smessages[i].Id,
			ToUserId:   smessages[i].ToUserId,
			FromUserId: smessages[i].FromUserId,
			Content:    smessages[i].Content,
			CreateTime: smessages[i].CreateTime,
		}
	}
	return
}