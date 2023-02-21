package service

import (
	"context"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/user/dal"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/user"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"golang.org/x/crypto/bcrypt"
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

	duser, flag := dal.UserIsExistByName(username)
	duserlock, flag2 := dal.UserLockInfoById(duser.Id)

	if flag == true && flag2 == true {
		if ComparePwd(duserlock.Password, password) {
			resp = &user.UserLoginResponse{
				UserId: duser.Id,
				Token: duserlock.Token,
			}
		} else {
			err = errno.PasswordErr
		}
	} else {
		err = errno.UserNotExistErr
	}
	return
}

// 比对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	// pwd1：数据库中的密码
	// pwd2：用户输入的密码
	if err != nil {
		return false
	} else {
		return true
	}
}