package controller

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	"sync/atomic"
	"time"
)

var tempChat = map[string][]Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	if user, exist := usersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(user.Id, int64(userIdB))

		atomic.AddInt64(&messageIdSequence, 1)
		curMessage := Message{
			Id:         messageIdSequence,
			Content:    content,
			CreateTime: time.Now().Unix(),
		}

		if messages, exist := tempChat[chatKey]; exist {
			tempChat[chatKey] = append(messages, curMessage)
		} else {
			tempChat[chatKey] = []Message{curMessage}
		}
		c.JSON(consts.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// MessageChat all users have same follow list
func MessageChat(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")

	if user, exist := usersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(user.Id, int64(userIdB))

		c.JSON(consts.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	} else {
		c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
