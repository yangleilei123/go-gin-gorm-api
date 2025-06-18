package database

import (
	"fmt"
	"go-gin-gorm-api/config"
	"go-gin-gorm-api/model"
	"log"

	"gorm.io/driver/mysql" // MySQL 驱动
	"gorm.io/gorm"         // GORM 主库
)

var DB *gorm.DB // 全局数据库连接对象

// InitMySQL 初始化 MySQL 数据库连接
func InitMySQL() {
	c := config.Conf
	// 构造 DSN 数据源名称
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.GetString("mysql.user"),
		c.GetString("mysql.password"),
		c.GetString("mysql.host"),
		c.GetString("mysql.port"),
		c.GetString("mysql.dbname"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("MySQL connection error: %v", err)
	}

	// 自动建表（不存在时）
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}
}
