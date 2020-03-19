package cnet

import "copaka/iface"

type Message struct{
	Data    []byte    //消息内容
	DataLen uint32
	ID      uint32
}

func NewMessages(buf []byte,id uint32)iface.IMessage{
	m := &Message{}

	m.SetData(buf)
	m.SetID(id)
	m.DataLen = uint32(len(buf))
	return m
}

func(p *Message)GetData()[]byte{return p.Data}

func(p *Message)GetDataLen()uint32{return p.DataLen}

func(p *Message)GetID()uint32{return p.ID}

func(p *Message)SetData(data []byte){p.Data = data}

func(p *Message)SetID(id uint32){p.ID = id}