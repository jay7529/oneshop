package utils

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func CheckErr(err error) bool {
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	return true
}

func UploadImage(c *gin.Context, dirname string, filename string) {

	//建立路徑
	os.MkdirAll("uploads/"+dirname, os.ModePerm)

	file, err := c.FormFile("image")
	if err != nil {
		Failed(c, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	path := "uploads/" + dirname + filename + filepath.Ext(file.Filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		Failed(c, err.Error())
		return
	}

	Success(c, map[string]interface{}{"filepath": filename}, "Upload Success")
}

func HandlerImage(c *gin.Context, path string) {
	c.File(path)
}
