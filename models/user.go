// models/user.go
package models

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"uniqueIndex;size:64;not null"`
    KeyMD5   string `gorm:"size:32;not null"`
}