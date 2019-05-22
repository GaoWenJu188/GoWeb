package main

import "zinx/Zinx/net"

func main(){
	//创建一个zinx server 对象
	s:=net.NewServer("zinx v0.1")
	//让server对象 启动服务
	s.Serve()
}
