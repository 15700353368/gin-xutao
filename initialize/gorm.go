package initialize

import (
	"gin-xutao/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Gorm() *gorm.DB {

	dsn := global.GVA_CONFIG.Mysql.Dsn()
	global.GVA_LOG.Info("MYSQL链接信息",zap.String("mysql",dsn))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "gormv2_",
			SingularTable: true,
		},
	})

	if err != nil{
		global.GVA_LOG.Error("数据库连接出错",zap.String("mysql",err.Error()))
	}
	return db
}