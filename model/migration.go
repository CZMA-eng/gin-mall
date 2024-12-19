package model

import "gorm.io/gorm"

// Migration 用于记录数据库迁移版本
type Migrations struct {
    gorm.Model
    Version int    `json:"version"`  // 迁移的版本号
    AppliedAt int64 `json:"applied_at"` // 迁移执行的时间戳
}