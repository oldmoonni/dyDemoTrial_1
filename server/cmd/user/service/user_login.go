package service

import (
	"context"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/user/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/user"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
)

type UserLoginService struct {
	ctx context.Context
}

func NewUserLoginService(ctx context.Context) *UserLoginService {
	return &UserLoginService{ctx: ctx}
}

func (s *UserLoginService) UserLogin(req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	username := req.Name
	password := req.Password

	token := username + password

	duser, flag := dal.UserIsExistByName(username)
	duserlock, flag2 := dal.UserLockInfoById(duser.Id)

	if flag == true && flag2 == true {
		if duserlock.Password == password {
			resp = &user.UserLoginResponse{
				UserId: duser.Id,
				Token: token,
			}
		} else {
			err = errno.PasswordErr
		}
	} else {
		err = errno.UserNotExistErr
	}
	return
}