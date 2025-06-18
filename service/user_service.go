package service

import (
	"errors"
	"go-gin-gorm-api/database"
	"go-gin-gorm-api/model"
	"go-gin-gorm-api/utils"
)

// GetAllUsers 查询所有用户
func GetAllUsers() []model.User {
	var users []model.User
	database.DB.Find(&users)
	return users
}

// CreateUser 创建一个新用户
func CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}

func Register(username, email, password string) error {
	db := database.DB

	// 检查用户名或邮箱是否已存在
	var count int64
	db.Model(&model.User{}).Where("username = ? OR email = ?", username, email).Count(&count)
	if count > 0 {
		return errors.New("用户名或邮箱已存在")
	}

	// 密码加密
	hashedPwd, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	// 创建用户
	user := model.User{
		Username: username,
		Email:    email,
		Password: hashedPwd,
	}
	return db.Create(&user).Error
}
