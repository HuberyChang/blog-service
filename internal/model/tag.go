/*
* @Author: HuberyChang
* @Date: 2021/5/8 21:55
 */

package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
