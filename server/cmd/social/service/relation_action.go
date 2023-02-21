package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	dal2 "github.com/trial_1/dyDemoTrial_1/server/cmd/social/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/social"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"log"
	"sync"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

var WaitGroup sync.WaitGroup
func (s *RelationActionService) RelationAction(req *social.RelationActionRequest) (err error) {
	token := req.Token
	toUserId := req.ToUserId
	actionType := req.ActionType

	var duserlock dal2.DUserLock
	var flag, flag2 bool
	WaitGroup.Add(2)
	go func() {
		duserlock, flag = dal2.UserLockInfoByToken(token)
		println("第一个协程完成")
		WaitGroup.Done()
	}()
	go func() {
		_, flag2 = dal2.UserIsExistById(toUserId)
		println("第二个协程完成")
		WaitGroup.Done()
	}()
	WaitGroup.Wait()

	fromId := duserlock.Id

	if flag == true && flag2 == true {

		if fromId == toUserId {
			err = errno.IllegalOperationErr
			return
		}

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