package utils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

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

	Success(c, map[string]interface{}{"filepath": filename + filepath.Ext(file.Filename)}, "Upload Success")
}

func HandlerImage(c *gin.Context, path string) {
	c.File(path)
}

type mail struct {
	user   string
	passwd string
}

func New(u string, p string) mail {
	temp := mail{user: u, passwd: p}
	return temp
}

func (m mail) Send(title string, text string, toId string) {
	auth := smtp.PlainAuth("", m.user, m.passwd, "smtp.gmail.com")

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.gmail.com",
	}

	conn, err := tls.Dial("tcp", "smtp.gmail.com:465", tlsconfig)
	CheckErr(err)

	client, err := smtp.NewClient(conn, "smtp.gmail.com")
	CheckErr(err)

	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = client.Mail(m.user); err != nil {
		log.Panic(err)
	}

	if err = client.Rcpt(toId); err != nil {
		log.Panic(err)
	}

	w, err := client.Data()
	CheckErr(err)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", m.user, toId, title, text)

	_, err = w.Write([]byte(msg))
	CheckErr(err)

	err = w.Close()
	CheckErr(err)

	client.Quit()
}

func SendEmail(subjectLine string, content string, email string) {
	foo := New("office7326068@gmail.com", "ycdz cctp roaf zpdv")
	foo.Send(subjectLine, content, email)
}
