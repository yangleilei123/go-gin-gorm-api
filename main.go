package main

import (
	"go-gin-gorm-api/config"   // 配置加载模块
	"go-gin-gorm-api/database" // 数据库初始化模块
	"go-gin-gorm-api/router"   // 路由模块
)

func main() {
	config.InitConfig()       // 加载 config.yaml 配置
	database.InitMySQL()      // 初始化数据库连接
	r := router.SetupRouter() // 初始化并设置路由
	r.Run(":8080")            // 启动 HTTP 服务
}
