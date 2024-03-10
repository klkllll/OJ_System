package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")                                                                                                                          //允许所有来源的请求
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")                                                                                   //允许的HTTP方法
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, AccessToken")                                                //允许的请求头
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type") //暴露给客户端的响应头
		c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                  // 允许发送包含凭据的请求。                                                                                                            // 允许发送包含凭据的请求
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent) //204表示请求成功，但不返回任何内容
		}
		c.Next()
	}
}
