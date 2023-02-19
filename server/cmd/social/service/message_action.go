package service

import (
	"context"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/social"
)

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

func (s *MessageActionService) MessageAction(req *social.MessageActionRequest) (err error) {
	return
}