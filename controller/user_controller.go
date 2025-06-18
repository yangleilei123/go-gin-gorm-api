package controller

import (
	"go-gin-gorm-api/database"
	"go-gin-gorm-api/model"
	"go-gin-gorm-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) {
	uidVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	userID := uidVal.(uint)
	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	// 不返回密码
	user.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": user})
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.Register(input.Username, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"` // 可为用户名或邮箱
	Password string `json:"password" binding:"required"` // 密码
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := service.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
	})
}

type UpdateUserInput struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

func UpdateUser(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}
	userID := userIDVal.(uint)

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数解析失败"})
		return
	}

	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	// ✅ 用户名唯一性校验（排除自己）
	if input.Username != "" && input.Username != user.Username {
		var count int64
		database.DB.Model(&model.User{}).
			Where("username = ? AND id != ?", input.Username, user.ID).
			Count(&count)
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已被占用"})
			return
		}
		user.Username = input.Username
	}

	// ✅ 邮箱唯一性校验（排除自己）
	if input.Email != "" && input.Email != user.Email {
		var count int64
		database.DB.Model(&model.User{}).
			Where("email = ? AND id != ?", input.Email, user.ID).
			Count(&count)
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱已被占用"})
			return
		}
		user.Email = input.Email
	}

	// 更新字段
	if input.Username != "" {
		user.Username = input.Username
	}
	if input.Phone != "" {
		user.Phone = input.Phone
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Avatar != "" {
		user.Avatar = input.Avatar
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败：" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功", "user": user})
}
