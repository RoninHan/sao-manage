package models

import "time"
type User struct {
	UserGuid string      `json:"userGuid" gorm:"primaryKey"`
	Password string    `json:"password"`
	UserName string    `json:"userName"`
	HeadUrl  string    `json:"head_url"`
	Birthday *time.Time `json:"birthday" gorm:"type:date"`
	Address  string    `json:"address"`
	Gender   string    `json:"gender"`
	Role     int       `json:"role"`
	Mobile   string    `json:"mobile"`
}

// 设置User的表名为`profiles`
func (User) TableName() string {
	return "user"
}
