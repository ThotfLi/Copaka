package cnet

import (
	"bytes"
	"copaka/iface"
	"copaka/utils"
	"encoding/binary"
	"github.com/pkg/errors"
)

type Packet struct{

}
func NewPack()iface.IPacket{
	return &Packet{}
}

func (p *Packet)Pack(message iface.IMessage) ([]byte,error){
	//写入的数据格式 DATALEN DATAID DATA
	iobuf := bytes.NewBuffer([]byte{})

	if err := binary.Write(iobuf,binary.LittleEndian,message.GetDataLen());err != nil{
		return nil,err
	}

	if err := binary.Write(iobuf,binary.LittleEndian,message.GetID());err != nil{
		return nil,err
	}

	if err := binary.Write(iobuf,binary.LittleEndian,message.GetData());err != nil{
		return nil,err
	}
	return iobuf.Bytes(),nil
}

func (p *Packet)UnPack(buf []byte)(iface.IMessage,error){
	iobuf := bytes.NewBuffer(buf)

	newmsg := &Message{}
	if err := binary.Read(iobuf,binary.LittleEndian,&newmsg.DataLen);err != nil || newmsg.GetDataLen() <= 0{
		return nil ,err
	}

	if err := binary.Read(iobuf,binary.LittleEndian,&newmsg.ID);err != nil{
		return nil,err
	}
	if utils.Globalogobject.MaxPacketSiz < newmsg.DataLen && utils.Globalogobject.MaxPacketSiz >0 {
		return nil,errors.New("datalen error")
	}

	return newmsg,nil
}

func (p *Packet)GetHeadLen()uint32{
	return 8
}