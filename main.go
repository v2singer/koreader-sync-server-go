package main

import (
    "koreader-sync-server-go/controllers"
    "koreader-sync-server-go/middleware"
    "koreader-sync-server-go/db"
    "github.com/gin-gonic/gin"
)

func main() {
    db.InitSQLite()
    r := gin.Default()

    v1 := r.Group("/")
    v1.POST("/users/create", controllers.CreateUser)
    v1.GET("/users/auth", middleware.AuthMiddleware, controllers.AuthUser)
    v1.PUT("/syncs/progress", middleware.AuthMiddleware, controllers.UpdateProgress)
    v1.GET("/syncs/progress/:document", middleware.AuthMiddleware, controllers.GetProgress)
    v1.GET("/healthcheck", controllers.HealthCheck)

    r.Run(":7200")
} 