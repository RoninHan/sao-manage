package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
	"sao-manage/Response"
)
// base64Captcha  缓存对象
var store = base64Captcha.DefaultMemStore
// GetCaptcha 获取验证码
func GetCaptcha(ctx *gin.Context) {
	//
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	// b64s是图片的base64编码
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误,:%s ", err.Error())
		Response.Err(ctx, http.StatusInternalServerError,  "生成验证码错误", "")
		return
	}
	Response.Success(ctx, "生成验证码成功", gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}