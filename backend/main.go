package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerifyRequest struct {
	Password string `json:"password"`
}

func main() {
	r := gin.Default()

	// 允许前端跨域（本地开发用）
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.Static("/", "./static")

	r.POST("/api/verify", func(c *gin.Context) {
		var req VerifyRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
			})
			return
		}

		if req.Password == "celebrity0302" {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
			})
		}
	})

	r.Run(":8080")
}
