// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/interact"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/interact/interactservice"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"
	"log"
)

var interactClient interactservice.Client

func initInteractRpc() {
	//r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	//if err != nil {
	//	log.Fatal(err)
	//	//panic(err)
	//}

	//var c, err = userservice.NewClient(
	//	"user",
	//	client.WithHostPorts("127.0.0.1:8801"),
	//	client.WithMiddleware(middleware.CommonMiddleware),
	//	client.WithInstanceMW(middleware.ClientMiddleware),
	//	client.WithMuxConnection(1),                    // mux
	//	client.WithRPCTimeout(3*time.Second),           // rpc timeout
	//	client.WithConnectTimeout(50*time.Millisecond), // conn timeout
	//	//client.WithFailureRetry(retry.NewFailurePolicy()), // retry
	//	//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
	//	//client.WithResolver(r),                            // resolver
	//)
	//if err != nil {
	//	log.Fatal(err)
	//	//panic(err)
	//}
	//userClient = c

	c, err := interactservice.NewClient("interactservice", client.WithHostPorts("127.0.0.1:8803"))
	if err != nil {
		log.Fatal(err)
	}
	//time.Sleep(time.Second)
	interactClient = c
}

func FavoriteAction(ctx context.Context, req *interact.FavoriteActionRequest) (resp *interact.FavoriteActionResponse, err error) {
	resp, err = interactClient.FavoriteAction(ctx, req)
	if err != nil {
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		err = errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		return
	}
	return
}

func FavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (resp *interact.FavoriteListResponse, err error) {
	resp, err = interactClient.FavoriteList(ctx, req)
	if err != nil {
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		err = errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		return
	}
	return
}

func CommentAction(ctx context.Context, req *interact.CommentActionRequest) (resp *interact.CommentActionResponse, err error) {
	resp, err = interactClient.CommentAction(ctx, req)
	if err != nil {
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		err = errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		return
	}
	return
}

func CommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	resp, err = interactClient.CommentList(ctx, req)
	if err != nil {
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		err = errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		return
	}
	return
}