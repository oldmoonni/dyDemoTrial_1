package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/trial_1/dyDemoTrial_1/server/kitex_gen/user/userservice"
	"github.com/trial_1/dyDemoTrial_1/server/pkg/middleware"
	"log"
	"net"
	"time"
)

func main() {
	//r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	//if err != nil {
	//	log.Fatal(err)
	//	//panic(err)
	//}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8801")
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userservice"}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                             // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithReadWriteTimeout(time.Hour),
		//server.WithMuxTransport(),                                          // Multiplex
		//server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		//server.WithRegistry(r),                                             // registry
		)
	//svr := userservice.NewServer(new(UserServiceImpl),
	//	server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
