package iface

type IPacket interface{
	Pack(message IMessage) ([]byte,error)
	UnPack([]byte)       (IMessage,error)
	GetHeadLen()uint32
}
