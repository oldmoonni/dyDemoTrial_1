package service

import (
	"context"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/user/dal"
	user1 "github.com/trial_1/dyDemoTrial_1/server/kitex_gen/user"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"log"
)

type UserInfoService struct {
	ctx context.Context
}

func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

func (s *UserInfoService) UserInfo(req *user1.UserInfoRequest) (user *user1.User, err error) {
	userId := req.UserId
	token := req.Token

	duserlock, flag := dal.UserLockInfoById(userId)
	duser, flag2 := dal.UserIsExistById(userId)

	if flag == true && flag2 == true && duserlock.Id == duser.Id {
		user = u2uplustoken(duser, token)
	} else {
		err = errno.UserNotExistErr
	}
	return
}

func u2uplustoken(duser dal.DUser, token string) (user *user1.User) {
	duserlock, flag := dal.UserLockInfoByToken(token)
	if flag == false {
		log.Fatal("wrong user information")
	}
	user = &user1.User{
		Id: duser.Id,
		Name: duser.Name,
		FollowCount: duser.FollowCount,
		FollowerCount: duser.FollowerCount,
		IsFollow: dal.Isfollow(duserlock.Id, duser.Id),
	}
	return
}