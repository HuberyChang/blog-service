/*
* @Author: HuberyChang
* @Date: 2021/5/8 22:11
 */

package v1

import (
	"github.com/HuberyChang/blog-service/pkg/app"
	"github.com/HuberyChang/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (a Article) List(c *gin.Context) {

}
func (a Article) Create(c *gin.Context) {

}
func (a Article) Update(c *gin.Context) {

}
func (a Article) Delete(c *gin.Context) {

}
