/*
* @Author: HuberyChang
* @Date: 2021/5/14 14:46
 */

package dao

import "github.com/HuberyChang/blog-service/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	return auth.Get(d.engine)
}
