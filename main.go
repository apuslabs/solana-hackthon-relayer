package main

import (
	"apus-relayer/relayer/config"
	"apus-relayer/relayer/db"
	"apus-relayer/relayer/server"
)

func main() {
	// 配置初始化
	config.Init()
	// 连接数据库
	db.Init()
	// 创建server
	server.Root()
}
