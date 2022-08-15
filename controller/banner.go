package controller

import (
	"github.com/gin-gonic/gin"
	"sao-manage/Response"
	"sao-manage/dao"
	"sao-manage/forms"
	"sao-manage/utils"
)

func GetBannerList(c *gin.Context){
	BannerList :=forms.BannerSearch{}
	if err := c.ShouldBind(&BannerList); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 获取数据
	total, bannerlist := dao.GetBannerList(BannerList.Page, BannerList.PageSize)
	// 判断
	if (total + len(bannerlist)) == 0 {
		Response.Err(c, 401, "未获取到到数据", map[string]interface{}{
			"total":    total,
			"userlist": bannerlist,
		})
		return
	}
	Response.Success(c,  "获取用户列表成功", map[string]interface{}{
		"total":    total,
		"userlist": bannerlist,
	})
}