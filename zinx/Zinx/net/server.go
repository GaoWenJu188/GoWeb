/**
server模块的实现层
 */
 package net

import (
	"zinx/Zinx/ziface"
	"fmt"
	"net"
)

type Server struct {

 	//服务器ip
 	IPVersion string
 	IP string
 	//服务器port
 	Port int
 	//服务器明名称
 	Name string
 }
 //初始化的new方法
 func NewServer (name string) ziface.IServer{
 	s:=&Server{
 		Name:name,
 		IPVersion:"tcp4",
 		IP:"0.0.0.0",
 		Port:8999,
	}
	return s
 }
//启动服务器
func (s *Server) Start() {
	fmt.Printf("[start] Server Listenner at IP:%s,Port:%d, is starting ...",s.IP,s.Port)
	//1 创建套接字  ：得到一个TCP的addr
 	addr,err:= net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.IP,s.Port))
	if err!=nil{
		fmt.Println("resolve tcp addr error:",err)
		return
	}
	//2 监听服务器地址
	listenner,err:= net.ListenTCP(s.IPVersion,addr)
	if err!=nil{
		fmt.Println("listen",s.IPVersion,"err",err)
		return
	}
	//3 阻塞等待客户端发送请求，
	go func(){
		for {
			//阻塞等待客户端请求,
			conn,err:= listenner.Accept()
			if err!=nil{
				fmt.Println("Accept err:",err)
				continue
			}
			//此时conn就已经和对端客户端连接
			go func(){
				for{
					buf:= make([]byte,512)
					cnt,err:= conn.Read(buf)
					if err!=nil{
						fmt.Println("recv buf err:",err)
						break
					}
					fmt.Printf("recv client buf %s, cnt = %d\n", buf, cnt)

					//回显功能
					if _,err:=conn.Write(buf[:cnt]);err!=nil{
						fmt.Println("write back buf err:",err)
						continue
					}
				}
			}()
		}
	}()
}
//停止服务器
func (s *Server) Stop() {
	//todo 将一些资源进行回收....
}
//运行服务器
func (s *Server )Serve() {
	//启动server的监听功能
	s.Start()//并不希望他永久的阻塞

	//TODO  做一些其他的扩展
	//阻塞//告诉CPU不再需要处理的，节省cpu资源
	select{}
}