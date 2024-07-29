# qrAPI

API that creates an encoded image of a QR code that leads to your url.

## example

bytes = qrapi.fly.dev/?url=youtube.com&size=256

image = "data:image/png;base64, " + bytes;
