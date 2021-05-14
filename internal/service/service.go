/*
* @Author: HuberyChang
* @Date: 2021/5/11 10:27
 */

package service

import (
	"context"

	"github.com/HuberyChang/blog-service/global"
	"github.com/HuberyChang/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
