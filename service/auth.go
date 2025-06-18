package service

import (
	"errors"
	"go-gin-gorm-api/database"
	"go-gin-gorm-api/model"
	"go-gin-gorm-api/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Login(usernameOrEmail, password string) (string, error) {
	var user model.User

	err := database.DB.Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).First(&user).Error
	if err != nil {
		return "", errors.New("用户不存在")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("密码错误")
	}

	// 生成 token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("生成 token 失败")
	}

	// 更新登录时间
	database.DB.Model(&user).Update("last_login", time.Now())

	return token, nil
}
