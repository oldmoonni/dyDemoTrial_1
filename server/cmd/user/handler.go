package main

import (
	"context"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/user/pack"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/user/service"
	user "github.com/trial_1/dyDemoTrial_1/server/kitex_gen/user"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// TODO: Your code here...

	resp = new(user.UserRegisterResponse)

	if len(req.Name) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	resp, err = service.NewUserRegisterService(ctx).UserRegister(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil

}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...

	resp = new(user.UserLoginResponse)

	if len(req.Name) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, errno.PasswordErr
	}

	resp, err = service.NewUserLoginService(ctx).UserLogin(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return 	resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserInfoResponse)

	if len(req.Token) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, errno.PasswordErr
	}

	user, err := service.NewUserInfoService(ctx).UserInfo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.User = user
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, err
}
