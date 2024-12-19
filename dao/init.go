package dao

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func Database(connRead, connWrite string){
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	}else{
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize: 256, // string type length
		DisableDatetimePrecision: true, // 禁止 datatime 精度，mysql 5.6 之前的不支持
		DontSupportRenameIndex: true, // 不要重命名索引，要么就先删掉再重建。
		DontSupportRenameColumn: true, // 用 change 重命名列， mysql8 之前的不支持
		SkipInitializeWithVersion: false, // auto configure based on current MySQL version
	}), &gorm.Config{
		Logger : ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to open database")
	}

	sqlDB , _ := db.DB()
	sqlDB.SetMaxOpenConns(20)	// 最大连接打开数
	sqlDB.SetMaxIdleConns(20) // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(time.Second*30)
	_db = db

	// master-slave configuration
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{mysql.Open(connWrite)}, // write operation
		Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, // read operation
		Policy: dbresolver.RandomPolicy{}, // 处理从库的操作策略
	}))

	migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}