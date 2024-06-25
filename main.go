package main

import (
	"fmt"
	qrCodeGenerator "qrAPI/generator"
)

func main() {
	qrCodeGenerator.GenerateQR("youtube.com", 256)
	fmt.Println("Qr generated")

}

