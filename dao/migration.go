package dao

import (
	"fmt"
	"gin_mall_tmp/model"
	"time"
)

func migration() {
    var version int
    err := _db.Model(&model.Migrations{}).Select("version").Order("id desc").Limit(1).Scan(&version).Error
    if err != nil && err.Error() != "record not found" {
        fmt.Println("Error checking migration version: ", err)
        return
    }

    if version == 0 { // 如果没有迁移记录或版本为0
        fmt.Println("Running migration for the first time...")
        err = _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
            &model.Address{},
            &model.Admin{},
            &model.Carousel{},
            &model.Cart{},
            &model.Category{},
            &model.Favorite{},
            &model.Notice{},
            &model.Order{},
            &model.Product{},
            &model.ProductImg{},
            &model.User{},
        )
        if err != nil {
            fmt.Println("fail to migrate: ", err)
            return
        }

        // 迁移成功后，记录迁移版本
        err = _db.Create(&model.Migrations{Version: 1, AppliedAt: time.Now().Unix()}).Error
        if err != nil {
            fmt.Println("Failed to update migration table: ", err)
        } else {
            fmt.Println("Database migration successful.")
        }
    } else {
        fmt.Println("Migration already applied. Skipping migration.")
    }
}