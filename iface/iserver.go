package iface

//服务器

type Iserver interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//开启服务器业务
	Serve()
	//路由功能
	AddRouter(msgid uint32,router IRouter)
}