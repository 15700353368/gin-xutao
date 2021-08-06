package global

import (
	"gin-xutao/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_CONFIG config.Server
	GVA_LOG *zap.Logger
	GVA_GORM *gorm.DB
	GVA_REDIS  *redis.Client
)