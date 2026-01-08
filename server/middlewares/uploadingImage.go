package middlewares

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"github.com/gin-gonic/gin"
	)
func UploadFileMiddleware(field string, maxSize int64) gin.HandlerFunc {
	return func(c *gin.Context) {

		file, err := c.FormFile(field)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "file is required"})
			return
		}

		// limit size
		if file.Size > maxSize {
			c.AbortWithStatusJSON(400, gin.H{"error": "file too large"})
			return
		}

		// create uploads folder if not exists
		if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "cannot create upload folder"})
			return
		}

		// safe filename
		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		dst := filepath.Join("uploads", filename)

		// save file
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "upload failed"})
			return
		}

		// pass data to next handler
		c.Set("file_path", dst)
		c.Set("file_name", filename)
		c.Set("file_size", file.Size)

		c.Next()
	}
}
