package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/user/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/user"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"log"
)

type UserRegisterService struct {
	ctx context.Context
}

func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{ctx: ctx}
}

func (s *UserRegisterService) UserRegister(req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	username := req.Name
	password := req.Password

	token := username + password

	//生成的id
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}
	id := int64(node.Generate())

	_, flag := dal.UserIsExistByName(username)

	if flag == true {
		err = errno.UserAlreadyExistErr
	} else {
		//将新用户加入数据库
		dal.UserInsert(id, username)
		dal.UserLockInsert(id, username, password)
		dal.DrecomInsert(token)

		resp = &user.UserRegisterResponse{
			UserId: id,
			Token: token,
		}
	}
	return
}