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

package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/rpc"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/user"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		c.JSON(consts.StatusOK, UserRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: "wrong Param"},
		})
		return
	}

	resp, err := rpc.UserRegister(context.Background(), &user.UserRegisterRequest{
		Name: username,
		Password: password,
	})
	if err != nil {
		c.JSON(consts.StatusOK, UserRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: errno.ConvertErr(err).ErrMsg},
		})
		return
	}

	c.JSON(consts.StatusOK, UserRegisterResponse{
		Response: Response{StatusCode: 0},
		UserId:   resp.UserId,
		Token:    resp.Token,
	})
}
