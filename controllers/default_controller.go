package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var charSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// RegisterRoutes ...
func RegisterRoutes(router *gin.Engine) {
	router.GET("/", showIndexPage)
	router.POST("/", encoder)
}

func encoder(c *gin.Context) {
	inputText := c.PostForm("inputTextArea")
	render(c, gin.H{
		"title":  "Base64",
		"value":  inputText,
		"output": base64Encoder(inputText),
		"year":   time.Now().Year(),
	},
		"index.html")
}

func showIndexPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Home Page",
		"year":  time.Now().Year(),
	},
		"index.html")
}

// base64Encoder encodes text to base64 format
func base64Encoder(input string) string {
	bin := ""
	inputLen := len(input)
	for _, c := range input {
		bin = fmt.Sprintf("%s%.8b", bin, c)
	}
	if len(bin)%6 != 0 {
		padding := (6 - len(bin)%6)
		bin += strings.Repeat("0", padding)
	}
	encodedOutput := ""
	n := len(bin)
	for i := 0; i < n; i += 6 {
		v, _ := strconv.ParseInt(bin[i:i+6], 2, 64)
		encodedOutput += string(charSet[v])
	}
	extra := inputLen % 3
	if extra != 0 {
		encodedOutput += strings.Repeat("=", 3-extra)
	}
	return encodedOutput
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	if _, ok := data["placeHolder"]; !ok {
		if _, ok := data["value"]; ok {
			data["placeHolder"] = data["value"]
		} else {
			data["placeHolder"] = "Welcome"
		}
	}
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
