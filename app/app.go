package app

import "gorm.io/gorm"
import "github.com/gin-gonic/gin"

func AppInit(db *gorm.DB) {
    // migrate models db.AutoMigrate()
    router := gin.Default()
    router.GET("/entity", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "data": "hello data",
        })
    })
    router.Run()
}
