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

var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	//生成的id
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}
	id := int64(node.Generate())

	_, flag := dao.UserIsExistByName(username)

	if flag == true {
		c.JSON(consts.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		//将新用户加入数据库
		//
		//
		//
		//应该是原子的
		dao.UserInsert(id, username)
		dao.UserLockInsert(id, username, password)
		dao.DrecomInsert(token)

		c.JSON(consts.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   id,
			Token:    token,
		})
	}
}

func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	duser, flag := dao.UserIsExistByName(username)
	duserlock, flag2 := dao.UserLockInfoById(duser.Id)

	if flag == true && flag2 == true {
		if duserlock.Password == password {
			c.JSON(consts.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   duser.Id,
				Token:    token,
			})
		} else {
			c.JSON(consts.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "wrong password"},
			})
		}
	} else {
		c.JSON(consts.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Fatal("wrong user_id")
	}
	token := c.Query("token")

	duserlock, flag := dao.UserLockInfoById(id)
	duser, flag2 := dao.UserIsExistById(id)

	if flag == true && flag2 == true && duserlock.Id == duser.Id {
		c.JSON(consts.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     u2uplustoken(duser, token),
		})
	} else {
		c.JSON(consts.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

//func UserInfo(ctx context.Context, c *app.RequestContext) {
//	id ,err := strconv.ParseInt(c.Query("user_id"), 10, 64)
//	if err != nil {
//		log.Fatal("wrong user_id")
//	}
//	token := c.Query("token")
//
//	duserlock, flag := dao.UserLockInfoById(id)
//	duser, flag2 := dao.UserIsExistById(id)
//
//	if flag == true && flag2 == true && duserlock.Id == duser.Id {
//		if token == duserlock.Name + duserlock.Password {
//			c.JSON(consts.StatusOK, UserResponse{
//				Response: Response{StatusCode: 0},
//				User:     u2u(duser),
//			})
//		} else {
//			c.JSON(consts.StatusOK, UserResponse{
//				Response: Response{StatusCode: 1, StatusMsg: "worng User information"},
//			})
//		}
//	} else {
//		c.JSON(consts.StatusOK, UserResponse{
//			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
//		})
//	}
//}

func u2u(duser dao.DUser) (user User) {
	user.Id = duser.Id
	user.Name = duser.Name
	user.FollowCount = duser.FollowCount
	user.FollowerCount = duser.FollowerCount
	return
}

//修改后
func u2uplustoken(duser dao.DUser, token string) (user User) {
	user.Id = duser.Id
	user.Name = duser.Name
	user.FollowCount = duser.FollowCount
	user.FollowerCount = duser.FollowerCount
	duserlock, flag := dao.UserLockInfoByToken(token)
	if flag == false {
		log.Fatal("wrong user information")
	}
	user = User{
		Id: duser.Id,
		Name: duser.Name,
		FollowCount: duser.FollowCount,
		FollowerCount: duser.FollowerCount,
		IsFollow: dao.Isfollow(duserlock.Id, duser.Id),
		Avatar: duser.Avatar,
		BackgroundImage: duser.BackgroundImage,
		Signature: duser.Signature,
		TotalFavorited: duser.TotalFavorited,
		WorkCount: duser.WorkCount,
		FavoriteCount: duser.FavoriteCount,
	}
	return
}

func u2uplustokenList(dusers []dao.DUser, token string) (users []User) {
	users = make([]User, len(dusers))
	for i := range dusers {
		users[i].Id = dusers[i].Id
		users[i].Name = dusers[i].Name
		users[i].FollowCount = dusers[i].FollowCount
		users[i].FollowerCount = dusers[i].FollowerCount
		duserlock, flag := dao.UserLockInfoByToken(token)
		if flag == false {
			log.Fatal("wrong user information")
		}
		users[i].IsFollow = dao.Isfollow(duserlock.Id, dusers[i].Id)
	}

	return
}
