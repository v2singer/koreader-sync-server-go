package db

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "koreader-sync-server-go/models"
)

var DB *gorm.DB

func InitSQLite() {
    var err error
    DB, err = gorm.Open(sqlite.Open("koreader.db"), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    // 自动迁移表结构
    DB.AutoMigrate(&models.User{}, &models.Progress{})
}