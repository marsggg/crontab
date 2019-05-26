package main

import (
	"flag"
	"fmt"
	"github.com/marsggg/crontab/master"
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
	//master -config ./master.json
	flag.StringVar(&confFile, "config", "./master.json", "指定master路径")
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
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}

	//初始化日志管理器
	if err = master.InitLogMgr(); err != nil {
		goto ERR
	}

	//任务管理器初始化
	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}

	//启动Api Http服务

	//启动Api HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	for {
		time.Sleep(1 * time.Second)
	}

	return

ERR:
	fmt.Println(err)
}
