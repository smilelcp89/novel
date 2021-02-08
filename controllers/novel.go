package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"novel/models"
	"novel/utils"
	"time"
)

//小说首页
func NovelIndex(c *gin.Context){
	c.HTML(http.StatusOK, "novel/index.html", gin.H{"title": "小说列表"})
}

//创建小说
func CreateNovel(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.HTML(http.StatusOK, "public/error.html", gin.H{"error": err, "redirectUrl": "/admin/novel/create"})
		}
	}()
	if c.Request.Method == "POST"{
		//创建小说
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "" || password == "" {
			panic("账号密码不能为空")
		}
		password = utils.Md5String(password)
		user := &models.TbUser{Username: username, Password: password, CreateTime: int(time.Now().Unix()), UpdateTime: int(time.Now().Unix())}
		if !user.Create() {
			panic("创建小说失败")
		}
		c.Redirect(http.StatusFound, "/admin/novel/index")
	}else{
		c.HTML(http.StatusOK, "novel/edit.html", gin.H{"title": "创建小说"})
	}
}
