package service

import (
	"context"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/user/dal"
	user1 "github.com/trial_1/dyDemoTrial_1/server/kitex_gen/user"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"sync"
)

type UserInfoService struct {
	ctx context.Context
}

func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

var WaitGroup sync.WaitGroup
func (s *UserInfoService) UserInfo(req *user1.UserInfoRequest) (user *user1.User, err error) {
	userId := req.UserId
	token := req.Token

	var duserlock dal.DUserLock
	var duser dal.DUser
	var flag, flag2 bool
	WaitGroup.Add(2)
	go func() {
		duserlock, flag = dal.UserLockInfoById(userId)
		println("第一个协程完成")
		WaitGroup.Done()
	}()
	go func() {
		duser, flag2 = dal.UserIsExistById(userId)
		println("第二个协程完成")
		WaitGroup.Done()
	}()
	WaitGroup.Wait()
	//duserlock, flag := dal.UserLockInfoById(userId)
	//duser, flag2 := dal.UserIsExistById(userId)

	if flag == true && flag2 == true && duserlock.Id == duser.Id {
		user, err = u2uplustoken(duser, token)
	} else {
		err = errno.UserNotExistErr
	}
	return
}

func u2uplustoken(duser dal.DUser, token string) (user *user1.User, err error) {
	duserlock, flag := dal.UserLockInfoByToken(token)
	if flag == false {
		err = errno.UserInfoWrongErr
		return
	}
	user = &user1.User{
		Id: duser.Id,
		Name: duser.Name,
		FollowCount: duser.FollowCount,
		FollowerCount: duser.FollowerCount,
		IsFollow: dal.Isfollow(duserlock.Id, duser.Id),
		Avatar: duser.Avatar,
		BackgroundImage: duser.BackgroundImage,
		Signature: duser.Signature,
		TotalFavorited: duser.TotalFavorited,
		WorkCount: duser.WorkCount,
		FavoriteCount: duser.FavoriteCount,
	}
	return
}