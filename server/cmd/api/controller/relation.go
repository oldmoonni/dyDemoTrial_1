package controller

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/dao"
	"log"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	to_user_id, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong to_user_id")
	}
	action_type, err := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err != nil {
		log.Fatal("wrong action_type")
	}

	duserlock, flag := dao.UserLockInfoByToken(token)
	_, flag2 := dao.UserIsExistById(to_user_id)
	from_id := duserlock.Id

	if flag == true && flag2 == true {
		//actionType==1 关注 actionType==2 取消关注
		if action_type == 1 {
			if dao.Isfollow(from_id, to_user_id) == true {
				c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: "you have followed this user"})
				return
			}
			//生成的id
			node, err := snowflake.NewNode(1)
			if err != nil {
				log.Fatal(err)
			}
			id := int64(node.Generate())
			dao.RelationAdd(id, from_id, to_user_id)
			c.JSON(consts.StatusOK, Response{StatusCode: 0})
		} else {
			if dao.Isfollow(from_id, to_user_id) == false {
				c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: "you have not followed this user"})
				return
			}
			dao.RelationSub(from_id, to_user_id)
			c.JSON(consts.StatusOK, Response{StatusCode: 0})
		}
	} else {
		c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(consts.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(ctx context.Context, c *app.RequestContext) {
	user_id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong user_id")
	}
	token := c.Query("token")
	dusers := dao.GetFollowList(user_id)
	users := u2uplustokenList(dusers, token)

	c.JSON(consts.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}

// FollowerList all users have same follower list
func FollowerList(ctx context.Context, c *app.RequestContext) {
	user_id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong user_id")
	}
	token := c.Query("token")
	dusers := dao.GetFollowerList(user_id)
	users := u2uplustokenList(dusers, token)

	c.JSON(consts.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}

// FriendList all users have same friend list
func FriendList(ctx context.Context, c *app.RequestContext) {
	user_id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong user_id")
	}
	token := c.Query("token")
	dusers := dao.GetFollowerList(user_id)
	users := u2uplustokenList(dusers, token)

	c.JSON(consts.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}
