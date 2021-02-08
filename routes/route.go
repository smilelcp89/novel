package routes

import (
	"github.com/gin-gonic/gin"
	"novel/controllers"
	"novel/middlewares"
)

func Setup(r *gin.Engine) *gin.Engine {
	r.GET("/login", controllers.Login)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)
	r.GET("/error", controllers.Error)
	r.GET("/captcha/:captchaId", controllers.LoginCaptcha)
	//后台路由，需要登录，使用中间件校验身份
	admin := r.Group("/admin",middlewares.LoginAuth())
	{
		admin.GET("/index", controllers.Index)
		//用户模块
		admin.GET("/user/index", controllers.UserIndex)
		admin.GET("/user/create", controllers.CreateUser)
		admin.POST("/user/create", controllers.CreateUser)
		//小说模块
		admin.GET("/novel/index", controllers.NovelIndex)
		admin.GET("/novel/create", controllers.CreateNovel)
	}
	return r
}
