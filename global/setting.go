/*
* @Author: HuberyChang
* @Date: 2021/5/9 14:30
 */

package global

import (
	"github.com/HuberyChang/blog-service/pkg/logger"
	"github.com/HuberyChang/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger
)
