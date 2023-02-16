// Code generated by hertz generator.

package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/controller"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/handlers"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	//r.GET("/ping", handler.Ping)

	// your code ...
	apiRouter := r.Group("/douyin")
	{
		//基础接口
		apiRouter.GET("/feed/", func(ctx context.Context, c *app.RequestContext) {
			controller.Feed(ctx, c)
			println("/feed:get")
		})
		apiRouter.GET("/user/", handlers.UserInfo)
		apiRouter.POST("/user/register/", handlers.UserRegister)
		apiRouter.POST("/user/login/", handlers.UserLogin)
		//apiRouter.GET("/user/", func(ctx context.Context, c *app.RequestContext) {
		//	controller.UserInfo(ctx, c)
		//	println("/user:get")
		//})
		//apiRouter.POST("/user/register/", func(ctx context.Context, c *app.RequestContext) {
		//	controller.Register(ctx, c)
		//	println("/user/register:post")
		//})
		//apiRouter.POST("/user/login/", func(ctx context.Context, c *app.RequestContext) {
		//	controller.Login(ctx, c)
		//	println("/user/login:post")
		//})
		apiRouter.POST("/publish/action/", func(ctx context.Context, c *app.RequestContext) {
			println("开始运行")
			controller.Publish(ctx, c)
			println("/publish/action:post")
		})
		apiRouter.GET("/publish/list/", func(ctx context.Context, c *app.RequestContext) {
			controller.PublishList(ctx, c)
			println("/publish/list:get")
		})

		//互动接口
		apiRouter.POST("/favorite/action/", func(ctx context.Context, c *app.RequestContext) {
			controller.FavoriteAction(ctx, c)
			println("/favorite/action:post")
		})
		apiRouter.GET("/favorite/list/", func(ctx context.Context, c *app.RequestContext) {
			controller.FavoriteList(ctx, c)
			println("/favorite/list:get")
		})
		apiRouter.POST("/comment/action/", func(ctx context.Context, c *app.RequestContext) {
			controller.CommentAction(ctx, c)
			println("/comment/action:post")
		})
		apiRouter.GET("/comment/list/", func(ctx context.Context, c *app.RequestContext) {
			controller.CommentList(ctx, c)
			println("/comment/list:get")
		})

		//社交接口
		apiRouter.POST("/relation/action/", func(ctx context.Context, c *app.RequestContext) {
			controller.RelationAction(ctx, c)
			println("/relation/action:post")
		})
		apiRouter.GET("/relation/follow/list/", func(ctx context.Context, c *app.RequestContext) {
			controller.FollowList(ctx, c)
			println("/relation/follow/list:get")
		})
		apiRouter.GET("/relation/follower/list/", func(ctx context.Context, c *app.RequestContext) {
			controller.FollowerList(ctx, c)
			println("/relation/follower/list:get")
		})
		apiRouter.GET("/relation/friend/list/", func(ctx context.Context, c *app.RequestContext) {
			controller.FriendList(ctx, c)
			println("/relation/friend/list:get")
		})
		apiRouter.GET("/message/chat/", func(ctx context.Context, c *app.RequestContext) {
			controller.MessageChat(ctx, c)
			println("/message/chat:get")
		})
		apiRouter.POST("/message/action/", func(ctx context.Context, c *app.RequestContext) {
			controller.MessageAction(ctx, c)
			println("/message/action:post")
		})

	}

}
