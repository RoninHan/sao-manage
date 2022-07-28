package models

import "time"

type Works struct {
	WorksGuid string `json:"worksGuid" gorm:"primaryKey"`
	WorksName string `json:"worksName"`
	WorksContent string `json:"worksContent"`
	WorksUrl string `json:"worksUrl"`
	CreateUser string `json:"createUser"`
	CreateUserName string `json:"createUserName"`
	CreateTime *time.Time `json:"createTime" gorm:"type:date"`
	UpdateUser string `json:"updateUser"`
	UpdateUserName string `json:"updateUserName"`
	UpdateTime *time.Time `json:"updateTime" gorm:"type:date"`
}

func (Works) TableName() string {
	return "works"
}

