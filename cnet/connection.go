package cnet

import (
	"copaka/iface"
	"copaka/utils"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net"
)

type Connection struct{
	//当前连接套接字
	Conn net.Conn
	//当前连接id
	ConnID   uint32
	//当前连接状态是否已关闭
	isClosed bool

	//告知该链接已经退出/停止的channel
	ExitBuffChan chan bool

	//路由
	Routers iface.IRoutersHandle
	msgChan chan []byte
}

func NewConnection(conn net.Conn,id uint32,routers iface.IRoutersHandle)iface.IConnection{
	return &Connection{Conn:conn,
		               ConnID:id,
		               isClosed:false,
		               ExitBuffChan:make(chan bool,1),
		               Routers:routers,
		               msgChan:make(chan []byte)}
}

func(p *Connection)GetConn()net.Conn{
	return p.Conn
}

func(p *Connection)StartReader(){
	fmt.Println("reader goroutine is running")
	defer fmt.Println("exit reader conn is ",p.Conn.RemoteAddr().String())
	defer p.Stop()

	pack := NewPack()
	for {
		headlen := make([]byte,pack.GetHeadLen())
		_,err := io.ReadFull(p.Conn,headlen)
		if err != nil{
			fmt.Println("recv buf err ",err)
			p.ExitBuffChan <- true
			break
		}
		msg,err := pack.UnPack(headlen)
		if err != nil {
			fmt.Println("message err ",err)
			break
		}
		data := make([]byte,msg.GetDataLen())
		_,err = io.ReadFull(p.Conn,data)
		if err != nil {
			fmt.Println("message err",err)
			p.ExitBuffChan <- true
			break
		}
		msg.SetData(data)


		//msg := NewMessages(buf[:cnt],1)

		req := NewRequest(p,msg)

		//不强制使用任务池
		if utils.Globalogobject.WorkerPoolSize >0 {
			go p.Routers.SendMsgToTaskQueue(req)
		}else {
			go p.Routers.RouterConverter(req)
		}



	}
}

func(p *Connection)StartWriter(){
	fmt.Println("[START]Writer goroutine is running")
	select {
	case data := <- p.msgChan:
		if _,err := p.Conn.Write(data);err != nil{
			fmt.Println("[ERROR] Writer Msg Err:",err)
		}
	case <-p.ExitBuffChan:
		return
	}
}

func(p *Connection)Start(){
	go p.StartReader()
	go p.StartWriter()

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

func(p *Connection)SendMsg(id uint32,data []byte)error{
	if p.isClosed == true{
		return errors.New("Connection closed when send msg")
	}

	dp := NewPack()
	msgByte,err :=dp.Pack(NewMessages(data,id))
	if err != nil {
		fmt.Println("Pack error msg id = ",id)
		return errors.New("Pack eror msg")
	}
	p.msgChan <- msgByte
	return nil
}