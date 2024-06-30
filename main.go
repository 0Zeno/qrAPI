package main

import (
	"net/http"
	qrCodeGenerator "qrAPI/generator"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Url string `form:"url"`
	Size string `form:"size"`
}

func convertToInt(size string) int {
	i, err := strconv.Atoi(size)
	if err != nil {
		panic(err)
	}

	if i < 256 {
		i = 256
	}

	if i > 1024 {
		i = 1024
	}
	
	return i;
}

func getResonse(c *gin.Context){
	var input Input
	if c.ShouldBind(&input) == nil {
		c.IndentedJSON(http.StatusOK, qrCodeGenerator.GenerateQR(input.Url, convertToInt(input.Size)))

	}
}

func main() {
	router := gin.Default()
	router.GET("/", getResonse)
	router.Run("localhost:8080")
}

