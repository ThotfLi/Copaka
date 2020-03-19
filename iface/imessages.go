package iface

type IMessage interface{
	//获取数据内容
	GetData() []byte
	//获取数据长度
	GetDataLen() uint32
	//获取数据id
	GetID() uint32

	//设置内容
	SetData([]byte)
	SetID(uint32)
}