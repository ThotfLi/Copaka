package cnet

import (
	"copaka/iface"
	"fmt"
	"net"
)

type Connection struct{
	//当前连接套接字
	Conn net.Conn
	//当前连接id
	ConnID   uint32
	//当前连接状态是否已关闭
	isClosed bool
	//该链接的处理API
	handleAPI iface.HandFunc
	//告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
}

func NewConnection(conn net.Conn,id uint32,callback iface.HandFunc)iface.IConnection{
	return &Connection{Conn:conn,
						ConnID:id,
							isClosed:false,
								ExitBuffChan:make(chan bool,1),
									handleAPI:callback}
}

func(p *Connection)GetConn()net.Conn{
	return p.Conn
}

func(p *Connection)StartReader(){
	fmt.Println("reader goroutine is running")
	defer fmt.Println("exit reader conn is ",p.Conn.RemoteAddr().String())
	defer p.Stop()

	for {
		buf := make([]byte,512)
		cnt,err := p.Conn.Read(buf)
		if err != nil{
			fmt.Println("recv buf err ",err)
			p.ExitBuffChan <- true
		}

		if err := p.handleAPI(p.Conn,buf,cnt); err != nil{
			fmt.Println("connID",p.ConnID,"handle is err",err)
			p.ExitBuffChan <- true
		}

	}
}

func(p *Connection)Start(){
	go p.StartReader()

	for {
		select {
		case <-p.ExitBuffChan:
			return
		}
	}
}

func(p *Connection)Stop(){
	if p.isClosed == true{
		return
	}
	p.isClosed = true

	p.Conn.Close()

	p.ExitBuffChan <- true

	close(p.ExitBuffChan)
}

func(p *Connection)GetAddr()net.Addr{
	return p.Conn.RemoteAddr()
}

func(p *Connection)GetConnID()uint32{
	return p.ConnID
}

