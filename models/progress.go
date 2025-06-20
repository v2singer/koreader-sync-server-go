// models/progress.go
package models

type Progress struct {
    ID         uint   `gorm:"primaryKey"`
    UserID     uint   `gorm:"index;not null"`
    Document   string `gorm:"size:64;not null;index:idx_user_doc,unique"`
    Percentage float64
    Progress   string
    Device     string
    DeviceID   string
    Timestamp  int64
}