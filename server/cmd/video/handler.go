package main

import (
	"context"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/video/service"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/video/pack"
	video "github.com/trial_1/dyDemoTrial_1/server/kitex_gen/video"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// VideoFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoFeed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...

	resp = new(video.FeedResponse)

	if len(req.Token) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	resp, err = service.NewVideoFeedService(ctx).VideoFeed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// VideoPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoPublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	resp = new(video.PublishListResponse)

	if len(req.Token) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	resp, err = service.NewVideoPublishListService(ctx).VideoPublishList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
