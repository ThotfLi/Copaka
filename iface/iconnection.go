package iface

import "net"

type IConnection interface {
	//获得连接的原始socket
	GetConn() net.Conn
	//开启conn开始工作
	Start()
	//关闭conn
	Stop()
	//获得地址连接信息
	GetAddr() net.Addr
	//获得连接id
	GetConnID() uint32

}

type HandFunc func(net.Conn,[]byte,int) error