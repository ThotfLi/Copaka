package cnet

import (
	"copaka/iface"
	"copaka/utils"
	"fmt"
)

type RoutersHandle struct{
	apis map[uint32]iface.IRouter  //每个uin32 对应一个路由
	WorkerPoolSize uint32
	TaskQueue      []chan iface.IRequest
}
func NewRoutersHandle()iface.IRoutersHandle{
	return &RoutersHandle{apis:make(map[uint32]iface.IRouter),
						WorkerPoolSize:utils.Globalogobject.WorkerPoolSize,
						TaskQueue:make([]chan iface.IRequest,utils.Globalogobject.WorkerPoolSize)}
}

func(p *RoutersHandle)AddRouter(msgid uint32,router iface.IRouter){
	//判断路由是否存在
	if _,ok := p.apis[msgid]; ok == true{
		panic(fmt.Sprintln("已存在路由id：",msgid))
	}
	//添加新路由
	p.apis[msgid] = router
}

func(p *RoutersHandle)RouterConverter(request iface.IRequest){
	router,ok := p.apis[request.GetData().GetID()]
	if ok != true{
		panic(fmt.Sprintln("路由不存在 id：",request.GetData().GetID()))
	}

	//执行路由任务
	func(request iface.IRequest){
		router.PreHandel(request)
		router.Handel(request)
		router.PostHandel(request)
	}(request)

}
func(p *RoutersHandle)StartWorkerPool(){
	fmt.Println("[START] Worker Pool")
	for i := 0;i < int(utils.Globalogobject.WorkerPoolSize);i++ {
		p.TaskQueue[i] = make(chan iface.IRequest,utils.Globalogobject.MaxWorkerTaskLen)
		p.StartOneWorker(i,p.TaskQueue[i])
	}
}
func(p *RoutersHandle)SendMsgToTaskQueue(request iface.IRequest){
	//根据ConnID来分配当前的连接应该由哪个worker负责处理
	//轮询的平均分配法则
	//得到需要处理此条连接的workerID
	workerID := request.GetConnection().GetConnID() % p.WorkerPoolSize
	fmt.Println("Add ConnID=", request.GetConnection().GetConnID()," request msgID=", request.GetData().GetID(), "to workerID=", workerID)
	//将请求消息发送给任务队列
	p.TaskQueue[workerID] <- request
}
func (p *RoutersHandle) StartOneWorker(workerID int, taskQueue chan iface.IRequest) {
	fmt.Println("Worker ID = ", workerID, " is started.")
	//不断的等待队列中的消息
	for {
		select {
		//有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			p.RouterConverter(request)
		}
	}
}