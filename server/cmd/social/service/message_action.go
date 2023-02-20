package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/social/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/social"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"log"
)

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

func (s *MessageActionService) MessageAction(req *social.MessageActionRequest) (err error) {
	token := req.Token
	toUserId := req.ToUserId
	//actionType := req.ActionType
	content := req.Content
	createTime := req.CreateTime

	//生成的id
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}
	messageId := int64(node.Generate())

	duserlock, flag := dal.UserLockInfoByToken(token)
	if flag == false {
		err = errno.UserNotExistErr
		return
	}
	fromUserId := duserlock.Id

	dal.MessageInsert(messageId, fromUserId, toUserId, content, createTime)

	return
}