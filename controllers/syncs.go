package controllers

import (
	"koreader-sync-server-go/db"
	"koreader-sync-server-go/models"
	"koreader-sync-server-go/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2003, "message": "Invalid request"})
		return
	}
	// 检查用户名是否已存在
	var count int64
	db.DB.Model(&models.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		c.JSON(http.StatusPaymentRequired, gin.H{"code": 2002, "message": "Username is already registered."})
		return
	}
	user := models.User{
		Username: req.Username,
		KeyMD5:   utils.MD5Hex(req.Password),
	}
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"code": 2000, "message": "Unknown server error."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"username": req.Username})
}

func AuthUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"authorized": "OK"})
}

func UpdateProgress(c *gin.Context) {
	username := c.GetString("username")

	var req struct {
		Document   string  `json:"document"`
		Percentage float64 `json:"percentage"`
		Progress   string  `json:"progress"`
		Device     string  `json:"device"`
		DeviceID   string  `json:"device_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Document == "" || req.Progress == "" || req.Device == "" {
		c.JSON(http.StatusForbidden, gin.H{"code": 2003, "message": "Invalid request"})
		return
	}
	// 查找用户
	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 2001, "message": "Unauthorized"})
		return
	}
	// 查找是否已有进度
	var progress models.Progress
	timestamp := time.Now().Unix()
	err := db.DB.Where("user_id = ? AND document = ?", user.ID, req.Document).First(&progress).Error
	if err == nil {
		// 已有则更新
		progress.Percentage = req.Percentage
		progress.Progress = req.Progress
		progress.Device = req.Device
		progress.DeviceID = req.DeviceID
		progress.Timestamp = timestamp
		if err := db.DB.Save(&progress).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"code": 2000, "message": "Unknown server error."})
			return
		}
	} else {
		// 没有则新建
		progress = models.Progress{
			UserID:     user.ID,
			Document:   req.Document,
			Percentage: req.Percentage,
			Progress:   req.Progress,
			Device:     req.Device,
			DeviceID:   req.DeviceID,
			Timestamp:  timestamp,
		}
		if err := db.DB.Create(&progress).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"code": 2000, "message": "Unknown server error."})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"document": req.Document, "timestamp": timestamp})
}

func GetProgress(c *gin.Context) {
	username := c.GetString("username")
	document := c.Param("document")
	if document == "" {
		c.JSON(http.StatusForbidden, gin.H{"code": 2004, "message": "Field 'document' not provided."})
		return
	}
	// 查找用户
	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 2001, "message": "Unauthorized"})
		return
	}
	// 查找进度
	var progress models.Progress
	if err := db.DB.Where("user_id = ? AND document = ?", user.ID, document).First(&progress).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	res := gin.H{
		"percentage": progress.Percentage,
		"progress":   progress.Progress,
		"device":     progress.Device,
		"device_id":  progress.DeviceID,
		"timestamp":  progress.Timestamp,
		"document":   progress.Document,
	}
	c.JSON(http.StatusOK, res)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"state": "OK"})
}
