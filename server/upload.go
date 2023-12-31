package server

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

func generateRandomString(length int) string {
	rnd := make([]byte, length)
	rand.Read(rnd)
	return base64.URLEncoding.EncodeToString(rnd)
}

func generateRandomFilePath(ext string) string {
	rnd := generateRandomString(32)
	return "./upload/" + rnd + "." + ext
}

func uploadFile(c *gin.Context) {
	f, _ := c.FormFile("file")
	if f.Size > 1024*1024*50 {
		c.JSON(400, gin.H{
			"error": "file size too big",
		})
		return
	}

	ext := f.Filename[len(f.Filename)-3:]
	filename := generateRandomFilePath(ext)

	c.SaveUploadedFile(f, filename)
	c.JSON(200, gin.H{
		"success": true,
		"file":    filename,
	})
}
