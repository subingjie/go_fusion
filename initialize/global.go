package initialize

import (
	"go_fusion/config"

	"gorm.io/gorm"
)

var (
	GlobalConfig *config.Config // 全局配置
	GormDb       *gorm.DB
)
