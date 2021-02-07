package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"novel/global"
	"time"
)

//后台首页
func Index(c *gin.Context) {
	sess := sessions.Default(c)
	loginUser := sess.Get(global.LoginSession).(global.MapData)
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"now":      time.Now().Format("2006-01-02 15:04:05"),
		"title":    "后台首页",
		"userName": loginUser["userName"],
	})
}
