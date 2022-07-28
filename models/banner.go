
package models

import "time"

type Banner struct {
	BannerGuid string `json:"bannerGuid" gorm:"primaryKey"`
	BanerTitle string `json:"banerTitle"`
	BanerContent string `json:"banerContent"`
	BannerUrl string `json:"bannerUrl"`
	CreateUser string `json:"createUser"`
	CreateUserName string `json:"createUserName"`
	CreateTime *time.Time `json:"createTime" gorm:"type:date"`
	UpdateUser string `json:"updateUser"`
	UpdateUserName string `json:"updateUserName"`
	UpdateTime *time.Time `json:"updateTime" gorm:"type:date"`
	Desc int `json:"desc"`
}

func (Banner) TableName() string {
	return "banner"
}

