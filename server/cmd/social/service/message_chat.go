package service

import (
	"bufio"
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	dal2 "github.com/trial_1/dyDemoTrial_1/server/cmd/social/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/social"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"log"
	"net"
	"os"
	"strings"
)

type MessageChatService struct {
	ctx context.Context
}

func NewMessageChatService(ctx context.Context) *MessageChatService {
	return &MessageChatService{ctx: ctx}
}

func (s *MessageChatService) MessageChat(req *social.MessageChatRequest) (resp *social.MessageChatResponse, err error) {

	return
}