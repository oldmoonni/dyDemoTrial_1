// Code generated by Kitex v0.4.4. DO NOT EDIT.

package socialservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	social "github.com/trial_1/dyDemoTrial_1/server/kitex_gen/social"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationAction(ctx context.Context, req *social.RelationActionRequest, callOptions ...callopt.Option) (r *social.RelationActionResponse, err error)
	FollowList(ctx context.Context, req *social.FollowListRequest, callOptions ...callopt.Option) (r *social.FollowListResponse, err error)
	FollowerList(ctx context.Context, req *social.FollowerListRequest, callOptions ...callopt.Option) (r *social.FollowerListResponse, err error)
	FriendList(ctx context.Context, req *social.FriendListRequest, callOptions ...callopt.Option) (r *social.FriendListResponse, err error)
	MessageChat(ctx context.Context, req *social.MessageChatRequest, callOptions ...callopt.Option) (r *social.MessageChatResponse, err error)
	MessageAction(ctx context.Context, req *social.MessageActionRequest, callOptions ...callopt.Option) (r *social.MessageActionResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kSocialServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kSocialServiceClient struct {
	*kClient
}

func (p *kSocialServiceClient) RelationAction(ctx context.Context, req *social.RelationActionRequest, callOptions ...callopt.Option) (r *social.RelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, req)
}

func (p *kSocialServiceClient) FollowList(ctx context.Context, req *social.FollowListRequest, callOptions ...callopt.Option) (r *social.FollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowList(ctx, req)
}

func (p *kSocialServiceClient) FollowerList(ctx context.Context, req *social.FollowerListRequest, callOptions ...callopt.Option) (r *social.FollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowerList(ctx, req)
}

func (p *kSocialServiceClient) FriendList(ctx context.Context, req *social.FriendListRequest, callOptions ...callopt.Option) (r *social.FriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FriendList(ctx, req)
}

func (p *kSocialServiceClient) MessageChat(ctx context.Context, req *social.MessageChatRequest, callOptions ...callopt.Option) (r *social.MessageChatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageChat(ctx, req)
}

func (p *kSocialServiceClient) MessageAction(ctx context.Context, req *social.MessageActionRequest, callOptions ...callopt.Option) (r *social.MessageActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageAction(ctx, req)
}
