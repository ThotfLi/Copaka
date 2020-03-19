package utils

import (
	"encoding/json"
	"io/ioutil"
)

type globalog struct {
	Name       string
	Host       string
	Port          int
	Version    string

	MaxPacketSiz  uint32  //数据包最大值
	Maxcon        int  //当前最大连接数
	WorkerPoolSize uint32   //工作池数量
	MaxWorkerTaskLen int //工作池任务存储最大数量
}

var Globalogobject *globalog

func init(){
	Globalogobject = &globalog{Name:"default server",
		                       Host:"localhost",
		                       Port:9191,
	                           Version:"copaka-v0.1",
	                           MaxPacketSiz:512,
	                           Maxcon:3,
	                           WorkerPoolSize:10,
	                           MaxWorkerTaskLen:1024}
}

func InitConfig(){
	buf,err := ioutil.ReadFile("conf/copaka.json")
	if err != nil {
		panic(err)
	}

	//初始化globalog
	err = json.Unmarshal(buf,Globalogobject)
	if err != nil{
		panic(err)
	}

}