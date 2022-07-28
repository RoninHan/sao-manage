package dao

import (
	"sao-manage/global"
	"sao-manage/models"
)

var banners []models.Banner

func GetBannerList(page int, page_size int)(int, []interface{}){
	// 分页用户列表数据
	bannerList := make([]interface{}, 0, len(banners))
	// 计算偏移量
	offset := (page - 1) * page_size
	result :=global.DB.Offset(offset).Limit(page_size).Find(&banners)
	// 查不到数据时
	if result.RowsAffected == 0{
		return 0, bannerList
	}
	// 获取user总数
	total := len(banners)
	// 查询数据
	result.Offset(offset).Limit(page_size).Find(&banners)
	//
	for _, useSingle := range banners {
		userItemMap := map[string]interface{}{
			"bannerGuid":useSingle.BannerGuid,
			"banerTitle" :useSingle.BanerTitle,
			"banerContent" :useSingle.BanerContent,
			"bannerUrl" :useSingle.BannerUrl,
			"createUser" :useSingle.CreateUser,
			"createUserName" :useSingle.CreateUserName,
			"createTime"  :useSingle.CreateTime,
			"updateUser" :useSingle.UpdateUser,
			"updateUserName" :useSingle.UpdateUserName,
			"updateTime" :useSingle.UpdateTime,
			"desc" :useSingle.Desc,
		}
		bannerList = append(bannerList, userItemMap)
	}
	return total, bannerList
}