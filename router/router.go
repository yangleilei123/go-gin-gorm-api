package router

import (
	"go-gin-gorm-api/controller" // 引入控制器模块
	"go-gin-gorm-api/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 用户相关路由分组
	user := r.Group("/users")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login) // 新增登录接口

		auth := user.Group("/")
		auth.Use(middleware.JWTAuth())
		auth.GET("/me", controller.GetCurrentUser) // 新增接口

		auth.PUT("/me", controller.UpdateUser)
	}

	return r
}
