package cnet

import (
	"copaka/iface"
	"fmt"
	"net"
	"time"
)

type Server struct{
	Ip string
	Port int
	Name string
}

func ClientHandel(conn net.Conn,buf []byte,i int)error{
	if _,err := conn.Write(buf[:i]);err != nil{
		return err
	}
	return nil
}

func NewServer(name,ip string,port int)iface.Iserver{
	return &Server{Name:name,
					Port:port,
						Ip:ip}
}

func (p *Server)Start(){
	fmt.Printf("[START]server,listen at IP: %s,Port %d ,is starting\n",p.Ip,p.Port)

	//监听服务器
	listener,err := net.Listen("tcp",fmt.Sprintf("%s:%d",p.Ip,p.Port))
	if err != nil{
		fmt.Printf("[ERROR]Monitoring failed :%s\n",err)
		return
	}

	//监听成功
	fmt.Printf("start copaka server,%s,now listening",p.Name)

	//启动网络连接服务
	for{
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err",err)
			continue
		}

	//简单数据回显业务
		go func() {
			for  {
				buf := make([]byte,512)
				n,err := conn.Read(buf)
				if err != nil{
					fmt.Println("recv buf err",err)
					continue
				}
				if _,err := conn.Write(buf[:n]); err != nil{
					fmt.Println("write buf err",err)
					continue
				}
			}
		}()

	}

}

func (p *Server)Stop(){
	fmt.Println("[STOP] copaka server,name ",p.Name)
}

func (p *Server)Serve(){
	go p.Start()

	time.Sleep(20*time.Second)
}