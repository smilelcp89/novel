package routes

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"novel/controllers"
	"novel/middlewares"
	"path"
	"time"
)

func Setup(r *gin.Engine) *gin.Engine {
	r.GET("/login", controllers.Login)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)
	r.GET("/error", controllers.Error)
	r.GET("/captcha/:captchaId", func(c *gin.Context) {
		//c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		//c.Writer.Header().Set("Pragma", "no-cache")
		//c.Writer.Header().Set("Expires", "0")
		c.Writer.Header().Set("Content-Type", "image/png")
		_, file := path.Split(c.Request.URL.Path)
		ext := path.Ext(file)
		captchaId := file[:len(file)-len(ext)]
		var content bytes.Buffer
		//var captchaId = c.Param("captchaId")
		log.Println(captchaId)
		_ = captcha.WriteImage(&content, captchaId, captcha.StdWidth, captcha.StdHeight)
		http.ServeContent(c.Writer, c.Request, captchaId + ext, time.Time{}, bytes.NewReader(content.Bytes()))
	})
	//后台路由
	admin := r.Group("/admin",middlewares.LoginAuth())
	{
		admin.GET("/index", controllers.Index)
		admin.GET("/user/index", controllers.UserIndex)
		admin.GET("/user/create", controllers.CreateUser)
		admin.POST("/user/create", controllers.CreateUser)


		admin.GET("/novel/create", func(c *gin.Context) {
			c.String(200,"novel_create")
		})
	}
	return r
}
