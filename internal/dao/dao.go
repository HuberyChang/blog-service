/*
* @Author: HuberyChang
* @Date: 2021/5/11 9:45
 */

package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
