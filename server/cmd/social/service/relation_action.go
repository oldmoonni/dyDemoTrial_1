package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	dal2 "github.com/trial_1/dyDemoTrial_1/server/cmd/social/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/social"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"log"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (s *RelationActionService) RelationAction(req *social.RelationActionRequest) (err error) {
	token := req.Token
	toUserId := req.ToUserId
	actionType := req.ActionType

	duserlock, flag := dal2.UserLockInfoByToken(token)
	_, flag2 := dal2.UserIsExistById(toUserId)
	fromId := duserlock.Id

	if flag == true && flag2 == true {
		//actionType==1 关注 actionType==2 取消关注
		if actionType == 1 {
			if dal2.Isfollow(fromId, toUserId) == true {
				err = errno.HaveDoneErr
				return
			}
			//生成的id
			node, err := snowflake.NewNode(1)
			if err != nil {
				log.Fatal(err)
			}
			id := int64(node.Generate())
			dal2.RelationAdd(id, fromId, toUserId)
			return nil
		} else {
			if dal2.Isfollow(fromId, toUserId) == false {
				err = errno.HaveDoneErr
				return
			}
			dal2.RelationSub(fromId, toUserId)
			return nil
		}
	} else {
		err = errno.UserNotExistErr
		return
	}
}