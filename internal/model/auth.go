/*
* @Author: HuberyChang
* @Date: 2021/5/14 10:58
 */

package model

type Auth struct {
	*Model
	Appkey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}
