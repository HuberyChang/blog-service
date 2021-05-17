/*
* @Author: HuberyChang
* @Date: 2021/5/16 23:01
 */

package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog_service")
		c.Set("spp_version", "1.0.0")
		c.Next()
	}
}
