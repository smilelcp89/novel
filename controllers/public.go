package controllers

import (
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"novel/global"
	"novel/models"
	"novel/utils"
)

//登录页面
func Login(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.HTML(http.StatusOK, "public/error.html", gin.H{"error": err, "redirectUrl": "/login"})
		}
	}()
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "public/login.html", gin.H{"CaptchaId": captcha.New()})
	} else {
		username := c.PostForm("username")
		password := c.PostForm("password")
		captchaId := c.PostForm("captchaId")
		code := c.PostForm("code")
		if captchaId == "" {
			panic("验证码不能为空")
		}
		if username == "" || password == "" {
			panic("账号密码不能为空")
		}
		if !captcha.VerifyString(captchaId, code) {
			panic("验证码错误")
		}
		user := (&models.TbUser{}).GetOne(0, username, "id,password")
		if user.Id <= 0 {
			panic("账号或密码错误")
		}
		if utils.Md5String(password) != user.Password {
			panic("账号或密码错误")
		}
		//登录成功，将信息存进session
		sess := sessions.Default(c)
		sess.Set(global.LoginSession, global.MapData{"userId": user.Id, "userName": user.Username, "test": 1})
		err := sess.Save()
		if err != nil {
			log.Println("保存session报错：" + err.Error())
		}
		c.Redirect(http.StatusFound, "/admin/index")
	}
}

//登出
func Logout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Clear()
	sess.Save()
	c.Redirect(http.StatusFound, "/login")
}

//错误页面
func Error(c *gin.Context) {
	c.HTML(http.StatusOK, "public/error.html", gin.H{})
}
