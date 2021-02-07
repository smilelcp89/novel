package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"novel/models"
	"novel/utils"
	"time"
)

//用户首页
func UserIndex(c *gin.Context){
	c.HTML(http.StatusOK, "user/index.html", gin.H{"title": "用户列表"})
}

//创建用户
func CreateUser(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.HTML(http.StatusOK, "public/error.html", gin.H{"error": "账号或密码不能为空", "redirectUrl": "/admin/user/create"})
		}
	}()
	if c.Request.Method == "POST"{
		//创建用户
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "" || password == "" {
			panic("账号密码不能为空")
		}
		password = utils.Md5String(password)
		user := &models.TbUser{Username: username, Password: password, CreateTime: int(time.Now().Unix()), UpdateTime: int(time.Now().Unix())}
		if !user.Create() {
			panic("创建用户失败")
		}
		c.Redirect(http.StatusFound, "/admin/user/index")
	}else{
		c.HTML(http.StatusOK, "user/edit.html", gin.H{"title": "创建用户"})
	}
}
