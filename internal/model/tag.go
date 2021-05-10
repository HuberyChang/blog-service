/*
* @Author: HuberyChang
* @Date: 2021/5/8 21:55
 */

package model

import "github.com/HuberyChang/blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
