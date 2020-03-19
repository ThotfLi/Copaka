package cnet

import "copaka/iface"

type Request struct{
	conn    iface.IConnection
	Message iface.IMessage
}

func NewRequest(conn iface.IConnection,message iface.IMessage)iface.IRequest{
	return &Request{
		conn: conn,
		Message: message,
	}
}

func(p *Request)GetConnection()iface.IConnection{
	return p.conn
}
func(p *Request)GetData()iface.IMessage{
	return p.Message
}
