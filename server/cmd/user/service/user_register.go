package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/user/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/user"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"golang.org/x/crypto/bcrypt"
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
		sPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		WaitGroup.Add(3)
		go func() {
			dal.UserInsert(id, username)
			println("第一个协程完成")
			WaitGroup.Done()
		}()
		go func() {
			dal.UserLockInsert(id, username, string(sPwd))
			println("第二个协程完成")
			WaitGroup.Done()
		}()
		go func() {
			dal.DrecomInsert(username + string(sPwd))
			println("第三个协程完成")
			WaitGroup.Done()
		}()
		WaitGroup.Wait()
		//dal.UserInsert(id, username)
		//dal.UserLockInsert(id, username, string(sPwd))
		//dal.DrecomInsert(username + string(sPwd))

		resp = &user.UserRegisterResponse{
			UserId: id,
			Token: username + string(sPwd),
		}
	}
	return
}