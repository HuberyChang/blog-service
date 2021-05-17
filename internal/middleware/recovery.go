/*
* @Author: HuberyChang
* @Date: 2021/5/16 22:23
 */

package middleware

import (
	"fmt"
	"time"

	"github.com/HuberyChang/blog-service/global"
	"github.com/HuberyChang/blog-service/pkg/app"
	"github.com/HuberyChang/blog-service/pkg/email"
	"github.com/HuberyChang/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSLL,
		UserName: global.EmailSetting.UserName,
		PassWord: global.EmailSetting.PassWord,
		From:     global.EmailSetting.From,
	})

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)

				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("抛出异常，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)

				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail: %v", err)

				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
