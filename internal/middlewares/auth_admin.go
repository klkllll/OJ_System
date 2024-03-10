package middlewares

import (
	"net/http"

	"getcharzp.cn/helper"
	"github.com/gin-gonic/gin"
)

// 中间件
// AuthAdminCheck is a middleware function that checks if the user is authenticated with admin role.
func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization") // 前端发送过来的请求头，Authorization带有token的
		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Authorization",
			})
			return
		}
		if userClaim == nil || userClaim.IsAdmin != 1 {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Admin",
			})
			return
		}
		c.Next()
	}
}
