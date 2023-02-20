package service

import (
	"context"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/social/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/social"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
)

type MessageChatService struct {
	ctx context.Context
}

func NewMessageChatService(ctx context.Context) *MessageChatService {
	return &MessageChatService{ctx: ctx}
}

func (s *MessageChatService) MessageChat(req *social.MessageChatRequest) (resp *social.MessageChatResponse, err error) {
	token := req.Token
	toUserId := req.ToUserId
	preMsgTime := req.PreMsgTime

	duserlock, flag := dal.UserLockInfoByToken(token)
	if flag == false {
		err = errno.UserNotExistErr
		return
	}
	fromUserId := duserlock.Id
	messages := dal.MessageSearch(fromUserId, toUserId, preMsgTime)
	resp = &social.MessageChatResponse{
		MessageList: m2mpList(messages),
	}

	return resp, nil
}

func m2mpList(dmessages []dal.DMessage) (messages []*social.Message) {
	messages = make([]*social.Message, len(dmessages))
	for i := range dmessages {
		messages[i] = &social.Message{
			Id:         dmessages[i].Message_id,
			ToUserId:   dmessages[i].To_user_id,
			FromUserId: dmessages[i].From_user_id,
			Content:    dmessages[i].Content,
			CreateTime: dmessages[i].Create_time,
		}
	}
	return
}