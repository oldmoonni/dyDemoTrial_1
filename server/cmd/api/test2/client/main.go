package main

import (
	"fmt"
	"net"
)

type Message struct {
	Id         int64  `json:"id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	FromUserId   int64  `json:"from_user_id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

//客户端
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer conn.Close() //关闭连接
	msg := "1,11,12,douyin,2-19"
	_, err = conn.Write([]byte(msg)) //发送数据
	if err != nil {
		return
	}
	//buf := [512]byte{}
	//n, err := conn.Read(buf[:])
	//if err != nil {
	//	fmt.Println("recv failed,err:", err)
	//	return
	//}
	//fmt.Println(string(buf[:n]))
}