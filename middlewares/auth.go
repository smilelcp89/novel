package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"novel/global"
)

// 定义中间
func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)
		if sess.Get(global.LoginSession) == nil{
			c.Abort()
			c.Redirect(http.StatusFound, "/login")
		}
		// 执行函数
		c.Next()
	}
}
