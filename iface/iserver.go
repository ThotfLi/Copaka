package iface

//服务器

type Iserver interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//开启服务器业务
	Serve()
}