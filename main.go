package main

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"novel/global"
	"novel/global/config"
	"novel/models"
	"novel/routes"
)

var Logger *logrus.Logger

func init() {
	global.Setup() //初始化全局变量
	config.Setup()
	models.Setup()
	gob.Register(global.MapData{})
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	r.LoadHTMLGlob("views/**/*")
	// 创建基于cookie的存储引擎，secretkey 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secretkey"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("GOSESSION", store))
	r = routes.Setup(r)
	r.Run()
}
