package main

import (
	"net/http"
	qrCodeGenerator "qrAPI/generator"

	"github.com/gin-gonic/gin"
)
func getResonse(context *gin.Context){
	context.IndentedJSON(http.StatusOK, qrCodeGenerator.GenerateQR("youtube.com", 256))
	qrCodeGenerator.GenerateQR("youtube.com", 256)

}

func main() {
	router := gin.Default()
	router.GET("/generateQR", getResonse)
	router.Run("localhost:8080")
}

