package iface

type IRouter interface {
	//处理业务之前的钩子
	PreHandel(request IRequest)
	//业务处理主函数
	Handel(request IRequest)
	//业务处理之后的钩子
	PostHandel(request IRequest)
}
