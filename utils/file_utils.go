package utils

import (
	"io"
	"os"
	"github.com/gin-gonic/gin"
)

// UploadFile handles file uploads.
func UploadFile(c *gin.Context) error {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	dst, err := os.Create("./uploads/uploaded_file")
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}
