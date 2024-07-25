package main

import (
	"net/http"

	qrcode "github.com/skip2/go-qrcode"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Url string `form:"url"`
	Size int `form:"size"`
}

func getResonse(c *gin.Context){
	var input Input
	if c.ShouldBind(&input) == nil {
		qrCode, err := qrcode.Encode(input.Url, qrcode.Medium, input.Size)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//c.Header("Content-Type", "image/png")
		c.JSON(http.StatusOK, qrCode)

	}
}

func main() {
	router := gin.Default()
	router.GET("/", getResonse)
	router.Run()
}

