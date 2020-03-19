package iface

type IRequest interface{
	GetConnection() IConnection
	GetData()       IMessage
}
