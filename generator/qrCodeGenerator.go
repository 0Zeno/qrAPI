package qrCodeGenerator

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
)

var image []byte
var err error

func GenerateQR(url string, size int) ([]byte){
	image, err = qrcode.Encode(url, qrcode.Medium, size)
	if err != nil {
		fmt.Println("Somthing went wrong generating the QR code")
	}
	return image
}
