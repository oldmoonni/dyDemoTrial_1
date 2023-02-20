package main

import (
	"context"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/social/pack"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/social/service"
	social "github.com/trial_1/dyDemoTrial_1/server/kitex_gen/social"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
)

// SocialServiceImpl implements the last service interface defined in the IDL.
type SocialServiceImpl struct{}

// RelationAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) RelationAction(ctx context.Context, req *social.RelationActionRequest) (resp *social.RelationActionResponse, err error) {
	// TODO: Your code here...
	resp = new(social.RelationActionResponse)

	if len(req.Token) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FollowList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FollowList(ctx context.Context, req *social.FollowListRequest) (resp *social.FollowListResponse, err error) {
	// TODO: Your code here...
	resp = new(social.FollowListResponse)

	if len(req.Token) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	resp, err = service.NewFollowListService(ctx).FollowList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FollowerList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FollowerList(ctx context.Context, req *social.FollowerListRequest) (resp *social.FollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new(social.FollowerListResponse)

	if len(req.Token) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	resp, err = service.NewFollowerListService(ctx).FollowerList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FriendList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FriendList(ctx context.Context, req *social.FriendListRequest) (resp *social.FriendListResponse, err error) {
	// TODO: Your code here...
	resp = new(social.FriendListResponse)

	if len(req.Token) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	resp, err = service.NewFriendListService(ctx).FriendList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// MessageChat implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) MessageChat(ctx context.Context, req *social.MessageChatRequest) (resp *social.MessageChatResponse, err error) {
	// TODO: Your code here...
	resp = new(social.MessageChatResponse)

	if len(req.Token) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	resp, err = service.NewMessageChatService(ctx).MessageChat(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// MessageAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) MessageAction(ctx context.Context, req *social.MessageActionRequest) (resp *social.MessageActionResponse, err error) {
	// TODO: Your code here...
	resp = new(social.MessageActionResponse)

	if len(req.Token) == 0 || len(req.Content) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	err = service.NewMessageActionService(ctx).MessageAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}