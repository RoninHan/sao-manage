package dao

import (
	"fmt"
	"sao-manage/forms"
	"sao-manage/global"
	"sao-manage/models"
)

var users []models.User

// GetUserList 获取用户列表(page第几页,page_size每页几条数据)
func GetUserListDao(page int, page_size int) (int, []interface{}) {
	// 分页用户列表数据
	userList := make([]interface{}, 0, len(users))
	// 计算偏移量
	offset := (page - 1) * page_size
	// 查询所有的user
	result := global.DB.Offset(offset).Limit(page_size).Find(&users)
	// 查不到数据时
	if result.RowsAffected == 0{
		return 0, userList
	}
	// 获取user总数
	total := len(users)
	// 查询数据
	result.Offset(offset).Limit(page_size).Find(&users)
	//
	for _, useSingle := range users {
		birthday := ""
		if useSingle.Birthday == nil {
			birthday = ""
		}else {
			// 给未设置生日的初始值
			birthday = useSingle.Birthday.Format("2006-01-02")
		}
		userItemMap := map[string]interface{}{
			"userGuid":        useSingle.UserGuid,
			"password":  useSingle.Password,
			"userName": useSingle.UserName,
			"head_url":  useSingle.HeadUrl,
			"birthday":  birthday,
			"address":   useSingle.Address,
			"gender":    useSingle.Gender,
			"role":      useSingle.Role,
			"mobile":    useSingle.Mobile,
		}
		userList = append(userList, userItemMap)
	}
	return total, userList
}

var user models.User
// UsernameFindUserInfo 通过username找到用户信息
func FindUserInfo(username string, password string) (*models.User, bool) {
	// 查询用户
	rows := global.DB.Where(&models.User{UserName: username,Password: password}).Find(&user)
	fmt.Println(&user)
	if rows.RowsAffected < 1 {
		return &user, false
	}
	return &user, true
}


//添加用户
func InsertUser(userInfo models.User) bool{
	row :=global.DB.Create(userInfo)
	if row.RowsAffected < 1 {
		return false
	}
	return true
}

//更新用户基本信息
func UpdateUserById(userInfo forms.UserForm) bool{
	userItemMap := map[string]interface{}{
		"userName": userInfo.UserName,
		"head_url":  userInfo.HeadUrl,
		"birthday":  userInfo.Birthday,
		"address":   userInfo.Address,
		"gender":    userInfo.Gender,
		"mobile":    userInfo.Mobile,
	}
	row :=global.DB.Where(&models.User{UserGuid: userInfo.UserGuid }).Updates(userItemMap)
	if row.RowsAffected < 1 {
		return false
	}
	return true
}