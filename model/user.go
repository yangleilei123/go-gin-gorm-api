package model

import "time"

// User 用户模型，适用于实际项目（含认证、审计字段）
type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"size:100;uniqueIndex" json:"username"` // varchar(100)，唯一索引
	Email     string     `gorm:"size:100;uniqueIndex" json:"email"`    // varchar(100)，唯一索引
	Password  string     `json:"-"`                                    // 密码，忽略 json 输出
	Phone     string     `gorm:"size:20" json:"phone"`
	Role      string     `gorm:"size:50" json:"role"`
	Avatar    string     `json:"avatar"`
	Status    int        `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	LastLogin *time.Time `json:"last_login"`
}
