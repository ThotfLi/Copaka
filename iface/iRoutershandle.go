package iface

//多路由处理器接口
type IRoutersHandle interface{
	AddRouter(msgid uint32,router IRouter)       //添加新的路由
	RouterConverter(request IRequest) //路由转换器

	StartWorkerPool() //启动worker工作池
	SendMsgToTaskQueue(request IRequest) //将消息交给TaskQueue,由worker进行处理
}
