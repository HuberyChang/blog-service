/*
* @Author: HuberyChang
* @Date: 2021/5/15 15:56
 */

package middleware

import (
	"bytes"
	"time"

	"github.com/HuberyChang/blog-service/global"
	"github.com/HuberyChang/blog-service/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		// next函数将中间件程序分为前后两个部分，next前面的按照顺序执行，next之后会在业务逻辑处理完毕后再执行
		// 先执行next函数前面的代码，执行到next函数时，转而去执行业务处理程序，执行完业务程序，然后再返回中间件执行next后面的代码
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		s := "access log: method: %s, status_code: %d, begin_time: %d, end_time: %d"
		global.Logger.WithFields(fields).Infof(c, s,
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
