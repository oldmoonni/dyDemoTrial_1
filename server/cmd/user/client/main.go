package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	user2 "github.com/trial_1/dyDemoTrial_1/server/kitex_gen/user"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/user/userservice"
	"log"
	"time"
)

func main()  {
	for {
		c, err := userservice.NewClient("user", client.WithHostPorts("127.0.0.1:8801"))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
		req := &user2.UserLoginRequest{Name: "zhangsan", Password: "123456"}
		resp, err := c.UserLogin(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
