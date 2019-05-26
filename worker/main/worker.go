package main

import (
	"flag"
	"fmt"
	"github.com/marsggg/crontab/worker"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var (
	confFile string
)

//解析命令行参数
func initArgs() {
	//worker -config ./master.json
	//worker -h
	flag.StringVar(&confFile, "config", "./worker.json", "指定master路径")
	flag.Parse()
}

func main() {
	var (
		err error
	)
	//初始化线程

	//初始化命令行参数
	initArgs()

	//初始化配置文件
	if err = worker.InitConfig(confFile); err != nil {
		goto ERR
	}

	//启动日志保存器
	if err = worker.InitLogSink(); err != nil {
		goto ERR
	}

	//启动执行器
	if err = worker.InitExecutor(); err != nil {
		goto ERR
	}

	//启动任务调度器
	if err = worker.InitScheduler(); err != nil {
		goto ERR
	}

	//初始化任务管理器
	if err = worker.InitJobMgr(); err != nil {
		goto ERR
	}

	for {
		time.Sleep(1 * time.Second)
	}

	return

ERR:
	fmt.Println(err)
}
