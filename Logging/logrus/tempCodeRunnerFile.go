package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AddToLibraryRequest struct {
	MangaID int `json:"manga_id"`
}

func main() {
	r := gin.Default()

	// Cấu hình logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	r.POST("/library/add", addToLibraryHandler)

	r.Run(":8080")
}

func addToLibraryHandler(c *gin.Context) {
	var req AddToLibraryRequest

	// Parse JSON
	if err := c.BindJSON(&req); err != nil {
		logrus.Warnf("BindJSON failed: %v", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}

	// Validate MangaID
	if req.MangaID <= 0 {
		logrus.Warn("invalid MangaID received")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "MangaID must be > 0",
		})
		return
	}

	// Log thành công
	logrus.Infof("Added manga to library, MangaID=%d", req.MangaID)

	c.JSON(http.StatusOK, gin.H{
		"message": "added successfully",
		"id":      req.MangaID,
	})
}
