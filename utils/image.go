package utils

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context, identity string, id string, filename string) {

	//建立路徑
	os.MkdirAll("uploads/"+identity+id, os.ModePerm)

	file, err := c.FormFile("image")
	if err != nil {
		Error(c, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	path := "uploads/" + identity + id + filename + filepath.Ext(file.Filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		Error(c, err.Error())
		return
	}

	Success(c, "", map[string]interface{}{"filepath": id + filename + filepath.Ext(file.Filename)}, "Upload Success")
}

func HandlerImage(c *gin.Context, path string) {
	c.File(path)
}
