package cnet

import (
	"copaka/iface"
	"fmt"
	"net"
	"time"
	"copaka/utils"
)

type Server struct{
	IP string
	Port int
	Name string
	Routers iface.IRoutersHandle
}

func NewServer()iface.Iserver{
	utils.InitConfig()
	return &Server{Name:utils.Globalogobject.Name,
		           Port:utils.Globalogobject.Port,
		           IP:utils.Globalogobject.Host,
		           Routers:nil}
}

func (p *Server)Start(){
	fmt.Printf("[START]server,listen at IP: %s,Port %d ,is starting\n",p.IP,p.Port)
	fmt.Println(p.Port,p.Name,p.IP)
	//监听服务器
	listener,err := net.Listen("tcp",fmt.Sprintf("%s:%d",p.IP,p.Port))
	if err != nil{
		fmt.Printf("[ERROR]Monitoring failed :%s\n",err)
		return
	}
	var AutoID uint32
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
		connectioner := NewConnection(conn,AutoID,p.Routers)
		go connectioner.Start()
		AutoID += 1
	}

}

func (p *Server)Stop(){
	fmt.Println("[STOP] copaka server,name ",p.Name)
}

func (p *Server)Serve(){
	go p.Start()

	time.Sleep(20*time.Second)
}

func (p *Server)AddRouter(msgid uint32,router iface.IRouter){
	p.Routers.AddRouter(msgid,router)
}