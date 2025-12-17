package main

// curl -X POST http://localhost:8080/library/add -H "Content-Type: application/json" -H "X-User-ID: 123" -d "{\"manga_id\": 10}"

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MangaRequest struct {
	MangaID int `json:"manga_id"`
}

type LibraryService struct{}

func (s *LibraryService) AddToLibrary(userID int, mangaID int) error {
	if mangaID == 0 {
		return errors.New("invalid manga ID")
	}
	return nil
}

func main() {
	// 1) JSON LOGGING
	logrus.SetFormatter(&logrus.JSONFormatter{})

	r := gin.Default()
	service := &LibraryService{}

	r.POST("/library/add", func(c *gin.Context) {
		var req MangaRequest

		// bind JSON
		if err := c.ShouldBindJSON(&req); err != nil {
			// 2) LOG WITH FIELDS + 3) ERROR LEVEL
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("Failed to parse JSON")

			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		// extract user id (mock)
		userID, _ := strconv.Atoi(c.GetHeader("X-User-ID"))

		err := service.AddToLibrary(userID, req.MangaID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"user_id":  userID,
				"manga_id": req.MangaID,
				"error":    err.Error(),
			}).Warn("Failed to add manga to library")

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// SUCCESS LOG
		logrus.WithFields(logrus.Fields{
			"user_id":  userID,
			"manga_id": req.MangaID,
		}).Info("Manga added to library")

		c.JSON(http.StatusOK, gin.H{"message": "added successfully"})
	})

	r.Run(":1213")
}
