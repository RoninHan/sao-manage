package controller

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"sao-manage/Response"
	"sao-manage/dao"
	"sao-manage/forms"
	"sao-manage/models"
	"sao-manage/utils"
	"time"
)

// PasswordLogin 登录
func PasswordLogin(c *gin.Context) {
	PasswordLoginForm := forms.PasswordLoginForm{

	}
	if err := c.ShouldBind(&PasswordLoginForm); err != nil {
		utils.HandleValidatorError(c,err)
		return
	}
	//if !store.Verify(PasswordLoginForm.CaptchaId,PasswordLoginForm.Captcha,true){
	//	Response.Err(c,400,400,"验证码错误","")
	//}
	md5Password :=utils.NewMd5ToString(PasswordLoginForm.PassWord)
	user,ok:=dao.FindUserInfo(PasswordLoginForm.UserName,md5Password)
	color.Blue(user.Password)
	if !ok {
		Response.Err(c,401,"This User is NOT FOUND","")
		return
	}
	token :=utils.CreateToken(c, user.UserGuid,user.UserName,user.Role)
	userinfoMap :=HandleUserModelToMap(user)
	userinfoMap["token"] = token
	Response.Success(c,  "success", userinfoMap)
}

// GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	// 获取参数
	UserListForm := forms.UserListForm{}
	if err := c.ShouldBind(&UserListForm); err != nil {
		color.Cyan(string(UserListForm.PageSize))
		utils.HandleValidatorError(c, err)
		return
	}
	// 获取数据
	total, userlist := dao.GetUserListDao(UserListForm.Current, UserListForm.PageSize)
	// 判断
	if (total + len(userlist)) == 0 {
		Response.Err(c, 401, "未获取到到数据", map[string]interface{}{
			"total":    total,
			"userlist": userlist,
		})
		return
	}
	Response.SuccessGrid(c,  "获取用户列表成功", userlist,total)
}

func HandleUserModelToMap(user *models.User) map[string]interface{} {
	birthday := ""
	if user.Birthday == nil {
		birthday = ""
	} else {
		birthday = user.Birthday.Format("2006-01-02")
	}
	userItemMap := map[string]interface{}{
		"id":        user.UserGuid,
		"userName": user.UserName,
		"head_url":  user.HeadUrl,
		"birthday":  birthday,
		"address":   user.Address,
		"gender":    user.Gender,
		"role":      user.Role,
		"mobile":    user.Mobile,
	}
	return userItemMap
}

// PutHeaderImage 上传用户头像
func PutHeaderImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileObj, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 把文件上传到minio对应的桶中
	ok := utils.UploadFile("userheader", file.Filename, fileObj, file.Size)
	if !ok {
		Response.Err(c, 401, "头像上传失败", "")
		return
	}
	headerUrl := utils.GetFileUrl("userheader", file.Filename, time.Second*24*60*60)
	if headerUrl == "" {
		Response.Err(c,  401, "获取用户头像失败", "headerMap")
		return
	}
	//TODO 把用户的头像地址存入到对应user表中head_url 中
	Response.Success(c,  "头像上传成功", map[string]interface{}{
		"userheaderUrl": headerUrl,
	})
}

func InsertUser(c *gin.Context){
	UserForm := models.User{}
	if err := c.ShouldBind(&UserForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	UserForm.UserGuid=utils.GetGuid()
	UserForm.Password=utils.NewMd5ToString(UserForm.Password)
	ok:=dao.InsertUser(UserForm)
	if !ok {
		Response.Err(c, 401,  "添加失败", ok)
		return
	}
	Response.Success(c, "添加成功", ok)
}

func UpdatetUser(c *gin.Context){
	UserForm := forms.UserForm{}
	if err := c.ShouldBind(&UserForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	ok:=dao.UpdateUserById(UserForm)

	if !ok {
		Response.Err(c, 401,  "添加失败", ok)
		return
	}
	Response.Success(c,  "添加成功", ok)
}