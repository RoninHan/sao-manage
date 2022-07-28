package forms

import "time"

type InsertOrUpdateBannerForm struct {
	BanerTitle string `form:"banerTitle" json:"banerTitle" binding:"required"`
	BanerContent string `form:"banerContent" json:"banerContent" binding:"required"`
	BannerUrl string `form:"bannerUrl" json:"bannerUrl" binding:"required"`
	CreateUser string `form:"createUser" json:"createUser" binding:"required"`
	CreateUserName string `form:"createUserName" json:"createUserName" binding:"required"`
	UpdateUser string `form:"updateUser" json:"updateUser" binding:"required"`
	UpdateUserName string `form:"updateUserName" json:"updateUserName" binding:"required"`
	Desc int `form:"desc" json:"desc" binding:"required"`
}

type  BannerSearch struct {
	BanerTitle string
	Page int
	PageSize int
}

type BannerList struct {
	BannerGuid string
	BanerTitle string
	BanerContent string
	BannerUrl string
	CreateUser string
	CreateUserName string
	CreateTime  *time.Time
	UpdateUser string
	UpdateUserName string
	UpdateTime *time.Time
	Desc int
}